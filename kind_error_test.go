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
	"encoding"
	"encoding/gob"
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	_ fmt.Stringer               = (*Kind)(nil)
	_ fmt.Stringer               = (*Kinds)(nil)
	_ encoding.BinaryMarshaler   = (*kindFundamental)(nil)
	_ encoding.BinaryMarshaler   = (*kindStacked)(nil)
	_ encoding.BinaryUnmarshaler = (*kindFundamental)(nil)
	_ encoding.BinaryUnmarshaler = (*kindStacked)(nil)
)

func TestAttachDetach(t *testing.T) {
	t.Run("stacked", func(t *testing.T) {
		err := errors.New("test basic error")
		bErr := AlreadyInUse.New(err, "Test Error %d!", 9876)
		bErr = Attach(bErr, Temporary)
		bErr = Attach(bErr, Unavailable)

		if !Is(bErr, Temporary) {
			t.Errorf("%#v should have Temporary", bErr)
		}
		if !Is(bErr, AlreadyInUse) {
			t.Errorf("%#v should have AlreadyInUse", bErr)
		}
		if !Is(bErr, Unavailable) {
			t.Errorf("%#v should have Unavailable", bErr)
		}

		bErr = Detach(bErr, AlreadyInUse)
		if Is(bErr, AlreadyInUse) {
			t.Errorf("%#v should have not AlreadyInUse", bErr)
		}

		bErr = Detach(bErr, Temporary)
		if Is(bErr, Temporary) {
			t.Errorf("%#v should have not Temporary", bErr)
		}
		bErr = Detach(bErr, ReadFailed)
		if Is(bErr, ReadFailed) {
			t.Errorf("%#v should have not ReadFailed", bErr)
		}
		if !Is(bErr, Unavailable) {
			t.Errorf("%#v should have Unavailable", bErr)
		}
	})
	t.Run("fundamental", func(t *testing.T) {

		bErr := AlreadyInUse.Newf("Test Error %d!", 5678)
		bErr = Attach(bErr, Temporary)
		bErr = Attach(bErr, WrongVersion)

		if !Is(bErr, WrongVersion) {
			t.Errorf("%#v should have WrongVersion", bErr)
		}
		if !Is(bErr, Temporary) {
			t.Errorf("%#v should have Temporary", bErr)
		}
		if !Is(bErr, AlreadyInUse) {
			t.Errorf("%#v should have AlreadyInUse", bErr)
		}

		bErr = Detach(bErr, AlreadyInUse)
		if Is(bErr, AlreadyInUse) {
			t.Errorf("%#v should have not AlreadyInUse", bErr)
		}

		bErr = Detach(bErr, Temporary)
		if Is(bErr, Temporary) {
			t.Errorf("%#v should have not Temporary", bErr)
		}
		bErr = Detach(bErr, ReadFailed)
		if Is(bErr, ReadFailed) {
			t.Errorf("%#v should have not ReadFailed", bErr)
		}
		if !Is(bErr, WrongVersion) {
			t.Errorf("%#v should have WrongVersion", bErr)
		}
		if uk := UnwrapKind(bErr); uk != WrongVersion {
			t.Errorf("bErr should have WrongVersion, but got: %q", uk)
		}

		bErr = Detach(bErr, WrongVersion)
		if uk := UnwrapKind(bErr); uk != 0 {
			t.Errorf("bErr should have no kind, but got: %q", uk)
		}
	})
}

type Error string

func (e Error) Error() string { return string(e) }

type testBehave struct{ kind Kind }

func (nf testBehave) Error() string {
	return fmt.Sprintf("Has Error Kind %s", nf.kind)
}

func TestCausedBehaviour(t *testing.T) {
	runner := func(err error, k Kind, want bool) func(*testing.T) {
		return func(t *testing.T) {
			have := CausedBehaviour(err, k)
			assert.Exactly(t, want, have, "%s", t.Name())
		}
	}
	t.Run("No cause", runner(errors.New("X"), Fatal, false))
	t.Run("Fatal1", runner(Fatal.Newf("X"), Fatal, true))
	t.Run("Fatal2", runner(Wrapf(Fatal.Newf("X"), "wrap"), Fatal, true))
	t.Run("Fatal3", runner(Empty.New(Wrapf(Fatal.Newf("X"), "wrap"), "empty"), Fatal, true))
	t.Run("Empty1", runner(Empty.New(Wrapf(Fatal.Newf("X"), "wrap"), "empty"), Empty, true))
	t.Run("Fatal4", runner(Empty.New(Wrapf(Fatal.New(errors.New("X"), "fatal"), "wrap"), "empty"), Fatal, true))
	t.Run("Empty2", runner(Empty.New(Wrapf(Fatal.New(errors.New("X"), "fatal"), "wrap"), "empty"), Empty, true))
	t.Run("AlreadyClosed", runner(Empty.New(AlreadyClosed.New(Wrapf(Fatal.New(errors.New("X"), "fatal"), "wrap"), "already closed"), "empty"), AlreadyClosed, true))
}

