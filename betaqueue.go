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

// BetaQueueService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaQueueService] method instead.
type BetaQueueService struct {
	Options []option.RequestOption
}

// NewBetaQueueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaQueueService(opts ...option.RequestOption) (r BetaQueueService) {
	r = BetaQueueService{}
	r.Options = opts
	return
}

// Check the status of a job using request_id and model query parameters.
func (r *BetaQueueService) Get(ctx context.Context, query BetaQueueGetParams, opts ...option.RequestOption) (res *BetaQueueGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cancel a pending or running job. Returns the job status after the cancellation
// attempt.
func (r *BetaQueueService) Cancel(ctx context.Context, body BetaQueueCancelParams, opts ...option.RequestOption) (res *BetaQueueCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/cancel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the current queue statistics including pending and running job counts.
func (r *BetaQueueService) Metrics(ctx context.Context, query BetaQueueMetricsParams, opts ...option.RequestOption) (res *BetaQueueMetricsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Submit a new job to the queue. Returns a request ID that can be used to check
// status.
func (r *BetaQueueService) Submit(ctx context.Context, body BetaQueueSubmitParams, opts ...option.RequestOption) (res *BetaQueueSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BetaQueueGetResponse struct {
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
func (r BetaQueueGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaQueueGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaQueueCancelResponse struct {
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaQueueCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaQueueCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaQueueMetricsResponse map[string]any

type BetaQueueSubmitResponse struct {
	Error     BetaQueueSubmitResponseError `json:"error"`
	RequestID string                       `json:"requestId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		RequestID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaQueueSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaQueueSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaQueueSubmitResponseError struct {
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
func (r BetaQueueSubmitResponseError) RawJSON() string { return r.JSON.raw }
func (r *BetaQueueSubmitResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaQueueGetParams struct {
	// Model name
	Model string `query:"model,required" json:"-"`
	// Request ID
	RequestID string `query:"request_id,required" json:"-"`
	paramObj
}

// URLQuery serializes [BetaQueueGetParams]'s query parameters as `url.Values`.
func (r BetaQueueGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaQueueCancelParams struct {
	Model     string `json:"model,required"`
	RequestID string `json:"request_id,required"`
	paramObj
}

func (r BetaQueueCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaQueueCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaQueueCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaQueueMetricsParams struct {
	// Model name to get metrics for
	Model string `query:"model,required" json:"-"`
	paramObj
}

// URLQuery serializes [BetaQueueMetricsParams]'s query parameters as `url.Values`.
func (r BetaQueueMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaQueueSubmitParams struct {
	// Required model identifier
	Model    string           `json:"model,required"`
	Payload  map[string]any   `json:"payload,omitzero,required"`
	Priority param.Opt[int64] `json:"priority,omitzero"`
	Info     map[string]any   `json:"info,omitzero"`
	paramObj
}

func (r BetaQueueSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaQueueSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaQueueSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
