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
	"github.com/togethercomputer/together-go/packages/ssestream"
)

// AudioSpeechService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioSpeechService] method instead.
type AudioSpeechService struct {
	Options []option.RequestOption
}

// NewAudioSpeechService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudioSpeechService(opts ...option.RequestOption) (r AudioSpeechService) {
	r = AudioSpeechService{}
	r.Options = opts
	return
}

// Generate audio from input text
func (r *AudioSpeechService) New(ctx context.Context, body AudioSpeechNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "audio/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate audio from input text
func (r *AudioSpeechService) NewStreaming(ctx context.Context, body AudioSpeechNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[AudioSpeechStreamChunk]) {
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

type AudioSpeechNewParams struct {
	// Input text to generate the audio for
	Input string `json:"input,required"`
	// The name of the model to query.
	//
	// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#audio-models)
	// The current supported tts models are: - cartesia/sonic - hexgrad/Kokoro-82M -
	// canopylabs/orpheus-3b-0.1-ft
	Model AudioSpeechNewParamsModel `json:"model,omitzero,required"`
	// The voice to use for generating the audio. The voices supported are different
	// for each model. For eg - for canopylabs/orpheus-3b-0.1-ft, one of the voices
	// supported is tara, for hexgrad/Kokoro-82M, one of the voices supported is
	// af_alloy and for cartesia/sonic, one of the voices supported is "friendly
	// sidekick".
	//
	// You can view the voices supported for each model using the /v1/voices endpoint
	// sending the model name as the query parameter.
	// [View all supported voices here](https://docs.together.ai/docs/text-to-speech#voices-available).
	Voice string `json:"voice,required"`
	// Sampling rate to use for the output audio. The default sampling rate for
	// canopylabs/orpheus-3b-0.1-ft and hexgrad/Kokoro-82M is 24000 and for
	// cartesia/sonic is 44100.
	SampleRate param.Opt[float64] `json:"sample_rate,omitzero"`
	// Language of input text.
	//
	// Any of "en", "de", "fr", "es", "hi", "it", "ja", "ko", "nl", "pl", "pt", "ru",
	// "sv", "tr", "zh".
	Language AudioSpeechNewParamsLanguage `json:"language,omitzero"`
	// Audio encoding of response
	//
	// Any of "pcm_f32le", "pcm_s16le", "pcm_mulaw", "pcm_alaw".
	ResponseEncoding AudioSpeechNewParamsResponseEncoding `json:"response_encoding,omitzero"`
	// The format of audio output. Supported formats are mp3, wav, raw if streaming is
	// false. If streaming is true, the only supported format is raw.
	//
	// Any of "mp3", "wav", "raw".
	ResponseFormat AudioSpeechNewParamsResponseFormat `json:"response_format,omitzero"`
	paramObj
}

func (r AudioSpeechNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AudioSpeechNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudioSpeechNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The name of the model to query.
//
// [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#audio-models)
// The current supported tts models are: - cartesia/sonic - hexgrad/Kokoro-82M -
// canopylabs/orpheus-3b-0.1-ft
type AudioSpeechNewParamsModel string

const (
	AudioSpeechNewParamsModelCartesiaSonic            AudioSpeechNewParamsModel = "cartesia/sonic"
	AudioSpeechNewParamsModelHexgradKokoro82M         AudioSpeechNewParamsModel = "hexgrad/Kokoro-82M"
	AudioSpeechNewParamsModelCanopylabsOrpheus3b0_1Ft AudioSpeechNewParamsModel = "canopylabs/orpheus-3b-0.1-ft"
)

// Language of input text.
type AudioSpeechNewParamsLanguage string

const (
	AudioSpeechNewParamsLanguageEn AudioSpeechNewParamsLanguage = "en"
	AudioSpeechNewParamsLanguageDe AudioSpeechNewParamsLanguage = "de"
	AudioSpeechNewParamsLanguageFr AudioSpeechNewParamsLanguage = "fr"
	AudioSpeechNewParamsLanguageEs AudioSpeechNewParamsLanguage = "es"
	AudioSpeechNewParamsLanguageHi AudioSpeechNewParamsLanguage = "hi"
	AudioSpeechNewParamsLanguageIt AudioSpeechNewParamsLanguage = "it"
	AudioSpeechNewParamsLanguageJa AudioSpeechNewParamsLanguage = "ja"
	AudioSpeechNewParamsLanguageKo AudioSpeechNewParamsLanguage = "ko"
	AudioSpeechNewParamsLanguageNl AudioSpeechNewParamsLanguage = "nl"
	AudioSpeechNewParamsLanguagePl AudioSpeechNewParamsLanguage = "pl"
	AudioSpeechNewParamsLanguagePt AudioSpeechNewParamsLanguage = "pt"
	AudioSpeechNewParamsLanguageRu AudioSpeechNewParamsLanguage = "ru"
	AudioSpeechNewParamsLanguageSv AudioSpeechNewParamsLanguage = "sv"
	AudioSpeechNewParamsLanguageTr AudioSpeechNewParamsLanguage = "tr"
	AudioSpeechNewParamsLanguageZh AudioSpeechNewParamsLanguage = "zh"
)

// Audio encoding of response
type AudioSpeechNewParamsResponseEncoding string

const (
	AudioSpeechNewParamsResponseEncodingPcmF32le AudioSpeechNewParamsResponseEncoding = "pcm_f32le"
	AudioSpeechNewParamsResponseEncodingPcmS16le AudioSpeechNewParamsResponseEncoding = "pcm_s16le"
	AudioSpeechNewParamsResponseEncodingPcmMulaw AudioSpeechNewParamsResponseEncoding = "pcm_mulaw"
	AudioSpeechNewParamsResponseEncodingPcmAlaw  AudioSpeechNewParamsResponseEncoding = "pcm_alaw"
)

// The format of audio output. Supported formats are mp3, wav, raw if streaming is
// false. If streaming is true, the only supported format is raw.
type AudioSpeechNewParamsResponseFormat string

const (
	AudioSpeechNewParamsResponseFormatMP3 AudioSpeechNewParamsResponseFormat = "mp3"
	AudioSpeechNewParamsResponseFormatWav AudioSpeechNewParamsResponseFormat = "wav"
	AudioSpeechNewParamsResponseFormatRaw AudioSpeechNewParamsResponseFormat = "raw"
)
