// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaRlOperationService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaRlOperationService] method instead.
type BetaRlOperationService struct {
	Options []option.RequestOption
}

// NewBetaRlOperationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaRlOperationService(opts ...option.RequestOption) (r BetaRlOperationService) {
	r = BetaRlOperationService{}
	r.Options = opts
	return
}

// Submits an operation that will asynchronously save the current LoRA adapter as
// an inference checkpoint and upload it to object storage.
func (r *BetaRlOperationService) NewInferenceCheckpoint(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *InferenceCheckpointOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/inference-checkpoint", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Submits an operation that will asynchronously save the full training state
// (adapter + optimizer + step).
func (r *BetaRlOperationService) NewTrainingCheckpoint(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *TrainingCheckpointOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/training-checkpoint", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Submits a forward-backward pass driven by externally computed gradients of the
// loss with respect to per-token log-probabilities.
func (r *BetaRlOperationService) CustomForwardBackward(ctx context.Context, sessionID string, body BetaRlOperationCustomForwardBackwardParams, opts ...option.RequestOption) (res *CustomForwardBackwardOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/custom-forward-backward", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Submits a forward-backward pass operation that will asynchronously compute
// gradients via backpropagation.
func (r *BetaRlOperationService) ForwardBackward(ctx context.Context, sessionID string, body BetaRlOperationForwardBackwardParams, opts ...option.RequestOption) (res *ForwardBackwardOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/forward-backward", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Submits an optimizer step operation that will asynchronously apply accumulated
// gradients to update model parameters. Does not make the updated parameters
// available for sampling; call `weights-sync` afterwards when you want subsequent
// samples to use the updated policy.
func (r *BetaRlOperationService) OptimStep(ctx context.Context, sessionID string, body BetaRlOperationOptimStepParams, opts ...option.RequestOption) (res *OptimStepOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/optim-step", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the current status and result of a custom forward-backward operation.
func (r *BetaRlOperationService) GetCustomForwardBackward(ctx context.Context, operationID string, query BetaRlOperationGetCustomForwardBackwardParams, opts ...option.RequestOption) (res *CustomForwardBackwardOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/custom-forward-backward/%s", query.SessionID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the current status and result of a forward-backward operation.
func (r *BetaRlOperationService) GetForwardBackward(ctx context.Context, operationID string, query BetaRlOperationGetForwardBackwardParams, opts ...option.RequestOption) (res *ForwardBackwardOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/forward-backward/%s", query.SessionID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the current status and result of an inference checkpoint operation.
func (r *BetaRlOperationService) GetInferenceCheckpoint(ctx context.Context, operationID string, query BetaRlOperationGetInferenceCheckpointParams, opts ...option.RequestOption) (res *InferenceCheckpointOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/inference-checkpoint/%s", query.SessionID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the current status and result of an optim-step operation.
func (r *BetaRlOperationService) GetOptimStep(ctx context.Context, operationID string, query BetaRlOperationGetOptimStepParams, opts ...option.RequestOption) (res *OptimStepOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/optim-step/%s", query.SessionID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the current status and result of a sample operation.
func (r *BetaRlOperationService) GetSample(ctx context.Context, operationID string, query BetaRlOperationGetSampleParams, opts ...option.RequestOption) (res *SampleOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/sample/%s", query.SessionID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the current status and result of a save training checkpoint operation.
func (r *BetaRlOperationService) GetTrainingCheckpoint(ctx context.Context, operationID string, query BetaRlOperationGetTrainingCheckpointParams, opts ...option.RequestOption) (res *TrainingCheckpointOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/training-checkpoint/%s", query.SessionID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the current status and result of a weights-sync operation.
func (r *BetaRlOperationService) GetWeightsSync(ctx context.Context, operationID string, query BetaRlOperationGetWeightsSyncParams, opts ...option.RequestOption) (res *WeightsSyncOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/weights-sync/%s", query.SessionID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Submits a sample operation that will asynchronously generate text completions
// with logprobs.
func (r *BetaRlOperationService) Sample(ctx context.Context, sessionID string, body BetaRlOperationSampleParams, opts ...option.RequestOption) (res *SampleOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/sample", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Submits a weights-sync operation that makes the session's current trained
// parameters available for sampling. Call this after `optim-step` when you want
// subsequent samples to use the updated policy.
func (r *BetaRlOperationService) WeightsSync(ctx context.Context, sessionID string, body BetaRlOperationWeightsSyncParams, opts ...option.RequestOption) (res *WeightsSyncOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/weights-sync", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Per-step Adam optimizer overrides.
type AdamParams struct {
	// Exponential decay rate for the first-moment estimate
	Beta1 param.Opt[float64] `json:"beta1,omitzero"`
	// Exponential decay rate for the second-moment estimate
	Beta2 param.Opt[float64] `json:"beta2,omitzero"`
	// Epsilon for numerical stability
	Eps param.Opt[float64] `json:"eps,omitzero"`
	// Maximum gradient norm for this step, gradients across all model parameters are
	// clipped to this value. Set to 0 to disable gradient clipping. When unset,
	// gradients are clipped to the session default (1.0).
	GradClipNorm param.Opt[float64] `json:"grad_clip_norm,omitzero"`
	// Learning rate for the Adam-tuned parameters
	LearningRate param.Opt[float64] `json:"learning_rate,omitzero"`
	// Weight decay coefficient
	WeightDecay param.Opt[float64] `json:"weight_decay,omitzero"`
	paramObj
}

func (r AdamParams) MarshalJSON() (data []byte, err error) {
	type shadow AdamParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdamParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CispoLossParams struct {
	// Upper absolute bound for the importance ratio; the clipped ratio is applied as a
	// detached coefficient
	ClipHighThreshold param.Opt[float64] `json:"clip_high_threshold,omitzero"`
	// Lower absolute bound for the importance ratio; the clipped ratio is applied as a
	// detached coefficient
	ClipLowThreshold param.Opt[float64] `json:"clip_low_threshold,omitzero"`
	paramObj
}

func (r CispoLossParams) MarshalJSON() (data []byte, err error) {
	type shadow CispoLossParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CispoLossParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cross-entropy loss parameters (currently empty).
type CrossEntropyLossParams struct {
	paramObj
}

func (r CrossEntropyLossParams) MarshalJSON() (data []byte, err error) {
	type shadow CrossEntropyLossParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrossEntropyLossParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Async custom forward-backward pass operation
type CustomForwardBackwardOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status OperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error OperationError `json:"error"`
	// Result on success
	Output CustomForwardBackwardResult `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomForwardBackwardOperation) RawJSON() string { return r.JSON.raw }
func (r *CustomForwardBackwardOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a custom forward-backward pass operation
type CustomForwardBackwardResult struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomForwardBackwardResult) RawJSON() string { return r.JSON.raw }
func (r *CustomForwardBackwardResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DType string

const (
	DTypeDTypeUnspecified DType = "D_TYPE_UNSPECIFIED"
	DTypeDTypeInt64       DType = "D_TYPE_INT64"
	DTypeDTypeFloat32     DType = "D_TYPE_FLOAT32"
	DTypeDTypeBfloat16    DType = "D_TYPE_BFLOAT16"
)

// The property Beta is required.
type DroLossParams struct {
	// Coefficient on the quadratic log-ratio penalty. Required; there is no default.
	Beta float64 `json:"beta" api:"required"`
	paramObj
}

func (r DroLossParams) MarshalJSON() (data []byte, err error) {
	type shadow DroLossParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DroLossParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pre-tokenized text content for a model input chunk.
//
// The property Tokens is required.
type EncodedTextChunkParam struct {
	// Pre-tokenized text input
	Tokens []EncodedTextChunkTokenUnionParam `json:"tokens,omitzero" api:"required"`
	paramObj
}

func (r EncodedTextChunkParam) MarshalJSON() (data []byte, err error) {
	type shadow EncodedTextChunkParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EncodedTextChunkParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EncodedTextChunkTokenUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u EncodedTextChunkTokenUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *EncodedTextChunkTokenUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EncodedTextChunkTokenUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Async forward-backward pass operation
type ForwardBackwardOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status OperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error OperationError `json:"error"`
	// Result on success
	Output ForwardBackwardResult `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForwardBackwardOperation) RawJSON() string { return r.JSON.raw }
func (r *ForwardBackwardOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a scored forward or forward-backward operation
type ForwardBackwardResult struct {
	// Loss value
	Loss float64 `json:"loss" api:"required"`
	// Per-sample loss function outputs, in request order. Empty unless the request set
	// `return_loss_fn_outputs`.
	LossFnOutputs []LossFnOutput `json:"loss_fn_outputs"`
	// Loss-specific metrics (e.g., KL divergence, clip fraction for GRPO)
	Metrics map[string]float64 `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Loss          respjson.Field
		LossFnOutputs respjson.Field
		Metrics       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForwardBackwardResult) RawJSON() string { return r.JSON.raw }
func (r *ForwardBackwardResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GrpoLossAggregationType string

const (
	GrpoLossAggregationTypeGrpoLossAggregationTypeUnspecified  GrpoLossAggregationType = "GRPO_LOSS_AGGREGATION_TYPE_UNSPECIFIED"
	GrpoLossAggregationTypeGrpoLossAggregationTypeFixedHorizon GrpoLossAggregationType = "GRPO_LOSS_AGGREGATION_TYPE_FIXED_HORIZON"
	GrpoLossAggregationTypeGrpoLossAggregationTypeTokenMean    GrpoLossAggregationType = "GRPO_LOSS_AGGREGATION_TYPE_TOKEN_MEAN"
	GrpoLossAggregationTypeGrpoLossAggregationTypeSequenceMean GrpoLossAggregationType = "GRPO_LOSS_AGGREGATION_TYPE_SEQUENCE_MEAN"
)

type GrpoLossParams struct {
	// KL penalty coefficient
	Beta param.Opt[float64] `json:"beta,omitzero"`
	// Upper clip threshold for the importance-sampling ratio. The ratio is clamped to
	// this bound; tighter clipping makes policy updates more conservative. Must
	// be >= 1.
	ClipHighThreshold param.Opt[float64] `json:"clip_high_threshold,omitzero"`
	// Lower clip threshold for the importance-sampling ratio. The ratio is clamped to
	// this bound; tighter clipping makes policy updates more conservative. Must be
	// <= 1.
	ClipLowThreshold param.Opt[float64] `json:"clip_low_threshold,omitzero"`
	// Aggregation type for loss computation
	//
	// Any of "GRPO_LOSS_AGGREGATION_TYPE_UNSPECIFIED",
	// "GRPO_LOSS_AGGREGATION_TYPE_FIXED_HORIZON",
	// "GRPO_LOSS_AGGREGATION_TYPE_TOKEN_MEAN",
	// "GRPO_LOSS_AGGREGATION_TYPE_SEQUENCE_MEAN".
	AggType GrpoLossAggregationType `json:"agg_type,omitzero"`
	// Controls how the importance-sampling ratio is computed in GRPO loss. Defaults to
	// token-level ratios, which is the standard GRPO behavior. Use sequence-level
	// ratios to enable GSPO-style loss calculation instead.
	//
	// Any of "GRPO_LOSS_RATIO_TYPE_TOKEN", "GRPO_LOSS_RATIO_TYPE_SEQUENCE".
	RatioType GrpoLossRatioType `json:"ratio_type,omitzero"`
	paramObj
}

func (r GrpoLossParams) MarshalJSON() (data []byte, err error) {
	type shadow GrpoLossParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GrpoLossParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls whether GRPO loss uses token-level or sequence-level importance ratios.
type GrpoLossRatioType string

const (
	GrpoLossRatioTypeGrpoLossRatioTypeToken    GrpoLossRatioType = "GRPO_LOSS_RATIO_TYPE_TOKEN"
	GrpoLossRatioTypeGrpoLossRatioTypeSequence GrpoLossRatioType = "GRPO_LOSS_RATIO_TYPE_SEQUENCE"
)

// Async inference checkpoint operation
type InferenceCheckpointOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status OperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error OperationError `json:"error"`
	// Result on success
	Output InferenceCheckpointResult `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceCheckpointOperation) RawJSON() string { return r.JSON.raw }
func (r *InferenceCheckpointOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an inference checkpoint operation
type InferenceCheckpointResult struct {
	// Registered model name for downloading the checkpoint
	ModelName string `json:"model_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceCheckpointResult) RawJSON() string { return r.JSON.raw }
func (r *InferenceCheckpointResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type LossConfigParam struct {
	// Type of loss function to use
	//
	// Any of "LOSS_TYPE_UNSPECIFIED", "LOSS_TYPE_CROSS_ENTROPY", "LOSS_TYPE_GRPO",
	// "LOSS_TYPE_IMPORTANCE_SAMPLING", "LOSS_TYPE_PPO", "LOSS_TYPE_CISPO",
	// "LOSS_TYPE_DRO".
	Type        LossType        `json:"type,omitzero" api:"required"`
	CispoParams CispoLossParams `json:"cispo_params,omitzero"`
	// Cross-entropy loss parameters (currently empty).
	CrossEntropyParams CrossEntropyLossParams `json:"cross_entropy_params,omitzero"`
	DroParams          DroLossParams          `json:"dro_params,omitzero"`
	GrpoParams         GrpoLossParams         `json:"grpo_params,omitzero"`
	PpoParams          PpoLossParams          `json:"ppo_params,omitzero"`
	paramObj
}

func (r LossConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow LossConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LossConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Output tensors produced by the loss function for one sample.
type LossFnOutput struct {
	// Output tensors keyed by name. Built-in losses return `logprobs`: the model's
	// float32 per-token log-probabilities under the current policy, one value per
	// token of the sample's input. Positions excluded from the loss, such as
	// zero-weight positions, are masked to zero rather than true log-probabilities.
	Tensors map[string]TensorData `json:"tensors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tensors     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LossFnOutput) RawJSON() string { return r.JSON.raw }
func (r *LossFnOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of loss function used for RL training.
type LossType string

const (
	LossTypeLossTypeUnspecified        LossType = "LOSS_TYPE_UNSPECIFIED"
	LossTypeLossTypeCrossEntropy       LossType = "LOSS_TYPE_CROSS_ENTROPY"
	LossTypeLossTypeGrpo               LossType = "LOSS_TYPE_GRPO"
	LossTypeLossTypeImportanceSampling LossType = "LOSS_TYPE_IMPORTANCE_SAMPLING"
	LossTypeLossTypePpo                LossType = "LOSS_TYPE_PPO"
	LossTypeLossTypeCispo              LossType = "LOSS_TYPE_CISPO"
	LossTypeLossTypeDro                LossType = "LOSS_TYPE_DRO"
)

// The property Chunks is required.
type ModelInputParam struct {
	// Input chunks for the model
	Chunks []ModelInputChunkParam `json:"chunks,omitzero" api:"required"`
	paramObj
}

func (r ModelInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single chunk of model input content.
//
// The property EncodedText is required.
type ModelInputChunkParam struct {
	// Pre-tokenized text content for this input chunk.
	EncodedText EncodedTextChunkParam `json:"encoded_text,omitzero" api:"required"`
	paramObj
}

func (r ModelInputChunkParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelInputChunkParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelInputChunkParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-step Muon optimizer overrides
type MuonParams struct {
	// Maximum gradient norm for this step, gradients across all model parameters are
	// clipped to this value. Set to 0 to disable gradient clipping. When unset,
	// gradients are clipped to the session default (1.0).
	GradClipNorm param.Opt[float64] `json:"grad_clip_norm,omitzero"`
	// Learning rate for this Muon optimizer step.
	LearningRate param.Opt[float64] `json:"learning_rate,omitzero"`
	// Momentum coefficient
	Momentum param.Opt[float64] `json:"momentum,omitzero"`
	// Number of Newton-Schulz iterations
	NewtonSchulzSteps param.Opt[int64] `json:"newton_schulz_steps,omitzero"`
	// Weight decay coefficient
	WeightDecay param.Opt[float64] `json:"weight_decay,omitzero"`
	// Per-step Adam optimizer overrides for the Adam-tuned parameters in a Muon-tuned
	// optimizer session.
	Adam AdamParams `json:"adam,omitzero"`
	paramObj
}

func (r MuonParams) MarshalJSON() (data []byte, err error) {
	type shadow MuonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MuonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details for a failed training operation
type OperationError struct {
	// Application error code
	//
	// Any of "TRAINING_OPERATION_ERROR_CODE_UNSPECIFIED",
	// "TRAINING_OPERATION_ERROR_CODE_RESOURCE_EXHAUSTED",
	// "TRAINING_OPERATION_ERROR_CODE_TIMEOUT",
	// "TRAINING_OPERATION_ERROR_CODE_INTERNAL_ERROR",
	// "TRAINING_OPERATION_ERROR_CODE_SESSION_NOT_ACTIVE",
	// "TRAINING_OPERATION_ERROR_CODE_INVALID_INPUT",
	// "TRAINING_OPERATION_ERROR_CODE_NON_FINITE_LOSS".
	Code OperationErrorCode `json:"code"`
	// Human-readable error message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationError) RawJSON() string { return r.JSON.raw }
func (r *OperationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Application error code for a failed training operation
type OperationErrorCode string

const (
	OperationErrorCodeTrainingOperationErrorCodeUnspecified       OperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_UNSPECIFIED"
	OperationErrorCodeTrainingOperationErrorCodeResourceExhausted OperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_RESOURCE_EXHAUSTED"
	OperationErrorCodeTrainingOperationErrorCodeTimeout           OperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_TIMEOUT"
	OperationErrorCodeTrainingOperationErrorCodeInternalError     OperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_INTERNAL_ERROR"
	OperationErrorCodeTrainingOperationErrorCodeSessionNotActive  OperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_SESSION_NOT_ACTIVE"
	OperationErrorCodeTrainingOperationErrorCodeInvalidInput      OperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_INVALID_INPUT"
	OperationErrorCodeTrainingOperationErrorCodeNonFiniteLoss     OperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_NON_FINITE_LOSS"
)

type OperationStatus string

const (
	OperationStatusTrainingOperationStatusUnspecified OperationStatus = "TRAINING_OPERATION_STATUS_UNSPECIFIED"
	OperationStatusTrainingOperationStatusPending     OperationStatus = "TRAINING_OPERATION_STATUS_PENDING"
	OperationStatusTrainingOperationStatusRunning     OperationStatus = "TRAINING_OPERATION_STATUS_RUNNING"
	OperationStatusTrainingOperationStatusCompleted   OperationStatus = "TRAINING_OPERATION_STATUS_COMPLETED"
	OperationStatusTrainingOperationStatusFailed      OperationStatus = "TRAINING_OPERATION_STATUS_FAILED"
)

// Async optimizer step operation
type OptimStepOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status OperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error OperationError `json:"error"`
	// Result on success
	Output OptimStepResult `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimStepOperation) RawJSON() string { return r.JSON.raw }
func (r *OptimStepOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an optimizer step operation
type OptimStepResult struct {
	// Step number
	Step OptimStepResultStepUnion `json:"step" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Step        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptimStepResult) RawJSON() string { return r.JSON.raw }
func (r *OptimStepResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OptimStepResultStepUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type OptimStepResultStepUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u OptimStepResultStepUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OptimStepResultStepUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OptimStepResultStepUnion) RawJSON() string { return u.JSON.raw }

func (r *OptimStepResultStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A (policy version, starting token) span within a sampled sequence. Version 0 is
// the initial model; each optim_step call increments the version by 1.
type PolicyVersionSegment struct {
	// Index of the first token of this segment within the sampled sequence. Always 0
	// for the first segment.
	StartToken int64 `json:"start_token" api:"required"`
	// Model version under which this segment of tokens was generated
	Version int64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartToken  respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyVersionSegment) RawJSON() string { return r.JSON.raw }
func (r *PolicyVersionSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PpoLossParams struct {
	// Upper absolute bound for the importance ratio in the clipped surrogate. Must
	// be >= 1.
	ClipHighThreshold param.Opt[float64] `json:"clip_high_threshold,omitzero"`
	// Lower absolute bound for the importance ratio in the clipped surrogate. Must be
	// <= 1.
	ClipLowThreshold param.Opt[float64] `json:"clip_low_threshold,omitzero"`
	paramObj
}

func (r PpoLossParams) MarshalJSON() (data []byte, err error) {
	type shadow PpoLossParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PpoLossParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The most likely alternative tokens at a single prompt position, as two parallel
// arrays of equal length. Both are empty for a position with no conditioning
// context, such as position 0.
type PromptTopLogprobs struct {
	// Log-probability of each alternative in `token_ids`, at the same index.
	Logprobs []float64 `json:"logprobs"`
	// Token IDs of the alternatives, ordered by descending log-probability.
	TokenIDs []int64 `json:"token_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logprobs    respjson.Field
		TokenIDs    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptTopLogprobs) RawJSON() string { return r.JSON.raw }
func (r *PromptTopLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mixture-of-experts routing decisions captured while generating, so training can
// reuse the same expert selection. Exactly one source is set—legacy inline `data`,
// or a backend-owned `object_uri` that the manager hydrates before training. The
// contiguous int32 buffer is reshaped by `shape`, which is always
// `[num_tokens, num_layers, width]`; packed buffers carry fp32-bitcast routing
// weights in the trailing top-k columns.
type RoutedExperts struct {
	// Buffer shape as `[num_tokens, num_layers, width]`.
	Shape []RoutedExpertsShapeUnion `json:"shape" api:"required"`
	// Legacy base64-encoded contiguous int32 routing buffer, row-major over (token,
	// layer, width).
	Data string `json:"data" format:"byte"`
	// Backend-owned S3/R2 object URI containing the contiguous int32 routing buffer.
	// Clients relay this URI unchanged; the manager validates and downloads it before
	// training.
	ObjectUri string `json:"object_uri" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Shape       respjson.Field
		Data        respjson.Field
		ObjectUri   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoutedExperts) RawJSON() string { return r.JSON.raw }
func (r *RoutedExperts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RoutedExperts to a RoutedExpertsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RoutedExpertsParam.Overrides()
func (r RoutedExperts) ToParam() RoutedExpertsParam {
	return param.Override[RoutedExpertsParam](json.RawMessage(r.RawJSON()))
}

// RoutedExpertsShapeUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type RoutedExpertsShapeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u RoutedExpertsShapeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RoutedExpertsShapeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RoutedExpertsShapeUnion) RawJSON() string { return u.JSON.raw }

func (r *RoutedExpertsShapeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mixture-of-experts routing decisions captured while generating, so training can
// reuse the same expert selection. Exactly one source is set—legacy inline `data`,
// or a backend-owned `object_uri` that the manager hydrates before training. The
// contiguous int32 buffer is reshaped by `shape`, which is always
// `[num_tokens, num_layers, width]`; packed buffers carry fp32-bitcast routing
// weights in the trailing top-k columns.
//
// The property Shape is required.
type RoutedExpertsParam struct {
	// Buffer shape as `[num_tokens, num_layers, width]`.
	Shape []RoutedExpertsShapeUnionParam `json:"shape,omitzero" api:"required"`
	// Legacy base64-encoded contiguous int32 routing buffer, row-major over (token,
	// layer, width).
	Data param.Opt[string] `json:"data,omitzero" format:"byte"`
	// Backend-owned S3/R2 object URI containing the contiguous int32 routing buffer.
	// Clients relay this URI unchanged; the manager validates and downloads it before
	// training.
	ObjectUri param.Opt[string] `json:"object_uri,omitzero" format:"uri"`
	paramObj
}

func (r RoutedExpertsParam) MarshalJSON() (data []byte, err error) {
	type shadow RoutedExpertsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RoutedExpertsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RoutedExpertsShapeUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u RoutedExpertsShapeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *RoutedExpertsShapeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RoutedExpertsShapeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Async sample operation
type SampleOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status OperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error OperationError `json:"error"`
	// Result on success
	Output SampleOperationOutput `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SampleOperation) RawJSON() string { return r.JSON.raw }
func (r *SampleOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result on success
type SampleOperationOutput struct {
	// One result per model input
	Results []SampleResult `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SampleOperationOutput) RawJSON() string { return r.JSON.raw }
func (r *SampleOperationOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Completions generated for a single model input
type SampleResult struct {
	// Policy versions that produced these completions
	PolicySegments []PolicyVersionSegment `json:"policy_segments" api:"required"`
	// Generated completions
	Sequences []SampledSequence `json:"sequences" api:"required"`
	// Teacher-forced log-probability of each model input token. Full prompt length;
	// entry i corresponds to prompt token i. Entry 0 is always 0 as a placeholder: the
	// first prompt token has no conditioning context, so it has no log-probability.
	// Present only when prompt_logprobs was set on the request.
	PromptLogprobs []float64 `json:"prompt_logprobs"`
	// The most likely alternative tokens at each model input token, up to
	// `topk_prompt_logprobs` per position. Full prompt length; entry i corresponds to
	// prompt token i, and entry 0 is empty. Present only when topk_prompt_logprobs was
	// set on the request.
	TopkPromptLogprobs []PromptTopLogprobs `json:"topk_prompt_logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PolicySegments     respjson.Field
		Sequences          respjson.Field
		PromptLogprobs     respjson.Field
		TopkPromptLogprobs respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SampleResult) RawJSON() string { return r.JSON.raw }
func (r *SampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single generated completion sequence with tokens and logprobs
type SampledSequence struct {
	// Number of model input tokens served from the prefix cache while generating this
	// sequence.
	PromptCacheHitTokens int64 `json:"prompt_cache_hit_tokens" api:"required"`
	// Reason for stopping generation
	//
	// Any of "STOP_REASON_LENGTH", "STOP_REASON_STOP".
	StopReason StopReason `json:"stop_reason" api:"required"`
	// Generated token IDs
	Tokens []SampledSequenceTokenUnion `json:"tokens" api:"required"`
	// Log probabilities for each generated token
	Logprobs []float64 `json:"logprobs"`
	// MoE per-token routing decisions captured during generation; absent for dense
	// models or when capture is disabled.
	RoutedExperts RoutedExperts `json:"routed_experts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PromptCacheHitTokens respjson.Field
		StopReason           respjson.Field
		Tokens               respjson.Field
		Logprobs             respjson.Field
		RoutedExperts        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SampledSequence) RawJSON() string { return r.JSON.raw }
func (r *SampledSequence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SampledSequenceTokenUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type SampledSequenceTokenUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u SampledSequenceTokenUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SampledSequenceTokenUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SampledSequenceTokenUnion) RawJSON() string { return u.JSON.raw }

func (r *SampledSequenceTokenUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SamplingParams struct {
	// Maximum number of tokens to generate per completion
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// Sampling temperature
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Top-k sampling limit
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Nucleus sampling probability threshold
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Random seed for reproducible sampling for the same prompt and model state.
	// Per-completion seeds remain stable if the request is split across generator
	// replicas.
	Seed SamplingParamsSeedUnion `json:"seed,omitzero"`
	// Generation stops when any of these strings is produced
	Stop []string `json:"stop,omitzero"`
	paramObj
}

func (r SamplingParams) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SamplingParamsSeedUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u SamplingParamsSeedUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *SamplingParamsSeedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SamplingParamsSeedUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Reason generation stopped.
type StopReason string

const (
	StopReasonStopReasonLength StopReason = "STOP_REASON_LENGTH"
	StopReasonStopReasonStop   StopReason = "STOP_REASON_STOP"
)

// A tensor encoded as flattened row-major values, with an optional shape.
type TensorData struct {
	// Flattened one-dimensional values encoded as JSON numbers.
	Data []float64 `json:"data" api:"required"`
	// Tensor element type, either `int64` or `float32`.
	//
	// Any of "int64", "float32".
	Dtype TensorDataDtype `json:"dtype" api:"required"`
	// Optional tensor shape; training operations accept one-dimensional tensors only,
	// and the dimension must match the data length.
	Shape []int64 `json:"shape"`
	// Reserved for Tinker schema compatibility; current training operations reject
	// sparse tensors.
	SparseColIndices []int64 `json:"sparse_col_indices"`
	// Reserved for Tinker schema compatibility; current training operations reject
	// sparse tensors.
	SparseCrowIndices []int64 `json:"sparse_crow_indices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data              respjson.Field
		Dtype             respjson.Field
		Shape             respjson.Field
		SparseColIndices  respjson.Field
		SparseCrowIndices respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TensorData) RawJSON() string { return r.JSON.raw }
func (r *TensorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TensorData to a TensorDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TensorDataParam.Overrides()
func (r TensorData) ToParam() TensorDataParam {
	return param.Override[TensorDataParam](json.RawMessage(r.RawJSON()))
}

// Tensor element type, either `int64` or `float32`.
type TensorDataDtype string

const (
	TensorDataDtypeInt64   TensorDataDtype = "int64"
	TensorDataDtypeFloat32 TensorDataDtype = "float32"
)

// A tensor encoded as flattened row-major values, with an optional shape.
//
// The properties Data, Dtype are required.
type TensorDataParam struct {
	// Flattened one-dimensional values encoded as JSON numbers.
	Data []float64 `json:"data,omitzero" api:"required"`
	// Tensor element type, either `int64` or `float32`.
	//
	// Any of "int64", "float32".
	Dtype TensorDataDtype `json:"dtype,omitzero" api:"required"`
	// Optional tensor shape; training operations accept one-dimensional tensors only,
	// and the dimension must match the data length.
	Shape []int64 `json:"shape,omitzero"`
	// Reserved for Tinker schema compatibility; current training operations reject
	// sparse tensors.
	SparseColIndices []int64 `json:"sparse_col_indices,omitzero"`
	// Reserved for Tinker schema compatibility; current training operations reject
	// sparse tensors.
	SparseCrowIndices []int64 `json:"sparse_crow_indices,omitzero"`
	paramObj
}

func (r TensorDataParam) MarshalJSON() (data []byte, err error) {
	type shadow TensorDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TensorDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Async save training checkpoint operation
type TrainingCheckpointOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status OperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error OperationError `json:"error"`
	// Result on success
	Output TrainingCheckpointResult `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrainingCheckpointOperation) RawJSON() string { return r.JSON.raw }
func (r *TrainingCheckpointOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a save training checkpoint operation
type TrainingCheckpointResult struct {
	// ID of the saved training checkpoint (use for resume via Start)
	CheckpointID string `json:"checkpoint_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckpointID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrainingCheckpointResult) RawJSON() string { return r.JSON.raw }
func (r *TrainingCheckpointResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How updated policy parameters are made available for sampling. SYNCHRONOUS waits
// for the policy update before returning; BACKGROUND_PUBLISH returns after
// scheduling the update; PIPELINE overlaps the update with in-flight sampling when
// possible.
type WeightSyncType string

const (
	WeightSyncTypeWeightSyncTypeSynchronous       WeightSyncType = "WEIGHT_SYNC_TYPE_SYNCHRONOUS"
	WeightSyncTypeWeightSyncTypeBackgroundPublish WeightSyncType = "WEIGHT_SYNC_TYPE_BACKGROUND_PUBLISH"
	WeightSyncTypeWeightSyncTypePipeline          WeightSyncType = "WEIGHT_SYNC_TYPE_PIPELINE"
)

// Async weights-sync operation
type WeightsSyncOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status OperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error OperationError `json:"error"`
	// Result on success
	Output WeightsSyncResult `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WeightsSyncOperation) RawJSON() string { return r.JSON.raw }
func (r *WeightsSyncOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a weights-sync operation
type WeightsSyncResult struct {
	// Policy version now available for sampling, or queued to become available for
	// deferred sync modes. Comparable to `policy_segments[].version` on sample
	// results.
	WeightsVersion WeightsSyncResultWeightsVersionUnion `json:"weights_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WeightsVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WeightsSyncResult) RawJSON() string { return r.JSON.raw }
func (r *WeightsSyncResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WeightsSyncResultWeightsVersionUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WeightsSyncResultWeightsVersionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WeightsSyncResultWeightsVersionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WeightsSyncResultWeightsVersionUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WeightsSyncResultWeightsVersionUnion) RawJSON() string { return u.JSON.raw }

func (r *WeightsSyncResultWeightsVersionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlOperationCustomForwardBackwardParams struct {
	// Per-sample per-token gradients of the loss with respect to log-probabilities
	Gradients []BetaRlOperationCustomForwardBackwardParamsGradient `json:"gradients,omitzero" api:"required"`
	// Batch of training samples
	Samples []BetaRlOperationCustomForwardBackwardParamsSample `json:"samples,omitzero" api:"required"`
	paramObj
}

func (r BetaRlOperationCustomForwardBackwardParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationCustomForwardBackwardParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationCustomForwardBackwardParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-token gradients of the loss with respect to target log-probabilities
//
// The property Data is required.
type BetaRlOperationCustomForwardBackwardParamsGradient struct {
	// Float array of per-token gradients (d loss / d log p)
	Data []float64 `json:"data,omitzero" api:"required"`
	// Data type of the float array
	//
	// Any of "D_TYPE_UNSPECIFIED", "D_TYPE_INT64", "D_TYPE_FLOAT32",
	// "D_TYPE_BFLOAT16".
	Dtype DType `json:"dtype,omitzero"`
	paramObj
}

func (r BetaRlOperationCustomForwardBackwardParamsGradient) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationCustomForwardBackwardParamsGradient
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationCustomForwardBackwardParamsGradient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties LossFnInputs, ModelInput are required.
type BetaRlOperationCustomForwardBackwardParamsSample struct {
	// Per-token loss tensors keyed by name. Include `target_tokens` and the inputs
	// required by the selected loss. Each tensor must declare `int64` or `float32`, be
	// one-dimensional, and have the same length.
	LossFnInputs map[string]TensorDataParam `json:"loss_fn_inputs,omitzero" api:"required"`
	// Model input
	ModelInput ModelInputParam `json:"model_input,omitzero" api:"required"`
	// Optional MoE per-token routing captured at sample time. Replayed on every
	// training operation, so expert selection matches the one used at sample time.
	// Must cover the whole sample, or all but its last token.
	RoutedExperts RoutedExpertsParam `json:"routed_experts,omitzero"`
	paramObj
}

func (r BetaRlOperationCustomForwardBackwardParamsSample) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationCustomForwardBackwardParamsSample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationCustomForwardBackwardParamsSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlOperationForwardBackwardParams struct {
	// Loss function configuration
	Loss LossConfigParam `json:"loss,omitzero" api:"required"`
	// Batch of training samples to process
	Samples []BetaRlOperationForwardBackwardParamsSample `json:"samples,omitzero" api:"required"`
	// Run the forward pass only: report the loss and metrics, and the per-sample
	// outputs when requested, without accumulating gradients. Defaults to false. Pair
	// it with `return_loss_fn_outputs` to score a batch and read back its per-token
	// log-probabilities.
	ForwardOnly param.Opt[bool] `json:"forward_only,omitzero"`
	// Return the loss function's per-sample output tensors alongside the loss and
	// metrics. Defaults to false. Enabling it increases the response size
	// substantially for large batches and reduces step throughput, so leave it unset
	// for ordinary training steps.
	ReturnLossFnOutputs param.Opt[bool] `json:"return_loss_fn_outputs,omitzero"`
	paramObj
}

func (r BetaRlOperationForwardBackwardParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationForwardBackwardParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationForwardBackwardParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties LossFnInputs, ModelInput are required.
type BetaRlOperationForwardBackwardParamsSample struct {
	// Per-token loss tensors keyed by name. Include `target_tokens` and the inputs
	// required by the selected loss. Each tensor must declare `int64` or `float32`, be
	// one-dimensional, and have the same length.
	LossFnInputs map[string]TensorDataParam `json:"loss_fn_inputs,omitzero" api:"required"`
	// Model input
	ModelInput ModelInputParam `json:"model_input,omitzero" api:"required"`
	// Optional MoE per-token routing captured at sample time. Replayed on every
	// training operation, so expert selection matches the one used at sample time.
	// Must cover the whole sample, or all but its last token.
	RoutedExperts RoutedExpertsParam `json:"routed_experts,omitzero"`
	paramObj
}

func (r BetaRlOperationForwardBackwardParamsSample) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationForwardBackwardParamsSample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationForwardBackwardParamsSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlOperationOptimStepParams struct {
	// Adam optimizer overrides for this step.
	AdamParams AdamParams `json:"adam_params,omitzero"`
	// Muon optimizer overrides for this step.
	MuonParams MuonParams `json:"muon_params,omitzero"`
	paramObj
}

func (r BetaRlOperationOptimStepParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationOptimStepParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationOptimStepParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlOperationGetCustomForwardBackwardParams struct {
	// Training session ID
	SessionID string `path:"session_id" api:"required" json:"-"`
	paramObj
}

type BetaRlOperationGetForwardBackwardParams struct {
	// Training session ID
	SessionID string `path:"session_id" api:"required" json:"-"`
	paramObj
}

type BetaRlOperationGetInferenceCheckpointParams struct {
	// Training session ID
	SessionID string `path:"session_id" api:"required" json:"-"`
	paramObj
}

type BetaRlOperationGetOptimStepParams struct {
	// Training session ID
	SessionID string `path:"session_id" api:"required" json:"-"`
	paramObj
}

type BetaRlOperationGetSampleParams struct {
	// Training session ID
	SessionID string `path:"session_id" api:"required" json:"-"`
	paramObj
}

type BetaRlOperationGetTrainingCheckpointParams struct {
	// Training session ID
	SessionID string `path:"session_id" api:"required" json:"-"`
	paramObj
}

type BetaRlOperationGetWeightsSyncParams struct {
	// Training session ID
	SessionID string `path:"session_id" api:"required" json:"-"`
	paramObj
}

type BetaRlOperationSampleParams struct {
	// Model inputs to sample from
	ModelInputs []ModelInputParam `json:"model_inputs,omitzero" api:"required"`
	// Number of completions to generate per prompt
	NumSamples param.Opt[int64] `json:"num_samples,omitzero"`
	// When true, also compute teacher-forced log-probabilities for the model input
	// tokens and return them in `SampleResult.prompt_logprobs`.
	PromptLogprobs param.Opt[bool] `json:"prompt_logprobs,omitzero"`
	// When true, capture the mixture-of-experts routing decisions made while
	// generating and return them in `SampledSequence.routed_experts`, so training can
	// reuse the same expert selection. Only available on mixture-of-experts models;
	// ignored otherwise. The captured buffer scales with sequence length, so leave it
	// off unless you replay routing during training.
	ReturnRoutedExperts param.Opt[bool] `json:"return_routed_experts,omitzero"`
	// When true together with `return_routed_experts`, return each routing capture as
	// a backend-owned `object_uri` plus shape instead of inline base64 data. Clients
	// that do not opt in keep the legacy inline response.
	ReturnRoutedExpertsObjectUri param.Opt[bool] `json:"return_routed_experts_object_uri,omitzero"`
	// Number of most likely alternative tokens to return per model input token in
	// `SampleResult.topk_prompt_logprobs`. 0 disables top-k prompt log-probabilities.
	// Maximum 20.
	TopkPromptLogprobs param.Opt[int64] `json:"topk_prompt_logprobs,omitzero"`
	// Optional sampling parameters
	SamplingParams SamplingParams `json:"sampling_params,omitzero"`
	paramObj
}

func (r BetaRlOperationSampleParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationSampleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationSampleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlOperationWeightsSyncParams struct {
	// How updated parameters are made available for sampling. See `WeightSyncType` for
	// accepted values.
	//
	// Any of "WEIGHT_SYNC_TYPE_SYNCHRONOUS", "WEIGHT_SYNC_TYPE_BACKGROUND_PUBLISH",
	// "WEIGHT_SYNC_TYPE_PIPELINE".
	WeightSyncType WeightSyncType `json:"weight_sync_type,omitzero" api:"required"`
	paramObj
}

func (r BetaRlOperationWeightsSyncParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationWeightsSyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationWeightsSyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
