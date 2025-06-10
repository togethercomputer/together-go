// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

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

// Create a fine-tuning job with the provided model and training data.
func (r *FineTuneService) New(ctx context.Context, body FineTuneNewParams, opts ...option.RequestOption) (res *FineTuneNewResponse, err error) {
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

// List the metadata for all fine-tuning jobs. Returns a list of
// FinetuneResponseTruncated objects.
func (r *FineTuneService) List(ctx context.Context, opts ...option.RequestOption) (res *FineTuneListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "fine-tunes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a currently running fine-tuning job. Returns a FinetuneResponseTruncated
// object.
func (r *FineTuneService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTuneCancelResponse, err error) {
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
func (r *FineTuneService) ListEvents(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTuneListEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("fine-tunes/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the checkpoints for a single fine-tuning job.
func (r *FineTuneService) GetCheckpoints(ctx context.Context, id string, opts ...option.RequestOption) (res *FineTuneGetCheckpointsResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	NumCycles float64                   `json:"num_cycles,required"`
	JSON      cosineLrSchedulerArgsJSON `json:"-"`
}

// cosineLrSchedulerArgsJSON contains the JSON metadata for the struct
// [CosineLrSchedulerArgs]
type cosineLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CosineLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cosineLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r CosineLrSchedulerArgs) implementsLrSchedulerLrSchedulerArgs() {}

type CosineLrSchedulerArgsParam struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio param.Field[float64] `json:"min_lr_ratio,required"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles param.Field[float64] `json:"num_cycles,required"`
}

func (r CosineLrSchedulerArgsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CosineLrSchedulerArgsParam) implementsLrSchedulerLrSchedulerArgsUnionParam() {}

type FineTune struct {
	ID                   string                     `json:"id,required" format:"uuid"`
	Status               FineTuneStatus             `json:"status,required"`
	BatchSize            FineTuneBatchSizeUnion     `json:"batch_size"`
	CreatedAt            string                     `json:"created_at"`
	EpochsCompleted      int64                      `json:"epochs_completed"`
	EvalSteps            int64                      `json:"eval_steps"`
	Events               []FineTuneEvent            `json:"events"`
	FromCheckpoint       string                     `json:"from_checkpoint"`
	JobID                string                     `json:"job_id"`
	LearningRate         float64                    `json:"learning_rate"`
	LrScheduler          LrScheduler                `json:"lr_scheduler"`
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
	TrainingMethod       FineTuneTrainingMethod     `json:"training_method"`
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
	TrainingMethod       apijson.Field
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

// Union satisfied by [shared.UnionInt] or [FineTuneBatchSizeString].
type FineTuneBatchSizeUnion interface {
	ImplementsFineTuneBatchSizeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneBatchSizeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(FineTuneBatchSizeString("")),
		},
	)
}

type FineTuneBatchSizeString string

const (
	FineTuneBatchSizeStringMax FineTuneBatchSizeString = "max"
)

func (r FineTuneBatchSizeString) IsKnown() bool {
	switch r {
	case FineTuneBatchSizeStringMax:
		return true
	}
	return false
}

func (r FineTuneBatchSizeString) ImplementsFineTuneBatchSizeUnion() {}

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

type FineTuneTrainingMethod struct {
	Method  FineTuneTrainingMethodMethod `json:"method,required"`
	DpoBeta float64                      `json:"dpo_beta"`
	// This field can have the runtime type of [TrainingMethodSftTrainOnInputsUnion].
	TrainOnInputs interface{}                `json:"train_on_inputs"`
	JSON          fineTuneTrainingMethodJSON `json:"-"`
	union         FineTuneTrainingMethodUnion
}

// fineTuneTrainingMethodJSON contains the JSON metadata for the struct
// [FineTuneTrainingMethod]
type fineTuneTrainingMethodJSON struct {
	Method        apijson.Field
	DpoBeta       apijson.Field
	TrainOnInputs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r fineTuneTrainingMethodJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneTrainingMethod) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneTrainingMethod{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneTrainingMethodUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [TrainingMethodSft],
// [TrainingMethodDpo].
func (r FineTuneTrainingMethod) AsUnion() FineTuneTrainingMethodUnion {
	return r.union
}

// Union satisfied by [TrainingMethodSft] or [TrainingMethodDpo].
type FineTuneTrainingMethodUnion interface {
	implementsFineTuneTrainingMethod()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneTrainingMethodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrainingMethodSft{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrainingMethodDpo{}),
		},
	)
}

type FineTuneTrainingMethodMethod string

const (
	FineTuneTrainingMethodMethodSft FineTuneTrainingMethodMethod = "sft"
	FineTuneTrainingMethodMethodDpo FineTuneTrainingMethodMethod = "dpo"
)

func (r FineTuneTrainingMethodMethod) IsKnown() bool {
	switch r {
	case FineTuneTrainingMethodMethodSft, FineTuneTrainingMethodMethodDpo:
		return true
	}
	return false
}

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
// Possible runtime types of the union are [FullTrainingType], [LoRaTrainingType].
func (r FineTuneTrainingType) AsUnion() FineTuneTrainingTypeUnion {
	return r.union
}

// Union satisfied by [FullTrainingType] or [LoRaTrainingType].
type FineTuneTrainingTypeUnion interface {
	implementsFineTuneTrainingType()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneTrainingTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FullTrainingType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LoRaTrainingType{}),
		},
	)
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

type FineTuneEvent struct {
	CheckpointPath string              `json:"checkpoint_path,required"`
	CreatedAt      string              `json:"created_at,required"`
	Hash           string              `json:"hash,required"`
	Message        string              `json:"message,required"`
	ModelPath      string              `json:"model_path,required"`
	Object         FineTuneEventObject `json:"object,required"`
	ParamCount     int64               `json:"param_count,required"`
	Step           int64               `json:"step,required"`
	TokenCount     int64               `json:"token_count,required"`
	TotalSteps     int64               `json:"total_steps,required"`
	TrainingOffset int64               `json:"training_offset,required"`
	Type           FineTuneEventType   `json:"type,required"`
	WandbURL       string              `json:"wandb_url,required"`
	Level          FineTuneEventLevel  `json:"level,nullable"`
	JSON           fineTuneEventJSON   `json:"-"`
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

type FineTuneEventObject string

const (
	FineTuneEventObjectFineTuneEvent FineTuneEventObject = "fine-tune-event"
)

func (r FineTuneEventObject) IsKnown() bool {
	switch r {
	case FineTuneEventObjectFineTuneEvent:
		return true
	}
	return false
}

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

func (r FineTuneEventType) IsKnown() bool {
	switch r {
	case FineTuneEventTypeJobPending, FineTuneEventTypeJobStart, FineTuneEventTypeJobStopped, FineTuneEventTypeModelDownloading, FineTuneEventTypeModelDownloadComplete, FineTuneEventTypeTrainingDataDownloading, FineTuneEventTypeTrainingDataDownloadComplete, FineTuneEventTypeValidationDataDownloading, FineTuneEventTypeValidationDataDownloadComplete, FineTuneEventTypeWandbInit, FineTuneEventTypeTrainingStart, FineTuneEventTypeCheckpointSave, FineTuneEventTypeBillingLimit, FineTuneEventTypeEpochComplete, FineTuneEventTypeTrainingComplete, FineTuneEventTypeModelCompressing, FineTuneEventTypeModelCompressionComplete, FineTuneEventTypeModelUploading, FineTuneEventTypeModelUploadComplete, FineTuneEventTypeJobComplete, FineTuneEventTypeJobError, FineTuneEventTypeCancelRequested, FineTuneEventTypeJobRestarted, FineTuneEventTypeRefund, FineTuneEventTypeWarning:
		return true
	}
	return false
}

type FineTuneEventLevel string

const (
	FineTuneEventLevelInfo           FineTuneEventLevel = "info"
	FineTuneEventLevelWarning        FineTuneEventLevel = "warning"
	FineTuneEventLevelError          FineTuneEventLevel = "error"
	FineTuneEventLevelLegacyInfo     FineTuneEventLevel = "legacy_info"
	FineTuneEventLevelLegacyIwarning FineTuneEventLevel = "legacy_iwarning"
	FineTuneEventLevelLegacyIerror   FineTuneEventLevel = "legacy_ierror"
)

func (r FineTuneEventLevel) IsKnown() bool {
	switch r {
	case FineTuneEventLevelInfo, FineTuneEventLevelWarning, FineTuneEventLevelError, FineTuneEventLevelLegacyInfo, FineTuneEventLevelLegacyIwarning, FineTuneEventLevelLegacyIerror:
		return true
	}
	return false
}

type FullTrainingType struct {
	Type FullTrainingTypeType `json:"type,required"`
	JSON fullTrainingTypeJSON `json:"-"`
}

// fullTrainingTypeJSON contains the JSON metadata for the struct
// [FullTrainingType]
type fullTrainingTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FullTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fullTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r FullTrainingType) implementsFineTuneTrainingType() {}

func (r FullTrainingType) implementsFineTuneNewResponseTrainingType() {}

func (r FullTrainingType) implementsFineTuneListResponseDataTrainingType() {}

func (r FullTrainingType) implementsFineTuneCancelResponseTrainingType() {}

type FullTrainingTypeType string

const (
	FullTrainingTypeTypeFull FullTrainingTypeType = "Full"
)

func (r FullTrainingTypeType) IsKnown() bool {
	switch r {
	case FullTrainingTypeTypeFull:
		return true
	}
	return false
}

type FullTrainingTypeParam struct {
	Type param.Field[FullTrainingTypeType] `json:"type,required"`
}

func (r FullTrainingTypeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FullTrainingTypeParam) implementsFineTuneNewParamsTrainingTypeUnion() {}

type LinearLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64                   `json:"min_lr_ratio"`
	JSON       linearLrSchedulerArgsJSON `json:"-"`
}

// linearLrSchedulerArgsJSON contains the JSON metadata for the struct
// [LinearLrSchedulerArgs]
type linearLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinearLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linearLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r LinearLrSchedulerArgs) implementsLrSchedulerLrSchedulerArgs() {}

type LinearLrSchedulerArgsParam struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio param.Field[float64] `json:"min_lr_ratio"`
}

func (r LinearLrSchedulerArgsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LinearLrSchedulerArgsParam) implementsLrSchedulerLrSchedulerArgsUnionParam() {}

type LoRaTrainingType struct {
	LoraAlpha            int64                `json:"lora_alpha,required"`
	LoraR                int64                `json:"lora_r,required"`
	Type                 LoRaTrainingTypeType `json:"type,required"`
	LoraDropout          float64              `json:"lora_dropout"`
	LoraTrainableModules string               `json:"lora_trainable_modules"`
	JSON                 loRaTrainingTypeJSON `json:"-"`
}

// loRaTrainingTypeJSON contains the JSON metadata for the struct
// [LoRaTrainingType]
type loRaTrainingTypeJSON struct {
	LoraAlpha            apijson.Field
	LoraR                apijson.Field
	Type                 apijson.Field
	LoraDropout          apijson.Field
	LoraTrainableModules apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoRaTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loRaTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r LoRaTrainingType) implementsFineTuneTrainingType() {}

func (r LoRaTrainingType) implementsFineTuneNewResponseTrainingType() {}

func (r LoRaTrainingType) implementsFineTuneListResponseDataTrainingType() {}

func (r LoRaTrainingType) implementsFineTuneCancelResponseTrainingType() {}

type LoRaTrainingTypeType string

const (
	LoRaTrainingTypeTypeLora LoRaTrainingTypeType = "Lora"
)

func (r LoRaTrainingTypeType) IsKnown() bool {
	switch r {
	case LoRaTrainingTypeTypeLora:
		return true
	}
	return false
}

type LoRaTrainingTypeParam struct {
	LoraAlpha            param.Field[int64]                `json:"lora_alpha,required"`
	LoraR                param.Field[int64]                `json:"lora_r,required"`
	Type                 param.Field[LoRaTrainingTypeType] `json:"type,required"`
	LoraDropout          param.Field[float64]              `json:"lora_dropout"`
	LoraTrainableModules param.Field[string]               `json:"lora_trainable_modules"`
}

func (r LoRaTrainingTypeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LoRaTrainingTypeParam) implementsFineTuneNewParamsTrainingTypeUnion() {}

type LrScheduler struct {
	LrSchedulerType LrSchedulerLrSchedulerType `json:"lr_scheduler_type,required"`
	LrSchedulerArgs LrSchedulerLrSchedulerArgs `json:"lr_scheduler_args"`
	JSON            lrSchedulerJSON            `json:"-"`
}

// lrSchedulerJSON contains the JSON metadata for the struct [LrScheduler]
type lrSchedulerJSON struct {
	LrSchedulerType apijson.Field
	LrSchedulerArgs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LrScheduler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lrSchedulerJSON) RawJSON() string {
	return r.raw
}

type LrSchedulerLrSchedulerType string

const (
	LrSchedulerLrSchedulerTypeLinear LrSchedulerLrSchedulerType = "linear"
	LrSchedulerLrSchedulerTypeCosine LrSchedulerLrSchedulerType = "cosine"
)

func (r LrSchedulerLrSchedulerType) IsKnown() bool {
	switch r {
	case LrSchedulerLrSchedulerTypeLinear, LrSchedulerLrSchedulerTypeCosine:
		return true
	}
	return false
}

type LrSchedulerLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64                        `json:"num_cycles"`
	JSON      lrSchedulerLrSchedulerArgsJSON `json:"-"`
	union     LrSchedulerLrSchedulerArgsUnion
}

// lrSchedulerLrSchedulerArgsJSON contains the JSON metadata for the struct
// [LrSchedulerLrSchedulerArgs]
type lrSchedulerLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r lrSchedulerLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r *LrSchedulerLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	*r = LrSchedulerLrSchedulerArgs{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [LrSchedulerLrSchedulerArgsUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [LinearLrSchedulerArgs],
// [CosineLrSchedulerArgs].
func (r LrSchedulerLrSchedulerArgs) AsUnion() LrSchedulerLrSchedulerArgsUnion {
	return r.union
}

// Union satisfied by [LinearLrSchedulerArgs] or [CosineLrSchedulerArgs].
type LrSchedulerLrSchedulerArgsUnion interface {
	implementsLrSchedulerLrSchedulerArgs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LrSchedulerLrSchedulerArgsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LinearLrSchedulerArgs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CosineLrSchedulerArgs{}),
		},
	)
}

type LrSchedulerParam struct {
	LrSchedulerType param.Field[LrSchedulerLrSchedulerType]           `json:"lr_scheduler_type,required"`
	LrSchedulerArgs param.Field[LrSchedulerLrSchedulerArgsUnionParam] `json:"lr_scheduler_args"`
}

func (r LrSchedulerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LrSchedulerLrSchedulerArgsParam struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio param.Field[float64] `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles param.Field[float64] `json:"num_cycles"`
}

func (r LrSchedulerLrSchedulerArgsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LrSchedulerLrSchedulerArgsParam) implementsLrSchedulerLrSchedulerArgsUnionParam() {}

// Satisfied by [LinearLrSchedulerArgsParam], [CosineLrSchedulerArgsParam],
// [LrSchedulerLrSchedulerArgsParam].
type LrSchedulerLrSchedulerArgsUnionParam interface {
	implementsLrSchedulerLrSchedulerArgsUnionParam()
}

type TrainingMethodDpo struct {
	Method  TrainingMethodDpoMethod `json:"method,required"`
	DpoBeta float64                 `json:"dpo_beta"`
	JSON    trainingMethodDpoJSON   `json:"-"`
}

// trainingMethodDpoJSON contains the JSON metadata for the struct
// [TrainingMethodDpo]
type trainingMethodDpoJSON struct {
	Method      apijson.Field
	DpoBeta     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrainingMethodDpo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trainingMethodDpoJSON) RawJSON() string {
	return r.raw
}

func (r TrainingMethodDpo) implementsFineTuneTrainingMethod() {}

func (r TrainingMethodDpo) implementsFineTuneNewResponseTrainingMethod() {}

func (r TrainingMethodDpo) implementsFineTuneListResponseDataTrainingMethod() {}

func (r TrainingMethodDpo) implementsFineTuneCancelResponseTrainingMethod() {}

type TrainingMethodDpoMethod string

const (
	TrainingMethodDpoMethodDpo TrainingMethodDpoMethod = "dpo"
)

func (r TrainingMethodDpoMethod) IsKnown() bool {
	switch r {
	case TrainingMethodDpoMethodDpo:
		return true
	}
	return false
}

type TrainingMethodDpoParam struct {
	Method  param.Field[TrainingMethodDpoMethod] `json:"method,required"`
	DpoBeta param.Field[float64]                 `json:"dpo_beta"`
}

func (r TrainingMethodDpoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TrainingMethodDpoParam) implementsFineTuneNewParamsTrainingMethodUnion() {}

type TrainingMethodSft struct {
	Method TrainingMethodSftMethod `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs TrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	JSON          trainingMethodSftJSON               `json:"-"`
}

// trainingMethodSftJSON contains the JSON metadata for the struct
// [TrainingMethodSft]
type trainingMethodSftJSON struct {
	Method        apijson.Field
	TrainOnInputs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TrainingMethodSft) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trainingMethodSftJSON) RawJSON() string {
	return r.raw
}

func (r TrainingMethodSft) implementsFineTuneTrainingMethod() {}

func (r TrainingMethodSft) implementsFineTuneNewResponseTrainingMethod() {}

func (r TrainingMethodSft) implementsFineTuneListResponseDataTrainingMethod() {}

func (r TrainingMethodSft) implementsFineTuneCancelResponseTrainingMethod() {}

type TrainingMethodSftMethod string

const (
	TrainingMethodSftMethodSft TrainingMethodSftMethod = "sft"
)

func (r TrainingMethodSftMethod) IsKnown() bool {
	switch r {
	case TrainingMethodSftMethodSft:
		return true
	}
	return false
}

// Whether to mask the user messages in conversational data or prompts in
// instruction data.
//
// Union satisfied by [shared.UnionBool] or [TrainingMethodSftTrainOnInputsString].
type TrainingMethodSftTrainOnInputsUnion interface {
	ImplementsTrainingMethodSftTrainOnInputsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TrainingMethodSftTrainOnInputsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(TrainingMethodSftTrainOnInputsString("")),
		},
	)
}

type TrainingMethodSftTrainOnInputsString string

const (
	TrainingMethodSftTrainOnInputsStringAuto TrainingMethodSftTrainOnInputsString = "auto"
)

func (r TrainingMethodSftTrainOnInputsString) IsKnown() bool {
	switch r {
	case TrainingMethodSftTrainOnInputsStringAuto:
		return true
	}
	return false
}

func (r TrainingMethodSftTrainOnInputsString) ImplementsTrainingMethodSftTrainOnInputsUnion() {}

func (r TrainingMethodSftTrainOnInputsString) ImplementsTrainingMethodSftTrainOnInputsUnionParam() {}

type TrainingMethodSftParam struct {
	Method param.Field[TrainingMethodSftMethod] `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs param.Field[TrainingMethodSftTrainOnInputsUnionParam] `json:"train_on_inputs,required"`
}

func (r TrainingMethodSftParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TrainingMethodSftParam) implementsFineTuneNewParamsTrainingMethodUnion() {}

// Whether to mask the user messages in conversational data or prompts in
// instruction data.
//
// Satisfied by [shared.UnionBool], [TrainingMethodSftTrainOnInputsString].
type TrainingMethodSftTrainOnInputsUnionParam interface {
	ImplementsTrainingMethodSftTrainOnInputsUnionParam()
}

// A truncated version of the fine-tune response, used for POST /fine-tunes, GET
// /fine-tunes and POST /fine-tunes/{id}/cancel endpoints
type FineTuneNewResponse struct {
	// Unique identifier for the fine-tune job
	ID string `json:"id,required"`
	// Creation timestamp of the fine-tune job
	CreatedAt time.Time                 `json:"created_at,required" format:"date-time"`
	Status    FineTuneNewResponseStatus `json:"status,required"`
	// Last update timestamp of the fine-tune job
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Batch size used for training
	BatchSize int64 `json:"batch_size"`
	// Events related to this fine-tune job
	Events []FineTuneEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
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
	TrainingMethod FineTuneNewResponseTrainingMethod `json:"training_method"`
	// Type of training used (full or LoRA)
	TrainingType FineTuneNewResponseTrainingType `json:"training_type"`
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
	WeightDecay float64                 `json:"weight_decay"`
	JSON        fineTuneNewResponseJSON `json:"-"`
}

// fineTuneNewResponseJSON contains the JSON metadata for the struct
// [FineTuneNewResponse]
type fineTuneNewResponseJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	Status           apijson.Field
	UpdatedAt        apijson.Field
	BatchSize        apijson.Field
	Events           apijson.Field
	FromCheckpoint   apijson.Field
	LearningRate     apijson.Field
	LrScheduler      apijson.Field
	MaxGradNorm      apijson.Field
	Model            apijson.Field
	ModelOutputName  apijson.Field
	NCheckpoints     apijson.Field
	NEpochs          apijson.Field
	NEvals           apijson.Field
	OwnerAddress     apijson.Field
	Suffix           apijson.Field
	TokenCount       apijson.Field
	TotalPrice       apijson.Field
	TrainingFile     apijson.Field
	TrainingMethod   apijson.Field
	TrainingType     apijson.Field
	UserID           apijson.Field
	ValidationFile   apijson.Field
	WandbName        apijson.Field
	WandbProjectName apijson.Field
	WarmupRatio      apijson.Field
	WeightDecay      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FineTuneNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneNewResponseJSON) RawJSON() string {
	return r.raw
}

type FineTuneNewResponseStatus string

const (
	FineTuneNewResponseStatusPending         FineTuneNewResponseStatus = "pending"
	FineTuneNewResponseStatusQueued          FineTuneNewResponseStatus = "queued"
	FineTuneNewResponseStatusRunning         FineTuneNewResponseStatus = "running"
	FineTuneNewResponseStatusCompressing     FineTuneNewResponseStatus = "compressing"
	FineTuneNewResponseStatusUploading       FineTuneNewResponseStatus = "uploading"
	FineTuneNewResponseStatusCancelRequested FineTuneNewResponseStatus = "cancel_requested"
	FineTuneNewResponseStatusCancelled       FineTuneNewResponseStatus = "cancelled"
	FineTuneNewResponseStatusError           FineTuneNewResponseStatus = "error"
	FineTuneNewResponseStatusCompleted       FineTuneNewResponseStatus = "completed"
)

func (r FineTuneNewResponseStatus) IsKnown() bool {
	switch r {
	case FineTuneNewResponseStatusPending, FineTuneNewResponseStatusQueued, FineTuneNewResponseStatusRunning, FineTuneNewResponseStatusCompressing, FineTuneNewResponseStatusUploading, FineTuneNewResponseStatusCancelRequested, FineTuneNewResponseStatusCancelled, FineTuneNewResponseStatusError, FineTuneNewResponseStatusCompleted:
		return true
	}
	return false
}

// Method of training used
type FineTuneNewResponseTrainingMethod struct {
	Method  FineTuneNewResponseTrainingMethodMethod `json:"method,required"`
	DpoBeta float64                                 `json:"dpo_beta"`
	// This field can have the runtime type of [TrainingMethodSftTrainOnInputsUnion].
	TrainOnInputs interface{}                           `json:"train_on_inputs"`
	JSON          fineTuneNewResponseTrainingMethodJSON `json:"-"`
	union         FineTuneNewResponseTrainingMethodUnion
}

// fineTuneNewResponseTrainingMethodJSON contains the JSON metadata for the struct
// [FineTuneNewResponseTrainingMethod]
type fineTuneNewResponseTrainingMethodJSON struct {
	Method        apijson.Field
	DpoBeta       apijson.Field
	TrainOnInputs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r fineTuneNewResponseTrainingMethodJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneNewResponseTrainingMethod) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneNewResponseTrainingMethod{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneNewResponseTrainingMethodUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [TrainingMethodSft],
// [TrainingMethodDpo].
func (r FineTuneNewResponseTrainingMethod) AsUnion() FineTuneNewResponseTrainingMethodUnion {
	return r.union
}

// Method of training used
//
// Union satisfied by [TrainingMethodSft] or [TrainingMethodDpo].
type FineTuneNewResponseTrainingMethodUnion interface {
	implementsFineTuneNewResponseTrainingMethod()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneNewResponseTrainingMethodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrainingMethodSft{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrainingMethodDpo{}),
		},
	)
}

type FineTuneNewResponseTrainingMethodMethod string

const (
	FineTuneNewResponseTrainingMethodMethodSft FineTuneNewResponseTrainingMethodMethod = "sft"
	FineTuneNewResponseTrainingMethodMethodDpo FineTuneNewResponseTrainingMethodMethod = "dpo"
)

func (r FineTuneNewResponseTrainingMethodMethod) IsKnown() bool {
	switch r {
	case FineTuneNewResponseTrainingMethodMethodSft, FineTuneNewResponseTrainingMethodMethodDpo:
		return true
	}
	return false
}

// Type of training used (full or LoRA)
type FineTuneNewResponseTrainingType struct {
	Type                 FineTuneNewResponseTrainingTypeType `json:"type,required"`
	LoraAlpha            int64                               `json:"lora_alpha"`
	LoraDropout          float64                             `json:"lora_dropout"`
	LoraR                int64                               `json:"lora_r"`
	LoraTrainableModules string                              `json:"lora_trainable_modules"`
	JSON                 fineTuneNewResponseTrainingTypeJSON `json:"-"`
	union                FineTuneNewResponseTrainingTypeUnion
}

// fineTuneNewResponseTrainingTypeJSON contains the JSON metadata for the struct
// [FineTuneNewResponseTrainingType]
type fineTuneNewResponseTrainingTypeJSON struct {
	Type                 apijson.Field
	LoraAlpha            apijson.Field
	LoraDropout          apijson.Field
	LoraR                apijson.Field
	LoraTrainableModules apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r fineTuneNewResponseTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneNewResponseTrainingType) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneNewResponseTrainingType{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneNewResponseTrainingTypeUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [FullTrainingType], [LoRaTrainingType].
func (r FineTuneNewResponseTrainingType) AsUnion() FineTuneNewResponseTrainingTypeUnion {
	return r.union
}

// Type of training used (full or LoRA)
//
// Union satisfied by [FullTrainingType] or [LoRaTrainingType].
type FineTuneNewResponseTrainingTypeUnion interface {
	implementsFineTuneNewResponseTrainingType()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneNewResponseTrainingTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FullTrainingType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LoRaTrainingType{}),
		},
	)
}

