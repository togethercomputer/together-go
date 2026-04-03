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
func (r *VideoService) New(ctx context.Context, body VideoNewParams, opts ...option.RequestOption) (res *VideoJob, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.xyz/v2/")}, opts...)
	path := "videos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch video metadata
func (r *VideoService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoJob, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.together.xyz/v2/")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("videos/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Structured information describing a generated video job.
type VideoJob struct {
	// Unique identifier for the video job.
	ID string `json:"id" api:"required"`
	// Unix timestamp (seconds) for when the job was created.
	CreatedAt float64 `json:"created_at" api:"required"`
	// The video generation model that produced the job.
	Model string `json:"model" api:"required"`
	// Duration of the generated clip in seconds.
	Seconds string `json:"seconds" api:"required"`
	// The resolution of the generated video.
	Size string `json:"size" api:"required"`
	// Current lifecycle status of the video job.
	//
	// Any of "in_progress", "completed", "failed".
	Status VideoJobStatus `json:"status" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Cost int64 `json:"cost" api:"required"`
	// URL hosting the generated video
	VideoURL string `json:"video_url" api:"required"`
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

type VideoNewParams struct {
	// The model to be used for the video creation request.
	Model string `json:"model" api:"required"`
	// Frames per second. Defaults to 24.
	Fps param.Opt[int64] `json:"fps,omitzero"`
	// Whether to generate audio for the video.
	GenerateAudio param.Opt[bool] `json:"generate_audio,omitzero"`
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
	// Aspect ratio of the video.
	Ratio param.Opt[string] `json:"ratio,omitzero"`
	// Video resolution.
	Resolution param.Opt[string] `json:"resolution,omitzero"`
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
	// Deprecated: use media.frame_images instead. Array of images to guide video
	// generation, similar to keyframes.
	FrameImages []VideoNewParamsFrameImage `json:"frame_images,omitzero"`
	// Media inputs for video generation. The accepted fields depend on the model type
	// (e.g. i2v, r2v, t2v, videoedit).
	Media VideoNewParamsMedia `json:"media,omitzero"`
	// Specifies the format of the output video. Defaults to MP4.
	//
	// Any of "MP4", "WEBM".
	OutputFormat VideoNewParamsOutputFormat `json:"output_format,omitzero"`
	// Deprecated: use media.reference_images instead. Unlike frame_images which
	// constrain specific timeline positions, reference images guide the general
	// appearance that should appear consistently across the video.
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
	InputImage string `json:"input_image" api:"required"`
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

// Media inputs for video generation. The accepted fields depend on the model type
// (e.g. i2v, r2v, t2v, videoedit).
type VideoNewParamsMedia struct {
	// Array of audio inputs.
	AudioInputs []VideoNewParamsMediaAudioInput `json:"audio_inputs,omitzero"`
	// Array of images to guide video generation at specific timeline positions.
	FrameImages []VideoNewParamsMediaFrameImage `json:"frame_images,omitzero"`
	// Array of video clips to use as starting clips.
	FrameVideos []VideoNewParamsMediaFrameVideo `json:"frame_videos,omitzero"`
	// Array of image URLs that guide the general appearance across the video.
	ReferenceImages []string `json:"reference_images,omitzero"`
	// Array of reference videos.
	ReferenceVideos []VideoNewParamsMediaReferenceVideo `json:"reference_videos,omitzero"`
	// Source video to edit.
	SourceVideo VideoNewParamsMediaSourceVideo `json:"source_video,omitzero"`
	paramObj
}

func (r VideoNewParamsMedia) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParamsMedia
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParamsMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Audio is required.
type VideoNewParamsMediaAudioInput struct {
	// URL of the audio.
	Audio string `json:"audio" api:"required"`
	paramObj
}

func (r VideoNewParamsMediaAudioInput) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParamsMediaAudioInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParamsMediaAudioInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property InputImage is required.
type VideoNewParamsMediaFrameImage struct {
	// URL path to hosted image that is used for a frame
	InputImage string `json:"input_image" api:"required"`
	// Optional param to specify where to insert the frame. If this is omitted, the
	// following heuristics are applied:
	//
	// - frame_images size is one, frame is first.
	// - If size is two, frames are first and last.
	// - If size is larger, frames are first, last and evenly spaced between.
	Frame VideoNewParamsMediaFrameImageFrameUnion `json:"frame,omitzero"`
	paramObj
}

func (r VideoNewParamsMediaFrameImage) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParamsMediaFrameImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParamsMediaFrameImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VideoNewParamsMediaFrameImageFrameUnion struct {
	OfFloat param.Opt[float64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfVideoNewsMediaFrameImageFrameString)
	OfVideoNewsMediaFrameImageFrameString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u VideoNewParamsMediaFrameImageFrameUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfVideoNewsMediaFrameImageFrameString)
}
func (u *VideoNewParamsMediaFrameImageFrameUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VideoNewParamsMediaFrameImageFrameUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfVideoNewsMediaFrameImageFrameString) {
		return &u.OfVideoNewsMediaFrameImageFrameString
	}
	return nil
}

type VideoNewParamsMediaFrameImageFrameString string

const (
	VideoNewParamsMediaFrameImageFrameStringFirst VideoNewParamsMediaFrameImageFrameString = "first"
	VideoNewParamsMediaFrameImageFrameStringLast  VideoNewParamsMediaFrameImageFrameString = "last"
)

// The property Video is required.
type VideoNewParamsMediaFrameVideo struct {
	// URL of the video.
	Video string `json:"video" api:"required"`
	paramObj
}

func (r VideoNewParamsMediaFrameVideo) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParamsMediaFrameVideo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParamsMediaFrameVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Video is required.
type VideoNewParamsMediaReferenceVideo struct {
	// URL of the video.
	Video string `json:"video" api:"required"`
	paramObj
}

func (r VideoNewParamsMediaReferenceVideo) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParamsMediaReferenceVideo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParamsMediaReferenceVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source video to edit.
//
// The property Video is required.
type VideoNewParamsMediaSourceVideo struct {
	// URL of the video.
	Video string `json:"video" api:"required"`
	paramObj
}

func (r VideoNewParamsMediaSourceVideo) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParamsMediaSourceVideo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParamsMediaSourceVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the format of the output video. Defaults to MP4.
type VideoNewParamsOutputFormat string

const (
	VideoNewParamsOutputFormatMP4  VideoNewParamsOutputFormat = "MP4"
	VideoNewParamsOutputFormatWebm VideoNewParamsOutputFormat = "WEBM"
)
