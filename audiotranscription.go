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

// AudioTranscriptionService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioTranscriptionService] method instead.
type AudioTranscriptionService struct {
	Options []option.RequestOption
}

// NewAudioTranscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAudioTranscriptionService(opts ...option.RequestOption) (r AudioTranscriptionService) {
	r = AudioTranscriptionService{}
	r.Options = opts
	return
}

// Transcribes audio into text
func (r *AudioTranscriptionService) New(ctx context.Context, body AudioTranscriptionNewParams, opts ...option.RequestOption) (res *AudioTranscriptionNewResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "audio/transcriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// AudioTranscriptionNewResponseUnion contains all possible properties and values
// from [AudioTranscriptionNewResponseAudioTranscriptionJsonResponse],
// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AudioTranscriptionNewResponseUnion struct {
	Text string `json:"text"`
	// This field is from variant
	// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse].
	Duration float64 `json:"duration"`
	// This field is from variant
	// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse].
	Language string `json:"language"`
	// This field is from variant
	// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse].
	Segments []AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegment `json:"segments"`
	// This field is from variant
	// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse].
	Task string `json:"task"`
	// This field is from variant
	// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse].
	Words []AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord `json:"words"`
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

func (u AudioTranscriptionNewResponseUnion) AsAudioTranscriptionNewResponseAudioTranscriptionJsonResponse() (v AudioTranscriptionNewResponseAudioTranscriptionJsonResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AudioTranscriptionNewResponseUnion) AsAudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse() (v AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AudioTranscriptionNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AudioTranscriptionNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranscriptionNewResponseAudioTranscriptionJsonResponse struct {
	// The transcribed text
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseAudioTranscriptionJsonResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseAudioTranscriptionJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse struct {
	// The duration of the audio in seconds
	Duration float64 `json:"duration,required"`
	// The language of the audio
	Language string `json:"language,required"`
	// Array of transcription segments
	Segments []AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegment `json:"segments,required"`
	// The task performed
	//
	// Any of "transcribe", "translate".
	Task string `json:"task,required"`
	// The transcribed text
	Text string `json:"text,required"`
	// Array of transcription words (only when timestamp_granularities includes 'word')
	Words []AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord `json:"words"`
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
func (r AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegment struct {
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
func (r AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord struct {
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
func (r AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranscriptionNewParams struct {
	// Audio file to transcribe
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Optional ISO 639-1 language code. If `auto` is provided, language is
	// auto-detected.
	Language param.Opt[string] `json:"language,omitzero"`
	// Optional text to bias decoding.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// Sampling temperature between 0.0 and 1.0
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Model to use for transcription
	//
	// Any of "openai/whisper-large-v3".
	Model AudioTranscriptionNewParamsModel `json:"model,omitzero"`
	// The format of the response
	//
	// Any of "json", "verbose_json".
	ResponseFormat AudioTranscriptionNewParamsResponseFormat `json:"response_format,omitzero"`
	// Controls level of timestamp detail in verbose_json. Only used when
	// response_format is verbose_json. Can be a single granularity or an array to get
	// multiple levels.
	TimestampGranularities AudioTranscriptionNewParamsTimestampGranularitiesUnion `json:"timestamp_granularities,omitzero"`
	paramObj
}

func (r AudioTranscriptionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// Model to use for transcription
type AudioTranscriptionNewParamsModel string

const (
	AudioTranscriptionNewParamsModelOpenAIWhisperLargeV3 AudioTranscriptionNewParamsModel = "openai/whisper-large-v3"
)

// The format of the response
type AudioTranscriptionNewParamsResponseFormat string

const (
	AudioTranscriptionNewParamsResponseFormatJson        AudioTranscriptionNewParamsResponseFormat = "json"
	AudioTranscriptionNewParamsResponseFormatVerboseJson AudioTranscriptionNewParamsResponseFormat = "verbose_json"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AudioTranscriptionNewParamsTimestampGranularitiesUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfAudioTranscriptionNewsTimestampGranularitiesString)
	OfAudioTranscriptionNewsTimestampGranularitiesString         param.Opt[string] `json:",omitzero,inline"`
	OfAudioTranscriptionNewsTimestampGranularitiesArrayItemArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AudioTranscriptionNewParamsTimestampGranularitiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAudioTranscriptionNewsTimestampGranularitiesString, u.OfAudioTranscriptionNewsTimestampGranularitiesArrayItemArray)
}
func (u *AudioTranscriptionNewParamsTimestampGranularitiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AudioTranscriptionNewParamsTimestampGranularitiesUnion) asAny() any {
	if !param.IsOmitted(u.OfAudioTranscriptionNewsTimestampGranularitiesString) {
		return &u.OfAudioTranscriptionNewsTimestampGranularitiesString
	} else if !param.IsOmitted(u.OfAudioTranscriptionNewsTimestampGranularitiesArrayItemArray) {
		return &u.OfAudioTranscriptionNewsTimestampGranularitiesArrayItemArray
	}
	return nil
}

type AudioTranscriptionNewParamsTimestampGranularitiesString string

const (
	AudioTranscriptionNewParamsTimestampGranularitiesStringSegment AudioTranscriptionNewParamsTimestampGranularitiesString = "segment"
	AudioTranscriptionNewParamsTimestampGranularitiesStringWord    AudioTranscriptionNewParamsTimestampGranularitiesString = "word"
)
