// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
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

// HardwareService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHardwareService] method instead.
type HardwareService struct {
	Options []option.RequestOption
}

// NewHardwareService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHardwareService(opts ...option.RequestOption) (r HardwareService) {
	r = HardwareService{}
	r.Options = opts
	return
}

// Returns a list of available hardware configurations for deploying models. When a
// model parameter is provided, it returns only hardware configurations compatible
// with that model, including their current availability status.
func (r *HardwareService) List(ctx context.Context, query HardwareListParams, opts ...option.RequestOption) (res *HardwareListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "hardware"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type HardwareListResponse struct {
	Data []HardwareListResponseData `json:"data,required"`
	// Any of "list".
	Object HardwareListResponseObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HardwareListResponse) RawJSON() string { return r.JSON.raw }
func (r *HardwareListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hardware configuration details with optional availability status
type HardwareListResponseData struct {
	// Unique identifier for the hardware configuration
	ID string `json:"id,required"`
	// Any of "hardware".
	Object string `json:"object,required"`
	// Pricing details for using an endpoint
	Pricing HardwareListResponseDataPricing `json:"pricing,required"`
	// Detailed specifications of a hardware configuration
	Specs HardwareListResponseDataSpecs `json:"specs,required"`
	// Timestamp of when the hardware status was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Indicates the current availability status of a hardware configuration
	Availability HardwareListResponseDataAvailability `json:"availability"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Object       respjson.Field
		Pricing      respjson.Field
		Specs        respjson.Field
		UpdatedAt    respjson.Field
		Availability respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HardwareListResponseData) RawJSON() string { return r.JSON.raw }
func (r *HardwareListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pricing details for using an endpoint
type HardwareListResponseDataPricing struct {
	// Cost per minute of endpoint uptime in cents
	CentsPerMinute float64 `json:"cents_per_minute,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CentsPerMinute respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HardwareListResponseDataPricing) RawJSON() string { return r.JSON.raw }
func (r *HardwareListResponseDataPricing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed specifications of a hardware configuration
type HardwareListResponseDataSpecs struct {
	// Number of GPUs in this configuration
	GPUCount int64 `json:"gpu_count,required"`
	// The GPU interconnect technology
	GPULink string `json:"gpu_link,required"`
	// Amount of GPU memory in GB
	GPUMemory float64 `json:"gpu_memory,required"`
	// The type/model of GPU
	GPUType string `json:"gpu_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GPUCount    respjson.Field
		GPULink     respjson.Field
		GPUMemory   respjson.Field
		GPUType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HardwareListResponseDataSpecs) RawJSON() string { return r.JSON.raw }
func (r *HardwareListResponseDataSpecs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the current availability status of a hardware configuration
type HardwareListResponseDataAvailability struct {
	// The availability status of the hardware configuration
	//
	// Any of "available", "unavailable", "insufficient".
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HardwareListResponseDataAvailability) RawJSON() string { return r.JSON.raw }
func (r *HardwareListResponseDataAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HardwareListResponseObject string

const (
	HardwareListResponseObjectList HardwareListResponseObject = "list"
)

type HardwareListParams struct {
	// Filter hardware configurations by model compatibility. When provided, the
	// response includes availability status for each compatible configuration.
	Model param.Opt[string] `query:"model,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HardwareListParams]'s query parameters as `url.Values`.
func (r HardwareListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
