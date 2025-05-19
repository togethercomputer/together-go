// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
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
func NewCompletionService(opts ...option.RequestOption) (r *CompletionService) {
	r = &CompletionService{}
	r.Options = opts
	return
}

// Query a language, code, or image model.
func (r *CompletionService) New(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (res *Completion, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[CompletionChunk](ssestream.NewDecoder(raw), err)
}

type Completion struct {
	ID      string              `json:"id,required"`
	Choices []CompletionChoice  `json:"choices,required"`
	Created int64               `json:"created,required"`
	Model   string              `json:"model,required"`
	Object  CompletionObject    `json:"object,required"`
	Usage   ChatCompletionUsage `json:"usage,required,nullable"`
	Prompt  []CompletionPrompt  `json:"prompt"`
	JSON    completionJSON      `json:"-"`
}

// completionJSON contains the JSON metadata for the struct [Completion]
type completionJSON struct {
	ID          apijson.Field
	Choices     apijson.Field
	Created     apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	Usage       apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Completion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionJSON) RawJSON() string {
	return r.raw
}

type CompletionChoice struct {
	FinishReason CompletionChoicesFinishReason `json:"finish_reason"`
	Logprobs     LogProbs                      `json:"logprobs"`
	Seed         int64                         `json:"seed"`
	Text         string                        `json:"text"`
	JSON         completionChoiceJSON          `json:"-"`
}

// completionChoiceJSON contains the JSON metadata for the struct
// [CompletionChoice]
type completionChoiceJSON struct {
	FinishReason apijson.Field
	Logprobs     apijson.Field
	Seed         apijson.Field
	Text         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoiceJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesFinishReason string

const (
	CompletionChoicesFinishReasonStop         CompletionChoicesFinishReason = "stop"
	CompletionChoicesFinishReasonEos          CompletionChoicesFinishReason = "eos"
	CompletionChoicesFinishReasonLength       CompletionChoicesFinishReason = "length"
	CompletionChoicesFinishReasonToolCalls    CompletionChoicesFinishReason = "tool_calls"
	CompletionChoicesFinishReasonFunctionCall CompletionChoicesFinishReason = "function_call"
)

func (r CompletionChoicesFinishReason) IsKnown() bool {
	switch r {
	case CompletionChoicesFinishReasonStop, CompletionChoicesFinishReasonEos, CompletionChoicesFinishReasonLength, CompletionChoicesFinishReasonToolCalls, CompletionChoicesFinishReasonFunctionCall:
		return true
	}
	return false
}

type CompletionObject string

const (
	CompletionObjectTextCompletion CompletionObject = "text_completion"
)

func (r CompletionObject) IsKnown() bool {
	switch r {
	case CompletionObjectTextCompletion:
		return true
	}
	return false
}

type CompletionPrompt struct {
	Logprobs LogProbs             `json:"logprobs"`
	Text     string               `json:"text"`
	JSON     completionPromptJSON `json:"-"`
}

// completionPromptJSON contains the JSON metadata for the struct
// [CompletionPrompt]
type completionPromptJSON struct {
	Logprobs    apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionPromptJSON) RawJSON() string {
	return r.raw
}

type CompletionChunk struct {
	ID           string                      `json:"id,required"`
	Token        CompletionChunkToken        `json:"token,required"`
	Choices      []CompletionChunkChoice     `json:"choices,required"`
	FinishReason CompletionChunkFinishReason `json:"finish_reason,required,nullable"`
	Usage        ChatCompletionUsage         `json:"usage,required,nullable"`
	Created      int64                       `json:"created"`
	Object       CompletionChunkObject       `json:"object"`
	Seed         int64                       `json:"seed"`
	JSON         completionChunkJSON         `json:"-"`
}

// completionChunkJSON contains the JSON metadata for the struct [CompletionChunk]
type completionChunkJSON struct {
	ID           apijson.Field
	Token        apijson.Field
	Choices      apijson.Field
	FinishReason apijson.Field
	Usage        apijson.Field
	Created      apijson.Field
	Object       apijson.Field
	Seed         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChunkJSON) RawJSON() string {
	return r.raw
}

type CompletionChunkToken struct {
	ID      int64                    `json:"id,required"`
	Logprob float64                  `json:"logprob,required"`
	Special bool                     `json:"special,required"`
	Text    string                   `json:"text,required"`
	JSON    completionChunkTokenJSON `json:"-"`
}

// completionChunkTokenJSON contains the JSON metadata for the struct
// [CompletionChunkToken]
type completionChunkTokenJSON struct {
	ID          apijson.Field
	Logprob     apijson.Field
	Special     apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChunkToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChunkTokenJSON) RawJSON() string {
	return r.raw
}

type CompletionChunkChoice struct {
	Index int64                       `json:"index,required"`
	Delta CompletionChunkChoicesDelta `json:"delta"`
	Text  string                      `json:"text"`
	JSON  completionChunkChoiceJSON   `json:"-"`
}

// completionChunkChoiceJSON contains the JSON metadata for the struct
// [CompletionChunkChoice]
type completionChunkChoiceJSON struct {
	Index       apijson.Field
	Delta       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChunkChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChunkChoiceJSON) RawJSON() string {
	return r.raw
}

type CompletionChunkChoicesDelta struct {
	Role    CompletionChunkChoicesDeltaRole `json:"role,required"`
	Content string                          `json:"content,nullable"`
	// Deprecated: deprecated
	FunctionCall CompletionChunkChoicesDeltaFunctionCall `json:"function_call,nullable"`
	TokenID      int64                                   `json:"token_id"`
	ToolCalls    []ToolChoice                            `json:"tool_calls"`
	JSON         completionChunkChoicesDeltaJSON         `json:"-"`
}

// completionChunkChoicesDeltaJSON contains the JSON metadata for the struct
// [CompletionChunkChoicesDelta]
type completionChunkChoicesDeltaJSON struct {
	Role         apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	TokenID      apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionChunkChoicesDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChunkChoicesDeltaJSON) RawJSON() string {
	return r.raw
}

type CompletionChunkChoicesDeltaRole string

const (
	CompletionChunkChoicesDeltaRoleSystem    CompletionChunkChoicesDeltaRole = "system"
	CompletionChunkChoicesDeltaRoleUser      CompletionChunkChoicesDeltaRole = "user"
	CompletionChunkChoicesDeltaRoleAssistant CompletionChunkChoicesDeltaRole = "assistant"
	CompletionChunkChoicesDeltaRoleFunction  CompletionChunkChoicesDeltaRole = "function"
	CompletionChunkChoicesDeltaRoleTool      CompletionChunkChoicesDeltaRole = "tool"
)

func (r CompletionChunkChoicesDeltaRole) IsKnown() bool {
	switch r {
	case CompletionChunkChoicesDeltaRoleSystem, CompletionChunkChoicesDeltaRoleUser, CompletionChunkChoicesDeltaRoleAssistant, CompletionChunkChoicesDeltaRoleFunction, CompletionChunkChoicesDeltaRoleTool:
		return true
	}
	return false
}

// Deprecated: deprecated
type CompletionChunkChoicesDeltaFunctionCall struct {
	Arguments string                                      `json:"arguments,required"`
	Name      string                                      `json:"name,required"`
	JSON      completionChunkChoicesDeltaFunctionCallJSON `json:"-"`
}

// completionChunkChoicesDeltaFunctionCallJSON contains the JSON metadata for the
// struct [CompletionChunkChoicesDeltaFunctionCall]
type completionChunkChoicesDeltaFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChunkChoicesDeltaFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChunkChoicesDeltaFunctionCallJSON) RawJSON() string {
	return r.raw
}

type CompletionChunkFinishReason string

const (
	CompletionChunkFinishReasonStop         CompletionChunkFinishReason = "stop"
	CompletionChunkFinishReasonEos          CompletionChunkFinishReason = "eos"
	CompletionChunkFinishReasonLength       CompletionChunkFinishReason = "length"
	CompletionChunkFinishReasonToolCalls    CompletionChunkFinishReason = "tool_calls"
	CompletionChunkFinishReasonFunctionCall CompletionChunkFinishReason = "function_call"
)

func (r CompletionChunkFinishReason) IsKnown() bool {
	switch r {
	case CompletionChunkFinishReasonStop, CompletionChunkFinishReasonEos, CompletionChunkFinishReasonLength, CompletionChunkFinishReasonToolCalls, CompletionChunkFinishReasonFunctionCall:
		return true
	}
	return false
}

type CompletionChunkObject string

const (
	CompletionChunkObjectCompletionChunk CompletionChunkObject = "completion.chunk"
)

func (r CompletionChunkObject) IsKnown() bool {
	switch r {
	case CompletionChunkObjectCompletionChunk:
		return true
	}
	return false
}

type LogProbs struct {
	// List of token IDs corresponding to the logprobs
	TokenIDs []float64 `json:"token_ids"`
	// List of token log probabilities
	TokenLogprobs []float64 `json:"token_logprobs"`
	// List of token strings
	Tokens []string     `json:"tokens"`
	JSON   logProbsJSON `json:"-"`
}

// logProbsJSON contains the JSON metadata for the struct [LogProbs]
type logProbsJSON struct {
	TokenIDs      apijson.Field
	TokenLogprobs apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LogProbs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logProbsJSON) RawJSON() string {
	return r.raw
}

type ToolChoice struct {
	ID       string             `json:"id,required"`
	Function ToolChoiceFunction `json:"function,required"`
	Index    float64            `json:"index,required"`
	Type     ToolChoiceType     `json:"type,required"`
	JSON     toolChoiceJSON     `json:"-"`
}

// toolChoiceJSON contains the JSON metadata for the struct [ToolChoice]
type toolChoiceJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Index       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolChoiceJSON) RawJSON() string {
	return r.raw
}

