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
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaClusterService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaClusterService] method instead.
type BetaClusterService struct {
	Options []option.RequestOption
	Storage BetaClusterStorageService
}

// NewBetaClusterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaClusterService(opts ...option.RequestOption) (r BetaClusterService) {
	r = BetaClusterService{}
	r.Options = opts
	r.Storage = NewBetaClusterStorageService(opts...)
	return
}

// Create an Instant Cluster on Together's high-performance GPU clusters. With
// features like on-demand scaling, long-lived resizable high-bandwidth shared
// DC-local storage, Kubernetes and Slurm cluster flavors, a REST API, and
// Terraform support, you can run workloads flexibly without complex infrastructure
// management.
func (r *BetaClusterService) New(ctx context.Context, body BetaClusterNewParams, opts ...option.RequestOption) (res *Cluster, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve information about a specific GPU cluster.
func (r *BetaClusterService) Get(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *Cluster, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update the configuration of an existing GPU cluster.
func (r *BetaClusterService) Update(ctx context.Context, clusterID string, body BetaClusterUpdateParams, opts ...option.RequestOption) (res *Cluster, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List all GPU clusters.
func (r *BetaClusterService) List(ctx context.Context, opts ...option.RequestOption) (res *BetaClusterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a GPU cluster by cluster ID.
func (r *BetaClusterService) Delete(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *BetaClusterDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List regions and corresponding supported driver versions
func (r *BetaClusterService) ListRegions(ctx context.Context, opts ...option.RequestOption) (res *BetaClusterListRegionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/regions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Cluster struct {
	ClusterID   string `json:"cluster_id" api:"required"`
	ClusterName string `json:"cluster_name" api:"required"`
	// Type of cluster.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType       ClusterClusterType        `json:"cluster_type" api:"required"`
	ControlPlaneNodes []ClusterControlPlaneNode `json:"control_plane_nodes" api:"required"`
	CudaVersion       string                    `json:"cuda_version" api:"required"`
	// Any of "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM",
	// "H100_SXM_INF".
	GPUType             ClusterGPUType         `json:"gpu_type" api:"required"`
	GPUWorkerNodes      []ClusterGPUWorkerNode `json:"gpu_worker_nodes" api:"required"`
	KubeConfig          string                 `json:"kube_config" api:"required"`
	NumGPUs             int64                  `json:"num_gpus" api:"required"`
	NvidiaDriverVersion string                 `json:"nvidia_driver_version" api:"required"`
	Region              string                 `json:"region" api:"required"`
	// Current status of the GPU cluster.
	//
	// Any of "WaitingForControlPlaneNodes", "WaitingForDataPlaneNodes",
	// "WaitingForSubnet", "WaitingForSharedVolume", "InstallingDrivers",
	// "RunningAcceptanceTests", "Paused", "OnDemandComputePaused", "Ready",
	// "Degraded", "Deleting".
	Status               ClusterStatus   `json:"status" api:"required"`
	Volumes              []ClusterVolume `json:"volumes" api:"required"`
	CapacityPoolID       string          `json:"capacity_pool_id"`
	CreatedAt            time.Time       `json:"created_at" format:"date-time"`
	DurationHours        int64           `json:"duration_hours"`
	InstallTraefik       bool            `json:"install_traefik"`
	ReservationEndTime   time.Time       `json:"reservation_end_time" format:"date-time"`
	ReservationStartTime time.Time       `json:"reservation_start_time" format:"date-time"`
	SlurmShmSizeGib      int64           `json:"slurm_shm_size_gib"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterID            respjson.Field
		ClusterName          respjson.Field
		ClusterType          respjson.Field
		ControlPlaneNodes    respjson.Field
		CudaVersion          respjson.Field
		GPUType              respjson.Field
		GPUWorkerNodes       respjson.Field
		KubeConfig           respjson.Field
		NumGPUs              respjson.Field
		NvidiaDriverVersion  respjson.Field
		Region               respjson.Field
		Status               respjson.Field
		Volumes              respjson.Field
		CapacityPoolID       respjson.Field
		CreatedAt            respjson.Field
		DurationHours        respjson.Field
		InstallTraefik       respjson.Field
		ReservationEndTime   respjson.Field
		ReservationStartTime respjson.Field
		SlurmShmSizeGib      respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Cluster) RawJSON() string { return r.JSON.raw }
func (r *Cluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cluster.
type ClusterClusterType string

const (
	ClusterClusterTypeKubernetes ClusterClusterType = "KUBERNETES"
	ClusterClusterTypeSlurm      ClusterClusterType = "SLURM"
)

type ClusterControlPlaneNode struct {
	HostName    string  `json:"host_name" api:"required"`
	MemoryGib   float64 `json:"memory_gib" api:"required"`
	Network     string  `json:"network" api:"required"`
	NodeID      string  `json:"node_id" api:"required"`
	NodeName    string  `json:"node_name" api:"required"`
	NumCPUCores int64   `json:"num_cpu_cores" api:"required"`
	Status      string  `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HostName    respjson.Field
		MemoryGib   respjson.Field
		Network     respjson.Field
		NodeID      respjson.Field
		NodeName    respjson.Field
		NumCPUCores respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterControlPlaneNode) RawJSON() string { return r.JSON.raw }
func (r *ClusterControlPlaneNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterGPUType string

const (
	ClusterGPUTypeH100Sxm    ClusterGPUType = "H100_SXM"
	ClusterGPUTypeH200Sxm    ClusterGPUType = "H200_SXM"
	ClusterGPUTypeRtx6000Pci ClusterGPUType = "RTX_6000_PCI"
	ClusterGPUTypeL40Pcie    ClusterGPUType = "L40_PCIE"
	ClusterGPUTypeB200Sxm    ClusterGPUType = "B200_SXM"
	ClusterGPUTypeH100SxmInf ClusterGPUType = "H100_SXM_INF"
)

type ClusterGPUWorkerNode struct {
	HostName    string   `json:"host_name" api:"required"`
	MemoryGib   float64  `json:"memory_gib" api:"required"`
	Networks    []string `json:"networks" api:"required"`
	NodeID      string   `json:"node_id" api:"required"`
	NodeName    string   `json:"node_name" api:"required"`
	NumCPUCores int64    `json:"num_cpu_cores" api:"required"`
	NumGPUs     int64    `json:"num_gpus" api:"required"`
	Status      string   `json:"status" api:"required"`
	InstanceID  string   `json:"instance_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HostName    respjson.Field
		MemoryGib   respjson.Field
		Networks    respjson.Field
		NodeID      respjson.Field
		NodeName    respjson.Field
		NumCPUCores respjson.Field
		NumGPUs     respjson.Field
		Status      respjson.Field
		InstanceID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterGPUWorkerNode) RawJSON() string { return r.JSON.raw }
func (r *ClusterGPUWorkerNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the GPU cluster.
type ClusterStatus string

const (
	ClusterStatusWaitingForControlPlaneNodes ClusterStatus = "WaitingForControlPlaneNodes"
	ClusterStatusWaitingForDataPlaneNodes    ClusterStatus = "WaitingForDataPlaneNodes"
	ClusterStatusWaitingForSubnet            ClusterStatus = "WaitingForSubnet"
	ClusterStatusWaitingForSharedVolume      ClusterStatus = "WaitingForSharedVolume"
	ClusterStatusInstallingDrivers           ClusterStatus = "InstallingDrivers"
	ClusterStatusRunningAcceptanceTests      ClusterStatus = "RunningAcceptanceTests"
	ClusterStatusPaused                      ClusterStatus = "Paused"
	ClusterStatusOnDemandComputePaused       ClusterStatus = "OnDemandComputePaused"
	ClusterStatusReady                       ClusterStatus = "Ready"
	ClusterStatusDegraded                    ClusterStatus = "Degraded"
	ClusterStatusDeleting                    ClusterStatus = "Deleting"
)

type ClusterVolume struct {
	SizeTib    int64  `json:"size_tib" api:"required"`
	Status     string `json:"status" api:"required"`
	VolumeID   string `json:"volume_id" api:"required"`
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
func (r ClusterVolume) RawJSON() string { return r.JSON.raw }
func (r *ClusterVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterListResponse struct {
	Clusters []Cluster `json:"clusters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clusters    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterDeleteResponse struct {
	ClusterID string `json:"cluster_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterListRegionsResponse struct {
	Regions []BetaClusterListRegionsResponseRegion `json:"regions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Regions     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterListRegionsResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterListRegionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterListRegionsResponseRegion struct {
	// List of supported identifiable cuda/nvidia driver versions pairs available in
	// the region.
	DriverVersions []BetaClusterListRegionsResponseRegionDriverVersion `json:"driver_versions" api:"required"`
	// Identifiable name of the region.
	Name string `json:"name" api:"required"`
	// List of supported identifiable gpus available in the region.
	SupportedInstanceTypes []string `json:"supported_instance_types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DriverVersions         respjson.Field
		Name                   respjson.Field
		SupportedInstanceTypes respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterListRegionsResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterListRegionsResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CUDA/NVIDIA driver versions pair available in the region to use in the create
// cluster request.
type BetaClusterListRegionsResponseRegionDriverVersion struct {
	// CUDA driver version.
	CudaVersion string `json:"cuda_version" api:"required"`
	// NVIDIA driver version.
	NvidiaDriverVersion string `json:"nvidia_driver_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CudaVersion         respjson.Field
		NvidiaDriverVersion respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterListRegionsResponseRegionDriverVersion) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterListRegionsResponseRegionDriverVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParams struct {
	// RESERVED billing types allow you to specify the duration of the cluster
	// reservation via the duration_days field. ON_DEMAND billing types will give you
	// ownership of the cluster until you delete it. SCHEDULED_CAPACITY billing types
	// allow you to reserve capacity for a scheduled time window. You must specify the
	// reservation_start_time and reservation_end_time with this request.
	//
	// Any of "RESERVED", "ON_DEMAND", "SCHEDULED_CAPACITY".
	BillingType BetaClusterNewParamsBillingType `json:"billing_type,omitzero" api:"required"`
	// Name of the GPU cluster.
	ClusterName string `json:"cluster_name" api:"required"`
	// CUDA version for this cluster. For example, 12.5
	CudaVersion string `json:"cuda_version" api:"required"`
	// Type of GPU to use in the cluster
	//
	// Any of "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM",
	// "H100_SXM_INF".
	GPUType BetaClusterNewParamsGPUType `json:"gpu_type,omitzero" api:"required"`
	// Number of GPUs to allocate in the cluster. This must be multiple of 8. For
	// example, 8, 16 or 24
	NumGPUs int64 `json:"num_gpus" api:"required"`
	// Nvidia driver version for this cluster. For example, 550. Only some combination
	// of cuda_version and nvidia_driver_version are supported.
	NvidiaDriverVersion string `json:"nvidia_driver_version" api:"required"`
	// Region to create the GPU cluster in. Usable regions can be found from
	// `client.clusters.list_regions()`
	Region string `json:"region" api:"required"`
	// Maximum number of GPUs to which the cluster can be auto-scaled up. This field is
	// required if auto_scaled is true.
	AutoScaleMaxGPUs param.Opt[int64] `json:"auto_scale_max_gpus,omitzero"`
	// Whether GPU cluster should be auto-scaled based on the workload. By default, it
	// is not auto-scaled.
	AutoScaled param.Opt[bool] `json:"auto_scaled,omitzero"`
	// ID of the capacity pool to use for the cluster. This field is optional and only
	// applicable if the cluster is created from a capacity pool.
	CapacityPoolID param.Opt[string] `json:"capacity_pool_id,omitzero"`
	// Duration in days to keep the cluster running.
	DurationDays param.Opt[int64] `json:"duration_days,omitzero"`
	// Whether automated GPU node failover should be enabled for this cluster. By
	// default, it is disabled.
	GPUNodeFailoverEnabled param.Opt[bool] `json:"gpu_node_failover_enabled,omitzero"`
	// Whether to install Traefik ingress controller in the cluster. This field is only
	// applicable for Kubernetes clusters and is false by default.
	InstallTraefik param.Opt[bool] `json:"install_traefik,omitzero"`
	// Reservation end time of the cluster. This field is required for SCHEDULED
	// billing to specify the reservation end time for the cluster.
	ReservationEndTime param.Opt[time.Time] `json:"reservation_end_time,omitzero" format:"date-time"`
	// Reservation start time of the cluster. This field is required for SCHEDULED
	// billing to specify the reservation start time for the cluster. If not provided,
	// the cluster will be provisioned immediately.
	ReservationStartTime param.Opt[time.Time] `json:"reservation_start_time,omitzero" format:"date-time"`
	// Custom Slurm image for Slurm clusters.
	SlurmImage param.Opt[string] `json:"slurm_image,omitzero"`
	// Shared memory size in GiB for Slurm cluster. This field is required if
	// cluster_type is SLURM.
	SlurmShmSizeGib param.Opt[int64] `json:"slurm_shm_size_gib,omitzero"`
	// ID of an existing volume to use with the cluster creation.
	VolumeID param.Opt[string] `json:"volume_id,omitzero"`
	// Type of cluster to create.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType BetaClusterNewParamsClusterType `json:"cluster_type,omitzero"`
	// Inline configuration to create a shared volume with the cluster creation.
	SharedVolume BetaClusterNewParamsSharedVolume `json:"shared_volume,omitzero"`
	paramObj
}

func (r BetaClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RESERVED billing types allow you to specify the duration of the cluster
// reservation via the duration_days field. ON_DEMAND billing types will give you
// ownership of the cluster until you delete it. SCHEDULED_CAPACITY billing types
// allow you to reserve capacity for a scheduled time window. You must specify the
// reservation_start_time and reservation_end_time with this request.
type BetaClusterNewParamsBillingType string

const (
	BetaClusterNewParamsBillingTypeReserved          BetaClusterNewParamsBillingType = "RESERVED"
	BetaClusterNewParamsBillingTypeOnDemand          BetaClusterNewParamsBillingType = "ON_DEMAND"
	BetaClusterNewParamsBillingTypeScheduledCapacity BetaClusterNewParamsBillingType = "SCHEDULED_CAPACITY"
)

// Type of GPU to use in the cluster
type BetaClusterNewParamsGPUType string

const (
	BetaClusterNewParamsGPUTypeH100Sxm    BetaClusterNewParamsGPUType = "H100_SXM"
	BetaClusterNewParamsGPUTypeH200Sxm    BetaClusterNewParamsGPUType = "H200_SXM"
	BetaClusterNewParamsGPUTypeRtx6000Pci BetaClusterNewParamsGPUType = "RTX_6000_PCI"
	BetaClusterNewParamsGPUTypeL40Pcie    BetaClusterNewParamsGPUType = "L40_PCIE"
	BetaClusterNewParamsGPUTypeB200Sxm    BetaClusterNewParamsGPUType = "B200_SXM"
	BetaClusterNewParamsGPUTypeH100SxmInf BetaClusterNewParamsGPUType = "H100_SXM_INF"
)

// Type of cluster to create.
type BetaClusterNewParamsClusterType string

const (
	BetaClusterNewParamsClusterTypeKubernetes BetaClusterNewParamsClusterType = "KUBERNETES"
	BetaClusterNewParamsClusterTypeSlurm      BetaClusterNewParamsClusterType = "SLURM"
)

// Inline configuration to create a shared volume with the cluster creation.
//
// The properties Region, SizeTib, VolumeName are required.
type BetaClusterNewParamsSharedVolume struct {
	// Region name. Usable regions can be found from `client.clusters.list_regions()`
	Region string `json:"region" api:"required"`
	// Volume size in whole tebibytes (TiB).
	SizeTib int64 `json:"size_tib" api:"required"`
	// Customizable name of the volume to create.
	VolumeName string `json:"volume_name" api:"required"`
	paramObj
}

func (r BetaClusterNewParamsSharedVolume) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsSharedVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsSharedVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterUpdateParams struct {
	// Number of GPUs to allocate in the cluster. This must be multiple of 8. For
	// example, 8, 16 or 24
	NumGPUs param.Opt[int64] `json:"num_gpus,omitzero"`
	// Timestamp at which the cluster should be decommissioned. Only accepted for
	// prepaid clusters.
	ReservationEndTime param.Opt[time.Time] `json:"reservation_end_time,omitzero" format:"date-time"`
	// Type of cluster to update.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType BetaClusterUpdateParamsClusterType `json:"cluster_type,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cluster to update.
type BetaClusterUpdateParamsClusterType string

const (
	BetaClusterUpdateParamsClusterTypeKubernetes BetaClusterUpdateParamsClusterType = "KUBERNETES"
	BetaClusterUpdateParamsClusterTypeSlurm      BetaClusterUpdateParamsClusterType = "SLURM"
)
