// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
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
func NewChatCompletionService(opts ...option.RequestOption) (r ChatCompletionService) {
	r = ChatCompletionService{}
	r.Options = opts
	return
}

// Query a chat model.
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *ChatCompletion, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ChatCompletionChunk](ssestream.NewDecoder(raw), err)
}

type ChatCompletion struct {
	ID      string                 `json:"id,required"`
	Choices []ChatCompletionChoice `json:"choices,required"`
	Created int64                  `json:"created,required"`
	Model   string                 `json:"model,required"`
	// Any of "chat.completion".
	Object   ChatCompletionObject    `json:"object,required"`
	Usage    ChatCompletionUsage     `json:"usage,nullable"`
	Warnings []ChatCompletionWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Choices     respjson.Field
		Created     respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		Usage       respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletion) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionChoice struct {
	// Any of "stop", "eos", "length", "tool_calls", "function_call".
	FinishReason string                      `json:"finish_reason"`
	Index        int64                       `json:"index"`
	Logprobs     LogProbs                    `json:"logprobs,nullable"`
	Message      ChatCompletionChoiceMessage `json:"message"`
	Seed         int64                       `json:"seed"`
	Text         string                      `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Logprobs     respjson.Field
		Message      respjson.Field
		Seed         respjson.Field
		Text         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionChoiceMessage struct {
	Content string `json:"content,required"`
	// Any of "assistant".
	Role string `json:"role,required"`
	// Deprecated: deprecated
	FunctionCall ChatCompletionChoiceMessageFunctionCall `json:"function_call"`
	Reasoning    string                                  `json:"reasoning,nullable"`
	ToolCalls    []ToolChoice                            `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content      respjson.Field
		Role         respjson.Field
		FunctionCall respjson.Field
		Reasoning    respjson.Field
		ToolCalls    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChoiceMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChoiceMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated: deprecated
type ChatCompletionChoiceMessageFunctionCall struct {
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
func (r ChatCompletionChoiceMessageFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChoiceMessageFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionObject string

const (
	ChatCompletionObjectChatCompletion ChatCompletionObject = "chat.completion"
)

type ChatCompletionChunk struct {
	ID      string                      `json:"id,required"`
	Choices []ChatCompletionChunkChoice `json:"choices,required"`
	Created int64                       `json:"created,required"`
	Model   string                      `json:"model,required"`
	// Any of "chat.completion.chunk".
	Object            ChatCompletionChunkObject `json:"object,required"`
	SystemFingerprint string                    `json:"system_fingerprint"`
	Usage             ChatCompletionUsage       `json:"usage,nullable"`
	Warnings          []ChatCompletionWarning   `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Choices           respjson.Field
		Created           respjson.Field
		Model             respjson.Field
		Object            respjson.Field
		SystemFingerprint respjson.Field
		Usage             respjson.Field
		Warnings          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunk) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionChunkChoice struct {
	Delta ChatCompletionChunkChoiceDelta `json:"delta,required"`
	// Any of "stop", "eos", "length", "tool_calls", "function_call".
	FinishReason string  `json:"finish_reason,required"`
	Index        int64   `json:"index,required"`
	Logprobs     float64 `json:"logprobs,nullable"`
	Seed         int64   `json:"seed,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta        respjson.Field
		FinishReason respjson.Field
		Index        respjson.Field
		Logprobs     respjson.Field
		Seed         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunkChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionChunkChoiceDelta struct {
	// Any of "system", "user", "assistant", "function", "tool".
	Role    string `json:"role,required"`
	Content string `json:"content,nullable"`
	// Deprecated: deprecated
	FunctionCall ChatCompletionChunkChoiceDeltaFunctionCall `json:"function_call,nullable"`
	Reasoning    string                                     `json:"reasoning,nullable"`
	TokenID      int64                                      `json:"token_id"`
	ToolCalls    []ToolChoice                               `json:"tool_calls"`
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
func (r ChatCompletionChunkChoiceDelta) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated: deprecated
type ChatCompletionChunkChoiceDeltaFunctionCall struct {
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
func (r ChatCompletionChunkChoiceDeltaFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkChoiceDeltaFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionChunkObject string

const (
	ChatCompletionChunkObjectChatCompletionChunk ChatCompletionChunkObject = "chat.completion.chunk"
)

type ChatCompletionStructuredMessageImageURLParam struct {
	ImageURL ChatCompletionStructuredMessageImageURLImageURLParam `json:"image_url,omitzero"`
	// Any of "image_url".
	Type ChatCompletionStructuredMessageImageURLType `json:"type,omitzero"`
	paramObj
}

func (r ChatCompletionStructuredMessageImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionStructuredMessageImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionStructuredMessageImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type ChatCompletionStructuredMessageImageURLImageURLParam struct {
	// The URL of the image
	URL string `json:"url,required"`
	paramObj
}

func (r ChatCompletionStructuredMessageImageURLImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionStructuredMessageImageURLImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionStructuredMessageImageURLImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionStructuredMessageImageURLType string

const (
	ChatCompletionStructuredMessageImageURLTypeImageURL ChatCompletionStructuredMessageImageURLType = "image_url"
)

// The properties Text, Type are required.
type ChatCompletionStructuredMessageTextParam struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type ChatCompletionStructuredMessageTextType `json:"type,omitzero,required"`
	paramObj
}

func (r ChatCompletionStructuredMessageTextParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionStructuredMessageTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionStructuredMessageTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionStructuredMessageTextType string

const (
	ChatCompletionStructuredMessageTextTypeText ChatCompletionStructuredMessageTextType = "text"
)

// The properties Type, VideoURL are required.
type ChatCompletionStructuredMessageVideoURLParam struct {
	// Any of "video_url".
	Type     ChatCompletionStructuredMessageVideoURLType          `json:"type,omitzero,required"`
	VideoURL ChatCompletionStructuredMessageVideoURLVideoURLParam `json:"video_url,omitzero,required"`
	paramObj
}

func (r ChatCompletionStructuredMessageVideoURLParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionStructuredMessageVideoURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionStructuredMessageVideoURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionStructuredMessageVideoURLType string

const (
	ChatCompletionStructuredMessageVideoURLTypeVideoURL ChatCompletionStructuredMessageVideoURLType = "video_url"
)

// The property URL is required.
type ChatCompletionStructuredMessageVideoURLVideoURLParam struct {
	// The URL of the video
	URL string `json:"url,required"`
	paramObj
}

func (r ChatCompletionStructuredMessageVideoURLVideoURLParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionStructuredMessageVideoURLVideoURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionStructuredMessageVideoURLVideoURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionUsage struct {
	CompletionTokens int64 `json:"completion_tokens,required"`
	PromptTokens     int64 `json:"prompt_tokens,required"`
	TotalTokens      int64 `json:"total_tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens respjson.Field
		PromptTokens     respjson.Field
		TotalTokens      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionUsage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionWarning struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionWarning) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewParams struct {
	// A list of messages comprising the conversation so far.
	Messages []ChatCompletionNewParamsMessageUnion `json:"messages,omitzero,required"`
	// The name of the model to query.
	//
	// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#chat-models)
	Model ChatCompletionNewParamsModel `json:"model,omitzero,required"`
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
	// A number between 0 and 1 that can be used as an alternative to top_p and top-k.
	MinP param.Opt[float64] `json:"min_p,omitzero"`
	// The number of completions to generate for each prompt.
	N param.Opt[int64] `json:"n,omitzero"`
	// A number between -2.0 and 2.0 where a positive value increases the likelihood of
	// a model talking about new topics.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// A number that controls the diversity of generated text by reducing the
	// likelihood of repeated sequences. Higher values decrease repetition.
	RepetitionPenalty param.Opt[float64] `json:"repetition_penalty,omitzero"`
	// The name of the moderation model used to validate tokens. Choose from the
	// available moderation models found
	// [here](https://docs.together.ai/docs/inference-models#moderation-models).
	SafetyModel param.Opt[string] `json:"safety_model,omitzero"`
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
	// Defined the behavior of the API when max_tokens exceed the maximum context
	// length of the model. When set to 'error', API will return 400 with appropriate
	// error message. When set to 'truncate', override the max_tokens with maximum
	// context length of the model.
	//
	// Any of "truncate", "error".
	ContextLengthExceededBehavior ChatCompletionNewParamsContextLengthExceededBehavior `json:"context_length_exceeded_behavior,omitzero"`
	FunctionCall                  ChatCompletionNewParamsFunctionCallUnion             `json:"function_call,omitzero"`
	// Adjusts the likelihood of specific tokens appearing in the generated output.
	LogitBias map[string]float64 `json:"logit_bias,omitzero"`
	// Controls the level of reasoning effort the model should apply when generating
	// responses. Higher values may result in more thoughtful and detailed responses
	// but may take longer to generate.
	//
	// Any of "low", "medium", "high".
	ReasoningEffort ChatCompletionNewParamsReasoningEffort `json:"reasoning_effort,omitzero"`
	// An object specifying the format that the model must output.
	ResponseFormat ChatCompletionNewParamsResponseFormat `json:"response_format,omitzero"`
	// A list of string sequences that will truncate (stop) inference text output. For
	// example, "</s>" will stop generation as soon as the model generates the given
	// token.
	Stop []string `json:"stop,omitzero"`
	// Controls which (if any) function is called by the model. By default uses `auto`,
	// which lets the model pick between generating a message or calling a function.
	ToolChoice ChatCompletionNewParamsToolChoiceUnion `json:"tool_choice,omitzero"`
	// A list of tools the model may call. Currently, only functions are supported as a
	// tool. Use this to provide a list of functions the model may generate JSON inputs
	// for.
	Tools []ToolsParam `json:"tools,omitzero"`
	paramObj
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageUnion struct {
	OfChatCompletionNewsMessageChatCompletionSystemMessageParam    *ChatCompletionNewParamsMessageChatCompletionSystemMessageParam    `json:",omitzero,inline"`
	OfChatCompletionNewsMessageChatCompletionUserMessageParam      *ChatCompletionNewParamsMessageChatCompletionUserMessageParam      `json:",omitzero,inline"`
	OfChatCompletionNewsMessageChatCompletionAssistantMessageParam *ChatCompletionNewParamsMessageChatCompletionAssistantMessageParam `json:",omitzero,inline"`
	OfChatCompletionNewsMessageChatCompletionToolMessageParam      *ChatCompletionNewParamsMessageChatCompletionToolMessageParam      `json:",omitzero,inline"`
	OfChatCompletionNewsMessageChatCompletionFunctionMessageParam  *ChatCompletionNewParamsMessageChatCompletionFunctionMessageParam  `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfChatCompletionNewsMessageChatCompletionSystemMessageParam,
		u.OfChatCompletionNewsMessageChatCompletionUserMessageParam,
		u.OfChatCompletionNewsMessageChatCompletionAssistantMessageParam,
		u.OfChatCompletionNewsMessageChatCompletionToolMessageParam,
		u.OfChatCompletionNewsMessageChatCompletionFunctionMessageParam)
}
func (u *ChatCompletionNewParamsMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageUnion) asAny() any {
	if !param.IsOmitted(u.OfChatCompletionNewsMessageChatCompletionSystemMessageParam) {
		return u.OfChatCompletionNewsMessageChatCompletionSystemMessageParam
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageChatCompletionUserMessageParam) {
		return u.OfChatCompletionNewsMessageChatCompletionUserMessageParam
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageChatCompletionAssistantMessageParam) {
		return u.OfChatCompletionNewsMessageChatCompletionAssistantMessageParam
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageChatCompletionToolMessageParam) {
		return u.OfChatCompletionNewsMessageChatCompletionToolMessageParam
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageChatCompletionFunctionMessageParam) {
		return u.OfChatCompletionNewsMessageChatCompletionFunctionMessageParam
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUnion) GetFunctionCall() *ChatCompletionNewParamsMessageChatCompletionAssistantMessageParamFunctionCall {
	if vt := u.OfChatCompletionNewsMessageChatCompletionAssistantMessageParam; vt != nil {
		return &vt.FunctionCall
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUnion) GetToolCalls() []ToolChoiceParam {
	if vt := u.OfChatCompletionNewsMessageChatCompletionAssistantMessageParam; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUnion) GetToolCallID() *string {
	if vt := u.OfChatCompletionNewsMessageChatCompletionToolMessageParam; vt != nil {
		return &vt.ToolCallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUnion) GetRole() *string {
	if vt := u.OfChatCompletionNewsMessageChatCompletionSystemMessageParam; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionUserMessageParam; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionAssistantMessageParam; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionToolMessageParam; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionFunctionMessageParam; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageUnion) GetName() *string {
	if vt := u.OfChatCompletionNewsMessageChatCompletionSystemMessageParam; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionUserMessageParam; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionAssistantMessageParam; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionToolMessageParam; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionFunctionMessageParam; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ChatCompletionNewParamsMessageUnion) GetContent() (res chatCompletionNewParamsMessageUnionContent) {
	if vt := u.OfChatCompletionNewsMessageChatCompletionSystemMessageParam; vt != nil {
		res.any = &vt.Content
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionUserMessageParam; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionAssistantMessageParam; vt != nil && vt.Content.Valid() {
		res.any = &vt.Content.Value
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionToolMessageParam; vt != nil {
		res.any = &vt.Content
	} else if vt := u.OfChatCompletionNewsMessageChatCompletionFunctionMessageParam; vt != nil {
		res.any = &vt.Content
	}
	return
}

// Can have the runtime types [*string],
// [\*[]ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion]
type chatCompletionNewParamsMessageUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]together.ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u chatCompletionNewParamsMessageUnionContent) AsAny() any { return u.any }

// The properties Content, Role are required.
type ChatCompletionNewParamsMessageChatCompletionSystemMessageParam struct {
	Content string `json:"content,required"`
	// Any of "system".
	Role string            `json:"role,omitzero,required"`
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionSystemMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionSystemMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionSystemMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageChatCompletionSystemMessageParam](
		"role", "system",
	)
}

// The properties Content, Role are required.
type ChatCompletionNewParamsMessageChatCompletionUserMessageParam struct {
	// The content of the message, which can either be a simple string or a structured
	// format.
	Content ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentUnion `json:"content,omitzero,required"`
	// Any of "user".
	Role string            `json:"role,omitzero,required"`
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionUserMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionUserMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionUserMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageChatCompletionUserMessageParam](
		"role", "user",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentUnion struct {
	OfString                                                                                                        param.Opt[string]                                                                                                        `json:",omitzero,inline"`
	OfChatCompletionNewsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalArray []ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatCompletionNewsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalArray)
}
func (u *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfChatCompletionNewsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalArray) {
		return &u.OfChatCompletionNewsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion struct {
	OfChatCompletionStructuredMessageText     *ChatCompletionStructuredMessageTextParam                                                                                    `json:",omitzero,inline"`
	OfChatCompletionStructuredMessageImageURL *ChatCompletionStructuredMessageImageURLParam                                                                                `json:",omitzero,inline"`
	OfVideo                                   *ChatCompletionStructuredMessageVideoURLParam                                                                                `json:",omitzero,inline"`
	OfAudio                                   *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio      `json:",omitzero,inline"`
	OfInputAudio                              *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfChatCompletionStructuredMessageText,
		u.OfChatCompletionStructuredMessageImageURL,
		u.OfVideo,
		u.OfAudio,
		u.OfInputAudio)
}
func (u *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) asAny() any {
	if !param.IsOmitted(u.OfChatCompletionStructuredMessageText) {
		return u.OfChatCompletionStructuredMessageText
	} else if !param.IsOmitted(u.OfChatCompletionStructuredMessageImageURL) {
		return u.OfChatCompletionStructuredMessageImageURL
	} else if !param.IsOmitted(u.OfVideo) {
		return u.OfVideo
	} else if !param.IsOmitted(u.OfAudio) {
		return u.OfAudio
	} else if !param.IsOmitted(u.OfInputAudio) {
		return u.OfInputAudio
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetText() *string {
	if vt := u.OfChatCompletionStructuredMessageText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetImageURL() *ChatCompletionStructuredMessageImageURLImageURLParam {
	if vt := u.OfChatCompletionStructuredMessageImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetVideoURL() *ChatCompletionStructuredMessageVideoURLVideoURLParam {
	if vt := u.OfVideo; vt != nil {
		return &vt.VideoURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetAudioURL() *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL {
	if vt := u.OfAudio; vt != nil {
		return &vt.AudioURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetInputAudio() *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio {
	if vt := u.OfInputAudio; vt != nil {
		return &vt.InputAudio
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetType() *string {
	if vt := u.OfChatCompletionStructuredMessageText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfChatCompletionStructuredMessageImageURL; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVideo; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAudio; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInputAudio; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The properties AudioURL, Type are required.
type ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio struct {
	AudioURL ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL `json:"audio_url,omitzero,required"`
	// Any of "audio_url".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio](
		"type", "audio_url",
	)
}

// The property URL is required.
type ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL struct {
	// The URL of the audio
	URL string `json:"url,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties InputAudio, Type are required.
type ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio struct {
	InputAudio ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio `json:"input_audio,omitzero,required"`
	// Any of "input_audio".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio](
		"type", "input_audio",
	)
}

// The properties Data, Format are required.
type ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio struct {
	// The base64 encoded audio data
	Data string `json:"data,required"`
	// The format of the audio data
	//
	// Any of "wav".
	Format string `json:"format,omitzero,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio](
		"format", "wav",
	)
}

// The property Role is required.
type ChatCompletionNewParamsMessageChatCompletionAssistantMessageParam struct {
	// Any of "assistant".
	Role    string            `json:"role,omitzero,required"`
	Content param.Opt[string] `json:"content,omitzero"`
	Name    param.Opt[string] `json:"name,omitzero"`
	// Deprecated: deprecated
	FunctionCall ChatCompletionNewParamsMessageChatCompletionAssistantMessageParamFunctionCall `json:"function_call,omitzero"`
	ToolCalls    []ToolChoiceParam                                                             `json:"tool_calls,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionAssistantMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionAssistantMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionAssistantMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageChatCompletionAssistantMessageParam](
		"role", "assistant",
	)
}

// Deprecated: deprecated
//
// The properties Arguments, Name are required.
type ChatCompletionNewParamsMessageChatCompletionAssistantMessageParamFunctionCall struct {
	Arguments string `json:"arguments,required"`
	Name      string `json:"name,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionAssistantMessageParamFunctionCall) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionAssistantMessageParamFunctionCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionAssistantMessageParamFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Role, ToolCallID are required.
type ChatCompletionNewParamsMessageChatCompletionToolMessageParam struct {
	Content string `json:"content,required"`
	// Any of "tool".
	Role       string            `json:"role,omitzero,required"`
	ToolCallID string            `json:"tool_call_id,required"`
	Name       param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionToolMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionToolMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionToolMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageChatCompletionToolMessageParam](
		"role", "tool",
	)
}

// Deprecated: deprecated
//
// The properties Content, Name, Role are required.
type ChatCompletionNewParamsMessageChatCompletionFunctionMessageParam struct {
	Content string `json:"content,required"`
	Name    string `json:"name,required"`
	// Any of "function".
	Role string `json:"role,omitzero,required"`
	paramObj
}

func (r ChatCompletionNewParamsMessageChatCompletionFunctionMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMessageChatCompletionFunctionMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsMessageChatCompletionFunctionMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsMessageChatCompletionFunctionMessageParam](
		"role", "function",
	)
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

// Defined the behavior of the API when max_tokens exceed the maximum context
// length of the model. When set to 'error', API will return 400 with appropriate
// error message. When set to 'truncate', override the max_tokens with maximum
// context length of the model.
type ChatCompletionNewParamsContextLengthExceededBehavior string

const (
	ChatCompletionNewParamsContextLengthExceededBehaviorTruncate ChatCompletionNewParamsContextLengthExceededBehavior = "truncate"
	ChatCompletionNewParamsContextLengthExceededBehaviorError    ChatCompletionNewParamsContextLengthExceededBehavior = "error"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsFunctionCallUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfChatCompletionNewsFunctionCallString)
	OfChatCompletionNewsFunctionCallString param.Opt[string]                        `json:",omitzero,inline"`
	OfChatCompletionNewsFunctionCallName   *ChatCompletionNewParamsFunctionCallName `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsFunctionCallUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfChatCompletionNewsFunctionCallString, u.OfChatCompletionNewsFunctionCallName)
}
func (u *ChatCompletionNewParamsFunctionCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsFunctionCallUnion) asAny() any {
	if !param.IsOmitted(u.OfChatCompletionNewsFunctionCallString) {
		return &u.OfChatCompletionNewsFunctionCallString
	} else if !param.IsOmitted(u.OfChatCompletionNewsFunctionCallName) {
		return u.OfChatCompletionNewsFunctionCallName
	}
	return nil
}

type ChatCompletionNewParamsFunctionCallString string

const (
	ChatCompletionNewParamsFunctionCallStringNone ChatCompletionNewParamsFunctionCallString = "none"
	ChatCompletionNewParamsFunctionCallStringAuto ChatCompletionNewParamsFunctionCallString = "auto"
)

// The property Name is required.
type ChatCompletionNewParamsFunctionCallName struct {
	Name string `json:"name,required"`
	paramObj
}

func (r ChatCompletionNewParamsFunctionCallName) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsFunctionCallName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsFunctionCallName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the level of reasoning effort the model should apply when generating
// responses. Higher values may result in more thoughtful and detailed responses
// but may take longer to generate.
type ChatCompletionNewParamsReasoningEffort string

const (
	ChatCompletionNewParamsReasoningEffortLow    ChatCompletionNewParamsReasoningEffort = "low"
	ChatCompletionNewParamsReasoningEffortMedium ChatCompletionNewParamsReasoningEffort = "medium"
	ChatCompletionNewParamsReasoningEffortHigh   ChatCompletionNewParamsReasoningEffort = "high"
)

// An object specifying the format that the model must output.
type ChatCompletionNewParamsResponseFormat struct {
	// The type of the response format.
	Type param.Opt[string] `json:"type,omitzero"`
	// The schema of the response format.
	Schema map[string]any `json:"schema,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsToolChoiceUnion struct {
	OfString     param.Opt[string] `json:",omitzero,inline"`
	OfToolChoice *ToolChoiceParam  `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsToolChoiceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfToolChoice)
}
func (u *ChatCompletionNewParamsToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsToolChoiceUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfToolChoice) {
		return u.OfToolChoice
	}
	return nil
}
