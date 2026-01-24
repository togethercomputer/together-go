// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
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
func (r *BetaJigVolumeService) New(ctx context.Context, body BetaJigVolumeNewParams, opts ...option.RequestOption) (res *BetaJigVolumeNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific volume by its ID or name
func (r *BetaJigVolumeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BetaJigVolumeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("storage/volumes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing volume's configuration or contents
func (r *BetaJigVolumeService) Update(ctx context.Context, id string, body BetaJigVolumeUpdateParams, opts ...option.RequestOption) (res *BetaJigVolumeUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("storage/volumes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve all volumes in your project
func (r *BetaJigVolumeService) List(ctx context.Context, opts ...option.RequestOption) (res *BetaJigVolumeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing volume
func (r *BetaJigVolumeService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *BetaJigVolumeDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("storage/volumes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type BetaJigVolumeNewResponse struct {
	// ID is the unique identifier for this volume
	ID string `json:"id"`
	// Content specifies the content that will be preloaded to this volume
	Content BetaJigVolumeNewResponseContent `json:"content"`
	// CreatedAt is the ISO8601 timestamp when this volume was created
	CreatedAt string `json:"created_at"`
	// Name is the name of the volume
	Name string `json:"name"`
	// Object is the type identifier for this response (always "volume")
	Object string `json:"object"`
	// Type is the volume type (e.g., "readOnly")
	//
	// Any of "readOnly".
	Type BetaJigVolumeNewResponseType `json:"type"`
	// UpdatedAt is the ISO8601 timestamp when this volume was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigVolumeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigVolumeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content specifies the content that will be preloaded to this volume
type BetaJigVolumeNewResponseContent struct {
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
		SourcePrefix respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigVolumeNewResponseContent) RawJSON() string { return r.JSON.raw }
func (r *BetaJigVolumeNewResponseContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type is the volume type (e.g., "readOnly")
type BetaJigVolumeNewResponseType string

const (
	BetaJigVolumeNewResponseTypeReadOnly BetaJigVolumeNewResponseType = "readOnly"
)

type BetaJigVolumeGetResponse struct {
	// ID is the unique identifier for this volume
	ID string `json:"id"`
	// Content specifies the content that will be preloaded to this volume
	Content BetaJigVolumeGetResponseContent `json:"content"`
	// CreatedAt is the ISO8601 timestamp when this volume was created
	CreatedAt string `json:"created_at"`
	// Name is the name of the volume
	Name string `json:"name"`
	// Object is the type identifier for this response (always "volume")
	Object string `json:"object"`
	// Type is the volume type (e.g., "readOnly")
	//
	// Any of "readOnly".
	Type BetaJigVolumeGetResponseType `json:"type"`
	// UpdatedAt is the ISO8601 timestamp when this volume was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigVolumeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigVolumeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content specifies the content that will be preloaded to this volume
type BetaJigVolumeGetResponseContent struct {
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
		SourcePrefix respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigVolumeGetResponseContent) RawJSON() string { return r.JSON.raw }
func (r *BetaJigVolumeGetResponseContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type is the volume type (e.g., "readOnly")
type BetaJigVolumeGetResponseType string

const (
	BetaJigVolumeGetResponseTypeReadOnly BetaJigVolumeGetResponseType = "readOnly"
)

type BetaJigVolumeUpdateResponse struct {
	// ID is the unique identifier for this volume
	ID string `json:"id"`
	// Content specifies the content that will be preloaded to this volume
	Content BetaJigVolumeUpdateResponseContent `json:"content"`
	// CreatedAt is the ISO8601 timestamp when this volume was created
	CreatedAt string `json:"created_at"`
	// Name is the name of the volume
	Name string `json:"name"`
	// Object is the type identifier for this response (always "volume")
	Object string `json:"object"`
	// Type is the volume type (e.g., "readOnly")
	//
	// Any of "readOnly".
	Type BetaJigVolumeUpdateResponseType `json:"type"`
	// UpdatedAt is the ISO8601 timestamp when this volume was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigVolumeUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigVolumeUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content specifies the content that will be preloaded to this volume
type BetaJigVolumeUpdateResponseContent struct {
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
		SourcePrefix respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigVolumeUpdateResponseContent) RawJSON() string { return r.JSON.raw }
func (r *BetaJigVolumeUpdateResponseContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type is the volume type (e.g., "readOnly")
type BetaJigVolumeUpdateResponseType string

const (
	BetaJigVolumeUpdateResponseTypeReadOnly BetaJigVolumeUpdateResponseType = "readOnly"
)

type BetaJigVolumeListResponse struct {
	// Data is the array of volume items
	Data []BetaJigVolumeListResponseData `json:"data"`
	// Object is the type identifier for this response (always "list")
	Object string `json:"object"`
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

type BetaJigVolumeListResponseData struct {
	// ID is the unique identifier for this volume
	ID string `json:"id"`
	// Content specifies the content that will be preloaded to this volume
	Content BetaJigVolumeListResponseDataContent `json:"content"`
	// CreatedAt is the ISO8601 timestamp when this volume was created
	CreatedAt string `json:"created_at"`
	// Name is the name of the volume
	Name string `json:"name"`
	// Object is the type identifier for this response (always "volume")
	Object string `json:"object"`
	// Type is the volume type (e.g., "readOnly")
	//
	// Any of "readOnly".
	Type string `json:"type"`
	// UpdatedAt is the ISO8601 timestamp when this volume was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigVolumeListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BetaJigVolumeListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content specifies the content that will be preloaded to this volume
type BetaJigVolumeListResponseDataContent struct {
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
		SourcePrefix respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigVolumeListResponseDataContent) RawJSON() string { return r.JSON.raw }
func (r *BetaJigVolumeListResponseDataContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigVolumeDeleteResponse = any

type BetaJigVolumeNewParams struct {
	// Content specifies the content configuration for this volume
	Content BetaJigVolumeNewParamsContent `json:"content,omitzero,required"`
	// Name is the unique identifier for the volume within the project
	Name string `json:"name,required"`
	// Type is the volume type (currently only "readOnly" is supported)
	//
	// Any of "readOnly".
	Type BetaJigVolumeNewParamsType `json:"type,omitzero,required"`
	paramObj
}

func (r BetaJigVolumeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigVolumeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigVolumeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content specifies the content configuration for this volume
type BetaJigVolumeNewParamsContent struct {
	// SourcePrefix is the file path prefix for the content to be preloaded into the
	// volume
	SourcePrefix param.Opt[string] `json:"source_prefix,omitzero"`
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

type BetaJigVolumeUpdateParams struct {
	// Name is the new unique identifier for the volume within the project
	Name param.Opt[string] `json:"name,omitzero"`
	// Content specifies the new content that will be preloaded to this volume
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

// Content specifies the new content that will be preloaded to this volume
type BetaJigVolumeUpdateParamsContent struct {
	// SourcePrefix is the file path prefix for the content to be preloaded into the
	// volume
	SourcePrefix param.Opt[string] `json:"source_prefix,omitzero"`
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
