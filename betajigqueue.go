// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

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

// Poll the current status of a previously submitted job. Provide the request_id
// and model as query parameters.
func (r *BetaJigQueueService) Get(ctx context.Context, query BetaJigQueueGetParams, opts ...option.RequestOption) (res *BetaJigQueueGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cancel a pending job. Only jobs in pending status can be canceled. Running jobs
// cannot be stopped. Returns the job status after the attempt. If the job is not
// pending, returns 409 with the current status unchanged.
func (r *BetaJigQueueService) Cancel(ctx context.Context, body BetaJigQueueCancelParams, opts ...option.RequestOption) (res *BetaJigQueueCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/cancel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the current queue statistics for a model, including pending and running job
// counts.
func (r *BetaJigQueueService) Metrics(ctx context.Context, query BetaJigQueueMetricsParams, opts ...option.RequestOption) (res *BetaJigQueueMetricsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Submit a new job to the queue for asynchronous processing. Jobs are processed in
// strict priority order (higher priority first, FIFO within the same priority).
// Returns a request ID that can be used to poll status or cancel the job.
func (r *BetaJigQueueService) Submit(ctx context.Context, body BetaJigQueueSubmitParams, opts ...option.RequestOption) (res *BetaJigQueueSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queue/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BetaJigQueueGetResponse struct {
	// Model identifier the job was submitted to
	Model string `json:"model,required"`
	// The request ID that was returned from the submit endpoint
	RequestID string `json:"request_id,required"`
	// Current job status. Transitions: pending → running → done/failed. A pending job
	// may also be canceled.
	//
	// Any of "pending", "running", "done", "failed", "canceled".
	Status BetaJigQueueGetResponseStatus `json:"status,required"`
	// Timestamp when a worker claimed the job
	ClaimedAt time.Time `json:"claimed_at" format:"date-time"`
	// Timestamp when the job was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp when the job completed (done or failed)
	DoneAt time.Time `json:"done_at" format:"date-time"`
	// Job metadata. Contains keys from the submit request, plus any modifications from
	// the model or system (e.g. progress, retry history).
	Info map[string]any `json:"info"`
	// Freeform model input, as submitted
	Inputs map[string]any `json:"inputs"`
	// Freeform model output, populated when the job reaches done status. Contents are
	// model-specific.
	Outputs map[string]any `json:"outputs"`
	// Job priority. Higher values are processed first.
	Priority int64 `json:"priority"`
	// Number of times this job has been retried. Workers set a claim timeout and must
	// send periodic status updates to keep the job alive. If no update is received
	// within the timeout, the job is returned to the queue and retried. After 3
	// retries the job is permanently failed. Jobs explicitly failed by the model are
	// not retried.
	Retries int64 `json:"retries"`
	// Non-fatal messages about the request (e.g. deprecation notices)
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Model       respjson.Field
		RequestID   respjson.Field
		Status      respjson.Field
		ClaimedAt   respjson.Field
		CreatedAt   respjson.Field
		DoneAt      respjson.Field
		Info        respjson.Field
		Inputs      respjson.Field
		Outputs     respjson.Field
		Priority    respjson.Field
		Retries     respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigQueueGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigQueueGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current job status. Transitions: pending → running → done/failed. A pending job
// may also be canceled.
type BetaJigQueueGetResponseStatus string

const (
	BetaJigQueueGetResponseStatusPending  BetaJigQueueGetResponseStatus = "pending"
	BetaJigQueueGetResponseStatusRunning  BetaJigQueueGetResponseStatus = "running"
	BetaJigQueueGetResponseStatusDone     BetaJigQueueGetResponseStatus = "done"
	BetaJigQueueGetResponseStatusFailed   BetaJigQueueGetResponseStatus = "failed"
	BetaJigQueueGetResponseStatusCanceled BetaJigQueueGetResponseStatus = "canceled"
)

type BetaJigQueueCancelResponse struct {
	// Job status after the cancel attempt. Only pending jobs can be canceled. If the
	// job is already running, done, or failed, the status is returned unchanged.
	//
	// Any of "canceled", "running", "done", "failed".
	Status BetaJigQueueCancelResponseStatus `json:"status,required"`
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

// Job status after the cancel attempt. Only pending jobs can be canceled. If the
// job is already running, done, or failed, the status is returned unchanged.
type BetaJigQueueCancelResponseStatus string

const (
	BetaJigQueueCancelResponseStatusCanceled BetaJigQueueCancelResponseStatus = "canceled"
	BetaJigQueueCancelResponseStatusRunning  BetaJigQueueCancelResponseStatus = "running"
	BetaJigQueueCancelResponseStatusDone     BetaJigQueueCancelResponseStatus = "done"
	BetaJigQueueCancelResponseStatusFailed   BetaJigQueueCancelResponseStatus = "failed"
)

type BetaJigQueueMetricsResponse struct {
	// Number of jobs currently being processed
	MessagesRunning int64 `json:"messages_running,required"`
	// Number of jobs waiting to be claimed by a worker
	MessagesWaiting int64 `json:"messages_waiting,required"`
	// Total number of active jobs (waiting + running)
	TotalJobs int64 `json:"total_jobs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessagesRunning respjson.Field
		MessagesWaiting respjson.Field
		TotalJobs       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigQueueMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigQueueMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigQueueSubmitResponse struct {
	Error BetaJigQueueSubmitResponseError `json:"error"`
	// Unique identifier for the submitted job. Use this to poll status or cancel.
	RequestID string `json:"requestId"`
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
	// Machine-readable error code
	Code string `json:"code"`
	// Human-readable error message
	Message string `json:"message"`
	// The parameter that caused the error, if applicable
	Param string `json:"param"`
	// Error category (e.g. "invalid_request_error", "not_found_error")
	Type string `json:"type"`
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

type BetaJigQueueGetParams struct {
	// Model name the job was submitted to
	Model string `query:"model,required" json:"-"`
	// Request ID returned from the submit endpoint
	RequestID string `query:"request_id,required" json:"-"`
	paramObj
}

// URLQuery serializes [BetaJigQueueGetParams]'s query parameters as `url.Values`.
func (r BetaJigQueueGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaJigQueueCancelParams struct {
	// Model identifier the job was submitted to
	Model string `json:"model,required"`
	// The request ID returned from the submit endpoint
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

type BetaJigQueueMetricsParams struct {
	// Model name to get metrics for
	Model string `query:"model,required" json:"-"`
	paramObj
}

// URLQuery serializes [BetaJigQueueMetricsParams]'s query parameters as
// `url.Values`.
func (r BetaJigQueueMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaJigQueueSubmitParams struct {
	// Required model identifier
	Model string `json:"model,required"`
	// Freeform model input. Passed unchanged to the model. Contents are
	// model-specific.
	Payload map[string]any `json:"payload,omitzero,required"`
	// Job priority. Higher values are processed first (strict priority ordering). Jobs
	// with equal priority are processed in submission order (FIFO).
	Priority param.Opt[int64] `json:"priority,omitzero"`
	// Arbitrary JSON metadata stored with the job and returned in status responses.
	// The model and system may add or update keys during processing.
	Info map[string]any `json:"info,omitzero"`
	paramObj
}

func (r BetaJigQueueSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigQueueSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigQueueSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
