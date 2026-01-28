// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaJigQueueService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaJigQueueService] method instead.
type BetaJigQueueService struct {
	Options []option.RequestOption
}

// NewBetaJigQueueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaJigQueueService(opts ...option.RequestOption) (r BetaJigQueueService) {
	r = BetaJigQueueService{}
	r.Options = opts
	return
}

// Cancel a pending or running job. Returns the job status after the cancellation
// attempt.
func (r *BetaJigQueueService) Cancel(ctx context.Context, body BetaJigQueueCancelParams, opts ...option.RequestOption) (res *BetaJigQueueCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/cancel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the current queue statistics including pending and running job counts.
func (r *BetaJigQueueService) GetMetrics(ctx context.Context, query BetaJigQueueGetMetricsParams, opts ...option.RequestOption) (res *BetaJigQueueGetMetricsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check the status of a job using request_id and model query parameters.
func (r *BetaJigQueueService) GetStatus(ctx context.Context, query BetaJigQueueGetStatusParams, opts ...option.RequestOption) (res *BetaJigQueueGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Submit a new job to the queue. Returns a request ID that can be used to check
// status.
func (r *BetaJigQueueService) Submit(ctx context.Context, body BetaJigQueueSubmitParams, opts ...option.RequestOption) (res *BetaJigQueueSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BetaJigQueueCancelResponse struct {
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigQueueCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigQueueCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigQueueGetMetricsResponse map[string]any

type BetaJigQueueGetStatusResponse struct {
	ClaimedAt string         `json:"claimed_at"`
	CreatedAt string         `json:"created_at"`
	DoneAt    string         `json:"done_at"`
	Info      map[string]any `json:"info"`
	Inputs    map[string]any `json:"inputs"`
	Model     string         `json:"model"`
	Outputs   map[string]any `json:"outputs"`
	// Additional fields for test compatibility
	Priority  int64  `json:"priority"`
	RequestID string `json:"request_id"`
	Retries   int64  `json:"retries"`
	// this should be the enum, but isn't for backwards compatability
	Status   string   `json:"status"`
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClaimedAt   respjson.Field
		CreatedAt   respjson.Field
		DoneAt      respjson.Field
		Info        respjson.Field
		Inputs      respjson.Field
		Model       respjson.Field
		Outputs     respjson.Field
		Priority    respjson.Field
		RequestID   respjson.Field
		Retries     respjson.Field
		Status      respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigQueueGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigQueueGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigQueueSubmitResponse struct {
	Error     BetaJigQueueSubmitResponseError `json:"error"`
	RequestID string                          `json:"requestId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigQueueSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigQueueSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigQueueSubmitResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Param       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigQueueSubmitResponseError) RawJSON() string { return r.JSON.raw }
func (r *BetaJigQueueSubmitResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigQueueCancelParams struct {
	Model     string `json:"model,required"`
	RequestID string `json:"request_id,required"`
	paramObj
}

func (r BetaJigQueueCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigQueueCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigQueueCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigQueueGetMetricsParams struct {
	// Model name to get metrics for
	Model string `query:"model,required" json:"-"`
	paramObj
}

// URLQuery serializes [BetaJigQueueGetMetricsParams]'s query parameters as
// `url.Values`.
func (r BetaJigQueueGetMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaJigQueueGetStatusParams struct {
	// Model name
	Model string `query:"model,required" json:"-"`
	// Request ID
	RequestID string `query:"request_id,required" json:"-"`
	paramObj
}

// URLQuery serializes [BetaJigQueueGetStatusParams]'s query parameters as
// `url.Values`.
func (r BetaJigQueueGetStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaJigQueueSubmitParams struct {
	// Required model identifier
	Model    string           `json:"model,required"`
	Payload  map[string]any   `json:"payload,omitzero,required"`
	Priority param.Opt[int64] `json:"priority,omitzero"`
	Info     map[string]any   `json:"info,omitzero"`
	paramObj
}

func (r BetaJigQueueSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigQueueSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigQueueSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
