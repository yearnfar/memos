package service

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (s *Service) getWorkspaceMemoRelatedSetting(ctx context.Context) (*model.WorkspaceMemoRelatedSetting, error) {
	settingCache, err := s.getWorkspaceSettingCache(ctx, string(model.WorkspaceSettingKeyMemoRelated))
	if err != nil {
		return nil, err
	} else if settingCache == nil {
		return nil, nil
	}
	if val, ok := settingCache.Value.(*model.WorkspaceMemoRelatedSetting); !ok {
		return nil, errors.New("type error")
	} else {
		return val, nil
	}
}
func (s *Service) getWorkspaceSettingCache(ctx context.Context, name string) (*model.WorkspaceSettingCache, error) {
	if cache, ok := s.workspaceSettingCache.Load(name); ok {
		workspaceSetting, ok := cache.(*model.WorkspaceSettingCache)
		if ok {
			return workspaceSetting, nil
		}
	}
	list, err := s.dao.FindWorkspaceSettings(ctx, &model.FindWorkspaceSettingsRequest{Name: name})
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
		return nil, errors.Errorf("found multiple workspace settings with key %s", name)
	}
	return settingCaches[0], nil
}

func (s *Service) convertWorkspaceSettingCache(wsSetting *model.WorkspaceSetting) (settingCache *model.WorkspaceSettingCache, err error) {
	settingCache = &model.WorkspaceSettingCache{
		Key: wsSetting.Name,
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
		settingCache.Value = relatedSetting
	default:
		// Skip unsupported workspace setting key.
		return nil, nil
	}
	return settingCache, nil
}
