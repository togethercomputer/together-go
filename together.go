// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"encoding/json"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/packages/respjson"
	"github.com/togethercomputer/together-go/shared/constant"
)

type WhoamiAPIKey struct {
	// The authenticated API key identifier.
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiAPIKey) RawJSON() string { return r.JSON.raw }
func (r *WhoamiAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhoamiOrgAPIKey struct {
	APIKey       WhoamiAPIKey       `json:"api_key" api:"required"`
	Organization WhoamiOrganization `json:"organization" api:"required"`
	// Always `org_api_key`.
	Type constant.OrgAPIKey `json:"type" default:"org_api_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey       respjson.Field
		Organization respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiOrgAPIKey) RawJSON() string { return r.JSON.raw }
func (r *WhoamiOrgAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhoamiOrganization struct {
	// Organization identifier.
	ID string `json:"id" api:"required"`
	// Organization display name.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiOrganization) RawJSON() string { return r.JSON.raw }
func (r *WhoamiOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhoamiProject struct {
	// Project identifier.
	ID string `json:"id" api:"required"`
	// Project display name.
	Name string `json:"name" api:"required"`
	// DNS-friendly project identifier. Used with an endpoint slug as
	// `<project_slug>/<endpoint_slug>` to form the `model` value in dedicated endpoint
	// inference calls.
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiProject) RawJSON() string { return r.JSON.raw }
func (r *WhoamiProject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhoamiProjectAPIKey struct {
	APIKey       WhoamiAPIKey       `json:"api_key" api:"required"`
	Organization WhoamiOrganization `json:"organization" api:"required"`
	Project      WhoamiProject      `json:"project" api:"required"`
	// Always `project_api_key`.
	Type constant.ProjectAPIKey `json:"type" default:"project_api_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey       respjson.Field
		Organization respjson.Field
		Project      respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiProjectAPIKey) RawJSON() string { return r.JSON.raw }
func (r *WhoamiProjectAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhoamiResponseUnion contains all possible properties and values from
// [WhoamiProjectAPIKey], [WhoamiOrgAPIKey], [WhoamiUserJwt], [WhoamiUserOrgJwt],
// [WhoamiUserOrgProjectJwt].
//
// Use the [WhoamiResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WhoamiResponseUnion struct {
	// This field is from variant [WhoamiProjectAPIKey].
	APIKey WhoamiAPIKey `json:"api_key"`
	// This field is from variant [WhoamiProjectAPIKey].
	Organization WhoamiOrganization `json:"organization"`
	// This field is from variant [WhoamiProjectAPIKey].
	Project WhoamiProject `json:"project"`
	// Any of "project_api_key", "org_api_key", "user_jwt", "user_org_jwt",
	// "user_org_project_jwt".
	Type string `json:"type"`
	// This field is from variant [WhoamiUserJwt].
	User WhoamiUser `json:"user"`
	JSON struct {
		APIKey       respjson.Field
		Organization respjson.Field
		Project      respjson.Field
		Type         respjson.Field
		User         respjson.Field
		raw          string
	} `json:"-"`
}

// anyWhoamiResponse is implemented by each variant of [WhoamiResponseUnion] to add
// type safety for the return type of [WhoamiResponseUnion.AsAny]
type anyWhoamiResponse interface {
	implWhoamiResponseUnion()
}

func (WhoamiProjectAPIKey) implWhoamiResponseUnion()     {}
func (WhoamiOrgAPIKey) implWhoamiResponseUnion()         {}
func (WhoamiUserJwt) implWhoamiResponseUnion()           {}
func (WhoamiUserOrgJwt) implWhoamiResponseUnion()        {}
func (WhoamiUserOrgProjectJwt) implWhoamiResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WhoamiResponseUnion.AsAny().(type) {
//	case together.WhoamiProjectAPIKey:
//	case together.WhoamiOrgAPIKey:
//	case together.WhoamiUserJwt:
//	case together.WhoamiUserOrgJwt:
//	case together.WhoamiUserOrgProjectJwt:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WhoamiResponseUnion) AsAny() anyWhoamiResponse {
	switch u.Type {
	case "project_api_key":
		return u.AsProjectAPIKey()
	case "org_api_key":
		return u.AsOrgAPIKey()
	case "user_jwt":
		return u.AsUserJwt()
	case "user_org_jwt":
		return u.AsUserOrgJwt()
	case "user_org_project_jwt":
		return u.AsUserOrgProjectJwt()
	}
	return nil
}

func (u WhoamiResponseUnion) AsProjectAPIKey() (v WhoamiProjectAPIKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WhoamiResponseUnion) AsOrgAPIKey() (v WhoamiOrgAPIKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WhoamiResponseUnion) AsUserJwt() (v WhoamiUserJwt) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WhoamiResponseUnion) AsUserOrgJwt() (v WhoamiUserOrgJwt) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WhoamiResponseUnion) AsUserOrgProjectJwt() (v WhoamiUserOrgProjectJwt) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WhoamiResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *WhoamiResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhoamiUser struct {
	// Authenticated user identifier.
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiUser) RawJSON() string { return r.JSON.raw }
func (r *WhoamiUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhoamiUserJwt struct {
	// Always `user_jwt`.
	Type constant.UserJwt `json:"type" default:"user_jwt"`
	User WhoamiUser       `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiUserJwt) RawJSON() string { return r.JSON.raw }
func (r *WhoamiUserJwt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhoamiUserOrgJwt struct {
	Organization WhoamiOrganization `json:"organization" api:"required"`
	// Always `user_org_jwt`.
	Type constant.UserOrgJwt `json:"type" default:"user_org_jwt"`
	User WhoamiUser          `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Type         respjson.Field
		User         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiUserOrgJwt) RawJSON() string { return r.JSON.raw }
func (r *WhoamiUserOrgJwt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhoamiUserOrgProjectJwt struct {
	Organization WhoamiOrganization `json:"organization" api:"required"`
	Project      WhoamiProject      `json:"project" api:"required"`
	// Always `user_org_project_jwt`.
	Type constant.UserOrgProjectJwt `json:"type" default:"user_org_project_jwt"`
	User WhoamiUser                 `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Project      respjson.Field
		Type         respjson.Field
		User         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhoamiUserOrgProjectJwt) RawJSON() string { return r.JSON.raw }
func (r *WhoamiUserOrgProjectJwt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
