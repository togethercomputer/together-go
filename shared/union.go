// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsChatCompletionNewParamsToolChoiceUnion() {}
func (UnionString) ImplementsEmbeddingNewParamsInputUnion()           {}

type UnionBool bool

func (UnionBool) ImplementsFineTuneTrainOnInputsUnion()          {}
func (UnionBool) ImplementsFineTuneNewParamsTrainOnInputsUnion() {}
