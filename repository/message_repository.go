package repository

import (
	"database/sql"

	"github.com/fernandojr999/go-learning/domain"
)

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepo{
		db: db,
	}
}

type MessageRepository interface {
	SendMessage(message *domain.Message) error
}

type messageRepo struct {
	db *sql.DB
}

func (r *messageRepo) SendMessage(message *domain.Message) error {
	_, err := r.db.Exec("INSERT INTO messages (user_id, message, to_id) VALUES ($1, $2, $3)", message.UserId, message.Message, message.To)
	return err
}
