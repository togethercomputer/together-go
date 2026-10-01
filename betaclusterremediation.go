// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	shimjson "github.com/togethercomputer/together-go/internal/encoding/json"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaClusterRemediationService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaClusterRemediationService] method instead.
type BetaClusterRemediationService struct {
	Options []option.RequestOption
}

// NewBetaClusterRemediationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaClusterRemediationService(opts ...option.RequestOption) (r BetaClusterRemediationService) {
	r = BetaClusterRemediationService{}
	r.Options = opts
	return
}

// Creates a new remediation for an instance.
//
// Remediations created via the API goes directly to PENDING state.
//
// Our system may trigger automated remediations that require approval. These
// remediations are created with PENDING_APPROVAL state. The user must call
// /approve to start the actual remediation process. These operations can also be
// rejected by calling /reject.
func (r *BetaClusterRemediationService) New(ctx context.Context, instanceID string, params BetaClusterRemediationNewParams, opts ...option.RequestOption) (res *Remediation, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ClusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/%s/instances/%s/remediations", params.ClusterID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve the status of a specific remdiation on a specific instance in a
// specific cluster.
func (r *BetaClusterRemediationService) Get(ctx context.Context, remediationID string, query BetaClusterRemediationGetParams, opts ...option.RequestOption) (res *Remediation, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ClusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if query.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if remediationID == "" {
		err = errors.New("missing required remediation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/%s/instances/%s/remediations/%s", query.ClusterID, query.InstanceID, remediationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists remediations for an instance or cluster.
func (r *BetaClusterRemediationService) List(ctx context.Context, instanceID string, params BetaClusterRemediationListParams, opts ...option.RequestOption) (res *BetaClusterRemediationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ClusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/%s/instances/%s/remediations", params.ClusterID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Approves a pending remediation.
//
// Only remediations with state PENDING_APPROVAL can be approved.
//
// On APPROVE: state changes to PENDING and the remediation process begins. The
// reviewed_by, review_time, and review_comment fields are populated on the
// remediation after approval.
func (r *BetaClusterRemediationService) Approve(ctx context.Context, remediationID string, params BetaClusterRemediationApproveParams, opts ...option.RequestOption) (res *Remediation, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ClusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if remediationID == "" {
		err = errors.New("missing required remediation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/%s/instances/%s/remediations/%s/approve", params.ClusterID, params.InstanceID, remediationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Cancels a pending remediation.
//
// Only remediations in PENDING_APPROVAL or PENDING state can be cancelled.
func (r *BetaClusterRemediationService) Cancel(ctx context.Context, remediationID string, body BetaClusterRemediationCancelParams, opts ...option.RequestOption) (res *Remediation, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ClusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if body.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if remediationID == "" {
		err = errors.New("missing required remediation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/%s/instances/%s/remediations/%s/cancel", body.ClusterID, body.InstanceID, remediationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Rejects a pending remediation.
//
// Only remediations with state PENDING_APPROVAL can be rejected.
//
// On REJECT: state changes to CANCELLED. The reviewed_by, review_time, and
// review_comment fields are populated on the remediation after rejection.
func (r *BetaClusterRemediationService) Reject(ctx context.Context, remediationID string, params BetaClusterRemediationRejectParams, opts ...option.RequestOption) (res *Remediation, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ClusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if remediationID == "" {
		err = errors.New("missing required remediation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/clusters/%s/instances/%s/remediations/%s/reject", params.ClusterID, params.InstanceID, remediationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// ListRemediationsResponse is the response for ListRemediations.
type BetaClusterRemediationListResponse struct {
	// Indicates if there are more results available.
	HasNext bool `json:"has_next" api:"required"`
	// Token for the next page.
	NextPageToken string `json:"next_page_token" api:"required"`
	// The list of remediations.
	Remediations []Remediation `json:"remediations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNext       respjson.Field
		NextPageToken respjson.Field
		Remediations  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterRemediationListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterRemediationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterRemediationNewParams struct {
	ClusterID string `path:"cluster_id" api:"required" json:"-"`
	// Remediation represents a node remediation request for an instance. An instance
	// can have multiple remediations over time (e.g., failed attempts followed by
	// retries).
	Remediation RemediationParam
	// Client-specified ID for idempotency.
	RemediationID param.Opt[string] `query:"remediation_id,omitzero" json:"-"`
	paramObj
}

func (r BetaClusterRemediationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Remediation)
}
func (r *BetaClusterRemediationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaClusterRemediationNewParams]'s query parameters as
// `url.Values`.
func (r BetaClusterRemediationNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaClusterRemediationGetParams struct {
	ClusterID  string `path:"cluster_id" api:"required" json:"-"`
	InstanceID string `path:"instance_id" api:"required" json:"-"`
	paramObj
}

type BetaClusterRemediationListParams struct {
	ClusterID string `path:"cluster_id" api:"required" json:"-"`
	// Order by expression.
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Maximum results to return.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Pagination token from previous request.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	// Filter by remediation mode(s). Returns remediations matching any of the
	// specified modes.
	//
	// Any of "REMEDIATION_MODE_VM_ONLY", "REMEDIATION_MODE_HOST_AWARE",
	// "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT", "REMEDIATION_MODE_REBOOT_VM",
	// "REMEDIATION_MODE_HOST_POWER_CYCLE".
	Mode []string `query:"mode,omitzero" json:"-"`
	// Filter by state(s). Returns remediations matching any of the specified states.
	//
	//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
	//   - `PENDING`: Approved and queued for processing.
	//   - `RUNNING`: Actively being processed.
	//   - `SUCCEEDED`: Successfully completed.
	//   - `FAILED`: Failed with an error.
	//   - `CANCELLED`: Cancelled by user or system.
	//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
	//     processing.
	//   - `QUARANTINING`: Cordoning or preparing the host before remediation.
	//   - `QUARANTINED`: Host has been cordoned or isolated for remediation.
	//
	// Any of "PENDING_APPROVAL", "PENDING", "RUNNING", "SUCCEEDED", "FAILED",
	// "CANCELLED", "AUTO_RESOLVED", "QUARANTINING", "QUARANTINED".
	State []string `query:"state,omitzero" json:"-"`
	// Filter by trigger type(s). Returns remediations matching any of the specified
	// triggers.
	//
	// Any of "REMEDIATION_TRIGGER_MANUAL", "REMEDIATION_TRIGGER_AUTOMATED".
	Trigger []string `query:"trigger,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaClusterRemediationListParams]'s query parameters as
// `url.Values`.
func (r BetaClusterRemediationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaClusterRemediationApproveParams struct {
	ClusterID  string `path:"cluster_id" api:"required" json:"-"`
	InstanceID string `path:"instance_id" api:"required" json:"-"`
	// Approval comment explaining the decision.
	Comment param.Opt[string] `json:"comment,omitzero"`
	// Remediation mode to use after approval. When omitted, the remediation keeps its
	// existing mode.
	//
	//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
	//     available host.
	//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
	//     provisions a new one on a different host.
	//   - `REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT`: Evicts the VM without
	//     provisioning a replacement.
	//   - `REMEDIATION_MODE_REBOOT_VM`: Reboots the VM in place.
	//   - `REMEDIATION_MODE_HOST_POWER_CYCLE`: Power-cycles the bare-metal host after
	//     cordoning it. This mode cannot be set as an approval override; create a host
	//     power-cycle remediation directly.
	//
	// Any of "REMEDIATION_MODE_VM_ONLY", "REMEDIATION_MODE_HOST_AWARE",
	// "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT", "REMEDIATION_MODE_REBOOT_VM",
	// "REMEDIATION_MODE_HOST_POWER_CYCLE".
	Mode BetaClusterRemediationApproveParamsMode `json:"mode,omitzero"`
	paramObj
}

func (r BetaClusterRemediationApproveParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterRemediationApproveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterRemediationApproveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remediation mode to use after approval. When omitted, the remediation keeps its
// existing mode.
//
//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
//     available host.
//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
//     provisions a new one on a different host.
//   - `REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT`: Evicts the VM without
//     provisioning a replacement.
//   - `REMEDIATION_MODE_REBOOT_VM`: Reboots the VM in place.
//   - `REMEDIATION_MODE_HOST_POWER_CYCLE`: Power-cycles the bare-metal host after
//     cordoning it. This mode cannot be set as an approval override; create a host
//     power-cycle remediation directly.
type BetaClusterRemediationApproveParamsMode string

const (
	BetaClusterRemediationApproveParamsModeRemediationModeVmOnly                  BetaClusterRemediationApproveParamsMode = "REMEDIATION_MODE_VM_ONLY"
	BetaClusterRemediationApproveParamsModeRemediationModeHostAware               BetaClusterRemediationApproveParamsMode = "REMEDIATION_MODE_HOST_AWARE"
	BetaClusterRemediationApproveParamsModeRemediationModeEvictWithoutReplacement BetaClusterRemediationApproveParamsMode = "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT"
	BetaClusterRemediationApproveParamsModeRemediationModeRebootVm                BetaClusterRemediationApproveParamsMode = "REMEDIATION_MODE_REBOOT_VM"
	BetaClusterRemediationApproveParamsModeRemediationModeHostPowerCycle          BetaClusterRemediationApproveParamsMode = "REMEDIATION_MODE_HOST_POWER_CYCLE"
)

type BetaClusterRemediationCancelParams struct {
	// The cluster ID.
	ClusterID string `path:"cluster_id" api:"required" json:"-"`
	// The instance ID.
	InstanceID string `path:"instance_id" api:"required" json:"-"`
	paramObj
}

type BetaClusterRemediationRejectParams struct {
	ClusterID  string `path:"cluster_id" api:"required" json:"-"`
	InstanceID string `path:"instance_id" api:"required" json:"-"`
	// Comment explaining the action.
	Comment param.Opt[string] `json:"comment,omitzero"`
	paramObj
}

func (r BetaClusterRemediationRejectParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterRemediationRejectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterRemediationRejectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
