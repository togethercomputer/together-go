// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
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
)

// EvaluationService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvaluationService] method instead.
type EvaluationService struct {
	Options []option.RequestOption
}

// NewEvaluationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvaluationService(opts ...option.RequestOption) (r *EvaluationService) {
	r = &EvaluationService{}
	r.Options = opts
	return
}

// Get a list of evaluation jobs with optional filtering
func (r *EvaluationService) List(ctx context.Context, query EvaluationListParams, opts ...option.RequestOption) (res *[]EvaluationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "evaluations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the list of models that are allowed for evaluation
func (r *EvaluationService) GetAllowedModels(ctx context.Context, opts ...option.RequestOption) (res *EvaluationGetAllowedModelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "evaluations/model-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EvaluationListResponse struct {
	// When the job was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ID of the job owner (admin only)
	OwnerID string `json:"owner_id"`
	// The parameters used for this evaluation
	Parameters map[string]interface{} `json:"parameters"`
	// Results of the evaluation (when completed)
	Results EvaluationListResponseResults `json:"results,nullable"`
	// Current status of the job
	Status EvaluationListResponseStatus `json:"status"`
	// History of status updates (admin only)
	StatusUpdates []EvaluationListResponseStatusUpdate `json:"status_updates"`
	// The type of evaluation
	Type EvaluationListResponseType `json:"type"`
	// When the job was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The evaluation job ID
	WorkflowID string                     `json:"workflow_id"`
	JSON       evaluationListResponseJSON `json:"-"`
}

// evaluationListResponseJSON contains the JSON metadata for the struct
// [EvaluationListResponse]
type evaluationListResponseJSON struct {
	CreatedAt     apijson.Field
	OwnerID       apijson.Field
	Parameters    apijson.Field
	Results       apijson.Field
	Status        apijson.Field
	StatusUpdates apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	WorkflowID    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EvaluationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseJSON) RawJSON() string {
	return r.raw
}

// Results of the evaluation (when completed)
type EvaluationListResponseResults struct {
	// Number of times model A won
	AWins int64 `json:"A_wins"`
	// This field can have the runtime type of
	// [EvaluationListResponseResultsEvaluationScoreResultsAggregatedScores].
	AggregatedScores interface{} `json:"aggregated_scores"`
	// Number of times model B won
	BWins int64  `json:"B_wins"`
	Error string `json:"error"`
	// number of failed samples generated from model
	FailedSamples float64 `json:"failed_samples"`
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// Number of invalid labels
	InvalidLabelCount float64 `json:"invalid_label_count,nullable"`
	// number of invalid scores generated from model
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// JSON string representing label counts
	LabelCounts string `json:"label_counts"`
	// Total number of samples compared
	NumSamples int64 `json:"num_samples"`
	// Pecentage of pass labels.
	PassPercentage float64 `json:"pass_percentage,nullable"`
	// Data File ID
	ResultFileID string `json:"result_file_id"`
	// Number of ties
	Ties  int64                             `json:"Ties"`
	JSON  evaluationListResponseResultsJSON `json:"-"`
	union EvaluationListResponseResultsUnion
}

// evaluationListResponseResultsJSON contains the JSON metadata for the struct
// [EvaluationListResponseResults]
type evaluationListResponseResultsJSON struct {
	AWins               apijson.Field
	AggregatedScores    apijson.Field
	BWins               apijson.Field
	Error               apijson.Field
	FailedSamples       apijson.Field
	GenerationFailCount apijson.Field
	InvalidLabelCount   apijson.Field
	InvalidScoreCount   apijson.Field
	JudgeFailCount      apijson.Field
	LabelCounts         apijson.Field
	NumSamples          apijson.Field
	PassPercentage      apijson.Field
	ResultFileID        apijson.Field
	Ties                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r evaluationListResponseResultsJSON) RawJSON() string {
	return r.raw
}

func (r *EvaluationListResponseResults) UnmarshalJSON(data []byte) (err error) {
	*r = EvaluationListResponseResults{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvaluationListResponseResultsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvaluationListResponseResultsEvaluationClassifyResults],
// [EvaluationListResponseResultsEvaluationScoreResults],
// [EvaluationListResponseResultsEvaluationCompareResults],
// [EvaluationListResponseResultsError].
func (r EvaluationListResponseResults) AsUnion() EvaluationListResponseResultsUnion {
	return r.union
}

// Results of the evaluation (when completed)
//
// Union satisfied by [EvaluationListResponseResultsEvaluationClassifyResults],
// [EvaluationListResponseResultsEvaluationScoreResults],
// [EvaluationListResponseResultsEvaluationCompareResults] or
// [EvaluationListResponseResultsError].
type EvaluationListResponseResultsUnion interface {
	implementsEvaluationListResponseResults()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvaluationListResponseResultsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationListResponseResultsEvaluationClassifyResults{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationListResponseResultsEvaluationScoreResults{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationListResponseResultsEvaluationCompareResults{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationListResponseResultsError{}),
		},
	)
}

type EvaluationListResponseResultsEvaluationClassifyResults struct {
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
	ResultFileID string                                                     `json:"result_file_id"`
	JSON         evaluationListResponseResultsEvaluationClassifyResultsJSON `json:"-"`
}

// evaluationListResponseResultsEvaluationClassifyResultsJSON contains the JSON
// metadata for the struct [EvaluationListResponseResultsEvaluationClassifyResults]
type evaluationListResponseResultsEvaluationClassifyResultsJSON struct {
	GenerationFailCount apijson.Field
	InvalidLabelCount   apijson.Field
	JudgeFailCount      apijson.Field
	LabelCounts         apijson.Field
	PassPercentage      apijson.Field
	ResultFileID        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvaluationListResponseResultsEvaluationClassifyResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseResultsEvaluationClassifyResultsJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationListResponseResultsEvaluationClassifyResults) implementsEvaluationListResponseResults() {
}

type EvaluationListResponseResultsEvaluationScoreResults struct {
	AggregatedScores EvaluationListResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// number of failed samples generated from model
	FailedSamples float64 `json:"failed_samples"`
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// number of invalid scores generated from model
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// Data File ID
	ResultFileID string                                                  `json:"result_file_id"`
	JSON         evaluationListResponseResultsEvaluationScoreResultsJSON `json:"-"`
}

// evaluationListResponseResultsEvaluationScoreResultsJSON contains the JSON
// metadata for the struct [EvaluationListResponseResultsEvaluationScoreResults]
type evaluationListResponseResultsEvaluationScoreResultsJSON struct {
	AggregatedScores    apijson.Field
	FailedSamples       apijson.Field
	GenerationFailCount apijson.Field
	InvalidScoreCount   apijson.Field
	JudgeFailCount      apijson.Field
	ResultFileID        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvaluationListResponseResultsEvaluationScoreResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseResultsEvaluationScoreResultsJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationListResponseResultsEvaluationScoreResults) implementsEvaluationListResponseResults() {
}

type EvaluationListResponseResultsEvaluationScoreResultsAggregatedScores struct {
	MeanScore      float64                                                                 `json:"mean_score"`
	PassPercentage float64                                                                 `json:"pass_percentage"`
	StdScore       float64                                                                 `json:"std_score"`
	JSON           evaluationListResponseResultsEvaluationScoreResultsAggregatedScoresJSON `json:"-"`
}

// evaluationListResponseResultsEvaluationScoreResultsAggregatedScoresJSON contains
// the JSON metadata for the struct
// [EvaluationListResponseResultsEvaluationScoreResultsAggregatedScores]
type evaluationListResponseResultsEvaluationScoreResultsAggregatedScoresJSON struct {
	MeanScore      apijson.Field
	PassPercentage apijson.Field
	StdScore       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvaluationListResponseResultsEvaluationScoreResultsAggregatedScores) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseResultsEvaluationScoreResultsAggregatedScoresJSON) RawJSON() string {
	return r.raw
}

type EvaluationListResponseResultsEvaluationCompareResults struct {
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
	Ties int64                                                     `json:"Ties"`
	JSON evaluationListResponseResultsEvaluationCompareResultsJSON `json:"-"`
}

// evaluationListResponseResultsEvaluationCompareResultsJSON contains the JSON
// metadata for the struct [EvaluationListResponseResultsEvaluationCompareResults]
type evaluationListResponseResultsEvaluationCompareResultsJSON struct {
	AWins               apijson.Field
	BWins               apijson.Field
	GenerationFailCount apijson.Field
	JudgeFailCount      apijson.Field
	NumSamples          apijson.Field
	ResultFileID        apijson.Field
	Ties                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvaluationListResponseResultsEvaluationCompareResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseResultsEvaluationCompareResultsJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationListResponseResultsEvaluationCompareResults) implementsEvaluationListResponseResults() {
}

type EvaluationListResponseResultsError struct {
	Error string                                 `json:"error"`
	JSON  evaluationListResponseResultsErrorJSON `json:"-"`
}

// evaluationListResponseResultsErrorJSON contains the JSON metadata for the struct
// [EvaluationListResponseResultsError]
type evaluationListResponseResultsErrorJSON struct {
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationListResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationListResponseResultsError) implementsEvaluationListResponseResults() {}

// Current status of the job
type EvaluationListResponseStatus string

const (
	EvaluationListResponseStatusPending   EvaluationListResponseStatus = "pending"
	EvaluationListResponseStatusQueued    EvaluationListResponseStatus = "queued"
	EvaluationListResponseStatusRunning   EvaluationListResponseStatus = "running"
	EvaluationListResponseStatusCompleted EvaluationListResponseStatus = "completed"
	EvaluationListResponseStatusError     EvaluationListResponseStatus = "error"
	EvaluationListResponseStatusUserError EvaluationListResponseStatus = "user_error"
)

func (r EvaluationListResponseStatus) IsKnown() bool {
	switch r {
	case EvaluationListResponseStatusPending, EvaluationListResponseStatusQueued, EvaluationListResponseStatusRunning, EvaluationListResponseStatusCompleted, EvaluationListResponseStatusError, EvaluationListResponseStatusUserError:
		return true
	}
	return false
}

type EvaluationListResponseStatusUpdate struct {
	// Additional message for this update
	Message string `json:"message"`
	// The status at this update
	Status string `json:"status"`
	// When this update occurred
	Timestamp time.Time                              `json:"timestamp" format:"date-time"`
	JSON      evaluationListResponseStatusUpdateJSON `json:"-"`
}

// evaluationListResponseStatusUpdateJSON contains the JSON metadata for the struct
// [EvaluationListResponseStatusUpdate]
type evaluationListResponseStatusUpdateJSON struct {
	Message     apijson.Field
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationListResponseStatusUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseStatusUpdateJSON) RawJSON() string {
	return r.raw
}

// The type of evaluation
type EvaluationListResponseType string

const (
	EvaluationListResponseTypeClassify EvaluationListResponseType = "classify"
	EvaluationListResponseTypeScore    EvaluationListResponseType = "score"
	EvaluationListResponseTypeCompare  EvaluationListResponseType = "compare"
)

func (r EvaluationListResponseType) IsKnown() bool {
	switch r {
	case EvaluationListResponseTypeClassify, EvaluationListResponseTypeScore, EvaluationListResponseTypeCompare:
		return true
	}
	return false
}

type EvaluationGetAllowedModelsResponse struct {
	ModelList []string                               `json:"model_list"`
	JSON      evaluationGetAllowedModelsResponseJSON `json:"-"`
}

// evaluationGetAllowedModelsResponseJSON contains the JSON metadata for the struct
// [EvaluationGetAllowedModelsResponse]
type evaluationGetAllowedModelsResponseJSON struct {
	ModelList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationGetAllowedModelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetAllowedModelsResponseJSON) RawJSON() string {
	return r.raw
}

type EvaluationListParams struct {
	// Maximum number of results to return (max 100)
	Limit param.Field[int64] `query:"limit"`
	// Filter by job status
	Status param.Field[EvaluationListParamsStatus] `query:"status"`
}

// URLQuery serializes [EvaluationListParams]'s query parameters as `url.Values`.
func (r EvaluationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by job status
type EvaluationListParamsStatus string

const (
	EvaluationListParamsStatusPending   EvaluationListParamsStatus = "pending"
	EvaluationListParamsStatusQueued    EvaluationListParamsStatus = "queued"
	EvaluationListParamsStatusRunning   EvaluationListParamsStatus = "running"
	EvaluationListParamsStatusCompleted EvaluationListParamsStatus = "completed"
	EvaluationListParamsStatusError     EvaluationListParamsStatus = "error"
	EvaluationListParamsStatusUserError EvaluationListParamsStatus = "user_error"
)

func (r EvaluationListParamsStatus) IsKnown() bool {
	switch r {
	case EvaluationListParamsStatusPending, EvaluationListParamsStatusQueued, EvaluationListParamsStatusRunning, EvaluationListParamsStatusCompleted, EvaluationListParamsStatusError, EvaluationListParamsStatusUserError:
		return true
	}
	return false
}
