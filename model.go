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
func NewModelService(opts ...option.RequestOption) (r *ModelService) {
	r = &ModelService{}
	r.Options = opts
	return
}

// Lists all of Together's open-source models
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ModelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a custom model from Hugging Face or S3
func (r *ModelService) Upload(ctx context.Context, body ModelUploadParams, opts ...option.RequestOption) (res *ModelUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ModelListResponse struct {
	ID            string                   `json:"id,required"`
	Created       int64                    `json:"created,required"`
	Object        string                   `json:"object,required"`
	Type          ModelListResponseType    `json:"type,required"`
	ContextLength int64                    `json:"context_length"`
	DisplayName   string                   `json:"display_name"`
	License       string                   `json:"license"`
	Link          string                   `json:"link"`
	Organization  string                   `json:"organization"`
	Pricing       ModelListResponsePricing `json:"pricing"`
	JSON          modelListResponseJSON    `json:"-"`
}

// modelListResponseJSON contains the JSON metadata for the struct
// [ModelListResponse]
type modelListResponseJSON struct {
	ID            apijson.Field
	Created       apijson.Field
	Object        apijson.Field
	Type          apijson.Field
	ContextLength apijson.Field
	DisplayName   apijson.Field
	License       apijson.Field
	Link          apijson.Field
	Organization  apijson.Field
	Pricing       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ModelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelListResponseJSON) RawJSON() string {
	return r.raw
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

func (r ModelListResponseType) IsKnown() bool {
	switch r {
	case ModelListResponseTypeChat, ModelListResponseTypeLanguage, ModelListResponseTypeCode, ModelListResponseTypeImage, ModelListResponseTypeEmbedding, ModelListResponseTypeModeration, ModelListResponseTypeRerank:
		return true
	}
	return false
}

type ModelListResponsePricing struct {
	Base     float64                      `json:"base,required"`
	Finetune float64                      `json:"finetune,required"`
	Hourly   float64                      `json:"hourly,required"`
	Input    float64                      `json:"input,required"`
	Output   float64                      `json:"output,required"`
	JSON     modelListResponsePricingJSON `json:"-"`
}

// modelListResponsePricingJSON contains the JSON metadata for the struct
// [ModelListResponsePricing]
type modelListResponsePricingJSON struct {
	Base        apijson.Field
	Finetune    apijson.Field
	Hourly      apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelListResponsePricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelListResponsePricingJSON) RawJSON() string {
	return r.raw
}

type ModelUploadResponse struct {
	Data    ModelUploadResponseData `json:"data,required"`
	Message string                  `json:"message,required"`
	JSON    modelUploadResponseJSON `json:"-"`
}

// modelUploadResponseJSON contains the JSON metadata for the struct
// [ModelUploadResponse]
type modelUploadResponseJSON struct {
	Data        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelUploadResponseJSON) RawJSON() string {
	return r.raw
}

type ModelUploadResponseData struct {
	JobID       string                      `json:"job_id,required"`
	ModelID     string                      `json:"model_id,required"`
	ModelName   string                      `json:"model_name,required"`
	ModelSource string                      `json:"model_source,required"`
	JSON        modelUploadResponseDataJSON `json:"-"`
}

// modelUploadResponseDataJSON contains the JSON metadata for the struct
// [ModelUploadResponseData]
type modelUploadResponseDataJSON struct {
	JobID       apijson.Field
	ModelID     apijson.Field
	ModelName   apijson.Field
	ModelSource apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelUploadResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelUploadResponseDataJSON) RawJSON() string {
	return r.raw
}

type ModelUploadParams struct {
	// The name to give to your uploaded model
	ModelName param.Field[string] `json:"model_name,required"`
	// The source location of the model (Hugging Face repo or S3 path)
	ModelSource param.Field[string] `json:"model_source,required"`
	// A description of your model
	Description param.Field[string] `json:"description"`
	// Hugging Face token (if uploading from Hugging Face)
	HfToken param.Field[string] `json:"hf_token"`
}

func (r ModelUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
