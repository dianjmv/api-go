package http

import (
	"context"
	"encoding/json"
	"github.com/dianjmv/api-go/internal/client/usecase"
	"net/http"
)

type ClientHandler struct {
	clientUsecase usecase.ClientUsecase
}

func NewClientHandler(cu usecase.ClientUsecase) *ClientHandler {
	return &ClientHandler{
		clientUsecase: cu,
	}
}

func (h *ClientHandler) GetClientByID(w http.ResponseWriter, r *http.Request) {
	clientId := r.URL.Query().Get("clientId")
	client, err := h.clientUsecase.GetClientByID(context.Background(), clientId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(client)
}
