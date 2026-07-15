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
	"github.com/togethercomputer/together-go/packages/pagination"
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// BetaEndpointShadowExperimentService contains methods and other services that
// help with interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaEndpointShadowExperimentService] method instead.
type BetaEndpointShadowExperimentService struct {
	Options []option.RequestOption
	Targets BetaEndpointShadowExperimentTargetService
}

// NewBetaEndpointShadowExperimentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBetaEndpointShadowExperimentService(opts ...option.RequestOption) (r BetaEndpointShadowExperimentService) {
	r = BetaEndpointShadowExperimentService{}
	r.Options = opts
	r.Targets = NewBetaEndpointShadowExperimentTargetService(opts...)
	return
}

// Creates an experiment that mirrors a sampled portion of endpoint traffic to one
// or more target deployments without returning their responses to clients. Add a
// description with the update operation after creation.
func (r *BetaEndpointShadowExperimentService) New(ctx context.Context, endpointID string, params BetaEndpointShadowExperimentNewParams, opts ...option.RequestOption) (res *ShadowExperiment, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments", params.ProjectID.Value, endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a shadow experiment, including its sampling strategy and target
// deployments.
func (r *BetaEndpointShadowExperimentService) Get(ctx context.Context, id string, query BetaEndpointShadowExperimentGetParams, opts ...option.RequestOption) (res *ShadowExperiment, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	if query.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if query.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments/%s", query.ProjectID.Value, query.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a shadow experiment's description or source sampling strategy.
// `updateMask` is required; source changes also require the current `etag` in the
// request body.
func (r *BetaEndpointShadowExperimentService) Update(ctx context.Context, id string, params BetaEndpointShadowExperimentUpdateParams, opts ...option.RequestOption) (res *ShadowExperiment, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments/%s", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists experiments that mirror sampled endpoint traffic to target deployments
// without affecting client responses. Set `includeTargets=true` to include target
// details inline.
func (r *BetaEndpointShadowExperimentService) List(ctx context.Context, endpointID string, params BetaEndpointShadowExperimentListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ShadowExperiment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if endpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments", params.ProjectID.Value, endpointID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists experiments that mirror sampled endpoint traffic to target deployments
// without affecting client responses. Set `includeTargets=true` to include target
// details inline.
func (r *BetaEndpointShadowExperimentService) ListAutoPaging(ctx context.Context, endpointID string, params BetaEndpointShadowExperimentListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ShadowExperiment] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, endpointID, params, opts...))
}

// Deletes a shadow experiment and its target records. The underlying deployments
// are not deleted.
func (r *BetaEndpointShadowExperimentService) Delete(ctx context.Context, id string, params BetaEndpointShadowExperimentDeleteParams, opts ...option.RequestOption) (res *BetaEndpointShadowExperimentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if params.ProjectID.Value == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if params.EndpointID == "" {
		err = errors.New("missing required endpointId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/endpoints/%s/shadowExperiments/%s", params.ProjectID.Value, params.EndpointID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Experiment that mirrors sampled endpoint requests to target deployments without
// changing client responses.
type ShadowExperiment struct {
	// Output only. Unique shadow experiment identifier.
	ID string `json:"id" api:"required"`
	// Timestamp when the experiment was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Identifier of the principal that created the experiment.
	CreatedBy string `json:"createdBy" api:"required"`
	// Output only. Endpoint whose traffic this experiment samples.
	EndpointID string `json:"endpointId" api:"required"`
	// Opaque version tag for optimistic concurrency control. Returned on read; set it
	// on update or delete requests for consistent read-modify-write.
	Etag string `json:"etag" api:"required"`
	// Human-readable shadow experiment name, unique within the endpoint. At most 256
	// characters.
	Name string `json:"name" api:"required"`
	// Output only. Project that owns the parent endpoint.
	ProjectID string `json:"projectId" api:"required"`
	// Endpoint traffic source returned for a shadow experiment.
	Source ShadowExperimentSource `json:"source" api:"required"`
	// Derived serving state, active when the experiment has at least one target.
	//
	// Any of "SHADOW_EXPERIMENT_STATE_ACTIVE", "SHADOW_EXPERIMENT_STATE_INACTIVE".
	State ShadowExperimentState `json:"state" api:"required"`
	// Target deployments that receive mirrored traffic.
	Targets []ShadowExperimentTarget `json:"targets" api:"required"`
	// Timestamp when the experiment was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// User defined description.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		EndpointID  respjson.Field
		Etag        respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		Source      respjson.Field
		State       respjson.Field
		Targets     respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperiment) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Endpoint traffic source returned for a shadow experiment.
type ShadowExperimentSource struct {
	// Endpoint-level source returned for a shadow experiment.
	Endpoint ShadowExperimentSourceEndpoint `json:"endpoint" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Endpoint    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSource) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperimentSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Endpoint-level source returned for a shadow experiment.
type ShadowExperimentSourceEndpoint struct {
	// Sampling strategy returned for endpoint-level shadow traffic.
	Sampling ShadowExperimentSourceEndpointSamplingUnion `json:"sampling" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sampling    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSourceEndpoint) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperimentSourceEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShadowExperimentSourceEndpointSamplingUnion contains all possible properties and
// values from [ShadowExperimentSourceEndpointSamplingUniform],
// [ShadowExperimentSourceEndpointSamplingKeyBased],
// [ShadowExperimentSourceEndpointSamplingAdaptiveUniform],
// [ShadowExperimentSourceEndpointSamplingAdaptiveKeyBased].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShadowExperimentSourceEndpointSamplingUnion struct {
	// This field is from variant [ShadowExperimentSourceEndpointSamplingUniform].
	Uniform ShadowExperimentSourceEndpointSamplingUniformUniform `json:"uniform"`
	// This field is from variant [ShadowExperimentSourceEndpointSamplingKeyBased].
	KeyBased ShadowExperimentSourceEndpointSamplingKeyBasedKeyBased `json:"keyBased"`
	// This field is from variant
	// [ShadowExperimentSourceEndpointSamplingAdaptiveUniform].
	AdaptiveUniform ShadowExperimentSourceEndpointSamplingAdaptiveUniformAdaptiveUniform `json:"adaptiveUniform"`
	// This field is from variant
	// [ShadowExperimentSourceEndpointSamplingAdaptiveKeyBased].
	AdaptiveKeyBased ShadowExperimentSourceEndpointSamplingAdaptiveKeyBasedAdaptiveKeyBased `json:"adaptiveKeyBased"`
	JSON             struct {
		Uniform          respjson.Field
		KeyBased         respjson.Field
		AdaptiveUniform  respjson.Field
		AdaptiveKeyBased respjson.Field
		raw              string
	} `json:"-"`
}

func (u ShadowExperimentSourceEndpointSamplingUnion) AsUniform() (v ShadowExperimentSourceEndpointSamplingUniform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShadowExperimentSourceEndpointSamplingUnion) AsKeyBased() (v ShadowExperimentSourceEndpointSamplingKeyBased) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShadowExperimentSourceEndpointSamplingUnion) AsAdaptiveUniform() (v ShadowExperimentSourceEndpointSamplingAdaptiveUniform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShadowExperimentSourceEndpointSamplingUnion) AsAdaptiveKeyBased() (v ShadowExperimentSourceEndpointSamplingAdaptiveKeyBased) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShadowExperimentSourceEndpointSamplingUnion) RawJSON() string { return u.JSON.raw }

func (r *ShadowExperimentSourceEndpointSamplingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShadowExperimentSourceEndpointSamplingUniform struct {
	// Fixed-rate random sampling returned by the API. A zero rate may be omitted by
	// JSON serialization.
	Uniform ShadowExperimentSourceEndpointSamplingUniformUniform `json:"uniform" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uniform     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSourceEndpointSamplingUniform) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperimentSourceEndpointSamplingUniform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fixed-rate random sampling returned by the API. A zero rate may be omitted by
// JSON serialization.
type ShadowExperimentSourceEndpointSamplingUniformUniform struct {
	// Fraction of requests sampled, from 0.0 to 1.0.
	Rate float64 `json:"rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rate        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSourceEndpointSamplingUniformUniform) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperimentSourceEndpointSamplingUniformUniform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShadowExperimentSourceEndpointSamplingKeyBased struct {
	// Fixed-rate sticky-key sampling returned by the API. A zero rate may be omitted
	// by JSON serialization.
	KeyBased ShadowExperimentSourceEndpointSamplingKeyBasedKeyBased `json:"keyBased" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeyBased    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSourceEndpointSamplingKeyBased) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperimentSourceEndpointSamplingKeyBased) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fixed-rate sticky-key sampling returned by the API. A zero rate may be omitted
// by JSON serialization.
type ShadowExperimentSourceEndpointSamplingKeyBasedKeyBased struct {
	// Request-body field used as the sticky sampling key.
	Key string `json:"key" api:"required"`
	// Fraction of distinct key values sampled, from 0.0 to 1.0.
	Rate float64 `json:"rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Rate        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSourceEndpointSamplingKeyBasedKeyBased) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperimentSourceEndpointSamplingKeyBasedKeyBased) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShadowExperimentSourceEndpointSamplingAdaptiveUniform struct {
	// Adaptive random sampling returned by the API.
	AdaptiveUniform ShadowExperimentSourceEndpointSamplingAdaptiveUniformAdaptiveUniform `json:"adaptiveUniform" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdaptiveUniform respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSourceEndpointSamplingAdaptiveUniform) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperimentSourceEndpointSamplingAdaptiveUniform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Adaptive random sampling returned by the API.
type ShadowExperimentSourceEndpointSamplingAdaptiveUniformAdaptiveUniform struct {
	// Per-gateway-replica target QPS.
	TargetQps float64 `json:"targetQps" api:"required"`
	// Sliding window for QPS observation when explicitly configured.
	Window string `json:"window"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TargetQps   respjson.Field
		Window      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSourceEndpointSamplingAdaptiveUniformAdaptiveUniform) RawJSON() string {
	return r.JSON.raw
}
func (r *ShadowExperimentSourceEndpointSamplingAdaptiveUniformAdaptiveUniform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShadowExperimentSourceEndpointSamplingAdaptiveKeyBased struct {
	// Adaptive sticky-key sampling returned by the API.
	AdaptiveKeyBased ShadowExperimentSourceEndpointSamplingAdaptiveKeyBasedAdaptiveKeyBased `json:"adaptiveKeyBased" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdaptiveKeyBased respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSourceEndpointSamplingAdaptiveKeyBased) RawJSON() string { return r.JSON.raw }
func (r *ShadowExperimentSourceEndpointSamplingAdaptiveKeyBased) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Adaptive sticky-key sampling returned by the API.
type ShadowExperimentSourceEndpointSamplingAdaptiveKeyBasedAdaptiveKeyBased struct {
	// Request-body field used as the sticky sampling key.
	Key string `json:"key" api:"required"`
	// Per-gateway-replica target QPS.
	TargetQps float64 `json:"targetQps" api:"required"`
	// Sliding window for QPS observation when explicitly configured.
	Window string `json:"window"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		TargetQps   respjson.Field
		Window      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShadowExperimentSourceEndpointSamplingAdaptiveKeyBasedAdaptiveKeyBased) RawJSON() string {
	return r.JSON.raw
}
func (r *ShadowExperimentSourceEndpointSamplingAdaptiveKeyBasedAdaptiveKeyBased) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Derived serving state, active when the experiment has at least one target.
type ShadowExperimentState string

const (
	ShadowExperimentStateShadowExperimentStateActive   ShadowExperimentState = "SHADOW_EXPERIMENT_STATE_ACTIVE"
	ShadowExperimentStateShadowExperimentStateInactive ShadowExperimentState = "SHADOW_EXPERIMENT_STATE_INACTIVE"
)

// Empty response returned after a successful delete operation.
type BetaEndpointShadowExperimentDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointShadowExperimentDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointShadowExperimentDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointShadowExperimentNewParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Human-readable shadow experiment name, unique within the endpoint. At most 256
	// characters.
	Name string `json:"name" api:"required"`
	// Traffic source for a shadow experiment. The public API supports endpoint sources
	// only.
	Source ShadowSourceParam `json:"source,omitzero" api:"required"`
	// Optional initial target deployments. At most 100 targets; manage later changes
	// through the target APIs.
	Targets []BetaEndpointShadowExperimentNewParamsTarget `json:"targets,omitzero"`
	paramObj
}

func (r BetaEndpointShadowExperimentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointShadowExperimentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointShadowExperimentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deployment under the parent endpoint that should receive mirrored requests from
// a shadow experiment.
//
// The properties Name, TargetDeploymentID are required.
type BetaEndpointShadowExperimentNewParamsTarget struct {
	// Human-readable target name, unique within the shadow experiment. At most 256
	// characters.
	Name string `json:"name" api:"required"`
	// Deployment under the parent endpoint that receives mirrored traffic. Exclude it
	// from the endpoint's live traffic split.
	TargetDeploymentID string `json:"targetDeploymentId" api:"required"`
	// Optional free-form target description.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r BetaEndpointShadowExperimentNewParamsTarget) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointShadowExperimentNewParamsTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointShadowExperimentNewParamsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaEndpointShadowExperimentGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	paramObj
}

type BetaEndpointShadowExperimentUpdateParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Required fields to update, such as description or source.
	UpdateMask string `query:"updateMask" api:"required" json:"-"`
	// Updated free-form description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Opaque version tag from a prior read for optimistic concurrency.
	Etag param.Opt[string] `json:"etag,omitzero"`
	// Traffic source for a shadow experiment. The public API supports endpoint sources
	// only.
	Source ShadowSourceParam `json:"source,omitzero"`
	paramObj
}

func (r BetaEndpointShadowExperimentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaEndpointShadowExperimentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaEndpointShadowExperimentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaEndpointShadowExperimentUpdateParams]'s query
// parameters as `url.Values`.
func (r BetaEndpointShadowExperimentUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointShadowExperimentListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous shadow experiment list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to include target deployments in each returned shadow experiment.
	IncludeTargets param.Opt[bool] `query:"includeTargets,omitzero" json:"-"`
	// Maximum number of shadow experiments to return. Max 500, defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointShadowExperimentListParams]'s query parameters
// as `url.Values`.
func (r BetaEndpointShadowExperimentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaEndpointShadowExperimentDeleteParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Etag for optimistic concurrency.
	Etag param.Opt[string] `query:"etag,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaEndpointShadowExperimentDeleteParams]'s query
// parameters as `url.Values`.
func (r BetaEndpointShadowExperimentDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
