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

// Query a chat model.
func (r *ChatCompletionService) NewStreaming(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[ChatCompletionChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ChatCompletionChunk](ssestream.NewDecoder(raw), err)
}

type ChatCompletion struct {
	ID       string                 `json:"id,required"`
	Choices  []ChatCompletionChoice `json:"choices,required"`
	Created  int64                  `json:"created,required"`
	Model    string                 `json:"model,required"`
	Object   ChatCompletionObject   `json:"object,required"`
	Usage    ChatCompletionUsage    `json:"usage,nullable"`
	Warnings ChatCompletionWarnings `json:"warnings"`
	JSON     chatCompletionJSON     `json:"-"`
}

// chatCompletionJSON contains the JSON metadata for the struct [ChatCompletion]
type chatCompletionJSON struct {
	ID          apijson.Field
	Choices     apijson.Field
	Created     apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	Usage       apijson.Field
	Warnings    apijson.Field
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
	Content string                           `json:"content,required,nullable"`
	Role    ChatCompletionChoicesMessageRole `json:"role,required"`
	// Deprecated: deprecated
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

// Deprecated: deprecated
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
	Warnings          ChatCompletionWarnings      `json:"warnings"`
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
	Warnings          apijson.Field
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
	FinishReason ChatCompletionChunkChoicesFinishReason `json:"finish_reason,required,nullable"`
	Index        int64                                  `json:"index,required"`
	Logprobs     float64                                `json:"logprobs,nullable"`
	Seed         int64                                  `json:"seed,nullable"`
	JSON         chatCompletionChunkChoiceJSON          `json:"-"`
}

// chatCompletionChunkChoiceJSON contains the JSON metadata for the struct
// [ChatCompletionChunkChoice]
type chatCompletionChunkChoiceJSON struct {
	Delta        apijson.Field
	FinishReason apijson.Field
	Index        apijson.Field
	Logprobs     apijson.Field
	Seed         apijson.Field
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
	Role    ChatCompletionChunkChoicesDeltaRole `json:"role,required"`
	Content string                              `json:"content,nullable"`
	// Deprecated: deprecated
	FunctionCall ChatCompletionChunkChoicesDeltaFunctionCall `json:"function_call,nullable"`
	TokenID      int64                                       `json:"token_id"`
	ToolCalls    []ToolChoice                                `json:"tool_calls"`
	JSON         chatCompletionChunkChoicesDeltaJSON         `json:"-"`
}

