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

type Cluster struct {
	// Enabled add-ons on this cluster. Only add-ons with enabled=true in their config
	// are returned.
	AddOns []ClusterAddOn `json:"add_ons" api:"required"`
	// Actual number of preemptible GPUs currently allocated to the cluster. Updated
	// asynchronously by the fulfillment and reclamation workers; may be less than
	// desired_preemptible_gpus when capacity is constrained.
	AllocatedPreemptibleGPUs int64 `json:"allocated_preemptible_gpus" api:"required"`
	// Billing type for the cluster (RESERVED, ON_DEMAND, or SCHEDULED_CAPACITY).
	//
	// Any of "RESERVED", "ON_DEMAND", "SCHEDULED_CAPACITY".
	BillingType ClusterBillingType `json:"billing_type" api:"required"`
	ClusterID   string             `json:"cluster_id" api:"required"`
	ClusterName string             `json:"cluster_name" api:"required"`
	// Type of cluster.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType       ClusterClusterType        `json:"cluster_type" api:"required"`
	ControlPlaneNodes []ClusterControlPlaneNode `json:"control_plane_nodes" api:"required"`
	CudaVersion       string                    `json:"cuda_version" api:"required"`
	// Customer's requested number of preemptible GPUs. Set on cluster create or
	// update; persists until changed.
	DesiredPreemptibleGPUs int64 `json:"desired_preemptible_gpus" api:"required"`
	// Any of "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM",
	// "H100_SXM_INF".
	GPUType        ClusterGPUType         `json:"gpu_type" api:"required"`
	GPUWorkerNodes []ClusterGPUWorkerNode `json:"gpu_worker_nodes" api:"required"`
	KubeConfig     string                 `json:"kube_config" api:"required"`
	// Number of CPU-only worker nodes in the cluster.
	NumCPUWorkers       int64  `json:"num_cpu_workers" api:"required"`
	NumGPUs             int64  `json:"num_gpus" api:"required"`
	NvidiaDriverVersion string `json:"nvidia_driver_version" api:"required"`
	// Cluster-level phase transition history.
	PhaseTransitions []ClusterPhaseTransition `json:"phase_transitions" api:"required"`
	ProjectID        string                   `json:"project_id" api:"required"`
	Region           string                   `json:"region" api:"required"`
	// Current status of the GPU cluster.
	//
	// Any of "WaitingForControlPlaneNodes", "WaitingForDataPlaneNodes",
	// "WaitingForSubnet", "WaitingForSharedVolume", "InstallingDrivers",
	// "RunningAcceptanceTests", "Paused", "OnDemandComputePaused", "Ready",
	// "Degraded", "Deleting".
	Status         ClusterStatus        `json:"status" api:"required"`
	Volumes        []ClusterVolume      `json:"volumes" api:"required"`
	CapacityPoolID string               `json:"capacity_pool_id"`
	ClusterConfig  ClusterClusterConfig `json:"cluster_config"`
	// Whether the control plane is currently ready.
	ControlPlaneReady bool      `json:"control_plane_ready"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	DurationHours     int64     `json:"duration_hours"`
	// Timestamp when the cluster first reached the Ready phase.
	FirstReadyAt   time.Time `json:"first_ready_at" format:"date-time"`
	InstallTraefik bool      `json:"install_traefik"`
	// Whether the cluster is managed inside a substrate environment.
	IsInSubstrate bool `json:"is_in_substrate"`
	// ID of the machine cluster backing this GPU cluster.
	MachineClusterID string `json:"machine_cluster_id"`
	// Internal NVIDIA version ID for this cluster's driver and CUDA combination.
	NvidiaDriverVersionID string            `json:"nvidia_driver_version_id"`
	OidcConfig            ClusterOidcConfig `json:"oidc_config"`
	// Data-volume image name for GPU worker nodes.
	OsImage              string    `json:"os_image"`
	ReservationEndTime   time.Time `json:"reservation_end_time" format:"date-time"`
	ReservationStartTime time.Time `json:"reservation_start_time" format:"date-time"`
	SlurmShmSizeGib      int64     `json:"slurm_shm_size_gib"`
	// UMS organization ID associated with this cluster.
	UmsOrgID string `json:"ums_org_id"`
	// UMS project ID associated with this cluster.
	UmsProjectID string `json:"ums_project_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddOns                   respjson.Field
		AllocatedPreemptibleGPUs respjson.Field
		BillingType              respjson.Field
		ClusterID                respjson.Field
		ClusterName              respjson.Field
		ClusterType              respjson.Field
		ControlPlaneNodes        respjson.Field
		CudaVersion              respjson.Field
		DesiredPreemptibleGPUs   respjson.Field
		GPUType                  respjson.Field
		GPUWorkerNodes           respjson.Field
		KubeConfig               respjson.Field
		NumCPUWorkers            respjson.Field
		NumGPUs                  respjson.Field
		NvidiaDriverVersion      respjson.Field
		PhaseTransitions         respjson.Field
		ProjectID                respjson.Field
		Region                   respjson.Field
		Status                   respjson.Field
		Volumes                  respjson.Field
		CapacityPoolID           respjson.Field
		ClusterConfig            respjson.Field
		ControlPlaneReady        respjson.Field
		CreatedAt                respjson.Field
		DurationHours            respjson.Field
		FirstReadyAt             respjson.Field
		InstallTraefik           respjson.Field
		IsInSubstrate            respjson.Field
		MachineClusterID         respjson.Field
		NvidiaDriverVersionID    respjson.Field
		OidcConfig               respjson.Field
		OsImage                  respjson.Field
		ReservationEndTime       respjson.Field
		ReservationStartTime     respjson.Field
		SlurmShmSizeGib          respjson.Field
		UmsOrgID                 respjson.Field
		UmsProjectID             respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Cluster) RawJSON() string { return r.JSON.raw }
