// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsChatCompletionNewParamsMessagesContentUnion() {}
func (UnionString) ImplementsChatCompletionNewParamsToolChoiceUnion()      {}
func (UnionString) ImplementsEmbeddingNewParamsInputUnion()                {}
func (UnionString) ImplementsExecuteResponseFailedExecutionErrorsUnion()   {}
func (UnionString) ImplementsSessionListResponseErrorsUnion()              {}

type UnionBool bool

func (UnionBool) ImplementsFineTuneTrainOnInputsUnion()          {}
func (UnionBool) ImplementsFineTuneNewParamsTrainOnInputsUnion() {}

type UnionInt int64

func (UnionInt) ImplementsFineTuneBatchSizeUnion()          {}
func (UnionInt) ImplementsFineTuneNewParamsBatchSizeUnion() {}
