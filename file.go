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

// List the metadata for a single uploaded data file.
func (r *FileService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("files/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the metadata for all uploaded data files.
func (r *FileService) List(ctx context.Context, opts ...option.RequestOption) (res *FileList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a previously uploaded data file.
func (r *FileService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *FileDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("files/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the contents of a single uploaded data file.
func (r *FileService) Content(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/binary")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("files/%s/content", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a file with specified purpose, file name, and file type.
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *FileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileList struct {
	Data []FileResponse `json:"data,required"`
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
	FilePurposeFineTune       FilePurpose = "fine-tune"
	FilePurposeEval           FilePurpose = "eval"
	FilePurposeEvalSample     FilePurpose = "eval-sample"
	FilePurposeEvalOutput     FilePurpose = "eval-output"
	FilePurposeEvalSummary    FilePurpose = "eval-summary"
	FilePurposeBatchGenerated FilePurpose = "batch-generated"
	FilePurposeBatchAPI       FilePurpose = "batch-api"
)

type FileResponse struct {
	ID        string `json:"id,required"`
	Bytes     int64  `json:"bytes,required"`
	CreatedAt int64  `json:"created_at,required"`
	Filename  string `json:"filename,required"`
	// The type of the file
	//
	// Any of "csv", "jsonl", "parquet".
	FileType  FileType `json:"FileType,required"`
	LineCount int64    `json:"LineCount,required"`
	Object    string   `json:"object,required"`
	Processed bool     `json:"Processed,required"`
	// The purpose of the file
	//
	// Any of "fine-tune", "eval", "eval-sample", "eval-output", "eval-summary",
	// "batch-generated", "batch-api".
	Purpose FilePurpose `json:"purpose,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Bytes       respjson.Field
		CreatedAt   respjson.Field
		Filename    respjson.Field
		FileType    respjson.Field
		LineCount   respjson.Field
		Object      respjson.Field
		Processed   respjson.Field
		Purpose     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileResponse) RawJSON() string { return r.JSON.raw }
func (r *FileResponse) UnmarshalJSON(data []byte) error {
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
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// The name of the file being uploaded
	FileName string `json:"file_name,required"`
	// The purpose of the file
	//
	// Any of "fine-tune", "eval", "eval-sample", "eval-output", "eval-summary",
	// "batch-generated", "batch-api".
	Purpose FilePurpose `json:"purpose,omitzero,required"`
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