// chatCompletionChunkChoicesDeltaJSON contains the JSON metadata for the struct
// [ChatCompletionChunkChoicesDelta]
type chatCompletionChunkChoicesDeltaJSON struct {
	Role         apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
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

// Deprecated: deprecated
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

type ChatCompletionStructuredMessageImageURLParam struct {
	ImageURL param.Field[ChatCompletionStructuredMessageImageURLImageURLParam] `json:"image_url"`
	Type     param.Field[ChatCompletionStructuredMessageImageURLType]          `json:"type"`
}

func (r ChatCompletionStructuredMessageImageURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionStructuredMessageImageURLParam) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion() {
}

type ChatCompletionStructuredMessageImageURLImageURLParam struct {
	// The URL of the image
	URL param.Field[string] `json:"url,required"`
}

func (r ChatCompletionStructuredMessageImageURLImageURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionStructuredMessageImageURLType string

const (
	ChatCompletionStructuredMessageImageURLTypeImageURL ChatCompletionStructuredMessageImageURLType = "image_url"
)

func (r ChatCompletionStructuredMessageImageURLType) IsKnown() bool {
	switch r {
	case ChatCompletionStructuredMessageImageURLTypeImageURL:
		return true
	}
	return false
}

type ChatCompletionStructuredMessageTextParam struct {
	Text param.Field[string]                                  `json:"text,required"`
	Type param.Field[ChatCompletionStructuredMessageTextType] `json:"type,required"`
}

func (r ChatCompletionStructuredMessageTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionStructuredMessageTextParam) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion() {
}

type ChatCompletionStructuredMessageTextType string

const (
	ChatCompletionStructuredMessageTextTypeText ChatCompletionStructuredMessageTextType = "text"
)

func (r ChatCompletionStructuredMessageTextType) IsKnown() bool {
	switch r {
	case ChatCompletionStructuredMessageTextTypeText:
		return true
	}
	return false
}

type ChatCompletionStructuredMessageVideoURLParam struct {
	Type     param.Field[ChatCompletionStructuredMessageVideoURLType]          `json:"type,required"`
	VideoURL param.Field[ChatCompletionStructuredMessageVideoURLVideoURLParam] `json:"video_url,required"`
}

func (r ChatCompletionStructuredMessageVideoURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionStructuredMessageVideoURLParam) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion() {
}

type ChatCompletionStructuredMessageVideoURLType string

const (
	ChatCompletionStructuredMessageVideoURLTypeVideoURL ChatCompletionStructuredMessageVideoURLType = "video_url"
)

func (r ChatCompletionStructuredMessageVideoURLType) IsKnown() bool {
	switch r {
	case ChatCompletionStructuredMessageVideoURLTypeVideoURL:
		return true
	}
	return false
}

type ChatCompletionStructuredMessageVideoURLVideoURLParam struct {
	// The URL of the video
	URL param.Field[string] `json:"url,required"`
}

func (r ChatCompletionStructuredMessageVideoURLVideoURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type ChatCompletionWarnings []ChatCompletionWarningsItem

type ChatCompletionWarningsItem struct {
	Message string                         `json:"message,required"`
	JSON    chatCompletionWarningsItemJSON `json:"-"`
}

// chatCompletionWarningsItemJSON contains the JSON metadata for the struct
// [ChatCompletionWarningsItem]
type chatCompletionWarningsItemJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatCompletionWarningsItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionWarningsItemJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionNewParams struct {
	// A list of messages comprising the conversation so far.
	Messages param.Field[[]ChatCompletionNewParamsMessageUnion] `json:"messages,required"`
	// The name of the model to query.
	//
	// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#chat-models)
	Model param.Field[ChatCompletionNewParamsModel] `json:"model,required"`
	// Defined the behavior of the API when max_tokens exceed the maximum context
	// length of the model. When set to 'error', API will return 400 with appropriate
	// error message. When set to 'truncate', override the max_tokens with maximum
	// context length of the model.
	ContextLengthExceededBehavior param.Field[ChatCompletionNewParamsContextLengthExceededBehavior] `json:"context_length_exceeded_behavior"`
	// If true, the response will contain the prompt. Can be used with `logprobs` to
	// return prompt logprobs.
	Echo param.Field[bool] `json:"echo"`
	// A number between -2.0 and 2.0 where a positive value decreases the likelihood of
	// repeating tokens that have already been mentioned.
	FrequencyPenalty param.Field[float64]                                  `json:"frequency_penalty"`
	FunctionCall     param.Field[ChatCompletionNewParamsFunctionCallUnion] `json:"function_call"`
	// Adjusts the likelihood of specific tokens appearing in the generated output.
	LogitBias param.Field[map[string]float64] `json:"logit_bias"`
	// An integer between 0 and 20 of the top k tokens to return log probabilities for
	// at each generation step, instead of just the sampled token. Log probabilities
	// help assess model confidence in token predictions.
	Logprobs param.Field[int64] `json:"logprobs"`
	// The maximum number of tokens to generate.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// A number between 0 and 1 that can be used as an alternative to top_p and top-k.
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
	ResponseFormat param.Field[ChatCompletionNewParamsResponseFormat] `json:"response_format"`
	// The name of the moderation model used to validate tokens. Choose from the
	// available moderation models found
	// [here](https://docs.together.ai/docs/inference-models#moderation-models).
	SafetyModel param.Field[string] `json:"safety_model"`
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
	// Controls which (if any) function is called by the model. By default uses `auto`,
	// which lets the model pick between generating a message or calling a function.
	ToolChoice param.Field[ChatCompletionNewParamsToolChoiceUnion] `json:"tool_choice"`
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

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessage struct {
	Role         param.Field[ChatCompletionNewParamsMessagesRole] `json:"role,required"`
	Content      param.Field[interface{}]                         `json:"content"`
	FunctionCall param.Field[interface{}]                         `json:"function_call"`
	Name         param.Field[string]                              `json:"name"`
	ToolCallID   param.Field[string]                              `json:"tool_call_id"`
	ToolCalls    param.Field[interface{}]                         `json:"tool_calls"`
}

func (r ChatCompletionNewParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessage) implementsChatCompletionNewParamsMessageUnion() {}

// Satisfied by [ChatCompletionNewParamsMessagesChatCompletionSystemMessageParam],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageParam],
// [ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParam],
// [ChatCompletionNewParamsMessagesChatCompletionToolMessageParam],
// [ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParam],
// [ChatCompletionNewParamsMessage].
type ChatCompletionNewParamsMessageUnion interface {
	implementsChatCompletionNewParamsMessageUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionSystemMessageParam struct {
	Content param.Field[string]                                                              `json:"content,required"`
	Role    param.Field[ChatCompletionNewParamsMessagesChatCompletionSystemMessageParamRole] `json:"role,required"`
	Name    param.Field[string]                                                              `json:"name"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionSystemMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionSystemMessageParam) implementsChatCompletionNewParamsMessageUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionSystemMessageParamRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionSystemMessageParamRoleSystem ChatCompletionNewParamsMessagesChatCompletionSystemMessageParamRole = "system"
)

func (r ChatCompletionNewParamsMessagesChatCompletionSystemMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionSystemMessageParamRoleSystem:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParam struct {
	// The content of the message, which can either be a simple string or a structured
	// format.
	Content param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentUnion] `json:"content,required"`
	Role    param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageParamRole]         `json:"role,required"`
	Name    param.Field[string]                                                                    `json:"name"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParam) implementsChatCompletionNewParamsMessageUnion() {
}

// The content of the message, which can either be a simple string or a structured
// format.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodal].
type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodal []ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodal) ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItem struct {
	AudioURL   param.Field[interface{}]                                                                                                        `json:"audio_url"`
	ImageURL   param.Field[interface{}]                                                                                                        `json:"image_url"`
	InputAudio param.Field[interface{}]                                                                                                        `json:"input_audio"`
	Text       param.Field[string]                                                                                                             `json:"text"`
	Type       param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalType] `json:"type"`
	VideoURL   param.Field[interface{}]                                                                                                        `json:"video_url"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItem) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion() {
}

// Satisfied by [ChatCompletionStructuredMessageTextParam],
// [ChatCompletionStructuredMessageImageURLParam],
// [ChatCompletionStructuredMessageVideoURLParam],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudio],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudio],
// [ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItem].
type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion interface {
	implementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion()
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudio struct {
	AudioURL param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudioAudioURL] `json:"audio_url,required"`
	Type     param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudioType]     `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudio) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudioAudioURL struct {
	// The URL of the audio
	URL param.Field[string] `json:"url,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudioAudioURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudioType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudioTypeAudioURL ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudioType = "audio_url"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudioType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalAudioTypeAudioURL:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudio struct {
	InputAudio param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioInputAudio] `json:"input_audio,required"`
	Type       param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioType]       `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudio) implementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioInputAudio struct {
	// The base64 encoded audio data
	Data param.Field[string] `json:"data,required"`
	// The format of the audio data
	Format param.Field[ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioInputAudioFormat] `json:"format,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioInputAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the audio data
type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioInputAudioFormat string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioInputAudioFormatWav ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioInputAudioFormat = "wav"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioInputAudioFormat) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioInputAudioFormatWav:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioTypeInputAudio ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioType = "input_audio"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalInputAudioTypeInputAudio:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalType string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeText       ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalType = "text"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeImageURL   ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalType = "image_url"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeVideoURL   ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalType = "video_url"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeAudioURL   ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalType = "audio_url"
	ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeInputAudio ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalType = "input_audio"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeText, ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeImageURL, ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeVideoURL, ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeAudioURL, ChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalTypeInputAudio:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionUserMessageParamRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionUserMessageParamRoleUser ChatCompletionNewParamsMessagesChatCompletionUserMessageParamRole = "user"
)

func (r ChatCompletionNewParamsMessagesChatCompletionUserMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionUserMessageParamRoleUser:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParam struct {
	Role    param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParamRole] `json:"role,required"`
	Content param.Field[string]                                                                 `json:"content"`
	// Deprecated: deprecated
	FunctionCall param.Field[ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParamFunctionCall] `json:"function_call"`
	Name         param.Field[string]                                                                         `json:"name"`
	ToolCalls    param.Field[[]ToolChoiceParam]                                                              `json:"tool_calls"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParam) implementsChatCompletionNewParamsMessageUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParamRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParamRoleAssistant ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParamRole = "assistant"
)

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParamRoleAssistant:
		return true
	}
	return false
}

