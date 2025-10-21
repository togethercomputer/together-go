// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

type RerankResponse struct {
	// The model to be used for the rerank request.
	Model string `json:"model,required"`
	// Object type
	//
	// Any of "rerank".
	Object  RerankResponseObject   `json:"object,required"`
	Results []RerankResponseResult `json:"results,required"`
	// Request ID
	ID    string              `json:"id"`
	Usage ChatCompletionUsage `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Model       respjson.Field
		Object      respjson.Field
		Results     respjson.Field
		ID          respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RerankResponse) RawJSON() string { return r.JSON.raw }
func (r *RerankResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object type
type RerankResponseObject string

const (
	RerankResponseObjectRerank RerankResponseObject = "rerank"
)

type RerankResponseResult struct {
	Document       RerankResponseResultDocument `json:"document,required"`
	Index          int64                        `json:"index,required"`
	RelevanceScore float64                      `json:"relevance_score,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document       respjson.Field
		Index          respjson.Field
		RelevanceScore respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RerankResponseResult) RawJSON() string { return r.JSON.raw }
func (r *RerankResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RerankResponseResultDocument struct {
	Text string `json:"text,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RerankResponseResultDocument) RawJSON() string { return r.JSON.raw }
func (r *RerankResponseResultDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RerankParams struct {
	// List of documents, which can be either strings or objects.
	Documents RerankParamsDocumentsUnion `json:"documents,omitzero,required"`
	// The model to be used for the rerank request.
	//
	// [See all of Together AI's rerank models](https://docs.together.ai/docs/serverless-models#rerank-models)
	Model RerankParamsModel `json:"model,omitzero,required"`
	// The search query to be used for ranking.
	Query string `json:"query,required"`
	// Whether to return supplied documents with the response.
	ReturnDocuments param.Opt[bool] `json:"return_documents,omitzero"`
	// The number of top results to return.
	TopN param.Opt[int64] `json:"top_n,omitzero"`
	// List of keys in the JSON Object document to rank by. Defaults to use all
	// supplied keys for ranking.
	RankFields []string `json:"rank_fields,omitzero"`
	paramObj
}

func (r RerankParams) MarshalJSON() (data []byte, err error) {
	type shadow RerankParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RerankParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RerankParamsDocumentsUnion struct {
	OfMapOfAnyMap []map[string]any `json:",omitzero,inline"`
	OfStringArray []string         `json:",omitzero,inline"`
	paramUnion
}

func (u RerankParamsDocumentsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMapOfAnyMap, u.OfStringArray)
}
func (u *RerankParamsDocumentsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RerankParamsDocumentsUnion) asAny() any {
	if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The model to be used for the rerank request.
//
// [See all of Together AI's rerank models](https://docs.together.ai/docs/serverless-models#rerank-models)
type RerankParamsModel string

const (
	RerankParamsModelSalesforceLlamaRankV1 RerankParamsModel = "Salesforce/Llama-Rank-v1"
)
