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

import "fmt"

// BehaviourFunc defines the signature needed for a function to check
// if an error has a specific behaviour attached.
type BehaviourFunc func(error) bool

func errWrapf(err error, format string, args ...interface{}) wrapper {
	ret := wrapper{
		cause: cause{
			cause: err,
			msg:   format,
		},
		stack: callers(),
	}
	if len(args) > 0 {
		ret.cause.msg = fmt.Sprintf(format, args...)
	}
	return ret
}

func errNewf(format string, args ...interface{}) (ret _error) {
	ret.msg = format
	ret.stack = callers()
	if len(args) > 0 {
		ret.msg = fmt.Sprintf(format, args...)
	}
	return
}

// HasBehaviour checks if err contains at least one of the provided behaviour
// functions. Does not traverse recursive into the error.
func HasBehaviour(err error, bfs ...BehaviourFunc) bool {
	for _, f := range bfs {
		if f(err) {
			return true
		}
	}
	return false
}

// CausedBehaviour returns the first underlying caused behaviour of the error,
// if possible. An error value has a cause if it implements the following
// interface:
//
//     type Causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause or is nil, false will be returned. The
// variable bhf gets called on each unwrapped "cause" error
func CausedBehaviour(err error, bhf BehaviourFunc) bool {
	if bhf(err) {
		return true
	}
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			return false
		}
		err = cause.Cause() // don't touch if you're unsure
		if bhf(err) {
			return true
		}
	}
	return false
}
