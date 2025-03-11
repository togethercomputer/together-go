// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/tidwall/gjson"
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/shared"
)

// FineTuneService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFineTuneService] method instead.
type FineTuneService struct {
	Options []option.RequestOption
}

// NewFineTuneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFineTuneService(opts ...option.RequestOption) (r *FineTuneService) {
	r = &FineTuneService{}
	r.Options = opts
	return
}

// Use a model to create a fine-tuning job.
func (r *FineTuneService) New(ctx context.Context, body FineTuneNewParams, opts ...option.RequestOption) (res *FineTune, err error) {
	opts = append(r.Options[:], opts...)
	path := "fine-tunes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List the metadata for a single fine-tuning job.
func (r *FineTuneService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTune, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the metadata for all fine-tuning jobs.
func (r *FineTuneService) List(ctx context.Context, opts ...option.RequestOption) (res *FineTuneListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "fine-tunes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a currently running fine-tuning job.
func (r *FineTuneService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTune, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Download a compressed fine-tuned model or checkpoint to local disk.
func (r *FineTuneService) Download(ctx context.Context, query FineTuneDownloadParams, opts ...option.RequestOption) (res *FineTuneDownloadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "finetune/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List the events for a single fine-tuning job.
func (r *FineTuneService) ListEvents(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTuneEvent, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FineTune struct {
	ID                   string                     `json:"id,required" format:"uuid"`
	Status               FineTuneStatus             `json:"status,required"`
	BatchSize            int64                      `json:"batch_size"`
	CreatedAt            string                     `json:"created_at"`
	EpochsCompleted      int64                      `json:"epochs_completed"`
	EvalSteps            int64                      `json:"eval_steps"`
	Events               []FineTuneEvent            `json:"events"`
	FromCheckpoint       string                     `json:"from_checkpoint"`
	JobID                string                     `json:"job_id"`
	LearningRate         float64                    `json:"learning_rate"`
	LrScheduler          FineTuneLrScheduler        `json:"lr_scheduler"`
	MaxGradNorm          float64                    `json:"max_grad_norm"`
	Model                string                     `json:"model"`
	ModelOutputName      string                     `json:"model_output_name"`
	ModelOutputPath      string                     `json:"model_output_path"`
	NCheckpoints         int64                      `json:"n_checkpoints"`
	NEpochs              int64                      `json:"n_epochs"`
	NEvals               int64                      `json:"n_evals"`
	ParamCount           int64                      `json:"param_count"`
	QueueDepth           int64                      `json:"queue_depth"`
	TokenCount           int64                      `json:"token_count"`
	TotalPrice           int64                      `json:"total_price"`
	TrainOnInputs        FineTuneTrainOnInputsUnion `json:"train_on_inputs"`
	TrainingFile         string                     `json:"training_file"`
	TrainingType         FineTuneTrainingType       `json:"training_type"`
	TrainingfileNumlines int64                      `json:"trainingfile_numlines"`
	TrainingfileSize     int64                      `json:"trainingfile_size"`
	UpdatedAt            string                     `json:"updated_at"`
	ValidationFile       string                     `json:"validation_file"`
	WandbProjectName     string                     `json:"wandb_project_name"`
	WandbURL             string                     `json:"wandb_url"`
	WarmupRatio          float64                    `json:"warmup_ratio"`
	WeightDecay          float64                    `json:"weight_decay"`
	JSON                 fineTuneJSON               `json:"-"`
}

// fineTuneJSON contains the JSON metadata for the struct [FineTune]
type fineTuneJSON struct {
	ID                   apijson.Field
	Status               apijson.Field
	BatchSize            apijson.Field
	CreatedAt            apijson.Field
	EpochsCompleted      apijson.Field
	EvalSteps            apijson.Field
	Events               apijson.Field
	FromCheckpoint       apijson.Field
	JobID                apijson.Field
	LearningRate         apijson.Field
	LrScheduler          apijson.Field
	MaxGradNorm          apijson.Field
	Model                apijson.Field
	ModelOutputName      apijson.Field
	ModelOutputPath      apijson.Field
	NCheckpoints         apijson.Field
	NEpochs              apijson.Field
	NEvals               apijson.Field
	ParamCount           apijson.Field
	QueueDepth           apijson.Field
	TokenCount           apijson.Field
	TotalPrice           apijson.Field
	TrainOnInputs        apijson.Field
	TrainingFile         apijson.Field
	TrainingType         apijson.Field
	TrainingfileNumlines apijson.Field
	TrainingfileSize     apijson.Field
	UpdatedAt            apijson.Field
	ValidationFile       apijson.Field
	WandbProjectName     apijson.Field
	WandbURL             apijson.Field
	WarmupRatio          apijson.Field
	WeightDecay          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FineTune) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneJSON) RawJSON() string {
	return r.raw
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

func (r FineTuneStatus) IsKnown() bool {
	switch r {
	case FineTuneStatusPending, FineTuneStatusQueued, FineTuneStatusRunning, FineTuneStatusCompressing, FineTuneStatusUploading, FineTuneStatusCancelRequested, FineTuneStatusCancelled, FineTuneStatusError, FineTuneStatusCompleted:
		return true
	}
	return false
}

type FineTuneEvent struct {
	CheckpointPath string               `json:"checkpoint_path,required"`
	CreatedAt      string               `json:"created_at,required"`
	Hash           string               `json:"hash,required"`
	Message        string               `json:"message,required"`
	ModelPath      string               `json:"model_path,required"`
	Object         FineTuneEventsObject `json:"object,required"`
	ParamCount     int64                `json:"param_count,required"`
	Step           int64                `json:"step,required"`
	TokenCount     int64                `json:"token_count,required"`
	TotalSteps     int64                `json:"total_steps,required"`
	TrainingOffset int64                `json:"training_offset,required"`
	Type           FineTuneEventsType   `json:"type,required"`
	WandbURL       string               `json:"wandb_url,required"`
	Level          FineTuneEventsLevel  `json:"level,nullable"`
	JSON           fineTuneEventJSON    `json:"-"`
}

// fineTuneEventJSON contains the JSON metadata for the struct [FineTuneEvent]
type fineTuneEventJSON struct {
	CheckpointPath apijson.Field
	CreatedAt      apijson.Field
	Hash           apijson.Field
	Message        apijson.Field
	ModelPath      apijson.Field
	Object         apijson.Field
	ParamCount     apijson.Field
	Step           apijson.Field
	TokenCount     apijson.Field
	TotalSteps     apijson.Field
	TrainingOffset apijson.Field
	Type           apijson.Field
	WandbURL       apijson.Field
	Level          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FineTuneEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneEventJSON) RawJSON() string {
	return r.raw
}

type FineTuneEventsObject string

const (
	FineTuneEventsObjectFineTuneEvent FineTuneEventsObject = "fine-tune-event"
)

func (r FineTuneEventsObject) IsKnown() bool {
	switch r {
	case FineTuneEventsObjectFineTuneEvent:
		return true
	}
	return false
}

type FineTuneEventsType string

const (
	FineTuneEventsTypeJobPending                     FineTuneEventsType = "job_pending"
	FineTuneEventsTypeJobStart                       FineTuneEventsType = "job_start"
	FineTuneEventsTypeJobStopped                     FineTuneEventsType = "job_stopped"
	FineTuneEventsTypeModelDownloading               FineTuneEventsType = "model_downloading"
	FineTuneEventsTypeModelDownloadComplete          FineTuneEventsType = "model_download_complete"
	FineTuneEventsTypeTrainingDataDownloading        FineTuneEventsType = "training_data_downloading"
	FineTuneEventsTypeTrainingDataDownloadComplete   FineTuneEventsType = "training_data_download_complete"
	FineTuneEventsTypeValidationDataDownloading      FineTuneEventsType = "validation_data_downloading"
	FineTuneEventsTypeValidationDataDownloadComplete FineTuneEventsType = "validation_data_download_complete"
	FineTuneEventsTypeWandbInit                      FineTuneEventsType = "wandb_init"
	FineTuneEventsTypeTrainingStart                  FineTuneEventsType = "training_start"
	FineTuneEventsTypeCheckpointSave                 FineTuneEventsType = "checkpoint_save"
	FineTuneEventsTypeBillingLimit                   FineTuneEventsType = "billing_limit"
	FineTuneEventsTypeEpochComplete                  FineTuneEventsType = "epoch_complete"
	FineTuneEventsTypeTrainingComplete               FineTuneEventsType = "training_complete"
	FineTuneEventsTypeModelCompressing               FineTuneEventsType = "model_compressing"
	FineTuneEventsTypeModelCompressionComplete       FineTuneEventsType = "model_compression_complete"
	FineTuneEventsTypeModelUploading                 FineTuneEventsType = "model_uploading"
	FineTuneEventsTypeModelUploadComplete            FineTuneEventsType = "model_upload_complete"
	FineTuneEventsTypeJobComplete                    FineTuneEventsType = "job_complete"
	FineTuneEventsTypeJobError                       FineTuneEventsType = "job_error"
	FineTuneEventsTypeCancelRequested                FineTuneEventsType = "cancel_requested"
	FineTuneEventsTypeJobRestarted                   FineTuneEventsType = "job_restarted"
	FineTuneEventsTypeRefund                         FineTuneEventsType = "refund"
	FineTuneEventsTypeWarning                        FineTuneEventsType = "warning"
)

func (r FineTuneEventsType) IsKnown() bool {
	switch r {
	case FineTuneEventsTypeJobPending, FineTuneEventsTypeJobStart, FineTuneEventsTypeJobStopped, FineTuneEventsTypeModelDownloading, FineTuneEventsTypeModelDownloadComplete, FineTuneEventsTypeTrainingDataDownloading, FineTuneEventsTypeTrainingDataDownloadComplete, FineTuneEventsTypeValidationDataDownloading, FineTuneEventsTypeValidationDataDownloadComplete, FineTuneEventsTypeWandbInit, FineTuneEventsTypeTrainingStart, FineTuneEventsTypeCheckpointSave, FineTuneEventsTypeBillingLimit, FineTuneEventsTypeEpochComplete, FineTuneEventsTypeTrainingComplete, FineTuneEventsTypeModelCompressing, FineTuneEventsTypeModelCompressionComplete, FineTuneEventsTypeModelUploading, FineTuneEventsTypeModelUploadComplete, FineTuneEventsTypeJobComplete, FineTuneEventsTypeJobError, FineTuneEventsTypeCancelRequested, FineTuneEventsTypeJobRestarted, FineTuneEventsTypeRefund, FineTuneEventsTypeWarning:
		return true
	}
	return false
}

type FineTuneEventsLevel string

const (
	FineTuneEventsLevelInfo           FineTuneEventsLevel = "info"
	FineTuneEventsLevelWarning        FineTuneEventsLevel = "warning"
	FineTuneEventsLevelError          FineTuneEventsLevel = "error"
	FineTuneEventsLevelLegacyInfo     FineTuneEventsLevel = "legacy_info"
	FineTuneEventsLevelLegacyIwarning FineTuneEventsLevel = "legacy_iwarning"
	FineTuneEventsLevelLegacyIerror   FineTuneEventsLevel = "legacy_ierror"
)

func (r FineTuneEventsLevel) IsKnown() bool {
	switch r {
	case FineTuneEventsLevelInfo, FineTuneEventsLevelWarning, FineTuneEventsLevelError, FineTuneEventsLevelLegacyInfo, FineTuneEventsLevelLegacyIwarning, FineTuneEventsLevelLegacyIerror:
		return true
	}
	return false
}

type FineTuneLrScheduler struct {
	LrSchedulerType string                             `json:"lr_scheduler_type,required"`
	LrSchedulerArgs FineTuneLrSchedulerLrSchedulerArgs `json:"lr_scheduler_args"`
	JSON            fineTuneLrSchedulerJSON            `json:"-"`
}

// fineTuneLrSchedulerJSON contains the JSON metadata for the struct
// [FineTuneLrScheduler]
type fineTuneLrSchedulerJSON struct {
	LrSchedulerType apijson.Field
	LrSchedulerArgs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FineTuneLrScheduler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneLrSchedulerJSON) RawJSON() string {
	return r.raw
}

type FineTuneLrSchedulerLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64                                `json:"min_lr_ratio"`
	JSON       fineTuneLrSchedulerLrSchedulerArgsJSON `json:"-"`
}

// fineTuneLrSchedulerLrSchedulerArgsJSON contains the JSON metadata for the struct
// [FineTuneLrSchedulerLrSchedulerArgs]
type fineTuneLrSchedulerLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneLrSchedulerLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneLrSchedulerLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool] or [FineTuneTrainOnInputsString].
type FineTuneTrainOnInputsUnion interface {
	ImplementsFineTuneTrainOnInputsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneTrainOnInputsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FineTuneTrainOnInputsString("")),
		},
	)
}

type FineTuneTrainOnInputsString string

const (
	FineTuneTrainOnInputsStringAuto FineTuneTrainOnInputsString = "auto"
)

func (r FineTuneTrainOnInputsString) IsKnown() bool {
	switch r {
	case FineTuneTrainOnInputsStringAuto:
		return true
	}
	return false
}

func (r FineTuneTrainOnInputsString) ImplementsFineTuneTrainOnInputsUnion() {}

type FineTuneTrainingType struct {
	Type                 FineTuneTrainingTypeType `json:"type,required"`
	LoraAlpha            int64                    `json:"lora_alpha"`
	LoraDropout          float64                  `json:"lora_dropout"`
	LoraR                int64                    `json:"lora_r"`
	LoraTrainableModules string                   `json:"lora_trainable_modules"`
	JSON                 fineTuneTrainingTypeJSON `json:"-"`
	union                FineTuneTrainingTypeUnion
}

// fineTuneTrainingTypeJSON contains the JSON metadata for the struct
// [FineTuneTrainingType]
type fineTuneTrainingTypeJSON struct {
	Type                 apijson.Field
	LoraAlpha            apijson.Field
	LoraDropout          apijson.Field
	LoraR                apijson.Field
	LoraTrainableModules apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r fineTuneTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneTrainingType) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneTrainingType{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneTrainingTypeUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [FineTuneTrainingTypeFullTrainingType],
// [FineTuneTrainingTypeLoRaTrainingType].
func (r FineTuneTrainingType) AsUnion() FineTuneTrainingTypeUnion {
	return r.union
}

// Union satisfied by [FineTuneTrainingTypeFullTrainingType] or
// [FineTuneTrainingTypeLoRaTrainingType].
type FineTuneTrainingTypeUnion interface {
	implementsFineTuneTrainingType()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneTrainingTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneTrainingTypeFullTrainingType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneTrainingTypeLoRaTrainingType{}),
		},
	)
}

