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

// Submits a forward operation that will asynchronously run a no-grad forward pass
// and return per-token log-probabilities for each sample.
func (r *BetaRlOperationService) Forward(ctx context.Context, sessionID string, body BetaRlOperationForwardParams, opts ...option.RequestOption) (res *ForwardOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/forward", sessionID)
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
// gradients to update model parameters.
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

// Retrieves the current status and result of a forward operation.
func (r *BetaRlOperationService) GetForward(ctx context.Context, operationID string, query BetaRlOperationGetForwardParams, opts ...option.RequestOption) (res *ForwardOperation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/forward/%s", query.SessionID, operationID)
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

// The properties Advantages, Logprobs are required.
type CispoLossInputsParam struct {
	// Per-token advantages for CISPO
	Advantages LossAdvantagesParam `json:"advantages,omitzero" api:"required"`
	// Log probabilities for CISPO
	Logprobs LossLogprobsParam `json:"logprobs,omitzero" api:"required"`
	paramObj
}

func (r CispoLossInputsParam) MarshalJSON() (data []byte, err error) {
	type shadow CispoLossInputsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CispoLossInputsParam) UnmarshalJSON(data []byte) error {
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
	Status TrainingOperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error TrainingOperationError `json:"error"`
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

type CustomForwardBackwardResult = any

type DType string

const (
	DTypeDTypeUnspecified DType = "D_TYPE_UNSPECIFIED"
	DTypeDTypeInt64       DType = "D_TYPE_INT64"
	DTypeDTypeFloat32     DType = "D_TYPE_FLOAT32"
	DTypeDTypeBfloat16    DType = "D_TYPE_BFLOAT16"
)

// The properties Advantages, Logprobs are required.
type DroLossInputsParam struct {
	// Per-token advantages for DRO
	Advantages LossAdvantagesParam `json:"advantages,omitzero" api:"required"`
	// Log probabilities for DRO
	Logprobs LossLogprobsParam `json:"logprobs,omitzero" api:"required"`
	paramObj
}

func (r DroLossInputsParam) MarshalJSON() (data []byte, err error) {
	type shadow DroLossInputsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DroLossInputsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	Status TrainingOperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error TrainingOperationError `json:"error"`
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

// Result of a forward-backward pass operation
type ForwardBackwardResult struct {
	// Loss value
	Loss float64 `json:"loss" api:"required"`
	// Loss-specific metrics (e.g., KL divergence, clip fraction for GRPO)
	Metrics map[string]float64 `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Loss        respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForwardBackwardResult) RawJSON() string { return r.JSON.raw }
func (r *ForwardBackwardResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Async forward pass operation
type ForwardOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status TrainingOperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error TrainingOperationError `json:"error"`
	// Result on success
	Output ForwardResult `json:"output"`
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
func (r ForwardOperation) RawJSON() string { return r.JSON.raw }
func (r *ForwardOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a forward pass operation
type ForwardResult struct {
	// Per-sample per-token log-probabilities
	Logprobs []ForwardResultLogprob `json:"logprobs" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logprobs    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForwardResult) RawJSON() string { return r.JSON.raw }
func (r *ForwardResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-token log-probabilities from the target model
type ForwardResultLogprob struct {
	// Float array of per-token log probabilities
	Data []float64 `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForwardResultLogprob) RawJSON() string { return r.JSON.raw }
func (r *ForwardResultLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GrpoLossAggregationType string

const (
	GrpoLossAggregationTypeGrpoLossAggregationTypeUnspecified  GrpoLossAggregationType = "GRPO_LOSS_AGGREGATION_TYPE_UNSPECIFIED"
	GrpoLossAggregationTypeGrpoLossAggregationTypeFixedHorizon GrpoLossAggregationType = "GRPO_LOSS_AGGREGATION_TYPE_FIXED_HORIZON"
	GrpoLossAggregationTypeGrpoLossAggregationTypeTokenMean    GrpoLossAggregationType = "GRPO_LOSS_AGGREGATION_TYPE_TOKEN_MEAN"
	GrpoLossAggregationTypeGrpoLossAggregationTypeSequenceMean GrpoLossAggregationType = "GRPO_LOSS_AGGREGATION_TYPE_SEQUENCE_MEAN"
)

// The properties Advantages, Logprobs are required.
type GrpoLossInputsParam struct {
	// Per-token advantages for GRPO
	Advantages LossAdvantagesParam `json:"advantages,omitzero" api:"required"`
	// Log probabilities for GRPO
	Logprobs LossLogprobsParam `json:"logprobs,omitzero" api:"required"`
	// Reference model log probabilities (required if beta > 0)
	ReferenceLogprobs LossLogprobsParam `json:"reference_logprobs,omitzero"`
	paramObj
}

func (r GrpoLossInputsParam) MarshalJSON() (data []byte, err error) {
	type shadow GrpoLossInputsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GrpoLossInputsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// Loss inputs for unclipped importance-sampling policy-gradient updates.
//
// The properties Advantages, Logprobs are required.
type ImportanceSamplingLossInputsParam struct {
	// Per-token advantages for importance sampling
	Advantages LossAdvantagesParam `json:"advantages,omitzero" api:"required"`
	// Log probabilities for importance sampling
	Logprobs LossLogprobsParam `json:"logprobs,omitzero" api:"required"`
	paramObj
}

func (r ImportanceSamplingLossInputsParam) MarshalJSON() (data []byte, err error) {
	type shadow ImportanceSamplingLossInputsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImportanceSamplingLossInputsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Async inference checkpoint operation
type InferenceCheckpointOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status TrainingOperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error TrainingOperationError `json:"error"`
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

// The property Data is required.
type LossAdvantagesParam struct {
	// Float array of per-token advantages
	Data []float64 `json:"data,omitzero" api:"required"`
	// Data type of the float array (D_TYPE_FLOAT32 or D_TYPE_BFLOAT16)
	//
	// Any of "D_TYPE_UNSPECIFIED", "D_TYPE_INT64", "D_TYPE_FLOAT32",
	// "D_TYPE_BFLOAT16".
	Dtype DType `json:"dtype,omitzero"`
	paramObj
}

func (r LossAdvantagesParam) MarshalJSON() (data []byte, err error) {
	type shadow LossAdvantagesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LossAdvantagesParam) UnmarshalJSON(data []byte) error {
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

// Token-level inputs used to compute the loss for one training sample.
//
// The property TargetTokens is required.
type LossInputsParam struct {
	// Target tokens for loss computation
	TargetTokens LossTargetTokensParam `json:"target_tokens,omitzero" api:"required"`
	CispoInputs  CispoLossInputsParam  `json:"cispo_inputs,omitzero"`
	DroInputs    DroLossInputsParam    `json:"dro_inputs,omitzero"`
	// Inputs required when the loss type is GRPO
	GrpoInputs GrpoLossInputsParam `json:"grpo_inputs,omitzero"`
	// Inputs required when the loss type is importance sampling
	ImportanceSamplingInputs ImportanceSamplingLossInputsParam `json:"importance_sampling_inputs,omitzero"`
	PpoInputs                PpoLossInputsParam                `json:"ppo_inputs,omitzero"`
	// Per-token weights (1=compute loss, 0=ignore).
	Weights WeightsParam `json:"weights,omitzero"`
	paramObj
}

func (r LossInputsParam) MarshalJSON() (data []byte, err error) {
	type shadow LossInputsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LossInputsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Data is required.
type LossLogprobsParam struct {
	// Float array of per-token log probabilities
	Data []float64 `json:"data,omitzero" api:"required"`
	// Data type of the float array (D_TYPE_FLOAT32 or D_TYPE_BFLOAT16)
	//
	// Any of "D_TYPE_UNSPECIFIED", "D_TYPE_INT64", "D_TYPE_FLOAT32",
	// "D_TYPE_BFLOAT16".
	Dtype DType `json:"dtype,omitzero"`
	paramObj
}

func (r LossLogprobsParam) MarshalJSON() (data []byte, err error) {
	type shadow LossLogprobsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LossLogprobsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Data is required.
type LossTargetTokensParam struct {
	// Integer array of target tokens
	Data []LossTargetTokensDataUnionParam `json:"data,omitzero" api:"required"`
	// Data type of the integer array
	//
	// Any of "D_TYPE_UNSPECIFIED", "D_TYPE_INT64", "D_TYPE_FLOAT32",
	// "D_TYPE_BFLOAT16".
	Dtype DType `json:"dtype,omitzero"`
	paramObj
}

func (r LossTargetTokensParam) MarshalJSON() (data []byte, err error) {
	type shadow LossTargetTokensParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LossTargetTokensParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LossTargetTokensDataUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u LossTargetTokensDataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LossTargetTokensDataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LossTargetTokensDataUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
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

// Async optimizer step operation
type OptimStepOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status TrainingOperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error TrainingOperationError `json:"error"`
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

// ToParam converts this PolicyVersionSegment to a PolicyVersionSegmentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicyVersionSegmentParam.Overrides()
func (r PolicyVersionSegment) ToParam() PolicyVersionSegmentParam {
	return param.Override[PolicyVersionSegmentParam](json.RawMessage(r.RawJSON()))
}

// A (policy version, starting token) span within a sampled sequence. Version 0 is
// the initial model; each optim_step call increments the version by 1.
//
// The properties StartToken, Version are required.
type PolicyVersionSegmentParam struct {
	// Index of the first token of this segment within the sampled sequence. Always 0
	// for the first segment.
	StartToken int64 `json:"start_token" api:"required"`
	// Model version under which this segment of tokens was generated
	Version int64 `json:"version" api:"required"`
	paramObj
}

func (r PolicyVersionSegmentParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyVersionSegmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyVersionSegmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Advantages, Logprobs are required.
type PpoLossInputsParam struct {
	// Per-token advantages for PPO
	Advantages LossAdvantagesParam `json:"advantages,omitzero" api:"required"`
	// Log probabilities for PPO
	Logprobs LossLogprobsParam `json:"logprobs,omitzero" api:"required"`
	paramObj
}

func (r PpoLossInputsParam) MarshalJSON() (data []byte, err error) {
	type shadow PpoLossInputsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PpoLossInputsParam) UnmarshalJSON(data []byte) error {
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

// Async sample operation
type SampleOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status TrainingOperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error TrainingOperationError `json:"error"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PromptCacheHitTokens respjson.Field
		StopReason           respjson.Field
		Tokens               respjson.Field
		Logprobs             respjson.Field
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
	// Random seed for reproducibility
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

// Async save training checkpoint operation
type TrainingCheckpointOperation struct {
	// Operation ID
	ID string `json:"id" api:"required"`
	// Operation status
	//
	// Any of "TRAINING_OPERATION_STATUS_UNSPECIFIED",
	// "TRAINING_OPERATION_STATUS_PENDING", "TRAINING_OPERATION_STATUS_RUNNING",
	// "TRAINING_OPERATION_STATUS_COMPLETED", "TRAINING_OPERATION_STATUS_FAILED".
	Status TrainingOperationStatus `json:"status" api:"required"`
	// Error details on failure
	Error TrainingOperationError `json:"error"`
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

// Error details for a failed training operation
type TrainingOperationError struct {
	// Application error code
	//
	// Any of "TRAINING_OPERATION_ERROR_CODE_UNSPECIFIED",
	// "TRAINING_OPERATION_ERROR_CODE_RESOURCE_EXHAUSTED",
	// "TRAINING_OPERATION_ERROR_CODE_TIMEOUT",
	// "TRAINING_OPERATION_ERROR_CODE_INTERNAL_ERROR",
	// "TRAINING_OPERATION_ERROR_CODE_SESSION_NOT_ACTIVE",
	// "TRAINING_OPERATION_ERROR_CODE_INVALID_INPUT".
	Code TrainingOperationErrorCode `json:"code"`
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
func (r TrainingOperationError) RawJSON() string { return r.JSON.raw }
func (r *TrainingOperationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrainingOperationErrorCode string

const (
	TrainingOperationErrorCodeTrainingOperationErrorCodeUnspecified       TrainingOperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_UNSPECIFIED"
	TrainingOperationErrorCodeTrainingOperationErrorCodeResourceExhausted TrainingOperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_RESOURCE_EXHAUSTED"
	TrainingOperationErrorCodeTrainingOperationErrorCodeTimeout           TrainingOperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_TIMEOUT"
	TrainingOperationErrorCodeTrainingOperationErrorCodeInternalError     TrainingOperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_INTERNAL_ERROR"
	TrainingOperationErrorCodeTrainingOperationErrorCodeSessionNotActive  TrainingOperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_SESSION_NOT_ACTIVE"
	TrainingOperationErrorCodeTrainingOperationErrorCodeInvalidInput      TrainingOperationErrorCode = "TRAINING_OPERATION_ERROR_CODE_INVALID_INPUT"
)

type TrainingOperationStatus string

const (
	TrainingOperationStatusTrainingOperationStatusUnspecified TrainingOperationStatus = "TRAINING_OPERATION_STATUS_UNSPECIFIED"
	TrainingOperationStatusTrainingOperationStatusPending     TrainingOperationStatus = "TRAINING_OPERATION_STATUS_PENDING"
	TrainingOperationStatusTrainingOperationStatusRunning     TrainingOperationStatus = "TRAINING_OPERATION_STATUS_RUNNING"
	TrainingOperationStatusTrainingOperationStatusCompleted   TrainingOperationStatus = "TRAINING_OPERATION_STATUS_COMPLETED"
	TrainingOperationStatusTrainingOperationStatusFailed      TrainingOperationStatus = "TRAINING_OPERATION_STATUS_FAILED"
)

// How the trainer's updated weights are propagated to the generator after an
// optimizer step. SYNCHRONOUS publishes inline before the step returns;
// BACKGROUND_PUBLISH returns immediately and publishes once in-flight rollouts
// drain; PIPELINE overlaps the publish with in-flight rollouts so they continue
// decoding under the new weights.
type WeightSyncType string

const (
	WeightSyncTypeWeightSyncTypeUnspecified       WeightSyncType = "WEIGHT_SYNC_TYPE_UNSPECIFIED"
	WeightSyncTypeWeightSyncTypeSynchronous       WeightSyncType = "WEIGHT_SYNC_TYPE_SYNCHRONOUS"
	WeightSyncTypeWeightSyncTypeBackgroundPublish WeightSyncType = "WEIGHT_SYNC_TYPE_BACKGROUND_PUBLISH"
	WeightSyncTypeWeightSyncTypePipeline          WeightSyncType = "WEIGHT_SYNC_TYPE_PIPELINE"
)

// The property Data is required.
type WeightsParam struct {
	// Per-token weights: 1 to include the token in the loss, 0 to ignore it.
	Data []WeightsDataUnionParam `json:"data,omitzero" api:"required"`
	// Data type of the integer array (must be D_TYPE_INT64)
	//
	// Any of "D_TYPE_UNSPECIFIED", "D_TYPE_INT64", "D_TYPE_FLOAT32",
	// "D_TYPE_BFLOAT16".
	Dtype DType `json:"dtype,omitzero"`
	paramObj
}

func (r WeightsParam) MarshalJSON() (data []byte, err error) {
	type shadow WeightsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WeightsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WeightsDataUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WeightsDataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WeightsDataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WeightsDataUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
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

// The properties LossInputs, ModelInput, PolicySegments are required.
type BetaRlOperationCustomForwardBackwardParamsSample struct {
	// Loss function inputs
	LossInputs LossInputsParam `json:"loss_inputs,omitzero" api:"required"`
	// Model input
	ModelInput ModelInputParam `json:"model_input,omitzero" api:"required"`
	// Policy versions associated with this sample's tokens
	PolicySegments []PolicyVersionSegmentParam `json:"policy_segments,omitzero" api:"required"`
	paramObj
}

func (r BetaRlOperationCustomForwardBackwardParamsSample) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationCustomForwardBackwardParamsSample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationCustomForwardBackwardParamsSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlOperationForwardParams struct {
	// Batch of training samples for which to compute per-token log-probabilities
	Samples []BetaRlOperationForwardParamsSample `json:"samples,omitzero" api:"required"`
	paramObj
}

func (r BetaRlOperationForwardParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationForwardParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationForwardParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties LossInputs, ModelInput, PolicySegments are required.
type BetaRlOperationForwardParamsSample struct {
	// Loss function inputs
	LossInputs LossInputsParam `json:"loss_inputs,omitzero" api:"required"`
	// Model input
	ModelInput ModelInputParam `json:"model_input,omitzero" api:"required"`
	// Policy versions associated with this sample's tokens
	PolicySegments []PolicyVersionSegmentParam `json:"policy_segments,omitzero" api:"required"`
	paramObj
}

func (r BetaRlOperationForwardParamsSample) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationForwardParamsSample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationForwardParamsSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlOperationForwardBackwardParams struct {
	// Loss function configuration
	Loss LossConfigParam `json:"loss,omitzero" api:"required"`
	// Batch of training samples to process
	Samples []BetaRlOperationForwardBackwardParamsSample `json:"samples,omitzero" api:"required"`
	paramObj
}

func (r BetaRlOperationForwardBackwardParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlOperationForwardBackwardParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlOperationForwardBackwardParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties LossInputs, ModelInput, PolicySegments are required.
type BetaRlOperationForwardBackwardParamsSample struct {
	// Loss function inputs
	LossInputs LossInputsParam `json:"loss_inputs,omitzero" api:"required"`
	// Model input
	ModelInput ModelInputParam `json:"model_input,omitzero" api:"required"`
	// Policy versions associated with this sample's tokens
	PolicySegments []PolicyVersionSegmentParam `json:"policy_segments,omitzero" api:"required"`
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
	// How the trainer's updated weights are propagated to the generator after this
	// optimizer step. See `WeightSyncType` for accepted values.
	//
	// Any of "WEIGHT_SYNC_TYPE_UNSPECIFIED", "WEIGHT_SYNC_TYPE_SYNCHRONOUS",
	// "WEIGHT_SYNC_TYPE_BACKGROUND_PUBLISH", "WEIGHT_SYNC_TYPE_PIPELINE".
	WeightSyncType WeightSyncType `json:"weight_sync_type,omitzero" api:"required"`
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

type BetaRlOperationGetForwardParams struct {
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

type BetaRlOperationSampleParams struct {
	// Model inputs to sample from
	ModelInputs []ModelInputParam `json:"model_inputs,omitzero" api:"required"`
	// Number of completions to generate per prompt
	NumSamples param.Opt[int64] `json:"num_samples,omitzero"`
	// When true, also compute teacher-forced log-probabilities for the model input
	// tokens and return them in `SampleResult.prompt_logprobs`.
	PromptLogprobs param.Opt[bool] `json:"prompt_logprobs,omitzero"`
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
