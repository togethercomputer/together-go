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

// BetaEndpointRolloutService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaEndpointRolloutService] method instead.
type BetaEndpointRolloutService struct {
	Options []option.RequestOption
}

// NewBetaEndpointRolloutService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaEndpointRolloutService(opts ...option.RequestOption) (r BetaEndpointRolloutService) {
	r = BetaEndpointRolloutService{}
	r.Options = opts
	return
}

// Creates a rollout in the pending state without shifting traffic. Start the
// rollout in a separate request after reviewing its strategy and metric gates.
func (r *BetaEndpointRolloutService) New(ctx context.Context, endpointID string, params BetaEndpointRolloutNewParams, opts ...option.RequestOption) (res *Rollout, err error) {
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
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts", params.ProjectID.Value, endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a rollout's strategy, lifecycle state, current traffic percentage,
// step history, and metric-gate results.
func (r *BetaEndpointRolloutService) Get(ctx context.Context, id string, query BetaEndpointRolloutGetParams, opts ...option.RequestOption) (res *Rollout, err error) {
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
	if query.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts/%s", query.ProjectID.Value, query.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists rollout histories for an endpoint. Use `filter=ROLLOUT_FILTER_ACTIVE` to
// return only the active rollout, if one exists.
func (r *BetaEndpointRolloutService) List(ctx context.Context, endpointID string, params BetaEndpointRolloutListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Rollout], err error) {
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
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts", params.ProjectID.Value, endpointID)
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

// Lists rollout histories for an endpoint. Use `filter=ROLLOUT_FILTER_ACTIVE` to
// return only the active rollout, if one exists.
func (r *BetaEndpointRolloutService) ListAutoPaging(ctx context.Context, endpointID string, params BetaEndpointRolloutListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Rollout] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, endpointID, params, opts...))
}

