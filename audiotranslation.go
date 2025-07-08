// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/tidwall/gjson"
	"github.com/togethercomputer/together-go/internal/apiform"
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/ssestream"
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
func NewAudioTranslationService(opts ...option.RequestOption) (r *AudioTranslationService) {
	r = &AudioTranslationService{}
	r.Options = opts
	return
}

// Translates audio into English
func (r *AudioTranslationService) NewStreaming(ctx context.Context, body AudioTranslationNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[AudioSpeechStreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	path := "audio/translations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[AudioSpeechStreamChunk](ssestream.NewDecoder(raw), err)
}

type AudioTranslationNewResponse struct {
	// The translated text
	Text string `json:"text,required"`
	// The duration of the audio in seconds
	Duration float64 `json:"duration"`
	// The target language of the translation
	Language string `json:"language"`
	// This field can have the runtime type of
	// [[]AudioTranslationNewResponseAudioTranslationVerboseJsonResponseSegment].
	Segments interface{} `json:"segments"`
	// The task performed
	Task AudioTranslationNewResponseTask `json:"task"`
	// This field can have the runtime type of
	// [[]AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord].
	Words interface{}                     `json:"words"`
	JSON  audioTranslationNewResponseJSON `json:"-"`
	union AudioTranslationNewResponseUnion
}

// audioTranslationNewResponseJSON contains the JSON metadata for the struct
// [AudioTranslationNewResponse]
type audioTranslationNewResponseJSON struct {
	Text        apijson.Field
	Duration    apijson.Field
	Language    apijson.Field
	Segments    apijson.Field
	Task        apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r audioTranslationNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AudioTranslationNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AudioTranslationNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AudioTranslationNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AudioTranslationNewResponseAudioTranslationJsonResponse],
// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponse].
func (r AudioTranslationNewResponse) AsUnion() AudioTranslationNewResponseUnion {
	return r.union
}

// Union satisfied by [AudioTranslationNewResponseAudioTranslationJsonResponse] or
// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponse].
type AudioTranslationNewResponseUnion interface {
	implementsAudioTranslationNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AudioTranslationNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AudioTranslationNewResponseAudioTranslationJsonResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AudioTranslationNewResponseAudioTranslationVerboseJsonResponse{}),
		},
	)
}

type AudioTranslationNewResponseAudioTranslationJsonResponse struct {
	// The translated text
	Text string                                                      `json:"text,required"`
	JSON audioTranslationNewResponseAudioTranslationJsonResponseJSON `json:"-"`
}

// audioTranslationNewResponseAudioTranslationJsonResponseJSON contains the JSON
// metadata for the struct
// [AudioTranslationNewResponseAudioTranslationJsonResponse]
type audioTranslationNewResponseAudioTranslationJsonResponseJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranslationNewResponseAudioTranslationJsonResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranslationNewResponseAudioTranslationJsonResponseJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranslationNewResponseAudioTranslationJsonResponse) implementsAudioTranslationNewResponse() {
}

type AudioTranslationNewResponseAudioTranslationVerboseJsonResponse struct {
	// The duration of the audio in seconds
	Duration float64 `json:"duration,required"`
	// The target language of the translation
	Language string `json:"language,required"`
	// Array of translation segments
	Segments []AudioTranslationNewResponseAudioTranslationVerboseJsonResponseSegment `json:"segments,required"`
	// The task performed
	Task AudioTranslationNewResponseAudioTranslationVerboseJsonResponseTask `json:"task,required"`
	// The translated text
	Text string `json:"text,required"`
	// Array of translation words (only when timestamp_granularities includes 'word')
	Words []AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord `json:"words"`
	JSON  audioTranslationNewResponseAudioTranslationVerboseJsonResponseJSON   `json:"-"`
}

// audioTranslationNewResponseAudioTranslationVerboseJsonResponseJSON contains the
// JSON metadata for the struct
// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponse]
type audioTranslationNewResponseAudioTranslationVerboseJsonResponseJSON struct {
	Duration    apijson.Field
	Language    apijson.Field
	Segments    apijson.Field
	Task        apijson.Field
	Text        apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranslationNewResponseAudioTranslationVerboseJsonResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranslationNewResponseAudioTranslationVerboseJsonResponseJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranslationNewResponseAudioTranslationVerboseJsonResponse) implementsAudioTranslationNewResponse() {
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
	// Array of token IDs for the segment
	Tokens []int64                                                                   `json:"tokens,required"`
	JSON   audioTranslationNewResponseAudioTranslationVerboseJsonResponseSegmentJSON `json:"-"`
}

