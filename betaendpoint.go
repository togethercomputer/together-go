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
	"github.com/togethercomputer/together-go/packages/pagination"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaEndpointService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaEndpointService] method instead.
type BetaEndpointService struct {
	Options           []option.RequestOption
	PlacementProfiles BetaEndpointPlacementProfileService
	AbExperiments     BetaEndpointAbExperimentService
	ShadowExperiments BetaEndpointShadowExperimentService
	Rollouts          BetaEndpointRolloutService
	Hardware          BetaEndpointHardwareService
	Adapters          BetaEndpointAdapterService
	Deployments       BetaEndpointDeploymentService
}

// NewBetaEndpointService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaEndpointService(opts ...option.RequestOption) (r BetaEndpointService) {
	r = BetaEndpointService{}
	r.Options = opts
	r.PlacementProfiles = NewBetaEndpointPlacementProfileService(opts...)
	r.AbExperiments = NewBetaEndpointAbExperimentService(opts...)
	r.ShadowExperiments = NewBetaEndpointShadowExperimentService(opts...)
	r.Rollouts = NewBetaEndpointRolloutService(opts...)
	r.Hardware = NewBetaEndpointHardwareService(opts...)
	r.Adapters = NewBetaEndpointAdapterService(opts...)
	r.Deployments = NewBetaEndpointDeploymentService(opts...)
	return
}

