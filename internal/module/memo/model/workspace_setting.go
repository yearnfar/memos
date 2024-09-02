package model

type WorkspaceSetting struct {
	Name        string
	Value       string
	Description string
}

func (WorkspaceSetting) TableName() string {
	return TableWorkspaceSetting
}

type WorkspaceSettingCache struct {
	Key   WorkspaceSettingKey
	Value any
}

type WorkspaceBasicSetting struct {
	SecretKey string `json:"secret_key"`
}

type WorkspaceGeneralSetting struct {
	// additional_script is the additional script.
	AdditionalScript string `json:"additional_script"`
	// additional_style is the additional style.
	AdditionalStyle string `json:"additional_style"`
	// custom_profile is the custom profile.
	CustomProfile *WorkspaceCustomProfile `json:"custom_profile"`
}

type WorkspaceCustomProfile struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LogoUrl     string `json:"logo_url"`
	Locale      string `json:"locale"`
	Appearance  string `json:"appearance"`
}

type WorkspaceStorageSetting struct {
	// storage_type is the storage type.
	StorageType StorageType `json:"storage_type"`
	// The template of file path.
	// e.g. assets/{timestamp}_{filename}
	FilepathTemplate string `json:"filepath_template"`
	// The max upload size in megabytes.
	UploadSizeLimitMb int64 `json:"upload_size_limit_mb"`
	// The S3 config.
	S3Config *StorageS3Config `json:"s3_config"`
}

type StorageS3Config struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Endpoint        string `json:"endpoint"`
	Region          string `json:"region"`
	Bucket          string `json:"bucket"`
}

type WorkspaceMemoRelatedSetting struct {
	// disallow_public_visibility disallows set memo as public visibility.
	DisallowPublicVisibility bool `json:"disallow_public_visibility"`
	// display_with_update_time orders and displays memo with update time.
	DisplayWithUpdateTime bool `json:"display_with_update_time"`
	// content_length_limit is the limit of content length. Unit is byte.
	ContentLengthLimit int32 `json:"content_length_limit"`
	// enable_auto_compact enables auto compact for large content.
	EnableAutoCompact bool `json:"enable_auto_compact"`
	// enable_double_click_edit enables editing on double click.
	EnableDoubleClickEdit bool `json:"enable_double_click_edit"`
	// enable_link_preview enables links preview.
	EnableLinkPreview bool `json:"enable_link_preview"`
}

type SetWorkspaceSettingRequest struct {
	Name        string
	Description string
	Basic       *WorkspaceBasicSetting
	General     *WorkspaceGeneralSetting
	Storage     *WorkspaceStorageSetting
	MemoRelated *WorkspaceMemoRelatedSetting
}

type GetWorkspaceSettingRequest struct {
	Name string
}
