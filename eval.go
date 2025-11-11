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

// Get evaluation job details
func (r *EvalService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EvalGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("evaluation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all evaluation jobs. Deprecated! Please use /evaluation
func (r *EvalService) List(ctx context.Context, query EvalListParams, opts ...option.RequestOption) (res *[]EvalListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "evaluations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get model list
func (r *EvalService) GetAllowedModels(ctx context.Context, query EvalGetAllowedModelsParams, opts ...option.RequestOption) (res *EvalGetAllowedModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "evaluations/model-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get evaluation job status and results
func (r *EvalService) GetStatus(ctx context.Context, id string, opts ...option.RequestOption) (res *EvalGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("evaluation/%s/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EvalGetResponse struct {
	// When the job was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ID of the job owner (admin only)
	OwnerID string `json:"owner_id"`
	// The parameters used for this evaluation
	Parameters map[string]any `json:"parameters"`
	// Results of the evaluation (when completed)
	Results EvalGetResponseResultsUnion `json:"results,nullable"`
	// Current status of the job
	//
	// Any of "pending", "queued", "running", "completed", "error", "user_error".
	Status EvalGetResponseStatus `json:"status"`
	// History of status updates (admin only)
	StatusUpdates []EvalGetResponseStatusUpdate `json:"status_updates"`
	// The type of evaluation
	//
	// Any of "classify", "score", "compare".
	Type EvalGetResponseType `json:"type"`
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
func (r EvalGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvalGetResponseResultsUnion contains all possible properties and values from
// [EvalGetResponseResultsEvaluationClassifyResults],
// [EvalGetResponseResultsEvaluationScoreResults],
// [EvalGetResponseResultsEvaluationCompareResults], [EvalGetResponseResultsError].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EvalGetResponseResultsUnion struct {
	GenerationFailCount float64 `json:"generation_fail_count"`
	// This field is from variant [EvalGetResponseResultsEvaluationClassifyResults].
	InvalidLabelCount float64 `json:"invalid_label_count"`
	JudgeFailCount    float64 `json:"judge_fail_count"`
	// This field is from variant [EvalGetResponseResultsEvaluationClassifyResults].
	LabelCounts string `json:"label_counts"`
	// This field is from variant [EvalGetResponseResultsEvaluationClassifyResults].
	PassPercentage float64 `json:"pass_percentage"`
	ResultFileID   string  `json:"result_file_id"`
	// This field is from variant [EvalGetResponseResultsEvaluationScoreResults].
	AggregatedScores EvalGetResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// This field is from variant [EvalGetResponseResultsEvaluationScoreResults].
	FailedSamples float64 `json:"failed_samples"`
	// This field is from variant [EvalGetResponseResultsEvaluationScoreResults].
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// This field is from variant [EvalGetResponseResultsEvaluationCompareResults].
	AWins int64 `json:"A_wins"`
	// This field is from variant [EvalGetResponseResultsEvaluationCompareResults].
	BWins int64 `json:"B_wins"`
	// This field is from variant [EvalGetResponseResultsEvaluationCompareResults].
	NumSamples int64 `json:"num_samples"`
	// This field is from variant [EvalGetResponseResultsEvaluationCompareResults].
	Ties int64 `json:"Ties"`
	// This field is from variant [EvalGetResponseResultsError].
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

func (u EvalGetResponseResultsUnion) AsEvalGetResponseResultsEvaluationClassifyResults() (v EvalGetResponseResultsEvaluationClassifyResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalGetResponseResultsUnion) AsEvalGetResponseResultsEvaluationScoreResults() (v EvalGetResponseResultsEvaluationScoreResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalGetResponseResultsUnion) AsEvalGetResponseResultsEvaluationCompareResults() (v EvalGetResponseResultsEvaluationCompareResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalGetResponseResultsUnion) AsEvalGetResponseResultsError() (v EvalGetResponseResultsError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvalGetResponseResultsUnion) RawJSON() string { return u.JSON.raw }

func (r *EvalGetResponseResultsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetResponseResultsEvaluationClassifyResults struct {
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
func (r EvalGetResponseResultsEvaluationClassifyResults) RawJSON() string { return r.JSON.raw }
func (r *EvalGetResponseResultsEvaluationClassifyResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetResponseResultsEvaluationScoreResults struct {
	AggregatedScores EvalGetResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
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
func (r EvalGetResponseResultsEvaluationScoreResults) RawJSON() string { return r.JSON.raw }
func (r *EvalGetResponseResultsEvaluationScoreResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetResponseResultsEvaluationScoreResultsAggregatedScores struct {
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
func (r EvalGetResponseResultsEvaluationScoreResultsAggregatedScores) RawJSON() string {
	return r.JSON.raw
}
func (r *EvalGetResponseResultsEvaluationScoreResultsAggregatedScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetResponseResultsEvaluationCompareResults struct {
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
func (r EvalGetResponseResultsEvaluationCompareResults) RawJSON() string { return r.JSON.raw }
func (r *EvalGetResponseResultsEvaluationCompareResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetResponseResultsError struct {
	Error string `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalGetResponseResultsError) RawJSON() string { return r.JSON.raw }
func (r *EvalGetResponseResultsError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the job
type EvalGetResponseStatus string

const (
	EvalGetResponseStatusPending   EvalGetResponseStatus = "pending"
	EvalGetResponseStatusQueued    EvalGetResponseStatus = "queued"
	EvalGetResponseStatusRunning   EvalGetResponseStatus = "running"
	EvalGetResponseStatusCompleted EvalGetResponseStatus = "completed"
	EvalGetResponseStatusError     EvalGetResponseStatus = "error"
	EvalGetResponseStatusUserError EvalGetResponseStatus = "user_error"
)

type EvalGetResponseStatusUpdate struct {
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
func (r EvalGetResponseStatusUpdate) RawJSON() string { return r.JSON.raw }
func (r *EvalGetResponseStatusUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of evaluation
type EvalGetResponseType string

const (
	EvalGetResponseTypeClassify EvalGetResponseType = "classify"
	EvalGetResponseTypeScore    EvalGetResponseType = "score"
	EvalGetResponseTypeCompare  EvalGetResponseType = "compare"
)

type EvalListResponse struct {
	// When the job was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ID of the job owner (admin only)
	OwnerID string `json:"owner_id"`
	// The parameters used for this evaluation
	Parameters map[string]any `json:"parameters"`
	// Results of the evaluation (when completed)
	Results EvalListResponseResultsUnion `json:"results,nullable"`
	// Current status of the job
	//
	// Any of "pending", "queued", "running", "completed", "error", "user_error".
	Status EvalListResponseStatus `json:"status"`
	// History of status updates (admin only)
	StatusUpdates []EvalListResponseStatusUpdate `json:"status_updates"`
	// The type of evaluation
	//
	// Any of "classify", "score", "compare".
	Type EvalListResponseType `json:"type"`
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
func (r EvalListResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvalListResponseResultsUnion contains all possible properties and values from
// [EvalListResponseResultsEvaluationClassifyResults],
// [EvalListResponseResultsEvaluationScoreResults],
// [EvalListResponseResultsEvaluationCompareResults],
// [EvalListResponseResultsError].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EvalListResponseResultsUnion struct {
	GenerationFailCount float64 `json:"generation_fail_count"`
	// This field is from variant [EvalListResponseResultsEvaluationClassifyResults].
	InvalidLabelCount float64 `json:"invalid_label_count"`
	JudgeFailCount    float64 `json:"judge_fail_count"`
	// This field is from variant [EvalListResponseResultsEvaluationClassifyResults].
	LabelCounts string `json:"label_counts"`
	// This field is from variant [EvalListResponseResultsEvaluationClassifyResults].
	PassPercentage float64 `json:"pass_percentage"`
	ResultFileID   string  `json:"result_file_id"`
	// This field is from variant [EvalListResponseResultsEvaluationScoreResults].
	AggregatedScores EvalListResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// This field is from variant [EvalListResponseResultsEvaluationScoreResults].
	FailedSamples float64 `json:"failed_samples"`
	// This field is from variant [EvalListResponseResultsEvaluationScoreResults].
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// This field is from variant [EvalListResponseResultsEvaluationCompareResults].
	AWins int64 `json:"A_wins"`
	// This field is from variant [EvalListResponseResultsEvaluationCompareResults].
	BWins int64 `json:"B_wins"`
	// This field is from variant [EvalListResponseResultsEvaluationCompareResults].
	NumSamples int64 `json:"num_samples"`
	// This field is from variant [EvalListResponseResultsEvaluationCompareResults].
	Ties int64 `json:"Ties"`
	// This field is from variant [EvalListResponseResultsError].
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

func (u EvalListResponseResultsUnion) AsEvalListResponseResultsEvaluationClassifyResults() (v EvalListResponseResultsEvaluationClassifyResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalListResponseResultsUnion) AsEvalListResponseResultsEvaluationScoreResults() (v EvalListResponseResultsEvaluationScoreResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalListResponseResultsUnion) AsEvalListResponseResultsEvaluationCompareResults() (v EvalListResponseResultsEvaluationCompareResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalListResponseResultsUnion) AsEvalListResponseResultsError() (v EvalListResponseResultsError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvalListResponseResultsUnion) RawJSON() string { return u.JSON.raw }

func (r *EvalListResponseResultsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalListResponseResultsEvaluationClassifyResults struct {
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
func (r EvalListResponseResultsEvaluationClassifyResults) RawJSON() string { return r.JSON.raw }
func (r *EvalListResponseResultsEvaluationClassifyResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalListResponseResultsEvaluationScoreResults struct {
	AggregatedScores EvalListResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
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
func (r EvalListResponseResultsEvaluationScoreResults) RawJSON() string { return r.JSON.raw }
func (r *EvalListResponseResultsEvaluationScoreResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalListResponseResultsEvaluationScoreResultsAggregatedScores struct {
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
func (r EvalListResponseResultsEvaluationScoreResultsAggregatedScores) RawJSON() string {
	return r.JSON.raw
}
func (r *EvalListResponseResultsEvaluationScoreResultsAggregatedScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalListResponseResultsEvaluationCompareResults struct {
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
func (r EvalListResponseResultsEvaluationCompareResults) RawJSON() string { return r.JSON.raw }
func (r *EvalListResponseResultsEvaluationCompareResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalListResponseResultsError struct {
	Error string `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalListResponseResultsError) RawJSON() string { return r.JSON.raw }
func (r *EvalListResponseResultsError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the job
type EvalListResponseStatus string

const (
	EvalListResponseStatusPending   EvalListResponseStatus = "pending"
	EvalListResponseStatusQueued    EvalListResponseStatus = "queued"
	EvalListResponseStatusRunning   EvalListResponseStatus = "running"
	EvalListResponseStatusCompleted EvalListResponseStatus = "completed"
	EvalListResponseStatusError     EvalListResponseStatus = "error"
	EvalListResponseStatusUserError EvalListResponseStatus = "user_error"
)

type EvalListResponseStatusUpdate struct {
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
func (r EvalListResponseStatusUpdate) RawJSON() string { return r.JSON.raw }
func (r *EvalListResponseStatusUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of evaluation
type EvalListResponseType string

const (
	EvalListResponseTypeClassify EvalListResponseType = "classify"
	EvalListResponseTypeScore    EvalListResponseType = "score"
	EvalListResponseTypeCompare  EvalListResponseType = "compare"
)

type EvalGetAllowedModelsResponse struct {
	ModelList []string `json:"model_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelList   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalGetAllowedModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalGetAllowedModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetStatusResponse struct {
	// The results of the evaluation job
	Results EvalGetStatusResponseResultsUnion `json:"results"`
	// The status of the evaluation job
	//
	// Any of "completed", "error", "user_error", "running", "queued", "pending".
	Status EvalGetStatusResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvalGetStatusResponseResultsUnion contains all possible properties and values
// from [EvalGetStatusResponseResultsEvaluationClassifyResults],
// [EvalGetStatusResponseResultsEvaluationScoreResults],
// [EvalGetStatusResponseResultsEvaluationCompareResults].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EvalGetStatusResponseResultsUnion struct {
	GenerationFailCount float64 `json:"generation_fail_count"`
	// This field is from variant
	// [EvalGetStatusResponseResultsEvaluationClassifyResults].
	InvalidLabelCount float64 `json:"invalid_label_count"`
	JudgeFailCount    float64 `json:"judge_fail_count"`
	// This field is from variant
	// [EvalGetStatusResponseResultsEvaluationClassifyResults].
	LabelCounts string `json:"label_counts"`
	// This field is from variant
	// [EvalGetStatusResponseResultsEvaluationClassifyResults].
	PassPercentage float64 `json:"pass_percentage"`
	ResultFileID   string  `json:"result_file_id"`
	// This field is from variant [EvalGetStatusResponseResultsEvaluationScoreResults].
	AggregatedScores EvalGetStatusResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// This field is from variant [EvalGetStatusResponseResultsEvaluationScoreResults].
	FailedSamples float64 `json:"failed_samples"`
	// This field is from variant [EvalGetStatusResponseResultsEvaluationScoreResults].
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// This field is from variant
	// [EvalGetStatusResponseResultsEvaluationCompareResults].
	AWins int64 `json:"A_wins"`
	// This field is from variant
	// [EvalGetStatusResponseResultsEvaluationCompareResults].
	BWins int64 `json:"B_wins"`
	// This field is from variant
	// [EvalGetStatusResponseResultsEvaluationCompareResults].
	NumSamples int64 `json:"num_samples"`
	// This field is from variant
	// [EvalGetStatusResponseResultsEvaluationCompareResults].
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

func (u EvalGetStatusResponseResultsUnion) AsEvalGetStatusResponseResultsEvaluationClassifyResults() (v EvalGetStatusResponseResultsEvaluationClassifyResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalGetStatusResponseResultsUnion) AsEvalGetStatusResponseResultsEvaluationScoreResults() (v EvalGetStatusResponseResultsEvaluationScoreResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalGetStatusResponseResultsUnion) AsEvalGetStatusResponseResultsEvaluationCompareResults() (v EvalGetStatusResponseResultsEvaluationCompareResults) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvalGetStatusResponseResultsUnion) RawJSON() string { return u.JSON.raw }

func (r *EvalGetStatusResponseResultsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetStatusResponseResultsEvaluationClassifyResults struct {
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
func (r EvalGetStatusResponseResultsEvaluationClassifyResults) RawJSON() string { return r.JSON.raw }
func (r *EvalGetStatusResponseResultsEvaluationClassifyResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetStatusResponseResultsEvaluationScoreResults struct {
	AggregatedScores EvalGetStatusResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
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
func (r EvalGetStatusResponseResultsEvaluationScoreResults) RawJSON() string { return r.JSON.raw }
func (r *EvalGetStatusResponseResultsEvaluationScoreResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetStatusResponseResultsEvaluationScoreResultsAggregatedScores struct {
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
func (r EvalGetStatusResponseResultsEvaluationScoreResultsAggregatedScores) RawJSON() string {
	return r.JSON.raw
}
func (r *EvalGetStatusResponseResultsEvaluationScoreResultsAggregatedScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalGetStatusResponseResultsEvaluationCompareResults struct {
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
func (r EvalGetStatusResponseResultsEvaluationCompareResults) RawJSON() string { return r.JSON.raw }
func (r *EvalGetStatusResponseResultsEvaluationCompareResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the evaluation job
type EvalGetStatusResponseStatus string

const (
	EvalGetStatusResponseStatusCompleted EvalGetStatusResponseStatus = "completed"
	EvalGetStatusResponseStatusError     EvalGetStatusResponseStatus = "error"
	EvalGetStatusResponseStatusUserError EvalGetStatusResponseStatus = "user_error"
	EvalGetStatusResponseStatusRunning   EvalGetStatusResponseStatus = "running"
	EvalGetStatusResponseStatusQueued    EvalGetStatusResponseStatus = "queued"
	EvalGetStatusResponseStatusPending   EvalGetStatusResponseStatus = "pending"
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

type EvalGetAllowedModelsParams struct {
	ModelSource param.Opt[string] `query:"model_source,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EvalGetAllowedModelsParams]'s query parameters as
// `url.Values`.
func (r EvalGetAllowedModelsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
