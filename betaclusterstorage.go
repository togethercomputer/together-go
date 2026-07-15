// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
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

// Instant Clusters supports long-lived, resizable in-DC shared storage with user
// data persistence. You can dynamically create and attach volumes to your cluster
// at cluster creation time, and resize as your data grows. All shared storage is
// backed by multi-NIC bare metal paths, ensuring high-throughput and low-latency
// performance for shared storage.
func (r *BetaClusterStorageService) New(ctx context.Context, body BetaClusterStorageNewParams, opts ...option.RequestOption) (res *ClusterStorage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/clusters/storage/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve information about a specific shared volume.
func (r *BetaClusterStorageService) Get(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *ClusterStorage, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/storage/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update the configuration of an existing shared volume.
func (r *BetaClusterStorageService) Update(ctx context.Context, body BetaClusterStorageUpdateParams, opts ...option.RequestOption) (res *ClusterStorage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/clusters/storage/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List all shared volumes.
func (r *BetaClusterStorageService) List(ctx context.Context, query BetaClusterStorageListParams, opts ...option.RequestOption) (res *BetaClusterStorageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/clusters/storage/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a shared volume. Note that if this volume is attached to a cluster,
// deleting will fail.
func (r *BetaClusterStorageService) Delete(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *BetaClusterStorageDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/storage/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ClusterStorage struct {
	// Size of the volume in TiB.
	SizeTib int64 `json:"size_tib" api:"required"`
	// Current status of the shared volume.
	//
	// Any of "scheduled", "available", "bound", "provisioning", "deleting", "failed",
	// "access_revoked", "unknown".
	Status ClusterStorageStatus `json:"status" api:"required"`
	// ID of the volume.
	VolumeID string `json:"volume_id" api:"required"`
	// User provided name of the volume.
	VolumeName string `json:"volume_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SizeTib     respjson.Field
		Status      respjson.Field
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

// Current status of the shared volume.
type ClusterStorageStatus string

const (
	ClusterStorageStatusScheduled     ClusterStorageStatus = "scheduled"
	ClusterStorageStatusAvailable     ClusterStorageStatus = "available"
	ClusterStorageStatusBound         ClusterStorageStatus = "bound"
	ClusterStorageStatusProvisioning  ClusterStorageStatus = "provisioning"
	ClusterStorageStatusDeleting      ClusterStorageStatus = "deleting"
	ClusterStorageStatusFailed        ClusterStorageStatus = "failed"
	ClusterStorageStatusAccessRevoked ClusterStorageStatus = "access_revoked"
	ClusterStorageStatusUnknown       ClusterStorageStatus = "unknown"
)

type BetaClusterStorageListResponse struct {
	Volumes []ClusterStorage `json:"volumes" api:"required"`
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
	Success bool `json:"success" api:"required"`
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
	// Region name. Usable regions can be found from `clusters.list_regions()`
	Region string `json:"region" api:"required"`
	// Volume size in whole tebibytes (TiB).
	SizeTib int64 `json:"size_tib" api:"required"`
	// User provided name of the volume.
	VolumeName string `json:"volume_name" api:"required"`
	// When true, the shared volume is not deleted when the cluster is decommissioned.
	IsLifecycleIndependent param.Opt[bool] `json:"is_lifecycle_independent,omitzero"`
	// Project ID that will own the volume. When omitted, the caller's default project
	// is used.
	ProjectID param.Opt[string] `json:"project_id,omitzero"`
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
	// ID of the volume.
	VolumeID string `json:"volume_id" api:"required"`
	// Size of the volume in TiB.
	SizeTib param.Opt[int64] `json:"size_tib,omitzero"`
	paramObj
}

func (r BetaClusterStorageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterStorageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterStorageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterStorageListParams struct {
	// Optional UMS project ID to filter volumes by. When set, only volumes belonging
	// to this project are returned. The caller must be a member of the project;
	// otherwise the result set will be empty.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `query:"projectId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaClusterStorageListParams]'s query parameters as
// `url.Values`.
func (r BetaClusterStorageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
