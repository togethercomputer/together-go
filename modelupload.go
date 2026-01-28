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

// ModelUploadService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelUploadService] method instead.
type ModelUploadService struct {
	Options []option.RequestOption
}

// NewModelUploadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModelUploadService(opts ...option.RequestOption) (r ModelUploadService) {
	r = ModelUploadService{}
	r.Options = opts
	return
}

// Get the status of a specific job
func (r *ModelUploadService) Status(ctx context.Context, jobID string, opts ...option.RequestOption) (res *ModelUploadStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ModelUploadStatusResponse struct {
	Args      ModelUploadStatusResponseArgs `json:"args,required"`
	CreatedAt time.Time                     `json:"created_at,required" format:"date-time"`
	JobID     string                        `json:"job_id,required"`
	// Any of "Queued", "Running", "Complete", "Failed".
	Status        ModelUploadStatusResponseStatus         `json:"status,required"`
	StatusUpdates []ModelUploadStatusResponseStatusUpdate `json:"status_updates,required"`
	Type          string                                  `json:"type,required"`
	UpdatedAt     time.Time                               `json:"updated_at,required" format:"date-time"`
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
func (r ModelUploadStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelUploadStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelUploadStatusResponseArgs struct {
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
func (r ModelUploadStatusResponseArgs) RawJSON() string { return r.JSON.raw }
func (r *ModelUploadStatusResponseArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelUploadStatusResponseStatus string

const (
	ModelUploadStatusResponseStatusQueued   ModelUploadStatusResponseStatus = "Queued"
	ModelUploadStatusResponseStatusRunning  ModelUploadStatusResponseStatus = "Running"
	ModelUploadStatusResponseStatusComplete ModelUploadStatusResponseStatus = "Complete"
	ModelUploadStatusResponseStatusFailed   ModelUploadStatusResponseStatus = "Failed"
)

type ModelUploadStatusResponseStatusUpdate struct {
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
func (r ModelUploadStatusResponseStatusUpdate) RawJSON() string { return r.JSON.raw }
func (r *ModelUploadStatusResponseStatusUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
