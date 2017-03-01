// Auto generated via behaviour_gen.go

package errors

type (
	aborted  struct{ wrapper }
	abortedf struct{ _error }
)

// NewAborted returns an error which wraps err that satisfies
// IsAborted().
func NewAborted(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &aborted{errWrapf(err, msg, args...)}
}

// NewAbortedf returns a formatted error that satisfies IsAborted().
func NewAbortedf(format string, args ...interface{}) error {
	return &abortedf{errNewf(format, args...)}
}

func isAborted(err error) (ok bool) {
	type iFace interface {
		Aborted() bool
	}
	switch et := err.(type) {
	case *aborted:
		ok = true
	case *abortedf:
		ok = true
	case iFace:
		ok = et.Aborted()
	}
	return
}

// IsAborted reports whether err was created with NewAborted() or
// implements interface:
//     type Aborteder interface {
//            Aborted() bool
//     }
func IsAborted(err error) bool {
	return CausedBehaviour(err, isAborted)
}

type (
	alreadyClosed  struct{ wrapper }
	alreadyClosedf struct{ _error }
)

// NewAlreadyClosed returns an error which wraps err that satisfies
// IsAlreadyClosed().
func NewAlreadyClosed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &alreadyClosed{errWrapf(err, msg, args...)}
}

// NewAlreadyClosedf returns a formatted error that satisfies IsAlreadyClosed().
func NewAlreadyClosedf(format string, args ...interface{}) error {
	return &alreadyClosedf{errNewf(format, args...)}
}

func isAlreadyClosed(err error) (ok bool) {
	type iFace interface {
		AlreadyClosed() bool
	}
	switch et := err.(type) {
	case *alreadyClosed:
		ok = true
	case *alreadyClosedf:
		ok = true
	case iFace:
		ok = et.AlreadyClosed()
	}
	return
}

// IsAlreadyClosed reports whether err was created with NewAlreadyClosed() or
// implements interface:
//     type AlreadyCloseder interface {
//            AlreadyClosed() bool
//     }
func IsAlreadyClosed(err error) bool {
	return CausedBehaviour(err, isAlreadyClosed)
}

type (
	alreadyExists  struct{ wrapper }
	alreadyExistsf struct{ _error }
)

// NewAlreadyExists returns an error which wraps err that satisfies
// IsAlreadyExists().
func NewAlreadyExists(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &alreadyExists{errWrapf(err, msg, args...)}
}

// NewAlreadyExistsf returns a formatted error that satisfies IsAlreadyExists().
func NewAlreadyExistsf(format string, args ...interface{}) error {
	return &alreadyExistsf{errNewf(format, args...)}
}

func isAlreadyExists(err error) (ok bool) {
	type iFace interface {
		AlreadyExists() bool
	}
	switch et := err.(type) {
	case *alreadyExists:
		ok = true
	case *alreadyExistsf:
		ok = true
	case iFace:
		ok = et.AlreadyExists()
	}
	return
}

// IsAlreadyExists reports whether err was created with NewAlreadyExists() or
// implements interface:
//     type AlreadyExistser interface {
//            AlreadyExists() bool
//     }
func IsAlreadyExists(err error) bool {
	return CausedBehaviour(err, isAlreadyExists)
}

type (
	alreadyInUse  struct{ wrapper }
	alreadyInUsef struct{ _error }
)

// NewAlreadyInUse returns an error which wraps err that satisfies
// IsAlreadyInUse().
func NewAlreadyInUse(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &alreadyInUse{errWrapf(err, msg, args...)}
}

// NewAlreadyInUsef returns a formatted error that satisfies IsAlreadyInUse().
func NewAlreadyInUsef(format string, args ...interface{}) error {
	return &alreadyInUsef{errNewf(format, args...)}
}

func isAlreadyInUse(err error) (ok bool) {
	type iFace interface {
		AlreadyInUse() bool
	}
	switch et := err.(type) {
	case *alreadyInUse:
		ok = true
	case *alreadyInUsef:
		ok = true
	case iFace:
		ok = et.AlreadyInUse()
	}
	return
}

