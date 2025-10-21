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
)

// ModelService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	return
}

// Lists all of Together's open-source models
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ModelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a custom model or adapter from Hugging Face or S3
func (r *ModelService) Upload(ctx context.Context, body ModelUploadParams, opts ...option.RequestOption) (res *ModelUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ModelListResponse struct {
	ID      string `json:"id,required"`
	Created int64  `json:"created,required"`
	Object  string `json:"object,required"`
	// Any of "chat", "language", "code", "image", "embedding", "moderation", "rerank".
	Type          ModelListResponseType    `json:"type,required"`
	ContextLength int64                    `json:"context_length"`
	DisplayName   string                   `json:"display_name"`
	License       string                   `json:"license"`
	Link          string                   `json:"link"`
	Organization  string                   `json:"organization"`
	Pricing       ModelListResponsePricing `json:"pricing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Created       respjson.Field
		Object        respjson.Field
		Type          respjson.Field
		ContextLength respjson.Field
		DisplayName   respjson.Field
		License       respjson.Field
		Link          respjson.Field
		Organization  respjson.Field
		Pricing       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListResponseType string

const (
	ModelListResponseTypeChat       ModelListResponseType = "chat"
	ModelListResponseTypeLanguage   ModelListResponseType = "language"
	ModelListResponseTypeCode       ModelListResponseType = "code"
	ModelListResponseTypeImage      ModelListResponseType = "image"
	ModelListResponseTypeEmbedding  ModelListResponseType = "embedding"
	ModelListResponseTypeModeration ModelListResponseType = "moderation"
	ModelListResponseTypeRerank     ModelListResponseType = "rerank"
)

type ModelListResponsePricing struct {
	Base     float64 `json:"base,required"`
	Finetune float64 `json:"finetune,required"`
	Hourly   float64 `json:"hourly,required"`
	Input    float64 `json:"input,required"`
	Output   float64 `json:"output,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base        respjson.Field
		Finetune    respjson.Field
		Hourly      respjson.Field
		Input       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponsePricing) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponsePricing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelUploadResponse struct {
	Data    ModelUploadResponseData `json:"data,required"`
	Message string                  `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelUploadResponseData struct {
	JobID       string `json:"job_id,required"`
	ModelID     string `json:"model_id,required"`
	ModelName   string `json:"model_name,required"`
	ModelSource string `json:"model_source,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		ModelID     respjson.Field
		ModelName   respjson.Field
		ModelSource respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelUploadResponseData) RawJSON() string { return r.JSON.raw }
func (r *ModelUploadResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelUploadParams struct {
	// The name to give to your uploaded model
	ModelName string `json:"model_name,required"`
	// The source location of the model (Hugging Face repo or S3 path)
	ModelSource string `json:"model_source,required"`
	// The base model to use for an adapter if setting it to run against a serverless
	// pool. Only used for model_type `adapter`.
	BaseModel param.Opt[string] `json:"base_model,omitzero"`
	// A description of your model
	Description param.Opt[string] `json:"description,omitzero"`
	// Hugging Face token (if uploading from Hugging Face)
	HfToken param.Opt[string] `json:"hf_token,omitzero"`
	// The lora pool to use for an adapter if setting it to run against, say, a
	// dedicated pool. Only used for model_type `adapter`.
	LoraModel param.Opt[string] `json:"lora_model,omitzero"`
	// Whether the model is a full model or an adapter
	//
	// Any of "model", "adapter".
	ModelType ModelUploadParamsModelType `json:"model_type,omitzero"`
	paramObj
}

func (r ModelUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow ModelUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the model is a full model or an adapter
type ModelUploadParamsModelType string

const (
	ModelUploadParamsModelTypeModel   ModelUploadParamsModelType = "model"
	ModelUploadParamsModelTypeAdapter ModelUploadParamsModelType = "adapter"
)
