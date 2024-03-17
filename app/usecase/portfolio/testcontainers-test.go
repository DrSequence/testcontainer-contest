package portfolio

import (
	"context"
	"errors"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testcontainer-contest/config"
	"testing"
	"time"
)

const (
	mongoPort  = "27017"
	mongoImage = "mongo:6"
	listener   = "27017/tcp"
)

func RunMongo(ctx context.Context, cfg config.Config) (cnt testcontainers.Container, err error) {
	cnt, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
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
		return nil, errors.New(fmt.Sprintf("failed to start container: %s", err))
	}

	return cnt, nil
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

func CreateCfg(database, collectionName string) config.Config {
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
	return cfg
}
