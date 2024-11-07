// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
)

// ImageService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageService] method instead.
type ImageService struct {
	Options []option.RequestOption
}

// NewImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageService(opts ...option.RequestOption) (r *ImageService) {
	r = &ImageService{}
	r.Options = opts
	return
}

// Use an image model to generate an image for a given prompt.
func (r *ImageService) New(ctx context.Context, body ImageNewParams, opts ...option.RequestOption) (res *ImageFile, err error) {
	opts = append(r.Options[:], opts...)
	path := "images/generations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ImageFile struct {
	ID     string          `json:"id,required"`
	Data   []ImageFileData `json:"data,required"`
	Model  string          `json:"model,required"`
	Object ImageFileObject `json:"object,required"`
	JSON   imageFileJSON   `json:"-"`
}

// imageFileJSON contains the JSON metadata for the struct [ImageFile]
type imageFileJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageFileJSON) RawJSON() string {
	return r.raw
}

type ImageFileData struct {
	B64Json string            `json:"b64_json,required"`
	Index   int64             `json:"index,required"`
	JSON    imageFileDataJSON `json:"-"`
}

// imageFileDataJSON contains the JSON metadata for the struct [ImageFileData]
type imageFileDataJSON struct {
	B64Json     apijson.Field
	Index       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageFileData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageFileDataJSON) RawJSON() string {
	return r.raw
}

type ImageFileObject string

const (
	ImageFileObjectList ImageFileObject = "list"
)

func (r ImageFileObject) IsKnown() bool {
	switch r {
	case ImageFileObjectList:
		return true
	}
	return false
}

type ImageNewParams struct {
	// The model to use for image generation.
	//
	// [See all of Together AI's image models](https://docs.together.ai/docs/serverless-models#image-models)
	Model param.Field[ImageNewParamsModel] `json:"model,required"`
	// A description of the desired images. Maximum length varies by model.
	Prompt param.Field[string] `json:"prompt,required"`
	// Height of the image to generate in number of pixels.
	Height param.Field[int64] `json:"height"`
	// Number of image results to generate.
	N param.Field[int64] `json:"n"`
	// The prompt or prompts not to guide the image generation.
	NegativePrompt param.Field[string] `json:"negative_prompt"`
	// Seed used for generation. Can be used to reproduce image generations.
	Seed param.Field[int64] `json:"seed"`
	// Number of generation steps.
	Steps param.Field[int64] `json:"steps"`
	// Width of the image to generate in number of pixels.
	Width param.Field[int64] `json:"width"`
}

func (r ImageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The model to use for image generation.
//
// [See all of Together AI's image models](https://docs.together.ai/docs/serverless-models#image-models)
type ImageNewParamsModel string

const (
	ImageNewParamsModelBlackForestLabsFlux1SchnellFree ImageNewParamsModel = "black-forest-labs/FLUX.1-schnell-Free"
	ImageNewParamsModelBlackForestLabsFlux1Schnell     ImageNewParamsModel = "black-forest-labs/FLUX.1-schnell"
	ImageNewParamsModelBlackForestLabsFlux1_1Pro       ImageNewParamsModel = "black-forest-labs/FLUX.1.1-pro"
)

func (r ImageNewParamsModel) IsKnown() bool {
	switch r {
	case ImageNewParamsModelBlackForestLabsFlux1SchnellFree, ImageNewParamsModelBlackForestLabsFlux1Schnell, ImageNewParamsModelBlackForestLabsFlux1_1Pro:
		return true
	}
	return false
}
