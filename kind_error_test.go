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
	"testing"
)

var (
	_ error                      = (*Kind)(nil)
	_ encoding.BinaryMarshaler   = (*kindFundamental)(nil)
	_ encoding.BinaryMarshaler   = (*kindStacked)(nil)
	_ encoding.BinaryUnmarshaler = (*kindFundamental)(nil)
	_ encoding.BinaryUnmarshaler = (*kindStacked)(nil)
)

func TestKind_Error(t *testing.T) {
	const xErr1 Kind = "xErr1"
	if have, want := xErr1.Error(), "xErr1"; have != want {
		t.Errorf("xErr1 should return itself unformatted")
	}
}

func TestErrorDetach(t *testing.T) {
	tests := []struct {
		have Kind
		rm   Kind
		want Kind
	}{
		{"Failed", "Failed", ""},
		{"Failed", "", "Failed"},
		{"", "Failed", ""},
		{"Failed|Aborted", "Failed", "Aborted"},
		{"Failed|Aborted", "Aborted", "Failed"},
		{"Failed|Aborted", "Abort3d", "Failed|Aborted"},
		{"Failed|Aborted|Stuck", "Aborted", "Failed|Stuck"},
		{"Failed|Aborted|Stuck", "Stuck", "Failed|Aborted"},
		{"Failed|Aborted|Stuck", "Failed", "Aborted|Stuck"},
		{"Failed|Aborted|Stuck|", "Failed", "Aborted|Stuck"},
		{"|Failed|Aborted|Stuck|", "Failed", "Aborted|Stuck"},
	}
	for i, test := range tests {
		if want, have := test.want, test.have.detach(test.rm); want != have {
			t.Fatalf("want %q have  %q at index %d", want, have, i)
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

type Error string

func (e Error) Error() string { return string(e) }

type testBehave struct{ kind Kind }

func (nf testBehave) Error() string {
	return fmt.Sprintf("Has Error Kind %s", nf.kind)
}

func (nf testBehave) ErrorKind() Kind {
	return nf.kind
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

func TestUnwrapKind(t *testing.T) {
	var err error = testBehave{kind: Aborted}
	if have, want := UnwrapKind(err), Aborted; have != want {
		t.Errorf("Have %q Want %q", have, want)
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
	if have, want := Kind(0).String(), ""; have != want {
		t.Errorf("Have %q Want %q", have, want)
	}
	if have, want := Kind("UserNotFound|Temporary|Locked").String(), "UserNotFound|Temporary|Locked"; have != want {
		t.Errorf("Have %q Want %q", have, want)
	}
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

func assertFalse(t *testing.T, value bool, msgAndArgs ...interface{}) {
	if value {
		t.Errorf("Should be false: %t. Message: %v", value, msgAndArgs)
	}
}

func assertTrue(t *testing.T, value bool, msgAndArgs ...interface{}) {
	if !value {
		t.Errorf("Should be true: %t. Message: %v", value, msgAndArgs)
	}
}

func TestErrorInterfaces(t *testing.T) {
	t.Parallel()

	assertFalse(t, causedBehaviourIFace(errAborted{Kind: Aborted}, ""), "Aborted")
	assertFalse(t, causedBehaviourIFace(errAborted{Kind: Aborted}, "TODO"), "Aborted")
	assertFalse(t, causedBehaviourIFace(errOutOfDate{Kind: OutOfDate}, "TODO"), "OutOfDate")
	assertFalse(t, causedBehaviourIFace(errOutOfDate{Kind: CorruptData}, CorruptData), "CorruptData OutOfDate")
	assertFalse(t, Aborted.MatchInterface(Aborted.Newf("Ups")), "Should not match because Aborted does not have an Aborted function")

	assertTrue(t, Aborted.MatchInterface(errAborted{Kind: Aborted}), "Aborted")
	assertTrue(t, AlreadyCaptured.MatchInterface(errAlreadyCaptured{Kind: AlreadyCaptured}), "AlreadyCaptured")
	assertTrue(t, AlreadyClosed.MatchInterface(errAlreadyClosed{Kind: AlreadyClosed}), "AlreadyClosed")
	assertTrue(t, AlreadyExists.MatchInterface(errAlreadyExists{Kind: AlreadyExists}), "AlreadyExists")
	assertTrue(t, AlreadyInUse.MatchInterface(errAlreadyInUse{Kind: AlreadyInUse}), "AlreadyInUse")
	assertTrue(t, AlreadyRefunded.MatchInterface(errAlreadyRefunded{Kind: AlreadyRefunded}), "AlreadyRefunded")
	assertTrue(t, Blocked.MatchInterface(errBlocked{Kind: Blocked}), "Blocked")
	assertTrue(t, ReadFailed.MatchInterface(errReadFailed{Kind: ReadFailed}), "ReadFailed")
	assertTrue(t, WriteFailed.MatchInterface(errWriteFailed{Kind: WriteFailed}), "WriteFailed")
	assertTrue(t, VerificationFailed.MatchInterface(errVerificationFailed{Kind: VerificationFailed}), "VerificationFailed")
	assertTrue(t, DecryptionFailed.MatchInterface(errDecryptionFailed{Kind: DecryptionFailed}), "DecryptionFailed")
	assertTrue(t, EncryptionFailed.MatchInterface(errEncryptionFailed{Kind: EncryptionFailed}), "EncryptionFailed")
	assertTrue(t, ConnectionFailed.MatchInterface(errConnectionFailed{Kind: ConnectionFailed}), "ConnectionFailed")
	assertTrue(t, BadEncoding.MatchInterface(errBadEncoding{Kind: BadEncoding}), "BadEncoding")
	assertTrue(t, ConnectionLost.MatchInterface(errConnectionLost{Kind: ConnectionLost}), "ConnectionLost")
	assertTrue(t, Declined.MatchInterface(errDeclined{Kind: Declined}), "Declined")
	assertTrue(t, Denied.MatchInterface(errDenied{Kind: Denied}), "Denied")
	assertTrue(t, Duplicated.MatchInterface(errDuplicated{Kind: Duplicated}), "Duplicated")
	assertTrue(t, NotEmpty.MatchInterface(errNotEmpty{Kind: NotEmpty}), "NotEmpty")
	assertTrue(t, Empty.MatchInterface(errEmpty{Kind: Empty}), "Empty")
	assertTrue(t, Exceeded.MatchInterface(errExceeded{Kind: Exceeded}), "Exceeded")
	assertTrue(t, Exists.MatchInterface(errExists{Kind: Exists}), "Exists")
	assertTrue(t, NotExists.MatchInterface(errNotExists{Kind: NotExists}), "NotExists")
	assertTrue(t, Expired.MatchInterface(errExpired{Kind: Expired}), "Expired")
	assertTrue(t, Fatal.MatchInterface(errFatal{Kind: Fatal}), "Fatal")
	assertTrue(t, InProgress.MatchInterface(errInProgress{Kind: InProgress}), "InProgress")
	assertTrue(t, Insufficient.MatchInterface(errInsufficient{Kind: Insufficient}), "Insufficient")
	assertTrue(t, Interrupted.MatchInterface(errInterrupted{Kind: Interrupted}), "Interrupted")
	assertTrue(t, IsDirectory.MatchInterface(errIsDirectory{Kind: IsDirectory}), "IsDirectory")
	assertTrue(t, IsFile.MatchInterface(errIsFile{Kind: IsFile}), "IsFile")
	assertTrue(t, NotDirectory.MatchInterface(errNotDirectory{Kind: NotDirectory}), "NotDirectory")
	assertTrue(t, NotFile.MatchInterface(errNotFile{Kind: NotFile}), "NotFile")
	assertTrue(t, Locked.MatchInterface(errLocked{Kind: Locked}), "Locked")
	assertTrue(t, Mismatch.MatchInterface(errMismatch{Kind: Mismatch}), "Mismatch")
	assertTrue(t, NotAcceptable.MatchInterface(errNotAcceptable{Kind: NotAcceptable}), "NotAcceptable")
	assertTrue(t, NotAllowed.MatchInterface(errNotAllowed{Kind: NotAllowed}), "NotAllowed")
	assertTrue(t, NotFound.MatchInterface(errNotFound{Kind: NotFound}), "NotFound")
	assertTrue(t, NotImplemented.MatchInterface(errNotImplemented{Kind: NotImplemented}), "NotImplemented")
	assertTrue(t, NotRecoverable.MatchInterface(errNotRecoverable{Kind: NotRecoverable}), "NotRecoverable")
	assertTrue(t, NotSupported.MatchInterface(errNotSupported{Kind: NotSupported}), "NotSupported")
	assertTrue(t, NotValid.MatchInterface(errNotValid{Kind: NotValid}), "NotValid")
	assertTrue(t, Overflowed.MatchInterface(errOverflowed{Kind: Overflowed}), "Overflowed")
	assertTrue(t, PermissionDenied.MatchInterface(errPermissionDenied{Kind: PermissionDenied}), "PermissionDenied")
	assertTrue(t, Unauthorized.MatchInterface(errUnauthorized{Kind: Unauthorized}), "Unauthorized")
	assertTrue(t, UserNotFound.MatchInterface(errUserNotFound{Kind: UserNotFound}), "UserNotFound")
	assertTrue(t, QuotaExceeded.MatchInterface(errQuotaExceeded{Kind: QuotaExceeded}), "QuotaExceeded")
	assertTrue(t, Rejected.MatchInterface(errRejected{Kind: Rejected}), "Rejected")
	assertTrue(t, Required.MatchInterface(errRequired{Kind: Required}), "Required")
	assertTrue(t, Restricted.MatchInterface(errRestricted{Kind: Restricted}), "Restricted")
	assertTrue(t, Revoked.MatchInterface(errRevoked{Kind: Revoked}), "Revoked")
	assertTrue(t, Temporary.MatchInterface(errTemporary{Kind: Temporary}), "Temporary")
	assertTrue(t, Terminated.MatchInterface(errTerminated{Kind: Terminated}), "Terminated")
	assertTrue(t, Timeout.MatchInterface(errTimeout{Kind: Timeout}), "Timeout")
	assertTrue(t, TooLarge.MatchInterface(errTooLarge{Kind: TooLarge}), "TooLarge")
	assertTrue(t, Unavailable.MatchInterface(errUnavailable{Kind: Unavailable}), "Unavailable")
	assertTrue(t, WrongVersion.MatchInterface(errWrongVersion{Kind: WrongVersion}), "WrongVersion")
	assertTrue(t, CorruptData.MatchInterface(errCorruptData{Kind: CorruptData}), "CorruptData")
	assertTrue(t, OutOfRange.MatchInterface(errOutOfRange{Kind: OutOfRange}), "OutOfRange")
	assertTrue(t, OutOfDate.MatchInterface(errOutOfDate{Kind: OutOfDate}), "OutOfDate")
}
