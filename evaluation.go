// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/tidwall/gjson"
	"github.com/togethercomputer/together-go/internal/apijson"
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

// Creates a new evaluation job for classify, score, or compare tasks
func (r *EvaluationService) New(ctx context.Context, body EvaluationNewParams, opts ...option.RequestOption) (res *EvaluationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "evaluation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details of a specific evaluation job
func (r *EvaluationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EvaluationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("evaluation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the status and results of a specific evaluation job
func (r *EvaluationService) GetStatus(ctx context.Context, id string, opts ...option.RequestOption) (res *EvaluationGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("evaluation/%s/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Internal callback endpoint for workflows to update job status and results
func (r *EvaluationService) UpdateStatus(ctx context.Context, id string, body EvaluationUpdateStatusParams, opts ...option.RequestOption) (res *EvaluationUpdateStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("evaluation/%s/update", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EvaluationJudgeModelConfigParam struct {
	// Name of the judge model
	ModelName param.Field[string] `json:"model_name,required"`
	// System prompt template for the judge
	SystemTemplate param.Field[string] `json:"system_template,required"`
}

func (r EvaluationJudgeModelConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvaluationModelRequestParam struct {
	// Input prompt template
	InputTemplate param.Field[string] `json:"input_template,required"`
	// Maximum number of tokens to generate
	MaxTokens param.Field[int64] `json:"max_tokens,required"`
	// Name of the model to evaluate
	ModelName param.Field[string] `json:"model_name,required"`
	// System prompt template
	SystemTemplate param.Field[string] `json:"system_template,required"`
	// Sampling temperature
	Temperature param.Field[float64] `json:"temperature,required"`
}

func (r EvaluationModelRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvaluationModelRequestParam) ImplementsEvaluationNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion() {
}

func (r EvaluationModelRequestParam) ImplementsEvaluationNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion() {
}

func (r EvaluationModelRequestParam) ImplementsEvaluationNewParamsParametersEvaluationCompareParametersModelAUnion() {
}

func (r EvaluationModelRequestParam) ImplementsEvaluationNewParamsParametersEvaluationCompareParametersModelBUnion() {
}

type EvaluationNewResponse struct {
	// Initial status of the job
	Status EvaluationNewResponseStatus `json:"status"`
	// The ID of the created evaluation job
	WorkflowID string                    `json:"workflow_id"`
	JSON       evaluationNewResponseJSON `json:"-"`
}

// evaluationNewResponseJSON contains the JSON metadata for the struct
// [EvaluationNewResponse]
type evaluationNewResponseJSON struct {
	Status      apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Initial status of the job
type EvaluationNewResponseStatus string

const (
	EvaluationNewResponseStatusPending EvaluationNewResponseStatus = "pending"
)

func (r EvaluationNewResponseStatus) IsKnown() bool {
	switch r {
	case EvaluationNewResponseStatusPending:
		return true
	}
	return false
}

type EvaluationGetResponse struct {
	// When the job was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ID of the job owner (admin only)
	OwnerID string `json:"owner_id"`
	// The parameters used for this evaluation
	Parameters map[string]interface{} `json:"parameters"`
	// Results of the evaluation (when completed)
	Results EvaluationGetResponseResults `json:"results,nullable"`
	// Current status of the job
	Status EvaluationGetResponseStatus `json:"status"`
	// History of status updates (admin only)
	StatusUpdates []EvaluationGetResponseStatusUpdate `json:"status_updates"`
	// The type of evaluation
	Type EvaluationGetResponseType `json:"type"`
	// When the job was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The evaluation job ID
	WorkflowID string                    `json:"workflow_id"`
	JSON       evaluationGetResponseJSON `json:"-"`
}

// evaluationGetResponseJSON contains the JSON metadata for the struct
// [EvaluationGetResponse]
type evaluationGetResponseJSON struct {
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

func (r *EvaluationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Results of the evaluation (when completed)
type EvaluationGetResponseResults struct {
	// Number of times model A won
	AWins int64 `json:"A_wins"`
	// This field can have the runtime type of
	// [EvaluationGetResponseResultsEvaluationScoreResultsAggregatedScores].
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
	Ties  int64                            `json:"Ties"`
	JSON  evaluationGetResponseResultsJSON `json:"-"`
	union EvaluationGetResponseResultsUnion
}

// evaluationGetResponseResultsJSON contains the JSON metadata for the struct
// [EvaluationGetResponseResults]
type evaluationGetResponseResultsJSON struct {
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

func (r evaluationGetResponseResultsJSON) RawJSON() string {
	return r.raw
}

func (r *EvaluationGetResponseResults) UnmarshalJSON(data []byte) (err error) {
	*r = EvaluationGetResponseResults{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvaluationGetResponseResultsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvaluationGetResponseResultsEvaluationClassifyResults],
// [EvaluationGetResponseResultsEvaluationScoreResults],
// [EvaluationGetResponseResultsEvaluationCompareResults],
// [EvaluationGetResponseResultsError].
func (r EvaluationGetResponseResults) AsUnion() EvaluationGetResponseResultsUnion {
	return r.union
}

// Results of the evaluation (when completed)
//
// Union satisfied by [EvaluationGetResponseResultsEvaluationClassifyResults],
// [EvaluationGetResponseResultsEvaluationScoreResults],
// [EvaluationGetResponseResultsEvaluationCompareResults] or
// [EvaluationGetResponseResultsError].
type EvaluationGetResponseResultsUnion interface {
	implementsEvaluationGetResponseResults()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvaluationGetResponseResultsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationGetResponseResultsEvaluationClassifyResults{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationGetResponseResultsEvaluationScoreResults{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationGetResponseResultsEvaluationCompareResults{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationGetResponseResultsError{}),
		},
	)
}

type EvaluationGetResponseResultsEvaluationClassifyResults struct {
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
	ResultFileID string                                                    `json:"result_file_id"`
	JSON         evaluationGetResponseResultsEvaluationClassifyResultsJSON `json:"-"`
}

// evaluationGetResponseResultsEvaluationClassifyResultsJSON contains the JSON
// metadata for the struct [EvaluationGetResponseResultsEvaluationClassifyResults]
type evaluationGetResponseResultsEvaluationClassifyResultsJSON struct {
	GenerationFailCount apijson.Field
	InvalidLabelCount   apijson.Field
	JudgeFailCount      apijson.Field
	LabelCounts         apijson.Field
	PassPercentage      apijson.Field
	ResultFileID        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvaluationGetResponseResultsEvaluationClassifyResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseResultsEvaluationClassifyResultsJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationGetResponseResultsEvaluationClassifyResults) implementsEvaluationGetResponseResults() {
}

type EvaluationGetResponseResultsEvaluationScoreResults struct {
	AggregatedScores EvaluationGetResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// number of failed samples generated from model
	FailedSamples float64 `json:"failed_samples"`
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// number of invalid scores generated from model
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// Data File ID
	ResultFileID string                                                 `json:"result_file_id"`
	JSON         evaluationGetResponseResultsEvaluationScoreResultsJSON `json:"-"`
}

// evaluationGetResponseResultsEvaluationScoreResultsJSON contains the JSON
// metadata for the struct [EvaluationGetResponseResultsEvaluationScoreResults]
type evaluationGetResponseResultsEvaluationScoreResultsJSON struct {
	AggregatedScores    apijson.Field
	FailedSamples       apijson.Field
	GenerationFailCount apijson.Field
	InvalidScoreCount   apijson.Field
	JudgeFailCount      apijson.Field
	ResultFileID        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvaluationGetResponseResultsEvaluationScoreResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseResultsEvaluationScoreResultsJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationGetResponseResultsEvaluationScoreResults) implementsEvaluationGetResponseResults() {
}

type EvaluationGetResponseResultsEvaluationScoreResultsAggregatedScores struct {
	MeanScore      float64                                                                `json:"mean_score"`
	PassPercentage float64                                                                `json:"pass_percentage"`
	StdScore       float64                                                                `json:"std_score"`
	JSON           evaluationGetResponseResultsEvaluationScoreResultsAggregatedScoresJSON `json:"-"`
}

// evaluationGetResponseResultsEvaluationScoreResultsAggregatedScoresJSON contains
// the JSON metadata for the struct
// [EvaluationGetResponseResultsEvaluationScoreResultsAggregatedScores]
type evaluationGetResponseResultsEvaluationScoreResultsAggregatedScoresJSON struct {
	MeanScore      apijson.Field
	PassPercentage apijson.Field
	StdScore       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvaluationGetResponseResultsEvaluationScoreResultsAggregatedScores) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseResultsEvaluationScoreResultsAggregatedScoresJSON) RawJSON() string {
	return r.raw
}

type EvaluationGetResponseResultsEvaluationCompareResults struct {
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
	Ties int64                                                    `json:"Ties"`
	JSON evaluationGetResponseResultsEvaluationCompareResultsJSON `json:"-"`
}

// evaluationGetResponseResultsEvaluationCompareResultsJSON contains the JSON
// metadata for the struct [EvaluationGetResponseResultsEvaluationCompareResults]
type evaluationGetResponseResultsEvaluationCompareResultsJSON struct {
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

func (r *EvaluationGetResponseResultsEvaluationCompareResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseResultsEvaluationCompareResultsJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationGetResponseResultsEvaluationCompareResults) implementsEvaluationGetResponseResults() {
}

type EvaluationGetResponseResultsError struct {
	Error string                                `json:"error"`
	JSON  evaluationGetResponseResultsErrorJSON `json:"-"`
}

// evaluationGetResponseResultsErrorJSON contains the JSON metadata for the struct
// [EvaluationGetResponseResultsError]
type evaluationGetResponseResultsErrorJSON struct {
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationGetResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationGetResponseResultsError) implementsEvaluationGetResponseResults() {}

// Current status of the job
type EvaluationGetResponseStatus string

const (
	EvaluationGetResponseStatusPending   EvaluationGetResponseStatus = "pending"
	EvaluationGetResponseStatusQueued    EvaluationGetResponseStatus = "queued"
	EvaluationGetResponseStatusRunning   EvaluationGetResponseStatus = "running"
	EvaluationGetResponseStatusCompleted EvaluationGetResponseStatus = "completed"
	EvaluationGetResponseStatusError     EvaluationGetResponseStatus = "error"
	EvaluationGetResponseStatusUserError EvaluationGetResponseStatus = "user_error"
)

func (r EvaluationGetResponseStatus) IsKnown() bool {
	switch r {
	case EvaluationGetResponseStatusPending, EvaluationGetResponseStatusQueued, EvaluationGetResponseStatusRunning, EvaluationGetResponseStatusCompleted, EvaluationGetResponseStatusError, EvaluationGetResponseStatusUserError:
		return true
	}
	return false
}

type EvaluationGetResponseStatusUpdate struct {
	// Additional message for this update
	Message string `json:"message"`
	// The status at this update
	Status string `json:"status"`
	// When this update occurred
	Timestamp time.Time                             `json:"timestamp" format:"date-time"`
	JSON      evaluationGetResponseStatusUpdateJSON `json:"-"`
}

// evaluationGetResponseStatusUpdateJSON contains the JSON metadata for the struct
// [EvaluationGetResponseStatusUpdate]
type evaluationGetResponseStatusUpdateJSON struct {
	Message     apijson.Field
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationGetResponseStatusUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseStatusUpdateJSON) RawJSON() string {
	return r.raw
}

// The type of evaluation
type EvaluationGetResponseType string

const (
	EvaluationGetResponseTypeClassify EvaluationGetResponseType = "classify"
	EvaluationGetResponseTypeScore    EvaluationGetResponseType = "score"
	EvaluationGetResponseTypeCompare  EvaluationGetResponseType = "compare"
)

func (r EvaluationGetResponseType) IsKnown() bool {
	switch r {
	case EvaluationGetResponseTypeClassify, EvaluationGetResponseTypeScore, EvaluationGetResponseTypeCompare:
		return true
	}
	return false
}

type EvaluationGetStatusResponse struct {
	Results EvaluationGetStatusResponseResults `json:"results,nullable"`
	Status  EvaluationGetStatusResponseStatus  `json:"status"`
	JSON    evaluationGetStatusResponseJSON    `json:"-"`
}

// evaluationGetStatusResponseJSON contains the JSON metadata for the struct
// [EvaluationGetStatusResponse]
type evaluationGetStatusResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetStatusResponseJSON) RawJSON() string {
	return r.raw
}

type EvaluationGetStatusResponseResults struct {
	// Number of times model A won
	AWins int64 `json:"A_wins"`
	// This field can have the runtime type of
	// [EvaluationGetStatusResponseResultsEvaluationScoreResultsAggregatedScores].
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
	Ties  int64                                  `json:"Ties"`
	JSON  evaluationGetStatusResponseResultsJSON `json:"-"`
	union EvaluationGetStatusResponseResultsUnion
}

// evaluationGetStatusResponseResultsJSON contains the JSON metadata for the struct
// [EvaluationGetStatusResponseResults]
type evaluationGetStatusResponseResultsJSON struct {
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

func (r evaluationGetStatusResponseResultsJSON) RawJSON() string {
	return r.raw
}

func (r *EvaluationGetStatusResponseResults) UnmarshalJSON(data []byte) (err error) {
	*r = EvaluationGetStatusResponseResults{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EvaluationGetStatusResponseResultsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [EvaluationGetStatusResponseResultsEvaluationClassifyResults],
// [EvaluationGetStatusResponseResultsEvaluationScoreResults],
// [EvaluationGetStatusResponseResultsEvaluationCompareResults],
// [EvaluationGetStatusResponseResultsError].
func (r EvaluationGetStatusResponseResults) AsUnion() EvaluationGetStatusResponseResultsUnion {
	return r.union
}

// Union satisfied by
// [EvaluationGetStatusResponseResultsEvaluationClassifyResults],
// [EvaluationGetStatusResponseResultsEvaluationScoreResults],
// [EvaluationGetStatusResponseResultsEvaluationCompareResults] or
// [EvaluationGetStatusResponseResultsError].
type EvaluationGetStatusResponseResultsUnion interface {
	implementsEvaluationGetStatusResponseResults()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvaluationGetStatusResponseResultsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationGetStatusResponseResultsEvaluationClassifyResults{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationGetStatusResponseResultsEvaluationScoreResults{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationGetStatusResponseResultsEvaluationCompareResults{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EvaluationGetStatusResponseResultsError{}),
		},
	)
}

type EvaluationGetStatusResponseResultsEvaluationClassifyResults struct {
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
	ResultFileID string                                                          `json:"result_file_id"`
	JSON         evaluationGetStatusResponseResultsEvaluationClassifyResultsJSON `json:"-"`
}

// evaluationGetStatusResponseResultsEvaluationClassifyResultsJSON contains the
// JSON metadata for the struct
// [EvaluationGetStatusResponseResultsEvaluationClassifyResults]
type evaluationGetStatusResponseResultsEvaluationClassifyResultsJSON struct {
	GenerationFailCount apijson.Field
	InvalidLabelCount   apijson.Field
	JudgeFailCount      apijson.Field
	LabelCounts         apijson.Field
	PassPercentage      apijson.Field
	ResultFileID        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvaluationGetStatusResponseResultsEvaluationClassifyResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetStatusResponseResultsEvaluationClassifyResultsJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationGetStatusResponseResultsEvaluationClassifyResults) implementsEvaluationGetStatusResponseResults() {
}

type EvaluationGetStatusResponseResultsEvaluationScoreResults struct {
	AggregatedScores EvaluationGetStatusResponseResultsEvaluationScoreResultsAggregatedScores `json:"aggregated_scores"`
	// number of failed samples generated from model
	FailedSamples float64 `json:"failed_samples"`
	// Number of failed generations.
	GenerationFailCount float64 `json:"generation_fail_count,nullable"`
	// number of invalid scores generated from model
	InvalidScoreCount float64 `json:"invalid_score_count"`
	// Number of failed judge generations
	JudgeFailCount float64 `json:"judge_fail_count,nullable"`
	// Data File ID
	ResultFileID string                                                       `json:"result_file_id"`
	JSON         evaluationGetStatusResponseResultsEvaluationScoreResultsJSON `json:"-"`
}

// evaluationGetStatusResponseResultsEvaluationScoreResultsJSON contains the JSON
// metadata for the struct
// [EvaluationGetStatusResponseResultsEvaluationScoreResults]
type evaluationGetStatusResponseResultsEvaluationScoreResultsJSON struct {
	AggregatedScores    apijson.Field
	FailedSamples       apijson.Field
	GenerationFailCount apijson.Field
	InvalidScoreCount   apijson.Field
	JudgeFailCount      apijson.Field
	ResultFileID        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EvaluationGetStatusResponseResultsEvaluationScoreResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetStatusResponseResultsEvaluationScoreResultsJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationGetStatusResponseResultsEvaluationScoreResults) implementsEvaluationGetStatusResponseResults() {
}

type EvaluationGetStatusResponseResultsEvaluationScoreResultsAggregatedScores struct {
	MeanScore      float64                                                                      `json:"mean_score"`
	PassPercentage float64                                                                      `json:"pass_percentage"`
	StdScore       float64                                                                      `json:"std_score"`
	JSON           evaluationGetStatusResponseResultsEvaluationScoreResultsAggregatedScoresJSON `json:"-"`
}

// evaluationGetStatusResponseResultsEvaluationScoreResultsAggregatedScoresJSON
// contains the JSON metadata for the struct
// [EvaluationGetStatusResponseResultsEvaluationScoreResultsAggregatedScores]
type evaluationGetStatusResponseResultsEvaluationScoreResultsAggregatedScoresJSON struct {
	MeanScore      apijson.Field
	PassPercentage apijson.Field
	StdScore       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvaluationGetStatusResponseResultsEvaluationScoreResultsAggregatedScores) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetStatusResponseResultsEvaluationScoreResultsAggregatedScoresJSON) RawJSON() string {
	return r.raw
}

type EvaluationGetStatusResponseResultsEvaluationCompareResults struct {
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
	Ties int64                                                          `json:"Ties"`
	JSON evaluationGetStatusResponseResultsEvaluationCompareResultsJSON `json:"-"`
}

// evaluationGetStatusResponseResultsEvaluationCompareResultsJSON contains the JSON
// metadata for the struct
// [EvaluationGetStatusResponseResultsEvaluationCompareResults]
type evaluationGetStatusResponseResultsEvaluationCompareResultsJSON struct {
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

func (r *EvaluationGetStatusResponseResultsEvaluationCompareResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetStatusResponseResultsEvaluationCompareResultsJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationGetStatusResponseResultsEvaluationCompareResults) implementsEvaluationGetStatusResponseResults() {
}

type EvaluationGetStatusResponseResultsError struct {
	Error string                                      `json:"error"`
	JSON  evaluationGetStatusResponseResultsErrorJSON `json:"-"`
}

// evaluationGetStatusResponseResultsErrorJSON contains the JSON metadata for the
// struct [EvaluationGetStatusResponseResultsError]
type evaluationGetStatusResponseResultsErrorJSON struct {
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationGetStatusResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetStatusResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

func (r EvaluationGetStatusResponseResultsError) implementsEvaluationGetStatusResponseResults() {}

type EvaluationGetStatusResponseStatus string

const (
	EvaluationGetStatusResponseStatusPending   EvaluationGetStatusResponseStatus = "pending"
	EvaluationGetStatusResponseStatusQueued    EvaluationGetStatusResponseStatus = "queued"
	EvaluationGetStatusResponseStatusRunning   EvaluationGetStatusResponseStatus = "running"
	EvaluationGetStatusResponseStatusCompleted EvaluationGetStatusResponseStatus = "completed"
	EvaluationGetStatusResponseStatusError     EvaluationGetStatusResponseStatus = "error"
	EvaluationGetStatusResponseStatusUserError EvaluationGetStatusResponseStatus = "user_error"
)

func (r EvaluationGetStatusResponseStatus) IsKnown() bool {
	switch r {
	case EvaluationGetStatusResponseStatusPending, EvaluationGetStatusResponseStatusQueued, EvaluationGetStatusResponseStatusRunning, EvaluationGetStatusResponseStatusCompleted, EvaluationGetStatusResponseStatusError, EvaluationGetStatusResponseStatusUserError:
		return true
	}
	return false
}

type EvaluationUpdateStatusResponse struct {
	Status     string                             `json:"status"`
	WorkflowID string                             `json:"workflow_id"`
	JSON       evaluationUpdateStatusResponseJSON `json:"-"`
}

// evaluationUpdateStatusResponseJSON contains the JSON metadata for the struct
// [EvaluationUpdateStatusResponse]
type evaluationUpdateStatusResponseJSON struct {
	Status      apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationUpdateStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationUpdateStatusResponseJSON) RawJSON() string {
	return r.raw
}

type EvaluationNewParams struct {
	// Type-specific parameters for the evaluation
	Parameters param.Field[EvaluationNewParamsParametersUnion] `json:"parameters,required"`
	// The type of evaluation to perform
	Type param.Field[EvaluationNewParamsType] `json:"type,required"`
}

func (r EvaluationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type-specific parameters for the evaluation
type EvaluationNewParamsParameters struct {
	// Data file ID
	InputDataFilePath param.Field[string]                          `json:"input_data_file_path,required"`
	Judge             param.Field[EvaluationJudgeModelConfigParam] `json:"judge,required"`
	Labels            param.Field[interface{}]                     `json:"labels"`
	// Maximum possible score
	MaxScore param.Field[float64] `json:"max_score"`
	// Minimum possible score
	MinScore        param.Field[float64]     `json:"min_score"`
	ModelA          param.Field[interface{}] `json:"model_a"`
	ModelB          param.Field[interface{}] `json:"model_b"`
	ModelToEvaluate param.Field[interface{}] `json:"model_to_evaluate"`
	PassLabels      param.Field[interface{}] `json:"pass_labels"`
	// Score threshold for passing
	PassThreshold param.Field[float64] `json:"pass_threshold"`
}

func (r EvaluationNewParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvaluationNewParamsParameters) implementsEvaluationNewParamsParametersUnion() {}

// Type-specific parameters for the evaluation
//
// Satisfied by [EvaluationNewParamsParametersEvaluationClassifyParameters],
// [EvaluationNewParamsParametersEvaluationScoreParameters],
// [EvaluationNewParamsParametersEvaluationCompareParameters],
// [EvaluationNewParamsParameters].
type EvaluationNewParamsParametersUnion interface {
	implementsEvaluationNewParamsParametersUnion()
}

type EvaluationNewParamsParametersEvaluationClassifyParameters struct {
	// Data file ID
	InputDataFilePath param.Field[string]                          `json:"input_data_file_path,required"`
	Judge             param.Field[EvaluationJudgeModelConfigParam] `json:"judge,required"`
	// List of possible classification labels
	Labels param.Field[[]string] `json:"labels,required"`
	// List of labels that are considered passing
	PassLabels param.Field[[]string] `json:"pass_labels,required"`
	// Field name in the input data
	ModelToEvaluate param.Field[EvaluationNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion] `json:"model_to_evaluate"`
}

func (r EvaluationNewParamsParametersEvaluationClassifyParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvaluationNewParamsParametersEvaluationClassifyParameters) implementsEvaluationNewParamsParametersUnion() {
}

// Field name in the input data
//
// Satisfied by [shared.UnionString], [EvaluationModelRequestParam].
type EvaluationNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion interface {
	ImplementsEvaluationNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion()
}

type EvaluationNewParamsParametersEvaluationScoreParameters struct {
	// Data file ID
	InputDataFilePath param.Field[string]                          `json:"input_data_file_path,required"`
	Judge             param.Field[EvaluationJudgeModelConfigParam] `json:"judge,required"`
	// Maximum possible score
	MaxScore param.Field[float64] `json:"max_score,required"`
	// Minimum possible score
	MinScore param.Field[float64] `json:"min_score,required"`
	// Score threshold for passing
	PassThreshold param.Field[float64] `json:"pass_threshold,required"`
	// Field name in the input data
	ModelToEvaluate param.Field[EvaluationNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion] `json:"model_to_evaluate"`
}

func (r EvaluationNewParamsParametersEvaluationScoreParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvaluationNewParamsParametersEvaluationScoreParameters) implementsEvaluationNewParamsParametersUnion() {
}

// Field name in the input data
//
// Satisfied by [shared.UnionString], [EvaluationModelRequestParam].
type EvaluationNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion interface {
	ImplementsEvaluationNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion()
}

type EvaluationNewParamsParametersEvaluationCompareParameters struct {
	// Data file name
	InputDataFilePath param.Field[string]                          `json:"input_data_file_path,required"`
	Judge             param.Field[EvaluationJudgeModelConfigParam] `json:"judge,required"`
	// Field name in the input data
	ModelA param.Field[EvaluationNewParamsParametersEvaluationCompareParametersModelAUnion] `json:"model_a"`
	// Field name in the input data
	ModelB param.Field[EvaluationNewParamsParametersEvaluationCompareParametersModelBUnion] `json:"model_b"`
}

func (r EvaluationNewParamsParametersEvaluationCompareParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvaluationNewParamsParametersEvaluationCompareParameters) implementsEvaluationNewParamsParametersUnion() {
}

// Field name in the input data
//
// Satisfied by [shared.UnionString], [EvaluationModelRequestParam].
type EvaluationNewParamsParametersEvaluationCompareParametersModelAUnion interface {
	ImplementsEvaluationNewParamsParametersEvaluationCompareParametersModelAUnion()
}

// Field name in the input data
//
// Satisfied by [shared.UnionString], [EvaluationModelRequestParam].
type EvaluationNewParamsParametersEvaluationCompareParametersModelBUnion interface {
	ImplementsEvaluationNewParamsParametersEvaluationCompareParametersModelBUnion()
}

// The type of evaluation to perform
type EvaluationNewParamsType string

const (
	EvaluationNewParamsTypeClassify EvaluationNewParamsType = "classify"
	EvaluationNewParamsTypeScore    EvaluationNewParamsType = "score"
	EvaluationNewParamsTypeCompare  EvaluationNewParamsType = "compare"
)

func (r EvaluationNewParamsType) IsKnown() bool {
	switch r {
	case EvaluationNewParamsTypeClassify, EvaluationNewParamsTypeScore, EvaluationNewParamsTypeCompare:
		return true
	}
	return false
}

type EvaluationUpdateStatusParams struct {
	// The new status for the job
	Status param.Field[EvaluationUpdateStatusParamsStatus] `json:"status,required"`
	// Error message
	Error   param.Field[string]                                   `json:"error"`
	Results param.Field[EvaluationUpdateStatusParamsResultsUnion] `json:"results"`
}

func (r EvaluationUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The new status for the job
type EvaluationUpdateStatusParamsStatus string

const (
	EvaluationUpdateStatusParamsStatusCompleted      EvaluationUpdateStatusParamsStatus = "completed"
	EvaluationUpdateStatusParamsStatusError          EvaluationUpdateStatusParamsStatus = "error"
	EvaluationUpdateStatusParamsStatusRunning        EvaluationUpdateStatusParamsStatus = "running"
	EvaluationUpdateStatusParamsStatusQueued         EvaluationUpdateStatusParamsStatus = "queued"
	EvaluationUpdateStatusParamsStatusUserError      EvaluationUpdateStatusParamsStatus = "user_error"
	EvaluationUpdateStatusParamsStatusInferenceError EvaluationUpdateStatusParamsStatus = "inference_error"
)

func (r EvaluationUpdateStatusParamsStatus) IsKnown() bool {
	switch r {
	case EvaluationUpdateStatusParamsStatusCompleted, EvaluationUpdateStatusParamsStatusError, EvaluationUpdateStatusParamsStatusRunning, EvaluationUpdateStatusParamsStatusQueued, EvaluationUpdateStatusParamsStatusUserError, EvaluationUpdateStatusParamsStatusInferenceError:
		return true
	}
	return false
}

type EvaluationUpdateStatusParamsResults struct {
	// Number of times model A won
	AWins            param.Field[int64]       `json:"A_wins"`
	AggregatedScores param.Field[interface{}] `json:"aggregated_scores"`
	// Number of times model B won
	BWins param.Field[int64] `json:"B_wins"`
	// number of failed samples generated from model
	FailedSamples param.Field[float64] `json:"failed_samples"`
	// Number of failed generations.
	GenerationFailCount param.Field[float64] `json:"generation_fail_count"`
	// Number of invalid labels
	InvalidLabelCount param.Field[float64] `json:"invalid_label_count"`
	// number of invalid scores generated from model
	InvalidScoreCount param.Field[float64] `json:"invalid_score_count"`
	// Number of failed judge generations
	JudgeFailCount param.Field[float64] `json:"judge_fail_count"`
	// JSON string representing label counts
	LabelCounts param.Field[string] `json:"label_counts"`
	// Total number of samples compared
	NumSamples param.Field[int64] `json:"num_samples"`
	// Pecentage of pass labels.
	PassPercentage param.Field[float64] `json:"pass_percentage"`
	// Data File ID
	ResultFileID param.Field[string] `json:"result_file_id"`
	// Number of ties
	Ties param.Field[int64] `json:"Ties"`
}

func (r EvaluationUpdateStatusParamsResults) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvaluationUpdateStatusParamsResults) implementsEvaluationUpdateStatusParamsResultsUnion() {}

// Satisfied by [EvaluationUpdateStatusParamsResultsEvaluationClassifyResults],
// [EvaluationUpdateStatusParamsResultsEvaluationScoreResults],
// [EvaluationUpdateStatusParamsResultsEvaluationCompareResults],
// [EvaluationUpdateStatusParamsResults].
type EvaluationUpdateStatusParamsResultsUnion interface {
	implementsEvaluationUpdateStatusParamsResultsUnion()
}

type EvaluationUpdateStatusParamsResultsEvaluationClassifyResults struct {
	// Number of failed generations.
	GenerationFailCount param.Field[float64] `json:"generation_fail_count"`
	// Number of invalid labels
	InvalidLabelCount param.Field[float64] `json:"invalid_label_count"`
	// Number of failed judge generations
	JudgeFailCount param.Field[float64] `json:"judge_fail_count"`
	// JSON string representing label counts
	LabelCounts param.Field[string] `json:"label_counts"`
	// Pecentage of pass labels.
	PassPercentage param.Field[float64] `json:"pass_percentage"`
	// Data File ID
	ResultFileID param.Field[string] `json:"result_file_id"`
}

func (r EvaluationUpdateStatusParamsResultsEvaluationClassifyResults) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvaluationUpdateStatusParamsResultsEvaluationClassifyResults) implementsEvaluationUpdateStatusParamsResultsUnion() {
}

type EvaluationUpdateStatusParamsResultsEvaluationScoreResults struct {
	AggregatedScores param.Field[EvaluationUpdateStatusParamsResultsEvaluationScoreResultsAggregatedScores] `json:"aggregated_scores"`
	// number of failed samples generated from model
	FailedSamples param.Field[float64] `json:"failed_samples"`
	// Number of failed generations.
	GenerationFailCount param.Field[float64] `json:"generation_fail_count"`
	// number of invalid scores generated from model
	InvalidScoreCount param.Field[float64] `json:"invalid_score_count"`
	// Number of failed judge generations
	JudgeFailCount param.Field[float64] `json:"judge_fail_count"`
	// Data File ID
	ResultFileID param.Field[string] `json:"result_file_id"`
}

func (r EvaluationUpdateStatusParamsResultsEvaluationScoreResults) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvaluationUpdateStatusParamsResultsEvaluationScoreResults) implementsEvaluationUpdateStatusParamsResultsUnion() {
}

type EvaluationUpdateStatusParamsResultsEvaluationScoreResultsAggregatedScores struct {
	MeanScore      param.Field[float64] `json:"mean_score"`
	PassPercentage param.Field[float64] `json:"pass_percentage"`
	StdScore       param.Field[float64] `json:"std_score"`
}

func (r EvaluationUpdateStatusParamsResultsEvaluationScoreResultsAggregatedScores) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvaluationUpdateStatusParamsResultsEvaluationCompareResults struct {
	// Number of times model A won
	AWins param.Field[int64] `json:"A_wins"`
	// Number of times model B won
	BWins param.Field[int64] `json:"B_wins"`
	// Number of failed generations.
	GenerationFailCount param.Field[float64] `json:"generation_fail_count"`
	// Number of failed judge generations
	JudgeFailCount param.Field[float64] `json:"judge_fail_count"`
	// Total number of samples compared
	NumSamples param.Field[int64] `json:"num_samples"`
	// Data File ID
	ResultFileID param.Field[string] `json:"result_file_id"`
	// Number of ties
	Ties param.Field[int64] `json:"Ties"`
}

func (r EvaluationUpdateStatusParamsResultsEvaluationCompareResults) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EvaluationUpdateStatusParamsResultsEvaluationCompareResults) implementsEvaluationUpdateStatusParamsResultsUnion() {
}
