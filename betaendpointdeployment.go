// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/pagination"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaEndpointDeploymentService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaEndpointDeploymentService] method instead.
type BetaEndpointDeploymentService struct {
	Options []option.RequestOption
}

// NewBetaEndpointDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaEndpointDeploymentService(opts ...option.RequestOption) (r BetaEndpointDeploymentService) {
	r = BetaEndpointDeploymentService{}
	r.Options = opts
	return
}

// Creates a model deployment under an endpoint. The deployment provisions
// asynchronously; monitor its status before routing live traffic to it.
func (r *BetaEndpointDeploymentService) New(ctx context.Context, endpointID string, params BetaEndpointDeploymentNewParams, opts ...option.RequestOption) (res *EndpointDeployment, err error) {
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
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/deployments", params.ProjectID.Value, endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a deployment's desired configuration, placement, runtime information,
// and current provisioning status.
func (r *BetaEndpointDeploymentService) Get(ctx context.Context, id string, query BetaEndpointDeploymentGetParams, opts ...option.RequestOption) (res *EndpointDeployment, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/deployments/%s", query.ProjectID.Value, query.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates mutable deployment fields such as its model, configuration, autoscaling
// bounds, or LoRA support. Changes that affect serving may trigger asynchronous
// reprovisioning.
func (r *BetaEndpointDeploymentService) Update(ctx context.Context, id string, params BetaEndpointDeploymentUpdateParams, opts ...option.RequestOption) (res *EndpointDeployment, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/deployments/%s", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the deployments attached to an endpoint, including their model,
// configuration, scaling settings, placement, and current status.
func (r *BetaEndpointDeploymentService) List(ctx context.Context, endpointID string, params BetaEndpointDeploymentListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[EndpointDeployment], err error) {
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
	path := fmt.Sprintf("projects/%s/endpoints/%s/deployments", params.ProjectID.Value, endpointID)
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

// Lists the deployments attached to an endpoint, including their model,
// configuration, scaling settings, placement, and current status.
func (r *BetaEndpointDeploymentService) ListAutoPaging(ctx context.Context, endpointID string, params BetaEndpointDeploymentListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[EndpointDeployment] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, endpointID, params, opts...))
}

// Permanently deletes a deployment from its endpoint. Remove the deployment from
// live traffic first; use `etag` to reject the request if it changed after it was
// read.
func (r *BetaEndpointDeploymentService) Delete(ctx context.Context, id string, params BetaEndpointDeploymentDeleteParams, opts ...option.RequestOption) (res *BetaEndpointDeploymentDeleteResponse, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/deployments/%s", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Empty response returned after a successful delete operation.
type BetaEndpointDeploymentDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointDeploymentDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointDeploymentDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointDeploymentNewParams struct {
	// ID of the project that owns the endpoint.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Autoscaling configuration for a deployment.
	Autoscaling DeploymentAutoscalingParam `json:"autoscaling,omitzero" api:"required"`
	// Name for the deployment within its endpoint. Returned as a project- and
	// endpoint-qualified inference name.
	Name string `json:"name" api:"required"`
	// When true, validates the request without creating or provisioning a deployment.
	ValidateOnly param.Opt[bool] `query:"validateOnly,omitzero" json:"-"`
	// Immutable config revision in the form
	// `projects/{projectId}/configs/{configRevisionId}`. The config must be compatible
	// with the model.
	Config param.Opt[string] `json:"config,omitzero"`
	// Deprecated. Use `config`. Config revision identifier to deploy, accepted when
	// `config` is unset.
	ConfigID param.Opt[string] `json:"configId,omitzero"`
	// Enables dynamic loading of LoRA adapters on the deployment.
	EnableLora param.Opt[bool] `json:"enableLora,omitzero"`
	// Model resource name in the form
	// `projects/{projectId}/models/{modelId}[/revisions/{revisionId}]`. Omit the
	// revision segment to pin the latest revision at creation time.
	Model param.Opt[string] `json:"model,omitzero"`
	// Deprecated. Use `model`. Model identifier to serve, accepted when `model` is
	// unset.
	ModelID param.Opt[string] `json:"modelId,omitzero"`
	// Deprecated. Use `model` with a /revisions/{revisionId} segment. If omitted, the
	// latest revision is resolved at creation.
	ModelRevisionID param.Opt[string] `json:"modelRevisionId,omitzero"`
	// Placement controls where a deployment is scheduled.
	Placement BetaEndpointDeploymentNewParamsPlacementUnion `json:"placement,omitzero"`
	paramObj
}

func (r BetaEndpointDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointDeploymentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointDeploymentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaEndpointDeploymentNewParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointDeploymentNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaEndpointDeploymentNewParamsPlacementUnion struct {
	OfInline  *BetaEndpointDeploymentNewParamsPlacementInline  `json:",omitzero,inline"`
	OfProfile *BetaEndpointDeploymentNewParamsPlacementProfile `json:",omitzero,inline"`
	paramUnion
}

func (u BetaEndpointDeploymentNewParamsPlacementUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInline, u.OfProfile)
}
func (u *BetaEndpointDeploymentNewParamsPlacementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BetaEndpointDeploymentNewParamsPlacementUnion) asAny() any {
	if !param.IsOmitted(u.OfInline) {
		return u.OfInline
	} else if !param.IsOmitted(u.OfProfile) {
		return u.OfProfile
	}
	return nil
}

// The property Inline is required.
type BetaEndpointDeploymentNewParamsPlacementInline struct {
	// Inline placement parameters expanded into scheduling rules by the server.
	Inline DeploymentPlacementConfigParam `json:"inline,omitzero" api:"required"`
	paramObj
}

func (r BetaEndpointDeploymentNewParamsPlacementInline) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointDeploymentNewParamsPlacementInline
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointDeploymentNewParamsPlacementInline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Profile is required.
type BetaEndpointDeploymentNewParamsPlacementProfile struct {
	// UID of a saved placement profile.
	Profile string `json:"profile" api:"required"`
	paramObj
}

func (r BetaEndpointDeploymentNewParamsPlacementProfile) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointDeploymentNewParamsPlacementProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointDeploymentNewParamsPlacementProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointDeploymentGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	paramObj
}

type BetaEndpointDeploymentUpdateParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Fields to update. If not set, the fields populated on `deployment` are updated.
	UpdateMask param.Opt[string] `query:"updateMask,omitzero" json:"-"`
	// Current deployment version. The update is rejected if this value no longer
	// matches.
	Etag param.Opt[string] `json:"etag,omitzero"`
	// Updated inference-addressable deployment name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Autoscaling configuration for a deployment.
	Autoscaling DeploymentAutoscalingParam `json:"autoscaling,omitzero"`
	paramObj
}

func (r BetaEndpointDeploymentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointDeploymentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointDeploymentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaEndpointDeploymentUpdateParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointDeploymentUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointDeploymentListParams struct {
	// ID of the project that owns the endpoint.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous deployment list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Filter expression using `name`, `state`, `model`, `created_at`, or `updated_at`
	// with comparison operators and AND/OR/NOT; `state` takes a DeploymentState enum
	// name and `model` takes a model resource name. `name` supports substring matching
	// with `:` and prefix/suffix wildcards with `*`, and accepts a bare deployment
	// name or `<project_slug>/<endpoint_name>/<deployment_name>`.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// Maximum number of deployments to return. Max 500, defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort field for the results. Supports `created_at` or `updated_at`, optionally
	// followed by `asc` or `desc`.
	OrderBy param.Opt[string] `query:"orderBy,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointDeploymentListParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointDeploymentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointDeploymentDeleteParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Etag for optimistic concurrency. If set, the delete is rejected if the current
	// etag does not match.
	Etag param.Opt[string] `query:"etag,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointDeploymentDeleteParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointDeploymentDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
