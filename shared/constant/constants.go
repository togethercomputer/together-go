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

type AudioTtsChunk string       // Always "audio.tts.chunk"
type ChatCompletion string      // Always "chat.completion"
type ChatCompletionChunk string // Always "chat.completion.chunk"
type Embedding string           // Always "embedding"
type Endpoint string            // Always "endpoint"
type Error string               // Always "error"
type File string                // Always "file"
type FineTuneEvent string       // Always "fine-tune-event"
type Hardware string            // Always "hardware"
type JsonObject string          // Always "json_object"
type JsonSchema string          // Always "json_schema"
type List string                // Always "list"
type Model string               // Always "model"
type Rerank string              // Always "rerank"
type Text string                // Always "text"
type TextCompletion string      // Always "text.completion"

func (c AudioTtsChunk) Default() AudioTtsChunk             { return "audio.tts.chunk" }
func (c ChatCompletion) Default() ChatCompletion           { return "chat.completion" }
func (c ChatCompletionChunk) Default() ChatCompletionChunk { return "chat.completion.chunk" }
func (c Embedding) Default() Embedding                     { return "embedding" }
func (c Endpoint) Default() Endpoint                       { return "endpoint" }
func (c Error) Default() Error                             { return "error" }
func (c File) Default() File                               { return "file" }
func (c FineTuneEvent) Default() FineTuneEvent             { return "fine-tune-event" }
func (c Hardware) Default() Hardware                       { return "hardware" }
func (c JsonObject) Default() JsonObject                   { return "json_object" }
func (c JsonSchema) Default() JsonSchema                   { return "json_schema" }
func (c List) Default() List                               { return "list" }
func (c Model) Default() Model                             { return "model" }
func (c Rerank) Default() Rerank                           { return "rerank" }
func (c Text) Default() Text                               { return "text" }
func (c TextCompletion) Default() TextCompletion           { return "text.completion" }

func (c AudioTtsChunk) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ChatCompletion) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ChatCompletionChunk) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Embedding) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Endpoint) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Error) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c FineTuneEvent) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Hardware) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c JsonObject) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c JsonSchema) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c List) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Model) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Rerank) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c TextCompletion) MarshalJSON() ([]byte, error)      { return marshalString(c) }

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
