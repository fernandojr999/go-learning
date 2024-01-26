package usecase

import (
	"github.com/fernandojr999/go-learning/domain"
	"github.com/fernandojr999/go-learning/repository"
)

type MessageUsecase struct {
	messageRepository repository.MessageRepository
}

func NewMessageUsecase(messageRepository repository.MessageRepository) *MessageUsecase {
	return &MessageUsecase{
		messageRepository: messageRepository,
	}
}

func (m *MessageUsecase) SendMessage(message *domain.Message) error {

	return m.messageRepository.SendMessage(message)
}
