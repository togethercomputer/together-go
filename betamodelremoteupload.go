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

// BetaModelRemoteUploadService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaModelRemoteUploadService] method instead.
type BetaModelRemoteUploadService struct {
	Options []option.RequestOption
}

// NewBetaModelRemoteUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaModelRemoteUploadService(opts ...option.RequestOption) (r BetaModelRemoteUploadService) {
	r = BetaModelRemoteUploadService{}
	r.Options = opts
	return
}

// Starts an asynchronous job that imports model files from Hugging Face or a
// presigned URL into a registered model and creates a model revision when the
// import completes.
func (r *BetaModelRemoteUploadService) New(ctx context.Context, params BetaModelRemoteUploadNewParams, opts ...option.RequestOption) (res *BetaModelRemoteUploadNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/models/uploads", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves the status, progress details, retry counts, and timestamps for a
// remote model import job.
func (r *BetaModelRemoteUploadService) Get(ctx context.Context, id string, query BetaModelRemoteUploadGetParams, opts ...option.RequestOption) (res *BetaModelRemoteUploadGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
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
	path := fmt.Sprintf("projects/%s/models/uploads/%s", query.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists asynchronous jobs that import model files from Hugging Face or a presigned
// remote URL.
func (r *BetaModelRemoteUploadService) List(ctx context.Context, params BetaModelRemoteUploadListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[BetaModelRemoteUploadListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/models/uploads", params.ProjectID.Value)
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
func (r *BetaModelRemoteUploadService) ListAutoPaging(ctx context.Context, params BetaModelRemoteUploadListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[BetaModelRemoteUploadListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Lists progress and diagnostic events for a remote model import job.
func (r *BetaModelRemoteUploadService) Events(ctx context.Context, id string, params BetaModelRemoteUploadEventsParams, opts ...option.RequestOption) (res *BetaModelRemoteUploadEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
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
	path := fmt.Sprintf("projects/%s/models/uploads/%s/events", params.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Asynchronous job that imports remote files into a registered model and creates a
// model revision.
type BetaModelRemoteUploadNewResponse struct {
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
	Status BetaModelRemoteUploadNewResponseStatus `json:"status" api:"required"`
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
func (r BetaModelRemoteUploadNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaModelRemoteUploadNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current lifecycle state of the asynchronous import job.
type BetaModelRemoteUploadNewResponseStatus string

const (
	BetaModelRemoteUploadNewResponseStatusRemoteUploadStatusPending   BetaModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_PENDING"
	BetaModelRemoteUploadNewResponseStatusRemoteUploadStatusRunning   BetaModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_RUNNING"
	BetaModelRemoteUploadNewResponseStatusRemoteUploadStatusError     BetaModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_ERROR"
	BetaModelRemoteUploadNewResponseStatusRemoteUploadStatusSucceeded BetaModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_SUCCEEDED"
	BetaModelRemoteUploadNewResponseStatusRemoteUploadStatusFailed    BetaModelRemoteUploadNewResponseStatus = "REMOTE_UPLOAD_STATUS_FAILED"
)

// Asynchronous job that imports remote files into a registered model and creates a
// model revision.
type BetaModelRemoteUploadGetResponse struct {
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
	Status BetaModelRemoteUploadGetResponseStatus `json:"status" api:"required"`
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
func (r BetaModelRemoteUploadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaModelRemoteUploadGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current lifecycle state of the asynchronous import job.
type BetaModelRemoteUploadGetResponseStatus string

const (
	BetaModelRemoteUploadGetResponseStatusRemoteUploadStatusPending   BetaModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_PENDING"
	BetaModelRemoteUploadGetResponseStatusRemoteUploadStatusRunning   BetaModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_RUNNING"
	BetaModelRemoteUploadGetResponseStatusRemoteUploadStatusError     BetaModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_ERROR"
	BetaModelRemoteUploadGetResponseStatusRemoteUploadStatusSucceeded BetaModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_SUCCEEDED"
	BetaModelRemoteUploadGetResponseStatusRemoteUploadStatusFailed    BetaModelRemoteUploadGetResponseStatus = "REMOTE_UPLOAD_STATUS_FAILED"
)

// Asynchronous job that imports remote files into a registered model and creates a
// model revision.
type BetaModelRemoteUploadListResponse struct {
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
	Status BetaModelRemoteUploadListResponseStatus `json:"status" api:"required"`
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
func (r BetaModelRemoteUploadListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaModelRemoteUploadListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current lifecycle state of the asynchronous import job.
type BetaModelRemoteUploadListResponseStatus string

const (
	BetaModelRemoteUploadListResponseStatusRemoteUploadStatusPending   BetaModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_PENDING"
	BetaModelRemoteUploadListResponseStatusRemoteUploadStatusRunning   BetaModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_RUNNING"
	BetaModelRemoteUploadListResponseStatusRemoteUploadStatusError     BetaModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_ERROR"
	BetaModelRemoteUploadListResponseStatusRemoteUploadStatusSucceeded BetaModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_SUCCEEDED"
	BetaModelRemoteUploadListResponseStatusRemoteUploadStatusFailed    BetaModelRemoteUploadListResponseStatus = "REMOTE_UPLOAD_STATUS_FAILED"
)

// Status and diagnostic events for a remote model import job.
type BetaModelRemoteUploadEventsResponse struct {
	// Events for the remote upload.
	Data []BetaModelRemoteUploadEventsResponseData `json:"data" api:"required"`
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
func (r BetaModelRemoteUploadEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaModelRemoteUploadEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Progress or diagnostic event emitted while importing remote model files.
type BetaModelRemoteUploadEventsResponseData struct {
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
func (r BetaModelRemoteUploadEventsResponseData) RawJSON() string { return r.JSON.raw }
func (r *BetaModelRemoteUploadEventsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaModelRemoteUploadNewParams struct {
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

func (r BetaModelRemoteUploadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaModelRemoteUploadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaModelRemoteUploadNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaModelRemoteUploadGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	paramObj
}

type BetaModelRemoteUploadListParams struct {
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

// URLQuery serializes [BetaModelRemoteUploadListParams]'s query parameters as
// `url.Values`.
func (r BetaModelRemoteUploadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaModelRemoteUploadEventsParams struct {
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

// URLQuery serializes [BetaModelRemoteUploadEventsParams]'s query parameters as
// `url.Values`.
func (r BetaModelRemoteUploadEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
