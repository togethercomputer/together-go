// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"reflect"

	"github.com/tidwall/gjson"
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/shared"
)

// CodeInterpreterService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCodeInterpreterService] method instead.
type CodeInterpreterService struct {
	Options  []option.RequestOption
	Sessions *CodeInterpreterSessionService
}

// NewCodeInterpreterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCodeInterpreterService(opts ...option.RequestOption) (r *CodeInterpreterService) {
	r = &CodeInterpreterService{}
	r.Options = opts
	r.Sessions = NewCodeInterpreterSessionService(opts...)
	return
}

// Executes the given code snippet and returns the output. Without a session_id, a
// new session will be created to run the code. If you do pass in a valid
// session_id, the code will be run in that session. This is useful for running
// multiple code snippets in the same environment, because dependencies and similar
// things are persisted between calls to the same session.
func (r *CodeInterpreterService) Execute(ctx context.Context, body CodeInterpreterExecuteParams, opts ...option.RequestOption) (res *ExecuteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "tci/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The result of the execution. If successful, `data` contains the result and
// `errors` will be null. If unsuccessful, `data` will be null and `errors` will
// contain the errors.
type ExecuteResponse struct {
	// This field can have the runtime type of
	// [ExecuteResponseSuccessfulExecutionData], [interface{}].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [interface{}],
	// [[]ExecuteResponseFailedExecutionErrorsUnion].
	Errors interface{}         `json:"errors,required"`
	JSON   executeResponseJSON `json:"-"`
	union  ExecuteResponseUnion
}

// executeResponseJSON contains the JSON metadata for the struct [ExecuteResponse]
type executeResponseJSON struct {
	Data        apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r executeResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ExecuteResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ExecuteResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ExecuteResponseSuccessfulExecution],
// [ExecuteResponseFailedExecution].
func (r ExecuteResponse) AsUnion() ExecuteResponseUnion {
	return r.union
}

// The result of the execution. If successful, `data` contains the result and
// `errors` will be null. If unsuccessful, `data` will be null and `errors` will
// contain the errors.
//
// Union satisfied by [ExecuteResponseSuccessfulExecution] or
// [ExecuteResponseFailedExecution].
type ExecuteResponseUnion interface {
	implementsExecuteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExecuteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExecuteResponseSuccessfulExecution{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExecuteResponseFailedExecution{}),
		},
	)
}

type ExecuteResponseSuccessfulExecution struct {
	Data   ExecuteResponseSuccessfulExecutionData `json:"data,required"`
	Errors interface{}                            `json:"errors,required,nullable"`
	JSON   executeResponseSuccessfulExecutionJSON `json:"-"`
}

// executeResponseSuccessfulExecutionJSON contains the JSON metadata for the struct
// [ExecuteResponseSuccessfulExecution]
type executeResponseSuccessfulExecutionJSON struct {
	Data        apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteResponseSuccessfulExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeResponseSuccessfulExecutionJSON) RawJSON() string {
	return r.raw
}

func (r ExecuteResponseSuccessfulExecution) implementsExecuteResponse() {}

type ExecuteResponseSuccessfulExecutionData struct {
	Outputs []ExecuteResponseSuccessfulExecutionDataOutput `json:"outputs,required"`
	// Identifier of the current session. Used to make follow-up calls.
	SessionID string                                     `json:"session_id,required"`
	JSON      executeResponseSuccessfulExecutionDataJSON `json:"-"`
}

// executeResponseSuccessfulExecutionDataJSON contains the JSON metadata for the
// struct [ExecuteResponseSuccessfulExecutionData]
type executeResponseSuccessfulExecutionDataJSON struct {
	Outputs     apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteResponseSuccessfulExecutionData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeResponseSuccessfulExecutionDataJSON) RawJSON() string {
	return r.raw
}

