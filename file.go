// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/togethercomputer/together-go/internal/apiform"
	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
	"github.com/togethercomputer/together-go/packages/respjson"
	"github.com/togethercomputer/together-go/shared/constant"
)

// FileService contains methods and other services that help with interacting with
// the together API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	return
}

// Retrieve the metadata for a single uploaded data file.
func (r *FileService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the metadata for all uploaded data files.
func (r *FileService) List(ctx context.Context, opts ...option.RequestOption) (res *FileList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a previously uploaded data file.
func (r *FileService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *FileDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get the contents of a single uploaded data file.
func (r *FileService) Content(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/binary")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s/content", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Upload a file with specified purpose, file name, and file type.
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *FileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FileList struct {
	Data []FileResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileList) RawJSON() string { return r.JSON.raw }
func (r *FileList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The purpose of the file
type FilePurpose string

const (
	FilePurposeFineTune FilePurpose = "fine-tune"
	FilePurposeEval     FilePurpose = "eval"
	FilePurposeBatchAPI FilePurpose = "batch-api"
)

// Structured information describing a file uploaded to Together.
type FileResponse struct {
	// ID of the file.
	ID string `json:"id" api:"required"`
	// The number of bytes in the file.
	Bytes int64 `json:"bytes" api:"required"`
	// The timestamp when the file was created.
	CreatedAt int64 `json:"created_at" api:"required"`
	// The name of the file as it was uploaded.
	Filename string `json:"filename" api:"required"`
	// The type of the file such as `jsonl`, `csv`, or `parquet`.
	//
	// Any of "csv", "jsonl", "parquet".
	FileType FileType `json:"FileType" api:"required"`
	// The object type, which is always `file`.
	Object constant.File `json:"object" default:"file"`
	// Deprecated. Whether file has been fully uploaded.
	//
	// Deprecated: deprecated
	Processed bool `json:"Processed" api:"required"`
	// The purpose of the file as it was uploaded.
	//
	// Any of "fine-tune", "eval", "batch-api".
	Purpose FilePurpose `json:"purpose" api:"required"`
	// Lifecycle state of the file validation pipeline. Files for non-`fine-tune`
	// purposes skip validation.
	//
	// Any of "PENDING", "QUEUED", "RUNNING", "COMPLETED", "FAILED", "INVALID_FORMAT".
	ProcessingStatus FileResponseProcessingStatus `json:"processing_status"`
	// Report produced by the file validation pipeline. Present once validation has
	// run; absent on files that bypassed validation (non-`fine-tune` purposes) or have
	// not yet been validated.
	ValidationReport FileResponseValidationReport `json:"validation_report"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Bytes            respjson.Field
		CreatedAt        respjson.Field
		Filename         respjson.Field
		FileType         respjson.Field
		Object           respjson.Field
		Processed        respjson.Field
		Purpose          respjson.Field
		ProcessingStatus respjson.Field
		ValidationReport respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileResponse) RawJSON() string { return r.JSON.raw }
func (r *FileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle state of the file validation pipeline. Files for non-`fine-tune`
// purposes skip validation.
type FileResponseProcessingStatus string

const (
	FileResponseProcessingStatusPending       FileResponseProcessingStatus = "PENDING"
	FileResponseProcessingStatusQueued        FileResponseProcessingStatus = "QUEUED"
	FileResponseProcessingStatusRunning       FileResponseProcessingStatus = "RUNNING"
	FileResponseProcessingStatusCompleted     FileResponseProcessingStatus = "COMPLETED"
	FileResponseProcessingStatusFailed        FileResponseProcessingStatus = "FAILED"
	FileResponseProcessingStatusInvalidFormat FileResponseProcessingStatus = "INVALID_FORMAT"
)

// Report produced by the file validation pipeline. Present once validation has
// run; absent on files that bypassed validation (non-`fine-tune` purposes) or have
// not yet been validated.
type FileResponseValidationReport struct {
	// Whether the file passed validation.
	Valid bool `json:"valid" api:"required"`
	// Detected dataset format (e.g. `CONVERSATION`, `INSTRUCTION`).
	DatasetFormat string `json:"dataset_format"`
	// Whether the dataset carries per-message weights (only possible for
	// `CONVERSATION` format).
	DatasetHasMessageWeights bool `json:"dataset_has_message_weights"`
	// Whether the dataset contains parallel tool-use messages.
	DatasetHasParallelToolCalls bool `json:"dataset_has_parallel_tool_calls"`
	// Whether the dataset contains reasoning content.
	DatasetHasReasoning bool `json:"dataset_has_reasoning"`
	// Whether the dataset carries per-sample weights.
	DatasetHasSampleWeights bool `json:"dataset_has_sample_weights"`
	// Whether the dataset contains tool-use messages.
	DatasetHasTools bool `json:"dataset_has_tools"`
	// Whether the dataset contains multimodal content.
	DatasetIsMultimodal bool `json:"dataset_is_multimodal"`
	// Human-readable validation error message. Only present when `error_type` is set
	// (i.e. user-correctable failures).
	Error string `json:"error"`
	// Category of validation failure.
	//
	// Any of "INVALID_FORMAT".
	ErrorType string `json:"error_type"`
	// ID of the file this report describes.
	FileID string `json:"file_id"`
	// Number of lines (records) in the dataset.
	Nlines int64 `json:"nlines"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Valid                       respjson.Field
		DatasetFormat               respjson.Field
		DatasetHasMessageWeights    respjson.Field
		DatasetHasParallelToolCalls respjson.Field
		DatasetHasReasoning         respjson.Field
		DatasetHasSampleWeights     respjson.Field
		DatasetHasTools             respjson.Field
		DatasetIsMultimodal         respjson.Field
		Error                       respjson.Field
		ErrorType                   respjson.Field
		FileID                      respjson.Field
		Nlines                      respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileResponseValidationReport) RawJSON() string { return r.JSON.raw }
func (r *FileResponseValidationReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the file
type FileType string

const (
	FileTypeCsv     FileType = "csv"
	FileTypeJSONL   FileType = "jsonl"
	FileTypeParquet FileType = "parquet"
)

type FileDeleteResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadParams struct {
	// The content of the file being uploaded
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	// The name of the file being uploaded
	FileName string `json:"file_name" api:"required"`
	// The purpose of the file
	//
	// Any of "fine-tune", "eval", "batch-api".
	Purpose FilePurpose `json:"purpose,omitzero" api:"required"`
	// The type of the file
	//
	// Any of "csv", "jsonl", "parquet".
	FileType FileType `json:"file_type,omitzero"`
	paramObj
}

func (r FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