// IsAlreadyInUse reports whether err was created with NewAlreadyInUse() or
// implements interface:
//     type AlreadyInUseer interface {
//            AlreadyInUse() bool
//     }
func IsAlreadyInUse(err error) bool {
	return CausedBehaviour(err, isAlreadyInUse)
}

type (
	connectionFailed  struct{ wrapper }
	connectionFailedf struct{ _error }
)

// NewConnectionFailed returns an error which wraps err that satisfies
// IsConnectionFailed().
func NewConnectionFailed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &connectionFailed{errWrapf(err, msg, args...)}
}

// NewConnectionFailedf returns a formatted error that satisfies IsConnectionFailed().
func NewConnectionFailedf(format string, args ...interface{}) error {
	return &connectionFailedf{errNewf(format, args...)}
}

func isConnectionFailed(err error) (ok bool) {
	type iFace interface {
		ConnectionFailed() bool
	}
	switch et := err.(type) {
	case *connectionFailed:
		ok = true
	case *connectionFailedf:
		ok = true
	case iFace:
		ok = et.ConnectionFailed()
	}
	return
}

// IsConnectionFailed reports whether err was created with NewConnectionFailed() or
// implements interface:
//     type ConnectionFaileder interface {
//            ConnectionFailed() bool
//     }
func IsConnectionFailed(err error) bool {
	return CausedBehaviour(err, isConnectionFailed)
}

type (
	empty  struct{ wrapper }
	emptyf struct{ _error }
)

// NewEmpty returns an error which wraps err that satisfies
// IsEmpty().
func NewEmpty(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &empty{errWrapf(err, msg, args...)}
}

// NewEmptyf returns a formatted error that satisfies IsEmpty().
func NewEmptyf(format string, args ...interface{}) error {
	return &emptyf{errNewf(format, args...)}
}

func isEmpty(err error) (ok bool) {
	type iFace interface {
		Empty() bool
	}
	switch et := err.(type) {
	case *empty:
		ok = true
	case *emptyf:
		ok = true
	case iFace:
		ok = et.Empty()
	}
	return
}

// IsEmpty reports whether err was created with NewEmpty() or
// implements interface:
//     type Emptyer interface {
//            Empty() bool
//     }
func IsEmpty(err error) bool {
	return CausedBehaviour(err, isEmpty)
}

type (
	expired  struct{ wrapper }
	expiredf struct{ _error }
)

// NewExpired returns an error which wraps err that satisfies
// IsExpired().
func NewExpired(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &expired{errWrapf(err, msg, args...)}
}

// NewExpiredf returns a formatted error that satisfies IsExpired().
func NewExpiredf(format string, args ...interface{}) error {
	return &expiredf{errNewf(format, args...)}
}

func isExpired(err error) (ok bool) {
	type iFace interface {
		Expired() bool
	}
	switch et := err.(type) {
	case *expired:
		ok = true
	case *expiredf:
		ok = true
	case iFace:
		ok = et.Expired()
	}
	return
}

// IsExpired reports whether err was created with NewExpired() or
// implements interface:
//     type Expireder interface {
//            Expired() bool
//     }
func IsExpired(err error) bool {
	return CausedBehaviour(err, isExpired)
}

type (
	fatal  struct{ wrapper }
	fatalf struct{ _error }
)

// NewFatal returns an error which wraps err that satisfies
// IsFatal().
func NewFatal(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &fatal{errWrapf(err, msg, args...)}
}

// NewFatalf returns a formatted error that satisfies IsFatal().
func NewFatalf(format string, args ...interface{}) error {
	return &fatalf{errNewf(format, args...)}
}

func isFatal(err error) (ok bool) {
	type iFace interface {
		Fatal() bool
	}
	switch et := err.(type) {
	case *fatal:
		ok = true
	case *fatalf:
		ok = true
	case iFace:
		ok = et.Fatal()
	}
	return
}

