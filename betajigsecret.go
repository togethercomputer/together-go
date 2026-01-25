// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaJigSecretService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaJigSecretService] method instead.
type BetaJigSecretService struct {
	Options []option.RequestOption
}

// NewBetaJigSecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaJigSecretService(opts ...option.RequestOption) (r BetaJigSecretService) {
	r = BetaJigSecretService{}
	r.Options = opts
	return
}

// Create a new secret to store sensitive configuration values
func (r *BetaJigSecretService) New(ctx context.Context, body BetaJigSecretNewParams, opts ...option.RequestOption) (res *Secret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific secret by its ID or name
func (r *BetaJigSecretService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Secret, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("secrets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing secret's value or metadata
func (r *BetaJigSecretService) Update(ctx context.Context, id string, body BetaJigSecretUpdateParams, opts ...option.RequestOption) (res *Secret, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("secrets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve all secrets in your project
func (r *BetaJigSecretService) List(ctx context.Context, opts ...option.RequestOption) (res *BetaJigSecretListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing secret
func (r *BetaJigSecretService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *BetaJigSecretDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("secrets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Secret struct {
	// ID is the unique identifier for this secret
	ID string `json:"id"`
	// CreatedAt is the ISO8601 timestamp when this secret was created
	CreatedAt string `json:"created_at"`
	// CreatedBy is the identifier of the user who created this secret
	CreatedBy string `json:"created_by"`
	// Description is a human-readable description of the secret's purpose
	Description string `json:"description"`
	// LastUpdatedBy is the identifier of the user who last updated this secret
	LastUpdatedBy string `json:"last_updated_by"`
	// Name is the name/key of the secret
	Name string `json:"name"`
	// Object is the type identifier for this response (always "secret")
	Object string `json:"object"`
	// UpdatedAt is the ISO8601 timestamp when this secret was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CreatedBy     respjson.Field
		Description   respjson.Field
		LastUpdatedBy respjson.Field
		Name          respjson.Field
		Object        respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Secret) RawJSON() string { return r.JSON.raw }
func (r *Secret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigSecretListResponse struct {
	// Data is the array of secret items
	Data []Secret `json:"data"`
	// Object is the type identifier for this response (always "list")
	Object string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaJigSecretListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaJigSecretListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigSecretDeleteResponse = any

type BetaJigSecretNewParams struct {
	// Name is the unique identifier for the secret. Can contain alphanumeric
	// characters, underscores, hyphens, forward slashes, and periods (1-100
	// characters)
	Name string `json:"name,required"`
	// Value is the sensitive data to store securely (e.g., API keys, passwords,
	// tokens). This value will be encrypted at rest
	Value string `json:"value,required"`
	// Description is an optional human-readable description of the secret's purpose
	// (max 500 characters)
	Description param.Opt[string] `json:"description,omitzero"`
	// ProjectID is ignored - the project is automatically determined from your
	// authentication
	ProjectID param.Opt[string] `json:"project_id,omitzero"`
	paramObj
}

func (r BetaJigSecretNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigSecretNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigSecretNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaJigSecretUpdateParams struct {
	// Description is an optional human-readable description of the secret's purpose
	// (max 500 characters)
	Description param.Opt[string] `json:"description,omitzero"`
	// Name is the new unique identifier for the secret. Can contain alphanumeric
	// characters, underscores, hyphens, forward slashes, and periods (1-100
	// characters)
	Name param.Opt[string] `json:"name,omitzero"`
	// ProjectID is ignored - the project is automatically determined from your
	// authentication
	ProjectID param.Opt[string] `json:"project_id,omitzero"`
	// Value is the new sensitive data to store securely. Updating this will replace
	// the existing secret value
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r BetaJigSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaJigSecretUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaJigSecretUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
