package portfolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"testcontainer-contest/config"
	pm "testcontainer-contest/domain"
)

const (
	database       = "test"
	collectionName = "portfolio"
)

func TestFindByIDWithTestContainer(t *testing.T) {
	ctx := context.Background()

	cfg := config.Config{}
	cfg.Server.Port = "8080"
	cfg.Server.Host = "localhost"
	cfg.Database.Username = "root"
	cfg.Database.Password = "example"
	cfg.Database.Database = database
	cfg.Database.Collection = collectionName
	cfg.Cache.Address = "localhost:6379"
	cfg.Cache.Exp = 5 * time.Minute
	cfg.Cache.Pass = "cachepassword"

	mongodbContainer := RunMongo(ctx, t, cfg)
	// Clean up the container
	defer func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}()

	mappedPort, err := mongodbContainer.MappedPort(ctx, "27017")
	address := "mongodb://localhost:" + mappedPort.Port()
	cfg.Database.Address = address

	client := GetClient(ctx, t, cfg)
	defer client.Disconnect(ctx)

	collection := client.Database(database).Collection(collectionName)

	testPortfolio := pm.Portfolio{
		Name:    "John Doe",
		Details: "Software Developer",
	}
	insertResult, err := collection.InsertOne(ctx, testPortfolio)
	if err != nil {
		t.Fatal(err)
	}

	savedObjectID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("InsertedID is not an ObjectID")
	}

	service, err := NewMongoPortfolioService(cfg)
	if err != nil {
		t.Fatal(err)
	}

	foundPortfolio, err := service.FindByID(ctx, savedObjectID.Hex())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testPortfolio.Name, foundPortfolio.Name)
	assert.Equal(t, testPortfolio.Details, foundPortfolio.Details)
}