type FineTuneTrainingTypeFullTrainingType struct {
	Type FineTuneTrainingTypeFullTrainingTypeType `json:"type,required"`
	JSON fineTuneTrainingTypeFullTrainingTypeJSON `json:"-"`
}

// fineTuneTrainingTypeFullTrainingTypeJSON contains the JSON metadata for the
// struct [FineTuneTrainingTypeFullTrainingType]
type fineTuneTrainingTypeFullTrainingTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneTrainingTypeFullTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneTrainingTypeFullTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneTrainingTypeFullTrainingType) implementsFineTuneTrainingType() {}

type FineTuneTrainingTypeFullTrainingTypeType string

const (
	FineTuneTrainingTypeFullTrainingTypeTypeFull FineTuneTrainingTypeFullTrainingTypeType = "Full"
)

func (r FineTuneTrainingTypeFullTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneTrainingTypeFullTrainingTypeTypeFull:
		return true
	}
	return false
}

type FineTuneTrainingTypeLoRaTrainingType struct {
	LoraAlpha            int64                                    `json:"lora_alpha,required"`
	LoraR                int64                                    `json:"lora_r,required"`
	Type                 FineTuneTrainingTypeLoRaTrainingTypeType `json:"type,required"`
	LoraDropout          float64                                  `json:"lora_dropout"`
	LoraTrainableModules string                                   `json:"lora_trainable_modules"`
	JSON                 fineTuneTrainingTypeLoRaTrainingTypeJSON `json:"-"`
}

