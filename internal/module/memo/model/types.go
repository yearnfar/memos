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

type WorkspaceSettingKey string

const (
	WorkspaceSettingKeyUnspecified WorkspaceSettingKey = "WORKSPACE_SETTING_KEY_UNSPECIFIED"
	// BASIC is the key for basic settings.
	WorkspaceSettingKeyBasic WorkspaceSettingKey = "BASIC"
	// GENERAL is the key for general settings.
	WorkspaceSettingKeyGeneral WorkspaceSettingKey = "GENERAL"
	// STORAGE is the key for storage settings.
	WorkspaceSettingKeyStorage WorkspaceSettingKey = "STORAGE"
	// MEMO_RELATED is the key for memo related settings.
	WorkspaceSettingKeyMemoRelated WorkspaceSettingKey = "MEMO_RELATED"
)

type StorageType string

const (
	StorageTypeUnspecified StorageType = "STORAGE_TYPE_UNSPECIFIED"
	// StorageTypeDatabase is the database storage type.
	StorageTypeDatabase StorageType = "DATABASE"
	// StorageTypeLocal is the local storage type.
	StorageTypeLocal StorageType = "LOCAL"
	// StorageTypeS3 is the S3 storage type.
	StorageTypeS3 StorageType = "S3"
)

const (
	TableInbox            = "inbox"
	TableMemo             = "memo"
	TableWorkspaceSetting = "system_setting"
)
