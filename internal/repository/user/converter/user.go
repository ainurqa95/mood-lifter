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
	userName := ""
	if user.UserName != nil {
		userName = *user.UserName
	}
	name := ""
	if user.Name != nil {
		name = *user.Name
	}
	return model.UserInfo{
		ChatID:   user.ChatID,
		UserName: userName,
		Name:     name,
	}
}
