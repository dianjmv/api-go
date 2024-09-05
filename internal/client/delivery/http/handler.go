package http

import (
	"api-go/internal/client/usecase"
	"context"
	"encoding/json"
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