// audioTranslationNewResponseAudioTranslationVerboseJsonResponseSegmentJSON
// contains the JSON metadata for the struct
// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponseSegment]
type audioTranslationNewResponseAudioTranslationVerboseJsonResponseSegmentJSON struct {
	ID          apijson.Field
	End         apijson.Field
	Start       apijson.Field
	Text        apijson.Field
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranslationNewResponseAudioTranslationVerboseJsonResponseSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranslationNewResponseAudioTranslationVerboseJsonResponseSegmentJSON) RawJSON() string {
	return r.raw
}

// The task performed
type AudioTranslationNewResponseAudioTranslationVerboseJsonResponseTask string

const (
	AudioTranslationNewResponseAudioTranslationVerboseJsonResponseTaskTranscribe AudioTranslationNewResponseAudioTranslationVerboseJsonResponseTask = "transcribe"
	AudioTranslationNewResponseAudioTranslationVerboseJsonResponseTaskTranslate  AudioTranslationNewResponseAudioTranslationVerboseJsonResponseTask = "translate"
)

func (r AudioTranslationNewResponseAudioTranslationVerboseJsonResponseTask) IsKnown() bool {
	switch r {
	case AudioTranslationNewResponseAudioTranslationVerboseJsonResponseTaskTranscribe, AudioTranslationNewResponseAudioTranslationVerboseJsonResponseTaskTranslate:
		return true
	}
	return false
}

type AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord struct {
	// End time of the word in seconds
	End float64 `json:"end,required"`
	// Start time of the word in seconds
	Start float64 `json:"start,required"`
	// The word
	Word string                                                                 `json:"word,required"`
	JSON audioTranslationNewResponseAudioTranslationVerboseJsonResponseWordJSON `json:"-"`
}

// audioTranslationNewResponseAudioTranslationVerboseJsonResponseWordJSON contains
// the JSON metadata for the struct
// [AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord]
type audioTranslationNewResponseAudioTranslationVerboseJsonResponseWordJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranslationNewResponseAudioTranslationVerboseJsonResponseWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranslationNewResponseAudioTranslationVerboseJsonResponseWordJSON) RawJSON() string {
	return r.raw
}

// The task performed
type AudioTranslationNewResponseTask string

const (
	AudioTranslationNewResponseTaskTranscribe AudioTranslationNewResponseTask = "transcribe"
	AudioTranslationNewResponseTaskTranslate  AudioTranslationNewResponseTask = "translate"
)

func (r AudioTranslationNewResponseTask) IsKnown() bool {
	switch r {
	case AudioTranslationNewResponseTaskTranscribe, AudioTranslationNewResponseTaskTranslate:
		return true
	}
	return false
}

type AudioTranslationNewParams struct {
	// Audio file to translate
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Target output language. Optional ISO 639-1 language code. If omitted, language
	// is set to English.
	Language param.Field[string] `json:"language"`
	// Model to use for translation
	Model param.Field[AudioTranslationNewParamsModel] `json:"model"`
	// Optional text to bias decoding.
	Prompt param.Field[string] `json:"prompt"`
	// The format of the response
	ResponseFormat param.Field[AudioTranslationNewParamsResponseFormat] `json:"response_format"`
	// Sampling temperature between 0.0 and 1.0
	Temperature param.Field[float64] `json:"temperature"`
	// Controls level of timestamp detail in verbose_json. Only used when
	// response_format is verbose_json.
	TimestampGranularities param.Field[AudioTranslationNewParamsTimestampGranularities] `json:"timestamp_granularities"`
}

func (r AudioTranslationNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
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

func (r AudioTranslationNewParamsModel) IsKnown() bool {
	switch r {
	case AudioTranslationNewParamsModelOpenAIWhisperLargeV3:
		return true
	}
	return false
}

// The format of the response
type AudioTranslationNewParamsResponseFormat string

const (
	AudioTranslationNewParamsResponseFormatJson        AudioTranslationNewParamsResponseFormat = "json"
	AudioTranslationNewParamsResponseFormatVerboseJson AudioTranslationNewParamsResponseFormat = "verbose_json"
)

func (r AudioTranslationNewParamsResponseFormat) IsKnown() bool {
	switch r {
	case AudioTranslationNewParamsResponseFormatJson, AudioTranslationNewParamsResponseFormatVerboseJson:
		return true
	}
	return false
}

// Controls level of timestamp detail in verbose_json. Only used when
// response_format is verbose_json.
type AudioTranslationNewParamsTimestampGranularities string

const (
	AudioTranslationNewParamsTimestampGranularitiesSegment AudioTranslationNewParamsTimestampGranularities = "segment"
	AudioTranslationNewParamsTimestampGranularitiesWord    AudioTranslationNewParamsTimestampGranularities = "word"
)

func (r AudioTranslationNewParamsTimestampGranularities) IsKnown() bool {
	switch r {
	case AudioTranslationNewParamsTimestampGranularitiesSegment, AudioTranslationNewParamsTimestampGranularitiesWord:
		return true
	}
	return false
}
