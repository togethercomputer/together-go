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

// ClusterService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterService] method instead.
type ClusterService struct {
	Options      []option.RequestOption
	Remediations ClusterRemediationService
	Storage      ClusterStorageService
}

// NewClusterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewClusterService(opts ...option.RequestOption) (r ClusterService) {
	r = ClusterService{}
	r.Options = opts
	r.Remediations = NewClusterRemediationService(opts...)
	r.Storage = NewClusterStorageService(opts...)
	return
}

// Create an Instant Cluster on Together's high-performance GPU clusters. With
// features like on-demand scaling, long-lived resizable high-bandwidth shared
// DC-local storage, Kubernetes and Slurm cluster flavors, a REST API, and
// Terraform support, you can run workloads flexibly without complex infrastructure
// management.
func (r *ClusterService) New(ctx context.Context, body ClusterNewParams, opts ...option.RequestOption) (res *Cluster, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve information about a specific GPU cluster.
func (r *ClusterService) Get(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *Cluster, err error) {
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
func (r *ClusterService) Update(ctx context.Context, clusterID string, body ClusterUpdateParams, opts ...option.RequestOption) (res *Cluster, err error) {
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
func (r *ClusterService) List(ctx context.Context, query ClusterListParams, opts ...option.RequestOption) (res *ClusterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a GPU cluster by cluster ID.
func (r *ClusterService) Delete(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *ClusterDeleteResponse, err error) {
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
func (r *ClusterService) ListRegions(ctx context.Context, opts ...option.RequestOption) (res *ClusterListRegionsResponse, err error) {
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
	// "H100_SXM_INF", "B300_SXM".
	GPUType        ClusterGPUType         `json:"gpu_type" api:"required"`
	GPUWorkerNodes []ClusterGPUWorkerNode `json:"gpu_worker_nodes" api:"required"`
	KubeConfig     string                 `json:"kube_config" api:"required"`
	// Number of GPUs to draw from a capacity pool. A component of the overall
	// num_gpus, alongside num_reserved_gpus.
	NumCapacityPoolGPUs int64 `json:"num_capacity_pool_gpus" api:"required"`
	// Number of CPU-only worker nodes in the cluster.
	NumCPUWorkers int64 `json:"num_cpu_workers" api:"required"`
	NumGPUs       int64 `json:"num_gpus" api:"required"`
	// Number of prepaid reserved GPUs for this cluster. A component of the overall
	// num_gpus, alongside num_capacity_pool_gpus.
	NumReservedGPUs     int64  `json:"num_reserved_gpus" api:"required"`
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
	// GPU worker nodes retained after they left the live data plane. These are
	// separate from gpu_worker_nodes and must not be counted as live capacity.
	DeletedGPUWorkerNodes []ClusterDeletedGPUWorkerNode `json:"deleted_gpu_worker_nodes"`
	DurationHours         int64                         `json:"duration_hours"`
	// Timestamp when the cluster first reached the Ready phase.
	FirstReadyAt   time.Time `json:"first_ready_at" format:"date-time"`
	InstallTraefik bool      `json:"install_traefik"`
	// Whether the cluster is managed inside a substrate environment.
	IsInSubstrate bool `json:"is_in_substrate"`
	// ID of the machine cluster backing this GPU cluster.
	MachineClusterID string `json:"machine_cluster_id"`
	// Recent node lifecycle events such as scale-up, scale-down, and preemption.
	// Combine these with live and deleted node lists to render the cluster timeline.
	NodeLifecycleEvents []ClusterNodeLifecycleEvent `json:"node_lifecycle_events"`
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
		NumCapacityPoolGPUs      respjson.Field
		NumCPUWorkers            respjson.Field
		NumGPUs                  respjson.Field
		NumReservedGPUs          respjson.Field
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
		DeletedGPUWorkerNodes    respjson.Field
		DurationHours            respjson.Field
		FirstReadyAt             respjson.Field
		InstallTraefik           respjson.Field
		IsInSubstrate            respjson.Field
		MachineClusterID         respjson.Field
		NodeLifecycleEvents      respjson.Field
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
	AddOnType string `json:"add_on_type" api:"required"`
	// Configuration for a cluster add-on.
	Config ClusterAddOnConfig `json:"config" api:"required"`
	Name   string             `json:"name" api:"required"`
	// State for a cluster add-on.
	State ClusterAddOnState `json:"state" api:"required"`
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

// Configuration for a cluster add-on.
type ClusterAddOnConfig struct {
	Dashboard ClusterAddOnConfigDashboard `json:"dashboard"`
	// Configuration for the Headlamp Kubernetes dashboard add-on.
	Headlamp ClusterAddOnConfigHeadlamp `json:"headlamp"`
	Ingress  ClusterAddOnConfigIngress  `json:"ingress"`
	// Configuration for the Slurm Web add-on.
	SlurmWeb ClusterAddOnConfigSlurmWeb `json:"slurm_web"`
	// Configuration for the Model Aware TorchPass add-on.
	Torchpass ClusterAddOnConfigTorchpass `json:"torchpass"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dashboard   respjson.Field
		Headlamp    respjson.Field
		Ingress     respjson.Field
		SlurmWeb    respjson.Field
		Torchpass   respjson.Field
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

// Configuration for the Headlamp Kubernetes dashboard add-on.
type ClusterAddOnConfigHeadlamp struct {
	// Whether to enable the Headlamp Kubernetes dashboard add-on.
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnConfigHeadlamp) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnConfigHeadlamp) UnmarshalJSON(data []byte) error {
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

// Configuration for the Slurm Web add-on.
type ClusterAddOnConfigSlurmWeb struct {
	// Whether to enable the Slurm Web add-on.
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnConfigSlurmWeb) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnConfigSlurmWeb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Model Aware TorchPass add-on.
type ClusterAddOnConfigTorchpass struct {
	// Whether to enable the Model Aware TorchPass add-on.
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnConfigTorchpass) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnConfigTorchpass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// State for a cluster add-on.
type ClusterAddOnState struct {
	Dashboard ClusterAddOnStateDashboard `json:"dashboard"`
	// State for the Headlamp Kubernetes dashboard add-on.
	Headlamp ClusterAddOnStateHeadlamp `json:"headlamp"`
	Ingress  ClusterAddOnStateIngress  `json:"ingress"`
	// State for the Slurm Web add-on.
	SlurmWeb ClusterAddOnStateSlurmWeb `json:"slurm_web"`
	// State for the Model Aware TorchPass add-on.
	Torchpass ClusterAddOnStateTorchpass `json:"torchpass"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dashboard   respjson.Field
		Headlamp    respjson.Field
		Ingress     respjson.Field
		SlurmWeb    respjson.Field
		Torchpass   respjson.Field
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

// State for the Headlamp Kubernetes dashboard add-on.
type ClusterAddOnStateHeadlamp struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnStateHeadlamp) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnStateHeadlamp) UnmarshalJSON(data []byte) error {
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

// State for the Slurm Web add-on.
type ClusterAddOnStateSlurmWeb struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnStateSlurmWeb) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnStateSlurmWeb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// State for the Model Aware TorchPass add-on.
type ClusterAddOnStateTorchpass struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAddOnStateTorchpass) RawJSON() string { return r.JSON.raw }
func (r *ClusterAddOnStateTorchpass) UnmarshalJSON(data []byte) error {
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
	ClusterGPUTypeB300Sxm    ClusterGPUType = "B300_SXM"
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
	// Timestamp when the node left the live data plane. Only set for
	// deleted_gpu_worker_nodes.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
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
		DeletedAt              respjson.Field
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
	GPUOperatorVersion         string                      `json:"gpu_operator_version"`
	Ingress                    ClusterClusterConfigIngress `json:"ingress"`
	JumphostEnabled            bool                        `json:"jumphost_enabled"`
	KubernetesDashboardEnabled bool                        `json:"kubernetes_dashboard_enabled"`
	// NVIDIA Network Operator chart/version for the tenant cluster (e.g. v24.7.0).
	// When omitted, a service default is applied.
	NetworkOperatorVersion string                            `json:"network_operator_version"`
	Observability          ClusterClusterConfigObservability `json:"observability"`
	// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
	// init, extra conf).
	SlurmStartupScripts ClusterClusterConfigSlurmStartupScripts `json:"slurm_startup_scripts"`
	// Whether this cluster uses a per-cluster SSH certificate authority for
	// OIDC-signed SSH access.
	SSHCaEnabled bool `json:"ssh_ca_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LoadBalancer               respjson.Field
		GPUOperatorVersion         respjson.Field
		Ingress                    respjson.Field
		JumphostEnabled            respjson.Field
		KubernetesDashboardEnabled respjson.Field
		NetworkOperatorVersion     respjson.Field
		Observability              respjson.Field
		SlurmStartupScripts        respjson.Field
		SSHCaEnabled               respjson.Field
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

type ClusterDeletedGPUWorkerNode struct {
	HostName    string   `json:"host_name" api:"required"`
	MemoryGib   float64  `json:"memory_gib" api:"required"`
	Networks    []string `json:"networks" api:"required"`
	NodeID      string   `json:"node_id" api:"required"`
	NumCPUCores int64    `json:"num_cpu_cores" api:"required"`
	NumGPUs     int64    `json:"num_gpus" api:"required"`
	// Phase transition history for this GPU worker node.
	PhaseTransitions []ClusterDeletedGPUWorkerNodePhaseTransition `json:"phase_transitions" api:"required"`
	Status           string                                       `json:"status" api:"required"`
	// Whether auto-remediation is enabled for this node's instance.
	AutoRemediationEnabled bool `json:"auto_remediation_enabled"`
	// Timestamp when the node left the live data plane. Only set for
	// deleted_gpu_worker_nodes.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
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
		DeletedAt              respjson.Field
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
func (r ClusterDeletedGPUWorkerNode) RawJSON() string { return r.JSON.raw }
func (r *ClusterDeletedGPUWorkerNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterDeletedGPUWorkerNodePhaseTransition struct {
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
func (r ClusterDeletedGPUWorkerNodePhaseTransition) RawJSON() string { return r.JSON.raw }
func (r *ClusterDeletedGPUWorkerNodePhaseTransition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Node lifecycle event included in a GPU cluster timeline.
type ClusterNodeLifecycleEvent struct {
	// Human-readable lifecycle event message.
	Message string `json:"message" api:"required"`
	// Tenant node name this lifecycle event applies to.
	NodeID string `json:"node_id" api:"required"`
	// Lifecycle event reason, for example TogetherScaledUp, TogetherScaledDown, or
	// TogetherPreempted.
	Reason string `json:"reason" api:"required"`
	// Event timestamp.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		NodeID      respjson.Field
		Reason      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterNodeLifecycleEvent) RawJSON() string { return r.JSON.raw }
func (r *ClusterNodeLifecycleEvent) UnmarshalJSON(data []byte) error {
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

type ClusterListResponse struct {
	Clusters []Cluster `json:"clusters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clusters    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterListResponse) RawJSON() string { return r.JSON.raw }
func (r *ClusterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterDeleteResponse struct {
	ClusterID string `json:"cluster_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ClusterDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterListRegionsResponse struct {
	Regions []ClusterListRegionsResponseRegion `json:"regions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Regions     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterListRegionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ClusterListRegionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterListRegionsResponseRegion struct {
	// List of supported identifiable cuda/nvidia driver versions pairs available in
	// the region.
	DriverVersions []ClusterListRegionsResponseRegionDriverVersion `json:"driver_versions" api:"required"`
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
func (r ClusterListRegionsResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *ClusterListRegionsResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NVIDIA software configuration available in the region.
type ClusterListRegionsResponseRegionDriverVersion struct {
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
func (r ClusterListRegionsResponseRegionDriverVersion) RawJSON() string { return r.JSON.raw }
func (r *ClusterListRegionsResponseRegionDriverVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Create
	// a cluster with a canonical NVIDIA version id. Do not also set cuda_version or
	// nvidia_driver_version.
	OfGPUClusterCreateRequestNvidiaVersion *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Create
	// a cluster with the legacy CUDA and NVIDIA driver selectors. nvidia_version_id
	// may also be set when it resolves to the same catalog entry.
	OfGPUClusterCreateRequestLegacyNvidia *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia `json:",inline"`

	paramObj
}

func (u ClusterNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGPUClusterCreateRequestNvidiaVersion, u.OfGPUClusterCreateRequestLegacyNvidia)
}
func (r *ClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create a cluster with a canonical NVIDIA version id. Do not also set
// cuda_version or nvidia_driver_version.
//
// The properties BillingType, ClusterName, GPUType, NumGPUs, NvidiaVersionID,
// Region are required.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion struct {
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
	AcceptanceTestsParams ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams `json:"acceptance_tests_params,omitzero"`
	// Add-ons to enable on the cluster at creation time.
	AddOns        []ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn       `json:"add_ons,omitzero"`
	ClusterConfig ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig `json:"cluster_config,omitzero"`
	// Type of cluster to create.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType string                                                             `json:"cluster_type,omitzero"`
	OidcConfig  ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig `json:"oidc_config,omitzero"`
	// Inline configuration to create a shared volume with the cluster creation.
	SharedVolume ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume `json:"shared_volume,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion](
		"billing_type", "RESERVED", "ON_DEMAND", "SCHEDULED_CAPACITY",
	)
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion](
		"gpu_type", "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM", "H100_SXM_INF", "B300_SXM",
	)
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersion](
		"cluster_type", "KUBERNETES", "SLURM",
	)
}

// AcceptanceTestsParams groups all GPU acceptance test options when enabled is
// true.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams struct {
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

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAcceptanceTestsParams](
		"dcgm_diag_level", "DCGM_DIAG_LEVEL_SHORT", "DCGM_DIAG_LEVEL_MEDIUM", "DCGM_DIAG_LEVEL_LONG", "DCGM_DIAG_LEVEL_EXTENDED",
	)
}

// The properties AddOnType, Name are required.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn struct {
	// Type of add-on. Valid values: 'dashboard', 'ingress', 'torchpass'.
	AddOnType string `json:"add_on_type" api:"required"`
	// Human-readable name for this add-on instance.
	Name string `json:"name" api:"required"`
	// Configuration for a cluster add-on.
	Config ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig `json:"config,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a cluster add-on.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig struct {
	Dashboard ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard `json:"dashboard,omitzero"`
	// Configuration for the Headlamp Kubernetes dashboard add-on.
	Headlamp ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp `json:"headlamp,omitzero"`
	Ingress  ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress  `json:"ingress,omitzero"`
	// Configuration for the Slurm Web add-on.
	SlurmWeb ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb `json:"slurm_web,omitzero"`
	// Configuration for the Model Aware TorchPass add-on.
	Torchpass ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass `json:"torchpass,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigDashboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Headlamp Kubernetes dashboard add-on.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp struct {
	// Whether to enable the Headlamp Kubernetes dashboard add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigHeadlamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Slurm Web add-on.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb struct {
	// Whether to enable the Slurm Web add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigSlurmWeb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Model Aware TorchPass add-on.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass struct {
	// Whether to enable the Model Aware TorchPass add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionAddOnConfigTorchpass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LoadBalancer is required.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig struct {
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
	SSHCaEnabled  param.Opt[bool]                                                                    `json:"ssh_ca_enabled,omitzero"`
	Ingress       ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress       `json:"ingress,omitzero"`
	Observability ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability `json:"observability,omitzero"`
	// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
	// init, extra conf).
	SlurmStartupScripts ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts `json:"slurm_startup_scripts,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfig](
		"load_balancer", "NONE", "TRAEFIK", "NGINX", "ISTIO",
	)
}

type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigObservability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
// init, extra conf).
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts struct {
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

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionClusterConfigSlurmStartupScripts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ClientID, GroupClaim, GroupPrefix, IssuerURL, UsernameClaim,
// UsernamePrefix are required.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig struct {
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

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionOidcConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline configuration to create a shared volume with the cluster creation.
//
// The properties Region, SizeTib, VolumeName are required.
type ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume struct {
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

func (r ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestNvidiaVersionSharedVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create a cluster with the legacy CUDA and NVIDIA driver selectors.
// nvidia_version_id may also be set when it resolves to the same catalog entry.
//
// The properties BillingType, ClusterName, CudaVersion, GPUType, NumGPUs,
// NvidiaDriverVersion, Region are required.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia struct {
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
	AcceptanceTestsParams ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams `json:"acceptance_tests_params,omitzero"`
	// Add-ons to enable on the cluster at creation time.
	AddOns        []ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn       `json:"add_ons,omitzero"`
	ClusterConfig ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig `json:"cluster_config,omitzero"`
	// Type of cluster to create.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType string                                                            `json:"cluster_type,omitzero"`
	OidcConfig  ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig `json:"oidc_config,omitzero"`
	// Inline configuration to create a shared volume with the cluster creation.
	SharedVolume ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume `json:"shared_volume,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia](
		"billing_type", "RESERVED", "ON_DEMAND", "SCHEDULED_CAPACITY",
	)
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia](
		"gpu_type", "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM", "H100_SXM_INF", "B300_SXM",
	)
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidia](
		"cluster_type", "KUBERNETES", "SLURM",
	)
}

// AcceptanceTestsParams groups all GPU acceptance test options when enabled is
// true.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams struct {
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

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAcceptanceTestsParams](
		"dcgm_diag_level", "DCGM_DIAG_LEVEL_SHORT", "DCGM_DIAG_LEVEL_MEDIUM", "DCGM_DIAG_LEVEL_LONG", "DCGM_DIAG_LEVEL_EXTENDED",
	)
}

// The properties AddOnType, Name are required.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn struct {
	// Type of add-on. Valid values: 'dashboard', 'ingress', 'torchpass'.
	AddOnType string `json:"add_on_type" api:"required"`
	// Human-readable name for this add-on instance.
	Name string `json:"name" api:"required"`
	// Configuration for a cluster add-on.
	Config ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig `json:"config,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a cluster add-on.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig struct {
	Dashboard ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard `json:"dashboard,omitzero"`
	// Configuration for the Headlamp Kubernetes dashboard add-on.
	Headlamp ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp `json:"headlamp,omitzero"`
	Ingress  ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress  `json:"ingress,omitzero"`
	// Configuration for the Slurm Web add-on.
	SlurmWeb ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb `json:"slurm_web,omitzero"`
	// Configuration for the Model Aware TorchPass add-on.
	Torchpass ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass `json:"torchpass,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigDashboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Headlamp Kubernetes dashboard add-on.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp struct {
	// Whether to enable the Headlamp Kubernetes dashboard add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigHeadlamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Slurm Web add-on.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb struct {
	// Whether to enable the Slurm Web add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigSlurmWeb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Model Aware TorchPass add-on.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass struct {
	// Whether to enable the Model Aware TorchPass add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaAddOnConfigTorchpass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LoadBalancer is required.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig struct {
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
	SSHCaEnabled  param.Opt[bool]                                                                   `json:"ssh_ca_enabled,omitzero"`
	Ingress       ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress       `json:"ingress,omitzero"`
	Observability ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability `json:"observability,omitzero"`
	// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
	// init, extra conf).
	SlurmStartupScripts ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts `json:"slurm_startup_scripts,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfig](
		"load_balancer", "NONE", "TRAEFIK", "NGINX", "ISTIO",
	)
}

type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigObservability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
// init, extra conf).
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts struct {
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

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaClusterConfigSlurmStartupScripts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ClientID, GroupClaim, GroupPrefix, IssuerURL, UsernameClaim,
// UsernamePrefix are required.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig struct {
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

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaOidcConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline configuration to create a shared volume with the cluster creation.
//
// The properties Region, SizeTib, VolumeName are required.
type ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume struct {
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

func (r ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParamsBodyGPUClusterCreateRequestLegacyNvidiaSharedVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterUpdateParams struct {
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
	AddOns        []ClusterUpdateParamsAddOn       `json:"add_ons,omitzero"`
	ClusterConfig ClusterUpdateParamsClusterConfig `json:"cluster_config,omitzero"`
	// Type of cluster to update.
	//
	// Any of "KUBERNETES", "SLURM".
	ClusterType ClusterUpdateParamsClusterType `json:"cluster_type,omitzero"`
	paramObj
}

func (r ClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type ClusterUpdateParamsAddOn struct {
	// Name of the add-on to update. Must match an existing add-on on the cluster.
	Name string `json:"name" api:"required"`
	// Configuration for a cluster add-on.
	Config ClusterUpdateParamsAddOnConfig `json:"config,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsAddOn) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsAddOn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsAddOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a cluster add-on.
type ClusterUpdateParamsAddOnConfig struct {
	Dashboard ClusterUpdateParamsAddOnConfigDashboard `json:"dashboard,omitzero"`
	// Configuration for the Headlamp Kubernetes dashboard add-on.
	Headlamp ClusterUpdateParamsAddOnConfigHeadlamp `json:"headlamp,omitzero"`
	Ingress  ClusterUpdateParamsAddOnConfigIngress  `json:"ingress,omitzero"`
	// Configuration for the Slurm Web add-on.
	SlurmWeb ClusterUpdateParamsAddOnConfigSlurmWeb `json:"slurm_web,omitzero"`
	// Configuration for the Model Aware TorchPass add-on.
	Torchpass ClusterUpdateParamsAddOnConfigTorchpass `json:"torchpass,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsAddOnConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsAddOnConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsAddOnConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterUpdateParamsAddOnConfigDashboard struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsAddOnConfigDashboard) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsAddOnConfigDashboard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsAddOnConfigDashboard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Headlamp Kubernetes dashboard add-on.
type ClusterUpdateParamsAddOnConfigHeadlamp struct {
	// Whether to enable the Headlamp Kubernetes dashboard add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsAddOnConfigHeadlamp) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsAddOnConfigHeadlamp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsAddOnConfigHeadlamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterUpdateParamsAddOnConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsAddOnConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsAddOnConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsAddOnConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Slurm Web add-on.
type ClusterUpdateParamsAddOnConfigSlurmWeb struct {
	// Whether to enable the Slurm Web add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsAddOnConfigSlurmWeb) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsAddOnConfigSlurmWeb
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsAddOnConfigSlurmWeb) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Model Aware TorchPass add-on.
type ClusterUpdateParamsAddOnConfigTorchpass struct {
	// Whether to enable the Model Aware TorchPass add-on.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsAddOnConfigTorchpass) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsAddOnConfigTorchpass
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsAddOnConfigTorchpass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LoadBalancer is required.
type ClusterUpdateParamsClusterConfig struct {
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
	SSHCaEnabled  param.Opt[bool]                               `json:"ssh_ca_enabled,omitzero"`
	Ingress       ClusterUpdateParamsClusterConfigIngress       `json:"ingress,omitzero"`
	Observability ClusterUpdateParamsClusterConfigObservability `json:"observability,omitzero"`
	// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
	// init, extra conf).
	SlurmStartupScripts ClusterUpdateParamsClusterConfigSlurmStartupScripts `json:"slurm_startup_scripts,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsClusterConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsClusterConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsClusterConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ClusterUpdateParamsClusterConfig](
		"load_balancer", "NONE", "TRAEFIK", "NGINX", "ISTIO",
	)
}

type ClusterUpdateParamsClusterConfigIngress struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsClusterConfigIngress) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsClusterConfigIngress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsClusterConfigIngress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterUpdateParamsClusterConfigObservability struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ClusterUpdateParamsClusterConfigObservability) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsClusterConfigObservability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsClusterConfigObservability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog,
// init, extra conf).
type ClusterUpdateParamsClusterConfigSlurmStartupScripts struct {
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

func (r ClusterUpdateParamsClusterConfigSlurmStartupScripts) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParamsClusterConfigSlurmStartupScripts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParamsClusterConfigSlurmStartupScripts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cluster to update.
type ClusterUpdateParamsClusterType string

const (
	ClusterUpdateParamsClusterTypeKubernetes ClusterUpdateParamsClusterType = "KUBERNETES"
	ClusterUpdateParamsClusterTypeSlurm      ClusterUpdateParamsClusterType = "SLURM"
)

type ClusterListParams struct {
	// Optional UMS project ID to filter clusters by. When set, only clusters belonging
	// to this project are returned. The caller must be a member of the project;
	// otherwise the result set will be empty.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `query:"projectId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterListParams]'s query parameters as `url.Values`.
func (r ClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
