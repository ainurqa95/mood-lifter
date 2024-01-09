package user

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
	def "github.com/ainurqa95/mood-lifter/internal/repository"
	"github.com/ainurqa95/mood-lifter/internal/repository/user/converter"
	repoModel "github.com/ainurqa95/mood-lifter/internal/repository/user/model"
	"sync"
	"time"
)

var _ def.UserRepository = (*mapRepository)(nil)

type mapRepository struct {
	data map[string]*repoModel.User
	m    sync.RWMutex
}

func NewRepository() *mapRepository {
	return &mapRepository{
		data: make(map[string]*repoModel.User),
	}
}

func (r *mapRepository) Create(_ context.Context, userUUID string, info *model.UserInfo) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.data[userUUID] = &repoModel.User{
		UUID:      userUUID,
		Info:      converter.ToUserInfoFromService(info),
		CreatedAt: time.Now(),
	}

	return nil
}

func (r *mapRepository) Get(_ context.Context, uuid string) (*model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	user, ok := r.data[uuid]
	if !ok {
		return nil, nil
	}

	return converter.ToUserFromRepo(user), nil
}

func (r *mapRepository) GetByChatId(ctx context.Context, botId int) (*model.User, error) {
	return nil, nil
}
