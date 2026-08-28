// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"encoding/json"
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

// BetaRlSessionService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaRlSessionService] method instead.
type BetaRlSessionService struct {
	Options []option.RequestOption
}

// NewBetaRlSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaRlSessionService(opts ...option.RequestOption) (r BetaRlSessionService) {
	r = BetaRlSessionService{}
	r.Options = opts
	return
}

// Creates a training session and returns its details.
func (r *BetaRlSessionService) New(ctx context.Context, body BetaRlSessionNewParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rl/training-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets a training session by its ID and returns its details.
func (r *BetaRlSessionService) Get(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all training sessions.
func (r *BetaRlSessionService) List(ctx context.Context, query BetaRlSessionListParams, opts ...option.RequestOption) (res *SessionsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rl/training-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Stops a training session.
func (r *BetaRlSessionService) Stop(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/training-sessions/%s/stop", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Saved inference checkpoint
type InferenceCheckpoint struct {
	// Unique identifier for the checkpoint
	ID string `json:"id" api:"required"`
	// Timestamp when the checkpoint was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Training step at time of save
	Step InferenceCheckpointStepUnion `json:"step" api:"required"`
	// Model registration details
	Registration InferenceCheckpointRegistration `json:"registration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Step         respjson.Field
		Registration respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceCheckpoint) RawJSON() string { return r.JSON.raw }
func (r *InferenceCheckpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InferenceCheckpointStepUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InferenceCheckpointStepUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InferenceCheckpointStepUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InferenceCheckpointStepUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InferenceCheckpointStepUnion) RawJSON() string { return u.JSON.raw }

func (r *InferenceCheckpointStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model registration details
type InferenceCheckpointRegistration struct {
	// Registered model name for downloading the checkpoint
	ModelName string `json:"model_name" api:"required"`
	// Timestamp when the model was registered
	RegisteredAt time.Time `json:"registered_at" api:"required" format:"date-time"`
	// Together model registry object ID for the adapter checkpoint (e.g. `ml_...`),
	// set on LoRA training sessions
	AdapterObjectID string `json:"adapter_object_id"`
	// Together model registry revision ID for the adapter checkpoint (e.g. `rv_...`)
	AdapterObjectRevisionID string `json:"adapter_object_revision_id"`
	// Together model registry object ID for the model checkpoint (e.g. `ml_...`), set
	// on full-weight training sessions
	ModelObjectID string `json:"model_object_id"`
	// Together model registry revision ID for the model checkpoint (e.g. `rv_...`)
	ModelObjectRevisionID string `json:"model_object_revision_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelName               respjson.Field
		RegisteredAt            respjson.Field
		AdapterObjectID         respjson.Field
		AdapterObjectRevisionID respjson.Field
		ModelObjectID           respjson.Field
		ModelObjectRevisionID   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceCheckpointRegistration) RawJSON() string { return r.JSON.raw }
func (r *InferenceCheckpointRegistration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LoRA adapter configuration
type LoraConfig struct {
	// Alpha of the LoRA adapter
	Alpha int64 `json:"alpha"`
	// Dropout of the LoRA adapter
	Dropout float64 `json:"dropout"`
	// Rank of the LoRA adapter
	Rank int64 `json:"rank"`
	// Random seed for initializing LoRA adapter weights. Ignored when LoRA is disabled
	// or the session resumes from a checkpoint.
	Seed LoraConfigSeedUnion `json:"seed"`
	// Whether to also train a LoRA adapter on the output head. Defaults to true.
	TrainUnembed bool `json:"train_unembed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alpha        respjson.Field
		Dropout      respjson.Field
		Rank         respjson.Field
		Seed         respjson.Field
		TrainUnembed respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoraConfig) RawJSON() string { return r.JSON.raw }
func (r *LoraConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LoraConfig to a LoraConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LoraConfigParam.Overrides()
func (r LoraConfig) ToParam() LoraConfigParam {
	return param.Override[LoraConfigParam](json.RawMessage(r.RawJSON()))
}

// LoraConfigSeedUnion contains all possible properties and values from [string],
// [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type LoraConfigSeedUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u LoraConfigSeedUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LoraConfigSeedUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LoraConfigSeedUnion) RawJSON() string { return u.JSON.raw }

func (r *LoraConfigSeedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LoRA adapter configuration
type LoraConfigParam struct {
	// Alpha of the LoRA adapter
	Alpha param.Opt[int64] `json:"alpha,omitzero"`
	// Dropout of the LoRA adapter
	Dropout param.Opt[float64] `json:"dropout,omitzero"`
	// Rank of the LoRA adapter
	Rank param.Opt[int64] `json:"rank,omitzero"`
	// Whether to also train a LoRA adapter on the output head. Defaults to true.
	TrainUnembed param.Opt[bool] `json:"train_unembed,omitzero"`
	// Random seed for initializing LoRA adapter weights. Ignored when LoRA is disabled
	// or the session resumes from a checkpoint.
	Seed LoraConfigSeedUnionParam `json:"seed,omitzero"`
	paramObj
}

func (r LoraConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow LoraConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoraConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LoraConfigSeedUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u LoraConfigSeedUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LoraConfigSeedUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LoraConfigSeedUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// A training session and its current state
type Session struct {
	// ID of the training session
	ID string `json:"id" api:"required"`
	// Base model the session trains, taken from the model resource it is attached to
	BaseModel string `json:"base_model" api:"required"`
	// Timestamp when the training session was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ID of the user who created the training session
	CreatedBy string `json:"created_by" api:"required"`
	// List of saved inference checkpoints for this session
	InferenceCheckpoints []InferenceCheckpoint `json:"inference_checkpoints" api:"required"`
	// Auxiliary metadata associated with the training session
	Metadata SessionMetadata `json:"metadata" api:"required"`
	// Model resource this session is attached to. The session runs on that resource's
	// GPU pods.
	ModelResourcesID string `json:"model_resources_id" api:"required"`
	// Session-scoped policy and weight versions for this session
	PolicyState SessionPolicyState `json:"policy_state" api:"required"`
	// Status of the training session
	//
	// Any of "TRAINING_SESSION_STATUS_UNSPECIFIED",
	// "TRAINING_SESSION_STATUS_CREATING", "TRAINING_SESSION_STATUS_RUNNING",
	// "TRAINING_SESSION_STATUS_STOPPED", "TRAINING_SESSION_STATUS_STOPPING",
	// "TRAINING_SESSION_STATUS_ERROR", "TRAINING_SESSION_STATUS_EXPIRED".
	Status SessionStatus `json:"status" api:"required"`
	// Current training step
	Step SessionStepUnion `json:"step" api:"required"`
	// List of saved training checkpoints for this session
	TrainingCheckpoints []TrainingCheckpoint `json:"training_checkpoints" api:"required"`
	// Timestamp when the training session was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Display name used to identify the training session
	DisplayName string `json:"display_name"`
	// Structured detail for the training session's current error. Set when the session
	// is in an error state.
	Error SessionError `json:"error"`
	// LoRA adapter configuration. Present only for sessions running on a LoRA-enabled
	// model resource.
	LoraConfig LoraConfig `json:"lora_config"`
	// Checkpoint ID this session was resumed from
	ResumeFromCheckpointID string `json:"resume_from_checkpoint_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		BaseModel              respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		InferenceCheckpoints   respjson.Field
		Metadata               respjson.Field
		ModelResourcesID       respjson.Field
		PolicyState            respjson.Field
		Status                 respjson.Field
		Step                   respjson.Field
		TrainingCheckpoints    respjson.Field
		UpdatedAt              respjson.Field
		DisplayName            respjson.Field
		Error                  respjson.Field
		LoraConfig             respjson.Field
		ResumeFromCheckpointID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Session) RawJSON() string { return r.JSON.raw }
func (r *Session) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Session-scoped policy and weight versions for this session
type SessionPolicyState struct {
	// Policy version successfully applied to the generator for this session.
	AppliedWeightsVersion SessionPolicyStateAppliedWeightsVersionUnion `json:"applied_weights_version" api:"required"`
	// True when a generator publish has been requested but has not finished.
	PendingPublish bool `json:"pending_publish" api:"required"`
	// Policy version promised to the generator by the latest weights-sync.
	TargetWeightsVersion SessionPolicyStateTargetWeightsVersionUnion `json:"target_weights_version" api:"required"`
	// Policy version produced by the last completed optimizer step. Distinct from
	// `TrainingSession.step`, which is the durable optimizer-step counter.
	TrainerStep SessionPolicyStateTrainerStepUnion `json:"trainer_step" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppliedWeightsVersion respjson.Field
		PendingPublish        respjson.Field
		TargetWeightsVersion  respjson.Field
		TrainerStep           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionPolicyState) RawJSON() string { return r.JSON.raw }
func (r *SessionPolicyState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SessionPolicyStateAppliedWeightsVersionUnion contains all possible properties
// and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type SessionPolicyStateAppliedWeightsVersionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u SessionPolicyStateAppliedWeightsVersionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SessionPolicyStateAppliedWeightsVersionUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SessionPolicyStateAppliedWeightsVersionUnion) RawJSON() string { return u.JSON.raw }

func (r *SessionPolicyStateAppliedWeightsVersionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SessionPolicyStateTargetWeightsVersionUnion contains all possible properties and
// values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type SessionPolicyStateTargetWeightsVersionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u SessionPolicyStateTargetWeightsVersionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SessionPolicyStateTargetWeightsVersionUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SessionPolicyStateTargetWeightsVersionUnion) RawJSON() string { return u.JSON.raw }

func (r *SessionPolicyStateTargetWeightsVersionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SessionPolicyStateTrainerStepUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type SessionPolicyStateTrainerStepUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u SessionPolicyStateTrainerStepUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SessionPolicyStateTrainerStepUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SessionPolicyStateTrainerStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SessionPolicyStateTrainerStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SessionStepUnion contains all possible properties and values from [string],
// [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type SessionStepUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u SessionStepUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SessionStepUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SessionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *SessionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured detail for the training session's current error
type SessionError struct {
	// Finite machine-readable reason code for UI branching
	//
	// Any of "TRAINING_SESSION_ERROR_CODE_RESOURCE_UNAVAILABLE",
	// "TRAINING_SESSION_ERROR_CODE_RESOURCE_AT_CAPACITY",
	// "TRAINING_SESSION_ERROR_CODE_TIMED_OUT",
	// "TRAINING_SESSION_ERROR_CODE_SESSION_FAILED".
	Code SessionErrorCode `json:"code" api:"required"`
	// User-safe human-readable detail for the current status
	Message string `json:"message" api:"required"`
	// Timestamp when this error was reported
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		OccurredAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionError) RawJSON() string { return r.JSON.raw }
func (r *SessionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Finite machine-readable training session lifecycle error code
type SessionErrorCode string

const (
	SessionErrorCodeTrainingSessionErrorCodeResourceUnavailable SessionErrorCode = "TRAINING_SESSION_ERROR_CODE_RESOURCE_UNAVAILABLE"
	SessionErrorCodeTrainingSessionErrorCodeResourceAtCapacity  SessionErrorCode = "TRAINING_SESSION_ERROR_CODE_RESOURCE_AT_CAPACITY"
	SessionErrorCodeTrainingSessionErrorCodeTimedOut            SessionErrorCode = "TRAINING_SESSION_ERROR_CODE_TIMED_OUT"
	SessionErrorCodeTrainingSessionErrorCodeSessionFailed       SessionErrorCode = "TRAINING_SESSION_ERROR_CODE_SESSION_FAILED"
)

// Auxiliary metadata associated with a training session
type SessionMetadata struct {
	// Weights & Biases details associated with the training session
	Wandb WandbMetadata `json:"wandb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Wandb       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionMetadata) RawJSON() string { return r.JSON.raw }
func (r *SessionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SessionMetadata to a SessionMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SessionMetadataParam.Overrides()
func (r SessionMetadata) ToParam() SessionMetadataParam {
	return param.Override[SessionMetadataParam](json.RawMessage(r.RawJSON()))
}

// Auxiliary metadata associated with a training session
type SessionMetadataParam struct {
	// Weights & Biases details associated with the training session
	Wandb WandbMetadataParam `json:"wandb,omitzero"`
	paramObj
}

func (r SessionMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow SessionMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the training session
type SessionStatus string

const (
	SessionStatusTrainingSessionStatusUnspecified SessionStatus = "TRAINING_SESSION_STATUS_UNSPECIFIED"
	SessionStatusTrainingSessionStatusCreating    SessionStatus = "TRAINING_SESSION_STATUS_CREATING"
	SessionStatusTrainingSessionStatusRunning     SessionStatus = "TRAINING_SESSION_STATUS_RUNNING"
	SessionStatusTrainingSessionStatusStopped     SessionStatus = "TRAINING_SESSION_STATUS_STOPPED"
	SessionStatusTrainingSessionStatusStopping    SessionStatus = "TRAINING_SESSION_STATUS_STOPPING"
	SessionStatusTrainingSessionStatusError       SessionStatus = "TRAINING_SESSION_STATUS_ERROR"
	SessionStatusTrainingSessionStatusExpired     SessionStatus = "TRAINING_SESSION_STATUS_EXPIRED"
)

// Paginated list of training sessions
type SessionsListResponse struct {
	// List of training sessions
	Data []Session `json:"data"`
	// Pagination metadata
	Meta SessionsListResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionsListResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionsListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type SessionsListResponseMeta struct {
	// Whether more items exist beyond this page
	HasMore bool `json:"has_more"`
	// Maximum number of items returned per page
	Limit int64 `json:"limit"`
	// Cursor to use as the 'after' parameter for the next page. Empty when has_more is
	// false.
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Limit       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionsListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SessionsListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Saved training checkpoint
type TrainingCheckpoint struct {
	// Unique identifier for the checkpoint
	ID string `json:"id" api:"required"`
	// Timestamp when the checkpoint was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Training step at time of save
	Step TrainingCheckpointStepUnion `json:"step" api:"required"`
	// Together model registry details, set when the checkpoint was uploaded to the
	// registry
	Registration TrainingCheckpointRegistration `json:"registration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Step         respjson.Field
		Registration respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrainingCheckpoint) RawJSON() string { return r.JSON.raw }
func (r *TrainingCheckpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TrainingCheckpointStepUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type TrainingCheckpointStepUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u TrainingCheckpointStepUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TrainingCheckpointStepUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TrainingCheckpointStepUnion) RawJSON() string { return u.JSON.raw }

func (r *TrainingCheckpointStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Together model registry details, set when the checkpoint was uploaded to the
// registry
type TrainingCheckpointRegistration struct {
	// Together model registry object ID for the training checkpoint artifact (e.g.
	// `ml_...`)
	ObjectID string `json:"object_id" api:"required"`
	// Together model registry revision ID for the training checkpoint artifact (e.g.
	// `rv_...`), empty when the upload reported no revision
	ObjectRevisionID string `json:"object_revision_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectID         respjson.Field
		ObjectRevisionID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrainingCheckpointRegistration) RawJSON() string { return r.JSON.raw }
func (r *TrainingCheckpointRegistration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details that associate a training session with a Weights & Biases run
type WandbMetadata struct {
	// Weights & Biases username or team that owns the project
	Entity string `json:"entity"`
	// Weights & Biases group used to organize related runs
	Group string `json:"group"`
	// Weights & Biases project containing the run
	Project string `json:"project"`
	// Unique identifier assigned to the run by Weights & Biases
	RunID string `json:"run_id"`
	// Human-readable name of the Weights & Biases run
	RunName string `json:"run_name"`
	// HTTPS URL for the Weights & Biases run
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entity      respjson.Field
		Group       respjson.Field
		Project     respjson.Field
		RunID       respjson.Field
		RunName     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WandbMetadata) RawJSON() string { return r.JSON.raw }
func (r *WandbMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WandbMetadata to a WandbMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WandbMetadataParam.Overrides()
func (r WandbMetadata) ToParam() WandbMetadataParam {
	return param.Override[WandbMetadataParam](json.RawMessage(r.RawJSON()))
}

// Details that associate a training session with a Weights & Biases run
type WandbMetadataParam struct {
	// Weights & Biases username or team that owns the project
	Entity param.Opt[string] `json:"entity,omitzero"`
	// Weights & Biases group used to organize related runs
	Group param.Opt[string] `json:"group,omitzero"`
	// Weights & Biases project containing the run
	Project param.Opt[string] `json:"project,omitzero"`
	// Unique identifier assigned to the run by Weights & Biases
	RunID param.Opt[string] `json:"run_id,omitzero"`
	// Human-readable name of the Weights & Biases run
	RunName param.Opt[string] `json:"run_name,omitzero"`
	// HTTPS URL for the Weights & Biases run
	URL param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r WandbMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow WandbMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WandbMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlSessionNewParams struct {
	// Model resource to attach the session to. The session runs on that resource's GPU
	// pods.
	ModelResourcesID string `json:"model_resources_id" api:"required"`
	// Optional display name used to identify the training session
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Whether to restore optimizer state and step from a training checkpoint. Omitted
	// or true restores them; false loads weights only with a fresh optimizer and
	// step 0. Not valid for inference or HuggingFace checkpoints, which have no
	// optimizer state.
	LoadOptimizer param.Opt[bool] `json:"load_optimizer,omitzero"`
	// Checkpoint ID to resume from
	ResumeFromCheckpointID param.Opt[string] `json:"resume_from_checkpoint_id,omitzero"`
	// HuggingFace repo (or hf://) to resume model weights from. Accepts either a full
	// model or a PEFT adapter directory. Mutually exclusive with
	// resume_from_checkpoint_id.
	ResumeFromHfCheckpoint param.Opt[string] `json:"resume_from_hf_checkpoint,omitzero"`
	// LoRA adapter configuration for the session
	LoraConfig LoraConfigParam `json:"lora_config,omitzero"`
	// Optional auxiliary metadata to associate with the training session
	Metadata SessionMetadataParam `json:"metadata,omitzero"`
	paramObj
}

func (r BetaRlSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaRlSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaRlSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlSessionListParams struct {
	// Cursor for pagination (ID of the last session from the previous page)
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Filter sessions in the current project by the creator ID. Pass "me" to show
	// sessions you created.
	CreatedBy param.Opt[string] `query:"created_by,omitzero" json:"-"`
	// Maximum number of sessions to return (1-100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter sessions by the model resource they are attached to
	ModelResourcesID param.Opt[string] `query:"model_resources_id,omitzero" json:"-"`
	// Status filters. When omitted, sessions in any status are returned.
	//
	// Any of "TRAINING_SESSION_STATUS_CREATING", "TRAINING_SESSION_STATUS_RUNNING",
	// "TRAINING_SESSION_STATUS_STOPPED", "TRAINING_SESSION_STATUS_STOPPING",
	// "TRAINING_SESSION_STATUS_ERROR", "TRAINING_SESSION_STATUS_EXPIRED".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaRlSessionListParams]'s query parameters as
// `url.Values`.
func (r BetaRlSessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
