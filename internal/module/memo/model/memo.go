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
	// Pinned   bool
	// ParentID *int32
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

type FindMemosRequest struct {
	Id              int32
	UID             int32
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
type FindMemoRequest struct {
	Id  int32
	UID string
}

type ListMemosRequest struct {
	CreatorId int32
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
