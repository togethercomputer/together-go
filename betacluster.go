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
	Options      []option.RequestOption
	Remediations BetaClusterRemediationService
	Storage      BetaClusterStorageService
}

// NewBetaClusterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaClusterService(opts ...option.RequestOption) (r BetaClusterService) {
	r = BetaClusterService{}
	r.Options = opts
	r.Remediations = NewBetaClusterRemediationService(opts...)
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
func (r *BetaClusterService) List(ctx context.Context, query BetaClusterListParams, opts ...option.RequestOption) (res *BetaClusterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

// NVIDIA software configuration available in the region.
type BetaClusterListRegionsResponseRegionDriverVersion struct {
	// Region-specific NVIDIA catalog ID to send as nvidia_version_id when creating a
	// cluster.
	ID string `json:"id" api:"required"`
	// Semantic CUDA version without operating system text.
	CudaVersion string `json:"cuda_version" api:"required"`
	// NVIDIA driver version.
	NvidiaDriverVersion string `json:"nvidia_driver_version" api:"required"`
	// Operating system image family for this catalog entry.
	Os string `json:"os" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CudaVersion         respjson.Field
		NvidiaDriverVersion respjson.Field
		Os                  respjson.Field
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

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Create
	// a cluster with a canonical NVIDIA version id. Do not also set cuda_version or
	// nvidia_driver_version.
	OfGPUClusterCreateRequestNvidiaVersion *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Create
	// a cluster with the legacy CUDA and NVIDIA driver selectors. nvidia_version_id
	// may also be set when it resolves to the same catalog entry.
	OfGPUClusterCreateRequestLegacyNvidia *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia `json:",inline"`

	paramObj
}

func (u BetaClusterNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGPUClusterCreateRequestNvidiaVersion, u.OfGPUClusterCreateRequestLegacyNvidia)
}
func (r *BetaClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create a cluster with a canonical NVIDIA version id. Do not also set
// cuda_version or nvidia_driver_version.
//
// The properties BillingType, ClusterName, GPUType, NumGPUs, NvidiaVersionID,
// Region are required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion struct {
	// RESERVED billing types allow you to specify the duration of the cluster
	// reservation via the duration_days field. ON_DEMAND billing types will give you
	// ownership of the cluster until you delete it. SCHEDULED_CAPACITY billing types
	// allow you to reserve capacity for a scheduled time window. You must specify the
	// reservation_start_time and reservation_end_time with this request.
	//
	// Any of "RESERVED", "ON_DEMAND", "SCHEDULED_CAPACITY".
	BillingType string `json:"billing_type,omitzero" api:"required"`
	// Name of the GPU cluster.
	ClusterName string `json:"cluster_name" api:"required"`
	// Type of GPU to use in the cluster
	//
	// Any of "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM",
	// "H100_SXM_INF", "B300_SXM".
	GPUType string `json:"gpu_type,omitzero" api:"required"`
	// Number of GPUs to allocate in the cluster. This must be multiple of 8. For
	// example, 8, 16 or 24
	NumGPUs int64 `json:"num_gpus" api:"required"`
	// Canonical region-specific NVIDIA version ID. If cuda_version and
	// nvidia_driver_version are also set, they must resolve to the same catalog entry.
	NvidiaVersionID string `json:"nvidia_version_id" api:"required"`
	// Region to create the GPU cluster in. Usable regions can be found from
	// `client.clusters.list_regions()`
	Region string `json:"region" api:"required"`
	// Whether to enable auto-scaling for the cluster. If true, the cluster will
	// automatically scale the number of GPU worker nodes between num_gpus and
	// auto_scale_max_gpus based on the workload.
	AutoScale param.Opt[bool] `json:"auto_scale,omitzero"`
	// Maximum number of GPUs to which the cluster can be auto-scaled up. This field is
	// required if auto_scaled is true.
	AutoScaleMaxGPUs param.Opt[int64] `json:"auto_scale_max_gpus,omitzero"`
	// Whether GPU cluster should be auto-scaled based on the workload. By default, it
	// is not auto-scaled.
	//
	// Deprecated: deprecated
	AutoScaled param.Opt[bool] `json:"auto_scaled,omitzero"`
	// ID of the capacity pool to use for the cluster. This field is optional and only
	// applicable if the cluster is created from a capacity pool.
	CapacityPoolID param.Opt[string] `json:"capacity_pool_id,omitzero"`
	// Legacy CUDA selector for this cluster. Bare semantic values such as 12.5 select
	// ubuntu-22.04; existing OS-suffixed values remain accepted for compatibility.
	// Must be paired with nvidia_driver_version. Prefer nvidia_version_id for new
	// integrations.
	CudaVersion param.Opt[string] `json:"cuda_version,omitzero"`
	// Duration in days to keep the cluster running.
	DurationDays param.Opt[int64] `json:"duration_days,omitzero"`
	// Whether to install Traefik ingress controller in the cluster. This field is only
	// applicable for Kubernetes clusters and is false by default.
	InstallTraefik param.Opt[bool] `json:"install_traefik,omitzero"`
	// Number of GPUs to allocate from the capacity pool. Must be a multiple of 8 and
	// not exceed num_gpus.
	NumCapacityPoolGPUs param.Opt[int64] `json:"num_capacity_pool_gpus,omitzero"`
	// Number of preemptible GPUs to request alongside on-demand capacity. Must be a
	// multiple of 8. Preemptible nodes are cheaper but may be reclaimed when on-demand
	// capacity is needed elsewhere; the system fulfills this asynchronously and
	// surfaces the actual count in allocated_preemptible_gpus.
	NumPreemptibleGPUs param.Opt[int64] `json:"num_preemptible_gpus,omitzero"`
	// Number of prepaid (PLG) reserved GPUs for this cluster. When omitted for
	// RESERVED billing on create, the server defaults this to num_gpus.
	NumReservedGPUs param.Opt[int64] `json:"num_reserved_gpus,omitzero"`
	// Legacy NVIDIA driver selector for this cluster. For example, 550. Must be paired
	// with cuda_version. Prefer nvidia_version_id for new integrations.
	NvidiaDriverVersion param.Opt[string] `json:"nvidia_driver_version,omitzero"`
	// Project ID for the cluster. If not set, the project from the request context is
	// used.
	ProjectID param.Opt[string] `json:"project_id,omitzero"`
	// Reservation end time of the cluster. This field is required for SCHEDULED
	// billing to specify the reservation end time for the cluster.
	ReservationEndTime param.Opt[time.Time] `json:"reservation_end_time,omitzero" format:"date-time"`
	// Reservation start time of the cluster. This field is required for SCHEDULED
	// billing to specify the reservation start time for the cluster. If not provided,
	// the cluster provisions immediately.
	ReservationStartTime param.Opt[time.Time] `json:"reservation_start_time,omitzero" format:"date-time"`
	// Custom Slurm image for Slurm clusters.
	SlurmImage param.Opt[string] `json:"slurm_image,omitzero"`
	// Shared memory size in GiB for Slurm cluster. This field is required if
	// cluster_type is SLURM.
	SlurmShmSizeGib param.Opt[int64] `json:"slurm_shm_size_gib,omitzero"`
	// ID of an existing volume to use with the cluster creation.
	VolumeID param.Opt[string] `json:"volume_id,omitzero"`
	// AcceptanceTestsParams groups all GPU acceptance test options when enabled is
	// true.
	AcceptanceTestsParams BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams `json:"acceptance_tests_params,omitzero"`
	// Add-ons to enable on the cluster at creation time.
	AddOns        []BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn       `json:"add_ons,omitzero"`
	ClusterConfig BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig `json:"cluster_config,omitzero"`
	// Type of cluster to create.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType string                                                                 `json:"cluster_type,omitzero"`
	OidcConfig  BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig `json:"oidc_config,omitzero"`
	// Inline configuration to create a shared volume with the cluster creation.
	SharedVolume BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume `json:"shared_volume,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion](
		"billing_type", "RESERVED", "ON_DEMAND", "SCHEDULED_CAPACITY",
	)
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion](
		"gpu_type", "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM", "H100_SXM_INF", "B300_SXM",
	)
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion](
		"cluster_type", "KUBERNETES", "SLURM",
	)
}

