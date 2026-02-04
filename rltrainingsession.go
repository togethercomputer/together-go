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
)

// RlTrainingSessionService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRlTrainingSessionService] method instead.
type RlTrainingSessionService struct {
	Options    []option.RequestOption
	Operations RlTrainingSessionOperationService
}

// NewRlTrainingSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRlTrainingSessionService(opts ...option.RequestOption) (r RlTrainingSessionService) {
	r = RlTrainingSessionService{}
	r.Options = opts
	r.Operations = NewRlTrainingSessionOperationService(opts...)
	return
}

// Creates a training session and returns its details.
func (r *RlTrainingSessionService) New(ctx context.Context, body RlTrainingSessionNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "rl/training-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Gets a training session by its ID and returns its details.
func (r *RlTrainingSessionService) Get(ctx context.Context, sessionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("rl/training-sessions/%s", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Lists all training sessions.
func (r *RlTrainingSessionService) List(ctx context.Context, query RlTrainingSessionListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "rl/training-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Submits a forward-backward pass operation that will asynchronously compute
// gradients via backpropagation.
func (r *RlTrainingSessionService) ForwardBackward(ctx context.Context, sessionID string, body RlTrainingSessionForwardBackwardParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("rl/training-sessions/%s:forward-backward", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Submits an optimizer step operation that will asynchronously apply accumulated
// gradients to update model parameters.
func (r *RlTrainingSessionService) OptimStep(ctx context.Context, sessionID string, body RlTrainingSessionOptimStepParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("rl/training-sessions/%s:optim-step", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Stops a training session.
func (r *RlTrainingSessionService) Stop(ctx context.Context, sessionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("rl/training-sessions/%s:stop", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type RlTrainingSessionNewParams struct {
	Body RlTrainingSessionNewParamsBody `json:"body,omitzero,required"`
	paramObj
}

// The property BaseModel is required.
type RlTrainingSessionNewParamsBody struct {
	// Base model to use for the training session
	BaseModel string `json:"base_model,required"`
	// Checkpoint ID to use for the training session
	CheckpointID param.Opt[string] `json:"checkpoint_id,omitzero"`
	// LoRA adapter configuration
	LoraConfig RlTrainingSessionNewParamsBodyLoraConfig `json:"lora_config,omitzero"`
	// Learning rate scheduler configuration
	LrSchedulerConfig RlTrainingSessionNewParamsBodyLrSchedulerConfig `json:"lr_scheduler_config,omitzero"`
	// Optimizer configuration. If omitted, defaults to AdamW with default parameters.
	OptimizerConfig RlTrainingSessionNewParamsBodyOptimizerConfig `json:"optimizer_config,omitzero"`
	paramObj
}

func (r RlTrainingSessionNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LoRA adapter configuration
type RlTrainingSessionNewParamsBodyLoraConfig struct {
	// Alpha of the LoRA adapter
	Alpha param.Opt[int64] `json:"alpha,omitzero"`
	// Dropout of the LoRA adapter
	Dropout param.Opt[float64] `json:"dropout,omitzero"`
	// Rank of the LoRA adapter
	Rank param.Opt[int64] `json:"rank,omitzero"`
	paramObj
}

func (r RlTrainingSessionNewParamsBodyLoraConfig) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionNewParamsBodyLoraConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionNewParamsBodyLoraConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Learning rate scheduler configuration
type RlTrainingSessionNewParamsBodyLrSchedulerConfig struct {
	// Linear learning rate scheduler configuration
	Linear RlTrainingSessionNewParamsBodyLrSchedulerConfigLinear `json:"linear,omitzero"`
	paramObj
}

func (r RlTrainingSessionNewParamsBodyLrSchedulerConfig) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionNewParamsBodyLrSchedulerConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionNewParamsBodyLrSchedulerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Linear learning rate scheduler configuration
type RlTrainingSessionNewParamsBodyLrSchedulerConfigLinear struct {
	// Linear learning rate scheduler parameters
	Params RlTrainingSessionNewParamsBodyLrSchedulerConfigLinearParams `json:"params,omitzero"`
	paramObj
}

func (r RlTrainingSessionNewParamsBodyLrSchedulerConfigLinear) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionNewParamsBodyLrSchedulerConfigLinear
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionNewParamsBodyLrSchedulerConfigLinear) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Linear learning rate scheduler parameters
type RlTrainingSessionNewParamsBodyLrSchedulerConfigLinearParams struct {
	// Minimum learning rate at the end of linear decay
	LrMin param.Opt[float64] `json:"lr_min,omitzero"`
	// Number of warmup steps
	WarmupSteps param.Opt[int64] `json:"warmup_steps,omitzero"`
	paramObj
}

func (r RlTrainingSessionNewParamsBodyLrSchedulerConfigLinearParams) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionNewParamsBodyLrSchedulerConfigLinearParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionNewParamsBodyLrSchedulerConfigLinearParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optimizer configuration. If omitted, defaults to AdamW with default parameters.
type RlTrainingSessionNewParamsBodyOptimizerConfig struct {
	// Maximum gradient norm for gradient clipping. Applies to all optimizer types.
	MaxGradNorm param.Opt[float64] `json:"max_grad_norm,omitzero"`
	// AdamW optimizer configuration
	Adamw RlTrainingSessionNewParamsBodyOptimizerConfigAdamw `json:"adamw,omitzero"`
	paramObj
}

func (r RlTrainingSessionNewParamsBodyOptimizerConfig) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionNewParamsBodyOptimizerConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionNewParamsBodyOptimizerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AdamW optimizer configuration
type RlTrainingSessionNewParamsBodyOptimizerConfigAdamw struct {
	// AdamW optimizer parameters
	Params RlTrainingSessionNewParamsBodyOptimizerConfigAdamwParams `json:"params,omitzero"`
	paramObj
}

func (r RlTrainingSessionNewParamsBodyOptimizerConfigAdamw) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionNewParamsBodyOptimizerConfigAdamw
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionNewParamsBodyOptimizerConfigAdamw) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AdamW optimizer parameters
type RlTrainingSessionNewParamsBodyOptimizerConfigAdamwParams struct {
	// First moment decay rate
	Beta1 param.Opt[float64] `json:"beta1,omitzero"`
	// Second moment decay rate
	Beta2 param.Opt[float64] `json:"beta2,omitzero"`
	// Epsilon for numerical stability
	Eps param.Opt[float64] `json:"eps,omitzero"`
	// Learning rate
	Lr param.Opt[float64] `json:"lr,omitzero"`
	// Weight decay coefficient
	WeightDecay param.Opt[float64] `json:"weight_decay,omitzero"`
	paramObj
}

func (r RlTrainingSessionNewParamsBodyOptimizerConfigAdamwParams) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionNewParamsBodyOptimizerConfigAdamwParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionNewParamsBodyOptimizerConfigAdamwParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RlTrainingSessionListParams struct {
	// Maximum number of sessions to return (1-100)
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// Number of sessions to skip
	Offset param.Opt[string] `query:"offset,omitzero" json:"-"`
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RlTrainingSessionListParams]'s query parameters as
// `url.Values`.
func (r RlTrainingSessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RlTrainingSessionForwardBackwardParams struct {
	Body RlTrainingSessionForwardBackwardParamsBody `json:"body,omitzero,required"`
	paramObj
}

// The properties LossFn, Samples are required.
type RlTrainingSessionForwardBackwardParamsBody struct {
	// Loss function to use for gradient computation
	//
	// Any of "LOSS_FN_UNSPECIFIED", "LOSS_FN_GRPO".
	LossFn string `json:"loss_fn,omitzero,required"`
	// Batch of training samples to process
	Samples []RlTrainingSessionForwardBackwardParamsBodySample `json:"samples,omitzero,required"`
	paramObj
}

func (r RlTrainingSessionForwardBackwardParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionForwardBackwardParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionForwardBackwardParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RlTrainingSessionForwardBackwardParamsBody](
		"loss_fn", "LOSS_FN_UNSPECIFIED", "LOSS_FN_GRPO",
	)
}

// The properties LossFnInputs, ModelInput are required.
type RlTrainingSessionForwardBackwardParamsBodySample struct {
	// Loss function inputs
	LossFnInputs RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputs `json:"loss_fn_inputs,omitzero,required"`
	// Model input
	ModelInput RlTrainingSessionForwardBackwardParamsBodySampleModelInput `json:"model_input,omitzero,required"`
	paramObj
}

func (r RlTrainingSessionForwardBackwardParamsBodySample) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionForwardBackwardParamsBodySample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionForwardBackwardParamsBodySample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Loss function inputs
//
// The properties TargetTokens, Weights are required.
type RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputs struct {
	// Target tokens for loss computation
	TargetTokens RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsTargetTokens `json:"target_tokens,omitzero,required"`
	// Per-token weights
	Weights RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsWeights `json:"weights,omitzero,required"`
	paramObj
}

func (r RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputs) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target tokens for loss computation
//
// The property Data is required.
type RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsTargetTokens struct {
	// Integer array of target tokens
	Data []float64 `json:"data,omitzero,required"`
	// Data type of the integer array
	//
	// Any of "D_TYPE_UNSPECIFIED", "D_TYPE_INT64", "D_TYPE_FLOAT32",
	// "D_TYPE_BFLOAT16".
	Dtype string `json:"dtype,omitzero"`
	paramObj
}

func (r RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsTargetTokens) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsTargetTokens
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsTargetTokens) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsTargetTokens](
		"dtype", "D_TYPE_UNSPECIFIED", "D_TYPE_INT64", "D_TYPE_FLOAT32", "D_TYPE_BFLOAT16",
	)
}

// Per-token weights
//
// The property Data is required.
type RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsWeights struct {
	// Float array of per-token weights
	Data []float64 `json:"data,omitzero,required"`
	// Data type of the float array
	//
	// Any of "D_TYPE_UNSPECIFIED", "D_TYPE_INT64", "D_TYPE_FLOAT32",
	// "D_TYPE_BFLOAT16".
	Dtype string `json:"dtype,omitzero"`
	paramObj
}

func (r RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsWeights) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsWeights
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsWeights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RlTrainingSessionForwardBackwardParamsBodySampleLossFnInputsWeights](
		"dtype", "D_TYPE_UNSPECIFIED", "D_TYPE_INT64", "D_TYPE_FLOAT32", "D_TYPE_BFLOAT16",
	)
}

// Model input
//
// The property Chunks is required.
type RlTrainingSessionForwardBackwardParamsBodySampleModelInput struct {
	// Input chunks for the model
	Chunks []RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunk `json:"chunks,omitzero,required"`
	paramObj
}

func (r RlTrainingSessionForwardBackwardParamsBodySampleModelInput) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionForwardBackwardParamsBodySampleModelInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionForwardBackwardParamsBodySampleModelInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunk struct {
	EncodedText RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunkEncodedText `json:"encoded_text,omitzero"`
	paramObj
}

func (r RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunk) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunk
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunkEncodedText struct {
	// Pre-tokenized text input
	Tokens []int64 `json:"tokens,omitzero"`
	paramObj
}

func (r RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunkEncodedText) MarshalJSON() (data []byte, err error) {
	type shadow RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunkEncodedText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RlTrainingSessionForwardBackwardParamsBodySampleModelInputChunkEncodedText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RlTrainingSessionOptimStepParams struct {
	Body any `json:"body,omitzero,required"`
	paramObj
}
