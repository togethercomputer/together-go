// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/respjson"
	"github.com/togethercomputer/together-go/shared/constant"
)

// BetaEndpointHardwareService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaEndpointHardwareService] method instead.
type BetaEndpointHardwareService struct {
	Options []option.RequestOption
}

// NewBetaEndpointHardwareService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaEndpointHardwareService(opts ...option.RequestOption) (r BetaEndpointHardwareService) {
	r = BetaEndpointHardwareService{}
	r.Options = opts
	return
}

// Retrieves the GPU resources, pricing, regional availability, and best-effort
// capacity headroom for one inference instance type.
func (r *BetaEndpointHardwareService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *InferenceInstanceType, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/inference-instance-types/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists hardware instance types currently available to inference deployments,
// including GPU resources, pricing, regions, and best-effort capacity headroom.
func (r *BetaEndpointHardwareService) List(ctx context.Context, opts ...option.RequestOption) (res *BetaEndpointHardwareListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.ai/v2/")}, opts...)
	path := "public/inference-instance-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// GPU hardware configuration on which one inference replica can run.
type InferenceInstanceType struct {
	// Stable hardware instance type identifier used by deployment configs.
	ID string `json:"id" api:"required"`
	// Human-readable summary of the hardware configuration.
	Description string `json:"description" api:"required"`
	// Number of GPUs in one replica of this instance type.
	GPUCount int64 `json:"gpuCount" api:"required"`
	// Memory available on each GPU, in GiB.
	GPUMemoryGib int64 `json:"gpuMemoryGib" api:"required"`
	// GPU accelerator model, such as `H100` or `B200`.
	GPUType string `json:"gpuType" api:"required"`
	// Human-readable instance type name.
	Name string `json:"name" api:"required"`
	// On-demand price for one running replica, in US cents per hour.
	PriceCentsPerHour int64 `json:"priceCentsPerHour" api:"required"`
	// Regions where this instance type is offered.
	Regions []InferenceInstanceTypeRegion `json:"regions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Description       respjson.Field
		GPUCount          respjson.Field
		GPUMemoryGib      respjson.Field
		GPUType           respjson.Field
		Name              respjson.Field
		PriceCentsPerHour respjson.Field
		Regions           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceInstanceType) RawJSON() string { return r.JSON.raw }
func (r *InferenceInstanceType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where an instance type is offered.
type InferenceInstanceTypeRegion struct {
	// Region name where an instance type is offered.
	Name string `json:"name" api:"required"`
	// Best-effort estimate of how many additional replicas currently fit in a region.
	Headroom InferenceInstanceTypeRegionHeadroom `json:"headroom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Headroom    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceInstanceTypeRegion) RawJSON() string { return r.JSON.raw }
func (r *InferenceInstanceTypeRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Best-effort estimate of how many additional replicas currently fit in a region.
type InferenceInstanceTypeRegionHeadroom struct {
	// Whether the value is exact or a lower bound.
	//
	// Any of "RELATION_EQ", "RELATION_GTE".
	Relation string `json:"relation" api:"required"`
	// Capped count of replicas that currently fit.
	Value int64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Relation    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceInstanceTypeRegionHeadroom) RawJSON() string { return r.JSON.raw }
func (r *InferenceInstanceTypeRegionHeadroom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hardware instance types available for inference deployments.
type BetaEndpointHardwareListResponse struct {
	// Instance types available for inference.
	Data []InferenceInstanceType `json:"data" api:"required"`
	// Object type. Always `list`.
	Object constant.List `json:"object" default:"list"`
	// Cursor for the next page. Always null today because this catalog is returned in
	// full.
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaEndpointHardwareListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaEndpointHardwareListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