// AcceptanceTestsParams groups all GPU acceptance test options when enabled is
// true.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams struct {
	// Skip DCGM diagnostics acceptance test.
	DcgmDiagSkipped param.Opt[bool] `json:"dcgm_diag_skipped,omitzero"`
	// Whether to run GPU acceptance tests during cluster bring-up.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// GPU burn duration in seconds; 0 means use the default when enabled.
	GPUBurnDuration param.Opt[int64] `json:"gpu_burn_duration,omitzero"`
	// Skip GPU burn acceptance test.
	GPUBurnSkipped param.Opt[bool] `json:"gpu_burn_skipped,omitzero"`
	// Skip NCCL multi-node acceptance test.
	NcclMultiNodeSkipped param.Opt[bool] `json:"nccl_multi_node_skipped,omitzero"`
	// Skip NCCL single-node acceptance test.
	NcclSingleNodeSkipped param.Opt[bool] `json:"nccl_single_node_skipped,omitzero"`
	// Skip storage-performance acceptance test.
	StorageSkipped param.Opt[bool] `json:"storage_skipped,omitzero"`
	// DCGM diagnostic depth. SHORT = readiness; MEDIUM = default; LONG = system
	// validation; EXTENDED = memtest. An omitted value selects MEDIUM when enabled.
	//
	// Any of "DCGM_DIAG_LEVEL_SHORT", "DCGM_DIAG_LEVEL_MEDIUM",
	// "DCGM_DIAG_LEVEL_LONG", "DCGM_DIAG_LEVEL_EXTENDED".
	DcgmDiagLevel string `json:"dcgm_diag_level,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams](
		"dcgm_diag_level", "DCGM_DIAG_LEVEL_SHORT", "DCGM_DIAG_LEVEL_MEDIUM", "DCGM_DIAG_LEVEL_LONG", "DCGM_DIAG_LEVEL_EXTENDED",
	)
}

// The properties AddOnType, Name are required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn struct {
	// Type of add-on. Valid values: 'dashboard', 'ingress', 'torchpass'.
	AddOnType string `json:"add_on_type" api:"required"`
	// Human-readable name for this add-on instance.
	Name string `json:"name" api:"required"`
	// Configuration for a cluster add-on.
	Config BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig `json:"config,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a cluster add-on.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig struct {
	Dashboard BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard `json:"dashboard,omitzero"`
	// Configuration for the Headlamp Kubernetes dashboard add-on.
	Headlamp BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp `json:"headlamp,omitzero"`
	Ingress  BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress  `json:"ingress,omitzero"`
	// Configuration for the Slurm Web add-on.
	SlurmWeb BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb `json:"slurm_web,omitzero"`
	// Configuration for the Model Aware TorchPass add-on.
	Torchpass BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass `json:"torchpass,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Headlamp Kubernetes dashboard add-on.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp struct {
	// Whether to enable the Headlamp Kubernetes dashboard add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Slurm Web add-on.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb struct {
	// Whether to enable the Slurm Web add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Model Aware TorchPass add-on.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass struct {
	// Whether to enable the Model Aware TorchPass add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LoadBalancer is required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig struct {
	// Any of "NONE", "TRAEFIK", "NGINX", "ISTIO".
	LoadBalancer string `json:"load_balancer,omitzero" api:"required"`
	// NVIDIA GPU Operator chart/version for the tenant cluster (e.g. v24.6.2). When
	// omitted, a service default is applied.
	GPUOperatorVersion         param.Opt[string] `json:"gpu_operator_version,omitzero"`
	JumphostEnabled            param.Opt[bool]   `json:"jumphost_enabled,omitzero"`
	KubernetesDashboardEnabled param.Opt[bool]   `json:"kubernetes_dashboard_enabled,omitzero"`
	// NVIDIA Network Operator chart/version for the tenant cluster (e.g. v24.7.0).
	// When omitted, a service default is applied.
	NetworkOperatorVersion param.Opt[string] `json:"network_operator_version,omitzero"`
	// Whether this cluster uses a per-cluster SSH certificate authority for
	// OIDC-signed SSH access.
	SSHCaEnabled  param.Opt[bool]                                                                        `json:"ssh_ca_enabled,omitzero"`
	Ingress       BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress       `json:"ingress,omitzero"`
	Observability BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability `json:"observability,omitzero"`
	// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
	// init, extra conf).
	SlurmStartupScripts BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts `json:"slurm_startup_scripts,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig](
		"load_balancer", "NONE", "TRAEFIK", "NGINX", "ISTIO",
	)
}