// IsFatal reports whether err was created with NewFatal() or
// implements interface:
//     type Fataler interface {
//            Fatal() bool
//     }
func IsFatal(err error) bool {
	return CausedBehaviour(err, isFatal)
}

type (
	inProgress  struct{ wrapper }
	inProgressf struct{ _error }
)

// NewInProgress returns an error which wraps err that satisfies
// IsInProgress().
func NewInProgress(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &inProgress{errWrapf(err, msg, args...)}
}

// NewInProgressf returns a formatted error that satisfies IsInProgress().
func NewInProgressf(format string, args ...interface{}) error {
	return &inProgressf{errNewf(format, args...)}
}

func isInProgress(err error) (ok bool) {
	type iFace interface {
		InProgress() bool
	}
	switch et := err.(type) {
	case *inProgress:
		ok = true
	case *inProgressf:
		ok = true
	case iFace:
		ok = et.InProgress()
	}
	return
}

// IsInProgress reports whether err was created with NewInProgress() or
// implements interface:
//     type InProgresser interface {
//            InProgress() bool
//     }
func IsInProgress(err error) bool {
	return CausedBehaviour(err, isInProgress)
}

type (
	interrupted  struct{ wrapper }
	interruptedf struct{ _error }
)

// NewInterrupted returns an error which wraps err that satisfies
// IsInterrupted().
func NewInterrupted(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &interrupted{errWrapf(err, msg, args...)}
}

// NewInterruptedf returns a formatted error that satisfies IsInterrupted().
func NewInterruptedf(format string, args ...interface{}) error {
	return &interruptedf{errNewf(format, args...)}
}

func isInterrupted(err error) (ok bool) {
	type iFace interface {
		Interrupted() bool
	}
	switch et := err.(type) {
	case *interrupted:
		ok = true
	case *interruptedf:
		ok = true
	case iFace:
		ok = et.Interrupted()
	}
	return
}

// IsInterrupted reports whether err was created with NewInterrupted() or
// implements interface:
//     type Interrupteder interface {
//            Interrupted() bool
//     }
func IsInterrupted(err error) bool {
	return CausedBehaviour(err, isInterrupted)
}

type (
	locked  struct{ wrapper }
	lockedf struct{ _error }
)

// NewLocked returns an error which wraps err that satisfies
// IsLocked().
func NewLocked(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &locked{errWrapf(err, msg, args...)}
}

// NewLockedf returns a formatted error that satisfies IsLocked().
func NewLockedf(format string, args ...interface{}) error {
	return &lockedf{errNewf(format, args...)}
}

func isLocked(err error) (ok bool) {
	type iFace interface {
		Locked() bool
	}
	switch et := err.(type) {
	case *locked:
		ok = true
	case *lockedf:
		ok = true
	case iFace:
		ok = et.Locked()
	}
	return
}

// IsLocked reports whether err was created with NewLocked() or
// implements interface:
//     type Lockeder interface {
//            Locked() bool
//     }
func IsLocked(err error) bool {
	return CausedBehaviour(err, isLocked)
}

type (
	methodNotAllowed  struct{ wrapper }
	methodNotAllowedf struct{ _error }
)

// NewMethodNotAllowed returns an error which wraps err that satisfies
// IsMethodNotAllowed().
func NewMethodNotAllowed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &methodNotAllowed{errWrapf(err, msg, args...)}
}

// NewMethodNotAllowedf returns a formatted error that satisfies IsMethodNotAllowed().
func NewMethodNotAllowedf(format string, args ...interface{}) error {
	return &methodNotAllowedf{errNewf(format, args...)}
}

func isMethodNotAllowed(err error) (ok bool) {
	type iFace interface {
		MethodNotAllowed() bool
	}
	switch et := err.(type) {
	case *methodNotAllowed:
		ok = true
	case *methodNotAllowedf:
		ok = true
	case iFace:
		ok = et.MethodNotAllowed()
	}
	return
}

