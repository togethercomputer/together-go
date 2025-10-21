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
	"github.com/togethercomputer/together-go/packages/param"
	"github.com/togethercomputer/together-go/packages/respjson"
)

// VideoService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVideoService] method instead.
type VideoService struct {
	Options []option.RequestOption
}

// NewVideoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVideoService(opts ...option.RequestOption) (r VideoService) {
	r = VideoService{}
	r.Options = opts
	return
}

// Create a video
func (r *VideoService) New(ctx context.Context, body VideoNewParams, opts ...option.RequestOption) (res *VideoNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.xyz/v2/")}, opts...)
	path := "videos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch video metadata
func (r *VideoService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoJob, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.xyz/v2/")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("videos/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Structured information describing a generated video job.
type VideoJob struct {
	// Unique identifier for the video job.
	ID string `json:"id,required"`
	// Unix timestamp (seconds) for when the job was created.
	CreatedAt float64 `json:"created_at,required"`
	// The video generation model that produced the job.
	Model string `json:"model,required"`
	// Duration of the generated clip in seconds.
	Seconds string `json:"seconds,required"`
	// The resolution of the generated video.
	Size string `json:"size,required"`
	// Current lifecycle status of the video job.
	//
	// Any of "in_progress", "completed", "failed".
	Status VideoJobStatus `json:"status,required"`
	// Unix timestamp (seconds) for when the job completed, if finished.
	CompletedAt float64 `json:"completed_at"`
	// Error payload that explains why generation failed, if applicable.
	Error VideoJobError `json:"error"`
	// The object type, which is always video.
	//
	// Any of "video".
	Object VideoJobObject `json:"object"`
	// Available upon completion, the outputs provides the cost charged and the hosted
	// url to access the video
	Outputs VideoJobOutputs `json:"outputs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Model       respjson.Field
		Seconds     respjson.Field
		Size        respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		Error       respjson.Field
		Object      respjson.Field
		Outputs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoJob) RawJSON() string { return r.JSON.raw }
func (r *VideoJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current lifecycle status of the video job.
type VideoJobStatus string

const (
	VideoJobStatusInProgress VideoJobStatus = "in_progress"
	VideoJobStatusCompleted  VideoJobStatus = "completed"
	VideoJobStatusFailed     VideoJobStatus = "failed"
)

// Error payload that explains why generation failed, if applicable.
type VideoJobError struct {
	Message string `json:"message,required"`
	Code    string `json:"code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Code        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoJobError) RawJSON() string { return r.JSON.raw }
func (r *VideoJobError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type, which is always video.
type VideoJobObject string

const (
	VideoJobObjectVideo VideoJobObject = "video"
)

// Available upon completion, the outputs provides the cost charged and the hosted
// url to access the video
type VideoJobOutputs struct {
	// The cost of generated video charged to the owners account.
	Cost int64 `json:"cost,required"`
	// URL hosting the generated video
	VideoURL string `json:"video_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		VideoURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoJobOutputs) RawJSON() string { return r.JSON.raw }
func (r *VideoJobOutputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoNewResponse struct {
	// Unique identifier for the video job.
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoNewParams struct {
	// The model to be used for the video creation request.
	Model string `json:"model,required"`
	// Frames per second. Defaults to 24.
	Fps param.Opt[int64] `json:"fps,omitzero"`
	// Controls how closely the video generation follows your prompt. Higher values
	// make the model adhere more strictly to your text description, while lower values
	// allow more creative freedom. guidence_scale affects both visual content and
	// temporal consistency.Recommended range is 6.0-10.0 for most video models. Values
	// above 12 may cause over-guidance artifacts or unnatural motion patterns.
	GuidanceScale param.Opt[int64] `json:"guidance_scale,omitzero"`
	Height        param.Opt[int64] `json:"height,omitzero"`
	// Similar to prompt, but specifies what to avoid instead of what to include
	NegativePrompt param.Opt[string] `json:"negative_prompt,omitzero"`
	// Compression quality. Defaults to 20.
	OutputQuality param.Opt[int64] `json:"output_quality,omitzero"`
	// Text prompt that describes the video to generate.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// Clip duration in seconds.
	Seconds param.Opt[string] `json:"seconds,omitzero"`
	// Seed to use in initializing the video generation. Using the same seed allows
	// deterministic video generation. If not provided a random seed is generated for
	// each request.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// The number of denoising steps the model performs during video generation. More
	// steps typically result in higher quality output but require longer processing
	// time.
	Steps param.Opt[int64] `json:"steps,omitzero"`
	Width param.Opt[int64] `json:"width,omitzero"`
	// Array of images to guide video generation, similar to keyframes.
	FrameImages []VideoNewParamsFrameImage `json:"frame_images,omitzero"`
	// Specifies the format of the output video. Defaults to MP4.
	//
	// Any of "MP4", "WEBM".
	OutputFormat VideoNewParamsOutputFormat `json:"output_format,omitzero"`
	// Unlike frame_images which constrain specific timeline positions, reference
	// images guide the general appearance that should appear consistently across the
	// video.
	ReferenceImages []string `json:"reference_images,omitzero"`
	paramObj
}

func (r VideoNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property InputImage is required.
type VideoNewParamsFrameImage struct {
	// URL path to hosted image that is used for a frame
	InputImage string `json:"input_image,required"`
	// Optional param to specify where to insert the frame. If this is omitted, the
	// following heuristics are applied:
	//
	// - frame_images size is one, frame is first.
	// - If size is two, frames are first and last.
	// - If size is larger, frames are first, last and evenly spaced between.
	Frame VideoNewParamsFrameImageFrameUnion `json:"frame,omitzero"`
	paramObj
}

func (r VideoNewParamsFrameImage) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParamsFrameImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParamsFrameImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VideoNewParamsFrameImageFrameUnion struct {
	OfFloat param.Opt[float64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfVideoNewsFrameImageFrameString)
	OfVideoNewsFrameImageFrameString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u VideoNewParamsFrameImageFrameUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfVideoNewsFrameImageFrameString)
}
func (u *VideoNewParamsFrameImageFrameUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VideoNewParamsFrameImageFrameUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfVideoNewsFrameImageFrameString) {
		return &u.OfVideoNewsFrameImageFrameString
	}
	return nil
}

type VideoNewParamsFrameImageFrameString string

const (
	VideoNewParamsFrameImageFrameStringFirst VideoNewParamsFrameImageFrameString = "first"
	VideoNewParamsFrameImageFrameStringLast  VideoNewParamsFrameImageFrameString = "last"
)

// Specifies the format of the output video. Defaults to MP4.
type VideoNewParamsOutputFormat string

const (
	VideoNewParamsOutputFormatMP4  VideoNewParamsOutputFormat = "MP4"
	VideoNewParamsOutputFormatWebm VideoNewParamsOutputFormat = "WEBM"
)
