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
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
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
func NewEndpointService(opts ...option.RequestOption) (r *EndpointService) {
	r = &EndpointService{}
	r.Options = opts
	return
}

// Creates a new dedicated endpoint for serving models. The endpoint will
// automatically start after creation. You can deploy any supported model on
// hardware configurations that meet the model's requirements.
func (r *EndpointService) New(ctx context.Context, body EndpointNewParams, opts ...option.RequestOption) (res *EndpointNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves details about a specific endpoint, including its current state,
// configuration, and scaling settings.
func (r *EndpointService) Get(ctx context.Context, endpointID string, opts ...option.RequestOption) (res *EndpointGetResponse, err error) {
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
func (r *EndpointService) Update(ctx context.Context, endpointID string, body EndpointUpdateParams, opts ...option.RequestOption) (res *EndpointUpdateResponse, err error) {
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
	MinReplicas int64           `json:"min_replicas,required"`
	JSON        autoscalingJSON `json:"-"`
}

// autoscalingJSON contains the JSON metadata for the struct [Autoscaling]
type autoscalingJSON struct {
	MaxReplicas apijson.Field
	MinReplicas apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Autoscaling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r autoscalingJSON) RawJSON() string {
	return r.raw
}

// Configuration for automatic scaling of replicas based on demand.
type AutoscalingParam struct {
	// The maximum number of replicas to scale up to under load
	MaxReplicas param.Field[int64] `json:"max_replicas,required"`
	// The minimum number of replicas to maintain, even when there is no load
	MinReplicas param.Field[int64] `json:"min_replicas,required"`
}

func (r AutoscalingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details about a dedicated endpoint deployment
type EndpointNewResponse struct {
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
	Object EndpointNewResponseObject `json:"object,required"`
	// The owner of this endpoint
	Owner string `json:"owner,required"`
	// Current state of the endpoint
	State EndpointNewResponseState `json:"state,required"`
	// The type of endpoint
	Type EndpointNewResponseType `json:"type,required"`
	JSON endpointNewResponseJSON `json:"-"`
}

// endpointNewResponseJSON contains the JSON metadata for the struct
// [EndpointNewResponse]
type endpointNewResponseJSON struct {
	ID          apijson.Field
	Autoscaling apijson.Field
	CreatedAt   apijson.Field
	DisplayName apijson.Field
	Hardware    apijson.Field
	Model       apijson.Field
	Name        apijson.Field
	Object      apijson.Field
	Owner       apijson.Field
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointNewResponseJSON) RawJSON() string {
	return r.raw
}

// The type of object
type EndpointNewResponseObject string

const (
	EndpointNewResponseObjectEndpoint EndpointNewResponseObject = "endpoint"
)

func (r EndpointNewResponseObject) IsKnown() bool {
	switch r {
	case EndpointNewResponseObjectEndpoint:
		return true
	}
	return false
}

// Current state of the endpoint
type EndpointNewResponseState string

const (
	EndpointNewResponseStatePending  EndpointNewResponseState = "PENDING"
	EndpointNewResponseStateStarting EndpointNewResponseState = "STARTING"
	EndpointNewResponseStateStarted  EndpointNewResponseState = "STARTED"
	EndpointNewResponseStateStopping EndpointNewResponseState = "STOPPING"
	EndpointNewResponseStateStopped  EndpointNewResponseState = "STOPPED"
	EndpointNewResponseStateError    EndpointNewResponseState = "ERROR"
)

func (r EndpointNewResponseState) IsKnown() bool {
	switch r {
	case EndpointNewResponseStatePending, EndpointNewResponseStateStarting, EndpointNewResponseStateStarted, EndpointNewResponseStateStopping, EndpointNewResponseStateStopped, EndpointNewResponseStateError:
		return true
	}
	return false
}

// The type of endpoint
type EndpointNewResponseType string

const (
	EndpointNewResponseTypeDedicated EndpointNewResponseType = "dedicated"
)

func (r EndpointNewResponseType) IsKnown() bool {
	switch r {
	case EndpointNewResponseTypeDedicated:
		return true
	}
	return false
}

// Details about a dedicated endpoint deployment
type EndpointGetResponse struct {
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
	Object EndpointGetResponseObject `json:"object,required"`
	// The owner of this endpoint
	Owner string `json:"owner,required"`
	// Current state of the endpoint
	State EndpointGetResponseState `json:"state,required"`
	// The type of endpoint
	Type EndpointGetResponseType `json:"type,required"`
	JSON endpointGetResponseJSON `json:"-"`
}

// endpointGetResponseJSON contains the JSON metadata for the struct
// [EndpointGetResponse]
type endpointGetResponseJSON struct {
	ID          apijson.Field
	Autoscaling apijson.Field
	CreatedAt   apijson.Field
	DisplayName apijson.Field
	Hardware    apijson.Field
	Model       apijson.Field
	Name        apijson.Field
	Object      apijson.Field
	Owner       apijson.Field
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointGetResponseJSON) RawJSON() string {
	return r.raw
}

// The type of object
type EndpointGetResponseObject string

const (
	EndpointGetResponseObjectEndpoint EndpointGetResponseObject = "endpoint"
)

func (r EndpointGetResponseObject) IsKnown() bool {
	switch r {
	case EndpointGetResponseObjectEndpoint:
		return true
	}
	return false
}

// Current state of the endpoint
type EndpointGetResponseState string

const (
	EndpointGetResponseStatePending  EndpointGetResponseState = "PENDING"
	EndpointGetResponseStateStarting EndpointGetResponseState = "STARTING"
	EndpointGetResponseStateStarted  EndpointGetResponseState = "STARTED"
	EndpointGetResponseStateStopping EndpointGetResponseState = "STOPPING"
	EndpointGetResponseStateStopped  EndpointGetResponseState = "STOPPED"
	EndpointGetResponseStateError    EndpointGetResponseState = "ERROR"
)

func (r EndpointGetResponseState) IsKnown() bool {
	switch r {
	case EndpointGetResponseStatePending, EndpointGetResponseStateStarting, EndpointGetResponseStateStarted, EndpointGetResponseStateStopping, EndpointGetResponseStateStopped, EndpointGetResponseStateError:
		return true
	}
	return false
}

// The type of endpoint
type EndpointGetResponseType string

const (
	EndpointGetResponseTypeDedicated EndpointGetResponseType = "dedicated"
)

func (r EndpointGetResponseType) IsKnown() bool {
	switch r {
	case EndpointGetResponseTypeDedicated:
		return true
	}
	return false
}

// Details about a dedicated endpoint deployment
type EndpointUpdateResponse struct {
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
	Object EndpointUpdateResponseObject `json:"object,required"`
	// The owner of this endpoint
	Owner string `json:"owner,required"`
	// Current state of the endpoint
	State EndpointUpdateResponseState `json:"state,required"`
	// The type of endpoint
	Type EndpointUpdateResponseType `json:"type,required"`
	JSON endpointUpdateResponseJSON `json:"-"`
}

// endpointUpdateResponseJSON contains the JSON metadata for the struct
// [EndpointUpdateResponse]
type endpointUpdateResponseJSON struct {
	ID          apijson.Field
	Autoscaling apijson.Field
	CreatedAt   apijson.Field
	DisplayName apijson.Field
	Hardware    apijson.Field
	Model       apijson.Field
	Name        apijson.Field
	Object      apijson.Field
	Owner       apijson.Field
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The type of object
type EndpointUpdateResponseObject string

const (
	EndpointUpdateResponseObjectEndpoint EndpointUpdateResponseObject = "endpoint"
)

func (r EndpointUpdateResponseObject) IsKnown() bool {
	switch r {
	case EndpointUpdateResponseObjectEndpoint:
		return true
	}
	return false
}

// Current state of the endpoint
type EndpointUpdateResponseState string

const (
	EndpointUpdateResponseStatePending  EndpointUpdateResponseState = "PENDING"
	EndpointUpdateResponseStateStarting EndpointUpdateResponseState = "STARTING"
	EndpointUpdateResponseStateStarted  EndpointUpdateResponseState = "STARTED"
	EndpointUpdateResponseStateStopping EndpointUpdateResponseState = "STOPPING"
	EndpointUpdateResponseStateStopped  EndpointUpdateResponseState = "STOPPED"
	EndpointUpdateResponseStateError    EndpointUpdateResponseState = "ERROR"
)

func (r EndpointUpdateResponseState) IsKnown() bool {
	switch r {
	case EndpointUpdateResponseStatePending, EndpointUpdateResponseStateStarting, EndpointUpdateResponseStateStarted, EndpointUpdateResponseStateStopping, EndpointUpdateResponseStateStopped, EndpointUpdateResponseStateError:
		return true
	}
	return false
}

// The type of endpoint
type EndpointUpdateResponseType string

const (
	EndpointUpdateResponseTypeDedicated EndpointUpdateResponseType = "dedicated"
)

func (r EndpointUpdateResponseType) IsKnown() bool {
	switch r {
	case EndpointUpdateResponseTypeDedicated:
		return true
	}
	return false
}

type EndpointListResponse struct {
	Data   []EndpointListResponseData `json:"data,required"`
	Object EndpointListResponseObject `json:"object,required"`
	JSON   endpointListResponseJSON   `json:"-"`
}

// endpointListResponseJSON contains the JSON metadata for the struct
// [EndpointListResponse]
type endpointListResponseJSON struct {
	Data        apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointListResponseJSON) RawJSON() string {
	return r.raw
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
	Object EndpointListResponseDataObject `json:"object,required"`
	// The owner of this endpoint
	Owner string `json:"owner,required"`
	// Current state of the endpoint
	State EndpointListResponseDataState `json:"state,required"`
	// The type of endpoint
	Type EndpointListResponseDataType `json:"type,required"`
	JSON endpointListResponseDataJSON `json:"-"`
}

// endpointListResponseDataJSON contains the JSON metadata for the struct
// [EndpointListResponseData]
type endpointListResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Model       apijson.Field
	Name        apijson.Field
	Object      apijson.Field
	Owner       apijson.Field
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointListResponseDataJSON) RawJSON() string {
	return r.raw
}

// The type of object
type EndpointListResponseDataObject string

const (
	EndpointListResponseDataObjectEndpoint EndpointListResponseDataObject = "endpoint"
)

func (r EndpointListResponseDataObject) IsKnown() bool {
	switch r {
	case EndpointListResponseDataObjectEndpoint:
		return true
	}
	return false
}

// Current state of the endpoint
type EndpointListResponseDataState string

const (
	EndpointListResponseDataStatePending  EndpointListResponseDataState = "PENDING"
	EndpointListResponseDataStateStarting EndpointListResponseDataState = "STARTING"
	EndpointListResponseDataStateStarted  EndpointListResponseDataState = "STARTED"
	EndpointListResponseDataStateStopping EndpointListResponseDataState = "STOPPING"
	EndpointListResponseDataStateStopped  EndpointListResponseDataState = "STOPPED"
	EndpointListResponseDataStateError    EndpointListResponseDataState = "ERROR"
)

func (r EndpointListResponseDataState) IsKnown() bool {
	switch r {
	case EndpointListResponseDataStatePending, EndpointListResponseDataStateStarting, EndpointListResponseDataStateStarted, EndpointListResponseDataStateStopping, EndpointListResponseDataStateStopped, EndpointListResponseDataStateError:
		return true
	}
	return false
}

// The type of endpoint
type EndpointListResponseDataType string

const (
	EndpointListResponseDataTypeServerless EndpointListResponseDataType = "serverless"
	EndpointListResponseDataTypeDedicated  EndpointListResponseDataType = "dedicated"
)

func (r EndpointListResponseDataType) IsKnown() bool {
	switch r {
	case EndpointListResponseDataTypeServerless, EndpointListResponseDataTypeDedicated:
		return true
	}
	return false
}

type EndpointListResponseObject string

const (
	EndpointListResponseObjectList EndpointListResponseObject = "list"
)

func (r EndpointListResponseObject) IsKnown() bool {
	switch r {
	case EndpointListResponseObjectList:
		return true
	}
	return false
}

type EndpointNewParams struct {
	// Configuration for automatic scaling of the endpoint
	Autoscaling param.Field[AutoscalingParam] `json:"autoscaling,required"`
	// The hardware configuration to use for this endpoint
	Hardware param.Field[string] `json:"hardware,required"`
	// The model to deploy on this endpoint
	Model param.Field[string] `json:"model,required"`
	// Whether to disable the prompt cache for this endpoint
	DisablePromptCache param.Field[bool] `json:"disable_prompt_cache"`
	// Whether to disable speculative decoding for this endpoint
	DisableSpeculativeDecoding param.Field[bool] `json:"disable_speculative_decoding"`
	// A human-readable name for the endpoint
	DisplayName param.Field[string] `json:"display_name"`
	// The number of minutes of inactivity after which the endpoint will be
	// automatically stopped. Set to null, omit or set to 0 to disable automatic
	// timeout.
	InactiveTimeout param.Field[int64] `json:"inactive_timeout"`
	// The desired state of the endpoint
	State param.Field[EndpointNewParamsState] `json:"state"`
}

func (r EndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The desired state of the endpoint
type EndpointNewParamsState string

const (
	EndpointNewParamsStateStarted EndpointNewParamsState = "STARTED"
	EndpointNewParamsStateStopped EndpointNewParamsState = "STOPPED"
)

func (r EndpointNewParamsState) IsKnown() bool {
	switch r {
	case EndpointNewParamsStateStarted, EndpointNewParamsStateStopped:
		return true
	}
	return false
}

type EndpointUpdateParams struct {
	// New autoscaling configuration for the endpoint
	Autoscaling param.Field[AutoscalingParam] `json:"autoscaling"`
	// A human-readable name for the endpoint
	DisplayName param.Field[string] `json:"display_name"`
	// The number of minutes of inactivity after which the endpoint will be
	// automatically stopped. Set to 0 to disable automatic timeout.
	InactiveTimeout param.Field[int64] `json:"inactive_timeout"`
	// The desired state of the endpoint
	State param.Field[EndpointUpdateParamsState] `json:"state"`
}

func (r EndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The desired state of the endpoint
type EndpointUpdateParamsState string

const (
	EndpointUpdateParamsStateStarted EndpointUpdateParamsState = "STARTED"
	EndpointUpdateParamsStateStopped EndpointUpdateParamsState = "STOPPED"
)

func (r EndpointUpdateParamsState) IsKnown() bool {
	switch r {
	case EndpointUpdateParamsStateStarted, EndpointUpdateParamsStateStopped:
		return true
	}
	return false
}

type EndpointListParams struct {
	// Filter endpoints by type
	Type param.Field[EndpointListParamsType] `query:"type"`
}

// URLQuery serializes [EndpointListParams]'s query parameters as `url.Values`.
func (r EndpointListParams) URLQuery() (v url.Values) {
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

func (r EndpointListParamsType) IsKnown() bool {
	switch r {
	case EndpointListParamsTypeDedicated, EndpointListParamsTypeServerless:
		return true
	}
	return false
}
