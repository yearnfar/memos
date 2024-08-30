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

type InboxMsgType string

const (
	InboxMsgTypeUnspecified   InboxMsgType = "TYPE_UNSPECIFIED"
	InboxMsgTypeMemoComment   InboxMsgType = "MEMO_COMMENT"
	InboxMsgTypeVersionUpdate InboxMsgType = "VERSION_UPDATE"
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

type ReactionType string

const (
	ReactionTypeUnspecified   ReactionType = "REACTION_TYPE_UNSPECIFIED"
	ReactionTypeThumbsUp      ReactionType = "THUMBS_UP"
	ReactionTypeThumbsDown    ReactionType = "THUMBS_DOWN"
	ReactionTypeHeart         ReactionType = "HEART"
	ReactionTypeFire          ReactionType = "FIRE"
	ReactionTypeClappingHands ReactionType = "CLAPPING_HANDS"
	ReactionTypeLaugh         ReactionType = "LAUGH"
	ReactionTypeOkHand        ReactionType = "OK_HAND"
	ReactionTypeRocket        ReactionType = "ROCKET"
	ReactionTypeEyes          ReactionType = "EYES"
	ReactionTypeThinkingFace  ReactionType = "THINKING_FACE"
	ReactionTypeClownFace     ReactionType = "CLOWN_FACE"
	ReactionTypeQuestionMark  ReactionType = "QUESTION_MARK"
)

func (r ReactionType) String() string {
	return string(r)
}

type ResourceStorageType string

const (
	ResourceStorageTypeUnspecified ResourceStorageType = "UNSPECIFIED"
	ResourceStorageTypeLocal       ResourceStorageType = "LOCAL"
	ResourceStorageTypeS3          ResourceStorageType = "S3"
	ResourceStorageTypeExternal    ResourceStorageType = "EXTERNAL"
)

type ActivityType string

const (
	ActivityTypeMemoComment   ActivityType = "MEMO_COMMENT"
	ActivityTypeVersionUpdate ActivityType = "VERSION_UPDATE"
)

func (t ActivityType) String() string {
	return string(t)
}

type ActivityLevel string

const (
	ActivityLevelInfo ActivityLevel = "INFO"
)

func (l ActivityLevel) String() string {
	return string(l)
}

const (
	TableInbox            = "inbox"
	TableMemo             = "memo"
	TableMemoOrganizer    = "memo_organizer"
	TableMemoRelation     = "memo_relation"
	TableWorkspaceSetting = "system_setting"
	TableReaction         = "reaction"
	TableResource         = "resource"
)