type FineTuneNewResponseTrainingTypeType string

const (
	FineTuneNewResponseTrainingTypeTypeFull FineTuneNewResponseTrainingTypeType = "Full"
	FineTuneNewResponseTrainingTypeTypeLora FineTuneNewResponseTrainingTypeType = "Lora"
)

func (r FineTuneNewResponseTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneNewResponseTrainingTypeTypeFull, FineTuneNewResponseTrainingTypeTypeLora:
		return true
	}
	return false
}

type FineTuneListResponse struct {
	Data []FineTuneListResponseData `json:"data,required"`
	JSON fineTuneListResponseJSON   `json:"-"`
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

// A truncated version of the fine-tune response, used for POST /fine-tunes, GET
// /fine-tunes and POST /fine-tunes/{id}/cancel endpoints
type FineTuneListResponseData struct {
	// Unique identifier for the fine-tune job
	ID string `json:"id,required"`
	// Creation timestamp of the fine-tune job
	CreatedAt time.Time                      `json:"created_at,required" format:"date-time"`
	Status    FineTuneListResponseDataStatus `json:"status,required"`
	// Last update timestamp of the fine-tune job
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Batch size used for training
	BatchSize int64 `json:"batch_size"`
	// Events related to this fine-tune job
	Events []FineTuneEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
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
	TrainingMethod FineTuneListResponseDataTrainingMethod `json:"training_method"`
	// Type of training used (full or LoRA)
	TrainingType FineTuneListResponseDataTrainingType `json:"training_type"`
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
	WeightDecay float64                      `json:"weight_decay"`
	JSON        fineTuneListResponseDataJSON `json:"-"`
}

// fineTuneListResponseDataJSON contains the JSON metadata for the struct
// [FineTuneListResponseData]
type fineTuneListResponseDataJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	Status           apijson.Field
	UpdatedAt        apijson.Field
	BatchSize        apijson.Field
	Events           apijson.Field
	FromCheckpoint   apijson.Field
	LearningRate     apijson.Field
	LrScheduler      apijson.Field
	MaxGradNorm      apijson.Field
	Model            apijson.Field
	ModelOutputName  apijson.Field
	NCheckpoints     apijson.Field
	NEpochs          apijson.Field
	NEvals           apijson.Field
	OwnerAddress     apijson.Field
	Suffix           apijson.Field
	TokenCount       apijson.Field
	TotalPrice       apijson.Field
	TrainingFile     apijson.Field
	TrainingMethod   apijson.Field
	TrainingType     apijson.Field
	UserID           apijson.Field
	ValidationFile   apijson.Field
	WandbName        apijson.Field
	WandbProjectName apijson.Field
	WarmupRatio      apijson.Field
	WeightDecay      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FineTuneListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseDataJSON) RawJSON() string {
	return r.raw
}

