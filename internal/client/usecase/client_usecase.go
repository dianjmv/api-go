package usecase

import (
	"api-go/internal/client"
	"context"
)

type ClientUsecase interface {
	GetClientByID(ctx context.Context, clientId string) (*client.Client, error)
}

type clientUsecase struct {
	clientRepository client.Repository
}

func NewClientUsecase(cr client.Repository) ClientUsecase {
	return &clientUsecase{
		clientRepository: cr,
	}
}

func (u *clientUsecase) GetClientByID(ctx context.Context, clientId string) (*client.Client, error) {
	return u.clientRepository.GetClientByID(ctx, clientId)
}
