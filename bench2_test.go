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
	"testing"
)

var benchmarkAsserted bool

type aExistsEStruct struct{}

func (a aExistsEStruct) Error() string   { return "Err" }
func (a aExistsEStruct) ErrorKind() Kind { return AlreadyExists }

func BenchmarkAssertBehaviourEmptyStruct(b *testing.B) {
	var ae = aExistsEStruct{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkAsserted = AlreadyExists.Match(ae)
		if !benchmarkAsserted {
			b.Errorf("Hell should already exists: %#v", ae)
		}
	}
}

type aExistsStr string

func (c aExistsStr) Error() string   { return string(c) }
func (c aExistsStr) ErrorKind() Kind { return AlreadyExists }

func BenchmarkAssertBehaviourConstant(b *testing.B) {
	const hell aExistsStr = "Hell"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkAsserted = AlreadyExists.Match(hell)
		if !benchmarkAsserted {
			b.Error("Hell should already exists.")
		}
	}
}

func BenchmarkAssertBehaviourPointer(b *testing.B) {
	var hell = AlreadyExists.New(errors.New("hell"), "There is already a place for you")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkAsserted = AlreadyExists.Match(hell)
		if !benchmarkAsserted {
			b.Error("Hell should already exists.")
		}
	}
}

func BenchmarkAssertBehaviourNoMatch(b *testing.B) {
	var hell = AlreadyClosed.New(errors.New("hell"), "There is already a place for you")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkAsserted = AlreadyExists.Match(hell)
		if benchmarkAsserted {
			b.Error("Hell should already be clsoed.")
		}
	}
}

var benchmarkMultiErr string

func BenchmarkMultiErr(b *testing.B) {
	e := NewMultiErr().
		AppendErrors(
			errors.New("Err5"),
			nil,
			errors.New("Err6"),
			errors.New("Err7"),
		)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkMultiErr = e.Error()
	}
}

var errorPointer = errors.New("I'm an error pointer")
var errorPointer2 = errors.New("I'm an error pointer2")

const errorConstant Error = `I'm an error constant`
const errorConstant2 Error = `I'm an error constant2`

var errorHave string

func BenchmarkMultiErrPointer(b *testing.B) {
	merr := NewMultiErr(errorPointer, errorPointer2)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errorHave = merr.Error()
		if errorHave == "" {
			b.Fatal("errorHave is empty")
		}
	}
}

func BenchmarkMultiErrConstant(b *testing.B) {
	merr := NewMultiErr(errorConstant, errorConstant2)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errorHave = merr.Error()
		if errorHave == "" {
			b.Fatal("errorHave is empty")
		}
	}
}
