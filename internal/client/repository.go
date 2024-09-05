package client

import (
	"context"
)

type Repository interface {
	GetClientByID(ctx context.Context, clientId string) (*Client, error)
}
