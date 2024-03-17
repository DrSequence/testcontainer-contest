package portfolio

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"testing"

	pm "testcontainer-contest/domain"
)

const (
	database       = "test"
	collectionName = "portfolio"
)

var mongoAddress string

func TestMain(m *testing.M) {
	ctx := context.Background()
	cfg := CreateCfg(database, collectionName)

	mongodbContainer, err := RunMongo(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	mappedPort, err := mongodbContainer.MappedPort(ctx, "27017")
	mongoAddress = "mongodb://localhost:" + mappedPort.Port()

	os.Exit(m.Run())
}

func TestFindByID(t *testing.T) {
	ctx := context.Background()
	cfg := CreateCfg(database, collectionName)

	cfg.Database.Address = mongoAddress

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

func TestShouldNptFindByIDW(t *testing.T) {
	ctx := context.Background()
	cfg := CreateCfg(database, collectionName)
	cfg.Database.Address = mongoAddress

	service, err := NewMongoPortfolioService(cfg)
	if err != nil {
		t.Fatal(err)
	}

	_, err = service.FindByID(ctx, "65f65ce1810a46fab20c69a5")
	assert.Error(t, err, "mongo: no documents in result")
}
