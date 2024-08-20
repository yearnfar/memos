package service

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/yearnfar/memos/internal/module/user/model"
	"github.com/yearnfar/memos/internal/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignUp(ctx context.Context, req *model.SignUpRequest) (user *model.User, err error) {
	if !util.UIDMatcher.MatchString(strings.ToLower(req.Username)) {
		err = errors.Errorf("invalid username: %s", req.Username)
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.Errorf("failed to generate password hash: %v", err)
		return
	}
	hostUserType := model.RoleHost
	existedHostUsers, err := s.dao.FindUsers(ctx, &model.FindUsersRequest{
		Role: hostUserType,
	})
	if err != nil {
		err = errors.Errorf("failed to list users, err: %s", err)
		return
	}
	role := model.RoleUser
	if len(existedHostUsers) == 0 {
		role = model.RoleHost
	}

	user = &model.User{
		Username:     req.Username,
		Role:         role,
		Nickname:     req.Username,
		PasswordHash: string(passwordHash),
		RowStatus:    model.Normal,
	}
	if err = s.dao.CreateUser(ctx, user); err != nil {
		err = errors.Errorf("failed to create user: %v", err)
		return
	}
	return user, nil
}

func (s *Service) CreateUser(ctx context.Context, req *model.CreateUserRequest) (user *model.User, err error) {
	if req.Role != model.RoleHost {
		err = errors.New("permission denied")
		return
	}
	if !util.UIDMatcher.MatchString(strings.ToLower(req.Username)) {
		err = errors.Errorf("invalid username: %s", req.Username)
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.Errorf("failed to generate password hash: %v", err)
		return
	}
	user = &model.User{
		Username:     req.Username,
		Role:         req.Role,
		Email:        req.Email,
		Nickname:     req.Nickname,
		PasswordHash: string(passwordHash),
	}
	if err = s.dao.CreateUser(ctx, user); err != nil {
		err = errors.Errorf("failed to create user: %v", err)
		return
	}
	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (user *model.User, err error) {
	user, err = s.dao.FindUserById(ctx, req.UserId)
	if err != nil {
		err = errors.Errorf("failed to get user: %v", err)
		return
	}
	if user == nil {
		err = errors.New("user not found")
		return
	}
	update := make(map[string]any)
	for _, field := range req.UpdateMasks {
		if field == "username" {
			if !util.UIDMatcher.MatchString(strings.ToLower(req.Username)) {
				err = errors.Errorf("invalid username: %s", req.Username)
				return
			}
			update["username"] = req.Username
		} else if field == "nickname" {
			update["nickname"] = req.Nickname
		} else if field == "email" {
			update["email"] = req.Email
		} else if field == "avatar_url" {
			update["avatar_url"] = req.AvatarURL
		} else if field == "description" {
			update["description"] = req.Description
		} else if field == "role" {
			update["role"] = req.Role
		} else if field == "password" {
			var passwordHash []byte
			passwordHash, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			if err != nil {
				err = errors.Errorf("failed to generate password hash: %v", err)
				return
			}
			update["password_hash"] = string(passwordHash)
		} else if field == "row_status" {
			update["row_status"] = req.RowStatus
		} else {
			err = errors.Errorf("invalid update path: %s", field)
			return
		}
	}
	err = s.dao.UpdateUser(ctx, user, update)
	if err != nil {
		err = errors.Errorf("failed to update user: %v", err)
		return
	}
	return nil, nil
}

func (s *Service) GetUserById(ctx context.Context, id int) (*model.User, error) {
	return s.dao.FindUserById(ctx, id)
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.dao.FindUserByUsername(ctx, username)
}
