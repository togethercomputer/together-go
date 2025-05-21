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
	BatchSize            FineTuneBatchSizeUnion     `json:"batch_size"`
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
	LrSchedulerType FineTuneLrSchedulerLrSchedulerType `json:"lr_scheduler_type,required"`
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

type FineTuneLrSchedulerLrSchedulerType string

const (
	FineTuneLrSchedulerLrSchedulerTypeLinear FineTuneLrSchedulerLrSchedulerType = "linear"
	FineTuneLrSchedulerLrSchedulerTypeCosine FineTuneLrSchedulerLrSchedulerType = "cosine"
)

func (r FineTuneLrSchedulerLrSchedulerType) IsKnown() bool {
	switch r {
	case FineTuneLrSchedulerLrSchedulerTypeLinear, FineTuneLrSchedulerLrSchedulerTypeCosine:
		return true
	}
	return false
}

type FineTuneLrSchedulerLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64                                `json:"num_cycles"`
	JSON      fineTuneLrSchedulerLrSchedulerArgsJSON `json:"-"`
	union     FineTuneLrSchedulerLrSchedulerArgsUnion
}

// fineTuneLrSchedulerLrSchedulerArgsJSON contains the JSON metadata for the struct
// [FineTuneLrSchedulerLrSchedulerArgs]
type fineTuneLrSchedulerLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r fineTuneLrSchedulerLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneLrSchedulerLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneLrSchedulerLrSchedulerArgs{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneLrSchedulerLrSchedulerArgsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs],
// [FineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
func (r FineTuneLrSchedulerLrSchedulerArgs) AsUnion() FineTuneLrSchedulerLrSchedulerArgsUnion {
	return r.union
}

// Union satisfied by [FineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs] or
// [FineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
type FineTuneLrSchedulerLrSchedulerArgsUnion interface {
	implementsFineTuneLrSchedulerLrSchedulerArgs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneLrSchedulerLrSchedulerArgsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs{}),
		},
	)
}

type FineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64                                                     `json:"min_lr_ratio"`
	JSON       fineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON `json:"-"`
}

// fineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON contains the JSON
// metadata for the struct
// [FineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs]
type fineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) implementsFineTuneLrSchedulerLrSchedulerArgs() {
}

type FineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64                                                     `json:"num_cycles"`
	JSON      fineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON `json:"-"`
}

// fineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON contains the JSON
// metadata for the struct
// [FineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs]
type fineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) implementsFineTuneLrSchedulerLrSchedulerArgs() {
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

type FineTuneTrainingMethod struct {
	Method  FineTuneTrainingMethodMethod `json:"method,required"`
	DpoBeta float64                      `json:"dpo_beta"`
	// This field can have the runtime type of
	// [FineTuneTrainingMethodTrainingMethodSftTrainOnInputsUnion].
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
// Possible runtime types of the union are
// [FineTuneTrainingMethodTrainingMethodSft],
// [FineTuneTrainingMethodTrainingMethodDpo].
func (r FineTuneTrainingMethod) AsUnion() FineTuneTrainingMethodUnion {
	return r.union
}

// Union satisfied by [FineTuneTrainingMethodTrainingMethodSft] or
// [FineTuneTrainingMethodTrainingMethodDpo].
type FineTuneTrainingMethodUnion interface {
	implementsFineTuneTrainingMethod()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneTrainingMethodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneTrainingMethodTrainingMethodSft{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneTrainingMethodTrainingMethodDpo{}),
		},
	)
}

