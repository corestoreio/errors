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

func BenchmarkAssertKindEmptyStruct(b *testing.B) {
	ae := aExistsEStruct{}
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

func BenchmarkAssertKindConstant(b *testing.B) {
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

func BenchmarkAssertKindPointer(b *testing.B) {
	var hell error = AlreadyExists.New(errors.New("hell"), "There is already a place for you")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkAsserted = AlreadyExists.Match(hell)
		if !benchmarkAsserted {
			b.Error("Hell should already exists.")
		}
	}
}

func BenchmarkAssertKindIFace(b *testing.B) {
	var hell error = errTerminated{Kind: Terminated}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkAsserted = Terminated.MatchInterface(hell)
		if !benchmarkAsserted {
			b.Error("Hell should ber terminated.")
		}
	}
}

func BenchmarkAssertKindNoMatch(b *testing.B) {
	b.Run("type", func(b *testing.B) {
		hell := AlreadyClosed.New(errors.New("hell"), "There is already a place for you")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchmarkAsserted = AlreadyExists.Match(hell)
			if benchmarkAsserted {
				b.Error("Hell should already be closed.")
			}
		}
	})
	b.Run("iface", func(b *testing.B) {
		hell := errors.New("not matched")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchmarkAsserted = AlreadyCaptured.MatchInterface(hell)
			if benchmarkAsserted {
				b.Error("Hell should not match")
			}
		}
	})
}

var benchmarkKindUnwrap Kinds

func BenchmarkKindUnwrap(b *testing.B) {
	b.ReportAllocs()
	const ek = Kind("Aborted|AlreadyCaptured|AlreadyClosed|AlreadyExists|AlreadyInUse|AlreadyRefunded|Blocked|ReadFailed|WriteFailed|VerificationFailed|DecryptionFailed|EncryptionFailed|ConnectionFailed|BadEncoding|ConnectionLost|Declined|Denied|Duplicated|NotEmpty|Empty|Exceeded|Exists|NotExists|Expired|Fatal|InProgress|Insufficient|Interrupted|IsDirectory|IsFile|NotDirectory|NotFile|Locked|Mismatch|NotAcceptable|NotAllowed|NotFound|NotImplemented|NotRecoverable|NotSupported|NotValid|Overflowed|PermissionDenied|Unauthorized|UserNotFound|QuotaExceeded|Rejected|Required|Restricted|Revoked|Temporary|Terminated|Timeout|TooLarge|Unavailable|WrongVersion|CorruptData|OutOfRange|OutOfDate|TooShort")
	for i := 0; i < b.N; i++ {
		benchmarkKindUnwrap = ek.Unwrap()
	}
}
