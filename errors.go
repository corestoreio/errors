// Copyright (c) 2015, Dave Cheney <dave@cheney.net>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
// list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package errors

import (
	"errors"
	"fmt"
	"io"
)

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(message string) error {
	return &fundamental{
		msg:   message,
		stack: callers(),
	}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...interface{}) error {
	return &fundamental{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(),
	}
}

// fundamental is an error that has a message and a stack, but no caller.
type fundamental struct {
	msg string
	*stack
}

func (f *fundamental) Error() string { return f.msg }

func (f *fundamental) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, f.msg)
			f.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, f.msg)
	case 'q':
		fmt.Fprintf(s, "%q", f.msg)
	}
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	return &withStack{
		err,
		callers(),
	}
}

type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error  { return w.error }
func (w *withStack) Unwrap() error { return w.error }

func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   message,
	}
	return &withStack{
		err,
		callers(),
	}
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
	return &withStack{
		err,
		callers(),
	}
}

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// As will panic if target is not a non-nil pointer to either a type that implements
// error, or to any interface type. As returns false if err is nil.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   message,
	}
}

// WithMessagef annotates err with the format specifier.
// If err is nil, WithMessagef returns nil.
func WithMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
}

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }
func (w *withMessage) Cause() error  { return w.cause }
func (w *withMessage) Unwrap() error { return w.cause }

func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			io.WriteString(s, w.msg)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error())
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	for err != nil {
		var c causer
		if !As(err, &c) {
			break
		}
		err = c.Cause()
	}
	return err
}

type causer interface {
	Cause() error
}
