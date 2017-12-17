// Copyright 2015-present, Cyrill @ Schumacher.fm and the CoreStore contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Kind defines the kind or behaviour of an error. An error can have multiple
// Kinds wrapped into each other via bit operations. A zero Kind represents an
// empty Kind.
type Kind uint

// Kinds a slice of Kind. Each Kind does not contain any other Kind.
type Kinds []Kind

// Kinder may be implemented by any other custom error type to define desired
// kinds/behaviours. Multiple Kinds can be returned.
type Kinder interface {
	ErrorKind() Kind
}

// Empty returns true if no behaviour/kind has been set.
func (k Kind) Empty() bool {
	return k == 0
}

func (k Kind) isSet(k2 Kind) bool {
	return k&k2 != 0
}

func (k Kind) attach(k2 Kind) Kind {
	return k | k2 // bit set
}

func (k Kind) detach(k2 Kind) Kind {
	return k & ^k2 // bit clear
}

func (ks Kinds) matchAll(k Kind) bool {
	match := 0
	for _, kss := range ks {
		if k > 0 && kss > 0 && k.isSet(kss) {
			match++
		}
	}
	return match > 0 && len(ks) == match
}

func (ks Kinds) matchAny(k Kind) bool {
	for _, kss := range ks {
		if k > 0 && kss > 0 && k.isSet(kss) {
			return true
		}
	}
	return false
}

// String creates a comma separated list of the Kind names.
func (ks Kinds) String() string {
	var buf strings.Builder
	for i, k := range ks {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(k.String())
	}
	return buf.String()
}

func (k Kind) match(err error) bool {
	switch e := err.(type) {
	case kindStacked:
		return e.Kind&k != 0
	case kindFundamental:
		return e.Kind&k != 0
	case Kinder:
		return e.ErrorKind()&k != 0
	}
	return false
}

// Match returns true if `err` matches the Kind.
func (k Kind) Match(err error) bool {
	return CausedBehaviour(err, k)
}

// New wraps err with the specified Kind. Allows to write an additional message
// which gets formatted by fmt.Sprintf.
func (k Kind) New(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return kindStacked{withStack: errWrapf(err, msg, args...), Kind: k}
}

// Newf creates a new error with a message formatted by fmt.Sprintf.
func (k Kind) Newf(format string, args ...interface{}) error {
	return kindFundamental{fundamental: errNewf(format, args...), Kind: k}
}

type (
	kindStacked struct {
		*withStack
		Kind     Kind
		rawStack []byte // only set in case Unmarshalling
	}
	kindFundamental struct {
		*fundamental
		Kind     Kind
		rawStack []byte // only set in case Unmarshalling
	}
)

// MarshalBinary marshals its receiver into a byte slice, which it returns. It
// returns nil if the error is nil. The returned error is always nil.
func (e kindStacked) MarshalBinary() ([]byte, error) {
	return e.MarshalAppend(nil), nil
}

