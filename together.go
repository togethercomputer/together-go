// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/param"
)

type RerankResponse struct {
	// The model to be used for the rerank request.
	Model string `json:"model,required"`
	// Object type
	Object  RerankResponseObject   `json:"object,required"`
	Results []RerankResponseResult `json:"results,required"`
	// Request ID
	ID    string              `json:"id"`
	Usage ChatCompletionUsage `json:"usage,nullable"`
	JSON  rerankResponseJSON  `json:"-"`
}

// rerankResponseJSON contains the JSON metadata for the struct [RerankResponse]
type rerankResponseJSON struct {
	Model       apijson.Field
	Object      apijson.Field
	Results     apijson.Field
	ID          apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RerankResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rerankResponseJSON) RawJSON() string {
	return r.raw
}

// Object type
type RerankResponseObject string

const (
	RerankResponseObjectRerank RerankResponseObject = "rerank"
)

func (r RerankResponseObject) IsKnown() bool {
	switch r {
	case RerankResponseObjectRerank:
		return true
	}
	return false
}

type RerankResponseResult struct {
	Document       RerankResponseResultsDocument `json:"document,required"`
	Index          int64                         `json:"index,required"`
	RelevanceScore float64                       `json:"relevance_score,required"`
	JSON           rerankResponseResultJSON      `json:"-"`
}

// rerankResponseResultJSON contains the JSON metadata for the struct
// [RerankResponseResult]
type rerankResponseResultJSON struct {
	Document       apijson.Field
	Index          apijson.Field
	RelevanceScore apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RerankResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rerankResponseResultJSON) RawJSON() string {
	return r.raw
}

type RerankResponseResultsDocument struct {
	Text string                            `json:"text,nullable"`
	JSON rerankResponseResultsDocumentJSON `json:"-"`
}

// rerankResponseResultsDocumentJSON contains the JSON metadata for the struct
// [RerankResponseResultsDocument]
type rerankResponseResultsDocumentJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RerankResponseResultsDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rerankResponseResultsDocumentJSON) RawJSON() string {
	return r.raw
}

type RerankParams struct {
	// List of documents, which can be either strings or objects.
	Documents param.Field[RerankParamsDocumentsUnion] `json:"documents,required"`
	// The model to be used for the rerank request.
	//
	// [See all of Together AI's rerank models](https://docs.together.ai/docs/serverless-models#rerank-models)
	Model param.Field[RerankParamsModel] `json:"model,required"`
	// The search query to be used for ranking.
	Query param.Field[string] `json:"query,required"`
	// List of keys in the JSON Object document to rank by. Defaults to use all
	// supplied keys for ranking.
	RankFields param.Field[[]string] `json:"rank_fields"`
	// Whether to return supplied documents with the response.
	ReturnDocuments param.Field[bool] `json:"return_documents"`
	// The number of top results to return.
	TopN param.Field[int64] `json:"top_n"`
}

func (r RerankParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// List of documents, which can be either strings or objects.
//
// Satisfied by [RerankParamsDocumentsArray], [RerankParamsDocumentsArray].
type RerankParamsDocumentsUnion interface {
	implementsRerankParamsDocumentsUnion()
}

type RerankParamsDocumentsArray []map[string]interface{}

func (r RerankParamsDocumentsArray) implementsRerankParamsDocumentsUnion() {}

// The model to be used for the rerank request.
//
// [See all of Together AI's rerank models](https://docs.together.ai/docs/serverless-models#rerank-models)
type RerankParamsModel string

const (
	RerankParamsModelSalesforceLlamaRankV1 RerankParamsModel = "Salesforce/Llama-Rank-v1"
)

func (r RerankParamsModel) IsKnown() bool {
	switch r {
	case RerankParamsModelSalesforceLlamaRankV1:
		return true
	}
	return false
}
