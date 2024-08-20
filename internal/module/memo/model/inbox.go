package model

type Inbox struct {
	ID         int
	CreatedTs  int64
	SenderID   int
	ReceiverID int
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

type FindInboxesRequest struct {
	Id         int
	SenderId   int
	ReceiverId int
	Status     InboxStatus
}

type ListInboxesRequest struct {
	ReceiverId int
}
