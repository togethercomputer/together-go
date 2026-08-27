// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"github.com/togethercomputer/together-go/option"
)

// BetaOrganizationService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaOrganizationService] method instead.
type BetaOrganizationService struct {
	Options []option.RequestOption
	Usage   BetaOrganizationUsageService
}

// NewBetaOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaOrganizationService(opts ...option.RequestOption) (r BetaOrganizationService) {
	r = BetaOrganizationService{}
	r.Options = opts
	r.Usage = NewBetaOrganizationUsageService(opts...)
	return
}
