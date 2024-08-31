package model

type Memo struct {
	// ID is the system generated unique identifier for the memo.
	ID int32
	// UID is the user defined unique identifier for the memo.
	UID string

	// Standard fields
	RowStatus RowStatus
	CreatorID int32
	CreatedTs int64 `gorm:"autoCreateTime"`
	UpdatedTs int64 `gorm:"autoUpdateTime"`

	// Domain specific fields
	Content    string
	Visibility Visibility
	Payload    *MemoPayload `gorm:"serializer:json"`

	// Composed fields
	Pinned   bool  `gorm:"-"`
	ParentID int32 `gorm:"-"`
}

func (Memo) TableName() string {
	return TableMemo
}

type MemoPayload struct {
	// property is the memo's property.
	Property *MemoPayloadProperty `json:"property"`
}

type MemoPayloadProperty struct {
	Tags               []string `json:"tags"`
	HasLink            bool     `json:"has_link"`
	HasTaskList        bool     `json:"Has_task_list"`
	HasCode            bool     `json:"has_code"`
	HasIncompleteTasks bool     `json:"has_incomplete_tasks"`
}

type FindMemoRequest struct {
	Id              int32
	UID             string
	CreatorId       int32
	RoeStatus       string
	CreatedTsBefore int64
	CreatedTsAfter  int64
	UpdatedTsBefore int64
	UpdatedTsAfter  int64
	ContentSearch   []string
	VisibilityList  []Visibility
	ExcludeContent  bool
	ExcludeComments bool
}

type ListMemosRequest struct {
	CreatorId       int32
	ExcludeComments bool
}

type UpdateMemoRequest struct {
	UpdateMasks []string
	UserId      int32
	ID          int32
	UID         string
	Content     string
	RowStatus   RowStatus
	Visibility  Visibility
	UpdatedTime int64
	CreatedTime int64
	DisplayTime int64
}

type GetMemoRequest struct {
	Id            int32
	UID           string
	CurrentUserId int32
}

type CreateMemoRequest struct {
	UserId     int32
	Content    string
	Visibility Visibility
}

type CreateMemoCommentRequest struct {
	ID      int32
	Comment *CreateMemoRequest
}
