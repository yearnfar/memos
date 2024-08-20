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
	Payload    *PayloadProperty `gorm:"serializer:json"`

	// Composed fields
	Pinned   bool
	ParentID *int32
}

func (Memo) TableName() string {
	return TableMemo
}

type PayloadProperty struct {
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
