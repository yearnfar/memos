package service

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

const (
	// DefaultContentLengthLimit is the default limit of content length in bytes. 8KB.
	DefaultContentLengthLimit = 8 * 1024
)

func (s *Service) SetWorkspaceSetting(ctx context.Context, req *model.SetWorkspaceSettingRequest) (setting *model.WorkspaceSettingCache, err error) {
	var data []byte
	switch req.Name {
	case string(model.WorkspaceSettingKeyBasic):
		data, err = json.Marshal(req.Basic)
		if err != nil {
			return
		}
	case string(model.WorkspaceSettingKeyGeneral):
		data, err = json.Marshal(req.General)
		if err != nil {
			return
		}
	case string(model.WorkspaceSettingKeyStorage):
		data, err = json.Marshal(req.Storage)
		if err != nil {
			return
		}
	case string(model.WorkspaceSettingKeyMemoRelated):
		data, err = json.Marshal(req.MemoRelated)
		if err != nil {
			return
		}
	default:
		return nil, errors.Errorf("unsupported workspace setting key: %s", req.Name)
	}
	wsSetting := &model.WorkspaceSetting{
		Name:        req.Name,
		Value:       string(data),
		Description: req.Description,
	}
	if err = s.dao.UpsertWorkspaceSetting(ctx, wsSetting); err != nil {
		err = errors.Wrap(err, "Failed to upsert workspace setting")
		return
	}
	settingCache, err := s.convertWorkspaceSettingCache(wsSetting)
	if err != nil {
		err = errors.Wrap(err, "Failed to convert workspace setting")
		return
	}
	s.workspaceSettingCache.Store(settingCache.Key, settingCache)
	return settingCache, nil
}

func (s *Service) GetWorkspaceSetting(ctx context.Context, req *model.GetWorkspaceSettingRequest) (setting *model.WorkspaceSettingCache, err error) {
	return s.getWorkspaceSettingCache(ctx, model.WorkspaceSettingKey(req.Name))
}

func (s *Service) getWorkspaceStorageSetting(ctx context.Context) (*model.WorkspaceStorageSetting, error) {
	settingCache, err := s.getWorkspaceSettingCache(ctx, model.WorkspaceSettingKeyStorage)
	if err != nil {
		return nil, err
	}
	setting := &model.WorkspaceStorageSetting{
		StorageType: model.StorageTypeLocal,
	}
	if settingCache != nil {
		var ok bool
		if setting, ok = settingCache.Value.(*model.WorkspaceStorageSetting); !ok {
			return nil, errors.New("type error")
		}
	}
	return setting, nil
}

func (s *Service) getWorkspaceMemoRelatedSetting(ctx context.Context) (*model.WorkspaceMemoRelatedSetting, error) {
	settingCache, err := s.getWorkspaceSettingCache(ctx, model.WorkspaceSettingKeyMemoRelated)
	if err != nil {
		return nil, err
	}
	setting := &model.WorkspaceMemoRelatedSetting{}
	if settingCache != nil {
		var ok bool
		if setting, ok = settingCache.Value.(*model.WorkspaceMemoRelatedSetting); !ok {
			return nil, errors.New("type error")
		}
	}
	setting.ContentLengthLimit = max(setting.ContentLengthLimit, DefaultContentLengthLimit)
	return setting, nil

}

func (s *Service) getWorkspaceSettingCache(ctx context.Context, key model.WorkspaceSettingKey) (*model.WorkspaceSettingCache, error) {
	if cache, ok := s.workspaceSettingCache.Load(key); ok {
		workspaceSetting, ok := cache.(*model.WorkspaceSettingCache)
		if ok {
			return workspaceSetting, nil
		}
	}
	list, err := s.dao.FindWorkspaceSettings(ctx, &model.FindWorkspaceSettingsRequest{Name: string(key)})
	if err != nil {
		return nil, err
	}
	settingCaches := []*model.WorkspaceSettingCache{}
	for _, item := range list {
		settingCache, err := s.convertWorkspaceSettingCache(item)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to convert workspace setting")
		}
		if settingCache == nil {
			continue
		}
		s.workspaceSettingCache.Store(settingCache.Key, settingCache)
		settingCaches = append(settingCaches, settingCache)
	}
	if len(settingCaches) == 0 {
		return nil, nil
	}
	if len(settingCaches) > 1 {
		return nil, errors.Errorf("found multiple workspace settings with key %s", key)
	}
	return settingCaches[0], nil
}

func (s *Service) convertWorkspaceSettingCache(wsSetting *model.WorkspaceSetting) (settingCache *model.WorkspaceSettingCache, err error) {
	settingCache = &model.WorkspaceSettingCache{
		Key: model.WorkspaceSettingKey(wsSetting.Name),
	}
	switch wsSetting.Name {
	case string(model.WorkspaceSettingKeyBasic):
		var basicSetting model.WorkspaceBasicSetting
		if err = json.Unmarshal([]byte(wsSetting.Value), &basicSetting); err != nil {
			return
		}
		settingCache.Value = &basicSetting
	case string(model.WorkspaceSettingKeyGeneral):
		var generalSetting model.WorkspaceGeneralSetting
		if err = json.Unmarshal([]byte(wsSetting.Value), &generalSetting); err != nil {
			return
		}
		settingCache.Value = &generalSetting
	case string(model.WorkspaceSettingKeyStorage):
		var storageSetting model.WorkspaceStorageSetting
		if err = json.Unmarshal([]byte(wsSetting.Value), &storageSetting); err != nil {
			return
		}
		settingCache.Value = &storageSetting
	case string(model.WorkspaceSettingKeyMemoRelated):
		var relatedSetting model.WorkspaceMemoRelatedSetting
		if err = json.Unmarshal([]byte(wsSetting.Value), &relatedSetting); err != nil {
			return
		}
		settingCache.Value = &relatedSetting
	default:
		// Skip unsupported workspace setting key.
		return nil, nil
	}
	return settingCache, nil
}
