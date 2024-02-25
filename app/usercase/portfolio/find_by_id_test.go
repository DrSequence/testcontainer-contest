package portfolio

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pm "testcontainer-contest/entity"
)

const (
	database       = "test"
	collectionName = "portfolio"
)

func TestFindByIDWithTestContainer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	mongoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Could not start mongo: %s", err)
	}
	defer func() {
		if err := mongoContainer.Terminate(ctx); err != nil {
			log.Fatalf("Could not stop mongo: %s", err)
		}
	}()

	ip, err := mongoContainer.ContainerIP(ctx)
	if err != nil {
		log.Fatal(err)
	}
	address := "mongodb://" + ip + ":27017"

	// Инициализация MongoDB клиента и создание коллекции
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(address))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("test").Collection("portfolios")

	testPortfolio := pm.Portfolio{
		Name:    "John Doe",
		Details: "Software Developer",
	}
	_, err = collection.InsertOne(context.Background(), testPortfolio)
	if err != nil {
		t.Fatal(err)
	}

	service, err := NewMongoPortfolioService(address, database, collectionName)
	if err != nil {
		t.Fatal(err)
	}

	foundPortfolio, err := service.FindByID(context.Background(), testPortfolio.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testPortfolio.Name, foundPortfolio.Name)
	assert.Equal(t, testPortfolio.Details, foundPortfolio.Details)
}