type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
// init, extra conf).
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts struct {
	// Slurm controller epilog script.
	ControllerEpilog param.Opt[string] `json:"controller_epilog,omitzero"`
	// Slurm controller prolog script.
	ControllerProlog param.Opt[string] `json:"controller_prolog,omitzero"`
	// Additional slurm.conf fragments.
	ExtraSlurmConf param.Opt[string] `json:"extra_slurm_conf,omitzero"`
	// Script run on Slurm login node init.
	LoginInitScript param.Opt[string] `json:"login_init_script,omitzero"`
	// Script run on Slurm nodeset init.
	NodesetInitScript param.Opt[string] `json:"nodeset_init_script,omitzero"`
	// Slurm worker node epilog script.
	WorkerEpilog param.Opt[string] `json:"worker_epilog,omitzero"`
	// Slurm worker node prolog script.
	WorkerProlog param.Opt[string] `json:"worker_prolog,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ClientID, GroupClaim, GroupPrefix, IssuerURL, UsernameClaim,
// UsernamePrefix are required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig struct {
	// OIDC client ID for authentication.
	ClientID string `json:"client_id" api:"required"`
	// JWT claim to use for user groups. For example, 'groups'
	GroupClaim string `json:"group_claim" api:"required"`
	// Prefix to add to the group claim to form the final group name. For example,
	// 'oidc:'
	GroupPrefix string `json:"group_prefix" api:"required"`
	// OIDC issuer URL for authentication. For example, https://accounts.google.com
	IssuerURL string `json:"issuer_url" api:"required"`
	// JWT claim to use as the username. For example, 'sub' or 'email'
	UsernameClaim string `json:"username_claim" api:"required"`
	// Prefix to add to the username claim to form the final username. For example,
	// 'oidc:'
	UsernamePrefix string `json:"username_prefix" api:"required"`
	// CA certificate in PEM format to validate the OIDC issuer's TLS certificate. This
	// field is optional but recommended if the issuer uses a private CA or self-signed
	// certificate.
	CaCert param.Opt[string] `json:"ca_cert,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline configuration to create a shared volume with the cluster creation.
//
// The properties Region, SizeTib, VolumeName are required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume struct {
	// Region name. Usable regions can be found from `clusters.list_regions()`
	Region string `json:"region" api:"required"`
	// Volume size in whole tebibytes (TiB).
	SizeTib int64 `json:"size_tib" api:"required"`
	// User provided name of the volume.
	VolumeName string `json:"volume_name" api:"required"`
	// Cluster ID to pin the volume to the same substrate as that GPU cluster.
	InstanceClusterID param.Opt[string] `json:"instance_cluster_id,omitzero"`
	// When true, the shared volume is not deleted when the cluster is decommissioned.
	IsLifecycleIndependent param.Opt[bool] `json:"is_lifecycle_independent,omitzero"`
	// Project ID that will own the volume. When omitted, the caller's default project
	// is used.
	ProjectID param.Opt[string] `json:"project_id,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create a cluster with the legacy CUDA and NVIDIA driver selectors.
// nvidia_version_id may also be set when it resolves to the same catalog entry.
//
// The properties BillingType, ClusterName, CudaVersion, GPUType, NumGPUs,
// NvidiaDriverVersion, Region are required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia struct {
	// RESERVED billing types allow you to specify the duration of the cluster
	// reservation via the duration_days field. ON_DEMAND billing types will give you
	// ownership of the cluster until you delete it. SCHEDULED_CAPACITY billing types
	// allow you to reserve capacity for a scheduled time window. You must specify the
	// reservation_start_time and reservation_end_time with this request.
	//
	// Any of "RESERVED", "ON_DEMAND", "SCHEDULED_CAPACITY".
	BillingType string `json:"billing_type,omitzero" api:"required"`
	// Name of the GPU cluster.
	ClusterName string `json:"cluster_name" api:"required"`
	// Legacy CUDA selector for this cluster. Bare semantic values such as 12.5 select
	// ubuntu-22.04; existing OS-suffixed values remain accepted for compatibility.
	// Must be paired with nvidia_driver_version. Prefer nvidia_version_id for new
	// integrations.
	CudaVersion string `json:"cuda_version" api:"required"`
	// Type of GPU to use in the cluster
	//
	// Any of "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM",
	// "H100_SXM_INF", "B300_SXM".
	GPUType string `json:"gpu_type,omitzero" api:"required"`
	// Number of GPUs to allocate in the cluster. This must be multiple of 8. For
	// example, 8, 16 or 24
	NumGPUs int64 `json:"num_gpus" api:"required"`
	// Legacy NVIDIA driver selector for this cluster. For example, 550. Must be paired
	// with cuda_version. Prefer nvidia_version_id for new integrations.
	NvidiaDriverVersion string `json:"nvidia_driver_version" api:"required"`
	// Region to create the GPU cluster in. Usable regions can be found from
	// `client.clusters.list_regions()`
	Region string `json:"region" api:"required"`
	// Whether to enable auto-scaling for the cluster. If true, the cluster will
	// automatically scale the number of GPU worker nodes between num_gpus and
	// auto_scale_max_gpus based on the workload.
	AutoScale param.Opt[bool] `json:"auto_scale,omitzero"`
	// Maximum number of GPUs to which the cluster can be auto-scaled up. This field is
	// required if auto_scaled is true.
	AutoScaleMaxGPUs param.Opt[int64] `json:"auto_scale_max_gpus,omitzero"`
	// Whether GPU cluster should be auto-scaled based on the workload. By default, it
	// is not auto-scaled.
	//
	// Deprecated: deprecated
	AutoScaled param.Opt[bool] `json:"auto_scaled,omitzero"`
	// ID of the capacity pool to use for the cluster. This field is optional and only
	// applicable if the cluster is created from a capacity pool.
	CapacityPoolID param.Opt[string] `json:"capacity_pool_id,omitzero"`
	// Duration in days to keep the cluster running.
	DurationDays param.Opt[int64] `json:"duration_days,omitzero"`
	// Whether to install Traefik ingress controller in the cluster. This field is only
	// applicable for Kubernetes clusters and is false by default.
	InstallTraefik param.Opt[bool] `json:"install_traefik,omitzero"`
	// Number of GPUs to allocate from the capacity pool. Must be a multiple of 8 and
	// not exceed num_gpus.
	NumCapacityPoolGPUs param.Opt[int64] `json:"num_capacity_pool_gpus,omitzero"`
	// Number of preemptible GPUs to request alongside on-demand capacity. Must be a
	// multiple of 8. Preemptible nodes are cheaper but may be reclaimed when on-demand
	// capacity is needed elsewhere; the system fulfills this asynchronously and
	// surfaces the actual count in allocated_preemptible_gpus.
	NumPreemptibleGPUs param.Opt[int64] `json:"num_preemptible_gpus,omitzero"`
	// Number of prepaid (PLG) reserved GPUs for this cluster. When omitted for
	// RESERVED billing on create, the server defaults this to num_gpus.
	NumReservedGPUs param.Opt[int64] `json:"num_reserved_gpus,omitzero"`
	// Canonical region-specific NVIDIA version ID. If cuda_version and
	// nvidia_driver_version are also set, they must resolve to the same catalog entry.
	NvidiaVersionID param.Opt[string] `json:"nvidia_version_id,omitzero"`
	// Project ID for the cluster. If not set, the project from the request context is
	// used.
	ProjectID param.Opt[string] `json:"project_id,omitzero"`
	// Reservation end time of the cluster. This field is required for SCHEDULED
	// billing to specify the reservation end time for the cluster.
	ReservationEndTime param.Opt[time.Time] `json:"reservation_end_time,omitzero" format:"date-time"`
	// Reservation start time of the cluster. This field is required for SCHEDULED
	// billing to specify the reservation start time for the cluster. If not provided,
	// the cluster provisions immediately.
	ReservationStartTime param.Opt[time.Time] `json:"reservation_start_time,omitzero" format:"date-time"`
	// Custom Slurm image for Slurm clusters.
	SlurmImage param.Opt[string] `json:"slurm_image,omitzero"`
	// Shared memory size in GiB for Slurm cluster. This field is required if
	// cluster_type is SLURM.
	SlurmShmSizeGib param.Opt[int64] `json:"slurm_shm_size_gib,omitzero"`
	// ID of an existing volume to use with the cluster creation.
	VolumeID param.Opt[string] `json:"volume_id,omitzero"`
	// AcceptanceTestsParams groups all GPU acceptance test options when enabled is
	// true.
	AcceptanceTestsParams BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams `json:"acceptance_tests_params,omitzero"`
	// Add-ons to enable on the cluster at creation time.
	AddOns        []BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn       `json:"add_ons,omitzero"`
	ClusterConfig BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig `json:"cluster_config,omitzero"`
	// Type of cluster to create.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType string                                                                `json:"cluster_type,omitzero"`
	OidcConfig  BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig `json:"oidc_config,omitzero"`
	// Inline configuration to create a shared volume with the cluster creation.
	SharedVolume BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume `json:"shared_volume,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia](
		"billing_type", "RESERVED", "ON_DEMAND", "SCHEDULED_CAPACITY",
	)
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia](
		"gpu_type", "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM", "H100_SXM_INF", "B300_SXM",
	)
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia](
		"cluster_type", "KUBERNETES", "SLURM",
	)
}