// MarshalAppend marshals kindStacked into a byte slice. The result is appended
// to b, which may be nil. It returns the argument slice unchanged if the error
// is nil.
func (e kindStacked) MarshalAppend(b []byte) []byte {
	// Encode: Kind, stack
	var tmp [16]byte // For use by PutVarint.
	N := binary.PutUvarint(tmp[:], uint64(e.Kind))
	b = append(b, tmp[:N]...)

	if e.withStack != nil && e.withStack.error != nil {
		b = appendString(b, e.withStack.error.Error())
	} else {
		b = appendString(b, "")
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%+v", e.stack)
	b = appendBytes(b, buf.Bytes())
	return b
}

// UnmarshalBinary unmarshals the byte slice into the receiver, which must be
// non-nil. The returned error is always nil.
func (e *kindStacked) UnmarshalBinary(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	// Decode: Kind, msg, stack
	k, N := binary.Uvarint(b)
	if N < 0 {
		return Fatal.Newf("[errors] Error Kind overflows uint64: %d with data: %q", N, b)
	}
	e.Kind = Kind(k)
	b = b[N:]

	orgErr, b, err := getBytes(b)
	if err != nil {
		return WithStack(err)
	}
	e.withStack = errWrapf(errors.New(string(orgErr)), "")

	e.rawStack, b, err = getBytes(b)
	if err != nil {
		return WithStack(err)
	}
	if len(b) != 0 {
		return Fatal.Newf("[errors] kindStacked.UnmarshalBinary error: trailing bytes: %q", b)
	}
	return nil
}

// MarshalAppend marshals kindStacked into a byte slice. The result is appended
// to b, which may be nil. It returns the argument slice unchanged if the error
// is nil.
func (e kindFundamental) MarshalAppend(b []byte) []byte {
	// Encode: Kind, msg, stack
	var tmp [16]byte // For use by PutVarint.
	N := binary.PutUvarint(tmp[:], uint64(e.Kind))
	b = append(b, tmp[:N]...)
	b = appendString(b, e.msg)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%+v", e.fundamental)
	b = appendBytes(b, buf.Bytes())
	return b
}

// MarshalBinary marshals its receiver into a byte slice, which it returns. It
// returns nil if the error is nil. The returned error is always nil.
func (e kindFundamental) MarshalBinary() ([]byte, error) {
	return e.MarshalAppend(nil), nil
}

// UnmarshalBinary unmarshals the byte slice into the receiver, which must be
// non-nil. The returned error is always nil.
func (e *kindFundamental) UnmarshalBinary(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	// Decode: Kind, msg, stack
	k, N := binary.Uvarint(b)
	if N < 0 {
		return Fatal.Newf("[errors] Error Kind overflows uint64: %d with data: %q", N, b)
	}
	e.Kind = Kind(k)
	b = b[N:]

	msg, b, err := getBytes(b)
	if err != nil {
		return WithStack(err)
	}
	if msg != nil {
		e.fundamental = errNewf(string(msg))
	}
	e.rawStack, b, err = getBytes(b)
	if err != nil {
		return WithStack(err)
	}
	if len(b) != 0 {
		return Fatal.Newf("[errors] kindFundamental.UnmarshalBinary error: trailing bytes: %q", b)
	}
	return nil
}

// Is returns true if `err` is of Kind `k`. It unwraps all underlying errors
// which implement the Causer interface.
func Is(err error, k Kind) bool {
	return k.Match(err)
}

// MatchAny checks if at least one Kind is included in `err`. It does not unwrap
// `err` by its `Causer` interface.
func MatchAny(err error, k ...Kind) bool {
	if err == nil {
		return false
	}
	uk := UnwrapKind(err)
	return Kinds(k).matchAny(uk)
}

// MatchAll checks if all Kinds are included in `err`. It does not unwrap `err`
// by its `Causer` interface.
func MatchAll(err error, k ...Kind) bool {
	if err == nil {
		return false
	}
	uk := UnwrapKind(err)
	return Kinds(k).matchAll(uk)
}

// UnwrapKinds checks if error has a Kind/behaviour and returns a slice with all
// found Kinds. The returned slice is sorted numerical.
func UnwrapKinds(err error) Kinds {
	if err == nil {
		return nil
	}
	uk := UnwrapKind(err)
	if uk == 0 {
		return nil
	}

	var ks Kinds
	const one Kind = 1 // Go type system ;-)
	for i := one - 1; i < maxKindExp; i++ {
		if k := one << i; uk.isSet(k) {
			ks = append(ks, k)
		}
	}
	return ks
}

// UnwrapKind extract the Kind of an error. If the error has not been created
// via this package or does not implement interface `Kinder`, this function will
// return zero.
func UnwrapKind(err error) (k Kind) {
	switch e := err.(type) {
	case kindStacked:
		k = e.Kind
	case kindFundamental:
		k = e.Kind
	case Kinder:
		k = e.ErrorKind()
	}
	return k
}

// UnwrapStack tries to extract the previous stack trace after unmarshalling a
// byte slice into an error. It can return nil.
func UnwrapStack(err error) []byte {
	switch e := err.(type) {
	case kindStacked:
		return e.rawStack
	case kindFundamental:
		return e.rawStack
	default:
		if err = Cause(err); err != nil {
			return UnwrapStack(err) // stack overflow?
		}
	}
	return nil
}

// Attach adds Kind `k` to an error but only if the error has been created with
// this package. For example you can create a Restricted error and then attach a
// Temporary kind. Now the error has two kinds. If the error hasn't been created
// by this package, then nothing happens.
func Attach(err error, k Kind) error {
	switch e := err.(type) {
	case kindStacked:
		e.Kind = e.Kind.attach(k)
		err = e
	case kindFundamental:
		e.Kind = e.Kind.attach(k)
		err = e
	}
	return err
}

// Detach opposite of Attach.
func Detach(err error, k Kind) error {
	switch e := err.(type) {
	case kindStacked:
		e.Kind = e.Kind.detach(k)
		err = e
	case kindFundamental:
		e.Kind = e.Kind.detach(k)
		err = e
	}
	return err
}

// MarshalAppend marshals an arbitrary error into a byte slice. The result is
// appended to b, which may be nil. It returns the argument slice unchanged if
// the error is nil. If the error is not an *Error, it just records the result
// of err.Error(). Otherwise it encodes the full Error struct.
func MarshalAppend(err error, b []byte) []byte {
	if err == nil {
		return b
	}
	switch e := err.(type) {
	case kindStacked:
		// This is an errors.kindStacked. Mark it as such.
		b = append(b, 'S')
		b = e.MarshalAppend(b)
	case kindFundamental:
		// This is an errors.kindFundamental. Mark it as such.
		b = append(b, 'F')
		b = e.MarshalAppend(b)
	default:
		// Ordinary error.
		b = append(b, 'e')
		b = appendString(b, err.Error())
	}
	return b
}

// Unmarshal unmarshals the byte slice into an error value. If the slice is nil
// or empty, it returns nil. Otherwise the byte slice must have been created by
// `Marshal` or `MarshalAppend`. If the encoded error was of type `errors`
// within this package, the returned error value will have that underlying type.
// Otherwise it will be just a simple value that implements the error interface.
func Unmarshal(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	code := b[0]
	b = b[1:]
	switch code {
	case 'e':
		// Plain error.
		var data []byte
		var err error
		data, b, err = getBytes(b)
		if err != nil {
			return WithStack(err)
		}
		if len(b) != 0 {
			return Fatal.Newf("[errors] Unmarshal error: trailing bytes: %q", b)
		}
		return errors.New(string(data))
	case 'S':
		// kindStacked value.
		var err kindStacked
		if err2 := err.UnmarshalBinary(b); err2 != nil {
			return WithStack(err2)
		}
		return err
	case 'F':
		// kindFundamental value.
		var err kindFundamental
		if err2 := err.UnmarshalBinary(b); err2 != nil {
			return WithStack(err2)
		}
		return err
	default:
		return CorruptData.Newf("[errors] Unmarshal error: corrupt data %q", b)
	}
}

func appendString(b []byte, str string) []byte {
	var tmp [16]byte // For use by PutUvarint.
	N := binary.PutUvarint(tmp[:], uint64(len(str)))
	b = append(b, tmp[:N]...)
	b = append(b, str...)
	return b
}

func appendBytes(b []byte, str []byte) []byte {
	var tmp [16]byte // For use by PutUvarint.
	N := binary.PutUvarint(tmp[:], uint64(len(str)))
	b = append(b, tmp[:N]...)
	b = append(b, str...)
	return b
}

// getBytes unmarshals the byte slice at b (uvarint count followed by bytes)
// and returns the slice followed by the remaining bytes.
// If there is insufficient data, both return values will be nil.
func getBytes(b []byte) (data, remaining []byte, _ error) {
	u, N := binary.Uvarint(b)
	if len(b) < N+int(u) {
		return nil, nil, BadEncoding.Newf("[errors] Unmarshal error[1]. Data length: %d", len(b))
	}
	if N == 0 {
		return nil, b, BadEncoding.Newf("[errors] Unmarshal error[2]. Data length: %d", len(b))
	}
	return b[N : N+int(u)], b[N+int(u):], nil
}

func fmtNoSprintf(format string, _ ...interface{}) string {
	return format
}

func errWrapf(err error, format string, args ...interface{}) *withStack {
	sprintf := fmtNoSprintf
	if len(args) > 0 {
		sprintf = fmt.Sprintf
	}
	return &withStack{
		error: &withMessage{
			cause: err,
			msg:   sprintf(format, args...),
		},
		stack: callers(),
	}
}

func errNewf(format string, args ...interface{}) *fundamental {
	sprintf := fmtNoSprintf
	if len(args) > 0 {
		sprintf = fmt.Sprintf
	}
	return &fundamental{
		msg:   sprintf(format, args...),
		stack: callers(),
	}
}

// CausedBehaviour returns the first underlying caused kind/behaviour of the
// error, if possible. An error value has a cause if it implements the following
// interface:
//
//     type Causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause or is nil, false will be returned. The
// variable `k` gets called on each unwrapped "cause" error.
func CausedBehaviour(err error, k Kind) bool {
	if k.match(err) {
		return true
	}
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			return false
		}
		err = cause.Cause() // don't touch if you're unsure
		if k.match(err) {
			return true
		}
	}
	return false
}

