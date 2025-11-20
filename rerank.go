// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// RerankService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRerankService] method instead.
type RerankService struct {
	Options []option.RequestOption
}

// NewRerankService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRerankService(opts ...option.RequestOption) (r RerankService) {
	r = RerankService{}
	r.Options = opts
	return
}

// Query a reranker model
func (r *RerankService) New(ctx context.Context, body RerankNewParams, opts ...option.RequestOption) (res *RerankNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rerank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RerankNewResponse struct {
	// The model to be used for the rerank request.
	Model string `json:"model,required"`
	// Object type
	//
	// Any of "rerank".
	Object  RerankNewResponseObject   `json:"object,required"`
	Results []RerankNewResponseResult `json:"results,required"`
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
func (r RerankNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RerankNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object type
type RerankNewResponseObject string

const (
	RerankNewResponseObjectRerank RerankNewResponseObject = "rerank"
)

type RerankNewResponseResult struct {
	Document       RerankNewResponseResultDocument `json:"document,required"`
	Index          int64                           `json:"index,required"`
	RelevanceScore float64                         `json:"relevance_score,required"`
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
func (r RerankNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *RerankNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RerankNewResponseResultDocument struct {
	Text string `json:"text,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RerankNewResponseResultDocument) RawJSON() string { return r.JSON.raw }
func (r *RerankNewResponseResultDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RerankNewParams struct {
	// List of documents, which can be either strings or objects.
	Documents RerankNewParamsDocumentsUnion `json:"documents,omitzero,required"`
	// The model to be used for the rerank request.
	//
	// [See all of Together AI's rerank models](https://docs.together.ai/docs/serverless-models#rerank-models)
	Model RerankNewParamsModel `json:"model,omitzero,required"`
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

func (r RerankNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RerankNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RerankNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RerankNewParamsDocumentsUnion struct {
	OfMapOfAnyMap []map[string]any `json:",omitzero,inline"`
	OfStringArray []string         `json:",omitzero,inline"`
	paramUnion
}

func (u RerankNewParamsDocumentsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMapOfAnyMap, u.OfStringArray)
}
func (u *RerankNewParamsDocumentsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RerankNewParamsDocumentsUnion) asAny() any {
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
type RerankNewParamsModel string

const (
	RerankNewParamsModelSalesforceLlamaRankV1 RerankNewParamsModel = "Salesforce/Llama-Rank-v1"
)
