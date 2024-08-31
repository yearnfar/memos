package model

type Inbox struct {
	ID         int
	CreatedTs  int64 `gorm:"autoCreateTime"`
	SenderID   int32
	ReceiverID int32
	Status     InboxStatus
	Message    *InboxMessage `gorm:"serializer:json"`
}

func (Inbox) TableName() string {
	return TableInbox
}

type InboxMessage struct {
	Type       InboxMsgType
	ActivityId int32
}

type FindInboxRequest struct {
	Id         int
	SenderId   int32
	ReceiverId int32
	Status     InboxStatus
}

type ListInboxesRequest struct {
	ReceiverId int32
}
