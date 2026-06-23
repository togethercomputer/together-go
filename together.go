// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/packages/respjson"
)

type WhoamiResponse struct {
	// The ID of the API key that authenticated the request.
	APIKeyID string `json:"api_key_id" api:"required"`
	// The ID of the organization that owns the project.
	OrganizationID string `json:"organization_id" api:"required"`
	// Human-readable name of the organization.
	OrganizationName string `json:"organization_name" api:"required"`
	// The ID of the project the API key is scoped to.
	ProjectID string `json:"project_id" api:"required"`
	// Human-readable name of the project.
	ProjectName string `json:"project_name" api:"required"`
	// DNS-friendly project identifier. Used with an endpoint slug as
	// `<project_slug>/<endpoint_slug>` to form the `model` value in dedicated endpoint
	// inference calls.
	ProjectSlug string `json:"project_slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyID         respjson.Field
		OrganizationID   respjson.Field
		OrganizationName respjson.Field
		ProjectID        respjson.Field
		ProjectName      respjson.Field
		ProjectSlug      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiResponse) RawJSON() string { return r.JSON.raw }
func (r *WhoamiResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