func TestError_Error(t *testing.T) {
	const e1 Error = "e1"
	assert.EqualError(t, e1, "e1")
}

func TestWrapf2(t *testing.T) {
	var e = Wrapf(nil, "Error %d")
	assert.Nil(t, e)
}

func TestKind_Match(t *testing.T) {
	tests := []struct {
		err       error
		k         Kind
		wantMatch bool
	}{
		{Wrap(AlreadyCaptured.Newf("error caused"), "outer"), AlreadyCaptured, true},
		{Wrap(AlreadyCaptured.New(NotAcceptable.Newf("error inner cause "), "error outer 1"), "outer2"), NotAcceptable, true},
		{AlreadyCaptured.Newf("error caused"), AlreadyCaptured, true},
	}
	for i, test := range tests {
		if have, want := test.k.Match(test.err), test.wantMatch; have != want {
			t.Errorf("Index %d: Have %t Want %t", i, have, want)
		}
	}
}

func TestMatchAll(t *testing.T) {
	tests := []struct {
		me   error
		ks   Kinds
		want bool
	}{
		{NotFound.Newf("e0"), Kinds{NotFound}, true},
		{NotFound.Newf("e1"), Kinds{NotValid}, false},
		{NotFound.Newf("e2"), Kinds{NotValid, NotFound}, false},
		{Attach(NotFound.Newf("e2"), NotValid), Kinds{NotValid, NotFound}, true},
		{NotFound.New(NotValid.Newf("NotValid inner"), "NotFound outer"), Kinds{NotValid, NotFound}, false},
		{NotFound.New(NotValid.Newf("NotValid inner"), "NotFound outer"), Kinds{NotFound}, true},
		{nil, Kinds{NotValid}, false},
		{nil, nil, false},
		{errors.New("hi there"), nil, false},
		{errors.New("hi there"), Kinds{NotValid}, false},
	}

	for i, test := range tests {
		if have, want := MatchAll(test.me, test.ks...), test.want; have != want {
			t.Errorf("Index %d: Have %t Want %t", i, have, want)
		}
	}
}

func TestMatchAny(t *testing.T) {

	tests := []struct {
		me   error
		ks   Kinds
		want bool
	}{
		{NotFound.Newf("e0"), Kinds{NotFound}, true},
		{NotFound.Newf("e1"), Kinds{NotValid}, false},
		{NotFound.Newf("e2"), Kinds{NotValid, NotFound}, true},
		{Attach(NotFound.Newf("e2"), NotValid), Kinds{NotValid, NotFound}, true},
		{NotFound.New(NotValid.Newf("NotValid inner"), "NotFound outer"), Kinds{NotValid, NotFound}, true},
		{NotFound.New(NotValid.Newf("NotValid inner"), "NotFound outer"), Kinds{NotFound}, true},
		{nil, Kinds{NotValid}, false},
		{nil, nil, false},
		{errors.New("hi there"), nil, false},
		{errors.New("hi there"), Kinds{NotValid}, false},
	}

	for i, test := range tests {
		if have, want := MatchAny(test.me, test.ks...), test.want; have != want {
			t.Errorf("Index %d: Have %t Want %t", i, have, want)
		}
	}
}

func TestErrWrapf(t *testing.T) {

	const e Error = "Error1"
	if haveEB, want := errWrapf(e, "Hello World %#v"), "Hello World %#v"; haveEB.error.(*withMessage).msg != want {
		t.Errorf("have %q want %q", haveEB.error.(*withMessage).msg, want)
	}
	if haveEB, want := errWrapf(e, "Hello World %d", 123), "Hello World 123"; haveEB.error.(*withMessage).msg != want {
		t.Errorf("have %q want %q", haveEB.error.(*withMessage).msg, want)
	}
}

func TestErrNewf(t *testing.T) {

	if have, want := errNewf("Hello World %d", 633), "Hello World 633"; have.msg != want {
		t.Errorf("have %q want %q", have.msg, want)
	}
	if have, want := errNewf("Hello World %d"), "Hello World %d"; have.msg != want {
		t.Errorf("have %q want %q", have.msg, want)
	}
}

func (nf testBehave) ErrorKind() Kind {
	return nf.kind
}

func TestUnwrapKind(t *testing.T) {
	var err error = testBehave{kind: Aborted}
	if have, want := UnwrapKind(err), Aborted; have != want {
		t.Errorf("Have %q Want %q", have, want)
	}
}

