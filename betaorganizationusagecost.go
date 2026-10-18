// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/pagination"
	"github.com/togethercomputer/together-go/packages/param"
)

// BetaOrganizationUsageCostService contains methods and other services that help
// with interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaOrganizationUsageCostService] method instead.
type BetaOrganizationUsageCostService struct {
	Options []option.RequestOption
}

// NewBetaOrganizationUsageCostService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBetaOrganizationUsageCostService(opts ...option.RequestOption) (r BetaOrganizationUsageCostService) {
	r = BetaOrganizationUsageCostService{}
	r.Options = opts
	return
}

// Returns an organization's billing usage for a month as cost-annotated line items
// grouped into time windows. Finalized windows are returned through the end of
// yesterday at daily granularity, or the last completed hour at hourly granularity
// (UTC).
func (r *BetaOrganizationUsageCostService) List(ctx context.Context, query BetaOrganizationUsageCostListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[BillingUsageWindow], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/billing/usage"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns an organization's billing usage for a month as cost-annotated line items
// grouped into time windows. Finalized windows are returned through the end of
// yesterday at daily granularity, or the last completed hour at hourly granularity
// (UTC).
func (r *BetaOrganizationUsageCostService) ListAutoPaging(ctx context.Context, query BetaOrganizationUsageCostListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[BillingUsageWindow] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

type BetaOrganizationUsageCostListParams struct {
	// Opaque cursor from a previous response's `next_cursor`. Only valid for the month
	// and granularity it was issued for. Current-month data is a moving snapshot: each
	// page reflects data as of its request, so newly completed windows may appear at
	// the end of the sequence.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of time windows per page (max 1000).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Billing month in YYYY-MM format. Defaults to the current month, which returns
	// data through yesterday (daily) or the last completed hour (hourly).
	Month param.Opt[string] `query:"month,omitzero" json:"-"`
	// Organization to report on. Usage is always reported for the organization that
	// owns the API key, so this is optional; supplying an ID that is not that
	// organization returns 403.
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" json:"-"`
	// Time window size for rows. 'hour' returns ~24x more rows than 'day'.
	//
	// Any of "day", "hour".
	Granularity BetaOrganizationUsageCostListParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaOrganizationUsageCostListParams]'s query parameters as
// `url.Values`.
func (r BetaOrganizationUsageCostListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time window size for rows. 'hour' returns ~24x more rows than 'day'.
type BetaOrganizationUsageCostListParamsGranularity string

const (
	BetaOrganizationUsageCostListParamsGranularityDay  BetaOrganizationUsageCostListParamsGranularity = "day"
	BetaOrganizationUsageCostListParamsGranularityHour BetaOrganizationUsageCostListParamsGranularity = "hour"
)