type FineTuneTrainingMethodTrainingMethodSft struct {
	Method FineTuneTrainingMethodTrainingMethodSftMethod `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FineTuneTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	JSON          fineTuneTrainingMethodTrainingMethodSftJSON               `json:"-"`
}

// fineTuneTrainingMethodTrainingMethodSftJSON contains the JSON metadata for the
// struct [FineTuneTrainingMethodTrainingMethodSft]
type fineTuneTrainingMethodTrainingMethodSftJSON struct {
	Method        apijson.Field
	TrainOnInputs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FineTuneTrainingMethodTrainingMethodSft) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneTrainingMethodTrainingMethodSftJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneTrainingMethodTrainingMethodSft) implementsFineTuneTrainingMethod() {}

type FineTuneTrainingMethodTrainingMethodSftMethod string

const (
	FineTuneTrainingMethodTrainingMethodSftMethodSft FineTuneTrainingMethodTrainingMethodSftMethod = "sft"
)

func (r FineTuneTrainingMethodTrainingMethodSftMethod) IsKnown() bool {
	switch r {
	case FineTuneTrainingMethodTrainingMethodSftMethodSft:
		return true
	}
	return false
}

// Whether to mask the user messages in conversational data or prompts in
// instruction data.
//
// Union satisfied by [shared.UnionBool] or
// [FineTuneTrainingMethodTrainingMethodSftTrainOnInputsString].
type FineTuneTrainingMethodTrainingMethodSftTrainOnInputsUnion interface {
	ImplementsFineTuneTrainingMethodTrainingMethodSftTrainOnInputsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneTrainingMethodTrainingMethodSftTrainOnInputsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(FineTuneTrainingMethodTrainingMethodSftTrainOnInputsString("")),
		},
	)
}

type FineTuneTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FineTuneTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FineTuneTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

func (r FineTuneTrainingMethodTrainingMethodSftTrainOnInputsString) IsKnown() bool {
	switch r {
	case FineTuneTrainingMethodTrainingMethodSftTrainOnInputsStringAuto:
		return true
	}
	return false
}

func (r FineTuneTrainingMethodTrainingMethodSftTrainOnInputsString) ImplementsFineTuneTrainingMethodTrainingMethodSftTrainOnInputsUnion() {
}

type FineTuneTrainingMethodTrainingMethodDpo struct {
	Method  FineTuneTrainingMethodTrainingMethodDpoMethod `json:"method,required"`
	DpoBeta float64                                       `json:"dpo_beta"`
	JSON    fineTuneTrainingMethodTrainingMethodDpoJSON   `json:"-"`
}

// fineTuneTrainingMethodTrainingMethodDpoJSON contains the JSON metadata for the
// struct [FineTuneTrainingMethodTrainingMethodDpo]
type fineTuneTrainingMethodTrainingMethodDpoJSON struct {
	Method      apijson.Field
	DpoBeta     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneTrainingMethodTrainingMethodDpo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneTrainingMethodTrainingMethodDpoJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneTrainingMethodTrainingMethodDpo) implementsFineTuneTrainingMethod() {}

type FineTuneTrainingMethodTrainingMethodDpoMethod string

const (
	FineTuneTrainingMethodTrainingMethodDpoMethodDpo FineTuneTrainingMethodTrainingMethodDpoMethod = "dpo"
)

func (r FineTuneTrainingMethodTrainingMethodDpoMethod) IsKnown() bool {
	switch r {
	case FineTuneTrainingMethodTrainingMethodDpoMethodDpo:
		return true
	}
	return false
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
	Events []FineTuneNewResponseEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
	// Learning rate used for training
	LearningRate float64 `json:"learning_rate"`
	// Learning rate scheduler configuration
	LrScheduler FineTuneNewResponseLrScheduler `json:"lr_scheduler"`
	// Maximum gradient norm for clipping
	MaxGradNorm float64 `json:"max_grad_norm"`
	// Base model used for fine-tuning
	Model string `json:"model"`
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

type FineTuneNewResponseEvent struct {
	CheckpointPath string                          `json:"checkpoint_path,required"`
	CreatedAt      string                          `json:"created_at,required"`
	Hash           string                          `json:"hash,required"`
	Message        string                          `json:"message,required"`
	ModelPath      string                          `json:"model_path,required"`
	Object         FineTuneNewResponseEventsObject `json:"object,required"`
	ParamCount     int64                           `json:"param_count,required"`
	Step           int64                           `json:"step,required"`
	TokenCount     int64                           `json:"token_count,required"`
	TotalSteps     int64                           `json:"total_steps,required"`
	TrainingOffset int64                           `json:"training_offset,required"`
	Type           FineTuneNewResponseEventsType   `json:"type,required"`
	WandbURL       string                          `json:"wandb_url,required"`
	Level          FineTuneNewResponseEventsLevel  `json:"level,nullable"`
	JSON           fineTuneNewResponseEventJSON    `json:"-"`
}

// fineTuneNewResponseEventJSON contains the JSON metadata for the struct
// [FineTuneNewResponseEvent]
type fineTuneNewResponseEventJSON struct {
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

func (r *FineTuneNewResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneNewResponseEventJSON) RawJSON() string {
	return r.raw
}

type FineTuneNewResponseEventsObject string

const (
	FineTuneNewResponseEventsObjectFineTuneEvent FineTuneNewResponseEventsObject = "fine-tune-event"
)

func (r FineTuneNewResponseEventsObject) IsKnown() bool {
	switch r {
	case FineTuneNewResponseEventsObjectFineTuneEvent:
		return true
	}
	return false
}

type FineTuneNewResponseEventsType string

const (
	FineTuneNewResponseEventsTypeJobPending                     FineTuneNewResponseEventsType = "job_pending"
	FineTuneNewResponseEventsTypeJobStart                       FineTuneNewResponseEventsType = "job_start"
	FineTuneNewResponseEventsTypeJobStopped                     FineTuneNewResponseEventsType = "job_stopped"
	FineTuneNewResponseEventsTypeModelDownloading               FineTuneNewResponseEventsType = "model_downloading"
	FineTuneNewResponseEventsTypeModelDownloadComplete          FineTuneNewResponseEventsType = "model_download_complete"
	FineTuneNewResponseEventsTypeTrainingDataDownloading        FineTuneNewResponseEventsType = "training_data_downloading"
	FineTuneNewResponseEventsTypeTrainingDataDownloadComplete   FineTuneNewResponseEventsType = "training_data_download_complete"
	FineTuneNewResponseEventsTypeValidationDataDownloading      FineTuneNewResponseEventsType = "validation_data_downloading"
	FineTuneNewResponseEventsTypeValidationDataDownloadComplete FineTuneNewResponseEventsType = "validation_data_download_complete"
	FineTuneNewResponseEventsTypeWandbInit                      FineTuneNewResponseEventsType = "wandb_init"
	FineTuneNewResponseEventsTypeTrainingStart                  FineTuneNewResponseEventsType = "training_start"
	FineTuneNewResponseEventsTypeCheckpointSave                 FineTuneNewResponseEventsType = "checkpoint_save"
	FineTuneNewResponseEventsTypeBillingLimit                   FineTuneNewResponseEventsType = "billing_limit"
	FineTuneNewResponseEventsTypeEpochComplete                  FineTuneNewResponseEventsType = "epoch_complete"
	FineTuneNewResponseEventsTypeTrainingComplete               FineTuneNewResponseEventsType = "training_complete"
	FineTuneNewResponseEventsTypeModelCompressing               FineTuneNewResponseEventsType = "model_compressing"
	FineTuneNewResponseEventsTypeModelCompressionComplete       FineTuneNewResponseEventsType = "model_compression_complete"
	FineTuneNewResponseEventsTypeModelUploading                 FineTuneNewResponseEventsType = "model_uploading"
	FineTuneNewResponseEventsTypeModelUploadComplete            FineTuneNewResponseEventsType = "model_upload_complete"
	FineTuneNewResponseEventsTypeJobComplete                    FineTuneNewResponseEventsType = "job_complete"
	FineTuneNewResponseEventsTypeJobError                       FineTuneNewResponseEventsType = "job_error"
	FineTuneNewResponseEventsTypeCancelRequested                FineTuneNewResponseEventsType = "cancel_requested"
	FineTuneNewResponseEventsTypeJobRestarted                   FineTuneNewResponseEventsType = "job_restarted"
	FineTuneNewResponseEventsTypeRefund                         FineTuneNewResponseEventsType = "refund"
	FineTuneNewResponseEventsTypeWarning                        FineTuneNewResponseEventsType = "warning"
)

func (r FineTuneNewResponseEventsType) IsKnown() bool {
	switch r {
	case FineTuneNewResponseEventsTypeJobPending, FineTuneNewResponseEventsTypeJobStart, FineTuneNewResponseEventsTypeJobStopped, FineTuneNewResponseEventsTypeModelDownloading, FineTuneNewResponseEventsTypeModelDownloadComplete, FineTuneNewResponseEventsTypeTrainingDataDownloading, FineTuneNewResponseEventsTypeTrainingDataDownloadComplete, FineTuneNewResponseEventsTypeValidationDataDownloading, FineTuneNewResponseEventsTypeValidationDataDownloadComplete, FineTuneNewResponseEventsTypeWandbInit, FineTuneNewResponseEventsTypeTrainingStart, FineTuneNewResponseEventsTypeCheckpointSave, FineTuneNewResponseEventsTypeBillingLimit, FineTuneNewResponseEventsTypeEpochComplete, FineTuneNewResponseEventsTypeTrainingComplete, FineTuneNewResponseEventsTypeModelCompressing, FineTuneNewResponseEventsTypeModelCompressionComplete, FineTuneNewResponseEventsTypeModelUploading, FineTuneNewResponseEventsTypeModelUploadComplete, FineTuneNewResponseEventsTypeJobComplete, FineTuneNewResponseEventsTypeJobError, FineTuneNewResponseEventsTypeCancelRequested, FineTuneNewResponseEventsTypeJobRestarted, FineTuneNewResponseEventsTypeRefund, FineTuneNewResponseEventsTypeWarning:
		return true
	}
	return false
}

type FineTuneNewResponseEventsLevel string

const (
	FineTuneNewResponseEventsLevelInfo           FineTuneNewResponseEventsLevel = "info"
	FineTuneNewResponseEventsLevelWarning        FineTuneNewResponseEventsLevel = "warning"
	FineTuneNewResponseEventsLevelError          FineTuneNewResponseEventsLevel = "error"
	FineTuneNewResponseEventsLevelLegacyInfo     FineTuneNewResponseEventsLevel = "legacy_info"
	FineTuneNewResponseEventsLevelLegacyIwarning FineTuneNewResponseEventsLevel = "legacy_iwarning"
	FineTuneNewResponseEventsLevelLegacyIerror   FineTuneNewResponseEventsLevel = "legacy_ierror"
)

func (r FineTuneNewResponseEventsLevel) IsKnown() bool {
	switch r {
	case FineTuneNewResponseEventsLevelInfo, FineTuneNewResponseEventsLevelWarning, FineTuneNewResponseEventsLevelError, FineTuneNewResponseEventsLevelLegacyInfo, FineTuneNewResponseEventsLevelLegacyIwarning, FineTuneNewResponseEventsLevelLegacyIerror:
		return true
	}
	return false
}

// Learning rate scheduler configuration
type FineTuneNewResponseLrScheduler struct {
	LrSchedulerType FineTuneNewResponseLrSchedulerLrSchedulerType `json:"lr_scheduler_type,required"`
	LrSchedulerArgs FineTuneNewResponseLrSchedulerLrSchedulerArgs `json:"lr_scheduler_args"`
	JSON            fineTuneNewResponseLrSchedulerJSON            `json:"-"`
}

// fineTuneNewResponseLrSchedulerJSON contains the JSON metadata for the struct
// [FineTuneNewResponseLrScheduler]
type fineTuneNewResponseLrSchedulerJSON struct {
	LrSchedulerType apijson.Field
	LrSchedulerArgs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FineTuneNewResponseLrScheduler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneNewResponseLrSchedulerJSON) RawJSON() string {
	return r.raw
}

type FineTuneNewResponseLrSchedulerLrSchedulerType string

const (
	FineTuneNewResponseLrSchedulerLrSchedulerTypeLinear FineTuneNewResponseLrSchedulerLrSchedulerType = "linear"
	FineTuneNewResponseLrSchedulerLrSchedulerTypeCosine FineTuneNewResponseLrSchedulerLrSchedulerType = "cosine"
)

func (r FineTuneNewResponseLrSchedulerLrSchedulerType) IsKnown() bool {
	switch r {
	case FineTuneNewResponseLrSchedulerLrSchedulerTypeLinear, FineTuneNewResponseLrSchedulerLrSchedulerTypeCosine:
		return true
	}
	return false
}

type FineTuneNewResponseLrSchedulerLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64                                           `json:"num_cycles"`
	JSON      fineTuneNewResponseLrSchedulerLrSchedulerArgsJSON `json:"-"`
	union     FineTuneNewResponseLrSchedulerLrSchedulerArgsUnion
}

// fineTuneNewResponseLrSchedulerLrSchedulerArgsJSON contains the JSON metadata for
// the struct [FineTuneNewResponseLrSchedulerLrSchedulerArgs]
type fineTuneNewResponseLrSchedulerLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r fineTuneNewResponseLrSchedulerLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneNewResponseLrSchedulerLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneNewResponseLrSchedulerLrSchedulerArgs{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneNewResponseLrSchedulerLrSchedulerArgsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs],
// [FineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
func (r FineTuneNewResponseLrSchedulerLrSchedulerArgs) AsUnion() FineTuneNewResponseLrSchedulerLrSchedulerArgsUnion {
	return r.union
}

// Union satisfied by
// [FineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs] or
// [FineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
type FineTuneNewResponseLrSchedulerLrSchedulerArgsUnion interface {
	implementsFineTuneNewResponseLrSchedulerLrSchedulerArgs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneNewResponseLrSchedulerLrSchedulerArgsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs{}),
		},
	)
}

type FineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64                                                                `json:"min_lr_ratio"`
	JSON       fineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON `json:"-"`
}

// fineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON contains
// the JSON metadata for the struct
// [FineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs]
type fineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneNewResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) implementsFineTuneNewResponseLrSchedulerLrSchedulerArgs() {
}

type FineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64                                                                `json:"num_cycles"`
	JSON      fineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON `json:"-"`
}

// fineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON contains
// the JSON metadata for the struct
// [FineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs]
type fineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneNewResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) implementsFineTuneNewResponseLrSchedulerLrSchedulerArgs() {
}

// Method of training used
type FineTuneNewResponseTrainingMethod struct {
	Method  FineTuneNewResponseTrainingMethodMethod `json:"method,required"`
	DpoBeta float64                                 `json:"dpo_beta"`
	// This field can have the runtime type of
	// [FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion].
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
// Possible runtime types of the union are
// [FineTuneNewResponseTrainingMethodTrainingMethodSft],
// [FineTuneNewResponseTrainingMethodTrainingMethodDpo].
func (r FineTuneNewResponseTrainingMethod) AsUnion() FineTuneNewResponseTrainingMethodUnion {
	return r.union
}

// Method of training used
//
// Union satisfied by [FineTuneNewResponseTrainingMethodTrainingMethodSft] or
// [FineTuneNewResponseTrainingMethodTrainingMethodDpo].
type FineTuneNewResponseTrainingMethodUnion interface {
	implementsFineTuneNewResponseTrainingMethod()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneNewResponseTrainingMethodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneNewResponseTrainingMethodTrainingMethodSft{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneNewResponseTrainingMethodTrainingMethodDpo{}),
		},
	)
}

type FineTuneNewResponseTrainingMethodTrainingMethodSft struct {
	Method FineTuneNewResponseTrainingMethodTrainingMethodSftMethod `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	JSON          fineTuneNewResponseTrainingMethodTrainingMethodSftJSON               `json:"-"`
}