// Creates a stable, inference-addressable endpoint. Add one or more deployments
// and configure its traffic split before sending inference requests to the
// endpoint name.
func (r *BetaEndpointService) New(ctx context.Context, params BetaEndpointNewParams, opts ...option.RequestOption) (res *Endpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves an endpoint and lightweight summaries of the deployments attached to
// it.
func (r *BetaEndpointService) Get(ctx context.Context, id string, query BetaEndpointGetParams, opts ...option.RequestOption) (res *Endpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	if query.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s", query.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates mutable endpoint fields such as its endpoint string, visibility, or
// deployment traffic split. Use `updateMask` to select fields explicitly and
// `etag` in the request body for optimistic concurrency.
func (r *BetaEndpointService) Update(ctx context.Context, id string, params BetaEndpointUpdateParams, opts ...option.RequestOption) (res *Endpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s", params.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the dedicated inference endpoints owned by the specified project.
func (r *BetaEndpointService) List(ctx context.Context, params BetaEndpointListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Endpoint], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints", params.ProjectID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists the dedicated inference endpoints owned by the specified project.
func (r *BetaEndpointService) ListAutoPaging(ctx context.Context, params BetaEndpointListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Endpoint] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Permanently deletes an endpoint. Delete its deployments first; use `etag` to
// reject the request if the endpoint changed after it was read.
func (r *BetaEndpointService) Delete(ctx context.Context, id string, params BetaEndpointDeleteParams, opts ...option.RequestOption) (res *BetaEndpointDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s", params.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Returns aggregated request, token, latency, throughput, error, and
// resource-utilization metrics for an endpoint over a time range. Optionally
// includes time-series buckets and a per-deployment breakdown.
func (r *BetaEndpointService) Analytics(ctx context.Context, id string, params BetaEndpointAnalyticsParams, opts ...option.RequestOption) (res *BetaEndpointAnalyticsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/analytics", params.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Lists an endpoint's audit and lifecycle events newest first. The feed combines
// endpoint changes with provisioning, scaling, readiness, rollout, and other
// events from deployments under the endpoint.
func (r *BetaEndpointService) ListEvents(ctx context.Context, id string, params BetaEndpointListEventsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[BetaEndpointListEventsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/events", params.ProjectID.Value, id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists an endpoint's audit and lifecycle events newest first. The feed combines
// endpoint changes with provisioning, scaling, readiness, rollout, and other
// events from deployments under the endpoint.
func (r *BetaEndpointService) ListEventsAutoPaging(ctx context.Context, id string, params BetaEndpointListEventsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[BetaEndpointListEventsResponse] {
	return pagination.NewCursorPaginationAutoPager(r.ListEvents(ctx, id, params, opts...))
}

// Lists endpoints shared with every project in the specified organization.
// Project-private and public endpoints are not included.
func (r *BetaEndpointService) ListOrgScoped(ctx context.Context, organizationID string, query BetaEndpointListOrgScopedParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Endpoint], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/endpoints", organizationID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists endpoints shared with every project in the specified organization.
// Project-private and public endpoints are not included.
func (r *BetaEndpointService) ListOrgScopedAutoPaging(ctx context.Context, organizationID string, query BetaEndpointListOrgScopedParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Endpoint] {
	return pagination.NewCursorPaginationAutoPager(r.ListOrgScoped(ctx, organizationID, query, opts...))
}

// Deployment participating in an A/B experiment.
type AbMember struct {
	// Deployment under the parent endpoint.
	DeploymentID string `json:"deploymentId" api:"required"`
	// Integer traffic percent in [1, 100]. Percentages across all members must sum
	// to 100.
	Percent int64 `json:"percent" api:"required"`
	// Role of this deployment within the A/B experiment.
	//
	// Any of "AB_EXPERIMENT_MEMBER_ROLE_CONTROL", "AB_EXPERIMENT_MEMBER_ROLE_VARIANT".
	Role AbMemberRole `json:"role" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeploymentID respjson.Field
		Percent      respjson.Field
		Role         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AbMember) RawJSON() string { return r.JSON.raw }
func (r *AbMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AbMember to a AbMemberParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AbMemberParam.Overrides()
func (r AbMember) ToParam() AbMemberParam {
	return param.Override[AbMemberParam](json.RawMessage(r.RawJSON()))
}

// Role of this deployment within the A/B experiment.
type AbMemberRole string

const (
	AbMemberRoleAbExperimentMemberRoleControl AbMemberRole = "AB_EXPERIMENT_MEMBER_ROLE_CONTROL"
	AbMemberRoleAbExperimentMemberRoleVariant AbMemberRole = "AB_EXPERIMENT_MEMBER_ROLE_VARIANT"
)

// Deployment participating in an A/B experiment.
//
// The properties DeploymentID, Percent, Role are required.
type AbMemberParam struct {
	// Deployment under the parent endpoint.
	DeploymentID string `json:"deploymentId" api:"required"`
	// Integer traffic percent in [1, 100]. Percentages across all members must sum
	// to 100.
	Percent int64 `json:"percent" api:"required"`
	// Role of this deployment within the A/B experiment.
	//
	// Any of "AB_EXPERIMENT_MEMBER_ROLE_CONTROL", "AB_EXPERIMENT_MEMBER_ROLE_VARIANT".
	Role AbMemberRole `json:"role,omitzero" api:"required"`
	paramObj
}

func (r AbMemberParam) MarshalJSON() (data []byte, err error) {
	type shadow AbMemberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AbMemberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Autoscaling configuration for a deployment.
type DeploymentAutoscaling struct {
	// Maximum number of replicas. Defaults to `minReplicas`; omitting it on update
	// preserves the current value.
	MaxReplicas int64 `json:"maxReplicas"`
	// Minimum number of replicas. Omit on update to preserve the current value. Set
	// both `minReplicas` and `maxReplicas` to `0` to stop the deployment.
	MinReplicas int64 `json:"minReplicas"`
	// Rate limits applied after stabilization and before replica bounds.
	ScaleDown ScalingRules `json:"scaleDown"`
	// Time a lower replica recommendation must remain stable before scaling down.
	// Defaults to `5m`.
	ScaleDownWindow string `json:"scaleDownWindow"`
	// Idle period after which the deployment automatically stops and releases its
	// replicas.
	ScaleToZeroWindow string `json:"scaleToZeroWindow"`
	// Rate limits applied after stabilization and before replica bounds.
	ScaleUp ScalingRules `json:"scaleUp"`
	// Stabilization window before scaling up.
	ScaleUpWindow string `json:"scaleUpWindow"`
	// Metrics and targets that drive replica recommendations. When omitted, the
	// platform uses concurrent in-flight requests per replica.
	ScalingMetrics []ScalingMetric `json:"scalingMetrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxReplicas       respjson.Field
		MinReplicas       respjson.Field
		ScaleDown         respjson.Field
		ScaleDownWindow   respjson.Field
		ScaleToZeroWindow respjson.Field
		ScaleUp           respjson.Field
		ScaleUpWindow     respjson.Field
		ScalingMetrics    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentAutoscaling) RawJSON() string { return r.JSON.raw }
func (r *DeploymentAutoscaling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeploymentAutoscaling to a DeploymentAutoscalingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeploymentAutoscalingParam.Overrides()
func (r DeploymentAutoscaling) ToParam() DeploymentAutoscalingParam {
	return param.Override[DeploymentAutoscalingParam](json.RawMessage(r.RawJSON()))
}

// Autoscaling configuration for a deployment.
type DeploymentAutoscalingParam struct {
	// Maximum number of replicas. Defaults to `minReplicas`; omitting it on update
	// preserves the current value.
	MaxReplicas param.Opt[int64] `json:"maxReplicas,omitzero"`
	// Minimum number of replicas. Omit on update to preserve the current value. Set
	// both `minReplicas` and `maxReplicas` to `0` to stop the deployment.
	MinReplicas param.Opt[int64] `json:"minReplicas,omitzero"`
	// Time a lower replica recommendation must remain stable before scaling down.
	// Defaults to `5m`.
	ScaleDownWindow param.Opt[string] `json:"scaleDownWindow,omitzero"`
	// Idle period after which the deployment automatically stops and releases its
	// replicas.
	ScaleToZeroWindow param.Opt[string] `json:"scaleToZeroWindow,omitzero"`
	// Stabilization window before scaling up.
	ScaleUpWindow param.Opt[string] `json:"scaleUpWindow,omitzero"`
	// Rate limits applied after stabilization and before replica bounds.
	ScaleDown ScalingRulesParam `json:"scaleDown,omitzero"`
	// Rate limits applied after stabilization and before replica bounds.
	ScaleUp ScalingRulesParam `json:"scaleUp,omitzero"`
	// Metrics and targets that drive replica recommendations. When omitted, the
	// platform uses concurrent in-flight requests per replica.
	ScalingMetrics []ScalingMetricParam `json:"scalingMetrics,omitzero"`
	paramObj
}

func (r DeploymentAutoscalingParam) MarshalJSON() (data []byte, err error) {
	type shadow DeploymentAutoscalingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeploymentAutoscalingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operational metrics for one deployment under an endpoint.
type DeploymentMetrics struct {
	// ID of the deployment summarized by these metrics.
	DeploymentID string `json:"deploymentId"`
	// ID of the deployment's parent endpoint.
	EndpointID string `json:"endpointId"`
	// Error rate and counts by error type.
	ErrorMetrics ErrorMetrics `json:"errorMetrics"`
	// Time-to-first-token, end-to-end, and inter-token latency percentiles.
	LatencyMetrics LatencyMetrics `json:"latencyMetrics"`
	// Request counts and rates.
	RequestMetrics RequestMetrics `json:"requestMetrics"`
	// Average CPU, GPU, memory, and network utilization.
	ResourceUtilization ResourceUtilization `json:"resourceUtilization"`
	// Token, request, and batching throughput.
	ThroughputMetrics ThroughputMetrics `json:"throughputMetrics"`
	// Closed-open time range covered by the metrics.
	TimeRange MetricsTimeRange `json:"timeRange"`
	// Input and output token totals and averages.
	TokenMetrics TokenMetrics `json:"tokenMetrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeploymentID        respjson.Field
		EndpointID          respjson.Field
		ErrorMetrics        respjson.Field
		LatencyMetrics      respjson.Field
		RequestMetrics      respjson.Field
		ResourceUtilization respjson.Field
		ThroughputMetrics   respjson.Field
		TimeRange           respjson.Field
		TokenMetrics        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentMetrics) RawJSON() string { return r.JSON.raw }
func (r *DeploymentMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline placement parameters expanded into scheduling rules by the server.
type DeploymentPlacementConfig struct {
	// Compliance regimes required by a deployment placement policy.
	CompliancePolicy DeploymentPlacementConfigCompliancePolicy `json:"compliancePolicy"`
	// How strictly the regions list is enforced.
	//
	// Any of "ENFORCEMENT_REQUIRED", "ENFORCEMENT_PREFERRED".
	Constraint DeploymentPlacementConfigConstraint `json:"constraint"`
	// Regions where the deployment is allowed to run. Multiple regions allow
	// best-effort replica spreading.
	Regions []string `json:"regions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompliancePolicy respjson.Field
		Constraint       respjson.Field
		Regions          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentPlacementConfig) RawJSON() string { return r.JSON.raw }
func (r *DeploymentPlacementConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeploymentPlacementConfig to a
// DeploymentPlacementConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeploymentPlacementConfigParam.Overrides()
func (r DeploymentPlacementConfig) ToParam() DeploymentPlacementConfigParam {
	return param.Override[DeploymentPlacementConfigParam](json.RawMessage(r.RawJSON()))
}

// Compliance regimes required by a deployment placement policy.
type DeploymentPlacementConfigCompliancePolicy struct {
	// Restrict placement to HIPAA-attested clusters.
	Hipaa bool `json:"hipaa"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hipaa       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentPlacementConfigCompliancePolicy) RawJSON() string { return r.JSON.raw }
func (r *DeploymentPlacementConfigCompliancePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How strictly the regions list is enforced.
type DeploymentPlacementConfigConstraint string

const (
	DeploymentPlacementConfigConstraintEnforcementRequired  DeploymentPlacementConfigConstraint = "ENFORCEMENT_REQUIRED"
	DeploymentPlacementConfigConstraintEnforcementPreferred DeploymentPlacementConfigConstraint = "ENFORCEMENT_PREFERRED"
)

// Inline placement parameters expanded into scheduling rules by the server.
type DeploymentPlacementConfigParam struct {
	// Compliance regimes required by a deployment placement policy.
	CompliancePolicy DeploymentPlacementConfigCompliancePolicyParam `json:"compliancePolicy,omitzero"`
	// How strictly the regions list is enforced.
	//
	// Any of "ENFORCEMENT_REQUIRED", "ENFORCEMENT_PREFERRED".
	Constraint DeploymentPlacementConfigConstraint `json:"constraint,omitzero"`
	// Regions where the deployment is allowed to run. Multiple regions allow
	// best-effort replica spreading.
	Regions []string `json:"regions,omitzero"`
	paramObj
}

func (r DeploymentPlacementConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow DeploymentPlacementConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeploymentPlacementConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compliance regimes required by a deployment placement policy.
type DeploymentPlacementConfigCompliancePolicyParam struct {
	// Restrict placement to HIPAA-attested clusters.
	Hipaa param.Opt[bool] `json:"hipaa,omitzero"`
	paramObj
}

func (r DeploymentPlacementConfigCompliancePolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow DeploymentPlacementConfigCompliancePolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeploymentPlacementConfigCompliancePolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of a deployment, derived at read time from internal state.
type DeploymentStatus struct {
	// Human-readable explanation of the current state.
	Message string `json:"message" api:"required"`
	// High-level lifecycle state.
	//
	// Any of "DEPLOYMENT_STATE_PROVISIONING", "DEPLOYMENT_STATE_READY",
	// "DEPLOYMENT_STATE_SCALING", "DEPLOYMENT_STATE_DEGRADED",
	// "DEPLOYMENT_STATE_FAILED", "DEPLOYMENT_STATE_STOPPED",
	// "DEPLOYMENT_STATE_STOPPING".
	State DeploymentStatusState `json:"state" api:"required"`
	// Total replicas actively serving traffic across all clusters.
	ReadyReplicas int64 `json:"readyReplicas"`
	// Replicas the scheduler has placed on clusters.
	ScheduledReplicas int64 `json:"scheduledReplicas"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message           respjson.Field
		State             respjson.Field
		ReadyReplicas     respjson.Field
		ScheduledReplicas respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentStatus) RawJSON() string { return r.JSON.raw }
func (r *DeploymentStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// High-level lifecycle state.
type DeploymentStatusState string

const (
	DeploymentStatusStateDeploymentStateProvisioning DeploymentStatusState = "DEPLOYMENT_STATE_PROVISIONING"
	DeploymentStatusStateDeploymentStateReady        DeploymentStatusState = "DEPLOYMENT_STATE_READY"
	DeploymentStatusStateDeploymentStateScaling      DeploymentStatusState = "DEPLOYMENT_STATE_SCALING"
	DeploymentStatusStateDeploymentStateDegraded     DeploymentStatusState = "DEPLOYMENT_STATE_DEGRADED"
	DeploymentStatusStateDeploymentStateFailed       DeploymentStatusState = "DEPLOYMENT_STATE_FAILED"
	DeploymentStatusStateDeploymentStateStopped      DeploymentStatusState = "DEPLOYMENT_STATE_STOPPED"
	DeploymentStatusStateDeploymentStateStopping     DeploymentStatusState = "DEPLOYMENT_STATE_STOPPING"
)

// Stable inference entry point that groups deployments and routes requests among
// them.
type Endpoint struct {
	// Unique endpoint identifier.
	ID string `json:"id" api:"required"`
	// Timestamp when the endpoint was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Lightweight summaries of deployments under this endpoint. Retrieve a deployment
	// through the endpoint's deployment API for full details.
	Deployments []EndpointDeploymentSummary `json:"deployments" api:"required"`
	// Serving class of the endpoint. Reserved endpoints use reserved capacity.
	//
	// Any of "ENDPOINT_TYPE_DEDICATED", "ENDPOINT_TYPE_SERVERLESS",
	// "ENDPOINT_TYPE_RESERVED".
	EndpointType EndpointEndpointType `json:"endpointType" api:"required"`
	// Opaque version tag for optimistic concurrency control. Supply on update/delete
	// to ensure consistent read-modify-write. If not set, the write overwrites based
	// on current state.
	Etag string `json:"etag" api:"required"`
	// Project-qualified endpoint name in the form `<project_slug>/<endpoint_name>`.
	// Pass this value as `model` in inference requests. Create and update requests may
	// use either a bare endpoint name or the qualified form; a supplied project slug
	// must match the project in the request path.
	Name string `json:"name" api:"required"`
	// ID of the project that owns the endpoint.
	ProjectID string `json:"projectId" api:"required"`
	// Deployments eligible for live traffic and their capacity weights. An empty list
	// leaves the endpoint unrouted.
	TrafficSplit []EndpointTrafficSplitEntry `json:"trafficSplit" api:"required"`
	// Output only. Timestamp when the endpoint was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Who can discover the endpoint. `VISIBILITY_PRIVATE` restricts it to the project;
	// `VISIBILITY_INTERNAL` shares it with the organization.
	//
	// Any of "VISIBILITY_PRIVATE", "VISIBILITY_INTERNAL".
	Visibility EndpointVisibility `json:"visibility" api:"required"`
	// ID of the currently active rollout in an in-flight state, including paused.
	ActiveRolloutID string `json:"activeRolloutId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Deployments     respjson.Field
		EndpointType    respjson.Field
		Etag            respjson.Field
		Name            respjson.Field
		ProjectID       respjson.Field
		TrafficSplit    respjson.Field
		UpdatedAt       respjson.Field
		Visibility      respjson.Field
		ActiveRolloutID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Endpoint) RawJSON() string { return r.JSON.raw }
func (r *Endpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Serving class of the endpoint. Reserved endpoints use reserved capacity.
type EndpointEndpointType string

const (
	EndpointEndpointTypeEndpointTypeDedicated  EndpointEndpointType = "ENDPOINT_TYPE_DEDICATED"
	EndpointEndpointTypeEndpointTypeServerless EndpointEndpointType = "ENDPOINT_TYPE_SERVERLESS"
	EndpointEndpointTypeEndpointTypeReserved   EndpointEndpointType = "ENDPOINT_TYPE_RESERVED"
)

// Who can discover the endpoint. `VISIBILITY_PRIVATE` restricts it to the project;
// `VISIBILITY_INTERNAL` shares it with the organization.
type EndpointVisibility string

const (
	EndpointVisibilityVisibilityPrivate  EndpointVisibility = "VISIBILITY_PRIVATE"
	EndpointVisibilityVisibilityInternal EndpointVisibility = "VISIBILITY_INTERNAL"
)

// Serving workload that binds a model and immutable config to an endpoint and
// manages its replicas.
type EndpointDeployment struct {
	// Unique deployment identifier.
	ID string `json:"id" api:"required"`
	// Replica bounds, timing windows, and metrics that control horizontal scaling.
	Autoscaling EndpointDeploymentAutoscaling `json:"autoscaling" api:"required"`
	// Immutable config revision in the form
	// `projects/{projectId}/configs/{configRevisionId}`.
	Config string `json:"config" api:"required"`
	// Deprecated. Use `config`. Config revision identifier, populated during
	// migration.
	ConfigID string `json:"configId" api:"required"`
	// Timestamp when the deployment was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// ID of the endpoint that contains the deployment.
	EndpointID string `json:"endpointId" api:"required"`
	// Opaque version tag for optimistic concurrency control. Supply on update/delete
	// to ensure consistent read-modify-write. If not set, the write overwrites based
	// on current state.
	Etag string `json:"etag" api:"required"`
	// Hardware selected by the deployment config, including GPU type and count.
	Hardware string `json:"hardware" api:"required"`
	// Pinned model resource in the form
	// `projects/{projectId}/models/{modelId}/revisions/{revisionId}`.
	Model string `json:"model" api:"required"`
	// Deprecated. Use `model`. Model identifier being served, populated during
	// migration.
	ModelID string `json:"modelId" api:"required"`
	// Deprecated. Use `model` with a /revisions/{revisionId} segment. Pin to a
	// specific model revision.
	ModelRevisionID string `json:"modelRevisionId" api:"required"`
	// Project- and endpoint-qualified deployment name in the form
	// `<project_slug>/<endpoint_name>/<deployment_name>`. Pass it as `model` in an
	// inference request to target this deployment directly instead of using the
	// endpoint's traffic split.
	Name string `json:"name" api:"required"`
	// ID of the project that owns the deployment.
	ProjectID string `json:"projectId" api:"required"`
	// Current status of a deployment, derived at read time from internal state.
	Status DeploymentStatus `json:"status" api:"required"`
	// Whether the deployment serves client-visible responses or only mirrored shadow
	// traffic.
	//
	// Any of "TRAFFIC_MODE_LIVE", "TRAFFIC_MODE_SHADOW".
	TrafficMode EndpointDeploymentTrafficMode `json:"trafficMode" api:"required"`
	// Timestamp when the deployment was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Number of replicas the autoscaler currently wants across all regions.
	DesiredReplicas int64 `json:"desiredReplicas"`
	// Estimated fraction in [0, 1] of endpoint traffic that reaches this deployment
	// under the current routing configuration. Absent or unrouted deployments are 0.
	EstimatedEffectiveTrafficShare float64 `json:"estimatedEffectiveTrafficShare"`
	// Minutes without an inference request before the deployment stops automatically.
	// Omitted or 0 means automatic stopping is disabled.
	InactiveTimeout int64 `json:"inactiveTimeout"`
	// Placement controls where a deployment is scheduled.
	Placement EndpointDeploymentPlacementUnion `json:"placement"`
	// Runtime information derived from the deployment's configuration.
	RuntimeInfo RuntimeInfo `json:"runtimeInfo"`
	// Pinned draft-model resource used for speculative decoding, in the same form as
	// `model`. Omitted when speculative decoding is disabled.
	Speculator string `json:"speculator"`
	// Deprecated. Use `speculator`. Speculative decoding model identifier derived from
	// the deployment config.
	SpeculatorID string `json:"speculatorId"`
	// Deprecated. Use `speculator`. ID of the speculative decoding draft-model
	// revision pinned at creation time.
	SpeculatorRevisionID string `json:"speculatorRevisionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                             respjson.Field
		Autoscaling                    respjson.Field
		Config                         respjson.Field
		ConfigID                       respjson.Field
		CreatedAt                      respjson.Field
		EndpointID                     respjson.Field
		Etag                           respjson.Field
		Hardware                       respjson.Field
		Model                          respjson.Field
		ModelID                        respjson.Field
		ModelRevisionID                respjson.Field
		Name                           respjson.Field
		ProjectID                      respjson.Field
		Status                         respjson.Field
		TrafficMode                    respjson.Field
		UpdatedAt                      respjson.Field
		DesiredReplicas                respjson.Field
		EstimatedEffectiveTrafficShare respjson.Field
		InactiveTimeout                respjson.Field
		Placement                      respjson.Field
		RuntimeInfo                    respjson.Field
		Speculator                     respjson.Field
		SpeculatorID                   respjson.Field
		SpeculatorRevisionID           respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointDeployment) RawJSON() string { return r.JSON.raw }
func (r *EndpointDeployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Replica bounds, timing windows, and metrics that control horizontal scaling.
type EndpointDeploymentAutoscaling struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DeploymentAutoscaling
}

// Returns the unmodified JSON received from the API
func (r EndpointDeploymentAutoscaling) RawJSON() string { return r.JSON.raw }
func (r *EndpointDeploymentAutoscaling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the deployment serves client-visible responses or only mirrored shadow
// traffic.
type EndpointDeploymentTrafficMode string

const (
	EndpointDeploymentTrafficModeTrafficModeLive   EndpointDeploymentTrafficMode = "TRAFFIC_MODE_LIVE"
	EndpointDeploymentTrafficModeTrafficModeShadow EndpointDeploymentTrafficMode = "TRAFFIC_MODE_SHADOW"
)

// EndpointDeploymentPlacementUnion contains all possible properties and values
// from [EndpointDeploymentPlacementInline], [EndpointDeploymentPlacementProfile].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: Inline]
type EndpointDeploymentPlacementUnion struct {
	// This field is from variant [EndpointDeploymentPlacementInline].
	Inline DeploymentPlacementConfig `json:"inline"`
	// This field is from variant [EndpointDeploymentPlacementProfile].
	Profile string `json:"profile"`
	JSON    struct {
		Inline  respjson.Field
		Profile respjson.Field
		raw     string
	} `json:"-"`
}

func (u EndpointDeploymentPlacementUnion) AsInline() (v EndpointDeploymentPlacementInline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EndpointDeploymentPlacementUnion) AsProfile() (v EndpointDeploymentPlacementProfile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EndpointDeploymentPlacementUnion) RawJSON() string { return u.JSON.raw }

func (r *EndpointDeploymentPlacementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointDeploymentPlacementInline struct {
	// Inline placement parameters expanded into scheduling rules by the server.
	Inline DeploymentPlacementConfig `json:"inline" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inline      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointDeploymentPlacementInline) RawJSON() string { return r.JSON.raw }
func (r *EndpointDeploymentPlacementInline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointDeploymentPlacementProfile struct {
	// UID of a saved placement profile.
	Profile string `json:"profile" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Profile     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointDeploymentPlacementProfile) RawJSON() string { return r.JSON.raw }
func (r *EndpointDeploymentPlacementProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compact deployment status embedded in an endpoint response.
type EndpointDeploymentSummary struct {
	// Deployment identifier.
	ID string `json:"id" api:"required"`
	// Autoscaling configuration for the deployment.
	Autoscaling EndpointDeploymentSummaryAutoscaling `json:"autoscaling" api:"required"`
	// Timestamp when the deployment was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Estimated fraction from 0 to 1 of endpoint traffic currently routed to this
	// deployment.
	EstimatedEffectiveTrafficShare float64 `json:"estimatedEffectiveTrafficShare" api:"required"`
	// Resource name of the served model in the form
	// `projects/{projectId}/models/{modelId}/revisions/{revisionId}`. For public
	// models, the model's owning project may differ from the deployment's project.
	Model string `json:"model" api:"required"`
	// Deprecated. Use `model`. Model identifier being served.
	ModelID string `json:"modelId" api:"required"`
	// Inference-addressable name in the fully-qualified form
	// "<project_slug>/<endpoint_name>/<deployment_name>". Pass it as the "model" field
	// when calling the inference API to pin to this deployment.
	Name string `json:"name" api:"required"`
	// Current state of the deployment.
	//
	// Any of "DEPLOYMENT_STATE_PROVISIONING", "DEPLOYMENT_STATE_READY",
	// "DEPLOYMENT_STATE_SCALING", "DEPLOYMENT_STATE_DEGRADED",
	// "DEPLOYMENT_STATE_FAILED", "DEPLOYMENT_STATE_STOPPED",
	// "DEPLOYMENT_STATE_STOPPING".
	State EndpointDeploymentSummaryState `json:"state" api:"required"`
	// Whether the deployment serves client-visible responses or only mirrored shadow
	// traffic.
	//
	// Any of "TRAFFIC_MODE_LIVE", "TRAFFIC_MODE_SHADOW".
	TrafficMode EndpointDeploymentSummaryTrafficMode `json:"trafficMode" api:"required"`
	// Number of replicas the autoscaler currently wants across all regions.
	DesiredReplicas int64 `json:"desiredReplicas"`
	// Hardware configuration selected by the deployment's config, such as its GPU type
	// and count.
	Hardware string `json:"hardware"`
	// Number of replicas currently ready to serve requests across all regions.
	ReadyReplicas int64 `json:"readyReplicas"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                             respjson.Field
		Autoscaling                    respjson.Field
		CreatedAt                      respjson.Field
		EstimatedEffectiveTrafficShare respjson.Field
		Model                          respjson.Field
		ModelID                        respjson.Field
		Name                           respjson.Field
		State                          respjson.Field
		TrafficMode                    respjson.Field
		DesiredReplicas                respjson.Field
		Hardware                       respjson.Field
		ReadyReplicas                  respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointDeploymentSummary) RawJSON() string { return r.JSON.raw }
func (r *EndpointDeploymentSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Autoscaling configuration for the deployment.
type EndpointDeploymentSummaryAutoscaling struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DeploymentAutoscaling
}

// Returns the unmodified JSON received from the API
func (r EndpointDeploymentSummaryAutoscaling) RawJSON() string { return r.JSON.raw }
func (r *EndpointDeploymentSummaryAutoscaling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of the deployment.
type EndpointDeploymentSummaryState string

const (
	EndpointDeploymentSummaryStateDeploymentStateProvisioning EndpointDeploymentSummaryState = "DEPLOYMENT_STATE_PROVISIONING"
	EndpointDeploymentSummaryStateDeploymentStateReady        EndpointDeploymentSummaryState = "DEPLOYMENT_STATE_READY"
	EndpointDeploymentSummaryStateDeploymentStateScaling      EndpointDeploymentSummaryState = "DEPLOYMENT_STATE_SCALING"
	EndpointDeploymentSummaryStateDeploymentStateDegraded     EndpointDeploymentSummaryState = "DEPLOYMENT_STATE_DEGRADED"
	EndpointDeploymentSummaryStateDeploymentStateFailed       EndpointDeploymentSummaryState = "DEPLOYMENT_STATE_FAILED"
	EndpointDeploymentSummaryStateDeploymentStateStopped      EndpointDeploymentSummaryState = "DEPLOYMENT_STATE_STOPPED"
	EndpointDeploymentSummaryStateDeploymentStateStopping     EndpointDeploymentSummaryState = "DEPLOYMENT_STATE_STOPPING"
)

// Whether the deployment serves client-visible responses or only mirrored shadow
// traffic.
type EndpointDeploymentSummaryTrafficMode string

const (
	EndpointDeploymentSummaryTrafficModeTrafficModeLive   EndpointDeploymentSummaryTrafficMode = "TRAFFIC_MODE_LIVE"
	EndpointDeploymentSummaryTrafficModeTrafficModeShadow EndpointDeploymentSummaryTrafficMode = "TRAFFIC_MODE_SHADOW"
)

// Capacity weight assigned to one deployment in an endpoint's live traffic split.
type EndpointTrafficSplitEntry struct {
	// ID of a deployment under the endpoint that can receive live traffic.
	DeploymentID string `json:"deploymentId" api:"required"`
	// Non-negative, finite weight applied to each ready replica. A deployment's
	// effective routing capacity is `weight * readyReplicas`, and requests are
	// distributed in proportion to that capacity. Set to `0` to remove the deployment
	// from the live traffic split.
	Weight float64 `json:"weight" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeploymentID respjson.Field
		Weight       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointTrafficSplitEntry) RawJSON() string { return r.JSON.raw }
func (r *EndpointTrafficSplitEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EndpointTrafficSplitEntry to a
// EndpointTrafficSplitEntryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EndpointTrafficSplitEntryParam.Overrides()
func (r EndpointTrafficSplitEntry) ToParam() EndpointTrafficSplitEntryParam {
	return param.Override[EndpointTrafficSplitEntryParam](json.RawMessage(r.RawJSON()))
}

// Capacity weight assigned to one deployment in an endpoint's live traffic split.
//
// The properties DeploymentID, Weight are required.
type EndpointTrafficSplitEntryParam struct {
	// ID of a deployment under the endpoint that can receive live traffic.
	DeploymentID string `json:"deploymentId" api:"required"`
	// Non-negative, finite weight applied to each ready replica. A deployment's
	// effective routing capacity is `weight * readyReplicas`, and requests are
	// distributed in proportion to that capacity. Set to `0` to remove the deployment
	// from the live traffic split.
	Weight float64 `json:"weight" api:"required"`
	paramObj
}

func (r EndpointTrafficSplitEntryParam) MarshalJSON() (data []byte, err error) {
	type shadow EndpointTrafficSplitEntryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointTrafficSplitEntryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error rate and aggregate counts by error type. Individual error samples are not
// included.
type ErrorMetrics struct {
	// Percentage in [0, 100].
	ErrorRate float64 `json:"errorRate"`
	// Counts of errors keyed by error type (e.g. HTTP status code or error kind).
	ErrorsByType map[string]string `json:"errorsByType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorRate    respjson.Field
		ErrorsByType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorMetrics) RawJSON() string { return r.JSON.raw }
func (r *ErrorMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time-to-first-token, end-to-end, and inter-token latency percentiles in
// milliseconds.
type LatencyMetrics struct {
	// 50th-percentile inter-token latency, in milliseconds.
	ItlP50Ms float64 `json:"itlP50Ms"`
	// 90th-percentile inter-token latency, in milliseconds.
	ItlP90Ms float64 `json:"itlP90Ms"`
	// 99th-percentile inter-token latency, in milliseconds.
	ItlP99Ms float64 `json:"itlP99Ms"`
	// 50th-percentile end-to-end request latency, in milliseconds.
	LatencyP50Ms float64 `json:"latencyP50Ms"`
	// 90th-percentile end-to-end request latency, in milliseconds.
	LatencyP90Ms float64 `json:"latencyP90Ms"`
	// 99th-percentile end-to-end request latency, in milliseconds.
	LatencyP99Ms float64 `json:"latencyP99Ms"`
	// 50th-percentile time to first token, in milliseconds.
	TtftP50Ms float64 `json:"ttftP50Ms"`
	// 90th-percentile time to first token, in milliseconds.
	TtftP90Ms float64 `json:"ttftP90Ms"`
	// 99th-percentile time to first token, in milliseconds.
	TtftP99Ms float64 `json:"ttftP99Ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItlP50Ms     respjson.Field
		ItlP90Ms     respjson.Field
		ItlP99Ms     respjson.Field
		LatencyP50Ms respjson.Field
		LatencyP90Ms respjson.Field
		LatencyP99Ms respjson.Field
		TtftP50Ms    respjson.Field
		TtftP90Ms    respjson.Field
		TtftP99Ms    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LatencyMetrics) RawJSON() string { return r.JSON.raw }
func (r *LatencyMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Closed-open time range used by metrics and analytics responses.
type MetricsTimeRange struct {
	// Exclusive end of the time range.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// Inclusive start of the time range.
	StartTime time.Time `json:"startTime" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsTimeRange) RawJSON() string { return r.JSON.raw }
func (r *MetricsTimeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request counts, rate, and status-code distribution over a time range.
type RequestMetrics struct {
	// Requests that failed during the time range.
	FailedRequests string `json:"failedRequests"`
	// Request counts keyed by HTTP status code.
	RequestsByStatusCode map[string]string `json:"requestsByStatusCode"`
	// Average requests per second over the time range.
	RequestsPerSecond float64 `json:"requestsPerSecond"`
	// Requests completed successfully during the time range.
	SuccessfulRequests string `json:"successfulRequests"`
	// Total requests received during the time range.
	TotalRequests string `json:"totalRequests"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FailedRequests       respjson.Field
		RequestsByStatusCode respjson.Field
		RequestsPerSecond    respjson.Field
		SuccessfulRequests   respjson.Field
		TotalRequests        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequestMetrics) RawJSON() string { return r.JSON.raw }
func (r *RequestMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Average compute, memory, and network utilization for replicas over a time range.
type ResourceUtilization struct {
	// Average CPU utilization across replicas, as a percentage.
	CPUUtilization float64 `json:"cpuUtilization"`
	// Average GPU memory utilization across replicas, as a percentage.
	GPUMemoryUtilization float64 `json:"gpuMemoryUtilization"`
	// Average GPU compute utilization across replicas, as a percentage.
	GPUUtilization float64 `json:"gpuUtilization"`
	// Average system memory utilization across replicas, as a percentage.
	MemoryUtilization float64 `json:"memoryUtilization"`
	// Average network throughput across replicas, in megabits per second.
	NetworkBandwidthMbps float64 `json:"networkBandwidthMbps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPUUtilization       respjson.Field
		GPUMemoryUtilization respjson.Field
		GPUUtilization       respjson.Field
		MemoryUtilization    respjson.Field
		NetworkBandwidthMbps respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResourceUtilization) RawJSON() string { return r.JSON.raw }
func (r *ResourceUtilization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime information derived from the deployment's configuration.
type RuntimeInfo struct {
	// Serving engine, such as `vllm`, `trtllm`, or `sglang`.
	EngineType string `json:"engineType"`
	// Version of the serving engine.
	EngineVersion string `json:"engineVersion"`
	// Whether the runtime accepts tool and function-calling requests.
	FunctionCallingSupported bool `json:"functionCallingSupported"`
	// Whether the runtime can constrain generation to a structured output schema.
	StructuredOutputSupported bool `json:"structuredOutputSupported"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EngineType                respjson.Field
		EngineVersion             respjson.Field
		FunctionCallingSupported  respjson.Field
		StructuredOutputSupported respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuntimeInfo) RawJSON() string { return r.JSON.raw }
func (r *RuntimeInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metric and target used by the autoscaler to recommend a replica count.
type ScalingMetric struct {
	// Autoscaling metric name from the server allowlist.
	//
	// Any of "active_sessions", "cache_hit_rate", "decoding_speed", "e2e_latency",
	// "gpu_utilization", "inflight_requests", "throughput_per_replica",
	// "token_utilization", "ttft".
	Name ScalingMetricName `json:"name" api:"required"`
	// Target interpreted according to `type`. Utilization uses a percentage from 0 to
	// 100, value uses an absolute measurement, and average value uses a per-replica
	// measurement.
	Target float64 `json:"target" api:"required"`
	// Whether `target` is an absolute value, a utilization percentage, or a
	// per-replica average.
	//
	// Any of "METRIC_TARGET_TYPE_VALUE", "METRIC_TARGET_TYPE_UTILIZATION",
	// "METRIC_TARGET_TYPE_AVERAGE_VALUE".
	Type ScalingMetricType `json:"type" api:"required"`
	// Percentile to evaluate for latency-based metrics: `p50`, `p90`, `p95`, or `p99`.
	Percentile string `json:"percentile"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Target      respjson.Field
		Type        respjson.Field
		Percentile  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScalingMetric) RawJSON() string { return r.JSON.raw }
func (r *ScalingMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ScalingMetric to a ScalingMetricParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ScalingMetricParam.Overrides()
func (r ScalingMetric) ToParam() ScalingMetricParam {
	return param.Override[ScalingMetricParam](json.RawMessage(r.RawJSON()))
}

// Autoscaling metric name from the server allowlist.
type ScalingMetricName string

const (
	ScalingMetricNameActiveSessions       ScalingMetricName = "active_sessions"
	ScalingMetricNameCacheHitRate         ScalingMetricName = "cache_hit_rate"
	ScalingMetricNameDecodingSpeed        ScalingMetricName = "decoding_speed"
	ScalingMetricNameE2ELatency           ScalingMetricName = "e2e_latency"
	ScalingMetricNameGPUUtilization       ScalingMetricName = "gpu_utilization"
	ScalingMetricNameInflightRequests     ScalingMetricName = "inflight_requests"
	ScalingMetricNameThroughputPerReplica ScalingMetricName = "throughput_per_replica"
	ScalingMetricNameTokenUtilization     ScalingMetricName = "token_utilization"
	ScalingMetricNameTtft                 ScalingMetricName = "ttft"
)

// Whether `target` is an absolute value, a utilization percentage, or a
// per-replica average.
type ScalingMetricType string

const (
	ScalingMetricTypeMetricTargetTypeValue        ScalingMetricType = "METRIC_TARGET_TYPE_VALUE"
	ScalingMetricTypeMetricTargetTypeUtilization  ScalingMetricType = "METRIC_TARGET_TYPE_UTILIZATION"
	ScalingMetricTypeMetricTargetTypeAverageValue ScalingMetricType = "METRIC_TARGET_TYPE_AVERAGE_VALUE"
)

// Metric and target used by the autoscaler to recommend a replica count.
//
// The properties Name, Target, Type are required.
type ScalingMetricParam struct {
	// Autoscaling metric name from the server allowlist.
	//
	// Any of "active_sessions", "cache_hit_rate", "decoding_speed", "e2e_latency",
	// "gpu_utilization", "inflight_requests", "throughput_per_replica",
	// "token_utilization", "ttft".
	Name ScalingMetricName `json:"name,omitzero" api:"required"`
	// Target interpreted according to `type`. Utilization uses a percentage from 0 to
	// 100, value uses an absolute measurement, and average value uses a per-replica
	// measurement.
	Target float64 `json:"target" api:"required"`
	// Whether `target` is an absolute value, a utilization percentage, or a
	// per-replica average.
	//
	// Any of "METRIC_TARGET_TYPE_VALUE", "METRIC_TARGET_TYPE_UTILIZATION",
	// "METRIC_TARGET_TYPE_AVERAGE_VALUE".
	Type ScalingMetricType `json:"type,omitzero" api:"required"`
	// Percentile to evaluate for latency-based metrics: `p50`, `p90`, `p95`, or `p99`.
	Percentile param.Opt[string] `json:"percentile,omitzero"`
	paramObj
}

func (r ScalingMetricParam) MarshalJSON() (data []byte, err error) {
	type shadow ScalingMetricParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScalingMetricParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Replica rate-limit policy applied over a trailing window.
type ScalingPolicy struct {
	// Trailing rate-limit window in seconds, from 1 to 1800.
	PeriodSeconds int64 `json:"periodSeconds" api:"required"`
	// Whether `value` is a replica count or a percentage of the replica count at the
	// start of the trailing period. Scaling events within that period count against
	// the allowance; percentages are rounded to whole replicas.
	//
	// Any of "SCALING_POLICY_TYPE_PODS", "SCALING_POLICY_TYPE_PERCENT".
	Type ScalingPolicyType `json:"type" api:"required"`
	// Positive replica count or percentage used as the rate-limit amount.
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PeriodSeconds respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScalingPolicy) RawJSON() string { return r.JSON.raw }
func (r *ScalingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ScalingPolicy to a ScalingPolicyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ScalingPolicyParam.Overrides()
func (r ScalingPolicy) ToParam() ScalingPolicyParam {
	return param.Override[ScalingPolicyParam](json.RawMessage(r.RawJSON()))
}

// Whether `value` is a replica count or a percentage of the replica count at the
// start of the trailing period. Scaling events within that period count against
// the allowance; percentages are rounded to whole replicas.
type ScalingPolicyType string

const (
	ScalingPolicyTypeScalingPolicyTypePods    ScalingPolicyType = "SCALING_POLICY_TYPE_PODS"
	ScalingPolicyTypeScalingPolicyTypePercent ScalingPolicyType = "SCALING_POLICY_TYPE_PERCENT"
)

// Replica rate-limit policy applied over a trailing window.
//
// The properties PeriodSeconds, Type, Value are required.
type ScalingPolicyParam struct {
	// Trailing rate-limit window in seconds, from 1 to 1800.
	PeriodSeconds int64 `json:"periodSeconds" api:"required"`
	// Whether `value` is a replica count or a percentage of the replica count at the
	// start of the trailing period. Scaling events within that period count against
	// the allowance; percentages are rounded to whole replicas.
	//
	// Any of "SCALING_POLICY_TYPE_PODS", "SCALING_POLICY_TYPE_PERCENT".
	Type ScalingPolicyType `json:"type,omitzero" api:"required"`
	// Positive replica count or percentage used as the rate-limit amount.
	Value int64 `json:"value" api:"required"`
	paramObj
}

func (r ScalingPolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow ScalingPolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScalingPolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rate limits applied after stabilization and before replica bounds.
type ScalingRules struct {
	// Non-empty lists replace the existing policies. To clear policies, include
	// `autoscaling.scaleDown.policies` or `autoscaling.scaleUp.policies` in the update
	// mask and supply an empty scaling rules object or `policies: []`.
	Policies []ScalingPolicy `json:"policies"`
	// `SCALING_POLICY_SELECT_MIN` chooses the policy allowing the smallest replica
	// change; `SCALING_POLICY_SELECT_MAX` chooses the largest. These are caps, not
	// guaranteed changes. `SCALING_POLICY_SELECT_DISABLED` holds this direction steady
	// while replica bounds still apply. Omitted preserves the existing selector on
	// update. When no selector is configured, authored policies use MAX; with no
	// policies configured, the platform defaults apply. To reset the selector, include
	// `autoscaling.scaleDown.selectPolicy` or `autoscaling.scaleUp.selectPolicy` in
	// the update mask and omit `selectPolicy`. Clear both policies and `selectPolicy`
	// to restore inherited defaults.
	//
	// Any of "SCALING_POLICY_SELECT_MAX", "SCALING_POLICY_SELECT_MIN",
	// "SCALING_POLICY_SELECT_DISABLED".
	SelectPolicy ScalingRulesSelectPolicy `json:"selectPolicy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Policies     respjson.Field
		SelectPolicy respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScalingRules) RawJSON() string { return r.JSON.raw }
func (r *ScalingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ScalingRules to a ScalingRulesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ScalingRulesParam.Overrides()
func (r ScalingRules) ToParam() ScalingRulesParam {
	return param.Override[ScalingRulesParam](json.RawMessage(r.RawJSON()))
}

// `SCALING_POLICY_SELECT_MIN` chooses the policy allowing the smallest replica
// change; `SCALING_POLICY_SELECT_MAX` chooses the largest. These are caps, not
// guaranteed changes. `SCALING_POLICY_SELECT_DISABLED` holds this direction steady
// while replica bounds still apply. Omitted preserves the existing selector on
// update. When no selector is configured, authored policies use MAX; with no
// policies configured, the platform defaults apply. To reset the selector, include
// `autoscaling.scaleDown.selectPolicy` or `autoscaling.scaleUp.selectPolicy` in
// the update mask and omit `selectPolicy`. Clear both policies and `selectPolicy`
// to restore inherited defaults.
type ScalingRulesSelectPolicy string

const (
	ScalingRulesSelectPolicyScalingPolicySelectMax      ScalingRulesSelectPolicy = "SCALING_POLICY_SELECT_MAX"
	ScalingRulesSelectPolicyScalingPolicySelectMin      ScalingRulesSelectPolicy = "SCALING_POLICY_SELECT_MIN"
	ScalingRulesSelectPolicyScalingPolicySelectDisabled ScalingRulesSelectPolicy = "SCALING_POLICY_SELECT_DISABLED"
)

// Rate limits applied after stabilization and before replica bounds.
type ScalingRulesParam struct {
	// Non-empty lists replace the existing policies. To clear policies, include
	// `autoscaling.scaleDown.policies` or `autoscaling.scaleUp.policies` in the update
	// mask and supply an empty scaling rules object or `policies: []`.
	Policies []ScalingPolicyParam `json:"policies,omitzero"`
	// `SCALING_POLICY_SELECT_MIN` chooses the policy allowing the smallest replica
	// change; `SCALING_POLICY_SELECT_MAX` chooses the largest. These are caps, not
	// guaranteed changes. `SCALING_POLICY_SELECT_DISABLED` holds this direction steady
	// while replica bounds still apply. Omitted preserves the existing selector on
	// update. When no selector is configured, authored policies use MAX; with no
	// policies configured, the platform defaults apply. To reset the selector, include
	// `autoscaling.scaleDown.selectPolicy` or `autoscaling.scaleUp.selectPolicy` in
	// the update mask and omit `selectPolicy`. Clear both policies and `selectPolicy`
	// to restore inherited defaults.
	//
	// Any of "SCALING_POLICY_SELECT_MAX", "SCALING_POLICY_SELECT_MIN",
	// "SCALING_POLICY_SELECT_DISABLED".
	SelectPolicy ScalingRulesSelectPolicy `json:"selectPolicy,omitzero"`
	paramObj
}

func (r ScalingRulesParam) MarshalJSON() (data []byte, err error) {
	type shadow ScalingRulesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScalingRulesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Adaptive sticky-key sampling that throttles toward a target QPS.
//
// The properties Key, TargetQps are required.
type ShadowAdaptiveKeyBasedSamplingParam struct {
	// Required request-body field used as the sticky sampling key.
	Key string `json:"key" api:"required"`
	// Required per-gateway-replica target QPS for adaptive sampling.
	TargetQps float64 `json:"targetQps" api:"required"`
	// Optional sliding window for QPS observation. Defaults to 60s and must not be
	// negative.
	Window param.Opt[string] `json:"window,omitzero"`
	paramObj
}

func (r ShadowAdaptiveKeyBasedSamplingParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowAdaptiveKeyBasedSamplingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowAdaptiveKeyBasedSamplingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Adaptive sticky-key sampling returned by the API.
type ShadowAdaptiveKeyBasedSamplingResponse struct {
	// Request-body field used as the sticky sampling key.
	Key string `json:"key" api:"required"`
	// Per-gateway-replica target QPS.
	TargetQps float64 `json:"targetQps" api:"required"`
	// Sliding window for QPS observation when explicitly configured.
	Window string `json:"window"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		TargetQps   respjson.Field
		Window      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowAdaptiveKeyBasedSamplingResponse) RawJSON() string { return r.JSON.raw }
func (r *ShadowAdaptiveKeyBasedSamplingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Adaptive random sampling that throttles toward a target QPS.
//
// The property TargetQps is required.
type ShadowAdaptiveUniformSamplingParam struct {
	// Required per-gateway-replica target QPS for adaptive sampling.
	TargetQps float64 `json:"targetQps" api:"required"`
	// Optional sliding window for QPS observation. Defaults to 60s and must not be
	// negative.
	Window param.Opt[string] `json:"window,omitzero"`
	paramObj
}

func (r ShadowAdaptiveUniformSamplingParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowAdaptiveUniformSamplingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowAdaptiveUniformSamplingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Adaptive random sampling returned by the API.
type ShadowAdaptiveUniformSamplingResponse struct {
	// Per-gateway-replica target QPS.
	TargetQps float64 `json:"targetQps" api:"required"`
	// Sliding window for QPS observation when explicitly configured.
	Window string `json:"window"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TargetQps   respjson.Field
		Window      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowAdaptiveUniformSamplingResponse) RawJSON() string { return r.JSON.raw }
func (r *ShadowAdaptiveUniformSamplingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Endpoint-level source that samples endpoint traffic at the API gateway.
//
// The property Sampling is required.
type ShadowEndpointSourceParam struct {
	// Sampling strategy for endpoint-level shadow traffic. Exactly one strategy must
	// be set.
	Sampling ShadowEndpointSourceSamplingUnionParam `json:"sampling,omitzero" api:"required"`
	paramObj
}

func (r ShadowEndpointSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowEndpointSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowEndpointSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ShadowEndpointSourceSamplingUnionParam struct {
	OfUniform          *ShadowEndpointSourceSamplingUniformParam          `json:",omitzero,inline"`
	OfKeyBased         *ShadowEndpointSourceSamplingKeyBasedParam         `json:",omitzero,inline"`
	OfAdaptiveUniform  *ShadowEndpointSourceSamplingAdaptiveUniformParam  `json:",omitzero,inline"`
	OfAdaptiveKeyBased *ShadowEndpointSourceSamplingAdaptiveKeyBasedParam `json:",omitzero,inline"`
	paramUnion
}

func (u ShadowEndpointSourceSamplingUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUniform, u.OfKeyBased, u.OfAdaptiveUniform, u.OfAdaptiveKeyBased)
}
func (u *ShadowEndpointSourceSamplingUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ShadowEndpointSourceSamplingUnionParam) asAny() any {
	if !param.IsOmitted(u.OfUniform) {
		return u.OfUniform
	} else if !param.IsOmitted(u.OfKeyBased) {
		return u.OfKeyBased
	} else if !param.IsOmitted(u.OfAdaptiveUniform) {
		return u.OfAdaptiveUniform
	} else if !param.IsOmitted(u.OfAdaptiveKeyBased) {
		return u.OfAdaptiveKeyBased
	}
	return nil
}

// The property Uniform is required.
type ShadowEndpointSourceSamplingUniformParam struct {
	// Fixed-rate random sampling of endpoint requests.
	Uniform ShadowUniformSamplingParam `json:"uniform,omitzero" api:"required"`
	paramObj
}

func (r ShadowEndpointSourceSamplingUniformParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowEndpointSourceSamplingUniformParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowEndpointSourceSamplingUniformParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property KeyBased is required.
type ShadowEndpointSourceSamplingKeyBasedParam struct {
	// Fixed-rate sampling of distinct key values with sticky decisions.
	KeyBased ShadowKeyBasedSamplingParam `json:"keyBased,omitzero" api:"required"`
	paramObj
}

func (r ShadowEndpointSourceSamplingKeyBasedParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowEndpointSourceSamplingKeyBasedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowEndpointSourceSamplingKeyBasedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AdaptiveUniform is required.
type ShadowEndpointSourceSamplingAdaptiveUniformParam struct {
	// Adaptive random sampling that throttles toward a target QPS.
	AdaptiveUniform ShadowAdaptiveUniformSamplingParam `json:"adaptiveUniform,omitzero" api:"required"`
	paramObj
}

func (r ShadowEndpointSourceSamplingAdaptiveUniformParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowEndpointSourceSamplingAdaptiveUniformParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowEndpointSourceSamplingAdaptiveUniformParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AdaptiveKeyBased is required.
type ShadowEndpointSourceSamplingAdaptiveKeyBasedParam struct {
	// Adaptive sticky-key sampling that throttles toward a target QPS.
	AdaptiveKeyBased ShadowAdaptiveKeyBasedSamplingParam `json:"adaptiveKeyBased,omitzero" api:"required"`
	paramObj
}

func (r ShadowEndpointSourceSamplingAdaptiveKeyBasedParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowEndpointSourceSamplingAdaptiveKeyBasedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowEndpointSourceSamplingAdaptiveKeyBasedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Endpoint-level source returned for a shadow experiment.
type ShadowEndpointSourceResponse struct {
	// Sampling strategy returned for endpoint-level shadow traffic.
	Sampling ShadowEndpointSourceResponseSamplingUnion `json:"sampling" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sampling    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowEndpointSourceResponse) RawJSON() string { return r.JSON.raw }
func (r *ShadowEndpointSourceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShadowEndpointSourceResponseSamplingUnion contains all possible properties and
// values from [ShadowEndpointSourceResponseSamplingUniform],
// [ShadowEndpointSourceResponseSamplingKeyBased],
// [ShadowEndpointSourceResponseSamplingAdaptiveUniform],
// [ShadowEndpointSourceResponseSamplingAdaptiveKeyBased].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShadowEndpointSourceResponseSamplingUnion struct {
	// This field is from variant [ShadowEndpointSourceResponseSamplingUniform].
	Uniform ShadowUniformSamplingResponse `json:"uniform"`
	// This field is from variant [ShadowEndpointSourceResponseSamplingKeyBased].
	KeyBased ShadowKeyBasedSamplingResponse `json:"keyBased"`
	// This field is from variant
	// [ShadowEndpointSourceResponseSamplingAdaptiveUniform].
	AdaptiveUniform ShadowAdaptiveUniformSamplingResponse `json:"adaptiveUniform"`
	// This field is from variant
	// [ShadowEndpointSourceResponseSamplingAdaptiveKeyBased].
	AdaptiveKeyBased ShadowAdaptiveKeyBasedSamplingResponse `json:"adaptiveKeyBased"`
	JSON             struct {
		Uniform          respjson.Field
		KeyBased         respjson.Field
		AdaptiveUniform  respjson.Field
		AdaptiveKeyBased respjson.Field
		raw              string
	} `json:"-"`
}

func (u ShadowEndpointSourceResponseSamplingUnion) AsUniform() (v ShadowEndpointSourceResponseSamplingUniform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShadowEndpointSourceResponseSamplingUnion) AsKeyBased() (v ShadowEndpointSourceResponseSamplingKeyBased) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShadowEndpointSourceResponseSamplingUnion) AsAdaptiveUniform() (v ShadowEndpointSourceResponseSamplingAdaptiveUniform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShadowEndpointSourceResponseSamplingUnion) AsAdaptiveKeyBased() (v ShadowEndpointSourceResponseSamplingAdaptiveKeyBased) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShadowEndpointSourceResponseSamplingUnion) RawJSON() string { return u.JSON.raw }

func (r *ShadowEndpointSourceResponseSamplingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShadowEndpointSourceResponseSamplingUniform struct {
	// Fixed-rate random sampling returned by the API. A zero rate may be omitted by
	// JSON serialization.
	Uniform ShadowUniformSamplingResponse `json:"uniform" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uniform     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowEndpointSourceResponseSamplingUniform) RawJSON() string { return r.JSON.raw }
func (r *ShadowEndpointSourceResponseSamplingUniform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShadowEndpointSourceResponseSamplingKeyBased struct {
	// Fixed-rate sticky-key sampling returned by the API. A zero rate may be omitted
	// by JSON serialization.
	KeyBased ShadowKeyBasedSamplingResponse `json:"keyBased" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeyBased    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowEndpointSourceResponseSamplingKeyBased) RawJSON() string { return r.JSON.raw }
func (r *ShadowEndpointSourceResponseSamplingKeyBased) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShadowEndpointSourceResponseSamplingAdaptiveUniform struct {
	// Adaptive random sampling returned by the API.
	AdaptiveUniform ShadowAdaptiveUniformSamplingResponse `json:"adaptiveUniform" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdaptiveUniform respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowEndpointSourceResponseSamplingAdaptiveUniform) RawJSON() string { return r.JSON.raw }
func (r *ShadowEndpointSourceResponseSamplingAdaptiveUniform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShadowEndpointSourceResponseSamplingAdaptiveKeyBased struct {
	// Adaptive sticky-key sampling returned by the API.
	AdaptiveKeyBased ShadowAdaptiveKeyBasedSamplingResponse `json:"adaptiveKeyBased" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdaptiveKeyBased respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowEndpointSourceResponseSamplingAdaptiveKeyBased) RawJSON() string { return r.JSON.raw }
func (r *ShadowEndpointSourceResponseSamplingAdaptiveKeyBased) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fixed-rate sampling of distinct key values with sticky decisions.
//
// The properties Key, Rate are required.
type ShadowKeyBasedSamplingParam struct {
	// Required request-body field used as the sticky sampling key.
	Key string `json:"key" api:"required"`
	// Required fraction of distinct key values to sample, from 0.0 to 1.0.
	Rate float64 `json:"rate" api:"required"`
	paramObj
}

func (r ShadowKeyBasedSamplingParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowKeyBasedSamplingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowKeyBasedSamplingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fixed-rate sticky-key sampling returned by the API. A zero rate may be omitted
// by JSON serialization.
type ShadowKeyBasedSamplingResponse struct {
	// Request-body field used as the sticky sampling key.
	Key string `json:"key" api:"required"`
	// Fraction of distinct key values sampled, from 0.0 to 1.0.
	Rate float64 `json:"rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Rate        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowKeyBasedSamplingResponse) RawJSON() string { return r.JSON.raw }
func (r *ShadowKeyBasedSamplingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Traffic source for a shadow experiment. The public API supports endpoint sources
// only.
//
// The property Endpoint is required.
type ShadowSourceParam struct {
	// Endpoint-level source that samples endpoint traffic at the API gateway.
	Endpoint ShadowEndpointSourceParam `json:"endpoint,omitzero" api:"required"`
	paramObj
}

func (r ShadowSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Endpoint traffic source returned for a shadow experiment.
type ShadowSourceResponse struct {
	// Endpoint-level source returned for a shadow experiment.
	Endpoint ShadowEndpointSourceResponse `json:"endpoint" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Endpoint    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowSourceResponse) RawJSON() string { return r.JSON.raw }
func (r *ShadowSourceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fixed-rate random sampling of endpoint requests.
//
// The property Rate is required.
type ShadowUniformSamplingParam struct {
	// Required fraction of requests to sample, from 0.0 to 1.0.
	Rate float64 `json:"rate" api:"required"`
	paramObj
}

func (r ShadowUniformSamplingParam) MarshalJSON() (data []byte, err error) {
	type shadow ShadowUniformSamplingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShadowUniformSamplingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fixed-rate random sampling returned by the API. A zero rate may be omitted by
// JSON serialization.
type ShadowUniformSamplingResponse struct {
	// Fraction of requests sampled, from 0.0 to 1.0.
	Rate float64 `json:"rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rate        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowUniformSamplingResponse) RawJSON() string { return r.JSON.raw }
func (r *ShadowUniformSamplingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token, request, and batching throughput over a time range.
type ThroughputMetrics struct {
	// Average number of batches queued or in flight in the serving engine.
	AvgBatchDepth float64 `json:"avgBatchDepth"`
	// Average number of requests processed in each runtime batch.
	AvgBatchSize float64 `json:"avgBatchSize"`
	// Average completed requests per second.
	RequestsPerSecond float64 `json:"requestsPerSecond"`
	// Average generated tokens per second.
	TokensPerSecond float64 `json:"tokensPerSecond"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgBatchDepth     respjson.Field
		AvgBatchSize      respjson.Field
		RequestsPerSecond respjson.Field
		TokensPerSecond   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThroughputMetrics) RawJSON() string { return r.JSON.raw }
func (r *ThroughputMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Timestamped bucket containing one or more named metric values.
type TimeSeriesDataPoint struct {
	// Start time of the metric bucket.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// Metric names mapped to their numeric values for this bucket.
	Values map[string]float64 `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeSeriesDataPoint) RawJSON() string { return r.JSON.raw }
func (r *TimeSeriesDataPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aggregate and per-request token usage over a time range.
type TokenMetrics struct {
	// Average input tokens per request.
	AvgInputTokens float64 `json:"avgInputTokens"`
	// Average output tokens per request.
	AvgOutputTokens float64 `json:"avgOutputTokens"`
	// Total input tokens processed during the time range.
	TotalInputTokens string `json:"totalInputTokens"`
	// Total output tokens generated during the time range.
	TotalOutputTokens string `json:"totalOutputTokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgInputTokens    respjson.Field
		AvgOutputTokens   respjson.Field
		TotalInputTokens  respjson.Field
		TotalOutputTokens respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenMetrics) RawJSON() string { return r.JSON.raw }
func (r *TokenMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Empty response returned after a successful delete operation.
type BetaEndpointDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Endpoint-wide usage and performance analytics with optional time-series and
// per-deployment breakdowns.
type BetaEndpointAnalyticsResponse struct {
	// Per-deployment analytics.
	DeploymentAnalytics []BetaEndpointAnalyticsResponseDeploymentAnalytics `json:"deploymentAnalytics"`
	// ID of the endpoint summarized by these analytics.
	EndpointID string `json:"endpointId"`
	// Operational metrics aggregated across all deployments receiving traffic for an
	// endpoint.
	Metrics BetaEndpointAnalyticsResponseMetrics `json:"metrics"`
	// Closed-open time range covered by the analytics.
	TimeRange MetricsTimeRange `json:"timeRange"`
	// Per-bucket metric samples, included only when `includeTimeSeries` is true.
	TimeSeries []TimeSeriesDataPoint `json:"timeSeries"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeploymentAnalytics respjson.Field
		EndpointID          respjson.Field
		Metrics             respjson.Field
		TimeRange           respjson.Field
		TimeSeries          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointAnalyticsResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointAnalyticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage and performance analytics for one deployment under an endpoint.
type BetaEndpointAnalyticsResponseDeploymentAnalytics struct {
	// ID of the deployment summarized by these analytics.
	DeploymentID string `json:"deploymentId"`
	// ID of the deployment's parent endpoint.
	EndpointID string `json:"endpointId"`
	// Aggregate operational metrics for the deployment.
	Metrics DeploymentMetrics `json:"metrics"`
	// Closed-open time range covered by the analytics.
	TimeRange MetricsTimeRange `json:"timeRange"`
	// Per-bucket metric samples for the deployment.
	TimeSeries []TimeSeriesDataPoint `json:"timeSeries"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeploymentID respjson.Field
		EndpointID   respjson.Field
		Metrics      respjson.Field
		TimeRange    respjson.Field
		TimeSeries   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointAnalyticsResponseDeploymentAnalytics) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointAnalyticsResponseDeploymentAnalytics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operational metrics aggregated across all deployments receiving traffic for an
// endpoint.
type BetaEndpointAnalyticsResponseMetrics struct {
	// Per-deployment breakdown, if the endpoint has multiple deployments.
	DeploymentMetrics []DeploymentMetrics `json:"deploymentMetrics"`
	// The endpoint these metrics describe.
	EndpointID string `json:"endpointId"`
	// Error rate and counts by error type.
	ErrorMetrics ErrorMetrics `json:"errorMetrics"`
	// Time-to-first-token, end-to-end, and inter-token latency percentiles.
	LatencyMetrics LatencyMetrics `json:"latencyMetrics"`
	// Request counts and rates.
	RequestMetrics RequestMetrics `json:"requestMetrics"`
	// Average CPU, GPU, memory, and network utilization.
	ResourceUtilization ResourceUtilization `json:"resourceUtilization"`
	// Token, request, and batching throughput.
	ThroughputMetrics ThroughputMetrics `json:"throughputMetrics"`
	// Closed-open time range used by metrics and analytics responses.
	TimeRange MetricsTimeRange `json:"timeRange"`
	// Input and output token totals and averages.
	TokenMetrics TokenMetrics `json:"tokenMetrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeploymentMetrics   respjson.Field
		EndpointID          respjson.Field
		ErrorMetrics        respjson.Field
		LatencyMetrics      respjson.Field
		RequestMetrics      respjson.Field
		ResourceUtilization respjson.Field
		ThroughputMetrics   respjson.Field
		TimeRange           respjson.Field
		TokenMetrics        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointAnalyticsResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointAnalyticsResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One endpoint- or deployment-scoped entry in an endpoint's combined audit and
// lifecycle feed.
type BetaEndpointListEventsResponse struct {
	// Output only. Unique event identifier.
	ID string `json:"id" api:"required"`
	// Output only. Event creation time.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Output only. The endpoint this event belongs to. Always set.
	EndpointID string `json:"endpointId" api:"required"`
	// Output only. Severity level.
	//
	// Any of "LEVEL_DEBUG", "LEVEL_INFO", "LEVEL_WARN", "LEVEL_ERROR".
	Level BetaEndpointListEventsResponseLevel `json:"level" api:"required"`
	// Output only. Service, cluster, or controller that emitted the event.
	Source string `json:"source" api:"required"`
	// Output only. Whether this row describes the endpoint or one of its deployments.
	//
	// Any of "SOURCE_KIND_ENDPOINT", "SOURCE_KIND_DEPLOYMENT".
	SourceKind BetaEndpointListEventsResponseSourceKind `json:"sourceKind" api:"required"`
	// Output only. Stable event type, such as `endpoint.updated`,
	// `deployment.created`, `deployment.scaled`, `condition.set`, or `pod.log`.
	Type string `json:"type" api:"required"`
	// ID of the cluster associated with a cluster-scoped event.
	ClusterID string `json:"clusterId"`
	// Stable public component label associated with a replica event, such as `engine`,
	// `model-download`, or `sidecar`.
	ContainerName string `json:"containerName"`
	// Output only. Deployment associated with the event when `sourceKind` is
	// `SOURCE_KIND_DEPLOYMENT`.
	DeploymentID string `json:"deploymentId"`
	// Short diagnostic log excerpt captured with a pod event, for example during a
	// crash, out-of-memory termination, or image pull failure. This field is truncated
	// and is not a streaming log API.
	LogExcerpt string `json:"logExcerpt"`
	// Output only. Human-readable description of the event. Short and stable; not
	// structured data.
	Message string `json:"message"`
	// Resource name at the time of the event. Populated by: deployment.created,
	// deployment.deleted, endpoint.created, endpoint.deleted
	Name string `json:"name"`
	// New replica count for a `deployment.scaled` event.
	NewReplicas int64 `json:"newReplicas"`
	// Opaque node handle for correlating replica failures on the same node. Omitted
	// when the replica is unscheduled or the node is unknown.
	NodeID string `json:"nodeId"`
	// Replica-count transition. Populated by: deployment.scaled
	OldReplicas int64 `json:"oldReplicas"`
	// Field-mask paths that were modified. Populated by: deployment.updated,
	// endpoint.updated
	Paths []string `json:"paths"`
	// Stable condition reason, such as `AllReplicasReady`, `ReplicasProgressing`, or
	// `ApplySuccessful`.
	Reason string `json:"reason"`
	// Opaque replica identity associated with a `pod.*` event, stable for grouping
	// events from the same replica.
	ReplicaID string `json:"replicaId"`
	// Deployment subservice associated with the event, such as `model-deployment` or
	// `speculator-deployment`.
	ServiceType string `json:"serviceType"`
	// Condition status for `condition.set` and `cluster_condition.set`: `True`,
	// `False`, or `Unknown`. The condition type is carried in `subjectId`.
	Status string `json:"status"`
	// Output only. ID of the event's subject, such as a rollout, shadow target, or
	// condition type.
	SubjectID string `json:"subjectId"`
	// Target version. Populated by `target.created`; the target ID is carried in
	// `subjectId`.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		EndpointID    respjson.Field
		Level         respjson.Field
		Source        respjson.Field
		SourceKind    respjson.Field
		Type          respjson.Field
		ClusterID     respjson.Field
		ContainerName respjson.Field
		DeploymentID  respjson.Field
		LogExcerpt    respjson.Field
		Message       respjson.Field
		Name          respjson.Field
		NewReplicas   respjson.Field
		NodeID        respjson.Field
		OldReplicas   respjson.Field
		Paths         respjson.Field
		Reason        respjson.Field
		ReplicaID     respjson.Field
		ServiceType   respjson.Field
		Status        respjson.Field
		SubjectID     respjson.Field
		Version       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointListEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointListEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Output only. Severity level.
type BetaEndpointListEventsResponseLevel string

const (
	BetaEndpointListEventsResponseLevelLevelDebug BetaEndpointListEventsResponseLevel = "LEVEL_DEBUG"
	BetaEndpointListEventsResponseLevelLevelInfo  BetaEndpointListEventsResponseLevel = "LEVEL_INFO"
	BetaEndpointListEventsResponseLevelLevelWarn  BetaEndpointListEventsResponseLevel = "LEVEL_WARN"
	BetaEndpointListEventsResponseLevelLevelError BetaEndpointListEventsResponseLevel = "LEVEL_ERROR"
)

// Output only. Whether this row describes the endpoint or one of its deployments.
type BetaEndpointListEventsResponseSourceKind string

const (
	BetaEndpointListEventsResponseSourceKindSourceKindEndpoint   BetaEndpointListEventsResponseSourceKind = "SOURCE_KIND_ENDPOINT"
	BetaEndpointListEventsResponseSourceKindSourceKindDeployment BetaEndpointListEventsResponseSourceKind = "SOURCE_KIND_DEPLOYMENT"
)

type BetaEndpointNewParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Inference-addressable endpoint name to create.
	Name string `json:"name" api:"required"`
	// Who can discover the endpoint. `VISIBILITY_PRIVATE` restricts it to the project;
	// `VISIBILITY_INTERNAL` shares it with the organization.
	//
	// Any of "VISIBILITY_PRIVATE", "VISIBILITY_INTERNAL".
	Visibility BetaEndpointNewParamsVisibility `json:"visibility,omitzero"`
	paramObj
}

func (r BetaEndpointNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who can discover the endpoint. `VISIBILITY_PRIVATE` restricts it to the project;
// `VISIBILITY_INTERNAL` shares it with the organization.
type BetaEndpointNewParamsVisibility string

const (
	BetaEndpointNewParamsVisibilityVisibilityPrivate  BetaEndpointNewParamsVisibility = "VISIBILITY_PRIVATE"
	BetaEndpointNewParamsVisibilityVisibilityInternal BetaEndpointNewParamsVisibility = "VISIBILITY_INTERNAL"
)

type BetaEndpointGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	paramObj
}

type BetaEndpointUpdateParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Fields to update. If not set, the fields populated are updated.
	UpdateMask param.Opt[string] `query:"updateMask,omitzero" json:"-"`
	// Current endpoint version. The update is rejected if this value no longer
	// matches.
	Etag param.Opt[string] `json:"etag,omitzero"`
	// Updated endpoint string.
	Name param.Opt[string] `json:"name,omitzero"`
	// Replacement live traffic split. Use an empty list to stop routing live traffic.
	TrafficSplit []EndpointTrafficSplitEntryParam `json:"trafficSplit,omitzero"`
	// Who can discover the endpoint. `VISIBILITY_PRIVATE` restricts it to the project;
	// `VISIBILITY_INTERNAL` shares it with the organization.
	//
	// Any of "VISIBILITY_PRIVATE", "VISIBILITY_INTERNAL".
	Visibility BetaEndpointUpdateParamsVisibility `json:"visibility,omitzero"`
	paramObj
}

func (r BetaEndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaEndpointUpdateParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Who can discover the endpoint. `VISIBILITY_PRIVATE` restricts it to the project;
// `VISIBILITY_INTERNAL` shares it with the organization.
type BetaEndpointUpdateParamsVisibility string

const (
	BetaEndpointUpdateParamsVisibilityVisibilityPrivate  BetaEndpointUpdateParamsVisibility = "VISIBILITY_PRIVATE"
	BetaEndpointUpdateParamsVisibilityVisibilityInternal BetaEndpointUpdateParamsVisibility = "VISIBILITY_INTERNAL"
)

type BetaEndpointListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Filter expression using `name`, `created_at`, or `updated_at` with comparison
	// operators and AND/OR/NOT; timestamps must be RFC 3339 strings. `name` supports
	// substring matching with `:` and prefix/suffix wildcards with `*`, and accepts a
	// bare endpoint name or `<project_slug>/<endpoint_name>`.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// Maximum number of endpoints to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort field for the results. Supports `created_at` or `updated_at`, optionally
	// followed by `asc` or `desc`.
	OrderBy param.Opt[string] `query:"orderBy,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointListParams]'s query parameters as `url.Values`.
func (r BetaEndpointListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointDeleteParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Etag for optimistic concurrency. If set, the delete is rejected if the current
	// etag does not match.
	Etag param.Opt[string] `query:"etag,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointDeleteParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointAnalyticsParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Restrict to a single deployment under this endpoint.
	DeploymentID param.Opt[string] `query:"deploymentId,omitzero" json:"-"`
	// Exclusive end of the time range. Defaults to now if unset.
	EndTime param.Opt[time.Time] `query:"endTime,omitzero" format:"date-time" json:"-"`
	// Time-series bucket duration, such as `1m`, `1h`, or `1d`. Defaults to `1d`.
	Granularity param.Opt[string] `query:"granularity,omitzero" json:"-"`
	// When true, include per-bucket time series in the response.
	IncludeTimeSeries param.Opt[bool] `query:"includeTimeSeries,omitzero" json:"-"`
	// Inclusive start of the time range. Defaults to 24 hours ago if unset.
	StartTime param.Opt[time.Time] `query:"startTime,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointAnalyticsParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointAnalyticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointListEventsParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous endpoint event list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of events to return. Max 10000, defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Return only events at or after this time.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// ID of a subject associated with the event, such as a rollout. Combined with
	// other filters using AND.
	SubjectID param.Opt[string] `query:"subjectId,omitzero" json:"-"`
	// Return only events strictly before this time.
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	// Deployment IDs whose events should be included. Every ID must belong to the
	// endpoint. Supplying this filter excludes endpoint-scoped events unless
	// `SOURCE_KIND_ENDPOINT` is also included in `sourceKinds`.
	DeploymentIDs []string `query:"deploymentIds,omitzero" json:"-"`
	// Minimum severity. Omit to disable severity filtering.
	//
	// Any of "LEVEL_DEBUG", "LEVEL_INFO", "LEVEL_WARN", "LEVEL_ERROR".
	MinLevel BetaEndpointListEventsParamsMinLevel `query:"minLevel,omitzero" json:"-"`
	// Resource kinds whose events should be included. Omit to include both endpoint-
	// and deployment-scoped events.
	//
	// Any of "SOURCE_KIND_ENDPOINT", "SOURCE_KIND_DEPLOYMENT".
	SourceKinds []string `query:"sourceKinds,omitzero" json:"-"`
	// Event types to include, such as `deployment.scaled` or `condition.set`. Combined
	// with other filters using AND.
	Types []string `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointListEventsParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Minimum severity. Omit to disable severity filtering.
type BetaEndpointListEventsParamsMinLevel string

const (
	BetaEndpointListEventsParamsMinLevelLevelDebug BetaEndpointListEventsParamsMinLevel = "LEVEL_DEBUG"
	BetaEndpointListEventsParamsMinLevelLevelInfo  BetaEndpointListEventsParamsMinLevel = "LEVEL_INFO"
	BetaEndpointListEventsParamsMinLevelLevelWarn  BetaEndpointListEventsParamsMinLevel = "LEVEL_WARN"
	BetaEndpointListEventsParamsMinLevelLevelError BetaEndpointListEventsParamsMinLevel = "LEVEL_ERROR"
)

type BetaEndpointListOrgScopedParams struct {
	// Cursor from a previous list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Filter expression using `name`, `created_at`, or `updated_at` with comparison
	// operators and AND/OR/NOT; timestamps must be RFC 3339 strings. `name` supports
	// substring matching with `:` and prefix/suffix wildcards with `*`, and must be a
	// bare endpoint name.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort field for the results. Supports `created_at` or `updated_at`, optionally
	// followed by `asc` or `desc`.
	OrderBy param.Opt[string] `query:"orderBy,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointListOrgScopedParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointListOrgScopedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
