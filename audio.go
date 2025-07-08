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

// AudioService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioService] method instead.
type AudioService struct {
	Options        []option.RequestOption
	Transcriptions *AudioTranscriptionService
	Translations   *AudioTranslationService
}

// NewAudioService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAudioService(opts ...option.RequestOption) (r *AudioService) {
	r = &AudioService{}
	r.Options = opts
	r.Transcriptions = NewAudioTranscriptionService(opts...)
	r.Translations = NewAudioTranslationService(opts...)
	return
}

// Generate audio from input text
func (r *AudioService) New(ctx context.Context, body AudioNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream"), option.WithJSONSet("stream", true)}, opts...)
	path := "audio/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[AudioSpeechStreamChunk](ssestream.NewDecoder(raw), err)
}

type AudioSpeechStreamChunk struct {
	// base64 encoded audio stream
	B64    string                       `json:"b64,required"`
	Model  string                       `json:"model,required"`
	Object AudioSpeechStreamChunkObject `json:"object,required"`
	JSON   audioSpeechStreamChunkJSON   `json:"-"`
}

// audioSpeechStreamChunkJSON contains the JSON metadata for the struct
// [AudioSpeechStreamChunk]
type audioSpeechStreamChunkJSON struct {
	B64         apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioSpeechStreamChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioSpeechStreamChunkJSON) RawJSON() string {
	return r.raw
}

type AudioSpeechStreamChunkObject string

const (
	AudioSpeechStreamChunkObjectAudioTtsChunk AudioSpeechStreamChunkObject = "audio.tts.chunk"
)

func (r AudioSpeechStreamChunkObject) IsKnown() bool {
	switch r {
	case AudioSpeechStreamChunkObjectAudioTtsChunk:
		return true
	}
	return false
}

type AudioNewParams struct {
	// Input text to generate the audio for
	Input param.Field[string] `json:"input,required"`
	// The name of the model to query.
	//
	// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#audio-models)
	Model param.Field[AudioNewParamsModel] `json:"model,required"`
	// The voice to use for generating the audio.
	// [View all supported voices here](https://docs.together.ai/docs/text-to-speech#voices-available).
	Voice param.Field[AudioNewParamsVoice] `json:"voice,required"`
	// Language of input text
	Language param.Field[AudioNewParamsLanguage] `json:"language"`
	// Audio encoding of response
	ResponseEncoding param.Field[AudioNewParamsResponseEncoding] `json:"response_encoding"`
	// The format of audio output
	ResponseFormat param.Field[AudioNewParamsResponseFormat] `json:"response_format"`
	// Sampling rate to use for the output audio
	SampleRate param.Field[float64] `json:"sample_rate"`
}

func (r AudioNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of the model to query.
//
// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#audio-models)
type AudioNewParamsModel string

const (
	AudioNewParamsModelCartesiaSonic AudioNewParamsModel = "cartesia/sonic"
)

func (r AudioNewParamsModel) IsKnown() bool {
	switch r {
	case AudioNewParamsModelCartesiaSonic:
		return true
	}
	return false
}

// The voice to use for generating the audio.
// [View all supported voices here](https://docs.together.ai/docs/text-to-speech#voices-available).
type AudioNewParamsVoice string

const (
	AudioNewParamsVoiceLaidbackWoman    AudioNewParamsVoice = "laidback woman"
	AudioNewParamsVoicePoliteMan        AudioNewParamsVoice = "polite man"
	AudioNewParamsVoiceStorytellerLady  AudioNewParamsVoice = "storyteller lady"
	AudioNewParamsVoiceFriendlySidekick AudioNewParamsVoice = "friendly sidekick"
)

func (r AudioNewParamsVoice) IsKnown() bool {
	switch r {
	case AudioNewParamsVoiceLaidbackWoman, AudioNewParamsVoicePoliteMan, AudioNewParamsVoiceStorytellerLady, AudioNewParamsVoiceFriendlySidekick:
		return true
	}
	return false
}

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

func (r AudioNewParamsLanguage) IsKnown() bool {
	switch r {
	case AudioNewParamsLanguageEn, AudioNewParamsLanguageDe, AudioNewParamsLanguageFr, AudioNewParamsLanguageEs, AudioNewParamsLanguageHi, AudioNewParamsLanguageIt, AudioNewParamsLanguageJa, AudioNewParamsLanguageKo, AudioNewParamsLanguageNl, AudioNewParamsLanguagePl, AudioNewParamsLanguagePt, AudioNewParamsLanguageRu, AudioNewParamsLanguageSv, AudioNewParamsLanguageTr, AudioNewParamsLanguageZh:
		return true
	}
	return false
}

// Audio encoding of response
type AudioNewParamsResponseEncoding string

const (
	AudioNewParamsResponseEncodingPcmF32le AudioNewParamsResponseEncoding = "pcm_f32le"
	AudioNewParamsResponseEncodingPcmS16le AudioNewParamsResponseEncoding = "pcm_s16le"
	AudioNewParamsResponseEncodingPcmMulaw AudioNewParamsResponseEncoding = "pcm_mulaw"
	AudioNewParamsResponseEncodingPcmAlaw  AudioNewParamsResponseEncoding = "pcm_alaw"
)

func (r AudioNewParamsResponseEncoding) IsKnown() bool {
	switch r {
	case AudioNewParamsResponseEncodingPcmF32le, AudioNewParamsResponseEncodingPcmS16le, AudioNewParamsResponseEncodingPcmMulaw, AudioNewParamsResponseEncodingPcmAlaw:
		return true
	}
	return false
}

// The format of audio output
type AudioNewParamsResponseFormat string

const (
	AudioNewParamsResponseFormatMP3 AudioNewParamsResponseFormat = "mp3"
	AudioNewParamsResponseFormatWav AudioNewParamsResponseFormat = "wav"
	AudioNewParamsResponseFormatRaw AudioNewParamsResponseFormat = "raw"
)

func (r AudioNewParamsResponseFormat) IsKnown() bool {
	switch r {
	case AudioNewParamsResponseFormatMP3, AudioNewParamsResponseFormatWav, AudioNewParamsResponseFormatRaw:
		return true
	}
	return false
}
