// Copyright 2015-2016, Cyrill @ Schumacher.fm and the CoreStore contributors
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
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testBehave struct{ ret bool }

func (nf testBehave) Error() string {
	return fmt.Sprintf("Has Error %t", nf.ret)
}

func TestCausedBehaviour(t *testing.T) {
	runner := func(err error, bhf BehaviourFunc, want bool) func(*testing.T) {
		return func(t *testing.T) {
			have := CausedBehaviour(err, bhf)
			assert.Exactly(t, want, have, "%s", t.Name())
		}
	}
	t.Run("No cause", runner(errors.New("X"), IsFatal, false))
	t.Run("IsFatal1", runner(NewFatalf("X"), IsFatal, true))
	t.Run("IsFatal2", runner(Wrapf(NewFatalf("X"), "wrap"), IsFatal, true))
	t.Run("IsFatal3", runner(NewEmpty(Wrapf(NewFatalf("X"), "wrap"), "empty"), IsFatal, true))
	t.Run("IsEmpty1", runner(NewEmpty(Wrapf(NewFatalf("X"), "wrap"), "empty"), IsEmpty, true))
	t.Run("IsFatal4", runner(NewEmpty(Wrapf(NewFatal(errors.New("X"), "fatal"), "wrap"), "empty"), IsFatal, true))
	t.Run("IsEmpty2", runner(NewEmpty(Wrapf(NewFatal(errors.New("X"), "fatal"), "wrap"), "empty"), IsEmpty, true))
	t.Run("IsAlreadyClosed", runner(NewEmpty(NewAlreadyClosed(Wrapf(NewFatal(errors.New("X"), "fatal"), "wrap"), "already closed"), "empty"), IsAlreadyClosed, true))
}

func TestError_Error(t *testing.T) {

	const e1 Error = "e1"
	assert.EqualError(t, e1, "e1")
}

func TestWrapf2(t *testing.T) {

	var e = Wrapf(nil, "Error %d")
	assert.Nil(t, e)
}

func TestErrorContainsAny(t *testing.T) {

	tests := []struct {
		me   error
		vf   []BehaviourFunc
		want bool
	}{
		{NewNotFoundf("e0"), []BehaviourFunc{IsNotFound}, true},
		{NewNotFoundf("e1"), []BehaviourFunc{IsNotValid}, false},
		{NewNotFoundf("e2"), []BehaviourFunc{IsNotValid, IsNotFound}, true},
		{NewNotFound(NewNotValidf("NotValid inner"), "NotFound outer"), []BehaviourFunc{IsNotValid, IsNotFound}, true},
		// once HasBehaviour acts recursive the next line will switch to true
		{NewNotFound(NewNotValidf("NotValid inner"), "NotFound outer"), []BehaviourFunc{IsNotValid}, true},
		{NewNotFound(NewNotValidf("NotValid inner"), "NotFound outer"), []BehaviourFunc{IsNotFound}, true},
		{NewNotFound(NewNotValidf("NotValid inner"), "NotFound outer"), []BehaviourFunc{IsAlreadyClosed}, false},
		{nil, []BehaviourFunc{IsNotValid}, false},
		{nil, nil, false},
	}

	for i, test := range tests {
		if have, want := HasBehaviour(test.me, test.vf...), test.want; have != want {
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
