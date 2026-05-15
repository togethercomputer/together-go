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
// If mode is unspecified, it defaults to VM_ONLY. If trigger is unspecified, it
// defaults to MANUAL.
//
// For MANUAL triggers: The remediation goes directly to PENDING state.
//
// For AUTOMATED triggers: The remediation is created with PENDING_APPROVAL state.
// The caller must then use ApproveRemediation to start the remediation process.
func (r *BetaClusterRemediationService) New(ctx context.Context, instanceID string, params BetaClusterRemediationNewParams, opts ...option.RequestOption) (res *BetaClusterRemediationNewResponse, err error) {
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

// Lists remediations for an instance or cluster. Use instances/- as wildcard to
// list all remediations in a cluster.
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
func (r *BetaClusterRemediationService) Approve(ctx context.Context, remediationID string, params BetaClusterRemediationApproveParams, opts ...option.RequestOption) (res *BetaClusterRemediationApproveResponse, err error) {
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
func (r *BetaClusterRemediationService) Cancel(ctx context.Context, remediationID string, body BetaClusterRemediationCancelParams, opts ...option.RequestOption) (res *BetaClusterRemediationCancelResponse, err error) {
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
func (r *BetaClusterRemediationService) Reject(ctx context.Context, remediationID string, params BetaClusterRemediationRejectParams, opts ...option.RequestOption) (res *BetaClusterRemediationRejectResponse, err error) {
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

// Remediation represents a node remediation request for an instance. An instance
// can have multiple remediations over time (e.g., failed attempts followed by
// retries).
type BetaClusterRemediationNewResponse struct {
	ID         string `json:"id" api:"required"`
	ClusterID  string `json:"cluster_id" api:"required"`
	InstanceID string `json:"instance_id" api:"required"`
	// Remediation mode specifies how the remediation should be performed.
	//
	//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
	//     available host.
	//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
	//     provisions a new one on a different host.
	//
	// Any of "REMEDIATION_MODE_VM_ONLY", "REMEDIATION_MODE_HOST_AWARE",
	// "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT", "REMEDIATION_MODE_REBOOT_VM".
	Mode BetaClusterRemediationNewResponseMode `json:"mode" api:"required"`
	// RemediationState represents the lifecycle state of a remediation.
	//
	//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
	//   - `PENDING`: Approved and queued for processing.
	//   - `RUNNING`: Actively being processed.
	//   - `SUCCEEDED`: Successfully completed.
	//   - `FAILED`: Failed with an error.
	//   - `CANCELLED`: Cancelled by user or system.
	//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
	//     processing.
	//
	// Any of "PENDING_APPROVAL", "PENDING", "RUNNING", "SUCCEEDED", "FAILED",
	// "CANCELLED", "AUTO_RESOLVED".
	State BetaClusterRemediationNewResponseState `json:"state" api:"required"`
	// RemediationTrigger specifies how the remediation was triggered.
	//
	//   - `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI
	//     or API call).
	//   - `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires
	//     approval.
	//
	// Any of "REMEDIATION_TRIGGER_MANUAL", "REMEDIATION_TRIGGER_AUTOMATED".
	Trigger BetaClusterRemediationNewResponseTrigger `json:"trigger" api:"required"`
	// Active health check run ID (UUID) that triggered this remediation.
	ActiveHealthCheckRunID string `json:"active_health_check_run_id"`
	// When the remediation was created.
	CreateTime time.Time `json:"create_time" format:"date-time"`
	// When the remediation completed.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Error message if the remediation failed.
	ErrorMessage string `json:"error_message"`
	// Passive health check event ID that triggered this remediation.
	PassiveHealthCheckEventID string `json:"passive_health_check_event_id"`
	// User-provided reason for the remediation.
	Reason string `json:"reason"`
	// Who requested the remediation.
	RequestedBy string `json:"requested_by"`
	// Review comment.
	ReviewComment string `json:"review_comment"`
	// When the remediation was reviewed.
	ReviewTime time.Time `json:"review_time" format:"date-time"`
	// Who reviewed the remediation.
	ReviewedBy string `json:"reviewed_by"`
	// When processing started.
	StartTime time.Time `json:"start_time" format:"date-time"`
	// When the remediation was last updated.
	UpdateTime time.Time `json:"update_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ClusterID                 respjson.Field
		InstanceID                respjson.Field
		Mode                      respjson.Field
		State                     respjson.Field
		Trigger                   respjson.Field
		ActiveHealthCheckRunID    respjson.Field
		CreateTime                respjson.Field
		EndTime                   respjson.Field
		ErrorMessage              respjson.Field
		PassiveHealthCheckEventID respjson.Field
		Reason                    respjson.Field
		RequestedBy               respjson.Field
		ReviewComment             respjson.Field
		ReviewTime                respjson.Field
		ReviewedBy                respjson.Field
		StartTime                 respjson.Field
		UpdateTime                respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterRemediationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterRemediationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remediation mode specifies how the remediation should be performed.
//
//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
//     available host.
//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
//     provisions a new one on a different host.
type BetaClusterRemediationNewResponseMode string

const (
	BetaClusterRemediationNewResponseModeRemediationModeVmOnly                  BetaClusterRemediationNewResponseMode = "REMEDIATION_MODE_VM_ONLY"
	BetaClusterRemediationNewResponseModeRemediationModeHostAware               BetaClusterRemediationNewResponseMode = "REMEDIATION_MODE_HOST_AWARE"
	BetaClusterRemediationNewResponseModeRemediationModeEvictWithoutReplacement BetaClusterRemediationNewResponseMode = "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT"
	BetaClusterRemediationNewResponseModeRemediationModeRebootVm                BetaClusterRemediationNewResponseMode = "REMEDIATION_MODE_REBOOT_VM"
)

// RemediationState represents the lifecycle state of a remediation.
//
//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
//   - `PENDING`: Approved and queued for processing.
//   - `RUNNING`: Actively being processed.
//   - `SUCCEEDED`: Successfully completed.
//   - `FAILED`: Failed with an error.
//   - `CANCELLED`: Cancelled by user or system.
//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
//     processing.
type BetaClusterRemediationNewResponseState string

const (
	BetaClusterRemediationNewResponseStatePendingApproval BetaClusterRemediationNewResponseState = "PENDING_APPROVAL"
	BetaClusterRemediationNewResponseStatePending         BetaClusterRemediationNewResponseState = "PENDING"
	BetaClusterRemediationNewResponseStateRunning         BetaClusterRemediationNewResponseState = "RUNNING"
	BetaClusterRemediationNewResponseStateSucceeded       BetaClusterRemediationNewResponseState = "SUCCEEDED"
	BetaClusterRemediationNewResponseStateFailed          BetaClusterRemediationNewResponseState = "FAILED"
	BetaClusterRemediationNewResponseStateCancelled       BetaClusterRemediationNewResponseState = "CANCELLED"
	BetaClusterRemediationNewResponseStateAutoResolved    BetaClusterRemediationNewResponseState = "AUTO_RESOLVED"
)

// RemediationTrigger specifies how the remediation was triggered.
//
//   - `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI
//     or API call).
//   - `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires
//     approval.
type BetaClusterRemediationNewResponseTrigger string

const (
	BetaClusterRemediationNewResponseTriggerRemediationTriggerManual    BetaClusterRemediationNewResponseTrigger = "REMEDIATION_TRIGGER_MANUAL"
	BetaClusterRemediationNewResponseTriggerRemediationTriggerAutomated BetaClusterRemediationNewResponseTrigger = "REMEDIATION_TRIGGER_AUTOMATED"
)

// ListRemediationsResponse is the response for ListRemediations.
type BetaClusterRemediationListResponse struct {
	// Indicates if there are more results available.
	HasNext bool `json:"has_next" api:"required"`
	// Token for the next page.
	NextPageToken string `json:"next_page_token" api:"required"`
	// The list of remediations.
	Remediations []BetaClusterRemediationListResponseRemediation `json:"remediations" api:"required"`
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

// Remediation represents a node remediation request for an instance. An instance
// can have multiple remediations over time (e.g., failed attempts followed by
// retries).
type BetaClusterRemediationListResponseRemediation struct {
	ID         string `json:"id" api:"required"`
	ClusterID  string `json:"cluster_id" api:"required"`
	InstanceID string `json:"instance_id" api:"required"`
	// Remediation mode specifies how the remediation should be performed.
	//
	//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
	//     available host.
	//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
	//     provisions a new one on a different host.
	//
	// Any of "REMEDIATION_MODE_VM_ONLY", "REMEDIATION_MODE_HOST_AWARE",
	// "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT", "REMEDIATION_MODE_REBOOT_VM".
	Mode string `json:"mode" api:"required"`
	// RemediationState represents the lifecycle state of a remediation.
	//
	//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
	//   - `PENDING`: Approved and queued for processing.
	//   - `RUNNING`: Actively being processed.
	//   - `SUCCEEDED`: Successfully completed.
	//   - `FAILED`: Failed with an error.
	//   - `CANCELLED`: Cancelled by user or system.
	//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
	//     processing.
	//
	// Any of "PENDING_APPROVAL", "PENDING", "RUNNING", "SUCCEEDED", "FAILED",
	// "CANCELLED", "AUTO_RESOLVED".
	State string `json:"state" api:"required"`
	// RemediationTrigger specifies how the remediation was triggered.
	//
	//   - `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI
	//     or API call).
	//   - `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires
	//     approval.
	//
	// Any of "REMEDIATION_TRIGGER_MANUAL", "REMEDIATION_TRIGGER_AUTOMATED".
	Trigger string `json:"trigger" api:"required"`
	// Active health check run ID (UUID) that triggered this remediation.
	ActiveHealthCheckRunID string `json:"active_health_check_run_id"`
	// When the remediation was created.
	CreateTime time.Time `json:"create_time" format:"date-time"`
	// When the remediation completed.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Error message if the remediation failed.
	ErrorMessage string `json:"error_message"`
	// Passive health check event ID that triggered this remediation.
	PassiveHealthCheckEventID string `json:"passive_health_check_event_id"`
	// User-provided reason for the remediation.
	Reason string `json:"reason"`
	// Who requested the remediation.
	RequestedBy string `json:"requested_by"`
	// Review comment.
	ReviewComment string `json:"review_comment"`
	// When the remediation was reviewed.
	ReviewTime time.Time `json:"review_time" format:"date-time"`
	// Who reviewed the remediation.
	ReviewedBy string `json:"reviewed_by"`
	// When processing started.
	StartTime time.Time `json:"start_time" format:"date-time"`
	// When the remediation was last updated.
	UpdateTime time.Time `json:"update_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ClusterID                 respjson.Field
		InstanceID                respjson.Field
		Mode                      respjson.Field
		State                     respjson.Field
		Trigger                   respjson.Field
		ActiveHealthCheckRunID    respjson.Field
		CreateTime                respjson.Field
		EndTime                   respjson.Field
		ErrorMessage              respjson.Field
		PassiveHealthCheckEventID respjson.Field
		Reason                    respjson.Field
		RequestedBy               respjson.Field
		ReviewComment             respjson.Field
		ReviewTime                respjson.Field
		ReviewedBy                respjson.Field
		StartTime                 respjson.Field
		UpdateTime                respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterRemediationListResponseRemediation) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterRemediationListResponseRemediation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remediation represents a node remediation request for an instance. An instance
// can have multiple remediations over time (e.g., failed attempts followed by
// retries).
type BetaClusterRemediationApproveResponse struct {
	ID         string `json:"id" api:"required"`
	ClusterID  string `json:"cluster_id" api:"required"`
	InstanceID string `json:"instance_id" api:"required"`
	// Remediation mode specifies how the remediation should be performed.
	//
	//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
	//     available host.
	//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
	//     provisions a new one on a different host.
	//
	// Any of "REMEDIATION_MODE_VM_ONLY", "REMEDIATION_MODE_HOST_AWARE",
	// "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT", "REMEDIATION_MODE_REBOOT_VM".
	Mode BetaClusterRemediationApproveResponseMode `json:"mode" api:"required"`
	// RemediationState represents the lifecycle state of a remediation.
	//
	//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
	//   - `PENDING`: Approved and queued for processing.
	//   - `RUNNING`: Actively being processed.
	//   - `SUCCEEDED`: Successfully completed.
	//   - `FAILED`: Failed with an error.
	//   - `CANCELLED`: Cancelled by user or system.
	//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
	//     processing.
	//
	// Any of "PENDING_APPROVAL", "PENDING", "RUNNING", "SUCCEEDED", "FAILED",
	// "CANCELLED", "AUTO_RESOLVED".
	State BetaClusterRemediationApproveResponseState `json:"state" api:"required"`
	// RemediationTrigger specifies how the remediation was triggered.
	//
	//   - `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI
	//     or API call).
	//   - `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires
	//     approval.
	//
	// Any of "REMEDIATION_TRIGGER_MANUAL", "REMEDIATION_TRIGGER_AUTOMATED".
	Trigger BetaClusterRemediationApproveResponseTrigger `json:"trigger" api:"required"`
	// Active health check run ID (UUID) that triggered this remediation.
	ActiveHealthCheckRunID string `json:"active_health_check_run_id"`
	// When the remediation was created.
	CreateTime time.Time `json:"create_time" format:"date-time"`
	// When the remediation completed.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Error message if the remediation failed.
	ErrorMessage string `json:"error_message"`
	// Passive health check event ID that triggered this remediation.
	PassiveHealthCheckEventID string `json:"passive_health_check_event_id"`
	// User-provided reason for the remediation.
	Reason string `json:"reason"`
	// Who requested the remediation.
	RequestedBy string `json:"requested_by"`
	// Review comment.
	ReviewComment string `json:"review_comment"`
	// When the remediation was reviewed.
	ReviewTime time.Time `json:"review_time" format:"date-time"`
	// Who reviewed the remediation.
	ReviewedBy string `json:"reviewed_by"`
	// When processing started.
	StartTime time.Time `json:"start_time" format:"date-time"`
	// When the remediation was last updated.
	UpdateTime time.Time `json:"update_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ClusterID                 respjson.Field
		InstanceID                respjson.Field
		Mode                      respjson.Field
		State                     respjson.Field
		Trigger                   respjson.Field
		ActiveHealthCheckRunID    respjson.Field
		CreateTime                respjson.Field
		EndTime                   respjson.Field
		ErrorMessage              respjson.Field
		PassiveHealthCheckEventID respjson.Field
		Reason                    respjson.Field
		RequestedBy               respjson.Field
		ReviewComment             respjson.Field
		ReviewTime                respjson.Field
		ReviewedBy                respjson.Field
		StartTime                 respjson.Field
		UpdateTime                respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterRemediationApproveResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterRemediationApproveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remediation mode specifies how the remediation should be performed.
//
//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
//     available host.
//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
//     provisions a new one on a different host.
type BetaClusterRemediationApproveResponseMode string

const (
	BetaClusterRemediationApproveResponseModeRemediationModeVmOnly                  BetaClusterRemediationApproveResponseMode = "REMEDIATION_MODE_VM_ONLY"
	BetaClusterRemediationApproveResponseModeRemediationModeHostAware               BetaClusterRemediationApproveResponseMode = "REMEDIATION_MODE_HOST_AWARE"
	BetaClusterRemediationApproveResponseModeRemediationModeEvictWithoutReplacement BetaClusterRemediationApproveResponseMode = "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT"
	BetaClusterRemediationApproveResponseModeRemediationModeRebootVm                BetaClusterRemediationApproveResponseMode = "REMEDIATION_MODE_REBOOT_VM"
)

// RemediationState represents the lifecycle state of a remediation.
//
//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
//   - `PENDING`: Approved and queued for processing.
//   - `RUNNING`: Actively being processed.
//   - `SUCCEEDED`: Successfully completed.
//   - `FAILED`: Failed with an error.
//   - `CANCELLED`: Cancelled by user or system.
//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
//     processing.
type BetaClusterRemediationApproveResponseState string

const (
	BetaClusterRemediationApproveResponseStatePendingApproval BetaClusterRemediationApproveResponseState = "PENDING_APPROVAL"
	BetaClusterRemediationApproveResponseStatePending         BetaClusterRemediationApproveResponseState = "PENDING"
	BetaClusterRemediationApproveResponseStateRunning         BetaClusterRemediationApproveResponseState = "RUNNING"
	BetaClusterRemediationApproveResponseStateSucceeded       BetaClusterRemediationApproveResponseState = "SUCCEEDED"
	BetaClusterRemediationApproveResponseStateFailed          BetaClusterRemediationApproveResponseState = "FAILED"
	BetaClusterRemediationApproveResponseStateCancelled       BetaClusterRemediationApproveResponseState = "CANCELLED"
	BetaClusterRemediationApproveResponseStateAutoResolved    BetaClusterRemediationApproveResponseState = "AUTO_RESOLVED"
)

// RemediationTrigger specifies how the remediation was triggered.
//
//   - `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI
//     or API call).
//   - `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires
//     approval.
type BetaClusterRemediationApproveResponseTrigger string

const (
	BetaClusterRemediationApproveResponseTriggerRemediationTriggerManual    BetaClusterRemediationApproveResponseTrigger = "REMEDIATION_TRIGGER_MANUAL"
	BetaClusterRemediationApproveResponseTriggerRemediationTriggerAutomated BetaClusterRemediationApproveResponseTrigger = "REMEDIATION_TRIGGER_AUTOMATED"
)

// Remediation represents a node remediation request for an instance. An instance
// can have multiple remediations over time (e.g., failed attempts followed by
// retries).
type BetaClusterRemediationCancelResponse struct {
	ID         string `json:"id" api:"required"`
	ClusterID  string `json:"cluster_id" api:"required"`
	InstanceID string `json:"instance_id" api:"required"`
	// Remediation mode specifies how the remediation should be performed.
	//
	//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
	//     available host.
	//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
	//     provisions a new one on a different host.
	//
	// Any of "REMEDIATION_MODE_VM_ONLY", "REMEDIATION_MODE_HOST_AWARE",
	// "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT", "REMEDIATION_MODE_REBOOT_VM".
	Mode BetaClusterRemediationCancelResponseMode `json:"mode" api:"required"`
	// RemediationState represents the lifecycle state of a remediation.
	//
	//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
	//   - `PENDING`: Approved and queued for processing.
	//   - `RUNNING`: Actively being processed.
	//   - `SUCCEEDED`: Successfully completed.
	//   - `FAILED`: Failed with an error.
	//   - `CANCELLED`: Cancelled by user or system.
	//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
	//     processing.
	//
	// Any of "PENDING_APPROVAL", "PENDING", "RUNNING", "SUCCEEDED", "FAILED",
	// "CANCELLED", "AUTO_RESOLVED".
	State BetaClusterRemediationCancelResponseState `json:"state" api:"required"`
	// RemediationTrigger specifies how the remediation was triggered.
	//
	//   - `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI
	//     or API call).
	//   - `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires
	//     approval.
	//
	// Any of "REMEDIATION_TRIGGER_MANUAL", "REMEDIATION_TRIGGER_AUTOMATED".
	Trigger BetaClusterRemediationCancelResponseTrigger `json:"trigger" api:"required"`
	// Active health check run ID (UUID) that triggered this remediation.
	ActiveHealthCheckRunID string `json:"active_health_check_run_id"`
	// When the remediation was created.
	CreateTime time.Time `json:"create_time" format:"date-time"`
	// When the remediation completed.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Error message if the remediation failed.
	ErrorMessage string `json:"error_message"`
	// Passive health check event ID that triggered this remediation.
	PassiveHealthCheckEventID string `json:"passive_health_check_event_id"`
	// User-provided reason for the remediation.
	Reason string `json:"reason"`
	// Who requested the remediation.
	RequestedBy string `json:"requested_by"`
	// Review comment.
	ReviewComment string `json:"review_comment"`
	// When the remediation was reviewed.
	ReviewTime time.Time `json:"review_time" format:"date-time"`
	// Who reviewed the remediation.
	ReviewedBy string `json:"reviewed_by"`
	// When processing started.
	StartTime time.Time `json:"start_time" format:"date-time"`
	// When the remediation was last updated.
	UpdateTime time.Time `json:"update_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ClusterID                 respjson.Field
		InstanceID                respjson.Field
		Mode                      respjson.Field
		State                     respjson.Field
		Trigger                   respjson.Field
		ActiveHealthCheckRunID    respjson.Field
		CreateTime                respjson.Field
		EndTime                   respjson.Field
		ErrorMessage              respjson.Field
		PassiveHealthCheckEventID respjson.Field
		Reason                    respjson.Field
		RequestedBy               respjson.Field
		ReviewComment             respjson.Field
		ReviewTime                respjson.Field
		ReviewedBy                respjson.Field
		StartTime                 respjson.Field
		UpdateTime                respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterRemediationCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterRemediationCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remediation mode specifies how the remediation should be performed.
//
//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
//     available host.
//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
//     provisions a new one on a different host.
type BetaClusterRemediationCancelResponseMode string

const (
	BetaClusterRemediationCancelResponseModeRemediationModeVmOnly                  BetaClusterRemediationCancelResponseMode = "REMEDIATION_MODE_VM_ONLY"
	BetaClusterRemediationCancelResponseModeRemediationModeHostAware               BetaClusterRemediationCancelResponseMode = "REMEDIATION_MODE_HOST_AWARE"
	BetaClusterRemediationCancelResponseModeRemediationModeEvictWithoutReplacement BetaClusterRemediationCancelResponseMode = "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT"
	BetaClusterRemediationCancelResponseModeRemediationModeRebootVm                BetaClusterRemediationCancelResponseMode = "REMEDIATION_MODE_REBOOT_VM"
)

// RemediationState represents the lifecycle state of a remediation.
//
//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
//   - `PENDING`: Approved and queued for processing.
//   - `RUNNING`: Actively being processed.
//   - `SUCCEEDED`: Successfully completed.
//   - `FAILED`: Failed with an error.
//   - `CANCELLED`: Cancelled by user or system.
//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
//     processing.
type BetaClusterRemediationCancelResponseState string

const (
	BetaClusterRemediationCancelResponseStatePendingApproval BetaClusterRemediationCancelResponseState = "PENDING_APPROVAL"
	BetaClusterRemediationCancelResponseStatePending         BetaClusterRemediationCancelResponseState = "PENDING"
	BetaClusterRemediationCancelResponseStateRunning         BetaClusterRemediationCancelResponseState = "RUNNING"
	BetaClusterRemediationCancelResponseStateSucceeded       BetaClusterRemediationCancelResponseState = "SUCCEEDED"
	BetaClusterRemediationCancelResponseStateFailed          BetaClusterRemediationCancelResponseState = "FAILED"
	BetaClusterRemediationCancelResponseStateCancelled       BetaClusterRemediationCancelResponseState = "CANCELLED"
	BetaClusterRemediationCancelResponseStateAutoResolved    BetaClusterRemediationCancelResponseState = "AUTO_RESOLVED"
)

// RemediationTrigger specifies how the remediation was triggered.
//
//   - `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI
//     or API call).
//   - `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires
//     approval.
type BetaClusterRemediationCancelResponseTrigger string

const (
	BetaClusterRemediationCancelResponseTriggerRemediationTriggerManual    BetaClusterRemediationCancelResponseTrigger = "REMEDIATION_TRIGGER_MANUAL"
	BetaClusterRemediationCancelResponseTriggerRemediationTriggerAutomated BetaClusterRemediationCancelResponseTrigger = "REMEDIATION_TRIGGER_AUTOMATED"
)

// Remediation represents a node remediation request for an instance. An instance
// can have multiple remediations over time (e.g., failed attempts followed by
// retries).
type BetaClusterRemediationRejectResponse struct {
	ID         string `json:"id" api:"required"`
	ClusterID  string `json:"cluster_id" api:"required"`
	InstanceID string `json:"instance_id" api:"required"`
	// Remediation mode specifies how the remediation should be performed.
	//
	//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
	//     available host.
	//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
	//     provisions a new one on a different host.
	//
	// Any of "REMEDIATION_MODE_VM_ONLY", "REMEDIATION_MODE_HOST_AWARE",
	// "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT", "REMEDIATION_MODE_REBOOT_VM".
	Mode BetaClusterRemediationRejectResponseMode `json:"mode" api:"required"`
	// RemediationState represents the lifecycle state of a remediation.
	//
	//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
	//   - `PENDING`: Approved and queued for processing.
	//   - `RUNNING`: Actively being processed.
	//   - `SUCCEEDED`: Successfully completed.
	//   - `FAILED`: Failed with an error.
	//   - `CANCELLED`: Cancelled by user or system.
	//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
	//     processing.
	//
	// Any of "PENDING_APPROVAL", "PENDING", "RUNNING", "SUCCEEDED", "FAILED",
	// "CANCELLED", "AUTO_RESOLVED".
	State BetaClusterRemediationRejectResponseState `json:"state" api:"required"`
	// RemediationTrigger specifies how the remediation was triggered.
	//
	//   - `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI
	//     or API call).
	//   - `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires
	//     approval.
	//
	// Any of "REMEDIATION_TRIGGER_MANUAL", "REMEDIATION_TRIGGER_AUTOMATED".
	Trigger BetaClusterRemediationRejectResponseTrigger `json:"trigger" api:"required"`
	// Active health check run ID (UUID) that triggered this remediation.
	ActiveHealthCheckRunID string `json:"active_health_check_run_id"`
	// When the remediation was created.
	CreateTime time.Time `json:"create_time" format:"date-time"`
	// When the remediation completed.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Error message if the remediation failed.
	ErrorMessage string `json:"error_message"`
	// Passive health check event ID that triggered this remediation.
	PassiveHealthCheckEventID string `json:"passive_health_check_event_id"`
	// User-provided reason for the remediation.
	Reason string `json:"reason"`
	// Who requested the remediation.
	RequestedBy string `json:"requested_by"`
	// Review comment.
	ReviewComment string `json:"review_comment"`
	// When the remediation was reviewed.
	ReviewTime time.Time `json:"review_time" format:"date-time"`
	// Who reviewed the remediation.
	ReviewedBy string `json:"reviewed_by"`
	// When processing started.
	StartTime time.Time `json:"start_time" format:"date-time"`
	// When the remediation was last updated.
	UpdateTime time.Time `json:"update_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ClusterID                 respjson.Field
		InstanceID                respjson.Field
		Mode                      respjson.Field
		State                     respjson.Field
		Trigger                   respjson.Field
		ActiveHealthCheckRunID    respjson.Field
		CreateTime                respjson.Field
		EndTime                   respjson.Field
		ErrorMessage              respjson.Field
		PassiveHealthCheckEventID respjson.Field
		Reason                    respjson.Field
		RequestedBy               respjson.Field
		ReviewComment             respjson.Field
		ReviewTime                respjson.Field
		ReviewedBy                respjson.Field
		StartTime                 respjson.Field
		UpdateTime                respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaClusterRemediationRejectResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaClusterRemediationRejectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remediation mode specifies how the remediation should be performed.
//
//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
//     available host.
//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
//     provisions a new one on a different host.
type BetaClusterRemediationRejectResponseMode string

const (
	BetaClusterRemediationRejectResponseModeRemediationModeVmOnly                  BetaClusterRemediationRejectResponseMode = "REMEDIATION_MODE_VM_ONLY"
	BetaClusterRemediationRejectResponseModeRemediationModeHostAware               BetaClusterRemediationRejectResponseMode = "REMEDIATION_MODE_HOST_AWARE"
	BetaClusterRemediationRejectResponseModeRemediationModeEvictWithoutReplacement BetaClusterRemediationRejectResponseMode = "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT"
	BetaClusterRemediationRejectResponseModeRemediationModeRebootVm                BetaClusterRemediationRejectResponseMode = "REMEDIATION_MODE_REBOOT_VM"
)

// RemediationState represents the lifecycle state of a remediation.
//
//   - `PENDING_APPROVAL`: Awaiting approval before processing can begin.
//   - `PENDING`: Approved and queued for processing.
//   - `RUNNING`: Actively being processed.
//   - `SUCCEEDED`: Successfully completed.
//   - `FAILED`: Failed with an error.
//   - `CANCELLED`: Cancelled by user or system.
//   - `AUTO_RESOLVED`: The underlying issue was automatically resolved before
//     processing.
type BetaClusterRemediationRejectResponseState string

const (
	BetaClusterRemediationRejectResponseStatePendingApproval BetaClusterRemediationRejectResponseState = "PENDING_APPROVAL"
	BetaClusterRemediationRejectResponseStatePending         BetaClusterRemediationRejectResponseState = "PENDING"
	BetaClusterRemediationRejectResponseStateRunning         BetaClusterRemediationRejectResponseState = "RUNNING"
	BetaClusterRemediationRejectResponseStateSucceeded       BetaClusterRemediationRejectResponseState = "SUCCEEDED"
	BetaClusterRemediationRejectResponseStateFailed          BetaClusterRemediationRejectResponseState = "FAILED"
	BetaClusterRemediationRejectResponseStateCancelled       BetaClusterRemediationRejectResponseState = "CANCELLED"
	BetaClusterRemediationRejectResponseStateAutoResolved    BetaClusterRemediationRejectResponseState = "AUTO_RESOLVED"
)

// RemediationTrigger specifies how the remediation was triggered.
//
//   - `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI
//     or API call).
//   - `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires
//     approval.
type BetaClusterRemediationRejectResponseTrigger string

const (
	BetaClusterRemediationRejectResponseTriggerRemediationTriggerManual    BetaClusterRemediationRejectResponseTrigger = "REMEDIATION_TRIGGER_MANUAL"
	BetaClusterRemediationRejectResponseTriggerRemediationTriggerAutomated BetaClusterRemediationRejectResponseTrigger = "REMEDIATION_TRIGGER_AUTOMATED"
)

type BetaClusterRemediationNewParams struct {
	// The cluster ID.
	ClusterID string `path:"cluster_id" api:"required" json:"-"`
	// Remediation mode specifies how the remediation should be performed.
	//
	//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
	//     available host.
	//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
	//     provisions a new one on a different host.
	//
	// Any of "REMEDIATION_MODE_VM_ONLY", "REMEDIATION_MODE_HOST_AWARE",
	// "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT", "REMEDIATION_MODE_REBOOT_VM".
	Mode BetaClusterRemediationNewParamsMode `json:"mode,omitzero" api:"required"`
	// Optional. Client-specified ID for idempotency.
	RemediationID param.Opt[string] `query:"remediation_id,omitzero" json:"-"`
	// User-provided reason for the remediation.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r BetaClusterRemediationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterRemediationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
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

// Remediation mode specifies how the remediation should be performed.
//
//   - `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any
//     available host.
//   - `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and
//     provisions a new one on a different host.
type BetaClusterRemediationNewParamsMode string

const (
	BetaClusterRemediationNewParamsModeRemediationModeVmOnly                  BetaClusterRemediationNewParamsMode = "REMEDIATION_MODE_VM_ONLY"
	BetaClusterRemediationNewParamsModeRemediationModeHostAware               BetaClusterRemediationNewParamsMode = "REMEDIATION_MODE_HOST_AWARE"
	BetaClusterRemediationNewParamsModeRemediationModeEvictWithoutReplacement BetaClusterRemediationNewParamsMode = "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT"
	BetaClusterRemediationNewParamsModeRemediationModeRebootVm                BetaClusterRemediationNewParamsMode = "REMEDIATION_MODE_REBOOT_VM"
)

type BetaClusterRemediationListParams struct {
	// The cluster ID.
	ClusterID string `path:"cluster_id" api:"required" json:"-"`
	// Optional. Order by expression.
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Optional. Maximum results to return.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Optional. Pagination token from previous request.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	// Optional. Filter by remediation mode. Returns only remediations matching the
	// specified mode.
	//
	// Any of "REMEDIATION_MODE_VM_ONLY", "REMEDIATION_MODE_HOST_AWARE",
	// "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT", "REMEDIATION_MODE_REBOOT_VM".
	Mode BetaClusterRemediationListParamsMode `query:"mode,omitzero" json:"-"`
	// Optional. Filter by state(s). Returns remediations matching any of the specified
	// states.
	//
	// Any of "PENDING_APPROVAL", "PENDING", "RUNNING", "SUCCEEDED", "FAILED",
	// "CANCELLED", "AUTO_RESOLVED".
	State []string `query:"state,omitzero" json:"-"`
	// Optional. Filter by trigger type. Returns only remediations matching the
	// specified trigger.
	//
	// Any of "REMEDIATION_TRIGGER_MANUAL", "REMEDIATION_TRIGGER_AUTOMATED".
	Trigger BetaClusterRemediationListParamsTrigger `query:"trigger,omitzero" json:"-"`
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

// Optional. Filter by remediation mode. Returns only remediations matching the
// specified mode.
type BetaClusterRemediationListParamsMode string

const (
	BetaClusterRemediationListParamsModeRemediationModeVmOnly                  BetaClusterRemediationListParamsMode = "REMEDIATION_MODE_VM_ONLY"
	BetaClusterRemediationListParamsModeRemediationModeHostAware               BetaClusterRemediationListParamsMode = "REMEDIATION_MODE_HOST_AWARE"
	BetaClusterRemediationListParamsModeRemediationModeEvictWithoutReplacement BetaClusterRemediationListParamsMode = "REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT"
	BetaClusterRemediationListParamsModeRemediationModeRebootVm                BetaClusterRemediationListParamsMode = "REMEDIATION_MODE_REBOOT_VM"
)

// Optional. Filter by trigger type. Returns only remediations matching the
// specified trigger.
type BetaClusterRemediationListParamsTrigger string

const (
	BetaClusterRemediationListParamsTriggerRemediationTriggerManual    BetaClusterRemediationListParamsTrigger = "REMEDIATION_TRIGGER_MANUAL"
	BetaClusterRemediationListParamsTriggerRemediationTriggerAutomated BetaClusterRemediationListParamsTrigger = "REMEDIATION_TRIGGER_AUTOMATED"
)

type BetaClusterRemediationApproveParams struct {
	// The cluster ID.
	ClusterID string `path:"cluster_id" api:"required" json:"-"`
	// The instance ID.
	InstanceID string `path:"instance_id" api:"required" json:"-"`
	// Comment explaining the action.
	Comment param.Opt[string] `json:"comment,omitzero"`
	paramObj
}

func (r BetaClusterRemediationApproveParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaClusterRemediationApproveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaClusterRemediationApproveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaClusterRemediationCancelParams struct {
	// The cluster ID.
	ClusterID string `path:"cluster_id" api:"required" json:"-"`
	// The instance ID.
	InstanceID string `path:"instance_id" api:"required" json:"-"`
	paramObj
}

type BetaClusterRemediationRejectParams struct {
	// The cluster ID.
	ClusterID string `path:"cluster_id" api:"required" json:"-"`
	// The instance ID.
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
