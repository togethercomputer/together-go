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

// BetaClusterStorageService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaClusterStorageService] method instead.
type BetaClusterStorageService struct {
	Options []option.RequestOption
}

// NewBetaClusterStorageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaClusterStorageService(opts ...option.RequestOption) (r BetaClusterStorageService) {
	r = BetaClusterStorageService{}
	r.Options = opts
	return
}

// Create a shared volume.
func (r *BetaClusterStorageService) New(ctx context.Context, body BetaClusterStorageNewParams, opts ...option.RequestOption) (res *BetaClusterStorageNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get shared volume by volume Id.
func (r *BetaClusterStorageService) Get(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *ClusterStorage, err error) {
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
func (r *BetaClusterStorageService) Update(ctx context.Context, body BetaClusterStorageUpdateParams, opts ...option.RequestOption) (res *ClusterStorage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all shared volumes.
func (r *BetaClusterStorageService) List(ctx context.Context, opts ...option.RequestOption) (res *BetaClusterStorageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters/storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete shared volume by volume id.
func (r *BetaClusterStorageService) Delete(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *BetaClusterStorageDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("clusters/storages/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ClusterStorage struct {
	SizeTib    int64  `json:"size_tib,required"`
	VolumeID   string `json:"volume_id,required"`
	VolumeName string `json:"volume_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SizeTib     respjson.Field
		VolumeID    respjson.Field
		VolumeName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterStorage) RawJSON() string { return r.JSON.raw }
func (r *ClusterStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterStorageNewResponse struct {
	VolumeID string `json:"volume_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VolumeID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterStorageNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterStorageNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterStorageListResponse struct {
	Volumes []ClusterStorage `json:"volumes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Volumes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterStorageListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterStorageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterStorageDeleteResponse struct {
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterStorageDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterStorageDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterStorageNewParams struct {
	// Region name. Usable regions can be found from `client.clusters.list_regions()`
	Region string `json:"region,required"`
	// Volume size in whole tebibytes (TiB).
	SizeTib    int64  `json:"size_tib,required"`
	VolumeName string `json:"volume_name,required"`
	paramObj
}

func (r BetaClusterStorageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterStorageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterStorageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterStorageUpdateParams struct {
	SizeTib  param.Opt[int64]  `json:"size_tib,omitzero"`
	VolumeID param.Opt[string] `json:"volume_id,omitzero"`
	paramObj
}

func (r BetaClusterStorageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterStorageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterStorageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
