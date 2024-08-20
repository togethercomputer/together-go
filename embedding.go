// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"net/http"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
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
func NewEmbeddingService(opts ...option.RequestOption) (r *EmbeddingService) {
	r = &EmbeddingService{}
	r.Options = opts
	return
}

// Query an embedding model for a given string of text.
func (r *EmbeddingService) New(ctx context.Context, body EmbeddingNewParams, opts ...option.RequestOption) (res *Embedding, err error) {
	opts = append(r.Options[:], opts...)
	path := "embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Embedding struct {
	Data   []EmbeddingData `json:"data,required"`
	Model  string          `json:"model,required"`
	Object EmbeddingObject `json:"object,required"`
	JSON   embeddingJSON   `json:"-"`
}

// embeddingJSON contains the JSON metadata for the struct [Embedding]
type embeddingJSON struct {
	Data        apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Embedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r embeddingJSON) RawJSON() string {
	return r.raw
}

type EmbeddingData struct {
	Embedding []float64           `json:"embedding,required"`
	Index     int64               `json:"index,required"`
	Object    EmbeddingDataObject `json:"object,required"`
	JSON      embeddingDataJSON   `json:"-"`
}

// embeddingDataJSON contains the JSON metadata for the struct [EmbeddingData]
type embeddingDataJSON struct {
	Embedding   apijson.Field
	Index       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmbeddingData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r embeddingDataJSON) RawJSON() string {
	return r.raw
}

type EmbeddingDataObject string

const (
	EmbeddingDataObjectEmbedding EmbeddingDataObject = "embedding"
)

func (r EmbeddingDataObject) IsKnown() bool {
	switch r {
	case EmbeddingDataObjectEmbedding:
		return true
	}
	return false
}

type EmbeddingObject string

const (
	EmbeddingObjectList EmbeddingObject = "list"
)

func (r EmbeddingObject) IsKnown() bool {
	switch r {
	case EmbeddingObjectList:
		return true
	}
	return false
}

type EmbeddingNewParams struct {
	// A string providing the text for the model to embed.
	Input param.Field[EmbeddingNewParamsInputUnion] `json:"input,required"`
	// The name of the embedding model to use.
	Model param.Field[string] `json:"model,required"`
}

func (r EmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A string providing the text for the model to embed.
//
// Satisfied by [shared.UnionString], [EmbeddingNewParamsInputArray].
type EmbeddingNewParamsInputUnion interface {
	ImplementsEmbeddingNewParamsInputUnion()
}

type EmbeddingNewParamsInputArray []string

func (r EmbeddingNewParamsInputArray) ImplementsEmbeddingNewParamsInputUnion() {}
