// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apiform"
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// AudioTranslationService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioTranslationService] method instead.
type AudioTranslationService struct {
	Options []option.RequestOption
}

// NewAudioTranslationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAudioTranslationService(opts ...option.RequestOption) (r AudioTranslationService) {
	r = AudioTranslationService{}
	r.Options = opts
	return
}

// Translates audio into English
func (r *AudioTranslationService) New(ctx context.Context, body AudioTranslationNewParams, opts ...option.RequestOption) (res *AudioTranslationNewResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "audio/translations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// AudioTranslationNewResponseUnion contains all possible properties and values
// from [AudioTranslationNewResponseAudioTranslationJsonResponse],
// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AudioTranslationNewResponseUnion struct {
	Text string `json:"text"`
	// This field is from variant
	// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponse].
	Duration float64 `json:"duration"`
	// This field is from variant
	// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponse].
	Language string `json:"language"`
	// This field is from variant
	// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponse].
	Segments []AudioTranslationNewResponseAudioTranslationVerboseJsonResponseSegment `json:"segments"`
	// This field is from variant
	// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponse].
	Task string `json:"task"`
	// This field is from variant
	// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponse].
	Words []AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord `json:"words"`
	JSON  struct {
		Text     respjson.Field
		Duration respjson.Field
		Language respjson.Field
		Segments respjson.Field
		Task     respjson.Field
		Words    respjson.Field
		raw      string
	} `json:"-"`
}

func (u AudioTranslationNewResponseUnion) AsAudioTranslationNewResponseAudioTranslationJsonResponse() (v AudioTranslationNewResponseAudioTranslationJsonResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AudioTranslationNewResponseUnion) AsAudioTranslationNewResponseAudioTranslationVerboseJsonResponse() (v AudioTranslationNewResponseAudioTranslationVerboseJsonResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AudioTranslationNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AudioTranslationNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranslationNewResponseAudioTranslationJsonResponse struct {
	// The translated text
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranslationNewResponseAudioTranslationJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *AudioTranslationNewResponseAudioTranslationJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranslationNewResponseAudioTranslationVerboseJsonResponse struct {
	// The duration of the audio in seconds
	Duration float64 `json:"duration,required"`
	// The target language of the translation
	Language string `json:"language,required"`
	// Array of translation segments
	Segments []AudioTranslationNewResponseAudioTranslationVerboseJsonResponseSegment `json:"segments,required"`
	// The task performed
	//
	// Any of "transcribe", "translate".
	Task string `json:"task,required"`
	// The translated text
	Text string `json:"text,required"`
	// Array of translation words (only when timestamp_granularities includes 'word')
	Words []AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord `json:"words"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Language    respjson.Field
		Segments    respjson.Field
		Task        respjson.Field
		Text        respjson.Field
		Words       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranslationNewResponseAudioTranslationVerboseJsonResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranslationNewResponseAudioTranslationVerboseJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranslationNewResponseAudioTranslationVerboseJsonResponseSegment struct {
	// Unique identifier for the segment
	ID int64 `json:"id,required"`
	// End time of the segment in seconds
	End float64 `json:"end,required"`
	// Start time of the segment in seconds
	Start float64 `json:"start,required"`
	// The text content of the segment
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		End         respjson.Field
		Start       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranslationNewResponseAudioTranslationVerboseJsonResponseSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranslationNewResponseAudioTranslationVerboseJsonResponseSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord struct {
	// End time of the word in seconds
	End float64 `json:"end,required"`
	// Start time of the word in seconds
	Start float64 `json:"start,required"`
	// The word
	Word string `json:"word,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		Word        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranslationNewParams struct {
	// Audio file to translate
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Target output language. Optional ISO 639-1 language code. If omitted, language
	// is set to English.
	Language param.Opt[string] `json:"language,omitzero"`
	// Optional text to bias decoding.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// Sampling temperature between 0.0 and 1.0
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Model to use for translation
	//
	// Any of "openai/whisper-large-v3".
	Model AudioTranslationNewParamsModel `json:"model,omitzero"`
	// The format of the response
	//
	// Any of "json", "verbose_json".
	ResponseFormat AudioTranslationNewParamsResponseFormat `json:"response_format,omitzero"`
	// Controls level of timestamp detail in verbose_json. Only used when
	// response_format is verbose_json. Can be a single granularity or an array to get
	// multiple levels.
	TimestampGranularities AudioTranslationNewParamsTimestampGranularitiesUnion `json:"timestamp_granularities,omitzero"`
	paramObj
}

func (r AudioTranslationNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Model to use for translation
type AudioTranslationNewParamsModel string

const (
	AudioTranslationNewParamsModelOpenAIWhisperLargeV3 AudioTranslationNewParamsModel = "openai/whisper-large-v3"
)

// The format of the response
type AudioTranslationNewParamsResponseFormat string

const (
	AudioTranslationNewParamsResponseFormatJson        AudioTranslationNewParamsResponseFormat = "json"
	AudioTranslationNewParamsResponseFormatVerboseJson AudioTranslationNewParamsResponseFormat = "verbose_json"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AudioTranslationNewParamsTimestampGranularitiesUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfAudioTranslationNewsTimestampGranularitiesString)
	OfAudioTranslationNewsTimestampGranularitiesString         param.Opt[string] `json:",omitzero,inline"`
	OfAudioTranslationNewsTimestampGranularitiesArrayItemArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AudioTranslationNewParamsTimestampGranularitiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAudioTranslationNewsTimestampGranularitiesString, u.OfAudioTranslationNewsTimestampGranularitiesArrayItemArray)
}
func (u *AudioTranslationNewParamsTimestampGranularitiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AudioTranslationNewParamsTimestampGranularitiesUnion) asAny() any {
	if !param.IsOmitted(u.OfAudioTranslationNewsTimestampGranularitiesString) {
		return &u.OfAudioTranslationNewsTimestampGranularitiesString
	} else if !param.IsOmitted(u.OfAudioTranslationNewsTimestampGranularitiesArrayItemArray) {
		return &u.OfAudioTranslationNewsTimestampGranularitiesArrayItemArray
	}
	return nil
}

type AudioTranslationNewParamsTimestampGranularitiesString string

const (
	AudioTranslationNewParamsTimestampGranularitiesStringSegment AudioTranslationNewParamsTimestampGranularitiesString = "segment"
	AudioTranslationNewParamsTimestampGranularitiesStringWord    AudioTranslationNewParamsTimestampGranularitiesString = "word"
)
