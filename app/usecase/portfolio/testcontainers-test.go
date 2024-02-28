package portfolio

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testcontainer-contest/config"
	"testing"
)

const (
	mongoPort  = "27017"
	mongoImage = "mongo:6"
	listener   = "27017/tcp"
)

func RunMongo(ctx context.Context, t *testing.T, cfg config.Config) testcontainers.Container {
	mongodbContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        mongoImage,
			ExposedPorts: []string{listener},
			WaitingFor:   wait.ForListeningPort(mongoPort),
			Env: map[string]string{
				"MONGO_INITDB_ROOT_USERNAME": cfg.Database.Username,
				"MONGO_INITDB_ROOT_PASSWORD": cfg.Database.Password,
			},
		},
		Started: true,
	})
	if err != nil {
		t.Fatalf("failed to start container: %s", err)
	}

	return mongodbContainer
}

func GetClient(ctx context.Context, t *testing.T, cfg config.Config) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Database.Address), options.Client().SetAuth(options.Credential{
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
	}))
	if err != nil {
		t.Fatal(err)
	}

	return client
}
