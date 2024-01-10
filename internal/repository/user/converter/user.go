package converter

import (
	"github.com/ainurqa95/mood-lifter/internal/model"
	repoModel "github.com/ainurqa95/mood-lifter/internal/repository/user/model"
)

func ToUserFromRepo(user repoModel.User) *model.User {
	return &model.User{
		Id:        user.Id,
		UUID:      user.UUID,
		Info:      ToUserInfoFromRepo(user),
		CreatedAt: user.CreatedAt,
	}
}

func ToUserInfoFromRepo(user repoModel.User) model.UserInfo {
	return model.UserInfo{
		Name:     user.Name,
		ChatID:   user.ChatID,
		UserName: user.UserName,
	}
}

func ToUserInfoFromService(info *model.UserInfo) repoModel.UserInfo {
	return repoModel.UserInfo{
		Name:     info.Name,
		ChatID:   info.ChatID,
		UserName: info.UserName,
	}
}
