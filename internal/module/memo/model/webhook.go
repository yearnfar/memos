package model

type Webhook struct {
	ID        int32
	CreatedTs int64 `gorm:"autoCreateTime"`
	UpdatedTs int64 `gorm:"autoCreateTime"`
	CreatorID int32
	RowStatus RowStatus
	Name      string
	URL       string
}

func (Webhook) TableName() string {
	return TableWebhook
}
