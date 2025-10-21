// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BatchService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchService] method instead.
type BatchService struct {
	Options []option.RequestOption
}

// NewBatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBatchService(opts ...option.RequestOption) (r BatchService) {
	r = BatchService{}
	r.Options = opts
	return
}

// Create a new batch job with the given input file and endpoint
func (r *BatchService) New(ctx context.Context, body BatchNewParams, opts ...option.RequestOption) (res *BatchNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details of a batch job by ID
func (r *BatchService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BatchGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("batches/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all batch jobs for the authenticated user
func (r *BatchService) List(ctx context.Context, opts ...option.RequestOption) (res *[]BatchListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BatchNewResponse struct {
	Job     BatchNewResponseJob `json:"job"`
	Warning string              `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		Warning     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchNewResponseJob struct {
	ID          string    `json:"id" format:"uuid"`
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	Endpoint    string    `json:"endpoint"`
	Error       string    `json:"error"`
	ErrorFileID string    `json:"error_file_id"`
	// Size of input file in bytes
	FileSizeBytes int64     `json:"file_size_bytes"`
	InputFileID   string    `json:"input_file_id"`
	JobDeadline   time.Time `json:"job_deadline" format:"date-time"`
	// Model used for processing requests
	ModelID      string `json:"model_id"`
	OutputFileID string `json:"output_file_id"`
	// Completion progress (0.0 to 100)
	Progress float64 `json:"progress"`
	// Current status of the batch job
	//
	// Any of "VALIDATING", "IN_PROGRESS", "COMPLETED", "FAILED", "EXPIRED",
	// "CANCELLED".
	Status string `json:"status"`
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CompletedAt   respjson.Field
		CreatedAt     respjson.Field
		Endpoint      respjson.Field
		Error         respjson.Field
		ErrorFileID   respjson.Field
		FileSizeBytes respjson.Field
		InputFileID   respjson.Field
		JobDeadline   respjson.Field
		ModelID       respjson.Field
		OutputFileID  respjson.Field
		Progress      respjson.Field
		Status        respjson.Field
		UserID        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseJob) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchGetResponse struct {
	ID          string    `json:"id" format:"uuid"`
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	Endpoint    string    `json:"endpoint"`
	Error       string    `json:"error"`
	ErrorFileID string    `json:"error_file_id"`
	// Size of input file in bytes
	FileSizeBytes int64     `json:"file_size_bytes"`
	InputFileID   string    `json:"input_file_id"`
	JobDeadline   time.Time `json:"job_deadline" format:"date-time"`
	// Model used for processing requests
	ModelID      string `json:"model_id"`
	OutputFileID string `json:"output_file_id"`
	// Completion progress (0.0 to 100)
	Progress float64 `json:"progress"`
	// Current status of the batch job
	//
	// Any of "VALIDATING", "IN_PROGRESS", "COMPLETED", "FAILED", "EXPIRED",
	// "CANCELLED".
	Status BatchGetResponseStatus `json:"status"`
	UserID string                 `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CompletedAt   respjson.Field
		CreatedAt     respjson.Field
		Endpoint      respjson.Field
		Error         respjson.Field
		ErrorFileID   respjson.Field
		FileSizeBytes respjson.Field
		InputFileID   respjson.Field
		JobDeadline   respjson.Field
		ModelID       respjson.Field
		OutputFileID  respjson.Field
		Progress      respjson.Field
		Status        respjson.Field
		UserID        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the batch job
type BatchGetResponseStatus string

const (
	BatchGetResponseStatusValidating BatchGetResponseStatus = "VALIDATING"
	BatchGetResponseStatusInProgress BatchGetResponseStatus = "IN_PROGRESS"
	BatchGetResponseStatusCompleted  BatchGetResponseStatus = "COMPLETED"
	BatchGetResponseStatusFailed     BatchGetResponseStatus = "FAILED"
	BatchGetResponseStatusExpired    BatchGetResponseStatus = "EXPIRED"
	BatchGetResponseStatusCancelled  BatchGetResponseStatus = "CANCELLED"
)

type BatchListResponse struct {
	ID          string    `json:"id" format:"uuid"`
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	Endpoint    string    `json:"endpoint"`
	Error       string    `json:"error"`
	ErrorFileID string    `json:"error_file_id"`
	// Size of input file in bytes
	FileSizeBytes int64     `json:"file_size_bytes"`
	InputFileID   string    `json:"input_file_id"`
	JobDeadline   time.Time `json:"job_deadline" format:"date-time"`
	// Model used for processing requests
	ModelID      string `json:"model_id"`
	OutputFileID string `json:"output_file_id"`
	// Completion progress (0.0 to 100)
	Progress float64 `json:"progress"`
	// Current status of the batch job
	//
	// Any of "VALIDATING", "IN_PROGRESS", "COMPLETED", "FAILED", "EXPIRED",
	// "CANCELLED".
	Status BatchListResponseStatus `json:"status"`
	UserID string                  `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CompletedAt   respjson.Field
		CreatedAt     respjson.Field
		Endpoint      respjson.Field
		Error         respjson.Field
		ErrorFileID   respjson.Field
		FileSizeBytes respjson.Field
		InputFileID   respjson.Field
		JobDeadline   respjson.Field
		ModelID       respjson.Field
		OutputFileID  respjson.Field
		Progress      respjson.Field
		Status        respjson.Field
		UserID        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the batch job
type BatchListResponseStatus string

const (
	BatchListResponseStatusValidating BatchListResponseStatus = "VALIDATING"
	BatchListResponseStatusInProgress BatchListResponseStatus = "IN_PROGRESS"
	BatchListResponseStatusCompleted  BatchListResponseStatus = "COMPLETED"
	BatchListResponseStatusFailed     BatchListResponseStatus = "FAILED"
	BatchListResponseStatusExpired    BatchListResponseStatus = "EXPIRED"
	BatchListResponseStatusCancelled  BatchListResponseStatus = "CANCELLED"
)

type BatchNewParams struct {
	// The endpoint to use for batch processing
	Endpoint string `json:"endpoint,required"`
	// ID of the uploaded input file containing batch requests
	InputFileID string `json:"input_file_id,required"`
	// Time window for batch completion (optional)
	CompletionWindow param.Opt[string] `json:"completion_window,omitzero"`
	// Model to use for processing batch requests
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// Priority for batch processing (optional)
	Priority param.Opt[int64] `json:"priority,omitzero"`
	paramObj
}

func (r BatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