// NoKind defines an empty Kind with no behaviour. This constant must be placed
// outside the constant block to avoid a conflict with iota.
const NoKind Kind = 0

// These constants define different behaviours. They are not sorted and new
// constants must be appended at the end. The zero kind defines empty.
const (
	Aborted Kind = 1 << iota
	AlreadyCaptured
	AlreadyClosed
	AlreadyExists
	AlreadyInUse
	AlreadyRefunded
	Blocked
	ReadFailed
	WriteFailed
	VerificationFailed
	DecryptionFailed
	EncryptionFailed
	ConnectionFailed
	BadEncoding
	ConnectionLost
	Declined
	Denied
	Duplicated
	NotEmpty
	Empty
	Exceeded
	Exists
	NotExists
	Expired
	Fatal
	InProgress
	Insufficient
	Interrupted
	IsDirectory
	IsFile
	NotDirectory
	NotFile
	Locked
	Mismatch
	NotAcceptable
	NotAllowed
	NotFound
	NotImplemented
	NotRecoverable
	NotSupported
	NotValid
	Overflowed
	PermissionDenied
	Unauthorized
	UserNotFound
	QuotaExceeded
	Rejected
	Required
	Restricted
	Revoked
	Temporary
	Terminated
	Timeout
	TooLarge
	Unavailable
	WrongVersion
	CorruptData
	maxKind
)