type FineTuneListResponseDataStatus string

const (
	FineTuneListResponseDataStatusPending         FineTuneListResponseDataStatus = "pending"
	FineTuneListResponseDataStatusQueued          FineTuneListResponseDataStatus = "queued"
	FineTuneListResponseDataStatusRunning         FineTuneListResponseDataStatus = "running"
	FineTuneListResponseDataStatusCompressing     FineTuneListResponseDataStatus = "compressing"
	FineTuneListResponseDataStatusUploading       FineTuneListResponseDataStatus = "uploading"
	FineTuneListResponseDataStatusCancelRequested FineTuneListResponseDataStatus = "cancel_requested"
	FineTuneListResponseDataStatusCancelled       FineTuneListResponseDataStatus = "cancelled"
	FineTuneListResponseDataStatusError           FineTuneListResponseDataStatus = "error"
	FineTuneListResponseDataStatusCompleted       FineTuneListResponseDataStatus = "completed"
)

func (r FineTuneListResponseDataStatus) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataStatusPending, FineTuneListResponseDataStatusQueued, FineTuneListResponseDataStatusRunning, FineTuneListResponseDataStatusCompressing, FineTuneListResponseDataStatusUploading, FineTuneListResponseDataStatusCancelRequested, FineTuneListResponseDataStatusCancelled, FineTuneListResponseDataStatusError, FineTuneListResponseDataStatusCompleted:
		return true
	}
	return false
}

