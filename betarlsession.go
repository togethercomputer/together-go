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
func (r *BetaRlSessionService) New(ctx context.Context, body BetaRlSessionNewParams, opts ...option.RequestOption) (res *TrainingSession, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rl/training-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Gets a training session by its ID and returns its details.
func (r *BetaRlSessionService) Get(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *TrainingSession, err error) {
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
func (r *BetaRlSessionService) List(ctx context.Context, query BetaRlSessionListParams, opts ...option.RequestOption) (res *TrainingSessionsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rl/training-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Stops a training session.
func (r *BetaRlSessionService) Stop(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *TrainingSession, err error) {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelName    respjson.Field
		RegisteredAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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
	// Whether to enable LoRA fine-tuning. If false, full fine-tuning is used.
	Enable bool `json:"enable"`
	// Rank of the LoRA adapter
	Rank int64 `json:"rank"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alpha       respjson.Field
		Dropout     respjson.Field
		Enable      respjson.Field
		Rank        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// LoRA adapter configuration
type LoraConfigParam struct {
	// Alpha of the LoRA adapter
	Alpha param.Opt[int64] `json:"alpha,omitzero"`
	// Dropout of the LoRA adapter
	Dropout param.Opt[float64] `json:"dropout,omitzero"`
	// Whether to enable LoRA fine-tuning. If false, full fine-tuning is used.
	Enable param.Opt[bool] `json:"enable,omitzero"`
	// Rank of the LoRA adapter
	Rank param.Opt[int64] `json:"rank,omitzero"`
	paramObj
}

func (r LoraConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow LoraConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoraConfigParam) UnmarshalJSON(data []byte) error {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Step        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// A training session and its current state
type TrainingSession struct {
	// ID of the training session
	ID string `json:"id" api:"required"`
	// Timestamp when the training session was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ID of the user who created the training session
	CreatedBy string `json:"created_by" api:"required"`
	// List of saved inference checkpoints for this session
	InferenceCheckpoints []InferenceCheckpoint `json:"inference_checkpoints" api:"required"`
	// LoRA adapter configuration for this session
	LoraConfig LoraConfig `json:"lora_config" api:"required"`
	// Model resource this session is attached to. The session runs on that resource's
	// GPU pods.
	ModelResourcesID string `json:"model_resources_id" api:"required"`
	// Status of the training session
	//
	// Any of "TRAINING_SESSION_STATUS_UNSPECIFIED",
	// "TRAINING_SESSION_STATUS_CREATING", "TRAINING_SESSION_STATUS_RUNNING",
	// "TRAINING_SESSION_STATUS_STOPPED", "TRAINING_SESSION_STATUS_STOPPING",
	// "TRAINING_SESSION_STATUS_ERROR", "TRAINING_SESSION_STATUS_EXPIRED".
	Status TrainingSessionStatus `json:"status" api:"required"`
	// Current training step
	Step TrainingSessionStepUnion `json:"step" api:"required"`
	// List of saved training checkpoints for this session
	TrainingCheckpoints []TrainingCheckpoint `json:"training_checkpoints" api:"required"`
	// Timestamp when the training session was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Structured detail for the training session's current error. Set when the session
	// is in an error state.
	Error TrainingSessionError `json:"error"`
	// Checkpoint ID this session was resumed from
	ResumeFromCheckpointID string `json:"resume_from_checkpoint_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CreatedAt              respjson.Field
		CreatedBy              respjson.Field
		InferenceCheckpoints   respjson.Field
		LoraConfig             respjson.Field
		ModelResourcesID       respjson.Field
		Status                 respjson.Field
		Step                   respjson.Field
		TrainingCheckpoints    respjson.Field
		UpdatedAt              respjson.Field
		Error                  respjson.Field
		ResumeFromCheckpointID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrainingSession) RawJSON() string { return r.JSON.raw }
func (r *TrainingSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TrainingSessionStepUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type TrainingSessionStepUnion struct {
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

func (u TrainingSessionStepUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TrainingSessionStepUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TrainingSessionStepUnion) RawJSON() string { return u.JSON.raw }

func (r *TrainingSessionStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured detail for the training session's current error
type TrainingSessionError struct {
	// Finite machine-readable reason code for UI branching
	//
	// Any of "TRAINING_SESSION_ERROR_CODE_RESOURCE_UNAVAILABLE",
	// "TRAINING_SESSION_ERROR_CODE_RESOURCE_AT_CAPACITY",
	// "TRAINING_SESSION_ERROR_CODE_TIMED_OUT",
	// "TRAINING_SESSION_ERROR_CODE_SESSION_FAILED".
	Code TrainingSessionErrorCode `json:"code" api:"required"`
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
func (r TrainingSessionError) RawJSON() string { return r.JSON.raw }
func (r *TrainingSessionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Finite machine-readable training session lifecycle error code
type TrainingSessionErrorCode string

const (
	TrainingSessionErrorCodeTrainingSessionErrorCodeResourceUnavailable TrainingSessionErrorCode = "TRAINING_SESSION_ERROR_CODE_RESOURCE_UNAVAILABLE"
	TrainingSessionErrorCodeTrainingSessionErrorCodeResourceAtCapacity  TrainingSessionErrorCode = "TRAINING_SESSION_ERROR_CODE_RESOURCE_AT_CAPACITY"
	TrainingSessionErrorCodeTrainingSessionErrorCodeTimedOut            TrainingSessionErrorCode = "TRAINING_SESSION_ERROR_CODE_TIMED_OUT"
	TrainingSessionErrorCodeTrainingSessionErrorCodeSessionFailed       TrainingSessionErrorCode = "TRAINING_SESSION_ERROR_CODE_SESSION_FAILED"
)

// Status of the training session
type TrainingSessionStatus string

const (
	TrainingSessionStatusTrainingSessionStatusUnspecified TrainingSessionStatus = "TRAINING_SESSION_STATUS_UNSPECIFIED"
	TrainingSessionStatusTrainingSessionStatusCreating    TrainingSessionStatus = "TRAINING_SESSION_STATUS_CREATING"
	TrainingSessionStatusTrainingSessionStatusRunning     TrainingSessionStatus = "TRAINING_SESSION_STATUS_RUNNING"
	TrainingSessionStatusTrainingSessionStatusStopped     TrainingSessionStatus = "TRAINING_SESSION_STATUS_STOPPED"
	TrainingSessionStatusTrainingSessionStatusStopping    TrainingSessionStatus = "TRAINING_SESSION_STATUS_STOPPING"
	TrainingSessionStatusTrainingSessionStatusError       TrainingSessionStatus = "TRAINING_SESSION_STATUS_ERROR"
	TrainingSessionStatusTrainingSessionStatusExpired     TrainingSessionStatus = "TRAINING_SESSION_STATUS_EXPIRED"
)

// Paginated list of training sessions
type TrainingSessionsListResponse struct {
	// List of training sessions
	Data []TrainingSession `json:"data"`
	// Pagination metadata
	Meta TrainingSessionsListResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrainingSessionsListResponse) RawJSON() string { return r.JSON.raw }
func (r *TrainingSessionsListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type TrainingSessionsListResponseMeta struct {
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
func (r TrainingSessionsListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *TrainingSessionsListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlSessionNewParams struct {
	// Model resource to attach the session to. The session runs on that resource's GPU
	// pods.
	ModelResourcesID string `json:"model_resources_id" api:"required"`
	// Checkpoint ID to resume from
	ResumeFromCheckpointID param.Opt[string] `json:"resume_from_checkpoint_id,omitzero"`
	// HuggingFace repo (or hf://) to resume model weights from. Accepts either a full
	// model or a PEFT adapter directory. Mutually exclusive with
	// resume_from_checkpoint_id.
	ResumeFromHfCheckpoint param.Opt[string] `json:"resume_from_hf_checkpoint,omitzero"`
	// LoRA adapter configuration for the session
	LoraConfig LoraConfigParam `json:"lora_config,omitzero"`
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
