package clientportfolio

import (
	"context"
)

type Repository interface {
	GetClientPortfolioByCode(ctx context.Context, customerCode string) (*ClientPortfolio, error)
}
