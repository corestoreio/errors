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

package errors_test

import (
	"fmt"
	"testing"

	"github.com/corestoreio/errors"
	"github.com/corestoreio/pkg/util/assert"
)

var (
	_ fmt.Stringer = (*errors.Kind)(nil)
	_ fmt.Stringer = (*errors.Kinds)(nil)
)

func TestAttachDetach(t *testing.T) {
	t.Run("stacked", func(t *testing.T) {
		err := errors.New("test basic error")
		bErr := errors.AlreadyInUse.New(err, "Test Error %d!", 9876)
		bErr = errors.Attach(bErr, errors.Temporary)
		bErr = errors.Attach(bErr, errors.Unavailable)

		if !errors.Is(bErr, errors.Temporary) {
			t.Errorf("%#v should have Temporary", bErr)
		}
		if !errors.Is(bErr, errors.AlreadyInUse) {
			t.Errorf("%#v should have AlreadyInUse", bErr)
		}
		if !errors.Is(bErr, errors.Unavailable) {
			t.Errorf("%#v should have Unavailable", bErr)
		}

		bErr = errors.Detach(bErr, errors.AlreadyInUse)
		if errors.Is(bErr, errors.AlreadyInUse) {
			t.Errorf("%#v should have not AlreadyInUse", bErr)
		}

		bErr = errors.Detach(bErr, errors.Temporary)
		if errors.Is(bErr, errors.Temporary) {
			t.Errorf("%#v should have not Temporary", bErr)
		}
		bErr = errors.Detach(bErr, errors.ReadFailed)
		if errors.Is(bErr, errors.ReadFailed) {
			t.Errorf("%#v should have not ReadFailed", bErr)
		}
		if !errors.Is(bErr, errors.Unavailable) {
			t.Errorf("%#v should have Unavailable", bErr)
		}
	})
	t.Run("fundamental", func(t *testing.T) {
		bErr := errors.AlreadyInUse.Newf("Test Error %d!", 5678)
		bErr = errors.Attach(bErr, errors.Temporary)
		bErr = errors.Attach(bErr, errors.WrongVersion)

		if !errors.Is(bErr, errors.WrongVersion) {
			t.Errorf("%#v should have WrongVersion", bErr)
		}
		if !errors.Is(bErr, errors.Temporary) {
			t.Errorf("%#v should have Temporary", bErr)
		}
		if !errors.Is(bErr, errors.AlreadyInUse) {
			t.Errorf("%#v should have AlreadyInUse", bErr)
		}

		bErr = errors.Detach(bErr, errors.AlreadyInUse)
		if errors.Is(bErr, errors.AlreadyInUse) {
			t.Errorf("%#v should have not AlreadyInUse", bErr)
		}

		bErr = errors.Detach(bErr, errors.Temporary)
		if errors.Is(bErr, errors.Temporary) {
			t.Errorf("%#v should have not Temporary", bErr)
		}
		bErr = errors.Detach(bErr, errors.ReadFailed)
		if errors.Is(bErr, errors.ReadFailed) {
			t.Errorf("%#v should have not ReadFailed", bErr)
		}
		if !errors.Is(bErr, errors.WrongVersion) {
			t.Errorf("%#v should have WrongVersion", bErr)
		}
		if uk := errors.UnwrapKind(bErr); uk != errors.WrongVersion {
			t.Errorf("bErr should have WrongVersion, but got: %q", uk)
		}

		bErr = errors.Detach(bErr, errors.WrongVersion)
		if uk := errors.UnwrapKind(bErr); uk != "" {
			t.Errorf("bErr should have no kind, but got: %q", uk)
		}
	})
}

func TestCausedBehaviour(t *testing.T) {
	runner := func(err error, k errors.Kind, want bool) func(*testing.T) {
		return func(t *testing.T) {
			have := errors.CausedBehaviour(err, k)
			assert.Exactly(t, want, have, "%s", t.Name())
		}
	}
	t.Run("nil", runner(nil, errors.Fatal, false))
	t.Run("No cause", runner(errors.New("X"), errors.Fatal, false))
	t.Run("Fatal1", runner(errors.Fatal.Newf("X"), errors.Fatal, true))
	t.Run("Fatal2", runner(errors.Wrapf(errors.Fatal.Newf("X"), "wrap"), errors.Fatal, true))
	t.Run("Fatal3", runner(errors.Empty.New(errors.Wrapf(errors.Fatal.Newf("X"), "wrap"), "empty"), errors.Fatal, true))
	t.Run("Empty1", runner(errors.Empty.New(errors.Wrapf(errors.Fatal.Newf("X"), "wrap"), "empty"), errors.Empty, true))
	t.Run("Fatal4", runner(errors.Empty.New(errors.Wrapf(errors.Fatal.New(errors.New("X"), "fatal"), "wrap"), "empty"), errors.Fatal, true))
	t.Run("Empty2", runner(errors.Empty.New(errors.Wrapf(errors.Fatal.New(errors.New("X"), "fatal"), "wrap"), "empty"), errors.Empty, true))
	t.Run("AlreadyClosed", runner(errors.Empty.New(errors.AlreadyClosed.New(errors.Wrapf(errors.Fatal.New(errors.New("X"), "fatal"), "wrap"), "already closed"), "empty"), errors.AlreadyClosed, true))
}

