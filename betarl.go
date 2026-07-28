// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"github.com/togethercomputer/together-go/option"
)

// BetaRlService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaRlService] method instead.
type BetaRlService struct {
	Options         []option.RequestOption
	Sessions        BetaRlSessionService
	Operations      BetaRlOperationService
	Checkpoints     BetaRlCheckpointService
	ModelResources  BetaRlModelResourceService
	SupportedModels BetaRlSupportedModelService
}

// NewBetaRlService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBetaRlService(opts ...option.RequestOption) (r BetaRlService) {
	r = BetaRlService{}
	r.Options = opts
	r.Sessions = NewBetaRlSessionService(opts...)
	r.Operations = NewBetaRlOperationService(opts...)
	r.Checkpoints = NewBetaRlCheckpointService(opts...)
	r.ModelResources = NewBetaRlModelResourceService(opts...)
	r.SupportedModels = NewBetaRlSupportedModelService(opts...)
	return
}