// fineTuneNewResponseTrainingMethodTrainingMethodSftJSON contains the JSON
// metadata for the struct [FineTuneNewResponseTrainingMethodTrainingMethodSft]
type fineTuneNewResponseTrainingMethodTrainingMethodSftJSON struct {
	Method        apijson.Field
	TrainOnInputs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FineTuneNewResponseTrainingMethodTrainingMethodSft) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneNewResponseTrainingMethodTrainingMethodSftJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneNewResponseTrainingMethodTrainingMethodSft) implementsFineTuneNewResponseTrainingMethod() {
}

type FineTuneNewResponseTrainingMethodTrainingMethodSftMethod string

const (
	FineTuneNewResponseTrainingMethodTrainingMethodSftMethodSft FineTuneNewResponseTrainingMethodTrainingMethodSftMethod = "sft"
)

func (r FineTuneNewResponseTrainingMethodTrainingMethodSftMethod) IsKnown() bool {
	switch r {
	case FineTuneNewResponseTrainingMethodTrainingMethodSftMethodSft:
		return true
	}
	return false
}

// Whether to mask the user messages in conversational data or prompts in
// instruction data.
//
// Union satisfied by [shared.UnionBool] or
// [FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString].
type FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion interface {
	ImplementsFineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString("")),
		},
	)
}

type FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

func (r FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString) IsKnown() bool {
	switch r {
	case FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsStringAuto:
		return true
	}
	return false
}

func (r FineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsString) ImplementsFineTuneNewResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion() {
}

type FineTuneNewResponseTrainingMethodTrainingMethodDpo struct {
	Method  FineTuneNewResponseTrainingMethodTrainingMethodDpoMethod `json:"method,required"`
	DpoBeta float64                                                  `json:"dpo_beta"`
	JSON    fineTuneNewResponseTrainingMethodTrainingMethodDpoJSON   `json:"-"`
}

// fineTuneNewResponseTrainingMethodTrainingMethodDpoJSON contains the JSON
// metadata for the struct [FineTuneNewResponseTrainingMethodTrainingMethodDpo]
type fineTuneNewResponseTrainingMethodTrainingMethodDpoJSON struct {
	Method      apijson.Field
	DpoBeta     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneNewResponseTrainingMethodTrainingMethodDpo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneNewResponseTrainingMethodTrainingMethodDpoJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneNewResponseTrainingMethodTrainingMethodDpo) implementsFineTuneNewResponseTrainingMethod() {
}

type FineTuneNewResponseTrainingMethodTrainingMethodDpoMethod string

const (
	FineTuneNewResponseTrainingMethodTrainingMethodDpoMethodDpo FineTuneNewResponseTrainingMethodTrainingMethodDpoMethod = "dpo"
)

func (r FineTuneNewResponseTrainingMethodTrainingMethodDpoMethod) IsKnown() bool {
	switch r {
	case FineTuneNewResponseTrainingMethodTrainingMethodDpoMethodDpo:
		return true
	}
	return false
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
// Possible runtime types of the union are
// [FineTuneNewResponseTrainingTypeFullTrainingType],
// [FineTuneNewResponseTrainingTypeLoRaTrainingType].
func (r FineTuneNewResponseTrainingType) AsUnion() FineTuneNewResponseTrainingTypeUnion {
	return r.union
}

// Type of training used (full or LoRA)
//
// Union satisfied by [FineTuneNewResponseTrainingTypeFullTrainingType] or
// [FineTuneNewResponseTrainingTypeLoRaTrainingType].
type FineTuneNewResponseTrainingTypeUnion interface {
	implementsFineTuneNewResponseTrainingType()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneNewResponseTrainingTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneNewResponseTrainingTypeFullTrainingType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneNewResponseTrainingTypeLoRaTrainingType{}),
		},
	)
}

type FineTuneNewResponseTrainingTypeFullTrainingType struct {
	Type FineTuneNewResponseTrainingTypeFullTrainingTypeType `json:"type,required"`
	JSON fineTuneNewResponseTrainingTypeFullTrainingTypeJSON `json:"-"`
}

// fineTuneNewResponseTrainingTypeFullTrainingTypeJSON contains the JSON metadata
// for the struct [FineTuneNewResponseTrainingTypeFullTrainingType]
type fineTuneNewResponseTrainingTypeFullTrainingTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneNewResponseTrainingTypeFullTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneNewResponseTrainingTypeFullTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneNewResponseTrainingTypeFullTrainingType) implementsFineTuneNewResponseTrainingType() {
}

type FineTuneNewResponseTrainingTypeFullTrainingTypeType string

const (
	FineTuneNewResponseTrainingTypeFullTrainingTypeTypeFull FineTuneNewResponseTrainingTypeFullTrainingTypeType = "Full"
)

func (r FineTuneNewResponseTrainingTypeFullTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneNewResponseTrainingTypeFullTrainingTypeTypeFull:
		return true
	}
	return false
}

type FineTuneNewResponseTrainingTypeLoRaTrainingType struct {
	LoraAlpha            int64                                               `json:"lora_alpha,required"`
	LoraR                int64                                               `json:"lora_r,required"`
	Type                 FineTuneNewResponseTrainingTypeLoRaTrainingTypeType `json:"type,required"`
	LoraDropout          float64                                             `json:"lora_dropout"`
	LoraTrainableModules string                                              `json:"lora_trainable_modules"`
	JSON                 fineTuneNewResponseTrainingTypeLoRaTrainingTypeJSON `json:"-"`
}

// fineTuneNewResponseTrainingTypeLoRaTrainingTypeJSON contains the JSON metadata
// for the struct [FineTuneNewResponseTrainingTypeLoRaTrainingType]
type fineTuneNewResponseTrainingTypeLoRaTrainingTypeJSON struct {
	LoraAlpha            apijson.Field
	LoraR                apijson.Field
	Type                 apijson.Field
	LoraDropout          apijson.Field
	LoraTrainableModules apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FineTuneNewResponseTrainingTypeLoRaTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneNewResponseTrainingTypeLoRaTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneNewResponseTrainingTypeLoRaTrainingType) implementsFineTuneNewResponseTrainingType() {
}

type FineTuneNewResponseTrainingTypeLoRaTrainingTypeType string

const (
	FineTuneNewResponseTrainingTypeLoRaTrainingTypeTypeLora FineTuneNewResponseTrainingTypeLoRaTrainingTypeType = "Lora"
)

func (r FineTuneNewResponseTrainingTypeLoRaTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneNewResponseTrainingTypeLoRaTrainingTypeTypeLora:
		return true
	}
	return false
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
	Events []FineTuneListResponseDataEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
	// Learning rate used for training
	LearningRate float64 `json:"learning_rate"`
	// Learning rate scheduler configuration
	LrScheduler FineTuneListResponseDataLrScheduler `json:"lr_scheduler"`
	// Maximum gradient norm for clipping
	MaxGradNorm float64 `json:"max_grad_norm"`
	// Base model used for fine-tuning
	Model string `json:"model"`
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

