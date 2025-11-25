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
func (r *FineTuningService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FinetuneResponse, err error) {
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

// Download a compressed fine-tuned model or checkpoint.
func (r *FineTuningService) Content(ctx context.Context, query FineTuningContentParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	path := "finetune/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List the checkpoints for a single fine-tuning job.
func (r *FineTuningService) ListCheckpoints(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTuningListCheckpointsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s/checkpoints", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

type FinetuneEvent struct {
	CheckpointPath string `json:"checkpoint_path,required"`
	CreatedAt      string `json:"created_at,required"`
	Hash           string `json:"hash,required"`
	Message        string `json:"message,required"`
	ModelPath      string `json:"model_path,required"`
	// Any of "fine-tune-event".
	Object         FinetuneEventObject `json:"object,required"`
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
	Type     FinetuneEventType `json:"type,required"`
	WandbURL string            `json:"wandb_url,required"`
	// Any of "info", "warning", "error", "legacy_info", "legacy_iwarning",
	// "legacy_ierror".
	Level FinetuneEventLevel `json:"level,nullable"`
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
func (r FinetuneEvent) RawJSON() string { return r.JSON.raw }
func (r *FinetuneEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneEventObject string

const (
	FinetuneEventObjectFineTuneEvent FinetuneEventObject = "fine-tune-event"
)

type FinetuneEventLevel string

const (
	FinetuneEventLevelInfo           FinetuneEventLevel = "info"
	FinetuneEventLevelWarning        FinetuneEventLevel = "warning"
	FinetuneEventLevelError          FinetuneEventLevel = "error"
	FinetuneEventLevelLegacyInfo     FinetuneEventLevel = "legacy_info"
	FinetuneEventLevelLegacyIwarning FinetuneEventLevel = "legacy_iwarning"
	FinetuneEventLevelLegacyIerror   FinetuneEventLevel = "legacy_ierror"
)

type FinetuneEventType string

const (
	FinetuneEventTypeJobPending                     FinetuneEventType = "job_pending"
	FinetuneEventTypeJobStart                       FinetuneEventType = "job_start"
	FinetuneEventTypeJobStopped                     FinetuneEventType = "job_stopped"
	FinetuneEventTypeModelDownloading               FinetuneEventType = "model_downloading"
	FinetuneEventTypeModelDownloadComplete          FinetuneEventType = "model_download_complete"
	FinetuneEventTypeTrainingDataDownloading        FinetuneEventType = "training_data_downloading"
	FinetuneEventTypeTrainingDataDownloadComplete   FinetuneEventType = "training_data_download_complete"
	FinetuneEventTypeValidationDataDownloading      FinetuneEventType = "validation_data_downloading"
	FinetuneEventTypeValidationDataDownloadComplete FinetuneEventType = "validation_data_download_complete"
	FinetuneEventTypeWandbInit                      FinetuneEventType = "wandb_init"
	FinetuneEventTypeTrainingStart                  FinetuneEventType = "training_start"
	FinetuneEventTypeCheckpointSave                 FinetuneEventType = "checkpoint_save"
	FinetuneEventTypeBillingLimit                   FinetuneEventType = "billing_limit"
	FinetuneEventTypeEpochComplete                  FinetuneEventType = "epoch_complete"
	FinetuneEventTypeTrainingComplete               FinetuneEventType = "training_complete"
	FinetuneEventTypeModelCompressing               FinetuneEventType = "model_compressing"
	FinetuneEventTypeModelCompressionComplete       FinetuneEventType = "model_compression_complete"
	FinetuneEventTypeModelUploading                 FinetuneEventType = "model_uploading"
	FinetuneEventTypeModelUploadComplete            FinetuneEventType = "model_upload_complete"
	FinetuneEventTypeJobComplete                    FinetuneEventType = "job_complete"
	FinetuneEventTypeJobError                       FinetuneEventType = "job_error"
	FinetuneEventTypeCancelRequested                FinetuneEventType = "cancel_requested"
	FinetuneEventTypeJobRestarted                   FinetuneEventType = "job_restarted"
	FinetuneEventTypeRefund                         FinetuneEventType = "refund"
	FinetuneEventTypeWarning                        FinetuneEventType = "warning"
)

type FinetuneResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Any of "pending", "queued", "running", "compressing", "uploading",
	// "cancel_requested", "cancelled", "error", "completed".
	Status               FinetuneResponseStatus              `json:"status,required"`
	BatchSize            FinetuneResponseBatchSizeUnion      `json:"batch_size"`
	CreatedAt            string                              `json:"created_at"`
	EpochsCompleted      int64                               `json:"epochs_completed"`
	EvalSteps            int64                               `json:"eval_steps"`
	Events               []FinetuneEvent                     `json:"events"`
	FromCheckpoint       string                              `json:"from_checkpoint"`
	FromHfModel          string                              `json:"from_hf_model"`
	HfModelRevision      string                              `json:"hf_model_revision"`
	JobID                string                              `json:"job_id"`
	LearningRate         float64                             `json:"learning_rate"`
	LrScheduler          FinetuneResponseLrScheduler         `json:"lr_scheduler"`
	MaxGradNorm          float64                             `json:"max_grad_norm"`
	Model                string                              `json:"model"`
	ModelOutputName      string                              `json:"model_output_name"`
	ModelOutputPath      string                              `json:"model_output_path"`
	NCheckpoints         int64                               `json:"n_checkpoints"`
	NEpochs              int64                               `json:"n_epochs"`
	NEvals               int64                               `json:"n_evals"`
	ParamCount           int64                               `json:"param_count"`
	QueueDepth           int64                               `json:"queue_depth"`
	TokenCount           int64                               `json:"token_count"`
	TotalPrice           int64                               `json:"total_price"`
	TrainOnInputs        FinetuneResponseTrainOnInputsUnion  `json:"train_on_inputs"`
	TrainingFile         string                              `json:"training_file"`
	TrainingMethod       FinetuneResponseTrainingMethodUnion `json:"training_method"`
	TrainingType         FinetuneResponseTrainingTypeUnion   `json:"training_type"`
	TrainingfileNumlines int64                               `json:"trainingfile_numlines"`
	TrainingfileSize     int64                               `json:"trainingfile_size"`
	UpdatedAt            string                              `json:"updated_at"`
	ValidationFile       string                              `json:"validation_file"`
	WandbProjectName     string                              `json:"wandb_project_name"`
	WandbURL             string                              `json:"wandb_url"`
	WarmupRatio          float64                             `json:"warmup_ratio"`
	WeightDecay          float64                             `json:"weight_decay"`
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
func (r FinetuneResponse) RawJSON() string { return r.JSON.raw }
func (r *FinetuneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneResponseStatus string

const (
	FinetuneResponseStatusPending         FinetuneResponseStatus = "pending"
	FinetuneResponseStatusQueued          FinetuneResponseStatus = "queued"
	FinetuneResponseStatusRunning         FinetuneResponseStatus = "running"
	FinetuneResponseStatusCompressing     FinetuneResponseStatus = "compressing"
	FinetuneResponseStatusUploading       FinetuneResponseStatus = "uploading"
	FinetuneResponseStatusCancelRequested FinetuneResponseStatus = "cancel_requested"
	FinetuneResponseStatusCancelled       FinetuneResponseStatus = "cancelled"
	FinetuneResponseStatusError           FinetuneResponseStatus = "error"
	FinetuneResponseStatusCompleted       FinetuneResponseStatus = "completed"
)

// FinetuneResponseBatchSizeUnion contains all possible properties and values from
// [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfFinetuneResponseBatchSizeString]
type FinetuneResponseBatchSizeUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfFinetuneResponseBatchSizeString string `json:",inline"`
	JSON                              struct {
		OfInt                             respjson.Field
		OfFinetuneResponseBatchSizeString respjson.Field
		raw                               string
	} `json:"-"`
}

func (u FinetuneResponseBatchSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FinetuneResponseBatchSizeUnion) AsFinetuneResponseBatchSizeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FinetuneResponseBatchSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *FinetuneResponseBatchSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneResponseBatchSizeString string

const (
	FinetuneResponseBatchSizeStringMax FinetuneResponseBatchSizeString = "max"
)

type FinetuneResponseLrScheduler struct {
	// Any of "linear", "cosine".
	LrSchedulerType string                                          `json:"lr_scheduler_type,required"`
	LrSchedulerArgs FinetuneResponseLrSchedulerLrSchedulerArgsUnion `json:"lr_scheduler_args"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LrSchedulerType respjson.Field
		LrSchedulerArgs respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FinetuneResponseLrScheduler) RawJSON() string { return r.JSON.raw }
func (r *FinetuneResponseLrScheduler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FinetuneResponseLrSchedulerLrSchedulerArgsUnion contains all possible properties
// and values from
// [FinetuneResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs],
// [FinetuneResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FinetuneResponseLrSchedulerLrSchedulerArgsUnion struct {
	MinLrRatio float64 `json:"min_lr_ratio"`
	// This field is from variant
	// [FinetuneResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
	NumCycles float64 `json:"num_cycles"`
	JSON      struct {
		MinLrRatio respjson.Field
		NumCycles  respjson.Field
		raw        string
	} `json:"-"`
}

func (u FinetuneResponseLrSchedulerLrSchedulerArgsUnion) AsFinetuneResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs() (v FinetuneResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FinetuneResponseLrSchedulerLrSchedulerArgsUnion) AsFinetuneResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs() (v FinetuneResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FinetuneResponseLrSchedulerLrSchedulerArgsUnion) RawJSON() string { return u.JSON.raw }

func (r *FinetuneResponseLrSchedulerLrSchedulerArgsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
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
func (r FinetuneResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) RawJSON() string {
	return r.JSON.raw
}
func (r *FinetuneResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
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
func (r FinetuneResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) RawJSON() string {
	return r.JSON.raw
}
func (r *FinetuneResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FinetuneResponseTrainOnInputsUnion contains all possible properties and values
// from [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFinetuneResponseTrainOnInputsString]
type FinetuneResponseTrainOnInputsUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfFinetuneResponseTrainOnInputsString string `json:",inline"`
	JSON                                  struct {
		OfBool                                respjson.Field
		OfFinetuneResponseTrainOnInputsString respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u FinetuneResponseTrainOnInputsUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FinetuneResponseTrainOnInputsUnion) AsFinetuneResponseTrainOnInputsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FinetuneResponseTrainOnInputsUnion) RawJSON() string { return u.JSON.raw }

func (r *FinetuneResponseTrainOnInputsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneResponseTrainOnInputsString string

const (
	FinetuneResponseTrainOnInputsStringAuto FinetuneResponseTrainOnInputsString = "auto"
)

// FinetuneResponseTrainingMethodUnion contains all possible properties and values
// from [FinetuneResponseTrainingMethodTrainingMethodSft],
// [FinetuneResponseTrainingMethodTrainingMethodDpo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FinetuneResponseTrainingMethodUnion struct {
	Method string `json:"method"`
	// This field is from variant [FinetuneResponseTrainingMethodTrainingMethodSft].
	TrainOnInputs FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs"`
	// This field is from variant [FinetuneResponseTrainingMethodTrainingMethodDpo].
	DpoBeta float64 `json:"dpo_beta"`
	// This field is from variant [FinetuneResponseTrainingMethodTrainingMethodDpo].
	DpoNormalizeLogratiosByLength bool `json:"dpo_normalize_logratios_by_length"`
	// This field is from variant [FinetuneResponseTrainingMethodTrainingMethodDpo].
	DpoReferenceFree bool `json:"dpo_reference_free"`
	// This field is from variant [FinetuneResponseTrainingMethodTrainingMethodDpo].
	RpoAlpha float64 `json:"rpo_alpha"`
	// This field is from variant [FinetuneResponseTrainingMethodTrainingMethodDpo].
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

func (u FinetuneResponseTrainingMethodUnion) AsFinetuneResponseTrainingMethodTrainingMethodSft() (v FinetuneResponseTrainingMethodTrainingMethodSft) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FinetuneResponseTrainingMethodUnion) AsFinetuneResponseTrainingMethodTrainingMethodDpo() (v FinetuneResponseTrainingMethodTrainingMethodDpo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FinetuneResponseTrainingMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *FinetuneResponseTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneResponseTrainingMethodTrainingMethodSft struct {
	// Any of "sft".
	Method string `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method        respjson.Field
		TrainOnInputs respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FinetuneResponseTrainingMethodTrainingMethodSft) RawJSON() string { return r.JSON.raw }
func (r *FinetuneResponseTrainingMethodTrainingMethodSft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion contains all
// possible properties and values from [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool
// OfFinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsString]
type FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfFinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsString string `json:",inline"`
	JSON                                                                 struct {
		OfBool                                                               respjson.Field
		OfFinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsString respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) AsFinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FinetuneResponseTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

type FinetuneResponseTrainingMethodTrainingMethodDpo struct {
	// Any of "dpo".
	Method                        string  `json:"method,required"`
	DpoBeta                       float64 `json:"dpo_beta"`
	DpoNormalizeLogratiosByLength bool    `json:"dpo_normalize_logratios_by_length"`
	DpoReferenceFree              bool    `json:"dpo_reference_free"`
	RpoAlpha                      float64 `json:"rpo_alpha"`
	SimpoGamma                    float64 `json:"simpo_gamma"`
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
func (r FinetuneResponseTrainingMethodTrainingMethodDpo) RawJSON() string { return r.JSON.raw }
func (r *FinetuneResponseTrainingMethodTrainingMethodDpo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FinetuneResponseTrainingTypeUnion contains all possible properties and values
// from [FinetuneResponseTrainingTypeFullTrainingType],
// [FinetuneResponseTrainingTypeLoRaTrainingType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FinetuneResponseTrainingTypeUnion struct {
	Type string `json:"type"`
	// This field is from variant [FinetuneResponseTrainingTypeLoRaTrainingType].
	LoraAlpha int64 `json:"lora_alpha"`
	// This field is from variant [FinetuneResponseTrainingTypeLoRaTrainingType].
	LoraR int64 `json:"lora_r"`
	// This field is from variant [FinetuneResponseTrainingTypeLoRaTrainingType].
	LoraDropout float64 `json:"lora_dropout"`
	// This field is from variant [FinetuneResponseTrainingTypeLoRaTrainingType].
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

func (u FinetuneResponseTrainingTypeUnion) AsFinetuneResponseTrainingTypeFullTrainingType() (v FinetuneResponseTrainingTypeFullTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FinetuneResponseTrainingTypeUnion) AsFinetuneResponseTrainingTypeLoRaTrainingType() (v FinetuneResponseTrainingTypeLoRaTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FinetuneResponseTrainingTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *FinetuneResponseTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneResponseTrainingTypeFullTrainingType struct {
	// Any of "Full".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FinetuneResponseTrainingTypeFullTrainingType) RawJSON() string { return r.JSON.raw }
func (r *FinetuneResponseTrainingTypeFullTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FinetuneResponseTrainingTypeLoRaTrainingType struct {
	LoraAlpha int64 `json:"lora_alpha,required"`
	LoraR     int64 `json:"lora_r,required"`
	// Any of "Lora".
	Type                 string  `json:"type,required"`
	LoraDropout          float64 `json:"lora_dropout"`
	LoraTrainableModules string  `json:"lora_trainable_modules"`
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
func (r FinetuneResponseTrainingTypeLoRaTrainingType) RawJSON() string { return r.JSON.raw }
func (r *FinetuneResponseTrainingTypeLoRaTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	Events []FinetuneEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
	// Hugging Face Hub repo to start training from
	FromHfModel string `json:"from_hf_model"`
	// The revision of the Hugging Face Hub model to continue training from
	HfModelRevision string `json:"hf_model_revision"`
	// Learning rate used for training
	LearningRate float64 `json:"learning_rate"`
	// Learning rate scheduler configuration
	LrScheduler FineTuningNewResponseLrScheduler `json:"lr_scheduler"`
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

// Learning rate scheduler configuration
type FineTuningNewResponseLrScheduler struct {
	// Any of "linear", "cosine".
	LrSchedulerType string                                               `json:"lr_scheduler_type,required"`
	LrSchedulerArgs FineTuningNewResponseLrSchedulerLrSchedulerArgsUnion `json:"lr_scheduler_args"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LrSchedulerType respjson.Field
		LrSchedulerArgs respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningNewResponseLrScheduler) RawJSON() string { return r.JSON.raw }
func (r *FineTuningNewResponseLrScheduler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningNewResponseLrSchedulerLrSchedulerArgsUnion contains all possible
// properties and values from
// [FineTuningNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs],
// [FineTuningNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningNewResponseLrSchedulerLrSchedulerArgsUnion struct {
	MinLrRatio float64 `json:"min_lr_ratio"`
	// This field is from variant
	// [FineTuningNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
	NumCycles float64 `json:"num_cycles"`
	JSON      struct {
		MinLrRatio respjson.Field
		NumCycles  respjson.Field
		raw        string
	} `json:"-"`
}

func (u FineTuningNewResponseLrSchedulerLrSchedulerArgsUnion) AsFineTuningNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs() (v FineTuningNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningNewResponseLrSchedulerLrSchedulerArgsUnion) AsFineTuningNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs() (v FineTuningNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningNewResponseLrSchedulerLrSchedulerArgsUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningNewResponseLrSchedulerLrSchedulerArgsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
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
func (r FineTuningNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) RawJSON() string {
	return r.JSON.raw
}
func (r *FineTuningNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
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
func (r FineTuningNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) RawJSON() string {
	return r.JSON.raw
}
func (r *FineTuningNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningNewResponseTrainingMethodUnion contains all possible properties and
// values from [FineTuningNewResponseTrainingMethodTrainingMethodSft],
// [FineTuningNewResponseTrainingMethodTrainingMethodDpo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningNewResponseTrainingMethodUnion struct {
	Method string `json:"method"`
	// This field is from variant
	// [FineTuningNewResponseTrainingMethodTrainingMethodSft].
	TrainOnInputs FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs"`
	// This field is from variant
	// [FineTuningNewResponseTrainingMethodTrainingMethodDpo].
	DpoBeta float64 `json:"dpo_beta"`
	// This field is from variant
	// [FineTuningNewResponseTrainingMethodTrainingMethodDpo].
	DpoNormalizeLogratiosByLength bool `json:"dpo_normalize_logratios_by_length"`
	// This field is from variant
	// [FineTuningNewResponseTrainingMethodTrainingMethodDpo].
	DpoReferenceFree bool `json:"dpo_reference_free"`
	// This field is from variant
	// [FineTuningNewResponseTrainingMethodTrainingMethodDpo].
	RpoAlpha float64 `json:"rpo_alpha"`
	// This field is from variant
	// [FineTuningNewResponseTrainingMethodTrainingMethodDpo].
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

func (u FineTuningNewResponseTrainingMethodUnion) AsFineTuningNewResponseTrainingMethodTrainingMethodSft() (v FineTuningNewResponseTrainingMethodTrainingMethodSft) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningNewResponseTrainingMethodUnion) AsFineTuningNewResponseTrainingMethodTrainingMethodDpo() (v FineTuningNewResponseTrainingMethodTrainingMethodDpo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningNewResponseTrainingMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningNewResponseTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningNewResponseTrainingMethodTrainingMethodSft struct {
	// Any of "sft".
	Method string `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method        respjson.Field
		TrainOnInputs respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningNewResponseTrainingMethodTrainingMethodSft) RawJSON() string { return r.JSON.raw }
func (r *FineTuningNewResponseTrainingMethodTrainingMethodSft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion contains
// all possible properties and values from [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool
// OfFineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString]
type FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfFineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString string `json:",inline"`
	JSON                                                                      struct {
		OfBool                                                                    respjson.Field
		OfFineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString respjson.Field
		raw                                                                       string
	} `json:"-"`
}

func (u FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) AsFineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FineTuningNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

type FineTuningNewResponseTrainingMethodTrainingMethodDpo struct {
	// Any of "dpo".
	Method                        string  `json:"method,required"`
	DpoBeta                       float64 `json:"dpo_beta"`
	DpoNormalizeLogratiosByLength bool    `json:"dpo_normalize_logratios_by_length"`
	DpoReferenceFree              bool    `json:"dpo_reference_free"`
	RpoAlpha                      float64 `json:"rpo_alpha"`
	SimpoGamma                    float64 `json:"simpo_gamma"`
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
func (r FineTuningNewResponseTrainingMethodTrainingMethodDpo) RawJSON() string { return r.JSON.raw }
func (r *FineTuningNewResponseTrainingMethodTrainingMethodDpo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningNewResponseTrainingTypeUnion contains all possible properties and
// values from [FineTuningNewResponseTrainingTypeFullTrainingType],
// [FineTuningNewResponseTrainingTypeLoRaTrainingType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningNewResponseTrainingTypeUnion struct {
	Type string `json:"type"`
	// This field is from variant [FineTuningNewResponseTrainingTypeLoRaTrainingType].
	LoraAlpha int64 `json:"lora_alpha"`
	// This field is from variant [FineTuningNewResponseTrainingTypeLoRaTrainingType].
	LoraR int64 `json:"lora_r"`
	// This field is from variant [FineTuningNewResponseTrainingTypeLoRaTrainingType].
	LoraDropout float64 `json:"lora_dropout"`
	// This field is from variant [FineTuningNewResponseTrainingTypeLoRaTrainingType].
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

func (u FineTuningNewResponseTrainingTypeUnion) AsFineTuningNewResponseTrainingTypeFullTrainingType() (v FineTuningNewResponseTrainingTypeFullTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningNewResponseTrainingTypeUnion) AsFineTuningNewResponseTrainingTypeLoRaTrainingType() (v FineTuningNewResponseTrainingTypeLoRaTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningNewResponseTrainingTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningNewResponseTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningNewResponseTrainingTypeFullTrainingType struct {
	// Any of "Full".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningNewResponseTrainingTypeFullTrainingType) RawJSON() string { return r.JSON.raw }
func (r *FineTuningNewResponseTrainingTypeFullTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningNewResponseTrainingTypeLoRaTrainingType struct {
	LoraAlpha int64 `json:"lora_alpha,required"`
	LoraR     int64 `json:"lora_r,required"`
	// Any of "Lora".
	Type                 string  `json:"type,required"`
	LoraDropout          float64 `json:"lora_dropout"`
	LoraTrainableModules string  `json:"lora_trainable_modules"`
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
func (r FineTuningNewResponseTrainingTypeLoRaTrainingType) RawJSON() string { return r.JSON.raw }
func (r *FineTuningNewResponseTrainingTypeLoRaTrainingType) UnmarshalJSON(data []byte) error {
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
	Events []FinetuneEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
	// Hugging Face Hub repo to start training from
	FromHfModel string `json:"from_hf_model"`
	// The revision of the Hugging Face Hub model to continue training from
	HfModelRevision string `json:"hf_model_revision"`
	// Learning rate used for training
	LearningRate float64 `json:"learning_rate"`
	// Learning rate scheduler configuration
	LrScheduler FineTuningListResponseDataLrScheduler `json:"lr_scheduler"`
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

// Learning rate scheduler configuration
type FineTuningListResponseDataLrScheduler struct {
	// Any of "linear", "cosine".
	LrSchedulerType string                                                    `json:"lr_scheduler_type,required"`
	LrSchedulerArgs FineTuningListResponseDataLrSchedulerLrSchedulerArgsUnion `json:"lr_scheduler_args"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LrSchedulerType respjson.Field
		LrSchedulerArgs respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningListResponseDataLrScheduler) RawJSON() string { return r.JSON.raw }
func (r *FineTuningListResponseDataLrScheduler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningListResponseDataLrSchedulerLrSchedulerArgsUnion contains all possible
// properties and values from
// [FineTuningListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs],
// [FineTuningListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningListResponseDataLrSchedulerLrSchedulerArgsUnion struct {
	MinLrRatio float64 `json:"min_lr_ratio"`
	// This field is from variant
	// [FineTuningListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
	NumCycles float64 `json:"num_cycles"`
	JSON      struct {
		MinLrRatio respjson.Field
		NumCycles  respjson.Field
		raw        string
	} `json:"-"`
}

func (u FineTuningListResponseDataLrSchedulerLrSchedulerArgsUnion) AsFineTuningListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs() (v FineTuningListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningListResponseDataLrSchedulerLrSchedulerArgsUnion) AsFineTuningListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs() (v FineTuningListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningListResponseDataLrSchedulerLrSchedulerArgsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FineTuningListResponseDataLrSchedulerLrSchedulerArgsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
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
func (r FineTuningListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) RawJSON() string {
	return r.JSON.raw
}
func (r *FineTuningListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
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
func (r FineTuningListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) RawJSON() string {
	return r.JSON.raw
}
func (r *FineTuningListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningListResponseDataTrainingMethodUnion contains all possible properties
// and values from [FineTuningListResponseDataTrainingMethodTrainingMethodSft],
// [FineTuningListResponseDataTrainingMethodTrainingMethodDpo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningListResponseDataTrainingMethodUnion struct {
	Method string `json:"method"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingMethodTrainingMethodSft].
	TrainOnInputs FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingMethodTrainingMethodDpo].
	DpoBeta float64 `json:"dpo_beta"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingMethodTrainingMethodDpo].
	DpoNormalizeLogratiosByLength bool `json:"dpo_normalize_logratios_by_length"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingMethodTrainingMethodDpo].
	DpoReferenceFree bool `json:"dpo_reference_free"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingMethodTrainingMethodDpo].
	RpoAlpha float64 `json:"rpo_alpha"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingMethodTrainingMethodDpo].
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

func (u FineTuningListResponseDataTrainingMethodUnion) AsFineTuningListResponseDataTrainingMethodTrainingMethodSft() (v FineTuningListResponseDataTrainingMethodTrainingMethodSft) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningListResponseDataTrainingMethodUnion) AsFineTuningListResponseDataTrainingMethodTrainingMethodDpo() (v FineTuningListResponseDataTrainingMethodTrainingMethodDpo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningListResponseDataTrainingMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningListResponseDataTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListResponseDataTrainingMethodTrainingMethodSft struct {
	// Any of "sft".
	Method string `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method        respjson.Field
		TrainOnInputs respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningListResponseDataTrainingMethodTrainingMethodSft) RawJSON() string {
	return r.JSON.raw
}
func (r *FineTuningListResponseDataTrainingMethodTrainingMethodSft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion
// contains all possible properties and values from [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool
// OfFineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString]
type FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfFineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString string `json:",inline"`
	JSON                                                                           struct {
		OfBool                                                                         respjson.Field
		OfFineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion) AsFineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FineTuningListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

type FineTuningListResponseDataTrainingMethodTrainingMethodDpo struct {
	// Any of "dpo".
	Method                        string  `json:"method,required"`
	DpoBeta                       float64 `json:"dpo_beta"`
	DpoNormalizeLogratiosByLength bool    `json:"dpo_normalize_logratios_by_length"`
	DpoReferenceFree              bool    `json:"dpo_reference_free"`
	RpoAlpha                      float64 `json:"rpo_alpha"`
	SimpoGamma                    float64 `json:"simpo_gamma"`
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
func (r FineTuningListResponseDataTrainingMethodTrainingMethodDpo) RawJSON() string {
	return r.JSON.raw
}
func (r *FineTuningListResponseDataTrainingMethodTrainingMethodDpo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningListResponseDataTrainingTypeUnion contains all possible properties and
// values from [FineTuningListResponseDataTrainingTypeFullTrainingType],
// [FineTuningListResponseDataTrainingTypeLoRaTrainingType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningListResponseDataTrainingTypeUnion struct {
	Type string `json:"type"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingTypeLoRaTrainingType].
	LoraAlpha int64 `json:"lora_alpha"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingTypeLoRaTrainingType].
	LoraR int64 `json:"lora_r"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingTypeLoRaTrainingType].
	LoraDropout float64 `json:"lora_dropout"`
	// This field is from variant
	// [FineTuningListResponseDataTrainingTypeLoRaTrainingType].
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

func (u FineTuningListResponseDataTrainingTypeUnion) AsFineTuningListResponseDataTrainingTypeFullTrainingType() (v FineTuningListResponseDataTrainingTypeFullTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningListResponseDataTrainingTypeUnion) AsFineTuningListResponseDataTrainingTypeLoRaTrainingType() (v FineTuningListResponseDataTrainingTypeLoRaTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningListResponseDataTrainingTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningListResponseDataTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListResponseDataTrainingTypeFullTrainingType struct {
	// Any of "Full".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningListResponseDataTrainingTypeFullTrainingType) RawJSON() string { return r.JSON.raw }
func (r *FineTuningListResponseDataTrainingTypeFullTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListResponseDataTrainingTypeLoRaTrainingType struct {
	LoraAlpha int64 `json:"lora_alpha,required"`
	LoraR     int64 `json:"lora_r,required"`
	// Any of "Lora".
	Type                 string  `json:"type,required"`
	LoraDropout          float64 `json:"lora_dropout"`
	LoraTrainableModules string  `json:"lora_trainable_modules"`
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
func (r FineTuningListResponseDataTrainingTypeLoRaTrainingType) RawJSON() string { return r.JSON.raw }
func (r *FineTuningListResponseDataTrainingTypeLoRaTrainingType) UnmarshalJSON(data []byte) error {
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
	Events []FinetuneEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
	// Hugging Face Hub repo to start training from
	FromHfModel string `json:"from_hf_model"`
	// The revision of the Hugging Face Hub model to continue training from
	HfModelRevision string `json:"hf_model_revision"`
	// Learning rate used for training
	LearningRate float64 `json:"learning_rate"`
	// Learning rate scheduler configuration
	LrScheduler FineTuningCancelResponseLrScheduler `json:"lr_scheduler"`
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

// Learning rate scheduler configuration
type FineTuningCancelResponseLrScheduler struct {
	// Any of "linear", "cosine".
	LrSchedulerType string                                                  `json:"lr_scheduler_type,required"`
	LrSchedulerArgs FineTuningCancelResponseLrSchedulerLrSchedulerArgsUnion `json:"lr_scheduler_args"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LrSchedulerType respjson.Field
		LrSchedulerArgs respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningCancelResponseLrScheduler) RawJSON() string { return r.JSON.raw }
func (r *FineTuningCancelResponseLrScheduler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningCancelResponseLrSchedulerLrSchedulerArgsUnion contains all possible
// properties and values from
// [FineTuningCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs],
// [FineTuningCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningCancelResponseLrSchedulerLrSchedulerArgsUnion struct {
	MinLrRatio float64 `json:"min_lr_ratio"`
	// This field is from variant
	// [FineTuningCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
	NumCycles float64 `json:"num_cycles"`
	JSON      struct {
		MinLrRatio respjson.Field
		NumCycles  respjson.Field
		raw        string
	} `json:"-"`
}

func (u FineTuningCancelResponseLrSchedulerLrSchedulerArgsUnion) AsFineTuningCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs() (v FineTuningCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningCancelResponseLrSchedulerLrSchedulerArgsUnion) AsFineTuningCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs() (v FineTuningCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningCancelResponseLrSchedulerLrSchedulerArgsUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningCancelResponseLrSchedulerLrSchedulerArgsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
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
func (r FineTuningCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) RawJSON() string {
	return r.JSON.raw
}
func (r *FineTuningCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
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
func (r FineTuningCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) RawJSON() string {
	return r.JSON.raw
}
func (r *FineTuningCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningCancelResponseTrainingMethodUnion contains all possible properties and
// values from [FineTuningCancelResponseTrainingMethodTrainingMethodSft],
// [FineTuningCancelResponseTrainingMethodTrainingMethodDpo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningCancelResponseTrainingMethodUnion struct {
	Method string `json:"method"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingMethodTrainingMethodSft].
	TrainOnInputs FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingMethodTrainingMethodDpo].
	DpoBeta float64 `json:"dpo_beta"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingMethodTrainingMethodDpo].
	DpoNormalizeLogratiosByLength bool `json:"dpo_normalize_logratios_by_length"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingMethodTrainingMethodDpo].
	DpoReferenceFree bool `json:"dpo_reference_free"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingMethodTrainingMethodDpo].
	RpoAlpha float64 `json:"rpo_alpha"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingMethodTrainingMethodDpo].
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

func (u FineTuningCancelResponseTrainingMethodUnion) AsFineTuningCancelResponseTrainingMethodTrainingMethodSft() (v FineTuningCancelResponseTrainingMethodTrainingMethodSft) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningCancelResponseTrainingMethodUnion) AsFineTuningCancelResponseTrainingMethodTrainingMethodDpo() (v FineTuningCancelResponseTrainingMethodTrainingMethodDpo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningCancelResponseTrainingMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningCancelResponseTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningCancelResponseTrainingMethodTrainingMethodSft struct {
	// Any of "sft".
	Method string `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method        respjson.Field
		TrainOnInputs respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningCancelResponseTrainingMethodTrainingMethodSft) RawJSON() string { return r.JSON.raw }
func (r *FineTuningCancelResponseTrainingMethodTrainingMethodSft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion
// contains all possible properties and values from [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool
// OfFineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString]
type FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfFineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString string `json:",inline"`
	JSON                                                                         struct {
		OfBool                                                                       respjson.Field
		OfFineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) AsFineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FineTuningCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

type FineTuningCancelResponseTrainingMethodTrainingMethodDpo struct {
	// Any of "dpo".
	Method                        string  `json:"method,required"`
	DpoBeta                       float64 `json:"dpo_beta"`
	DpoNormalizeLogratiosByLength bool    `json:"dpo_normalize_logratios_by_length"`
	DpoReferenceFree              bool    `json:"dpo_reference_free"`
	RpoAlpha                      float64 `json:"rpo_alpha"`
	SimpoGamma                    float64 `json:"simpo_gamma"`
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
func (r FineTuningCancelResponseTrainingMethodTrainingMethodDpo) RawJSON() string { return r.JSON.raw }
func (r *FineTuningCancelResponseTrainingMethodTrainingMethodDpo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FineTuningCancelResponseTrainingTypeUnion contains all possible properties and
// values from [FineTuningCancelResponseTrainingTypeFullTrainingType],
// [FineTuningCancelResponseTrainingTypeLoRaTrainingType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FineTuningCancelResponseTrainingTypeUnion struct {
	Type string `json:"type"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingTypeLoRaTrainingType].
	LoraAlpha int64 `json:"lora_alpha"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingTypeLoRaTrainingType].
	LoraR int64 `json:"lora_r"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingTypeLoRaTrainingType].
	LoraDropout float64 `json:"lora_dropout"`
	// This field is from variant
	// [FineTuningCancelResponseTrainingTypeLoRaTrainingType].
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

func (u FineTuningCancelResponseTrainingTypeUnion) AsFineTuningCancelResponseTrainingTypeFullTrainingType() (v FineTuningCancelResponseTrainingTypeFullTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FineTuningCancelResponseTrainingTypeUnion) AsFineTuningCancelResponseTrainingTypeLoRaTrainingType() (v FineTuningCancelResponseTrainingTypeLoRaTrainingType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FineTuningCancelResponseTrainingTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *FineTuningCancelResponseTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningCancelResponseTrainingTypeFullTrainingType struct {
	// Any of "Full".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningCancelResponseTrainingTypeFullTrainingType) RawJSON() string { return r.JSON.raw }
func (r *FineTuningCancelResponseTrainingTypeFullTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningCancelResponseTrainingTypeLoRaTrainingType struct {
	LoraAlpha int64 `json:"lora_alpha,required"`
	LoraR     int64 `json:"lora_r,required"`
	// Any of "Lora".
	Type                 string  `json:"type,required"`
	LoraDropout          float64 `json:"lora_dropout"`
	LoraTrainableModules string  `json:"lora_trainable_modules"`
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
func (r FineTuningCancelResponseTrainingTypeLoRaTrainingType) RawJSON() string { return r.JSON.raw }
func (r *FineTuningCancelResponseTrainingTypeLoRaTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListCheckpointsResponse struct {
	Data []FineTuningListCheckpointsResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningListCheckpointsResponse) RawJSON() string { return r.JSON.raw }
func (r *FineTuningListCheckpointsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListCheckpointsResponseData struct {
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
func (r FineTuningListCheckpointsResponseData) RawJSON() string { return r.JSON.raw }
func (r *FineTuningListCheckpointsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FineTuningListEventsResponse struct {
	Data []FinetuneEvent `json:"data,required"`
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
	LrScheduler FineTuningNewParamsLrScheduler `json:"lr_scheduler,omitzero"`
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

// The learning rate scheduler to use. It specifies how the learning rate is
// adjusted during training.
//
// The property LrSchedulerType is required.
type FineTuningNewParamsLrScheduler struct {
	// Any of "linear", "cosine".
	LrSchedulerType string                                             `json:"lr_scheduler_type,omitzero,required"`
	LrSchedulerArgs FineTuningNewParamsLrSchedulerLrSchedulerArgsUnion `json:"lr_scheduler_args,omitzero"`
	paramObj
}

func (r FineTuningNewParamsLrScheduler) MarshalJSON() (data []byte, err error) {
	type shadow FineTuningNewParamsLrScheduler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FineTuningNewParamsLrScheduler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FineTuningNewParamsLrScheduler](
		"lr_scheduler_type", "linear", "cosine",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FineTuningNewParamsLrSchedulerLrSchedulerArgsUnion struct {
	OfFineTuningNewsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs *FineTuningNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs `json:",omitzero,inline"`
	OfFineTuningNewsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs *FineTuningNewParamsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs `json:",omitzero,inline"`
	paramUnion
}

func (u FineTuningNewParamsLrSchedulerLrSchedulerArgsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFineTuningNewsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs, u.OfFineTuningNewsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs)
}
func (u *FineTuningNewParamsLrSchedulerLrSchedulerArgsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FineTuningNewParamsLrSchedulerLrSchedulerArgsUnion) asAny() any {
	if !param.IsOmitted(u.OfFineTuningNewsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) {
		return u.OfFineTuningNewsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs
	} else if !param.IsOmitted(u.OfFineTuningNewsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) {
		return u.OfFineTuningNewsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsLrSchedulerLrSchedulerArgsUnion) GetNumCycles() *float64 {
	if vt := u.OfFineTuningNewsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs; vt != nil {
		return &vt.NumCycles
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsLrSchedulerLrSchedulerArgsUnion) GetMinLrRatio() *float64 {
	if vt := u.OfFineTuningNewsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs; vt != nil && vt.MinLrRatio.Valid() {
		return &vt.MinLrRatio.Value
	} else if vt := u.OfFineTuningNewsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs; vt != nil {
		return (*float64)(&vt.MinLrRatio)
	}
	return nil
}

type FineTuningNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio param.Opt[float64] `json:"min_lr_ratio,omitzero"`
	paramObj
}

func (r FineTuningNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) MarshalJSON() (data []byte, err error) {
	type shadow FineTuningNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FineTuningNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MinLrRatio, NumCycles are required.
type FineTuningNewParamsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio,required"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64 `json:"num_cycles,required"`
	paramObj
}

func (r FineTuningNewParamsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) MarshalJSON() (data []byte, err error) {
	type shadow FineTuningNewParamsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FineTuningNewParamsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	OfFineTuningNewsTrainingMethodTrainingMethodSft *FineTuningNewParamsTrainingMethodTrainingMethodSft `json:",omitzero,inline"`
	OfFineTuningNewsTrainingMethodTrainingMethodDpo *FineTuningNewParamsTrainingMethodTrainingMethodDpo `json:",omitzero,inline"`
	paramUnion
}

func (u FineTuningNewParamsTrainingMethodUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFineTuningNewsTrainingMethodTrainingMethodSft, u.OfFineTuningNewsTrainingMethodTrainingMethodDpo)
}
func (u *FineTuningNewParamsTrainingMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FineTuningNewParamsTrainingMethodUnion) asAny() any {
	if !param.IsOmitted(u.OfFineTuningNewsTrainingMethodTrainingMethodSft) {
		return u.OfFineTuningNewsTrainingMethodTrainingMethodSft
	} else if !param.IsOmitted(u.OfFineTuningNewsTrainingMethodTrainingMethodDpo) {
		return u.OfFineTuningNewsTrainingMethodTrainingMethodDpo
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetTrainOnInputs() *FineTuningNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion {
	if vt := u.OfFineTuningNewsTrainingMethodTrainingMethodSft; vt != nil {
		return &vt.TrainOnInputs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetDpoBeta() *float64 {
	if vt := u.OfFineTuningNewsTrainingMethodTrainingMethodDpo; vt != nil && vt.DpoBeta.Valid() {
		return &vt.DpoBeta.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetDpoNormalizeLogratiosByLength() *bool {
	if vt := u.OfFineTuningNewsTrainingMethodTrainingMethodDpo; vt != nil && vt.DpoNormalizeLogratiosByLength.Valid() {
		return &vt.DpoNormalizeLogratiosByLength.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetDpoReferenceFree() *bool {
	if vt := u.OfFineTuningNewsTrainingMethodTrainingMethodDpo; vt != nil && vt.DpoReferenceFree.Valid() {
		return &vt.DpoReferenceFree.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetRpoAlpha() *float64 {
	if vt := u.OfFineTuningNewsTrainingMethodTrainingMethodDpo; vt != nil && vt.RpoAlpha.Valid() {
		return &vt.RpoAlpha.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetSimpoGamma() *float64 {
	if vt := u.OfFineTuningNewsTrainingMethodTrainingMethodDpo; vt != nil && vt.SimpoGamma.Valid() {
		return &vt.SimpoGamma.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingMethodUnion) GetMethod() *string {
	if vt := u.OfFineTuningNewsTrainingMethodTrainingMethodSft; vt != nil {
		return (*string)(&vt.Method)
	} else if vt := u.OfFineTuningNewsTrainingMethodTrainingMethodDpo; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// The properties Method, TrainOnInputs are required.
type FineTuningNewParamsTrainingMethodTrainingMethodSft struct {
	// Any of "sft".
	Method string `json:"method,omitzero,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FineTuningNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,omitzero,required"`
	paramObj
}

func (r FineTuningNewParamsTrainingMethodTrainingMethodSft) MarshalJSON() (data []byte, err error) {
	type shadow FineTuningNewParamsTrainingMethodTrainingMethodSft
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FineTuningNewParamsTrainingMethodTrainingMethodSft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FineTuningNewParamsTrainingMethodTrainingMethodSft](
		"method", "sft",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FineTuningNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFineTuningNewsTrainingMethodTrainingMethodSftTrainOnInputsString)
	OfFineTuningNewsTrainingMethodTrainingMethodSftTrainOnInputsString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u FineTuningNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFineTuningNewsTrainingMethodTrainingMethodSftTrainOnInputsString)
}
func (u *FineTuningNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FineTuningNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFineTuningNewsTrainingMethodTrainingMethodSftTrainOnInputsString) {
		return &u.OfFineTuningNewsTrainingMethodTrainingMethodSftTrainOnInputsString
	}
	return nil
}

type FineTuningNewParamsTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FineTuningNewParamsTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FineTuningNewParamsTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

// The property Method is required.
type FineTuningNewParamsTrainingMethodTrainingMethodDpo struct {
	// Any of "dpo".
	Method                        string             `json:"method,omitzero,required"`
	DpoBeta                       param.Opt[float64] `json:"dpo_beta,omitzero"`
	DpoNormalizeLogratiosByLength param.Opt[bool]    `json:"dpo_normalize_logratios_by_length,omitzero"`
	DpoReferenceFree              param.Opt[bool]    `json:"dpo_reference_free,omitzero"`
	RpoAlpha                      param.Opt[float64] `json:"rpo_alpha,omitzero"`
	SimpoGamma                    param.Opt[float64] `json:"simpo_gamma,omitzero"`
	paramObj
}

func (r FineTuningNewParamsTrainingMethodTrainingMethodDpo) MarshalJSON() (data []byte, err error) {
	type shadow FineTuningNewParamsTrainingMethodTrainingMethodDpo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FineTuningNewParamsTrainingMethodTrainingMethodDpo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FineTuningNewParamsTrainingMethodTrainingMethodDpo](
		"method", "dpo",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FineTuningNewParamsTrainingTypeUnion struct {
	OfFineTuningNewsTrainingTypeFullTrainingType *FineTuningNewParamsTrainingTypeFullTrainingType `json:",omitzero,inline"`
	OfFineTuningNewsTrainingTypeLoRaTrainingType *FineTuningNewParamsTrainingTypeLoRaTrainingType `json:",omitzero,inline"`
	paramUnion
}

func (u FineTuningNewParamsTrainingTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFineTuningNewsTrainingTypeFullTrainingType, u.OfFineTuningNewsTrainingTypeLoRaTrainingType)
}
func (u *FineTuningNewParamsTrainingTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FineTuningNewParamsTrainingTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfFineTuningNewsTrainingTypeFullTrainingType) {
		return u.OfFineTuningNewsTrainingTypeFullTrainingType
	} else if !param.IsOmitted(u.OfFineTuningNewsTrainingTypeLoRaTrainingType) {
		return u.OfFineTuningNewsTrainingTypeLoRaTrainingType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetLoraAlpha() *int64 {
	if vt := u.OfFineTuningNewsTrainingTypeLoRaTrainingType; vt != nil {
		return &vt.LoraAlpha
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetLoraR() *int64 {
	if vt := u.OfFineTuningNewsTrainingTypeLoRaTrainingType; vt != nil {
		return &vt.LoraR
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetLoraDropout() *float64 {
	if vt := u.OfFineTuningNewsTrainingTypeLoRaTrainingType; vt != nil && vt.LoraDropout.Valid() {
		return &vt.LoraDropout.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetLoraTrainableModules() *string {
	if vt := u.OfFineTuningNewsTrainingTypeLoRaTrainingType; vt != nil && vt.LoraTrainableModules.Valid() {
		return &vt.LoraTrainableModules.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FineTuningNewParamsTrainingTypeUnion) GetType() *string {
	if vt := u.OfFineTuningNewsTrainingTypeFullTrainingType; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFineTuningNewsTrainingTypeLoRaTrainingType; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The property Type is required.
type FineTuningNewParamsTrainingTypeFullTrainingType struct {
	// Any of "Full".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r FineTuningNewParamsTrainingTypeFullTrainingType) MarshalJSON() (data []byte, err error) {
	type shadow FineTuningNewParamsTrainingTypeFullTrainingType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FineTuningNewParamsTrainingTypeFullTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FineTuningNewParamsTrainingTypeFullTrainingType](
		"type", "Full",
	)
}

// The properties LoraAlpha, LoraR, Type are required.
type FineTuningNewParamsTrainingTypeLoRaTrainingType struct {
	LoraAlpha int64 `json:"lora_alpha,required"`
	LoraR     int64 `json:"lora_r,required"`
	// Any of "Lora".
	Type                 string             `json:"type,omitzero,required"`
	LoraDropout          param.Opt[float64] `json:"lora_dropout,omitzero"`
	LoraTrainableModules param.Opt[string]  `json:"lora_trainable_modules,omitzero"`
	paramObj
}

func (r FineTuningNewParamsTrainingTypeLoRaTrainingType) MarshalJSON() (data []byte, err error) {
	type shadow FineTuningNewParamsTrainingTypeLoRaTrainingType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FineTuningNewParamsTrainingTypeLoRaTrainingType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FineTuningNewParamsTrainingTypeLoRaTrainingType](
		"type", "Lora",
	)
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

type FineTuningContentParams struct {
	// Fine-tune ID to download. A string that starts with `ft-`.
	FtID string `query:"ft_id,required" json:"-"`
	// Specifies step number for checkpoint to download. Ignores `checkpoint` value if
	// set.
	CheckpointStep param.Opt[int64] `query:"checkpoint_step,omitzero" json:"-"`
	// Specifies checkpoint type to download - `merged` vs `adapter`. This field is
	// required if the checkpoint_step is not set.
	//
	// Any of "merged", "adapter", "model_output_path".
	Checkpoint FineTuningContentParamsCheckpoint `query:"checkpoint,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FineTuningContentParams]'s query parameters as
// `url.Values`.
func (r FineTuningContentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies checkpoint type to download - `merged` vs `adapter`. This field is
// required if the checkpoint_step is not set.
type FineTuningContentParamsCheckpoint string

const (
	FineTuningContentParamsCheckpointMerged          FineTuningContentParamsCheckpoint = "merged"
	FineTuningContentParamsCheckpointAdapter         FineTuningContentParamsCheckpoint = "adapter"
	FineTuningContentParamsCheckpointModelOutputPath FineTuningContentParamsCheckpoint = "model_output_path"
)
