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

var compliments = []string{
	"%s ты самый лучший человек на планете!",
	"%s ты прекрасный человечек!",
	"%s у тебя все получится!",
	"%s улыбнись!",
}

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

func (s *complimentService) MassCreate(ctx context.Context) error {
	compliment, err := s.GetRandom(ctx)
	if compliment != nil {
		return nil
	}

	// TODO load compliments from source
	err = s.complimentRepository.MassCreate(ctx, compliments)
	if err != nil {
		log.Printf("ошибка создания комплиментов: %v\n", err)
		return fmt.Errorf("ошибка создания комплиментов: %v\n", err)
	}
	return nil
}

func (s *complimentService) GetRandom(ctx context.Context) (*model.Compliment, error) {
	compliment, err := s.complimentRepository.GetRandom(ctx)
	if err != nil {
		log.Printf("ошибка получения комплиментов: %v\n", err)
		return nil, fmt.Errorf("ошибка получения комплиментов: %v\n", err)
	}

	return compliment, nil
}
