// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
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
	Logprobs     LogProbs                      `json:"logprobs,nullable"`
	Text         string                        `json:"text"`
	JSON         completionChoiceJSON          `json:"-"`
}

// completionChoiceJSON contains the JSON metadata for the struct
// [CompletionChoice]
type completionChoiceJSON struct {
	FinishReason apijson.Field
	Logprobs     apijson.Field
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
	CompletionChoicesFinishReasonStop      CompletionChoicesFinishReason = "stop"
	CompletionChoicesFinishReasonEos       CompletionChoicesFinishReason = "eos"
	CompletionChoicesFinishReasonLength    CompletionChoicesFinishReason = "length"
	CompletionChoicesFinishReasonToolCalls CompletionChoicesFinishReason = "tool_calls"
)

func (r CompletionChoicesFinishReason) IsKnown() bool {
	switch r {
	case CompletionChoicesFinishReasonStop, CompletionChoicesFinishReasonEos, CompletionChoicesFinishReasonLength, CompletionChoicesFinishReasonToolCalls:
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

type LogProbs struct {
	// List of token log probabilities
	TokenLogprobs []float64 `json:"token_logprobs"`
	// List of token strings
	Tokens []string     `json:"tokens"`
	JSON   logProbsJSON `json:"-"`
}

// logProbsJSON contains the JSON metadata for the struct [LogProbs]
type logProbsJSON struct {
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

type ToolChoiceParam struct {
	Function param.Field[ToolChoiceFunctionParam] `json:"function"`
	Type     param.Field[string]                  `json:"type"`
}

func (r ToolChoiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceParam) ImplementsChatCompletionNewParamsChatCompletionRequestToolChoiceUnion() {}

type ToolChoiceFunctionParam struct {
	Name param.Field[string] `json:"name"`
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

// This interface is a union satisfied by one of the following:
// [CompletionNewParamsCompletionRequest], [CompletionNewParamsCompletionRequest].
type CompletionNewParams interface {
	ImplementsCompletionNewParams()
}

type CompletionNewParamsCompletionRequest struct {
	// The name of the model to query.
	Model param.Field[string] `json:"model,required"`
	// A string providing context for the model to complete.
	Prompt param.Field[string] `json:"prompt,required"`
	// If set, the response will contain the prompt, and will also return prompt
	// logprobs if set with logprobs.
	Echo param.Field[bool] `json:"echo"`
	// The `frequency_penalty` parameter is a number between -2.0 and 2.0 where a
	// positive value will decrease the likelihood of repeating tokens that were
	// mentioned prior.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// The `logit_bias` parameter allows us to adjust the likelihood of specific tokens
	// appearing in the generated output.
	LogitBias param.Field[map[string]float64] `json:"logit_bias"`
	// Determines the number of most likely tokens to return at each token position log
	// probabilities to return
	Logprobs param.Field[int64] `json:"logprobs"`
	// The maximum number of tokens to generate.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// The `min_p` parameter is a number between 0 and 1 and an alternative to
	// `temperature`.
	MinP param.Field[float64] `json:"min_p"`
	// Number of generations to return
	N param.Field[int64] `json:"n"`
	// The `presence_penalty` parameter is a number between -2.0 and 2.0 where a
	// positive value will increase the likelihood of a model talking about new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// A number that controls the diversity of generated text by reducing the
	// likelihood of repeated sequences. Higher values decrease repetition.
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	// The name of the safety model to use.
	SafetyModel param.Field[string] `json:"safety_model"`
	// A list of string sequences that will truncate (stop) inference text output.
	Stop param.Field[[]string] `json:"stop"`
	// If set, tokens are returned as Server-Sent Events as they are made available.
	// Stream terminates with `data: [DONE]`
	Stream param.Field[CompletionNewParamsCompletionRequestStream] `json:"stream"`
	// Determines the degree of randomness in the response.
	Temperature param.Field[float64] `json:"temperature"`
	// The `top_k` parameter is used to limit the number of choices for the next
	// predicted word or token.
	TopK param.Field[int64] `json:"top_k"`
	// The `top_p` (nucleus) parameter is used to dynamically adjust the number of
	// choices for each predicted token based on the cumulative probabilities.
	TopP param.Field[float64] `json:"top_p"`
}

func (r CompletionNewParamsCompletionRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CompletionNewParamsCompletionRequest) ImplementsCompletionNewParams() {

}

// If set, tokens are returned as Server-Sent Events as they are made available.
// Stream terminates with `data: [DONE]`
type CompletionNewParamsCompletionRequestStream bool

const (
	CompletionNewParamsCompletionRequestStreamFalse CompletionNewParamsCompletionRequestStream = false
)

func (r CompletionNewParamsCompletionRequestStream) IsKnown() bool {
	switch r {
	case CompletionNewParamsCompletionRequestStreamFalse:
		return true
	}
	return false
}

type CompletionNewParamsCompletionRequest struct {
	// The name of the model to query.
	Model param.Field[string] `json:"model,required"`
	// A string providing context for the model to complete.
	Prompt param.Field[string] `json:"prompt,required"`
	// If set, tokens are returned as Server-Sent Events as they are made available.
	// Stream terminates with `data: [DONE]`
	Stream param.Field[CompletionNewParamsCompletionRequestStream] `json:"stream,required"`
	// If set, the response will contain the prompt, and will also return prompt
	// logprobs if set with logprobs.
	Echo param.Field[bool] `json:"echo"`
	// The `frequency_penalty` parameter is a number between -2.0 and 2.0 where a
	// positive value will decrease the likelihood of repeating tokens that were
	// mentioned prior.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// The `logit_bias` parameter allows us to adjust the likelihood of specific tokens
	// appearing in the generated output.
	LogitBias param.Field[map[string]float64] `json:"logit_bias"`
	// Determines the number of most likely tokens to return at each token position log
	// probabilities to return
	Logprobs param.Field[int64] `json:"logprobs"`
	// The maximum number of tokens to generate.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// The `min_p` parameter is a number between 0 and 1 and an alternative to
	// `temperature`.
	MinP param.Field[float64] `json:"min_p"`
	// Number of generations to return
	N param.Field[int64] `json:"n"`
	// The `presence_penalty` parameter is a number between -2.0 and 2.0 where a
	// positive value will increase the likelihood of a model talking about new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// A number that controls the diversity of generated text by reducing the
	// likelihood of repeated sequences. Higher values decrease repetition.
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	// The name of the safety model to use.
	SafetyModel param.Field[string] `json:"safety_model"`
	// A list of string sequences that will truncate (stop) inference text output.
	Stop param.Field[[]string] `json:"stop"`
	// Determines the degree of randomness in the response.
	Temperature param.Field[float64] `json:"temperature"`
	// The `top_k` parameter is used to limit the number of choices for the next
	// predicted word or token.
	TopK param.Field[int64] `json:"top_k"`
	// The `top_p` (nucleus) parameter is used to dynamically adjust the number of
	// choices for each predicted token based on the cumulative probabilities.
	TopP param.Field[float64] `json:"top_p"`
}

func (r CompletionNewParamsCompletionRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (CompletionNewParamsCompletionRequest) ImplementsCompletionNewParams() {

}
