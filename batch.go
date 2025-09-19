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
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
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
func NewBatchService(opts ...option.RequestOption) (r *BatchService) {
	r = &BatchService{}
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
	Job     BatchNewResponseJob  `json:"job"`
	Warning string               `json:"warning"`
	JSON    batchNewResponseJSON `json:"-"`
}

// batchNewResponseJSON contains the JSON metadata for the struct
// [BatchNewResponse]
type batchNewResponseJSON struct {
	Job         apijson.Field
	Warning     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BatchNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r batchNewResponseJSON) RawJSON() string {
	return r.raw
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
	Status BatchNewResponseJobStatus `json:"status"`
	UserID string                    `json:"user_id"`
	JSON   batchNewResponseJobJSON   `json:"-"`
}

// batchNewResponseJobJSON contains the JSON metadata for the struct
// [BatchNewResponseJob]
type batchNewResponseJobJSON struct {
	ID            apijson.Field
	CompletedAt   apijson.Field
	CreatedAt     apijson.Field
	Endpoint      apijson.Field
	Error         apijson.Field
	ErrorFileID   apijson.Field
	FileSizeBytes apijson.Field
	InputFileID   apijson.Field
	JobDeadline   apijson.Field
	ModelID       apijson.Field
	OutputFileID  apijson.Field
	Progress      apijson.Field
	Status        apijson.Field
	UserID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BatchNewResponseJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r batchNewResponseJobJSON) RawJSON() string {
	return r.raw
}

// Current status of the batch job
type BatchNewResponseJobStatus string

const (
	BatchNewResponseJobStatusValidating BatchNewResponseJobStatus = "VALIDATING"
	BatchNewResponseJobStatusInProgress BatchNewResponseJobStatus = "IN_PROGRESS"
	BatchNewResponseJobStatusCompleted  BatchNewResponseJobStatus = "COMPLETED"
	BatchNewResponseJobStatusFailed     BatchNewResponseJobStatus = "FAILED"
	BatchNewResponseJobStatusExpired    BatchNewResponseJobStatus = "EXPIRED"
	BatchNewResponseJobStatusCancelled  BatchNewResponseJobStatus = "CANCELLED"
)

func (r BatchNewResponseJobStatus) IsKnown() bool {
	switch r {
	case BatchNewResponseJobStatusValidating, BatchNewResponseJobStatusInProgress, BatchNewResponseJobStatusCompleted, BatchNewResponseJobStatusFailed, BatchNewResponseJobStatusExpired, BatchNewResponseJobStatusCancelled:
		return true
	}
	return false
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
	Status BatchGetResponseStatus `json:"status"`
	UserID string                 `json:"user_id"`
	JSON   batchGetResponseJSON   `json:"-"`
}

// batchGetResponseJSON contains the JSON metadata for the struct
// [BatchGetResponse]
type batchGetResponseJSON struct {
	ID            apijson.Field
	CompletedAt   apijson.Field
	CreatedAt     apijson.Field
	Endpoint      apijson.Field
	Error         apijson.Field
	ErrorFileID   apijson.Field
	FileSizeBytes apijson.Field
	InputFileID   apijson.Field
	JobDeadline   apijson.Field
	ModelID       apijson.Field
	OutputFileID  apijson.Field
	Progress      apijson.Field
	Status        apijson.Field
	UserID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BatchGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r batchGetResponseJSON) RawJSON() string {
	return r.raw
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

func (r BatchGetResponseStatus) IsKnown() bool {
	switch r {
	case BatchGetResponseStatusValidating, BatchGetResponseStatusInProgress, BatchGetResponseStatusCompleted, BatchGetResponseStatusFailed, BatchGetResponseStatusExpired, BatchGetResponseStatusCancelled:
		return true
	}
	return false
}

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
	Status BatchListResponseStatus `json:"status"`
	UserID string                  `json:"user_id"`
	JSON   batchListResponseJSON   `json:"-"`
}

// batchListResponseJSON contains the JSON metadata for the struct
// [BatchListResponse]
type batchListResponseJSON struct {
	ID            apijson.Field
	CompletedAt   apijson.Field
	CreatedAt     apijson.Field
	Endpoint      apijson.Field
	Error         apijson.Field
	ErrorFileID   apijson.Field
	FileSizeBytes apijson.Field
	InputFileID   apijson.Field
	JobDeadline   apijson.Field
	ModelID       apijson.Field
	OutputFileID  apijson.Field
	Progress      apijson.Field
	Status        apijson.Field
	UserID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BatchListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r batchListResponseJSON) RawJSON() string {
	return r.raw
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

func (r BatchListResponseStatus) IsKnown() bool {
	switch r {
	case BatchListResponseStatusValidating, BatchListResponseStatusInProgress, BatchListResponseStatusCompleted, BatchListResponseStatusFailed, BatchListResponseStatusExpired, BatchListResponseStatusCancelled:
		return true
	}
	return false
}

type BatchNewParams struct {
	// The endpoint to use for batch processing
	Endpoint param.Field[string] `json:"endpoint,required"`
	// ID of the uploaded input file containing batch requests
	InputFileID param.Field[string] `json:"input_file_id,required"`
	// Time window for batch completion (optional)
	CompletionWindow param.Field[string] `json:"completion_window"`
	// Model to use for processing batch requests
	ModelID param.Field[string] `json:"model_id"`
	// Priority for batch processing (optional)
	Priority param.Field[int64] `json:"priority"`
}

func (r BatchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
