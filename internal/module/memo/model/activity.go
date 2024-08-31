package model

type Activity struct {
	ID int32

	// Standard fields
	CreatorID int32
	CreatedTs int64 `gorm:"autoCreateTime"`

	// Domain specific fields
	Type    ActivityType
	Level   ActivityLevel
	Payload *ActivityPayload `gorm:"serializer:json"`
}

type ActivityPayload struct {
	MemoComment   *ActivityMemoCommentPayload   `json:"memo_comment"`
	VersionUpdate *ActivityVersionUpdatePayload `json:"version_update"`
}

type ActivityMemoCommentPayload struct {
	MemoId        int32 `json:"memo_id"`
	RelatedMemoId int32 `json:"related_memo_id"`
}

type ActivityVersionUpdatePayload struct {
	Version string `json:"version"`
}

type FindActivityRequest struct {
	ID   int32
	Type ActivityType
}