// Method of training used
type FineTuneListResponseDataTrainingMethod struct {
	Method  FineTuneListResponseDataTrainingMethodMethod `json:"method,required"`
	DpoBeta float64                                      `json:"dpo_beta"`
	// This field can have the runtime type of [TrainingMethodSftTrainOnInputsUnion].
	TrainOnInputs interface{}                                `json:"train_on_inputs"`
	JSON          fineTuneListResponseDataTrainingMethodJSON `json:"-"`
	union         FineTuneListResponseDataTrainingMethodUnion
}

// fineTuneListResponseDataTrainingMethodJSON contains the JSON metadata for the
// struct [FineTuneListResponseDataTrainingMethod]
type fineTuneListResponseDataTrainingMethodJSON struct {
	Method        apijson.Field
	DpoBeta       apijson.Field
	TrainOnInputs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r fineTuneListResponseDataTrainingMethodJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneListResponseDataTrainingMethod) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneListResponseDataTrainingMethod{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneListResponseDataTrainingMethodUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [TrainingMethodSft],
// [TrainingMethodDpo].
func (r FineTuneListResponseDataTrainingMethod) AsUnion() FineTuneListResponseDataTrainingMethodUnion {
	return r.union
}

// Method of training used
//
// Union satisfied by [TrainingMethodSft] or [TrainingMethodDpo].
type FineTuneListResponseDataTrainingMethodUnion interface {
	implementsFineTuneListResponseDataTrainingMethod()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneListResponseDataTrainingMethodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrainingMethodSft{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrainingMethodDpo{}),
		},
	)
}

