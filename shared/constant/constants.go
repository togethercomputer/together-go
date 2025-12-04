// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/togethercomputer/together-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Error string      // Always "error"
type JsonObject string // Always "json_object"
type JsonSchema string // Always "json_schema"
type Text string       // Always "text"

func (c Error) Default() Error           { return "error" }
func (c JsonObject) Default() JsonObject { return "json_object" }
func (c JsonSchema) Default() JsonSchema { return "json_schema" }
func (c Text) Default() Text             { return "text" }

func (c Error) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c JsonObject) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c JsonSchema) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)       { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