func TestUnwrapKinds(t *testing.T) {
	tests := []struct {
		err  error
		want string
	}{
		{(UserNotFound | Temporary | Locked).Newf("User NF, Temp, Locked"), Kinds{Locked, UserNotFound, Temporary}.String()},
		{(UserNotFound).Newf("User NF"), Kinds{UserNotFound}.String()},
		{errors.New("usual error"), Kinds{}.String()},
		{nil, Kinds{}.String()},
		{Kind(math.MaxUint64).Newf("all constants"), Kinds{Aborted, AlreadyCaptured, AlreadyClosed, AlreadyExists, AlreadyInUse, AlreadyRefunded, Blocked, ReadFailed, WriteFailed, VerificationFailed, DecryptionFailed, EncryptionFailed, ConnectionFailed, BadEncoding, ConnectionLost, Declined, Denied, Duplicated, NotEmpty, Empty, Exceeded, Exists, NotExists, Expired, Fatal, InProgress, Insufficient, Interrupted, IsDirectory, IsFile, NotDirectory, NotFile, Locked, Mismatch, NotAcceptable, NotAllowed, NotFound, NotImplemented, NotRecoverable, NotSupported, NotValid, Overflowed, PermissionDenied, Unauthorized, UserNotFound, QuotaExceeded, Rejected, Required, Restricted, Revoked, Temporary, Terminated, Timeout, TooLarge, Unavailable, WrongVersion, CorruptData, OutOfRange, OutOfDate, TooShort}.String()},
	}
	for _, test := range tests {
		want, have := test.want, UnwrapKinds(test.err)
		assert.Exactly(t, want, have.String())
	}
}

