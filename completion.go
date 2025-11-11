// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
	"github.com/togethercomputer/together-go/packages/ssestream"
)

// CompletionService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompletionService] method instead.
type CompletionService struct {
	Options []option.RequestOption
}

// NewCompletionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCompletionService(opts ...option.RequestOption) (r CompletionService) {
	r = CompletionService{}
	r.Options = opts
	return
}

// Query a language, code, or image model.
func (r *CompletionService) New(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (res *Completion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Query a language, code, or image model.
func (r *CompletionService) NewStreaming(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[CompletionChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[CompletionChunk](ssestream.NewDecoder(raw), err)
}

type Completion struct {
	ID      string             `json:"id,required"`
	Choices []CompletionChoice `json:"choices,required"`
	Created int64              `json:"created,required"`
	Model   string             `json:"model,required"`
	// Any of "text.completion".
	Object CompletionObject    `json:"object,required"`
	Usage  ChatCompletionUsage `json:"usage,required"`
	Prompt []CompletionPrompt  `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Choices     respjson.Field
		Created     respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		Usage       respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Completion) RawJSON() string { return r.JSON.raw }
func (r *Completion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoice struct {
	// Any of "stop", "eos", "length", "tool_calls", "function_call".
	FinishReason string   `json:"finish_reason"`
	Logprobs     LogProbs `json:"logprobs"`
	Seed         int64    `json:"seed"`
	Text         string   `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Logprobs     respjson.Field
		Seed         respjson.Field
		Text         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoice) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionObject string

const (
	CompletionObjectTextCompletion CompletionObject = "text.completion"
)

type CompletionPrompt struct {
	Logprobs LogProbs `json:"logprobs"`
	Text     string   `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logprobs    respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionPrompt) RawJSON() string { return r.JSON.raw }
func (r *CompletionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChunk struct {
	ID      string                  `json:"id,required"`
	Token   CompletionChunkToken    `json:"token,required"`
	Choices []CompletionChunkChoice `json:"choices,required"`
	// Any of "stop", "eos", "length", "tool_calls", "function_call".
	FinishReason CompletionChunkFinishReason `json:"finish_reason,required"`
	Usage        ChatCompletionUsage         `json:"usage,required"`
	Created      int64                       `json:"created"`
	// Any of "completion.chunk".
	Object CompletionChunkObject `json:"object"`
	Seed   int64                 `json:"seed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Token        respjson.Field
		Choices      respjson.Field
		FinishReason respjson.Field
		Usage        respjson.Field
		Created      respjson.Field
		Object       respjson.Field
		Seed         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChunk) RawJSON() string { return r.JSON.raw }
func (r *CompletionChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChunkToken struct {
	ID      int64   `json:"id,required"`
	Logprob float64 `json:"logprob,required"`
	Special bool    `json:"special,required"`
	Text    string  `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Logprob     respjson.Field
		Special     respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChunkToken) RawJSON() string { return r.JSON.raw }
func (r *CompletionChunkToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChunkChoice struct {
	Index int64                      `json:"index,required"`
	Delta CompletionChunkChoiceDelta `json:"delta"`
	Text  string                     `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		Delta       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChunkChoice) RawJSON() string { return r.JSON.raw }
func (r *CompletionChunkChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChunkChoiceDelta struct {
	// Any of "system", "user", "assistant", "function", "tool".
	Role    string `json:"role,required"`
	Content string `json:"content,nullable"`
	// Deprecated: deprecated
	FunctionCall CompletionChunkChoiceDeltaFunctionCall `json:"function_call,nullable"`
	Reasoning    string                                 `json:"reasoning,nullable"`
	TokenID      int64                                  `json:"token_id"`
	ToolCalls    []ToolChoice                           `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role         respjson.Field
		Content      respjson.Field
		FunctionCall respjson.Field
		Reasoning    respjson.Field
		TokenID      respjson.Field
		ToolCalls    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChunkChoiceDelta) RawJSON() string { return r.JSON.raw }
func (r *CompletionChunkChoiceDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated: deprecated
type CompletionChunkChoiceDeltaFunctionCall struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChunkChoiceDeltaFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *CompletionChunkChoiceDeltaFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChunkFinishReason string

const (
	CompletionChunkFinishReasonStop         CompletionChunkFinishReason = "stop"
	CompletionChunkFinishReasonEos          CompletionChunkFinishReason = "eos"
	CompletionChunkFinishReasonLength       CompletionChunkFinishReason = "length"
	CompletionChunkFinishReasonToolCalls    CompletionChunkFinishReason = "tool_calls"
	CompletionChunkFinishReasonFunctionCall CompletionChunkFinishReason = "function_call"
)

type CompletionChunkObject string

const (
	CompletionChunkObjectCompletionChunk CompletionChunkObject = "completion.chunk"
)

type LogProbs struct {
	// List of token IDs corresponding to the logprobs
	TokenIDs []float64 `json:"token_ids"`
	// List of token log probabilities
	TokenLogprobs []float64 `json:"token_logprobs"`
	// List of token strings
	Tokens []string `json:"tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TokenIDs      respjson.Field
		TokenLogprobs respjson.Field
		Tokens        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogProbs) RawJSON() string { return r.JSON.raw }
func (r *LogProbs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolChoice struct {
	ID       string             `json:"id,required"`
	Function ToolChoiceFunction `json:"function,required"`
	Index    float64            `json:"index,required"`
	// Any of "function".
	Type ToolChoiceType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolChoice) RawJSON() string { return r.JSON.raw }
func (r *ToolChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolChoice to a ToolChoiceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolChoiceParam.Overrides()
func (r ToolChoice) ToParam() ToolChoiceParam {
	return param.Override[ToolChoiceParam](json.RawMessage(r.RawJSON()))
}

type ToolChoiceFunction struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolChoiceFunction) RawJSON() string { return r.JSON.raw }
func (r *ToolChoiceFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolChoiceType string

const (
	ToolChoiceTypeFunction ToolChoiceType = "function"
)

// The properties ID, Function, Index, Type are required.
type ToolChoiceParam struct {
	ID       string                  `json:"id,required"`
	Function ToolChoiceFunctionParam `json:"function,omitzero,required"`
	Index    float64                 `json:"index,required"`
	// Any of "function".
	Type ToolChoiceType `json:"type,omitzero,required"`
	paramObj
}

func (r ToolChoiceParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolChoiceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolChoiceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Arguments, Name are required.
type ToolChoiceFunctionParam struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	paramObj
}

func (r ToolChoiceFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolChoiceFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolChoiceFunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolsParam struct {
	Type     param.Opt[string]  `json:"type,omitzero"`
	Function ToolsFunctionParam `json:"function,omitzero"`
	paramObj
}

func (r ToolsParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolsFunctionParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	// A map of parameter names to their values.
	Parameters map[string]any `json:"parameters,omitzero"`
	paramObj
}

func (r ToolsFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolsFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolsFunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionNewParams struct {
	// The name of the model to query.
	//
	// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#chat-models)
	Model CompletionNewParamsModel `json:"model,omitzero,required"`
	// A string providing context for the model to complete.
	Prompt string `json:"prompt,required"`
	// If true, the response will contain the prompt. Can be used with `logprobs` to
	// return prompt logprobs.
	Echo param.Opt[bool] `json:"echo,omitzero"`
	// A number between -2.0 and 2.0 where a positive value decreases the likelihood of
	// repeating tokens that have already been mentioned.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// An integer between 0 and 20 of the top k tokens to return log probabilities for
	// at each generation step, instead of just the sampled token. Log probabilities
	// help assess model confidence in token predictions.
	Logprobs param.Opt[int64] `json:"logprobs,omitzero"`
	// The maximum number of tokens to generate.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// A number between 0 and 1 that can be used as an alternative to top-p and top-k.
	MinP param.Opt[float64] `json:"min_p,omitzero"`
	// The number of completions to generate for each prompt.
	N param.Opt[int64] `json:"n,omitzero"`
	// A number between -2.0 and 2.0 where a positive value increases the likelihood of
	// a model talking about new topics.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// A number that controls the diversity of generated text by reducing the
	// likelihood of repeated sequences. Higher values decrease repetition.
	RepetitionPenalty param.Opt[float64] `json:"repetition_penalty,omitzero"`
	// Seed value for reproducibility.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// A decimal number from 0-1 that determines the degree of randomness in the
	// response. A temperature less than 1 favors more correctness and is appropriate
	// for question answering or summarization. A value closer to 1 introduces more
	// randomness in the output.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// An integer that's used to limit the number of choices for the next predicted
	// word or token. It specifies the maximum number of tokens to consider at each
	// step, based on their probability of occurrence. This technique helps to speed up
	// the generation process and can improve the quality of the generated text by
	// focusing on the most likely options.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// A percentage (also called the nucleus parameter) that's used to dynamically
	// adjust the number of choices for each predicted token based on the cumulative
	// probabilities. It specifies a probability threshold below which all less likely
	// tokens are filtered out. This technique helps maintain diversity and generate
	// more fluent and natural-sounding text.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Adjusts the likelihood of specific tokens appearing in the generated output.
	LogitBias map[string]float64 `json:"logit_bias,omitzero"`
	// The name of the moderation model used to validate tokens. Choose from the
	// available moderation models found
	// [here](https://docs.together.ai/docs/inference-models#moderation-models).
	SafetyModel CompletionNewParamsSafetyModel `json:"safety_model,omitzero"`
	// A list of string sequences that will truncate (stop) inference text output. For
	// example, "</s>" will stop generation as soon as the model generates the given
	// token.
	Stop []string `json:"stop,omitzero"`
	paramObj
}

func (r CompletionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CompletionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The name of the model to query.
//
// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#chat-models)
type CompletionNewParamsModel string

const (
	CompletionNewParamsModelMetaLlamaLlama2_70bHf    CompletionNewParamsModel = "meta-llama/Llama-2-70b-hf"
	CompletionNewParamsModelMistralaiMistral7BV0_1   CompletionNewParamsModel = "mistralai/Mistral-7B-v0.1"
	CompletionNewParamsModelMistralaiMixtral8x7BV0_1 CompletionNewParamsModel = "mistralai/Mixtral-8x7B-v0.1"
	CompletionNewParamsModelMetaLlamaLlamaGuard7b    CompletionNewParamsModel = "Meta-Llama/Llama-Guard-7b"
)

// The name of the moderation model used to validate tokens. Choose from the
// available moderation models found
// [here](https://docs.together.ai/docs/inference-models#moderation-models).
type CompletionNewParamsSafetyModel string

const (
	CompletionNewParamsSafetyModelMetaLlamaLlamaGuard7b CompletionNewParamsSafetyModel = "Meta-Llama/Llama-Guard-7b"
)
