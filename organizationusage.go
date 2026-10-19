// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"time"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/respjson"
	"github.com/togethercomputer/together-go/shared/constant"
)

// OrganizationUsageService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationUsageService] method instead.
type OrganizationUsageService struct {
	Options []option.RequestOption
	Costs   OrganizationUsageCostService
}

// NewOrganizationUsageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationUsageService(opts ...option.RequestOption) (r OrganizationUsageService) {
	r = OrganizationUsageService{}
	r.Options = opts
	r.Costs = NewOrganizationUsageCostService(opts...)
	return
}

type BillingUsageLineItem struct {
	// Resource identifiers for attribution, varying by product.
	Attributes map[string]string `json:"attributes" api:"required"`
	// Total cost for this line item in USD as a decimal string.
	Cost string `json:"cost" api:"required"`
	// Rate-determining dimensions (varies by product). Passthrough from Metronome
	// pricing group values.
	PricingDimensions map[string]string `json:"pricing_dimensions" api:"required"`
	// Metronome product name (e.g. 'Serverless Inference - Input Tokens').
	ProductName string `json:"product_name" api:"required"`
	// Total usage for the window in the product's native unit (GPU-hours, tokens, ...)
	// as a decimal string.
	Quantity string `json:"quantity" api:"required"`
	// Per-unit price in USD as a decimal string.
	UnitPrice string `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attributes        respjson.Field
		Cost              respjson.Field
		PricingDimensions respjson.Field
		ProductName       respjson.Field
		Quantity          respjson.Field
		UnitPrice         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingUsageLineItem) RawJSON() string { return r.JSON.raw }
func (r *BillingUsageLineItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingUsageReport struct {
	// Billing month (YYYY-MM).
	BillingPeriod string       `json:"billing_period" api:"required"`
	Currency      constant.Usd `json:"currency" default:"USD"`
	// Time windows containing usage line items.
	Data []BillingUsageWindow `json:"data" api:"required"`
	// Start of the earliest time window with usage in the month (UTC, ISO 8601); null
	// when the month has no usage.
	EarliestWindowStart time.Time `json:"earliest_window_start" api:"required" format:"date-time"`
	// Exclusive end of the latest time window with usage in the month (UTC, ISO 8601);
	// null when the month has no usage. Describes the whole month, not the current
	// page.
	LatestWindowEnd time.Time `json:"latest_window_end" api:"required" format:"date-time"`
	// Opaque cursor for the next page; null when this is the last page.
	NextCursor string        `json:"next_cursor" api:"required"`
	Object     constant.List `json:"object" default:"list"`
	// ID of the organization the report belongs to.
	OrganizationID string `json:"organization_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingPeriod       respjson.Field
		Currency            respjson.Field
		Data                respjson.Field
		EarliestWindowStart respjson.Field
		LatestWindowEnd     respjson.Field
		NextCursor          respjson.Field
		Object              respjson.Field
		OrganizationID      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingUsageReport) RawJSON() string { return r.JSON.raw }
func (r *BillingUsageReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingUsageWindow struct {
	// Day this window covers (UTC, YYYY-MM-DD). Present for both granularities.
	Date string `json:"date" api:"required"`
	// Window end (UTC, ISO 8601), exclusive. Start of the next day (daily) or hour
	// (hourly).
	EndTime time.Time `json:"end_time" api:"required" format:"date-time"`
	// Usage line items in this time window.
	LineItems []BillingUsageLineItem `json:"line_items" api:"required"`
	// Window start (UTC, ISO 8601). Start of the day (daily) or hour (hourly).
	StartTime time.Time `json:"start_time" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		EndTime     respjson.Field
		LineItems   respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingUsageWindow) RawJSON() string { return r.JSON.raw }
func (r *BillingUsageWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
