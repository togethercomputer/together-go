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

// EmbeddingService contains methods and other services that help with interacting
// with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddingService] method instead.
type EmbeddingService struct {
	Options []option.RequestOption
}

// NewEmbeddingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddingService(opts ...option.RequestOption) (r EmbeddingService) {
	r = EmbeddingService{}
	r.Options = opts
	return
}

// Query an embedding model for a given string of text.
func (r *EmbeddingService) New(ctx context.Context, body EmbeddingNewParams, opts ...option.RequestOption) (res *Embedding, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Embedding struct {
	Data  []EmbeddingData `json:"data,required"`
	Model string          `json:"model,required"`
	// Any of "list".
	Object EmbeddingObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Embedding) RawJSON() string { return r.JSON.raw }
func (r *Embedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddingData struct {
	Embedding []float64 `json:"embedding,required"`
	Index     int64     `json:"index,required"`
	// Any of "embedding".
	Object string `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embedding   respjson.Field
		Index       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddingData) RawJSON() string { return r.JSON.raw }
func (r *EmbeddingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddingObject string

const (
	EmbeddingObjectList EmbeddingObject = "list"
)

type EmbeddingNewParams struct {
	// A string providing the text for the model to embed.
	Input EmbeddingNewParamsInputUnion `json:"input,omitzero,required"`
	// The name of the embedding model to use.
	//
	// [See all of Together AI's embedding models](https://docs.together.ai/docs/serverless-models#embedding-models)
	Model EmbeddingNewParamsModel `json:"model,omitzero,required"`
	paramObj
}

func (r EmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbeddingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbeddingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EmbeddingNewParamsInputUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u EmbeddingNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *EmbeddingNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EmbeddingNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The name of the embedding model to use.
//
// [See all of Together AI's embedding models](https://docs.together.ai/docs/serverless-models#embedding-models)
type EmbeddingNewParamsModel string

const (
	EmbeddingNewParamsModelWhereIsAIUaeLargeV1                  EmbeddingNewParamsModel = "WhereIsAI/UAE-Large-V1"
	EmbeddingNewParamsModelBaaiBgeLargeEnV1_5                   EmbeddingNewParamsModel = "BAAI/bge-large-en-v1.5"
	EmbeddingNewParamsModelBaaiBgeBaseEnV1_5                    EmbeddingNewParamsModel = "BAAI/bge-base-en-v1.5"
	EmbeddingNewParamsModelTogethercomputerM2Bert80M8kRetrieval EmbeddingNewParamsModel = "togethercomputer/m2-bert-80M-8k-retrieval"
)
