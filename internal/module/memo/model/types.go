package model

// Visibility is the type of a visibility.
type Visibility string

const (
	// Public is the PUBLIC visibility.
	Public Visibility = "PUBLIC"
	// Protected is the PROTECTED visibility.
	Protected Visibility = "PROTECTED"
	// Private is the PRIVATE visibility.
	Private Visibility = "PRIVATE"
)

func (v Visibility) String() string {
	return string(v)
}

type InboxStatus string

const (
	InboxStatusUnspecified InboxStatus = "STATUS_UNSPECIFIED"
	InboxStatusUnread      InboxStatus = "UNREAD"
	InboxStatusArchived    InboxStatus = "ARCHIVED"
)

type InboxMsgType int32

const (
	InboxMsgTypeUnspecified   InboxMsgType = 0
	InboxMsgTypeMemoComment   InboxMsgType = 1
	InboxMsgTypeVersionUpdate InboxMsgType = 2
)

// RowStatus is the status for a row.
type RowStatus string

const (
	// Normal is the status for a normal row.
	Normal RowStatus = "NORMAL"
	// Archived is the status for an archived row.
	Archived RowStatus = "ARCHIVED"
)

func (r RowStatus) String() string {
	return string(r)
}

const (
	TableInbox = "inbox"
	TableMemo  = "memo"
)
