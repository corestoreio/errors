// Auto generated via behaviour_gen.go

package errors

type (
	aborted  struct{ *withStack }
	abortedf struct{ *fundamental }
)

// NewAborted returns an error which wraps err that satisfies
// IsAborted().
func NewAborted(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return aborted{errWrapf(err, msg, args...)}
}

// NewAbortedf returns a formatted error that satisfies IsAborted().
func NewAbortedf(format string, args ...interface{}) error {
	return abortedf{errNewf(format, args...)}
}

func isAborted(err error) (ok bool) {
	type iFace interface {
		Aborted() bool
	}
	switch et := err.(type) {
	case aborted:
		ok = true
	case abortedf:
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
	alreadyClosed  struct{ *withStack }
	alreadyClosedf struct{ *fundamental }
)

// NewAlreadyClosed returns an error which wraps err that satisfies
// IsAlreadyClosed().
func NewAlreadyClosed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return alreadyClosed{errWrapf(err, msg, args...)}
}

// NewAlreadyClosedf returns a formatted error that satisfies IsAlreadyClosed().
func NewAlreadyClosedf(format string, args ...interface{}) error {
	return alreadyClosedf{errNewf(format, args...)}
}

func isAlreadyClosed(err error) (ok bool) {
	type iFace interface {
		AlreadyClosed() bool
	}
	switch et := err.(type) {
	case alreadyClosed:
		ok = true
	case alreadyClosedf:
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
	alreadyExists  struct{ *withStack }
	alreadyExistsf struct{ *fundamental }
)

// NewAlreadyExists returns an error which wraps err that satisfies
// IsAlreadyExists().
func NewAlreadyExists(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return alreadyExists{errWrapf(err, msg, args...)}
}

// NewAlreadyExistsf returns a formatted error that satisfies IsAlreadyExists().
func NewAlreadyExistsf(format string, args ...interface{}) error {
	return alreadyExistsf{errNewf(format, args...)}
}

func isAlreadyExists(err error) (ok bool) {
	type iFace interface {
		AlreadyExists() bool
	}
	switch et := err.(type) {
	case alreadyExists:
		ok = true
	case alreadyExistsf:
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
	alreadyInUse  struct{ *withStack }
	alreadyInUsef struct{ *fundamental }
)

// NewAlreadyInUse returns an error which wraps err that satisfies
// IsAlreadyInUse().
func NewAlreadyInUse(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return alreadyInUse{errWrapf(err, msg, args...)}
}

// NewAlreadyInUsef returns a formatted error that satisfies IsAlreadyInUse().
func NewAlreadyInUsef(format string, args ...interface{}) error {
	return alreadyInUsef{errNewf(format, args...)}
}

func isAlreadyInUse(err error) (ok bool) {
	type iFace interface {
		AlreadyInUse() bool
	}
	switch et := err.(type) {
	case alreadyInUse:
		ok = true
	case alreadyInUsef:
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
	alreadyCaptured  struct{ *withStack }
	alreadyCapturedf struct{ *fundamental }
)

// NewAlreadyCaptured returns an error which wraps err that satisfies
// IsAlreadyCaptured().
func NewAlreadyCaptured(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return alreadyCaptured{errWrapf(err, msg, args...)}
}

// NewAlreadyCapturedf returns a formatted error that satisfies IsAlreadyCaptured().
func NewAlreadyCapturedf(format string, args ...interface{}) error {
	return alreadyCapturedf{errNewf(format, args...)}
}

func isAlreadyCaptured(err error) (ok bool) {
	type iFace interface {
		AlreadyCaptured() bool
	}
	switch et := err.(type) {
	case alreadyCaptured:
		ok = true
	case alreadyCapturedf:
		ok = true
	case iFace:
		ok = et.AlreadyCaptured()
	}
	return
}

// IsAlreadyCaptured reports whether err was created with NewAlreadyCaptured() or
// implements interface:
//     type AlreadyCaptureder interface {
//            AlreadyCaptured() bool
//     }
func IsAlreadyCaptured(err error) bool {
	return CausedBehaviour(err, isAlreadyCaptured)
}

type (
	alreadyRefunded  struct{ *withStack }
	alreadyRefundedf struct{ *fundamental }
)

// NewAlreadyRefunded returns an error which wraps err that satisfies
// IsAlreadyRefunded().
func NewAlreadyRefunded(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return alreadyRefunded{errWrapf(err, msg, args...)}
}

// NewAlreadyRefundedf returns a formatted error that satisfies IsAlreadyRefunded().
func NewAlreadyRefundedf(format string, args ...interface{}) error {
	return alreadyRefundedf{errNewf(format, args...)}
}

func isAlreadyRefunded(err error) (ok bool) {
	type iFace interface {
		AlreadyRefunded() bool
	}
	switch et := err.(type) {
	case alreadyRefunded:
		ok = true
	case alreadyRefundedf:
		ok = true
	case iFace:
		ok = et.AlreadyRefunded()
	}
	return
}

// IsAlreadyRefunded reports whether err was created with NewAlreadyRefunded() or
// implements interface:
//     type AlreadyRefundeder interface {
//            AlreadyRefunded() bool
//     }
func IsAlreadyRefunded(err error) bool {
	return CausedBehaviour(err, isAlreadyRefunded)
}

type (
	blocked  struct{ *withStack }
	blockedf struct{ *fundamental }
)

// NewBlocked returns an error which wraps err that satisfies
// IsBlocked().
func NewBlocked(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return blocked{errWrapf(err, msg, args...)}
}

// NewBlockedf returns a formatted error that satisfies IsBlocked().
func NewBlockedf(format string, args ...interface{}) error {
	return blockedf{errNewf(format, args...)}
}

func isBlocked(err error) (ok bool) {
	type iFace interface {
		Blocked() bool
	}
	switch et := err.(type) {
	case blocked:
		ok = true
	case blockedf:
		ok = true
	case iFace:
		ok = et.Blocked()
	}
	return
}

// IsBlocked reports whether err was created with NewBlocked() or
// implements interface:
//     type Blockeder interface {
//            Blocked() bool
//     }
func IsBlocked(err error) bool {
	return CausedBehaviour(err, isBlocked)
}

type (
	connectionFailed  struct{ *withStack }
	connectionFailedf struct{ *fundamental }
)

// NewConnectionFailed returns an error which wraps err that satisfies
// IsConnectionFailed().
func NewConnectionFailed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return connectionFailed{errWrapf(err, msg, args...)}
}

// NewConnectionFailedf returns a formatted error that satisfies IsConnectionFailed().
func NewConnectionFailedf(format string, args ...interface{}) error {
	return connectionFailedf{errNewf(format, args...)}
}

func isConnectionFailed(err error) (ok bool) {
	type iFace interface {
		ConnectionFailed() bool
	}
	switch et := err.(type) {
	case connectionFailed:
		ok = true
	case connectionFailedf:
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
	declined  struct{ *withStack }
	declinedf struct{ *fundamental }
)

// NewDeclined returns an error which wraps err that satisfies
// IsDeclined().
func NewDeclined(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return declined{errWrapf(err, msg, args...)}
}

// NewDeclinedf returns a formatted error that satisfies IsDeclined().
func NewDeclinedf(format string, args ...interface{}) error {
	return declinedf{errNewf(format, args...)}
}

func isDeclined(err error) (ok bool) {
	type iFace interface {
		Declined() bool
	}
	switch et := err.(type) {
	case declined:
		ok = true
	case declinedf:
		ok = true
	case iFace:
		ok = et.Declined()
	}
	return
}

// IsDeclined reports whether err was created with NewDeclined() or
// implements interface:
//     type Declineder interface {
//            Declined() bool
//     }
func IsDeclined(err error) bool {
	return CausedBehaviour(err, isDeclined)
}

type (
	denied  struct{ *withStack }
	deniedf struct{ *fundamental }
)

// NewDenied returns an error which wraps err that satisfies
// IsDenied().
func NewDenied(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return denied{errWrapf(err, msg, args...)}
}

// NewDeniedf returns a formatted error that satisfies IsDenied().
func NewDeniedf(format string, args ...interface{}) error {
	return deniedf{errNewf(format, args...)}
}

func isDenied(err error) (ok bool) {
	type iFace interface {
		Denied() bool
	}
	switch et := err.(type) {
	case denied:
		ok = true
	case deniedf:
		ok = true
	case iFace:
		ok = et.Denied()
	}
	return
}

// IsDenied reports whether err was created with NewDenied() or
// implements interface:
//     type Denieder interface {
//            Denied() bool
//     }
func IsDenied(err error) bool {
	return CausedBehaviour(err, isDenied)
}

type (
	duplicated  struct{ *withStack }
	duplicatedf struct{ *fundamental }
)

// NewDuplicated returns an error which wraps err that satisfies
// IsDuplicated().
func NewDuplicated(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return duplicated{errWrapf(err, msg, args...)}
}

// NewDuplicatedf returns a formatted error that satisfies IsDuplicated().
func NewDuplicatedf(format string, args ...interface{}) error {
	return duplicatedf{errNewf(format, args...)}
}

func isDuplicated(err error) (ok bool) {
	type iFace interface {
		Duplicated() bool
	}
	switch et := err.(type) {
	case duplicated:
		ok = true
	case duplicatedf:
		ok = true
	case iFace:
		ok = et.Duplicated()
	}
	return
}

// IsDuplicated reports whether err was created with NewDuplicated() or
// implements interface:
//     type Duplicateder interface {
//            Duplicated() bool
//     }
func IsDuplicated(err error) bool {
	return CausedBehaviour(err, isDuplicated)
}

type (
	empty  struct{ *withStack }
	emptyf struct{ *fundamental }
)

// NewEmpty returns an error which wraps err that satisfies
// IsEmpty().
func NewEmpty(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return empty{errWrapf(err, msg, args...)}
}

// NewEmptyf returns a formatted error that satisfies IsEmpty().
func NewEmptyf(format string, args ...interface{}) error {
	return emptyf{errNewf(format, args...)}
}

func isEmpty(err error) (ok bool) {
	type iFace interface {
		Empty() bool
	}
	switch et := err.(type) {
	case empty:
		ok = true
	case emptyf:
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
	exceeded  struct{ *withStack }
	exceededf struct{ *fundamental }
)

// NewExceeded returns an error which wraps err that satisfies
// IsExceeded().
func NewExceeded(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return exceeded{errWrapf(err, msg, args...)}
}

// NewExceededf returns a formatted error that satisfies IsExceeded().
func NewExceededf(format string, args ...interface{}) error {
	return exceededf{errNewf(format, args...)}
}

func isExceeded(err error) (ok bool) {
	type iFace interface {
		Exceeded() bool
	}
	switch et := err.(type) {
	case exceeded:
		ok = true
	case exceededf:
		ok = true
	case iFace:
		ok = et.Exceeded()
	}
	return
}

// IsExceeded reports whether err was created with NewExceeded() or
// implements interface:
//     type Exceededer interface {
//            Exceeded() bool
//     }
func IsExceeded(err error) bool {
	return CausedBehaviour(err, isExceeded)
}

type (
	expired  struct{ *withStack }
	expiredf struct{ *fundamental }
)

// NewExpired returns an error which wraps err that satisfies
// IsExpired().
func NewExpired(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return expired{errWrapf(err, msg, args...)}
}

// NewExpiredf returns a formatted error that satisfies IsExpired().
func NewExpiredf(format string, args ...interface{}) error {
	return expiredf{errNewf(format, args...)}
}

func isExpired(err error) (ok bool) {
	type iFace interface {
		Expired() bool
	}
	switch et := err.(type) {
	case expired:
		ok = true
	case expiredf:
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
	fatal  struct{ *withStack }
	fatalf struct{ *fundamental }
)

// NewFatal returns an error which wraps err that satisfies
// IsFatal().
func NewFatal(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fatal{errWrapf(err, msg, args...)}
}

// NewFatalf returns a formatted error that satisfies IsFatal().
func NewFatalf(format string, args ...interface{}) error {
	return fatalf{errNewf(format, args...)}
}

func isFatal(err error) (ok bool) {
	type iFace interface {
		Fatal() bool
	}
	switch et := err.(type) {
	case fatal:
		ok = true
	case fatalf:
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
	inProgress  struct{ *withStack }
	inProgressf struct{ *fundamental }
)

// NewInProgress returns an error which wraps err that satisfies
// IsInProgress().
func NewInProgress(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return inProgress{errWrapf(err, msg, args...)}
}

// NewInProgressf returns a formatted error that satisfies IsInProgress().
func NewInProgressf(format string, args ...interface{}) error {
	return inProgressf{errNewf(format, args...)}
}

func isInProgress(err error) (ok bool) {
	type iFace interface {
		InProgress() bool
	}
	switch et := err.(type) {
	case inProgress:
		ok = true
	case inProgressf:
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
	insufficient  struct{ *withStack }
	insufficientf struct{ *fundamental }
)

// NewInsufficient returns an error which wraps err that satisfies
// IsInsufficient().
func NewInsufficient(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return insufficient{errWrapf(err, msg, args...)}
}

// NewInsufficientf returns a formatted error that satisfies IsInsufficient().
func NewInsufficientf(format string, args ...interface{}) error {
	return insufficientf{errNewf(format, args...)}
}

func isInsufficient(err error) (ok bool) {
	type iFace interface {
		Insufficient() bool
	}
	switch et := err.(type) {
	case insufficient:
		ok = true
	case insufficientf:
		ok = true
	case iFace:
		ok = et.Insufficient()
	}
	return
}

// IsInsufficient reports whether err was created with NewInsufficient() or
// implements interface:
//     type Insufficienter interface {
//            Insufficient() bool
//     }
func IsInsufficient(err error) bool {
	return CausedBehaviour(err, isInsufficient)
}

type (
	interrupted  struct{ *withStack }
	interruptedf struct{ *fundamental }
)

// NewInterrupted returns an error which wraps err that satisfies
// IsInterrupted().
func NewInterrupted(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return interrupted{errWrapf(err, msg, args...)}
}

// NewInterruptedf returns a formatted error that satisfies IsInterrupted().
func NewInterruptedf(format string, args ...interface{}) error {
	return interruptedf{errNewf(format, args...)}
}

func isInterrupted(err error) (ok bool) {
	type iFace interface {
		Interrupted() bool
	}
	switch et := err.(type) {
	case interrupted:
		ok = true
	case interruptedf:
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
	locked  struct{ *withStack }
	lockedf struct{ *fundamental }
)

// NewLocked returns an error which wraps err that satisfies
// IsLocked().
func NewLocked(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return locked{errWrapf(err, msg, args...)}
}

// NewLockedf returns a formatted error that satisfies IsLocked().
func NewLockedf(format string, args ...interface{}) error {
	return lockedf{errNewf(format, args...)}
}

func isLocked(err error) (ok bool) {
	type iFace interface {
		Locked() bool
	}
	switch et := err.(type) {
	case locked:
		ok = true
	case lockedf:
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
	mismatch  struct{ *withStack }
	mismatchf struct{ *fundamental }
)

// NewMismatch returns an error which wraps err that satisfies
// IsMismatch().
func NewMismatch(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return mismatch{errWrapf(err, msg, args...)}
}

// NewMismatchf returns a formatted error that satisfies IsMismatch().
func NewMismatchf(format string, args ...interface{}) error {
	return mismatchf{errNewf(format, args...)}
}

func isMismatch(err error) (ok bool) {
	type iFace interface {
		Mismatch() bool
	}
	switch et := err.(type) {
	case mismatch:
		ok = true
	case mismatchf:
		ok = true
	case iFace:
		ok = et.Mismatch()
	}
	return
}

// IsMismatch reports whether err was created with NewMismatch() or
// implements interface:
//     type Mismatcher interface {
//            Mismatch() bool
//     }
func IsMismatch(err error) bool {
	return CausedBehaviour(err, isMismatch)
}

type (
	notAcceptable  struct{ *withStack }
	notAcceptablef struct{ *fundamental }
)

// NewNotAcceptable returns an error which wraps err that satisfies
// IsNotAcceptable().
func NewNotAcceptable(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return notAcceptable{errWrapf(err, msg, args...)}
}

// NewNotAcceptablef returns a formatted error that satisfies IsNotAcceptable().
func NewNotAcceptablef(format string, args ...interface{}) error {
	return notAcceptablef{errNewf(format, args...)}
}

func isNotAcceptable(err error) (ok bool) {
	type iFace interface {
		NotAcceptable() bool
	}
	switch et := err.(type) {
	case notAcceptable:
		ok = true
	case notAcceptablef:
		ok = true
	case iFace:
		ok = et.NotAcceptable()
	}
	return
}

// IsNotAcceptable reports whether err was created with NewNotAcceptable() or
// implements interface:
//     type NotAcceptableer interface {
//            NotAcceptable() bool
//     }
func IsNotAcceptable(err error) bool {
	return CausedBehaviour(err, isNotAcceptable)
}

type (
	notAllowed  struct{ *withStack }
	notAllowedf struct{ *fundamental }
)

// NewNotAllowed returns an error which wraps err that satisfies
// IsNotAllowed().
func NewNotAllowed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return notAllowed{errWrapf(err, msg, args...)}
}

// NewNotAllowedf returns a formatted error that satisfies IsNotAllowed().
func NewNotAllowedf(format string, args ...interface{}) error {
	return notAllowedf{errNewf(format, args...)}
}

func isNotAllowed(err error) (ok bool) {
	type iFace interface {
		NotAllowed() bool
	}
	switch et := err.(type) {
	case notAllowed:
		ok = true
	case notAllowedf:
		ok = true
	case iFace:
		ok = et.NotAllowed()
	}
	return
}

// IsNotAllowed reports whether err was created with NewNotAllowed() or
// implements interface:
//     type NotAlloweder interface {
//            NotAllowed() bool
//     }
func IsNotAllowed(err error) bool {
	return CausedBehaviour(err, isNotAllowed)
}

type (
	notFound  struct{ *withStack }
	notFoundf struct{ *fundamental }
)

// NewNotFound returns an error which wraps err that satisfies
// IsNotFound().
func NewNotFound(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return notFound{errWrapf(err, msg, args...)}
}

// NewNotFoundf returns a formatted error that satisfies IsNotFound().
func NewNotFoundf(format string, args ...interface{}) error {
	return notFoundf{errNewf(format, args...)}
}

func isNotFound(err error) (ok bool) {
	type iFace interface {
		NotFound() bool
	}
	switch et := err.(type) {
	case notFound:
		ok = true
	case notFoundf:
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
	notImplemented  struct{ *withStack }
	notImplementedf struct{ *fundamental }
)

// NewNotImplemented returns an error which wraps err that satisfies
// IsNotImplemented().
func NewNotImplemented(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return notImplemented{errWrapf(err, msg, args...)}
}

// NewNotImplementedf returns a formatted error that satisfies IsNotImplemented().
func NewNotImplementedf(format string, args ...interface{}) error {
	return notImplementedf{errNewf(format, args...)}
}

func isNotImplemented(err error) (ok bool) {
	type iFace interface {
		NotImplemented() bool
	}
	switch et := err.(type) {
	case notImplemented:
		ok = true
	case notImplementedf:
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
	notRecoverable  struct{ *withStack }
	notRecoverablef struct{ *fundamental }
)

// NewNotRecoverable returns an error which wraps err that satisfies
// IsNotRecoverable().
func NewNotRecoverable(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return notRecoverable{errWrapf(err, msg, args...)}
}

// NewNotRecoverablef returns a formatted error that satisfies IsNotRecoverable().
func NewNotRecoverablef(format string, args ...interface{}) error {
	return notRecoverablef{errNewf(format, args...)}
}

func isNotRecoverable(err error) (ok bool) {
	type iFace interface {
		NotRecoverable() bool
	}
	switch et := err.(type) {
	case notRecoverable:
		ok = true
	case notRecoverablef:
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
	notSupported  struct{ *withStack }
	notSupportedf struct{ *fundamental }
)

// NewNotSupported returns an error which wraps err that satisfies
// IsNotSupported().
func NewNotSupported(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return notSupported{errWrapf(err, msg, args...)}
}

// NewNotSupportedf returns a formatted error that satisfies IsNotSupported().
func NewNotSupportedf(format string, args ...interface{}) error {
	return notSupportedf{errNewf(format, args...)}
}

func isNotSupported(err error) (ok bool) {
	type iFace interface {
		NotSupported() bool
	}
	switch et := err.(type) {
	case notSupported:
		ok = true
	case notSupportedf:
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
	notValid  struct{ *withStack }
	notValidf struct{ *fundamental }
)

// NewNotValid returns an error which wraps err that satisfies
// IsNotValid().
func NewNotValid(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return notValid{errWrapf(err, msg, args...)}
}

// NewNotValidf returns a formatted error that satisfies IsNotValid().
func NewNotValidf(format string, args ...interface{}) error {
	return notValidf{errNewf(format, args...)}
}

func isNotValid(err error) (ok bool) {
	type iFace interface {
		NotValid() bool
	}
	switch et := err.(type) {
	case notValid:
		ok = true
	case notValidf:
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
	permissionDenied  struct{ *withStack }
	permissionDeniedf struct{ *fundamental }
)

// NewPermissionDenied returns an error which wraps err that satisfies
// IsPermissionDenied().
func NewPermissionDenied(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return permissionDenied{errWrapf(err, msg, args...)}
}

// NewPermissionDeniedf returns a formatted error that satisfies IsPermissionDenied().
func NewPermissionDeniedf(format string, args ...interface{}) error {
	return permissionDeniedf{errNewf(format, args...)}
}

func isPermissionDenied(err error) (ok bool) {
	type iFace interface {
		PermissionDenied() bool
	}
	switch et := err.(type) {
	case permissionDenied:
		ok = true
	case permissionDeniedf:
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
	quotaExceeded  struct{ *withStack }
	quotaExceededf struct{ *fundamental }
)

// NewQuotaExceeded returns an error which wraps err that satisfies
// IsQuotaExceeded().
func NewQuotaExceeded(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return quotaExceeded{errWrapf(err, msg, args...)}
}

// NewQuotaExceededf returns a formatted error that satisfies IsQuotaExceeded().
func NewQuotaExceededf(format string, args ...interface{}) error {
	return quotaExceededf{errNewf(format, args...)}
}

func isQuotaExceeded(err error) (ok bool) {
	type iFace interface {
		QuotaExceeded() bool
	}
	switch et := err.(type) {
	case quotaExceeded:
		ok = true
	case quotaExceededf:
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
	readFailed  struct{ *withStack }
	readFailedf struct{ *fundamental }
)

// NewReadFailed returns an error which wraps err that satisfies
// IsReadFailed().
func NewReadFailed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return readFailed{errWrapf(err, msg, args...)}
}

// NewReadFailedf returns a formatted error that satisfies IsReadFailed().
func NewReadFailedf(format string, args ...interface{}) error {
	return readFailedf{errNewf(format, args...)}
}

func isReadFailed(err error) (ok bool) {
	type iFace interface {
		ReadFailed() bool
	}
	switch et := err.(type) {
	case readFailed:
		ok = true
	case readFailedf:
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
	rejected  struct{ *withStack }
	rejectedf struct{ *fundamental }
)

// NewRejected returns an error which wraps err that satisfies
// IsRejected().
func NewRejected(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return rejected{errWrapf(err, msg, args...)}
}

// NewRejectedf returns a formatted error that satisfies IsRejected().
func NewRejectedf(format string, args ...interface{}) error {
	return rejectedf{errNewf(format, args...)}
}

func isRejected(err error) (ok bool) {
	type iFace interface {
		Rejected() bool
	}
	switch et := err.(type) {
	case rejected:
		ok = true
	case rejectedf:
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
	required  struct{ *withStack }
	requiredf struct{ *fundamental }
)

// NewRequired returns an error which wraps err that satisfies
// IsRequired().
func NewRequired(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return required{errWrapf(err, msg, args...)}
}

// NewRequiredf returns a formatted error that satisfies IsRequired().
func NewRequiredf(format string, args ...interface{}) error {
	return requiredf{errNewf(format, args...)}
}

func isRequired(err error) (ok bool) {
	type iFace interface {
		Required() bool
	}
	switch et := err.(type) {
	case required:
		ok = true
	case requiredf:
		ok = true
	case iFace:
		ok = et.Required()
	}
	return
}

// IsRequired reports whether err was created with NewRequired() or
// implements interface:
//     type Requireder interface {
//            Required() bool
//     }
func IsRequired(err error) bool {
	return CausedBehaviour(err, isRequired)
}

type (
	restricted  struct{ *withStack }
	restrictedf struct{ *fundamental }
)

// NewRestricted returns an error which wraps err that satisfies
// IsRestricted().
func NewRestricted(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return restricted{errWrapf(err, msg, args...)}
}

// NewRestrictedf returns a formatted error that satisfies IsRestricted().
func NewRestrictedf(format string, args ...interface{}) error {
	return restrictedf{errNewf(format, args...)}
}

func isRestricted(err error) (ok bool) {
	type iFace interface {
		Restricted() bool
	}
	switch et := err.(type) {
	case restricted:
		ok = true
	case restrictedf:
		ok = true
	case iFace:
		ok = et.Restricted()
	}
	return
}

// IsRestricted reports whether err was created with NewRestricted() or
// implements interface:
//     type Restricteder interface {
//            Restricted() bool
//     }
func IsRestricted(err error) bool {
	return CausedBehaviour(err, isRestricted)
}

type (
	revoked  struct{ *withStack }
	revokedf struct{ *fundamental }
)

// NewRevoked returns an error which wraps err that satisfies
// IsRevoked().
func NewRevoked(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return revoked{errWrapf(err, msg, args...)}
}

// NewRevokedf returns a formatted error that satisfies IsRevoked().
func NewRevokedf(format string, args ...interface{}) error {
	return revokedf{errNewf(format, args...)}
}

func isRevoked(err error) (ok bool) {
	type iFace interface {
		Revoked() bool
	}
	switch et := err.(type) {
	case revoked:
		ok = true
	case revokedf:
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
	temporary  struct{ *withStack }
	temporaryf struct{ *fundamental }
)

// NewTemporary returns an error which wraps err that satisfies
// IsTemporary().
func NewTemporary(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return temporary{errWrapf(err, msg, args...)}
}

// NewTemporaryf returns a formatted error that satisfies IsTemporary().
func NewTemporaryf(format string, args ...interface{}) error {
	return temporaryf{errNewf(format, args...)}
}

func isTemporary(err error) (ok bool) {
	type iFace interface {
		Temporary() bool
	}
	switch et := err.(type) {
	case temporary:
		ok = true
	case temporaryf:
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
	terminated  struct{ *withStack }
	terminatedf struct{ *fundamental }
)

// NewTerminated returns an error which wraps err that satisfies
// IsTerminated().
func NewTerminated(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return terminated{errWrapf(err, msg, args...)}
}

// NewTerminatedf returns a formatted error that satisfies IsTerminated().
func NewTerminatedf(format string, args ...interface{}) error {
	return terminatedf{errNewf(format, args...)}
}

func isTerminated(err error) (ok bool) {
	type iFace interface {
		Terminated() bool
	}
	switch et := err.(type) {
	case terminated:
		ok = true
	case terminatedf:
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
	timeout  struct{ *withStack }
	timeoutf struct{ *fundamental }
)

// NewTimeout returns an error which wraps err that satisfies
// IsTimeout().
func NewTimeout(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return timeout{errWrapf(err, msg, args...)}
}

// NewTimeoutf returns a formatted error that satisfies IsTimeout().
func NewTimeoutf(format string, args ...interface{}) error {
	return timeoutf{errNewf(format, args...)}
}

func isTimeout(err error) (ok bool) {
	type iFace interface {
		Timeout() bool
	}
	switch et := err.(type) {
	case timeout:
		ok = true
	case timeoutf:
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
	tooLarge  struct{ *withStack }
	tooLargef struct{ *fundamental }
)

// NewTooLarge returns an error which wraps err that satisfies
// IsTooLarge().
func NewTooLarge(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return tooLarge{errWrapf(err, msg, args...)}
}

// NewTooLargef returns a formatted error that satisfies IsTooLarge().
func NewTooLargef(format string, args ...interface{}) error {
	return tooLargef{errNewf(format, args...)}
}

func isTooLarge(err error) (ok bool) {
	type iFace interface {
		TooLarge() bool
	}
	switch et := err.(type) {
	case tooLarge:
		ok = true
	case tooLargef:
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
	unauthorized  struct{ *withStack }
	unauthorizedf struct{ *fundamental }
)

// NewUnauthorized returns an error which wraps err that satisfies
// IsUnauthorized().
func NewUnauthorized(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return unauthorized{errWrapf(err, msg, args...)}
}

// NewUnauthorizedf returns a formatted error that satisfies IsUnauthorized().
func NewUnauthorizedf(format string, args ...interface{}) error {
	return unauthorizedf{errNewf(format, args...)}
}

func isUnauthorized(err error) (ok bool) {
	type iFace interface {
		Unauthorized() bool
	}
	switch et := err.(type) {
	case unauthorized:
		ok = true
	case unauthorizedf:
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
	unavailable  struct{ *withStack }
	unavailablef struct{ *fundamental }
)

// NewUnavailable returns an error which wraps err that satisfies
// IsUnavailable().
func NewUnavailable(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return unavailable{errWrapf(err, msg, args...)}
}

// NewUnavailablef returns a formatted error that satisfies IsUnavailable().
func NewUnavailablef(format string, args ...interface{}) error {
	return unavailablef{errNewf(format, args...)}
}

func isUnavailable(err error) (ok bool) {
	type iFace interface {
		Unavailable() bool
	}
	switch et := err.(type) {
	case unavailable:
		ok = true
	case unavailablef:
		ok = true
	case iFace:
		ok = et.Unavailable()
	}
	return
}

// IsUnavailable reports whether err was created with NewUnavailable() or
// implements interface:
//     type Unavailableer interface {
//            Unavailable() bool
//     }
func IsUnavailable(err error) bool {
	return CausedBehaviour(err, isUnavailable)
}

type (
	userNotFound  struct{ *withStack }
	userNotFoundf struct{ *fundamental }
)

// NewUserNotFound returns an error which wraps err that satisfies
// IsUserNotFound().
func NewUserNotFound(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return userNotFound{errWrapf(err, msg, args...)}
}

// NewUserNotFoundf returns a formatted error that satisfies IsUserNotFound().
func NewUserNotFoundf(format string, args ...interface{}) error {
	return userNotFoundf{errNewf(format, args...)}
}

func isUserNotFound(err error) (ok bool) {
	type iFace interface {
		UserNotFound() bool
	}
	switch et := err.(type) {
	case userNotFound:
		ok = true
	case userNotFoundf:
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
	verificationFailed  struct{ *withStack }
	verificationFailedf struct{ *fundamental }
)

// NewVerificationFailed returns an error which wraps err that satisfies
// IsVerificationFailed().
func NewVerificationFailed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return verificationFailed{errWrapf(err, msg, args...)}
}

// NewVerificationFailedf returns a formatted error that satisfies IsVerificationFailed().
func NewVerificationFailedf(format string, args ...interface{}) error {
	return verificationFailedf{errNewf(format, args...)}
}

func isVerificationFailed(err error) (ok bool) {
	type iFace interface {
		VerificationFailed() bool
	}
	switch et := err.(type) {
	case verificationFailed:
		ok = true
	case verificationFailedf:
		ok = true
	case iFace:
		ok = et.VerificationFailed()
	}
	return
}

// IsVerificationFailed reports whether err was created with NewVerificationFailed() or
// implements interface:
//     type VerificationFaileder interface {
//            VerificationFailed() bool
//     }
func IsVerificationFailed(err error) bool {
	return CausedBehaviour(err, isVerificationFailed)
}

type (
	writeFailed  struct{ *withStack }
	writeFailedf struct{ *fundamental }
)

// NewWriteFailed returns an error which wraps err that satisfies
// IsWriteFailed().
func NewWriteFailed(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return writeFailed{errWrapf(err, msg, args...)}
}

// NewWriteFailedf returns a formatted error that satisfies IsWriteFailed().
func NewWriteFailedf(format string, args ...interface{}) error {
	return writeFailedf{errNewf(format, args...)}
}

func isWriteFailed(err error) (ok bool) {
	type iFace interface {
		WriteFailed() bool
	}
	switch et := err.(type) {
	case writeFailed:
		ok = true
	case writeFailedf:
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
	wrongVersion  struct{ *withStack }
	wrongVersionf struct{ *fundamental }
)

// NewWrongVersion returns an error which wraps err that satisfies
// IsWrongVersion().
func NewWrongVersion(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrongVersion{errWrapf(err, msg, args...)}
}

// NewWrongVersionf returns a formatted error that satisfies IsWrongVersion().
func NewWrongVersionf(format string, args ...interface{}) error {
	return wrongVersionf{errNewf(format, args...)}
}

func isWrongVersion(err error) (ok bool) {
	type iFace interface {
		WrongVersion() bool
	}
	switch et := err.(type) {
	case wrongVersion:
		ok = true
	case wrongVersionf:
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
