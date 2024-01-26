package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/fernandojr999/go-learning/domain"
	"github.com/fernandojr999/go-learning/usecase"
)

type MessageHandler struct {
	messageUsecase *usecase.MessageUsecase
}

func NewMessageHandler(messageUsecase *usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{
		messageUsecase: messageUsecase,
	}
}

func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var message domain.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.messageUsecase.SendMessage(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
