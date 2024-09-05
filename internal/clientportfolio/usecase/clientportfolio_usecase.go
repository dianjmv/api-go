package usecase

import (
	"context"
	"github.com/dianjmv/api-go/internal/clientportfolio"
)

type ClientPortfolioUsecase interface {
	GetClientPortfolioByCode(ctx context.Context, customerCode string) (*clientportfolio.ClientPortfolio, error)
}

type clientPortfolioUsecase struct {
	clientPortfolioRepository clientportfolio.Repository
}

func NewClientPortfolioUsecase(cr clientportfolio.Repository) ClientPortfolioUsecase {
	return &clientPortfolioUsecase{
		clientPortfolioRepository: cr,
	}
}

func (u *clientPortfolioUsecase) GetClientPortfolioByCode(ctx context.Context, customerCode string) (*clientportfolio.ClientPortfolio, error) {
	return u.clientPortfolioRepository.GetClientPortfolioByCode(ctx, customerCode)
}
