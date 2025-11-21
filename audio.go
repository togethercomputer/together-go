// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// AudioService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioService] method instead.
type AudioService struct {
	Options        []option.RequestOption
	Speech         AudioSpeechService
	Voices         AudioVoiceService
	Transcriptions AudioTranscriptionService
	Translations   AudioTranslationService
}

// NewAudioService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAudioService(opts ...option.RequestOption) (r AudioService) {
	r = AudioService{}
	r.Options = opts
	r.Speech = NewAudioSpeechService(opts...)
	r.Voices = NewAudioVoiceService(opts...)
	r.Transcriptions = NewAudioTranscriptionService(opts...)
	r.Translations = NewAudioTranslationService(opts...)
	return
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