// Deletes a rollout record. An active rollout must be aborted or completed before
// it can be deleted.
func (r *BetaEndpointRolloutService) Delete(ctx context.Context, id string, params BetaEndpointRolloutDeleteParams, opts ...option.RequestOption) (res *BetaEndpointRolloutDeleteResponse, err error) {
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
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts/%s", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Cancels a running, pausing, paused, system-paused, or stabilizing rollout by
// freezing the current traffic split into standing weights. Revert is removed and
// rejected; after canceling, start another canary rollout in either direction or
// rebalance the traffic split. The response is the accepted rollout snapshot; poll
// GetRollout until it reaches CANCELED.
func (r *BetaEndpointRolloutService) Cancel(ctx context.Context, id string, params BetaEndpointRolloutCancelParams, opts ...option.RequestOption) (res *Rollout, err error) {
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
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts/%s/cancel", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Requests a running or stabilizing rollout to pause and records an optional
// reason. The response returns the PAUSING snapshot; poll GetRollout until state
// is PAUSED to confirm the executor has parked.
func (r *BetaEndpointRolloutService) Pause(ctx context.Context, id string, params BetaEndpointRolloutPauseParams, opts ...option.RequestOption) (res *Rollout, err error) {
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
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts/%s/pause", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the values a create request would pick for any field left unset, plus
// the capacity context needed to display them, without creating a rollout.
// Responses are display state only and re-validated authoritatively at create and
// start; do not copy response values back into a create request.
func (r *BetaEndpointRolloutService) PreviewDefaults(ctx context.Context, endpointID string, params BetaEndpointRolloutPreviewDefaultsParams, opts ...option.RequestOption) (res *RolloutDefaultsPreview, err error) {
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
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts/preview-defaults", params.ProjectID.Value, endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Completes a running or paused rollout immediately by sending all live traffic to
// the target deployment.
func (r *BetaEndpointRolloutService) Promote(ctx context.Context, id string, params BetaEndpointRolloutPromoteParams, opts ...option.RequestOption) (res *Rollout, err error) {
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
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts/%s/promote", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Resumes a pausing, paused, or system-paused rollout from its current step and
// traffic split.
func (r *BetaEndpointRolloutService) Resume(ctx context.Context, id string, params BetaEndpointRolloutResumeParams, opts ...option.RequestOption) (res *Rollout, err error) {
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
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts/%s/resume", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Starts a pending rollout and begins its configured traffic-shifting workflow.
func (r *BetaEndpointRolloutService) Start(ctx context.Context, id string, body BetaEndpointRolloutStartParams, opts ...option.RequestOption) (res *Rollout, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.ProjectID)
	if body.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if body.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/rollouts/%s/start", body.ProjectID.Value, body.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Blue-green strategy configuration for a single cutover to the target deployment.
type BlueGreenConfig struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlueGreenConfig) RawJSON() string { return r.JSON.raw }
func (r *BlueGreenConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BlueGreenConfig to a BlueGreenConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BlueGreenConfigParam.Overrides()
func (r BlueGreenConfig) ToParam() BlueGreenConfigParam {
	return param.Override[BlueGreenConfigParam](json.RawMessage(r.RawJSON()))
}

// Blue-green strategy configuration for a single cutover to the target deployment.
type BlueGreenConfigParam struct {
	paramObj
}

func (r BlueGreenConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow BlueGreenConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlueGreenConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canary strategy configuration for gradual traffic progression. An empty config
// uses the default 5, 25, 50, 100 percent ladder; over a frozen traffic-split pair
// left by cancel, the default ladder is derived at start from the pair's current
// served share so it begins above it.
type CanaryConfig struct {
	// Optional positive soak between steps. Defaults to 3m if omitted, and grows to
	// cover metric rule windows plus ingestion lag.
	StepInterval string `json:"stepInterval"`
	// Optional progression steps. Defaults to 5, 25, 50, 100 percent when empty;
	// explicit steps must increase and end at 100 percent.
	Steps []RolloutStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StepInterval respjson.Field
		Steps        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CanaryConfig) RawJSON() string { return r.JSON.raw }
func (r *CanaryConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CanaryConfig to a CanaryConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CanaryConfigParam.Overrides()
func (r CanaryConfig) ToParam() CanaryConfigParam {
	return param.Override[CanaryConfigParam](json.RawMessage(r.RawJSON()))
}

// Canary strategy configuration for gradual traffic progression. An empty config
// uses the default 5, 25, 50, 100 percent ladder; over a frozen traffic-split pair
// left by cancel, the default ladder is derived at start from the pair's current
// served share so it begins above it.
type CanaryConfigParam struct {
	// Optional positive soak between steps. Defaults to 3m if omitted, and grows to
	// cover metric rule windows plus ingestion lag.
	StepInterval param.Opt[string] `json:"stepInterval,omitzero"`
	// Optional progression steps. Defaults to 5, 25, 50, 100 percent when empty;
	// explicit steps must increase and end at 100 percent.
	Steps []RolloutStepParam `json:"steps,omitzero"`
	paramObj
}

func (r CanaryConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow CanaryConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CanaryConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Observed metric result enriched with rollout rule criteria and the rule's
// recorded verdict. Unmeasured rules are synthesized with verdict
// METRIC_VERDICT_UNAVAILABLE and no source or target value.
type MetricResult struct {
	// Evaluation form used by the metric rule.
	//
	// Any of "METRIC_CHECK_TYPE_THRESHOLD", "METRIC_CHECK_TYPE_REGRESSION".
	Check MetricResultCheck `json:"check"`
	// Direction that indicates whether higher or lower values are worse.
	//
	// Any of "REGRESSION_DIRECTION_HIGHER_IS_WORSE",
	// "REGRESSION_DIRECTION_LOWER_IS_WORSE".
	Direction MetricResultDirection `json:"direction"`
	// Rule-specific failure text. Set only when verdict is METRIC_VERDICT_BREACHED and
	// the gate recorded one.
	FailureReason string `json:"failureReason"`
	// Regression percentage limit used when check is METRIC_CHECK_TYPE_REGRESSION.
	MaxRegressionPercent float64 `json:"maxRegressionPercent"`
	// Metric name as exported to the observability backend.
	Name string `json:"name"`
	// Threshold comparison operator.
	//
	// Any of "THRESHOLD_OPERATOR_GT", "THRESHOLD_OPERATOR_GTE",
	// "THRESHOLD_OPERATOR_LT", "THRESHOLD_OPERATOR_LTE".
	Operator MetricResultOperator `json:"operator"`
	// Percentile value, such as 99. Set only when stat is METRIC_STAT_TYPE_PERCENTILE.
	Percentile int64 `json:"percentile"`
	// Observed source baseline. Set only for regression checks with a recorded
	// observation; a 0 reading serializes explicitly.
	SourceValue float64 `json:"sourceValue"`
	// Aggregation used for the metric.
	//
	// Any of "METRIC_STAT_TYPE_AVG", "METRIC_STAT_TYPE_PERCENTILE".
	Stat MetricResultStat `json:"stat"`
	// Observed target value. Set when the gate recorded an observation; absent on
	// synthesized unavailable results. A 0 reading serializes explicitly.
	TargetValue float64 `json:"targetValue"`
	// Threshold criteria used when check is METRIC_CHECK_TYPE_THRESHOLD.
	Threshold float64 `json:"threshold"`
	// Rule decision recorded by the metric gate. Absent when no decision was recorded.
	//
	// Any of "METRIC_VERDICT_PASS", "METRIC_VERDICT_BREACHED",
	// "METRIC_VERDICT_UNAVAILABLE".
	Verdict MetricResultVerdict `json:"verdict"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Check                respjson.Field
		Direction            respjson.Field
		FailureReason        respjson.Field
		MaxRegressionPercent respjson.Field
		Name                 respjson.Field
		Operator             respjson.Field
		Percentile           respjson.Field
		SourceValue          respjson.Field
		Stat                 respjson.Field
		TargetValue          respjson.Field
		Threshold            respjson.Field
		Verdict              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricResult) RawJSON() string { return r.JSON.raw }
func (r *MetricResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Evaluation form used by the metric rule.
type MetricResultCheck string

const (
	MetricResultCheckMetricCheckTypeThreshold  MetricResultCheck = "METRIC_CHECK_TYPE_THRESHOLD"
	MetricResultCheckMetricCheckTypeRegression MetricResultCheck = "METRIC_CHECK_TYPE_REGRESSION"
)

// Direction that indicates whether higher or lower values are worse.
type MetricResultDirection string

const (
	MetricResultDirectionRegressionDirectionHigherIsWorse MetricResultDirection = "REGRESSION_DIRECTION_HIGHER_IS_WORSE"
	MetricResultDirectionRegressionDirectionLowerIsWorse  MetricResultDirection = "REGRESSION_DIRECTION_LOWER_IS_WORSE"
)

// Threshold comparison operator.
type MetricResultOperator string

const (
	MetricResultOperatorThresholdOperatorGt  MetricResultOperator = "THRESHOLD_OPERATOR_GT"
	MetricResultOperatorThresholdOperatorGte MetricResultOperator = "THRESHOLD_OPERATOR_GTE"
	MetricResultOperatorThresholdOperatorLt  MetricResultOperator = "THRESHOLD_OPERATOR_LT"
	MetricResultOperatorThresholdOperatorLte MetricResultOperator = "THRESHOLD_OPERATOR_LTE"
)

// Aggregation used for the metric.
type MetricResultStat string

const (
	MetricResultStatMetricStatTypeAvg        MetricResultStat = "METRIC_STAT_TYPE_AVG"
	MetricResultStatMetricStatTypePercentile MetricResultStat = "METRIC_STAT_TYPE_PERCENTILE"
)

// Rule decision recorded by the metric gate. Absent when no decision was recorded.
type MetricResultVerdict string

const (
	MetricResultVerdictMetricVerdictPass        MetricResultVerdict = "METRIC_VERDICT_PASS"
	MetricResultVerdictMetricVerdictBreached    MetricResultVerdict = "METRIC_VERDICT_BREACHED"
	MetricResultVerdictMetricVerdictUnavailable MetricResultVerdict = "METRIC_VERDICT_UNAVAILABLE"
)

// Metric gate evaluated during a rollout.
type MetricRule struct {
	// Required catalogue key for the metric to gate on. `serving_latency` is retired.
	//
	// Any of "inflight_requests", "router_error_rate", "router_latency".
	Name MetricRuleName `json:"name" api:"required"`
	// Percentile value, such as 99. Set only when stat is METRIC_STAT_TYPE_PERCENTILE.
	Percentile int64 `json:"percentile"`
	// Regression criteria that fail when the target regresses against the source
	// beyond a limit.
	RegressionCheck RegressionCheck `json:"regressionCheck"`
	// Aggregation used for the metric. Optional for router_error_rate and
	// inflight_requests; omitted values default to METRIC_STAT_TYPE_AVG. Required for
	// router_latency, where AVG or PERCENTILE may be used.
	//
	// Any of "METRIC_STAT_TYPE_AVG", "METRIC_STAT_TYPE_PERCENTILE".
	Stat MetricRuleStat `json:"stat"`
	// Threshold criteria that fail when the target metric violates the configured
	// bound.
	ThresholdCheck ThresholdCheck `json:"thresholdCheck"`
	// Optional query window for the metric. Defaults to the step soak duration.
	Window string `json:"window"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Percentile      respjson.Field
		RegressionCheck respjson.Field
		Stat            respjson.Field
		ThresholdCheck  respjson.Field
		Window          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricRule) RawJSON() string { return r.JSON.raw }
func (r *MetricRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MetricRule to a MetricRuleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MetricRuleParam.Overrides()
func (r MetricRule) ToParam() MetricRuleParam {
	return param.Override[MetricRuleParam](json.RawMessage(r.RawJSON()))
}

// Required catalogue key for the metric to gate on. `serving_latency` is retired.
type MetricRuleName string

const (
	MetricRuleNameInflightRequests MetricRuleName = "inflight_requests"
	MetricRuleNameRouterErrorRate  MetricRuleName = "router_error_rate"
	MetricRuleNameRouterLatency    MetricRuleName = "router_latency"
)

// Aggregation used for the metric. Optional for router_error_rate and
// inflight_requests; omitted values default to METRIC_STAT_TYPE_AVG. Required for
// router_latency, where AVG or PERCENTILE may be used.
type MetricRuleStat string

const (
	MetricRuleStatMetricStatTypeAvg        MetricRuleStat = "METRIC_STAT_TYPE_AVG"
	MetricRuleStatMetricStatTypePercentile MetricRuleStat = "METRIC_STAT_TYPE_PERCENTILE"
)

// Metric gate evaluated during a rollout.
//
// The property Name is required.
type MetricRuleParam struct {
	// Required catalogue key for the metric to gate on. `serving_latency` is retired.
	//
	// Any of "inflight_requests", "router_error_rate", "router_latency".
	Name MetricRuleName `json:"name,omitzero" api:"required"`
	// Percentile value, such as 99. Set only when stat is METRIC_STAT_TYPE_PERCENTILE.
	Percentile param.Opt[int64] `json:"percentile,omitzero"`
	// Optional query window for the metric. Defaults to the step soak duration.
	Window param.Opt[string] `json:"window,omitzero"`
	// Regression criteria that fail when the target regresses against the source
	// beyond a limit.
	RegressionCheck RegressionCheckParam `json:"regressionCheck,omitzero"`
	// Aggregation used for the metric. Optional for router_error_rate and
	// inflight_requests; omitted values default to METRIC_STAT_TYPE_AVG. Required for
	// router_latency, where AVG or PERCENTILE may be used.
	//
	// Any of "METRIC_STAT_TYPE_AVG", "METRIC_STAT_TYPE_PERCENTILE".
	Stat MetricRuleStat `json:"stat,omitzero"`
	// Threshold criteria that fail when the target metric violates the configured
	// bound.
	ThresholdCheck ThresholdCheckParam `json:"thresholdCheck,omitzero"`
	paramObj
}

func (r MetricRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow MetricRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MetricRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pause metadata returned while a rollout is paused.
type PauseInfo struct {
	// Timestamp when the rollout was paused.
	PausedAt time.Time `json:"pausedAt" api:"required" format:"date-time"`
	// Human-readable reason recorded when the rollout was paused.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PausedAt    respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PauseInfo) RawJSON() string { return r.JSON.raw }
func (r *PauseInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Regression criteria that fail when the target regresses against the source
// beyond a limit.
type RegressionCheck struct {
	// Required direction that indicates whether higher or lower metric values are
	// worse.
	//
	// Any of "REGRESSION_DIRECTION_HIGHER_IS_WORSE",
	// "REGRESSION_DIRECTION_LOWER_IS_WORSE".
	Direction RegressionCheckDirection `json:"direction" api:"required"`
	// Finite maximum allowed regression percentage, greater than or equal to 0.
	// Omitting this value is read as 0. A value of 0 is the strictest budget; any
	// regression fails, and exactly-at-budget passes.
	MaxRegressionPercent float64 `json:"maxRegressionPercent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction            respjson.Field
		MaxRegressionPercent respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegressionCheck) RawJSON() string { return r.JSON.raw }
func (r *RegressionCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RegressionCheck to a RegressionCheckParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RegressionCheckParam.Overrides()
func (r RegressionCheck) ToParam() RegressionCheckParam {
	return param.Override[RegressionCheckParam](json.RawMessage(r.RawJSON()))
}

// Required direction that indicates whether higher or lower metric values are
// worse.
type RegressionCheckDirection string

const (
	RegressionCheckDirectionRegressionDirectionHigherIsWorse RegressionCheckDirection = "REGRESSION_DIRECTION_HIGHER_IS_WORSE"
	RegressionCheckDirectionRegressionDirectionLowerIsWorse  RegressionCheckDirection = "REGRESSION_DIRECTION_LOWER_IS_WORSE"
)

// Regression criteria that fail when the target regresses against the source
// beyond a limit.
//
// The property Direction is required.
type RegressionCheckParam struct {
	// Required direction that indicates whether higher or lower metric values are
	// worse.
	//
	// Any of "REGRESSION_DIRECTION_HIGHER_IS_WORSE",
	// "REGRESSION_DIRECTION_LOWER_IS_WORSE".
	Direction RegressionCheckDirection `json:"direction,omitzero" api:"required"`
	// Finite maximum allowed regression percentage, greater than or equal to 0.
	// Omitting this value is read as 0. A value of 0 is the strictest budget; any
	// regression fails, and exactly-at-budget passes.
	MaxRegressionPercent param.Opt[float64] `json:"maxRegressionPercent,omitzero"`
	paramObj
}

func (r RegressionCheckParam) MarshalJSON() (data []byte, err error) {
	type shadow RegressionCheckParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegressionCheckParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rolling strategy configuration for capacity-preserving batches that ramp target
// replicas up while draining source replicas.
type RollingConfig struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RollingConfig) RawJSON() string { return r.JSON.raw }
func (r *RollingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RollingConfig to a RollingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RollingConfigParam.Overrides()
func (r RollingConfig) ToParam() RollingConfigParam {
	return param.Override[RollingConfigParam](json.RawMessage(r.RawJSON()))
}

// Rolling strategy configuration for capacity-preserving batches that ramp target
// replicas up while draining source replicas.
type RollingConfigParam struct {
	paramObj
}

func (r RollingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow RollingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RollingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public view of a rollout resource, including runtime progress and any pause or
// abort reason.
type Rollout struct {
	// Output only. Unique rollout identifier.
	ID string `json:"id" api:"required"`
	// Output only. Timestamp when the rollout was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Output only. Endpoint this rollout belongs to.
	EndpointID string `json:"endpointId" api:"required"`
	// Output only. Deployment that traffic is shifting away from.
	SourceDeploymentID string `json:"sourceDeploymentId" api:"required"`
	// Output only. High-level rollout lifecycle state.
	//
	// Any of "ROLLOUT_STATE_RUNNING", "ROLLOUT_STATE_PAUSED",
	// "ROLLOUT_STATE_STABILIZING", "ROLLOUT_STATE_COMPLETED", "ROLLOUT_STATE_PENDING",
	// "ROLLOUT_STATE_SYSTEM_PAUSED", "ROLLOUT_STATE_CANCELLING",
	// "ROLLOUT_STATE_CANCELED", "ROLLOUT_STATE_PAUSING".
	State RolloutState `json:"state" api:"required"`
	// Derived runtime progress for a rollout.
	Status RolloutStatus `json:"status" api:"required"`
	// Output only. Rollout strategy selected at creation.
	//
	// Any of "ROLLOUT_STRATEGY_TYPE_ROLLING", "ROLLOUT_STRATEGY_TYPE_CANARY",
	// "ROLLOUT_STRATEGY_TYPE_BLUE_GREEN".
	Strategy RolloutStrategy `json:"strategy" api:"required"`
	// Output only. Deployment that traffic is shifting toward.
	TargetDeploymentID string `json:"targetDeploymentId" api:"required"`
	// Output only. Timestamp when the rollout reached a terminal state.
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Output only. Zero-based index of the current step. Unset while PENDING; step 0
	// is reported explicitly after start.
	CurrentStep int64 `json:"currentStep"`
	// Output only. Applied percentage of traffic on the target deployment.
	CurrentTrafficPercent int64 `json:"currentTrafficPercent"`
	// Output only. Opaque version tag for optimistic concurrency control.
	Etag string `json:"etag"`
	// Pause metadata returned while a rollout is paused.
	PauseInfo PauseInfo `json:"pauseInfo"`
	// Output only. Timestamp when the rollout started running.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		EndpointID            respjson.Field
		SourceDeploymentID    respjson.Field
		State                 respjson.Field
		Status                respjson.Field
		Strategy              respjson.Field
		TargetDeploymentID    respjson.Field
		CompletedAt           respjson.Field
		CurrentStep           respjson.Field
		CurrentTrafficPercent respjson.Field
		Etag                  respjson.Field
		PauseInfo             respjson.Field
		StartedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Rollout) RawJSON() string { return r.JSON.raw }
func (r *Rollout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Output only. High-level rollout lifecycle state.
type RolloutState string

const (
	RolloutStateRolloutStateRunning      RolloutState = "ROLLOUT_STATE_RUNNING"
	RolloutStateRolloutStatePaused       RolloutState = "ROLLOUT_STATE_PAUSED"
	RolloutStateRolloutStateStabilizing  RolloutState = "ROLLOUT_STATE_STABILIZING"
	RolloutStateRolloutStateCompleted    RolloutState = "ROLLOUT_STATE_COMPLETED"
	RolloutStateRolloutStatePending      RolloutState = "ROLLOUT_STATE_PENDING"
	RolloutStateRolloutStateSystemPaused RolloutState = "ROLLOUT_STATE_SYSTEM_PAUSED"
	RolloutStateRolloutStateCancelling   RolloutState = "ROLLOUT_STATE_CANCELLING"
	RolloutStateRolloutStateCanceled     RolloutState = "ROLLOUT_STATE_CANCELED"
	RolloutStateRolloutStatePausing      RolloutState = "ROLLOUT_STATE_PAUSING"
)

// Output only. Rollout strategy selected at creation.
type RolloutStrategy string

const (
	RolloutStrategyRolloutStrategyTypeRolling   RolloutStrategy = "ROLLOUT_STRATEGY_TYPE_ROLLING"
	RolloutStrategyRolloutStrategyTypeCanary    RolloutStrategy = "ROLLOUT_STRATEGY_TYPE_CANARY"
	RolloutStrategyRolloutStrategyTypeBlueGreen RolloutStrategy = "ROLLOUT_STRATEGY_TYPE_BLUE_GREEN"
)

// Structured reason a rollout stopped progressing.
type RolloutCondition struct {
	// Step index where the condition arose. Step 0 serializes explicitly.
	AtStep int64 `json:"atStep"`
	// Category that classifies why the rollout stopped.
	//
	// Any of "ROLLOUT_FAILURE_CATEGORY_METRIC_REGRESSION",
	// "ROLLOUT_FAILURE_CATEGORY_METRICS_UNAVAILABLE",
	// "ROLLOUT_FAILURE_CATEGORY_TARGET_NOT_READY",
	// "ROLLOUT_FAILURE_CATEGORY_SOURCE_NOT_DRAINED",
	// "ROLLOUT_FAILURE_CATEGORY_HEALTH_REGRESSION",
	// "ROLLOUT_FAILURE_CATEGORY_CAPACITY_EXHAUSTED",
	// "ROLLOUT_FAILURE_CATEGORY_ROUTING_ERROR",
	// "ROLLOUT_FAILURE_CATEGORY_DEPENDENCY_OUTAGE",
	// "ROLLOUT_FAILURE_CATEGORY_ABORTED_BY_OPERATOR",
	// "ROLLOUT_FAILURE_CATEGORY_INTERNAL",
	// "ROLLOUT_FAILURE_CATEGORY_POLICY_INFEASIBLE",
	// "ROLLOUT_FAILURE_CATEGORY_UNDER_SERVED",
	// "ROLLOUT_FAILURE_CATEGORY_ENTITLEMENT_LAPSED".
	Category RolloutConditionCategory `json:"category"`
	// Human-readable explanation for the condition.
	Message string `json:"message"`
	// Metrics observed at the failing gate, enriched with their criteria. Unmeasured
	// rules appear as synthesized rows with verdict METRIC_VERDICT_UNAVAILABLE and no
	// measured values.
	Metrics []MetricResult `json:"metrics"`
	// Timestamp when the condition was observed.
	ObservedAt time.Time `json:"observedAt" format:"date-time"`
	// Informational condition type. `CapacityLimited` means the current step advanced
	// partially because full capacity was not placeable.
	//
	// Any of "CapacityLimited".
	Type RolloutConditionType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtStep      respjson.Field
		Category    respjson.Field
		Message     respjson.Field
		Metrics     respjson.Field
		ObservedAt  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RolloutCondition) RawJSON() string { return r.JSON.raw }
func (r *RolloutCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Category that classifies why the rollout stopped.
type RolloutConditionCategory string

const (
	RolloutConditionCategoryRolloutFailureCategoryMetricRegression   RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_METRIC_REGRESSION"
	RolloutConditionCategoryRolloutFailureCategoryMetricsUnavailable RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_METRICS_UNAVAILABLE"
	RolloutConditionCategoryRolloutFailureCategoryTargetNotReady     RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_TARGET_NOT_READY"
	RolloutConditionCategoryRolloutFailureCategorySourceNotDrained   RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_SOURCE_NOT_DRAINED"
	RolloutConditionCategoryRolloutFailureCategoryHealthRegression   RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_HEALTH_REGRESSION"
	RolloutConditionCategoryRolloutFailureCategoryCapacityExhausted  RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_CAPACITY_EXHAUSTED"
	RolloutConditionCategoryRolloutFailureCategoryRoutingError       RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_ROUTING_ERROR"
	RolloutConditionCategoryRolloutFailureCategoryDependencyOutage   RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_DEPENDENCY_OUTAGE"
	RolloutConditionCategoryRolloutFailureCategoryAbortedByOperator  RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_ABORTED_BY_OPERATOR"
	RolloutConditionCategoryRolloutFailureCategoryInternal           RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_INTERNAL"
	RolloutConditionCategoryRolloutFailureCategoryPolicyInfeasible   RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_POLICY_INFEASIBLE"
	RolloutConditionCategoryRolloutFailureCategoryUnderServed        RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_UNDER_SERVED"
	RolloutConditionCategoryRolloutFailureCategoryEntitlementLapsed  RolloutConditionCategory = "ROLLOUT_FAILURE_CATEGORY_ENTITLEMENT_LAPSED"
)

// Informational condition type. `CapacityLimited` means the current step advanced
// partially because full capacity was not placeable.
type RolloutConditionType string

const (
	RolloutConditionTypeCapacityLimited RolloutConditionType = "CapacityLimited"
)

// Completed create-form state — the caller's spec with defaulted values filled in,
// the steps the rollout is expected to walk, and the capacity context the defaults
// were computed from. Display only.
type RolloutDefaultsPreview struct {
	// Source deployment replica count the defaults were computed from. Zero is a real
	// value.
	SourceReplicas int64 `json:"sourceReplicas" api:"required"`
	// Strategy, metric gates, timing, and cleanup policy for shifting traffic between
	// two deployments under one endpoint.
	Spec RolloutDefaultsPreviewSpec `json:"spec" api:"required"`
	// Target deployment autoscaling maximum replica count. Zero is a real value.
	TargetMaxReplicas int64 `json:"targetMaxReplicas" api:"required"`
	// Target deployment autoscaling minimum replica count. Zero is a real value.
	TargetMinReplicas int64 `json:"targetMinReplicas" api:"required"`
	// Target deployment replica count the defaults were computed from. Zero is a real
	// value.
	TargetReplicas int64 `json:"targetReplicas" api:"required"`
	// Findings to surface next to the form when a later gate will refuse the spec or a
	// standing guarantee is lost. An empty list means the shown values are safe to
	// submit as-is; render message for unrecognized codes.
	Warnings []RolloutDefaultsPreviewWarning `json:"warnings" api:"required"`
	// Steps the rollout is expected to walk when the caller leaves steps unset.
	// Display only. Empty when the caller supplied steps or no ladder applies.
	EstimatedEffectiveSteps []RolloutStep `json:"estimatedEffectiveSteps"`
	// Percentage of the pair's traffic currently reaching the target, the floor the
	// suggested steps start above. Unset when not a frozen pair or unknown; 0 is a
	// real measurement.
	EstimatedSeedPercent int64 `json:"estimatedSeedPercent"`
	// True when both deployments stand in the endpoint traffic split, so the rollout
	// resumes from the current split rather than from zero. See warnings for standing
	// split shapes that StartRollout will still reject.
	FrozenPair bool `json:"frozenPair"`
	// Expected autoscaling maximum replicas for the completed target; unset while the
	// final target replicas cannot be resolved.
	LandingMaxReplicas int64 `json:"landingMaxReplicas"`
	// Expected autoscaling minimum replicas for the completed target; unset while the
	// final target replicas cannot be resolved.
	LandingMinReplicas int64 `json:"landingMinReplicas"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceReplicas          respjson.Field
		Spec                    respjson.Field
		TargetMaxReplicas       respjson.Field
		TargetMinReplicas       respjson.Field
		TargetReplicas          respjson.Field
		Warnings                respjson.Field
		EstimatedEffectiveSteps respjson.Field
		EstimatedSeedPercent    respjson.Field
		FrozenPair              respjson.Field
		LandingMaxReplicas      respjson.Field
		LandingMinReplicas      respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RolloutDefaultsPreview) RawJSON() string { return r.JSON.raw }
func (r *RolloutDefaultsPreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Strategy, metric gates, timing, and cleanup policy for shifting traffic between
// two deployments under one endpoint.
type RolloutDefaultsPreviewSpec struct {
	// Deployment that traffic shifts away from.
	SourceDeploymentID string `json:"sourceDeploymentId" api:"required"`
	// Deployment that traffic shifts toward.
	TargetDeploymentID string `json:"targetDeploymentId" api:"required"`
	// Blue-green strategy configuration for a single cutover to the target deployment.
	BlueGreen BlueGreenConfig `json:"blueGreen"`
	// Canary strategy configuration for gradual traffic progression. An empty config
	// uses the default 5, 25, 50, 100 percent ladder; over a frozen traffic-split pair
	// left by cancel, the default ladder is derived at start from the pair's current
	// served share so it begins above it.
	Canary CanaryConfig `json:"canary"`
	// Optional final replica count for the source deployment. Defaults to 0, which
	// drains and stops the source.
	FinalSourceReplicas int64 `json:"finalSourceReplicas"`
	// Optional target replica floor at completion. Must be at least 1 when set;
	// defaults to the source deployment's replica count at create time, or to the
	// source and target deployments' combined replica count when both already stand in
	// the endpoint traffic split after a cancel. The completed target's autoscaling
	// max lands at the landing ceiling, max(this value, the source max, the target's
	// own max); the rollout may lift the target max at first wake, at the first step
	// that needs it, or at completion unless an operator changes max mid-run. The
	// lifted ceiling remains after completion, and PreviewRolloutDefaults reports a
	// coming lift as ROLLOUT_WILL_RAISE_TARGET_MAX. A pre-existing target whose own
	// autoscaling min is higher keeps that floor, reported as
	// FINAL_BELOW_INHERITED_MIN. A target that starts stopped lands exactly at this
	// value; if the source min was higher, PreviewRolloutDefaults reports
	// FINAL_BELOW_SOURCE_MIN.
	FinalTargetReplicas int64 `json:"finalTargetReplicas"`
	// Optional metric gates evaluated after each step's soak. Canary only; rejected on
	// rolling and blue-green rollouts.
	Metrics []MetricRule `json:"metrics"`
	// Rolling strategy configuration for capacity-preserving batches that ramp target
	// replicas up while draining source replicas.
	Rolling RollingConfig `json:"rolling"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceDeploymentID  respjson.Field
		TargetDeploymentID  respjson.Field
		BlueGreen           respjson.Field
		Canary              respjson.Field
		FinalSourceReplicas respjson.Field
		FinalTargetReplicas respjson.Field
		Metrics             respjson.Field
		Rolling             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RolloutDefaultsPreviewSpec) RawJSON() string { return r.JSON.raw }
func (r *RolloutDefaultsPreviewSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A non-blocking finding attached to a rollout defaults preview; expected
// end-state facts are structured fields on RolloutDefaultsPreview.
type RolloutDefaultsPreviewWarning struct {
	// Machine-readable warning code. Current vocabulary is START_WILL_REJECT,
	// FINAL_BELOW_SOURCE_MIN, and FIRST_STEP_AT_SEED; render message for unrecognized
	// codes.
	Code string `json:"code" api:"required"`
	// Plain-language description of the finding, safe to show users as-is.
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RolloutDefaultsPreviewWarning) RawJSON() string { return r.JSON.raw }
func (r *RolloutDefaultsPreviewWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Derived runtime progress for a rollout.
type RolloutStatus struct {
	// Per-step rollout execution summaries.
	Steps []RolloutStepStatus `json:"steps" api:"required"`
	// Total number of steps in the rollout progression. Always serializes when status
	// is present.
	TotalSteps int64 `json:"totalSteps" api:"required"`
	// Structured reason a rollout stopped progressing.
	Condition RolloutCondition `json:"condition"`
	// Informational conditions that describe the rollout's current state. Omitted when
	// empty; clients should treat an absent key as an empty list.
	Conditions []RolloutCondition `json:"conditions"`
	// Timestamp of the most recent progress update.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Steps       respjson.Field
		TotalSteps  respjson.Field
		Condition   respjson.Field
		Conditions  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RolloutStatus) RawJSON() string { return r.JSON.raw }
func (r *RolloutStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One stage of a canary rollout progression.
type RolloutStep struct {
	// Required percentage of traffic on the target deployment for this step.
	Traffic int64 `json:"traffic" api:"required"`
	// Optional explicit target replica count for this step.
	Replicas int64 `json:"replicas"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Traffic     respjson.Field
		Replicas    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RolloutStep) RawJSON() string { return r.JSON.raw }
func (r *RolloutStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RolloutStep to a RolloutStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RolloutStepParam.Overrides()
func (r RolloutStep) ToParam() RolloutStepParam {
	return param.Override[RolloutStepParam](json.RawMessage(r.RawJSON()))
}

// One stage of a canary rollout progression.
//
// The property Traffic is required.
type RolloutStepParam struct {
	// Required percentage of traffic on the target deployment for this step.
	Traffic int64 `json:"traffic" api:"required"`
	// Optional explicit target replica count for this step.
	Replicas param.Opt[int64] `json:"replicas,omitzero"`
	paramObj
}

func (r RolloutStepParam) MarshalJSON() (data []byte, err error) {
	type shadow RolloutStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RolloutStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collapsed execution state for one rollout step.
type RolloutStepStatus struct {
	// Timestamp when this step finished, was skipped over, or the rollout ended on it.
	// Unset while in progress.
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Failure reason set only when this step failed.
	FailureReason string `json:"failureReason"`
	// Metric gate results for this step, enriched with criteria and verdict.
	// Unmeasured rules appear as synthesized rows with verdict
	// METRIC_VERDICT_UNAVAILABLE and no measured values.
	Metrics []MetricResult `json:"metrics"`
	// Timestamp when this step's first sub-step ran. Unset for steps no sub-step
	// reached.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// Outcome of this step. Finished steps are PASSED, the live step mirrors the
	// rollout state, skipped-over steps are SKIPPED, and unreached steps are PENDING.
	//
	// Any of "ROLLOUT_STEP_STATE_PENDING", "ROLLOUT_STEP_STATE_RUNNING",
	// "ROLLOUT_STEP_STATE_PASSED", "ROLLOUT_STEP_STATE_FAILED",
	// "ROLLOUT_STEP_STATE_PAUSED", "ROLLOUT_STEP_STATE_CANCELED",
	// "ROLLOUT_STEP_STATE_SKIPPED".
	State RolloutStepStatusState `json:"state"`
	// Index of this step in the rollout progression. Step 0 serializes explicitly.
	StepIndex int64 `json:"stepIndex"`
	// Target traffic percentage configured for this step. Always serializes for
	// recorded steps.
	TargetTrafficPercent int64 `json:"targetTrafficPercent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt          respjson.Field
		FailureReason        respjson.Field
		Metrics              respjson.Field
		StartedAt            respjson.Field
		State                respjson.Field
		StepIndex            respjson.Field
		TargetTrafficPercent respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RolloutStepStatus) RawJSON() string { return r.JSON.raw }
func (r *RolloutStepStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Outcome of this step. Finished steps are PASSED, the live step mirrors the
// rollout state, skipped-over steps are SKIPPED, and unreached steps are PENDING.
type RolloutStepStatusState string

const (
	RolloutStepStatusStateRolloutStepStatePending  RolloutStepStatusState = "ROLLOUT_STEP_STATE_PENDING"
	RolloutStepStatusStateRolloutStepStateRunning  RolloutStepStatusState = "ROLLOUT_STEP_STATE_RUNNING"
	RolloutStepStatusStateRolloutStepStatePassed   RolloutStepStatusState = "ROLLOUT_STEP_STATE_PASSED"
	RolloutStepStatusStateRolloutStepStateFailed   RolloutStepStatusState = "ROLLOUT_STEP_STATE_FAILED"
	RolloutStepStatusStateRolloutStepStatePaused   RolloutStepStatusState = "ROLLOUT_STEP_STATE_PAUSED"
	RolloutStepStatusStateRolloutStepStateCanceled RolloutStepStatusState = "ROLLOUT_STEP_STATE_CANCELED"
	RolloutStepStatusStateRolloutStepStateSkipped  RolloutStepStatusState = "ROLLOUT_STEP_STATE_SKIPPED"
)

// Threshold criteria that fail when the target metric violates the configured
// bound.
type ThresholdCheck struct {
	// Required comparison operator applied to the target metric value.
	//
	// Any of "THRESHOLD_OPERATOR_GT", "THRESHOLD_OPERATOR_GTE",
	// "THRESHOLD_OPERATOR_LT", "THRESHOLD_OPERATOR_LTE".
	Operator ThresholdCheckOperator `json:"operator" api:"required"`
	// Finite threshold value. Interpreted in the metric's unit: router_error_rate is a
	// ratio in [0, 1], router_latency is milliseconds, and inflight_requests is
	// in-flight requests per ready replica averaged over the rule window. Thresholds
	// that no achievable value could pass, or that every achievable value passes, are
	// rejected at create.
	//
	// Omitting this value is read as 0. Set 0 explicitly for the strictest threshold:
	// nothing at all is tolerated.
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThresholdCheck) RawJSON() string { return r.JSON.raw }
func (r *ThresholdCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ThresholdCheck to a ThresholdCheckParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ThresholdCheckParam.Overrides()
func (r ThresholdCheck) ToParam() ThresholdCheckParam {
	return param.Override[ThresholdCheckParam](json.RawMessage(r.RawJSON()))
}

// Required comparison operator applied to the target metric value.
type ThresholdCheckOperator string

const (
	ThresholdCheckOperatorThresholdOperatorGt  ThresholdCheckOperator = "THRESHOLD_OPERATOR_GT"
	ThresholdCheckOperatorThresholdOperatorGte ThresholdCheckOperator = "THRESHOLD_OPERATOR_GTE"
	ThresholdCheckOperatorThresholdOperatorLt  ThresholdCheckOperator = "THRESHOLD_OPERATOR_LT"
	ThresholdCheckOperatorThresholdOperatorLte ThresholdCheckOperator = "THRESHOLD_OPERATOR_LTE"
)

// Threshold criteria that fail when the target metric violates the configured
// bound.
//
// The property Operator is required.
type ThresholdCheckParam struct {
	// Required comparison operator applied to the target metric value.
	//
	// Any of "THRESHOLD_OPERATOR_GT", "THRESHOLD_OPERATOR_GTE",
	// "THRESHOLD_OPERATOR_LT", "THRESHOLD_OPERATOR_LTE".
	Operator ThresholdCheckOperator `json:"operator,omitzero" api:"required"`
	// Finite threshold value. Interpreted in the metric's unit: router_error_rate is a
	// ratio in [0, 1], router_latency is milliseconds, and inflight_requests is
	// in-flight requests per ready replica averaged over the rule window. Thresholds
	// that no achievable value could pass, or that every achievable value passes, are
	// rejected at create.
	//
	// Omitting this value is read as 0. Set 0 explicitly for the strictest threshold:
	// nothing at all is tolerated.
	Value param.Opt[float64] `json:"value,omitzero"`
	paramObj
}

func (r ThresholdCheckParam) MarshalJSON() (data []byte, err error) {
	type shadow ThresholdCheckParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ThresholdCheckParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Empty response returned after a successful delete operation.
type BetaEndpointRolloutDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointRolloutDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointRolloutDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointRolloutNewParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Deployment that traffic shifts away from.
	SourceDeploymentID string `json:"sourceDeploymentId" api:"required"`
	// Deployment that traffic shifts toward.
	TargetDeploymentID string `json:"targetDeploymentId" api:"required"`
	// Optional final replica count for the source deployment. Defaults to 0, which
	// drains and stops the source.
	FinalSourceReplicas param.Opt[int64] `json:"finalSourceReplicas,omitzero"`
	// Optional target replica floor at completion. Must be at least 1 when set;
	// defaults to the source deployment's replica count at create time, or to the
	// source and target deployments' combined replica count when both already stand in
	// the endpoint traffic split after a cancel. The completed target's autoscaling
	// max lands at the landing ceiling, max(this value, the source max, the target's
	// own max); the rollout may lift the target max at first wake, at the first step
	// that needs it, or at completion unless an operator changes max mid-run. The
	// lifted ceiling remains after completion, and PreviewRolloutDefaults reports a
	// coming lift as ROLLOUT_WILL_RAISE_TARGET_MAX. A pre-existing target whose own
	// autoscaling min is higher keeps that floor, reported as
	// FINAL_BELOW_INHERITED_MIN. A target that starts stopped lands exactly at this
	// value; if the source min was higher, PreviewRolloutDefaults reports
	// FINAL_BELOW_SOURCE_MIN.
	FinalTargetReplicas param.Opt[int64] `json:"finalTargetReplicas,omitzero"`
	// Blue-green strategy configuration for a single cutover to the target deployment.
	BlueGreen BlueGreenConfigParam `json:"blueGreen,omitzero"`
	// Canary strategy configuration for gradual traffic progression. An empty config
	// uses the default 5, 25, 50, 100 percent ladder; over a frozen traffic-split pair
	// left by cancel, the default ladder is derived at start from the pair's current
	// served share so it begins above it.
	Canary CanaryConfigParam `json:"canary,omitzero"`
	// Optional metric gates evaluated after each step's soak. Canary only; rejected on
	// rolling and blue-green rollouts.
	Metrics []MetricRuleParam `json:"metrics,omitzero"`
	// Rolling strategy configuration for capacity-preserving batches that ramp target
	// replicas up while draining source replicas.
	Rolling RollingConfigParam `json:"rolling,omitzero"`
	paramObj
}

func (r BetaEndpointRolloutNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointRolloutNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointRolloutNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointRolloutGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	paramObj
}

type BetaEndpointRolloutListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous rollout list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of rollouts to return. Max 500, defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Narrow results to active or terminal rollouts. Omit to list all rollouts.
	//
	// Any of "ROLLOUT_FILTER_ACTIVE", "ROLLOUT_FILTER_TERMINAL".
	Filter BetaEndpointRolloutListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointRolloutListParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointRolloutListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Narrow results to active or terminal rollouts. Omit to list all rollouts.
type BetaEndpointRolloutListParamsFilter string

const (
	BetaEndpointRolloutListParamsFilterRolloutFilterActive   BetaEndpointRolloutListParamsFilter = "ROLLOUT_FILTER_ACTIVE"
	BetaEndpointRolloutListParamsFilterRolloutFilterTerminal BetaEndpointRolloutListParamsFilter = "ROLLOUT_FILTER_TERMINAL"
)

type BetaEndpointRolloutDeleteParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Etag for optimistic concurrency.
	Etag param.Opt[string] `query:"etag,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointRolloutDeleteParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointRolloutDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointRolloutCancelParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Required human-readable reason recorded in the rollout audit trail.
	Reason string `json:"reason" api:"required"`
	// Optional etag for optimistic concurrency.
	Etag param.Opt[string] `json:"etag,omitzero"`
	// Optional cancel behavior. Absent defaults to freeze, which preserves the current
	// traffic split. Revert is removed and rejected with FAILED_PRECONDITION; cancel
	// with freeze, then run a reverse rollout back to the source.
	//
	// Any of "CANCEL_DISPOSITION_FREEZE", "CANCEL_DISPOSITION_REVERT".
	Disposition BetaEndpointRolloutCancelParamsDisposition `json:"disposition,omitzero"`
	paramObj
}

func (r BetaEndpointRolloutCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointRolloutCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointRolloutCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional cancel behavior. Absent defaults to freeze, which preserves the current
// traffic split. Revert is removed and rejected with FAILED_PRECONDITION; cancel
// with freeze, then run a reverse rollout back to the source.
type BetaEndpointRolloutCancelParamsDisposition string

const (
	BetaEndpointRolloutCancelParamsDispositionCancelDispositionFreeze BetaEndpointRolloutCancelParamsDisposition = "CANCEL_DISPOSITION_FREEZE"
	BetaEndpointRolloutCancelParamsDispositionCancelDispositionRevert BetaEndpointRolloutCancelParamsDisposition = "CANCEL_DISPOSITION_REVERT"
)

type BetaEndpointRolloutPauseParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Optional etag for optimistic concurrency.
	Etag param.Opt[string] `json:"etag,omitzero"`
	// Optional human-readable reason recorded on the rollout pause metadata.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r BetaEndpointRolloutPauseParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointRolloutPauseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointRolloutPauseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointRolloutPreviewDefaultsParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Deployment that traffic shifts away from.
	SourceDeploymentID string `json:"sourceDeploymentId" api:"required"`
	// Deployment that traffic shifts toward.
	TargetDeploymentID string `json:"targetDeploymentId" api:"required"`
	// Optional final replica count for the source deployment. Defaults to 0, which
	// drains and stops the source.
	FinalSourceReplicas param.Opt[int64] `json:"finalSourceReplicas,omitzero"`
	// Optional target replica floor at completion. Must be at least 1 when set;
	// defaults to the source deployment's replica count at create time, or to the
	// source and target deployments' combined replica count when both already stand in
	// the endpoint traffic split after a cancel. The completed target's autoscaling
	// max lands at the landing ceiling, max(this value, the source max, the target's
	// own max); the rollout may lift the target max at first wake, at the first step
	// that needs it, or at completion unless an operator changes max mid-run. The
	// lifted ceiling remains after completion, and PreviewRolloutDefaults reports a
	// coming lift as ROLLOUT_WILL_RAISE_TARGET_MAX. A pre-existing target whose own
	// autoscaling min is higher keeps that floor, reported as
	// FINAL_BELOW_INHERITED_MIN. A target that starts stopped lands exactly at this
	// value; if the source min was higher, PreviewRolloutDefaults reports
	// FINAL_BELOW_SOURCE_MIN.
	FinalTargetReplicas param.Opt[int64] `json:"finalTargetReplicas,omitzero"`
	// Blue-green strategy configuration for a single cutover to the target deployment.
	BlueGreen BlueGreenConfigParam `json:"blueGreen,omitzero"`
	// Canary strategy configuration for gradual traffic progression. An empty config
	// uses the default 5, 25, 50, 100 percent ladder; over a frozen traffic-split pair
	// left by cancel, the default ladder is derived at start from the pair's current
	// served share so it begins above it.
	Canary CanaryConfigParam `json:"canary,omitzero"`
	// Optional metric gates evaluated after each step's soak. Canary only; rejected on
	// rolling and blue-green rollouts.
	Metrics []MetricRuleParam `json:"metrics,omitzero"`
	// Rolling strategy configuration for capacity-preserving batches that ramp target
	// replicas up while draining source replicas.
	Rolling RollingConfigParam `json:"rolling,omitzero"`
	paramObj
}

func (r BetaEndpointRolloutPreviewDefaultsParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointRolloutPreviewDefaultsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointRolloutPreviewDefaultsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointRolloutPromoteParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Optional etag for optimistic concurrency.
	Etag param.Opt[string] `json:"etag,omitzero"`
	paramObj
}

func (r BetaEndpointRolloutPromoteParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointRolloutPromoteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointRolloutPromoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointRolloutResumeParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Optional etag for optimistic concurrency.
	Etag param.Opt[string] `json:"etag,omitzero"`
	paramObj
}

func (r BetaEndpointRolloutResumeParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointRolloutResumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointRolloutResumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointRolloutStartParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	paramObj
}
