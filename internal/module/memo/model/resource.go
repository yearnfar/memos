package model

import (
	"encoding/json"
)

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

type ResourceBinary struct {
	Resource *Resource
	Blob     []byte
}

type MemoResource struct {
	ID int32 `json:"id"`
	// The user defined id of the resource.
	Uid          string `json:"uid"`
	CreateTime   int64  `json:"create_time"`
	Filename     string `json:"filename"`
	Content      []byte `json:"content"`
	ExternalLink string `json:"external_link"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`

	MemeID int32
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

type SetMemoResourcesRequest struct {
	MemoID    int32
	Resources []*MemoResource
}

type CreateResourceRequest struct {
	UserId       int32
	Name         string `json:"name"`
	Uid          string `json:"uid"`
	CreateTime   int64  `json:"create_time"`
	Filename     string `json:"filename"`
	Content      []byte `json:"content"`
	ExternalLink string `json:"external_link"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`
	MemoID       int32  `json:"memo_id"`
}

type GetResourceBinaryRequest struct {
	Id     int32
	UserId int32
}

type GetResourceRequest struct {
	Id     int32
	UserId int32
}

type ListResourcesRequest struct {
	CreatorID int32
	MemoID    int32
}

type DeleteResourceRequest struct {
	ID     int32
	UserID int32
	MemoID int32
}

type UpdateResourceRequest struct {
	ID        int32
	MemoID    int32
	UpdatedTs int64
}
