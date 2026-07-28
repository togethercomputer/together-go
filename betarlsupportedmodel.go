// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaRlSupportedModelService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaRlSupportedModelService] method instead.
type BetaRlSupportedModelService struct {
	Options []option.RequestOption
}

// NewBetaRlSupportedModelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaRlSupportedModelService(opts ...option.RequestOption) (r BetaRlSupportedModelService) {
	r = BetaRlSupportedModelService{}
	r.Options = opts
	return
}

// Returns the models supported by the RL service and their limits for
// training/sampling operations.
func (r *BetaRlSupportedModelService) Get(ctx context.Context, opts ...option.RequestOption) (res *RlSupportedModels, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rl/supported-models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A base model supported by the RL service. Per-mode configs are present only when
// the model supports that mode.
type RlSupportedModel struct {
	// Base model identifier to pass as base_model when creating a model resource
	BaseModel string `json:"base_model" api:"required"`
	// Inference config. Set when the model can be provisioned with generator replicas.
	GeneratorConfig RlSupportedModelGeneratorConfig `json:"generator_config"`
	// Training config. Set when the model supports at least one training mode.
	TrainerConfig RlSupportedModelTrainerConfig `json:"trainer_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseModel       respjson.Field
		GeneratorConfig respjson.Field
		TrainerConfig   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModel) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference config. Set when the model can be provisioned with generator replicas.
type RlSupportedModelGeneratorConfig struct {
	// Maximum tokens in a single inference request (prompt + completion)
	ContextLength int64 `json:"context_length" api:"required"`
	// Default sampling parameters used for sample requests.
	SamplingDefaults RlSupportedModelGeneratorConfigSamplingDefaults `json:"sampling_defaults" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContextLength    respjson.Field
		SamplingDefaults respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModelGeneratorConfig) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelGeneratorConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default sampling parameters used for sample requests.
type RlSupportedModelGeneratorConfigSamplingDefaults struct {
	// Number of logprobs to return per token
	Logprobs int64 `json:"logprobs" api:"required"`
	// Maximum tokens generated per completion
	MaxTokens int64 `json:"max_tokens" api:"required"`
	// Number of completions per prompt
	N int64 `json:"n" api:"required"`
	// Sampling temperature
	Temperature float64 `json:"temperature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logprobs    respjson.Field
		MaxTokens   respjson.Field
		N           respjson.Field
		Temperature respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModelGeneratorConfigSamplingDefaults) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelGeneratorConfigSamplingDefaults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Training config. Set when the model supports at least one training mode.
type RlSupportedModelTrainerConfig struct {
	// Full-weight training config. Set when the model supports full-weight training.
	Full RlSupportedModelTrainerConfigFull `json:"full"`
	// LoRA training config. Set when the model supports LoRA training.
	Lora RlSupportedModelTrainerConfigLora `json:"lora"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Full        respjson.Field
		Lora        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModelTrainerConfig) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelTrainerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full-weight training config. Set when the model supports full-weight training.
type RlSupportedModelTrainerConfigFull struct {
	// Maximum global batch size accepted by a forward-backward step
	MaxBatchSize int64 `json:"max_batch_size" api:"required"`
	// Maximum sequence length in tokens
	MaxSeqLength int64 `json:"max_seq_length" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxBatchSize respjson.Field
		MaxSeqLength respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModelTrainerConfigFull) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelTrainerConfigFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LoRA training config. Set when the model supports LoRA training.
type RlSupportedModelTrainerConfigLora struct {
	// Maximum global batch size accepted by a forward-backward step
	MaxBatchSize int64 `json:"max_batch_size" api:"required"`
	// Maximum LoRA rank
	MaxRank int64 `json:"max_rank" api:"required"`
	// Maximum sequence length in tokens
	MaxSeqLength int64 `json:"max_seq_length" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxBatchSize respjson.Field
		MaxRank      respjson.Field
		MaxSeqLength respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModelTrainerConfigLora) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelTrainerConfigLora) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of base models supported by the RL service
type RlSupportedModels struct {
	// Supported base models for RL
	Data []RlSupportedModel `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModels) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModels) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
