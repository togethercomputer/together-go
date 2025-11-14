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

// FineTuningService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFineTuningService] method instead.
type FineTuningService struct {
	Options []option.RequestOption
}

// NewFineTuningService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFineTuningService(opts ...option.RequestOption) (r FineTuningService) {
	r = FineTuningService{}
	r.Options = opts
	return
}

// Create a fine-tuning job with the provided model and training data.
func (r *FineTuningService) New(ctx context.Context, body FineTuningNewParams, opts ...option.RequestOption) (res *FineTuningNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fine-tunes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List the metadata for a single fine-tuning job.
func (r *FineTuningService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTune, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the metadata for all fine-tuning jobs. Returns a list of
// FinetuneResponseTruncated objects.
func (r *FineTuningService) List(ctx context.Context, opts ...option.RequestOption) (res *FineTuningListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fine-tunes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a fine-tuning job.
func (r *FineTuningService) Delete(ctx context.Context, id string, body FineTuningDeleteParams, opts ...option.RequestOption) (res *FineTuningDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Cancel a currently running fine-tuning job. Returns a FinetuneResponseTruncated
// object.
func (r *FineTuningService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTuningCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Download a compressed fine-tuned model or checkpoint to local disk.
func (r *FineTuningService) Download(ctx context.Context, query FineTuningDownloadParams, opts ...option.RequestOption) (res *FineTuningDownloadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "finetune/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List the events for a single fine-tuning job.
func (r *FineTuningService) ListEvents(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTuningListEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the checkpoints for a single fine-tuning job.
func (r *FineTuningService) GetCheckpoints(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTuningGetCheckpointsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s/checkpoints", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CosineLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio,required"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64 `json:"num_cycles,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MinLrRatio  respjson.Field
		NumCycles   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CosineLrSchedulerArgs) RawJSON() string { return r.JSON.raw }
func (r *CosineLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CosineLrSchedulerArgs to a CosineLrSchedulerArgsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CosineLrSchedulerArgsParam.Overrides()
func (r CosineLrSchedulerArgs) ToParam() CosineLrSchedulerArgsParam {
	return param.Override[CosineLrSchedulerArgsParam](json.RawMessage(r.RawJSON()))
}

// The properties MinLrRatio, NumCycles are required.
type CosineLrSchedulerArgsParam struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio,required"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64 `json:"num_cycles,required"`
	paramObj
}

func (r CosineLrSchedulerArgsParam) MarshalJSON() (data []byte, err error) {
	type shadow CosineLrSchedulerArgsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CosineLrSchedulerArgsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTune struct {
	ID string `json:"id,required" format:"uuid"`
	// Any of "pending", "queued", "running", "compressing", "uploading",
	// "cancel_requested", "cancelled", "error", "completed".
	Status               FineTuneStatus              `json:"status,required"`
	BatchSize            FineTuneBatchSizeUnion      `json:"batch_size"`
	CreatedAt            string                      `json:"created_at"`
	EpochsCompleted      int64                       `json:"epochs_completed"`
	EvalSteps            int64                       `json:"eval_steps"`
	Events               []FineTuneEvent             `json:"events"`
	FromCheckpoint       string                      `json:"from_checkpoint"`
	FromHfModel          string                      `json:"from_hf_model"`
	HfModelRevision      string                      `json:"hf_model_revision"`
	JobID                string                      `json:"job_id"`
	LearningRate         float64                     `json:"learning_rate"`
	LrScheduler          LrScheduler                 `json:"lr_scheduler"`
	MaxGradNorm          float64                     `json:"max_grad_norm"`
	Model                string                      `json:"model"`
	ModelOutputName      string                      `json:"model_output_name"`
	ModelOutputPath      string                      `json:"model_output_path"`
	NCheckpoints         int64                       `json:"n_checkpoints"`
	NEpochs              int64                       `json:"n_epochs"`
	NEvals               int64                       `json:"n_evals"`
	ParamCount           int64                       `json:"param_count"`
	QueueDepth           int64                       `json:"queue_depth"`
	TokenCount           int64                       `json:"token_count"`
	TotalPrice           int64                       `json:"total_price"`
	TrainOnInputs        FineTuneTrainOnInputsUnion  `json:"train_on_inputs"`
	TrainingFile         string                      `json:"training_file"`
	TrainingMethod       FineTuneTrainingMethodUnion `json:"training_method"`
	TrainingType         FineTuneTrainingTypeUnion   `json:"training_type"`
	TrainingfileNumlines int64                       `json:"trainingfile_numlines"`
	TrainingfileSize     int64                       `json:"trainingfile_size"`
	UpdatedAt            string                      `json:"updated_at"`
	ValidationFile       string                      `json:"validation_file"`
	WandbProjectName     string                      `json:"wandb_project_name"`
	WandbURL             string                      `json:"wandb_url"`
	WarmupRatio          float64                     `json:"warmup_ratio"`
	WeightDecay          float64                     `json:"weight_decay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Status               respjson.Field
		BatchSize            respjson.Field
		CreatedAt            respjson.Field
		EpochsCompleted      respjson.Field
		EvalSteps            respjson.Field
		Events               respjson.Field
		FromCheckpoint       respjson.Field
		FromHfModel          respjson.Field
		HfModelRevision      respjson.Field
		JobID                respjson.Field
		LearningRate         respjson.Field
		LrScheduler          respjson.Field
		MaxGradNorm          respjson.Field
		Model                respjson.Field
		ModelOutputName      respjson.Field
		ModelOutputPath      respjson.Field
		NCheckpoints         respjson.Field
		NEpochs              respjson.Field
		NEvals               respjson.Field
		ParamCount           respjson.Field
		QueueDepth           respjson.Field
		TokenCount           respjson.Field
		TotalPrice           respjson.Field
		TrainOnInputs        respjson.Field
		TrainingFile         respjson.Field
		TrainingMethod       respjson.Field
		TrainingType         respjson.Field
		TrainingfileNumlines respjson.Field
		TrainingfileSize     respjson.Field
		UpdatedAt            respjson.Field
		ValidationFile       respjson.Field
		WandbProjectName     respjson.Field
		WandbURL             respjson.Field
		WarmupRatio          respjson.Field
		WeightDecay          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTune) RawJSON() string { return r.JSON.raw }
func (r *FineTune) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuneStatus string

const (
	FineTuneStatusPending         FineTuneStatus = "pending"
	FineTuneStatusQueued          FineTuneStatus = "queued"
	FineTuneStatusRunning         FineTuneStatus = "running"
	FineTuneStatusCompressing     FineTuneStatus = "compressing"
	FineTuneStatusUploading       FineTuneStatus = "uploading"
	FineTuneStatusCancelRequested FineTuneStatus = "cancel_requested"
	FineTuneStatusCancelled       FineTuneStatus = "cancelled"
	FineTuneStatusError           FineTuneStatus = "error"
	FineTuneStatusCompleted       FineTuneStatus = "completed"
)

// FineTuneBatchSizeUnion contains all possible properties and values from [int64],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfFineTuneBatchSizeString]
type FineTuneBatchSizeUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfFineTuneBatchSizeString string `json:",inline"`
	JSON                      struct {
		OfInt                     respjson.Field
		OfFineTuneBatchSizeString respjson.Field
		raw                       string
	} `json:"-"`
}

func (u FineTuneBatchSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuneBatchSizeUnion) AsFineTuneBatchSizeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuneBatchSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuneBatchSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuneBatchSizeString string

const (
	FineTuneBatchSizeStringMax FineTuneBatchSizeString = "max"
)

// FineTuneTrainOnInputsUnion contains all possible properties and values from
// [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFineTuneTrainOnInputsString]
type FineTuneTrainOnInputsUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfFineTuneTrainOnInputsString string `json:",inline"`
	JSON                          struct {
		OfBool                        respjson.Field
		OfFineTuneTrainOnInputsString respjson.Field
		raw                           string
	} `json:"-"`
}

func (u FineTuneTrainOnInputsUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuneTrainOnInputsUnion) AsFineTuneTrainOnInputsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuneTrainOnInputsUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuneTrainOnInputsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuneTrainOnInputsString string

const (
	FineTuneTrainOnInputsStringAuto FineTuneTrainOnInputsString = "auto"
)

// FineTuneTrainingMethodUnion contains all possible properties and values from
// [TrainingMethodSft], [TrainingMethodDpo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuneTrainingMethodUnion struct {
	Method string `json:"method"`
	// This field is from variant [TrainingMethodSft].
	TrainOnInputs TrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs"`
	// This field is from variant [TrainingMethodDpo].
	DpoBeta float64 `json:"dpo_beta"`
	// This field is from variant [TrainingMethodDpo].
	DpoNormalizeLogratiosByLength bool `json:"dpo_normalize_logratios_by_length"`
	// This field is from variant [TrainingMethodDpo].
	DpoReferenceFree bool `json:"dpo_reference_free"`
	// This field is from variant [TrainingMethodDpo].
	RpoAlpha float64 `json:"rpo_alpha"`
	// This field is from variant [TrainingMethodDpo].
	SimpoGamma float64 `json:"simpo_gamma"`
	JSON       struct {
		Method                        respjson.Field
		TrainOnInputs                 respjson.Field
		DpoBeta                       respjson.Field
		DpoNormalizeLogratiosByLength respjson.Field
		DpoReferenceFree              respjson.Field
		RpoAlpha                      respjson.Field
		SimpoGamma                    respjson.Field
		raw                           string
	} `json:"-"`
}

func (u FineTuneTrainingMethodUnion) AsTrainingMethodSft() (v TrainingMethodSft) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuneTrainingMethodUnion) AsTrainingMethodDpo() (v TrainingMethodDpo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuneTrainingMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuneTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuneTrainingTypeUnion contains all possible properties and values from
// [FullTrainingType], [LoRaTrainingType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuneTrainingTypeUnion struct {
	Type string `json:"type"`
	// This field is from variant [LoRaTrainingType].
	LoraAlpha int64 `json:"lora_alpha"`
	// This field is from variant [LoRaTrainingType].
	LoraR int64 `json:"lora_r"`
	// This field is from variant [LoRaTrainingType].
	LoraDropout float64 `json:"lora_dropout"`
	// This field is from variant [LoRaTrainingType].
	LoraTrainableModules string `json:"lora_trainable_modules"`
	JSON                 struct {
		Type                 respjson.Field
		LoraAlpha            respjson.Field
		LoraR                respjson.Field
		LoraDropout          respjson.Field
		LoraTrainableModules respjson.Field
		raw                  string
	} `json:"-"`
}

func (u FineTuneTrainingTypeUnion) AsFullTrainingType() (v FullTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuneTrainingTypeUnion) AsLoRaTrainingType() (v LoRaTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuneTrainingTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuneTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuneEvent struct {
	CheckpointPath string `json:"checkpoint_path,required"`
	CreatedAt      string `json:"created_at,required"`
	Hash           string `json:"hash,required"`
	Message        string `json:"message,required"`
	ModelPath      string `json:"model_path,required"`
	// Any of "fine-tune-event".
	Object         FineTuneEventObject `json:"object,required"`
	ParamCount     int64               `json:"param_count,required"`
	Step           int64               `json:"step,required"`
	TokenCount     int64               `json:"token_count,required"`
	TotalSteps     int64               `json:"total_steps,required"`
	TrainingOffset int64               `json:"training_offset,required"`
	// Any of "job_pending", "job_start", "job_stopped", "model_downloading",
	// "model_download_complete", "training_data_downloading",
	// "training_data_download_complete", "validation_data_downloading",
	// "validation_data_download_complete", "wandb_init", "training_start",
	// "checkpoint_save", "billing_limit", "epoch_complete", "training_complete",
	// "model_compressing", "model_compression_complete", "model_uploading",
	// "model_upload_complete", "job_complete", "job_error", "cancel_requested",
	// "job_restarted", "refund", "warning".
	Type     FineTuneEventType `json:"type,required"`
	WandbURL string            `json:"wandb_url,required"`
	// Any of "info", "warning", "error", "legacy_info", "legacy_iwarning",
	// "legacy_ierror".
	Level FineTuneEventLevel `json:"level,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckpointPath respjson.Field
		CreatedAt      respjson.Field
		Hash           respjson.Field
		Message        respjson.Field
		ModelPath      respjson.Field
		Object         respjson.Field
		ParamCount     respjson.Field
		Step           respjson.Field
		TokenCount     respjson.Field
		TotalSteps     respjson.Field
		TrainingOffset respjson.Field
		Type           respjson.Field
		WandbURL       respjson.Field
		Level          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuneEvent) RawJSON() string { return r.JSON.raw }
func (r *FineTuneEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuneEventObject string

const (
	FineTuneEventObjectFineTuneEvent FineTuneEventObject = "fine-tune-event"
)

type FineTuneEventType string

const (
	FineTuneEventTypeJobPending                     FineTuneEventType = "job_pending"
	FineTuneEventTypeJobStart                       FineTuneEventType = "job_start"
	FineTuneEventTypeJobStopped                     FineTuneEventType = "job_stopped"
	FineTuneEventTypeModelDownloading               FineTuneEventType = "model_downloading"
	FineTuneEventTypeModelDownloadComplete          FineTuneEventType = "model_download_complete"
	FineTuneEventTypeTrainingDataDownloading        FineTuneEventType = "training_data_downloading"
	FineTuneEventTypeTrainingDataDownloadComplete   FineTuneEventType = "training_data_download_complete"
	FineTuneEventTypeValidationDataDownloading      FineTuneEventType = "validation_data_downloading"
	FineTuneEventTypeValidationDataDownloadComplete FineTuneEventType = "validation_data_download_complete"
	FineTuneEventTypeWandbInit                      FineTuneEventType = "wandb_init"
	FineTuneEventTypeTrainingStart                  FineTuneEventType = "training_start"
	FineTuneEventTypeCheckpointSave                 FineTuneEventType = "checkpoint_save"
	FineTuneEventTypeBillingLimit                   FineTuneEventType = "billing_limit"
	FineTuneEventTypeEpochComplete                  FineTuneEventType = "epoch_complete"
	FineTuneEventTypeTrainingComplete               FineTuneEventType = "training_complete"
	FineTuneEventTypeModelCompressing               FineTuneEventType = "model_compressing"
	FineTuneEventTypeModelCompressionComplete       FineTuneEventType = "model_compression_complete"
	FineTuneEventTypeModelUploading                 FineTuneEventType = "model_uploading"
	FineTuneEventTypeModelUploadComplete            FineTuneEventType = "model_upload_complete"
	FineTuneEventTypeJobComplete                    FineTuneEventType = "job_complete"
	FineTuneEventTypeJobError                       FineTuneEventType = "job_error"
	FineTuneEventTypeCancelRequested                FineTuneEventType = "cancel_requested"
	FineTuneEventTypeJobRestarted                   FineTuneEventType = "job_restarted"
	FineTuneEventTypeRefund                         FineTuneEventType = "refund"
	FineTuneEventTypeWarning                        FineTuneEventType = "warning"
)

type FineTuneEventLevel string

const (
	FineTuneEventLevelInfo           FineTuneEventLevel = "info"
	FineTuneEventLevelWarning        FineTuneEventLevel = "warning"
	FineTuneEventLevelError          FineTuneEventLevel = "error"
	FineTuneEventLevelLegacyInfo     FineTuneEventLevel = "legacy_info"
	FineTuneEventLevelLegacyIwarning FineTuneEventLevel = "legacy_iwarning"
	FineTuneEventLevelLegacyIerror   FineTuneEventLevel = "legacy_ierror"
)

type FullTrainingType struct {
	// Any of "Full".
	Type FullTrainingTypeType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FullTrainingType) RawJSON() string { return r.JSON.raw }
func (r *FullTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FullTrainingType to a FullTrainingTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FullTrainingTypeParam.Overrides()
func (r FullTrainingType) ToParam() FullTrainingTypeParam {
	return param.Override[FullTrainingTypeParam](json.RawMessage(r.RawJSON()))
}

type FullTrainingTypeType string

const (
	FullTrainingTypeTypeFull FullTrainingTypeType = "Full"
)

// The property Type is required.
type FullTrainingTypeParam struct {
	// Any of "Full".
	Type FullTrainingTypeType `json:"type,omitzero,required"`
	paramObj
}

func (r FullTrainingTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow FullTrainingTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FullTrainingTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinearLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MinLrRatio  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinearLrSchedulerArgs) RawJSON() string { return r.JSON.raw }
func (r *LinearLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LinearLrSchedulerArgs to a LinearLrSchedulerArgsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LinearLrSchedulerArgsParam.Overrides()
func (r LinearLrSchedulerArgs) ToParam() LinearLrSchedulerArgsParam {
	return param.Override[LinearLrSchedulerArgsParam](json.RawMessage(r.RawJSON()))
}

type LinearLrSchedulerArgsParam struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio param.Opt[float64] `json:"min_lr_ratio,omitzero"`
	paramObj
}

func (r LinearLrSchedulerArgsParam) MarshalJSON() (data []byte, err error) {
	type shadow LinearLrSchedulerArgsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinearLrSchedulerArgsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoRaTrainingType struct {
	LoraAlpha int64 `json:"lora_alpha,required"`
	LoraR     int64 `json:"lora_r,required"`
	// Any of "Lora".
	Type                 LoRaTrainingTypeType `json:"type,required"`
	LoraDropout          float64              `json:"lora_dropout"`
	LoraTrainableModules string               `json:"lora_trainable_modules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LoraAlpha            respjson.Field
		LoraR                respjson.Field
		Type                 respjson.Field
		LoraDropout          respjson.Field
		LoraTrainableModules respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoRaTrainingType) RawJSON() string { return r.JSON.raw }
func (r *LoRaTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LoRaTrainingType to a LoRaTrainingTypeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LoRaTrainingTypeParam.Overrides()
func (r LoRaTrainingType) ToParam() LoRaTrainingTypeParam {
	return param.Override[LoRaTrainingTypeParam](json.RawMessage(r.RawJSON()))
}

type LoRaTrainingTypeType string

const (
	LoRaTrainingTypeTypeLora LoRaTrainingTypeType = "Lora"
)

// The properties LoraAlpha, LoraR, Type are required.
type LoRaTrainingTypeParam struct {
	LoraAlpha int64 `json:"lora_alpha,required"`
	LoraR     int64 `json:"lora_r,required"`
	// Any of "Lora".
	Type                 LoRaTrainingTypeType `json:"type,omitzero,required"`
	LoraDropout          param.Opt[float64]   `json:"lora_dropout,omitzero"`
	LoraTrainableModules param.Opt[string]    `json:"lora_trainable_modules,omitzero"`
	paramObj
}

func (r LoRaTrainingTypeParam) MarshalJSON() (data []byte, err error) {
	type shadow LoRaTrainingTypeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoRaTrainingTypeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LrScheduler struct {
	// Any of "linear", "cosine".
	LrSchedulerType LrSchedulerLrSchedulerType      `json:"lr_scheduler_type,required"`
	LrSchedulerArgs LrSchedulerLrSchedulerArgsUnion `json:"lr_scheduler_args"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LrSchedulerType respjson.Field
		LrSchedulerArgs respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LrScheduler) RawJSON() string { return r.JSON.raw }
func (r *LrScheduler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LrScheduler to a LrSchedulerParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LrSchedulerParam.Overrides()
func (r LrScheduler) ToParam() LrSchedulerParam {
	return param.Override[LrSchedulerParam](json.RawMessage(r.RawJSON()))
}

type LrSchedulerLrSchedulerType string

const (
	LrSchedulerLrSchedulerTypeLinear LrSchedulerLrSchedulerType = "linear"
	LrSchedulerLrSchedulerTypeCosine LrSchedulerLrSchedulerType = "cosine"
)

// LrSchedulerLrSchedulerArgsUnion contains all possible properties and values from
// [LinearLrSchedulerArgs], [CosineLrSchedulerArgs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LrSchedulerLrSchedulerArgsUnion struct {
	MinLrRatio float64 `json:"min_lr_ratio"`
	// This field is from variant [CosineLrSchedulerArgs].
	NumCycles float64 `json:"num_cycles"`
	JSON      struct {
		MinLrRatio respjson.Field
		NumCycles  respjson.Field
		raw        string
	} `json:"-"`
}

func (u LrSchedulerLrSchedulerArgsUnion) AsLinearLrSchedulerArgs() (v LinearLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LrSchedulerLrSchedulerArgsUnion) AsCosineLrSchedulerArgs() (v CosineLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LrSchedulerLrSchedulerArgsUnion) RawJSON() string { return u.JSON.raw }

func (r *LrSchedulerLrSchedulerArgsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LrSchedulerType is required.
type LrSchedulerParam struct {
	// Any of "linear", "cosine".
	LrSchedulerType LrSchedulerLrSchedulerType           `json:"lr_scheduler_type,omitzero,required"`
	LrSchedulerArgs LrSchedulerLrSchedulerArgsUnionParam `json:"lr_scheduler_args,omitzero"`
	paramObj
}

func (r LrSchedulerParam) MarshalJSON() (data []byte, err error) {
	type shadow LrSchedulerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LrSchedulerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LrSchedulerLrSchedulerArgsUnionParam struct {
	OfLinearLrSchedulerArgs *LinearLrSchedulerArgsParam `json:",omitzero,inline"`
	OfCosineLrSchedulerArgs *CosineLrSchedulerArgsParam `json:",omitzero,inline"`
	paramUnion
}

func (u LrSchedulerLrSchedulerArgsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLinearLrSchedulerArgs, u.OfCosineLrSchedulerArgs)
}
func (u *LrSchedulerLrSchedulerArgsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LrSchedulerLrSchedulerArgsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfLinearLrSchedulerArgs) {
		return u.OfLinearLrSchedulerArgs
	} else if !param.IsOmitted(u.OfCosineLrSchedulerArgs) {
		return u.OfCosineLrSchedulerArgs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LrSchedulerLrSchedulerArgsUnionParam) GetNumCycles() *float64 {
	if vt := u.OfCosineLrSchedulerArgs; vt != nil {
		return &vt.NumCycles
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LrSchedulerLrSchedulerArgsUnionParam) GetMinLrRatio() *float64 {
	if vt := u.OfLinearLrSchedulerArgs; vt != nil && vt.MinLrRatio.Valid() {
		return &vt.MinLrRatio.Value
	} else if vt := u.OfCosineLrSchedulerArgs; vt != nil {
		return (*float64)(&vt.MinLrRatio)
	}
	return nil
}

type TrainingMethodDpo struct {
	// Any of "dpo".
	Method                        TrainingMethodDpoMethod `json:"method,required"`
	DpoBeta                       float64                 `json:"dpo_beta"`
	DpoNormalizeLogratiosByLength bool                    `json:"dpo_normalize_logratios_by_length"`
	DpoReferenceFree              bool                    `json:"dpo_reference_free"`
	RpoAlpha                      float64                 `json:"rpo_alpha"`
	SimpoGamma                    float64                 `json:"simpo_gamma"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method                        respjson.Field
		DpoBeta                       respjson.Field
		DpoNormalizeLogratiosByLength respjson.Field
		DpoReferenceFree              respjson.Field
		RpoAlpha                      respjson.Field
		SimpoGamma                    respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrainingMethodDpo) RawJSON() string { return r.JSON.raw }
func (r *TrainingMethodDpo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TrainingMethodDpo to a TrainingMethodDpoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TrainingMethodDpoParam.Overrides()
func (r TrainingMethodDpo) ToParam() TrainingMethodDpoParam {
	return param.Override[TrainingMethodDpoParam](json.RawMessage(r.RawJSON()))
}

type TrainingMethodDpoMethod string

const (
	TrainingMethodDpoMethodDpo TrainingMethodDpoMethod = "dpo"
)

// The property Method is required.
type TrainingMethodDpoParam struct {
	// Any of "dpo".
	Method                        TrainingMethodDpoMethod `json:"method,omitzero,required"`
	DpoBeta                       param.Opt[float64]      `json:"dpo_beta,omitzero"`
	DpoNormalizeLogratiosByLength param.Opt[bool]         `json:"dpo_normalize_logratios_by_length,omitzero"`
	DpoReferenceFree              param.Opt[bool]         `json:"dpo_reference_free,omitzero"`
	RpoAlpha                      param.Opt[float64]      `json:"rpo_alpha,omitzero"`
	SimpoGamma                    param.Opt[float64]      `json:"simpo_gamma,omitzero"`
	paramObj
}

func (r TrainingMethodDpoParam) MarshalJSON() (data []byte, err error) {
	type shadow TrainingMethodDpoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrainingMethodDpoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrainingMethodSft struct {
	// Any of "sft".
	Method TrainingMethodSftMethod `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs TrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method        respjson.Field
		TrainOnInputs respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrainingMethodSft) RawJSON() string { return r.JSON.raw }
func (r *TrainingMethodSft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TrainingMethodSft to a TrainingMethodSftParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TrainingMethodSftParam.Overrides()
func (r TrainingMethodSft) ToParam() TrainingMethodSftParam {
	return param.Override[TrainingMethodSftParam](json.RawMessage(r.RawJSON()))
}

type TrainingMethodSftMethod string

const (
	TrainingMethodSftMethodSft TrainingMethodSftMethod = "sft"
)

// TrainingMethodSftTrainOnInputsUnion contains all possible properties and values
// from [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfTrainingMethodSftTrainOnInputsString]
type TrainingMethodSftTrainOnInputsUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfTrainingMethodSftTrainOnInputsString string `json:",inline"`
	JSON                                   struct {
		OfBool                                 respjson.Field
		OfTrainingMethodSftTrainOnInputsString respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u TrainingMethodSftTrainOnInputsUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TrainingMethodSftTrainOnInputsUnion) AsTrainingMethodSftTrainOnInputsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TrainingMethodSftTrainOnInputsUnion) RawJSON() string { return u.JSON.raw }

func (r *TrainingMethodSftTrainOnInputsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrainingMethodSftTrainOnInputsString string

const (
	TrainingMethodSftTrainOnInputsStringAuto TrainingMethodSftTrainOnInputsString = "auto"
)

// The properties Method, TrainOnInputs are required.
type TrainingMethodSftParam struct {
	// Any of "sft".
	Method TrainingMethodSftMethod `json:"method,omitzero,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs TrainingMethodSftTrainOnInputsUnionParam `json:"train_on_inputs,omitzero,required"`
	paramObj
}

func (r TrainingMethodSftParam) MarshalJSON() (data []byte, err error) {
	type shadow TrainingMethodSftParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrainingMethodSftParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TrainingMethodSftTrainOnInputsUnionParam struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfTrainingMethodSftTrainOnInputsString)
	OfTrainingMethodSftTrainOnInputsString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u TrainingMethodSftTrainOnInputsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfTrainingMethodSftTrainOnInputsString)
}
func (u *TrainingMethodSftTrainOnInputsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TrainingMethodSftTrainOnInputsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfTrainingMethodSftTrainOnInputsString) {
		return &u.OfTrainingMethodSftTrainOnInputsString
	}
	return nil
}

// A truncated version of the fine-tune response, used for POST /fine-tunes, GET
// /fine-tunes and POST /fine-tunes/{id}/cancel endpoints
type FineTuningNewResponse struct {
	// Unique identifier for the fine-tune job
	ID string `json:"id,required"`
	// Creation timestamp of the fine-tune job
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Any of "pending", "queued", "running", "compressing", "uploading",
	// "cancel_requested", "cancelled", "error", "completed".
	Status FineTuningNewResponseStatus `json:"status,required"`
	// Last update timestamp of the fine-tune job
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Batch size used for training
	BatchSize int64 `json:"batch_size"`
	// Events related to this fine-tune job
	Events []FineTuneEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
	// Hugging Face Hub repo to start training from
	FromHfModel string `json:"from_hf_model"`
	// The revision of the Hugging Face Hub model to continue training from
	HfModelRevision string `json:"hf_model_revision"`
	// Learning rate used for training
	LearningRate float64 `json:"learning_rate"`
	// Learning rate scheduler configuration
	LrScheduler LrScheduler `json:"lr_scheduler"`
	// Maximum gradient norm for clipping
	MaxGradNorm float64 `json:"max_grad_norm"`
	// Base model used for fine-tuning
	Model           string `json:"model"`
	ModelOutputName string `json:"model_output_name"`
	// Number of checkpoints saved during training
	NCheckpoints int64 `json:"n_checkpoints"`
	// Number of training epochs
	NEpochs int64 `json:"n_epochs"`
	// Number of evaluations during training
	NEvals int64 `json:"n_evals"`
	// Owner address information
	OwnerAddress string `json:"owner_address"`
	// Suffix added to the fine-tuned model name
	Suffix string `json:"suffix"`
	// Count of tokens processed
	TokenCount int64 `json:"token_count"`
	// Total price for the fine-tuning job
	TotalPrice int64 `json:"total_price"`
	// File-ID of the training file
	TrainingFile string `json:"training_file"`
	// Method of training used
	TrainingMethod FineTuningNewResponseTrainingMethodUnion `json:"training_method"`
	// Type of training used (full or LoRA)
	TrainingType FineTuningNewResponseTrainingTypeUnion `json:"training_type"`
	// Identifier for the user who created the job
	UserID string `json:"user_id"`
	// File-ID of the validation file
	ValidationFile string `json:"validation_file"`
	// Weights & Biases run name
	WandbName string `json:"wandb_name"`
	// Weights & Biases project name
	WandbProjectName string `json:"wandb_project_name"`
	// Ratio of warmup steps
	WarmupRatio float64 `json:"warmup_ratio"`
	// Weight decay value used
	WeightDecay float64 `json:"weight_decay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		BatchSize        respjson.Field
		Events           respjson.Field
		FromCheckpoint   respjson.Field
		FromHfModel      respjson.Field
		HfModelRevision  respjson.Field
		LearningRate     respjson.Field
		LrScheduler      respjson.Field
		MaxGradNorm      respjson.Field
		Model            respjson.Field
		ModelOutputName  respjson.Field
		NCheckpoints     respjson.Field
		NEpochs          respjson.Field
		NEvals           respjson.Field
		OwnerAddress     respjson.Field
		Suffix           respjson.Field
		TokenCount       respjson.Field
		TotalPrice       respjson.Field
		TrainingFile     respjson.Field
		TrainingMethod   respjson.Field
		TrainingType     respjson.Field
		UserID           respjson.Field
		ValidationFile   respjson.Field
		WandbName        respjson.Field
		WandbProjectName respjson.Field
		WarmupRatio      respjson.Field
		WeightDecay      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FineTuningNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningNewResponseStatus string

const (
	FineTuningNewResponseStatusPending         FineTuningNewResponseStatus = "pending"
	FineTuningNewResponseStatusQueued          FineTuningNewResponseStatus = "queued"
	FineTuningNewResponseStatusRunning         FineTuningNewResponseStatus = "running"
	FineTuningNewResponseStatusCompressing     FineTuningNewResponseStatus = "compressing"
	FineTuningNewResponseStatusUploading       FineTuningNewResponseStatus = "uploading"
	FineTuningNewResponseStatusCancelRequested FineTuningNewResponseStatus = "cancel_requested"
	FineTuningNewResponseStatusCancelled       FineTuningNewResponseStatus = "cancelled"
	FineTuningNewResponseStatusError           FineTuningNewResponseStatus = "error"
	FineTuningNewResponseStatusCompleted       FineTuningNewResponseStatus = "completed"
)

// FineTuningNewResponseTrainingMethodUnion contains all possible properties and
// values from [TrainingMethodSft], [TrainingMethodDpo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningNewResponseTrainingMethodUnion struct {
	Method string `json:"method"`
	// This field is from variant [TrainingMethodSft].
	TrainOnInputs TrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs"`
	// This field is from variant [TrainingMethodDpo].
	DpoBeta float64 `json:"dpo_beta"`
	// This field is from variant [TrainingMethodDpo].
	DpoNormalizeLogratiosByLength bool `json:"dpo_normalize_logratios_by_length"`
	// This field is from variant [TrainingMethodDpo].
	DpoReferenceFree bool `json:"dpo_reference_free"`
	// This field is from variant [TrainingMethodDpo].
	RpoAlpha float64 `json:"rpo_alpha"`
	// This field is from variant [TrainingMethodDpo].
	SimpoGamma float64 `json:"simpo_gamma"`
	JSON       struct {
		Method                        respjson.Field
		TrainOnInputs                 respjson.Field
		DpoBeta                       respjson.Field
		DpoNormalizeLogratiosByLength respjson.Field
		DpoReferenceFree              respjson.Field
		RpoAlpha                      respjson.Field
		SimpoGamma                    respjson.Field
		raw                           string
	} `json:"-"`
}

func (u FineTuningNewResponseTrainingMethodUnion) AsTrainingMethodSft() (v TrainingMethodSft) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningNewResponseTrainingMethodUnion) AsTrainingMethodDpo() (v TrainingMethodDpo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningNewResponseTrainingMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningNewResponseTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningNewResponseTrainingTypeUnion contains all possible properties and
// values from [FullTrainingType], [LoRaTrainingType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningNewResponseTrainingTypeUnion struct {
	Type string `json:"type"`
	// This field is from variant [LoRaTrainingType].
	LoraAlpha int64 `json:"lora_alpha"`
	// This field is from variant [LoRaTrainingType].
	LoraR int64 `json:"lora_r"`
	// This field is from variant [LoRaTrainingType].
	LoraDropout float64 `json:"lora_dropout"`
	// This field is from variant [LoRaTrainingType].
	LoraTrainableModules string `json:"lora_trainable_modules"`
	JSON                 struct {
		Type                 respjson.Field
		LoraAlpha            respjson.Field
		LoraR                respjson.Field
		LoraDropout          respjson.Field
		LoraTrainableModules respjson.Field
		raw                  string
	} `json:"-"`
}

func (u FineTuningNewResponseTrainingTypeUnion) AsFullTrainingType() (v FullTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningNewResponseTrainingTypeUnion) AsLoRaTrainingType() (v LoRaTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningNewResponseTrainingTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningNewResponseTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListResponse struct {
	Data []FineTuningListResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningListResponse) RawJSON() string { return r.JSON.raw }
func (r *FineTuningListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A truncated version of the fine-tune response, used for POST /fine-tunes, GET
// /fine-tunes and POST /fine-tunes/{id}/cancel endpoints
type FineTuningListResponseData struct {
	// Unique identifier for the fine-tune job
	ID string `json:"id,required"`
	// Creation timestamp of the fine-tune job
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Any of "pending", "queued", "running", "compressing", "uploading",
	// "cancel_requested", "cancelled", "error", "completed".
	Status string `json:"status,required"`
	// Last update timestamp of the fine-tune job
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Batch size used for training
	BatchSize int64 `json:"batch_size"`
	// Events related to this fine-tune job
	Events []FineTuneEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
	// Hugging Face Hub repo to start training from
	FromHfModel string `json:"from_hf_model"`
	// The revision of the Hugging Face Hub model to continue training from
	HfModelRevision string `json:"hf_model_revision"`
	// Learning rate used for training
	LearningRate float64 `json:"learning_rate"`
	// Learning rate scheduler configuration
	LrScheduler LrScheduler `json:"lr_scheduler"`
	// Maximum gradient norm for clipping
	MaxGradNorm float64 `json:"max_grad_norm"`
	// Base model used for fine-tuning
	Model           string `json:"model"`
	ModelOutputName string `json:"model_output_name"`
	// Number of checkpoints saved during training
	NCheckpoints int64 `json:"n_checkpoints"`
	// Number of training epochs
	NEpochs int64 `json:"n_epochs"`
	// Number of evaluations during training
	NEvals int64 `json:"n_evals"`
	// Owner address information
	OwnerAddress string `json:"owner_address"`
	// Suffix added to the fine-tuned model name
	Suffix string `json:"suffix"`
	// Count of tokens processed
	TokenCount int64 `json:"token_count"`
	// Total price for the fine-tuning job
	TotalPrice int64 `json:"total_price"`
	// File-ID of the training file
	TrainingFile string `json:"training_file"`
	// Method of training used
	TrainingMethod FineTuningListResponseDataTrainingMethodUnion `json:"training_method"`
	// Type of training used (full or LoRA)
	TrainingType FineTuningListResponseDataTrainingTypeUnion `json:"training_type"`
	// Identifier for the user who created the job
	UserID string `json:"user_id"`
	// File-ID of the validation file
	ValidationFile string `json:"validation_file"`
	// Weights & Biases run name
	WandbName string `json:"wandb_name"`
	// Weights & Biases project name
	WandbProjectName string `json:"wandb_project_name"`
	// Ratio of warmup steps
	WarmupRatio float64 `json:"warmup_ratio"`
	// Weight decay value used
	WeightDecay float64 `json:"weight_decay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		BatchSize        respjson.Field
		Events           respjson.Field
		FromCheckpoint   respjson.Field
		FromHfModel      respjson.Field
		HfModelRevision  respjson.Field
		LearningRate     respjson.Field
		LrScheduler      respjson.Field
		MaxGradNorm      respjson.Field
		Model            respjson.Field
		ModelOutputName  respjson.Field
		NCheckpoints     respjson.Field
		NEpochs          respjson.Field
		NEvals           respjson.Field
		OwnerAddress     respjson.Field
		Suffix           respjson.Field
		TokenCount       respjson.Field
		TotalPrice       respjson.Field
		TrainingFile     respjson.Field
		TrainingMethod   respjson.Field
		TrainingType     respjson.Field
		UserID           respjson.Field
		ValidationFile   respjson.Field
		WandbName        respjson.Field
		WandbProjectName respjson.Field
		WarmupRatio      respjson.Field
		WeightDecay      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningListResponseData) RawJSON() string { return r.JSON.raw }
func (r *FineTuningListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningListResponseDataTrainingMethodUnion contains all possible properties
// and values from [TrainingMethodSft], [TrainingMethodDpo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningListResponseDataTrainingMethodUnion struct {
	Method string `json:"method"`
	// This field is from variant [TrainingMethodSft].
	TrainOnInputs TrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs"`
	// This field is from variant [TrainingMethodDpo].
	DpoBeta float64 `json:"dpo_beta"`
	// This field is from variant [TrainingMethodDpo].
	DpoNormalizeLogratiosByLength bool `json:"dpo_normalize_logratios_by_length"`
	// This field is from variant [TrainingMethodDpo].
	DpoReferenceFree bool `json:"dpo_reference_free"`
	// This field is from variant [TrainingMethodDpo].
	RpoAlpha float64 `json:"rpo_alpha"`
	// This field is from variant [TrainingMethodDpo].
	SimpoGamma float64 `json:"simpo_gamma"`
	JSON       struct {
		Method                        respjson.Field
		TrainOnInputs                 respjson.Field
		DpoBeta                       respjson.Field
		DpoNormalizeLogratiosByLength respjson.Field
		DpoReferenceFree              respjson.Field
		RpoAlpha                      respjson.Field
		SimpoGamma                    respjson.Field
		raw                           string
	} `json:"-"`
}

func (u FineTuningListResponseDataTrainingMethodUnion) AsTrainingMethodSft() (v TrainingMethodSft) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningListResponseDataTrainingMethodUnion) AsTrainingMethodDpo() (v TrainingMethodDpo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningListResponseDataTrainingMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningListResponseDataTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningListResponseDataTrainingTypeUnion contains all possible properties and
// values from [FullTrainingType], [LoRaTrainingType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningListResponseDataTrainingTypeUnion struct {
	Type string `json:"type"`
	// This field is from variant [LoRaTrainingType].
	LoraAlpha int64 `json:"lora_alpha"`
	// This field is from variant [LoRaTrainingType].
	LoraR int64 `json:"lora_r"`
	// This field is from variant [LoRaTrainingType].
	LoraDropout float64 `json:"lora_dropout"`
	// This field is from variant [LoRaTrainingType].
	LoraTrainableModules string `json:"lora_trainable_modules"`
	JSON                 struct {
		Type                 respjson.Field
		LoraAlpha            respjson.Field
		LoraR                respjson.Field
		LoraDropout          respjson.Field
		LoraTrainableModules respjson.Field
		raw                  string
	} `json:"-"`
}

func (u FineTuningListResponseDataTrainingTypeUnion) AsFullTrainingType() (v FullTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningListResponseDataTrainingTypeUnion) AsLoRaTrainingType() (v LoRaTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningListResponseDataTrainingTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningListResponseDataTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningDeleteResponse struct {
	// Message indicating the result of the deletion
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FineTuningDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A truncated version of the fine-tune response, used for POST /fine-tunes, GET
// /fine-tunes and POST /fine-tunes/{id}/cancel endpoints
type FineTuningCancelResponse struct {
	// Unique identifier for the fine-tune job
	ID string `json:"id,required"`
	// Creation timestamp of the fine-tune job
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Any of "pending", "queued", "running", "compressing", "uploading",
	// "cancel_requested", "cancelled", "error", "completed".
	Status FineTuningCancelResponseStatus `json:"status,required"`
	// Last update timestamp of the fine-tune job
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Batch size used for training
	BatchSize int64 `json:"batch_size"`
	// Events related to this fine-tune job
	Events []FineTuneEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
	// Hugging Face Hub repo to start training from
	FromHfModel string `json:"from_hf_model"`
	// The revision of the Hugging Face Hub model to continue training from
	HfModelRevision string `json:"hf_model_revision"`
	// Learning rate used for training
	LearningRate float64 `json:"learning_rate"`
	// Learning rate scheduler configuration
	LrScheduler LrScheduler `json:"lr_scheduler"`
	// Maximum gradient norm for clipping
	MaxGradNorm float64 `json:"max_grad_norm"`
	// Base model used for fine-tuning
	Model           string `json:"model"`
	ModelOutputName string `json:"model_output_name"`
	// Number of checkpoints saved during training
	NCheckpoints int64 `json:"n_checkpoints"`
	// Number of training epochs
	NEpochs int64 `json:"n_epochs"`
	// Number of evaluations during training
	NEvals int64 `json:"n_evals"`
	// Owner address information
	OwnerAddress string `json:"owner_address"`
	// Suffix added to the fine-tuned model name
	Suffix string `json:"suffix"`
	// Count of tokens processed
	TokenCount int64 `json:"token_count"`
	// Total price for the fine-tuning job
	TotalPrice int64 `json:"total_price"`
	// File-ID of the training file
	TrainingFile string `json:"training_file"`
	// Method of training used
	TrainingMethod FineTuningCancelResponseTrainingMethodUnion `json:"training_method"`
	// Type of training used (full or LoRA)
	TrainingType FineTuningCancelResponseTrainingTypeUnion `json:"training_type"`
	// Identifier for the user who created the job
	UserID string `json:"user_id"`
	// File-ID of the validation file
	ValidationFile string `json:"validation_file"`
	// Weights & Biases run name
	WandbName string `json:"wandb_name"`
	// Weights & Biases project name
	WandbProjectName string `json:"wandb_project_name"`
	// Ratio of warmup steps
	WarmupRatio float64 `json:"warmup_ratio"`
	// Weight decay value used
	WeightDecay float64 `json:"weight_decay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		BatchSize        respjson.Field
		Events           respjson.Field
		FromCheckpoint   respjson.Field
		FromHfModel      respjson.Field
		HfModelRevision  respjson.Field
		LearningRate     respjson.Field
		LrScheduler      respjson.Field
		MaxGradNorm      respjson.Field
		Model            respjson.Field
		ModelOutputName  respjson.Field
		NCheckpoints     respjson.Field
		NEpochs          respjson.Field
		NEvals           respjson.Field
		OwnerAddress     respjson.Field
		Suffix           respjson.Field
		TokenCount       respjson.Field
		TotalPrice       respjson.Field
		TrainingFile     respjson.Field
		TrainingMethod   respjson.Field
		TrainingType     respjson.Field
		UserID           respjson.Field
		ValidationFile   respjson.Field
		WandbName        respjson.Field
		WandbProjectName respjson.Field
		WarmupRatio      respjson.Field
		WeightDecay      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *FineTuningCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningCancelResponseStatus string

const (
	FineTuningCancelResponseStatusPending         FineTuningCancelResponseStatus = "pending"
	FineTuningCancelResponseStatusQueued          FineTuningCancelResponseStatus = "queued"
	FineTuningCancelResponseStatusRunning         FineTuningCancelResponseStatus = "running"
	FineTuningCancelResponseStatusCompressing     FineTuningCancelResponseStatus = "compressing"
	FineTuningCancelResponseStatusUploading       FineTuningCancelResponseStatus = "uploading"
	FineTuningCancelResponseStatusCancelRequested FineTuningCancelResponseStatus = "cancel_requested"
	FineTuningCancelResponseStatusCancelled       FineTuningCancelResponseStatus = "cancelled"
	FineTuningCancelResponseStatusError           FineTuningCancelResponseStatus = "error"
	FineTuningCancelResponseStatusCompleted       FineTuningCancelResponseStatus = "completed"
)

// FineTuningCancelResponseTrainingMethodUnion contains all possible properties and
// values from [TrainingMethodSft], [TrainingMethodDpo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningCancelResponseTrainingMethodUnion struct {
	Method string `json:"method"`
	// This field is from variant [TrainingMethodSft].
	TrainOnInputs TrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs"`
	// This field is from variant [TrainingMethodDpo].
	DpoBeta float64 `json:"dpo_beta"`
	// This field is from variant [TrainingMethodDpo].
	DpoNormalizeLogratiosByLength bool `json:"dpo_normalize_logratios_by_length"`
	// This field is from variant [TrainingMethodDpo].
	DpoReferenceFree bool `json:"dpo_reference_free"`
	// This field is from variant [TrainingMethodDpo].
	RpoAlpha float64 `json:"rpo_alpha"`
	// This field is from variant [TrainingMethodDpo].
	SimpoGamma float64 `json:"simpo_gamma"`
	JSON       struct {
		Method                        respjson.Field
		TrainOnInputs                 respjson.Field
		DpoBeta                       respjson.Field
		DpoNormalizeLogratiosByLength respjson.Field
		DpoReferenceFree              respjson.Field
		RpoAlpha                      respjson.Field
		SimpoGamma                    respjson.Field
		raw                           string
	} `json:"-"`
}

func (u FineTuningCancelResponseTrainingMethodUnion) AsTrainingMethodSft() (v TrainingMethodSft) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningCancelResponseTrainingMethodUnion) AsTrainingMethodDpo() (v TrainingMethodDpo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningCancelResponseTrainingMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningCancelResponseTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningCancelResponseTrainingTypeUnion contains all possible properties and
// values from [FullTrainingType], [LoRaTrainingType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningCancelResponseTrainingTypeUnion struct {
	Type string `json:"type"`
	// This field is from variant [LoRaTrainingType].
	LoraAlpha int64 `json:"lora_alpha"`
	// This field is from variant [LoRaTrainingType].
	LoraR int64 `json:"lora_r"`
	// This field is from variant [LoRaTrainingType].
	LoraDropout float64 `json:"lora_dropout"`
	// This field is from variant [LoRaTrainingType].
	LoraTrainableModules string `json:"lora_trainable_modules"`
	JSON                 struct {
		Type                 respjson.Field
		LoraAlpha            respjson.Field
		LoraR                respjson.Field
		LoraDropout          respjson.Field
		LoraTrainableModules respjson.Field
		raw                  string
	} `json:"-"`
}

func (u FineTuningCancelResponseTrainingTypeUnion) AsFullTrainingType() (v FullTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningCancelResponseTrainingTypeUnion) AsLoRaTrainingType() (v LoRaTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningCancelResponseTrainingTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningCancelResponseTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningDownloadResponse struct {
	ID             string `json:"id"`
	CheckpointStep int64  `json:"checkpoint_step"`
	Filename       string `json:"filename"`
	// Any of "local".
	Object FineTuningDownloadResponseObject `json:"object,nullable"`
	Size   int64                            `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CheckpointStep respjson.Field
		Filename       respjson.Field
		Object         respjson.Field
		Size           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningDownloadResponse) RawJSON() string { return r.JSON.raw }
func (r *FineTuningDownloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningDownloadResponseObject string

const (
	FineTuningDownloadResponseObjectLocal FineTuningDownloadResponseObject = "local"
)

type FineTuningListEventsResponse struct {
	Data []FineTuneEvent `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningListEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *FineTuningListEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningGetCheckpointsResponse struct {
	Data []FineTuningGetCheckpointsResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningGetCheckpointsResponse) RawJSON() string { return r.JSON.raw }
func (r *FineTuningGetCheckpointsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningGetCheckpointsResponseData struct {
	CheckpointType string `json:"checkpoint_type,required"`
	CreatedAt      string `json:"created_at,required"`
	Path           string `json:"path,required"`
	Step           int64  `json:"step,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckpointType respjson.Field
		CreatedAt      respjson.Field
		Path           respjson.Field
		Step           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningGetCheckpointsResponseData) RawJSON() string { return r.JSON.raw }
func (r *FineTuningGetCheckpointsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningNewParams struct {
	// Name of the base model to run fine-tune job on
	Model string `json:"model,required"`
	// File-ID of a training file uploaded to the Together API
	TrainingFile string `json:"training_file,required"`
	// The checkpoint identifier to continue training from a previous fine-tuning job.
	// Format is `{$JOB_ID}` or `{$OUTPUT_MODEL_NAME}` or `{$JOB_ID}:{$STEP}` or
	// `{$OUTPUT_MODEL_NAME}:{$STEP}`. The step value is optional; without it, the
	// final checkpoint will be used.
	FromCheckpoint param.Opt[string] `json:"from_checkpoint,omitzero"`
	// The Hugging Face Hub repo to start training from. Should be as close as possible
	// to the base model (specified by the `model` argument) in terms of architecture
	// and size.
	FromHfModel param.Opt[string] `json:"from_hf_model,omitzero"`
	// The API token for the Hugging Face Hub.
	HfAPIToken param.Opt[string] `json:"hf_api_token,omitzero"`
	// The revision of the Hugging Face Hub model to continue training from. E.g.,
	// hf_model_revision=main (default, used if the argument is not provided) or
	// hf_model_revision='607a30d783dfa663caf39e06633721c8d4cfcd7e' (specific commit).
	HfModelRevision param.Opt[string] `json:"hf_model_revision,omitzero"`
	// The name of the Hugging Face repository to upload the fine-tuned model to.
	HfOutputRepoName param.Opt[string] `json:"hf_output_repo_name,omitzero"`
	// Controls how quickly the model adapts to new information (too high may cause
	// instability, too low may slow convergence)
	LearningRate param.Opt[float64] `json:"learning_rate,omitzero"`
	// Max gradient norm to be used for gradient clipping. Set to 0 to disable.
	MaxGradNorm param.Opt[float64] `json:"max_grad_norm,omitzero"`
	// Number of intermediate model versions saved during training for evaluation
	NCheckpoints param.Opt[int64] `json:"n_checkpoints,omitzero"`
	// Number of complete passes through the training dataset (higher values may
	// improve results but increase cost and risk of overfitting)
	NEpochs param.Opt[int64] `json:"n_epochs,omitzero"`
	// Number of evaluations to be run on a given validation set during training
	NEvals param.Opt[int64] `json:"n_evals,omitzero"`
	// Suffix that will be added to your fine-tuned model name
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// File-ID of a validation file uploaded to the Together API
	ValidationFile param.Opt[string] `json:"validation_file,omitzero"`
	// Integration key for tracking experiments and model metrics on W&B platform
	WandbAPIKey param.Opt[string] `json:"wandb_api_key,omitzero"`
	// The base URL of a dedicated Weights & Biases instance.
	WandbBaseURL param.Opt[string] `json:"wandb_base_url,omitzero"`
	// The Weights & Biases name for your run.
	WandbName param.Opt[string] `json:"wandb_name,omitzero"`
	// The Weights & Biases project for your run. If not specified, will use `together`
	// as the project name.
	WandbProjectName param.Opt[string] `json:"wandb_project_name,omitzero"`
	// The percent of steps at the start of training to linearly increase the learning
	// rate.
	WarmupRatio param.Opt[float64] `json:"warmup_ratio,omitzero"`
	// Weight decay. Regularization parameter for the optimizer.
	WeightDecay param.Opt[float64] `json:"weight_decay,omitzero"`
	// Number of training examples processed together (larger batches use more memory
	// but may train faster). Defaults to "max". We use training optimizations like
	// packing, so the effective batch size may be different than the value you set.
	BatchSize FineTuningNewParamsBatchSizeUnion `json:"batch_size,omitzero"`
	// The learning rate scheduler to use. It specifies how the learning rate is
	// adjusted during training.
	LrScheduler LrSchedulerParam `json:"lr_scheduler,omitzero"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FineTuningNewParamsTrainOnInputsUnion `json:"train_on_inputs,omitzero"`
	// The training method to use. 'sft' for Supervised Fine-Tuning or 'dpo' for Direct
	// Preference Optimization.
	TrainingMethod FineTuningNewParamsTrainingMethodUnion `json:"training_method,omitzero"`
	TrainingType   FineTuningNewParamsTrainingTypeUnion   `json:"training_type,omitzero"`
	paramObj
}

func (r FineTuningNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FineTuningNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FineTuningNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FineTuningNewParamsBatchSizeUnion struct {
	OfInt param.Opt[int64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFineTuningNewsBatchSizeString)
	OfFineTuningNewsBatchSizeString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u FineTuningNewParamsBatchSizeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfFineTuningNewsBatchSizeString)
}
func (u *FineTuningNewParamsBatchSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FineTuningNewParamsBatchSizeUnion) asAny() any {
	if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfFineTuningNewsBatchSizeString) {
		return &u.OfFineTuningNewsBatchSizeString
	}
	return nil
}

type FineTuningNewParamsBatchSizeString string

const (
	FineTuningNewParamsBatchSizeStringMax FineTuningNewParamsBatchSizeString = "max"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FineTuningNewParamsTrainOnInputsUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFineTuningNewsTrainOnInputsString)
	OfFineTuningNewsTrainOnInputsString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u FineTuningNewParamsTrainOnInputsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFineTuningNewsTrainOnInputsString)
}
func (u *FineTuningNewParamsTrainOnInputsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FineTuningNewParamsTrainOnInputsUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFineTuningNewsTrainOnInputsString) {
		return &u.OfFineTuningNewsTrainOnInputsString
	}
	return nil
}

type FineTuningNewParamsTrainOnInputsString string

const (
	FineTuningNewParamsTrainOnInputsStringAuto FineTuningNewParamsTrainOnInputsString = "auto"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FineTuningNewParamsTrainingMethodUnion struct {
	OfTrainingMethodSft *TrainingMethodSftParam `json:",omitzero,inline"`
	OfTrainingMethodDpo *TrainingMethodDpoParam `json:",omitzero,inline"`
	paramUnion
}

func (u FineTuningNewParamsTrainingMethodUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTrainingMethodSft, u.OfTrainingMethodDpo)
}
func (u *FineTuningNewParamsTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FineTuningNewParamsTrainingMethodUnion) asAny() any {
	if !param.IsOmitted(u.OfTrainingMethodSft) {
		return u.OfTrainingMethodSft
	} else if !param.IsOmitted(u.OfTrainingMethodDpo) {
		return u.OfTrainingMethodDpo
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetTrainOnInputs() *TrainingMethodSftTrainOnInputsUnionParam {
	if vt := u.OfTrainingMethodSft; vt != nil {
		return &vt.TrainOnInputs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetDpoBeta() *float64 {
	if vt := u.OfTrainingMethodDpo; vt != nil && vt.DpoBeta.Valid() {
		return &vt.DpoBeta.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetDpoNormalizeLogratiosByLength() *bool {
	if vt := u.OfTrainingMethodDpo; vt != nil && vt.DpoNormalizeLogratiosByLength.Valid() {
		return &vt.DpoNormalizeLogratiosByLength.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetDpoReferenceFree() *bool {
	if vt := u.OfTrainingMethodDpo; vt != nil && vt.DpoReferenceFree.Valid() {
		return &vt.DpoReferenceFree.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetRpoAlpha() *float64 {
	if vt := u.OfTrainingMethodDpo; vt != nil && vt.RpoAlpha.Valid() {
		return &vt.RpoAlpha.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetSimpoGamma() *float64 {
	if vt := u.OfTrainingMethodDpo; vt != nil && vt.SimpoGamma.Valid() {
		return &vt.SimpoGamma.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetMethod() *string {
	if vt := u.OfTrainingMethodSft; vt != nil {
		return (*string)(&vt.Method)
	} else if vt := u.OfTrainingMethodDpo; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FineTuningNewParamsTrainingTypeUnion struct {
	OfFullTrainingType *FullTrainingTypeParam `json:",omitzero,inline"`
	OfLoRaTrainingType *LoRaTrainingTypeParam `json:",omitzero,inline"`
	paramUnion
}

func (u FineTuningNewParamsTrainingTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFullTrainingType, u.OfLoRaTrainingType)
}
func (u *FineTuningNewParamsTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FineTuningNewParamsTrainingTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfFullTrainingType) {
		return u.OfFullTrainingType
	} else if !param.IsOmitted(u.OfLoRaTrainingType) {
		return u.OfLoRaTrainingType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetLoraAlpha() *int64 {
	if vt := u.OfLoRaTrainingType; vt != nil {
		return &vt.LoraAlpha
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetLoraR() *int64 {
	if vt := u.OfLoRaTrainingType; vt != nil {
		return &vt.LoraR
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetLoraDropout() *float64 {
	if vt := u.OfLoRaTrainingType; vt != nil && vt.LoraDropout.Valid() {
		return &vt.LoraDropout.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetLoraTrainableModules() *string {
	if vt := u.OfLoRaTrainingType; vt != nil && vt.LoraTrainableModules.Valid() {
		return &vt.LoraTrainableModules.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetType() *string {
	if vt := u.OfFullTrainingType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLoRaTrainingType; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

type FineTuningDeleteParams struct {
	Force bool `query:"force,required" json:"-"`
	paramObj
}

// URLQuery serializes [FineTuningDeleteParams]'s query parameters as `url.Values`.
func (r FineTuningDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FineTuningDownloadParams struct {
	// Fine-tune ID to download. A string that starts with `ft-`.
	FtID string `query:"ft_id,required" json:"-"`
	// Specifies step number for checkpoint to download. Ignores `checkpoint` value if
	// set.
	CheckpointStep param.Opt[int64] `query:"checkpoint_step,omitzero" json:"-"`
	// Specifies output file name for downloaded model. Defaults to
	// `$PWD/{model_name}.{extension}`.
	Output param.Opt[string] `query:"output,omitzero" json:"-"`
	// Specifies checkpoint type to download - `merged` vs `adapter`. This field is
	// required if the checkpoint_step is not set.
	//
	// Any of "merged", "adapter".
	Checkpoint FineTuningDownloadParamsCheckpoint `query:"checkpoint,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FineTuningDownloadParams]'s query parameters as
// `url.Values`.
func (r FineTuningDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies checkpoint type to download - `merged` vs `adapter`. This field is
// required if the checkpoint_step is not set.
type FineTuningDownloadParamsCheckpoint string

const (
	FineTuningDownloadParamsCheckpointMerged  FineTuningDownloadParamsCheckpoint = "merged"
	FineTuningDownloadParamsCheckpointAdapter FineTuningDownloadParamsCheckpoint = "adapter"
)
