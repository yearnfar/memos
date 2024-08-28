package service

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/lithammer/shortuuid/v4"
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
