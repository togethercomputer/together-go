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

// ChatCompletionService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatCompletionService] method instead.
type ChatCompletionService struct {
	Options []option.RequestOption
}

// NewChatCompletionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatCompletionService(opts ...option.RequestOption) (r *ChatCompletionService) {
	r = &ChatCompletionService{}
	r.Options = opts
	return
}

// Query a chat model.
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *ChatCompletion, err error) {
	opts = append(r.Options[:], opts...)
	path := "chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChatCompletion struct {
	ID      string                 `json:"id,required"`
	Choices []ChatCompletionChoice `json:"choices,required"`
	Created int64                  `json:"created,required"`
	Model   string                 `json:"model,required"`
	Object  ChatCompletionObject   `json:"object,required"`
	Usage   ChatCompletionUsage    `json:"usage,nullable"`
	JSON    chatCompletionJSON     `json:"-"`
}

// chatCompletionJSON contains the JSON metadata for the struct [ChatCompletion]
type chatCompletionJSON struct {
	ID          apijson.Field
	Choices     apijson.Field
	Created     apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatCompletion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionChoice struct {
	FinishReason ChatCompletionChoicesFinishReason `json:"finish_reason"`
	Index        int64                             `json:"index"`
	Logprobs     LogProbs                          `json:"logprobs,nullable"`
	Message      ChatCompletionChoicesMessage      `json:"message"`
	Seed         int64                             `json:"seed"`
	Text         string                            `json:"text"`
	JSON         chatCompletionChoiceJSON          `json:"-"`
}

// chatCompletionChoiceJSON contains the JSON metadata for the struct
// [ChatCompletionChoice]
type chatCompletionChoiceJSON struct {
	FinishReason apijson.Field
	Index        apijson.Field
	Logprobs     apijson.Field
	Message      apijson.Field
	Seed         apijson.Field
	Text         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChatCompletionChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionChoiceJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionChoicesFinishReason string

const (
	ChatCompletionChoicesFinishReasonStop         ChatCompletionChoicesFinishReason = "stop"
	ChatCompletionChoicesFinishReasonEos          ChatCompletionChoicesFinishReason = "eos"
	ChatCompletionChoicesFinishReasonLength       ChatCompletionChoicesFinishReason = "length"
	ChatCompletionChoicesFinishReasonToolCalls    ChatCompletionChoicesFinishReason = "tool_calls"
	ChatCompletionChoicesFinishReasonFunctionCall ChatCompletionChoicesFinishReason = "function_call"
)

func (r ChatCompletionChoicesFinishReason) IsKnown() bool {
	switch r {
	case ChatCompletionChoicesFinishReasonStop, ChatCompletionChoicesFinishReasonEos, ChatCompletionChoicesFinishReasonLength, ChatCompletionChoicesFinishReasonToolCalls, ChatCompletionChoicesFinishReasonFunctionCall:
		return true
	}
	return false
}

type ChatCompletionChoicesMessage struct {
	Content      string                                   `json:"content,required,nullable"`
	Role         ChatCompletionChoicesMessageRole         `json:"role,required"`
	FunctionCall ChatCompletionChoicesMessageFunctionCall `json:"function_call"`
	ToolCalls    []ToolChoice                             `json:"tool_calls"`
	JSON         chatCompletionChoicesMessageJSON         `json:"-"`
}

// chatCompletionChoicesMessageJSON contains the JSON metadata for the struct
// [ChatCompletionChoicesMessage]
type chatCompletionChoicesMessageJSON struct {
	Content      apijson.Field
	Role         apijson.Field
	FunctionCall apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChatCompletionChoicesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionChoicesMessageJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionChoicesMessageRole string

const (
	ChatCompletionChoicesMessageRoleAssistant ChatCompletionChoicesMessageRole = "assistant"
)

func (r ChatCompletionChoicesMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionChoicesMessageRoleAssistant:
		return true
	}
	return false
}

type ChatCompletionChoicesMessageFunctionCall struct {
	Arguments string                                       `json:"arguments,required"`
	Name      string                                       `json:"name,required"`
	JSON      chatCompletionChoicesMessageFunctionCallJSON `json:"-"`
}

// chatCompletionChoicesMessageFunctionCallJSON contains the JSON metadata for the
// struct [ChatCompletionChoicesMessageFunctionCall]
type chatCompletionChoicesMessageFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatCompletionChoicesMessageFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionChoicesMessageFunctionCallJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionObject string

const (
	ChatCompletionObjectChatCompletion ChatCompletionObject = "chat.completion"
)

func (r ChatCompletionObject) IsKnown() bool {
	switch r {
	case ChatCompletionObjectChatCompletion:
		return true
	}
	return false
}

type ChatCompletionChunk struct {
	ID                string                      `json:"id,required"`
	Choices           []ChatCompletionChunkChoice `json:"choices,required"`
	Created           int64                       `json:"created,required"`
	Model             string                      `json:"model,required"`
	Object            ChatCompletionChunkObject   `json:"object,required"`
	SystemFingerprint string                      `json:"system_fingerprint"`
	Usage             ChatCompletionUsage         `json:"usage,nullable"`
	JSON              chatCompletionChunkJSON     `json:"-"`
}

// chatCompletionChunkJSON contains the JSON metadata for the struct
// [ChatCompletionChunk]
type chatCompletionChunkJSON struct {
	ID                apijson.Field
	Choices           apijson.Field
	Created           apijson.Field
	Model             apijson.Field
	Object            apijson.Field
	SystemFingerprint apijson.Field
	Usage             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ChatCompletionChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionChunkJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionChunkChoice struct {
	Delta        ChatCompletionChunkChoicesDelta        `json:"delta,required"`
	FinishReason ChatCompletionChunkChoicesFinishReason `json:"finish_reason,required"`
	Index        int64                                  `json:"index,required"`
	Logprobs     LogProbs                               `json:"logprobs"`
	JSON         chatCompletionChunkChoiceJSON          `json:"-"`
}

// chatCompletionChunkChoiceJSON contains the JSON metadata for the struct
// [ChatCompletionChunkChoice]
type chatCompletionChunkChoiceJSON struct {
	Delta        apijson.Field
	FinishReason apijson.Field
	Index        apijson.Field
	Logprobs     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChatCompletionChunkChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionChunkChoiceJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionChunkChoicesDelta struct {
	Content      string                                      `json:"content,nullable"`
	FunctionCall ChatCompletionChunkChoicesDeltaFunctionCall `json:"function_call,nullable"`
	Role         ChatCompletionChunkChoicesDeltaRole         `json:"role"`
	TokenID      int64                                       `json:"token_id"`
	ToolCalls    []ToolChoice                                `json:"tool_calls"`
	JSON         chatCompletionChunkChoicesDeltaJSON         `json:"-"`
}

// chatCompletionChunkChoicesDeltaJSON contains the JSON metadata for the struct
// [ChatCompletionChunkChoicesDelta]
type chatCompletionChunkChoicesDeltaJSON struct {
	Content      apijson.Field
	FunctionCall apijson.Field
	Role         apijson.Field
	TokenID      apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChatCompletionChunkChoicesDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionChunkChoicesDeltaJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionChunkChoicesDeltaFunctionCall struct {
	Arguments string                                          `json:"arguments,required"`
	Name      string                                          `json:"name,required"`
	JSON      chatCompletionChunkChoicesDeltaFunctionCallJSON `json:"-"`
}

// chatCompletionChunkChoicesDeltaFunctionCallJSON contains the JSON metadata for
// the struct [ChatCompletionChunkChoicesDeltaFunctionCall]
type chatCompletionChunkChoicesDeltaFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatCompletionChunkChoicesDeltaFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionChunkChoicesDeltaFunctionCallJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionChunkChoicesDeltaRole string

const (
	ChatCompletionChunkChoicesDeltaRoleSystem    ChatCompletionChunkChoicesDeltaRole = "system"
	ChatCompletionChunkChoicesDeltaRoleUser      ChatCompletionChunkChoicesDeltaRole = "user"
	ChatCompletionChunkChoicesDeltaRoleAssistant ChatCompletionChunkChoicesDeltaRole = "assistant"
	ChatCompletionChunkChoicesDeltaRoleFunction  ChatCompletionChunkChoicesDeltaRole = "function"
	ChatCompletionChunkChoicesDeltaRoleTool      ChatCompletionChunkChoicesDeltaRole = "tool"
)

func (r ChatCompletionChunkChoicesDeltaRole) IsKnown() bool {
	switch r {
	case ChatCompletionChunkChoicesDeltaRoleSystem, ChatCompletionChunkChoicesDeltaRoleUser, ChatCompletionChunkChoicesDeltaRoleAssistant, ChatCompletionChunkChoicesDeltaRoleFunction, ChatCompletionChunkChoicesDeltaRoleTool:
		return true
	}
	return false
}

type ChatCompletionChunkChoicesFinishReason string

const (
	ChatCompletionChunkChoicesFinishReasonStop         ChatCompletionChunkChoicesFinishReason = "stop"
	ChatCompletionChunkChoicesFinishReasonEos          ChatCompletionChunkChoicesFinishReason = "eos"
	ChatCompletionChunkChoicesFinishReasonLength       ChatCompletionChunkChoicesFinishReason = "length"
	ChatCompletionChunkChoicesFinishReasonToolCalls    ChatCompletionChunkChoicesFinishReason = "tool_calls"
	ChatCompletionChunkChoicesFinishReasonFunctionCall ChatCompletionChunkChoicesFinishReason = "function_call"
)

func (r ChatCompletionChunkChoicesFinishReason) IsKnown() bool {
	switch r {
	case ChatCompletionChunkChoicesFinishReasonStop, ChatCompletionChunkChoicesFinishReasonEos, ChatCompletionChunkChoicesFinishReasonLength, ChatCompletionChunkChoicesFinishReasonToolCalls, ChatCompletionChunkChoicesFinishReasonFunctionCall:
		return true
	}
	return false
}

type ChatCompletionChunkObject string

const (
	ChatCompletionChunkObjectChatCompletionChunk ChatCompletionChunkObject = "chat.completion.chunk"
)

func (r ChatCompletionChunkObject) IsKnown() bool {
	switch r {
	case ChatCompletionChunkObjectChatCompletionChunk:
		return true
	}
	return false
}

type ChatCompletionUsage struct {
	CompletionTokens int64                   `json:"completion_tokens,required"`
	PromptTokens     int64                   `json:"prompt_tokens,required"`
	TotalTokens      int64                   `json:"total_tokens,required"`
	JSON             chatCompletionUsageJSON `json:"-"`
}

// chatCompletionUsageJSON contains the JSON metadata for the struct
// [ChatCompletionUsage]
type chatCompletionUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ChatCompletionUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionUsageJSON) RawJSON() string {
	return r.raw
}

// This interface is a union satisfied by one of the following:
// [ChatCompletionNewParamsChatCompletionRequest],
// [ChatCompletionNewParamsChatCompletionRequest].
type ChatCompletionNewParams interface {
	ImplementsChatCompletionNewParams()
}

type ChatCompletionNewParamsChatCompletionRequest struct {
	// A list of messages comprising the conversation so far.
	Messages param.Field[[]ChatCompletionNewParamsChatCompletionRequestMessage] `json:"messages,required"`
	// The name of the model to query.
	Model param.Field[string] `json:"model,required"`
	// If true, the response will contain the prompt. Can be used with `logprobs` to
	// return prompt logprobs.
	Echo param.Field[bool] `json:"echo"`
	// A number between -2.0 and 2.0 where a positive value decreases the likelihood of
	// repeating tokens that have already been mentioned.
	FrequencyPenalty param.Field[float64]                                                       `json:"frequency_penalty"`
	FunctionCall     param.Field[ChatCompletionNewParamsChatCompletionRequestFunctionCallUnion] `json:"function_call"`
	// Adjusts the likelihood of specific tokens appearing in the generated output.
	LogitBias param.Field[map[string]float64] `json:"logit_bias"`
	// Determines the number of most likely tokens to return at each token position log
	// probabilities to return.
	Logprobs param.Field[int64] `json:"logprobs"`
	// The maximum number of tokens to generate.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// A number between 0 and 1 that can be used as an alternative to temperature.
	MinP param.Field[float64] `json:"min_p"`
	// The number of completions to generate for each prompt.
	N param.Field[int64] `json:"n"`
	// A number between -2.0 and 2.0 where a positive value increases the likelihood of
	// a model talking about new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// A number that controls the diversity of generated text by reducing the
	// likelihood of repeated sequences. Higher values decrease repetition.
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	// An object specifying the format that the model must output.
	ResponseFormat param.Field[ChatCompletionNewParamsChatCompletionRequestResponseFormat] `json:"response_format"`
	// The name of the moderation model used to validate tokens. Choose from the
	// available moderation models found
	// [here](https://docs.together.ai/docs/inference-models#moderation-models).
	SafetyModel param.Field[string] `json:"safety_model"`
	// A list of string sequences that will truncate (stop) inference text output. For
	// example, "</s>" will stop generation as soon as the model generates the given
	// token.
	Stop param.Field[[]string] `json:"stop"`
	// If true, stream tokens as Server-Sent Events as the model generates them instead
	// of waiting for the full model response. The stream terminates with
	// `data: [DONE]`. If false, return a single JSON object containing the results.
	Stream param.Field[ChatCompletionNewParamsChatCompletionRequestStream] `json:"stream"`
	// A decimal number from 0-1 that determines the degree of randomness in the
	// response. A temperature less than 1 favors more correctness and is appropriate
	// for question answering or summarization. A value closer to 1 introduces more
	// randomness in the output.
	Temperature param.Field[float64] `json:"temperature"`
	// Controls which (if any) function is called by the model. By default uses `auto`,
	// which lets the model pick between generating a message or calling a function.
	ToolChoice param.Field[ChatCompletionNewParamsChatCompletionRequestToolChoiceUnion] `json:"tool_choice"`
	// A list of tools the model may call. Currently, only functions are supported as a
	// tool. Use this to provide a list of functions the model may generate JSON inputs
	// for.
	Tools param.Field[[]ToolsParam] `json:"tools"`
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

func (r ChatCompletionNewParamsChatCompletionRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ChatCompletionNewParamsChatCompletionRequest) ImplementsChatCompletionNewParams() {

}

type ChatCompletionNewParamsChatCompletionRequestMessage struct {
	// The contents of the message.
	Content param.Field[string] `json:"content,required"`
	// The role of the messages author. Choice between: system, user, or assistant.
	Role param.Field[ChatCompletionNewParamsChatCompletionRequestMessagesRole] `json:"role,required"`
}

func (r ChatCompletionNewParamsChatCompletionRequestMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The role of the messages author. Choice between: system, user, or assistant.
type ChatCompletionNewParamsChatCompletionRequestMessagesRole string

const (
	ChatCompletionNewParamsChatCompletionRequestMessagesRoleSystem    ChatCompletionNewParamsChatCompletionRequestMessagesRole = "system"
	ChatCompletionNewParamsChatCompletionRequestMessagesRoleUser      ChatCompletionNewParamsChatCompletionRequestMessagesRole = "user"
	ChatCompletionNewParamsChatCompletionRequestMessagesRoleAssistant ChatCompletionNewParamsChatCompletionRequestMessagesRole = "assistant"
)

func (r ChatCompletionNewParamsChatCompletionRequestMessagesRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsChatCompletionRequestMessagesRoleSystem, ChatCompletionNewParamsChatCompletionRequestMessagesRoleUser, ChatCompletionNewParamsChatCompletionRequestMessagesRoleAssistant:
		return true
	}
	return false
}

type ChatCompletionNewParamsChatCompletionRequestFunctionCall struct {
	Name param.Field[string] `json:"name,required"`
}

func (r ChatCompletionNewParamsChatCompletionRequestFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsChatCompletionRequestFunctionCall) implementsChatCompletionNewParamsChatCompletionRequestFunctionCallUnion() {
}

// Satisfied by [ChatCompletionNewParamsChatCompletionRequestFunctionCallString],
// [ChatCompletionNewParamsChatCompletionRequestFunctionCallObject],
// [ChatCompletionNewParamsChatCompletionRequestFunctionCall].
type ChatCompletionNewParamsChatCompletionRequestFunctionCallUnion interface {
	implementsChatCompletionNewParamsChatCompletionRequestFunctionCallUnion()
}

type ChatCompletionNewParamsChatCompletionRequestFunctionCallString string

const (
	ChatCompletionNewParamsChatCompletionRequestFunctionCallStringNone ChatCompletionNewParamsChatCompletionRequestFunctionCallString = "none"
	ChatCompletionNewParamsChatCompletionRequestFunctionCallStringAuto ChatCompletionNewParamsChatCompletionRequestFunctionCallString = "auto"
)

func (r ChatCompletionNewParamsChatCompletionRequestFunctionCallString) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsChatCompletionRequestFunctionCallStringNone, ChatCompletionNewParamsChatCompletionRequestFunctionCallStringAuto:
		return true
	}
	return false
}

func (r ChatCompletionNewParamsChatCompletionRequestFunctionCallString) implementsChatCompletionNewParamsChatCompletionRequestFunctionCallUnion() {
}

type ChatCompletionNewParamsChatCompletionRequestFunctionCallObject struct {
	Name param.Field[string] `json:"name,required"`
}

func (r ChatCompletionNewParamsChatCompletionRequestFunctionCallObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsChatCompletionRequestFunctionCallObject) implementsChatCompletionNewParamsChatCompletionRequestFunctionCallUnion() {
}

// An object specifying the format that the model must output.
type ChatCompletionNewParamsChatCompletionRequestResponseFormat struct {
	// The schema of the response format.
	Schema param.Field[map[string]string] `json:"schema"`
	// The type of the response format.
	Type param.Field[string] `json:"type"`
}

func (r ChatCompletionNewParamsChatCompletionRequestResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If true, stream tokens as Server-Sent Events as the model generates them instead
// of waiting for the full model response. The stream terminates with
// `data: [DONE]`. If false, return a single JSON object containing the results.
type ChatCompletionNewParamsChatCompletionRequestStream bool

const (
	ChatCompletionNewParamsChatCompletionRequestStreamFalse ChatCompletionNewParamsChatCompletionRequestStream = false
)

func (r ChatCompletionNewParamsChatCompletionRequestStream) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsChatCompletionRequestStreamFalse:
		return true
	}
	return false
}

// Controls which (if any) function is called by the model. By default uses `auto`,
// which lets the model pick between generating a message or calling a function.
type ChatCompletionNewParamsChatCompletionRequestToolChoice struct {
	Index    param.Field[float64]                                                    `json:"index,required"`
	ID       param.Field[string]                                                     `json:"id,required"`
	Type     param.Field[ChatCompletionNewParamsChatCompletionRequestToolChoiceType] `json:"type,required"`
	Function param.Field[interface{}]                                                `json:"function"`
}

func (r ChatCompletionNewParamsChatCompletionRequestToolChoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsChatCompletionRequestToolChoice) ImplementsChatCompletionNewParamsChatCompletionRequestToolChoiceUnion() {
}

// Controls which (if any) function is called by the model. By default uses `auto`,
// which lets the model pick between generating a message or calling a function.
//
// Satisfied by [shared.UnionString], [ToolChoiceParam],
// [ChatCompletionNewParamsChatCompletionRequestToolChoice].
type ChatCompletionNewParamsChatCompletionRequestToolChoiceUnion interface {
	ImplementsChatCompletionNewParamsChatCompletionRequestToolChoiceUnion()
}

type ChatCompletionNewParamsChatCompletionRequest struct {
	// A list of messages comprising the conversation so far.
	Messages param.Field[[]ChatCompletionNewParamsChatCompletionRequestMessage] `json:"messages,required"`
	// The name of the model to query.
	Model param.Field[string] `json:"model,required"`
	// If true, stream tokens as Server-Sent Events as the model generates them instead
	// of waiting for the full model response. The stream terminates with
	// `data: [DONE]`. If false, return a single JSON object containing the results.
	Stream param.Field[ChatCompletionNewParamsChatCompletionRequestStream] `json:"stream,required"`
	// If true, the response will contain the prompt. Can be used with `logprobs` to
	// return prompt logprobs.
	Echo param.Field[bool] `json:"echo"`
	// A number between -2.0 and 2.0 where a positive value decreases the likelihood of
	// repeating tokens that have already been mentioned.
	FrequencyPenalty param.Field[float64]                                                       `json:"frequency_penalty"`
	FunctionCall     param.Field[ChatCompletionNewParamsChatCompletionRequestFunctionCallUnion] `json:"function_call"`
	// Adjusts the likelihood of specific tokens appearing in the generated output.
	LogitBias param.Field[map[string]float64] `json:"logit_bias"`
	// Determines the number of most likely tokens to return at each token position log
	// probabilities to return.
	Logprobs param.Field[int64] `json:"logprobs"`
	// The maximum number of tokens to generate.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// A number between 0 and 1 that can be used as an alternative to temperature.
	MinP param.Field[float64] `json:"min_p"`
	// The number of completions to generate for each prompt.
	N param.Field[int64] `json:"n"`
	// A number between -2.0 and 2.0 where a positive value increases the likelihood of
	// a model talking about new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// A number that controls the diversity of generated text by reducing the
	// likelihood of repeated sequences. Higher values decrease repetition.
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	// An object specifying the format that the model must output.
	ResponseFormat param.Field[ChatCompletionNewParamsChatCompletionRequestResponseFormat] `json:"response_format"`
	// The name of the moderation model used to validate tokens. Choose from the
	// available moderation models found
	// [here](https://docs.together.ai/docs/inference-models#moderation-models).
	SafetyModel param.Field[string] `json:"safety_model"`
	// A list of string sequences that will truncate (stop) inference text output. For
	// example, "</s>" will stop generation as soon as the model generates the given
	// token.
	Stop param.Field[[]string] `json:"stop"`
	// A decimal number from 0-1 that determines the degree of randomness in the
	// response. A temperature less than 1 favors more correctness and is appropriate
	// for question answering or summarization. A value closer to 1 introduces more
	// randomness in the output.
	Temperature param.Field[float64] `json:"temperature"`
	// Controls which (if any) function is called by the model. By default uses `auto`,
	// which lets the model pick between generating a message or calling a function.
	ToolChoice param.Field[ChatCompletionNewParamsChatCompletionRequestToolChoiceUnion] `json:"tool_choice"`
	// A list of tools the model may call. Currently, only functions are supported as a
	// tool. Use this to provide a list of functions the model may generate JSON inputs
	// for.
	Tools param.Field[[]ToolsParam] `json:"tools"`
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

func (r ChatCompletionNewParamsChatCompletionRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ChatCompletionNewParamsChatCompletionRequest) ImplementsChatCompletionNewParams() {

}
