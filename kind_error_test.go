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
	"encoding"
	"errors"
	"fmt"
	"math"
	"testing"

	"bytes"
	"encoding/gob"
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
		{Kind(math.MaxUint64).Newf("all constants"), Kinds{Aborted, AlreadyCaptured, AlreadyClosed, AlreadyExists, AlreadyInUse, AlreadyRefunded, Blocked, ReadFailed, WriteFailed, VerificationFailed, DecryptionFailed, EncryptionFailed, ConnectionFailed, BadEncoding, ConnectionLost, Declined, Denied, Duplicated, NotEmpty, Empty, Exceeded, Exists, NotExists, Expired, Fatal, InProgress, Insufficient, Interrupted, IsDirectory, IsFile, NotDirectory, NotFile, Locked, Mismatch, NotAcceptable, NotAllowed, NotFound, NotImplemented, NotRecoverable, NotSupported, NotValid, Overflowed, PermissionDenied, Unauthorized, UserNotFound, QuotaExceeded, Rejected, Required, Restricted, Revoked, Temporary, Terminated, Timeout, TooLarge, Unavailable, WrongVersion, CorruptData, OutofRange}.String()},
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
		assert.Contains(t, sBuf, "F\x80\x80\x80\x80\x90\x80\x84\x02\x15"+errTxt+"\xca\x02"+errTxt+"\ngithub.com")
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
