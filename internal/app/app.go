package app

import (
	"context"
	clientHttp "github.com/dianjmv/api-go/internal/client/delivery/http"
	"github.com/dianjmv/api-go/internal/client/repository/mongo"
	"github.com/dianjmv/api-go/internal/client/usecase"
	portfolioHttp "github.com/dianjmv/api-go/internal/clientportfolio/delivery/http"
	portfolioMongo "github.com/dianjmv/api-go/internal/clientportfolio/repository/mongo"
	portfolioUsecase "github.com/dianjmv/api-go/internal/clientportfolio/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	clientRepo := mongo.NewClientRepository(client.Database("clientDB"))
	clientUsecase := usecase.NewClientUsecase(clientRepo)
	clientHandler := clientHttp.NewClientHandler(clientUsecase)

	portfolioRepo := portfolioMongo.NewClientPortfolioRepository(client.Database("clientPortfolioDB"))
	portfolioUsecase := portfolioUsecase.NewClientPortfolioUsecase(portfolioRepo)
	portfolioHandler := portfolioHttp.NewClientPortfolioHandler(portfolioUsecase)

	http.HandleFunc("/client", clientHandler.GetClientByID)
	http.HandleFunc("/client-portfolio", portfolioHandler.GetClientPortfolioByCode)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
