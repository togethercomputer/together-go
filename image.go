// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
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
func NewImageService(opts ...option.RequestOption) (r ImageService) {
	r = ImageService{}
	r.Options = opts
	return
}

// Use an image model to generate an image for a given prompt.
func (r *ImageService) Generate(ctx context.Context, body ImageGenerateParams, opts ...option.RequestOption) (res *ImageFile, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "images/generations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ImageDataB64 struct {
	B64Json string `json:"b64_json,required"`
	Index   int64  `json:"index,required"`
	// Any of "b64_json".
	Type ImageDataB64Type `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		B64Json     respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageDataB64) RawJSON() string { return r.JSON.raw }
func (r *ImageDataB64) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageDataB64Type string

const (
	ImageDataB64TypeB64Json ImageDataB64Type = "b64_json"
)

type ImageDataURL struct {
	Index int64 `json:"index,required"`
	// Any of "url".
	Type ImageDataURLType `json:"type,required"`
	URL  string           `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageDataURL) RawJSON() string { return r.JSON.raw }
func (r *ImageDataURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageDataURLType string

const (
	ImageDataURLTypeURL ImageDataURLType = "url"
)

type ImageFile struct {
	ID    string               `json:"id,required"`
	Data  []ImageFileDataUnion `json:"data,required"`
	Model string               `json:"model,required"`
	// Any of "list".
	Object ImageFileObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Data        respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageFile) RawJSON() string { return r.JSON.raw }
func (r *ImageFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ImageFileDataUnion contains all possible properties and values from
// [ImageDataB64], [ImageDataURL].
//
// Use the [ImageFileDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ImageFileDataUnion struct {
	// This field is from variant [ImageDataB64].
	B64Json string `json:"b64_json"`
	Index   int64  `json:"index"`
	// Any of "b64_json", "url".
	Type string `json:"type"`
	// This field is from variant [ImageDataURL].
	URL  string `json:"url"`
	JSON struct {
		B64Json respjson.Field
		Index   respjson.Field
		Type    respjson.Field
		URL     respjson.Field
		raw     string
	} `json:"-"`
}

// anyImageFileData is implemented by each variant of [ImageFileDataUnion] to add
// type safety for the return type of [ImageFileDataUnion.AsAny]
type anyImageFileData interface {
	implImageFileDataUnion()
}

func (ImageDataB64) implImageFileDataUnion() {}
func (ImageDataURL) implImageFileDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ImageFileDataUnion.AsAny().(type) {
//	case together.ImageDataB64:
//	case together.ImageDataURL:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ImageFileDataUnion) AsAny() anyImageFileData {
	switch u.Type {
	case "b64_json":
		return u.AsB64Json()
	case "url":
		return u.AsURL()
	}
	return nil
}

func (u ImageFileDataUnion) AsB64Json() (v ImageDataB64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ImageFileDataUnion) AsURL() (v ImageDataURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ImageFileDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ImageFileDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageFileObject string

const (
	ImageFileObjectList ImageFileObject = "list"
)

type ImageGenerateParams struct {
	// The model to use for image generation.
	//
	// [See all of Together AI's image models](https://docs.together.ai/docs/serverless-models#image-models)
	Model ImageGenerateParamsModel `json:"model,omitzero,required"`
	// A description of the desired images. Maximum length varies by model.
	Prompt string `json:"prompt,required"`
	// If true, disables the safety checker for image generation.
	DisableSafetyChecker param.Opt[bool] `json:"disable_safety_checker,omitzero"`
	// Adjusts the alignment of the generated image with the input prompt. Higher
	// values (e.g., 8-10) make the output more faithful to the prompt, while lower
	// values (e.g., 1-5) encourage more creative freedom.
	GuidanceScale param.Opt[float64] `json:"guidance_scale,omitzero"`
	// Height of the image to generate in number of pixels.
	Height param.Opt[int64] `json:"height,omitzero"`
	// URL of an image to use for image models that support it.
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Number of image results to generate.
	N param.Opt[int64] `json:"n,omitzero"`
	// The prompt or prompts not to guide the image generation.
	NegativePrompt param.Opt[string] `json:"negative_prompt,omitzero"`
	// Seed used for generation. Can be used to reproduce image generations.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Number of generation steps.
	Steps param.Opt[int64] `json:"steps,omitzero"`
	// Width of the image to generate in number of pixels.
	Width param.Opt[int64] `json:"width,omitzero"`
	// An array of objects that define LoRAs (Low-Rank Adaptations) to influence the
	// generated image.
	ImageLoras []ImageGenerateParamsImageLora `json:"image_loras,omitzero"`
	// The format of the image response. Can be either be `jpeg` or `png`. Defaults to
	// `jpeg`.
	//
	// Any of "jpeg", "png".
	OutputFormat ImageGenerateParamsOutputFormat `json:"output_format,omitzero"`
	// Format of the image response. Can be either a base64 string or a URL.
	//
	// Any of "base64", "url".
	ResponseFormat ImageGenerateParamsResponseFormat `json:"response_format,omitzero"`
	paramObj
}

func (r ImageGenerateParams) MarshalJSON() (data []byte, err error) {
	type shadow ImageGenerateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImageGenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The model to use for image generation.
//
// [See all of Together AI's image models](https://docs.together.ai/docs/serverless-models#image-models)
type ImageGenerateParamsModel string

const (
	ImageGenerateParamsModelBlackForestLabsFlux1SchnellFree ImageGenerateParamsModel = "black-forest-labs/FLUX.1-schnell-Free"
	ImageGenerateParamsModelBlackForestLabsFlux1Schnell     ImageGenerateParamsModel = "black-forest-labs/FLUX.1-schnell"
	ImageGenerateParamsModelBlackForestLabsFlux1_1Pro       ImageGenerateParamsModel = "black-forest-labs/FLUX.1.1-pro"
)

// The properties Path, Scale are required.
type ImageGenerateParamsImageLora struct {
	// The URL of the LoRA to apply (e.g.
	// https://huggingface.co/strangerzonehf/Flux-Midjourney-Mix2-LoRA).
	Path string `json:"path,required"`
	// The strength of the LoRA's influence. Most LoRA's recommend a value of 1.
	Scale float64 `json:"scale,required"`
	paramObj
}

func (r ImageGenerateParamsImageLora) MarshalJSON() (data []byte, err error) {
	type shadow ImageGenerateParamsImageLora
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImageGenerateParamsImageLora) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The format of the image response. Can be either be `jpeg` or `png`. Defaults to
// `jpeg`.
type ImageGenerateParamsOutputFormat string

const (
	ImageGenerateParamsOutputFormatJpeg ImageGenerateParamsOutputFormat = "jpeg"
	ImageGenerateParamsOutputFormatPng  ImageGenerateParamsOutputFormat = "png"
)

// Format of the image response. Can be either a base64 string or a URL.
type ImageGenerateParamsResponseFormat string

const (
	ImageGenerateParamsResponseFormatBase64 ImageGenerateParamsResponseFormat = "base64"
	ImageGenerateParamsResponseFormatURL    ImageGenerateParamsResponseFormat = "url"
)
