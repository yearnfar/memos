package model

type MemoRelationType string

const (
	// MemoRelationReference is the type for a reference memo relation.
	MemoRelationReference MemoRelationType = "REFERENCE"
	// MemoRelationComment is the type for a comment memo relation.
	MemoRelationComment MemoRelationType = "COMMENT"
)

type MemoRelation struct {
	MemoID        int32
	RelatedMemoID int32
	Type          MemoRelationType
}

func (MemoRelation) TableName() string {
	return TableMemoRelation
}

type FindMemoRelationsRequest struct {
	MemoId        int
	RelatedMemoId int
}

type ListMemoRelationsRequest struct {
	Id int
}

type DeleteMemoRelationsRequest struct {
	MemoID        int32
	RelatedMemoID int32
	Type          MemoRelationType
}

type SetMemoRelationsRequest struct {
	MemoID    int32
	Relations []*MemoRelation
}
