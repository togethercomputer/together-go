// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageParamContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsToolChoiceUnion()    {}
func (UnionString) ImplementsEmbeddingNewParamsInputUnion()              {}
func (UnionString) ImplementsExecuteResponseFailedExecutionErrorsUnion() {}
func (UnionString) ImplementsSessionListResponseErrorsUnion()            {}
func (UnionString) ImplementsEvaluationNewParamsParametersEvaluationClassifyParametersModelToEvaluateUnion() {
}
func (UnionString) ImplementsEvaluationNewParamsParametersEvaluationScoreParametersModelToEvaluateUnion() {
}
func (UnionString) ImplementsEvaluationNewParamsParametersEvaluationCompareParametersModelAUnion() {}
func (UnionString) ImplementsEvaluationNewParamsParametersEvaluationCompareParametersModelBUnion() {}

type UnionBool bool

func (UnionBool) ImplementsFineTuneTrainOnInputsUnion()               {}
func (UnionBool) ImplementsTrainingMethodSftTrainOnInputsUnionParam() {}
func (UnionBool) ImplementsTrainingMethodSftTrainOnInputsUnion()      {}
func (UnionBool) ImplementsFineTuneNewParamsTrainOnInputsUnion()      {}

type UnionInt int64

func (UnionInt) ImplementsFineTuneBatchSizeUnion()          {}
func (UnionInt) ImplementsFineTuneNewParamsBatchSizeUnion() {}
