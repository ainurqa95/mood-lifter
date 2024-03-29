package compliment

import (
	"context"
	"fmt"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/ainurqa95/mood-lifter/internal/repository"
	def "github.com/ainurqa95/mood-lifter/internal/service"
	"log"
)

var _ def.ComplimentService = (*complimentService)(nil)

type complimentService struct {
	complimentRepository repository.ComplimentRepository
}

func NewComplimentService(
	complimentRepository repository.ComplimentRepository,
) *complimentService {
	return &complimentService{
		complimentRepository: complimentRepository,
	}
}

func (s *complimentService) GetRandom(ctx context.Context) (*model.Compliment, error) {
	compliment, err := s.complimentRepository.GetRandom(ctx)
	if err != nil {
		log.Printf("ошибка получения комплиментов: %v\n", err)
		return nil, fmt.Errorf("ошибка получения комплиментов: %v\n", err)
	}

	return compliment, nil
}