type FineTuneListResponseDataTrainingMethodMethod string

const (
	FineTuneListResponseDataTrainingMethodMethodSft FineTuneListResponseDataTrainingMethodMethod = "sft"
	FineTuneListResponseDataTrainingMethodMethodDpo FineTuneListResponseDataTrainingMethodMethod = "dpo"
)

func (r FineTuneListResponseDataTrainingMethodMethod) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataTrainingMethodMethodSft, FineTuneListResponseDataTrainingMethodMethodDpo:
		return true
	}
	return false
}

// Type of training used (full or LoRA)
type FineTuneListResponseDataTrainingType struct {
	Type                 FineTuneListResponseDataTrainingTypeType `json:"type,required"`
	LoraAlpha            int64                                    `json:"lora_alpha"`
	LoraDropout          float64                                  `json:"lora_dropout"`
	LoraR                int64                                    `json:"lora_r"`
	LoraTrainableModules string                                   `json:"lora_trainable_modules"`
	JSON                 fineTuneListResponseDataTrainingTypeJSON `json:"-"`
	union                FineTuneListResponseDataTrainingTypeUnion
}

// fineTuneListResponseDataTrainingTypeJSON contains the JSON metadata for the
// struct [FineTuneListResponseDataTrainingType]
type fineTuneListResponseDataTrainingTypeJSON struct {
	Type                 apijson.Field
	LoraAlpha            apijson.Field
	LoraDropout          apijson.Field
	LoraR                apijson.Field
	LoraTrainableModules apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r fineTuneListResponseDataTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneListResponseDataTrainingType) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneListResponseDataTrainingType{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneListResponseDataTrainingTypeUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [FullTrainingType], [LoRaTrainingType].
func (r FineTuneListResponseDataTrainingType) AsUnion() FineTuneListResponseDataTrainingTypeUnion {
	return r.union
}

// Type of training used (full or LoRA)
//
// Union satisfied by [FullTrainingType] or [LoRaTrainingType].
type FineTuneListResponseDataTrainingTypeUnion interface {
	implementsFineTuneListResponseDataTrainingType()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneListResponseDataTrainingTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FullTrainingType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LoRaTrainingType{}),
		},
	)
}

