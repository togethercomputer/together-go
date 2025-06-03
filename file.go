// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package together

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/togethercomputer/together-go/internal/apijson"
	"github.com/togethercomputer/together-go/internal/requestconfig"
	"github.com/togethercomputer/together-go/option"
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
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// List the metadata for a single uploaded data file.
func (r *FileService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("files/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the metadata for all uploaded data files.
func (r *FileService) List(ctx context.Context, opts ...option.RequestOption) (res *FileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a previously uploaded data file.
func (r *FileService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *FileDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/binary")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("files/%s/content", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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

func (r FilePurpose) IsKnown() bool {
	switch r {
	case FilePurposeFineTune, FilePurposeEval, FilePurposeEvalSample, FilePurposeEvalOutput, FilePurposeEvalSummary, FilePurposeBatchGenerated, FilePurposeBatchAPI:
		return true
	}
	return false
}

// The type of the file
type FileType string

const (
	FileTypeCsv     FileType = "csv"
	FileTypeJSONL   FileType = "jsonl"
	FileTypeParquet FileType = "parquet"
)

func (r FileType) IsKnown() bool {
	switch r {
	case FileTypeCsv, FileTypeJSONL, FileTypeParquet:
		return true
	}
	return false
}

type FileGetResponse struct {
	ID        string `json:"id,required"`
	Bytes     int64  `json:"bytes,required"`
	CreatedAt int64  `json:"created_at,required"`
	Filename  string `json:"filename,required"`
	// The type of the file
	FileType  FileType `json:"FileType,required"`
	LineCount int64    `json:"LineCount,required"`
	Object    string   `json:"object,required"`
	Processed bool     `json:"Processed,required"`
	// The purpose of the file
	Purpose FilePurpose         `json:"purpose,required"`
	JSON    fileGetResponseJSON `json:"-"`
}

// fileGetResponseJSON contains the JSON metadata for the struct [FileGetResponse]
type fileGetResponseJSON struct {
	ID          apijson.Field
	Bytes       apijson.Field
	CreatedAt   apijson.Field
	Filename    apijson.Field
	FileType    apijson.Field
	LineCount   apijson.Field
	Object      apijson.Field
	Processed   apijson.Field
	Purpose     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileGetResponseJSON) RawJSON() string {
	return r.raw
}

type FileListResponse struct {
	Data []FileListResponseData `json:"data,required"`
	JSON fileListResponseJSON   `json:"-"`
}

// fileListResponseJSON contains the JSON metadata for the struct
// [FileListResponse]
type fileListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileListResponseJSON) RawJSON() string {
	return r.raw
}

type FileListResponseData struct {
	ID        string `json:"id,required"`
	Bytes     int64  `json:"bytes,required"`
	CreatedAt int64  `json:"created_at,required"`
	Filename  string `json:"filename,required"`
	// The type of the file
	FileType  FileType `json:"FileType,required"`
	LineCount int64    `json:"LineCount,required"`
	Object    string   `json:"object,required"`
	Processed bool     `json:"Processed,required"`
	// The purpose of the file
	Purpose FilePurpose              `json:"purpose,required"`
	JSON    fileListResponseDataJSON `json:"-"`
}

// fileListResponseDataJSON contains the JSON metadata for the struct
// [FileListResponseData]
type fileListResponseDataJSON struct {
	ID          apijson.Field
	Bytes       apijson.Field
	CreatedAt   apijson.Field
	Filename    apijson.Field
	FileType    apijson.Field
	LineCount   apijson.Field
	Object      apijson.Field
	Processed   apijson.Field
	Purpose     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileListResponseDataJSON) RawJSON() string {
	return r.raw
}

type FileDeleteResponse struct {
	ID      string                 `json:"id"`
	Deleted bool                   `json:"deleted"`
	JSON    fileDeleteResponseJSON `json:"-"`
}

// fileDeleteResponseJSON contains the JSON metadata for the struct
// [FileDeleteResponse]
type fileDeleteResponseJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileDeleteResponseJSON) RawJSON() string {
	return r.raw
}
