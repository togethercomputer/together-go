// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
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

// EndpointAdapterService contains methods and other services that help with
// interacting with the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEndpointAdapterService] method instead.
type EndpointAdapterService struct {
	Options []option.RequestOption
}

// NewEndpointAdapterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEndpointAdapterService(opts ...option.RequestOption) (r EndpointAdapterService) {
	r = EndpointAdapterService{}
	r.Options = opts
	return
}

// Attaches a LoRA adapter to a deployment. If the deployment is at adapter
// capacity, force can evict the oldest adapter.
func (r *EndpointAdapterService) New(ctx context.Context, params EndpointAdapterNewParams, opts ...option.RequestOption) (res *EndpointAdapterNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if params.DeploymentID == "" {
		err = errors.New("missing required deploymentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/deployments/%s/adapters", params.ProjectID.Value, params.EndpointID, params.DeploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Gets an adapter attachment by its `dad_` id and returns its per-cluster load
// state.
func (r *EndpointAdapterService) Get(ctx context.Context, id string, params EndpointAdapterGetParams, opts ...option.RequestOption) (res *EndpointAdapterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if params.DeploymentID == "" {
		err = errors.New("missing required deploymentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/deployments/%s/adapters/%s", params.ProjectID.Value, params.EndpointID, params.DeploymentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Updates the pinned revision of an adapter attachment by its `dad_` id using its
// row-level etag for optimistic concurrency.
func (r *EndpointAdapterService) Update(ctx context.Context, id string, params EndpointAdapterUpdateParams, opts ...option.RequestOption) (res *EndpointAdapterUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if params.DeploymentID == "" {
		err = errors.New("missing required deploymentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/deployments/%s/adapters/%s", params.ProjectID.Value, params.EndpointID, params.DeploymentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists LoRA adapters attached to a deployment with per-cluster load state.
func (r *EndpointAdapterService) List(ctx context.Context, endpointID string, deploymentID string, params EndpointAdapterListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[EndpointAdapterListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	if deploymentID == "" {
		err = errors.New("missing required deploymentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/deployments/%s/adapters", params.ProjectID.Value, endpointID, deploymentID)
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

// Lists LoRA adapters attached to a deployment with per-cluster load state.
func (r *EndpointAdapterService) ListAutoPaging(ctx context.Context, endpointID string, deploymentID string, params EndpointAdapterListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[EndpointAdapterListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, endpointID, deploymentID, params, opts...))
}

// Detaches an adapter attachment by its `dad_` id using its row-level etag for
// optimistic concurrency.
func (r *EndpointAdapterService) Delete(ctx context.Context, id string, params EndpointAdapterDeleteParams, opts ...option.RequestOption) (res *EndpointAdapterDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if params.DeploymentID == "" {
		err = errors.New("missing required deploymentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/projects/%s/endpoints/%s/deployments/%s/adapters/%s", params.ProjectID.Value, params.EndpointID, params.DeploymentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Controller-reported load state for an adapter on one deployment cluster.
type DeploymentAdapterStatus struct {
	// Adapter model identifier for this status row.
	AdapterModelID string `json:"adapterModelId" api:"required"`
	// Cluster reporting this adapter status.
	ClusterID string `json:"clusterId" api:"required"`
	// Number of pods that failed to load the adapter.
	FailedPodCount int64 `json:"failedPodCount" api:"required"`
	// Number of pods with the adapter ready to serve.
	ReadyPodCount int64 `json:"readyPodCount" api:"required"`
	// Current adapter load state in this cluster.
	//
	// Any of "ADAPTER_LOAD_STATE_PENDING", "ADAPTER_LOAD_STATE_LOADING",
	// "ADAPTER_LOAD_STATE_READY", "ADAPTER_LOAD_STATE_REMOVING",
	// "ADAPTER_LOAD_STATE_FAILED".
	State DeploymentAdapterStatusState `json:"state" api:"required"`
	// Total pods expected to report adapter load state.
	TotalPodCount int64 `json:"totalPodCount" api:"required"`
	// Resource name of the adapter model, using
	// projects/{projectId}/models/{adapterModelId}.
	AdapterModel string `json:"adapterModel"`
	// Optional inference-probe verdict for this adapter on this cluster; absent until
	// validation concludes.
	AdapterValid bool `json:"adapterValid"`
	// Human-readable probe rejection detail when adapterValid is false.
	AdapterValidReason string `json:"adapterValidReason"`
	// Time when the adapter first reached READY in this cluster.
	LoadedAt time.Time `json:"loadedAt" format:"date-time"`
	// Human-readable details about the current adapter state.
	Message string `json:"message"`
	// Adapter row etag observed by the controller when it wrote this status.
	RealizedEtag string `json:"realizedEtag"`
	// Resource name of the adapter model revision currently loaded in this cluster,
	// using projects/{projectId}/models/{adapterModelId}/revisions/{revisionId}.
	RealizedRevision string `json:"realizedRevision"`
	// Adapter revision currently loaded on pods in this cluster.
	RealizedRevisionID string `json:"realizedRevisionId"`
	// Stable reason code for the current adapter state.
	Reason string `json:"reason"`
	// Time when this adapter status was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdapterModelID     respjson.Field
		ClusterID          respjson.Field
		FailedPodCount     respjson.Field
		ReadyPodCount      respjson.Field
		State              respjson.Field
		TotalPodCount      respjson.Field
		AdapterModel       respjson.Field
		AdapterValid       respjson.Field
		AdapterValidReason respjson.Field
		LoadedAt           respjson.Field
		Message            respjson.Field
		RealizedEtag       respjson.Field
		RealizedRevision   respjson.Field
		RealizedRevisionID respjson.Field
		Reason             respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeploymentAdapterStatus) RawJSON() string { return r.JSON.raw }
func (r *DeploymentAdapterStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current adapter load state in this cluster.
type DeploymentAdapterStatusState string

const (
	DeploymentAdapterStatusStateAdapterLoadStatePending  DeploymentAdapterStatusState = "ADAPTER_LOAD_STATE_PENDING"
	DeploymentAdapterStatusStateAdapterLoadStateLoading  DeploymentAdapterStatusState = "ADAPTER_LOAD_STATE_LOADING"
	DeploymentAdapterStatusStateAdapterLoadStateReady    DeploymentAdapterStatusState = "ADAPTER_LOAD_STATE_READY"
	DeploymentAdapterStatusStateAdapterLoadStateRemoving DeploymentAdapterStatusState = "ADAPTER_LOAD_STATE_REMOVING"
	DeploymentAdapterStatusStateAdapterLoadStateFailed   DeploymentAdapterStatusState = "ADAPTER_LOAD_STATE_FAILED"
)

// Adapter attached to a deployment with desired revision and observed load state.
type EndpointAdapterNewResponse struct {
	// Adapter attachment `dad_` id used by get, update, and remove; stale ids never
	// address a replacement after remove and re-add.
	ID string `json:"id" api:"required"`
	// Adapter model identifier attached to the deployment.
	AdapterModelID string `json:"adapterModelId" api:"required"`
	// Adapter revision pinned on the deployment.
	DesiredRevisionID string `json:"desiredRevisionId" api:"required"`
	// Row-level etag required for UpdateAdapter and RemoveAdapter.
	Etag string `json:"etag" api:"required"`
	// Per-cluster adapter load state reported by the controller.
	PerCluster []DeploymentAdapterStatus `json:"perCluster" api:"required"`
	// Resource name of the adapter model, using
	// projects/{projectId}/models/{adapterModelId}.
	AdapterModel string `json:"adapterModel"`
	// Resource name of the adapter model revision pinned on the deployment, using
	// projects/{projectId}/models/{adapterModelId}/revisions/{revisionId}.
	DesiredRevision string `json:"desiredRevision"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AdapterModelID    respjson.Field
		DesiredRevisionID respjson.Field
		Etag              respjson.Field
		PerCluster        respjson.Field
		AdapterModel      respjson.Field
		DesiredRevision   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointAdapterNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointAdapterNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Adapter attached to a deployment with desired revision and observed load state.
type EndpointAdapterGetResponse struct {
	// Adapter attachment `dad_` id used by get, update, and remove; stale ids never
	// address a replacement after remove and re-add.
	ID string `json:"id" api:"required"`
	// Adapter model identifier attached to the deployment.
	AdapterModelID string `json:"adapterModelId" api:"required"`
	// Adapter revision pinned on the deployment.
	DesiredRevisionID string `json:"desiredRevisionId" api:"required"`
	// Row-level etag required for UpdateAdapter and RemoveAdapter.
	Etag string `json:"etag" api:"required"`
	// Per-cluster adapter load state reported by the controller.
	PerCluster []DeploymentAdapterStatus `json:"perCluster" api:"required"`
	// Resource name of the adapter model, using
	// projects/{projectId}/models/{adapterModelId}.
	AdapterModel string `json:"adapterModel"`
	// Resource name of the adapter model revision pinned on the deployment, using
	// projects/{projectId}/models/{adapterModelId}/revisions/{revisionId}.
	DesiredRevision string `json:"desiredRevision"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AdapterModelID    respjson.Field
		DesiredRevisionID respjson.Field
		Etag              respjson.Field
		PerCluster        respjson.Field
		AdapterModel      respjson.Field
		DesiredRevision   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointAdapterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointAdapterGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Adapter attached to a deployment with desired revision and observed load state.
type EndpointAdapterUpdateResponse struct {
	// Adapter attachment `dad_` id used by get, update, and remove; stale ids never
	// address a replacement after remove and re-add.
	ID string `json:"id" api:"required"`
	// Adapter model identifier attached to the deployment.
	AdapterModelID string `json:"adapterModelId" api:"required"`
	// Adapter revision pinned on the deployment.
	DesiredRevisionID string `json:"desiredRevisionId" api:"required"`
	// Row-level etag required for UpdateAdapter and RemoveAdapter.
	Etag string `json:"etag" api:"required"`
	// Per-cluster adapter load state reported by the controller.
	PerCluster []DeploymentAdapterStatus `json:"perCluster" api:"required"`
	// Resource name of the adapter model, using
	// projects/{projectId}/models/{adapterModelId}.
	AdapterModel string `json:"adapterModel"`
	// Resource name of the adapter model revision pinned on the deployment, using
	// projects/{projectId}/models/{adapterModelId}/revisions/{revisionId}.
	DesiredRevision string `json:"desiredRevision"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AdapterModelID    respjson.Field
		DesiredRevisionID respjson.Field
		Etag              respjson.Field
		PerCluster        respjson.Field
		AdapterModel      respjson.Field
		DesiredRevision   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointAdapterUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointAdapterUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Adapter attached to a deployment with desired revision and observed load state.
type EndpointAdapterListResponse struct {
	// Adapter attachment `dad_` id used by get, update, and remove; stale ids never
	// address a replacement after remove and re-add.
	ID string `json:"id" api:"required"`
	// Adapter model identifier attached to the deployment.
	AdapterModelID string `json:"adapterModelId" api:"required"`
	// Adapter revision pinned on the deployment.
	DesiredRevisionID string `json:"desiredRevisionId" api:"required"`
	// Row-level etag required for UpdateAdapter and RemoveAdapter.
	Etag string `json:"etag" api:"required"`
	// Per-cluster adapter load state reported by the controller.
	PerCluster []DeploymentAdapterStatus `json:"perCluster" api:"required"`
	// Resource name of the adapter model, using
	// projects/{projectId}/models/{adapterModelId}.
	AdapterModel string `json:"adapterModel"`
	// Resource name of the adapter model revision pinned on the deployment, using
	// projects/{projectId}/models/{adapterModelId}/revisions/{revisionId}.
	DesiredRevision string `json:"desiredRevision"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AdapterModelID    respjson.Field
		DesiredRevisionID respjson.Field
		Etag              respjson.Field
		PerCluster        respjson.Field
		AdapterModel      respjson.Field
		DesiredRevision   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointAdapterListResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointAdapterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Empty response returned after a successful delete operation.
type EndpointAdapterDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndpointAdapterDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *EndpointAdapterDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointAdapterNewParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Deployment identifier.
	DeploymentID string `path:"deploymentId" api:"required" json:"-"`
	// Adapter model identifier to attach.
	AdapterModelID string `json:"adapterModelId" api:"required"`
	// Optional adapter revision to pin. If omitted, the latest revision is resolved at
	// request time.
	AdapterRevisionID param.Opt[string] `json:"adapterRevisionId,omitzero"`
	// Whether to evict the oldest adapter if the deployment is at adapter capacity.
	Force param.Opt[bool] `json:"force,omitzero"`
	paramObj
}

func (r EndpointAdapterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointAdapterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointAdapterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointAdapterGetParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Deployment identifier.
	DeploymentID string `path:"deploymentId" api:"required" json:"-"`
	// Deprecated optional cross-check. When set, this must equal the `ml_` id of the
	// adapter model pinned by the live attachment; model names are not accepted.
	AdapterModelID param.Opt[string] `query:"adapterModelId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EndpointAdapterGetParams]'s query parameters as
// `url.Values`.
func (r EndpointAdapterGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EndpointAdapterUpdateParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Deployment identifier.
	DeploymentID string `path:"deploymentId" api:"required" json:"-"`
	// New adapter revision to pin.
	AdapterRevisionID string `json:"adapterRevisionId" api:"required"`
	// Row-level etag from a prior AddAdapter, UpdateAdapter, GetAdapter, or
	// ListAdapters response.
	Etag string `json:"etag" api:"required"`
	// Deprecated optional cross-check. When set, this must equal the `ml_` id of the
	// adapter model pinned by the live attachment; model names are not accepted.
	AdapterModelID param.Opt[string] `json:"adapterModelId,omitzero"`
	paramObj
}

func (r EndpointAdapterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EndpointAdapterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndpointAdapterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndpointAdapterListParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Cursor from a previous adapter list response.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Maximum number of adapters to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EndpointAdapterListParams]'s query parameters as
// `url.Values`.
func (r EndpointAdapterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EndpointAdapterDeleteParams struct {
	// Project identifier.
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[string] `path:"projectId,omitzero" api:"required" json:"-"`
	// Endpoint identifier.
	EndpointID string `path:"endpointId" api:"required" json:"-"`
	// Deployment identifier.
	DeploymentID string `path:"deploymentId" api:"required" json:"-"`
	// Adapter etag from a previous add, update, get, or list response. The removal is
	// rejected if the adapter changed after that response.
	Etag string `query:"etag" api:"required" json:"-"`
	// Deprecated optional cross-check. When set, this must equal the `ml_` id of the
	// adapter model pinned by the live attachment; model names are not accepted.
	AdapterModelID param.Opt[string] `query:"adapterModelId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EndpointAdapterDeleteParams]'s query parameters as
// `url.Values`.
func (r EndpointAdapterDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