// fineTuneTrainingTypeLoRaTrainingTypeJSON contains the JSON metadata for the
// struct [FineTuneTrainingTypeLoRaTrainingType]
type fineTuneTrainingTypeLoRaTrainingTypeJSON struct {
	LoraAlpha            apijson.Field
	LoraR                apijson.Field
	Type                 apijson.Field
	LoraDropout          apijson.Field
	LoraTrainableModules apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FineTuneTrainingTypeLoRaTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneTrainingTypeLoRaTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneTrainingTypeLoRaTrainingType) implementsFineTuneTrainingType() {}

type FineTuneTrainingTypeLoRaTrainingTypeType string

const (
	FineTuneTrainingTypeLoRaTrainingTypeTypeLora FineTuneTrainingTypeLoRaTrainingTypeType = "Lora"
)

func (r FineTuneTrainingTypeLoRaTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneTrainingTypeLoRaTrainingTypeTypeLora:
		return true
	}
	return false
}

type FineTuneTrainingTypeType string

const (
	FineTuneTrainingTypeTypeFull FineTuneTrainingTypeType = "Full"
	FineTuneTrainingTypeTypeLora FineTuneTrainingTypeType = "Lora"
)

func (r FineTuneTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneTrainingTypeTypeFull, FineTuneTrainingTypeTypeLora:
		return true
	}
	return false
}

