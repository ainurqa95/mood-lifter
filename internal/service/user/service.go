package user

import (
	"context"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/ainurqa95/mood-lifter/internal/repository"
	def "github.com/ainurqa95/mood-lifter/internal/service"
	"github.com/google/uuid"
	"log"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) *service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) CreateIfNotExists(ctx context.Context, info *model.UserInfo) (string, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("ошибка генерации user UUID: %v\n", err)
		return "", err
	}

	err = s.userRepository.Create(ctx, userUUID.String(), info)
	if err != nil {
		log.Printf("ошибка создания пользователя: %v\n", err)
		return "", err
	}

	return userUUID.String(), nil
}

func (s *service) GetUsersByOffset(ctx context.Context, limit int, offset int) ([]model.UserInfo, error) {
	users, err := s.userRepository.GetByLimitOffset(ctx, limit, offset)
	if err != nil {
		log.Printf("ошибка создания пользователя: %v\n", err)
		return nil, err
	}

	return users, nil
}
