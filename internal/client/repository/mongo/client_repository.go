package mongoRepository

import (
	"api-go/internal/client"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct {
	collection *mongo.Collection
}

func NewClientRepository(db *mongo.Database) *ClientRepository {
	return &ClientRepository{
		collection: db.Collection("client"),
	}
}

func (r *ClientRepository) GetClientByID(ctx context.Context, clientId string) (*client.Client, error) {
	var result client.Client
	err := r.collection.FindOne(ctx, bson.M{"clientId": clientId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