func TestWrapf_Nil(t *testing.T) {
	e := errors.Wrapf(nil, "Error %d", 987654321)
	assert.Nil(t, e)
	e = errors.Wrapf(errors.WriteFailed.Newf("Damn it"), "Error %d", 987654321)
	assert.Error(t, e)
	e = errors.Wrapf(errors.WriteFailed, "Error")
	assert.Error(t, e)
}

func TestKind_Match(t *testing.T) {
	tests := []struct {
		err       error
		k         errors.Kind
		wantMatch bool
	}{
		{errors.Wrap(errors.AlreadyCaptured.Newf("error caused"), "outer"), errors.AlreadyCaptured, true},
		{errors.Wrap(errors.AlreadyCaptured.New(errors.NotAcceptable.Newf("error inner cause "), "error outer 1"), "outer2"), errors.NotAcceptable, true},
		{errors.AlreadyCaptured.Newf("error caused"), errors.AlreadyCaptured, true},
	}
	for i, test := range tests {
		if have, want := test.k.Match(test.err), test.wantMatch; have != want {
			t.Errorf("Index %d: Have %t Want %t", i, have, want)
		}
	}
}

func TestUnwrapKinds(t *testing.T) {
	tests := []struct {
		err  error
		want string
	}{
		{errors.Kind("UserNotFound|Temporary|Locked").Newf("User NF, Temp, Lock3d"), errors.Kinds{errors.UserNotFound, errors.Temporary, errors.Locked}.String()},
		{errors.Kind("UserNotFound").Newf("User NF"), errors.Kinds{errors.UserNotFound}.String()},
		{errors.New("usual error"), errors.Kinds{}.String()},
		{nil, errors.Kinds{}.String()},
		{
			errors.Kind("Aborted|AlreadyCaptured|AlreadyClosed|AlreadyExists|AlreadyInUse|AlreadyRefunded|Blocked|ReadFailed|WriteFailed|VerificationFailed|DecryptionFailed|EncryptionFailed|ConnectionFailed|BadEncoding|ConnectionLost|Declined|Denied|Duplicated|NotEmpty|Empty|Exceeded|Exists|NotExists|Expired|Fatal|InProgress|Insufficient|Interrupted|IsDirectory|IsFile|NotDirectory|NotFile|Locked|Mismatch|NotAcceptable|NotAllowed|NotFound|NotImplemented|NotRecoverable|NotSupported|NotValid|Overflowed|PermissionDenied|Unauthorized|UserNotFound|QuotaExceeded|Rejected|Required|Restricted|Revoked|Temporary|Terminated|Timeout|TooLarge|Unavailable|WrongVersion|CorruptData|OutOfRange|OutOfDate|TooShort").Newf("all constants"),
			errors.Kinds{errors.Aborted, errors.AlreadyCaptured, errors.AlreadyClosed, errors.AlreadyExists, errors.AlreadyInUse, errors.AlreadyRefunded, errors.Blocked, errors.ReadFailed, errors.WriteFailed, errors.VerificationFailed, errors.DecryptionFailed, errors.EncryptionFailed, errors.ConnectionFailed, errors.BadEncoding, errors.ConnectionLost, errors.Declined, errors.Denied, errors.Duplicated, errors.NotEmpty, errors.Empty, errors.Exceeded, errors.Exists, errors.NotExists, errors.Expired, errors.Fatal, errors.InProgress, errors.Insufficient, errors.Interrupted, errors.IsDirectory, errors.IsFile, errors.NotDirectory, errors.NotFile, errors.Locked, errors.Mismatch, errors.NotAcceptable, errors.NotAllowed, errors.NotFound, errors.NotImplemented, errors.NotRecoverable, errors.NotSupported, errors.NotValid, errors.Overflowed, errors.PermissionDenied, errors.Unauthorized, errors.UserNotFound, errors.QuotaExceeded, errors.Rejected, errors.Required, errors.Restricted, errors.Revoked, errors.Temporary, errors.Terminated, errors.Timeout, errors.TooLarge, errors.Unavailable, errors.WrongVersion, errors.CorruptData, errors.OutOfRange, errors.OutOfDate, errors.TooShort}.String(),
		},
	}
	for _, test := range tests {
		want, have := test.want, errors.UnwrapKinds(test.err)
		assert.Exactly(t, want, have.String())
	}
}

