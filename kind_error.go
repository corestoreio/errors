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
	"strings"
)

// KindSeparator separates multiple Kinds packed into one Kind.
const (
	KindSeparator     = "|"
	kindSeparatorRune = '|'
)

// Kind defines the kind or behaviour of an error. An error can have multiple
// Kinds wrapped into each other via bit operations. A zero Kind represents an
// empty Kind. The underlying type uint might change, so please use the
// functions provided in this package to manipulate the Kind. 63 different
// constants are currently supported.
type Kind string

// Kinds a slice of Kind. Each Kind does not contain any other Kind.
type Kinds []Kind

// Kinder may be implemented by any other custom error type to define desired
// kinds/behaviours. Multiple Kinds can be returned.
type Kinder interface {
	ErrorKind() Kind
}

// Empty returns true if no behaviour/kind has been set.
func (k Kind) Empty() bool {
	return k == ""
}

// Unwrap returns all Kind where the bit flag is set. If Kind is empty, returns
// nil.
func (k Kind) Unwrap() Kinds {
	if k.Empty() {
		return nil
	}
	c := strings.Count(string(k), KindSeparator)
	if c == 0 {
		return Kinds{k}
	}
	// the following code is an optimized strings.Split
	ks := make(Kinds, c+1)
	e := strings.IndexByte(string(k), kindSeparatorRune)
	for i := 0; i < c; i++ {
		if e < 0 {
			e = len(k)
		}
		ks[i] = k[0:e]
		k = k[e+1:]
		e = strings.IndexByte(string(k), kindSeparatorRune)
	}
	ks[c] = k
	return ks
}

func (k Kind) isSet(k2 Kind) bool {
	return strings.Contains(string(k), string(k2))
}

func (k Kind) attach(k2 Kind) Kind {
	return k + KindSeparator + k2
}

func (k Kind) detach(k2 Kind) Kind {
	switch {
	case k == k2:
		return ""
	case k == "":
		return ""
	case k2 == "":
		return k
	}

	idx := strings.Index(string(k), string(k2))
	if idx < 0 {
		return k
	}
	a := Kind(strings.TrimPrefix(strings.TrimSuffix(string(k[0:idx]), KindSeparator), KindSeparator))
	b := Kind(strings.TrimPrefix(strings.TrimSuffix(string(k[idx+len(k2):]), KindSeparator), KindSeparator))

	switch {
	case len(a) == 0:
		return b
	case len(b) == 0:
		return a
	default:
		return a + KindSeparator + b
	}
}

func (ks Kinds) matchAll(k Kind) bool {
	match := 0
	for _, kss := range ks {
		if k != "" && kss != "" && k.isSet(kss) {
			match++
		}
	}
	return match > 0 && len(ks) == match
}

