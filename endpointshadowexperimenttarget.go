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

// EndpointShadowExperimentTargetService contains methods and other services that
// help with interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEndpointShadowExperimentTargetService] method instead.
type EndpointShadowExperimentTargetService struct {
	Options []option.RequestOption
}

// NewEndpointShadowExperimentTargetService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEndpointShadowExperimentTargetService(opts ...option.RequestOption) (r EndpointShadowExperimentTargetService) {
	r = EndpointShadowExperimentTargetService{}
	r.Options = opts
	return
}

// Adds a deployment under the same endpoint as a target for mirrored requests.
func (r *EndpointShadowExperimentTargetService) New(ctx context.Context, params EndpointShadowExperimentTargetNewParams, opts ...option.RequestOption) (res *ShadowExperimentTarget, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/shadowExperiments/%s/targets", params.ProjectID.Value, params.EndpointID, params.ExperimentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves one target configured to receive mirrored requests from a shadow
// experiment.
func (r *EndpointShadowExperimentTargetService) Get(ctx context.Context, id string, query EndpointShadowExperimentTargetGetParams, opts ...option.RequestOption) (res *ShadowExperimentTarget, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/shadowExperiments/%s/targets/%s", query.ProjectID.Value, query.EndpointID, query.ExperimentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a shadow target's name, deployment, or description. `updateMask` is
// required and must select at least one mutable field.
func (r *EndpointShadowExperimentTargetService) Update(ctx context.Context, id string, params EndpointShadowExperimentTargetUpdateParams, opts ...option.RequestOption) (res *ShadowExperimentTarget, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/shadowExperiments/%s/targets/%s", params.ProjectID.Value, params.EndpointID, params.ExperimentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the deployments that receive mirrored requests from a shadow experiment.
func (r *EndpointShadowExperimentTargetService) List(ctx context.Context, endpointID string, experimentID string, params EndpointShadowExperimentTargetListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ShadowExperimentTarget], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/shadowExperiments/%s/targets", params.ProjectID.Value, endpointID, experimentID)
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
func (r *EndpointShadowExperimentTargetService) ListAutoPaging(ctx context.Context, endpointID string, experimentID string, params EndpointShadowExperimentTargetListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ShadowExperimentTarget] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, endpointID, experimentID, params, opts...))
}

// Removes a target from a shadow experiment without deleting the underlying
// deployment.
func (r *EndpointShadowExperimentTargetService) Delete(ctx context.Context, id string, params EndpointShadowExperimentTargetDeleteParams, opts ...option.RequestOption) (res *EndpointShadowExperimentTargetDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/shadowExperiments/%s/targets/%s", params.ProjectID.Value, params.EndpointID, params.ExperimentID, id)
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
	// Deployment under the parent endpoint that receives mirrored traffic. It must not
	// be a live traffic-split member or the source or target of an active rollout;
	// traffic-split weight 0 warm-up targets are allowed.
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
type EndpointShadowExperimentTargetDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointShadowExperimentTargetDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointShadowExperimentTargetDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointShadowExperimentTargetNewParams struct {
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
	// Deployment under the parent endpoint that receives mirrored traffic. It must not
	// be a live traffic-split member or the source or target of an active rollout;
	// traffic-split weight 0 warm-up targets are allowed.
	TargetDeploymentID string `json:"targetDeploymentId" api:"required"`
	// Optional free-form target description.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r EndpointShadowExperimentTargetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointShadowExperimentTargetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointShadowExperimentTargetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointShadowExperimentTargetGetParams struct {
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

type EndpointShadowExperimentTargetUpdateParams struct {
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
	// Replacement deployment under the parent endpoint. It must not be a live
	// traffic-split member or the source or target of an active rollout; traffic-split
	// weight 0 warm-up targets are allowed.
	TargetDeploymentID param.Opt[string] `json:"targetDeploymentId,omitzero"`
	paramObj
}

func (r EndpointShadowExperimentTargetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointShadowExperimentTargetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointShadowExperimentTargetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [EndpointShadowExperimentTargetUpdateParams]'s query
// parameters as `url.Values`.
func (r EndpointShadowExperimentTargetUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EndpointShadowExperimentTargetListParams struct {
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

// URLQuery serializes [EndpointShadowExperimentTargetListParams]'s query
// parameters as `url.Values`.
func (r EndpointShadowExperimentTargetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EndpointShadowExperimentTargetDeleteParams struct {
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

// URLQuery serializes [EndpointShadowExperimentTargetDeleteParams]'s query
// parameters as `url.Values`.
func (r EndpointShadowExperimentTargetDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
