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

// BetaJigService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaJigService] method instead.
type BetaJigService struct {
	Options []option.RequestOption
	Queue   BetaJigQueueService
	Volumes BetaJigVolumeService
	Secrets BetaJigSecretService
}

// NewBetaJigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBetaJigService(opts ...option.RequestOption) (r BetaJigService) {
	r = BetaJigService{}
	r.Options = opts
	r.Queue = NewBetaJigQueueService(opts...)
	r.Volumes = NewBetaJigVolumeService(opts...)
	r.Secrets = NewBetaJigSecretService(opts...)
	return
}

// Retrieve details of a specific deployment by its ID or name
func (r *BetaJigService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing deployment configuration
func (r *BetaJigService) Update(ctx context.Context, id string, body BetaJigUpdateParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get a list of all deployments in your project
func (r *BetaJigService) List(ctx context.Context, opts ...option.RequestOption) (res *BetaJigListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a new deployment with specified configuration
func (r *BetaJigService) Deploy(ctx context.Context, body BetaJigDeployParams, opts ...option.RequestOption) (res *Deployment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete an existing deployment
func (r *BetaJigService) Destroy(ctx context.Context, id string, opts ...option.RequestOption) (res *BetaJigDestroyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieve logs from a deployment, optionally filtered by replica ID.
func (r *BetaJigService) GetLogs(ctx context.Context, id string, query BetaJigGetLogsParams, opts ...option.RequestOption) (res *DeploymentLogs, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("deployments/%s/logs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Deployment struct {
	// ID is the unique identifier of the deployment
	ID string `json:"id"`
	// Args are the arguments passed to the container's command
	Args []string `json:"args"`
	// Autoscaling contains autoscaling configuration parameters for this deployment.
	// Omitted when autoscaling is disabled (nil)
	Autoscaling DeploymentAutoscalingUnion `json:"autoscaling"`
	// Command is the entrypoint command run in the container
	Command []string `json:"command"`
	// CPU is the amount of CPU resource allocated to each replica in cores (fractional
	// value is allowed)
	CPU float64 `json:"cpu"`
	// CreatedAt is the ISO8601 timestamp when this deployment was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
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
	// The object type, which is always `deployment`.
	//
	// Any of "deployment".
	Object DeploymentObject `json:"object"`
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
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
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

// DeploymentAutoscalingUnion contains all possible properties and values from
// [DeploymentAutoscalingHTTPAutoscalingConfig],
// [DeploymentAutoscalingQueueAutoscalingConfig],
// [DeploymentAutoscalingCustomMetricAutoscalingConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DeploymentAutoscalingUnion struct {
	Metric string  `json:"metric"`
	Target float64 `json:"target"`
	// This field is from variant [DeploymentAutoscalingHTTPAutoscalingConfig].
	TimeIntervalMinutes int64 `json:"time_interval_minutes"`
	// This field is from variant [DeploymentAutoscalingQueueAutoscalingConfig].
	Model string `json:"model"`
	// This field is from variant [DeploymentAutoscalingCustomMetricAutoscalingConfig].
	CustomMetricName string `json:"custom_metric_name"`
	JSON             struct {
		Metric              respjson.Field
		Target              respjson.Field
		TimeIntervalMinutes respjson.Field
		Model               respjson.Field
		CustomMetricName    respjson.Field
		raw                 string
	} `json:"-"`
}

func (u DeploymentAutoscalingUnion) AsDeploymentAutoscalingHTTPAutoscalingConfig() (v DeploymentAutoscalingHTTPAutoscalingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DeploymentAutoscalingUnion) AsDeploymentAutoscalingQueueAutoscalingConfig() (v DeploymentAutoscalingQueueAutoscalingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DeploymentAutoscalingUnion) AsDeploymentAutoscalingCustomMetricAutoscalingConfig() (v DeploymentAutoscalingCustomMetricAutoscalingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DeploymentAutoscalingUnion) RawJSON() string { return u.JSON.raw }

func (r *DeploymentAutoscalingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Autoscaling config for HTTPTotalRequests and HTTPAvgRequestDuration metrics
type DeploymentAutoscalingHTTPAutoscalingConfig struct {
	// Metric must be HTTPTotalRequests or HTTPAvgRequestDuration
	//
	// Any of "HTTPTotalRequests", "HTTPAvgRequestDuration".
	Metric string `json:"metric"`
	// Target is the threshold value. Default: 100 for HTTPTotalRequests, 500 (ms) for
	// HTTPAvgRequestDuration
	Target float64 `json:"target"`
	// TimeIntervalMinutes is the rate window in minutes. Default: 10
	TimeIntervalMinutes int64 `json:"time_interval_minutes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metric              respjson.Field
		Target              respjson.Field
		TimeIntervalMinutes respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentAutoscalingHTTPAutoscalingConfig) RawJSON() string { return r.JSON.raw }
func (r *DeploymentAutoscalingHTTPAutoscalingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Autoscaling config for QueueBacklogPerWorker metric
type DeploymentAutoscalingQueueAutoscalingConfig struct {
	// Metric must be QueueBacklogPerWorker
	//
	// Any of "QueueBacklogPerWorker".
	Metric string `json:"metric"`
	// Model overrides the model name for queue status lookup. Defaults to the
	// deployment app name
	Model string `json:"model"`
	// Target is the threshold value. Default: 1.01
	Target float64 `json:"target"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metric      respjson.Field
		Model       respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentAutoscalingQueueAutoscalingConfig) RawJSON() string { return r.JSON.raw }
func (r *DeploymentAutoscalingQueueAutoscalingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Autoscaling config for CustomMetric metric
type DeploymentAutoscalingCustomMetricAutoscalingConfig struct {
	// CustomMetricName is the Prometheus metric name. Required. Must match
	// [a-zA-Z\_:][a-zA-Z0-9_:]\*
	CustomMetricName string `json:"custom_metric_name"`
	// Metric must be CustomMetric
	//
	// Any of "CustomMetric".
	Metric string `json:"metric"`
	// Target is the threshold value. Default: 500
	Target float64 `json:"target"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomMetricName respjson.Field
		Metric           respjson.Field
		Target           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentAutoscalingCustomMetricAutoscalingConfig) RawJSON() string { return r.JSON.raw }
func (r *DeploymentAutoscalingCustomMetricAutoscalingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentEnvironmentVariable struct {
	// Name is the environment variable name (e.g., "DATABASE_URL"). Must start with a
	// letter or underscore, followed by letters, numbers, or underscores
	Name string `json:"name" api:"required"`
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

// The object type, which is always `deployment`.
type DeploymentObject string

const (
	DeploymentObjectDeployment DeploymentObject = "deployment"
)

type DeploymentReplicaEvent struct {
	// Image is the container image used for this replica
	Image string `json:"image"`
	// ReplicaReadySince is the timestamp when the replica became ready to serve
	// traffic
	ReplicaReadySince string `json:"replica_ready_since"`
	// ReplicaStatus is the current status of the replica (e.g., "Running", "Waiting",
	// "Terminated")
	ReplicaStatus string `json:"replica_status"`
	// ReplicaStatusMessage provides a human-readable message explaining the replica's
	// status
	ReplicaStatusMessage string `json:"replica_status_message"`
	// ReplicaStatusReason provides a brief machine-readable reason for the replica's
	// status
	ReplicaStatusReason string `json:"replica_status_reason"`
	// RevisionID is the deployment revision ID associated with this replica
	RevisionID string `json:"revision_id"`
	// VolumePreloadCompletedAt is the timestamp when the volume preload completed
	VolumePreloadCompletedAt string `json:"volume_preload_completed_at"`
	// VolumePreloadStartedAt is the timestamp when the volume preload started
	VolumePreloadStartedAt string `json:"volume_preload_started_at"`
	// VolumePreloadStatus is the status of the volume preload (e.g., "InProgress",
	// "Completed", "Failed")
	VolumePreloadStatus string `json:"volume_preload_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image                    respjson.Field
		ReplicaReadySince        respjson.Field
		ReplicaStatus            respjson.Field
		ReplicaStatusMessage     respjson.Field
		ReplicaStatusReason      respjson.Field
		RevisionID               respjson.Field
		VolumePreloadCompletedAt respjson.Field
		VolumePreloadStartedAt   respjson.Field
		VolumePreloadStatus      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentReplicaEvent) RawJSON() string { return r.JSON.raw }
func (r *DeploymentReplicaEvent) UnmarshalJSON(data []byte) error {
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
	MountPath string `json:"mount_path" api:"required"`
	// Name is the name of the volume to mount. Must reference an existing volume by
	// name or ID
	Name string `json:"name" api:"required"`
	// Version is the volume version to mount. On create, defaults to the latest
	// version. On update, defaults to the currently mounted version.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MountPath   respjson.Field
		Name        respjson.Field
		Version     respjson.Field
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
	// The object type, which is always `list`.
	//
	// Any of "list".
	Object BetaJigListResponseObject `json:"object"`
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

// The object type, which is always `list`.
type BetaJigListResponseObject string

const (
	BetaJigListResponseObjectList BetaJigListResponseObject = "list"
)

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
	// Autoscaling configuration for the deployment. Set to {} to disable autoscaling
	Autoscaling BetaJigUpdateParamsAutoscalingUnion `json:"autoscaling,omitzero"`
	// Command overrides the container's ENTRYPOINT. Provide as an array (e.g.,
	// ["/bin/sh", "-c"])
	Command []string `json:"command,omitzero"`
	// EnvironmentVariables is a list of environment variables to set in the container.
	// This will replace all existing environment variables
	EnvironmentVariables []BetaJigUpdateParamsEnvironmentVariable `json:"environment_variables,omitzero"`
	// GPUType specifies the GPU hardware to use (e.g., "h100-80gb")
	//
	// Any of "h100-80gb".
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaJigUpdateParamsAutoscalingUnion struct {
	OfBetaJigUpdatesAutoscalingHTTPAutoscalingConfig         *BetaJigUpdateParamsAutoscalingHTTPAutoscalingConfig         `json:",omitzero,inline"`
	OfBetaJigUpdatesAutoscalingQueueAutoscalingConfig        *BetaJigUpdateParamsAutoscalingQueueAutoscalingConfig        `json:",omitzero,inline"`
	OfBetaJigUpdatesAutoscalingCustomMetricAutoscalingConfig *BetaJigUpdateParamsAutoscalingCustomMetricAutoscalingConfig `json:",omitzero,inline"`
	paramUnion
}

func (u BetaJigUpdateParamsAutoscalingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBetaJigUpdatesAutoscalingHTTPAutoscalingConfig, u.OfBetaJigUpdatesAutoscalingQueueAutoscalingConfig, u.OfBetaJigUpdatesAutoscalingCustomMetricAutoscalingConfig)
}
func (u *BetaJigUpdateParamsAutoscalingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BetaJigUpdateParamsAutoscalingUnion) asAny() any {
	if !param.IsOmitted(u.OfBetaJigUpdatesAutoscalingHTTPAutoscalingConfig) {
		return u.OfBetaJigUpdatesAutoscalingHTTPAutoscalingConfig
	} else if !param.IsOmitted(u.OfBetaJigUpdatesAutoscalingQueueAutoscalingConfig) {
		return u.OfBetaJigUpdatesAutoscalingQueueAutoscalingConfig
	} else if !param.IsOmitted(u.OfBetaJigUpdatesAutoscalingCustomMetricAutoscalingConfig) {
		return u.OfBetaJigUpdatesAutoscalingCustomMetricAutoscalingConfig
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigUpdateParamsAutoscalingUnion) GetTimeIntervalMinutes() *int64 {
	if vt := u.OfBetaJigUpdatesAutoscalingHTTPAutoscalingConfig; vt != nil && vt.TimeIntervalMinutes.Valid() {
		return &vt.TimeIntervalMinutes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigUpdateParamsAutoscalingUnion) GetModel() *string {
	if vt := u.OfBetaJigUpdatesAutoscalingQueueAutoscalingConfig; vt != nil && vt.Model.Valid() {
		return &vt.Model.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigUpdateParamsAutoscalingUnion) GetCustomMetricName() *string {
	if vt := u.OfBetaJigUpdatesAutoscalingCustomMetricAutoscalingConfig; vt != nil && vt.CustomMetricName.Valid() {
		return &vt.CustomMetricName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigUpdateParamsAutoscalingUnion) GetMetric() *string {
	if vt := u.OfBetaJigUpdatesAutoscalingHTTPAutoscalingConfig; vt != nil {
		return (*string)(&vt.Metric)
	} else if vt := u.OfBetaJigUpdatesAutoscalingQueueAutoscalingConfig; vt != nil {
		return (*string)(&vt.Metric)
	} else if vt := u.OfBetaJigUpdatesAutoscalingCustomMetricAutoscalingConfig; vt != nil {
		return (*string)(&vt.Metric)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigUpdateParamsAutoscalingUnion) GetTarget() *float64 {
	if vt := u.OfBetaJigUpdatesAutoscalingHTTPAutoscalingConfig; vt != nil && vt.Target.Valid() {
		return &vt.Target.Value
	} else if vt := u.OfBetaJigUpdatesAutoscalingQueueAutoscalingConfig; vt != nil && vt.Target.Valid() {
		return &vt.Target.Value
	} else if vt := u.OfBetaJigUpdatesAutoscalingCustomMetricAutoscalingConfig; vt != nil && vt.Target.Valid() {
		return &vt.Target.Value
	}
	return nil
}

// Autoscaling config for HTTPTotalRequests and HTTPAvgRequestDuration metrics
type BetaJigUpdateParamsAutoscalingHTTPAutoscalingConfig struct {
	// Target is the threshold value. Default: 100 for HTTPTotalRequests, 500 (ms) for
	// HTTPAvgRequestDuration
	Target param.Opt[float64] `json:"target,omitzero"`
	// TimeIntervalMinutes is the rate window in minutes. Default: 10
	TimeIntervalMinutes param.Opt[int64] `json:"time_interval_minutes,omitzero"`
	// Metric must be HTTPTotalRequests or HTTPAvgRequestDuration
	//
	// Any of "HTTPTotalRequests", "HTTPAvgRequestDuration".
	Metric string `json:"metric,omitzero"`
	paramObj
}

func (r BetaJigUpdateParamsAutoscalingHTTPAutoscalingConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigUpdateParamsAutoscalingHTTPAutoscalingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigUpdateParamsAutoscalingHTTPAutoscalingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaJigUpdateParamsAutoscalingHTTPAutoscalingConfig](
		"metric", "HTTPTotalRequests", "HTTPAvgRequestDuration",
	)
}

// Autoscaling config for QueueBacklogPerWorker metric
type BetaJigUpdateParamsAutoscalingQueueAutoscalingConfig struct {
	// Model overrides the model name for queue status lookup. Defaults to the
	// deployment app name
	Model param.Opt[string] `json:"model,omitzero"`
	// Target is the threshold value. Default: 1.01
	Target param.Opt[float64] `json:"target,omitzero"`
	// Metric must be QueueBacklogPerWorker
	//
	// Any of "QueueBacklogPerWorker".
	Metric string `json:"metric,omitzero"`
	paramObj
}

func (r BetaJigUpdateParamsAutoscalingQueueAutoscalingConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigUpdateParamsAutoscalingQueueAutoscalingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigUpdateParamsAutoscalingQueueAutoscalingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaJigUpdateParamsAutoscalingQueueAutoscalingConfig](
		"metric", "QueueBacklogPerWorker",
	)
}

// Autoscaling config for CustomMetric metric
type BetaJigUpdateParamsAutoscalingCustomMetricAutoscalingConfig struct {
	// CustomMetricName is the Prometheus metric name. Required. Must match
	// [a-zA-Z\_:][a-zA-Z0-9_:]\*
	CustomMetricName param.Opt[string] `json:"custom_metric_name,omitzero"`
	// Target is the threshold value. Default: 500
	Target param.Opt[float64] `json:"target,omitzero"`
	// Metric must be CustomMetric
	//
	// Any of "CustomMetric".
	Metric string `json:"metric,omitzero"`
	paramObj
}

func (r BetaJigUpdateParamsAutoscalingCustomMetricAutoscalingConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigUpdateParamsAutoscalingCustomMetricAutoscalingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigUpdateParamsAutoscalingCustomMetricAutoscalingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaJigUpdateParamsAutoscalingCustomMetricAutoscalingConfig](
		"metric", "CustomMetric",
	)
}

// The property Name is required.
type BetaJigUpdateParamsEnvironmentVariable struct {
	// Name is the environment variable name (e.g., "DATABASE_URL"). Must start with a
	// letter or underscore, followed by letters, numbers, or underscores
	Name string `json:"name" api:"required"`
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
)

// The properties MountPath, Name are required.
type BetaJigUpdateParamsVolume struct {
	// MountPath is the path in the container where the volume will be mounted (e.g.,
	// "/data")
	MountPath string `json:"mount_path" api:"required"`
	// Name is the name of the volume to mount. Must reference an existing volume by
	// name or ID
	Name string `json:"name" api:"required"`
	// Version is the volume version to mount. On create, defaults to the latest
	// version. On update, defaults to the currently mounted version.
	Version param.Opt[int64] `json:"version,omitzero"`
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
	// Any of "h100-80gb".
	GPUType BetaJigDeployParamsGPUType `json:"gpu_type,omitzero" api:"required"`
	// Image is the container image to deploy from registry.together.ai.
	Image string `json:"image" api:"required"`
	// Name is the unique identifier for your deployment. Must contain only
	// alphanumeric characters, underscores, or hyphens (1-100 characters)
	Name string `json:"name" api:"required"`
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
	// Autoscaling configuration. Example: {"metric": "QueueBacklogPerWorker",
	// "target": 1.01} to scale based on queue backlog. Omit or set to null to disable
	// autoscaling
	Autoscaling BetaJigDeployParamsAutoscalingUnion `json:"autoscaling,omitzero"`
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
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaJigDeployParamsAutoscalingUnion struct {
	OfBetaJigDeploysAutoscalingHTTPAutoscalingConfig         *BetaJigDeployParamsAutoscalingHTTPAutoscalingConfig         `json:",omitzero,inline"`
	OfBetaJigDeploysAutoscalingQueueAutoscalingConfig        *BetaJigDeployParamsAutoscalingQueueAutoscalingConfig        `json:",omitzero,inline"`
	OfBetaJigDeploysAutoscalingCustomMetricAutoscalingConfig *BetaJigDeployParamsAutoscalingCustomMetricAutoscalingConfig `json:",omitzero,inline"`
	paramUnion
}

func (u BetaJigDeployParamsAutoscalingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBetaJigDeploysAutoscalingHTTPAutoscalingConfig, u.OfBetaJigDeploysAutoscalingQueueAutoscalingConfig, u.OfBetaJigDeploysAutoscalingCustomMetricAutoscalingConfig)
}
func (u *BetaJigDeployParamsAutoscalingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BetaJigDeployParamsAutoscalingUnion) asAny() any {
	if !param.IsOmitted(u.OfBetaJigDeploysAutoscalingHTTPAutoscalingConfig) {
		return u.OfBetaJigDeploysAutoscalingHTTPAutoscalingConfig
	} else if !param.IsOmitted(u.OfBetaJigDeploysAutoscalingQueueAutoscalingConfig) {
		return u.OfBetaJigDeploysAutoscalingQueueAutoscalingConfig
	} else if !param.IsOmitted(u.OfBetaJigDeploysAutoscalingCustomMetricAutoscalingConfig) {
		return u.OfBetaJigDeploysAutoscalingCustomMetricAutoscalingConfig
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigDeployParamsAutoscalingUnion) GetTimeIntervalMinutes() *int64 {
	if vt := u.OfBetaJigDeploysAutoscalingHTTPAutoscalingConfig; vt != nil && vt.TimeIntervalMinutes.Valid() {
		return &vt.TimeIntervalMinutes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigDeployParamsAutoscalingUnion) GetModel() *string {
	if vt := u.OfBetaJigDeploysAutoscalingQueueAutoscalingConfig; vt != nil && vt.Model.Valid() {
		return &vt.Model.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigDeployParamsAutoscalingUnion) GetCustomMetricName() *string {
	if vt := u.OfBetaJigDeploysAutoscalingCustomMetricAutoscalingConfig; vt != nil && vt.CustomMetricName.Valid() {
		return &vt.CustomMetricName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigDeployParamsAutoscalingUnion) GetMetric() *string {
	if vt := u.OfBetaJigDeploysAutoscalingHTTPAutoscalingConfig; vt != nil {
		return (*string)(&vt.Metric)
	} else if vt := u.OfBetaJigDeploysAutoscalingQueueAutoscalingConfig; vt != nil {
		return (*string)(&vt.Metric)
	} else if vt := u.OfBetaJigDeploysAutoscalingCustomMetricAutoscalingConfig; vt != nil {
		return (*string)(&vt.Metric)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaJigDeployParamsAutoscalingUnion) GetTarget() *float64 {
	if vt := u.OfBetaJigDeploysAutoscalingHTTPAutoscalingConfig; vt != nil && vt.Target.Valid() {
		return &vt.Target.Value
	} else if vt := u.OfBetaJigDeploysAutoscalingQueueAutoscalingConfig; vt != nil && vt.Target.Valid() {
		return &vt.Target.Value
	} else if vt := u.OfBetaJigDeploysAutoscalingCustomMetricAutoscalingConfig; vt != nil && vt.Target.Valid() {
		return &vt.Target.Value
	}
	return nil
}

// Autoscaling config for HTTPTotalRequests and HTTPAvgRequestDuration metrics
type BetaJigDeployParamsAutoscalingHTTPAutoscalingConfig struct {
	// Target is the threshold value. Default: 100 for HTTPTotalRequests, 500 (ms) for
	// HTTPAvgRequestDuration
	Target param.Opt[float64] `json:"target,omitzero"`
	// TimeIntervalMinutes is the rate window in minutes. Default: 10
	TimeIntervalMinutes param.Opt[int64] `json:"time_interval_minutes,omitzero"`
	// Metric must be HTTPTotalRequests or HTTPAvgRequestDuration
	//
	// Any of "HTTPTotalRequests", "HTTPAvgRequestDuration".
	Metric string `json:"metric,omitzero"`
	paramObj
}

func (r BetaJigDeployParamsAutoscalingHTTPAutoscalingConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigDeployParamsAutoscalingHTTPAutoscalingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigDeployParamsAutoscalingHTTPAutoscalingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaJigDeployParamsAutoscalingHTTPAutoscalingConfig](
		"metric", "HTTPTotalRequests", "HTTPAvgRequestDuration",
	)
}

// Autoscaling config for QueueBacklogPerWorker metric
type BetaJigDeployParamsAutoscalingQueueAutoscalingConfig struct {
	// Model overrides the model name for queue status lookup. Defaults to the
	// deployment app name
	Model param.Opt[string] `json:"model,omitzero"`
	// Target is the threshold value. Default: 1.01
	Target param.Opt[float64] `json:"target,omitzero"`
	// Metric must be QueueBacklogPerWorker
	//
	// Any of "QueueBacklogPerWorker".
	Metric string `json:"metric,omitzero"`
	paramObj
}

func (r BetaJigDeployParamsAutoscalingQueueAutoscalingConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigDeployParamsAutoscalingQueueAutoscalingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigDeployParamsAutoscalingQueueAutoscalingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaJigDeployParamsAutoscalingQueueAutoscalingConfig](
		"metric", "QueueBacklogPerWorker",
	)
}

// Autoscaling config for CustomMetric metric
type BetaJigDeployParamsAutoscalingCustomMetricAutoscalingConfig struct {
	// CustomMetricName is the Prometheus metric name. Required. Must match
	// [a-zA-Z\_:][a-zA-Z0-9_:]\*
	CustomMetricName param.Opt[string] `json:"custom_metric_name,omitzero"`
	// Target is the threshold value. Default: 500
	Target param.Opt[float64] `json:"target,omitzero"`
	// Metric must be CustomMetric
	//
	// Any of "CustomMetric".
	Metric string `json:"metric,omitzero"`
	paramObj
}

func (r BetaJigDeployParamsAutoscalingCustomMetricAutoscalingConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigDeployParamsAutoscalingCustomMetricAutoscalingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigDeployParamsAutoscalingCustomMetricAutoscalingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaJigDeployParamsAutoscalingCustomMetricAutoscalingConfig](
		"metric", "CustomMetric",
	)
}

// The property Name is required.
type BetaJigDeployParamsEnvironmentVariable struct {
	// Name is the environment variable name (e.g., "DATABASE_URL"). Must start with a
	// letter or underscore, followed by letters, numbers, or underscores
	Name string `json:"name" api:"required"`
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
	MountPath string `json:"mount_path" api:"required"`
	// Name is the name of the volume to mount. Must reference an existing volume by
	// name or ID
	Name string `json:"name" api:"required"`
	// Version is the volume version to mount. On create, defaults to the latest
	// version. On update, defaults to the currently mounted version.
	Version param.Opt[int64] `json:"version,omitzero"`
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
