package mongo

import (
	"context"
	"github.com/dianjmv/api-go/internal/clientportfolio"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientPortfolioRepository struct {
	collection *mongo.Collection
}

func NewClientPortfolioRepository(db *mongo.Database) *ClientPortfolioRepository {
	return &ClientPortfolioRepository{
		collection: db.Collection("client-portfolio"),
	}
}

func (r *ClientPortfolioRepository) GetClientPortfolioByCode(ctx context.Context, customerCode string) (*clientportfolio.ClientPortfolio, error) {
	var result clientportfolio.ClientPortfolio
	err := r.collection.FindOne(ctx, bson.M{"customerCode": customerCode}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
