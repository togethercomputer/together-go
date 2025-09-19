// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/tidwall/gjson"
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/shared"
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
func NewCodeInterpreterSessionService(opts ...option.RequestOption) (r *CodeInterpreterSessionService) {
	r = &CodeInterpreterSessionService{}
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
	Data   SessionListResponseData          `json:"data"`
	Errors []SessionListResponseErrorsUnion `json:"errors"`
	JSON   sessionListResponseJSON          `json:"-"`
}

// sessionListResponseJSON contains the JSON metadata for the struct
// [SessionListResponse]
type sessionListResponseJSON struct {
	Data        apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListResponseJSON) RawJSON() string {
	return r.raw
}

type SessionListResponseData struct {
	Sessions []SessionListResponseDataSession `json:"sessions,required"`
	JSON     sessionListResponseDataJSON      `json:"-"`
}

// sessionListResponseDataJSON contains the JSON metadata for the struct
// [SessionListResponseData]
type sessionListResponseDataJSON struct {
	Sessions    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListResponseDataJSON) RawJSON() string {
	return r.raw
}

type SessionListResponseDataSession struct {
	// Session Identifier. Used to make follow-up calls.
	ID            string                             `json:"id,required"`
	ExecuteCount  int64                              `json:"execute_count,required"`
	ExpiresAt     time.Time                          `json:"expires_at,required" format:"date-time"`
	LastExecuteAt time.Time                          `json:"last_execute_at,required" format:"date-time"`
	StartedAt     time.Time                          `json:"started_at,required" format:"date-time"`
	JSON          sessionListResponseDataSessionJSON `json:"-"`
}

// sessionListResponseDataSessionJSON contains the JSON metadata for the struct
// [SessionListResponseDataSession]
type sessionListResponseDataSessionJSON struct {
	ID            apijson.Field
	ExecuteCount  apijson.Field
	ExpiresAt     apijson.Field
	LastExecuteAt apijson.Field
	StartedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SessionListResponseDataSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListResponseDataSessionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [SessionListResponseErrorsMap].
type SessionListResponseErrorsUnion interface {
	ImplementsSessionListResponseErrorsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionListResponseErrorsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SessionListResponseErrorsMap{}),
		},
	)
}

type SessionListResponseErrorsMap map[string]interface{}

func (r SessionListResponseErrorsMap) ImplementsSessionListResponseErrorsUnion() {}
