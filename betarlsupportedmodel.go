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
	// GPU type used when model-resource creation omits gpu_type.
	//
	// Any of "H100-80GB", "B200-SXM".
	DefaultGPUType RlSupportedModelDefaultGPUType `json:"default_gpu_type" api:"required"`
	// Validated GPU configurations available for this base model.
	ComputeConfigs []RlSupportedModelComputeConfig `json:"compute_configs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseModel      respjson.Field
		DefaultGPUType respjson.Field
		ComputeConfigs respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModel) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU type used when model-resource creation omits gpu_type.
type RlSupportedModelDefaultGPUType string

const (
	RlSupportedModelDefaultGPUTypeH100_80GB RlSupportedModelDefaultGPUType = "H100-80GB"
	RlSupportedModelDefaultGPUTypeB200Sxm   RlSupportedModelDefaultGPUType = "B200-SXM"
)

// A validated hardware configuration available for an RL base model.
type RlSupportedModelComputeConfig struct {
	// GPU type this configuration provisions.
	//
	// Any of "H100-80GB", "B200-SXM".
	GPUType string `json:"gpu_type" api:"required"`
	// Inference config for this GPU type. Set when the model can be provisioned with
	// generator replicas on this GPU type.
	GeneratorConfig RlSupportedModelComputeConfigGeneratorConfig `json:"generator_config"`
	// Training config for this GPU type. Set when the model supports at least one
	// training mode on this GPU type.
	TrainerConfig RlSupportedModelComputeConfigTrainerConfig `json:"trainer_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GPUType         respjson.Field
		GeneratorConfig respjson.Field
		TrainerConfig   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModelComputeConfig) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelComputeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference config for this GPU type. Set when the model can be provisioned with
// generator replicas on this GPU type.
type RlSupportedModelComputeConfigGeneratorConfig struct {
	// Maximum tokens in a single inference request (prompt + completion)
	ContextLength int64 `json:"context_length" api:"required"`
	// Default sampling parameters used for sample requests.
	SamplingDefaults RlSupportedModelComputeConfigGeneratorConfigSamplingDefaults `json:"sampling_defaults" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContextLength    respjson.Field
		SamplingDefaults respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModelComputeConfigGeneratorConfig) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelComputeConfigGeneratorConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default sampling parameters used for sample requests.
type RlSupportedModelComputeConfigGeneratorConfigSamplingDefaults struct {
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
func (r RlSupportedModelComputeConfigGeneratorConfigSamplingDefaults) RawJSON() string {
	return r.JSON.raw
}
func (r *RlSupportedModelComputeConfigGeneratorConfigSamplingDefaults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Training config for this GPU type. Set when the model supports at least one
// training mode on this GPU type.
type RlSupportedModelComputeConfigTrainerConfig struct {
	// Full-weight training config. Set when the model supports full-weight training.
	Full RlSupportedModelComputeConfigTrainerConfigFull `json:"full"`
	// LoRA training config. Set when the model supports LoRA training.
	Lora RlSupportedModelComputeConfigTrainerConfigLora `json:"lora"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Full        respjson.Field
		Lora        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RlSupportedModelComputeConfigTrainerConfig) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelComputeConfigTrainerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full-weight training config. Set when the model supports full-weight training.
type RlSupportedModelComputeConfigTrainerConfigFull struct {
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
func (r RlSupportedModelComputeConfigTrainerConfigFull) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelComputeConfigTrainerConfigFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LoRA training config. Set when the model supports LoRA training.
type RlSupportedModelComputeConfigTrainerConfigLora struct {
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
func (r RlSupportedModelComputeConfigTrainerConfigLora) RawJSON() string { return r.JSON.raw }
func (r *RlSupportedModelComputeConfigTrainerConfigLora) UnmarshalJSON(data []byte) error {
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