type FineTuneListResponseDataTrainingTypeType string

const (
	FineTuneListResponseDataTrainingTypeTypeFull FineTuneListResponseDataTrainingTypeType = "Full"
	FineTuneListResponseDataTrainingTypeTypeLora FineTuneListResponseDataTrainingTypeType = "Lora"
)

func (r FineTuneListResponseDataTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataTrainingTypeTypeFull, FineTuneListResponseDataTrainingTypeTypeLora:
		return true
	}
	return false
}

// A truncated version of the fine-tune response, used for POST /fine-tunes, GET
// /fine-tunes and POST /fine-tunes/{id}/cancel endpoints
type FineTuneCancelResponse struct {
	// Unique identifier for the fine-tune job
	ID string `json:"id,required"`
	// Creation timestamp of the fine-tune job
	CreatedAt time.Time                    `json:"created_at,required" format:"date-time"`
	Status    FineTuneCancelResponseStatus `json:"status,required"`
	// Last update timestamp of the fine-tune job
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Batch size used for training
	BatchSize int64 `json:"batch_size"`
	// Events related to this fine-tune job
	Events []FineTuneEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
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
	TrainingMethod FineTuneCancelResponseTrainingMethod `json:"training_method"`
	// Type of training used (full or LoRA)
	TrainingType FineTuneCancelResponseTrainingType `json:"training_type"`
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
	WeightDecay float64                    `json:"weight_decay"`
	JSON        fineTuneCancelResponseJSON `json:"-"`
}

// fineTuneCancelResponseJSON contains the JSON metadata for the struct
// [FineTuneCancelResponse]
type fineTuneCancelResponseJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	Status           apijson.Field
	UpdatedAt        apijson.Field
	BatchSize        apijson.Field
	Events           apijson.Field
	FromCheckpoint   apijson.Field
	LearningRate     apijson.Field
	LrScheduler      apijson.Field
	MaxGradNorm      apijson.Field
	Model            apijson.Field
	ModelOutputName  apijson.Field
	NCheckpoints     apijson.Field
	NEpochs          apijson.Field
	NEvals           apijson.Field
	OwnerAddress     apijson.Field
	Suffix           apijson.Field
	TokenCount       apijson.Field
	TotalPrice       apijson.Field
	TrainingFile     apijson.Field
	TrainingMethod   apijson.Field
	TrainingType     apijson.Field
	UserID           apijson.Field
	ValidationFile   apijson.Field
	WandbName        apijson.Field
	WandbProjectName apijson.Field
	WarmupRatio      apijson.Field
	WeightDecay      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FineTuneCancelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneCancelResponseJSON) RawJSON() string {
	return r.raw
}

type FineTuneCancelResponseStatus string

