# README for the Testcontainers-Enabled Service

## Service Description

This service is a REST API for managing portfolios, implemented in Go. MongoDB is used as the data store to save
portfolio information, and Redis is used for caching database queries, which helps to speed up data retrieval and reduce
the load on the main storage.

### Main functions of the service:

- `GET /api/v1/portfolio/{id}`: Retrieve a portfolio by its identifier.
- `POST /api/v1/save-portfolio`: Save or update a portfolio.
- `GET /api/v1/portfolios`: Retrieve a list of all portfolios with pagination.

## Technologies

- **Go** - Programming language for developing the service.
- **MongoDB** - Document-oriented database for storing portfolio data.
- **Redis** - In-memory database management system for caching.
- **Docker** and **Docker Compose** - For containerization and local deployment of the service and its dependencies.
- **Testcontainers-go** - Library for integration testing using Docker containers in Go tests.

## Configuration and Launch

### Requirements

Make sure you have:

- Docker
- Docker Compose
- Go (version 1.16 or higher)

### Launch Dependencies using Docker Compose

In the project's root, there is a `docker-compose.yml` file describing MongoDB and Redis containers needed for the
service. To start these containers, run the following command:

```bash
docker-compose up -d
```

This will create and start the `mongodb` and `redis` containers in the background.

### Launch the Service

After starting the dependencies, you can run the service using standard Go commands:

```bash
go build -o portfolio_service .
./portfolio_service
```

## Testing using Testcontainers

Tests with Testcontainers allow integration testing of the service under conditions closely resembling a real
environment, using Docker containers for dependencies.

```go
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
```

### Running Tests

To run integration tests, execute:

```bash
go test ./... -v
```

This command automatically starts the necessary Docker containers for the tests, runs the tests, and stops the
containers.

## References

- [Docker Compose](https://docs.docker.com/compose/)
- [Testcontainers-go](https://www.testcontainers.org/languages/go/)

## Example `docker-compose.yml`

```yaml
# Use root/example as user/password credentials
version: '3.1'

services:

  mongo:
    image: mongo
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: example

  redis-cache:
    image: redis:latest
    ports:
      - "6379:6379"
    command: redis-server --save 20 1 --loglevel warning --requirepass dsf231dasd123w12ddxhyjtj
```
