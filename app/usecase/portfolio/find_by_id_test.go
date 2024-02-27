package portfolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	pm "testcontainer-contest/entity"
)

const (
	database       = "test"
	collectionName = "portfolio"
)

func TestFindByIDWithTestContainer(t *testing.T) {
	ctx := context.Background()

	mongodbContainer := RunMongo(ctx, t)
	// Clean up the container
	defer func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}()

	mappedPort, err := mongodbContainer.MappedPort(ctx, "27017")
	address := "mongodb://localhost:" + mappedPort.Port()

	client := GetClient(ctx, t, address)
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

	service, err := NewMongoPortfolioService(address, database, collectionName)
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
