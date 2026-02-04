// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
)

// RlTrainingSessionOperationService contains methods and other services that help
// with interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRlTrainingSessionOperationService] method instead.
type RlTrainingSessionOperationService struct {
	Options []option.RequestOption
}

// NewRlTrainingSessionOperationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRlTrainingSessionOperationService(opts ...option.RequestOption) (r RlTrainingSessionOperationService) {
	r = RlTrainingSessionOperationService{}
	r.Options = opts
	return
}

// Retrieves the current status and result of a forward-backward operation.
func (r *RlTrainingSessionOperationService) GetForwardBackward(ctx context.Context, operationID string, query RlTrainingSessionOperationGetForwardBackwardParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/forward-backward/%s", query.SessionID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Retrieves the current status and result of an optim-step operation.
func (r *RlTrainingSessionOperationService) GetOptimStep(ctx context.Context, operationID string, query RlTrainingSessionOperationGetOptimStepParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("rl/training-sessions/%s/operations/optim-step/%s", query.SessionID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type RlTrainingSessionOperationGetForwardBackwardParams struct {
	SessionID string `path:"session_id,required" json:"-"`
	paramObj
}

type RlTrainingSessionOperationGetOptimStepParams struct {
	SessionID string `path:"session_id,required" json:"-"`
	paramObj
}