// IsMethodNotAllowed reports whether err was created with NewMethodNotAllowed() or
// implements interface:
//     type MethodNotAlloweder interface {
//            MethodNotAllowed() bool
//     }
func IsMethodNotAllowed(err error) bool {
	return CausedBehaviour(err, isMethodNotAllowed)
}

type (
	notFound  struct{ wrapper }
	notFoundf struct{ _error }
)

// NewNotFound returns an error which wraps err that satisfies
// IsNotFound().
func NewNotFound(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &notFound{errWrapf(err, msg, args...)}
}

// NewNotFoundf returns a formatted error that satisfies IsNotFound().
func NewNotFoundf(format string, args ...interface{}) error {
	return &notFoundf{errNewf(format, args...)}
}

func isNotFound(err error) (ok bool) {
	type iFace interface {
		NotFound() bool
	}
	switch et := err.(type) {
	case *notFound:
		ok = true
	case *notFoundf:
		ok = true
	case iFace:
		ok = et.NotFound()
	}
	return
}

// IsNotFound reports whether err was created with NewNotFound() or
// implements interface:
//     type NotFounder interface {
//            NotFound() bool
//     }
func IsNotFound(err error) bool {
	return CausedBehaviour(err, isNotFound)
}

type (
	notImplemented  struct{ wrapper }
	notImplementedf struct{ _error }
)

// NewNotImplemented returns an error which wraps err that satisfies
// IsNotImplemented().
func NewNotImplemented(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &notImplemented{errWrapf(err, msg, args...)}
}

// NewNotImplementedf returns a formatted error that satisfies IsNotImplemented().
func NewNotImplementedf(format string, args ...interface{}) error {
	return &notImplementedf{errNewf(format, args...)}
}

func isNotImplemented(err error) (ok bool) {
	type iFace interface {
		NotImplemented() bool
	}
	switch et := err.(type) {
	case *notImplemented:
		ok = true
	case *notImplementedf:
		ok = true
	case iFace:
		ok = et.NotImplemented()
	}
	return
}

// IsNotImplemented reports whether err was created with NewNotImplemented() or
// implements interface:
//     type NotImplementeder interface {
//            NotImplemented() bool
//     }
func IsNotImplemented(err error) bool {
	return CausedBehaviour(err, isNotImplemented)
}

type (
	notRecoverable  struct{ wrapper }
	notRecoverablef struct{ _error }
)

// NewNotRecoverable returns an error which wraps err that satisfies
// IsNotRecoverable().
func NewNotRecoverable(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &notRecoverable{errWrapf(err, msg, args...)}
}

// NewNotRecoverablef returns a formatted error that satisfies IsNotRecoverable().
func NewNotRecoverablef(format string, args ...interface{}) error {
	return &notRecoverablef{errNewf(format, args...)}
}

func isNotRecoverable(err error) (ok bool) {
	type iFace interface {
		NotRecoverable() bool
	}
	switch et := err.(type) {
	case *notRecoverable:
		ok = true
	case *notRecoverablef:
		ok = true
	case iFace:
		ok = et.NotRecoverable()
	}
	return
}

// IsNotRecoverable reports whether err was created with NewNotRecoverable() or
// implements interface:
//     type NotRecoverableer interface {
//            NotRecoverable() bool
//     }
func IsNotRecoverable(err error) bool {
	return CausedBehaviour(err, isNotRecoverable)
}

type (
	notSupported  struct{ wrapper }
	notSupportedf struct{ _error }
)

// NewNotSupported returns an error which wraps err that satisfies
// IsNotSupported().
func NewNotSupported(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &notSupported{errWrapf(err, msg, args...)}
}

// NewNotSupportedf returns a formatted error that satisfies IsNotSupported().
func NewNotSupportedf(format string, args ...interface{}) error {
	return &notSupportedf{errNewf(format, args...)}
}

