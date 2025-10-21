// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
	"github.com/togethercomputer/together-go/shared/constant"
)

// CodeInterpreterService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCodeInterpreterService] method instead.
type CodeInterpreterService struct {
	Options  []option.RequestOption
	Sessions CodeInterpreterSessionService
}

// NewCodeInterpreterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCodeInterpreterService(opts ...option.RequestOption) (r CodeInterpreterService) {
	r = CodeInterpreterService{}
	r.Options = opts
	r.Sessions = NewCodeInterpreterSessionService(opts...)
	return
}

// Executes the given code snippet and returns the output. Without a session_id, a
// new session will be created to run the code. If you do pass in a valid
// session_id, the code will be run in that session. This is useful for running
// multiple code snippets in the same environment, because dependencies and similar
// things are persisted between calls to the same session.
func (r *CodeInterpreterService) Execute(ctx context.Context, body CodeInterpreterExecuteParams, opts ...option.RequestOption) (res *ExecuteResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tci/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ExecuteResponseUnion contains all possible properties and values from
// [ExecuteResponseSuccessfulExecution], [ExecuteResponseFailedExecution].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExecuteResponseUnion struct {
	// This field is a union of [ExecuteResponseSuccessfulExecutionData], [any]
	Data ExecuteResponseUnionData `json:"data"`
	// This field is a union of [any], [[]ExecuteResponseFailedExecutionErrorUnion]
	Errors ExecuteResponseUnionErrors `json:"errors"`
	JSON   struct {
		Data   respjson.Field
		Errors respjson.Field
		raw    string
	} `json:"-"`
}

func (u ExecuteResponseUnion) AsSuccessfulExecution() (v ExecuteResponseSuccessfulExecution) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecuteResponseUnion) AsFailedExecution() (v ExecuteResponseFailedExecution) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExecuteResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ExecuteResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecuteResponseUnionData is an implicit subunion of [ExecuteResponseUnion].
// ExecuteResponseUnionData provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecuteResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExecuteResponseFailedExecutionData]
type ExecuteResponseUnionData struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExecuteResponseFailedExecutionData any `json:",inline"`
	// This field is from variant [ExecuteResponseSuccessfulExecutionData].
	Outputs []ExecuteResponseSuccessfulExecutionDataOutputUnion `json:"outputs"`
	// This field is from variant [ExecuteResponseSuccessfulExecutionData].
	SessionID string `json:"session_id"`
	// This field is from variant [ExecuteResponseSuccessfulExecutionData].
	Status string `json:"status"`
	JSON   struct {
		OfExecuteResponseFailedExecutionData respjson.Field
		Outputs                              respjson.Field
		SessionID                            respjson.Field
		Status                               respjson.Field
		raw                                  string
	} `json:"-"`
}

func (r *ExecuteResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecuteResponseUnionErrors is an implicit subunion of [ExecuteResponseUnion].
// ExecuteResponseUnionErrors provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecuteResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExecuteResponseSuccessfulExecutionErrors
// OfExecuteResponseFailedExecutionErrors]
type ExecuteResponseUnionErrors struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExecuteResponseSuccessfulExecutionErrors any `json:",inline"`
	// This field will be present if the value is a
	// [[]ExecuteResponseFailedExecutionErrorUnion] instead of an object.
	OfExecuteResponseFailedExecutionErrors []ExecuteResponseFailedExecutionErrorUnion `json:",inline"`
	JSON                                   struct {
		OfExecuteResponseSuccessfulExecutionErrors respjson.Field
		OfExecuteResponseFailedExecutionErrors     respjson.Field
		raw                                        string
	} `json:"-"`
}

