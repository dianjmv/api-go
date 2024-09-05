package http

import (
	"context"
	"encoding/json"
	"github.com/dianjmv/api-go/internal/clientportfolio/usecase"
	"net/http"
)

type ClientPortfolioHandler struct {
	clientPortfolioUsecase usecase.ClientPortfolioUsecase
}

func NewClientPortfolioHandler(cu usecase.ClientPortfolioUsecase) *ClientPortfolioHandler {
	return &ClientPortfolioHandler{
		clientPortfolioUsecase: cu,
	}
}

func (h *ClientPortfolioHandler) GetClientPortfolioByCode(w http.ResponseWriter, r *http.Request) {
	customerCode := r.URL.Query().Get("customerCode")
	portfolio, err := h.clientPortfolioUsecase.GetClientPortfolioByCode(context.Background(), customerCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(portfolio)
}
