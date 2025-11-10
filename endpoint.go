// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"encoding/json"
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
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// EndpointService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEndpointService] method instead.
type EndpointService struct {
	Options []option.RequestOption
}

// NewEndpointService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEndpointService(opts ...option.RequestOption) (r EndpointService) {
	r = EndpointService{}
	r.Options = opts
	return
}

// Creates a new dedicated endpoint for serving models. The endpoint will
// automatically start after creation. You can deploy any supported model on
// hardware configurations that meet the model's requirements.
func (r *EndpointService) New(ctx context.Context, body EndpointNewParams, opts ...option.RequestOption) (res *DedicatedEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves details about a specific endpoint, including its current state,
// configuration, and scaling settings.
func (r *EndpointService) Get(ctx context.Context, endpointID string, opts ...option.RequestOption) (res *DedicatedEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return
	}
	path := fmt.Sprintf("endpoints/%s", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing endpoint's configuration. You can modify the display name,
// autoscaling settings, or change the endpoint's state (start/stop).
func (r *EndpointService) Update(ctx context.Context, endpointID string, body EndpointUpdateParams, opts ...option.RequestOption) (res *DedicatedEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return
	}
	path := fmt.Sprintf("endpoints/%s", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of all endpoints associated with your account. You can filter the
// results by type (dedicated or serverless).
func (r *EndpointService) List(ctx context.Context, query EndpointListParams, opts ...option.RequestOption) (res *EndpointListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently deletes an endpoint. This action cannot be undone.
func (r *EndpointService) Delete(ctx context.Context, endpointID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return
	}
	path := fmt.Sprintf("endpoints/%s", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Configuration for automatic scaling of replicas based on demand.
type Autoscaling struct {
	// The maximum number of replicas to scale up to under load
	MaxReplicas int64 `json:"max_replicas,required"`
	// The minimum number of replicas to maintain, even when there is no load
	MinReplicas int64 `json:"min_replicas,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxReplicas respjson.Field
		MinReplicas respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Autoscaling) RawJSON() string { return r.JSON.raw }
func (r *Autoscaling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Autoscaling to a AutoscalingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AutoscalingParam.Overrides()
func (r Autoscaling) ToParam() AutoscalingParam {
	return param.Override[AutoscalingParam](json.RawMessage(r.RawJSON()))
}

// Configuration for automatic scaling of replicas based on demand.
//
// The properties MaxReplicas, MinReplicas are required.
type AutoscalingParam struct {
	// The maximum number of replicas to scale up to under load
	MaxReplicas int64 `json:"max_replicas,required"`
	// The minimum number of replicas to maintain, even when there is no load
	MinReplicas int64 `json:"min_replicas,required"`
	paramObj
}

func (r AutoscalingParam) MarshalJSON() (data []byte, err error) {
	type shadow AutoscalingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutoscalingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a dedicated endpoint deployment
type DedicatedEndpoint struct {
	// Unique identifier for the endpoint
	ID string `json:"id,required"`
	// Configuration for automatic scaling of the endpoint
	Autoscaling Autoscaling `json:"autoscaling,required"`
	// Timestamp when the endpoint was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable name for the endpoint
	DisplayName string `json:"display_name,required"`
	// The hardware configuration used for this endpoint
	Hardware string `json:"hardware,required"`
	// The model deployed on this endpoint
	Model string `json:"model,required"`
	// System name for the endpoint
	Name string `json:"name,required"`
	// The type of object
	//
	// Any of "endpoint".
	Object DedicatedEndpointObject `json:"object,required"`
	// The owner of this endpoint
	Owner string `json:"owner,required"`
	// Current state of the endpoint
	//
	// Any of "PENDING", "STARTING", "STARTED", "STOPPING", "STOPPED", "ERROR".
	State DedicatedEndpointState `json:"state,required"`
	// The type of endpoint
	//
	// Any of "dedicated".
	Type DedicatedEndpointType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Autoscaling respjson.Field
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		Hardware    respjson.Field
		Model       respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedEndpoint) RawJSON() string { return r.JSON.raw }
func (r *DedicatedEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of object
type DedicatedEndpointObject string

const (
	DedicatedEndpointObjectEndpoint DedicatedEndpointObject = "endpoint"
)

// Current state of the endpoint
type DedicatedEndpointState string

const (
	DedicatedEndpointStatePending  DedicatedEndpointState = "PENDING"
	DedicatedEndpointStateStarting DedicatedEndpointState = "STARTING"
	DedicatedEndpointStateStarted  DedicatedEndpointState = "STARTED"
	DedicatedEndpointStateStopping DedicatedEndpointState = "STOPPING"
	DedicatedEndpointStateStopped  DedicatedEndpointState = "STOPPED"
	DedicatedEndpointStateError    DedicatedEndpointState = "ERROR"
)

// The type of endpoint
type DedicatedEndpointType string

const (
	DedicatedEndpointTypeDedicated DedicatedEndpointType = "dedicated"
)

type EndpointListResponse struct {
	Data []EndpointListResponseData `json:"data,required"`
	// Any of "list".
	Object EndpointListResponseObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointListResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about an endpoint when listed via the list endpoint
type EndpointListResponseData struct {
	// Unique identifier for the endpoint
	ID string `json:"id,required"`
	// Timestamp when the endpoint was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The model deployed on this endpoint
	Model string `json:"model,required"`
	// System name for the endpoint
	Name string `json:"name,required"`
	// The type of object
	//
	// Any of "endpoint".
	Object string `json:"object,required"`
	// The owner of this endpoint
	Owner string `json:"owner,required"`
	// Current state of the endpoint
	//
	// Any of "PENDING", "STARTING", "STARTED", "STOPPING", "STOPPED", "ERROR".
	State string `json:"state,required"`
	// The type of endpoint
	//
	// Any of "serverless", "dedicated".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Model       respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointListResponseData) RawJSON() string { return r.JSON.raw }
func (r *EndpointListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointListResponseObject string

const (
	EndpointListResponseObjectList EndpointListResponseObject = "list"
)

type EndpointNewParams struct {
	// Configuration for automatic scaling of the endpoint
	Autoscaling AutoscalingParam `json:"autoscaling,omitzero,required"`
	// The hardware configuration to use for this endpoint
	Hardware string `json:"hardware,required"`
	// The model to deploy on this endpoint
	Model string `json:"model,required"`
	// The number of minutes of inactivity after which the endpoint will be
	// automatically stopped. Set to null, omit or set to 0 to disable automatic
	// timeout.
	InactiveTimeout param.Opt[int64] `json:"inactive_timeout,omitzero"`
	// Whether to disable the prompt cache for this endpoint
	DisablePromptCache param.Opt[bool] `json:"disable_prompt_cache,omitzero"`
	// Whether to disable speculative decoding for this endpoint
	DisableSpeculativeDecoding param.Opt[bool] `json:"disable_speculative_decoding,omitzero"`
	// A human-readable name for the endpoint
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// The desired state of the endpoint
	//
	// Any of "STARTED", "STOPPED".
	State EndpointNewParamsState `json:"state,omitzero"`
	paramObj
}

func (r EndpointNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The desired state of the endpoint
type EndpointNewParamsState string

const (
	EndpointNewParamsStateStarted EndpointNewParamsState = "STARTED"
	EndpointNewParamsStateStopped EndpointNewParamsState = "STOPPED"
)

type EndpointUpdateParams struct {
	// The number of minutes of inactivity after which the endpoint will be
	// automatically stopped. Set to 0 to disable automatic timeout.
	InactiveTimeout param.Opt[int64] `json:"inactive_timeout,omitzero"`
	// A human-readable name for the endpoint
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// New autoscaling configuration for the endpoint
	Autoscaling AutoscalingParam `json:"autoscaling,omitzero"`
	// The desired state of the endpoint
	//
	// Any of "STARTED", "STOPPED".
	State EndpointUpdateParamsState `json:"state,omitzero"`
	paramObj
}

func (r EndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The desired state of the endpoint
type EndpointUpdateParamsState string

const (
	EndpointUpdateParamsStateStarted EndpointUpdateParamsState = "STARTED"
	EndpointUpdateParamsStateStopped EndpointUpdateParamsState = "STOPPED"
)

type EndpointListParams struct {
	// If true, return only endpoints owned by the caller
	Mine param.Opt[bool] `query:"mine,omitzero" json:"-"`
	// Filter endpoints by type
	//
	// Any of "dedicated", "serverless".
	Type EndpointListParamsType `query:"type,omitzero" json:"-"`
	// Filter endpoints by usage type
	//
	// Any of "on-demand", "reserved".
	UsageType EndpointListParamsUsageType `query:"usage_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EndpointListParams]'s query parameters as `url.Values`.
func (r EndpointListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter endpoints by type
type EndpointListParamsType string

const (
	EndpointListParamsTypeDedicated  EndpointListParamsType = "dedicated"
	EndpointListParamsTypeServerless EndpointListParamsType = "serverless"
)

// Filter endpoints by usage type
type EndpointListParamsUsageType string

const (
	EndpointListParamsUsageTypeOnDemand EndpointListParamsUsageType = "on-demand"
	EndpointListParamsUsageTypeReserved EndpointListParamsUsageType = "reserved"
)