type FineTuneListResponseDataEvent struct {
	CheckpointPath string                               `json:"checkpoint_path,required"`
	CreatedAt      string                               `json:"created_at,required"`
	Hash           string                               `json:"hash,required"`
	Message        string                               `json:"message,required"`
	ModelPath      string                               `json:"model_path,required"`
	Object         FineTuneListResponseDataEventsObject `json:"object,required"`
	ParamCount     int64                                `json:"param_count,required"`
	Step           int64                                `json:"step,required"`
	TokenCount     int64                                `json:"token_count,required"`
	TotalSteps     int64                                `json:"total_steps,required"`
	TrainingOffset int64                                `json:"training_offset,required"`
	Type           FineTuneListResponseDataEventsType   `json:"type,required"`
	WandbURL       string                               `json:"wandb_url,required"`
	Level          FineTuneListResponseDataEventsLevel  `json:"level,nullable"`
	JSON           fineTuneListResponseDataEventJSON    `json:"-"`
}

// fineTuneListResponseDataEventJSON contains the JSON metadata for the struct
// [FineTuneListResponseDataEvent]
type fineTuneListResponseDataEventJSON struct {
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

func (r *FineTuneListResponseDataEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseDataEventJSON) RawJSON() string {
	return r.raw
}

type FineTuneListResponseDataEventsObject string

const (
	FineTuneListResponseDataEventsObjectFineTuneEvent FineTuneListResponseDataEventsObject = "fine-tune-event"
)

func (r FineTuneListResponseDataEventsObject) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataEventsObjectFineTuneEvent:
		return true
	}
	return false
}

type FineTuneListResponseDataEventsType string

const (
	FineTuneListResponseDataEventsTypeJobPending                     FineTuneListResponseDataEventsType = "job_pending"
	FineTuneListResponseDataEventsTypeJobStart                       FineTuneListResponseDataEventsType = "job_start"
	FineTuneListResponseDataEventsTypeJobStopped                     FineTuneListResponseDataEventsType = "job_stopped"
	FineTuneListResponseDataEventsTypeModelDownloading               FineTuneListResponseDataEventsType = "model_downloading"
	FineTuneListResponseDataEventsTypeModelDownloadComplete          FineTuneListResponseDataEventsType = "model_download_complete"
	FineTuneListResponseDataEventsTypeTrainingDataDownloading        FineTuneListResponseDataEventsType = "training_data_downloading"
	FineTuneListResponseDataEventsTypeTrainingDataDownloadComplete   FineTuneListResponseDataEventsType = "training_data_download_complete"
	FineTuneListResponseDataEventsTypeValidationDataDownloading      FineTuneListResponseDataEventsType = "validation_data_downloading"
	FineTuneListResponseDataEventsTypeValidationDataDownloadComplete FineTuneListResponseDataEventsType = "validation_data_download_complete"
	FineTuneListResponseDataEventsTypeWandbInit                      FineTuneListResponseDataEventsType = "wandb_init"
	FineTuneListResponseDataEventsTypeTrainingStart                  FineTuneListResponseDataEventsType = "training_start"
	FineTuneListResponseDataEventsTypeCheckpointSave                 FineTuneListResponseDataEventsType = "checkpoint_save"
	FineTuneListResponseDataEventsTypeBillingLimit                   FineTuneListResponseDataEventsType = "billing_limit"
	FineTuneListResponseDataEventsTypeEpochComplete                  FineTuneListResponseDataEventsType = "epoch_complete"
	FineTuneListResponseDataEventsTypeTrainingComplete               FineTuneListResponseDataEventsType = "training_complete"
	FineTuneListResponseDataEventsTypeModelCompressing               FineTuneListResponseDataEventsType = "model_compressing"
	FineTuneListResponseDataEventsTypeModelCompressionComplete       FineTuneListResponseDataEventsType = "model_compression_complete"
	FineTuneListResponseDataEventsTypeModelUploading                 FineTuneListResponseDataEventsType = "model_uploading"
	FineTuneListResponseDataEventsTypeModelUploadComplete            FineTuneListResponseDataEventsType = "model_upload_complete"
	FineTuneListResponseDataEventsTypeJobComplete                    FineTuneListResponseDataEventsType = "job_complete"
	FineTuneListResponseDataEventsTypeJobError                       FineTuneListResponseDataEventsType = "job_error"
	FineTuneListResponseDataEventsTypeCancelRequested                FineTuneListResponseDataEventsType = "cancel_requested"
	FineTuneListResponseDataEventsTypeJobRestarted                   FineTuneListResponseDataEventsType = "job_restarted"
	FineTuneListResponseDataEventsTypeRefund                         FineTuneListResponseDataEventsType = "refund"
	FineTuneListResponseDataEventsTypeWarning                        FineTuneListResponseDataEventsType = "warning"
)

func (r FineTuneListResponseDataEventsType) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataEventsTypeJobPending, FineTuneListResponseDataEventsTypeJobStart, FineTuneListResponseDataEventsTypeJobStopped, FineTuneListResponseDataEventsTypeModelDownloading, FineTuneListResponseDataEventsTypeModelDownloadComplete, FineTuneListResponseDataEventsTypeTrainingDataDownloading, FineTuneListResponseDataEventsTypeTrainingDataDownloadComplete, FineTuneListResponseDataEventsTypeValidationDataDownloading, FineTuneListResponseDataEventsTypeValidationDataDownloadComplete, FineTuneListResponseDataEventsTypeWandbInit, FineTuneListResponseDataEventsTypeTrainingStart, FineTuneListResponseDataEventsTypeCheckpointSave, FineTuneListResponseDataEventsTypeBillingLimit, FineTuneListResponseDataEventsTypeEpochComplete, FineTuneListResponseDataEventsTypeTrainingComplete, FineTuneListResponseDataEventsTypeModelCompressing, FineTuneListResponseDataEventsTypeModelCompressionComplete, FineTuneListResponseDataEventsTypeModelUploading, FineTuneListResponseDataEventsTypeModelUploadComplete, FineTuneListResponseDataEventsTypeJobComplete, FineTuneListResponseDataEventsTypeJobError, FineTuneListResponseDataEventsTypeCancelRequested, FineTuneListResponseDataEventsTypeJobRestarted, FineTuneListResponseDataEventsTypeRefund, FineTuneListResponseDataEventsTypeWarning:
		return true
	}
	return false
}

type FineTuneListResponseDataEventsLevel string

const (
	FineTuneListResponseDataEventsLevelInfo           FineTuneListResponseDataEventsLevel = "info"
	FineTuneListResponseDataEventsLevelWarning        FineTuneListResponseDataEventsLevel = "warning"
	FineTuneListResponseDataEventsLevelError          FineTuneListResponseDataEventsLevel = "error"
	FineTuneListResponseDataEventsLevelLegacyInfo     FineTuneListResponseDataEventsLevel = "legacy_info"
	FineTuneListResponseDataEventsLevelLegacyIwarning FineTuneListResponseDataEventsLevel = "legacy_iwarning"
	FineTuneListResponseDataEventsLevelLegacyIerror   FineTuneListResponseDataEventsLevel = "legacy_ierror"
)

func (r FineTuneListResponseDataEventsLevel) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataEventsLevelInfo, FineTuneListResponseDataEventsLevelWarning, FineTuneListResponseDataEventsLevelError, FineTuneListResponseDataEventsLevelLegacyInfo, FineTuneListResponseDataEventsLevelLegacyIwarning, FineTuneListResponseDataEventsLevelLegacyIerror:
		return true
	}
	return false
}

// Learning rate scheduler configuration
type FineTuneListResponseDataLrScheduler struct {
	LrSchedulerType FineTuneListResponseDataLrSchedulerLrSchedulerType `json:"lr_scheduler_type,required"`
	LrSchedulerArgs FineTuneListResponseDataLrSchedulerLrSchedulerArgs `json:"lr_scheduler_args"`
	JSON            fineTuneListResponseDataLrSchedulerJSON            `json:"-"`
}

// fineTuneListResponseDataLrSchedulerJSON contains the JSON metadata for the
// struct [FineTuneListResponseDataLrScheduler]
type fineTuneListResponseDataLrSchedulerJSON struct {
	LrSchedulerType apijson.Field
	LrSchedulerArgs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FineTuneListResponseDataLrScheduler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseDataLrSchedulerJSON) RawJSON() string {
	return r.raw
}

type FineTuneListResponseDataLrSchedulerLrSchedulerType string

const (
	FineTuneListResponseDataLrSchedulerLrSchedulerTypeLinear FineTuneListResponseDataLrSchedulerLrSchedulerType = "linear"
	FineTuneListResponseDataLrSchedulerLrSchedulerTypeCosine FineTuneListResponseDataLrSchedulerLrSchedulerType = "cosine"
)

func (r FineTuneListResponseDataLrSchedulerLrSchedulerType) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataLrSchedulerLrSchedulerTypeLinear, FineTuneListResponseDataLrSchedulerLrSchedulerTypeCosine:
		return true
	}
	return false
}

type FineTuneListResponseDataLrSchedulerLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64                                                `json:"num_cycles"`
	JSON      fineTuneListResponseDataLrSchedulerLrSchedulerArgsJSON `json:"-"`
	union     FineTuneListResponseDataLrSchedulerLrSchedulerArgsUnion
}

// fineTuneListResponseDataLrSchedulerLrSchedulerArgsJSON contains the JSON
// metadata for the struct [FineTuneListResponseDataLrSchedulerLrSchedulerArgs]
type fineTuneListResponseDataLrSchedulerLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r fineTuneListResponseDataLrSchedulerLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneListResponseDataLrSchedulerLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneListResponseDataLrSchedulerLrSchedulerArgs{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneListResponseDataLrSchedulerLrSchedulerArgsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs],
// [FineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
func (r FineTuneListResponseDataLrSchedulerLrSchedulerArgs) AsUnion() FineTuneListResponseDataLrSchedulerLrSchedulerArgsUnion {
	return r.union
}

// Union satisfied by
// [FineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs] or
// [FineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
type FineTuneListResponseDataLrSchedulerLrSchedulerArgsUnion interface {
	implementsFineTuneListResponseDataLrSchedulerLrSchedulerArgs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneListResponseDataLrSchedulerLrSchedulerArgsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs{}),
		},
	)
}

type FineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64                                                                     `json:"min_lr_ratio"`
	JSON       fineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON `json:"-"`
}

// fineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON
// contains the JSON metadata for the struct
// [FineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs]
type fineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneListResponseDataLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) implementsFineTuneListResponseDataLrSchedulerLrSchedulerArgs() {
}

type FineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64                                                                     `json:"num_cycles"`
	JSON      fineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON `json:"-"`
}

// fineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON
// contains the JSON metadata for the struct
// [FineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs]
type fineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneListResponseDataLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) implementsFineTuneListResponseDataLrSchedulerLrSchedulerArgs() {
}

// Method of training used
type FineTuneListResponseDataTrainingMethod struct {
	Method  FineTuneListResponseDataTrainingMethodMethod `json:"method,required"`
	DpoBeta float64                                      `json:"dpo_beta"`
	// This field can have the runtime type of
	// [FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion].
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
// Possible runtime types of the union are
// [FineTuneListResponseDataTrainingMethodTrainingMethodSft],
// [FineTuneListResponseDataTrainingMethodTrainingMethodDpo].
func (r FineTuneListResponseDataTrainingMethod) AsUnion() FineTuneListResponseDataTrainingMethodUnion {
	return r.union
}

// Method of training used
//
// Union satisfied by [FineTuneListResponseDataTrainingMethodTrainingMethodSft] or
// [FineTuneListResponseDataTrainingMethodTrainingMethodDpo].
type FineTuneListResponseDataTrainingMethodUnion interface {
	implementsFineTuneListResponseDataTrainingMethod()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneListResponseDataTrainingMethodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneListResponseDataTrainingMethodTrainingMethodSft{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneListResponseDataTrainingMethodTrainingMethodDpo{}),
		},
	)
}

