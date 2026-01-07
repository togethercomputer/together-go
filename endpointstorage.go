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

// EndpointStorageService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEndpointStorageService] method instead.
type EndpointStorageService struct {
	Options []option.RequestOption
}

// NewEndpointStorageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEndpointStorageService(opts ...option.RequestOption) (r EndpointStorageService) {
	r = EndpointStorageService{}
	r.Options = opts
	return
}

// Create a shared volume.
func (r *EndpointStorageService) NewSharedVolume(ctx context.Context, body EndpointStorageNewSharedVolumeParams, opts ...option.RequestOption) (res *EndpointStorageNewSharedVolumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete shared volume by volume id.
func (r *EndpointStorageService) DeleteSharedVolume(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *EndpointStorageDeleteSharedVolumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("clusters/storages/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List all shared volumes.
func (r *EndpointStorageService) ListSharedVolumes(ctx context.Context, opts ...option.RequestOption) (res *EndpointStorageListSharedVolumesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get shared volume by volume Id.
func (r *EndpointStorageService) GetSharedVolume(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *ClusterStorage, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("clusters/storages/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a shared volume.
func (r *EndpointStorageService) UpdateSharedVolume(ctx context.Context, body EndpointStorageUpdateSharedVolumeParams, opts ...option.RequestOption) (res *ClusterStorage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type EndpointStorageNewSharedVolumeResponse struct {
	VolumeID string `json:"volume_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VolumeID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointStorageNewSharedVolumeResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointStorageNewSharedVolumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointStorageDeleteSharedVolumeResponse struct {
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointStorageDeleteSharedVolumeResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointStorageDeleteSharedVolumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointStorageListSharedVolumesResponse struct {
	Volumes []ClusterStorage `json:"volumes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Volumes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointStorageListSharedVolumesResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointStorageListSharedVolumesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointStorageNewSharedVolumeParams struct {
	// Region name. Usable regions can be found from `client.clusters.list_regions()`
	Region string `json:"region,required"`
	// Volume size in whole tebibytes (TiB).
	SizeTib    int64  `json:"size_tib,required"`
	VolumeName string `json:"volume_name,required"`
	paramObj
}

func (r EndpointStorageNewSharedVolumeParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointStorageNewSharedVolumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointStorageNewSharedVolumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointStorageUpdateSharedVolumeParams struct {
	SizeTib  param.Opt[int64]  `json:"size_tib,omitzero"`
	VolumeID param.Opt[string] `json:"volume_id,omitzero"`
	paramObj
}

func (r EndpointStorageUpdateSharedVolumeParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointStorageUpdateSharedVolumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointStorageUpdateSharedVolumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
