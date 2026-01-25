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

// BetaJigService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaJigService] method instead.
type BetaJigService struct {
	Options []option.RequestOption
	Volumes BetaJigVolumeService
	Secrets BetaJigSecretService
}

// NewBetaJigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBetaJigService(opts ...option.RequestOption) (r BetaJigService) {
	r = BetaJigService{}
	r.Options = opts
	r.Volumes = NewBetaJigVolumeService(opts...)
	r.Secrets = NewBetaJigSecretService(opts...)
	return
}

// Retrieve details of a specific deployment by its ID or name
func (r *BetaJigService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("deployments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing deployment configuration
func (r *BetaJigService) Update(ctx context.Context, id string, body BetaJigUpdateParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("deployments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all deployments in your project
func (r *BetaJigService) List(ctx context.Context, opts ...option.RequestOption) (res *BetaJigListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create a new deployment with specified configuration
func (r *BetaJigService) Deploy(ctx context.Context, body BetaJigDeployParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete an existing deployment
func (r *BetaJigService) Destroy(ctx context.Context, id string, opts ...option.RequestOption) (res *BetaJigDestroyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("deployments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve logs from a deployment, optionally filtered by replica ID. Use
// follow=true to stream logs in real-time.
func (r *BetaJigService) GetLogs(ctx context.Context, id string, query BetaJigGetLogsParams, opts ...option.RequestOption) (res *DeploymentLogs, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("deployments/%s/logs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Deployment struct {
	// ID is the unique identifier of the deployment
	ID string `json:"id"`
	// Args are the arguments passed to the container's command
	Args []string `json:"args"`
	// Autoscaling contains autoscaling configuration parameters for this deployment
	Autoscaling map[string]string `json:"autoscaling"`
	// Command is the entrypoint command run in the container
	Command []string `json:"command"`
	// CPU is the amount of CPU resource allocated to each replica in cores (fractional
	// value is allowed)
	CPU float64 `json:"cpu"`
	// CreatedAt is the ISO8601 timestamp when this deployment was created
	CreatedAt string `json:"created_at"`
	// Description provides a human-readable explanation of the deployment's purpose or
	// content
	Description string `json:"description"`
	// DesiredReplicas is the number of replicas that the orchestrator is targeting
	DesiredReplicas int64 `json:"desired_replicas"`
	// EnvironmentVariables is a list of environment variables set in the container
	EnvironmentVariables []DeploymentEnvironmentVariable `json:"environment_variables"`
	// GPUCount is the number of GPUs allocated to each replica in this deployment
	GPUCount int64 `json:"gpu_count"`
	// GPUType specifies the type of GPU requested (if any) for this deployment
	//
	// Any of "h100-80gb", " a100-80gb".
	GPUType DeploymentGPUType `json:"gpu_type"`
	// HealthCheckPath is the HTTP path used for health checks of the application
	HealthCheckPath string `json:"health_check_path"`
	// Image specifies the container image used for this deployment
	Image string `json:"image"`
	// MaxReplicas is the maximum number of replicas to run for this deployment
	MaxReplicas int64 `json:"max_replicas"`
	// Memory is the amount of memory allocated to each replica in GiB (fractional
	// value is allowed)
	Memory float64 `json:"memory"`
	// MinReplicas is the minimum number of replicas to run for this deployment
	MinReplicas int64 `json:"min_replicas"`
	// Name is the name of the deployment
	Name string `json:"name"`
	// Object is the type identifier for this response (always "deployment")
	Object string `json:"object"`
	// Port is the container port that the deployment exposes
	Port int64 `json:"port"`
	// ReadyReplicas is the current number of replicas that are in the Ready state
	ReadyReplicas int64 `json:"ready_replicas"`
	// ReplicaEvents is a mapping of replica names or IDs to their status events
	ReplicaEvents map[string]DeploymentReplicaEvent `json:"replica_events"`
	// Status represents the overall status of the deployment (e.g., Updating, Scaling,
	// Ready, Failed)
	//
	// Any of "Updating", "Scaling", "Ready", "Failed".
	Status DeploymentStatus `json:"status"`
	// Storage is the amount of storage (in MB or units as defined by the platform)
	// allocated to each replica
	Storage int64 `json:"storage"`
	// UpdatedAt is the ISO8601 timestamp when this deployment was last updated
	UpdatedAt string `json:"updated_at"`
	// Volumes is a list of volume mounts for this deployment
	Volumes []DeploymentVolume `json:"volumes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Args                 respjson.Field
		Autoscaling          respjson.Field
		Command              respjson.Field
		CPU                  respjson.Field
		CreatedAt            respjson.Field
		Description          respjson.Field
		DesiredReplicas      respjson.Field
		EnvironmentVariables respjson.Field
		GPUCount             respjson.Field
		GPUType              respjson.Field
		HealthCheckPath      respjson.Field
		Image                respjson.Field
		MaxReplicas          respjson.Field
		Memory               respjson.Field
		MinReplicas          respjson.Field
		Name                 respjson.Field
		Object               respjson.Field
		Port                 respjson.Field
		ReadyReplicas        respjson.Field
		ReplicaEvents        respjson.Field
		Status               respjson.Field
		Storage              respjson.Field
		UpdatedAt            respjson.Field
		Volumes              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Deployment) RawJSON() string { return r.JSON.raw }
func (r *Deployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentEnvironmentVariable struct {
	// Name is the environment variable name (e.g., "DATABASE_URL"). Must start with a
	// letter or underscore, followed by letters, numbers, or underscores
	Name string `json:"name,required"`
	// Value is the plain text value for the environment variable. Use this for
	// non-sensitive values. Either Value or ValueFromSecret must be set, but not both
	Value string `json:"value"`
	// ValueFromSecret references a secret by name or ID to use as the value. Use this
	// for sensitive values like API keys or passwords. Either Value or ValueFromSecret
	// must be set, but not both
	ValueFromSecret string `json:"value_from_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Value           respjson.Field
		ValueFromSecret respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentEnvironmentVariable) RawJSON() string { return r.JSON.raw }
func (r *DeploymentEnvironmentVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUType specifies the type of GPU requested (if any) for this deployment
type DeploymentGPUType string

const (
	DeploymentGPUTypeH100_80gb DeploymentGPUType = "h100-80gb"
	DeploymentGPUTypeA100_80gb DeploymentGPUType = " a100-80gb"
)

type DeploymentReplicaEvent struct {
	// ContainerStatus provides detailed status information about the container within
	// this replica
	ContainerStatus DeploymentReplicaEventContainerStatus `json:"container_status"`
	// Events is a list of Kubernetes events related to this replica for
	// troubleshooting
	Events []DeploymentReplicaEventEvent `json:"events"`
	// ReplicaCompletedAt is the timestamp when the replica finished execution
	ReplicaCompletedAt string `json:"replica_completed_at"`
	// ReplicaMarkedForTerminationAt is the timestamp when the replica was marked for
	// termination
	ReplicaMarkedForTerminationAt string `json:"replica_marked_for_termination_at"`
	// ReplicaReadySince is the timestamp when the replica became ready to serve
	// traffic
	ReplicaReadySince string `json:"replica_ready_since"`
	// ReplicaRunningSince is the timestamp when the replica entered the running state
	ReplicaRunningSince string `json:"replica_running_since"`
	// ReplicaStartedAt is the timestamp when the replica was created
	ReplicaStartedAt string `json:"replica_started_at"`
	// ReplicaStatus is the current status of the replica (e.g., "Running", "Pending",
	// "Failed")
	ReplicaStatus string `json:"replica_status"`
	// ReplicaStatusMessage provides a human-readable message explaining the replica's
	// status
	ReplicaStatusMessage string `json:"replica_status_message"`
	// ReplicaStatusReason provides a brief machine-readable reason for the replica's
	// status
	ReplicaStatusReason string `json:"replica_status_reason"`
	// ScheduledOnCluster identifies which cluster this replica is scheduled on
	ScheduledOnCluster string `json:"scheduled_on_cluster"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContainerStatus               respjson.Field
		Events                        respjson.Field
		ReplicaCompletedAt            respjson.Field
		ReplicaMarkedForTerminationAt respjson.Field
		ReplicaReadySince             respjson.Field
		ReplicaRunningSince           respjson.Field
		ReplicaStartedAt              respjson.Field
		ReplicaStatus                 respjson.Field
		ReplicaStatusMessage          respjson.Field
		ReplicaStatusReason           respjson.Field
		ScheduledOnCluster            respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentReplicaEvent) RawJSON() string { return r.JSON.raw }
func (r *DeploymentReplicaEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContainerStatus provides detailed status information about the container within
// this replica
type DeploymentReplicaEventContainerStatus struct {
	// FinishedAt is the timestamp when the container finished execution (if
	// terminated)
	FinishedAt string `json:"finishedAt"`
	// Message provides a human-readable message with details about the container's
	// status
	Message string `json:"message"`
	// Name is the name of the container
	Name string `json:"name"`
	// Reason provides a brief machine-readable reason for the container's current
	// status
	Reason string `json:"reason"`
	// StartedAt is the timestamp when the container started execution
	StartedAt string `json:"startedAt"`
	// Status is the current state of the container (e.g., "Running", "Terminated",
	// "Waiting")
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishedAt  respjson.Field
		Message     respjson.Field
		Name        respjson.Field
		Reason      respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentReplicaEventContainerStatus) RawJSON() string { return r.JSON.raw }
func (r *DeploymentReplicaEventContainerStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentReplicaEventEvent struct {
	// Action is the action taken or reported by this event
	Action string `json:"action"`
	// Count is the number of times this event has occurred
	Count int64 `json:"count"`
	// FirstSeen is the timestamp when this event was first observed
	FirstSeen string `json:"first_seen"`
	// LastSeen is the timestamp when this event was last observed
	LastSeen string `json:"last_seen"`
	// Message is a human-readable description of the event
	Message string `json:"message"`
	// Reason is a brief machine-readable reason for this event (e.g., "Pulling",
	// "Started", "Failed")
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Count       respjson.Field
		FirstSeen   respjson.Field
		LastSeen    respjson.Field
		Message     respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentReplicaEventEvent) RawJSON() string { return r.JSON.raw }
func (r *DeploymentReplicaEventEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status represents the overall status of the deployment (e.g., Updating, Scaling,
// Ready, Failed)
type DeploymentStatus string

const (
	DeploymentStatusUpdating DeploymentStatus = "Updating"
	DeploymentStatusScaling  DeploymentStatus = "Scaling"
	DeploymentStatusReady    DeploymentStatus = "Ready"
	DeploymentStatusFailed   DeploymentStatus = "Failed"
)

type DeploymentVolume struct {
	// MountPath is the path in the container where the volume will be mounted (e.g.,
	// "/data")
	MountPath string `json:"mount_path,required"`
	// Name is the name of the volume to mount. Must reference an existing volume by
	// name or ID
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MountPath   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentVolume) RawJSON() string { return r.JSON.raw }
func (r *DeploymentVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentLogs struct {
	Lines []string `json:"lines"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lines       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentLogs) RawJSON() string { return r.JSON.raw }
func (r *DeploymentLogs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigListResponse struct {
	// Data is the array of deployment items
	Data []Deployment `json:"data"`
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
func (r BetaJigListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigDestroyResponse = any

type BetaJigUpdateParams struct {
	// CPU is the number of CPU cores to allocate per container instance (e.g., 0.1 =
	// 100 milli cores)
	CPU param.Opt[float64] `json:"cpu,omitzero"`
	// Description is an optional human-readable description of your deployment
	Description param.Opt[string] `json:"description,omitzero"`
	// GPUCount is the number of GPUs to allocate per container instance
	GPUCount param.Opt[int64] `json:"gpu_count,omitzero"`
	// HealthCheckPath is the HTTP path for health checks (e.g., "/health"). Set to
	// empty string to disable health checks
	HealthCheckPath param.Opt[string] `json:"health_check_path,omitzero"`
	// Image is the container image to deploy from registry.together.ai.
	Image param.Opt[string] `json:"image,omitzero"`
	// MaxReplicas is the maximum number of replicas that can be scaled up to.
	MaxReplicas param.Opt[int64] `json:"max_replicas,omitzero"`
	// Memory is the amount of RAM to allocate per container instance in GiB (e.g., 0.5
	// = 512MiB)
	Memory param.Opt[float64] `json:"memory,omitzero"`
	// MinReplicas is the minimum number of replicas to run
	MinReplicas param.Opt[int64] `json:"min_replicas,omitzero"`
	// Name is the new unique identifier for your deployment. Must contain only
	// alphanumeric characters, underscores, or hyphens (1-100 characters)
	Name param.Opt[string] `json:"name,omitzero"`
	// Port is the container port your application listens on (e.g., 8080 for web
	// servers)
	Port param.Opt[int64] `json:"port,omitzero"`
	// Storage is the amount of ephemeral disk storage to allocate per container
	// instance (e.g., 10 = 10GiB)
	Storage param.Opt[int64] `json:"storage,omitzero"`
	// TerminationGracePeriodSeconds is the time in seconds to wait for graceful
	// shutdown before forcefully terminating the replica
	TerminationGracePeriodSeconds param.Opt[int64] `json:"termination_grace_period_seconds,omitzero"`
	// Args overrides the container's CMD. Provide as an array of arguments (e.g.,
	// ["python", "app.py"])
	Args []string `json:"args,omitzero"`
	// Autoscaling configuration as key-value pairs. Example: {"metric":
	// "QueueBacklogPerWorker", "target": "10"} to scale based on queue backlog
	Autoscaling map[string]string `json:"autoscaling,omitzero"`
	// Command overrides the container's ENTRYPOINT. Provide as an array (e.g.,
	// ["/bin/sh", "-c"])
	Command []string `json:"command,omitzero"`
	// EnvironmentVariables is a list of environment variables to set in the container.
	// This will replace all existing environment variables
	EnvironmentVariables []BetaJigUpdateParamsEnvironmentVariable `json:"environment_variables,omitzero"`
	// GPUType specifies the GPU hardware to use (e.g., "h100-80gb")
	//
	// Any of "h100-80gb", " a100-80gb".
	GPUType BetaJigUpdateParamsGPUType `json:"gpu_type,omitzero"`
	// Volumes is a list of volume mounts to attach to the container. This will replace
	// all existing volumes
	Volumes []BetaJigUpdateParamsVolume `json:"volumes,omitzero"`
	paramObj
}

func (r BetaJigUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type BetaJigUpdateParamsEnvironmentVariable struct {
	// Name is the environment variable name (e.g., "DATABASE_URL"). Must start with a
	// letter or underscore, followed by letters, numbers, or underscores
	Name string `json:"name,required"`
	// Value is the plain text value for the environment variable. Use this for
	// non-sensitive values. Either Value or ValueFromSecret must be set, but not both
	Value param.Opt[string] `json:"value,omitzero"`
	// ValueFromSecret references a secret by name or ID to use as the value. Use this
	// for sensitive values like API keys or passwords. Either Value or ValueFromSecret
	// must be set, but not both
	ValueFromSecret param.Opt[string] `json:"value_from_secret,omitzero"`
	paramObj
}

func (r BetaJigUpdateParamsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigUpdateParamsEnvironmentVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigUpdateParamsEnvironmentVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUType specifies the GPU hardware to use (e.g., "h100-80gb")
type BetaJigUpdateParamsGPUType string

const (
	BetaJigUpdateParamsGPUTypeH100_80gb BetaJigUpdateParamsGPUType = "h100-80gb"
	BetaJigUpdateParamsGPUTypeA100_80gb BetaJigUpdateParamsGPUType = " a100-80gb"
)

// The properties MountPath, Name are required.
type BetaJigUpdateParamsVolume struct {
	// MountPath is the path in the container where the volume will be mounted (e.g.,
	// "/data")
	MountPath string `json:"mount_path,required"`
	// Name is the name of the volume to mount. Must reference an existing volume by
	// name or ID
	Name string `json:"name,required"`
	paramObj
}

func (r BetaJigUpdateParamsVolume) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigUpdateParamsVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigUpdateParamsVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigDeployParams struct {
	// GPUType specifies the GPU hardware to use (e.g., "h100-80gb").
	//
	// Any of "h100-80gb", " a100-80gb".
	GPUType BetaJigDeployParamsGPUType `json:"gpu_type,omitzero,required"`
	// Image is the container image to deploy from registry.together.ai.
	Image string `json:"image,required"`
	// Name is the unique identifier for your deployment. Must contain only
	// alphanumeric characters, underscores, or hyphens (1-100 characters)
	Name string `json:"name,required"`
	// CPU is the number of CPU cores to allocate per container instance (e.g., 0.1 =
	// 100 milli cores)
	CPU param.Opt[float64] `json:"cpu,omitzero"`
	// Description is an optional human-readable description of your deployment
	Description param.Opt[string] `json:"description,omitzero"`
	// GPUCount is the number of GPUs to allocate per container instance. Defaults to 0
	// if not specified
	GPUCount param.Opt[int64] `json:"gpu_count,omitzero"`
	// HealthCheckPath is the HTTP path for health checks (e.g., "/health"). If set,
	// the platform will check this endpoint to determine container health
	HealthCheckPath param.Opt[string] `json:"health_check_path,omitzero"`
	// MaxReplicas is the maximum number of container instances that can be scaled up
	// to. If not set, will be set to MinReplicas
	MaxReplicas param.Opt[int64] `json:"max_replicas,omitzero"`
	// Memory is the amount of RAM to allocate per container instance in GiB (e.g., 0.5
	// = 512MiB)
	Memory param.Opt[float64] `json:"memory,omitzero"`
	// MinReplicas is the minimum number of container instances to run. Defaults to 1
	// if not specified
	MinReplicas param.Opt[int64] `json:"min_replicas,omitzero"`
	// Port is the container port your application listens on (e.g., 8080 for web
	// servers). Required if your application serves traffic
	Port param.Opt[int64] `json:"port,omitzero"`
	// Storage is the amount of ephemeral disk storage to allocate per container
	// instance (e.g., 10 = 10GiB)
	Storage param.Opt[int64] `json:"storage,omitzero"`
	// TerminationGracePeriodSeconds is the time in seconds to wait for graceful
	// shutdown before forcefully terminating the replica
	TerminationGracePeriodSeconds param.Opt[int64] `json:"termination_grace_period_seconds,omitzero"`
	// Args overrides the container's CMD. Provide as an array of arguments (e.g.,
	// ["python", "app.py"])
	Args []string `json:"args,omitzero"`
	// Autoscaling configuration as key-value pairs. Example: {"metric":
	// "QueueBacklogPerWorker", "target": "10"} to scale based on queue backlog
	Autoscaling map[string]string `json:"autoscaling,omitzero"`
	// Command overrides the container's ENTRYPOINT. Provide as an array (e.g.,
	// ["/bin/sh", "-c"])
	Command []string `json:"command,omitzero"`
	// EnvironmentVariables is a list of environment variables to set in the container.
	// Each must have a name and either a value or value_from_secret
	EnvironmentVariables []BetaJigDeployParamsEnvironmentVariable `json:"environment_variables,omitzero"`
	// Volumes is a list of volume mounts to attach to the container. Each mount must
	// reference an existing volume by name
	Volumes []BetaJigDeployParamsVolume `json:"volumes,omitzero"`
	paramObj
}

func (r BetaJigDeployParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigDeployParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigDeployParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUType specifies the GPU hardware to use (e.g., "h100-80gb").
type BetaJigDeployParamsGPUType string

const (
	BetaJigDeployParamsGPUTypeH100_80gb BetaJigDeployParamsGPUType = "h100-80gb"
	BetaJigDeployParamsGPUTypeA100_80gb BetaJigDeployParamsGPUType = " a100-80gb"
)

// The property Name is required.
type BetaJigDeployParamsEnvironmentVariable struct {
	// Name is the environment variable name (e.g., "DATABASE_URL"). Must start with a
	// letter or underscore, followed by letters, numbers, or underscores
	Name string `json:"name,required"`
	// Value is the plain text value for the environment variable. Use this for
	// non-sensitive values. Either Value or ValueFromSecret must be set, but not both
	Value param.Opt[string] `json:"value,omitzero"`
	// ValueFromSecret references a secret by name or ID to use as the value. Use this
	// for sensitive values like API keys or passwords. Either Value or ValueFromSecret
	// must be set, but not both
	ValueFromSecret param.Opt[string] `json:"value_from_secret,omitzero"`
	paramObj
}

func (r BetaJigDeployParamsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigDeployParamsEnvironmentVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigDeployParamsEnvironmentVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MountPath, Name are required.
type BetaJigDeployParamsVolume struct {
	// MountPath is the path in the container where the volume will be mounted (e.g.,
	// "/data")
	MountPath string `json:"mount_path,required"`
	// Name is the name of the volume to mount. Must reference an existing volume by
	// name or ID
	Name string `json:"name,required"`
	paramObj
}

func (r BetaJigDeployParamsVolume) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigDeployParamsVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigDeployParamsVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigGetLogsParams struct {
	// Stream logs in real-time (ndjson format)
	Follow param.Opt[bool] `query:"follow,omitzero" json:"-"`
	// Replica ID to filter logs
	ReplicaID param.Opt[string] `query:"replica_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaJigGetLogsParams]'s query parameters as `url.Values`.
func (r BetaJigGetLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