// Deprecated: deprecated
type ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParamFunctionCall struct {
	Arguments param.Field[string] `json:"arguments,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionAssistantMessageParamFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsMessagesChatCompletionToolMessageParam struct {
	Content    param.Field[string]                                                            `json:"content,required"`
	Role       param.Field[ChatCompletionNewParamsMessagesChatCompletionToolMessageParamRole] `json:"role,required"`
	ToolCallID param.Field[string]                                                            `json:"tool_call_id,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessageParam) implementsChatCompletionNewParamsMessageUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionToolMessageParamRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionToolMessageParamRoleTool ChatCompletionNewParamsMessagesChatCompletionToolMessageParamRole = "tool"
)

func (r ChatCompletionNewParamsMessagesChatCompletionToolMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionToolMessageParamRoleTool:
		return true
	}
	return false
}

// Deprecated: deprecated
type ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParam struct {
	Content param.Field[string]                                                                `json:"content,required"`
	Name    param.Field[string]                                                                `json:"name,required"`
	Role    param.Field[ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParamRole] `json:"role,required"`
}

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParam) implementsChatCompletionNewParamsMessageUnion() {
}

type ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParamRole string

const (
	ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParamRoleFunction ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParamRole = "function"
)

func (r ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesChatCompletionFunctionMessageParamRoleFunction:
		return true
	}
	return false
}