func (r *Cluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AddOnInfo is returned in cluster responses and add-on CRUD operations.
type ClusterAddOn struct {
	AddOnType string             `json:"add_on_type" api:"required"`
	Config    ClusterAddOnConfig `json:"config" api:"required"`
	Name      string             `json:"name" api:"required"`
	State     ClusterAddOnState  `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddOnType   respjson.Field
		Config      respjson.Field
		Name        respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOn) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterAddOnConfig struct {
	Dashboard ClusterAddOnConfigDashboard `json:"dashboard"`
	Ingress   ClusterAddOnConfigIngress   `json:"ingress"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dashboard   respjson.Field
		Ingress     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnConfig) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterAddOnConfigDashboard struct {
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnConfigDashboard) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnConfigDashboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterAddOnConfigIngress struct {
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnConfigIngress) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterAddOnState struct {
	Dashboard ClusterAddOnStateDashboard `json:"dashboard"`
	Ingress   ClusterAddOnStateIngress   `json:"ingress"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dashboard   respjson.Field
		Ingress     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnState) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterAddOnStateDashboard struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnStateDashboard) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnStateDashboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterAddOnStateIngress struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnStateIngress) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnStateIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing type for the cluster (RESERVED, ON_DEMAND, or SCHEDULED_CAPACITY).
type ClusterBillingType string

const (
	ClusterBillingTypeReserved          ClusterBillingType = "RESERVED"
	ClusterBillingTypeOnDemand          ClusterBillingType = "ON_DEMAND"
	ClusterBillingTypeScheduledCapacity ClusterBillingType = "SCHEDULED_CAPACITY"
)

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
	NumCPUCores int64   `json:"num_cpu_cores" api:"required"`
	// Phase transition history for this control plane node.
	PhaseTransitions []ClusterControlPlaneNodePhaseTransition `json:"phase_transitions" api:"required"`
	Status           string                                   `json:"status" api:"required"`
	// Public IPv4 address of the control plane node.
	PublicIpv4 string `json:"public_ipv4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HostName         respjson.Field
		MemoryGib        respjson.Field
		Network          respjson.Field
		NodeID           respjson.Field
		NumCPUCores      respjson.Field
		PhaseTransitions respjson.Field
		Status           respjson.Field
		PublicIpv4       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterControlPlaneNode) RawJSON() string { return r.JSON.raw }
func (r *ClusterControlPlaneNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterControlPlaneNodePhaseTransition struct {
	// Node phase.
	//
	// Any of "NODE_PHASE_PENDING", "NODE_PHASE_SCHEDULING", "NODE_PHASE_BOOTING",
	// "NODE_PHASE_BOOTSTRAPPING", "NODE_PHASE_RUNNING", "NODE_PHASE_SUCCEEDED",
	// "NODE_PHASE_FAILED", "NODE_PHASE_PAUSED".
	Phase string `json:"phase" api:"required"`
	// Timestamp when the phase transition occurred.
	TransitionTime time.Time `json:"transition_time" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Phase          respjson.Field
		TransitionTime respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterControlPlaneNodePhaseTransition) RawJSON() string { return r.JSON.raw }
func (r *ClusterControlPlaneNodePhaseTransition) UnmarshalJSON(data []byte) error {
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
	NumCPUCores int64    `json:"num_cpu_cores" api:"required"`
	NumGPUs     int64    `json:"num_gpus" api:"required"`
	// Phase transition history for this GPU worker node.
	PhaseTransitions []ClusterGPUWorkerNodePhaseTransition `json:"phase_transitions" api:"required"`
	Status           string                                `json:"status" api:"required"`
	// Whether auto-remediation is enabled for this node's instance.
	AutoRemediationEnabled bool `json:"auto_remediation_enabled"`
	// Ephemeral storage size, such as 1Ti.
	EphemeralStorage string `json:"ephemeral_storage"`
	// Number of InfiniBand HCAs.
	IbHcaCount int64 `json:"ib_hca_count"`
	// InfiniBand HCA type.
	IbHcaType  string `json:"ib_hca_type"`
	InstanceID string `json:"instance_id"`
	// Remediation represents a node remediation request for an instance. An instance
	// can have multiple remediations over time (e.g., failed attempts followed by
	// retries).
	LatestRemediation Remediation `json:"latest_remediation"`
	// Whether this node is marked for deletion by the operator.
	MarkedForDeletion bool `json:"marked_for_deletion"`
	// Number of NVSwitches.
	NvswitchCount int64 `json:"nvswitch_count"`
	// NVSwitch type.
	NvswitchType string `json:"nvswitch_type"`
	// Public IPv4 address of the GPU worker node.
	PublicIpv4          string `json:"public_ipv4"`
	SlurmWorkerHostname string `json:"slurm_worker_hostname"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HostName               respjson.Field
		MemoryGib              respjson.Field
		Networks               respjson.Field
		NodeID                 respjson.Field
		NumCPUCores            respjson.Field
		NumGPUs                respjson.Field
		PhaseTransitions       respjson.Field
		Status                 respjson.Field
		AutoRemediationEnabled respjson.Field
		EphemeralStorage       respjson.Field
		IbHcaCount             respjson.Field
		IbHcaType              respjson.Field
		InstanceID             respjson.Field
		LatestRemediation      respjson.Field
		MarkedForDeletion      respjson.Field
		NvswitchCount          respjson.Field
		NvswitchType           respjson.Field
		PublicIpv4             respjson.Field
		SlurmWorkerHostname    respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterGPUWorkerNode) RawJSON() string { return r.JSON.raw }
func (r *ClusterGPUWorkerNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterGPUWorkerNodePhaseTransition struct {
	// Node phase.
	//
	// Any of "NODE_PHASE_PENDING", "NODE_PHASE_SCHEDULING", "NODE_PHASE_BOOTING",
	// "NODE_PHASE_BOOTSTRAPPING", "NODE_PHASE_RUNNING", "NODE_PHASE_SUCCEEDED",
	// "NODE_PHASE_FAILED", "NODE_PHASE_PAUSED".
	Phase string `json:"phase" api:"required"`
	// Timestamp when the phase transition occurred.
	TransitionTime time.Time `json:"transition_time" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Phase          respjson.Field
		TransitionTime respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterGPUWorkerNodePhaseTransition) RawJSON() string { return r.JSON.raw }
func (r *ClusterGPUWorkerNodePhaseTransition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterPhaseTransition struct {
	// Cluster phase.
	//
	// Any of "CLUSTER_PHASE_QUEUED", "CLUSTER_PHASE_SCHEDULED",
	// "CLUSTER_PHASE_WAITING_FOR_CONTROL_PLANE_NODES",
	// "CLUSTER_PHASE_WAITING_FOR_DATA_PLANE_NODES",
	// "CLUSTER_PHASE_WAITING_FOR_SUBNET", "CLUSTER_PHASE_WAITING_FOR_SHARED_VOLUME",
	// "CLUSTER_PHASE_WAITING_FOR_AUTO_SCALER", "CLUSTER_PHASE_INSTALLING_DRIVERS",
	// "CLUSTER_PHASE_RUNNING_ACCEPTANCE_TESTS",
	// "CLUSTER_PHASE_ACCEPTANCE_TESTS_FAILED", "CLUSTER_PHASE_RUNNING_NCCL_TESTS",
	// "CLUSTER_PHASE_NCCL_TESTS_FAILED", "CLUSTER_PHASE_READY",
	// "CLUSTER_PHASE_PAUSED", "CLUSTER_PHASE_ON_DEMAND_COMPUTE_PAUSED",
	// "CLUSTER_PHASE_DEGRADED", "CLUSTER_PHASE_DELETING".
	Phase string `json:"phase" api:"required"`
	// Timestamp when the phase transition occurred.
	TransitionTime time.Time `json:"transition_time" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Phase          respjson.Field
		TransitionTime respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterPhaseTransition) RawJSON() string { return r.JSON.raw }
func (r *ClusterPhaseTransition) UnmarshalJSON(data []byte) error {
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
	// Size of the volume in TiB.
	SizeTib int64 `json:"size_tib" api:"required"`
	// Current status of the volume.
	Status string `json:"status" api:"required"`
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
func (r ClusterVolume) RawJSON() string { return r.JSON.raw }
func (r *ClusterVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterClusterConfig struct {
	// Any of "NONE", "TRAEFIK", "NGINX", "ISTIO".
	LoadBalancer string `json:"load_balancer" api:"required"`
	// NVIDIA GPU Operator chart/version for the tenant cluster (e.g. v24.6.2). When
	// omitted, a service default is applied.
	GPUOperatorVersion         string                            `json:"gpu_operator_version"`
	Ingress                    ClusterClusterConfigIngress       `json:"ingress"`
	JumphostEnabled            bool                              `json:"jumphost_enabled"`
	KubernetesDashboardEnabled bool                              `json:"kubernetes_dashboard_enabled"`
	Observability              ClusterClusterConfigObservability `json:"observability"`
	// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
	// init, extra conf).
	SlurmStartupScripts ClusterClusterConfigSlurmStartupScripts `json:"slurm_startup_scripts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LoadBalancer               respjson.Field
		GPUOperatorVersion         respjson.Field
		Ingress                    respjson.Field
		JumphostEnabled            respjson.Field
		KubernetesDashboardEnabled respjson.Field
		Observability              respjson.Field
		SlurmStartupScripts        respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterClusterConfig) RawJSON() string { return r.JSON.raw }
func (r *ClusterClusterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterClusterConfigIngress struct {
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterClusterConfigIngress) RawJSON() string { return r.JSON.raw }
func (r *ClusterClusterConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterClusterConfigObservability struct {
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterClusterConfigObservability) RawJSON() string { return r.JSON.raw }
func (r *ClusterClusterConfigObservability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
// init, extra conf).
type ClusterClusterConfigSlurmStartupScripts struct {
	// Slurm controller epilog script.
	ControllerEpilog string `json:"controller_epilog"`
	// Slurm controller prolog script.
	ControllerProlog string `json:"controller_prolog"`
	// Additional slurm.conf fragments.
	ExtraSlurmConf string `json:"extra_slurm_conf"`
	// Script run on Slurm login node init.
	LoginInitScript string `json:"login_init_script"`
	// Script run on Slurm nodeset init.
	NodesetInitScript string `json:"nodeset_init_script"`
	// Slurm worker node epilog script.
	WorkerEpilog string `json:"worker_epilog"`
	// Slurm worker node prolog script.
	WorkerProlog string `json:"worker_prolog"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ControllerEpilog  respjson.Field
		ControllerProlog  respjson.Field
		ExtraSlurmConf    respjson.Field
		LoginInitScript   respjson.Field
		NodesetInitScript respjson.Field
		WorkerEpilog      respjson.Field
		WorkerProlog      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterClusterConfigSlurmStartupScripts) RawJSON() string { return r.JSON.raw }
func (r *ClusterClusterConfigSlurmStartupScripts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterOidcConfig struct {
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
	CaCert string `json:"ca_cert"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID       respjson.Field
		GroupClaim     respjson.Field
		GroupPrefix    respjson.Field
		IssuerURL      respjson.Field
		UsernameClaim  respjson.Field
		UsernamePrefix respjson.Field
		CaCert         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterOidcConfig) RawJSON() string { return r.JSON.raw }
func (r *ClusterOidcConfig) UnmarshalJSON(data []byte) error {
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
	// Whether to enable auto-scaling for the cluster. If true, the cluster will
	// automatically scale the number of GPU worker nodes between num_gpus and
	// auto_scale_max_gpus based on the workload.
	AutoScale param.Opt[bool] `json:"auto_scale,omitzero"`
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
	AcceptanceTestsParams BetaClusterNewParamsAcceptanceTestsParams `json:"acceptance_tests_params,omitzero"`
	// Add-ons to enable on the cluster at creation time.
	AddOns        []BetaClusterNewParamsAddOn       `json:"add_ons,omitzero"`
	ClusterConfig BetaClusterNewParamsClusterConfig `json:"cluster_config,omitzero"`
	// Type of cluster to create.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType BetaClusterNewParamsClusterType `json:"cluster_type,omitzero"`
	OidcConfig  BetaClusterNewParamsOidcConfig  `json:"oidc_config,omitzero"`
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

// AcceptanceTestsParams groups all GPU acceptance test options when enabled is
// true.
type BetaClusterNewParamsAcceptanceTestsParams struct {
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
	// DCGM diagnostic depth. SHORT = readiness; MEDIUM = default; LONG = system
	// validation; EXTENDED = memtest. An omitted value selects MEDIUM when enabled.
	//
	// Any of "DCGM_DIAG_LEVEL_SHORT", "DCGM_DIAG_LEVEL_MEDIUM",
	// "DCGM_DIAG_LEVEL_LONG", "DCGM_DIAG_LEVEL_EXTENDED".
	DcgmDiagLevel string `json:"dcgm_diag_level,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsAcceptanceTestsParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsAcceptanceTestsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsAcceptanceTestsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaClusterNewParamsAcceptanceTestsParams](
		"dcgm_diag_level", "DCGM_DIAG_LEVEL_SHORT", "DCGM_DIAG_LEVEL_MEDIUM", "DCGM_DIAG_LEVEL_LONG", "DCGM_DIAG_LEVEL_EXTENDED",
	)
}

// The properties AddOnType, Name are required.
type BetaClusterNewParamsAddOn struct {
	// Type of add-on. Valid values: 'dashboard', 'ingress'.
	AddOnType string `json:"add_on_type" api:"required"`
	// Human-readable name for this add-on instance.
	Name   string                          `json:"name" api:"required"`
	Config BetaClusterNewParamsAddOnConfig `json:"config,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsAddOn) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsAddOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsAddOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsAddOnConfig struct {
	Dashboard BetaClusterNewParamsAddOnConfigDashboard `json:"dashboard,omitzero"`
	Ingress   BetaClusterNewParamsAddOnConfigIngress   `json:"ingress,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsAddOnConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsAddOnConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsAddOnConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsAddOnConfigDashboard struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsAddOnConfigDashboard) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsAddOnConfigDashboard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsAddOnConfigDashboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsAddOnConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsAddOnConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsAddOnConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsAddOnConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LoadBalancer is required.
type BetaClusterNewParamsClusterConfig struct {
	// Any of "NONE", "TRAEFIK", "NGINX", "ISTIO".
	LoadBalancer string `json:"load_balancer,omitzero" api:"required"`
	// NVIDIA GPU Operator chart/version for the tenant cluster (e.g. v24.6.2). When
	// omitted, a service default is applied.
	GPUOperatorVersion         param.Opt[string]                              `json:"gpu_operator_version,omitzero"`
	JumphostEnabled            param.Opt[bool]                                `json:"jumphost_enabled,omitzero"`
	KubernetesDashboardEnabled param.Opt[bool]                                `json:"kubernetes_dashboard_enabled,omitzero"`
	Ingress                    BetaClusterNewParamsClusterConfigIngress       `json:"ingress,omitzero"`
	Observability              BetaClusterNewParamsClusterConfigObservability `json:"observability,omitzero"`
	// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
	// init, extra conf).
	SlurmStartupScripts BetaClusterNewParamsClusterConfigSlurmStartupScripts `json:"slurm_startup_scripts,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsClusterConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsClusterConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsClusterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaClusterNewParamsClusterConfig](
		"load_balancer", "NONE", "TRAEFIK", "NGINX", "ISTIO",
	)
}

type BetaClusterNewParamsClusterConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsClusterConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsClusterConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsClusterConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterNewParamsClusterConfigObservability struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r BetaClusterNewParamsClusterConfigObservability) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsClusterConfigObservability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsClusterConfigObservability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
// init, extra conf).
type BetaClusterNewParamsClusterConfigSlurmStartupScripts struct {
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

func (r BetaClusterNewParamsClusterConfigSlurmStartupScripts) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsClusterConfigSlurmStartupScripts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsClusterConfigSlurmStartupScripts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cluster to create.
type BetaClusterNewParamsClusterType string

const (
	BetaClusterNewParamsClusterTypeKubernetes BetaClusterNewParamsClusterType = "KUBERNETES"
	BetaClusterNewParamsClusterTypeSlurm      BetaClusterNewParamsClusterType = "SLURM"
)

// The properties ClientID, GroupClaim, GroupPrefix, IssuerURL, UsernameClaim,
// UsernamePrefix are required.
type BetaClusterNewParamsOidcConfig struct {
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

func (r BetaClusterNewParamsOidcConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterNewParamsOidcConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterNewParamsOidcConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline configuration to create a shared volume with the cluster creation.
//
// The properties Region, SizeTib, VolumeName are required.
type BetaClusterNewParamsSharedVolume struct {
	// Region name. Usable regions can be found from `clusters.list_regions()`
	Region string `json:"region" api:"required"`
	// Volume size in whole tebibytes (TiB).
	SizeTib int64 `json:"size_tib" api:"required"`
	// User provided name of the volume.
	VolumeName string `json:"volume_name" api:"required"`
	// When true, the shared volume is not deleted when the cluster is decommissioned.
	IsLifecycleIndependent param.Opt[bool] `json:"is_lifecycle_independent,omitzero"`
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
	Name   string                             `json:"name" api:"required"`
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

type BetaClusterUpdateParamsAddOnConfig struct {
	Dashboard BetaClusterUpdateParamsAddOnConfigDashboard `json:"dashboard,omitzero"`
	Ingress   BetaClusterUpdateParamsAddOnConfigIngress   `json:"ingress,omitzero"`
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

// The property LoadBalancer is required.
type BetaClusterUpdateParamsClusterConfig struct {
	// Any of "NONE", "TRAEFIK", "NGINX", "ISTIO".
	LoadBalancer string `json:"load_balancer,omitzero" api:"required"`
	// NVIDIA GPU Operator chart/version for the tenant cluster (e.g. v24.6.2). When
	// omitted, a service default is applied.
	GPUOperatorVersion         param.Opt[string]                                 `json:"gpu_operator_version,omitzero"`
	JumphostEnabled            param.Opt[bool]                                   `json:"jumphost_enabled,omitzero"`
	KubernetesDashboardEnabled param.Opt[bool]                                   `json:"kubernetes_dashboard_enabled,omitzero"`
	Ingress                    BetaClusterUpdateParamsClusterConfigIngress       `json:"ingress,omitzero"`
	Observability              BetaClusterUpdateParamsClusterConfigObservability `json:"observability,omitzero"`
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
	ProjectID param.Opt[string] `query:"project_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaClusterListParams]'s query parameters as `url.Values`.
func (r BetaClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
