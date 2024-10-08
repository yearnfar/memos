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
}

func (Memo) TableName() string {
	return TableMemo
}

type MemoInfo struct {
	Memo

	// Composed fields
	Pinned   bool
	ParentID int32
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

type ListMemosRequest struct {
	ID        int32
	UID       string
	CreatorID int32
	RowStatus RowStatus

	// Domain specific fields
	ContentSearch   []string
	VisibilityList  []Visibility
	PayloadFind     *FindMemoPayload
	ExcludeContent  bool
	ExcludeComments bool
	Random          bool

	CreatedTsAfter  int64
	CreatedTsBefore int64
	UpdatedTsAfter  int64
	UpdatedTsBefore int64

	// Pagination
	Limit  int
	Offset int

	// Ordering
	OrderByUpdatedTs bool
	OrderByPinned    bool
	OrderByTimeAsc   bool
}

type FindMemoPayload struct {
	Raw                string
	TagSearch          []string
	HasLink            bool
	HasTaskList        bool
	HasCode            bool
	HasIncompleteTasks bool
}

type UpdateMemoRequest struct {
	UpdateMasks []string
	UserId      int32
	ID          int32
	UID         string
	Content     string
	RowStatus   RowStatus
	Visibility  Visibility
	Pinned      bool
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