type ChatCompletionNewParamsMessagesRole string

const (
	ChatCompletionNewParamsMessagesRoleSystem    ChatCompletionNewParamsMessagesRole = "system"
	ChatCompletionNewParamsMessagesRoleUser      ChatCompletionNewParamsMessagesRole = "user"
	ChatCompletionNewParamsMessagesRoleAssistant ChatCompletionNewParamsMessagesRole = "assistant"
	ChatCompletionNewParamsMessagesRoleTool      ChatCompletionNewParamsMessagesRole = "tool"
	ChatCompletionNewParamsMessagesRoleFunction  ChatCompletionNewParamsMessagesRole = "function"
)

func (r ChatCompletionNewParamsMessagesRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesRoleSystem, ChatCompletionNewParamsMessagesRoleUser, ChatCompletionNewParamsMessagesRoleAssistant, ChatCompletionNewParamsMessagesRoleTool, ChatCompletionNewParamsMessagesRoleFunction:
		return true
	}
	return false
}

// The name of the model to query.
//
// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#chat-models)
type ChatCompletionNewParamsModel string

const (
	ChatCompletionNewParamsModelQwenQwen2_5_72BInstructTurbo            ChatCompletionNewParamsModel = "Qwen/Qwen2.5-72B-Instruct-Turbo"
	ChatCompletionNewParamsModelQwenQwen2_5_7BInstructTurbo             ChatCompletionNewParamsModel = "Qwen/Qwen2.5-7B-Instruct-Turbo"
	ChatCompletionNewParamsModelMetaLlamaMetaLlama3_1_405BInstructTurbo ChatCompletionNewParamsModel = "meta-llama/Meta-Llama-3.1-405B-Instruct-Turbo"
	ChatCompletionNewParamsModelMetaLlamaMetaLlama3_1_70BInstructTurbo  ChatCompletionNewParamsModel = "meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo"
	ChatCompletionNewParamsModelMetaLlamaMetaLlama3_1_8BInstructTurbo   ChatCompletionNewParamsModel = "meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo"
)

