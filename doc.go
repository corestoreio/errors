// Package errors provides simple error handling primitives and behavioral errors.
//
// Hello there! Read the presentation http://dave.cheney.net/paste/gocon-spring-2016.pdf
// to see what the big deal is.
//
// http://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
//
// Read this for asserting errors for their behaviour http://dave.cheney.net/2014/12/24/inspecting-errors
//
// The traditional error handling idiom in Go is roughly akin to
//
//     if err != nil {
//             return err
//     }
//
// which applied recursively up the call stack results in error reports
// without context or debugging information. The errors package allows
// programmers to add context to the failure path in their code in a way
// that does not destroy the original value of the error.
//
// Adding context to an error
//
// The errors.Wrap function returns a new error that adds context to the
// original error by recording a stack trace at the point Wrap is called,
// and the supplied message. For example
//
//     _, err := ioutil.ReadAll(r)
//     if err != nil {
//             return errors.Wrap(err, "read failed")
//     }
//
// If additional control is required the errors.WithStack and errors.WithMessage
// functions destructure errors.Wrap into its component operations of annotating
// an error with a stack trace and an a message, respectively.
//
// Retrieving the cause of an error
//
// Using errors.Wrap constructs a stack of errors, adding context to the
// preceding error. Depending on the nature of the error it may be necessary
// to reverse the operation of errors.Wrap to retrieve the original error
// for inspection. Any error value which implements this interface
//
//     type causer interface {
//             Cause() error
//     }
//
// can be inspected by errors.Cause. errors.Cause will recursively retrieve
// the topmost error which does not implement causer, which is assumed to be
// the original cause. For example:
//
//     switch err := errors.Cause(err).(type) {
//     case *MyError:
//             // handle specifically
//     default:
//             // unknown error
//     }
//
// causer interface is not exported by this package, but is considered a part
// of stable public API.
//
// Formatted printing of errors
//
// All error values returned from this package implement fmt.Formatter and can
// be formatted by the fmt package. The following verbs are supported
//
//     %s    print the error. If the error has a Cause it will be
//           printed recursively
//     %v    see %s
//     %+v   extended format. Each Frame of the error's StackTrace will
//           be printed in detail.
//
// Retrieving the stack trace of an error or wrapper
//
// New, Errorf, Wrap, and Wrapf record a stack trace at the point they are
// invoked. This information can be retrieved with the following interface.
//
//     type stackTracer interface {
//             StackTrace() errors.StackTrace
//     }
//
// Where errors.StackTrace is defined as
//
//     type StackTrace []Frame
//
// The Frame type represents a call site in the stack trace. Frame supports
// the fmt.Formatter interface that can be used for printing information about
// the stack trace of this error. For example:
//
//     if err, ok := err.(stackTracer); ok {
//             for _, f := range err.StackTrace() {
//                     fmt.Printf("%+s:%d", f)
//             }
//     }
//
// stackTracer interface is not exported by this package, but is considered a part
// of stable public API.
//
// See the documentation for Frame.Format for more details.
package errors