func isNotSupported(err error) (ok bool) {
	type iFace interface {
		NotSupported() bool
	}
	switch et := err.(type) {
	case *notSupported:
		ok = true
	case *notSupportedf:
		ok = true
	case iFace:
		ok = et.NotSupported()
	}
	return
}

// IsNotSupported reports whether err was created with NewNotSupported() or
// implements interface:
//     type NotSupporteder interface {
//            NotSupported() bool
//     }
func IsNotSupported(err error) bool {
	return CausedBehaviour(err, isNotSupported)
}

type (
	notValid  struct{ wrapper }
	notValidf struct{ _error }
)

// NewNotValid returns an error which wraps err that satisfies
// IsNotValid().
func NewNotValid(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &notValid{errWrapf(err, msg, args...)}
}

// NewNotValidf returns a formatted error that satisfies IsNotValid().
func NewNotValidf(format string, args ...interface{}) error {
	return &notValidf{errNewf(format, args...)}
}

func isNotValid(err error) (ok bool) {
	type iFace interface {
		NotValid() bool
	}
	switch et := err.(type) {
	case *notValid:
		ok = true
	case *notValidf:
		ok = true
	case iFace:
		ok = et.NotValid()
	}
	return
}

// IsNotValid reports whether err was created with NewNotValid() or
// implements interface:
//     type NotValider interface {
//            NotValid() bool
//     }
func IsNotValid(err error) bool {
	return CausedBehaviour(err, isNotValid)
}

type (
	permissionDenied  struct{ wrapper }
	permissionDeniedf struct{ _error }
)

// NewPermissionDenied returns an error which wraps err that satisfies
// IsPermissionDenied().
func NewPermissionDenied(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &permissionDenied{errWrapf(err, msg, args...)}
}

// NewPermissionDeniedf returns a formatted error that satisfies IsPermissionDenied().
func NewPermissionDeniedf(format string, args ...interface{}) error {
	return &permissionDeniedf{errNewf(format, args...)}
}

func isPermissionDenied(err error) (ok bool) {
	type iFace interface {
		PermissionDenied() bool
	}
	switch et := err.(type) {
	case *permissionDenied:
		ok = true
	case *permissionDeniedf:
		ok = true
	case iFace:
		ok = et.PermissionDenied()
	}
	return
}

// IsPermissionDenied reports whether err was created with NewPermissionDenied() or
// implements interface:
//     type PermissionDenieder interface {
//            PermissionDenied() bool
//     }
func IsPermissionDenied(err error) bool {
	return CausedBehaviour(err, isPermissionDenied)
}

type (
	previousOwnerDied  struct{ wrapper }
	previousOwnerDiedf struct{ _error }
)

// NewPreviousOwnerDied returns an error which wraps err that satisfies
// IsPreviousOwnerDied().
func NewPreviousOwnerDied(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &previousOwnerDied{errWrapf(err, msg, args...)}
}

// NewPreviousOwnerDiedf returns a formatted error that satisfies IsPreviousOwnerDied().
func NewPreviousOwnerDiedf(format string, args ...interface{}) error {
	return &previousOwnerDiedf{errNewf(format, args...)}
}

func isPreviousOwnerDied(err error) (ok bool) {
	type iFace interface {
		PreviousOwnerDied() bool
	}
	switch et := err.(type) {
	case *previousOwnerDied:
		ok = true
	case *previousOwnerDiedf:
		ok = true
	case iFace:
		ok = et.PreviousOwnerDied()
	}
	return
}

// IsPreviousOwnerDied reports whether err was created with NewPreviousOwnerDied() or
// implements interface:
//     type PreviousOwnerDieder interface {
//            PreviousOwnerDied() bool
//     }
func IsPreviousOwnerDied(err error) bool {
	return CausedBehaviour(err, isPreviousOwnerDied)
}

type (
	quotaExceeded  struct{ wrapper }
	quotaExceededf struct{ _error }
)

