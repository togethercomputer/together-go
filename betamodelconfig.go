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
)

// BetaModelConfigService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaModelConfigService] method instead.
type BetaModelConfigService struct {
	Options []option.RequestOption
}

// NewBetaModelConfigService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaModelConfigService(opts ...option.RequestOption) (r BetaModelConfigService) {
	r = BetaModelConfigService{}
	r.Options = opts
	return
}

// Retrieves a model configuration revision by ID, including its runtime selectors
// and certifications.
func (r *BetaModelConfigService) Get(ctx context.Context, id string, query BetaModelConfigGetParams, opts ...option.RequestOption) (res *Config, err error) {
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
	path := fmt.Sprintf("projects/%s/configs/%s", query.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists production-ready configuration revisions compatible with a reference
// model. Specify the model with `referenceModel` or the deprecated
// `referenceModelId`; if both are supplied, they must identify the same model.
// Results include public configurations and configurations owned by the specified
// project.
func (r *BetaModelConfigService) List(ctx context.Context, params BetaModelConfigListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Config], err error) {
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
	path := fmt.Sprintf("projects/%s/configs", params.ProjectID.Value)
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

// Lists production-ready configuration revisions compatible with a reference
// model. Specify the model with `referenceModel` or the deprecated
// `referenceModelId`; if both are supplied, they must identify the same model.
// Results include public configurations and configurations owned by the specified
// project.
func (r *BetaModelConfigService) ListAutoPaging(ctx context.Context, params BetaModelConfigListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Config] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Immutable, user-facing configuration revision that defines how a compatible
// model runs, including engine and hardware selectors.
type Config struct {
	// Config revision identifier.
	ID string `json:"id" api:"required"`
	// Model, hardware, and runtime combinations certified for this config revision.
	Certifications []ConfigCertification `json:"certifications" api:"required"`
	// ID of the project that owns the config revision. Public configs may be owned by
	// a different project than the deployment.
	ProjectID string `json:"projectId" api:"required"`
	// Resource name of the referenced model, using
	// `projects/{modelProject}/models/{modelId}`.
	ReferenceModel string `json:"referenceModel" api:"required"`
	// Deprecated. Use `referenceModel`. Reference model identifier.
	ReferenceModelID string `json:"referenceModelId" api:"required"`
	// Hardware and runtime selectors used to place and configure replicas.
	Selectors []ConfigSelector `json:"selectors" api:"required"`
	// Resource name of the draft model, using
	// `projects/{draftProject}/models/{modelId}`; empty when speculative decoding is
	// not enabled.
	DraftModel string `json:"draftModel"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Certifications   respjson.Field
		ProjectID        respjson.Field
		ReferenceModel   respjson.Field
		ReferenceModelID respjson.Field
		Selectors        respjson.Field
		DraftModel       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Config) RawJSON() string { return r.JSON.raw }
func (r *Config) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Certification result for a model, config, and optional draft-model combination.
type ConfigCertification struct {
	// Whether the model and config combination passed certification.
	//
	// Any of "CERTIFICATION_TYPE_CERTIFIED", "CERTIFICATION_TYPE_UNCERTIFIED".
	CertificationType string `json:"certificationType" api:"required"`
	// Time when the certification decision was recorded.
	CertifiedAt time.Time `json:"certifiedAt" api:"required" format:"date-time"`
	// Service or reviewer that recorded the certification.
	CertifiedBy string `json:"certifiedBy" api:"required"`
	// Resource name of the certified model.
	Model string `json:"model" api:"required"`
	// Revision identifier of the certified model.
	ModelRevisionID string `json:"modelRevisionId" api:"required"`
	// Product or serving environment for which the combination was evaluated.
	//
	// Any of "CERTIFICATION_TARGET_DE_SERVERLESS", "CERTIFICATION_TARGET_MRE".
	Target string `json:"target" api:"required"`
	// Resource name of the certified draft model.
	DraftModel string `json:"draftModel"`
	// Revision identifier of the certified draft model.
	DraftModelRevisionID string `json:"draftModelRevisionId"`
	// Human-readable certification notes or limitations.
	Notes string `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CertificationType    respjson.Field
		CertifiedAt          respjson.Field
		CertifiedBy          respjson.Field
		Model                respjson.Field
		ModelRevisionID      respjson.Field
		Target               respjson.Field
		DraftModel           respjson.Field
		DraftModelRevisionID respjson.Field
		Notes                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigCertification) RawJSON() string { return r.JSON.raw }
func (r *ConfigCertification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hardware or runtime requirement expressed as a key-value pair.
type ConfigSelector struct {
	// Selector name, such as GPU type, GPU count, or optimization profile.
	Key string `json:"key" api:"required"`
	// Required value for the selector.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigSelector) RawJSON() string { return r.JSON.raw }
func (r *ConfigSelector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaModelConfigGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	paramObj
}

type BetaModelConfigListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Model resource-name filter using `projects/{projectId}/models/{modelId}`;
	// alternative to `referenceModelId`. If both are set, they must agree.
	ReferenceModel param.Opt[string] `query:"referenceModel,omitzero" json:"-"`
	// Deprecated. Use `referenceModel`. Reference model identifier filter; if both are
	// set, they must agree.
	ReferenceModelID param.Opt[string] `query:"referenceModelId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaModelConfigListParams]'s query parameters as
// `url.Values`.
func (r BetaModelConfigListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