type FineTuneListResponseDataTrainingMethodTrainingMethodSft struct {
	Method FineTuneListResponseDataTrainingMethodTrainingMethodSftMethod `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	JSON          fineTuneListResponseDataTrainingMethodTrainingMethodSftJSON               `json:"-"`
}

// fineTuneListResponseDataTrainingMethodTrainingMethodSftJSON contains the JSON
// metadata for the struct
// [FineTuneListResponseDataTrainingMethodTrainingMethodSft]
type fineTuneListResponseDataTrainingMethodTrainingMethodSftJSON struct {
	Method        apijson.Field
	TrainOnInputs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FineTuneListResponseDataTrainingMethodTrainingMethodSft) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseDataTrainingMethodTrainingMethodSftJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneListResponseDataTrainingMethodTrainingMethodSft) implementsFineTuneListResponseDataTrainingMethod() {
}

type FineTuneListResponseDataTrainingMethodTrainingMethodSftMethod string

const (
	FineTuneListResponseDataTrainingMethodTrainingMethodSftMethodSft FineTuneListResponseDataTrainingMethodTrainingMethodSftMethod = "sft"
)

func (r FineTuneListResponseDataTrainingMethodTrainingMethodSftMethod) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataTrainingMethodTrainingMethodSftMethodSft:
		return true
	}
	return false
}

// Whether to mask the user messages in conversational data or prompts in
// instruction data.
//
// Union satisfied by [shared.UnionBool] or
// [FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString].
type FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion interface {
	ImplementsFineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString("")),
		},
	)
}

type FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

func (r FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsStringAuto:
		return true
	}
	return false
}

func (r FineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsString) ImplementsFineTuneListResponseDataTrainingMethodTrainingMethodSftTrainOnInputsUnion() {
}

type FineTuneListResponseDataTrainingMethodTrainingMethodDpo struct {
	Method  FineTuneListResponseDataTrainingMethodTrainingMethodDpoMethod `json:"method,required"`
	DpoBeta float64                                                       `json:"dpo_beta"`
	JSON    fineTuneListResponseDataTrainingMethodTrainingMethodDpoJSON   `json:"-"`
}

// fineTuneListResponseDataTrainingMethodTrainingMethodDpoJSON contains the JSON
// metadata for the struct
// [FineTuneListResponseDataTrainingMethodTrainingMethodDpo]
type fineTuneListResponseDataTrainingMethodTrainingMethodDpoJSON struct {
	Method      apijson.Field
	DpoBeta     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneListResponseDataTrainingMethodTrainingMethodDpo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseDataTrainingMethodTrainingMethodDpoJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneListResponseDataTrainingMethodTrainingMethodDpo) implementsFineTuneListResponseDataTrainingMethod() {
}

type FineTuneListResponseDataTrainingMethodTrainingMethodDpoMethod string

const (
	FineTuneListResponseDataTrainingMethodTrainingMethodDpoMethodDpo FineTuneListResponseDataTrainingMethodTrainingMethodDpoMethod = "dpo"
)

func (r FineTuneListResponseDataTrainingMethodTrainingMethodDpoMethod) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataTrainingMethodTrainingMethodDpoMethodDpo:
		return true
	}
	return false
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
// Possible runtime types of the union are
// [FineTuneListResponseDataTrainingTypeFullTrainingType],
// [FineTuneListResponseDataTrainingTypeLoRaTrainingType].
func (r FineTuneListResponseDataTrainingType) AsUnion() FineTuneListResponseDataTrainingTypeUnion {
	return r.union
}

// Type of training used (full or LoRA)
//
// Union satisfied by [FineTuneListResponseDataTrainingTypeFullTrainingType] or
// [FineTuneListResponseDataTrainingTypeLoRaTrainingType].
type FineTuneListResponseDataTrainingTypeUnion interface {
	implementsFineTuneListResponseDataTrainingType()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneListResponseDataTrainingTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneListResponseDataTrainingTypeFullTrainingType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneListResponseDataTrainingTypeLoRaTrainingType{}),
		},
	)
}

type FineTuneListResponseDataTrainingTypeFullTrainingType struct {
	Type FineTuneListResponseDataTrainingTypeFullTrainingTypeType `json:"type,required"`
	JSON fineTuneListResponseDataTrainingTypeFullTrainingTypeJSON `json:"-"`
}

// fineTuneListResponseDataTrainingTypeFullTrainingTypeJSON contains the JSON
// metadata for the struct [FineTuneListResponseDataTrainingTypeFullTrainingType]
type fineTuneListResponseDataTrainingTypeFullTrainingTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneListResponseDataTrainingTypeFullTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseDataTrainingTypeFullTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneListResponseDataTrainingTypeFullTrainingType) implementsFineTuneListResponseDataTrainingType() {
}

type FineTuneListResponseDataTrainingTypeFullTrainingTypeType string

const (
	FineTuneListResponseDataTrainingTypeFullTrainingTypeTypeFull FineTuneListResponseDataTrainingTypeFullTrainingTypeType = "Full"
)

func (r FineTuneListResponseDataTrainingTypeFullTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataTrainingTypeFullTrainingTypeTypeFull:
		return true
	}
	return false
}

type FineTuneListResponseDataTrainingTypeLoRaTrainingType struct {
	LoraAlpha            int64                                                    `json:"lora_alpha,required"`
	LoraR                int64                                                    `json:"lora_r,required"`
	Type                 FineTuneListResponseDataTrainingTypeLoRaTrainingTypeType `json:"type,required"`
	LoraDropout          float64                                                  `json:"lora_dropout"`
	LoraTrainableModules string                                                   `json:"lora_trainable_modules"`
	JSON                 fineTuneListResponseDataTrainingTypeLoRaTrainingTypeJSON `json:"-"`
}

// fineTuneListResponseDataTrainingTypeLoRaTrainingTypeJSON contains the JSON
// metadata for the struct [FineTuneListResponseDataTrainingTypeLoRaTrainingType]
type fineTuneListResponseDataTrainingTypeLoRaTrainingTypeJSON struct {
	LoraAlpha            apijson.Field
	LoraR                apijson.Field
	Type                 apijson.Field
	LoraDropout          apijson.Field
	LoraTrainableModules apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FineTuneListResponseDataTrainingTypeLoRaTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneListResponseDataTrainingTypeLoRaTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneListResponseDataTrainingTypeLoRaTrainingType) implementsFineTuneListResponseDataTrainingType() {
}

type FineTuneListResponseDataTrainingTypeLoRaTrainingTypeType string

const (
	FineTuneListResponseDataTrainingTypeLoRaTrainingTypeTypeLora FineTuneListResponseDataTrainingTypeLoRaTrainingTypeType = "Lora"
)

func (r FineTuneListResponseDataTrainingTypeLoRaTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneListResponseDataTrainingTypeLoRaTrainingTypeTypeLora:
		return true
	}
	return false
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
	Events []FineTuneCancelResponseEvent `json:"events"`
	// Checkpoint used to continue training
	FromCheckpoint string `json:"from_checkpoint"`
	// Learning rate used for training
	LearningRate float64 `json:"learning_rate"`
	// Learning rate scheduler configuration
	LrScheduler FineTuneCancelResponseLrScheduler `json:"lr_scheduler"`
	// Maximum gradient norm for clipping
	MaxGradNorm float64 `json:"max_grad_norm"`
	// Base model used for fine-tuning
	Model string `json:"model"`
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

type FineTuneCancelResponseEvent struct {
	CheckpointPath string                             `json:"checkpoint_path,required"`
	CreatedAt      string                             `json:"created_at,required"`
	Hash           string                             `json:"hash,required"`
	Message        string                             `json:"message,required"`
	ModelPath      string                             `json:"model_path,required"`
	Object         FineTuneCancelResponseEventsObject `json:"object,required"`
	ParamCount     int64                              `json:"param_count,required"`
	Step           int64                              `json:"step,required"`
	TokenCount     int64                              `json:"token_count,required"`
	TotalSteps     int64                              `json:"total_steps,required"`
	TrainingOffset int64                              `json:"training_offset,required"`
	Type           FineTuneCancelResponseEventsType   `json:"type,required"`
	WandbURL       string                             `json:"wandb_url,required"`
	Level          FineTuneCancelResponseEventsLevel  `json:"level,nullable"`
	JSON           fineTuneCancelResponseEventJSON    `json:"-"`
}

// fineTuneCancelResponseEventJSON contains the JSON metadata for the struct
// [FineTuneCancelResponseEvent]
type fineTuneCancelResponseEventJSON struct {
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

func (r *FineTuneCancelResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneCancelResponseEventJSON) RawJSON() string {
	return r.raw
}

type FineTuneCancelResponseEventsObject string

const (
	FineTuneCancelResponseEventsObjectFineTuneEvent FineTuneCancelResponseEventsObject = "fine-tune-event"
)

func (r FineTuneCancelResponseEventsObject) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseEventsObjectFineTuneEvent:
		return true
	}
	return false
}

type FineTuneCancelResponseEventsType string

const (
	FineTuneCancelResponseEventsTypeJobPending                     FineTuneCancelResponseEventsType = "job_pending"
	FineTuneCancelResponseEventsTypeJobStart                       FineTuneCancelResponseEventsType = "job_start"
	FineTuneCancelResponseEventsTypeJobStopped                     FineTuneCancelResponseEventsType = "job_stopped"
	FineTuneCancelResponseEventsTypeModelDownloading               FineTuneCancelResponseEventsType = "model_downloading"
	FineTuneCancelResponseEventsTypeModelDownloadComplete          FineTuneCancelResponseEventsType = "model_download_complete"
	FineTuneCancelResponseEventsTypeTrainingDataDownloading        FineTuneCancelResponseEventsType = "training_data_downloading"
	FineTuneCancelResponseEventsTypeTrainingDataDownloadComplete   FineTuneCancelResponseEventsType = "training_data_download_complete"
	FineTuneCancelResponseEventsTypeValidationDataDownloading      FineTuneCancelResponseEventsType = "validation_data_downloading"
	FineTuneCancelResponseEventsTypeValidationDataDownloadComplete FineTuneCancelResponseEventsType = "validation_data_download_complete"
	FineTuneCancelResponseEventsTypeWandbInit                      FineTuneCancelResponseEventsType = "wandb_init"
	FineTuneCancelResponseEventsTypeTrainingStart                  FineTuneCancelResponseEventsType = "training_start"
	FineTuneCancelResponseEventsTypeCheckpointSave                 FineTuneCancelResponseEventsType = "checkpoint_save"
	FineTuneCancelResponseEventsTypeBillingLimit                   FineTuneCancelResponseEventsType = "billing_limit"
	FineTuneCancelResponseEventsTypeEpochComplete                  FineTuneCancelResponseEventsType = "epoch_complete"
	FineTuneCancelResponseEventsTypeTrainingComplete               FineTuneCancelResponseEventsType = "training_complete"
	FineTuneCancelResponseEventsTypeModelCompressing               FineTuneCancelResponseEventsType = "model_compressing"
	FineTuneCancelResponseEventsTypeModelCompressionComplete       FineTuneCancelResponseEventsType = "model_compression_complete"
	FineTuneCancelResponseEventsTypeModelUploading                 FineTuneCancelResponseEventsType = "model_uploading"
	FineTuneCancelResponseEventsTypeModelUploadComplete            FineTuneCancelResponseEventsType = "model_upload_complete"
	FineTuneCancelResponseEventsTypeJobComplete                    FineTuneCancelResponseEventsType = "job_complete"
	FineTuneCancelResponseEventsTypeJobError                       FineTuneCancelResponseEventsType = "job_error"
	FineTuneCancelResponseEventsTypeCancelRequested                FineTuneCancelResponseEventsType = "cancel_requested"
	FineTuneCancelResponseEventsTypeJobRestarted                   FineTuneCancelResponseEventsType = "job_restarted"
	FineTuneCancelResponseEventsTypeRefund                         FineTuneCancelResponseEventsType = "refund"
	FineTuneCancelResponseEventsTypeWarning                        FineTuneCancelResponseEventsType = "warning"
)

func (r FineTuneCancelResponseEventsType) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseEventsTypeJobPending, FineTuneCancelResponseEventsTypeJobStart, FineTuneCancelResponseEventsTypeJobStopped, FineTuneCancelResponseEventsTypeModelDownloading, FineTuneCancelResponseEventsTypeModelDownloadComplete, FineTuneCancelResponseEventsTypeTrainingDataDownloading, FineTuneCancelResponseEventsTypeTrainingDataDownloadComplete, FineTuneCancelResponseEventsTypeValidationDataDownloading, FineTuneCancelResponseEventsTypeValidationDataDownloadComplete, FineTuneCancelResponseEventsTypeWandbInit, FineTuneCancelResponseEventsTypeTrainingStart, FineTuneCancelResponseEventsTypeCheckpointSave, FineTuneCancelResponseEventsTypeBillingLimit, FineTuneCancelResponseEventsTypeEpochComplete, FineTuneCancelResponseEventsTypeTrainingComplete, FineTuneCancelResponseEventsTypeModelCompressing, FineTuneCancelResponseEventsTypeModelCompressionComplete, FineTuneCancelResponseEventsTypeModelUploading, FineTuneCancelResponseEventsTypeModelUploadComplete, FineTuneCancelResponseEventsTypeJobComplete, FineTuneCancelResponseEventsTypeJobError, FineTuneCancelResponseEventsTypeCancelRequested, FineTuneCancelResponseEventsTypeJobRestarted, FineTuneCancelResponseEventsTypeRefund, FineTuneCancelResponseEventsTypeWarning:
		return true
	}
	return false
}

type FineTuneCancelResponseEventsLevel string

const (
	FineTuneCancelResponseEventsLevelInfo           FineTuneCancelResponseEventsLevel = "info"
	FineTuneCancelResponseEventsLevelWarning        FineTuneCancelResponseEventsLevel = "warning"
	FineTuneCancelResponseEventsLevelError          FineTuneCancelResponseEventsLevel = "error"
	FineTuneCancelResponseEventsLevelLegacyInfo     FineTuneCancelResponseEventsLevel = "legacy_info"
	FineTuneCancelResponseEventsLevelLegacyIwarning FineTuneCancelResponseEventsLevel = "legacy_iwarning"
	FineTuneCancelResponseEventsLevelLegacyIerror   FineTuneCancelResponseEventsLevel = "legacy_ierror"
)

func (r FineTuneCancelResponseEventsLevel) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseEventsLevelInfo, FineTuneCancelResponseEventsLevelWarning, FineTuneCancelResponseEventsLevelError, FineTuneCancelResponseEventsLevelLegacyInfo, FineTuneCancelResponseEventsLevelLegacyIwarning, FineTuneCancelResponseEventsLevelLegacyIerror:
		return true
	}
	return false
}

// Learning rate scheduler configuration
type FineTuneCancelResponseLrScheduler struct {
	LrSchedulerType FineTuneCancelResponseLrSchedulerLrSchedulerType `json:"lr_scheduler_type,required"`
	LrSchedulerArgs FineTuneCancelResponseLrSchedulerLrSchedulerArgs `json:"lr_scheduler_args"`
	JSON            fineTuneCancelResponseLrSchedulerJSON            `json:"-"`
}

// fineTuneCancelResponseLrSchedulerJSON contains the JSON metadata for the struct
// [FineTuneCancelResponseLrScheduler]
type fineTuneCancelResponseLrSchedulerJSON struct {
	LrSchedulerType apijson.Field
	LrSchedulerArgs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FineTuneCancelResponseLrScheduler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneCancelResponseLrSchedulerJSON) RawJSON() string {
	return r.raw
}

type FineTuneCancelResponseLrSchedulerLrSchedulerType string

const (
	FineTuneCancelResponseLrSchedulerLrSchedulerTypeLinear FineTuneCancelResponseLrSchedulerLrSchedulerType = "linear"
	FineTuneCancelResponseLrSchedulerLrSchedulerTypeCosine FineTuneCancelResponseLrSchedulerLrSchedulerType = "cosine"
)

func (r FineTuneCancelResponseLrSchedulerLrSchedulerType) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseLrSchedulerLrSchedulerTypeLinear, FineTuneCancelResponseLrSchedulerLrSchedulerTypeCosine:
		return true
	}
	return false
}

type FineTuneCancelResponseLrSchedulerLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64                                              `json:"num_cycles"`
	JSON      fineTuneCancelResponseLrSchedulerLrSchedulerArgsJSON `json:"-"`
	union     FineTuneCancelResponseLrSchedulerLrSchedulerArgsUnion
}

// fineTuneCancelResponseLrSchedulerLrSchedulerArgsJSON contains the JSON metadata
// for the struct [FineTuneCancelResponseLrSchedulerLrSchedulerArgs]
type fineTuneCancelResponseLrSchedulerLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r fineTuneCancelResponseLrSchedulerLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r *FineTuneCancelResponseLrSchedulerLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	*r = FineTuneCancelResponseLrSchedulerLrSchedulerArgs{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FineTuneCancelResponseLrSchedulerLrSchedulerArgsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs],
// [FineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
func (r FineTuneCancelResponseLrSchedulerLrSchedulerArgs) AsUnion() FineTuneCancelResponseLrSchedulerLrSchedulerArgsUnion {
	return r.union
}

// Union satisfied by
// [FineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs] or
// [FineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs].
type FineTuneCancelResponseLrSchedulerLrSchedulerArgsUnion interface {
	implementsFineTuneCancelResponseLrSchedulerLrSchedulerArgs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneCancelResponseLrSchedulerLrSchedulerArgsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs{}),
		},
	)
}

type FineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64                                                                   `json:"min_lr_ratio"`
	JSON       fineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON `json:"-"`
}

// fineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON
// contains the JSON metadata for the struct
// [FineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs]
type fineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneCancelResponseLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) implementsFineTuneCancelResponseLrSchedulerLrSchedulerArgs() {
}

type FineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio float64 `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles float64                                                                   `json:"num_cycles"`
	JSON      fineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON `json:"-"`
}

// fineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON
// contains the JSON metadata for the struct
// [FineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs]
type fineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON struct {
	MinLrRatio  apijson.Field
	NumCycles   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgsJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneCancelResponseLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) implementsFineTuneCancelResponseLrSchedulerLrSchedulerArgs() {
}

// Method of training used
type FineTuneCancelResponseTrainingMethod struct {
	Method  FineTuneCancelResponseTrainingMethodMethod `json:"method,required"`
	DpoBeta float64                                    `json:"dpo_beta"`
	// This field can have the runtime type of
	// [FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion].
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
// Possible runtime types of the union are
// [FineTuneCancelResponseTrainingMethodTrainingMethodSft],
// [FineTuneCancelResponseTrainingMethodTrainingMethodDpo].
func (r FineTuneCancelResponseTrainingMethod) AsUnion() FineTuneCancelResponseTrainingMethodUnion {
	return r.union
}

// Method of training used
//
// Union satisfied by [FineTuneCancelResponseTrainingMethodTrainingMethodSft] or
// [FineTuneCancelResponseTrainingMethodTrainingMethodDpo].
type FineTuneCancelResponseTrainingMethodUnion interface {
	implementsFineTuneCancelResponseTrainingMethod()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneCancelResponseTrainingMethodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneCancelResponseTrainingMethodTrainingMethodSft{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneCancelResponseTrainingMethodTrainingMethodDpo{}),
		},
	)
}

