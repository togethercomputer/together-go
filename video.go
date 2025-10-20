// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/param"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
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
func NewVideoService(opts ...option.RequestOption) (r *VideoService) {
	r = &VideoService{}
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
func (r *VideoService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoGetResponse, err error) {
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

type VideoNewResponse struct {
	// Unique identifier for the video job.
	ID   string               `json:"id"`
	JSON videoNewResponseJSON `json:"-"`
}

// videoNewResponseJSON contains the JSON metadata for the struct
// [VideoNewResponse]
type videoNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoNewResponseJSON) RawJSON() string {
	return r.raw
}

// Structured information describing a generated video job.
type VideoGetResponse struct {
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
	Status VideoGetResponseStatus `json:"status,required"`
	// Unix timestamp (seconds) for when the job completed, if finished.
	CompletedAt float64 `json:"completed_at"`
	// Error payload that explains why generation failed, if applicable.
	Error VideoGetResponseError `json:"error"`
	// The object type, which is always video.
	Object VideoGetResponseObject `json:"object"`
	// Available upon completion, the outputs provides the cost charged and the hosted
	// url to access the video
	Outputs VideoGetResponseOutputs `json:"outputs"`
	JSON    videoGetResponseJSON    `json:"-"`
}

// videoGetResponseJSON contains the JSON metadata for the struct
// [VideoGetResponse]
type videoGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Model       apijson.Field
	Seconds     apijson.Field
	Size        apijson.Field
	Status      apijson.Field
	CompletedAt apijson.Field
	Error       apijson.Field
	Object      apijson.Field
	Outputs     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoGetResponseJSON) RawJSON() string {
	return r.raw
}

// Current lifecycle status of the video job.
type VideoGetResponseStatus string

const (
	VideoGetResponseStatusInProgress VideoGetResponseStatus = "in_progress"
	VideoGetResponseStatusCompleted  VideoGetResponseStatus = "completed"
	VideoGetResponseStatusFailed     VideoGetResponseStatus = "failed"
)

func (r VideoGetResponseStatus) IsKnown() bool {
	switch r {
	case VideoGetResponseStatusInProgress, VideoGetResponseStatusCompleted, VideoGetResponseStatusFailed:
		return true
	}
	return false
}

// Error payload that explains why generation failed, if applicable.
type VideoGetResponseError struct {
	Message string                    `json:"message,required"`
	Code    string                    `json:"code"`
	JSON    videoGetResponseErrorJSON `json:"-"`
}

// videoGetResponseErrorJSON contains the JSON metadata for the struct
// [VideoGetResponseError]
type videoGetResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

// The object type, which is always video.
type VideoGetResponseObject string

const (
	VideoGetResponseObjectVideo VideoGetResponseObject = "video"
)

func (r VideoGetResponseObject) IsKnown() bool {
	switch r {
	case VideoGetResponseObjectVideo:
		return true
	}
	return false
}

// Available upon completion, the outputs provides the cost charged and the hosted
// url to access the video
type VideoGetResponseOutputs struct {
	// The cost of generated video charged to the owners account.
	Cost int64 `json:"cost,required"`
	// URL hosting the generated video
	VideoURL string                      `json:"video_url,required"`
	JSON     videoGetResponseOutputsJSON `json:"-"`
}

// videoGetResponseOutputsJSON contains the JSON metadata for the struct
// [VideoGetResponseOutputs]
type videoGetResponseOutputsJSON struct {
	Cost        apijson.Field
	VideoURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VideoGetResponseOutputs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r videoGetResponseOutputsJSON) RawJSON() string {
	return r.raw
}

type VideoNewParams struct {
	// The model to be used for the video creation request.
	Model param.Field[string] `json:"model,required"`
	// Frames per second. Defaults to 24.
	Fps param.Field[int64] `json:"fps"`
	// Array of images to guide video generation, like keyframes. If size 1, starting
	// frame, if size 2, starting and ending frame, if more than 2 then frame must be
	// specified
	FrameImages param.Field[[]VideoNewParamsFrameImage] `json:"frame_images"`
	// Controls how closely the video generation follows your prompt. Higher values
	// make the model adhere more strictly to your text description, while lower values
	// allow more creative freedom. guidence_scale affects both visual content and
	// temporal consistency.Recommended range is 6.0-10.0 for most video models. Values
	// above 12 may cause over-guidance artifacts or unnatural motion patterns.
	GuidanceScale param.Field[int64] `json:"guidance_scale"`
	Height        param.Field[int64] `json:"height"`
	// Similar to prompt, but specifies what to avoid instead of what to include
	NegativePrompt param.Field[string] `json:"negative_prompt"`
	// Specifies the format of the output video. Defaults to MP4.
	OutputFormat param.Field[VideoNewParamsOutputFormat] `json:"output_format"`
	// Compression quality. Defaults to 20.
	OutputQuality param.Field[int64] `json:"output_quality"`
	// Text prompt that describes the video to generate.
	Prompt param.Field[string] `json:"prompt"`
	// TODO need to figure this out
	ReferenceImages param.Field[[]interface{}] `json:"reference_images"`
	// Clip duration in seconds.
	Seconds param.Field[string] `json:"seconds"`
	// Seed to use in initializing the video generation. Using the same seed allows
	// deterministic video generation. If not provided a random seed is generated for
	// each request.
	Seed param.Field[int64] `json:"seed"`
	// The number of denoising steps the model performs during video generation. More
	// steps typically result in higher quality output but require longer processing
	// time.
	Steps param.Field[int64] `json:"steps"`
	Width param.Field[int64] `json:"width"`
}

func (r VideoNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VideoNewParamsFrameImage struct {
	Frame param.Field[VideoNewParamsFrameImagesFrameUnion] `json:"frame,required"`
	// idk
	InputImage param.Field[string] `json:"input_image,required"`
}

func (r VideoNewParamsFrameImage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionFloat], [VideoNewParamsFrameImagesFrameString].
type VideoNewParamsFrameImagesFrameUnion interface {
	ImplementsVideoNewParamsFrameImagesFrameUnion()
}

type VideoNewParamsFrameImagesFrameString string

const (
	VideoNewParamsFrameImagesFrameStringFirst VideoNewParamsFrameImagesFrameString = "first"
	VideoNewParamsFrameImagesFrameStringLast  VideoNewParamsFrameImagesFrameString = "last"
)

func (r VideoNewParamsFrameImagesFrameString) IsKnown() bool {
	switch r {
	case VideoNewParamsFrameImagesFrameStringFirst, VideoNewParamsFrameImagesFrameStringLast:
		return true
	}
	return false
}

func (r VideoNewParamsFrameImagesFrameString) ImplementsVideoNewParamsFrameImagesFrameUnion() {}

// Specifies the format of the output video. Defaults to MP4.
type VideoNewParamsOutputFormat string

const (
	VideoNewParamsOutputFormatMP4  VideoNewParamsOutputFormat = "MP4"
	VideoNewParamsOutputFormatWebm VideoNewParamsOutputFormat = "WEBM"
)

func (r VideoNewParamsOutputFormat) IsKnown() bool {
	switch r {
	case VideoNewParamsOutputFormatMP4, VideoNewParamsOutputFormatWebm:
		return true
	}
	return false
}
