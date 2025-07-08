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
func NewAudioTranscriptionService(opts ...option.RequestOption) (r *AudioTranscriptionService) {
	r = &AudioTranscriptionService{}
	r.Options = opts
	return
}

// Transcribes audio into text
func (r *AudioTranscriptionService) NewStreaming(ctx context.Context, body AudioTranscriptionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[AudioSpeechStreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	path := "audio/transcriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[AudioSpeechStreamChunk](ssestream.NewDecoder(raw), err)
}

type AudioTranscriptionNewResponse struct {
	// The transcribed text
	Text string `json:"text,required"`
	// The duration of the audio in seconds
	Duration float64 `json:"duration"`
	// The language of the audio
	Language string `json:"language"`
	// This field can have the runtime type of
	// [[]AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegment].
	Segments interface{} `json:"segments"`
	// The task performed
	Task AudioTranscriptionNewResponseTask `json:"task"`
	// This field can have the runtime type of
	// [[]AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord].
	Words interface{}                       `json:"words"`
	JSON  audioTranscriptionNewResponseJSON `json:"-"`
	union AudioTranscriptionNewResponseUnion
}

// audioTranscriptionNewResponseJSON contains the JSON metadata for the struct
// [AudioTranscriptionNewResponse]
type audioTranscriptionNewResponseJSON struct {
	Text        apijson.Field
	Duration    apijson.Field
	Language    apijson.Field
	Segments    apijson.Field
	Task        apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r audioTranscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AudioTranscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AudioTranscriptionNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AudioTranscriptionNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AudioTranscriptionNewResponseAudioTranscriptionJsonResponse],
// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse].
func (r AudioTranscriptionNewResponse) AsUnion() AudioTranscriptionNewResponseUnion {
	return r.union
}

// Union satisfied by [AudioTranscriptionNewResponseAudioTranscriptionJsonResponse]
// or [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse].
type AudioTranscriptionNewResponseUnion interface {
	implementsAudioTranscriptionNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AudioTranscriptionNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AudioTranscriptionNewResponseAudioTranscriptionJsonResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse{}),
		},
	)
}

type AudioTranscriptionNewResponseAudioTranscriptionJsonResponse struct {
	// The transcribed text
	Text string                                                          `json:"text,required"`
	JSON audioTranscriptionNewResponseAudioTranscriptionJsonResponseJSON `json:"-"`
}

// audioTranscriptionNewResponseAudioTranscriptionJsonResponseJSON contains the
// JSON metadata for the struct
// [AudioTranscriptionNewResponseAudioTranscriptionJsonResponse]
type audioTranscriptionNewResponseAudioTranscriptionJsonResponseJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseAudioTranscriptionJsonResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseAudioTranscriptionJsonResponseJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranscriptionNewResponseAudioTranscriptionJsonResponse) implementsAudioTranscriptionNewResponse() {
}

type AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse struct {
	// The duration of the audio in seconds
	Duration float64 `json:"duration,required"`
	// The language of the audio
	Language string `json:"language,required"`
	// Array of transcription segments
	Segments []AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegment `json:"segments,required"`
	// The task performed
	Task AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseTask `json:"task,required"`
	// The transcribed text
	Text string `json:"text,required"`
	// Array of transcription words (only when timestamp_granularities includes 'word')
	Words []AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord `json:"words"`
	JSON  audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseJSON   `json:"-"`
}

// audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseJSON contains
// the JSON metadata for the struct
// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse]
type audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseJSON struct {
	Duration    apijson.Field
	Language    apijson.Field
	Segments    apijson.Field
	Task        apijson.Field
	Text        apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponse) implementsAudioTranscriptionNewResponse() {
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
	// Array of token IDs for the segment
	Tokens []int64                                                                       `json:"tokens,required"`
	JSON   audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegmentJSON `json:"-"`
}

// audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegmentJSON
// contains the JSON metadata for the struct
// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegment]
type audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegmentJSON struct {
	ID          apijson.Field
	End         apijson.Field
	Start       apijson.Field
	Text        apijson.Field
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseSegmentJSON) RawJSON() string {
	return r.raw
}