type FineTuneListResponse struct {
	Data []FineTune               `json:"data,required"`
	JSON fineTuneListResponseJSON `json:"-"`
}

// fineTuneListResponseJSON contains the JSON metadata for the struct
// [FineTuneListResponse]
type fineTuneListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseJSON) RawJSON() string {
	return r.raw
}

type FineTuneDownloadResponse struct {
	ID             string                       `json:"id"`
	CheckpointStep int64                        `json:"checkpoint_step"`
	Filename       string                       `json:"filename"`
	Object         interface{}                  `json:"object"`
	Size           int64                        `json:"size"`
	JSON           fineTuneDownloadResponseJSON `json:"-"`
}

// fineTuneDownloadResponseJSON contains the JSON metadata for the struct
// [FineTuneDownloadResponse]
type fineTuneDownloadResponseJSON struct {
	ID             apijson.Field
	CheckpointStep apijson.Field
	Filename       apijson.Field
	Object         apijson.Field
	Size           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FineTuneDownloadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneDownloadResponseJSON) RawJSON() string {
	return r.raw
}

type FineTuneNewParams struct {
	// Name of the base model to run fine-tune job on
	Model param.Field[string] `json:"model,required"`
	// File-ID of a training file uploaded to the Together API
	TrainingFile param.Field[string] `json:"training_file,required"`
	// Number of training examples processed together (larger batches use more memory
	// but may train faster)
	BatchSize param.Field[int64] `json:"batch_size"`
	// The checkpoint identifier to continue training from a previous fine-tuning job.
	// Format `{$JOB_ID/$OUTPUT_MODEL_NAME}:{$STEP}`. The step value is optional,
	// without it the final checkpoint will be used.
	FromCheckpoint param.Field[string] `json:"from_checkpoint"`
	// Controls how quickly the model adapts to new information (too high may cause
	// instability, too low may slow convergence)
	LearningRate param.Field[float64]                      `json:"learning_rate"`
	LrScheduler  param.Field[FineTuneNewParamsLrScheduler] `json:"lr_scheduler"`
	// Max gradient norm to be used for gradient clipping. Set to 0 to disable.
	MaxGradNorm param.Field[float64] `json:"max_grad_norm"`
	// Number of intermediate model versions saved during training for evaluation
	NCheckpoints param.Field[int64] `json:"n_checkpoints"`
	// Number of complete passes through the training dataset (higher values may
	// improve results but increase cost and risk of overfitting)
	NEpochs param.Field[int64] `json:"n_epochs"`
	// Number of evaluations to be run on a given validation set during training
	NEvals param.Field[int64] `json:"n_evals"`
	// Suffix that will be added to your fine-tuned model name
	Suffix param.Field[string] `json:"suffix"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs param.Field[FineTuneNewParamsTrainOnInputsUnion] `json:"train_on_inputs"`
	TrainingType  param.Field[FineTuneNewParamsTrainingTypeUnion]  `json:"training_type"`
	// File-ID of a validation file uploaded to the Together API
	ValidationFile param.Field[string] `json:"validation_file"`
	// Integration key for tracking experiments and model metrics on W&B platform
	WandbAPIKey param.Field[string] `json:"wandb_api_key"`
	// The base URL of a dedicated Weights & Biases instance.
	WandbBaseURL param.Field[string] `json:"wandb_base_url"`
	// The Weights & Biases name for your run.
	WandbName param.Field[string] `json:"wandb_name"`
	// The Weights & Biases project for your run. If not specified, will use `together`
	// as the project name.
	WandbProjectName param.Field[string] `json:"wandb_project_name"`
	// The percent of steps at the start of training to linearly increase the learning
	// rate.
	WarmupRatio param.Field[float64] `json:"warmup_ratio"`
	// Weight decay
	WeightDecay param.Field[float64] `json:"weight_decay"`
}

func (r FineTuneNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FineTuneNewParamsLrScheduler struct {
	LrSchedulerType param.Field[string]                                      `json:"lr_scheduler_type,required"`
	LrSchedulerArgs param.Field[FineTuneNewParamsLrSchedulerLrSchedulerArgs] `json:"lr_scheduler_args"`
}

func (r FineTuneNewParamsLrScheduler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FineTuneNewParamsLrSchedulerLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio param.Field[float64] `json:"min_lr_ratio"`
}

func (r FineTuneNewParamsLrSchedulerLrSchedulerArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to mask the user messages in conversational data or prompts in
// instruction data.
//
// Satisfied by [shared.UnionBool], [FineTuneNewParamsTrainOnInputsString].
type FineTuneNewParamsTrainOnInputsUnion interface {
	ImplementsFineTuneNewParamsTrainOnInputsUnion()
}

type FineTuneNewParamsTrainOnInputsString string

const (
	FineTuneNewParamsTrainOnInputsStringAuto FineTuneNewParamsTrainOnInputsString = "auto"
)

func (r FineTuneNewParamsTrainOnInputsString) IsKnown() bool {
	switch r {
	case FineTuneNewParamsTrainOnInputsStringAuto:
		return true
	}
	return false
}

func (r FineTuneNewParamsTrainOnInputsString) ImplementsFineTuneNewParamsTrainOnInputsUnion() {}

type FineTuneNewParamsTrainingType struct {
	Type                 param.Field[FineTuneNewParamsTrainingTypeType] `json:"type,required"`
	LoraAlpha            param.Field[int64]                             `json:"lora_alpha"`
	LoraDropout          param.Field[float64]                           `json:"lora_dropout"`
	LoraR                param.Field[int64]                             `json:"lora_r"`
	LoraTrainableModules param.Field[string]                            `json:"lora_trainable_modules"`
}

func (r FineTuneNewParamsTrainingType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FineTuneNewParamsTrainingType) implementsFineTuneNewParamsTrainingTypeUnion() {}

// Satisfied by [FineTuneNewParamsTrainingTypeFullTrainingType],
// [FineTuneNewParamsTrainingTypeLoRaTrainingType],
// [FineTuneNewParamsTrainingType].
type FineTuneNewParamsTrainingTypeUnion interface {
	implementsFineTuneNewParamsTrainingTypeUnion()
}

type FineTuneNewParamsTrainingTypeFullTrainingType struct {
	Type param.Field[FineTuneNewParamsTrainingTypeFullTrainingTypeType] `json:"type,required"`
}

func (r FineTuneNewParamsTrainingTypeFullTrainingType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FineTuneNewParamsTrainingTypeFullTrainingType) implementsFineTuneNewParamsTrainingTypeUnion() {
}

type FineTuneNewParamsTrainingTypeFullTrainingTypeType string

const (
	FineTuneNewParamsTrainingTypeFullTrainingTypeTypeFull FineTuneNewParamsTrainingTypeFullTrainingTypeType = "Full"
)

func (r FineTuneNewParamsTrainingTypeFullTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneNewParamsTrainingTypeFullTrainingTypeTypeFull:
		return true
	}
	return false
}

type FineTuneNewParamsTrainingTypeLoRaTrainingType struct {
	LoraAlpha            param.Field[int64]                                             `json:"lora_alpha,required"`
	LoraR                param.Field[int64]                                             `json:"lora_r,required"`
	Type                 param.Field[FineTuneNewParamsTrainingTypeLoRaTrainingTypeType] `json:"type,required"`
	LoraDropout          param.Field[float64]                                           `json:"lora_dropout"`
	LoraTrainableModules param.Field[string]                                            `json:"lora_trainable_modules"`
}

func (r FineTuneNewParamsTrainingTypeLoRaTrainingType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FineTuneNewParamsTrainingTypeLoRaTrainingType) implementsFineTuneNewParamsTrainingTypeUnion() {
}

type FineTuneNewParamsTrainingTypeLoRaTrainingTypeType string

const (
	FineTuneNewParamsTrainingTypeLoRaTrainingTypeTypeLora FineTuneNewParamsTrainingTypeLoRaTrainingTypeType = "Lora"
)

func (r FineTuneNewParamsTrainingTypeLoRaTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneNewParamsTrainingTypeLoRaTrainingTypeTypeLora:
		return true
	}
	return false
}

type FineTuneNewParamsTrainingTypeType string

const (
	FineTuneNewParamsTrainingTypeTypeFull FineTuneNewParamsTrainingTypeType = "Full"
	FineTuneNewParamsTrainingTypeTypeLora FineTuneNewParamsTrainingTypeType = "Lora"
)

func (r FineTuneNewParamsTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneNewParamsTrainingTypeTypeFull, FineTuneNewParamsTrainingTypeTypeLora:
		return true
	}
	return false
}

type FineTuneDownloadParams struct {
	// Fine-tune ID to download. A string that starts with `ft-`.
	FtID param.Field[string] `query:"ft_id,required"`
	// Specifies checkpoint type to download - `merged` vs `adapter`. This field is
	// required if the checkpoint_step is not set.
	Checkpoint param.Field[FineTuneDownloadParamsCheckpoint] `query:"checkpoint"`
	// Specifies step number for checkpoint to download. Ignores `checkpoint` value if
	// set.
	CheckpointStep param.Field[int64] `query:"checkpoint_step"`
	// Specifies output file name for downloaded model. Defaults to
	// `$PWD/{model_name}.{extension}`.
	Output param.Field[string] `query:"output"`
}

// URLQuery serializes [FineTuneDownloadParams]'s query parameters as `url.Values`.
func (r FineTuneDownloadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies checkpoint type to download - `merged` vs `adapter`. This field is
// required if the checkpoint_step is not set.
type FineTuneDownloadParamsCheckpoint string

const (
	FineTuneDownloadParamsCheckpointMerged  FineTuneDownloadParamsCheckpoint = "merged"
	FineTuneDownloadParamsCheckpointAdapter FineTuneDownloadParamsCheckpoint = "adapter"
)

func (r FineTuneDownloadParamsCheckpoint) IsKnown() bool {
	switch r {
	case FineTuneDownloadParamsCheckpointMerged, FineTuneDownloadParamsCheckpointAdapter:
		return true
	}
	return false
}
