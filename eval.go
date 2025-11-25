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
	"github.com/togethercomputer/together-go/internal/paramutil"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// EvalService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalService] method instead.
type EvalService struct {
	Options []option.RequestOption
}

// NewEvalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvalService(opts ...option.RequestOption) (r EvalService) {
	r = EvalService{}
	r.Options = opts
	return
}

// Create an evaluation job
func (r *EvalService) New(ctx context.Context, body EvalNewParams, opts ...option.RequestOption) (res *EvalNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "evaluation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get evaluation job details
func (r *EvalService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EvaluationJob, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("evaluation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all evaluation jobs
func (r *EvalService) List(ctx context.Context, query EvalListParams, opts ...option.RequestOption) (res *[]EvaluationJob, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "evaluation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get evaluation job status and results
func (r *EvalService) Status(ctx context.Context, id string, opts ...option.RequestOption) (res *EvalStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("evaluation/%s/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EvaluationJob struct {
	// When the job was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ID of the job owner (admin only)
	OwnerID string `json:"owner_id"`
	// The parameters used for this evaluation
	Parameters map[string]any `json:"parameters"`
	// Results of the evaluation (when completed)
	Results EvaluationJobResultsUnion `json:"results,nullable"`
	// Current status of the job
	//
	// Any of "pending", "queued", "running", "completed", "error", "user_error".
	Status EvaluationJobStatus `json:"status"`
	// History of status updates (admin only)
	StatusUpdates []EvaluationJobStatusUpdate `json:"status_updates"`
	// The type of evaluation
	//
	// Any of "classify", "score", "compare".
	Type EvaluationJobType `json:"type"`
	// When the job was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The evaluation job ID
	WorkflowID string `json:"workflow_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		OwnerID       respjson.Field
		Parameters    respjson.Field
		Results       respjson.Field
		Status        respjson.Field
		StatusUpdates respjson.Field
		Type          respjson.Field
		UpdatedAt     respjson.Field
		WorkflowID    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationJob) RawJSON() string { return r.JSON.raw }
func (r *EvaluationJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvaluationJobResultsUnion contains all possible properties and values from
// [EvaluationJobResultsEvaluationClassifyResults],
// [EvaluationJobResultsEvaluationScoreResults],
// [EvaluationJobResultsEvaluationCompareResults], [EvaluationJobResultsError].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EvaluationJobResultsUnion struct {
	GenerationFailCount float64 `json:"generation_fail_count"`
	// This field is from variant [EvaluationJobResultsEvaluationClassifyResults].
	InvalidLabelCount float64 `json:"invalid_label_count"`
	JudgeFailCount    float64 `json:"judge_fail_count"`
	// This field is from variant [EvaluationJobResultsEvaluationClassifyResults].
	LabelCounts string `json:"label_counts"`
	// This field is from variant [EvaluationJobResultsEvaluationClassifyResults].
	PassPercentage float64 `json:"pass_percentage"`
	ResultFileID   string  `json:"result_file_id"`
	// This field is from variant [EvaluationJobResultsEvaluationScoreResults].
	AggregatedScores EvaluationJobResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// This field is from variant [EvaluationJobResultsEvaluationScoreResults].
	FailedSamples float64 `json:"failed_samples"`
	// This field is from variant [EvaluationJobResultsEvaluationScoreResults].
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// This field is from variant [EvaluationJobResultsEvaluationCompareResults].
	AWins int64 `json:"A_wins"`
	// This field is from variant [EvaluationJobResultsEvaluationCompareResults].
	BWins int64 `json:"B_wins"`
	// This field is from variant [EvaluationJobResultsEvaluationCompareResults].
	NumSamples int64 `json:"num_samples"`
	// This field is from variant [EvaluationJobResultsEvaluationCompareResults].
	Ties int64 `json:"Ties"`
	// This field is from variant [EvaluationJobResultsError].
	Error string `json:"error"`
	JSON  struct {
		GenerationFailCount respjson.Field
		InvalidLabelCount   respjson.Field
		JudgeFailCount      respjson.Field
		LabelCounts         respjson.Field
		PassPercentage      respjson.Field
		ResultFileID        respjson.Field
		AggregatedScores    respjson.Field
		FailedSamples       respjson.Field
		InvalidScoreCount   respjson.Field
		AWins               respjson.Field
		BWins               respjson.Field
		NumSamples          respjson.Field
		Ties                respjson.Field
		Error               respjson.Field
		raw                 string
	} `json:"-"`
}

func (u EvaluationJobResultsUnion) AsEvaluationJobResultsEvaluationClassifyResults() (v EvaluationJobResultsEvaluationClassifyResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluationJobResultsUnion) AsEvaluationJobResultsEvaluationScoreResults() (v EvaluationJobResultsEvaluationScoreResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluationJobResultsUnion) AsEvaluationJobResultsEvaluationCompareResults() (v EvaluationJobResultsEvaluationCompareResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvaluationJobResultsUnion) AsEvaluationJobResultsError() (v EvaluationJobResultsError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvaluationJobResultsUnion) RawJSON() string { return u.JSON.raw }

func (r *EvaluationJobResultsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluationJobResultsEvaluationClassifyResults struct {
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// Number of invalid labels
	InvalidLabelCount float64 `json:"invalid_label_count,nullable"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// JSON string representing label counts
	LabelCounts string `json:"label_counts"`
	// Pecentage of pass labels.
	PassPercentage float64 `json:"pass_percentage,nullable"`
	// Data File ID
	ResultFileID string `json:"result_file_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GenerationFailCount respjson.Field
		InvalidLabelCount   respjson.Field
		JudgeFailCount      respjson.Field
		LabelCounts         respjson.Field
		PassPercentage      respjson.Field
		ResultFileID        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationJobResultsEvaluationClassifyResults) RawJSON() string { return r.JSON.raw }
func (r *EvaluationJobResultsEvaluationClassifyResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluationJobResultsEvaluationScoreResults struct {
	AggregatedScores EvaluationJobResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// number of failed samples generated from model
	FailedSamples float64 `json:"failed_samples"`
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// number of invalid scores generated from model
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// Data File ID
	ResultFileID string `json:"result_file_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregatedScores    respjson.Field
		FailedSamples       respjson.Field
		GenerationFailCount respjson.Field
		InvalidScoreCount   respjson.Field
		JudgeFailCount      respjson.Field
		ResultFileID        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationJobResultsEvaluationScoreResults) RawJSON() string { return r.JSON.raw }
func (r *EvaluationJobResultsEvaluationScoreResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluationJobResultsEvaluationScoreResultsAggregatedScores struct {
	MeanScore      float64 `json:"mean_score"`
	PassPercentage float64 `json:"pass_percentage"`
	StdScore       float64 `json:"std_score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MeanScore      respjson.Field
		PassPercentage respjson.Field
		StdScore       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationJobResultsEvaluationScoreResultsAggregatedScores) RawJSON() string {
	return r.JSON.raw
}
func (r *EvaluationJobResultsEvaluationScoreResultsAggregatedScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluationJobResultsEvaluationCompareResults struct {
	// Number of times model A won
	AWins int64 `json:"A_wins"`
	// Number of times model B won
	BWins int64 `json:"B_wins"`
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// Total number of samples compared
	NumSamples int64 `json:"num_samples"`
	// Data File ID
	ResultFileID string `json:"result_file_id"`
	// Number of ties
	Ties int64 `json:"Ties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AWins               respjson.Field
		BWins               respjson.Field
		GenerationFailCount respjson.Field
		JudgeFailCount      respjson.Field
		NumSamples          respjson.Field
		ResultFileID        respjson.Field
		Ties                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationJobResultsEvaluationCompareResults) RawJSON() string { return r.JSON.raw }
func (r *EvaluationJobResultsEvaluationCompareResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluationJobResultsError struct {
	Error string `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationJobResultsError) RawJSON() string { return r.JSON.raw }
func (r *EvaluationJobResultsError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the job
type EvaluationJobStatus string

const (
	EvaluationJobStatusPending   EvaluationJobStatus = "pending"
	EvaluationJobStatusQueued    EvaluationJobStatus = "queued"
	EvaluationJobStatusRunning   EvaluationJobStatus = "running"
	EvaluationJobStatusCompleted EvaluationJobStatus = "completed"
	EvaluationJobStatusError     EvaluationJobStatus = "error"
	EvaluationJobStatusUserError EvaluationJobStatus = "user_error"
)

type EvaluationJobStatusUpdate struct {
	// Additional message for this update
	Message string `json:"message"`
	// The status at this update
	Status string `json:"status"`
	// When this update occurred
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationJobStatusUpdate) RawJSON() string { return r.JSON.raw }
func (r *EvaluationJobStatusUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of evaluation
type EvaluationJobType string

const (
	EvaluationJobTypeClassify EvaluationJobType = "classify"
	EvaluationJobTypeScore    EvaluationJobType = "score"
	EvaluationJobTypeCompare  EvaluationJobType = "compare"
)

type EvalNewResponse struct {
	// Initial status of the job
	//
	// Any of "pending".
	Status EvalNewResponseStatus `json:"status"`
	// The ID of the created evaluation job
	WorkflowID string `json:"workflow_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		WorkflowID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial status of the job
type EvalNewResponseStatus string

const (
	EvalNewResponseStatusPending EvalNewResponseStatus = "pending"
)

type EvalStatusResponse struct {
	// The results of the evaluation job
	Results EvalStatusResponseResultsUnion `json:"results"`
	// The status of the evaluation job
	//
	// Any of "completed", "error", "user_error", "running", "queued", "pending".
	Status EvalStatusResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvalStatusResponseResultsUnion contains all possible properties and values from
// [EvalStatusResponseResultsEvaluationClassifyResults],
// [EvalStatusResponseResultsEvaluationScoreResults],
// [EvalStatusResponseResultsEvaluationCompareResults].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EvalStatusResponseResultsUnion struct {
	GenerationFailCount float64 `json:"generation_fail_count"`
	// This field is from variant [EvalStatusResponseResultsEvaluationClassifyResults].
	InvalidLabelCount float64 `json:"invalid_label_count"`
	JudgeFailCount    float64 `json:"judge_fail_count"`
	// This field is from variant [EvalStatusResponseResultsEvaluationClassifyResults].
	LabelCounts string `json:"label_counts"`
	// This field is from variant [EvalStatusResponseResultsEvaluationClassifyResults].
	PassPercentage float64 `json:"pass_percentage"`
	ResultFileID   string  `json:"result_file_id"`
	// This field is from variant [EvalStatusResponseResultsEvaluationScoreResults].
	AggregatedScores EvalStatusResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// This field is from variant [EvalStatusResponseResultsEvaluationScoreResults].
	FailedSamples float64 `json:"failed_samples"`
	// This field is from variant [EvalStatusResponseResultsEvaluationScoreResults].
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// This field is from variant [EvalStatusResponseResultsEvaluationCompareResults].
	AWins int64 `json:"A_wins"`
	// This field is from variant [EvalStatusResponseResultsEvaluationCompareResults].
	BWins int64 `json:"B_wins"`
	// This field is from variant [EvalStatusResponseResultsEvaluationCompareResults].
	NumSamples int64 `json:"num_samples"`
	// This field is from variant [EvalStatusResponseResultsEvaluationCompareResults].
	Ties int64 `json:"Ties"`
	JSON struct {
		GenerationFailCount respjson.Field
		InvalidLabelCount   respjson.Field
		JudgeFailCount      respjson.Field
		LabelCounts         respjson.Field
		PassPercentage      respjson.Field
		ResultFileID        respjson.Field
		AggregatedScores    respjson.Field
		FailedSamples       respjson.Field
		InvalidScoreCount   respjson.Field
		AWins               respjson.Field
		BWins               respjson.Field
		NumSamples          respjson.Field
		Ties                respjson.Field
		raw                 string
	} `json:"-"`
}

func (u EvalStatusResponseResultsUnion) AsEvalStatusResponseResultsEvaluationClassifyResults() (v EvalStatusResponseResultsEvaluationClassifyResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalStatusResponseResultsUnion) AsEvalStatusResponseResultsEvaluationScoreResults() (v EvalStatusResponseResultsEvaluationScoreResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalStatusResponseResultsUnion) AsEvalStatusResponseResultsEvaluationCompareResults() (v EvalStatusResponseResultsEvaluationCompareResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvalStatusResponseResultsUnion) RawJSON() string { return u.JSON.raw }

func (r *EvalStatusResponseResultsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalStatusResponseResultsEvaluationClassifyResults struct {
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// Number of invalid labels
	InvalidLabelCount float64 `json:"invalid_label_count,nullable"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// JSON string representing label counts
	LabelCounts string `json:"label_counts"`
	// Pecentage of pass labels.
	PassPercentage float64 `json:"pass_percentage,nullable"`
	// Data File ID
	ResultFileID string `json:"result_file_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GenerationFailCount respjson.Field
		InvalidLabelCount   respjson.Field
		JudgeFailCount      respjson.Field
		LabelCounts         respjson.Field
		PassPercentage      respjson.Field
		ResultFileID        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalStatusResponseResultsEvaluationClassifyResults) RawJSON() string { return r.JSON.raw }
func (r *EvalStatusResponseResultsEvaluationClassifyResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalStatusResponseResultsEvaluationScoreResults struct {
	AggregatedScores EvalStatusResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// number of failed samples generated from model
	FailedSamples float64 `json:"failed_samples"`
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// number of invalid scores generated from model
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// Data File ID
	ResultFileID string `json:"result_file_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregatedScores    respjson.Field
		FailedSamples       respjson.Field
		GenerationFailCount respjson.Field
		InvalidScoreCount   respjson.Field
		JudgeFailCount      respjson.Field
		ResultFileID        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalStatusResponseResultsEvaluationScoreResults) RawJSON() string { return r.JSON.raw }
func (r *EvalStatusResponseResultsEvaluationScoreResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalStatusResponseResultsEvaluationScoreResultsAggregatedScores struct {
	MeanScore      float64 `json:"mean_score"`
	PassPercentage float64 `json:"pass_percentage"`
	StdScore       float64 `json:"std_score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MeanScore      respjson.Field
		PassPercentage respjson.Field
		StdScore       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalStatusResponseResultsEvaluationScoreResultsAggregatedScores) RawJSON() string {
	return r.JSON.raw
}
func (r *EvalStatusResponseResultsEvaluationScoreResultsAggregatedScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalStatusResponseResultsEvaluationCompareResults struct {
	// Number of times model A won
	AWins int64 `json:"A_wins"`
	// Number of times model B won
	BWins int64 `json:"B_wins"`
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// Total number of samples compared
	NumSamples int64 `json:"num_samples"`
	// Data File ID
	ResultFileID string `json:"result_file_id"`
	// Number of ties
	Ties int64 `json:"Ties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AWins               respjson.Field
		BWins               respjson.Field
		GenerationFailCount respjson.Field
		JudgeFailCount      respjson.Field
		NumSamples          respjson.Field
		ResultFileID        respjson.Field
		Ties                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalStatusResponseResultsEvaluationCompareResults) RawJSON() string { return r.JSON.raw }
func (r *EvalStatusResponseResultsEvaluationCompareResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the evaluation job
type EvalStatusResponseStatus string

const (
	EvalStatusResponseStatusCompleted EvalStatusResponseStatus = "completed"
	EvalStatusResponseStatusError     EvalStatusResponseStatus = "error"
	EvalStatusResponseStatusUserError EvalStatusResponseStatus = "user_error"
	EvalStatusResponseStatusRunning   EvalStatusResponseStatus = "running"
	EvalStatusResponseStatusQueued    EvalStatusResponseStatus = "queued"
	EvalStatusResponseStatusPending   EvalStatusResponseStatus = "pending"
)

type EvalNewParams struct {
	// Type-specific parameters for the evaluation
	Parameters EvalNewParamsParametersUnion `json:"parameters,omitzero,required"`
	// The type of evaluation to perform
	//
	// Any of "classify", "score", "compare".
	Type EvalNewParamsType `json:"type,omitzero,required"`
	paramObj
}

func (r EvalNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalNewParamsParametersUnion struct {
	OfEvalNewsParametersEvaluationClassifyParameters *EvalNewParamsParametersEvaluationClassifyParameters `json:",omitzero,inline"`
	OfEvalNewsParametersEvaluationScoreParameters    *EvalNewParamsParametersEvaluationScoreParameters    `json:",omitzero,inline"`
	OfEvalNewsParametersEvaluationCompareParameters  *EvalNewParamsParametersEvaluationCompareParameters  `json:",omitzero,inline"`
	paramUnion
}

func (u EvalNewParamsParametersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEvalNewsParametersEvaluationClassifyParameters, u.OfEvalNewsParametersEvaluationScoreParameters, u.OfEvalNewsParametersEvaluationCompareParameters)
}
func (u *EvalNewParamsParametersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalNewParamsParametersUnion) asAny() any {
	if !param.IsOmitted(u.OfEvalNewsParametersEvaluationClassifyParameters) {
		return u.OfEvalNewsParametersEvaluationClassifyParameters
	} else if !param.IsOmitted(u.OfEvalNewsParametersEvaluationScoreParameters) {
		return u.OfEvalNewsParametersEvaluationScoreParameters
	} else if !param.IsOmitted(u.OfEvalNewsParametersEvaluationCompareParameters) {
		return u.OfEvalNewsParametersEvaluationCompareParameters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParametersUnion) GetLabels() []string {
	if vt := u.OfEvalNewsParametersEvaluationClassifyParameters; vt != nil {
		return vt.Labels
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParametersUnion) GetPassLabels() []string {
	if vt := u.OfEvalNewsParametersEvaluationClassifyParameters; vt != nil {
		return vt.PassLabels
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParametersUnion) GetMaxScore() *float64 {
	if vt := u.OfEvalNewsParametersEvaluationScoreParameters; vt != nil {
		return &vt.MaxScore
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParametersUnion) GetMinScore() *float64 {
	if vt := u.OfEvalNewsParametersEvaluationScoreParameters; vt != nil {
		return &vt.MinScore
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParametersUnion) GetPassThreshold() *float64 {
	if vt := u.OfEvalNewsParametersEvaluationScoreParameters; vt != nil {
		return &vt.PassThreshold
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParametersUnion) GetModelA() *EvalNewParamsParametersEvaluationCompareParametersModelAUnion {
	if vt := u.OfEvalNewsParametersEvaluationCompareParameters; vt != nil {
		return &vt.ModelA
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParametersUnion) GetModelB() *EvalNewParamsParametersEvaluationCompareParametersModelBUnion {
	if vt := u.OfEvalNewsParametersEvaluationCompareParameters; vt != nil {
		return &vt.ModelB
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EvalNewParamsParametersUnion) GetInputDataFilePath() *string {
	if vt := u.OfEvalNewsParametersEvaluationClassifyParameters; vt != nil {
		return (*string)(&vt.InputDataFilePath)
	} else if vt := u.OfEvalNewsParametersEvaluationScoreParameters; vt != nil {
		return (*string)(&vt.InputDataFilePath)
	} else if vt := u.OfEvalNewsParametersEvaluationCompareParameters; vt != nil {
		return (*string)(&vt.InputDataFilePath)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u EvalNewParamsParametersUnion) GetJudge() (res evalNewParamsParametersUnionJudge) {
	if vt := u.OfEvalNewsParametersEvaluationClassifyParameters; vt != nil {
		res.any = &vt.Judge
	} else if vt := u.OfEvalNewsParametersEvaluationScoreParameters; vt != nil {
		res.any = &vt.Judge
	} else if vt := u.OfEvalNewsParametersEvaluationCompareParameters; vt != nil {
		res.any = &vt.Judge
	}
	return
}

// Can have the runtime types
// [*EvalNewParamsParametersEvaluationClassifyParametersJudge],
// [*EvalNewParamsParametersEvaluationScoreParametersJudge],
// [*EvalNewParamsParametersEvaluationCompareParametersJudge]
type evalNewParamsParametersUnionJudge struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *together.EvalNewParamsParametersEvaluationClassifyParametersJudge:
//	case *together.EvalNewParamsParametersEvaluationScoreParametersJudge:
//	case *together.EvalNewParamsParametersEvaluationCompareParametersJudge:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u evalNewParamsParametersUnionJudge) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionJudge) GetModel() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersJudge:
		return (*string)(&vt.Model)
	case *EvalNewParamsParametersEvaluationScoreParametersJudge:
		return (*string)(&vt.Model)
	case *EvalNewParamsParametersEvaluationCompareParametersJudge:
		return (*string)(&vt.Model)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionJudge) GetModelSource() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersJudge:
		return (*string)(&vt.ModelSource)
	case *EvalNewParamsParametersEvaluationScoreParametersJudge:
		return (*string)(&vt.ModelSource)
	case *EvalNewParamsParametersEvaluationCompareParametersJudge:
		return (*string)(&vt.ModelSource)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionJudge) GetSystemTemplate() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersJudge:
		return (*string)(&vt.SystemTemplate)
	case *EvalNewParamsParametersEvaluationScoreParametersJudge:
		return (*string)(&vt.SystemTemplate)
	case *EvalNewParamsParametersEvaluationCompareParametersJudge:
		return (*string)(&vt.SystemTemplate)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionJudge) GetExternalAPIToken() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersJudge:
		return paramutil.AddrIfPresent(vt.ExternalAPIToken)
	case *EvalNewParamsParametersEvaluationScoreParametersJudge:
		return paramutil.AddrIfPresent(vt.ExternalAPIToken)
	case *EvalNewParamsParametersEvaluationCompareParametersJudge:
		return paramutil.AddrIfPresent(vt.ExternalAPIToken)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionJudge) GetExternalBaseURL() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersJudge:
		return paramutil.AddrIfPresent(vt.ExternalBaseURL)
	case *EvalNewParamsParametersEvaluationScoreParametersJudge:
		return paramutil.AddrIfPresent(vt.ExternalBaseURL)
	case *EvalNewParamsParametersEvaluationCompareParametersJudge:
		return paramutil.AddrIfPresent(vt.ExternalBaseURL)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u EvalNewParamsParametersUnion) GetModelToEvaluate() (res evalNewParamsParametersUnionModelToEvaluate) {
	if vt := u.OfEvalNewsParametersEvaluationClassifyParameters; vt != nil {
		res.any = vt.ModelToEvaluate.asAny()
	} else if vt := u.OfEvalNewsParametersEvaluationScoreParameters; vt != nil {
		res.any = vt.ModelToEvaluate.asAny()
	}
	return
}

// Can have the runtime types [*string]
type evalNewParamsParametersUnionModelToEvaluate struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u evalNewParamsParametersUnionModelToEvaluate) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionModelToEvaluate) GetInputTemplate() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*string)(&vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest.InputTemplate)
		}
	case *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*string)(&vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest.InputTemplate)
		}
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionModelToEvaluate) GetMaxTokens() *int64 {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*int64)(&vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest.MaxTokens)
		}
	case *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*int64)(&vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest.MaxTokens)
		}
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionModelToEvaluate) GetModel() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*string)(&vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest.Model)
		}
	case *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*string)(&vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest.Model)
		}
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionModelToEvaluate) GetModelSource() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*string)(&vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest.ModelSource)
		}
	case *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*string)(&vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest.ModelSource)
		}
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionModelToEvaluate) GetSystemTemplate() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*string)(&vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest.SystemTemplate)
		}
	case *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*string)(&vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest.SystemTemplate)
		}
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionModelToEvaluate) GetTemperature() *float64 {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*float64)(&vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest.Temperature)
		}
	case *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest != nil {
			return (*float64)(&vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest.Temperature)
		}
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionModelToEvaluate) GetExternalAPIToken() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest != nil {
			return paramutil.AddrIfPresent((*string)(vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest.ExternalAPIToken))
		}
	case *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest != nil {
			return paramutil.AddrIfPresent((*string)(vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest.ExternalAPIToken))
		}
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u evalNewParamsParametersUnionModelToEvaluate) GetExternalBaseURL() *string {
	switch vt := u.any.(type) {
	case *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest != nil {
			return paramutil.AddrIfPresent((*string)(vt.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest.ExternalBaseURL))
		}
	case *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion:
		if vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest != nil {
			return paramutil.AddrIfPresent((*string)(vt.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest.ExternalBaseURL))
		}
	}
	return nil
}

// The properties InputDataFilePath, Judge, Labels, PassLabels are required.
type EvalNewParamsParametersEvaluationClassifyParameters struct {
	// Data file ID
	InputDataFilePath string                                                   `json:"input_data_file_path,required"`
	Judge             EvalNewParamsParametersEvaluationClassifyParametersJudge `json:"judge,omitzero,required"`
	// List of possible classification labels
	Labels []string `json:"labels,omitzero,required"`
	// List of labels that are considered passing
	PassLabels []string `json:"pass_labels,omitzero,required"`
	// Field name in the input data
	ModelToEvaluate EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion `json:"model_to_evaluate,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationClassifyParameters) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationClassifyParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationClassifyParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Model, ModelSource, SystemTemplate are required.
type EvalNewParamsParametersEvaluationClassifyParametersJudge struct {
	// Name of the judge model
	Model string `json:"model,required"`
	// Source of the judge model.
	//
	// Any of "serverless", "dedicated", "external".
	ModelSource string `json:"model_source,omitzero,required"`
	// System prompt template for the judge
	SystemTemplate string `json:"system_template,required"`
	// Bearer/API token for external judge models.
	ExternalAPIToken param.Opt[string] `json:"external_api_token,omitzero"`
	// Base URL for external judge models. Must be OpenAI-compatible base URL.
	ExternalBaseURL param.Opt[string] `json:"external_base_url,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationClassifyParametersJudge) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationClassifyParametersJudge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationClassifyParametersJudge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsParametersEvaluationClassifyParametersJudge](
		"model_source", "serverless", "dedicated", "external",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion struct {
	OfString                                                                              param.Opt[string]                                                                         `json:",omitzero,inline"`
	OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest `json:",omitzero,inline"`
	paramUnion
}

func (u EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest)
}
func (u *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest) {
		return u.OfEvalNewsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest
	}
	return nil
}

// The properties InputTemplate, MaxTokens, Model, ModelSource, SystemTemplate,
// Temperature are required.
type EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest struct {
	// Input prompt template
	InputTemplate string `json:"input_template,required"`
	// Maximum number of tokens to generate
	MaxTokens int64 `json:"max_tokens,required"`
	// Name of the model to evaluate
	Model string `json:"model,required"`
	// Source of the model.
	//
	// Any of "serverless", "dedicated", "external".
	ModelSource string `json:"model_source,omitzero,required"`
	// System prompt template
	SystemTemplate string `json:"system_template,required"`
	// Sampling temperature
	Temperature float64 `json:"temperature,required"`
	// Bearer/API token for external models.
	ExternalAPIToken param.Opt[string] `json:"external_api_token,omitzero"`
	// Base URL for external models. Must be OpenAI-compatible base URL
	ExternalBaseURL param.Opt[string] `json:"external_base_url,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsParametersEvaluationClassifyParametersModelToEvaluateEvaluationModelRequest](
		"model_source", "serverless", "dedicated", "external",
	)
}

// The properties InputDataFilePath, Judge, MaxScore, MinScore, PassThreshold are
// required.
type EvalNewParamsParametersEvaluationScoreParameters struct {
	// Data file ID
	InputDataFilePath string                                                `json:"input_data_file_path,required"`
	Judge             EvalNewParamsParametersEvaluationScoreParametersJudge `json:"judge,omitzero,required"`
	// Maximum possible score
	MaxScore float64 `json:"max_score,required"`
	// Minimum possible score
	MinScore float64 `json:"min_score,required"`
	// Score threshold for passing
	PassThreshold float64 `json:"pass_threshold,required"`
	// Field name in the input data
	ModelToEvaluate EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion `json:"model_to_evaluate,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationScoreParameters) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationScoreParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationScoreParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Model, ModelSource, SystemTemplate are required.
type EvalNewParamsParametersEvaluationScoreParametersJudge struct {
	// Name of the judge model
	Model string `json:"model,required"`
	// Source of the judge model.
	//
	// Any of "serverless", "dedicated", "external".
	ModelSource string `json:"model_source,omitzero,required"`
	// System prompt template for the judge
	SystemTemplate string `json:"system_template,required"`
	// Bearer/API token for external judge models.
	ExternalAPIToken param.Opt[string] `json:"external_api_token,omitzero"`
	// Base URL for external judge models. Must be OpenAI-compatible base URL.
	ExternalBaseURL param.Opt[string] `json:"external_base_url,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationScoreParametersJudge) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationScoreParametersJudge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationScoreParametersJudge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsParametersEvaluationScoreParametersJudge](
		"model_source", "serverless", "dedicated", "external",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion struct {
	OfString                                                                           param.Opt[string]                                                                      `json:",omitzero,inline"`
	OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest `json:",omitzero,inline"`
	paramUnion
}

func (u EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest)
}
func (u *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest) {
		return u.OfEvalNewsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest
	}
	return nil
}

// The properties InputTemplate, MaxTokens, Model, ModelSource, SystemTemplate,
// Temperature are required.
type EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest struct {
	// Input prompt template
	InputTemplate string `json:"input_template,required"`
	// Maximum number of tokens to generate
	MaxTokens int64 `json:"max_tokens,required"`
	// Name of the model to evaluate
	Model string `json:"model,required"`
	// Source of the model.
	//
	// Any of "serverless", "dedicated", "external".
	ModelSource string `json:"model_source,omitzero,required"`
	// System prompt template
	SystemTemplate string `json:"system_template,required"`
	// Sampling temperature
	Temperature float64 `json:"temperature,required"`
	// Bearer/API token for external models.
	ExternalAPIToken param.Opt[string] `json:"external_api_token,omitzero"`
	// Base URL for external models. Must be OpenAI-compatible base URL
	ExternalBaseURL param.Opt[string] `json:"external_base_url,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsParametersEvaluationScoreParametersModelToEvaluateEvaluationModelRequest](
		"model_source", "serverless", "dedicated", "external",
	)
}

// The properties InputDataFilePath, Judge are required.
type EvalNewParamsParametersEvaluationCompareParameters struct {
	// Data file name
	InputDataFilePath string                                                  `json:"input_data_file_path,required"`
	Judge             EvalNewParamsParametersEvaluationCompareParametersJudge `json:"judge,omitzero,required"`
	// Field name in the input data
	ModelA EvalNewParamsParametersEvaluationCompareParametersModelAUnion `json:"model_a,omitzero"`
	// Field name in the input data
	ModelB EvalNewParamsParametersEvaluationCompareParametersModelBUnion `json:"model_b,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationCompareParameters) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationCompareParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationCompareParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Model, ModelSource, SystemTemplate are required.
type EvalNewParamsParametersEvaluationCompareParametersJudge struct {
	// Name of the judge model
	Model string `json:"model,required"`
	// Source of the judge model.
	//
	// Any of "serverless", "dedicated", "external".
	ModelSource string `json:"model_source,omitzero,required"`
	// System prompt template for the judge
	SystemTemplate string `json:"system_template,required"`
	// Bearer/API token for external judge models.
	ExternalAPIToken param.Opt[string] `json:"external_api_token,omitzero"`
	// Base URL for external judge models. Must be OpenAI-compatible base URL.
	ExternalBaseURL param.Opt[string] `json:"external_base_url,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationCompareParametersJudge) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationCompareParametersJudge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationCompareParametersJudge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsParametersEvaluationCompareParametersJudge](
		"model_source", "serverless", "dedicated", "external",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalNewParamsParametersEvaluationCompareParametersModelAUnion struct {
	OfString                                                                    param.Opt[string]                                                               `json:",omitzero,inline"`
	OfEvalNewsParametersEvaluationCompareParametersModelAEvaluationModelRequest *EvalNewParamsParametersEvaluationCompareParametersModelAEvaluationModelRequest `json:",omitzero,inline"`
	paramUnion
}

func (u EvalNewParamsParametersEvaluationCompareParametersModelAUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfEvalNewsParametersEvaluationCompareParametersModelAEvaluationModelRequest)
}
func (u *EvalNewParamsParametersEvaluationCompareParametersModelAUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalNewParamsParametersEvaluationCompareParametersModelAUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfEvalNewsParametersEvaluationCompareParametersModelAEvaluationModelRequest) {
		return u.OfEvalNewsParametersEvaluationCompareParametersModelAEvaluationModelRequest
	}
	return nil
}

// The properties InputTemplate, MaxTokens, Model, ModelSource, SystemTemplate,
// Temperature are required.
type EvalNewParamsParametersEvaluationCompareParametersModelAEvaluationModelRequest struct {
	// Input prompt template
	InputTemplate string `json:"input_template,required"`
	// Maximum number of tokens to generate
	MaxTokens int64 `json:"max_tokens,required"`
	// Name of the model to evaluate
	Model string `json:"model,required"`
	// Source of the model.
	//
	// Any of "serverless", "dedicated", "external".
	ModelSource string `json:"model_source,omitzero,required"`
	// System prompt template
	SystemTemplate string `json:"system_template,required"`
	// Sampling temperature
	Temperature float64 `json:"temperature,required"`
	// Bearer/API token for external models.
	ExternalAPIToken param.Opt[string] `json:"external_api_token,omitzero"`
	// Base URL for external models. Must be OpenAI-compatible base URL
	ExternalBaseURL param.Opt[string] `json:"external_base_url,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationCompareParametersModelAEvaluationModelRequest) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationCompareParametersModelAEvaluationModelRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationCompareParametersModelAEvaluationModelRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsParametersEvaluationCompareParametersModelAEvaluationModelRequest](
		"model_source", "serverless", "dedicated", "external",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EvalNewParamsParametersEvaluationCompareParametersModelBUnion struct {
	OfString                                                                    param.Opt[string]                                                               `json:",omitzero,inline"`
	OfEvalNewsParametersEvaluationCompareParametersModelBEvaluationModelRequest *EvalNewParamsParametersEvaluationCompareParametersModelBEvaluationModelRequest `json:",omitzero,inline"`
	paramUnion
}

func (u EvalNewParamsParametersEvaluationCompareParametersModelBUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfEvalNewsParametersEvaluationCompareParametersModelBEvaluationModelRequest)
}
func (u *EvalNewParamsParametersEvaluationCompareParametersModelBUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EvalNewParamsParametersEvaluationCompareParametersModelBUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfEvalNewsParametersEvaluationCompareParametersModelBEvaluationModelRequest) {
		return u.OfEvalNewsParametersEvaluationCompareParametersModelBEvaluationModelRequest
	}
	return nil
}

// The properties InputTemplate, MaxTokens, Model, ModelSource, SystemTemplate,
// Temperature are required.
type EvalNewParamsParametersEvaluationCompareParametersModelBEvaluationModelRequest struct {
	// Input prompt template
	InputTemplate string `json:"input_template,required"`
	// Maximum number of tokens to generate
	MaxTokens int64 `json:"max_tokens,required"`
	// Name of the model to evaluate
	Model string `json:"model,required"`
	// Source of the model.
	//
	// Any of "serverless", "dedicated", "external".
	ModelSource string `json:"model_source,omitzero,required"`
	// System prompt template
	SystemTemplate string `json:"system_template,required"`
	// Sampling temperature
	Temperature float64 `json:"temperature,required"`
	// Bearer/API token for external models.
	ExternalAPIToken param.Opt[string] `json:"external_api_token,omitzero"`
	// Base URL for external models. Must be OpenAI-compatible base URL
	ExternalBaseURL param.Opt[string] `json:"external_base_url,omitzero"`
	paramObj
}

func (r EvalNewParamsParametersEvaluationCompareParametersModelBEvaluationModelRequest) MarshalJSON() (data []byte, err error) {
	type shadow EvalNewParamsParametersEvaluationCompareParametersModelBEvaluationModelRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalNewParamsParametersEvaluationCompareParametersModelBEvaluationModelRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EvalNewParamsParametersEvaluationCompareParametersModelBEvaluationModelRequest](
		"model_source", "serverless", "dedicated", "external",
	)
}

// The type of evaluation to perform
type EvalNewParamsType string

const (
	EvalNewParamsTypeClassify EvalNewParamsType = "classify"
	EvalNewParamsTypeScore    EvalNewParamsType = "score"
	EvalNewParamsTypeCompare  EvalNewParamsType = "compare"
)

type EvalListParams struct {
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Admin users can specify a user ID to filter jobs. Pass empty string to get all
	// jobs.
	UserID param.Opt[string] `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EvalListParams]'s query parameters as `url.Values`.
func (r EvalListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