func (r ChatCompletionNewParamsModel) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsModelQwenQwen2_5_72BInstructTurbo, ChatCompletionNewParamsModelQwenQwen2_5_7BInstructTurbo, ChatCompletionNewParamsModelMetaLlamaMetaLlama3_1_405BInstructTurbo, ChatCompletionNewParamsModelMetaLlamaMetaLlama3_1_70BInstructTurbo, ChatCompletionNewParamsModelMetaLlamaMetaLlama3_1_8BInstructTurbo:
		return true
	}
	return false
}

// Defined the behavior of the API when max_tokens exceed the maximum context
// length of the model. When set to 'error', API will return 400 with appropriate
// error message. When set to 'truncate', override the max_tokens with maximum
// context length of the model.
type ChatCompletionNewParamsContextLengthExceededBehavior string

const (
	ChatCompletionNewParamsContextLengthExceededBehaviorTruncate ChatCompletionNewParamsContextLengthExceededBehavior = "truncate"
	ChatCompletionNewParamsContextLengthExceededBehaviorError    ChatCompletionNewParamsContextLengthExceededBehavior = "error"
)

func (r ChatCompletionNewParamsContextLengthExceededBehavior) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsContextLengthExceededBehaviorTruncate, ChatCompletionNewParamsContextLengthExceededBehaviorError:
		return true
	}
	return false
}

// Satisfied by [ChatCompletionNewParamsFunctionCallString],
// [ChatCompletionNewParamsFunctionCallName].
type ChatCompletionNewParamsFunctionCallUnion interface {
	implementsChatCompletionNewParamsFunctionCallUnion()
}

type ChatCompletionNewParamsFunctionCallString string

const (
	ChatCompletionNewParamsFunctionCallStringNone ChatCompletionNewParamsFunctionCallString = "none"
	ChatCompletionNewParamsFunctionCallStringAuto ChatCompletionNewParamsFunctionCallString = "auto"
)

func (r ChatCompletionNewParamsFunctionCallString) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsFunctionCallStringNone, ChatCompletionNewParamsFunctionCallStringAuto:
		return true
	}
	return false
}

func (r ChatCompletionNewParamsFunctionCallString) implementsChatCompletionNewParamsFunctionCallUnion() {
}

type ChatCompletionNewParamsFunctionCallName struct {
	Name param.Field[string] `json:"name,required"`
}

func (r ChatCompletionNewParamsFunctionCallName) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsFunctionCallName) implementsChatCompletionNewParamsFunctionCallUnion() {
}

// An object specifying the format that the model must output.
type ChatCompletionNewParamsResponseFormat struct {
	// The schema of the response format.
	Schema param.Field[map[string]interface{}] `json:"schema"`
	// The type of the response format.
	Type param.Field[string] `json:"type"`
}

func (r ChatCompletionNewParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which (if any) function is called by the model. By default uses `auto`,
// which lets the model pick between generating a message or calling a function.
//
// Satisfied by [shared.UnionString], [ToolChoiceParam].
type ChatCompletionNewParamsToolChoiceUnion interface {
	ImplementsChatCompletionNewParamsToolChoiceUnion()
}