type ToolChoiceFunction struct {
	Arguments string                 `json:"arguments,required"`
	Name      string                 `json:"name,required"`
	JSON      toolChoiceFunctionJSON `json:"-"`
}

// toolChoiceFunctionJSON contains the JSON metadata for the struct
// [ToolChoiceFunction]
type toolChoiceFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolChoiceFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolChoiceFunctionJSON) RawJSON() string {
	return r.raw
}

type ToolChoiceType string

const (
	ToolChoiceTypeFunction ToolChoiceType = "function"
)

func (r ToolChoiceType) IsKnown() bool {
	switch r {
	case ToolChoiceTypeFunction:
		return true
	}
	return false
}

type ToolChoiceParam struct {
	ID       param.Field[string]                  `json:"id,required"`
	Function param.Field[ToolChoiceFunctionParam] `json:"function,required"`
	Index    param.Field[float64]                 `json:"index,required"`
	Type     param.Field[ToolChoiceType]          `json:"type,required"`
}

func (r ToolChoiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceParam) ImplementsChatCompletionNewParamsToolChoiceUnion() {}

type ToolChoiceFunctionParam struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r ToolChoiceFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolsParam struct {
	Function param.Field[ToolsFunctionParam] `json:"function"`
	Type     param.Field[string]             `json:"type"`
}