// AcceptanceTestsParams groups all GPU acceptance test options when enabled is
// true.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams struct {
	// Skip DCGM diagnostics acceptance test.
	DcgmDiagSkipped param.Opt[bool] `json:"dcgm_diag_skipped,omitzero"`
	// Whether to run GPU acceptance tests during cluster bring-up.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// GPU burn duration in seconds; 0 means use the default when enabled.
	GPUBurnDuration param.Opt[int64] `json:"gpu_burn_duration,omitzero"`
	// Skip GPU burn acceptance test.
	GPUBurnSkipped param.Opt[bool] `json:"gpu_burn_skipped,omitzero"`
	// Skip NCCL multi-node acceptance test.
	NcclMultiNodeSkipped param.Opt[bool] `json:"nccl_multi_node_skipped,omitzero"`
	// Skip NCCL single-node acceptance test.
	NcclSingleNodeSkipped param.Opt[bool] `json:"nccl_single_node_skipped,omitzero"`
	// Skip storage-performance acceptance test.
	StorageSkipped param.Opt[bool] `json:"storage_skipped,omitzero"`
	// DCGM diagnostic depth. SHORT = readiness; MEDIUM = default; LONG = system
	// validation; EXTENDED = memtest. An omitted value selects MEDIUM when enabled.
	//
	// Any of "DCGM_DIAG_LEVEL_SHORT", "DCGM_DIAG_LEVEL_MEDIUM",
	// "DCGM_DIAG_LEVEL_LONG", "DCGM_DIAG_LEVEL_EXTENDED".
	DcgmDiagLevel string `json:"dcgm_diag_level,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams](
		"dcgm_diag_level", "DCGM_DIAG_LEVEL_SHORT", "DCGM_DIAG_LEVEL_MEDIUM", "DCGM_DIAG_LEVEL_LONG", "DCGM_DIAG_LEVEL_EXTENDED",
	)
}

// The properties AddOnType, Name are required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn struct {
	// Type of add-on. Valid values: 'dashboard', 'ingress', 'torchpass'.
	AddOnType string `json:"add_on_type" api:"required"`
	// Human-readable name for this add-on instance.
	Name string `json:"name" api:"required"`
	// Configuration for a cluster add-on.
	Config BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig `json:"config,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a cluster add-on.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig struct {
	Dashboard BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard `json:"dashboard,omitzero"`
	// Configuration for the Headlamp Kubernetes dashboard add-on.
	Headlamp BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp `json:"headlamp,omitzero"`
	Ingress  BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress  `json:"ingress,omitzero"`
	// Configuration for the Slurm Web add-on.
	SlurmWeb BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb `json:"slurm_web,omitzero"`
	// Configuration for the Model Aware TorchPass add-on.
	Torchpass BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass `json:"torchpass,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Headlamp Kubernetes dashboard add-on.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp struct {
	// Whether to enable the Headlamp Kubernetes dashboard add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Slurm Web add-on.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb struct {
	// Whether to enable the Slurm Web add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Model Aware TorchPass add-on.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass struct {
	// Whether to enable the Model Aware TorchPass add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LoadBalancer is required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig struct {
	// Any of "NONE", "TRAEFIK", "NGINX", "ISTIO".
	LoadBalancer string `json:"load_balancer,omitzero" api:"required"`
	// NVIDIA GPU Operator chart/version for the tenant cluster (e.g. v24.6.2). When
	// omitted, a service default is applied.
	GPUOperatorVersion         param.Opt[string] `json:"gpu_operator_version,omitzero"`
	JumphostEnabled            param.Opt[bool]   `json:"jumphost_enabled,omitzero"`
	KubernetesDashboardEnabled param.Opt[bool]   `json:"kubernetes_dashboard_enabled,omitzero"`
	// NVIDIA Network Operator chart/version for the tenant cluster (e.g. v24.7.0).
	// When omitted, a service default is applied.
	NetworkOperatorVersion param.Opt[string] `json:"network_operator_version,omitzero"`
	// Whether this cluster uses a per-cluster SSH certificate authority for
	// OIDC-signed SSH access.
	SSHCaEnabled  param.Opt[bool]                                                                       `json:"ssh_ca_enabled,omitzero"`
	Ingress       BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress       `json:"ingress,omitzero"`
	Observability BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability `json:"observability,omitzero"`
	// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
	// init, extra conf).
	SlurmStartupScripts BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts `json:"slurm_startup_scripts,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig](
		"load_balancer", "NONE", "TRAEFIK", "NGINX", "ISTIO",
	)
}