func TestKind_Nesting(t *testing.T) {
	tests := []struct {
		err  error
		kind Kind
		want bool
	}{
		{errors.New("Error1"), Aborted, false},
		{Aborted.New(nil, "Error2"), Aborted, false},
		{Aborted.New(Error("Error3a"), "Error3"), Aborted, true},
		{Wrap(Aborted.Newf("Err4"), "Wrap4"), Aborted, true},
		{NotImplemented.New(Wrap(Aborted.Newf("Err5"), "Wrap5"), "NotImplemend5"), Aborted, true},
		{Wrap(Aborted.New(Wrap(NotImplemented.Newf("Err6"), "Wrap6"), "Aborted6"), "Wrap6a"), Aborted, true},
		{Wrap(Aborted.New(errors.New("the cause7"), "Aborted7"), "Wrap7"), Aborted, true},
		{Aborted.Newf("Error8"), Aborted, true},
		{nil, Aborted, false},
		{testBehave{}, Aborted, false},
		{testBehave{kind: Aborted}, Aborted, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.kind.Match(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func TestKind_String(t *testing.T) {
	if have, want := Aborted.String(), "Aborted"; have != want {
		t.Errorf("Have %q Want %q", have, want)
	}
	if have, want := Kind(math.MaxUint32).String(), "Aborted,AlreadyCaptured,AlreadyClosed,AlreadyExists,AlreadyInUse,AlreadyRefunded,Blocked,ReadFailed,WriteFailed,VerificationFailed,DecryptionFailed,EncryptionFailed,ConnectionFailed,BadEncoding,ConnectionLost,Declined,Denied,Duplicated,NotEmpty,Empty,Exceeded,Exists,NotExists,Expired,Fatal,InProgress,Insufficient,Interrupted,IsDirectory,IsFile,NotDirectory,NotFile"; have != want {
		t.Errorf("Have %q Want %q", have, want)
	}
	if have, want := Kind(0).String(), "Kind(0)"; have != want {
		t.Errorf("Have %q Want %q", have, want)
	}
	if have, want := (UserNotFound | Temporary | Locked).String(), "Locked,UserNotFound,Temporary"; have != want {
		t.Errorf("Have %q Want %q", have, want)
	}
}

func TestKinds_String(t *testing.T) {
	assert.Exactly(t, "UserNotFound,Temporary,Locked", Kinds{UserNotFound, Temporary, Locked}.String())
	assert.Exactly(t, "", Kinds{}.String())
}

func TestUnwrapStack(t *testing.T) {
	stack := UnwrapStack(nil)
	assert.Nil(t, stack)
}

func TestMarshalling(t *testing.T) {
	t.Run("kindFundamental", func(t *testing.T) {
		const errTxt = "User NF, Temp, Locked"
		err := (UserNotFound | Temporary | Locked).Newf(errTxt)
		buf := MarshalAppend(err, nil)
		sBuf := string(buf)
		//t.Logf("%q", sBuf)
		assert.Contains(t, sBuf, "F\x80\x80\x80\x80\x90\x80\x84\x02\x15User NF, Temp, Locked\xca\x02User NF, Temp, Locked")
		assert.Contains(t, sBuf, "errors.Kind.Newf\n\t")
		assert.Contains(t, sBuf, "errors/kind_error.go:")

		err = Unmarshal(buf)
		assert.Exactly(t, "Locked,UserNotFound,Temporary", UnwrapKinds(err).String())
		assert.Exactly(t, errTxt, err.Error())

		stack := UnwrapStack(WithStack(err))
		assert.Contains(t, string(stack), errTxt+"\n")
		assert.Contains(t, string(stack), "errors/kind_error.go:")
	})
	t.Run("kindStacked", func(t *testing.T) {
		const errTxt = "User NF, Temp, Locked"
		baseErr := errors.New(errTxt)
		err := (UserNotFound | Temporary | Locked).New(baseErr, "Marked the base error")
		buf := MarshalAppend(err, nil)
		sBuf := string(buf)
		assert.Contains(t, sBuf, "S\x80\x80\x80\x80\x90\x80\x84\x02,Marked the base error: "+errTxt)
		assert.Contains(t, sBuf, "errors/kind_error.go:")

		err = Unmarshal(buf)
		assert.Exactly(t, "Locked,UserNotFound,Temporary", UnwrapKinds(err).String())
		assert.Exactly(t, ": Marked the base error: User NF, Temp, Locked", err.Error())

		stack := UnwrapStack(err)
		assert.Contains(t, string(stack), "/errors.Kind.New\n")
	})
	t.Run("nil error", func(t *testing.T) {
		data := MarshalAppend(nil, nil)
		assert.Nil(t, data)
	})
	t.Run("nil data", func(t *testing.T) {
		data := Unmarshal(nil)
		assert.Nil(t, data)
	})
	t.Run("incorrect formatted data", func(t *testing.T) {
		err := Unmarshal([]byte("x\x0eordinary error"))
		assert.True(t, CorruptData.Match(err))
		assert.EqualError(t, err, "[errors] Unmarshal error: corrupt data \"\\x0eordinary error\"")

		err = Unmarshal([]byte("e\xffordinary error"))
		assert.EqualError(t, err, "[errors] Unmarshal error[1]. Data length: 15")
		assert.True(t, Is(err, BadEncoding), "Should be of Kind BadEncoding")

		err = Unmarshal([]byte("eordinary error"))
		assert.True(t, Is(err, BadEncoding), "Should be of Kind BadEncoding")
	})

	t.Run("ordinary error", func(t *testing.T) {
		oErr := errors.New("ordinary error")
		buf := MarshalAppend(oErr, nil)
		assert.Exactly(t, "e\x0eordinary error", string(buf))

		err := Unmarshal(buf)
		assert.EqualError(t, err, "ordinary error")
		assert.Exactly(t, Kind(0), UnwrapKind(err))
	})
}

type gobEncErr struct {
	FieldA string
	Err    error
}

func init() {
	gob.Register(gobEncErr{})
	gob.Register(kindFundamental{})
	gob.Register(kindStacked{})
}

func TestKindFundamental_MarshalBinary(t *testing.T) {
	ge := gobEncErr{
		FieldA: "Hello World",
		Err:    InProgress.Newf("A process is in progress!"),
	}

	var buf bytes.Buffer
	require.NoError(t, gob.NewEncoder(&buf).Encode(ge), "Encoding should not fail")
	//t.Logf("%q", buf.Bytes())

	var ge2 gobEncErr
	require.NoError(t, gob.NewDecoder(&buf).Decode(&ge2), "Decoding should not fail")
	assert.Exactly(t, ge.FieldA, ge2.FieldA)
	assert.True(t, Is(ge2.Err, InProgress), "Should be a InProgress Kind")
	stack := UnwrapStack(ge2.Err)
	assert.Contains(t, string(stack), "A process is in progress!\ngithub.com/corestoreio/errors.Kind.Newf\n")
}

func TestKindStacked_MarshalBinary(t *testing.T) {

	baseErr := errors.New("@base error@")
	ge := gobEncErr{
		FieldA: "Hello Universe",
		Err:    NotAllowed.New(baseErr, "A process is not allowed!"),
	}

	var buf bytes.Buffer
	require.NoError(t, gob.NewEncoder(&buf).Encode(ge), "Encoding should not fail")
	//t.Logf("%q", buf.Bytes())

	var ge2 gobEncErr
	require.NoError(t, gob.NewDecoder(&buf).Decode(&ge2), "Decoding should not fail")
	assert.Exactly(t, ge.FieldA, ge2.FieldA)
	assert.True(t, Is(ge2.Err, NotAllowed), "Should be a NotAllowed Kind")
	stack := UnwrapStack(ge2.Err)
	assert.Contains(t, string(stack), "github.com/corestoreio/errors.Kind.New\n")
}

func TestUnwrapKind_With(t *testing.T) {

	topErr := Interrupted.Newf("Something has been interrupted")
	err := Wrapf(topErr, "Oh ha")
	k := UnwrapKind(err)
	assert.Exactly(t, Interrupted, k)

	err = WithMessage(topErr, "Another message")
	k = UnwrapKind(err)
	assert.Exactly(t, Interrupted, k)
}

type (
	// These struct fix a very weird bug or my understanding of Go. If the
	// struct would be just `struct{ Kind }` then the interface assertion in
	// function Kind.match would match to errEmpty.Kind.Empty instead of
	// errEmpty.Empty because in test packages errEmpty is private and the
	// tested function Kind.match cannot access the errEmpty.Empty function.
	errAborted            struct{ Kind Kind }
	errAlreadyCaptured    struct{ Kind Kind }
	errAlreadyClosed      struct{ Kind Kind }
	errAlreadyExists      struct{ Kind Kind }
	errAlreadyInUse       struct{ Kind Kind }
	errAlreadyRefunded    struct{ Kind Kind }
	errBlocked            struct{ Kind Kind }
	errReadFailed         struct{ Kind Kind }
	errWriteFailed        struct{ Kind Kind }
	errVerificationFailed struct{ Kind Kind }
	errDecryptionFailed   struct{ Kind Kind }
	errEncryptionFailed   struct{ Kind Kind }
	errConnectionFailed   struct{ Kind Kind }
	errBadEncoding        struct{ Kind Kind }
	errConnectionLost     struct{ Kind Kind }
	errDeclined           struct{ Kind Kind }
	errDenied             struct{ Kind Kind }
	errDuplicated         struct{ Kind Kind }
	errNotEmpty           struct{ Kind Kind }
	errEmpty              struct{ Kind Kind }
	errExceeded           struct{ Kind Kind }
	errExists             struct{ Kind Kind }
	errNotExists          struct{ Kind Kind }
	errExpired            struct{ Kind Kind }
	errFatal              struct{ Kind Kind }
	errInProgress         struct{ Kind Kind }
	errInsufficient       struct{ Kind Kind }
	errInterrupted        struct{ Kind Kind }
	errIsDirectory        struct{ Kind Kind }
	errIsFile             struct{ Kind Kind }
	errNotDirectory       struct{ Kind Kind }
	errNotFile            struct{ Kind Kind }
	errLocked             struct{ Kind Kind }
	errMismatch           struct{ Kind Kind }
	errNotAcceptable      struct{ Kind Kind }
	errNotAllowed         struct{ Kind Kind }
	errNotFound           struct{ Kind Kind }
	errNotImplemented     struct{ Kind Kind }
	errNotRecoverable     struct{ Kind Kind }
	errNotSupported       struct{ Kind Kind }
	errNotValid           struct{ Kind Kind }
	errOverflowed         struct{ Kind Kind }
	errPermissionDenied   struct{ Kind Kind }
	errUnauthorized       struct{ Kind Kind }
	errUserNotFound       struct{ Kind Kind }
	errQuotaExceeded      struct{ Kind Kind }
	errRejected           struct{ Kind Kind }
	errRequired           struct{ Kind Kind }
	errRestricted         struct{ Kind Kind }
	errRevoked            struct{ Kind Kind }
	errTemporary          struct{ Kind Kind }
	errTerminated         struct{ Kind Kind }
	errTimeout            struct{ Kind Kind }
	errTooLarge           struct{ Kind Kind }
	errUnavailable        struct{ Kind Kind }
	errWrongVersion       struct{ Kind Kind }
	errCorruptData        struct{ Kind Kind }
	errOutOfRange         struct{ Kind Kind }
	errOutOfDate          struct{ Kind Kind }
)

func (ae errAborted) Aborted() bool                       { return ae.Kind == Aborted }
func (ae errAlreadyCaptured) AlreadyCaptured() bool       { return ae.Kind == AlreadyCaptured }
func (ae errAlreadyClosed) AlreadyClosed() bool           { return ae.Kind == AlreadyClosed }
func (ae errAlreadyExists) AlreadyExists() bool           { return ae.Kind == AlreadyExists }
func (ae errAlreadyInUse) AlreadyInUse() bool             { return ae.Kind == AlreadyInUse }
func (ae errAlreadyRefunded) AlreadyRefunded() bool       { return ae.Kind == AlreadyRefunded }
func (ae errBlocked) Blocked() bool                       { return ae.Kind == Blocked }
func (ae errReadFailed) ReadFailed() bool                 { return ae.Kind == ReadFailed }
func (ae errWriteFailed) WriteFailed() bool               { return ae.Kind == WriteFailed }
func (ae errVerificationFailed) VerificationFailed() bool { return ae.Kind == VerificationFailed }
func (ae errDecryptionFailed) DecryptionFailed() bool     { return ae.Kind == DecryptionFailed }
func (ae errEncryptionFailed) EncryptionFailed() bool     { return ae.Kind == EncryptionFailed }
func (ae errConnectionFailed) ConnectionFailed() bool     { return ae.Kind == ConnectionFailed }
func (ae errBadEncoding) BadEncoding() bool               { return ae.Kind == BadEncoding }
func (ae errConnectionLost) ConnectionLost() bool         { return ae.Kind == ConnectionLost }
func (ae errDeclined) Declined() bool                     { return ae.Kind == Declined }
func (ae errDenied) Denied() bool                         { return ae.Kind == Denied }
func (ae errDuplicated) Duplicated() bool                 { return ae.Kind == Duplicated }
func (ae errNotEmpty) NotEmpty() bool                     { return ae.Kind == NotEmpty }
func (ae errEmpty) Empty() bool                           { return ae.Kind == Empty }
func (ae errExceeded) Exceeded() bool                     { return ae.Kind == Exceeded }
func (ae errExists) Exists() bool                         { return ae.Kind == Exists }
func (ae errNotExists) NotExists() bool                   { return ae.Kind == NotExists }
func (ae errExpired) Expired() bool                       { return ae.Kind == Expired }
func (ae errFatal) Fatal() bool                           { return ae.Kind == Fatal }
func (ae errInProgress) InProgress() bool                 { return ae.Kind == InProgress }
func (ae errInsufficient) Insufficient() bool             { return ae.Kind == Insufficient }
func (ae errInterrupted) Interrupted() bool               { return ae.Kind == Interrupted }
func (ae errIsDirectory) IsDirectory() bool               { return ae.Kind == IsDirectory }
func (ae errIsFile) IsFile() bool                         { return ae.Kind == IsFile }
func (ae errNotDirectory) NotDirectory() bool             { return ae.Kind == NotDirectory }
func (ae errNotFile) NotFile() bool                       { return ae.Kind == NotFile }
func (ae errLocked) Locked() bool                         { return ae.Kind == Locked }
func (ae errMismatch) Mismatch() bool                     { return ae.Kind == Mismatch }
func (ae errNotAcceptable) NotAcceptable() bool           { return ae.Kind == NotAcceptable }
func (ae errNotAllowed) NotAllowed() bool                 { return ae.Kind == NotAllowed }
func (ae errNotFound) NotFound() bool                     { return ae.Kind == NotFound }
func (ae errNotImplemented) NotImplemented() bool         { return ae.Kind == NotImplemented }
func (ae errNotRecoverable) NotRecoverable() bool         { return ae.Kind == NotRecoverable }
func (ae errNotSupported) NotSupported() bool             { return ae.Kind == NotSupported }
func (ae errNotValid) NotValid() bool                     { return ae.Kind == NotValid }
func (ae errOverflowed) Overflowed() bool                 { return ae.Kind == Overflowed }
func (ae errPermissionDenied) PermissionDenied() bool     { return ae.Kind == PermissionDenied }
func (ae errUnauthorized) Unauthorized() bool             { return ae.Kind == Unauthorized }
func (ae errUserNotFound) UserNotFound() bool             { return ae.Kind == UserNotFound }
func (ae errQuotaExceeded) QuotaExceeded() bool           { return ae.Kind == QuotaExceeded }
func (ae errRejected) Rejected() bool                     { return ae.Kind == Rejected }
func (ae errRequired) Required() bool                     { return ae.Kind == Required }
func (ae errRestricted) Restricted() bool                 { return ae.Kind == Restricted }
func (ae errRevoked) Revoked() bool                       { return ae.Kind == Revoked }
func (ae errTemporary) Temporary() bool                   { return ae.Kind == Temporary }
func (ae errTerminated) Terminated() bool                 { return ae.Kind == Terminated }
func (ae errTimeout) Timeout() bool                       { return ae.Kind == Timeout }
func (ae errTooLarge) TooLarge() bool                     { return ae.Kind == TooLarge }
func (ae errUnavailable) Unavailable() bool               { return ae.Kind == Unavailable }
func (ae errWrongVersion) WrongVersion() bool             { return ae.Kind == WrongVersion }
func (ae errCorruptData) CorruptData() bool               { return ae.Kind == CorruptData }
func (ae errOutOfRange) OutOfRange() bool                 { return ae.Kind == OutOfRange }
func (ae errOutOfDate) OutOfDate() bool                   { return ae.Kind == OutOfDate }

func (ae errAborted) Error() string            { return "Aborted!" }
func (ae errAlreadyCaptured) Error() string    { return "AlreadyCaptured!" }
func (ae errAlreadyClosed) Error() string      { return "AlreadyClosed!" }
func (ae errAlreadyExists) Error() string      { return "AlreadyExists!" }
func (ae errAlreadyInUse) Error() string       { return "AlreadyInUse!" }
func (ae errAlreadyRefunded) Error() string    { return "AlreadyRefunded!" }
func (ae errBlocked) Error() string            { return "Blocked!" }
func (ae errReadFailed) Error() string         { return "ReadFailed!" }
func (ae errWriteFailed) Error() string        { return "WriteFailed!" }
func (ae errVerificationFailed) Error() string { return "VerificationFailed!" }
func (ae errDecryptionFailed) Error() string   { return "DecryptionFailed!" }
func (ae errEncryptionFailed) Error() string   { return "EncryptionFailed!" }
func (ae errConnectionFailed) Error() string   { return "ConnectionFailed!" }
func (ae errBadEncoding) Error() string        { return "BadEncoding!" }
func (ae errConnectionLost) Error() string     { return "ConnectionLost!" }
func (ae errDeclined) Error() string           { return "Declined!" }
func (ae errDenied) Error() string             { return "Denied!" }
func (ae errDuplicated) Error() string         { return "Duplicated!" }
func (ae errNotEmpty) Error() string           { return "NotEmpty!" }
func (ae errEmpty) Error() string              { return "Empty!" }
func (ae errExceeded) Error() string           { return "Exceeded!" }
func (ae errExists) Error() string             { return "Exists!" }
func (ae errNotExists) Error() string          { return "NotExists!" }
func (ae errExpired) Error() string            { return "Expired!" }
func (ae errFatal) Error() string              { return "Fatal!" }
func (ae errInProgress) Error() string         { return "InProgress!" }
func (ae errInsufficient) Error() string       { return "Insufficient!" }
func (ae errInterrupted) Error() string        { return "Interrupted!" }
func (ae errIsDirectory) Error() string        { return "IsDirectory!" }
func (ae errIsFile) Error() string             { return "IsFile!" }
func (ae errNotDirectory) Error() string       { return "NotDirectory!" }
func (ae errNotFile) Error() string            { return "NotFile!" }
func (ae errLocked) Error() string             { return "Locked!" }
func (ae errMismatch) Error() string           { return "Mismatch!" }
func (ae errNotAcceptable) Error() string      { return "NotAcceptable!" }
func (ae errNotAllowed) Error() string         { return "NotAllowed!" }
func (ae errNotFound) Error() string           { return "NotFound!" }
func (ae errNotImplemented) Error() string     { return "NotImplemented!" }
func (ae errNotRecoverable) Error() string     { return "NotRecoverable!" }
func (ae errNotSupported) Error() string       { return "NotSupported!" }
func (ae errNotValid) Error() string           { return "NotValid!" }
func (ae errOverflowed) Error() string         { return "Overflowed!" }
func (ae errPermissionDenied) Error() string   { return "PermissionDenied!" }
func (ae errUnauthorized) Error() string       { return "Unauthorized!" }
func (ae errUserNotFound) Error() string       { return "UserNotFound!" }
func (ae errQuotaExceeded) Error() string      { return "QuotaExceeded!" }
func (ae errRejected) Error() string           { return "Rejected!" }
func (ae errRequired) Error() string           { return "Required!" }
func (ae errRestricted) Error() string         { return "Restricted!" }
func (ae errRevoked) Error() string            { return "Revoked!" }
func (ae errTemporary) Error() string          { return "Temporary!" }
func (ae errTerminated) Error() string         { return "Terminated!" }
func (ae errTimeout) Error() string            { return "Timeout!" }
func (ae errTooLarge) Error() string           { return "TooLarge!" }
func (ae errUnavailable) Error() string        { return "Unavailable!" }
func (ae errWrongVersion) Error() string       { return "WrongVersion!" }
func (ae errCorruptData) Error() string        { return "CorruptData!" }
func (ae errOutOfRange) Error() string         { return "OutOfRange!" }
func (ae errOutOfDate) Error() string          { return "OutOfDate!" }

func TestErrorInterfaces(t *testing.T) {
	t.Parallel()

	assert.False(t, causedBehaviourIFace(errAborted{Kind: Aborted}, 0), "Aborted")
	assert.False(t, causedBehaviourIFace(errAborted{Kind: Aborted}, maxKind), "Aborted")
	assert.False(t, causedBehaviourIFace(errOutOfDate{Kind: OutOfDate}, maxKind), "OutOfDate")
	assert.False(t, causedBehaviourIFace(errOutOfDate{Kind: CorruptData}, CorruptData), "CorruptData OutOfDate")
	assert.False(t, Aborted.MatchInterface(Aborted.Newf("Ups")), "Should not match because Aborted does not have an Aborted function")

	assert.True(t, Aborted.MatchInterface(errAborted{Kind: Aborted}), "Aborted")
	assert.True(t, AlreadyCaptured.MatchInterface(errAlreadyCaptured{Kind: AlreadyCaptured}), "AlreadyCaptured")
	assert.True(t, AlreadyClosed.MatchInterface(errAlreadyClosed{Kind: AlreadyClosed}), "AlreadyClosed")
	assert.True(t, AlreadyExists.MatchInterface(errAlreadyExists{Kind: AlreadyExists}), "AlreadyExists")
	assert.True(t, AlreadyInUse.MatchInterface(errAlreadyInUse{Kind: AlreadyInUse}), "AlreadyInUse")
	assert.True(t, AlreadyRefunded.MatchInterface(errAlreadyRefunded{Kind: AlreadyRefunded}), "AlreadyRefunded")
	assert.True(t, Blocked.MatchInterface(errBlocked{Kind: Blocked}), "Blocked")
	assert.True(t, ReadFailed.MatchInterface(errReadFailed{Kind: ReadFailed}), "ReadFailed")
	assert.True(t, WriteFailed.MatchInterface(errWriteFailed{Kind: WriteFailed}), "WriteFailed")
	assert.True(t, VerificationFailed.MatchInterface(errVerificationFailed{Kind: VerificationFailed}), "VerificationFailed")
	assert.True(t, DecryptionFailed.MatchInterface(errDecryptionFailed{Kind: DecryptionFailed}), "DecryptionFailed")
	assert.True(t, EncryptionFailed.MatchInterface(errEncryptionFailed{Kind: EncryptionFailed}), "EncryptionFailed")
	assert.True(t, ConnectionFailed.MatchInterface(errConnectionFailed{Kind: ConnectionFailed}), "ConnectionFailed")
	assert.True(t, BadEncoding.MatchInterface(errBadEncoding{Kind: BadEncoding}), "BadEncoding")
	assert.True(t, ConnectionLost.MatchInterface(errConnectionLost{Kind: ConnectionLost}), "ConnectionLost")
	assert.True(t, Declined.MatchInterface(errDeclined{Kind: Declined}), "Declined")
	assert.True(t, Denied.MatchInterface(errDenied{Kind: Denied}), "Denied")
	assert.True(t, Duplicated.MatchInterface(errDuplicated{Kind: Duplicated}), "Duplicated")
	assert.True(t, NotEmpty.MatchInterface(errNotEmpty{Kind: NotEmpty}), "NotEmpty")
	assert.True(t, Empty.MatchInterface(errEmpty{Kind: Empty}), "Empty")
	assert.True(t, Exceeded.MatchInterface(errExceeded{Kind: Exceeded}), "Exceeded")
	assert.True(t, Exists.MatchInterface(errExists{Kind: Exists}), "Exists")
	assert.True(t, NotExists.MatchInterface(errNotExists{Kind: NotExists}), "NotExists")
	assert.True(t, Expired.MatchInterface(errExpired{Kind: Expired}), "Expired")
	assert.True(t, Fatal.MatchInterface(errFatal{Kind: Fatal}), "Fatal")
	assert.True(t, InProgress.MatchInterface(errInProgress{Kind: InProgress}), "InProgress")
	assert.True(t, Insufficient.MatchInterface(errInsufficient{Kind: Insufficient}), "Insufficient")
	assert.True(t, Interrupted.MatchInterface(errInterrupted{Kind: Interrupted}), "Interrupted")
	assert.True(t, IsDirectory.MatchInterface(errIsDirectory{Kind: IsDirectory}), "IsDirectory")
	assert.True(t, IsFile.MatchInterface(errIsFile{Kind: IsFile}), "IsFile")
	assert.True(t, NotDirectory.MatchInterface(errNotDirectory{Kind: NotDirectory}), "NotDirectory")
	assert.True(t, NotFile.MatchInterface(errNotFile{Kind: NotFile}), "NotFile")
	assert.True(t, Locked.MatchInterface(errLocked{Kind: Locked}), "Locked")
	assert.True(t, Mismatch.MatchInterface(errMismatch{Kind: Mismatch}), "Mismatch")
	assert.True(t, NotAcceptable.MatchInterface(errNotAcceptable{Kind: NotAcceptable}), "NotAcceptable")
	assert.True(t, NotAllowed.MatchInterface(errNotAllowed{Kind: NotAllowed}), "NotAllowed")
	assert.True(t, NotFound.MatchInterface(errNotFound{Kind: NotFound}), "NotFound")
	assert.True(t, NotImplemented.MatchInterface(errNotImplemented{Kind: NotImplemented}), "NotImplemented")
	assert.True(t, NotRecoverable.MatchInterface(errNotRecoverable{Kind: NotRecoverable}), "NotRecoverable")
	assert.True(t, NotSupported.MatchInterface(errNotSupported{Kind: NotSupported}), "NotSupported")
	assert.True(t, NotValid.MatchInterface(errNotValid{Kind: NotValid}), "NotValid")
	assert.True(t, Overflowed.MatchInterface(errOverflowed{Kind: Overflowed}), "Overflowed")
	assert.True(t, PermissionDenied.MatchInterface(errPermissionDenied{Kind: PermissionDenied}), "PermissionDenied")
	assert.True(t, Unauthorized.MatchInterface(errUnauthorized{Kind: Unauthorized}), "Unauthorized")
	assert.True(t, UserNotFound.MatchInterface(errUserNotFound{Kind: UserNotFound}), "UserNotFound")
	assert.True(t, QuotaExceeded.MatchInterface(errQuotaExceeded{Kind: QuotaExceeded}), "QuotaExceeded")
	assert.True(t, Rejected.MatchInterface(errRejected{Kind: Rejected}), "Rejected")
	assert.True(t, Required.MatchInterface(errRequired{Kind: Required}), "Required")
	assert.True(t, Restricted.MatchInterface(errRestricted{Kind: Restricted}), "Restricted")
	assert.True(t, Revoked.MatchInterface(errRevoked{Kind: Revoked}), "Revoked")
	assert.True(t, Temporary.MatchInterface(errTemporary{Kind: Temporary}), "Temporary")
	assert.True(t, Terminated.MatchInterface(errTerminated{Kind: Terminated}), "Terminated")
	assert.True(t, Timeout.MatchInterface(errTimeout{Kind: Timeout}), "Timeout")
	assert.True(t, TooLarge.MatchInterface(errTooLarge{Kind: TooLarge}), "TooLarge")
	assert.True(t, Unavailable.MatchInterface(errUnavailable{Kind: Unavailable}), "Unavailable")
	assert.True(t, WrongVersion.MatchInterface(errWrongVersion{Kind: WrongVersion}), "WrongVersion")
	assert.True(t, CorruptData.MatchInterface(errCorruptData{Kind: CorruptData}), "CorruptData")
	assert.True(t, OutOfRange.MatchInterface(errOutOfRange{Kind: OutOfRange}), "OutOfRange")
	assert.True(t, OutOfDate.MatchInterface(errOutOfDate{Kind: OutOfDate}), "OutOfDate")
}
