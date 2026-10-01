// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

// BetaJigVolumeService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaJigVolumeService] method instead.
type BetaJigVolumeService struct {
	Options []option.RequestOption
}

// NewBetaJigVolumeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaJigVolumeService(opts ...option.RequestOption) (r BetaJigVolumeService) {
	r = BetaJigVolumeService{}
	r.Options = opts
	return
}

// Create a new volume to preload files in deployments
func (r *BetaJigVolumeService) New(ctx context.Context, body BetaJigVolumeNewParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments/storage/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve details of a specific volume by its ID or name
func (r *BetaJigVolumeService) Get(ctx context.Context, id string, query BetaJigVolumeGetParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/storage/volumes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update an existing volume's configuration or contents
func (r *BetaJigVolumeService) Update(ctx context.Context, id string, body BetaJigVolumeUpdateParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/storage/volumes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve all volumes in your project
func (r *BetaJigVolumeService) List(ctx context.Context, opts ...option.RequestOption) (res *BetaJigVolumeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments/storage/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete an existing volume
func (r *BetaJigVolumeService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *BetaJigVolumeDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/storage/volumes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// S3 source configuration for volume sync.
type S3Origin struct {
	// IAM role ARN Together assumes to read the S3 bucket or prefix.
	RoleArn string `json:"role_arn" api:"required"`
	// S3 bucket or prefix to copy into the volume.
	Uri string `json:"uri" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RoleArn     respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r S3Origin) RawJSON() string { return r.JSON.raw }
func (r *S3Origin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this S3Origin to a S3OriginParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// S3OriginParam.Overrides()
func (r S3Origin) ToParam() S3OriginParam {
	return param.Override[S3OriginParam](json.RawMessage(r.RawJSON()))
}

// S3 source configuration for volume sync.
//
// The properties RoleArn, Uri are required.
type S3OriginParam struct {
	// IAM role ARN Together assumes to read the S3 bucket or prefix.
	RoleArn string `json:"role_arn" api:"required"`
	// S3 bucket or prefix to copy into the volume.
	Uri string `json:"uri" api:"required"`
	paramObj
}

func (r S3OriginParam) MarshalJSON() (data []byte, err error) {
	type shadow S3OriginParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *S3OriginParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Volume struct {
	// ID is the unique identifier for this volume
	ID string `json:"id"`
	// Content currently available on a volume version.
	Content VolumeContent `json:"content"`
	// CreatedAt is the ISO8601 timestamp when this volume was created
	CreatedAt string `json:"created_at"`
	// CurrentVersion is the current version number of this volume
	CurrentVersion int64 `json:"current_version"`
	// MountedBy is the list of deployment IDs currently mounting current volume
	// version
	MountedBy []string `json:"mounted_by"`
	// Name is the name of the volume
	Name string `json:"name"`
	// Object is the type identifier for this response (always "volume")
	Object string `json:"object"`
	// Status of the current volume version.
	//
	// Any of "ready", "pending", "syncing", "failed".
	Status VolumeStatus `json:"status"`
	// Message explaining why the current volume version failed, when applicable.
	StatusMessage string `json:"status_message"`
	// Any of "readOnly".
	Type VolumeType `json:"type"`
	// UpdatedAt is the ISO8601 timestamp when this volume was last updated
	UpdatedAt string `json:"updated_at"`
	// VersionHistory contains previous versions of this volume, keyed by version
	// number
	VersionHistory map[string]VolumeVersionHistory `json:"version_history"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Content        respjson.Field
		CreatedAt      respjson.Field
		CurrentVersion respjson.Field
		MountedBy      respjson.Field
		Name           respjson.Field
		Object         respjson.Field
		Status         respjson.Field
		StatusMessage  respjson.Field
		Type           respjson.Field
		UpdatedAt      respjson.Field
		VersionHistory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Volume) RawJSON() string { return r.JSON.raw }
func (r *Volume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content currently available on a volume version.
type VolumeContent struct {
	// Files is the list of files to preload into the volume, if the volume content
	// type is "files".
	Files []VolumeContentFile `json:"files"`
	// External source Together copied into this volume version.
	Origin VolumeOrigin `json:"origin"`
	// SourcePrefix is the file path prefix for the content to be preloaded into the
	// volume
	SourcePrefix string `json:"source_prefix"`
	// Type is the content type (currently only "files" is supported which allows
	// preloading files uploaded via Files API into the volume)
	//
	// Any of "files".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files        respjson.Field
		Origin       respjson.Field
		SourcePrefix respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeContent) RawJSON() string { return r.JSON.raw }
func (r *VolumeContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeContentFile struct {
	// LastModified is the timestamp when the file was last modified
	LastModified string `json:"last_modified"`
	// Name is the filename including extension (e.g., "model_weights.bin")
	Name string `json:"name"`
	// Size is the file size in bytes
	Size int64 `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Name         respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeContentFile) RawJSON() string { return r.JSON.raw }
func (r *VolumeContentFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeType string

const (
	VolumeTypeReadOnly VolumeType = "readOnly"
)

// Metadata for a previous volume version.
type VolumeVersionHistory struct {
	// Content configuration used to create this version.
	Content VolumeVersionHistoryContent `json:"content"`
	// Deployment IDs currently mounting this version.
	MountedBy []string `json:"mounted_by"`
	// Status of this volume version.
	//
	// Any of "ready", "pending", "syncing", "failed".
	Status VolumeStatus `json:"status"`
	// Message explaining why this volume version failed, when applicable.
	StatusMessage string `json:"status_message"`
	// Numeric version identifier for this volume content.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content       respjson.Field
		MountedBy     respjson.Field
		Status        respjson.Field
		StatusMessage respjson.Field
		Version       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeVersionHistory) RawJSON() string { return r.JSON.raw }
func (r *VolumeVersionHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content configuration used to create this version.
type VolumeVersionHistoryContent struct {
	// External source Together copies into a new volume version; mutually exclusive
	// with source_prefix.
	Origin VolumeOrigin `json:"origin"`
	// SourcePrefix is the file path prefix for the content to be preloaded into the
	// volume. Mutually exclusive with Origin
	SourcePrefix string `json:"source_prefix"`
	// Type is the content type (currently only "files" is supported which allows
	// preloading files uploaded via Files API into the volume)
	//
	// Any of "files".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Origin       respjson.Field
		SourcePrefix respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeVersionHistoryContent) RawJSON() string { return r.JSON.raw }
func (r *VolumeVersionHistoryContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External source Together copies into a new volume version.
type VolumeOrigin struct {
	// S3 bucket or prefix source for the volume sync.
	S3 S3Origin `json:"s3" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		S3          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeOrigin) RawJSON() string { return r.JSON.raw }
func (r *VolumeOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VolumeOrigin to a VolumeOriginParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeOriginParam.Overrides()
func (r VolumeOrigin) ToParam() VolumeOriginParam {
	return param.Override[VolumeOriginParam](json.RawMessage(r.RawJSON()))
}

// External source Together copies into a new volume version.
//
// The property S3 is required.
type VolumeOriginParam struct {
	// S3 bucket or prefix source for the volume sync.
	S3 S3OriginParam `json:"s3,omitzero" api:"required"`
	paramObj
}

func (r VolumeOriginParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeOriginParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeOriginParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a volume version. Only ready versions can be mounted.
type VolumeStatus string

const (
	VolumeStatusReady   VolumeStatus = "ready"
	VolumeStatusPending VolumeStatus = "pending"
	VolumeStatusSyncing VolumeStatus = "syncing"
	VolumeStatusFailed  VolumeStatus = "failed"
)

type BetaJigVolumeListResponse struct {
	// Data is the array of volume items
	Data []Volume `json:"data"`
	// The object type, which is always `list`.
	//
	// Any of "list".
	Object BetaJigVolumeListResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigVolumeListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigVolumeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type, which is always `list`.
type BetaJigVolumeListResponseObject string

const (
	BetaJigVolumeListResponseObjectList BetaJigVolumeListResponseObject = "list"
)

type BetaJigVolumeDeleteResponse = any

type BetaJigVolumeNewParams struct {
	// Content specifies the new content to preload to this volume.
	Content BetaJigVolumeNewParamsContent `json:"content,omitzero" api:"required"`
	// Name is the unique identifier for the volume within the project
	Name string `json:"name" api:"required"`
	// Type is the volume type (currently only "readOnly" is supported)
	//
	// Any of "readOnly".
	Type BetaJigVolumeNewParamsType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r BetaJigVolumeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigVolumeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigVolumeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content specifies the new content to preload to this volume.
type BetaJigVolumeNewParamsContent struct {
	// SourcePrefix is the file path prefix for the content to be preloaded into the
	// volume. Mutually exclusive with Origin
	SourcePrefix param.Opt[string] `json:"source_prefix,omitzero"`
	// External source Together copies into a new volume version; mutually exclusive
	// with source_prefix.
	Origin VolumeOriginParam `json:"origin,omitzero"`
	// Type is the content type (currently only "files" is supported which allows
	// preloading files uploaded via Files API into the volume)
	//
	// Any of "files".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BetaJigVolumeNewParamsContent) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigVolumeNewParamsContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigVolumeNewParamsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaJigVolumeNewParamsContent](
		"type", "files",
	)
}

// Type is the volume type (currently only "readOnly" is supported)
type BetaJigVolumeNewParamsType string

const (
	BetaJigVolumeNewParamsTypeReadOnly BetaJigVolumeNewParamsType = "readOnly"
)

type BetaJigVolumeGetParams struct {
	// Volume version to describe (defaults to current version)
	Version param.Opt[int64] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaJigVolumeGetParams]'s query parameters as `url.Values`.
func (r BetaJigVolumeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaJigVolumeUpdateParams struct {
	// Name is the new unique identifier for the volume within the project
	Name param.Opt[string] `json:"name,omitzero"`
	// Content specifies the new content to preload to this volume.
	Content BetaJigVolumeUpdateParamsContent `json:"content,omitzero"`
	// Type is the new volume type (currently only "readOnly" is supported)
	//
	// Any of "readOnly".
	Type BetaJigVolumeUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r BetaJigVolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigVolumeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigVolumeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content specifies the new content to preload to this volume.
type BetaJigVolumeUpdateParamsContent struct {
	// SourcePrefix is the file path prefix for the content to be preloaded into the
	// volume. Mutually exclusive with Origin
	SourcePrefix param.Opt[string] `json:"source_prefix,omitzero"`
	// External source Together copies into a new volume version; mutually exclusive
	// with source_prefix.
	Origin VolumeOriginParam `json:"origin,omitzero"`
	// Type is the content type (currently only "files" is supported which allows
	// preloading files uploaded via Files API into the volume)
	//
	// Any of "files".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BetaJigVolumeUpdateParamsContent) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigVolumeUpdateParamsContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigVolumeUpdateParamsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaJigVolumeUpdateParamsContent](
		"type", "files",
	)
}

// Type is the new volume type (currently only "readOnly" is supported)
type BetaJigVolumeUpdateParamsType string

const (
	BetaJigVolumeUpdateParamsTypeReadOnly BetaJigVolumeUpdateParamsType = "readOnly"
)