// Outputs that were printed to stdout or stderr
type ExecuteResponseSuccessfulExecutionDataOutput struct {
	// This field can have the runtime type of [string],
	// [ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputData].
	Data  interface{}                                       `json:"data,required"`
	Type  ExecuteResponseSuccessfulExecutionDataOutputsType `json:"type,required"`
	JSON  executeResponseSuccessfulExecutionDataOutputJSON  `json:"-"`
	union ExecuteResponseSuccessfulExecutionDataOutputsUnion
}

// executeResponseSuccessfulExecutionDataOutputJSON contains the JSON metadata for
// the struct [ExecuteResponseSuccessfulExecutionDataOutput]
type executeResponseSuccessfulExecutionDataOutputJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r executeResponseSuccessfulExecutionDataOutputJSON) RawJSON() string {
	return r.raw
}

func (r *ExecuteResponseSuccessfulExecutionDataOutput) UnmarshalJSON(data []byte) (err error) {
	*r = ExecuteResponseSuccessfulExecutionDataOutput{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ExecuteResponseSuccessfulExecutionDataOutputsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ExecuteResponseSuccessfulExecutionDataOutputsStreamOutput],
// [ExecuteResponseSuccessfulExecutionDataOutputsError],
// [ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutput].
func (r ExecuteResponseSuccessfulExecutionDataOutput) AsUnion() ExecuteResponseSuccessfulExecutionDataOutputsUnion {
	return r.union
}

// Outputs that were printed to stdout or stderr
//
// Union satisfied by [ExecuteResponseSuccessfulExecutionDataOutputsStreamOutput],
// [ExecuteResponseSuccessfulExecutionDataOutputsError] or
// [ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutput].
type ExecuteResponseSuccessfulExecutionDataOutputsUnion interface {
	implementsExecuteResponseSuccessfulExecutionDataOutput()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExecuteResponseSuccessfulExecutionDataOutputsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteResponseSuccessfulExecutionDataOutputsStreamOutput{}),
			DiscriminatorValue: "stdout",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteResponseSuccessfulExecutionDataOutputsStreamOutput{}),
			DiscriminatorValue: "stderr",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteResponseSuccessfulExecutionDataOutputsError{}),
			DiscriminatorValue: "error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutput{}),
			DiscriminatorValue: "display_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutput{}),
			DiscriminatorValue: "execute_result",
		},
	)
}

// Outputs that were printed to stdout or stderr
type ExecuteResponseSuccessfulExecutionDataOutputsStreamOutput struct {
	Data string                                                        `json:"data,required"`
	Type ExecuteResponseSuccessfulExecutionDataOutputsStreamOutputType `json:"type,required"`
	JSON executeResponseSuccessfulExecutionDataOutputsStreamOutputJSON `json:"-"`
}

// executeResponseSuccessfulExecutionDataOutputsStreamOutputJSON contains the JSON
// metadata for the struct
// [ExecuteResponseSuccessfulExecutionDataOutputsStreamOutput]
type executeResponseSuccessfulExecutionDataOutputsStreamOutputJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteResponseSuccessfulExecutionDataOutputsStreamOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeResponseSuccessfulExecutionDataOutputsStreamOutputJSON) RawJSON() string {
	return r.raw
}

func (r ExecuteResponseSuccessfulExecutionDataOutputsStreamOutput) implementsExecuteResponseSuccessfulExecutionDataOutput() {
}

type ExecuteResponseSuccessfulExecutionDataOutputsStreamOutputType string

const (
	ExecuteResponseSuccessfulExecutionDataOutputsStreamOutputTypeStdout ExecuteResponseSuccessfulExecutionDataOutputsStreamOutputType = "stdout"
	ExecuteResponseSuccessfulExecutionDataOutputsStreamOutputTypeStderr ExecuteResponseSuccessfulExecutionDataOutputsStreamOutputType = "stderr"
)

func (r ExecuteResponseSuccessfulExecutionDataOutputsStreamOutputType) IsKnown() bool {
	switch r {
	case ExecuteResponseSuccessfulExecutionDataOutputsStreamOutputTypeStdout, ExecuteResponseSuccessfulExecutionDataOutputsStreamOutputTypeStderr:
		return true
	}
	return false
}

