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

/*
	BRAIN FART of structured errors with additional information.
	For now commented out.
*/

//
//type fieldType uint8
//
//// Type* constants define all available types which a field can contain.
//const (
//	typeBool fieldType = iota + 1
//	typeInt
//	typeInts
//	typeInt64
//	typeInt64s
//	typeUint64
//	typeFloat64
//	typeString
//	typeStrings
//	typeStringer
//	typeStringFn
//	typeGoStringer
//	typeObject
//	typeObjectTypeOf
//	typeMarshaler
//	typeFields
//)
//
//// Fields a slice of n Field types. Fields implements the Field interface and
//// can be added to a function in the Logger interface in a simple way.
//type fields []Field
//
//// field is a deferred marshaling operation used to add a key-value pair to
//// a logger's context. Keys and values are appropriately escaped for the current
//// encoding scheme (e.g., JSON).
//type Field struct {
//	key string
//	// fieldType specifies the used type. If 0 this struct is empty
//	fieldType
//	int64
//	uint64
//	float64
//	string
//	obj interface{}
//}
//
//func (k Kind) WithFields(err error, fs ...Field) {
//	// add type fields to the structs kindStacked and kindFundamental
//}
