// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// AudioVoiceService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioVoiceService] method instead.
type AudioVoiceService struct {
	Options []option.RequestOption
}

// NewAudioVoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudioVoiceService(opts ...option.RequestOption) (r AudioVoiceService) {
	r = AudioVoiceService{}
	r.Options = opts
	return
}

// Fetch available voices for each model
func (r *AudioVoiceService) List(ctx context.Context, opts ...option.RequestOption) (res *AudioVoiceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response containing a list of models and their available voices.
type AudioVoiceListResponse struct {
	Data []AudioVoiceListResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioVoiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *AudioVoiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a model with its available voices.
type AudioVoiceListResponseData struct {
	Model  string                            `json:"model,required"`
	Voices []AudioVoiceListResponseDataVoice `json:"voices,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Model       respjson.Field
		Voices      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioVoiceListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AudioVoiceListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioVoiceListResponseDataVoice struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioVoiceListResponseDataVoice) RawJSON() string { return r.JSON.raw }
func (r *AudioVoiceListResponseDataVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
