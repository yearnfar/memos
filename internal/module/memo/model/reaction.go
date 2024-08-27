package model

type Reaction struct {
	ID        int32
	CreatedTs int64 `gorm:"autoCreateTime"`
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

type UpsertReactionRequest struct {
	CreatorID    int32
	ContentID    string
	ReactionType ReactionType
}

type ListReactionsRequest struct {
	Id        int
	CreatorId int
	ContentId string
}