// NewQuotaExceeded returns an error which wraps err that satisfies
// IsQuotaExceeded().
func NewQuotaExceeded(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &quotaExceeded{errWrapf(err, msg, args...)}
}

// NewQuotaExceededf returns a formatted error that satisfies IsQuotaExceeded().
func NewQuotaExceededf(format string, args ...interface{}) error {
	return &quotaExceededf{errNewf(format, args...)}
}

func isQuotaExceeded(err error) (ok bool) {
	type iFace interface {
		QuotaExceeded() bool
	}
	switch et := err.(type) {
	case *quotaExceeded:
		ok = true
	case *quotaExceededf:
		ok = true
	case iFace:
		ok = et.QuotaExceeded()
	}
	return
}

// IsQuotaExceeded reports whether err was created with NewQuotaExceeded() or
// implements interface:
//     type QuotaExceededer interface {
//            QuotaExceeded() bool
//     }
func IsQuotaExceeded(err error) bool {
	return CausedBehaviour(err, isQuotaExceeded)
}

type (
	readFailed  struct{ wrapper }
	readFailedf struct{ _error }
)

// NewReadFailed returns an error which wraps err that satisfies
// IsReadFailed().
func NewReadFailed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &readFailed{errWrapf(err, msg, args...)}
}

// NewReadFailedf returns a formatted error that satisfies IsReadFailed().
func NewReadFailedf(format string, args ...interface{}) error {
	return &readFailedf{errNewf(format, args...)}
}

func isReadFailed(err error) (ok bool) {
	type iFace interface {
		ReadFailed() bool
	}
	switch et := err.(type) {
	case *readFailed:
		ok = true
	case *readFailedf:
		ok = true
	case iFace:
		ok = et.ReadFailed()
	}
	return
}

// IsReadFailed reports whether err was created with NewReadFailed() or
// implements interface:
//     type ReadFaileder interface {
//            ReadFailed() bool
//     }
func IsReadFailed(err error) bool {
	return CausedBehaviour(err, isReadFailed)
}

type (
	rejected  struct{ wrapper }
	rejectedf struct{ _error }
)

// NewRejected returns an error which wraps err that satisfies
// IsRejected().
func NewRejected(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &rejected{errWrapf(err, msg, args...)}
}

// NewRejectedf returns a formatted error that satisfies IsRejected().
func NewRejectedf(format string, args ...interface{}) error {
	return &rejectedf{errNewf(format, args...)}
}

func isRejected(err error) (ok bool) {
	type iFace interface {
		Rejected() bool
	}
	switch et := err.(type) {
	case *rejected:
		ok = true
	case *rejectedf:
		ok = true
	case iFace:
		ok = et.Rejected()
	}
	return
}

// IsRejected reports whether err was created with NewRejected() or
// implements interface:
//     type Rejecteder interface {
//            Rejected() bool
//     }
func IsRejected(err error) bool {
	return CausedBehaviour(err, isRejected)
}

type (
	revoked  struct{ wrapper }
	revokedf struct{ _error }
)

// NewRevoked returns an error which wraps err that satisfies
// IsRevoked().
func NewRevoked(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &revoked{errWrapf(err, msg, args...)}
}

// NewRevokedf returns a formatted error that satisfies IsRevoked().
func NewRevokedf(format string, args ...interface{}) error {
	return &revokedf{errNewf(format, args...)}
}

func isRevoked(err error) (ok bool) {
	type iFace interface {
		Revoked() bool
	}
	switch et := err.(type) {
	case *revoked:
		ok = true
	case *revokedf:
		ok = true
	case iFace:
		ok = et.Revoked()
	}
	return
}

// IsRevoked reports whether err was created with NewRevoked() or
// implements interface:
//     type Revokeder interface {
//            Revoked() bool
//     }
func IsRevoked(err error) bool {
	return CausedBehaviour(err, isRevoked)
}

type (
	temporary  struct{ wrapper }
	temporaryf struct{ _error }
)

