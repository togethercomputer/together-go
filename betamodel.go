// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/pagination"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
	"github.com/togethercomputer/together-go/shared/constant"
)

// BetaModelService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaModelService] method instead.
type BetaModelService struct {
	Options       []option.RequestOption
	RemoteUploads BetaModelRemoteUploadService
	Configs       BetaModelConfigService
}

// NewBetaModelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaModelService(opts ...option.RequestOption) (r BetaModelService) {
	r = BetaModelService{}
	r.Options = opts
	r.RemoteUploads = NewBetaModelRemoteUploadService(opts...)
	r.Configs = NewBetaModelConfigService(opts...)
	return
}

// Registers a custom model resource in the project. Registration creates the
// model's metadata; upload or import model files separately before deploying it.
func (r *BetaModelService) New(ctx context.Context, params BetaModelNewParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/models", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a custom model's metadata, visibility, weight information, and
// base-model relationship.
func (r *BetaModelService) Get(ctx context.Context, id string, query BetaModelGetParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	if query.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/models/%s", query.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates mutable model metadata such as its inference name, description, base
// model, or visibility.
func (r *BetaModelService) Update(ctx context.Context, id string, params BetaModelUpdateParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/models/%s", params.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists custom model resources owned by the specified project. Use the
// organization endpoint to list models shared across projects or the
// supported-model catalog to discover Together-hosted base models.
func (r *BetaModelService) List(ctx context.Context, params BetaModelListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Model], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/models", params.ProjectID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists custom model resources owned by the specified project. Use the
// organization endpoint to list models shared across projects or the
// supported-model catalog to discover Together-hosted base models.
func (r *BetaModelService) ListAutoPaging(ctx context.Context, params BetaModelListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Model] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Permanently deletes a custom model resource. The model must not be in use by an
// active deployment.
func (r *BetaModelService) Delete(ctx context.Context, id string, body BetaModelDeleteParams, opts ...option.RequestOption) (res *BetaModelDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.ProjectID)
	if body.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/models/%s", body.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Lists files in the latest or specified revision of a model, including paths,
// sizes, and content hashes.
func (r *BetaModelService) ListFiles(ctx context.Context, id string, params BetaModelListFilesParams, opts ...option.RequestOption) (res *BetaModelListFilesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/models/%s/files", params.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Lists custom models shared with every project in the specified organization.
// Project-private and public models are not included.
func (r *BetaModelService) ListOrgScoped(ctx context.Context, organizationID string, query BetaModelListOrgScopedParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Model], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/models", organizationID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists custom models shared with every project in the specified organization.
// Project-private and public models are not included.
func (r *BetaModelService) ListOrgScopedAutoPaging(ctx context.Context, organizationID string, query BetaModelListOrgScopedParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Model] {
	return pagination.NewCursorPaginationAutoPager(r.ListOrgScoped(ctx, organizationID, query, opts...))
}

// Lists the immutable file revisions available for a custom model, newest first.
func (r *BetaModelService) ListRevisions(ctx context.Context, id string, query BetaModelListRevisionsParams, opts ...option.RequestOption) (res *BetaModelListRevisionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	if query.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/models/%s/revisions", query.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists Together-hosted base models that can be deployed for dedicated inference,
// together with their capabilities and certified deployment profiles.
func (r *BetaModelService) ListSupported(ctx context.Context, query BetaModelListSupportedParams, opts ...option.RequestOption) (res *pagination.CursorPagination[SupportedModel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	path := "supported-models"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists Together-hosted base models that can be deployed for dedicated inference,
// together with their capabilities and certified deployment profiles.
func (r *BetaModelService) ListSupportedAutoPaging(ctx context.Context, query BetaModelListSupportedParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[SupportedModel] {
	return pagination.NewCursorPaginationAutoPager(r.ListSupported(ctx, query, opts...))
}

// Retrieves a Together-hosted base model and the certified model, configuration,
// hardware, and performance profiles available for deployment.
func (r *BetaModelService) GetSupported(ctx context.Context, id string, opts ...option.RequestOption) (res *SupportedModel, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("supported-models/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Custom or derived model registered in a project and backed by versioned weight
// files.
type Model struct {
	// Unique model identifier.
	ID string `json:"id" api:"required"`
	// Project-qualified model name in the form `<project_slug>/<model_name>`. Create
	// and update requests may use the bare or qualified form.
	Name string `json:"name" api:"required"`
	// ID of the organization that owns the model's project.
	OrganizationID string `json:"organizationId" api:"required"`
	// ID of the project that owns the model.
	ProjectID string `json:"projectId" api:"required"`
	// Who can discover the model. `VISIBILITY_PRIVATE` restricts it to the project;
	// `VISIBILITY_INTERNAL` shares it with the organization.
	//
	// Any of "VISIBILITY_PRIVATE", "VISIBILITY_INTERNAL".
	Visibility ModelVisibility `json:"visibility" api:"required"`
	// Architecture, size, precision, and speculative-decoding metadata derived from
	// the model files.
	Weights ModelWeights `json:"weights" api:"required"`
	// Resource name of the base model, using
	// `projects/{baseProject}/models/{baseModelId}`; empty when the model has no base.
	BaseModel string `json:"baseModel"`
	// ID of the supported or custom base model from which this model was derived.
	BaseModelID string `json:"baseModelId"`
	// Human-readable description of the model and its intended use.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		ProjectID      respjson.Field
		Visibility     respjson.Field
		Weights        respjson.Field
		BaseModel      respjson.Field
		BaseModelID    respjson.Field
		Description    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who can discover the model. `VISIBILITY_PRIVATE` restricts it to the project;
// `VISIBILITY_INTERNAL` shares it with the organization.
type ModelVisibility string

const (
	ModelVisibilityVisibilityPrivate  ModelVisibility = "VISIBILITY_PRIVATE"
	ModelVisibilityVisibilityInternal ModelVisibility = "VISIBILITY_INTERNAL"
)

// Architecture, size, precision, and speculative-decoding metadata derived from
// the model files.
type ModelWeights struct {
	// Model architecture detected from the weight metadata.
	Architecture string `json:"architecture"`
	// Maximum context length reported by the model metadata.
	ContextLength string `json:"contextLength"`
	// Draft-model speculator family for draft speculative decoding.
	//
	// Any of "DRAFT_SPECULATOR_TYPE_EAGLE", "DRAFT_SPECULATOR_TYPE_PHOENIX".
	DraftSpeculatorType string `json:"draftSpeculatorType"`
	// Total parameter count and breakdown by numerical data type.
	Parameters ModelWeightsParameters `json:"parameters"`
	// Speculative decoding mechanism for speculator weights.
	//
	// Any of "SPECULATOR_MECHANISM_DRAFT", "SPECULATOR_MECHANISM_LOOKAHEAD",
	// "SPECULATOR_MECHANISM_MTP".
	SpeculatorMechanism string `json:"speculatorMechanism"`
	// Role of the weights: full model, speculative draft model, or LoRA adapter.
	//
	// Any of "WEIGHTS_TYPE_DEFAULT", "WEIGHTS_TYPE_SPECULATOR",
	// "WEIGHTS_TYPE_ADAPTER".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		ContextLength       respjson.Field
		DraftSpeculatorType respjson.Field
		Parameters          respjson.Field
		SpeculatorMechanism respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelWeights) RawJSON() string { return r.JSON.raw }
func (r *ModelWeights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Total parameter count and breakdown by numerical data type.
type ModelWeightsParameters struct {
	// Parameter counts grouped by numerical data type.
	ByDtype []ModelWeightsParametersByDtype `json:"byDtype" api:"required"`
	// Total number of parameters in the model weights.
	Total string `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByDtype     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelWeightsParameters) RawJSON() string { return r.JSON.raw }
func (r *ModelWeightsParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Number of model parameters stored in one numerical data type.
type ModelWeightsParametersByDtype struct {
	// Number of model parameters stored with this data type.
	Count string `json:"count" api:"required"`
	// Numerical data type, such as `float16`, `bfloat16`, or `int8`.
	Dtype string `json:"dtype" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Dtype       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelWeightsParametersByDtype) RawJSON() string { return r.JSON.raw }
func (r *ModelWeightsParametersByDtype) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Curated catalog entry for a platform-supported model.
type SupportedModel struct {
	// Unique ID of the deployable Together-hosted base model.
	ID string `json:"id" api:"required"`
	// Resource name for the base model as `projects/{projectId}/models/{modelId}`;
	// empty when unresolved.
	BaseModel string `json:"baseModel" api:"required"`
	// Bare model ID for the architecture's base model; empty when no base model is
	// linked.
	BaseModelID string `json:"baseModelId" api:"required"`
	// High-level tasks the model supports.
	//
	// Any of "CAPABILITY_CHAT", "CAPABILITY_EMBEDDING", "CAPABILITY_RERANKING",
	// "CAPABILITY_IMAGE_GENERATION", "CAPABILITY_VIDEO_GENERATION".
	Capabilities []string `json:"capabilities" api:"required"`
	// Timestamp when the catalog entry was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Certified deployment profiles available for the model.
	DeploymentProfiles []SupportedModelDeploymentProfile `json:"deploymentProfiles" api:"required"`
	// Catalog-controlled human-readable model name.
	DisplayName string `json:"displayName" api:"required"`
	// UI-facing model type badge, such as chat, language, code, image, embedding,
	// rerank, moderation, audio, video, or transcribe.
	DisplayType string `json:"displayType" api:"required"`
	// Input modalities supported by the model.
	//
	// Any of "MODALITY_TEXT", "MODALITY_IMAGE", "MODALITY_AUDIO", "MODALITY_VIDEO".
	InputModalities []string `json:"inputModalities" api:"required"`
	// Catalog-controlled HF model ID used for inference.
	Name string `json:"name" api:"required"`
	// Output modalities produced by the model.
	//
	// Any of "MODALITY_TEXT", "MODALITY_IMAGE", "MODALITY_AUDIO", "MODALITY_VIDEO".
	OutputModalities []string `json:"outputModalities" api:"required"`
	// Product surfaces where the model is offered.
	//
	// Any of "PRODUCT_SERVERLESS", "PRODUCT_DEDICATED", "PRODUCT_FINE_TUNING".
	Products []string `json:"products" api:"required"`
	// Organization or publisher associated with the model.
	Publisher string `json:"publisher" api:"required"`
	// Catalog recommendation status for the model.
	//
	// Any of "SUPPORTED_MODEL_STATUS_RECOMMENDED", "SUPPORTED_MODEL_STATUS_SUPPORTED",
	// "SUPPORTED_MODEL_STATUS_DEPRECATED", "SUPPORTED_MODEL_STATUS_HIDDEN".
	Status SupportedModelStatus `json:"status" api:"required"`
	// Timestamp when the catalog entry was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Model architecture from the underlying weights metadata.
	Architecture string `json:"architecture"`
	// Maximum context length from the underlying weights metadata.
	ContextLength string `json:"contextLength"`
	// Human-readable model description.
	Description string `json:"description"`
	// Model family identifier for related catalog entries.
	FamilyID string `json:"familyId"`
	// Advanced features exposed by the model.
	//
	// Any of "FEATURE_TOOL_CALLING", "FEATURE_STRUCTURED_OUTPUT", "FEATURE_REASONING".
	Features []string `json:"features"`
	// Preferred input format for the model.
	InputFormat string `json:"inputFormat"`
	// Preferred output format for the model.
	OutputFormat string `json:"outputFormat"`
	// Serverless endpoint name for inference, if available.
	ServerlessEndpoint string `json:"serverlessEndpoint"`
	// Searchable catalog tags for the model.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		BaseModel          respjson.Field
		BaseModelID        respjson.Field
		Capabilities       respjson.Field
		CreatedAt          respjson.Field
		DeploymentProfiles respjson.Field
		DisplayName        respjson.Field
		DisplayType        respjson.Field
		InputModalities    respjson.Field
		Name               respjson.Field
		OutputModalities   respjson.Field
		Products           respjson.Field
		Publisher          respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		Architecture       respjson.Field
		ContextLength      respjson.Field
		Description        respjson.Field
		FamilyID           respjson.Field
		Features           respjson.Field
		InputFormat        respjson.Field
		OutputFormat       respjson.Field
		ServerlessEndpoint respjson.Field
		Tags               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportedModel) RawJSON() string { return r.JSON.raw }
func (r *SupportedModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Catalog recommendation status for the model.
type SupportedModelStatus string

const (
	SupportedModelStatusSupportedModelStatusRecommended SupportedModelStatus = "SUPPORTED_MODEL_STATUS_RECOMMENDED"
	SupportedModelStatusSupportedModelStatusSupported   SupportedModelStatus = "SUPPORTED_MODEL_STATUS_SUPPORTED"
	SupportedModelStatusSupportedModelStatusDeprecated  SupportedModelStatus = "SUPPORTED_MODEL_STATUS_DEPRECATED"
	SupportedModelStatusSupportedModelStatusHidden      SupportedModelStatus = "SUPPORTED_MODEL_STATUS_HIDDEN"
)

// Certified deployment profile for a supported model.
type SupportedModelDeploymentProfile struct {
	// Certified configuration revision identifier.
	CertifiedConfigRevisionID string `json:"certifiedConfigRevisionId" api:"required"`
	// Certified model weight revision identifier, if available.
	CertifiedModelRevisionID string `json:"certifiedModelRevisionId" api:"required"`
	// Certified config revision in the form
	// `projects/{projectId}/configs/{configRevisionId}`. Omitted when the profile does
	// not pin a config.
	Config string `json:"config" api:"required"`
	// Number of GPUs required by the profile.
	GPUCount int64 `json:"gpuCount" api:"required"`
	// GPU instance type for the profile.
	GPUType string `json:"gpuType" api:"required"`
	// Deployable model resource in the form
	// `projects/{projectId}/models/{modelId}[/revisions/{revisionId}]`. Omitted when
	// the profile does not pin model weights.
	Model string `json:"model" api:"required"`
	// Free-form parallelism spec for the profile, such as TP8, TP4, EP, or PD;
	// supersedes tensor_parallel_size.
	Parallelism string `json:"parallelism" api:"required"`
	// Performance benchmarks for the profile, if available.
	PerformanceBenchmarks SupportedModelPerformanceBenchmarks `json:"performanceBenchmarks" api:"required"`
	// Stable profile identifier, usually the certified config id.
	ProfileID string `json:"profileId" api:"required"`
	// Quantization method for the profile, if available.
	Quantization string `json:"quantization" api:"required"`
	// Deprecated. Use `parallelism`. Legacy tensor-parallel shard count for the
	// profile.
	//
	// Deprecated: deprecated
	TensorParallelSize int64 `json:"tensorParallelSize"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CertifiedConfigRevisionID respjson.Field
		CertifiedModelRevisionID  respjson.Field
		Config                    respjson.Field
		GPUCount                  respjson.Field
		GPUType                   respjson.Field
		Model                     respjson.Field
		Parallelism               respjson.Field
		PerformanceBenchmarks     respjson.Field
		ProfileID                 respjson.Field
		Quantization              respjson.Field
		TensorParallelSize        respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportedModelDeploymentProfile) RawJSON() string { return r.JSON.raw }
func (r *SupportedModelDeploymentProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Performance benchmark metrics for a supported model profile.
type SupportedModelPerformanceBenchmarks struct {
	// Decoding throughput in tokens per second.
	DecodingSpeedTps float64 `json:"decodingSpeedTps"`
	// Maximum context length supported by the profile.
	MaxContextLength string `json:"maxContextLength"`
	// Time to first token in milliseconds.
	TimeToFirstTokenMs int64 `json:"timeToFirstTokenMs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DecodingSpeedTps   respjson.Field
		MaxContextLength   respjson.Field
		TimeToFirstTokenMs respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportedModelPerformanceBenchmarks) RawJSON() string { return r.JSON.raw }
func (r *SupportedModelPerformanceBenchmarks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Empty response returned after a successful delete operation.
type BetaModelDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaModelDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaModelDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Files and aggregate size information for one model revision.
type BetaModelListFilesResponse struct {
	// Files in the selected model revision.
	Data   []BetaModelListFilesResponseData `json:"data" api:"required"`
	Object constant.List                    `json:"object" default:"list"`
	// Cursor for the next page. Null if there are no more results.
	NextCursor string `json:"next_cursor"`
	// Time when the listed model revision was created.
	RevisionCreatedAt time.Time `json:"revisionCreatedAt" format:"date-time"`
	// ID of the model revision whose files are listed.
	RevisionID string `json:"revisionId"`
	// Total size of all files in the revision, in bytes.
	TotalSizeBytes string `json:"totalSizeBytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data              respjson.Field
		Object            respjson.Field
		NextCursor        respjson.Field
		RevisionCreatedAt respjson.Field
		RevisionID        respjson.Field
		TotalSizeBytes    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaModelListFilesResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaModelListFilesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for one file in a model revision.
type BetaModelListFilesResponseData struct {
	// Content hash for integrity verification and upload deduplication.
	Hash string `json:"hash"`
	// File path within the model revision.
	Path string `json:"path"`
	// File size in bytes.
	SizeBytes string `json:"sizeBytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hash        respjson.Field
		Path        respjson.Field
		SizeBytes   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaModelListFilesResponseData) RawJSON() string { return r.JSON.raw }
func (r *BetaModelListFilesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Immutable model revisions and pagination metadata.
type BetaModelListRevisionsResponse struct {
	// Immutable revisions available for the model.
	Data []BetaModelListRevisionsResponseData `json:"data"`
	// Cursor for the next page. Null if there are no more results.
	NextCursor string `json:"next_cursor"`
	// Object type. Always `list`.
	//
	// Any of "list".
	Object BetaModelListRevisionsResponseObject `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextCursor  respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaModelListRevisionsResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaModelListRevisionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Revision metadata for a volume object.
type BetaModelListRevisionsResponseData struct {
	// Timestamp when the revision was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Revision identifier.
	RevisionID string `json:"revisionId" api:"required"`
	// Timestamp when validation most recently ran for the revision.
	LastValidatedAt time.Time `json:"lastValidatedAt" format:"date-time"`
	// Validation errors reported for the revision.
	ValidationErrors []BetaModelListRevisionsResponseDataValidationError `json:"validationErrors"`
	// Current validation status for the revision.
	//
	// Any of "REVISION_VALIDATION_STATUS_PENDING",
	// "REVISION_VALIDATION_STATUS_SUCCESS", "REVISION_VALIDATION_STATUS_FAILED",
	// "REVISION_VALIDATION_STATUS_ERROR".
	ValidationStatus string `json:"validationStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		RevisionID       respjson.Field
		LastValidatedAt  respjson.Field
		ValidationErrors respjson.Field
		ValidationStatus respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaModelListRevisionsResponseData) RawJSON() string { return r.JSON.raw }
func (r *BetaModelListRevisionsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One validation error reported for a model revision.
type BetaModelListRevisionsResponseDataValidationError struct {
	// Human-readable validation error message.
	Message string `json:"message"`
	// Validation rule that produced the error.
	Rule string `json:"rule"`
	// Severity level reported by the validation rule.
	Severity string `json:"severity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Rule        respjson.Field
		Severity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaModelListRevisionsResponseDataValidationError) RawJSON() string { return r.JSON.raw }
func (r *BetaModelListRevisionsResponseDataValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object type. Always `list`.
type BetaModelListRevisionsResponseObject string

const (
	BetaModelListRevisionsResponseObjectList BetaModelListRevisionsResponseObject = "list"
)

type BetaModelNewParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// ID of the supported base model from which this model was derived.
	BaseModelID string `json:"baseModelId" api:"required"`
	// Name for the custom model. May be bare or qualified as
	// `<project_slug>/<model_name>`; a supplied project slug must match the project in
	// the request path.
	Name string `json:"name" api:"required"`
	// Volume type to create. Use `model` or `adapter`; plural `models` and `adapters`
	// are also accepted.
	Type string `json:"type" api:"required"`
	// Human-readable description of the model and its intended use.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r BetaModelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaModelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaModelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaModelGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	paramObj
}

type BetaModelUpdateParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Fields to update. If omitted, all mutable fields are overwritten.
	UpdateMask param.Opt[string] `query:"updateMask,omitzero" format:"field-mask" json:"-"`
	// Updated user-facing model description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Updated inference-addressable model name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Who can discover the model. `VISIBILITY_PRIVATE` restricts it to the project;
	// `VISIBILITY_INTERNAL` shares it with the organization.
	//
	// Any of "VISIBILITY_PRIVATE", "VISIBILITY_INTERNAL".
	Visibility BetaModelUpdateParamsVisibility `json:"visibility,omitzero"`
	paramObj
}

func (r BetaModelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaModelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaModelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaModelUpdateParams]'s query parameters as `url.Values`.
func (r BetaModelUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Who can discover the model. `VISIBILITY_PRIVATE` restricts it to the project;
// `VISIBILITY_INTERNAL` shares it with the organization.
type BetaModelUpdateParamsVisibility string

const (
	BetaModelUpdateParamsVisibilityVisibilityPrivate  BetaModelUpdateParamsVisibility = "VISIBILITY_PRIVATE"
	BetaModelUpdateParamsVisibilityVisibilityInternal BetaModelUpdateParamsVisibility = "VISIBILITY_INTERNAL"
)

type BetaModelListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous model list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of models to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Organization whose shared models should be included. Defaults to the
	// authenticated project's organization.
	OrganizationID param.Opt[string] `query:"organizationId,omitzero" json:"-"`
	// Model visibility. Private means it is scoped to the project. Internal means it
	// is scoped to the organization.
	//
	// Any of "VISIBILITY_PRIVATE", "VISIBILITY_INTERNAL".
	Visibility BetaModelListParamsVisibility `query:"visibility,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaModelListParams]'s query parameters as `url.Values`.
func (r BetaModelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Model visibility. Private means it is scoped to the project. Internal means it
// is scoped to the organization.
type BetaModelListParamsVisibility string

const (
	BetaModelListParamsVisibilityVisibilityPrivate  BetaModelListParamsVisibility = "VISIBILITY_PRIVATE"
	BetaModelListParamsVisibilityVisibilityInternal BetaModelListParamsVisibility = "VISIBILITY_INTERNAL"
)

type BetaModelDeleteParams struct {
	// ID of the project that owns the model.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	paramObj
}

type BetaModelListFilesParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Revision identifier to read from.
	RevisionID param.Opt[string] `query:"revisionId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaModelListFilesParams]'s query parameters as
// `url.Values`.
func (r BetaModelListFilesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaModelListOrgScopedParams struct {
	// Cursor from a previous list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaModelListOrgScopedParams]'s query parameters as
// `url.Values`.
func (r BetaModelListOrgScopedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaModelListRevisionsParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	paramObj
}

type BetaModelListSupportedParams struct {
	// Cursor from a previous supported-model list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of models to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Case-insensitive search across model IDs, names, and descriptions.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter models by input modality.
	//
	// Any of "MODALITY_TEXT", "MODALITY_IMAGE", "MODALITY_AUDIO", "MODALITY_VIDEO".
	Modality BetaModelListSupportedParamsModality `query:"modality,omitzero" json:"-"`
	// Filter models by product surface.
	//
	// Any of "PRODUCT_SERVERLESS", "PRODUCT_DEDICATED", "PRODUCT_FINE_TUNING".
	Product BetaModelListSupportedParamsProduct `query:"product,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaModelListSupportedParams]'s query parameters as
// `url.Values`.
func (r BetaModelListSupportedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter models by input modality.
type BetaModelListSupportedParamsModality string

const (
	BetaModelListSupportedParamsModalityModalityText  BetaModelListSupportedParamsModality = "MODALITY_TEXT"
	BetaModelListSupportedParamsModalityModalityImage BetaModelListSupportedParamsModality = "MODALITY_IMAGE"
	BetaModelListSupportedParamsModalityModalityAudio BetaModelListSupportedParamsModality = "MODALITY_AUDIO"
	BetaModelListSupportedParamsModalityModalityVideo BetaModelListSupportedParamsModality = "MODALITY_VIDEO"
)

// Filter models by product surface.
type BetaModelListSupportedParamsProduct string

const (
	BetaModelListSupportedParamsProductProductServerless BetaModelListSupportedParamsProduct = "PRODUCT_SERVERLESS"
	BetaModelListSupportedParamsProductProductDedicated  BetaModelListSupportedParamsProduct = "PRODUCT_DEDICATED"
	BetaModelListSupportedParamsProductProductFineTuning BetaModelListSupportedParamsProduct = "PRODUCT_FINE_TUNING"
)
