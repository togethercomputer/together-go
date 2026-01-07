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
	"time"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// EndpointService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEndpointService] method instead.
type EndpointService struct {
	Options  []option.RequestOption
	Storages EndpointStorageService
}

// NewEndpointService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEndpointService(opts ...option.RequestOption) (r EndpointService) {
	r = EndpointService{}
	r.Options = opts
	r.Storages = NewEndpointStorageService(opts...)
	return
}

// Creates a new dedicated endpoint for serving models. The endpoint will
// automatically start after creation. You can deploy any supported model on
// hardware configurations that meet the model's requirements.
func (r *EndpointService) New(ctx context.Context, body EndpointNewParams, opts ...option.RequestOption) (res *DedicatedEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves details about a specific endpoint, including its current state,
// configuration, and scaling settings.
func (r *EndpointService) Get(ctx context.Context, endpointID string, opts ...option.RequestOption) (res *DedicatedEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return
	}
	path := fmt.Sprintf("endpoints/%s", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing endpoint's configuration. You can modify the display name,
// autoscaling settings, or change the endpoint's state (start/stop).
func (r *EndpointService) Update(ctx context.Context, endpointID string, body EndpointUpdateParams, opts ...option.RequestOption) (res *DedicatedEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return
	}
	path := fmt.Sprintf("endpoints/%s", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of all endpoints associated with your account. You can filter the
// results by type (dedicated or serverless).
func (r *EndpointService) List(ctx context.Context, query EndpointListParams, opts ...option.RequestOption) (res *EndpointListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently deletes an endpoint. This action cannot be undone.
func (r *EndpointService) Delete(ctx context.Context, endpointID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return
	}
	path := fmt.Sprintf("endpoints/%s", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Create GPU Cluster
func (r *EndpointService) NewCluster(ctx context.Context, body EndpointNewClusterParams, opts ...option.RequestOption) (res *EndpointNewClusterResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete GPU cluster by cluster ID
func (r *EndpointService) DeleteCluster(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *EndpointDeleteClusterResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("clusters/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List all available availability zones.
func (r *EndpointService) ListAvzones(ctx context.Context, opts ...option.RequestOption) (res *EndpointListAvzonesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters/availability-zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all GPU clusters.
func (r *EndpointService) ListClusters(ctx context.Context, opts ...option.RequestOption) (res *EndpointListClustersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List regions and corresponding supported driver versions
func (r *EndpointService) ListRegions(ctx context.Context, opts ...option.RequestOption) (res *EndpointListRegionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clusters/regions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get GPU cluster by cluster ID
func (r *EndpointService) GetCluster(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *Cluster, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("clusters/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a GPU Cluster.
func (r *EndpointService) UpdateCluster(ctx context.Context, clusterID string, body EndpointUpdateClusterParams, opts ...option.RequestOption) (res *EndpointUpdateClusterResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("clusters/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Configuration for automatic scaling of replicas based on demand.
type Autoscaling struct {
	// The maximum number of replicas to scale up to under load
	MaxReplicas int64 `json:"max_replicas,required"`
	// The minimum number of replicas to maintain, even when there is no load
	MinReplicas int64 `json:"min_replicas,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxReplicas respjson.Field
		MinReplicas respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Autoscaling) RawJSON() string { return r.JSON.raw }
func (r *Autoscaling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Autoscaling to a AutoscalingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AutoscalingParam.Overrides()
func (r Autoscaling) ToParam() AutoscalingParam {
	return param.Override[AutoscalingParam](json.RawMessage(r.RawJSON()))
}

// Configuration for automatic scaling of replicas based on demand.
//
// The properties MaxReplicas, MinReplicas are required.
type AutoscalingParam struct {
	// The maximum number of replicas to scale up to under load
	MaxReplicas int64 `json:"max_replicas,required"`
	// The minimum number of replicas to maintain, even when there is no load
	MinReplicas int64 `json:"min_replicas,required"`
	paramObj
}

func (r AutoscalingParam) MarshalJSON() (data []byte, err error) {
	type shadow AutoscalingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutoscalingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a dedicated endpoint deployment
type DedicatedEndpoint struct {
	// Unique identifier for the endpoint
	ID string `json:"id,required"`
	// Configuration for automatic scaling of the endpoint
	Autoscaling Autoscaling `json:"autoscaling,required"`
	// Timestamp when the endpoint was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable name for the endpoint
	DisplayName string `json:"display_name,required"`
	// The hardware configuration used for this endpoint
	Hardware string `json:"hardware,required"`
	// The model deployed on this endpoint
	Model string `json:"model,required"`
	// System name for the endpoint
	Name string `json:"name,required"`
	// The type of object
	//
	// Any of "endpoint".
	Object DedicatedEndpointObject `json:"object,required"`
	// The owner of this endpoint
	Owner string `json:"owner,required"`
	// Current state of the endpoint
	//
	// Any of "PENDING", "STARTING", "STARTED", "STOPPING", "STOPPED", "ERROR".
	State DedicatedEndpointState `json:"state,required"`
	// The type of endpoint
	//
	// Any of "dedicated".
	Type DedicatedEndpointType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Autoscaling respjson.Field
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		Hardware    respjson.Field
		Model       respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedEndpoint) RawJSON() string { return r.JSON.raw }
func (r *DedicatedEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of object
type DedicatedEndpointObject string

const (
	DedicatedEndpointObjectEndpoint DedicatedEndpointObject = "endpoint"
)

// Current state of the endpoint
type DedicatedEndpointState string

const (
	DedicatedEndpointStatePending  DedicatedEndpointState = "PENDING"
	DedicatedEndpointStateStarting DedicatedEndpointState = "STARTING"
	DedicatedEndpointStateStarted  DedicatedEndpointState = "STARTED"
	DedicatedEndpointStateStopping DedicatedEndpointState = "STOPPING"
	DedicatedEndpointStateStopped  DedicatedEndpointState = "STOPPED"
	DedicatedEndpointStateError    DedicatedEndpointState = "ERROR"
)

// The type of endpoint
type DedicatedEndpointType string

const (
	DedicatedEndpointTypeDedicated DedicatedEndpointType = "dedicated"
)

type EndpointListResponse struct {
	Data []EndpointListResponseData `json:"data,required"`
	// Any of "list".
	Object EndpointListResponseObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointListResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about an endpoint when listed via the list endpoint
type EndpointListResponseData struct {
	// Unique identifier for the endpoint
	ID string `json:"id,required"`
	// Timestamp when the endpoint was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The model deployed on this endpoint
	Model string `json:"model,required"`
	// System name for the endpoint
	Name string `json:"name,required"`
	// The type of object
	//
	// Any of "endpoint".
	Object string `json:"object,required"`
	// The owner of this endpoint
	Owner string `json:"owner,required"`
	// Current state of the endpoint
	//
	// Any of "PENDING", "STARTING", "STARTED", "STOPPING", "STOPPED", "ERROR".
	State string `json:"state,required"`
	// The type of endpoint
	//
	// Any of "serverless", "dedicated".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Model       respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointListResponseData) RawJSON() string { return r.JSON.raw }
func (r *EndpointListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointListResponseObject string

const (
	EndpointListResponseObjectList EndpointListResponseObject = "list"
)

type EndpointNewClusterResponse struct {
	ClusterID string `json:"cluster_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointNewClusterResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointNewClusterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointDeleteClusterResponse struct {
	ClusterID string `json:"cluster_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointDeleteClusterResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointDeleteClusterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of unique availability zones
type EndpointListAvzonesResponse struct {
	Avzones []string `json:"avzones,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Avzones     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointListAvzonesResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointListAvzonesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointListClustersResponse struct {
	Clusters []Cluster `json:"clusters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clusters    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointListClustersResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointListClustersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointListRegionsResponse struct {
	Regions []EndpointListRegionsResponseRegion `json:"regions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Regions     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointListRegionsResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointListRegionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointListRegionsResponseRegion struct {
	ID                string   `json:"id,required"`
	AvailabilityZones []string `json:"availability_zones,required"`
	DriverVersions    []string `json:"driver_versions,required"`
	Name              string   `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AvailabilityZones respjson.Field
		DriverVersions    respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointListRegionsResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *EndpointListRegionsResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointUpdateClusterResponse struct {
	ClusterID string `json:"cluster_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointUpdateClusterResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointUpdateClusterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointNewParams struct {
	// Configuration for automatic scaling of the endpoint
	Autoscaling AutoscalingParam `json:"autoscaling,omitzero,required"`
	// The hardware configuration to use for this endpoint
	Hardware string `json:"hardware,required"`
	// The model to deploy on this endpoint
	Model string `json:"model,required"`
	// The number of minutes of inactivity after which the endpoint will be
	// automatically stopped. Set to null, omit or set to 0 to disable automatic
	// timeout.
	InactiveTimeout param.Opt[int64] `json:"inactive_timeout,omitzero"`
	// Create the endpoint in a specified availability zone (e.g., us-central-4b)
	AvailabilityZone param.Opt[string] `json:"availability_zone,omitzero"`
	// Whether to disable the prompt cache for this endpoint
	DisablePromptCache param.Opt[bool] `json:"disable_prompt_cache,omitzero"`
	// Whether to disable speculative decoding for this endpoint
	DisableSpeculativeDecoding param.Opt[bool] `json:"disable_speculative_decoding,omitzero"`
	// A human-readable name for the endpoint
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// The desired state of the endpoint
	//
	// Any of "STARTED", "STOPPED".
	State EndpointNewParamsState `json:"state,omitzero"`
	paramObj
}

func (r EndpointNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The desired state of the endpoint
type EndpointNewParamsState string

const (
	EndpointNewParamsStateStarted EndpointNewParamsState = "STARTED"
	EndpointNewParamsStateStopped EndpointNewParamsState = "STOPPED"
)

type EndpointUpdateParams struct {
	// The number of minutes of inactivity after which the endpoint will be
	// automatically stopped. Set to 0 to disable automatic timeout.
	InactiveTimeout param.Opt[int64] `json:"inactive_timeout,omitzero"`
	// A human-readable name for the endpoint
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// New autoscaling configuration for the endpoint
	Autoscaling AutoscalingParam `json:"autoscaling,omitzero"`
	// The desired state of the endpoint
	//
	// Any of "STARTED", "STOPPED".
	State EndpointUpdateParamsState `json:"state,omitzero"`
	paramObj
}

func (r EndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The desired state of the endpoint
type EndpointUpdateParamsState string

const (
	EndpointUpdateParamsStateStarted EndpointUpdateParamsState = "STARTED"
	EndpointUpdateParamsStateStopped EndpointUpdateParamsState = "STOPPED"
)

type EndpointListParams struct {
	// If true, return only endpoints owned by the caller
	Mine param.Opt[bool] `query:"mine,omitzero" json:"-"`
	// Filter endpoints by type
	//
	// Any of "dedicated", "serverless".
	Type EndpointListParamsType `query:"type,omitzero" json:"-"`
	// Filter endpoints by usage type
	//
	// Any of "on-demand", "reserved".
	UsageType EndpointListParamsUsageType `query:"usage_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EndpointListParams]'s query parameters as `url.Values`.
func (r EndpointListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter endpoints by type
type EndpointListParamsType string

const (
	EndpointListParamsTypeDedicated  EndpointListParamsType = "dedicated"
	EndpointListParamsTypeServerless EndpointListParamsType = "serverless"
)

// Filter endpoints by usage type
type EndpointListParamsUsageType string

const (
	EndpointListParamsUsageTypeOnDemand EndpointListParamsUsageType = "on-demand"
	EndpointListParamsUsageTypeReserved EndpointListParamsUsageType = "reserved"
)

type EndpointNewClusterParams struct {
	// Any of "RESERVED", "ON_DEMAND".
	BillingType EndpointNewClusterParamsBillingType `json:"billing_type,omitzero,required"`
	// Name of the GPU cluster.
	ClusterName string `json:"cluster_name,required"`
	// NVIDIA driver version to use in the cluster.
	//
	// Any of "CUDA_12_5_555", "CUDA_12_6_560", "CUDA_12_6_565", "CUDA_12_8_570".
	DriverVersion EndpointNewClusterParamsDriverVersion `json:"driver_version,omitzero,required"`
	// Duration in days to keep the cluster running.
	DurationDays int64 `json:"duration_days,required"`
	// Type of GPU to use in the cluster
	//
	// Any of "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM",
	// "H100_SXM_INF".
	GPUType EndpointNewClusterParamsGPUType `json:"gpu_type,omitzero,required"`
	// Number of GPUs to allocate in the cluster. This must be multiple of 8. For
	// example, 8, 16 or 24
	NumGPUs int64 `json:"num_gpus,required"`
	// Region to create the GPU cluster in. Valid values are us-central-8 and
	// us-central-4.
	//
	// Any of "us-central-8", "us-central-4".
	Region   EndpointNewClusterParamsRegion `json:"region,omitzero,required"`
	VolumeID param.Opt[string]              `json:"volume_id,omitzero"`
	// Any of "KUBERNETES", "SLURM".
	ClusterType  EndpointNewClusterParamsClusterType  `json:"cluster_type,omitzero"`
	SharedVolume EndpointNewClusterParamsSharedVolume `json:"shared_volume,omitzero"`
	paramObj
}

func (r EndpointNewClusterParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointNewClusterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointNewClusterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointNewClusterParamsBillingType string

const (
	EndpointNewClusterParamsBillingTypeReserved EndpointNewClusterParamsBillingType = "RESERVED"
	EndpointNewClusterParamsBillingTypeOnDemand EndpointNewClusterParamsBillingType = "ON_DEMAND"
)

// NVIDIA driver version to use in the cluster.
type EndpointNewClusterParamsDriverVersion string

const (
	EndpointNewClusterParamsDriverVersionCuda12_5_555 EndpointNewClusterParamsDriverVersion = "CUDA_12_5_555"
	EndpointNewClusterParamsDriverVersionCuda12_6_560 EndpointNewClusterParamsDriverVersion = "CUDA_12_6_560"
	EndpointNewClusterParamsDriverVersionCuda12_6_565 EndpointNewClusterParamsDriverVersion = "CUDA_12_6_565"
	EndpointNewClusterParamsDriverVersionCuda12_8_570 EndpointNewClusterParamsDriverVersion = "CUDA_12_8_570"
)

// Type of GPU to use in the cluster
type EndpointNewClusterParamsGPUType string

const (
	EndpointNewClusterParamsGPUTypeH100Sxm    EndpointNewClusterParamsGPUType = "H100_SXM"
	EndpointNewClusterParamsGPUTypeH200Sxm    EndpointNewClusterParamsGPUType = "H200_SXM"
	EndpointNewClusterParamsGPUTypeRtx6000Pci EndpointNewClusterParamsGPUType = "RTX_6000_PCI"
	EndpointNewClusterParamsGPUTypeL40Pcie    EndpointNewClusterParamsGPUType = "L40_PCIE"
	EndpointNewClusterParamsGPUTypeB200Sxm    EndpointNewClusterParamsGPUType = "B200_SXM"
	EndpointNewClusterParamsGPUTypeH100SxmInf EndpointNewClusterParamsGPUType = "H100_SXM_INF"
)

// Region to create the GPU cluster in. Valid values are us-central-8 and
// us-central-4.
type EndpointNewClusterParamsRegion string

const (
	EndpointNewClusterParamsRegionUsCentral8 EndpointNewClusterParamsRegion = "us-central-8"
	EndpointNewClusterParamsRegionUsCentral4 EndpointNewClusterParamsRegion = "us-central-4"
)

type EndpointNewClusterParamsClusterType string

const (
	EndpointNewClusterParamsClusterTypeKubernetes EndpointNewClusterParamsClusterType = "KUBERNETES"
	EndpointNewClusterParamsClusterTypeSlurm      EndpointNewClusterParamsClusterType = "SLURM"
)

// The properties Region, SizeTib, VolumeName are required.
type EndpointNewClusterParamsSharedVolume struct {
	// Region name. Usable regions can be found from `client.clusters.list_regions()`
	Region string `json:"region,required"`
	// Volume size in whole tebibytes (TiB).
	SizeTib    int64  `json:"size_tib,required"`
	VolumeName string `json:"volume_name,required"`
	paramObj
}

func (r EndpointNewClusterParamsSharedVolume) MarshalJSON() (data []byte, err error) {
	type shadow EndpointNewClusterParamsSharedVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointNewClusterParamsSharedVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointUpdateClusterParams struct {
	NumGPUs param.Opt[int64] `json:"num_gpus,omitzero"`
	// Any of "KUBERNETES", "SLURM".
	ClusterType EndpointUpdateClusterParamsClusterType `json:"cluster_type,omitzero"`
	paramObj
}

func (r EndpointUpdateClusterParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointUpdateClusterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointUpdateClusterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointUpdateClusterParamsClusterType string

const (
	EndpointUpdateClusterParamsClusterTypeKubernetes EndpointUpdateClusterParamsClusterType = "KUBERNETES"
	EndpointUpdateClusterParamsClusterTypeSlurm      EndpointUpdateClusterParamsClusterType = "SLURM"
)
