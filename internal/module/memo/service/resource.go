package service

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/lithammer/shortuuid/v4"
	"github.com/pkg/errors"

	"github.com/yearnfar/memos/internal/config"
	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/util"
)

func (s *Service) CreateResource(ctx context.Context, req *model.CreateResourceRequest) (resource *model.Resource, err error) {
	storageSetting, err := s.getWorkspaceStorageSetting(ctx)
	if err != nil {
		return
	}
	var reference string
	if storageSetting.StorageType == model.StorageTypeLocal {
		filepathTemplate := "assets/{timestamp}_{filename}"
		if storageSetting.FilepathTemplate != "" {
			filepathTemplate = storageSetting.FilepathTemplate
		}

		internalPath := filepathTemplate
		if !strings.Contains(internalPath, "{filename}") {
			internalPath = filepath.Join(internalPath, "{filename}")
		}
		internalPath = replaceFilenameWithPathTemplate(internalPath, req.Filename)
		internalPath = filepath.ToSlash(internalPath)

		// Ensure the directory exists.
		osPath := filepath.FromSlash(internalPath)
		if !filepath.IsAbs(osPath) {
			fs := config.GetApp().FileSystem
			osPath = filepath.Join(fs.Path, osPath)
		}
		if err = s.dao.SaveLocalFile(ctx, osPath, req.Content); err != nil {
			return
		}
		reference = internalPath
	}

	resource = &model.Resource{
		UID:         shortuuid.New(),
		CreatorID:   req.UserId,
		Filename:    req.Filename,
		Type:        req.Type,
		Size:        req.Size,
		StorageType: model.ResourceStorageType(storageSetting.StorageType),
		Reference:   reference,
	}
	err = s.dao.CreateResource(ctx, resource)
	return
}

func (s *Service) ListResources(ctx context.Context, req *model.ListResourcesRequest) (list []*model.Resource, err error) {
	list, err = s.dao.FindResources(ctx, []string{"memo_id=?"}, []any{req.MemoID})
	return
}

func (s *Service) GetResource(ctx context.Context, req *model.GetResourceRequest) (rb *model.Resource, err error) {
	return
}

func (s *Service) DeleteResource(ctx context.Context, req *model.DeleteResourceRequest) (err error) {
	resource, err := s.dao.FindResourceByID(ctx, req.ID)
	if err != nil {
		return errors.Wrap(err, "failed to get resource")
	}
	if resource.StorageType == model.ResourceStorageTypeLocal {
		osPath := filepath.FromSlash(resource.Reference)
		if !filepath.IsAbs(osPath) {
			fs := config.GetApp().FileSystem
			osPath = filepath.Join(fs.Path, osPath)
		}
		if err = s.dao.RemoveLocalFile(ctx, osPath); err != nil {
			return
		}
	} else if resource.StorageType == model.ResourceStorageTypeS3 {
		// TODO
	}

	err = s.dao.DeleteResourceById(ctx, req.ID)
	return
}

func (s *Service) GetResourceBinary(ctx context.Context, req *model.GetResourceBinaryRequest) (rb *model.ResourceBinary, err error) {
	resource, err := s.dao.FindResourceByID(ctx, req.Id)
	if err != nil {
		err = errors.Errorf("failed to get resource: %v", err)
		return
	}
	// Check the related memo visibility.
	if resource.MemoID != 0 {
		var memo *model.MemoInfo
		memo, err = s.dao.FindMemoByID(ctx, resource.MemoID)
		if err != nil {
			err = errors.Errorf("failed to find memo by ID: %v", resource.MemoID)
			return
		}
		if memo.Visibility == model.Private && req.UserId != resource.CreatorID {
			err = errors.New("unauthorized access")
			return
		}
	}
	var blob []byte
	if resource.StorageType == model.ResourceStorageTypeLocal {
		fpath := filepath.FromSlash(resource.Reference)
		if !filepath.IsAbs(fpath) {
			fs := config.GetApp().FileSystem
			fpath = filepath.Join(fs.Path, fpath)
		}
		blob, err = s.dao.ReadLocalFile(ctx, fpath, resource.Filename)
		if err != nil {
			return
		}
	}
	rb = &model.ResourceBinary{
		Resource: resource,
		Blob:     blob,
	}
	return
}

var fileKeyPattern = regexp.MustCompile(`\{[a-z]{1,9}\}`)

func replaceFilenameWithPathTemplate(path, filename string) string {
	t := time.Now()
	path = fileKeyPattern.ReplaceAllStringFunc(path, func(s string) string {
		switch s {
		case "{filename}":
			return filename
		case "{timestamp}":
			return fmt.Sprintf("%d", t.Unix())
		case "{year}":
			return fmt.Sprintf("%d", t.Year())
		case "{month}":
			return fmt.Sprintf("%02d", t.Month())
		case "{day}":
			return fmt.Sprintf("%02d", t.Day())
		case "{hour}":
			return fmt.Sprintf("%02d", t.Hour())
		case "{minute}":
			return fmt.Sprintf("%02d", t.Minute())
		case "{second}":
			return fmt.Sprintf("%02d", t.Second())
		case "{uuid}":
			return util.GenUUID()
		}
		return s
	})
	return path
}