// The task performed
type AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseTask string

const (
	AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseTaskTranscribe AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseTask = "transcribe"
	AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseTaskTranslate  AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseTask = "translate"
)

func (r AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseTask) IsKnown() bool {
	switch r {
	case AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseTaskTranscribe, AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseTaskTranslate:
		return true
	}
	return false
}

type AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord struct {
	// End time of the word in seconds
	End float64 `json:"end,required"`
	// Start time of the word in seconds
	Start float64 `json:"start,required"`
	// The word
	Word string                                                                     `json:"word,required"`
	JSON audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWordJSON `json:"-"`
}

// audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWordJSON
// contains the JSON metadata for the struct
// [AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord]
type audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWordJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseAudioTranscriptionVerboseJsonResponseWordJSON) RawJSON() string {
	return r.raw
}

// The task performed
type AudioTranscriptionNewResponseTask string

const (
	AudioTranscriptionNewResponseTaskTranscribe AudioTranscriptionNewResponseTask = "transcribe"
	AudioTranscriptionNewResponseTaskTranslate  AudioTranscriptionNewResponseTask = "translate"
)

func (r AudioTranscriptionNewResponseTask) IsKnown() bool {
	switch r {
	case AudioTranscriptionNewResponseTaskTranscribe, AudioTranscriptionNewResponseTaskTranslate:
		return true
	}
	return false
}

type AudioTranscriptionNewParams struct {
	// Audio file to transcribe
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Optional ISO 639-1 language code. If `auto` is provided, language is
	// auto-detected.
	Language param.Field[string] `json:"language"`
	// Model to use for transcription
	Model param.Field[AudioTranscriptionNewParamsModel] `json:"model"`
	// Optional text to bias decoding.
	Prompt param.Field[string] `json:"prompt"`
	// The format of the response
	ResponseFormat param.Field[AudioTranscriptionNewParamsResponseFormat] `json:"response_format"`
	// Sampling temperature between 0.0 and 1.0
	Temperature param.Field[float64] `json:"temperature"`
	// Controls level of timestamp detail in verbose_json. Only used when
	// response_format is verbose_json.
	TimestampGranularities param.Field[AudioTranscriptionNewParamsTimestampGranularities] `json:"timestamp_granularities"`
}

func (r AudioTranscriptionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// Model to use for transcription
type AudioTranscriptionNewParamsModel string

const (
	AudioTranscriptionNewParamsModelOpenAIWhisperLargeV3 AudioTranscriptionNewParamsModel = "openai/whisper-large-v3"
)

func (r AudioTranscriptionNewParamsModel) IsKnown() bool {
	switch r {
	case AudioTranscriptionNewParamsModelOpenAIWhisperLargeV3:
		return true
	}
	return false
}

// The format of the response
type AudioTranscriptionNewParamsResponseFormat string

const (
	AudioTranscriptionNewParamsResponseFormatJson        AudioTranscriptionNewParamsResponseFormat = "json"
	AudioTranscriptionNewParamsResponseFormatVerboseJson AudioTranscriptionNewParamsResponseFormat = "verbose_json"
)

func (r AudioTranscriptionNewParamsResponseFormat) IsKnown() bool {
	switch r {
	case AudioTranscriptionNewParamsResponseFormatJson, AudioTranscriptionNewParamsResponseFormatVerboseJson:
		return true
	}
	return false
}

// Controls level of timestamp detail in verbose_json. Only used when
// response_format is verbose_json.
type AudioTranscriptionNewParamsTimestampGranularities string

const (
	AudioTranscriptionNewParamsTimestampGranularitiesSegment AudioTranscriptionNewParamsTimestampGranularities = "segment"
	AudioTranscriptionNewParamsTimestampGranularitiesWord    AudioTranscriptionNewParamsTimestampGranularities = "word"
)

func (r AudioTranscriptionNewParamsTimestampGranularities) IsKnown() bool {
	switch r {
	case AudioTranscriptionNewParamsTimestampGranularitiesSegment, AudioTranscriptionNewParamsTimestampGranularitiesWord:
		return true
	}
	return false
}
