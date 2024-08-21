package model

type Memo struct {
	// ID is the system generated unique identifier for the memo.
	ID int32
	// UID is the user defined unique identifier for the memo.
	UID string

	// Standard fields
	RowStatus RowStatus
	CreatorID int32
	CreatedTs int64
	UpdatedTs int64

	// Domain specific fields
	Content    string
	Visibility Visibility
	Payload    *MemoPayload `gorm:"serializer:json"`

	// Composed fields
	Pinned   bool
	ParentID *int32
}

func (Memo) TableName() string {
	return TableMemo
}

type MemoPayload struct {
	// property is the memo's property.
	Property *MemoPayloadProperty
}

type MemoPayloadProperty struct {
	Tags               []string
	HasLink            bool
	HasTaskList        bool
	HasCode            bool
	HasIncompleteTasks bool
}

type FindMemosRequest struct {
}
type ListMemosRequest struct {
	CreatorId int
}

type CreateMemoRequest struct {
	UserId     int
	Content    string
	Visibility Visibility
}
