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

// AudioService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioService] method instead.
type AudioService struct {
	Options        []option.RequestOption
	Transcriptions AudioTranscriptionService
	Translations   AudioTranslationService
}

// NewAudioService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAudioService(opts ...option.RequestOption) (r AudioService) {
	r = AudioService{}
	r.Options = opts
	r.Transcriptions = NewAudioTranscriptionService(opts...)
	r.Translations = NewAudioTranslationService(opts...)
	return
}

// Generate audio from input text
func (r *AudioService) New(ctx context.Context, body AudioNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "audio/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate audio from input text
func (r *AudioService) NewStreaming(ctx context.Context, body AudioNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[AudioSpeechStreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream"), option.WithJSONSet("stream", true)}, opts...)
	path := "audio/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[AudioSpeechStreamChunk](ssestream.NewDecoder(raw), err)
}

type AudioSpeechStreamChunk struct {
	// base64 encoded audio stream
	B64   string `json:"b64,required"`
	Model string `json:"model,required"`
	// Any of "audio.tts.chunk".
	Object AudioSpeechStreamChunkObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		B64         respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioSpeechStreamChunk) RawJSON() string { return r.JSON.raw }
func (r *AudioSpeechStreamChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioSpeechStreamChunkObject string

const (
	AudioSpeechStreamChunkObjectAudioTtsChunk AudioSpeechStreamChunkObject = "audio.tts.chunk"
)

type AudioNewParams struct {
	// Input text to generate the audio for
	Input string `json:"input,required"`
	// The name of the model to query.
	//
	// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#audio-models)
	Model AudioNewParamsModel `json:"model,omitzero,required"`
	// The voice to use for generating the audio.
	// [View all supported voices here](https://docs.together.ai/docs/text-to-speech#voices-available).
	Voice AudioNewParamsVoice `json:"voice,omitzero,required"`
	// Sampling rate to use for the output audio
	SampleRate param.Opt[float64] `json:"sample_rate,omitzero"`
	// Language of input text
	//
	// Any of "en", "de", "fr", "es", "hi", "it", "ja", "ko", "nl", "pl", "pt", "ru",
	// "sv", "tr", "zh".
	Language AudioNewParamsLanguage `json:"language,omitzero"`
	// Audio encoding of response
	//
	// Any of "pcm_f32le", "pcm_s16le", "pcm_mulaw", "pcm_alaw".
	ResponseEncoding AudioNewParamsResponseEncoding `json:"response_encoding,omitzero"`
	// The format of audio output
	//
	// Any of "mp3", "wav", "raw".
	ResponseFormat AudioNewParamsResponseFormat `json:"response_format,omitzero"`
	paramObj
}

func (r AudioNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AudioNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudioNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The name of the model to query.
//
// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#audio-models)
type AudioNewParamsModel string

const (
	AudioNewParamsModelCartesiaSonic AudioNewParamsModel = "cartesia/sonic"
)

// The voice to use for generating the audio.
// [View all supported voices here](https://docs.together.ai/docs/text-to-speech#voices-available).
type AudioNewParamsVoice string

const (
	AudioNewParamsVoiceLaidbackWoman    AudioNewParamsVoice = "laidback woman"
	AudioNewParamsVoicePoliteMan        AudioNewParamsVoice = "polite man"
	AudioNewParamsVoiceStorytellerLady  AudioNewParamsVoice = "storyteller lady"
	AudioNewParamsVoiceFriendlySidekick AudioNewParamsVoice = "friendly sidekick"
)

// Language of input text
type AudioNewParamsLanguage string

const (
	AudioNewParamsLanguageEn AudioNewParamsLanguage = "en"
	AudioNewParamsLanguageDe AudioNewParamsLanguage = "de"
	AudioNewParamsLanguageFr AudioNewParamsLanguage = "fr"
	AudioNewParamsLanguageEs AudioNewParamsLanguage = "es"
	AudioNewParamsLanguageHi AudioNewParamsLanguage = "hi"
	AudioNewParamsLanguageIt AudioNewParamsLanguage = "it"
	AudioNewParamsLanguageJa AudioNewParamsLanguage = "ja"
	AudioNewParamsLanguageKo AudioNewParamsLanguage = "ko"
	AudioNewParamsLanguageNl AudioNewParamsLanguage = "nl"
	AudioNewParamsLanguagePl AudioNewParamsLanguage = "pl"
	AudioNewParamsLanguagePt AudioNewParamsLanguage = "pt"
	AudioNewParamsLanguageRu AudioNewParamsLanguage = "ru"
	AudioNewParamsLanguageSv AudioNewParamsLanguage = "sv"
	AudioNewParamsLanguageTr AudioNewParamsLanguage = "tr"
	AudioNewParamsLanguageZh AudioNewParamsLanguage = "zh"
)

// Audio encoding of response
type AudioNewParamsResponseEncoding string

const (
	AudioNewParamsResponseEncodingPcmF32le AudioNewParamsResponseEncoding = "pcm_f32le"
	AudioNewParamsResponseEncodingPcmS16le AudioNewParamsResponseEncoding = "pcm_s16le"
	AudioNewParamsResponseEncodingPcmMulaw AudioNewParamsResponseEncoding = "pcm_mulaw"
	AudioNewParamsResponseEncodingPcmAlaw  AudioNewParamsResponseEncoding = "pcm_alaw"
)

// The format of audio output
type AudioNewParamsResponseFormat string

const (
	AudioNewParamsResponseFormatMP3 AudioNewParamsResponseFormat = "mp3"
	AudioNewParamsResponseFormatWav AudioNewParamsResponseFormat = "wav"
	AudioNewParamsResponseFormatRaw AudioNewParamsResponseFormat = "raw"
)
