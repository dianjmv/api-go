package app

import (
	clientHttp "api-go/internal/client/delivery/http"
	"api-go/internal/client/repository/mongo"
	"api-go/internal/client/usecase"
	portfolioHttp "api-go/internal/clientportfolio/delivery/http"
	portfolioMongo "api-go/internal/clientportfolio/repository/mongo"
	portfolioUsecase "api-go/internal/clientportfolio/usecase"
	"context"
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

	clientRepo := mongoRepository.NewClientRepository(client.Database("test"))
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