// NewTemporary returns an error which wraps err that satisfies
// IsTemporary().
func NewTemporary(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &temporary{errWrapf(err, msg, args...)}
}

// NewTemporaryf returns a formatted error that satisfies IsTemporary().
func NewTemporaryf(format string, args ...interface{}) error {
	return &temporaryf{errNewf(format, args...)}
}

func isTemporary(err error) (ok bool) {
	type iFace interface {
		Temporary() bool
	}
	switch et := err.(type) {
	case *temporary:
		ok = true
	case *temporaryf:
		ok = true
	case iFace:
		ok = et.Temporary()
	}
	return
}

// IsTemporary reports whether err was created with NewTemporary() or
// implements interface:
//     type Temporaryer interface {
//            Temporary() bool
//     }
func IsTemporary(err error) bool {
	return CausedBehaviour(err, isTemporary)
}

type (
	terminated  struct{ wrapper }
	terminatedf struct{ _error }
)

// NewTerminated returns an error which wraps err that satisfies
// IsTerminated().
func NewTerminated(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &terminated{errWrapf(err, msg, args...)}
}

// NewTerminatedf returns a formatted error that satisfies IsTerminated().
func NewTerminatedf(format string, args ...interface{}) error {
	return &terminatedf{errNewf(format, args...)}
}

func isTerminated(err error) (ok bool) {
	type iFace interface {
		Terminated() bool
	}
	switch et := err.(type) {
	case *terminated:
		ok = true
	case *terminatedf:
		ok = true
	case iFace:
		ok = et.Terminated()
	}
	return
}

// IsTerminated reports whether err was created with NewTerminated() or
// implements interface:
//     type Terminateder interface {
//            Terminated() bool
//     }
func IsTerminated(err error) bool {
	return CausedBehaviour(err, isTerminated)
}

type (
	timeout  struct{ wrapper }
	timeoutf struct{ _error }
)

// NewTimeout returns an error which wraps err that satisfies
// IsTimeout().
func NewTimeout(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &timeout{errWrapf(err, msg, args...)}
}

// NewTimeoutf returns a formatted error that satisfies IsTimeout().
func NewTimeoutf(format string, args ...interface{}) error {
	return &timeoutf{errNewf(format, args...)}
}

func isTimeout(err error) (ok bool) {
	type iFace interface {
		Timeout() bool
	}
	switch et := err.(type) {
	case *timeout:
		ok = true
	case *timeoutf:
		ok = true
	case iFace:
		ok = et.Timeout()
	}
	return
}

// IsTimeout reports whether err was created with NewTimeout() or
// implements interface:
//     type Timeouter interface {
//            Timeout() bool
//     }
func IsTimeout(err error) bool {
	return CausedBehaviour(err, isTimeout)
}

type (
	tooLarge  struct{ wrapper }
	tooLargef struct{ _error }
)

// NewTooLarge returns an error which wraps err that satisfies
// IsTooLarge().
func NewTooLarge(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &tooLarge{errWrapf(err, msg, args...)}
}

// NewTooLargef returns a formatted error that satisfies IsTooLarge().
func NewTooLargef(format string, args ...interface{}) error {
	return &tooLargef{errNewf(format, args...)}
}

func isTooLarge(err error) (ok bool) {
	type iFace interface {
		TooLarge() bool
	}
	switch et := err.(type) {
	case *tooLarge:
		ok = true
	case *tooLargef:
		ok = true
	case iFace:
		ok = et.TooLarge()
	}
	return
}

// IsTooLarge reports whether err was created with NewTooLarge() or
// implements interface:
//     type TooLargeer interface {
//            TooLarge() bool
//     }
func IsTooLarge(err error) bool {
	return CausedBehaviour(err, isTooLarge)
}

type (
	unauthorized  struct{ wrapper }
	unauthorizedf struct{ _error }
)

// NewUnauthorized returns an error which wraps err that satisfies
// IsUnauthorized().
func NewUnauthorized(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &unauthorized{errWrapf(err, msg, args...)}
}

