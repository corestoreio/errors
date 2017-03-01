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

// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"unicode"
)

var behaviours = []string{
	"aborted",
	"alreadyClosed",
	"alreadyExists",
	"alreadyInUse",
	"connectionFailed",
	"empty",
	"expired",
	"fatal",
	"inProgress",
	"interrupted",
	"locked",
	"methodNotAllowed",
	"notFound",
	"notImplemented",
	"notRecoverable",
	"notSupported",
	"notValid",
	"permissionDenied",
	"previousOwnerDied", // not idea but sounds funny
	"quotaExceeded",
	"readFailed",
	"rejected",
	"revoked",
	"temporary",
	"terminated",
	"timeout",
	"tooLarge",
	"unauthorized",
	"userNotFound",
	"writeFailed",
	"wrongVersion",
}

func main() {

	tpl := template.Must(template.New("behaviour").Parse(behaviourTemplate))
	tplTest := template.Must(template.New("behaviour_test").Parse(behaviourTestTemplate))

	var buf bytes.Buffer
	var bufTest bytes.Buffer
	buf.WriteString("// Auto generated via behaviour_gen.go\n\npackage errors\n")
	bufTest.WriteString(`// Auto generated via behaviour_gen.go

package errors

import (
	"errors"
	"testing"
)
`)
	if len(os.Args) > 1 {
		behaviours = os.Args[1:]
	}

	for _, b := range behaviours {

		if containsSeparator(b) {
			panic(fmt.Sprintf("%q contains an illegal character", b))
		}

		data := struct {
			LcName string
			TcName string
		}{
			LcName: makeFirstLC(b),
			TcName: strings.Title(b),
		}
		if err := tpl.Execute(&buf, data); err != nil {
			panic(fmt.Sprintf("Data: %v\nError: %s\n\n", data, err))
		}
		if err := tplTest.Execute(&bufTest, data); err != nil {
			panic(fmt.Sprintf("Data: %v\nError: %s\n\n", data, err))
		}
	}

	writeFormatted("behaviour.go", buf.Bytes())
	writeFormatted("behaviour_test.go", bufTest.Bytes())

	fmt.Printf("Wrote behavioural errors: %v\n", behaviours)
}

func writeFormatted(fn string, data []byte) {
	var err error
	data, err = format.Source(data)
	if err != nil {
		panic(fmt.Sprintf("%s\n%s\n\n", err, data))
	}
	if err := ioutil.WriteFile(fn, data, 0644); err != nil {
		panic(err)
	}
}

func makeFirstLC(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

func containsSeparator(s string) bool {
	for _, r := range s {
		if isSeparator(r) {
			return true
		}
	}
	return false
}

func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	return unicode.IsSpace(r)
}

const behaviourTemplate = `
type (
	{{.LcName}}  struct{ wrapper }
	{{.LcName}}f struct{ _error }
)

// New{{.TcName}} returns an error which wraps err that satisfies
// Is{{.TcName}}().
func New{{.TcName}}(err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &{{.LcName}}{errWrapf(err, msg, args...)}
}

// New{{.TcName}}f returns a formatted error that satisfies Is{{.TcName}}().
func New{{.TcName}}f(format string, args ...interface{}) error {
	return &{{.LcName}}f{errNewf(format, args...)}
}

func is{{.TcName}}(err error) (ok bool) {
	type iFace interface {
		{{.TcName}}() bool
	}
	switch et := err.(type) {
	case *{{.LcName}}:
		ok = true
	case *{{.LcName}}f:
		ok = true
	case iFace:
		ok = et.{{.TcName}}()
	}
	return
}

// Is{{.TcName}} reports whether err was created with New{{.TcName}}() or
// implements interface:
//     type {{.TcName}}er interface {
//            {{.TcName}}() bool
//     }
func Is{{.TcName}}(err error) bool {
	return CausedBehaviour(err, is{{.TcName}})
}
`

const behaviourTestTemplate = `
func (nf testBehave) {{.TcName}}() bool {
	return nf.ret
}

func Test{{.TcName}}(t *testing.T) {
	t.Parallel()

	tests := []struct {
		err  error
		is   BehaviourFunc
		want bool
	}{
		{ errors.New("Error1"), Is{{.TcName}}, false},
		{ New{{.TcName}}(nil, "Error2"), Is{{.TcName}}, false },
		{ New{{.TcName}}(Error("Error3a"), "Error3"), Is{{.TcName}}, true },
		{ Wrap(New{{.TcName}}f("Err4"), "Wrap4"), Is{{.TcName}}, true },
		{ NewNotImplemented(Wrap(New{{.TcName}}f("Err5"), "Wrap5"), "NotImplemend5"), Is{{.TcName}}, true },
		{ Wrap(New{{.TcName}}(Wrap(NewNotImplementedf("Err6"), "Wrap6"), "{{.TcName}}6"), "Wrap6a"), Is{{.TcName}}, true },
		{ Wrap(New{{.TcName}}(errors.New("I'm the cause7"), "{{.TcName}}7"), "Wrap7"), Is{{.TcName}}, true },
		{ New{{.TcName}}f("Error8"), Is{{.TcName}}, true },
		{ nil, Is{{.TcName}}, false },
		{ testBehave{}, Is{{.TcName}}, false },
		{ testBehave{ret: true}, Is{{.TcName}}, true },
	}
	for i, test := range tests {
		if want, have := test.want, test.is(test.err); want != have {
			t.Errorf("Index %d: Want %t Have %t", i, want, have)
		}
	}
}
`
