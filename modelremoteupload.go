// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/pagination"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
	"github.com/togethercomputer/together-go/shared/constant"
)

// ModelRemoteUploadService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelRemoteUploadService] method instead.
type ModelRemoteUploadService struct {
	Options []option.RequestOption
}

// NewModelRemoteUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewModelRemoteUploadService(opts ...option.RequestOption) (r ModelRemoteUploadService) {
	r = ModelRemoteUploadService{}
	r.Options = opts
	return
}

// Starts an asynchronous job that imports model files from Hugging Face or a
// presigned URL into a registered model and creates a model revision when the
// import completes.
func (r *ModelRemoteUploadService) New(ctx context.Context, params ModelRemoteUploadNewParams, opts ...option.RequestOption) (res *ModelRemoteUploadNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/projects/%s/models/uploads", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves the status, progress details, retry counts, and timestamps for a
// remote model import job.
func (r *ModelRemoteUploadService) Get(ctx context.Context, id string, query ModelRemoteUploadGetParams, opts ...option.RequestOption) (res *ModelRemoteUploadGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	if query.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/projects/%s/models/uploads/%s", query.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists asynchronous jobs that import model files from Hugging Face or a presigned
// remote URL.
func (r *ModelRemoteUploadService) List(ctx context.Context, params ModelRemoteUploadListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ModelRemoteUploadListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/projects/%s/models/uploads", params.ProjectID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists asynchronous jobs that import model files from Hugging Face or a presigned
// remote URL.
func (r *ModelRemoteUploadService) ListAutoPaging(ctx context.Context, params ModelRemoteUploadListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ModelRemoteUploadListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Lists progress and diagnostic events for a remote model import job.
func (r *ModelRemoteUploadService) Events(ctx context.Context, id string, params ModelRemoteUploadEventsParams, opts ...option.RequestOption) (res *ModelRemoteUploadEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/projects/%s/models/uploads/%s/events", params.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Asynchronous job that imports remote files into a registered model and creates a
// model revision.
type ModelRemoteUploadNewResponse struct {
	// Unique ID of the remote model import job.
	ID string `json:"id" api:"required"`
	// Time when the import job was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// ID of the registered model receiving the imported files.
	ModelID string `json:"modelId" api:"required"`
	// ID of the project that owns the import job.
	ProjectID string `json:"projectId" api:"required"`
	// Hugging Face repository or presigned URL being imported.
	RemoteURL string `json:"remoteUrl" api:"required"`
	// Current lifecycle state of the asynchronous import job.
	//
	// Any of "REMOTE_UPLOAD_STATUS_PENDING", "REMOTE_UPLOAD_STATUS_RUNNING",
	// "REMOTE_UPLOAD_STATUS_ERROR", "REMOTE_UPLOAD_STATUS_SUCCEEDED",
	// "REMOTE_UPLOAD_STATUS_FAILED".
	Status ModelRemoteUploadNewResponseStatus `json:"status" api:"required"`
	// Maximum worker restarts allowed before the job fails permanently.
	MaxRestarts int64 `json:"maxRestarts"`
	// Number of times the import worker has restarted this job.
	RestartCount int64 `json:"restartCount"`
	// Human-readable progress or failure detail for the current status.
	StatusMessage string `json:"statusMessage"`
	// Time when the import job was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		ModelID       respjson.Field
		ProjectID     respjson.Field
		RemoteURL     respjson.Field
		Status        respjson.Field
		MaxRestarts   respjson.Field
		RestartCount  respjson.Field
		StatusMessage respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelRemoteUploadNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelRemoteUploadNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current lifecycle state of the asynchronous import job.
type ModelRemoteUploadNewResponseStatus string

const (
	ModelRemoteUploadNewResponseStatusRemoteUploadStatusPending   ModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_PENDING"
	ModelRemoteUploadNewResponseStatusRemoteUploadStatusRunning   ModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_RUNNING"
	ModelRemoteUploadNewResponseStatusRemoteUploadStatusError     ModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_ERROR"
	ModelRemoteUploadNewResponseStatusRemoteUploadStatusSucceeded ModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_SUCCEEDED"
	ModelRemoteUploadNewResponseStatusRemoteUploadStatusFailed    ModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_FAILED"
)

// Asynchronous job that imports remote files into a registered model and creates a
// model revision.
type ModelRemoteUploadGetResponse struct {
	// Unique ID of the remote model import job.
	ID string `json:"id" api:"required"`
	// Time when the import job was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// ID of the registered model receiving the imported files.
	ModelID string `json:"modelId" api:"required"`
	// ID of the project that owns the import job.
	ProjectID string `json:"projectId" api:"required"`
	// Hugging Face repository or presigned URL being imported.
	RemoteURL string `json:"remoteUrl" api:"required"`
	// Current lifecycle state of the asynchronous import job.
	//
	// Any of "REMOTE_UPLOAD_STATUS_PENDING", "REMOTE_UPLOAD_STATUS_RUNNING",
	// "REMOTE_UPLOAD_STATUS_ERROR", "REMOTE_UPLOAD_STATUS_SUCCEEDED",
	// "REMOTE_UPLOAD_STATUS_FAILED".
	Status ModelRemoteUploadGetResponseStatus `json:"status" api:"required"`
	// Maximum worker restarts allowed before the job fails permanently.
	MaxRestarts int64 `json:"maxRestarts"`
	// Number of times the import worker has restarted this job.
	RestartCount int64 `json:"restartCount"`
	// Human-readable progress or failure detail for the current status.
	StatusMessage string `json:"statusMessage"`
	// Time when the import job was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		ModelID       respjson.Field
		ProjectID     respjson.Field
		RemoteURL     respjson.Field
		Status        respjson.Field
		MaxRestarts   respjson.Field
		RestartCount  respjson.Field
		StatusMessage respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelRemoteUploadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelRemoteUploadGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current lifecycle state of the asynchronous import job.
type ModelRemoteUploadGetResponseStatus string

const (
	ModelRemoteUploadGetResponseStatusRemoteUploadStatusPending   ModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_PENDING"
	ModelRemoteUploadGetResponseStatusRemoteUploadStatusRunning   ModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_RUNNING"
	ModelRemoteUploadGetResponseStatusRemoteUploadStatusError     ModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_ERROR"
	ModelRemoteUploadGetResponseStatusRemoteUploadStatusSucceeded ModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_SUCCEEDED"
	ModelRemoteUploadGetResponseStatusRemoteUploadStatusFailed    ModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_FAILED"
)

// Asynchronous job that imports remote files into a registered model and creates a
// model revision.
type ModelRemoteUploadListResponse struct {
	// Unique ID of the remote model import job.
	ID string `json:"id" api:"required"`
	// Time when the import job was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// ID of the registered model receiving the imported files.
	ModelID string `json:"modelId" api:"required"`
	// ID of the project that owns the import job.
	ProjectID string `json:"projectId" api:"required"`
	// Hugging Face repository or presigned URL being imported.
	RemoteURL string `json:"remoteUrl" api:"required"`
	// Current lifecycle state of the asynchronous import job.
	//
	// Any of "REMOTE_UPLOAD_STATUS_PENDING", "REMOTE_UPLOAD_STATUS_RUNNING",
	// "REMOTE_UPLOAD_STATUS_ERROR", "REMOTE_UPLOAD_STATUS_SUCCEEDED",
	// "REMOTE_UPLOAD_STATUS_FAILED".
	Status ModelRemoteUploadListResponseStatus `json:"status" api:"required"`
	// Maximum worker restarts allowed before the job fails permanently.
	MaxRestarts int64 `json:"maxRestarts"`
	// Number of times the import worker has restarted this job.
	RestartCount int64 `json:"restartCount"`
	// Human-readable progress or failure detail for the current status.
	StatusMessage string `json:"statusMessage"`
	// Time when the import job was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		ModelID       respjson.Field
		ProjectID     respjson.Field
		RemoteURL     respjson.Field
		Status        respjson.Field
		MaxRestarts   respjson.Field
		RestartCount  respjson.Field
		StatusMessage respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelRemoteUploadListResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelRemoteUploadListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current lifecycle state of the asynchronous import job.
type ModelRemoteUploadListResponseStatus string

const (
	ModelRemoteUploadListResponseStatusRemoteUploadStatusPending   ModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_PENDING"
	ModelRemoteUploadListResponseStatusRemoteUploadStatusRunning   ModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_RUNNING"
	ModelRemoteUploadListResponseStatusRemoteUploadStatusError     ModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_ERROR"
	ModelRemoteUploadListResponseStatusRemoteUploadStatusSucceeded ModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_SUCCEEDED"
	ModelRemoteUploadListResponseStatusRemoteUploadStatusFailed    ModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_FAILED"
)

// Status and diagnostic events for a remote model import job.
type ModelRemoteUploadEventsResponse struct {
	// Events for the remote upload.
	Data []ModelRemoteUploadEventsResponseData `json:"data" api:"required"`
	// Object type. Always `list`.
	Object constant.List `json:"object" default:"list"`
	// Cursor for the next page. Null if there are no more results.
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelRemoteUploadEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelRemoteUploadEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Progress or diagnostic event emitted while importing remote model files.
type ModelRemoteUploadEventsResponseData struct {
	// Unique event identifier.
	ID string `json:"id" api:"required"`
	// Time when the event was recorded.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Severity of the event.
	//
	// Any of "LEVEL_DEBUG", "LEVEL_INFO", "LEVEL_WARN", "LEVEL_ERROR".
	Level string `json:"level" api:"required"`
	// Human-readable progress or diagnostic message.
	Message string `json:"message" api:"required"`
	// Stable event type emitted by the importer, such as `download.started`.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelRemoteUploadEventsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ModelRemoteUploadEventsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelRemoteUploadNewParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// ID of the registered model that will receive the imported files.
	ModelID string `json:"modelId" api:"required"`
	// Hugging Face repository URL or presigned archive URL to import.
	RemoteURL string `json:"remoteUrl" api:"required"`
	// Optional source credential used to access a private remote location. The value
	// is write-only and is not returned.
	Token param.Opt[string] `json:"token,omitzero"`
	paramObj
}

func (r ModelRemoteUploadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ModelRemoteUploadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelRemoteUploadNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelRemoteUploadGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	paramObj
}

type ModelRemoteUploadListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous remote upload list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of uploads to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ModelRemoteUploadListParams]'s query parameters as
// `url.Values`.
func (r ModelRemoteUploadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ModelRemoteUploadEventsParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous remote upload event list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of events to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ModelRemoteUploadEventsParams]'s query parameters as
// `url.Values`.
func (r ModelRemoteUploadEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