// NewUnauthorizedf returns a formatted error that satisfies IsUnauthorized().
func NewUnauthorizedf(format string, args ...interface{}) error {
	return &unauthorizedf{errNewf(format, args...)}
}

func isUnauthorized(err error) (ok bool) {
	type iFace interface {
		Unauthorized() bool
	}
	switch et := err.(type) {
	case *unauthorized:
		ok = true
	case *unauthorizedf:
		ok = true
	case iFace:
		ok = et.Unauthorized()
	}
	return
}

// IsUnauthorized reports whether err was created with NewUnauthorized() or
// implements interface:
//     type Unauthorizeder interface {
//            Unauthorized() bool
//     }
func IsUnauthorized(err error) bool {
	return CausedBehaviour(err, isUnauthorized)
}

type (
	userNotFound  struct{ wrapper }
	userNotFoundf struct{ _error }
)

// NewUserNotFound returns an error which wraps err that satisfies
// IsUserNotFound().
func NewUserNotFound(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &userNotFound{errWrapf(err, msg, args...)}
}

// NewUserNotFoundf returns a formatted error that satisfies IsUserNotFound().
func NewUserNotFoundf(format string, args ...interface{}) error {
	return &userNotFoundf{errNewf(format, args...)}
}

func isUserNotFound(err error) (ok bool) {
	type iFace interface {
		UserNotFound() bool
	}
	switch et := err.(type) {
	case *userNotFound:
		ok = true
	case *userNotFoundf:
		ok = true
	case iFace:
		ok = et.UserNotFound()
	}
	return
}

// IsUserNotFound reports whether err was created with NewUserNotFound() or
// implements interface:
//     type UserNotFounder interface {
//            UserNotFound() bool
//     }
func IsUserNotFound(err error) bool {
	return CausedBehaviour(err, isUserNotFound)
}

type (
	writeFailed  struct{ wrapper }
	writeFailedf struct{ _error }
)

// NewWriteFailed returns an error which wraps err that satisfies
// IsWriteFailed().
func NewWriteFailed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &writeFailed{errWrapf(err, msg, args...)}
}

// NewWriteFailedf returns a formatted error that satisfies IsWriteFailed().
func NewWriteFailedf(format string, args ...interface{}) error {
	return &writeFailedf{errNewf(format, args...)}
}

func isWriteFailed(err error) (ok bool) {
	type iFace interface {
		WriteFailed() bool
	}
	switch et := err.(type) {
	case *writeFailed:
		ok = true
	case *writeFailedf:
		ok = true
	case iFace:
		ok = et.WriteFailed()
	}
	return
}

// IsWriteFailed reports whether err was created with NewWriteFailed() or
// implements interface:
//     type WriteFaileder interface {
//            WriteFailed() bool
//     }
func IsWriteFailed(err error) bool {
	return CausedBehaviour(err, isWriteFailed)
}

type (
	wrongVersion  struct{ wrapper }
	wrongVersionf struct{ _error }
)

// NewWrongVersion returns an error which wraps err that satisfies
// IsWrongVersion().
func NewWrongVersion(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &wrongVersion{errWrapf(err, msg, args...)}
}

// NewWrongVersionf returns a formatted error that satisfies IsWrongVersion().
func NewWrongVersionf(format string, args ...interface{}) error {
	return &wrongVersionf{errNewf(format, args...)}
}

func isWrongVersion(err error) (ok bool) {
	type iFace interface {
		WrongVersion() bool
	}
	switch et := err.(type) {
	case *wrongVersion:
		ok = true
	case *wrongVersionf:
		ok = true
	case iFace:
		ok = et.WrongVersion()
	}
	return
}

// IsWrongVersion reports whether err was created with NewWrongVersion() or
// implements interface:
//     type WrongVersioner interface {
//            WrongVersion() bool
//     }
func IsWrongVersion(err error) bool {
	return CausedBehaviour(err, isWrongVersion)
}
