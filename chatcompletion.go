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
	"github.com/togethercomputer/together-go/shared/constant"
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

// Generate a model response for a given chat conversation. Supports single queries
// and multi-turn conversations with system, user, and assistant messages.
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *ChatCompletion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Generate a model response for a given chat conversation. Supports single queries
// and multi-turn conversations with system, user, and assistant messages.
func (r *ChatCompletionService) NewStreaming(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[ChatCompletionChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append(opts, option.WithJSONSet("stream", true))
	path := "chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ChatCompletionChunk](ssestream.NewDecoder(raw), err)
}

type ChatCompletion struct {
	ID      string                 `json:"id" api:"required"`
	Choices []ChatCompletionChoice `json:"choices" api:"required"`
	Created int64                  `json:"created" api:"required"`
	Model   string                 `json:"model" api:"required"`
	// The object type, which is always `chat.completion`.
	Object constant.ChatCompletion `json:"object" api:"required"`
	// When `echo` is true, the prompt is included in the response. Additionally, when
	// `logprobs` is also provided, log probability information is provided on the
	// prompt.
	Prompt   ChatCompletionPrompt    `json:"prompt" api:"required"`
	Usage    ChatCompletionUsage     `json:"usage" api:"nullable"`
	Warnings []ChatCompletionWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Choices     respjson.Field
		Created     respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		Prompt      respjson.Field
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
	FinishReason string                `json:"finish_reason"`
	Index        int64                 `json:"index"`
	Logprobs     LogProbs              `json:"logprobs" api:"nullable"`
	Message      ChatCompletionMessage `json:"message"`
	Seed         int64                 `json:"seed"`
	Text         string                `json:"text"`
	// Top log probabilities for the tokens.
	TopLogprobs map[string]float64 `json:"top_logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Logprobs     respjson.Field
		Message      respjson.Field
		Seed         respjson.Field
		Text         respjson.Field
		TopLogprobs  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionChunk struct {
	ID      string                      `json:"id" api:"required"`
	Choices []ChatCompletionChunkChoice `json:"choices" api:"required"`
	Created int64                       `json:"created" api:"required"`
	Model   string                      `json:"model" api:"required"`
	// The object type, which is always `chat.completion.chunk`.
	Object            constant.ChatCompletionChunk `json:"object" api:"required"`
	SystemFingerprint string                       `json:"system_fingerprint"`
	Usage             ChatCompletionUsage          `json:"usage" api:"nullable"`
	Warnings          []ChatCompletionWarning      `json:"warnings"`
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
	Delta ChatCompletionChunkChoiceDelta `json:"delta" api:"required"`
	// Any of "stop", "eos", "length", "tool_calls", "function_call".
	FinishReason string  `json:"finish_reason" api:"required"`
	Index        int64   `json:"index" api:"required"`
	Logprobs     float64 `json:"logprobs" api:"nullable"`
	Seed         int64   `json:"seed" api:"nullable"`
	// Top log probabilities for the tokens.
	TopLogprobs map[string]float64 `json:"top_logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta        respjson.Field
		FinishReason respjson.Field
		Index        respjson.Field
		Logprobs     respjson.Field
		Seed         respjson.Field
		TopLogprobs  respjson.Field
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
	Role    string `json:"role" api:"required"`
	Content string `json:"content" api:"nullable"`
	// Deprecated: deprecated
	FunctionCall ChatCompletionChunkChoiceDeltaFunctionCall `json:"function_call" api:"nullable"`
	Reasoning    string                                     `json:"reasoning" api:"nullable"`
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
	Arguments string `json:"arguments" api:"required"`
	Name      string `json:"name" api:"required"`
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

// Deprecated: deprecated
//
// The properties Content, Name, Role are required.
type ChatCompletionFunctionMessageParam struct {
	Content string `json:"content" api:"required"`
	Name    string `json:"name" api:"required"`
	// Any of "function".
	Role ChatCompletionFunctionMessageParamRole `json:"role,omitzero" api:"required"`
	paramObj
}

func (r ChatCompletionFunctionMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionFunctionMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionFunctionMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionFunctionMessageParamRole string

const (
	ChatCompletionFunctionMessageParamRoleFunction ChatCompletionFunctionMessageParamRole = "function"
)

type ChatCompletionMessage struct {
	Content string `json:"content" api:"required"`
	// Any of "assistant".
	Role ChatCompletionMessageRole `json:"role" api:"required"`
	// Deprecated: deprecated
	FunctionCall ChatCompletionMessageFunctionCall `json:"function_call"`
	Reasoning    string                            `json:"reasoning" api:"nullable"`
	ToolCalls    []ToolChoice                      `json:"tool_calls"`
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
func (r ChatCompletionMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionMessageRole string

const (
	ChatCompletionMessageRoleAssistant ChatCompletionMessageRole = "assistant"
)

// Deprecated: deprecated
type ChatCompletionMessageFunctionCall struct {
	Arguments string `json:"arguments" api:"required"`
	Name      string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionMessageFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionMessageFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ChatCompletionMessageParamOfChatCompletionSystemMessage(content string, role ChatCompletionSystemMessageParamRole) ChatCompletionMessageParamUnion {
	var variant ChatCompletionSystemMessageParam
	variant.Content = content
	variant.Role = role
	return ChatCompletionMessageParamUnion{OfChatCompletionSystemMessage: &variant}
}

func ChatCompletionMessageParamOfChatCompletionMessageChatCompletionUserMessageParam[
	T string | []ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion,
](content T, role string) ChatCompletionMessageParamUnion {
	var variant ChatCompletionMessageParamChatCompletionUserMessageParam
	switch v := any(content).(type) {
	case string:
		variant.Content.OfString = param.NewOpt(v)
	case []ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion:
		variant.Content.OfChatCompletionMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalArray = v
	}
	variant.Role = role
	return ChatCompletionMessageParamUnion{OfChatCompletionMessageChatCompletionUserMessageParam: &variant}
}

func ChatCompletionMessageParamOfChatCompletionMessageChatCompletionAssistantMessageParam(role string) ChatCompletionMessageParamUnion {
	var variant ChatCompletionMessageParamChatCompletionAssistantMessageParam
	variant.Role = role
	return ChatCompletionMessageParamUnion{OfChatCompletionMessageChatCompletionAssistantMessageParam: &variant}
}

func ChatCompletionMessageParamOfChatCompletionToolMessage(content string, role ChatCompletionToolMessageParamRole, toolCallID string) ChatCompletionMessageParamUnion {
	var variant ChatCompletionToolMessageParam
	variant.Content = content
	variant.Role = role
	variant.ToolCallID = toolCallID
	return ChatCompletionMessageParamUnion{OfChatCompletionToolMessage: &variant}
}

func ChatCompletionMessageParamOfChatCompletionFunctionMessage(content string, name string, role ChatCompletionFunctionMessageParamRole) ChatCompletionMessageParamUnion {
	var variant ChatCompletionFunctionMessageParam
	variant.Content = content
	variant.Name = name
	variant.Role = role
	return ChatCompletionMessageParamUnion{OfChatCompletionFunctionMessage: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionMessageParamUnion struct {
	OfChatCompletionSystemMessage                              *ChatCompletionSystemMessageParam                              `json:",omitzero,inline"`
	OfChatCompletionMessageChatCompletionUserMessageParam      *ChatCompletionMessageParamChatCompletionUserMessageParam      `json:",omitzero,inline"`
	OfChatCompletionMessageChatCompletionAssistantMessageParam *ChatCompletionMessageParamChatCompletionAssistantMessageParam `json:",omitzero,inline"`
	OfChatCompletionToolMessage                                *ChatCompletionToolMessageParam                                `json:",omitzero,inline"`
	OfChatCompletionFunctionMessage                            *ChatCompletionFunctionMessageParam                            `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionMessageParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfChatCompletionSystemMessage,
		u.OfChatCompletionMessageChatCompletionUserMessageParam,
		u.OfChatCompletionMessageChatCompletionAssistantMessageParam,
		u.OfChatCompletionToolMessage,
		u.OfChatCompletionFunctionMessage)
}
func (u *ChatCompletionMessageParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionMessageParamUnion) asAny() any {
	if !param.IsOmitted(u.OfChatCompletionSystemMessage) {
		return u.OfChatCompletionSystemMessage
	} else if !param.IsOmitted(u.OfChatCompletionMessageChatCompletionUserMessageParam) {
		return u.OfChatCompletionMessageChatCompletionUserMessageParam
	} else if !param.IsOmitted(u.OfChatCompletionMessageChatCompletionAssistantMessageParam) {
		return u.OfChatCompletionMessageChatCompletionAssistantMessageParam
	} else if !param.IsOmitted(u.OfChatCompletionToolMessage) {
		return u.OfChatCompletionToolMessage
	} else if !param.IsOmitted(u.OfChatCompletionFunctionMessage) {
		return u.OfChatCompletionFunctionMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamUnion) GetFunctionCall() *ChatCompletionMessageParamChatCompletionAssistantMessageParamFunctionCall {
	if vt := u.OfChatCompletionMessageChatCompletionAssistantMessageParam; vt != nil {
		return &vt.FunctionCall
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamUnion) GetToolCalls() []ToolChoiceParam {
	if vt := u.OfChatCompletionMessageChatCompletionAssistantMessageParam; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamUnion) GetToolCallID() *string {
	if vt := u.OfChatCompletionToolMessage; vt != nil {
		return &vt.ToolCallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamUnion) GetRole() *string {
	if vt := u.OfChatCompletionSystemMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfChatCompletionMessageChatCompletionUserMessageParam; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfChatCompletionMessageChatCompletionAssistantMessageParam; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfChatCompletionToolMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfChatCompletionFunctionMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamUnion) GetName() *string {
	if vt := u.OfChatCompletionSystemMessage; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfChatCompletionMessageChatCompletionUserMessageParam; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfChatCompletionMessageChatCompletionAssistantMessageParam; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfChatCompletionToolMessage; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfChatCompletionFunctionMessage; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ChatCompletionMessageParamUnion) GetContent() (res chatCompletionMessageParamUnionContent) {
	if vt := u.OfChatCompletionSystemMessage; vt != nil {
		res.any = &vt.Content
	} else if vt := u.OfChatCompletionMessageChatCompletionUserMessageParam; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfChatCompletionMessageChatCompletionAssistantMessageParam; vt != nil && vt.Content.Valid() {
		res.any = &vt.Content.Value
	} else if vt := u.OfChatCompletionToolMessage; vt != nil {
		res.any = &vt.Content
	} else if vt := u.OfChatCompletionFunctionMessage; vt != nil {
		res.any = &vt.Content
	}
	return
}

// Can have the runtime types [*string],
// [\*[]ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion]
type chatCompletionMessageParamUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]together.ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u chatCompletionMessageParamUnionContent) AsAny() any { return u.any }

// The properties Content, Role are required.
type ChatCompletionMessageParamChatCompletionUserMessageParam struct {
	// The content of the message, which can either be a simple string or a structured
	// format.
	Content ChatCompletionMessageParamChatCompletionUserMessageParamContentUnion `json:"content,omitzero" api:"required"`
	// Any of "user".
	Role string            `json:"role,omitzero" api:"required"`
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ChatCompletionMessageParamChatCompletionUserMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageParamChatCompletionUserMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageParamChatCompletionUserMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionMessageParamChatCompletionUserMessageParam](
		"role", "user",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionMessageParamChatCompletionUserMessageParamContentUnion struct {
	OfString                                                                                                    param.Opt[string]                                                                                                    `json:",omitzero,inline"`
	OfChatCompletionMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalArray []ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionMessageParamChatCompletionUserMessageParamContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfChatCompletionMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalArray)
}
func (u *ChatCompletionMessageParamChatCompletionUserMessageParamContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionMessageParamChatCompletionUserMessageParamContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfChatCompletionMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalArray) {
		return &u.OfChatCompletionMessageChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion struct {
	OfChatCompletionStructuredMessageText     *ChatCompletionStructuredMessageTextParam                                                                                `json:",omitzero,inline"`
	OfChatCompletionStructuredMessageImageURL *ChatCompletionStructuredMessageImageURLParam                                                                            `json:",omitzero,inline"`
	OfVideo                                   *ChatCompletionStructuredMessageVideoURLParam                                                                            `json:",omitzero,inline"`
	OfAudio                                   *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio      `json:",omitzero,inline"`
	OfInputAudio                              *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfChatCompletionStructuredMessageText,
		u.OfChatCompletionStructuredMessageImageURL,
		u.OfVideo,
		u.OfAudio,
		u.OfInputAudio)
}
func (u *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) asAny() any {
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
func (u ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetText() *string {
	if vt := u.OfChatCompletionStructuredMessageText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetImageURL() *ChatCompletionStructuredMessageImageURLImageURLParam {
	if vt := u.OfChatCompletionStructuredMessageImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetVideoURL() *ChatCompletionStructuredMessageVideoURLVideoURLParam {
	if vt := u.OfVideo; vt != nil {
		return &vt.VideoURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetAudioURL() *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL {
	if vt := u.OfAudio; vt != nil {
		return &vt.AudioURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetInputAudio() *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio {
	if vt := u.OfInputAudio; vt != nil {
		return &vt.InputAudio
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemUnion) GetType() *string {
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
type ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio struct {
	AudioURL ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL `json:"audio_url,omitzero" api:"required"`
	// Any of "audio_url".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudio](
		"type", "audio_url",
	)
}

// The property URL is required.
type ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL struct {
	// The URL of the audio
	URL string `json:"url" api:"required"`
	paramObj
}

func (r ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemAudioAudioURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties InputAudio, Type are required.
type ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio struct {
	InputAudio ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio `json:"input_audio,omitzero" api:"required"`
	// Any of "input_audio".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudio](
		"type", "input_audio",
	)
}

// The properties Data, Format are required.
type ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio struct {
	// The base64 encoded audio data
	Data string `json:"data" api:"required"`
	// The format of the audio data
	//
	// Any of "wav".
	Format string `json:"format,omitzero" api:"required"`
	paramObj
}

func (r ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionMessageParamChatCompletionUserMessageParamContentChatCompletionUserMessageContentMultimodalItemInputAudioInputAudio](
		"format", "wav",
	)
}

// The property Role is required.
type ChatCompletionMessageParamChatCompletionAssistantMessageParam struct {
	// Any of "assistant".
	Role    string            `json:"role,omitzero" api:"required"`
	Content param.Opt[string] `json:"content,omitzero"`
	Name    param.Opt[string] `json:"name,omitzero"`
	// Deprecated: deprecated
	FunctionCall ChatCompletionMessageParamChatCompletionAssistantMessageParamFunctionCall `json:"function_call,omitzero"`
	ToolCalls    []ToolChoiceParam                                                         `json:"tool_calls,omitzero"`
	paramObj
}

func (r ChatCompletionMessageParamChatCompletionAssistantMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageParamChatCompletionAssistantMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageParamChatCompletionAssistantMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionMessageParamChatCompletionAssistantMessageParam](
		"role", "assistant",
	)
}

// Deprecated: deprecated
//
// The properties Arguments, Name are required.
type ChatCompletionMessageParamChatCompletionAssistantMessageParamFunctionCall struct {
	Arguments string `json:"arguments" api:"required"`
	Name      string `json:"name" api:"required"`
	paramObj
}

func (r ChatCompletionMessageParamChatCompletionAssistantMessageParamFunctionCall) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageParamChatCompletionAssistantMessageParamFunctionCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageParamChatCompletionAssistantMessageParamFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionPrompt []ChatCompletionPromptItem

type ChatCompletionPromptItem struct {
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
func (r ChatCompletionPromptItem) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionPromptItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	URL string `json:"url" api:"required"`
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
	Text string `json:"text" api:"required"`
	// Any of "text".
	Type ChatCompletionStructuredMessageTextType `json:"type,omitzero" api:"required"`
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
	Type     ChatCompletionStructuredMessageVideoURLType          `json:"type,omitzero" api:"required"`
	VideoURL ChatCompletionStructuredMessageVideoURLVideoURLParam `json:"video_url,omitzero" api:"required"`
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
	URL string `json:"url" api:"required"`
	paramObj
}

func (r ChatCompletionStructuredMessageVideoURLVideoURLParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionStructuredMessageVideoURLVideoURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionStructuredMessageVideoURLVideoURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Role are required.
type ChatCompletionSystemMessageParam struct {
	Content string `json:"content" api:"required"`
	// Any of "system".
	Role ChatCompletionSystemMessageParamRole `json:"role,omitzero" api:"required"`
	Name param.Opt[string]                    `json:"name,omitzero"`
	paramObj
}

func (r ChatCompletionSystemMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionSystemMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionSystemMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionSystemMessageParamRole string

const (
	ChatCompletionSystemMessageParamRoleSystem ChatCompletionSystemMessageParamRole = "system"
)

// The properties Content, Role, ToolCallID are required.
type ChatCompletionToolMessageParam struct {
	Content string `json:"content" api:"required"`
	// Any of "tool".
	Role       ChatCompletionToolMessageParamRole `json:"role,omitzero" api:"required"`
	ToolCallID string                             `json:"tool_call_id" api:"required"`
	Name       param.Opt[string]                  `json:"name,omitzero"`
	paramObj
}

func (r ChatCompletionToolMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionToolMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionToolMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionToolMessageParamRole string

const (
	ChatCompletionToolMessageParamRoleTool ChatCompletionToolMessageParamRole = "tool"
)

type ChatCompletionUsage struct {
	CompletionTokens int64 `json:"completion_tokens" api:"required"`
	PromptTokens     int64 `json:"prompt_tokens" api:"required"`
	TotalTokens      int64 `json:"total_tokens" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Messages []ChatCompletionMessageParamUnion `json:"messages,omitzero" api:"required"`
	// The name of the model to query.
	//
	// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#chat-models)
	Model string `json:"model" api:"required"`
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
	// Additional configuration to pass to model engine.
	ChatTemplateKwargs any `json:"chat_template_kwargs,omitzero"`
	// Any of "hipaa".
	Compliance ChatCompletionNewParamsCompliance `json:"compliance,omitzero"`
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
	// For models that support toggling reasoning functionality, this object can be
	// used to control that functionality.
	Reasoning ChatCompletionNewParamsReasoning `json:"reasoning,omitzero"`
	// Controls the level of reasoning effort the model should apply when generating
	// responses. Higher values may result in more thoughtful and detailed responses
	// but may take longer to generate.
	//
	// Any of "low", "medium", "high".
	ReasoningEffort ChatCompletionNewParamsReasoningEffort `json:"reasoning_effort,omitzero"`
	// An object specifying the format that the model must output.
	//
	// Setting to `{ "type": "json_schema", "json_schema": {...} }` enables Structured
	// Outputs which ensures the model will match your supplied JSON schema. Learn more
	// in the [Structured Outputs guide](https://docs.together.ai/docs/json-mode).
	//
	// Setting to `{ "type": "json_object" }` enables the older JSON mode, which
	// ensures the message the model generates is valid JSON. Using `json_schema` is
	// preferred for models that support it.
	ResponseFormat ChatCompletionNewParamsResponseFormatUnion `json:"response_format,omitzero"`
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

type ChatCompletionNewParamsCompliance string

const (
	ChatCompletionNewParamsComplianceHipaa ChatCompletionNewParamsCompliance = "hipaa"
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
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ChatCompletionNewParamsFunctionCallName) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsFunctionCallName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsFunctionCallName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// For models that support toggling reasoning functionality, this object can be
// used to control that functionality.
type ChatCompletionNewParamsReasoning struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsReasoning) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsReasoning
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsReasoning) UnmarshalJSON(data []byte) error {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsResponseFormatUnion struct {
	OfText       *ChatCompletionNewParamsResponseFormatText       `json:",omitzero,inline"`
	OfJsonSchema *ChatCompletionNewParamsResponseFormatJsonSchema `json:",omitzero,inline"`
	OfJsonObject *ChatCompletionNewParamsResponseFormatJsonObject `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsResponseFormatUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfJsonSchema, u.OfJsonObject)
}
func (u *ChatCompletionNewParamsResponseFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsResponseFormatUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfJsonSchema) {
		return u.OfJsonSchema
	} else if !param.IsOmitted(u.OfJsonObject) {
		return u.OfJsonObject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsResponseFormatUnion) GetJsonSchema() *ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema {
	if vt := u.OfJsonSchema; vt != nil {
		return &vt.JsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsResponseFormatUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJsonSchema; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfJsonObject; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsResponseFormatUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsResponseFormatText]("text"),
		apijson.Discriminator[ChatCompletionNewParamsResponseFormatJsonSchema]("json_schema"),
		apijson.Discriminator[ChatCompletionNewParamsResponseFormatJsonObject]("json_object"),
	)
}

func NewChatCompletionNewParamsResponseFormatText() ChatCompletionNewParamsResponseFormatText {
	return ChatCompletionNewParamsResponseFormatText{
		Type: "text",
	}
}

// Default response format. Used to generate text responses.
//
// This struct has a constant value, construct it with
// [NewChatCompletionNewParamsResponseFormatText].
type ChatCompletionNewParamsResponseFormatText struct {
	// The type of response format being defined. Always `text`.
	Type constant.Text `json:"type" api:"required"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatText) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JSON Schema response format. Used to generate structured JSON responses. Learn
// more about [Structured Outputs](https://docs.together.ai/docs/json-mode).
//
// The properties JsonSchema, Type are required.
type ChatCompletionNewParamsResponseFormatJsonSchema struct {
	// Structured Outputs configuration options, including a JSON Schema.
	JsonSchema ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema `json:"json_schema,omitzero" api:"required"`
	// The type of response format being defined. Always `json_schema`.
	//
	// This field can be elided, and will marshal its zero value as "json_schema".
	Type constant.JsonSchema `json:"type" api:"required"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured Outputs configuration options, including a JSON Schema.
//
// The property Name is required.
type ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema struct {
	// The name of the response format. Must be a-z, A-Z, 0-9, or contain underscores
	// and dashes, with a maximum length of 64.
	Name string `json:"name" api:"required"`
	// Whether to enable strict schema adherence when generating the output. If set to
	// true, the model will always follow the exact schema defined in the `schema`
	// field. Only a subset of JSON Schema is supported when `strict` is `true`. To
	// learn more, read the
	// [Structured Outputs guide](https://docs.together.ai/docs/json-mode).
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// A description of what the response format is for, used by the model to determine
	// how to respond in the format.
	Description param.Opt[string] `json:"description,omitzero"`
	// The schema for the response format, described as a JSON Schema object. Learn how
	// to build JSON schemas [here](https://json-schema.org/).
	Schema map[string]any `json:"schema,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatJsonSchemaJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewChatCompletionNewParamsResponseFormatJsonObject() ChatCompletionNewParamsResponseFormatJsonObject {
	return ChatCompletionNewParamsResponseFormatJsonObject{
		Type: "json_object",
	}
}

// JSON object response format. An older method of generating JSON responses. Using
// `json_schema` is recommended for models that support it. Note that the model
// will not generate JSON without a system or user message instructing it to do so.
//
// This struct has a constant value, construct it with
// [NewChatCompletionNewParamsResponseFormatJsonObject].
type ChatCompletionNewParamsResponseFormatJsonObject struct {
	// The type of response format being defined. Always `json_object`.
	Type constant.JsonObject `json:"type" api:"required"`
	paramObj
}

func (r ChatCompletionNewParamsResponseFormatJsonObject) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsResponseFormatJsonObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsResponseFormatJsonObject) UnmarshalJSON(data []byte) error {
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
