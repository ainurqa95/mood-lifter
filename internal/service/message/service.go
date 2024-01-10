package message

import (
	"context"
	"fmt"
	"github.com/ainurqa95/mood-lifter/internal/model"
	"github.com/ainurqa95/mood-lifter/internal/repository"
	"log"
)

type messageService struct {
	messageRepository repository.MessageRepository
}

func NewMessageService(repository repository.MessageRepository) *messageService {
	return &messageService{
		messageRepository: repository,
	}
}

func (m *messageService) Create(ctx context.Context, chatId int64, text string) error {
	err := m.messageRepository.Create(ctx, model.Message{
		ChatId: chatId,
		Text:   text,
	})
	if err != nil {
		log.Printf("ошибка при добавлении message: %v", err)
		return fmt.Errorf("ошибка при добавлении message: %v", err)
	}
	return nil
}