// Errors and exceptions that occurred. If this output type is present, your code
// did not execute successfully.
type ExecuteResponseSuccessfulExecutionDataOutputsError struct {
	Data string                                                 `json:"data,required"`
	Type ExecuteResponseSuccessfulExecutionDataOutputsErrorType `json:"type,required"`
	JSON executeResponseSuccessfulExecutionDataOutputsErrorJSON `json:"-"`
}

// executeResponseSuccessfulExecutionDataOutputsErrorJSON contains the JSON
// metadata for the struct [ExecuteResponseSuccessfulExecutionDataOutputsError]
type executeResponseSuccessfulExecutionDataOutputsErrorJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteResponseSuccessfulExecutionDataOutputsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeResponseSuccessfulExecutionDataOutputsErrorJSON) RawJSON() string {
	return r.raw
}

func (r ExecuteResponseSuccessfulExecutionDataOutputsError) implementsExecuteResponseSuccessfulExecutionDataOutput() {
}

type ExecuteResponseSuccessfulExecutionDataOutputsErrorType string

const (
	ExecuteResponseSuccessfulExecutionDataOutputsErrorTypeError ExecuteResponseSuccessfulExecutionDataOutputsErrorType = "error"
)

func (r ExecuteResponseSuccessfulExecutionDataOutputsErrorType) IsKnown() bool {
	switch r {
	case ExecuteResponseSuccessfulExecutionDataOutputsErrorTypeError:
		return true
	}
	return false
}

type ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutput struct {
	Data ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputData `json:"data,required"`
	Type ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputType `json:"type,required"`
	JSON executeResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputJSON `json:"-"`
}

// executeResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputJSON contains
// the JSON metadata for the struct
// [ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutput]
type executeResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputJSON) RawJSON() string {
	return r.raw
}

func (r ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutput) implementsExecuteResponseSuccessfulExecutionDataOutput() {
}

type ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputData struct {
	ApplicationGeoJson           map[string]interface{}                                                      `json:"application/geo+json"`
	ApplicationJavascript        string                                                                      `json:"application/javascript"`
	ApplicationJson              map[string]interface{}                                                      `json:"application/json"`
	ApplicationPdf               string                                                                      `json:"application/pdf" format:"byte"`
	ApplicationVndVegaV5Json     map[string]interface{}                                                      `json:"application/vnd.vega.v5+json"`
	ApplicationVndVegaliteV4Json map[string]interface{}                                                      `json:"application/vnd.vegalite.v4+json"`
	ImageGif                     string                                                                      `json:"image/gif" format:"byte"`
	ImageJpeg                    string                                                                      `json:"image/jpeg" format:"byte"`
	ImagePng                     string                                                                      `json:"image/png" format:"byte"`
	ImageSvgXml                  string                                                                      `json:"image/svg+xml"`
	TextHTML                     string                                                                      `json:"text/html"`
	TextLatex                    string                                                                      `json:"text/latex"`
	TextMarkdown                 string                                                                      `json:"text/markdown"`
	TextPlain                    string                                                                      `json:"text/plain"`
	JSON                         executeResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputDataJSON `json:"-"`
}

// executeResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputDataJSON
// contains the JSON metadata for the struct
// [ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputData]
type executeResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputDataJSON struct {
	ApplicationGeoJson           apijson.Field
	ApplicationJavascript        apijson.Field
	ApplicationJson              apijson.Field
	ApplicationPdf               apijson.Field
	ApplicationVndVegaV5Json     apijson.Field
	ApplicationVndVegaliteV4Json apijson.Field
	ImageGif                     apijson.Field
	ImageJpeg                    apijson.Field
	ImagePng                     apijson.Field
	ImageSvgXml                  apijson.Field
	TextHTML                     apijson.Field
	TextLatex                    apijson.Field
	TextMarkdown                 apijson.Field
	TextPlain                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputDataJSON) RawJSON() string {
	return r.raw
}

type ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputType string

const (
	ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputTypeDisplayData   ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputType = "display_data"
	ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputTypeExecuteResult ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputType = "execute_result"
)

func (r ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputType) IsKnown() bool {
	switch r {
	case ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputTypeDisplayData, ExecuteResponseSuccessfulExecutionDataOutputsDisplayorExecuteOutputTypeExecuteResult:
		return true
	}
	return false
}

type ExecuteResponseSuccessfulExecutionDataOutputsType string

const (
	ExecuteResponseSuccessfulExecutionDataOutputsTypeStdout        ExecuteResponseSuccessfulExecutionDataOutputsType = "stdout"
	ExecuteResponseSuccessfulExecutionDataOutputsTypeStderr        ExecuteResponseSuccessfulExecutionDataOutputsType = "stderr"
	ExecuteResponseSuccessfulExecutionDataOutputsTypeError         ExecuteResponseSuccessfulExecutionDataOutputsType = "error"
	ExecuteResponseSuccessfulExecutionDataOutputsTypeDisplayData   ExecuteResponseSuccessfulExecutionDataOutputsType = "display_data"
	ExecuteResponseSuccessfulExecutionDataOutputsTypeExecuteResult ExecuteResponseSuccessfulExecutionDataOutputsType = "execute_result"
)

func (r ExecuteResponseSuccessfulExecutionDataOutputsType) IsKnown() bool {
	switch r {
	case ExecuteResponseSuccessfulExecutionDataOutputsTypeStdout, ExecuteResponseSuccessfulExecutionDataOutputsTypeStderr, ExecuteResponseSuccessfulExecutionDataOutputsTypeError, ExecuteResponseSuccessfulExecutionDataOutputsTypeDisplayData, ExecuteResponseSuccessfulExecutionDataOutputsTypeExecuteResult:
		return true
	}
	return false
}

type ExecuteResponseFailedExecution struct {
	Data   interface{}                                 `json:"data,required,nullable"`
	Errors []ExecuteResponseFailedExecutionErrorsUnion `json:"errors,required"`
	JSON   executeResponseFailedExecutionJSON          `json:"-"`
}

// executeResponseFailedExecutionJSON contains the JSON metadata for the struct
// [ExecuteResponseFailedExecution]
type executeResponseFailedExecutionJSON struct {
	Data        apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteResponseFailedExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeResponseFailedExecutionJSON) RawJSON() string {
	return r.raw
}

func (r ExecuteResponseFailedExecution) implementsExecuteResponse() {}

// Union satisfied by [shared.UnionString] or
// [ExecuteResponseFailedExecutionErrorsMap].
type ExecuteResponseFailedExecutionErrorsUnion interface {
	ImplementsExecuteResponseFailedExecutionErrorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExecuteResponseFailedExecutionErrorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ExecuteResponseFailedExecutionErrorsMap map[string]interface{}

func (r ExecuteResponseFailedExecutionErrorsMap) ImplementsExecuteResponseFailedExecutionErrorsUnion() {
}

type CodeInterpreterExecuteParams struct {
	// Code snippet to execute.
	Code param.Field[string] `json:"code,required"`
	// Programming language for the code to execute. Currently only supports Python,
	// but more will be added.
	Language param.Field[CodeInterpreterExecuteParamsLanguage] `json:"language,required"`
	// Identifier of the current session. Used to make follow-up calls. Requests will
	// return an error if the session does not belong to the caller or has expired.
	SessionID param.Field[string] `json:"session_id"`
}

func (r CodeInterpreterExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Programming language for the code to execute. Currently only supports Python,
// but more will be added.
type CodeInterpreterExecuteParamsLanguage string

const (
	CodeInterpreterExecuteParamsLanguagePython CodeInterpreterExecuteParamsLanguage = "python"
)

func (r CodeInterpreterExecuteParamsLanguage) IsKnown() bool {
	switch r {
	case CodeInterpreterExecuteParamsLanguagePython:
		return true
	}
	return false
}
