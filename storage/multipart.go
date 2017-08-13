package storage

import (
	"context"
	"io"

	"github.com/joyent/triton-go/client"
)

type MultipartClient struct {
	client *client.Client
}

type UploadUUID string
type UploadState string
type PartETag string

type UploadHeaders struct {
	DurabilityLevel uint64
	ContentLength   string
	ContentMD5      string
	UserDefined     map[string]string
}

type CreateMultipartInput struct {
	ObjectPath string
}

// https://github.com/joyent/rfd/blob/master/rfd/0065/README.md#1-create-mpu
func (c *MultipartClient) Create(ctx context.Context, input *CreateMultipartInput) (UploadUUID, error) {
	// Generate a new upload uuid
	// Insert a directory record for /$ACCOUNT/uploads/[0-f]+/uuid with its state set to CREATED
	return UploadUUID{"not implemented"}, nil
}

type UploadPartInput struct {
	ID         UploadUUID
	PartNumber int
	PartData   io.ReadSeeker
}

// https://github.com/joyent/rfd/blob/master/rfd/0065/README.md#2-upload-part
func (c *MultipartClient) UploadPart(ctx context.Context, input *UploadPartInput) (UploadETag, error) {
	// Read the upload state in the object record for /$ACCOUNT/uploads/[0-f]+/uuidd.
	// - If the state is CREATED, insert object record for /$ACCOUNT/uploads/[0-f]+/uuid/N, where N is the part number.
	// - If the state is FINALIZING, return an error.
	return PartETag{"not implemented"}, nil
}

type CommitMultipartInput struct {
	ID          UploadUUID
	PartNumbers []int
	ETags       []PartETag
}

// https://github.com/joyent/rfd/blob/master/rfd/0065/README.md#3-commit
func (c *MultipartClient) Commit(ctx context.Context, input *CommitMultipartInput) error {
	return nil
}

type AbortMultipartInput struct {
	ID UploadUUID
}

// https://github.com/joyent/rfd/blob/master/rfd/0065/README.md#4-abort
func (c *MultipartClient) Abort(ctx context.Context, input *AbortMultipartInput) error {
	return nil
}

type ListPartsInput struct {
	ID UploadUUID
}

type ListPartOutput struct {
	ID         UploadUUID
	PartNumber int
	ETag       PartETag
	Size       int
}

// https://github.com/joyent/rfd/blob/master/rfd/0065/README.md#5-list-parts
func (c *MultipartClient) ListParts(ctx context.Context, input *ListPartsInput) ([]ListPartOutput, error) {
	return []ListPartOutput{}, nil
}

type ListMultipartInput struct {
	ID UploadUUID
}

type ListMultipartOutput struct {
	ID          UploadUUID
	ObjectPath  string
	UploadState string
}

// https://github.com/joyent/rfd/blob/master/rfd/0065/README.md#6-list-mpu
func (c *MultipartClient) List(ctx context.Context, input *ListMultipartInput) ([]ListMultipartInput, error) {
	return []ListMultipartInput{}, nil
}

type GetMultipartInput struct {
	ID UploadUUID
}

// https://github.com/joyent/rfd/blob/master/rfd/0065/README.md#7-get-mpu
func (c *MultipartClient) Get(ctx context.Context, input *GetMultipartInput) (UploadState, error) {
	return UploadState{"not implemented"}, nil
}
