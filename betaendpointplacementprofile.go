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

// BetaEndpointPlacementProfileService contains methods and other services that
// help with interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaEndpointPlacementProfileService] method instead.
type BetaEndpointPlacementProfileService struct {
	Options []option.RequestOption
}

// NewBetaEndpointPlacementProfileService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBetaEndpointPlacementProfileService(opts ...option.RequestOption) (r BetaEndpointPlacementProfileService) {
	r = BetaEndpointPlacementProfileService{}
	r.Options = opts
	return
}

// Retrieves a reusable placement profile and its ordered region preferences.
func (r *BetaEndpointPlacementProfileService) Get(ctx context.Context, id string, query BetaEndpointPlacementProfileGetParams, opts ...option.RequestOption) (res *PlacementProfile, err error) {
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
	path := fmt.Sprintf("projects/%s/placement-profiles/%s", query.ProjectID.Value, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists reusable, project-visible placement policies that control the regions
// where deployments may be scheduled.
func (r *BetaEndpointPlacementProfileService) List(ctx context.Context, params BetaEndpointPlacementProfileListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[PlacementProfile], err error) {
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
	path := fmt.Sprintf("projects/%s/placement-profiles", params.ProjectID.Value)
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

// Lists reusable, project-visible placement policies that control the regions
// where deployments may be scheduled.
func (r *BetaEndpointPlacementProfileService) ListAutoPaging(ctx context.Context, params BetaEndpointPlacementProfileListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[PlacementProfile] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Reusable ordered region preferences for scheduling a project's deployments.
type PlacementProfile struct {
	// Unique placement profile identifier.
	ID string `json:"id" api:"required"`
	// Human-readable placement profile name.
	Name string `json:"name" api:"required"`
	// Organization that owns the placement profile.
	OrganizationID string `json:"organizationId" api:"required"`
	// Preferred deployment regions in descending priority order.
	PreferredRegions []string `json:"preferredRegions" api:"required"`
	// Project that owns the placement profile.
	ProjectID string `json:"projectId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Name             respjson.Field
		OrganizationID   respjson.Field
		PreferredRegions respjson.Field
		ProjectID        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementProfile) RawJSON() string { return r.JSON.raw }
func (r *PlacementProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointPlacementProfileGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	paramObj
}

type BetaEndpointPlacementProfileListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous placement profile list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of profiles to return. Max 500, defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointPlacementProfileListParams]'s query parameters
// as `url.Values`.
func (r BetaEndpointPlacementProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