func (r *ExecuteResponseUnionErrors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteResponseSuccessfulExecution struct {
	Data   ExecuteResponseSuccessfulExecutionData `json:"data,required"`
	Errors any                                    `json:"errors,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteResponseSuccessfulExecution) RawJSON() string { return r.JSON.raw }
func (r *ExecuteResponseSuccessfulExecution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteResponseSuccessfulExecutionData struct {
	Outputs []ExecuteResponseSuccessfulExecutionDataOutputUnion `json:"outputs,required"`
	// Identifier of the current session. Used to make follow-up calls.
	SessionID string `json:"session_id,required"`
	// Status of the execution. Currently only supports success.
	//
	// Any of "success".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outputs     respjson.Field
		SessionID   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteResponseSuccessfulExecutionData) RawJSON() string { return r.JSON.raw }
func (r *ExecuteResponseSuccessfulExecutionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecuteResponseSuccessfulExecutionDataOutputUnion contains all possible
// properties and values from
// [ExecuteResponseSuccessfulExecutionDataOutputStreamOutput],
// [ExecuteResponseSuccessfulExecutionDataOutputError],
// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutput].
//
// Use the [ExecuteResponseSuccessfulExecutionDataOutputUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExecuteResponseSuccessfulExecutionDataOutputUnion struct {
	// This field is a union of [string], [string],
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData]
	Data ExecuteResponseSuccessfulExecutionDataOutputUnionData `json:"data"`
	// Any of nil, "error", nil.
	Type string `json:"type"`
	JSON struct {
		Data respjson.Field
		Type respjson.Field
		raw  string
	} `json:"-"`
}

func (u ExecuteResponseSuccessfulExecutionDataOutputUnion) AsStreamOutput() (v ExecuteResponseSuccessfulExecutionDataOutputStreamOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecuteResponseSuccessfulExecutionDataOutputUnion) AsError() (v ExecuteResponseSuccessfulExecutionDataOutputError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecuteResponseSuccessfulExecutionDataOutputUnion) AsDisplayorExecuteOutput() (v ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExecuteResponseSuccessfulExecutionDataOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ExecuteResponseSuccessfulExecutionDataOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecuteResponseSuccessfulExecutionDataOutputUnionData is an implicit subunion of
// [ExecuteResponseSuccessfulExecutionDataOutputUnion].
// ExecuteResponseSuccessfulExecutionDataOutputUnionData provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecuteResponseSuccessfulExecutionDataOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type ExecuteResponseSuccessfulExecutionDataOutputUnionData struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ApplicationGeoJson map[string]any `json:"application/geo+json"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ApplicationJavascript string `json:"application/javascript"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ApplicationJson map[string]any `json:"application/json"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ApplicationPdf string `json:"application/pdf"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ApplicationVndVegaV5Json map[string]any `json:"application/vnd.vega.v5+json"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ApplicationVndVegaliteV4Json map[string]any `json:"application/vnd.vegalite.v4+json"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ImageGif string `json:"image/gif"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ImageJpeg string `json:"image/jpeg"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ImagePng string `json:"image/png"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	ImageSvgXml string `json:"image/svg+xml"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	TextHTML string `json:"text/html"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	TextLatex string `json:"text/latex"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	TextMarkdown string `json:"text/markdown"`
	// This field is from variant
	// [ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData].
	TextPlain string `json:"text/plain"`
	JSON      struct {
		OfString                     respjson.Field
		ApplicationGeoJson           respjson.Field
		ApplicationJavascript        respjson.Field
		ApplicationJson              respjson.Field
		ApplicationPdf               respjson.Field
		ApplicationVndVegaV5Json     respjson.Field
		ApplicationVndVegaliteV4Json respjson.Field
		ImageGif                     respjson.Field
		ImageJpeg                    respjson.Field
		ImagePng                     respjson.Field
		ImageSvgXml                  respjson.Field
		TextHTML                     respjson.Field
		TextLatex                    respjson.Field
		TextMarkdown                 respjson.Field
		TextPlain                    respjson.Field
		raw                          string
	} `json:"-"`
}

func (r *ExecuteResponseSuccessfulExecutionDataOutputUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Outputs that were printed to stdout or stderr
type ExecuteResponseSuccessfulExecutionDataOutputStreamOutput struct {
	Data string `json:"data,required"`
	// Any of "stdout", "stderr".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteResponseSuccessfulExecutionDataOutputStreamOutput) RawJSON() string { return r.JSON.raw }
func (r *ExecuteResponseSuccessfulExecutionDataOutputStreamOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Errors and exceptions that occurred. If this output type is present, your code
// did not execute successfully.
type ExecuteResponseSuccessfulExecutionDataOutputError struct {
	Data string         `json:"data,required"`
	Type constant.Error `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteResponseSuccessfulExecutionDataOutputError) RawJSON() string { return r.JSON.raw }
func (r *ExecuteResponseSuccessfulExecutionDataOutputError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutput struct {
	Data ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData `json:"data,required"`
	// Any of "display_data", "execute_result".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData struct {
	ApplicationGeoJson           map[string]any `json:"application/geo+json"`
	ApplicationJavascript        string         `json:"application/javascript"`
	ApplicationJson              map[string]any `json:"application/json"`
	ApplicationPdf               string         `json:"application/pdf" format:"byte"`
	ApplicationVndVegaV5Json     map[string]any `json:"application/vnd.vega.v5+json"`
	ApplicationVndVegaliteV4Json map[string]any `json:"application/vnd.vegalite.v4+json"`
	ImageGif                     string         `json:"image/gif" format:"byte"`
	ImageJpeg                    string         `json:"image/jpeg" format:"byte"`
	ImagePng                     string         `json:"image/png" format:"byte"`
	ImageSvgXml                  string         `json:"image/svg+xml"`
	TextHTML                     string         `json:"text/html"`
	TextLatex                    string         `json:"text/latex"`
	TextMarkdown                 string         `json:"text/markdown"`
	TextPlain                    string         `json:"text/plain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationGeoJson           respjson.Field
		ApplicationJavascript        respjson.Field
		ApplicationJson              respjson.Field
		ApplicationPdf               respjson.Field
		ApplicationVndVegaV5Json     respjson.Field
		ApplicationVndVegaliteV4Json respjson.Field
		ImageGif                     respjson.Field
		ImageJpeg                    respjson.Field
		ImagePng                     respjson.Field
		ImageSvgXml                  respjson.Field
		TextHTML                     respjson.Field
		TextLatex                    respjson.Field
		TextMarkdown                 respjson.Field
		TextPlain                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecuteResponseSuccessfulExecutionDataOutputDisplayorExecuteOutputData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteResponseFailedExecution struct {
	Data   any                                        `json:"data,required"`
	Errors []ExecuteResponseFailedExecutionErrorUnion `json:"errors,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteResponseFailedExecution) RawJSON() string { return r.JSON.raw }
func (r *ExecuteResponseFailedExecution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecuteResponseFailedExecutionErrorUnion contains all possible properties and
// values from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfExecuteResponseFailedExecutionErrorMapItem]
type ExecuteResponseFailedExecutionErrorUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfExecuteResponseFailedExecutionErrorMapItem any `json:",inline"`
	JSON                                         struct {
		OfString                                     respjson.Field
		OfExecuteResponseFailedExecutionErrorMapItem respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u ExecuteResponseFailedExecutionErrorUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecuteResponseFailedExecutionErrorUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExecuteResponseFailedExecutionErrorUnion) RawJSON() string { return u.JSON.raw }

func (r *ExecuteResponseFailedExecutionErrorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CodeInterpreterExecuteParams struct {
	// Code snippet to execute.
	Code string `json:"code,required"`
	// Programming language for the code to execute. Currently only supports Python,
	// but more will be added.
	//
	// Any of "python".
	Language CodeInterpreterExecuteParamsLanguage `json:"language,omitzero,required"`
	// Identifier of the current session. Used to make follow-up calls. Requests will
	// return an error if the session does not belong to the caller or has expired.
	SessionID param.Opt[string] `json:"session_id,omitzero"`
	// Files to upload to the session. If present, files will be uploaded before
	// executing the given code.
	Files []CodeInterpreterExecuteParamsFile `json:"files,omitzero"`
	paramObj
}

func (r CodeInterpreterExecuteParams) MarshalJSON() (data []byte, err error) {
	type shadow CodeInterpreterExecuteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CodeInterpreterExecuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Programming language for the code to execute. Currently only supports Python,
// but more will be added.
type CodeInterpreterExecuteParamsLanguage string

const (
	CodeInterpreterExecuteParamsLanguagePython CodeInterpreterExecuteParamsLanguage = "python"
)

// The properties Content, Encoding, Name are required.
type CodeInterpreterExecuteParamsFile struct {
	Content string `json:"content,required"`
	// Encoding of the file content. Use `string` for text files such as code, and
	// `base64` for binary files, such as images.
	//
	// Any of "string", "base64".
	Encoding string `json:"encoding,omitzero,required"`
	Name     string `json:"name,required"`
	paramObj
}

func (r CodeInterpreterExecuteParamsFile) MarshalJSON() (data []byte, err error) {
	type shadow CodeInterpreterExecuteParamsFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CodeInterpreterExecuteParamsFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CodeInterpreterExecuteParamsFile](
		"encoding", "string", "base64",
	)
}