func (r ToolsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolsFunctionParam struct {
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
	// A map of parameter names to their values.
	Parameters param.Field[map[string]interface{}] `json:"parameters"`
}

func (r ToolsFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CompletionNewParams struct {
	// The name of the model to query.
	//
	// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#chat-models)
	Model param.Field[CompletionNewParamsModel] `json:"model,required"`
	// A string providing context for the model to complete.
	Prompt param.Field[string] `json:"prompt,required"`
	// If true, the response will contain the prompt. Can be used with `logprobs` to
	// return prompt logprobs.
	Echo param.Field[bool] `json:"echo"`
	// A number between -2.0 and 2.0 where a positive value decreases the likelihood of
	// repeating tokens that have already been mentioned.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Adjusts the likelihood of specific tokens appearing in the generated output.
	LogitBias param.Field[map[string]float64] `json:"logit_bias"`
	// Integer (0 or 1) that controls whether log probabilities of generated tokens are
	// returned. Log probabilities help assess model confidence in token predictions.
	Logprobs param.Field[int64] `json:"logprobs"`
	// The maximum number of tokens to generate.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// A number between 0 and 1 that can be used as an alternative to top-p and top-k.
	MinP param.Field[float64] `json:"min_p"`
	// The number of completions to generate for each prompt.
	N param.Field[int64] `json:"n"`
	// A number between -2.0 and 2.0 where a positive value increases the likelihood of
	// a model talking about new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// A number that controls the diversity of generated text by reducing the
	// likelihood of repeated sequences. Higher values decrease repetition.
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	// The name of the moderation model used to validate tokens. Choose from the
	// available moderation models found
	// [here](https://docs.together.ai/docs/inference-models#moderation-models).
	SafetyModel param.Field[CompletionNewParamsSafetyModel] `json:"safety_model"`
	// Seed value for reproducibility.
	Seed param.Field[int64] `json:"seed"`
	// A list of string sequences that will truncate (stop) inference text output. For
	// example, "</s>" will stop generation as soon as the model generates the given
	// token.
	Stop param.Field[[]string] `json:"stop"`
	// A decimal number from 0-1 that determines the degree of randomness in the
	// response. A temperature less than 1 favors more correctness and is appropriate
	// for question answering or summarization. A value closer to 1 introduces more
	// randomness in the output.
	Temperature param.Field[float64] `json:"temperature"`
	// An integer that's used to limit the number of choices for the next predicted
	// word or token. It specifies the maximum number of tokens to consider at each
	// step, based on their probability of occurrence. This technique helps to speed up
	// the generation process and can improve the quality of the generated text by
	// focusing on the most likely options.
	TopK param.Field[int64] `json:"top_k"`
	// A percentage (also called the nucleus parameter) that's used to dynamically
	// adjust the number of choices for each predicted token based on the cumulative
	// probabilities. It specifies a probability threshold below which all less likely
	// tokens are filtered out. This technique helps maintain diversity and generate
	// more fluent and natural-sounding text.
	TopP param.Field[float64] `json:"top_p"`
}

func (r CompletionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

func (r CompletionNewParamsModel) IsKnown() bool {
	switch r {
	case CompletionNewParamsModelMetaLlamaLlama2_70bHf, CompletionNewParamsModelMistralaiMistral7BV0_1, CompletionNewParamsModelMistralaiMixtral8x7BV0_1, CompletionNewParamsModelMetaLlamaLlamaGuard7b:
		return true
	}
	return false
}

// The name of the moderation model used to validate tokens. Choose from the
// available moderation models found
// [here](https://docs.together.ai/docs/inference-models#moderation-models).
type CompletionNewParamsSafetyModel string

const (
	CompletionNewParamsSafetyModelMetaLlamaLlamaGuard7b CompletionNewParamsSafetyModel = "Meta-Llama/Llama-Guard-7b"
)

func (r CompletionNewParamsSafetyModel) IsKnown() bool {
	switch r {
	case CompletionNewParamsSafetyModelMetaLlamaLlamaGuard7b:
		return true
	}
	return false
}
