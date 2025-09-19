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
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
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
func NewHardwareService(opts ...option.RequestOption) (r *HardwareService) {
	r = &HardwareService{}
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
	Data   []HardwareListResponseData `json:"data,required"`
	Object HardwareListResponseObject `json:"object,required"`
	JSON   hardwareListResponseJSON   `json:"-"`
}

// hardwareListResponseJSON contains the JSON metadata for the struct
// [HardwareListResponse]
type hardwareListResponseJSON struct {
	Data        apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HardwareListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hardwareListResponseJSON) RawJSON() string {
	return r.raw
}

// Hardware configuration details with optional availability status
type HardwareListResponseData struct {
	// Unique identifier for the hardware configuration
	ID     string                         `json:"id,required"`
	Object HardwareListResponseDataObject `json:"object,required"`
	// Pricing details for using an endpoint
	Pricing HardwareListResponseDataPricing `json:"pricing,required"`
	// Detailed specifications of a hardware configuration
	Specs HardwareListResponseDataSpecs `json:"specs,required"`
	// Timestamp of when the hardware status was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Indicates the current availability status of a hardware configuration
	Availability HardwareListResponseDataAvailability `json:"availability"`
	JSON         hardwareListResponseDataJSON         `json:"-"`
}

// hardwareListResponseDataJSON contains the JSON metadata for the struct
// [HardwareListResponseData]
type hardwareListResponseDataJSON struct {
	ID           apijson.Field
	Object       apijson.Field
	Pricing      apijson.Field
	Specs        apijson.Field
	UpdatedAt    apijson.Field
	Availability apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HardwareListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hardwareListResponseDataJSON) RawJSON() string {
	return r.raw
}

type HardwareListResponseDataObject string

const (
	HardwareListResponseDataObjectHardware HardwareListResponseDataObject = "hardware"
)

func (r HardwareListResponseDataObject) IsKnown() bool {
	switch r {
	case HardwareListResponseDataObjectHardware:
		return true
	}
	return false
}

// Pricing details for using an endpoint
type HardwareListResponseDataPricing struct {
	// Cost per minute of endpoint uptime in cents
	CentsPerMinute float64                             `json:"cents_per_minute,required"`
	JSON           hardwareListResponseDataPricingJSON `json:"-"`
}

// hardwareListResponseDataPricingJSON contains the JSON metadata for the struct
// [HardwareListResponseDataPricing]
type hardwareListResponseDataPricingJSON struct {
	CentsPerMinute apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HardwareListResponseDataPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hardwareListResponseDataPricingJSON) RawJSON() string {
	return r.raw
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
	GPUType string                            `json:"gpu_type,required"`
	JSON    hardwareListResponseDataSpecsJSON `json:"-"`
}

// hardwareListResponseDataSpecsJSON contains the JSON metadata for the struct
// [HardwareListResponseDataSpecs]
type hardwareListResponseDataSpecsJSON struct {
	GPUCount    apijson.Field
	GPULink     apijson.Field
	GPUMemory   apijson.Field
	GPUType     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HardwareListResponseDataSpecs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hardwareListResponseDataSpecsJSON) RawJSON() string {
	return r.raw
}

// Indicates the current availability status of a hardware configuration
type HardwareListResponseDataAvailability struct {
	// The availability status of the hardware configuration
	Status HardwareListResponseDataAvailabilityStatus `json:"status,required"`
	JSON   hardwareListResponseDataAvailabilityJSON   `json:"-"`
}

// hardwareListResponseDataAvailabilityJSON contains the JSON metadata for the
// struct [HardwareListResponseDataAvailability]
type hardwareListResponseDataAvailabilityJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HardwareListResponseDataAvailability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hardwareListResponseDataAvailabilityJSON) RawJSON() string {
	return r.raw
}

// The availability status of the hardware configuration
type HardwareListResponseDataAvailabilityStatus string

const (
	HardwareListResponseDataAvailabilityStatusAvailable    HardwareListResponseDataAvailabilityStatus = "available"
	HardwareListResponseDataAvailabilityStatusUnavailable  HardwareListResponseDataAvailabilityStatus = "unavailable"
	HardwareListResponseDataAvailabilityStatusInsufficient HardwareListResponseDataAvailabilityStatus = "insufficient"
)

func (r HardwareListResponseDataAvailabilityStatus) IsKnown() bool {
	switch r {
	case HardwareListResponseDataAvailabilityStatusAvailable, HardwareListResponseDataAvailabilityStatusUnavailable, HardwareListResponseDataAvailabilityStatusInsufficient:
		return true
	}
	return false
}

type HardwareListResponseObject string

const (
	HardwareListResponseObjectList HardwareListResponseObject = "list"
)

func (r HardwareListResponseObject) IsKnown() bool {
	switch r {
	case HardwareListResponseObjectList:
		return true
	}
	return false
}

type HardwareListParams struct {
	// Filter hardware configurations by model compatibility. When provided, the
	// response includes availability status for each compatible configuration.
	Model param.Field[string] `query:"model"`
}

// URLQuery serializes [HardwareListParams]'s query parameters as `url.Values`.
func (r HardwareListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
