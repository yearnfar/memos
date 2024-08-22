package model

type MemoOrganizer struct {
	MemoID int32
	UserID int32
	Pinned bool
}

func (MemoOrganizer) TableName() string {
	return TableMemoOrganizer
}

type FindMemoOrganizersRequest struct{}
