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

// BetaEndpointAbExperimentService contains methods and other services that help
// with interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaEndpointAbExperimentService] method instead.
type BetaEndpointAbExperimentService struct {
	Options []option.RequestOption
}

// NewBetaEndpointAbExperimentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBetaEndpointAbExperimentService(opts ...option.RequestOption) (r BetaEndpointAbExperimentService) {
	r = BetaEndpointAbExperimentService{}
	r.Options = opts
	return
}

// Creates a managed control/variant split across two to 20 deployments under the
// same endpoint. Exactly one member is the control, member percentages must add up
// to 100, and the split applies only to traffic that the endpoint would otherwise
// send to the control.
func (r *BetaEndpointAbExperimentService) New(ctx context.Context, endpointID string, params BetaEndpointAbExperimentNewParams, opts ...option.RequestOption) (res *AbExperiment, err error) {
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
	path := fmt.Sprintf("projects/%s/endpoints/%s/abExperiments", params.ProjectID.Value, endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves an A/B experiment and its participating deployments, roles, and
// traffic percentages.
func (r *BetaEndpointAbExperimentService) Get(ctx context.Context, id string, query BetaEndpointAbExperimentGetParams, opts ...option.RequestOption) (res *AbExperiment, err error) {
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
	path := fmt.Sprintf("projects/%s/endpoints/%s/abExperiments/%s", query.ProjectID.Value, query.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an experiment's description or member traffic percentages. Use the
// experiment etag for optimistic concurrency.
func (r *BetaEndpointAbExperimentService) Update(ctx context.Context, id string, params BetaEndpointAbExperimentUpdateParams, opts ...option.RequestOption) (res *AbExperiment, err error) {
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
	path := fmt.Sprintf("projects/%s/endpoints/%s/abExperiments/%s", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the managed live-traffic experiments configured for an endpoint.
func (r *BetaEndpointAbExperimentService) List(ctx context.Context, endpointID string, params BetaEndpointAbExperimentListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[AbExperiment], err error) {
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
	path := fmt.Sprintf("projects/%s/endpoints/%s/abExperiments", params.ProjectID.Value, endpointID)
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

// Lists the managed live-traffic experiments configured for an endpoint.
func (r *BetaEndpointAbExperimentService) ListAutoPaging(ctx context.Context, endpointID string, params BetaEndpointAbExperimentListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[AbExperiment] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, endpointID, params, opts...))
}

// Deletes an A/B experiment and removes its managed traffic split. The deployments
// themselves are not deleted.
func (r *BetaEndpointAbExperimentService) Delete(ctx context.Context, id string, params BetaEndpointAbExperimentDeleteParams, opts ...option.RequestOption) (res *BetaEndpointAbExperimentDeleteResponse, err error) {
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
	path := fmt.Sprintf("projects/%s/endpoints/%s/abExperiments/%s", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Managed cohort split that subdivides a control deployment's live traffic among
// the control and one or more variants.
type AbExperiment struct {
	// Output only. Unique A/B experiment identifier.
	ID string `json:"id" api:"required"`
	// Output only. Timestamp when the A/B experiment was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Output only. Identifier of the principal that created the A/B experiment.
	CreatedBy string `json:"createdBy" api:"required"`
	// Output only. Endpoint this A/B experiment belongs to.
	EndpointID string `json:"endpointId" api:"required"`
	// Optional opaque version tag for optimistic concurrency control.
	Etag string `json:"etag" api:"required"`
	// Two to 20 participating deployments with exactly one control and percentages
	// that add up to 100.
	Members []AbMember `json:"members" api:"required"`
	// Human-readable A/B experiment name, unique within the endpoint.
	Name string `json:"name" api:"required"`
	// Output only. Project that owns the parent endpoint.
	ProjectID string `json:"projectId" api:"required"`
	// Output only. Timestamp when the A/B experiment was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Optional free-form description.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		EndpointID  respjson.Field
		Etag        respjson.Field
		Members     respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AbExperiment) RawJSON() string { return r.JSON.raw }
func (r *AbExperiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Empty response returned after a successful delete operation.
type BetaEndpointAbExperimentDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointAbExperimentDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointAbExperimentDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointAbExperimentNewParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Two to 20 participating deployments with exactly one control. Integer traffic
	// percentages across all members must add up to 100.
	Members []AbMemberParam `json:"members,omitzero" api:"required"`
	// Human-readable A/B experiment name, unique within the endpoint.
	Name string `json:"name" api:"required"`
	// Optional free-form description.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r BetaEndpointAbExperimentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointAbExperimentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointAbExperimentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointAbExperimentGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	paramObj
}

type BetaEndpointAbExperimentUpdateParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Fields to update. If omitted, all mutable fields are overwritten.
	UpdateMask param.Opt[string] `query:"updateMask,omitzero" json:"-"`
	// Updated free-form description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Opaque version tag from a prior read for optimistic concurrency.
	Etag param.Opt[string] `json:"etag,omitzero"`
	// Complete replacement member set. Requires two to 20 deployments, exactly one
	// control, and percentages that add up to 100.
	Members []AbMemberParam `json:"members,omitzero"`
	paramObj
}

func (r BetaEndpointAbExperimentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointAbExperimentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointAbExperimentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaEndpointAbExperimentUpdateParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointAbExperimentUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointAbExperimentListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous A/B experiment list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of A/B experiments to return. Max 500, defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointAbExperimentListParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointAbExperimentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointAbExperimentDeleteParams struct {
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

// URLQuery serializes [BetaEndpointAbExperimentDeleteParams]'s query parameters as
// `url.Values`.
func (r BetaEndpointAbExperimentDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
