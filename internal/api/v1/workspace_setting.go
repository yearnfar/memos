package v1

import (
	"context"
	"fmt"

	"github.com/yearnfar/memos/internal/api"
	memomod "github.com/yearnfar/memos/internal/module/memo"
	"github.com/yearnfar/memos/internal/module/memo/model"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkspaceSettingService struct {
	api.BaseService
	v1pb.UnimplementedWorkspaceSettingServiceServer
}

func (s *WorkspaceSettingService) GetWorkspaceSetting(ctx context.Context, request *v1pb.GetWorkspaceSettingRequest) (response *v1pb.WorkspaceSetting, err error) {
	key, err := api.ExtractWorkspaceSettingKeyFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid workspace setting name: %v", err)
	}
	wsSetting, err := memomod.GetWorkspaceSetting(ctx, &model.GetWorkspaceSettingRequest{Name: key})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get workspace setting: %v", err)
	}
	if wsSetting == nil {
		return nil, status.Errorf(codes.NotFound, "workspace setting not found")
	}
	// For storage setting, only host can get it.
	if wsSetting.Key == model.WorkspaceSettingKeyStorage {
		user, err := s.GetCurrentUser(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get current user: %v", err)
		}
		if user == nil || user.Role != usermodel.RoleHost {
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}
	}
	response = convertWorkspaceSettingFromStore(wsSetting)
	return
}

func (s *WorkspaceSettingService) SetWorkspaceSetting(ctx context.Context, request *v1pb.SetWorkspaceSettingRequest) (response *v1pb.WorkspaceSetting, err error) {
	key, _ := api.ExtractWorkspaceSettingKeyFromName(request.Setting.Name)
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get current user: %v", err)
	}
	if user.Role != usermodel.RoleHost {
		return nil, status.Errorf(codes.PermissionDenied, "permission denied")
	}
	req := &model.SetWorkspaceSettingRequest{Name: key}
	switch key {
	case string(model.WorkspaceSettingKeyBasic):
		// req.Basic = convertWorkspaceBasicSettingToStore(request.GetSetting().GetGeneralSetting())
	case string(model.WorkspaceSettingKeyGeneral):
		req.General = convertWorkspaceGeneralSettingToStore(request.GetSetting().GetGeneralSetting())
	case string(model.WorkspaceSettingKeyStorage):
		req.Storage = convertWorkspaceStorageSettingToStore(request.GetSetting().GetStorageSetting())
	case string(model.WorkspaceSettingKeyMemoRelated):
		req.MemoRelated = convertWorkspaceMemoRelatedSettingToStore(request.GetSetting().GetMemoRelatedSetting())
	default:
		err = status.Errorf(codes.PermissionDenied, "unsupported workspace setting key: %s", req.Name)
		return
	}
	settingCache, err := memomod.SetWorkspaceSetting(ctx, req)
	if err != nil {
		return
	}
	response = convertWorkspaceSettingFromStore(settingCache)
	return
}

func convertWorkspaceSettingFromStore(setting *model.WorkspaceSettingCache) *v1pb.WorkspaceSetting {
	workspaceSetting := &v1pb.WorkspaceSetting{
		Name: fmt.Sprintf("%s%s", api.WorkspaceSettingNamePrefix, setting.Key),
	}

	switch val := setting.Value.(type) {
	case *model.WorkspaceGeneralSetting:
		workspaceSetting.Value = &v1pb.WorkspaceSetting_GeneralSetting{
			GeneralSetting: convertWorkspaceGeneralSettingFromStore(val),
		}
	case *model.WorkspaceStorageSetting:
		workspaceSetting.Value = &v1pb.WorkspaceSetting_StorageSetting{
			StorageSetting: convertWorkspaceStorageSettingFromStore(val),
		}
	case *model.WorkspaceMemoRelatedSetting:
		workspaceSetting.Value = &v1pb.WorkspaceSetting_MemoRelatedSetting{
			MemoRelatedSetting: convertWorkspaceMemoRelatedSettingFromStore(val),
		}
	}
	return workspaceSetting
}

func convertWorkspaceGeneralSettingFromStore(setting *model.WorkspaceGeneralSetting) *v1pb.WorkspaceGeneralSetting {
	if setting == nil {
		return nil
	}
	generalSetting := &v1pb.WorkspaceGeneralSetting{
		AdditionalScript: setting.AdditionalScript,
		AdditionalStyle:  setting.AdditionalStyle,
	}
	if setting.CustomProfile != nil {
		generalSetting.CustomProfile = &v1pb.WorkspaceCustomProfile{
			Title:       setting.CustomProfile.Title,
			Description: setting.CustomProfile.Description,
			LogoUrl:     setting.CustomProfile.LogoUrl,
			Locale:      setting.CustomProfile.Locale,
			Appearance:  setting.CustomProfile.Appearance,
		}
	}
	return generalSetting
}

