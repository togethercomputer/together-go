// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// CodeInterpreterSessionService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCodeInterpreterSessionService] method instead.
type CodeInterpreterSessionService struct {
	Options []option.RequestOption
}

// NewCodeInterpreterSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCodeInterpreterSessionService(opts ...option.RequestOption) (r CodeInterpreterSessionService) {
	r = CodeInterpreterSessionService{}
	r.Options = opts
	return
}

// Lists all your currently active sessions.
func (r *CodeInterpreterSessionService) List(ctx context.Context, opts ...option.RequestOption) (res *SessionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tci/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SessionListResponse struct {
	Data   SessionListResponseData         `json:"data"`
	Errors []SessionListResponseErrorUnion `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionListResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionListResponseData struct {
	Sessions []SessionListResponseDataSession `json:"sessions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sessions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *SessionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionListResponseDataSession struct {
	// Session Identifier. Used to make follow-up calls.
	ID            string    `json:"id,required"`
	ExecuteCount  int64     `json:"execute_count,required"`
	ExpiresAt     time.Time `json:"expires_at,required" format:"date-time"`
	LastExecuteAt time.Time `json:"last_execute_at,required" format:"date-time"`
	StartedAt     time.Time `json:"started_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ExecuteCount  respjson.Field
		ExpiresAt     respjson.Field
		LastExecuteAt respjson.Field
		StartedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionListResponseDataSession) RawJSON() string { return r.JSON.raw }
func (r *SessionListResponseDataSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SessionListResponseErrorUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfSessionListResponseErrorMapItem]
type SessionListResponseErrorUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfSessionListResponseErrorMapItem any `json:",inline"`
	JSON                              struct {
		OfString                          respjson.Field
		OfSessionListResponseErrorMapItem respjson.Field
		raw                               string
	} `json:"-"`
}

func (u SessionListResponseErrorUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SessionListResponseErrorUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SessionListResponseErrorUnion) RawJSON() string { return u.JSON.raw }

func (r *SessionListResponseErrorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
