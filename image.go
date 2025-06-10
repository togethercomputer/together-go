// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"reflect"

	"github.com/tidwall/gjson"
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

type ImageDataB64 struct {
	B64Json string           `json:"b64_json,required"`
	Index   int64            `json:"index,required"`
	JSON    imageDataB64JSON `json:"-"`
}

// imageDataB64JSON contains the JSON metadata for the struct [ImageDataB64]
type imageDataB64JSON struct {
	B64Json     apijson.Field
	Index       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageDataB64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageDataB64JSON) RawJSON() string {
	return r.raw
}

func (r ImageDataB64) implementsImageFileData() {}

type ImageDataURL struct {
	Index int64            `json:"index,required"`
	URL   string           `json:"url,required"`
	JSON  imageDataURLJSON `json:"-"`
}

// imageDataURLJSON contains the JSON metadata for the struct [ImageDataURL]
type imageDataURLJSON struct {
	Index       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageDataURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageDataURLJSON) RawJSON() string {
	return r.raw
}

func (r ImageDataURL) implementsImageFileData() {}

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
	Index   int64             `json:"index,required"`
	B64Json string            `json:"b64_json"`
	URL     string            `json:"url"`
	JSON    imageFileDataJSON `json:"-"`
	union   ImageFileDataUnion
}

// imageFileDataJSON contains the JSON metadata for the struct [ImageFileData]
type imageFileDataJSON struct {
	Index       apijson.Field
	B64Json     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r imageFileDataJSON) RawJSON() string {
	return r.raw
}

func (r *ImageFileData) UnmarshalJSON(data []byte) (err error) {
	*r = ImageFileData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ImageFileDataUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ImageDataB64], [ImageDataURL].
func (r ImageFileData) AsUnion() ImageFileDataUnion {
	return r.union
}

// Union satisfied by [ImageDataB64] or [ImageDataURL].
type ImageFileDataUnion interface {
	implementsImageFileData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageFileDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ImageDataB64{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ImageDataURL{}),
		},
	)
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
	// Adjusts the alignment of the generated image with the input prompt. Higher
	// values (e.g., 8-10) make the output more faithful to the prompt, while lower
	// values (e.g., 1-5) encourage more creative freedom.
	Guidance param.Field[float64] `json:"guidance"`
	// Height of the image to generate in number of pixels.
	Height param.Field[int64] `json:"height"`
	// An array of objects that define LoRAs (Low-Rank Adaptations) to influence the
	// generated image.
	ImageLoras param.Field[[]ImageNewParamsImageLora] `json:"image_loras"`
	// URL of an image to use for image models that support it.
	ImageURL param.Field[string] `json:"image_url"`
	// Number of image results to generate.
	N param.Field[int64] `json:"n"`
	// The prompt or prompts not to guide the image generation.
	NegativePrompt param.Field[string] `json:"negative_prompt"`
	// The format of the image response. Can be either be `jpeg` or `png`. Defaults to
	// `jpeg`.
	OutputFormat param.Field[ImageNewParamsOutputFormat] `json:"output_format"`
	// Format of the image response. Can be either a base64 string or a URL.
	ResponseFormat param.Field[ImageNewParamsResponseFormat] `json:"response_format"`
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

type ImageNewParamsImageLora struct {
	// The URL of the LoRA to apply (e.g.
	// https://huggingface.co/strangerzonehf/Flux-Midjourney-Mix2-LoRA).
	Path param.Field[string] `json:"path,required"`
	// The strength of the LoRA's influence. Most LoRA's recommend a value of 1.
	Scale param.Field[float64] `json:"scale,required"`
}

func (r ImageNewParamsImageLora) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the image response. Can be either be `jpeg` or `png`. Defaults to
// `jpeg`.
type ImageNewParamsOutputFormat string

const (
	ImageNewParamsOutputFormatJpeg ImageNewParamsOutputFormat = "jpeg"
	ImageNewParamsOutputFormatPng  ImageNewParamsOutputFormat = "png"
)

func (r ImageNewParamsOutputFormat) IsKnown() bool {
	switch r {
	case ImageNewParamsOutputFormatJpeg, ImageNewParamsOutputFormatPng:
		return true
	}
	return false
}

// Format of the image response. Can be either a base64 string or a URL.
type ImageNewParamsResponseFormat string

const (
	ImageNewParamsResponseFormatBase64 ImageNewParamsResponseFormat = "base64"
	ImageNewParamsResponseFormatURL    ImageNewParamsResponseFormat = "url"
)

func (r ImageNewParamsResponseFormat) IsKnown() bool {
	switch r {
	case ImageNewParamsResponseFormatBase64, ImageNewParamsResponseFormatURL:
		return true
	}
	return false
}