func convertWorkspaceGeneralSettingToStore(setting *v1pb.WorkspaceGeneralSetting) *model.WorkspaceGeneralSetting {
	if setting == nil {
		return nil
	}
	generalSetting := &model.WorkspaceGeneralSetting{
		AdditionalScript: setting.AdditionalScript,
		AdditionalStyle:  setting.AdditionalStyle,
	}
	if setting.CustomProfile != nil {
		generalSetting.CustomProfile = &model.WorkspaceCustomProfile{
			Title:       setting.CustomProfile.Title,
			Description: setting.CustomProfile.Description,
			LogoUrl:     setting.CustomProfile.LogoUrl,
			Locale:      setting.CustomProfile.Locale,
			Appearance:  setting.CustomProfile.Appearance,
		}
	}
	return generalSetting
}

func convertWorkspaceStorageSettingToStore(setting *v1pb.WorkspaceStorageSetting) *model.WorkspaceStorageSetting {
	if setting == nil {
		return nil
	}
	settingpb := &model.WorkspaceStorageSetting{
		StorageType:       model.StorageType(v1pb.WorkspaceStorageSetting_StorageType_name[int32(setting.StorageType)]),
		FilepathTemplate:  setting.FilepathTemplate,
		UploadSizeLimitMb: setting.UploadSizeLimitMb,
	}
	if setting.S3Config != nil {
		settingpb.S3Config = &model.StorageS3Config{
			AccessKeyId:     setting.S3Config.AccessKeyId,
			AccessKeySecret: setting.S3Config.AccessKeySecret,
			Endpoint:        setting.S3Config.Endpoint,
			Region:          setting.S3Config.Region,
			Bucket:          setting.S3Config.Bucket,
		}
	}
	return settingpb
}

func convertWorkspaceStorageSettingFromStore(settingCache *model.WorkspaceStorageSetting) *v1pb.WorkspaceStorageSetting {
	if settingCache == nil {
		return nil
	}
	setting := &v1pb.WorkspaceStorageSetting{
		StorageType:       v1pb.WorkspaceStorageSetting_StorageType(v1pb.WorkspaceStorageSetting_StorageType_value[string(settingCache.StorageType)]),
		FilepathTemplate:  settingCache.FilepathTemplate,
		UploadSizeLimitMb: settingCache.UploadSizeLimitMb,
	}
	if settingCache.S3Config != nil {
		setting.S3Config = &v1pb.WorkspaceStorageSetting_S3Config{
			AccessKeyId:     settingCache.S3Config.AccessKeyId,
			AccessKeySecret: settingCache.S3Config.AccessKeySecret,
			Endpoint:        settingCache.S3Config.Endpoint,
			Region:          settingCache.S3Config.Region,
			Bucket:          settingCache.S3Config.Bucket,
		}
	}
	return setting
}

func convertWorkspaceMemoRelatedSettingFromStore(setting *model.WorkspaceMemoRelatedSetting) *v1pb.WorkspaceMemoRelatedSetting {
	if setting == nil {
		return nil
	}
	return &v1pb.WorkspaceMemoRelatedSetting{
		DisallowPublicVisibility: setting.DisallowPublicVisibility,
		DisplayWithUpdateTime:    setting.DisplayWithUpdateTime,
		ContentLengthLimit:       setting.ContentLengthLimit,
		EnableAutoCompact:        setting.EnableAutoCompact,
		EnableDoubleClickEdit:    setting.EnableDoubleClickEdit,
		EnableLinkPreview:        setting.EnableLinkPreview,
	}
}

func convertWorkspaceMemoRelatedSettingToStore(setting *v1pb.WorkspaceMemoRelatedSetting) *model.WorkspaceMemoRelatedSetting {
	if setting == nil {
		return nil
	}
	return &model.WorkspaceMemoRelatedSetting{
		DisallowPublicVisibility: setting.DisallowPublicVisibility,
		DisplayWithUpdateTime:    setting.DisplayWithUpdateTime,
		ContentLengthLimit:       setting.ContentLengthLimit,
		EnableAutoCompact:        setting.EnableAutoCompact,
		EnableDoubleClickEdit:    setting.EnableDoubleClickEdit,
		EnableLinkPreview:        setting.EnableLinkPreview,
	}
}
