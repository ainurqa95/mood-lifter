package converter

import (
	"github.com/ainurqa95/mood-lifter/internal/model"
	repoModel "github.com/ainurqa95/mood-lifter/internal/repository/user/model"
	"time"
)

func ToUserFromRepo(user *repoModel.User) *model.User {
	var updatedAt *time.Time
	if user.UpdatedAt.Valid {
		updatedAt = &user.UpdatedAt.Time
	}

	return &model.User{
		UUID:      user.UUID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromRepo(info repoModel.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:     info.Name,
		ChatID:   info.ChatID,
		UserName: info.UserName,
	}
}

func ToUserInfoFromService(info *model.UserInfo) repoModel.UserInfo {
	return repoModel.UserInfo{
		Name:     info.Name,
		ChatID:   info.ChatID,
		UserName: info.UserName,
	}
}
