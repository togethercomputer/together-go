// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaRlCheckpointService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaRlCheckpointService] method instead.
type BetaRlCheckpointService struct {
	Options []option.RequestOption
}

// NewBetaRlCheckpointService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaRlCheckpointService(opts ...option.RequestOption) (r BetaRlCheckpointService) {
	r = BetaRlCheckpointService{}
	r.Options = opts
	return
}

// Returns presigned URLs for downloading a checkpoint's model files. Only
// inference checkpoints support downloading.
func (r *BetaRlCheckpointService) Download(ctx context.Context, id string, query BetaRlCheckpointDownloadParams, opts ...option.RequestOption) (res *CheckpointDownloadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/checkpoints/%s/download", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Presigned download URLs for a checkpoint's files
type CheckpointDownloadResponse struct {
	// List of files with presigned download URLs
	Data []CheckpointFile `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckpointDownloadResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckpointDownloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A downloadable file within a checkpoint
type CheckpointFile struct {
	// Name of the file
	Filename string `json:"filename" api:"required"`
	// File size in bytes
	Size CheckpointFileSizeUnion `json:"size" api:"required"`
	// Presigned URL for downloading the file
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		Size        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckpointFile) RawJSON() string { return r.JSON.raw }
func (r *CheckpointFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CheckpointFileSizeUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type CheckpointFileSizeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u CheckpointFileSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CheckpointFileSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CheckpointFileSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *CheckpointFileSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Checkpoint variant: merged (full model) or adapter (LoRA weights only)
type CheckpointVariant string

const (
	CheckpointVariantCheckpointVariantUnspecified CheckpointVariant = "CHECKPOINT_VARIANT_UNSPECIFIED"
	CheckpointVariantCheckpointVariantMerged      CheckpointVariant = "CHECKPOINT_VARIANT_MERGED"
	CheckpointVariantCheckpointVariantAdapter     CheckpointVariant = "CHECKPOINT_VARIANT_ADAPTER"
)

type BetaRlCheckpointDownloadParams struct {
	// Checkpoint variant to download: merged (full model) or adapter (LoRA weights
	// only)
	//
	// Any of "CHECKPOINT_VARIANT_UNSPECIFIED", "CHECKPOINT_VARIANT_MERGED",
	// "CHECKPOINT_VARIANT_ADAPTER".
	Variant CheckpointVariant `query:"variant,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [BetaRlCheckpointDownloadParams]'s query parameters as
// `url.Values`.
func (r BetaRlCheckpointDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
