package model

type Reaction struct {
	ID        int32
	CreatedTs int64
	CreatorID int32
	// ContentID is the id of the content that the reaction is for.
	// This can be a memo. e.g. memos/101
	ContentID    string
	ReactionType ReactionType
}

func (Reaction) TableName() string {
	return TableReaction
}

type FindReactionsRequest struct {
	Id        int
	CreatorId int
	ContentId string
}

type ListReactionsRequest struct {
	Id        int
	CreatorId int
	ContentId string
}