func (ks Kinds) matchAny(k Kind) bool {
	for _, kss := range ks {
		if k != "" && kss != "" && k.isSet(kss) {
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

// String returns the underlying string.
func (k Kind) String() string {
	if k == Kind("\x00") {
		return ""
	}
	return string(k)
}

// Error implements error interface and returns the string without any
// formatting or stack trace. It calls under the hood the string function.
func (k Kind) Error() string {
	return k.String()
}

func (k Kind) match(err error) bool {
	switch e := err.(type) {
	case kindStacked:
		return e.Kind.isSet(k)
	case kindFundamental:
		return e.Kind.isSet(k)
	case Kinder:
		return e.ErrorKind().isSet(k)
	case Kind:
		return e == k
	case error:
		return e.Error() == k.String()
	}
	return false
}

func (k Kind) matchInterface(err error) bool {
	switch e := err.(type) {
	case iFaceAborted:
		return e.Aborted() && k.isSet(Aborted)
	case iFaceAlreadyCaptured:
		return e.AlreadyCaptured() && k.isSet(AlreadyCaptured)
	case iFaceAlreadyClosed:
		return e.AlreadyClosed() && k.isSet(AlreadyClosed)
	case iFaceAlreadyExists:
		return e.AlreadyExists() && k.isSet(AlreadyExists)
	case iFaceAlreadyInUse:
		return e.AlreadyInUse() && k.isSet(AlreadyInUse)
	case iFaceAlreadyRefunded:
		return e.AlreadyRefunded() && k.isSet(AlreadyRefunded)
	case iFaceBlocked:
		return e.Blocked() && k.isSet(Blocked)
	case iFaceReadFailed:
		return e.ReadFailed() && k.isSet(ReadFailed)
	case iFaceWriteFailed:
		return e.WriteFailed() && k.isSet(WriteFailed)
	case iFaceVerificationFailed:
		return e.VerificationFailed() && k.isSet(VerificationFailed)
	case iFaceDecryptionFailed:
		return e.DecryptionFailed() && k.isSet(DecryptionFailed)
	case iFaceEncryptionFailed:
		return e.EncryptionFailed() && k.isSet(EncryptionFailed)
	case iFaceConnectionFailed:
		return e.ConnectionFailed() && k.isSet(ConnectionFailed)
	case iFaceBadEncoding:
		return e.BadEncoding() && k.isSet(BadEncoding)
	case iFaceConnectionLost:
		return e.ConnectionLost() && k.isSet(ConnectionLost)
	case iFaceDeclined:
		return e.Declined() && k.isSet(Declined)
	case iFaceDenied:
		return e.Denied() && k.isSet(Denied)
	case iFaceDuplicated:
		return e.Duplicated() && k.isSet(Duplicated)
	case iFaceNotEmpty:
		return e.NotEmpty() && k.isSet(NotEmpty)
	case iFaceEmpty:
		return e.Empty() && k.isSet(Empty)
	case iFaceExceeded:
		return e.Exceeded() && k.isSet(Exceeded)
	case iFaceExists:
		return e.Exists() && k.isSet(Exists)
	case iFaceNotExists:
		return e.NotExists() && k.isSet(NotExists)
	case iFaceExpired:
		return e.Expired() && k.isSet(Expired)
	case iFaceFatal:
		return e.Fatal() && k.isSet(Fatal)
	case iFaceInProgress:
		return e.InProgress() && k.isSet(InProgress)
	case iFaceInsufficient:
		return e.Insufficient() && k.isSet(Insufficient)
	case iFaceInterrupted:
		return e.Interrupted() && k.isSet(Interrupted)
	case iFaceIsDirectory:
		return e.IsDirectory() && k.isSet(IsDirectory)
	case iFaceIsFile:
		return e.IsFile() && k.isSet(IsFile)
	case iFaceNotDirectory:
		return e.NotDirectory() && k.isSet(NotDirectory)
	case iFaceNotFile:
		return e.NotFile() && k.isSet(NotFile)
	case iFaceLocked:
		return e.Locked() && k.isSet(Locked)
	case iFaceMismatch:
		return e.Mismatch() && k.isSet(Mismatch)
	case iFaceNotAcceptable:
		return e.NotAcceptable() && k.isSet(NotAcceptable)
	case iFaceNotAllowed:
		return e.NotAllowed() && k.isSet(NotAllowed)
	case iFaceNotFound:
		return e.NotFound() && k.isSet(NotFound)
	case iFaceNotImplemented:
		return e.NotImplemented() && k.isSet(NotImplemented)
	case iFaceNotRecoverable:
		return e.NotRecoverable() && k.isSet(NotRecoverable)
	case iFaceNotSupported:
		return e.NotSupported() && k.isSet(NotSupported)
	case iFaceNotValid:
		return e.NotValid() && k.isSet(NotValid)
	case iFaceOverflowed:
		return e.Overflowed() && k.isSet(Overflowed)
	case iFacePermissionDenied:
		return e.PermissionDenied() && k.isSet(PermissionDenied)
	case iFaceUnauthorized:
		return e.Unauthorized() && k.isSet(Unauthorized)
	case iFaceUserNotFound:
		return e.UserNotFound() && k.isSet(UserNotFound)
	case iFaceQuotaExceeded:
		return e.QuotaExceeded() && k.isSet(QuotaExceeded)
	case iFaceRejected:
		return e.Rejected() && k.isSet(Rejected)
	case iFaceRequired:
		return e.Required() && k.isSet(Required)
	case iFaceRestricted:
		return e.Restricted() && k.isSet(Restricted)
	case iFaceRevoked:
		return e.Revoked() && k.isSet(Revoked)
	case iFaceTemporary:
		return e.Temporary() && k.isSet(Temporary)
	case iFaceTerminated:
		return e.Terminated() && k.isSet(Terminated)
	case iFaceTimeout:
		return e.Timeout() && k.isSet(Timeout)
	case iFaceTooLarge:
		return e.TooLarge() && k.isSet(TooLarge)
	case iFaceUnavailable:
		return e.Unavailable() && k.isSet(Unavailable)
	case iFaceWrongVersion:
		return e.WrongVersion() && k.isSet(WrongVersion)
	case iFaceCorruptData:
		return e.CorruptData() && k.isSet(CorruptData)
	case iFaceOutOfRange:
		return e.OutOfRange() && k.isSet(OutOfRange)
	case iFaceOutOfDate:
		return e.OutOfDate() && k.isSet(OutOfDate)
	case iFaceTooShort:
		return e.TooShort() && k.isSet(TooShort)
	}
	return false
}

// Match returns true if `err` matches the Kind.
// Very fast matching.
func (k Kind) Match(err error) bool {
	return CausedBehaviour(err, k)
}

// MatchInterface supports interface behaviour type matching to test for a kind.
// This allows a package to define a behaviour/kind of an error without
// importing this package. An error type should implement a function like:
// 		interface{ Fatal() bool }
// Where `Fatal` can be any behaviour name like the constants in this package.
// MatchInterface is 40x slower than function `Match`.
func (k Kind) MatchInterface(err error) bool {
	return causedBehaviourIFace(err, k)
}

// New wraps err with the specified Kind. Allows to write an additional message
// which gets formatted by fmt.Sprintf. It attaches a stack trace.
func (k Kind) New(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return kindStacked{withStack: errWrapf(err, msg, args...), Kind: k}
}

// Newf creates a new error with a message formatted by fmt.Sprintf and a stack trace.
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
	b = appendString(b, e.Kind.String())
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
	kRaw, b, err := getBytes(b)
	e.Kind = Kind(kRaw)
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
	b = appendString(b, e.Kind.String())
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
	kRaw, b, err := getBytes(b)
	e.Kind = Kind(kRaw)
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

// MatchKind returns true if `err` is of Kind `k`. It unwraps all underlying
// errors which implement the Causer interface. Does not supported implemented
// behaviour functions.
func MatchKind(err error, k Kind) bool {
	return k.Match(err)
}

// MatchAnyKind checks if at least one Kind is included in `err`. It does not
// unwrap `err` by its `Causer` interface. Does not supported implemented
// behaviour functions.
func MatchAnyKind(err error, k ...Kind) bool {
	if err == nil {
		return false
	}
	uk := UnwrapKind(err)
	return Kinds(k).matchAny(uk)
}

// MatchAllKind checks if all Kinds are included in `err`. It does not unwrap `err`
// by its `Causer` interface. Does not supported implemented behaviour
// functions.
func MatchAllKind(err error, k ...Kind) bool {
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
	return UnwrapKind(err).Unwrap()
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
	case *withMessage:
		k = UnwrapKind(e.Cause())
	case *withStack:
		k = UnwrapKind(e.Cause())
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
	if err == nil {
		return false
	}
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

func causedBehaviourIFace(err error, k Kind) bool {
	if k.matchInterface(err) {
		return true
	}
	for err != nil {
		var c causer
		if !As(err, &c) {
			break
		}
		err = c.Cause() // don't touch if you're unsure
		if k.matchInterface(err) {
			return true
		}
	}
	return false
}

// NoKind defines an empty Kind with no behaviour. This constant must be placed
// outside the constant block to avoid a conflict with iota.
const NoKind Kind = ""

// These constants define different behaviours. They are not sorted and new
// constants must be appended at the end. The zero kind defines empty.
const (
	Aborted            Kind = "Aborted"
	AlreadyCaptured    Kind = "AlreadyCaptured"
	AlreadyClosed      Kind = "AlreadyClosed"
	AlreadyExists      Kind = "AlreadyExists"
	AlreadyInUse       Kind = "AlreadyInUse"
	AlreadyRefunded    Kind = "AlreadyRefunded"
	BadEncoding        Kind = "BadEncoding"
	Blocked            Kind = "Blocked"
	ConnectionFailed   Kind = "ConnectionFailed"
	ConnectionLost     Kind = "ConnectionLost"
	CorruptData        Kind = "CorruptData"
	Declined           Kind = "Declined"
	DecryptionFailed   Kind = "DecryptionFailed"
	Denied             Kind = "Denied"
	Duplicated         Kind = "Duplicated"
	Empty              Kind = "Empty"
	EncryptionFailed   Kind = "EncryptionFailed"
	Exceeded           Kind = "Exceeded"
	Exists             Kind = "Exists"
	Expired            Kind = "Expired"
	Fatal              Kind = "Fatal"
	InProgress         Kind = "InProgress"
	Insufficient       Kind = "Insufficient"
	Interrupted        Kind = "Interrupted"
	IsDirectory        Kind = "IsDirectory"
	IsFile             Kind = "IsFile"
	Locked             Kind = "Locked"
	Mismatch           Kind = "Mismatch"
	NotAcceptable      Kind = "NotAcceptable"
	NotAllowed         Kind = "NotAllowed"
	NotDirectory       Kind = "NotDirectory"
	NotEmpty           Kind = "NotEmpty"
	NotExists          Kind = "NotExists"
	NotFile            Kind = "NotFile"
	NotFound           Kind = "NotFound"
	NotImplemented     Kind = "NotImplemented"
	NotRecoverable     Kind = "NotRecoverable"
	NotSupported       Kind = "NotSupported"
	NotValid           Kind = "NotValid"
	OutOfDate          Kind = "OutOfDate"
	OutOfRange         Kind = "OutOfRange"
	Overflowed         Kind = "Overflowed"
	PermissionDenied   Kind = "PermissionDenied"
	QuotaExceeded      Kind = "QuotaExceeded"
	ReadFailed         Kind = "ReadFailed"
	Rejected           Kind = "Rejected"
	Required           Kind = "Required"
	Restricted         Kind = "Restricted"
	Revoked            Kind = "Revoked"
	Temporary          Kind = "Temporary"
	Terminated         Kind = "Terminated"
	Timeout            Kind = "Timeout"
	TooLarge           Kind = "TooLarge"
	TooShort           Kind = "TooShort"
	Unauthorized       Kind = "Unauthorized"
	Unavailable        Kind = "Unavailable"
	UserNotFound       Kind = "UserNotFound"
	VerificationFailed Kind = "VerificationFailed"
	WriteFailed        Kind = "WriteFailed"
	WrongVersion       Kind = "WrongVersion"
)

type (
	iFaceAborted            interface{ Aborted() bool }
	iFaceAlreadyCaptured    interface{ AlreadyCaptured() bool }
	iFaceAlreadyClosed      interface{ AlreadyClosed() bool }
	iFaceAlreadyExists      interface{ AlreadyExists() bool }
	iFaceAlreadyInUse       interface{ AlreadyInUse() bool }
	iFaceAlreadyRefunded    interface{ AlreadyRefunded() bool }
	iFaceBadEncoding        interface{ BadEncoding() bool }
	iFaceBlocked            interface{ Blocked() bool }
	iFaceConnectionFailed   interface{ ConnectionFailed() bool }
	iFaceConnectionLost     interface{ ConnectionLost() bool }
	iFaceCorruptData        interface{ CorruptData() bool }
	iFaceDeclined           interface{ Declined() bool }
	iFaceDecryptionFailed   interface{ DecryptionFailed() bool }
	iFaceDenied             interface{ Denied() bool }
	iFaceDuplicated         interface{ Duplicated() bool }
	iFaceEmpty              interface{ Empty() bool }
	iFaceEncryptionFailed   interface{ EncryptionFailed() bool }
	iFaceExceeded           interface{ Exceeded() bool }
	iFaceExists             interface{ Exists() bool }
	iFaceExpired            interface{ Expired() bool }
	iFaceFatal              interface{ Fatal() bool }
	iFaceInProgress         interface{ InProgress() bool }
	iFaceInsufficient       interface{ Insufficient() bool }
	iFaceInterrupted        interface{ Interrupted() bool }
	iFaceIsDirectory        interface{ IsDirectory() bool }
	iFaceIsFile             interface{ IsFile() bool }
	iFaceLocked             interface{ Locked() bool }
	iFaceMismatch           interface{ Mismatch() bool }
	iFaceNotAcceptable      interface{ NotAcceptable() bool }
	iFaceNotAllowed         interface{ NotAllowed() bool }
	iFaceNotDirectory       interface{ NotDirectory() bool }
	iFaceNotEmpty           interface{ NotEmpty() bool }
	iFaceNotExists          interface{ NotExists() bool }
	iFaceNotFile            interface{ NotFile() bool }
	iFaceNotFound           interface{ NotFound() bool }
	iFaceNotImplemented     interface{ NotImplemented() bool }
	iFaceNotRecoverable     interface{ NotRecoverable() bool }
	iFaceNotSupported       interface{ NotSupported() bool }
	iFaceNotValid           interface{ NotValid() bool }
	iFaceOutOfDate          interface{ OutOfDate() bool }
	iFaceOutOfRange         interface{ OutOfRange() bool }
	iFaceOverflowed         interface{ Overflowed() bool }
	iFacePermissionDenied   interface{ PermissionDenied() bool }
	iFaceQuotaExceeded      interface{ QuotaExceeded() bool }
	iFaceReadFailed         interface{ ReadFailed() bool }
	iFaceRejected           interface{ Rejected() bool }
	iFaceRequired           interface{ Required() bool }
	iFaceRestricted         interface{ Restricted() bool }
	iFaceRevoked            interface{ Revoked() bool }
	iFaceTemporary          interface{ Temporary() bool }
	iFaceTerminated         interface{ Terminated() bool }
	iFaceTimeout            interface{ Timeout() bool }
	iFaceTooLarge           interface{ TooLarge() bool }
	iFaceTooShort           interface{ TooShort() bool }
	iFaceUnauthorized       interface{ Unauthorized() bool }
	iFaceUnavailable        interface{ Unavailable() bool }
	iFaceUserNotFound       interface{ UserNotFound() bool }
	iFaceVerificationFailed interface{ VerificationFailed() bool }
	iFaceWriteFailed        interface{ WriteFailed() bool }
	iFaceWrongVersion       interface{ WrongVersion() bool }
)
