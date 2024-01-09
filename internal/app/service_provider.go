package app

import (
	"github.com/ainurqa95/mood-lifter/internal/config"
	"github.com/ainurqa95/mood-lifter/internal/repository"
	complimentRepository "github.com/ainurqa95/mood-lifter/internal/repository/compliment"
	userRepository "github.com/ainurqa95/mood-lifter/internal/repository/user"
	"github.com/ainurqa95/mood-lifter/internal/service"
	complimentService "github.com/ainurqa95/mood-lifter/internal/service/compliment"
	userService "github.com/ainurqa95/mood-lifter/internal/service/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	cfg                  config.Config
	pool                 *pgxpool.Pool
	userRepository       repository.UserRepository
	complimentRepository repository.ComplimentRepository

	userService       service.UserService
	complimentService service.ComplimentService
}

func newServiceProvider(cfg config.Config, pool *pgxpool.Pool) *serviceProvider {
	return &serviceProvider{cfg: cfg, pool: pool}
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewDbUserRepository(s.pool)
	}

	return s.userRepository
}

func (s *serviceProvider) ComplimentService() service.ComplimentService {
	if s.complimentService == nil {
		s.complimentService = complimentService.NewComplimentService(
			s.ComplimentRepository(),
		)
	}

	return s.complimentService
}

func (s *serviceProvider) ComplimentRepository() repository.ComplimentRepository {
	if s.complimentRepository == nil {
		s.complimentRepository = complimentRepository.NewDbComplimentRepository(s.pool)
	}

	return s.complimentRepository
}
