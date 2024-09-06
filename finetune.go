// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
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
	ID                   string          `json:"id,required" format:"uuid"`
	Status               FineTuneStatus  `json:"status,required"`
	BatchSize            int64           `json:"batch_size"`
	CreatedAt            string          `json:"created_at"`
	EpochsCompleted      int64           `json:"epochs_completed"`
	EvalSteps            int64           `json:"eval_steps"`
	Events               []FineTuneEvent `json:"events"`
	JobID                string          `json:"job_id"`
	LearningRate         float64         `json:"learning_rate"`
	Lora                 bool            `json:"lora"`
	LoraAlpha            int64           `json:"lora_alpha"`
	LoraDropout          float64         `json:"lora_dropout"`
	LoraR                int64           `json:"lora_r"`
	LoraTrainableModules string          `json:"lora_trainable_modules"`
	Model                string          `json:"model"`
	ModelOutputName      string          `json:"model_output_name"`
	ModelOutputPath      string          `json:"model_output_path"`
	NCheckpoints         int64           `json:"n_checkpoints"`
	NEpochs              int64           `json:"n_epochs"`
	NEvals               int64           `json:"n_evals"`
	ParamCount           int64           `json:"param_count"`
	QueueDepth           int64           `json:"queue_depth"`
	TokenCount           int64           `json:"token_count"`
	TotalPrice           int64           `json:"total_price"`
	TrainingFile         string          `json:"training_file"`
	TrainingfileNumlines int64           `json:"trainingfile_numlines"`
	TrainingfileSize     int64           `json:"trainingfile_size"`
	UpdatedAt            string          `json:"updated_at"`
	ValidationFile       string          `json:"validation_file"`
	WandbProjectName     string          `json:"wandb_project_name"`
	WandbURL             string          `json:"wandb_url"`
	JSON                 fineTuneJSON    `json:"-"`
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
	JobID                apijson.Field
	LearningRate         apijson.Field
	Lora                 apijson.Field
	LoraAlpha            apijson.Field
	LoraDropout          apijson.Field
	LoraR                apijson.Field
	LoraTrainableModules apijson.Field
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
	TrainingFile         apijson.Field
	TrainingfileNumlines apijson.Field
	TrainingfileSize     apijson.Field
	UpdatedAt            apijson.Field
	ValidationFile       apijson.Field
	WandbProjectName     apijson.Field
	WandbURL             apijson.Field
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
	CreatedAt  string               `json:"created_at"`
	Hash       string               `json:"hash"`
	Level      FineTuneEventsLevel  `json:"level,nullable"`
	Message    string               `json:"message"`
	Object     FineTuneEventsObject `json:"object"`
	ParamCount int64                `json:"param_count"`
	TokenCount int64                `json:"token_count"`
	Type       FineTuneEventsType   `json:"type"`
	WandbURL   string               `json:"wandb_url"`
	JSON       fineTuneEventJSON    `json:"-"`
}

// fineTuneEventJSON contains the JSON metadata for the struct [FineTuneEvent]
type fineTuneEventJSON struct {
	CreatedAt   apijson.Field
	Hash        apijson.Field
	Level       apijson.Field
	Message     apijson.Field
	Object      apijson.Field
	ParamCount  apijson.Field
	TokenCount  apijson.Field
	Type        apijson.Field
	WandbURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneEventJSON) RawJSON() string {
	return r.raw
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

type FineTuneEventsObject string

const (
	FineTuneEventsObjectFinetuneEvent FineTuneEventsObject = "FinetuneEvent"
)

func (r FineTuneEventsObject) IsKnown() bool {
	switch r {
	case FineTuneEventsObjectFinetuneEvent:
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
	// Batch size for fine-tuning
	BatchSize param.Field[int64] `json:"batch_size"`
	// Learning rate multiplier to use for training
	LearningRate param.Field[float64] `json:"learning_rate"`
	// Whether to enable LoRA training. If not provided, full fine-tuning will be
	// applied.
	Lora param.Field[bool] `json:"lora"`
	// The alpha value for LoRA adapter training.
	LoraAlpha param.Field[int64] `json:"lora_alpha"`
	// The dropout probability for Lora layers.
	LoraDropout param.Field[float64] `json:"lora_dropout"`
	// Rank for LoRA adapter weights
	LoraR param.Field[int64] `json:"lora_r"`
	// A list of LoRA trainable modules, separated by a comma
	LoraTrainableModules param.Field[string] `json:"lora_trainable_modules"`
	// Number of checkpoints to save during fine-tuning
	NCheckpoints param.Field[int64] `json:"n_checkpoints"`
	// Number of epochs for fine-tuning
	NEpochs param.Field[int64] `json:"n_epochs"`
	// Number of evaluations to be run on a given validation set during training
	NEvals param.Field[int64] `json:"n_evals"`
	// Suffix that will be added to your fine-tuned model name
	Suffix param.Field[string] `json:"suffix"`
	// File-ID of a validation file uploaded to the Together API
	ValidationFile param.Field[string] `json:"validation_file"`
	// API key for Weights & Biases integration
	WandbAPIKey param.Field[string] `json:"wandb_api_key"`
}

func (r FineTuneNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FineTuneDownloadParams struct {
	// Fine-tune ID to download. A string that starts with `ft-`.
	FtID param.Field[string] `query:"ft_id,required"`
	// Specifies step number for checkpoint to download.
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
