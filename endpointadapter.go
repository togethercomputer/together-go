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

// EndpointAdapterService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEndpointAdapterService] method instead.
type EndpointAdapterService struct {
	Options []option.RequestOption
}

// NewEndpointAdapterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEndpointAdapterService(opts ...option.RequestOption) (r EndpointAdapterService) {
	r = EndpointAdapterService{}
	r.Options = opts
	return
}

// Returns all LoRA adapters bound to the specified dedicated endpoint.
func (r *EndpointAdapterService) List(ctx context.Context, endpointID string, opts ...option.RequestOption) (res *EndpointAdapterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	path := fmt.Sprintf("endpoints/%s/adapters", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Adds a LoRA adapter model to a dedicated endpoint. After this call, inference
// requests to the adapter model name will be routed to the specified endpoint. The
// endpoint must have LoRA enabled, and the adapter's base model must be compatible
// with the endpoint's model. The endpoint name prefix in model_id must match the
// resolved endpoint.
func (r *EndpointAdapterService) Add(ctx context.Context, endpointID string, body EndpointAdapterAddParams, opts ...option.RequestOption) (res *EndpointAdapterAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	path := fmt.Sprintf("endpoints/%s/adapters", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Removes the routing rule that binds an adapter to an endpoint. The adapter must
// be currently bound to this specific endpoint.
func (r *EndpointAdapterService) Remove(ctx context.Context, endpointID string, body EndpointAdapterRemoveParams, opts ...option.RequestOption) (res *EndpointAdapterRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	path := fmt.Sprintf("endpoints/%s/adapters", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type EndpointAdapterListResponse struct {
	Data   []EndpointAdapterListResponseData `json:"data"`
	Object string                            `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointAdapterListResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointAdapterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointAdapterListResponseData struct {
	AdapterName  string `json:"adapter_name"`
	EndpointName string `json:"endpoint_name"`
	// Combined endpoint:adapter identifier
	ModelID string `json:"model_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdapterName  respjson.Field
		EndpointName respjson.Field
		ModelID      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointAdapterListResponseData) RawJSON() string { return r.JSON.raw }
func (r *EndpointAdapterListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointAdapterAddResponse struct {
	ModelID string `json:"model_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointAdapterAddResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointAdapterAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointAdapterRemoveResponse struct {
	Deleted bool   `json:"deleted"`
	ModelID string `json:"model_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ModelID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointAdapterRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointAdapterRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointAdapterAddParams struct {
	// Combined identifier in format "endpoint_name:adapter_model_name".
	ModelID string `json:"model_id" api:"required"`
	paramObj
}

func (r EndpointAdapterAddParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointAdapterAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointAdapterAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointAdapterRemoveParams struct {
	// Combined identifier in format "endpoint_name:adapter_model_name".
	ModelID string `json:"model_id" api:"required"`
	paramObj
}

func (r EndpointAdapterRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointAdapterRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointAdapterRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
