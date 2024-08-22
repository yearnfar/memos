package model

import "encoding/json"

type Resource struct {
	// ID is the system generated unique identifier for the resource.
	ID int32
	// UID is the user defined unique identifier for the resource.
	UID string

	// Standard fields
	CreatorID int32
	CreatedTs int64 `gorm:"autoCreateTime"`
	UpdatedTs int64 `gorm:"autoUpdateTime"`

	// Domain specific fields
	Filename    string
	Blob        []byte
	Type        string
	Size        int64
	StorageType ResourceStorageType
	Reference   string
	Payload     ResourcePayload `gorm:"serializer:json"`

	// The related memo ID.
	MemoID int32
}

func (Resource) TableName() string {
	return TableResource
}

type ResourcePayload struct {
	Payload *json.RawMessage `json:"payload"`
}

type ResourcePayloadS3Object struct {
	S3Config *StorageS3Config `json:"s3_config"`
	// key is the S3 object key.
	Key string `json:"key"`
	// last_presigned_time is the last time the object was presigned.
	// This is used to determine if the presigned URL is still valid.
	LastPresignedTime int64 `json:"last_presigned_time"`
}

type FindResourcesRequest struct {
}

type ListResourcesRequest struct {
}