var maxKindExp = Kind(math.Log2(float64(maxKind))) // should be a constant ...

// _KindMap contains alphabetically sorted error constants.
var _KindMap = map[Kind]string{
	Aborted:            "Aborted",
	AlreadyCaptured:    "AlreadyCaptured",
	AlreadyClosed:      "AlreadyClosed",
	AlreadyExists:      "AlreadyExists",
	AlreadyInUse:       "AlreadyInUse",
	AlreadyRefunded:    "AlreadyRefunded",
	BadEncoding:        "BadEncoding",
	Blocked:            "Blocked",
	ConnectionFailed:   "ConnectionFailed",
	ConnectionLost:     "ConnectionLost",
	CorruptData:        "CorruptData",
	Declined:           "Declined",
	DecryptionFailed:   "DecryptionFailed",
	Denied:             "Denied",
	Duplicated:         "Duplicated",
	Empty:              "Empty",
	EncryptionFailed:   "EncryptionFailed",
	Exceeded:           "Exceeded",
	Exists:             "Exists",
	Expired:            "Expired",
	Fatal:              "Fatal",
	InProgress:         "InProgress",
	Insufficient:       "Insufficient",
	Interrupted:        "Interrupted",
	IsDirectory:        "IsDirectory",
	IsFile:             "IsFile",
	Locked:             "Locked",
	Mismatch:           "Mismatch",
	NotAcceptable:      "NotAcceptable",
	NotAllowed:         "NotAllowed",
	NotDirectory:       "NotDirectory",
	NotEmpty:           "NotEmpty",
	NotExists:          "NotExists",
	NotFile:            "NotFile",
	NotFound:           "NotFound",
	NotImplemented:     "NotImplemented",
	NotRecoverable:     "NotRecoverable",
	NotSupported:       "NotSupported",
	NotValid:           "NotValid",
	Overflowed:         "Overflowed",
	PermissionDenied:   "PermissionDenied",
	QuotaExceeded:      "QuotaExceeded",
	ReadFailed:         "ReadFailed",
	Rejected:           "Rejected",
	Required:           "Required",
	Restricted:         "Restricted",
	Revoked:            "Revoked",
	Temporary:          "Temporary",
	Terminated:         "Terminated",
	Timeout:            "Timeout",
	TooLarge:           "TooLarge",
	Unauthorized:       "Unauthorized",
	Unavailable:        "Unavailable",
	UserNotFound:       "UserNotFound",
	VerificationFailed: "VerificationFailed",
	WriteFailed:        "WriteFailed",
	WrongVersion:       "WrongVersion",
}

func (k Kind) String() string {
	if str, ok := _KindMap[k]; ok {
		return str
	}
	return "Kind(" + strconv.FormatUint(uint64(k), 10) + ")"
}