const (
	FineTuneCancelResponseStatusPending         FineTuneCancelResponseStatus = "pending"
	FineTuneCancelResponseStatusQueued          FineTuneCancelResponseStatus = "queued"
	FineTuneCancelResponseStatusRunning         FineTuneCancelResponseStatus = "running"
	FineTuneCancelResponseStatusCompressing     FineTuneCancelResponseStatus = "compressing"
	FineTuneCancelResponseStatusUploading       FineTuneCancelResponseStatus = "uploading"
	FineTuneCancelResponseStatusCancelRequested FineTuneCancelResponseStatus = "cancel_requested"
	FineTuneCancelResponseStatusCancelled       FineTuneCancelResponseStatus = "cancelled"
	FineTuneCancelResponseStatusError           FineTuneCancelResponseStatus = "error"
	FineTuneCancelResponseStatusCompleted       FineTuneCancelResponseStatus = "completed"
)

func (r FineTuneCancelResponseStatus) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseStatusPending, FineTuneCancelResponseStatusQueued, FineTuneCancelResponseStatusRunning, FineTuneCancelResponseStatusCompressing, FineTuneCancelResponseStatusUploading, FineTuneCancelResponseStatusCancelRequested, FineTuneCancelResponseStatusCancelled, FineTuneCancelResponseStatusError, FineTuneCancelResponseStatusCompleted:
		return true
	}
	return false
}

// Method of training used
type FineTuneCancelResponseTrainingMethod struct {
	Method  FineTuneCancelResponseTrainingMethodMethod `json:"method,required"`
	DpoBeta float64                                    `json:"dpo_beta"`
	// This field can have the runtime type of [TrainingMethodSftTrainOnInputsUnion].
	TrainOnInputs interface{}                              `json:"train_on_inputs"`
	JSON          fineTuneCancelResponseTrainingMethodJSON `json:"-"`
	union         FineTuneCancelResponseTrainingMethodUnion
}

// fineTuneCancelResponseTrainingMethodJSON contains the JSON metadata for the
// struct [FineTuneCancelResponseTrainingMethod]
type fineTuneCancelResponseTrainingMethodJSON struct {
	Method        apijson.Field
	DpoBeta       apijson.Field
	TrainOnInputs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r fineTuneCancelResponseTrainingMethodJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneCancelResponseTrainingMethod) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneCancelResponseTrainingMethod{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneCancelResponseTrainingMethodUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [TrainingMethodSft],
// [TrainingMethodDpo].
func (r FineTuneCancelResponseTrainingMethod) AsUnion() FineTuneCancelResponseTrainingMethodUnion {
	return r.union
}

// Method of training used
//
// Union satisfied by [TrainingMethodSft] or [TrainingMethodDpo].
type FineTuneCancelResponseTrainingMethodUnion interface {
	implementsFineTuneCancelResponseTrainingMethod()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneCancelResponseTrainingMethodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrainingMethodSft{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrainingMethodDpo{}),
		},
	)
}

type FineTuneCancelResponseTrainingMethodMethod string

const (
	FineTuneCancelResponseTrainingMethodMethodSft FineTuneCancelResponseTrainingMethodMethod = "sft"
	FineTuneCancelResponseTrainingMethodMethodDpo FineTuneCancelResponseTrainingMethodMethod = "dpo"
)

func (r FineTuneCancelResponseTrainingMethodMethod) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseTrainingMethodMethodSft, FineTuneCancelResponseTrainingMethodMethodDpo:
		return true
	}
	return false
}

// Type of training used (full or LoRA)
type FineTuneCancelResponseTrainingType struct {
	Type                 FineTuneCancelResponseTrainingTypeType `json:"type,required"`
	LoraAlpha            int64                                  `json:"lora_alpha"`
	LoraDropout          float64                                `json:"lora_dropout"`
	LoraR                int64                                  `json:"lora_r"`
	LoraTrainableModules string                                 `json:"lora_trainable_modules"`
	JSON                 fineTuneCancelResponseTrainingTypeJSON `json:"-"`
	union                FineTuneCancelResponseTrainingTypeUnion
}

// fineTuneCancelResponseTrainingTypeJSON contains the JSON metadata for the struct
// [FineTuneCancelResponseTrainingType]
type fineTuneCancelResponseTrainingTypeJSON struct {
	Type                 apijson.Field
	LoraAlpha            apijson.Field
	LoraDropout          apijson.Field
	LoraR                apijson.Field
	LoraTrainableModules apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r fineTuneCancelResponseTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneCancelResponseTrainingType) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneCancelResponseTrainingType{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneCancelResponseTrainingTypeUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [FullTrainingType], [LoRaTrainingType].
func (r FineTuneCancelResponseTrainingType) AsUnion() FineTuneCancelResponseTrainingTypeUnion {
	return r.union
}

// Type of training used (full or LoRA)
//
// Union satisfied by [FullTrainingType] or [LoRaTrainingType].
type FineTuneCancelResponseTrainingTypeUnion interface {
	implementsFineTuneCancelResponseTrainingType()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneCancelResponseTrainingTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FullTrainingType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LoRaTrainingType{}),
		},
	)
}

type FineTuneCancelResponseTrainingTypeType string

const (
	FineTuneCancelResponseTrainingTypeTypeFull FineTuneCancelResponseTrainingTypeType = "Full"
	FineTuneCancelResponseTrainingTypeTypeLora FineTuneCancelResponseTrainingTypeType = "Lora"
)

func (r FineTuneCancelResponseTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseTrainingTypeTypeFull, FineTuneCancelResponseTrainingTypeTypeLora:
		return true
	}
	return false
}

type FineTuneDownloadResponse struct {
	ID             string                         `json:"id"`
	CheckpointStep int64                          `json:"checkpoint_step"`
	Filename       string                         `json:"filename"`
	Object         FineTuneDownloadResponseObject `json:"object,nullable"`
	Size           int64                          `json:"size"`
	JSON           fineTuneDownloadResponseJSON   `json:"-"`
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

type FineTuneDownloadResponseObject string

const (
	FineTuneDownloadResponseObjectLocal FineTuneDownloadResponseObject = "local"
)

func (r FineTuneDownloadResponseObject) IsKnown() bool {
	switch r {
	case FineTuneDownloadResponseObjectLocal:
		return true
	}
	return false
}

type FineTuneListEventsResponse struct {
	Data []FineTuneEvent                `json:"data,required"`
	JSON fineTuneListEventsResponseJSON `json:"-"`
}

// fineTuneListEventsResponseJSON contains the JSON metadata for the struct
// [FineTuneListEventsResponse]
type fineTuneListEventsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneListEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListEventsResponseJSON) RawJSON() string {
	return r.raw
}

