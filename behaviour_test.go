// Auto generated via behaviour_gen.go

package errors

import (
	"errors"
	"testing"
)

func (nf testBehave) Aborted() bool {
	return nf.ret
}

func TestAborted(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsAborted, false},
		{NewAborted(nil, "Error2"), IsAborted, false},
		{NewAborted(Error("Error3a"), "Error3"), IsAborted, true},
		{Wrap(NewAbortedf("Err4"), "Wrap4"), IsAborted, true},
		{NewNotImplemented(Wrap(NewAbortedf("Err5"), "Wrap5"), "NotImplemend5"), IsAborted, true},
		{Wrap(NewAborted(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Aborted6"), "Wrap6a"), IsAborted, true},
		{Wrap(NewAborted(errors.New("I'm the cause7"), "Aborted7"), "Wrap7"), IsAborted, true},
		{NewAbortedf("Error8"), IsAborted, true},
		{nil, IsAborted, false},
		{testBehave{}, IsAborted, false},
		{testBehave{ret: true}, IsAborted, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) AlreadyClosed() bool {
	return nf.ret
}

func TestAlreadyClosed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsAlreadyClosed, false},
		{NewAlreadyClosed(nil, "Error2"), IsAlreadyClosed, false},
		{NewAlreadyClosed(Error("Error3a"), "Error3"), IsAlreadyClosed, true},
		{Wrap(NewAlreadyClosedf("Err4"), "Wrap4"), IsAlreadyClosed, true},
		{NewNotImplemented(Wrap(NewAlreadyClosedf("Err5"), "Wrap5"), "NotImplemend5"), IsAlreadyClosed, true},
		{Wrap(NewAlreadyClosed(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "AlreadyClosed6"), "Wrap6a"), IsAlreadyClosed, true},
		{Wrap(NewAlreadyClosed(errors.New("I'm the cause7"), "AlreadyClosed7"), "Wrap7"), IsAlreadyClosed, true},
		{NewAlreadyClosedf("Error8"), IsAlreadyClosed, true},
		{nil, IsAlreadyClosed, false},
		{testBehave{}, IsAlreadyClosed, false},
		{testBehave{ret: true}, IsAlreadyClosed, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) AlreadyExists() bool {
	return nf.ret
}

func TestAlreadyExists(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsAlreadyExists, false},
		{NewAlreadyExists(nil, "Error2"), IsAlreadyExists, false},
		{NewAlreadyExists(Error("Error3a"), "Error3"), IsAlreadyExists, true},
		{Wrap(NewAlreadyExistsf("Err4"), "Wrap4"), IsAlreadyExists, true},
		{NewNotImplemented(Wrap(NewAlreadyExistsf("Err5"), "Wrap5"), "NotImplemend5"), IsAlreadyExists, true},
		{Wrap(NewAlreadyExists(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "AlreadyExists6"), "Wrap6a"), IsAlreadyExists, true},
		{Wrap(NewAlreadyExists(errors.New("I'm the cause7"), "AlreadyExists7"), "Wrap7"), IsAlreadyExists, true},
		{NewAlreadyExistsf("Error8"), IsAlreadyExists, true},
		{nil, IsAlreadyExists, false},
		{testBehave{}, IsAlreadyExists, false},
		{testBehave{ret: true}, IsAlreadyExists, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) AlreadyInUse() bool {
	return nf.ret
}

func TestAlreadyInUse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsAlreadyInUse, false},
		{NewAlreadyInUse(nil, "Error2"), IsAlreadyInUse, false},
		{NewAlreadyInUse(Error("Error3a"), "Error3"), IsAlreadyInUse, true},
		{Wrap(NewAlreadyInUsef("Err4"), "Wrap4"), IsAlreadyInUse, true},
		{NewNotImplemented(Wrap(NewAlreadyInUsef("Err5"), "Wrap5"), "NotImplemend5"), IsAlreadyInUse, true},
		{Wrap(NewAlreadyInUse(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "AlreadyInUse6"), "Wrap6a"), IsAlreadyInUse, true},
		{Wrap(NewAlreadyInUse(errors.New("I'm the cause7"), "AlreadyInUse7"), "Wrap7"), IsAlreadyInUse, true},
		{NewAlreadyInUsef("Error8"), IsAlreadyInUse, true},
		{nil, IsAlreadyInUse, false},
		{testBehave{}, IsAlreadyInUse, false},
		{testBehave{ret: true}, IsAlreadyInUse, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) AlreadyCaptured() bool {
	return nf.ret
}

func TestAlreadyCaptured(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsAlreadyCaptured, false},
		{NewAlreadyCaptured(nil, "Error2"), IsAlreadyCaptured, false},
		{NewAlreadyCaptured(Error("Error3a"), "Error3"), IsAlreadyCaptured, true},
		{Wrap(NewAlreadyCapturedf("Err4"), "Wrap4"), IsAlreadyCaptured, true},
		{NewNotImplemented(Wrap(NewAlreadyCapturedf("Err5"), "Wrap5"), "NotImplemend5"), IsAlreadyCaptured, true},
		{Wrap(NewAlreadyCaptured(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "AlreadyCaptured6"), "Wrap6a"), IsAlreadyCaptured, true},
		{Wrap(NewAlreadyCaptured(errors.New("I'm the cause7"), "AlreadyCaptured7"), "Wrap7"), IsAlreadyCaptured, true},
		{NewAlreadyCapturedf("Error8"), IsAlreadyCaptured, true},
		{nil, IsAlreadyCaptured, false},
		{testBehave{}, IsAlreadyCaptured, false},
		{testBehave{ret: true}, IsAlreadyCaptured, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) AlreadyRefunded() bool {
	return nf.ret
}

func TestAlreadyRefunded(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsAlreadyRefunded, false},
		{NewAlreadyRefunded(nil, "Error2"), IsAlreadyRefunded, false},
		{NewAlreadyRefunded(Error("Error3a"), "Error3"), IsAlreadyRefunded, true},
		{Wrap(NewAlreadyRefundedf("Err4"), "Wrap4"), IsAlreadyRefunded, true},
		{NewNotImplemented(Wrap(NewAlreadyRefundedf("Err5"), "Wrap5"), "NotImplemend5"), IsAlreadyRefunded, true},
		{Wrap(NewAlreadyRefunded(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "AlreadyRefunded6"), "Wrap6a"), IsAlreadyRefunded, true},
		{Wrap(NewAlreadyRefunded(errors.New("I'm the cause7"), "AlreadyRefunded7"), "Wrap7"), IsAlreadyRefunded, true},
		{NewAlreadyRefundedf("Error8"), IsAlreadyRefunded, true},
		{nil, IsAlreadyRefunded, false},
		{testBehave{}, IsAlreadyRefunded, false},
		{testBehave{ret: true}, IsAlreadyRefunded, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Blocked() bool {
	return nf.ret
}

func TestBlocked(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsBlocked, false},
		{NewBlocked(nil, "Error2"), IsBlocked, false},
		{NewBlocked(Error("Error3a"), "Error3"), IsBlocked, true},
		{Wrap(NewBlockedf("Err4"), "Wrap4"), IsBlocked, true},
		{NewNotImplemented(Wrap(NewBlockedf("Err5"), "Wrap5"), "NotImplemend5"), IsBlocked, true},
		{Wrap(NewBlocked(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Blocked6"), "Wrap6a"), IsBlocked, true},
		{Wrap(NewBlocked(errors.New("I'm the cause7"), "Blocked7"), "Wrap7"), IsBlocked, true},
		{NewBlockedf("Error8"), IsBlocked, true},
		{nil, IsBlocked, false},
		{testBehave{}, IsBlocked, false},
		{testBehave{ret: true}, IsBlocked, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) ConnectionFailed() bool {
	return nf.ret
}

func TestConnectionFailed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsConnectionFailed, false},
		{NewConnectionFailed(nil, "Error2"), IsConnectionFailed, false},
		{NewConnectionFailed(Error("Error3a"), "Error3"), IsConnectionFailed, true},
		{Wrap(NewConnectionFailedf("Err4"), "Wrap4"), IsConnectionFailed, true},
		{NewNotImplemented(Wrap(NewConnectionFailedf("Err5"), "Wrap5"), "NotImplemend5"), IsConnectionFailed, true},
		{Wrap(NewConnectionFailed(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "ConnectionFailed6"), "Wrap6a"), IsConnectionFailed, true},
		{Wrap(NewConnectionFailed(errors.New("I'm the cause7"), "ConnectionFailed7"), "Wrap7"), IsConnectionFailed, true},
		{NewConnectionFailedf("Error8"), IsConnectionFailed, true},
		{nil, IsConnectionFailed, false},
		{testBehave{}, IsConnectionFailed, false},
		{testBehave{ret: true}, IsConnectionFailed, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Declined() bool {
	return nf.ret
}

func TestDeclined(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsDeclined, false},
		{NewDeclined(nil, "Error2"), IsDeclined, false},
		{NewDeclined(Error("Error3a"), "Error3"), IsDeclined, true},
		{Wrap(NewDeclinedf("Err4"), "Wrap4"), IsDeclined, true},
		{NewNotImplemented(Wrap(NewDeclinedf("Err5"), "Wrap5"), "NotImplemend5"), IsDeclined, true},
		{Wrap(NewDeclined(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Declined6"), "Wrap6a"), IsDeclined, true},
		{Wrap(NewDeclined(errors.New("I'm the cause7"), "Declined7"), "Wrap7"), IsDeclined, true},
		{NewDeclinedf("Error8"), IsDeclined, true},
		{nil, IsDeclined, false},
		{testBehave{}, IsDeclined, false},
		{testBehave{ret: true}, IsDeclined, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Denied() bool {
	return nf.ret
}

func TestDenied(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsDenied, false},
		{NewDenied(nil, "Error2"), IsDenied, false},
		{NewDenied(Error("Error3a"), "Error3"), IsDenied, true},
		{Wrap(NewDeniedf("Err4"), "Wrap4"), IsDenied, true},
		{NewNotImplemented(Wrap(NewDeniedf("Err5"), "Wrap5"), "NotImplemend5"), IsDenied, true},
		{Wrap(NewDenied(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Denied6"), "Wrap6a"), IsDenied, true},
		{Wrap(NewDenied(errors.New("I'm the cause7"), "Denied7"), "Wrap7"), IsDenied, true},
		{NewDeniedf("Error8"), IsDenied, true},
		{nil, IsDenied, false},
		{testBehave{}, IsDenied, false},
		{testBehave{ret: true}, IsDenied, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Duplicated() bool {
	return nf.ret
}

func TestDuplicated(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsDuplicated, false},
		{NewDuplicated(nil, "Error2"), IsDuplicated, false},
		{NewDuplicated(Error("Error3a"), "Error3"), IsDuplicated, true},
		{Wrap(NewDuplicatedf("Err4"), "Wrap4"), IsDuplicated, true},
		{NewNotImplemented(Wrap(NewDuplicatedf("Err5"), "Wrap5"), "NotImplemend5"), IsDuplicated, true},
		{Wrap(NewDuplicated(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Duplicated6"), "Wrap6a"), IsDuplicated, true},
		{Wrap(NewDuplicated(errors.New("I'm the cause7"), "Duplicated7"), "Wrap7"), IsDuplicated, true},
		{NewDuplicatedf("Error8"), IsDuplicated, true},
		{nil, IsDuplicated, false},
		{testBehave{}, IsDuplicated, false},
		{testBehave{ret: true}, IsDuplicated, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Empty() bool {
	return nf.ret
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsEmpty, false},
		{NewEmpty(nil, "Error2"), IsEmpty, false},
		{NewEmpty(Error("Error3a"), "Error3"), IsEmpty, true},
		{Wrap(NewEmptyf("Err4"), "Wrap4"), IsEmpty, true},
		{NewNotImplemented(Wrap(NewEmptyf("Err5"), "Wrap5"), "NotImplemend5"), IsEmpty, true},
		{Wrap(NewEmpty(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Empty6"), "Wrap6a"), IsEmpty, true},
		{Wrap(NewEmpty(errors.New("I'm the cause7"), "Empty7"), "Wrap7"), IsEmpty, true},
		{NewEmptyf("Error8"), IsEmpty, true},
		{nil, IsEmpty, false},
		{testBehave{}, IsEmpty, false},
		{testBehave{ret: true}, IsEmpty, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Exceeded() bool {
	return nf.ret
}

func TestExceeded(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsExceeded, false},
		{NewExceeded(nil, "Error2"), IsExceeded, false},
		{NewExceeded(Error("Error3a"), "Error3"), IsExceeded, true},
		{Wrap(NewExceededf("Err4"), "Wrap4"), IsExceeded, true},
		{NewNotImplemented(Wrap(NewExceededf("Err5"), "Wrap5"), "NotImplemend5"), IsExceeded, true},
		{Wrap(NewExceeded(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Exceeded6"), "Wrap6a"), IsExceeded, true},
		{Wrap(NewExceeded(errors.New("I'm the cause7"), "Exceeded7"), "Wrap7"), IsExceeded, true},
		{NewExceededf("Error8"), IsExceeded, true},
		{nil, IsExceeded, false},
		{testBehave{}, IsExceeded, false},
		{testBehave{ret: true}, IsExceeded, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Expired() bool {
	return nf.ret
}

func TestExpired(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsExpired, false},
		{NewExpired(nil, "Error2"), IsExpired, false},
		{NewExpired(Error("Error3a"), "Error3"), IsExpired, true},
		{Wrap(NewExpiredf("Err4"), "Wrap4"), IsExpired, true},
		{NewNotImplemented(Wrap(NewExpiredf("Err5"), "Wrap5"), "NotImplemend5"), IsExpired, true},
		{Wrap(NewExpired(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Expired6"), "Wrap6a"), IsExpired, true},
		{Wrap(NewExpired(errors.New("I'm the cause7"), "Expired7"), "Wrap7"), IsExpired, true},
		{NewExpiredf("Error8"), IsExpired, true},
		{nil, IsExpired, false},
		{testBehave{}, IsExpired, false},
		{testBehave{ret: true}, IsExpired, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Fatal() bool {
	return nf.ret
}

func TestFatal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsFatal, false},
		{NewFatal(nil, "Error2"), IsFatal, false},
		{NewFatal(Error("Error3a"), "Error3"), IsFatal, true},
		{Wrap(NewFatalf("Err4"), "Wrap4"), IsFatal, true},
		{NewNotImplemented(Wrap(NewFatalf("Err5"), "Wrap5"), "NotImplemend5"), IsFatal, true},
		{Wrap(NewFatal(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Fatal6"), "Wrap6a"), IsFatal, true},
		{Wrap(NewFatal(errors.New("I'm the cause7"), "Fatal7"), "Wrap7"), IsFatal, true},
		{NewFatalf("Error8"), IsFatal, true},
		{nil, IsFatal, false},
		{testBehave{}, IsFatal, false},
		{testBehave{ret: true}, IsFatal, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) InProgress() bool {
	return nf.ret
}

func TestInProgress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsInProgress, false},
		{NewInProgress(nil, "Error2"), IsInProgress, false},
		{NewInProgress(Error("Error3a"), "Error3"), IsInProgress, true},
		{Wrap(NewInProgressf("Err4"), "Wrap4"), IsInProgress, true},
		{NewNotImplemented(Wrap(NewInProgressf("Err5"), "Wrap5"), "NotImplemend5"), IsInProgress, true},
		{Wrap(NewInProgress(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "InProgress6"), "Wrap6a"), IsInProgress, true},
		{Wrap(NewInProgress(errors.New("I'm the cause7"), "InProgress7"), "Wrap7"), IsInProgress, true},
		{NewInProgressf("Error8"), IsInProgress, true},
		{nil, IsInProgress, false},
		{testBehave{}, IsInProgress, false},
		{testBehave{ret: true}, IsInProgress, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Insufficient() bool {
	return nf.ret
}

func TestInsufficient(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsInsufficient, false},
		{NewInsufficient(nil, "Error2"), IsInsufficient, false},
		{NewInsufficient(Error("Error3a"), "Error3"), IsInsufficient, true},
		{Wrap(NewInsufficientf("Err4"), "Wrap4"), IsInsufficient, true},
		{NewNotImplemented(Wrap(NewInsufficientf("Err5"), "Wrap5"), "NotImplemend5"), IsInsufficient, true},
		{Wrap(NewInsufficient(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Insufficient6"), "Wrap6a"), IsInsufficient, true},
		{Wrap(NewInsufficient(errors.New("I'm the cause7"), "Insufficient7"), "Wrap7"), IsInsufficient, true},
		{NewInsufficientf("Error8"), IsInsufficient, true},
		{nil, IsInsufficient, false},
		{testBehave{}, IsInsufficient, false},
		{testBehave{ret: true}, IsInsufficient, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Interrupted() bool {
	return nf.ret
}

func TestInterrupted(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsInterrupted, false},
		{NewInterrupted(nil, "Error2"), IsInterrupted, false},
		{NewInterrupted(Error("Error3a"), "Error3"), IsInterrupted, true},
		{Wrap(NewInterruptedf("Err4"), "Wrap4"), IsInterrupted, true},
		{NewNotImplemented(Wrap(NewInterruptedf("Err5"), "Wrap5"), "NotImplemend5"), IsInterrupted, true},
		{Wrap(NewInterrupted(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Interrupted6"), "Wrap6a"), IsInterrupted, true},
		{Wrap(NewInterrupted(errors.New("I'm the cause7"), "Interrupted7"), "Wrap7"), IsInterrupted, true},
		{NewInterruptedf("Error8"), IsInterrupted, true},
		{nil, IsInterrupted, false},
		{testBehave{}, IsInterrupted, false},
		{testBehave{ret: true}, IsInterrupted, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Locked() bool {
	return nf.ret
}

func TestLocked(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsLocked, false},
		{NewLocked(nil, "Error2"), IsLocked, false},
		{NewLocked(Error("Error3a"), "Error3"), IsLocked, true},
		{Wrap(NewLockedf("Err4"), "Wrap4"), IsLocked, true},
		{NewNotImplemented(Wrap(NewLockedf("Err5"), "Wrap5"), "NotImplemend5"), IsLocked, true},
		{Wrap(NewLocked(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Locked6"), "Wrap6a"), IsLocked, true},
		{Wrap(NewLocked(errors.New("I'm the cause7"), "Locked7"), "Wrap7"), IsLocked, true},
		{NewLockedf("Error8"), IsLocked, true},
		{nil, IsLocked, false},
		{testBehave{}, IsLocked, false},
		{testBehave{ret: true}, IsLocked, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Mismatch() bool {
	return nf.ret
}

func TestMismatch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsMismatch, false},
		{NewMismatch(nil, "Error2"), IsMismatch, false},
		{NewMismatch(Error("Error3a"), "Error3"), IsMismatch, true},
		{Wrap(NewMismatchf("Err4"), "Wrap4"), IsMismatch, true},
		{NewNotImplemented(Wrap(NewMismatchf("Err5"), "Wrap5"), "NotImplemend5"), IsMismatch, true},
		{Wrap(NewMismatch(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Mismatch6"), "Wrap6a"), IsMismatch, true},
		{Wrap(NewMismatch(errors.New("I'm the cause7"), "Mismatch7"), "Wrap7"), IsMismatch, true},
		{NewMismatchf("Error8"), IsMismatch, true},
		{nil, IsMismatch, false},
		{testBehave{}, IsMismatch, false},
		{testBehave{ret: true}, IsMismatch, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) NotAcceptable() bool {
	return nf.ret
}

func TestNotAcceptable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsNotAcceptable, false},
		{NewNotAcceptable(nil, "Error2"), IsNotAcceptable, false},
		{NewNotAcceptable(Error("Error3a"), "Error3"), IsNotAcceptable, true},
		{Wrap(NewNotAcceptablef("Err4"), "Wrap4"), IsNotAcceptable, true},
		{NewNotImplemented(Wrap(NewNotAcceptablef("Err5"), "Wrap5"), "NotImplemend5"), IsNotAcceptable, true},
		{Wrap(NewNotAcceptable(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "NotAcceptable6"), "Wrap6a"), IsNotAcceptable, true},
		{Wrap(NewNotAcceptable(errors.New("I'm the cause7"), "NotAcceptable7"), "Wrap7"), IsNotAcceptable, true},
		{NewNotAcceptablef("Error8"), IsNotAcceptable, true},
		{nil, IsNotAcceptable, false},
		{testBehave{}, IsNotAcceptable, false},
		{testBehave{ret: true}, IsNotAcceptable, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) NotAllowed() bool {
	return nf.ret
}

func TestNotAllowed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsNotAllowed, false},
		{NewNotAllowed(nil, "Error2"), IsNotAllowed, false},
		{NewNotAllowed(Error("Error3a"), "Error3"), IsNotAllowed, true},
		{Wrap(NewNotAllowedf("Err4"), "Wrap4"), IsNotAllowed, true},
		{NewNotImplemented(Wrap(NewNotAllowedf("Err5"), "Wrap5"), "NotImplemend5"), IsNotAllowed, true},
		{Wrap(NewNotAllowed(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "NotAllowed6"), "Wrap6a"), IsNotAllowed, true},
		{Wrap(NewNotAllowed(errors.New("I'm the cause7"), "NotAllowed7"), "Wrap7"), IsNotAllowed, true},
		{NewNotAllowedf("Error8"), IsNotAllowed, true},
		{nil, IsNotAllowed, false},
		{testBehave{}, IsNotAllowed, false},
		{testBehave{ret: true}, IsNotAllowed, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) NotFound() bool {
	return nf.ret
}

func TestNotFound(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsNotFound, false},
		{NewNotFound(nil, "Error2"), IsNotFound, false},
		{NewNotFound(Error("Error3a"), "Error3"), IsNotFound, true},
		{Wrap(NewNotFoundf("Err4"), "Wrap4"), IsNotFound, true},
		{NewNotImplemented(Wrap(NewNotFoundf("Err5"), "Wrap5"), "NotImplemend5"), IsNotFound, true},
		{Wrap(NewNotFound(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "NotFound6"), "Wrap6a"), IsNotFound, true},
		{Wrap(NewNotFound(errors.New("I'm the cause7"), "NotFound7"), "Wrap7"), IsNotFound, true},
		{NewNotFoundf("Error8"), IsNotFound, true},
		{nil, IsNotFound, false},
		{testBehave{}, IsNotFound, false},
		{testBehave{ret: true}, IsNotFound, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) NotImplemented() bool {
	return nf.ret
}

func TestNotImplemented(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsNotImplemented, false},
		{NewNotImplemented(nil, "Error2"), IsNotImplemented, false},
		{NewNotImplemented(Error("Error3a"), "Error3"), IsNotImplemented, true},
		{Wrap(NewNotImplementedf("Err4"), "Wrap4"), IsNotImplemented, true},
		{NewNotImplemented(Wrap(NewNotImplementedf("Err5"), "Wrap5"), "NotImplemend5"), IsNotImplemented, true},
		{Wrap(NewNotImplemented(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "NotImplemented6"), "Wrap6a"), IsNotImplemented, true},
		{Wrap(NewNotImplemented(errors.New("I'm the cause7"), "NotImplemented7"), "Wrap7"), IsNotImplemented, true},
		{NewNotImplementedf("Error8"), IsNotImplemented, true},
		{nil, IsNotImplemented, false},
		{testBehave{}, IsNotImplemented, false},
		{testBehave{ret: true}, IsNotImplemented, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) NotRecoverable() bool {
	return nf.ret
}

func TestNotRecoverable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsNotRecoverable, false},
		{NewNotRecoverable(nil, "Error2"), IsNotRecoverable, false},
		{NewNotRecoverable(Error("Error3a"), "Error3"), IsNotRecoverable, true},
		{Wrap(NewNotRecoverablef("Err4"), "Wrap4"), IsNotRecoverable, true},
		{NewNotImplemented(Wrap(NewNotRecoverablef("Err5"), "Wrap5"), "NotImplemend5"), IsNotRecoverable, true},
		{Wrap(NewNotRecoverable(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "NotRecoverable6"), "Wrap6a"), IsNotRecoverable, true},
		{Wrap(NewNotRecoverable(errors.New("I'm the cause7"), "NotRecoverable7"), "Wrap7"), IsNotRecoverable, true},
		{NewNotRecoverablef("Error8"), IsNotRecoverable, true},
		{nil, IsNotRecoverable, false},
		{testBehave{}, IsNotRecoverable, false},
		{testBehave{ret: true}, IsNotRecoverable, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) NotSupported() bool {
	return nf.ret
}

func TestNotSupported(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsNotSupported, false},
		{NewNotSupported(nil, "Error2"), IsNotSupported, false},
		{NewNotSupported(Error("Error3a"), "Error3"), IsNotSupported, true},
		{Wrap(NewNotSupportedf("Err4"), "Wrap4"), IsNotSupported, true},
		{NewNotImplemented(Wrap(NewNotSupportedf("Err5"), "Wrap5"), "NotImplemend5"), IsNotSupported, true},
		{Wrap(NewNotSupported(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "NotSupported6"), "Wrap6a"), IsNotSupported, true},
		{Wrap(NewNotSupported(errors.New("I'm the cause7"), "NotSupported7"), "Wrap7"), IsNotSupported, true},
		{NewNotSupportedf("Error8"), IsNotSupported, true},
		{nil, IsNotSupported, false},
		{testBehave{}, IsNotSupported, false},
		{testBehave{ret: true}, IsNotSupported, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) NotValid() bool {
	return nf.ret
}

func TestNotValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsNotValid, false},
		{NewNotValid(nil, "Error2"), IsNotValid, false},
		{NewNotValid(Error("Error3a"), "Error3"), IsNotValid, true},
		{Wrap(NewNotValidf("Err4"), "Wrap4"), IsNotValid, true},
		{NewNotImplemented(Wrap(NewNotValidf("Err5"), "Wrap5"), "NotImplemend5"), IsNotValid, true},
		{Wrap(NewNotValid(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "NotValid6"), "Wrap6a"), IsNotValid, true},
		{Wrap(NewNotValid(errors.New("I'm the cause7"), "NotValid7"), "Wrap7"), IsNotValid, true},
		{NewNotValidf("Error8"), IsNotValid, true},
		{nil, IsNotValid, false},
		{testBehave{}, IsNotValid, false},
		{testBehave{ret: true}, IsNotValid, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) PermissionDenied() bool {
	return nf.ret
}

func TestPermissionDenied(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsPermissionDenied, false},
		{NewPermissionDenied(nil, "Error2"), IsPermissionDenied, false},
		{NewPermissionDenied(Error("Error3a"), "Error3"), IsPermissionDenied, true},
		{Wrap(NewPermissionDeniedf("Err4"), "Wrap4"), IsPermissionDenied, true},
		{NewNotImplemented(Wrap(NewPermissionDeniedf("Err5"), "Wrap5"), "NotImplemend5"), IsPermissionDenied, true},
		{Wrap(NewPermissionDenied(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "PermissionDenied6"), "Wrap6a"), IsPermissionDenied, true},
		{Wrap(NewPermissionDenied(errors.New("I'm the cause7"), "PermissionDenied7"), "Wrap7"), IsPermissionDenied, true},
		{NewPermissionDeniedf("Error8"), IsPermissionDenied, true},
		{nil, IsPermissionDenied, false},
		{testBehave{}, IsPermissionDenied, false},
		{testBehave{ret: true}, IsPermissionDenied, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) QuotaExceeded() bool {
	return nf.ret
}

func TestQuotaExceeded(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsQuotaExceeded, false},
		{NewQuotaExceeded(nil, "Error2"), IsQuotaExceeded, false},
		{NewQuotaExceeded(Error("Error3a"), "Error3"), IsQuotaExceeded, true},
		{Wrap(NewQuotaExceededf("Err4"), "Wrap4"), IsQuotaExceeded, true},
		{NewNotImplemented(Wrap(NewQuotaExceededf("Err5"), "Wrap5"), "NotImplemend5"), IsQuotaExceeded, true},
		{Wrap(NewQuotaExceeded(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "QuotaExceeded6"), "Wrap6a"), IsQuotaExceeded, true},
		{Wrap(NewQuotaExceeded(errors.New("I'm the cause7"), "QuotaExceeded7"), "Wrap7"), IsQuotaExceeded, true},
		{NewQuotaExceededf("Error8"), IsQuotaExceeded, true},
		{nil, IsQuotaExceeded, false},
		{testBehave{}, IsQuotaExceeded, false},
		{testBehave{ret: true}, IsQuotaExceeded, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) ReadFailed() bool {
	return nf.ret
}

func TestReadFailed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsReadFailed, false},
		{NewReadFailed(nil, "Error2"), IsReadFailed, false},
		{NewReadFailed(Error("Error3a"), "Error3"), IsReadFailed, true},
		{Wrap(NewReadFailedf("Err4"), "Wrap4"), IsReadFailed, true},
		{NewNotImplemented(Wrap(NewReadFailedf("Err5"), "Wrap5"), "NotImplemend5"), IsReadFailed, true},
		{Wrap(NewReadFailed(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "ReadFailed6"), "Wrap6a"), IsReadFailed, true},
		{Wrap(NewReadFailed(errors.New("I'm the cause7"), "ReadFailed7"), "Wrap7"), IsReadFailed, true},
		{NewReadFailedf("Error8"), IsReadFailed, true},
		{nil, IsReadFailed, false},
		{testBehave{}, IsReadFailed, false},
		{testBehave{ret: true}, IsReadFailed, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Rejected() bool {
	return nf.ret
}

func TestRejected(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsRejected, false},
		{NewRejected(nil, "Error2"), IsRejected, false},
		{NewRejected(Error("Error3a"), "Error3"), IsRejected, true},
		{Wrap(NewRejectedf("Err4"), "Wrap4"), IsRejected, true},
		{NewNotImplemented(Wrap(NewRejectedf("Err5"), "Wrap5"), "NotImplemend5"), IsRejected, true},
		{Wrap(NewRejected(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Rejected6"), "Wrap6a"), IsRejected, true},
		{Wrap(NewRejected(errors.New("I'm the cause7"), "Rejected7"), "Wrap7"), IsRejected, true},
		{NewRejectedf("Error8"), IsRejected, true},
		{nil, IsRejected, false},
		{testBehave{}, IsRejected, false},
		{testBehave{ret: true}, IsRejected, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Required() bool {
	return nf.ret
}

func TestRequired(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsRequired, false},
		{NewRequired(nil, "Error2"), IsRequired, false},
		{NewRequired(Error("Error3a"), "Error3"), IsRequired, true},
		{Wrap(NewRequiredf("Err4"), "Wrap4"), IsRequired, true},
		{NewNotImplemented(Wrap(NewRequiredf("Err5"), "Wrap5"), "NotImplemend5"), IsRequired, true},
		{Wrap(NewRequired(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Required6"), "Wrap6a"), IsRequired, true},
		{Wrap(NewRequired(errors.New("I'm the cause7"), "Required7"), "Wrap7"), IsRequired, true},
		{NewRequiredf("Error8"), IsRequired, true},
		{nil, IsRequired, false},
		{testBehave{}, IsRequired, false},
		{testBehave{ret: true}, IsRequired, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Restricted() bool {
	return nf.ret
}

func TestRestricted(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsRestricted, false},
		{NewRestricted(nil, "Error2"), IsRestricted, false},
		{NewRestricted(Error("Error3a"), "Error3"), IsRestricted, true},
		{Wrap(NewRestrictedf("Err4"), "Wrap4"), IsRestricted, true},
		{NewNotImplemented(Wrap(NewRestrictedf("Err5"), "Wrap5"), "NotImplemend5"), IsRestricted, true},
		{Wrap(NewRestricted(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Restricted6"), "Wrap6a"), IsRestricted, true},
		{Wrap(NewRestricted(errors.New("I'm the cause7"), "Restricted7"), "Wrap7"), IsRestricted, true},
		{NewRestrictedf("Error8"), IsRestricted, true},
		{nil, IsRestricted, false},
		{testBehave{}, IsRestricted, false},
		{testBehave{ret: true}, IsRestricted, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Revoked() bool {
	return nf.ret
}

func TestRevoked(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsRevoked, false},
		{NewRevoked(nil, "Error2"), IsRevoked, false},
		{NewRevoked(Error("Error3a"), "Error3"), IsRevoked, true},
		{Wrap(NewRevokedf("Err4"), "Wrap4"), IsRevoked, true},
		{NewNotImplemented(Wrap(NewRevokedf("Err5"), "Wrap5"), "NotImplemend5"), IsRevoked, true},
		{Wrap(NewRevoked(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Revoked6"), "Wrap6a"), IsRevoked, true},
		{Wrap(NewRevoked(errors.New("I'm the cause7"), "Revoked7"), "Wrap7"), IsRevoked, true},
		{NewRevokedf("Error8"), IsRevoked, true},
		{nil, IsRevoked, false},
		{testBehave{}, IsRevoked, false},
		{testBehave{ret: true}, IsRevoked, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Temporary() bool {
	return nf.ret
}

func TestTemporary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsTemporary, false},
		{NewTemporary(nil, "Error2"), IsTemporary, false},
		{NewTemporary(Error("Error3a"), "Error3"), IsTemporary, true},
		{Wrap(NewTemporaryf("Err4"), "Wrap4"), IsTemporary, true},
		{NewNotImplemented(Wrap(NewTemporaryf("Err5"), "Wrap5"), "NotImplemend5"), IsTemporary, true},
		{Wrap(NewTemporary(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Temporary6"), "Wrap6a"), IsTemporary, true},
		{Wrap(NewTemporary(errors.New("I'm the cause7"), "Temporary7"), "Wrap7"), IsTemporary, true},
		{NewTemporaryf("Error8"), IsTemporary, true},
		{nil, IsTemporary, false},
		{testBehave{}, IsTemporary, false},
		{testBehave{ret: true}, IsTemporary, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Terminated() bool {
	return nf.ret
}

func TestTerminated(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsTerminated, false},
		{NewTerminated(nil, "Error2"), IsTerminated, false},
		{NewTerminated(Error("Error3a"), "Error3"), IsTerminated, true},
		{Wrap(NewTerminatedf("Err4"), "Wrap4"), IsTerminated, true},
		{NewNotImplemented(Wrap(NewTerminatedf("Err5"), "Wrap5"), "NotImplemend5"), IsTerminated, true},
		{Wrap(NewTerminated(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Terminated6"), "Wrap6a"), IsTerminated, true},
		{Wrap(NewTerminated(errors.New("I'm the cause7"), "Terminated7"), "Wrap7"), IsTerminated, true},
		{NewTerminatedf("Error8"), IsTerminated, true},
		{nil, IsTerminated, false},
		{testBehave{}, IsTerminated, false},
		{testBehave{ret: true}, IsTerminated, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Timeout() bool {
	return nf.ret
}

func TestTimeout(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsTimeout, false},
		{NewTimeout(nil, "Error2"), IsTimeout, false},
		{NewTimeout(Error("Error3a"), "Error3"), IsTimeout, true},
		{Wrap(NewTimeoutf("Err4"), "Wrap4"), IsTimeout, true},
		{NewNotImplemented(Wrap(NewTimeoutf("Err5"), "Wrap5"), "NotImplemend5"), IsTimeout, true},
		{Wrap(NewTimeout(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Timeout6"), "Wrap6a"), IsTimeout, true},
		{Wrap(NewTimeout(errors.New("I'm the cause7"), "Timeout7"), "Wrap7"), IsTimeout, true},
		{NewTimeoutf("Error8"), IsTimeout, true},
		{nil, IsTimeout, false},
		{testBehave{}, IsTimeout, false},
		{testBehave{ret: true}, IsTimeout, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) TooLarge() bool {
	return nf.ret
}

func TestTooLarge(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsTooLarge, false},
		{NewTooLarge(nil, "Error2"), IsTooLarge, false},
		{NewTooLarge(Error("Error3a"), "Error3"), IsTooLarge, true},
		{Wrap(NewTooLargef("Err4"), "Wrap4"), IsTooLarge, true},
		{NewNotImplemented(Wrap(NewTooLargef("Err5"), "Wrap5"), "NotImplemend5"), IsTooLarge, true},
		{Wrap(NewTooLarge(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "TooLarge6"), "Wrap6a"), IsTooLarge, true},
		{Wrap(NewTooLarge(errors.New("I'm the cause7"), "TooLarge7"), "Wrap7"), IsTooLarge, true},
		{NewTooLargef("Error8"), IsTooLarge, true},
		{nil, IsTooLarge, false},
		{testBehave{}, IsTooLarge, false},
		{testBehave{ret: true}, IsTooLarge, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Unauthorized() bool {
	return nf.ret
}

func TestUnauthorized(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsUnauthorized, false},
		{NewUnauthorized(nil, "Error2"), IsUnauthorized, false},
		{NewUnauthorized(Error("Error3a"), "Error3"), IsUnauthorized, true},
		{Wrap(NewUnauthorizedf("Err4"), "Wrap4"), IsUnauthorized, true},
		{NewNotImplemented(Wrap(NewUnauthorizedf("Err5"), "Wrap5"), "NotImplemend5"), IsUnauthorized, true},
		{Wrap(NewUnauthorized(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Unauthorized6"), "Wrap6a"), IsUnauthorized, true},
		{Wrap(NewUnauthorized(errors.New("I'm the cause7"), "Unauthorized7"), "Wrap7"), IsUnauthorized, true},
		{NewUnauthorizedf("Error8"), IsUnauthorized, true},
		{nil, IsUnauthorized, false},
		{testBehave{}, IsUnauthorized, false},
		{testBehave{ret: true}, IsUnauthorized, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) Unavailable() bool {
	return nf.ret
}

func TestUnavailable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsUnavailable, false},
		{NewUnavailable(nil, "Error2"), IsUnavailable, false},
		{NewUnavailable(Error("Error3a"), "Error3"), IsUnavailable, true},
		{Wrap(NewUnavailablef("Err4"), "Wrap4"), IsUnavailable, true},
		{NewNotImplemented(Wrap(NewUnavailablef("Err5"), "Wrap5"), "NotImplemend5"), IsUnavailable, true},
		{Wrap(NewUnavailable(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "Unavailable6"), "Wrap6a"), IsUnavailable, true},
		{Wrap(NewUnavailable(errors.New("I'm the cause7"), "Unavailable7"), "Wrap7"), IsUnavailable, true},
		{NewUnavailablef("Error8"), IsUnavailable, true},
		{nil, IsUnavailable, false},
		{testBehave{}, IsUnavailable, false},
		{testBehave{ret: true}, IsUnavailable, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) UserNotFound() bool {
	return nf.ret
}

func TestUserNotFound(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsUserNotFound, false},
		{NewUserNotFound(nil, "Error2"), IsUserNotFound, false},
		{NewUserNotFound(Error("Error3a"), "Error3"), IsUserNotFound, true},
		{Wrap(NewUserNotFoundf("Err4"), "Wrap4"), IsUserNotFound, true},
		{NewNotImplemented(Wrap(NewUserNotFoundf("Err5"), "Wrap5"), "NotImplemend5"), IsUserNotFound, true},
		{Wrap(NewUserNotFound(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "UserNotFound6"), "Wrap6a"), IsUserNotFound, true},
		{Wrap(NewUserNotFound(errors.New("I'm the cause7"), "UserNotFound7"), "Wrap7"), IsUserNotFound, true},
		{NewUserNotFoundf("Error8"), IsUserNotFound, true},
		{nil, IsUserNotFound, false},
		{testBehave{}, IsUserNotFound, false},
		{testBehave{ret: true}, IsUserNotFound, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) VerificationFailed() bool {
	return nf.ret
}

func TestVerificationFailed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsVerificationFailed, false},
		{NewVerificationFailed(nil, "Error2"), IsVerificationFailed, false},
		{NewVerificationFailed(Error("Error3a"), "Error3"), IsVerificationFailed, true},
		{Wrap(NewVerificationFailedf("Err4"), "Wrap4"), IsVerificationFailed, true},
		{NewNotImplemented(Wrap(NewVerificationFailedf("Err5"), "Wrap5"), "NotImplemend5"), IsVerificationFailed, true},
		{Wrap(NewVerificationFailed(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "VerificationFailed6"), "Wrap6a"), IsVerificationFailed, true},
		{Wrap(NewVerificationFailed(errors.New("I'm the cause7"), "VerificationFailed7"), "Wrap7"), IsVerificationFailed, true},
		{NewVerificationFailedf("Error8"), IsVerificationFailed, true},
		{nil, IsVerificationFailed, false},
		{testBehave{}, IsVerificationFailed, false},
		{testBehave{ret: true}, IsVerificationFailed, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) WriteFailed() bool {
	return nf.ret
}

func TestWriteFailed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsWriteFailed, false},
		{NewWriteFailed(nil, "Error2"), IsWriteFailed, false},
		{NewWriteFailed(Error("Error3a"), "Error3"), IsWriteFailed, true},
		{Wrap(NewWriteFailedf("Err4"), "Wrap4"), IsWriteFailed, true},
		{NewNotImplemented(Wrap(NewWriteFailedf("Err5"), "Wrap5"), "NotImplemend5"), IsWriteFailed, true},
		{Wrap(NewWriteFailed(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "WriteFailed6"), "Wrap6a"), IsWriteFailed, true},
		{Wrap(NewWriteFailed(errors.New("I'm the cause7"), "WriteFailed7"), "Wrap7"), IsWriteFailed, true},
		{NewWriteFailedf("Error8"), IsWriteFailed, true},
		{nil, IsWriteFailed, false},
		{testBehave{}, IsWriteFailed, false},
		{testBehave{ret: true}, IsWriteFailed, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}

func (nf testBehave) WrongVersion() bool {
	return nf.ret
}

func TestWrongVersion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{errors.New("Error1"), IsWrongVersion, false},
		{NewWrongVersion(nil, "Error2"), IsWrongVersion, false},
		{NewWrongVersion(Error("Error3a"), "Error3"), IsWrongVersion, true},
		{Wrap(NewWrongVersionf("Err4"), "Wrap4"), IsWrongVersion, true},
		{NewNotImplemented(Wrap(NewWrongVersionf("Err5"), "Wrap5"), "NotImplemend5"), IsWrongVersion, true},
		{Wrap(NewWrongVersion(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "WrongVersion6"), "Wrap6a"), IsWrongVersion, true},
		{Wrap(NewWrongVersion(errors.New("I'm the cause7"), "WrongVersion7"), "Wrap7"), IsWrongVersion, true},
		{NewWrongVersionf("Error8"), IsWrongVersion, true},
		{nil, IsWrongVersion, false},
		{testBehave{}, IsWrongVersion, false},
		{testBehave{ret: true}, IsWrongVersion, true},
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}