type FineTuneCancelResponseTrainingMethodTrainingMethodSft struct {
	Method FineTuneCancelResponseTrainingMethodTrainingMethodSftMethod `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion `json:"train_on_inputs,required"`
	JSON          fineTuneCancelResponseTrainingMethodTrainingMethodSftJSON               `json:"-"`
}

// fineTuneCancelResponseTrainingMethodTrainingMethodSftJSON contains the JSON
// metadata for the struct [FineTuneCancelResponseTrainingMethodTrainingMethodSft]
type fineTuneCancelResponseTrainingMethodTrainingMethodSftJSON struct {
	Method        apijson.Field
	TrainOnInputs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FineTuneCancelResponseTrainingMethodTrainingMethodSft) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneCancelResponseTrainingMethodTrainingMethodSftJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneCancelResponseTrainingMethodTrainingMethodSft) implementsFineTuneCancelResponseTrainingMethod() {
}

type FineTuneCancelResponseTrainingMethodTrainingMethodSftMethod string

const (
	FineTuneCancelResponseTrainingMethodTrainingMethodSftMethodSft FineTuneCancelResponseTrainingMethodTrainingMethodSftMethod = "sft"
)

func (r FineTuneCancelResponseTrainingMethodTrainingMethodSftMethod) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseTrainingMethodTrainingMethodSftMethodSft:
		return true
	}
	return false
}

// Whether to mask the user messages in conversational data or prompts in
// instruction data.
//
// Union satisfied by [shared.UnionBool] or
// [FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString].
type FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion interface {
	ImplementsFineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString("")),
		},
	)
}

type FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

func (r FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsStringAuto:
		return true
	}
	return false
}

func (r FineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsString) ImplementsFineTuneCancelResponseTrainingMethodTrainingMethodSftTrainOnInputsUnion() {
}

type FineTuneCancelResponseTrainingMethodTrainingMethodDpo struct {
	Method  FineTuneCancelResponseTrainingMethodTrainingMethodDpoMethod `json:"method,required"`
	DpoBeta float64                                                     `json:"dpo_beta"`
	JSON    fineTuneCancelResponseTrainingMethodTrainingMethodDpoJSON   `json:"-"`
}

// fineTuneCancelResponseTrainingMethodTrainingMethodDpoJSON contains the JSON
// metadata for the struct [FineTuneCancelResponseTrainingMethodTrainingMethodDpo]
type fineTuneCancelResponseTrainingMethodTrainingMethodDpoJSON struct {
	Method      apijson.Field
	DpoBeta     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneCancelResponseTrainingMethodTrainingMethodDpo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneCancelResponseTrainingMethodTrainingMethodDpoJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneCancelResponseTrainingMethodTrainingMethodDpo) implementsFineTuneCancelResponseTrainingMethod() {
}

type FineTuneCancelResponseTrainingMethodTrainingMethodDpoMethod string

const (
	FineTuneCancelResponseTrainingMethodTrainingMethodDpoMethodDpo FineTuneCancelResponseTrainingMethodTrainingMethodDpoMethod = "dpo"
)

func (r FineTuneCancelResponseTrainingMethodTrainingMethodDpoMethod) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseTrainingMethodTrainingMethodDpoMethodDpo:
		return true
	}
	return false
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
// Possible runtime types of the union are
// [FineTuneCancelResponseTrainingTypeFullTrainingType],
// [FineTuneCancelResponseTrainingTypeLoRaTrainingType].
func (r FineTuneCancelResponseTrainingType) AsUnion() FineTuneCancelResponseTrainingTypeUnion {
	return r.union
}

// Type of training used (full or LoRA)
//
// Union satisfied by [FineTuneCancelResponseTrainingTypeFullTrainingType] or
// [FineTuneCancelResponseTrainingTypeLoRaTrainingType].
type FineTuneCancelResponseTrainingTypeUnion interface {
	implementsFineTuneCancelResponseTrainingType()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FineTuneCancelResponseTrainingTypeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneCancelResponseTrainingTypeFullTrainingType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FineTuneCancelResponseTrainingTypeLoRaTrainingType{}),
		},
	)
}

type FineTuneCancelResponseTrainingTypeFullTrainingType struct {
	Type FineTuneCancelResponseTrainingTypeFullTrainingTypeType `json:"type,required"`
	JSON fineTuneCancelResponseTrainingTypeFullTrainingTypeJSON `json:"-"`
}

// fineTuneCancelResponseTrainingTypeFullTrainingTypeJSON contains the JSON
// metadata for the struct [FineTuneCancelResponseTrainingTypeFullTrainingType]
type fineTuneCancelResponseTrainingTypeFullTrainingTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FineTuneCancelResponseTrainingTypeFullTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneCancelResponseTrainingTypeFullTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneCancelResponseTrainingTypeFullTrainingType) implementsFineTuneCancelResponseTrainingType() {
}

type FineTuneCancelResponseTrainingTypeFullTrainingTypeType string

const (
	FineTuneCancelResponseTrainingTypeFullTrainingTypeTypeFull FineTuneCancelResponseTrainingTypeFullTrainingTypeType = "Full"
)

func (r FineTuneCancelResponseTrainingTypeFullTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseTrainingTypeFullTrainingTypeTypeFull:
		return true
	}
	return false
}

type FineTuneCancelResponseTrainingTypeLoRaTrainingType struct {
	LoraAlpha            int64                                                  `json:"lora_alpha,required"`
	LoraR                int64                                                  `json:"lora_r,required"`
	Type                 FineTuneCancelResponseTrainingTypeLoRaTrainingTypeType `json:"type,required"`
	LoraDropout          float64                                                `json:"lora_dropout"`
	LoraTrainableModules string                                                 `json:"lora_trainable_modules"`
	JSON                 fineTuneCancelResponseTrainingTypeLoRaTrainingTypeJSON `json:"-"`
}

// fineTuneCancelResponseTrainingTypeLoRaTrainingTypeJSON contains the JSON
// metadata for the struct [FineTuneCancelResponseTrainingTypeLoRaTrainingType]
type fineTuneCancelResponseTrainingTypeLoRaTrainingTypeJSON struct {
	LoraAlpha            apijson.Field
	LoraR                apijson.Field
	Type                 apijson.Field
	LoraDropout          apijson.Field
	LoraTrainableModules apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FineTuneCancelResponseTrainingTypeLoRaTrainingType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fineTuneCancelResponseTrainingTypeLoRaTrainingTypeJSON) RawJSON() string {
	return r.raw
}

func (r FineTuneCancelResponseTrainingTypeLoRaTrainingType) implementsFineTuneCancelResponseTrainingType() {
}

type FineTuneCancelResponseTrainingTypeLoRaTrainingTypeType string

const (
	FineTuneCancelResponseTrainingTypeLoRaTrainingTypeTypeLora FineTuneCancelResponseTrainingTypeLoRaTrainingTypeType = "Lora"
)

func (r FineTuneCancelResponseTrainingTypeLoRaTrainingTypeType) IsKnown() bool {
	switch r {
	case FineTuneCancelResponseTrainingTypeLoRaTrainingTypeTypeLora:
		return true
	}
	return false
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
	LrScheduler param.Field[FineTuneNewParamsLrScheduler] `json:"lr_scheduler"`
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

// The learning rate scheduler to use. It specifies how the learning rate is
// adjusted during training.
type FineTuneNewParamsLrScheduler struct {
	LrSchedulerType param.Field[FineTuneNewParamsLrSchedulerLrSchedulerType]      `json:"lr_scheduler_type,required"`
	LrSchedulerArgs param.Field[FineTuneNewParamsLrSchedulerLrSchedulerArgsUnion] `json:"lr_scheduler_args"`
}

func (r FineTuneNewParamsLrScheduler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FineTuneNewParamsLrSchedulerLrSchedulerType string

const (
	FineTuneNewParamsLrSchedulerLrSchedulerTypeLinear FineTuneNewParamsLrSchedulerLrSchedulerType = "linear"
	FineTuneNewParamsLrSchedulerLrSchedulerTypeCosine FineTuneNewParamsLrSchedulerLrSchedulerType = "cosine"
)

func (r FineTuneNewParamsLrSchedulerLrSchedulerType) IsKnown() bool {
	switch r {
	case FineTuneNewParamsLrSchedulerLrSchedulerTypeLinear, FineTuneNewParamsLrSchedulerLrSchedulerTypeCosine:
		return true
	}
	return false
}

type FineTuneNewParamsLrSchedulerLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio param.Field[float64] `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles param.Field[float64] `json:"num_cycles"`
}

func (r FineTuneNewParamsLrSchedulerLrSchedulerArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FineTuneNewParamsLrSchedulerLrSchedulerArgs) implementsFineTuneNewParamsLrSchedulerLrSchedulerArgsUnion() {
}

// Satisfied by [FineTuneNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs],
// [FineTuneNewParamsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs],
// [FineTuneNewParamsLrSchedulerLrSchedulerArgs].
type FineTuneNewParamsLrSchedulerLrSchedulerArgsUnion interface {
	implementsFineTuneNewParamsLrSchedulerLrSchedulerArgsUnion()
}

type FineTuneNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio param.Field[float64] `json:"min_lr_ratio"`
}

func (r FineTuneNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FineTuneNewParamsLrSchedulerLrSchedulerArgsLinearLrSchedulerArgs) implementsFineTuneNewParamsLrSchedulerLrSchedulerArgsUnion() {
}

type FineTuneNewParamsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs struct {
	// The ratio of the final learning rate to the peak learning rate
	MinLrRatio param.Field[float64] `json:"min_lr_ratio"`
	// Number or fraction of cycles for the cosine learning rate scheduler
	NumCycles param.Field[float64] `json:"num_cycles"`
}

func (r FineTuneNewParamsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FineTuneNewParamsLrSchedulerLrSchedulerArgsCosineLrSchedulerArgs) implementsFineTuneNewParamsLrSchedulerLrSchedulerArgsUnion() {
}

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
// Satisfied by [FineTuneNewParamsTrainingMethodTrainingMethodSft],
// [FineTuneNewParamsTrainingMethodTrainingMethodDpo],
// [FineTuneNewParamsTrainingMethod].
type FineTuneNewParamsTrainingMethodUnion interface {
	implementsFineTuneNewParamsTrainingMethodUnion()
}

type FineTuneNewParamsTrainingMethodTrainingMethodSft struct {
	Method param.Field[FineTuneNewParamsTrainingMethodTrainingMethodSftMethod] `json:"method,required"`
	// Whether to mask the user messages in conversational data or prompts in
	// instruction data.
	TrainOnInputs param.Field[FineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion] `json:"train_on_inputs,required"`
}

func (r FineTuneNewParamsTrainingMethodTrainingMethodSft) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FineTuneNewParamsTrainingMethodTrainingMethodSft) implementsFineTuneNewParamsTrainingMethodUnion() {
}

type FineTuneNewParamsTrainingMethodTrainingMethodSftMethod string

const (
	FineTuneNewParamsTrainingMethodTrainingMethodSftMethodSft FineTuneNewParamsTrainingMethodTrainingMethodSftMethod = "sft"
)

func (r FineTuneNewParamsTrainingMethodTrainingMethodSftMethod) IsKnown() bool {
	switch r {
	case FineTuneNewParamsTrainingMethodTrainingMethodSftMethodSft:
		return true
	}
	return false
}

// Whether to mask the user messages in conversational data or prompts in
// instruction data.
//
// Satisfied by [shared.UnionBool],
// [FineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsString].
type FineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion interface {
	ImplementsFineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion()
}

type FineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsString string

const (
	FineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsStringAuto FineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsString = "auto"
)

func (r FineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsString) IsKnown() bool {
	switch r {
	case FineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsStringAuto:
		return true
	}
	return false
}

func (r FineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsString) ImplementsFineTuneNewParamsTrainingMethodTrainingMethodSftTrainOnInputsUnion() {
}

type FineTuneNewParamsTrainingMethodTrainingMethodDpo struct {
	Method  param.Field[FineTuneNewParamsTrainingMethodTrainingMethodDpoMethod] `json:"method,required"`
	DpoBeta param.Field[float64]                                                `json:"dpo_beta"`
}

func (r FineTuneNewParamsTrainingMethodTrainingMethodDpo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FineTuneNewParamsTrainingMethodTrainingMethodDpo) implementsFineTuneNewParamsTrainingMethodUnion() {
}

type FineTuneNewParamsTrainingMethodTrainingMethodDpoMethod string

const (
	FineTuneNewParamsTrainingMethodTrainingMethodDpoMethodDpo FineTuneNewParamsTrainingMethodTrainingMethodDpoMethod = "dpo"
)

func (r FineTuneNewParamsTrainingMethodTrainingMethodDpoMethod) IsKnown() bool {
	switch r {
	case FineTuneNewParamsTrainingMethodTrainingMethodDpoMethodDpo:
		return true
	}
	return false
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