type FineTuneGetCheckpointsResponse struct {
	Data []FineTuneGetCheckpointsResponseData `json:"data,required"`
	JSON fineTuneGetCheckpointsResponseJSON   `json:"-"`
}

// fineTuneGetCheckpointsResponseJSON contains the JSON metadata for the struct
// [FineTuneGetCheckpointsResponse]
type fineTuneGetCheckpointsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneGetCheckpointsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneGetCheckpointsResponseJSON) RawJSON() string {
	return r.raw
}

type FineTuneGetCheckpointsResponseData struct {
	CheckpointType string                                 `json:"checkpoint_type,required"`
	CreatedAt      string                                 `json:"created_at,required"`
	Path           string                                 `json:"path,required"`
	Step           int64                                  `json:"step,required"`
	JSON           fineTuneGetCheckpointsResponseDataJSON `json:"-"`
}

// fineTuneGetCheckpointsResponseDataJSON contains the JSON metadata for the struct
// [FineTuneGetCheckpointsResponseData]
type fineTuneGetCheckpointsResponseDataJSON struct {
	CheckpointType apijson.Field
	CreatedAt      apijson.Field
	Path           apijson.Field
	Step           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FineTuneGetCheckpointsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneGetCheckpointsResponseDataJSON) RawJSON() string {
	return r.raw
}

type FineTuneNewParams struct {
	// Name of the base model to run fine-tune job on
	Model param.Field[string] `json:"model,required"`
	// File-ID of a training file uploaded to the Together API
	TrainingFile param.Field[string] `json:"training_file,required"`
	// Number of training examples processed together (larger batches use more memory
	// but may train faster). Defaults to "max". We use training optimizations like
	// packing, so the effective batch size may be different than the value you set.
	BatchSize param.Field[FineTuneNewParamsBatchSizeUnion] `json:"batch_size"`
	// The checkpoint identifier to continue training from a previous fine-tuning job.
	// Format is `{$JOB_ID}` or `{$OUTPUT_MODEL_NAME}` or `{$JOB_ID}:{$STEP}` or
	// `{$OUTPUT_MODEL_NAME}:{$STEP}`. The step value is optional; without it, the
	// final checkpoint will be used.
	FromCheckpoint param.Field[string] `json:"from_checkpoint"`
	// Controls how quickly the model adapts to new information (too high may cause
	// instability, too low may slow convergence)
	LearningRate param.Field[float64] `json:"learning_rate"`
	// The learning rate scheduler to use. It specifies how the learning rate is
	// adjusted during training.
	LrScheduler param.Field[LrSchedulerParam] `json:"lr_scheduler"`
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
	// The training method to use. 'sft' for Supervised Fine-Tuning or 'dpo' for Direct
	// Preference Optimization.
	TrainingMethod param.Field[FineTuneNewParamsTrainingMethodUnion] `json:"training_method"`
	TrainingType   param.Field[FineTuneNewParamsTrainingTypeUnion]   `json:"training_type"`
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
	// Weight decay. Regularization parameter for the optimizer.
	WeightDecay param.Field[float64] `json:"weight_decay"`
}

func (r FineTuneNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Number of training examples processed together (larger batches use more memory
// but may train faster). Defaults to "max". We use training optimizations like
// packing, so the effective batch size may be different than the value you set.
//
// Satisfied by [shared.UnionInt], [FineTuneNewParamsBatchSizeString].
type FineTuneNewParamsBatchSizeUnion interface {
	ImplementsFineTuneNewParamsBatchSizeUnion()
}

type FineTuneNewParamsBatchSizeString string

const (
	FineTuneNewParamsBatchSizeStringMax FineTuneNewParamsBatchSizeString = "max"
)

func (r FineTuneNewParamsBatchSizeString) IsKnown() bool {
	switch r {
	case FineTuneNewParamsBatchSizeStringMax:
		return true
	}
	return false
}

func (r FineTuneNewParamsBatchSizeString) ImplementsFineTuneNewParamsBatchSizeUnion() {}

// Whether to mask the user messages in conversational data or prompts in
// instruction data.
//
// Satisfied by [shared.UnionBool], [FineTuneNewParamsTrainOnInputsString].
//
// Deprecated: deprecated
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

// The training method to use. 'sft' for Supervised Fine-Tuning or 'dpo' for Direct
// Preference Optimization.
type FineTuneNewParamsTrainingMethod struct {
	Method        param.Field[FineTuneNewParamsTrainingMethodMethod] `json:"method,required"`
	DpoBeta       param.Field[float64]                               `json:"dpo_beta"`
	TrainOnInputs param.Field[interface{}]                           `json:"train_on_inputs"`
}

func (r FineTuneNewParamsTrainingMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FineTuneNewParamsTrainingMethod) implementsFineTuneNewParamsTrainingMethodUnion() {}

// The training method to use. 'sft' for Supervised Fine-Tuning or 'dpo' for Direct
// Preference Optimization.
//
// Satisfied by [TrainingMethodSftParam], [TrainingMethodDpoParam],
// [FineTuneNewParamsTrainingMethod].
type FineTuneNewParamsTrainingMethodUnion interface {
	implementsFineTuneNewParamsTrainingMethodUnion()
}

type FineTuneNewParamsTrainingMethodMethod string

const (
	FineTuneNewParamsTrainingMethodMethodSft FineTuneNewParamsTrainingMethodMethod = "sft"
	FineTuneNewParamsTrainingMethodMethodDpo FineTuneNewParamsTrainingMethodMethod = "dpo"
)

func (r FineTuneNewParamsTrainingMethodMethod) IsKnown() bool {
	switch r {
	case FineTuneNewParamsTrainingMethodMethodSft, FineTuneNewParamsTrainingMethodMethodDpo:
		return true
	}
	return false
}

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

// Satisfied by [FullTrainingTypeParam], [LoRaTrainingTypeParam],
// [FineTuneNewParamsTrainingType].
type FineTuneNewParamsTrainingTypeUnion interface {
	implementsFineTuneNewParamsTrainingTypeUnion()
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
