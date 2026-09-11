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
	"time"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/apiquery"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/param"
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

// Returns metadata for a checkpoint: type, base model, LoRA rank, step, and owning
// session.
func (r *BetaRlCheckpointService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Checkpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rl/checkpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists training checkpoints owned by the caller. Filter by session or base model
// to recover a checkpoint ID for resume. Inference checkpoints are not included;
// they remain on the training session and in the model catalog.
func (r *BetaRlCheckpointService) List(ctx context.Context, query BetaRlCheckpointListParams, opts ...option.RequestOption) (res *CheckpointsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rl/checkpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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

// Metadata for a saved checkpoint
type Checkpoint struct {
	// Unique identifier for the checkpoint
	ID string `json:"id" api:"required"`
	// Base model the checkpoint was trained from
	BaseModel string `json:"base_model" api:"required"`
	// Timestamp when the checkpoint was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Training session that produced the checkpoint
	SessionID string `json:"session_id" api:"required"`
	// Training step at time of save
	Step CheckpointStepUnion `json:"step" api:"required"`
	// Whether this is a training checkpoint or an inference checkpoint
	//
	// Any of "CHECKPOINT_TYPE_TRAINING", "CHECKPOINT_TYPE_INFERENCE".
	Type CheckpointType `json:"type" api:"required"`
	// LoRA rank of the session that produced this checkpoint. Absent for full-weight
	// sessions and for checkpoints saved before this field was recorded.
	LoraRank int64 `json:"lora_rank"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BaseModel   respjson.Field
		CreatedAt   respjson.Field
		SessionID   respjson.Field
		Step        respjson.Field
		Type        respjson.Field
		LoraRank    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Checkpoint) RawJSON() string { return r.JSON.raw }
func (r *Checkpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CheckpointStepUnion contains all possible properties and values from [string],
// [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type CheckpointStepUnion struct {
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

func (u CheckpointStepUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CheckpointStepUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CheckpointStepUnion) RawJSON() string { return u.JSON.raw }

func (r *CheckpointStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// Whether a checkpoint is saved for training resume or inference download.
type CheckpointType string

const (
	CheckpointTypeCheckpointTypeTraining  CheckpointType = "CHECKPOINT_TYPE_TRAINING"
	CheckpointTypeCheckpointTypeInference CheckpointType = "CHECKPOINT_TYPE_INFERENCE"
)

// Checkpoint variant: merged (full model) or adapter (LoRA weights only)
type CheckpointVariant string

const (
	CheckpointVariantCheckpointVariantUnspecified CheckpointVariant = "CHECKPOINT_VARIANT_UNSPECIFIED"
	CheckpointVariantCheckpointVariantMerged      CheckpointVariant = "CHECKPOINT_VARIANT_MERGED"
	CheckpointVariantCheckpointVariantAdapter     CheckpointVariant = "CHECKPOINT_VARIANT_ADAPTER"
)

// A page of training checkpoints
type CheckpointsListResponse struct {
	// Training checkpoints in this page
	Data []Checkpoint `json:"data" api:"required"`
	// Pagination metadata
	Meta CheckpointsListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckpointsListResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckpointsListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type CheckpointsListResponseMeta struct {
	// Whether more items exist beyond this page
	HasMore bool `json:"has_more"`
	// Maximum number of items returned per page
	Limit int64 `json:"limit"`
	// Cursor to use as the 'after' parameter for the next page. Empty when has_more is
	// false.
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Limit       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckpointsListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *CheckpointsListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaRlCheckpointListParams struct {
	// Cursor for pagination (ID of the last checkpoint from the previous page)
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Only return checkpoints trained from this base model. Match is exact.
	BaseModel param.Opt[string] `query:"base_model,omitzero" json:"-"`
	// Maximum number of checkpoints to return (1-100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return checkpoints produced by this training session
	SessionID param.Opt[string] `query:"session_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaRlCheckpointListParams]'s query parameters as
// `url.Values`.
func (r BetaRlCheckpointListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