func TestKinds_String(t *testing.T) {
	assert.Exactly(t, "UserNotFound,Temporary,Locked", errors.Kinds{errors.UserNotFound, errors.Temporary, errors.Locked}.String())
	assert.Exactly(t, "", errors.Kinds{}.String())
}

func TestUnwrapStack(t *testing.T) {
	stack := errors.UnwrapStack(nil)
	assert.Nil(t, stack)
}

func TestMarshalling(t *testing.T) {
	t.Run("kindFundamental", func(t *testing.T) {
		const errTxt = "User NF, Temp, Locked"
		err := errors.Kind("UserNotFound|Temporary|Locked").Newf(errTxt)
		buf := errors.MarshalAppend(err, nil)
		sBuf := string(buf)
		// t.Logf("%q", sBuf)
		assert.Contains(t, sBuf, "F\x1dUserNotFound|Temporary|Locked\x15User NF, Temp, Locked\x9a\x03User NF, Temp, Locked")
		assert.Contains(t, sBuf, "errors.Kind.Newf\n\t")
		assert.Contains(t, sBuf, "errors/kind_error.go:")

		err = errors.Unmarshal(buf)
		assert.Exactly(t, "UserNotFound,Temporary,Locked", errors.UnwrapKinds(err).String())
		assert.Exactly(t, errTxt, err.Error())

		stack := errors.UnwrapStack(errors.WithStack(err))
		assert.Contains(t, string(stack), errTxt+"\n")
		assert.Contains(t, string(stack), "errors/kind_error.go:")
	})
	t.Run("kindStacked", func(t *testing.T) {
		const errTxt = "User NF, Temp, Locked"
		baseErr := errors.New(errTxt)
		err := errors.Kind("UserNotFound|Temporary|Locked").New(baseErr, "Marked the base error")
		buf := errors.MarshalAppend(err, nil)
		sBuf := string(buf)
		// t.Logf("%q", sBuf)
		assert.Contains(t, sBuf, "S\x1dUserNotFound|Temporary|Locked,Marked the base error: "+errTxt)
		assert.Contains(t, sBuf, "errors/kind_error.go:")

		err = errors.Unmarshal(buf)
		assert.Exactly(t, "UserNotFound,Temporary,Locked", errors.UnwrapKinds(err).String(), "%+v", err)
		assert.Exactly(t, ": Marked the base error: User NF, Temp, Locked", err.Error())

		stack := errors.UnwrapStack(err)
		assert.Contains(t, string(stack), "/errors.Kind.New\n")
	})
	t.Run("nil error", func(t *testing.T) {
		data := errors.MarshalAppend(nil, nil)
		assert.Nil(t, data)
	})
	t.Run("nil data", func(t *testing.T) {
		data := errors.Unmarshal(nil)
		assert.Nil(t, data)
	})
	t.Run("incorrect formatted data", func(t *testing.T) {
		err := errors.Unmarshal([]byte("x\x0eordinary error"))
		assert.True(t, errors.CorruptData.Match(err))
		assert.EqualError(t, err, "[errors] Unmarshal error: corrupt data \"\\x0eordinary error\"")

		err = errors.Unmarshal([]byte("e\xffordinary error"))
		assert.EqualError(t, err, "[errors] Unmarshal error[1]. Data length: 15")
		assert.True(t, errors.Is(err, errors.BadEncoding), "Should be of Kind BadEncoding")

		err = errors.Unmarshal([]byte("eordinary error"))
		assert.True(t, errors.Is(err, errors.BadEncoding), "Should be of Kind BadEncoding")
	})

	t.Run("ordinary error", func(t *testing.T) {
		oErr := errors.New("ordinary error")
		buf := errors.MarshalAppend(oErr, nil)
		assert.Exactly(t, "e\x0eordinary error", string(buf))

		err := errors.Unmarshal(buf)
		assert.EqualError(t, err, "ordinary error")
		assert.Exactly(t, errors.Kind(""), errors.UnwrapKind(err))
	})
}

func TestUnwrapKind_With(t *testing.T) {
	topErr := errors.Interrupted.Newf("Something has been interrupted")
	err := errors.Wrapf(topErr, "Oh ha")
	k := errors.UnwrapKind(err)
	assert.Exactly(t, errors.Interrupted, k)

	err = errors.WithMessage(topErr, "Another message")
	k = errors.UnwrapKind(err)
	assert.Exactly(t, errors.Interrupted, k)
}
