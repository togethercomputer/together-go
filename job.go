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
func NewJobService(opts ...option.RequestOption) (r *JobService) {
	r = &JobService{}
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
	Args          JobGetResponseArgs           `json:"args,required"`
	CreatedAt     time.Time                    `json:"created_at,required" format:"date-time"`
	JobID         string                       `json:"job_id,required"`
	Status        JobGetResponseStatus         `json:"status,required"`
	StatusUpdates []JobGetResponseStatusUpdate `json:"status_updates,required"`
	Type          string                       `json:"type,required"`
	UpdatedAt     time.Time                    `json:"updated_at,required" format:"date-time"`
	JSON          jobGetResponseJSON           `json:"-"`
}

// jobGetResponseJSON contains the JSON metadata for the struct [JobGetResponse]
type jobGetResponseJSON struct {
	Args          apijson.Field
	CreatedAt     apijson.Field
	JobID         apijson.Field
	Status        apijson.Field
	StatusUpdates apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *JobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseArgs struct {
	Description string                 `json:"description"`
	ModelName   string                 `json:"modelName"`
	ModelSource string                 `json:"modelSource"`
	JSON        jobGetResponseArgsJSON `json:"-"`
}

// jobGetResponseArgsJSON contains the JSON metadata for the struct
// [JobGetResponseArgs]
type jobGetResponseArgsJSON struct {
	Description apijson.Field
	ModelName   apijson.Field
	ModelSource apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseArgsJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseStatus string

const (
	JobGetResponseStatusQueued   JobGetResponseStatus = "Queued"
	JobGetResponseStatusRunning  JobGetResponseStatus = "Running"
	JobGetResponseStatusComplete JobGetResponseStatus = "Complete"
	JobGetResponseStatusFailed   JobGetResponseStatus = "Failed"
)

func (r JobGetResponseStatus) IsKnown() bool {
	switch r {
	case JobGetResponseStatusQueued, JobGetResponseStatusRunning, JobGetResponseStatusComplete, JobGetResponseStatusFailed:
		return true
	}
	return false
}

type JobGetResponseStatusUpdate struct {
	Message   string                         `json:"message,required"`
	Status    string                         `json:"status,required"`
	Timestamp time.Time                      `json:"timestamp,required" format:"date-time"`
	JSON      jobGetResponseStatusUpdateJSON `json:"-"`
}

// jobGetResponseStatusUpdateJSON contains the JSON metadata for the struct
// [JobGetResponseStatusUpdate]
type jobGetResponseStatusUpdateJSON struct {
	Message     apijson.Field
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseStatusUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseStatusUpdateJSON) RawJSON() string {
	return r.raw
}

type JobListResponse struct {
	Data []JobListResponseData `json:"data,required"`
	JSON jobListResponseJSON   `json:"-"`
}

// jobListResponseJSON contains the JSON metadata for the struct [JobListResponse]
type jobListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobListResponseJSON) RawJSON() string {
	return r.raw
}

type JobListResponseData struct {
	Args          JobListResponseDataArgs           `json:"args,required"`
	CreatedAt     time.Time                         `json:"created_at,required" format:"date-time"`
	JobID         string                            `json:"job_id,required"`
	Status        JobListResponseDataStatus         `json:"status,required"`
	StatusUpdates []JobListResponseDataStatusUpdate `json:"status_updates,required"`
	Type          string                            `json:"type,required"`
	UpdatedAt     time.Time                         `json:"updated_at,required" format:"date-time"`
	JSON          jobListResponseDataJSON           `json:"-"`
}

// jobListResponseDataJSON contains the JSON metadata for the struct
// [JobListResponseData]
type jobListResponseDataJSON struct {
	Args          apijson.Field
	CreatedAt     apijson.Field
	JobID         apijson.Field
	Status        apijson.Field
	StatusUpdates apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *JobListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobListResponseDataJSON) RawJSON() string {
	return r.raw
}

type JobListResponseDataArgs struct {
	Description string                      `json:"description"`
	ModelName   string                      `json:"modelName"`
	ModelSource string                      `json:"modelSource"`
	JSON        jobListResponseDataArgsJSON `json:"-"`
}

// jobListResponseDataArgsJSON contains the JSON metadata for the struct
// [JobListResponseDataArgs]
type jobListResponseDataArgsJSON struct {
	Description apijson.Field
	ModelName   apijson.Field
	ModelSource apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobListResponseDataArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobListResponseDataArgsJSON) RawJSON() string {
	return r.raw
}

type JobListResponseDataStatus string

const (
	JobListResponseDataStatusQueued   JobListResponseDataStatus = "Queued"
	JobListResponseDataStatusRunning  JobListResponseDataStatus = "Running"
	JobListResponseDataStatusComplete JobListResponseDataStatus = "Complete"
	JobListResponseDataStatusFailed   JobListResponseDataStatus = "Failed"
)

func (r JobListResponseDataStatus) IsKnown() bool {
	switch r {
	case JobListResponseDataStatusQueued, JobListResponseDataStatusRunning, JobListResponseDataStatusComplete, JobListResponseDataStatusFailed:
		return true
	}
	return false
}

type JobListResponseDataStatusUpdate struct {
	Message   string                              `json:"message,required"`
	Status    string                              `json:"status,required"`
	Timestamp time.Time                           `json:"timestamp,required" format:"date-time"`
	JSON      jobListResponseDataStatusUpdateJSON `json:"-"`
}

// jobListResponseDataStatusUpdateJSON contains the JSON metadata for the struct
// [JobListResponseDataStatusUpdate]
type jobListResponseDataStatusUpdateJSON struct {
	Message     apijson.Field
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobListResponseDataStatusUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobListResponseDataStatusUpdateJSON) RawJSON() string {
	return r.raw
}
