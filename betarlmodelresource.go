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

// BetaRlModelResourceService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaRlModelResourceService] method instead.
type BetaRlModelResourceService struct {
	Options []option.RequestOption
}

// NewBetaRlModelResourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaRlModelResourceService(opts ...option.RequestOption) (r BetaRlModelResourceService) {
	r = BetaRlModelResourceService{}
	r.Options = opts
	return
}

// Provisions a standalone model resource that training sessions can attach to.
func (r *BetaRlModelResourceService) New(ctx context.Context, body BetaRlModelResourceNewParams, opts ...option.RequestOption) (res *ModelResources, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rl/model-resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets a model resource by its ID and returns its details.
func (r *BetaRlModelResourceService) Get(ctx context.Context, modelResourcesID string, opts ...option.RequestOption) (res *ModelResources, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelResourcesID == "" {
		err = errors.New("missing required model_resources_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/model-resources/%s", modelResourcesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the caller's model resources.
func (r *BetaRlModelResourceService) List(ctx context.Context, query BetaRlModelResourceListParams, opts ...option.RequestOption) (res *ModelResourcesListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rl/model-resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Estimates a model resource's on-demand hourly price without creating it.
func (r *BetaRlModelResourceService) EstimateCost(ctx context.Context, body BetaRlModelResourceEstimateCostParams, opts ...option.RequestOption) (res *ModelResourcesEstimateCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rl/model-resources/estimate-cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stops every session attached to the resource and tears down its GPU pods.
func (r *BetaRlModelResourceService) Stop(ctx context.Context, modelResourcesID string, body BetaRlModelResourceStopParams, opts ...option.RequestOption) (res *ModelResources, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelResourcesID == "" {
		err = errors.New("missing required model_resources_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/model-resources/%s/stop", modelResourcesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Allocated GPU resources that training sessions attach to
type ModelResources struct {
	// Unique identifier for the model resource
	ID string `json:"id" api:"required"`
	// Base model the resource is provisioned for
	BaseModel string `json:"base_model" api:"required"`
	// Compute layout provisioned for the resource.
	ComputeConfig ModelResourcesComputeConfig `json:"compute_config" api:"required"`
	// Timestamp when the model resource was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ID of the user who created the model resource
	CreatedBy string `json:"created_by" api:"required"`
	// Whether the resource hosts LoRA sessions or a full-weight session
	LoraEnabled bool `json:"lora_enabled" api:"required"`
	// Optimizer configuration for this resource.
	OptimizerConfig OptimizerConfig `json:"optimizer_config" api:"required"`
	// Lifecycle status of the model resource
	//
	// Any of "MODEL_RESOURCES_STATUS_PENDING", "MODEL_RESOURCES_STATUS_CREATING",
	// "MODEL_RESOURCES_STATUS_READY", "MODEL_RESOURCES_STATUS_ERROR",
	// "MODEL_RESOURCES_STATUS_STOPPED", "MODEL_RESOURCES_STATUS_STOPPING".
	Status ModelResourcesStatus `json:"status" api:"required"`
	// Timestamp when the model resource was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Structured detail for the model resource's current error. Set when the resource
	// is in an error state.
	Error ModelResourcesError `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BaseModel       respjson.Field
		ComputeConfig   respjson.Field
		CreatedAt       respjson.Field
		CreatedBy       respjson.Field
		LoraEnabled     respjson.Field
		OptimizerConfig respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Error           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelResources) RawJSON() string { return r.JSON.raw }
func (r *ModelResources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compute layout provisioned for the resource.
type ModelResourcesComputeConfig struct {
	// Number of generator replicas. 0 means the resource runs the trainer only, with
	// no generator.
	NumGeneratorReplicas int64 `json:"num_generator_replicas" api:"required"`
	// GPU type selected for this resource.
	//
	// Any of "H100-80GB", "B200-SXM".
	GPUType string `json:"gpu_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumGeneratorReplicas respjson.Field
		GPUType              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelResourcesComputeConfig) RawJSON() string { return r.JSON.raw }
func (r *ModelResourcesComputeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured detail for the model resource's current error
type ModelResourcesError struct {
	// Finite machine-readable reason code for UI branching
	//
	// Any of "MODEL_RESOURCES_ERROR_CODE_CAPACITY_WAIT_TIMEOUT",
	// "MODEL_RESOURCES_ERROR_CODE_PROVISIONING_FAILED".
	Code ModelResourcesErrorCode `json:"code" api:"required"`
	// User-safe human-readable detail for the current status
	Message string `json:"message" api:"required"`
	// Timestamp when this error was reported
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		OccurredAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelResourcesError) RawJSON() string { return r.JSON.raw }
func (r *ModelResourcesError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Finite machine-readable model resource lifecycle error code
type ModelResourcesErrorCode string

const (
	ModelResourcesErrorCodeModelResourcesErrorCodeCapacityWaitTimeout ModelResourcesErrorCode = "MODEL_RESOURCES_ERROR_CODE_CAPACITY_WAIT_TIMEOUT"
	ModelResourcesErrorCodeModelResourcesErrorCodeProvisioningFailed  ModelResourcesErrorCode = "MODEL_RESOURCES_ERROR_CODE_PROVISIONING_FAILED"
)

type ModelResourcesEstimateCostResponse struct {
	// ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	// Estimated on-demand price per hour in the currency's major unit.
	PricePerHour float64 `json:"price_per_hour" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency     respjson.Field
		PricePerHour respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelResourcesEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelResourcesEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of model resources
type ModelResourcesListResponse struct {
	// List of model resources
	Data []ModelResources `json:"data" api:"required"`
	// Pagination metadata
	Meta ModelResourcesListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelResourcesListResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelResourcesListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type ModelResourcesListResponseMeta struct {
	// Whether more items exist beyond this page
	HasMore bool `json:"has_more"`
	// Maximum number of items returned per page
	Limit int64 `json:"limit"`
	// Cursor to use as the 'after' parameter for the next page. Empty when has_more is
	// false.
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Limit       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelResourcesListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ModelResourcesListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle status of a model resource
type ModelResourcesStatus string

const (
	ModelResourcesStatusModelResourcesStatusPending  ModelResourcesStatus = "MODEL_RESOURCES_STATUS_PENDING"
	ModelResourcesStatusModelResourcesStatusCreating ModelResourcesStatus = "MODEL_RESOURCES_STATUS_CREATING"
	ModelResourcesStatusModelResourcesStatusReady    ModelResourcesStatus = "MODEL_RESOURCES_STATUS_READY"
	ModelResourcesStatusModelResourcesStatusError    ModelResourcesStatus = "MODEL_RESOURCES_STATUS_ERROR"
	ModelResourcesStatusModelResourcesStatusStopped  ModelResourcesStatus = "MODEL_RESOURCES_STATUS_STOPPED"
	ModelResourcesStatusModelResourcesStatusStopping ModelResourcesStatus = "MODEL_RESOURCES_STATUS_STOPPING"
)

type MuonScalingStrategy string

const (
	MuonScalingStrategyMuonScalingStrategyUnspecified MuonScalingStrategy = "MUON_SCALING_STRATEGY_UNSPECIFIED"
	MuonScalingStrategyMuonScalingStrategyMatchAdam   MuonScalingStrategy = "MUON_SCALING_STRATEGY_MATCH_ADAM"
	MuonScalingStrategyMuonScalingStrategyOriginal    MuonScalingStrategy = "MUON_SCALING_STRATEGY_ORIGINAL"
)

// Optimizer configuration
type OptimizerConfig struct {
	// Use the Adam optimizer.
	Adam OptimizerConfigAdam `json:"adam"`
	// Use the Muon optimizer.
	Muon OptimizerConfigMuon `json:"muon"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Adam        respjson.Field
		Muon        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizerConfig) RawJSON() string { return r.JSON.raw }
func (r *OptimizerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OptimizerConfig to a OptimizerConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OptimizerConfigParam.Overrides()
func (r OptimizerConfig) ToParam() OptimizerConfigParam {
	return param.Override[OptimizerConfigParam](json.RawMessage(r.RawJSON()))
}

// Use the Adam optimizer.
type OptimizerConfigAdam struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizerConfigAdam) RawJSON() string { return r.JSON.raw }
func (r *OptimizerConfigAdam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use the Muon optimizer.
type OptimizerConfigMuon struct {
	// Scaling strategy for the Muon optimizer.
	//
	// Any of "MUON_SCALING_STRATEGY_UNSPECIFIED", "MUON_SCALING_STRATEGY_MATCH_ADAM",
	// "MUON_SCALING_STRATEGY_ORIGINAL".
	ScalingStrategy MuonScalingStrategy `json:"scaling_strategy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScalingStrategy respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimizerConfigMuon) RawJSON() string { return r.JSON.raw }
func (r *OptimizerConfigMuon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optimizer configuration
type OptimizerConfigParam struct {
	// Use the Adam optimizer.
	Adam OptimizerConfigAdamParam `json:"adam,omitzero"`
	// Use the Muon optimizer.
	Muon OptimizerConfigMuonParam `json:"muon,omitzero"`
	paramObj
}

func (r OptimizerConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow OptimizerConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizerConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use the Adam optimizer.
type OptimizerConfigAdamParam struct {
	paramObj
}

func (r OptimizerConfigAdamParam) MarshalJSON() (data []byte, err error) {
	type shadow OptimizerConfigAdamParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizerConfigAdamParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use the Muon optimizer.
type OptimizerConfigMuonParam struct {
	// Scaling strategy for the Muon optimizer.
	//
	// Any of "MUON_SCALING_STRATEGY_UNSPECIFIED", "MUON_SCALING_STRATEGY_MATCH_ADAM",
	// "MUON_SCALING_STRATEGY_ORIGINAL".
	ScalingStrategy MuonScalingStrategy `json:"scaling_strategy,omitzero"`
	paramObj
}

func (r OptimizerConfigMuonParam) MarshalJSON() (data []byte, err error) {
	type shadow OptimizerConfigMuonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OptimizerConfigMuonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlModelResourceNewParams struct {
	// Base model to provision the resource for
	BaseModel string `json:"base_model" api:"required"`
	// Whether the resource hosts LoRA sessions or a single full-weight session
	LoraEnabled param.Opt[bool] `json:"lora_enabled,omitzero"`
	// Compute layout to provision.
	ComputeConfig BetaRlModelResourceNewParamsComputeConfig `json:"compute_config,omitzero"`
	// Optimizer configuration for this resource.
	OptimizerConfig OptimizerConfigParam `json:"optimizer_config,omitzero"`
	paramObj
}

func (r BetaRlModelResourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlModelResourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlModelResourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compute layout to provision.
type BetaRlModelResourceNewParamsComputeConfig struct {
	// Number of generator replicas. 0 runs the trainer only, with no generator.
	NumGeneratorReplicas param.Opt[int64] `json:"num_generator_replicas,omitzero"`
	// GPU type to provision. Omit to use the model's default GPU type.
	//
	// Any of "H100-80GB", "B200-SXM".
	GPUType string `json:"gpu_type,omitzero"`
	paramObj
}

func (r BetaRlModelResourceNewParamsComputeConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlModelResourceNewParamsComputeConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlModelResourceNewParamsComputeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaRlModelResourceNewParamsComputeConfig](
		"gpu_type", "H100-80GB", "B200-SXM",
	)
}

type BetaRlModelResourceListParams struct {
	// Cursor for pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Filter resources in the current project by the creator ID. Pass "me" to show
	// resources you created.
	CreatedBy param.Opt[string] `query:"created_by,omitzero" json:"-"`
	// Maximum number of resources to return (1-100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Status filters. When omitted, resources in any status are returned.
	//
	// Any of "MODEL_RESOURCES_STATUS_PENDING", "MODEL_RESOURCES_STATUS_CREATING",
	// "MODEL_RESOURCES_STATUS_READY", "MODEL_RESOURCES_STATUS_ERROR",
	// "MODEL_RESOURCES_STATUS_STOPPED", "MODEL_RESOURCES_STATUS_STOPPING".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaRlModelResourceListParams]'s query parameters as
// `url.Values`.
func (r BetaRlModelResourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaRlModelResourceEstimateCostParams struct {
	// Base model to provision the resource for
	BaseModel string `json:"base_model" api:"required"`
	// Whether the resource hosts LoRA sessions or a single full-weight session
	LoraEnabled param.Opt[bool] `json:"lora_enabled,omitzero"`
	// Compute layout to provision.
	ComputeConfig BetaRlModelResourceEstimateCostParamsComputeConfig `json:"compute_config,omitzero"`
	// Optimizer configuration for this resource.
	OptimizerConfig OptimizerConfigParam `json:"optimizer_config,omitzero"`
	paramObj
}

func (r BetaRlModelResourceEstimateCostParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlModelResourceEstimateCostParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlModelResourceEstimateCostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compute layout to provision.
type BetaRlModelResourceEstimateCostParamsComputeConfig struct {
	// Number of generator replicas. 0 runs the trainer only, with no generator.
	NumGeneratorReplicas param.Opt[int64] `json:"num_generator_replicas,omitzero"`
	// GPU type to provision. Omit to use the model's default GPU type.
	//
	// Any of "H100-80GB", "B200-SXM".
	GPUType string `json:"gpu_type,omitzero"`
	paramObj
}

func (r BetaRlModelResourceEstimateCostParamsComputeConfig) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlModelResourceEstimateCostParamsComputeConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlModelResourceEstimateCostParamsComputeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaRlModelResourceEstimateCostParamsComputeConfig](
		"gpu_type", "H100-80GB", "B200-SXM",
	)
}

type BetaRlModelResourceStopParams struct {
	// Stop the resource even if active training sessions are attached
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaRlModelResourceStopParams]'s query parameters as
// `url.Values`.
func (r BetaRlModelResourceStopParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