type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
// init, extra conf).
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts struct {
	// Slurm controller epilog script.
	ControllerEpilog param.Opt[string] `json:"controller_epilog,omitzero"`
	// Slurm controller prolog script.
	ControllerProlog param.Opt[string] `json:"controller_prolog,omitzero"`
	// Additional slurm.conf fragments.
	ExtraSlurmConf param.Opt[string] `json:"extra_slurm_conf,omitzero"`
	// Script run on Slurm login node init.
	LoginInitScript param.Opt[string] `json:"login_init_script,omitzero"`
	// Script run on Slurm nodeset init.
	NodesetInitScript param.Opt[string] `json:"nodeset_init_script,omitzero"`
	// Slurm worker node epilog script.
	WorkerEpilog param.Opt[string] `json:"worker_epilog,omitzero"`
	// Slurm worker node prolog script.
	WorkerProlog param.Opt[string] `json:"worker_prolog,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ClientID, GroupClaim, GroupPrefix, IssuerURL, UsernameClaim,
// UsernamePrefix are required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig struct {
	// OIDC client ID for authentication.
	ClientID string `json:"client_id" api:"required"`
	// JWT claim to use for user groups. For example, 'groups'
	GroupClaim string `json:"group_claim" api:"required"`
	// Prefix to add to the group claim to form the final group name. For example,
	// 'oidc:'
	GroupPrefix string `json:"group_prefix" api:"required"`
	// OIDC issuer URL for authentication. For example, https://accounts.google.com
	IssuerURL string `json:"issuer_url" api:"required"`
	// JWT claim to use as the username. For example, 'sub' or 'email'
	UsernameClaim string `json:"username_claim" api:"required"`
	// Prefix to add to the username claim to form the final username. For example,
	// 'oidc:'
	UsernamePrefix string `json:"username_prefix" api:"required"`
	// CA certificate in PEM format to validate the OIDC issuer's TLS certificate. This
	// field is optional but recommended if the issuer uses a private CA or self-signed
	// certificate.
	CaCert param.Opt[string] `json:"ca_cert,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline configuration to create a shared volume with the cluster creation.
//
// The properties Region, SizeTib, VolumeName are required.
type BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume struct {
	// Region name. Usable regions can be found from `clusters.list_regions()`
	Region string `json:"region" api:"required"`
	// Volume size in whole tebibytes (TiB).
	SizeTib int64 `json:"size_tib" api:"required"`
	// User provided name of the volume.
	VolumeName string `json:"volume_name" api:"required"`
	// Cluster ID to pin the volume to the same substrate as that GPU cluster.
	InstanceClusterID param.Opt[string] `json:"instance_cluster_id,omitzero"`
	// When true, the shared volume is not deleted when the cluster is decommissioned.
	IsLifecycleIndependent param.Opt[bool] `json:"is_lifecycle_independent,omitzero"`
	// Project ID that will own the volume. When omitted, the caller's default project
	// is used.
	ProjectID param.Opt[string] `json:"project_id,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterUpdateParams struct {
	// Number of GPUs to draw from the cluster's capacity pool. Only valid for clusters
	// created with a capacity_pool_id. Must be a multiple of 8 and not exceed
	// num_gpus. When omitted, the current value is preserved.
	NumCapacityPoolGPUs param.Opt[int64] `json:"num_capacity_pool_gpus,omitzero"`
	// Target GPU count for the cluster. When omitted, the server keeps the current GPU
	// count from cluster metadata (use for config-only or decommission-time-only
	// updates).
	NumGPUs param.Opt[int64] `json:"num_gpus,omitzero"`
	// Updated desired number of preemptible GPUs for the cluster. When omitted, the
	// current value is preserved. Must be a multiple of 8.
	NumPreemptibleGPUs param.Opt[int64] `json:"num_preemptible_gpus,omitzero"`
	// Number of reserved GPUs to update to. This field is only applicable for clusters
	// with RESERVED billing type.
	NumReservedGPUs param.Opt[int64] `json:"num_reserved_gpus,omitzero"`
	// Timestamp at which the cluster should be decommissioned. Only accepted for
	// prepaid clusters.
	ReservationEndTime param.Opt[time.Time] `json:"reservation_end_time,omitzero" format:"date-time"`
	// Add-ons to update on the cluster. Each entry identifies an existing add-on by
	// name and provides the new external config to merge.
	AddOns        []BetaClusterUpdateParamsAddOn       `json:"add_ons,omitzero"`
	ClusterConfig BetaClusterUpdateParamsClusterConfig `json:"cluster_config,omitzero"`
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

// The property Name is required.
type BetaClusterUpdateParamsAddOn struct {
	// Name of the add-on to update. Must match an existing add-on on the cluster.
	Name string `json:"name" api:"required"`
	// Configuration for a cluster add-on.
	Config BetaClusterUpdateParamsAddOnConfig `json:"config,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsAddOn) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsAddOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsAddOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a cluster add-on.
type BetaClusterUpdateParamsAddOnConfig struct {
	Dashboard BetaClusterUpdateParamsAddOnConfigDashboard `json:"dashboard,omitzero"`
	// Configuration for the Headlamp Kubernetes dashboard add-on.
	Headlamp BetaClusterUpdateParamsAddOnConfigHeadlamp `json:"headlamp,omitzero"`
	Ingress  BetaClusterUpdateParamsAddOnConfigIngress  `json:"ingress,omitzero"`
	// Configuration for the Slurm Web add-on.
	SlurmWeb BetaClusterUpdateParamsAddOnConfigSlurmWeb `json:"slurm_web,omitzero"`
	// Configuration for the Model Aware TorchPass add-on.
	Torchpass BetaClusterUpdateParamsAddOnConfigTorchpass `json:"torchpass,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsAddOnConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsAddOnConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsAddOnConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterUpdateParamsAddOnConfigDashboard struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsAddOnConfigDashboard) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsAddOnConfigDashboard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsAddOnConfigDashboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Headlamp Kubernetes dashboard add-on.
type BetaClusterUpdateParamsAddOnConfigHeadlamp struct {
	// Whether to enable the Headlamp Kubernetes dashboard add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsAddOnConfigHeadlamp) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsAddOnConfigHeadlamp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsAddOnConfigHeadlamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterUpdateParamsAddOnConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsAddOnConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsAddOnConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsAddOnConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Slurm Web add-on.
type BetaClusterUpdateParamsAddOnConfigSlurmWeb struct {
	// Whether to enable the Slurm Web add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsAddOnConfigSlurmWeb) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsAddOnConfigSlurmWeb
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsAddOnConfigSlurmWeb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Model Aware TorchPass add-on.
type BetaClusterUpdateParamsAddOnConfigTorchpass struct {
	// Whether to enable the Model Aware TorchPass add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsAddOnConfigTorchpass) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsAddOnConfigTorchpass
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsAddOnConfigTorchpass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LoadBalancer is required.
type BetaClusterUpdateParamsClusterConfig struct {
	// Any of "NONE", "TRAEFIK", "NGINX", "ISTIO".
	LoadBalancer string `json:"load_balancer,omitzero" api:"required"`
	// NVIDIA GPU Operator chart/version for the tenant cluster (e.g. v24.6.2). When
	// omitted, a service default is applied.
	GPUOperatorVersion         param.Opt[string] `json:"gpu_operator_version,omitzero"`
	JumphostEnabled            param.Opt[bool]   `json:"jumphost_enabled,omitzero"`
	KubernetesDashboardEnabled param.Opt[bool]   `json:"kubernetes_dashboard_enabled,omitzero"`
	// NVIDIA Network Operator chart/version for the tenant cluster (e.g. v24.7.0).
	// When omitted, a service default is applied.
	NetworkOperatorVersion param.Opt[string] `json:"network_operator_version,omitzero"`
	// Whether this cluster uses a per-cluster SSH certificate authority for
	// OIDC-signed SSH access.
	SSHCaEnabled  param.Opt[bool]                                   `json:"ssh_ca_enabled,omitzero"`
	Ingress       BetaClusterUpdateParamsClusterConfigIngress       `json:"ingress,omitzero"`
	Observability BetaClusterUpdateParamsClusterConfigObservability `json:"observability,omitzero"`
	// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
	// init, extra conf).
	SlurmStartupScripts BetaClusterUpdateParamsClusterConfigSlurmStartupScripts `json:"slurm_startup_scripts,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsClusterConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsClusterConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsClusterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaClusterUpdateParamsClusterConfig](
		"load_balancer", "NONE", "TRAEFIK", "NGINX", "ISTIO",
	)
}

type BetaClusterUpdateParamsClusterConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsClusterConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsClusterConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsClusterConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterUpdateParamsClusterConfigObservability struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsClusterConfigObservability) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsClusterConfigObservability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsClusterConfigObservability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
// init, extra conf).
type BetaClusterUpdateParamsClusterConfigSlurmStartupScripts struct {
	// Slurm controller epilog script.
	ControllerEpilog param.Opt[string] `json:"controller_epilog,omitzero"`
	// Slurm controller prolog script.
	ControllerProlog param.Opt[string] `json:"controller_prolog,omitzero"`
	// Additional slurm.conf fragments.
	ExtraSlurmConf param.Opt[string] `json:"extra_slurm_conf,omitzero"`
	// Script run on Slurm login node init.
	LoginInitScript param.Opt[string] `json:"login_init_script,omitzero"`
	// Script run on Slurm nodeset init.
	NodesetInitScript param.Opt[string] `json:"nodeset_init_script,omitzero"`
	// Slurm worker node epilog script.
	WorkerEpilog param.Opt[string] `json:"worker_epilog,omitzero"`
	// Slurm worker node prolog script.
	WorkerProlog param.Opt[string] `json:"worker_prolog,omitzero"`
	paramObj
}

func (r BetaClusterUpdateParamsClusterConfigSlurmStartupScripts) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterUpdateParamsClusterConfigSlurmStartupScripts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterUpdateParamsClusterConfigSlurmStartupScripts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cluster to update.
type BetaClusterUpdateParamsClusterType string

const (
	BetaClusterUpdateParamsClusterTypeKubernetes BetaClusterUpdateParamsClusterType = "KUBERNETES"
	BetaClusterUpdateParamsClusterTypeSlurm      BetaClusterUpdateParamsClusterType = "SLURM"
)

type BetaClusterListParams struct {
	// Optional UMS project ID to filter clusters by. When set, only clusters belonging
	// to this project are returned. The caller must be a member of the project;
	// otherwise the result set will be empty.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `query:"projectId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaClusterListParams]'s query parameters as `url.Values`.
func (r BetaClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
