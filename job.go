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
	"github.com/togethercomputer/together-go/packages/respjson"
)

// JobService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobService] method instead.
type JobService struct {
	Options []option.RequestOption
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r JobService) {
	r = JobService{}
	r.Options = opts
	return
}

// Get the status of a specific job
func (r *JobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all jobs and their statuses
func (r *JobService) List(ctx context.Context, opts ...option.RequestOption) (res *JobListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type JobGetResponse struct {
	Args      JobGetResponseArgs `json:"args,required"`
	CreatedAt time.Time          `json:"created_at,required" format:"date-time"`
	JobID     string             `json:"job_id,required"`
	// Any of "Queued", "Running", "Complete", "Failed".
	Status        JobGetResponseStatus         `json:"status,required"`
	StatusUpdates []JobGetResponseStatusUpdate `json:"status_updates,required"`
	Type          string                       `json:"type,required"`
	UpdatedAt     time.Time                    `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args          respjson.Field
		CreatedAt     respjson.Field
		JobID         respjson.Field
		Status        respjson.Field
		StatusUpdates respjson.Field
		Type          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetResponseArgs struct {
	Description string `json:"description"`
	ModelName   string `json:"modelName"`
	ModelSource string `json:"modelSource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		ModelName   respjson.Field
		ModelSource respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponseArgs) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponseArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetResponseStatus string

const (
	JobGetResponseStatusQueued   JobGetResponseStatus = "Queued"
	JobGetResponseStatusRunning  JobGetResponseStatus = "Running"
	JobGetResponseStatusComplete JobGetResponseStatus = "Complete"
	JobGetResponseStatusFailed   JobGetResponseStatus = "Failed"
)

type JobGetResponseStatusUpdate struct {
	Message   string    `json:"message,required"`
	Status    string    `json:"status,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponseStatusUpdate) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponseStatusUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListResponse struct {
	Data []JobListResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobListResponse) RawJSON() string { return r.JSON.raw }
func (r *JobListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListResponseData struct {
	Args      JobListResponseDataArgs `json:"args,required"`
	CreatedAt time.Time               `json:"created_at,required" format:"date-time"`
	JobID     string                  `json:"job_id,required"`
	// Any of "Queued", "Running", "Complete", "Failed".
	Status        string                            `json:"status,required"`
	StatusUpdates []JobListResponseDataStatusUpdate `json:"status_updates,required"`
	Type          string                            `json:"type,required"`
	UpdatedAt     time.Time                         `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args          respjson.Field
		CreatedAt     respjson.Field
		JobID         respjson.Field
		Status        respjson.Field
		StatusUpdates respjson.Field
		Type          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobListResponseData) RawJSON() string { return r.JSON.raw }
func (r *JobListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListResponseDataArgs struct {
	Description string `json:"description"`
	ModelName   string `json:"modelName"`
	ModelSource string `json:"modelSource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		ModelName   respjson.Field
		ModelSource respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobListResponseDataArgs) RawJSON() string { return r.JSON.raw }
func (r *JobListResponseDataArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListResponseDataStatusUpdate struct {
	Message   string    `json:"message,required"`
	Status    string    `json:"status,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobListResponseDataStatusUpdate) RawJSON() string { return r.JSON.raw }
func (r *JobListResponseDataStatusUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
