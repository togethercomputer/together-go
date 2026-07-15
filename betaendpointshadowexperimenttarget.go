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

// BetaEndpointShadowExperimentTargetService contains methods and other services
// that help with interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaEndpointShadowExperimentTargetService] method instead.
type BetaEndpointShadowExperimentTargetService struct {
	Options []option.RequestOption
}

// NewBetaEndpointShadowExperimentTargetService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewBetaEndpointShadowExperimentTargetService(opts ...option.RequestOption) (r BetaEndpointShadowExperimentTargetService) {
	r = BetaEndpointShadowExperimentTargetService{}
	r.Options = opts
	return
}

// Adds a deployment under the same endpoint as a target for mirrored requests.
func (r *BetaEndpointShadowExperimentTargetService) New(ctx context.Context, params BetaEndpointShadowExperimentTargetNewParams, opts ...option.RequestOption) (res *ShadowExperimentTarget, err error) {
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
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if params.ExperimentID == "" {
		err = errors.New("missing required experimentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments/%s/targets", params.ProjectID.Value, params.EndpointID, params.ExperimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves one target configured to receive mirrored requests from a shadow
// experiment.
func (r *BetaEndpointShadowExperimentTargetService) Get(ctx context.Context, id string, query BetaEndpointShadowExperimentTargetGetParams, opts ...option.RequestOption) (res *ShadowExperimentTarget, err error) {
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
	if query.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if query.ExperimentID == "" {
		err = errors.New("missing required experimentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments/%s/targets/%s", query.ProjectID.Value, query.EndpointID, query.ExperimentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a shadow target's name, deployment, or description. `updateMask` is
// required and must select at least one mutable field.
func (r *BetaEndpointShadowExperimentTargetService) Update(ctx context.Context, id string, params BetaEndpointShadowExperimentTargetUpdateParams, opts ...option.RequestOption) (res *ShadowExperimentTarget, err error) {
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
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if params.ExperimentID == "" {
		err = errors.New("missing required experimentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments/%s/targets/%s", params.ProjectID.Value, params.EndpointID, params.ExperimentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the deployments that receive mirrored requests from a shadow experiment.
func (r *BetaEndpointShadowExperimentTargetService) List(ctx context.Context, endpointID string, experimentID string, params BetaEndpointShadowExperimentTargetListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ShadowExperimentTarget], err error) {
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
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if experimentID == "" {
		err = errors.New("missing required experimentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments/%s/targets", params.ProjectID.Value, endpointID, experimentID)
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

// Lists the deployments that receive mirrored requests from a shadow experiment.
func (r *BetaEndpointShadowExperimentTargetService) ListAutoPaging(ctx context.Context, endpointID string, experimentID string, params BetaEndpointShadowExperimentTargetListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ShadowExperimentTarget] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, endpointID, experimentID, params, opts...))
}

// Removes a target from a shadow experiment without deleting the underlying
// deployment.
func (r *BetaEndpointShadowExperimentTargetService) Delete(ctx context.Context, id string, params BetaEndpointShadowExperimentTargetDeleteParams, opts ...option.RequestOption) (res *BetaEndpointShadowExperimentTargetDeleteResponse, err error) {
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
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if params.ExperimentID == "" {
		err = errors.New("missing required experimentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments/%s/targets/%s", params.ProjectID.Value, params.EndpointID, params.ExperimentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Deployment that receives mirrored traffic for a shadow experiment.
type ShadowExperimentTarget struct {
	// Output only. Unique shadow experiment target identifier.
	ID string `json:"id" api:"required"`
	// Output only. Timestamp when the target was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Opaque version tag for optimistic concurrency control. Returned on read; set it
	// on update or delete requests for consistent read-modify-write.
	Etag string `json:"etag" api:"required"`
	// Output only. Shadow experiment this target belongs to.
	ExperimentID string `json:"experimentId" api:"required"`
	// Human-readable target name, unique within the shadow experiment. At most 256
	// characters.
	Name string `json:"name" api:"required"`
	// Deployment under the parent endpoint that receives mirrored traffic. Shadow
	// targets should be excluded from the endpoint's live traffic split.
	TargetDeploymentID string `json:"targetDeploymentId" api:"required"`
	// Output only. Timestamp when the target was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Optional free-form target description.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Etag               respjson.Field
		ExperimentID       respjson.Field
		Name               respjson.Field
		TargetDeploymentID respjson.Field
		UpdatedAt          respjson.Field
		Description        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentTarget) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperimentTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Empty response returned after a successful delete operation.
type BetaEndpointShadowExperimentTargetDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointShadowExperimentTargetDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointShadowExperimentTargetDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointShadowExperimentTargetNewParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Shadow experiment identifier.
	ExperimentID string `path:"experimentId" api:"required" json:"-"`
	// Human-readable target name, unique within the shadow experiment. At most 256
	// characters.
	Name string `json:"name" api:"required"`
	// Deployment under the parent endpoint that receives mirrored traffic. Exclude it
	// from the endpoint's live traffic split.
	TargetDeploymentID string `json:"targetDeploymentId" api:"required"`
	// Optional free-form target description.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r BetaEndpointShadowExperimentTargetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointShadowExperimentTargetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointShadowExperimentTargetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointShadowExperimentTargetGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Shadow experiment identifier.
	ExperimentID string `path:"experimentId" api:"required" json:"-"`
	paramObj
}

type BetaEndpointShadowExperimentTargetUpdateParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Shadow experiment identifier.
	ExperimentID string `path:"experimentId" api:"required" json:"-"`
	// Comma-separated fields to update. Supported fields are `name`,
	// `targetDeploymentId`, and `description`.
	UpdateMask string `query:"updateMask" api:"required" json:"-"`
	// Updated free-form target description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Opaque version tag from a prior read for optimistic concurrency.
	Etag param.Opt[string] `json:"etag,omitzero"`
	// Updated human-readable target name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Replacement deployment under the parent endpoint. Exclude it from the endpoint's
	// live traffic split.
	TargetDeploymentID param.Opt[string] `json:"targetDeploymentId,omitzero"`
	paramObj
}

func (r BetaEndpointShadowExperimentTargetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointShadowExperimentTargetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointShadowExperimentTargetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaEndpointShadowExperimentTargetUpdateParams]'s query
// parameters as `url.Values`.
func (r BetaEndpointShadowExperimentTargetUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointShadowExperimentTargetListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous shadow experiment target list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of targets to return. Max 500, defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointShadowExperimentTargetListParams]'s query
// parameters as `url.Values`.
func (r BetaEndpointShadowExperimentTargetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointShadowExperimentTargetDeleteParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Shadow experiment identifier.
	ExperimentID string `path:"experimentId" api:"required" json:"-"`
	// Etag for optimistic concurrency.
	Etag param.Opt[string] `query:"etag,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointShadowExperimentTargetDeleteParams]'s query
// parameters as `url.Values`.
func (r BetaEndpointShadowExperimentTargetDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
