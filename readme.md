# Go GraphQL

## Install

```bash
go install
```

## Run

```bash
docker-compose up -d
```

## Queries

### Get Products

```graphql
query {
    products {
        id,
        name,
        description,
        quantity,
        status
    }
}
```

### Create Product

```graphql
mutation {
    createProduct(
        name: "My Product"
        description: "This is My Product",
        quantity: 1,
        status: true
    ) {
        id,
        name,
        description,
        quantity,
        status
    }
}
```

### Get Product

```graphql
query {
    product(
        id: "5f2411dfd5b8b1928e356a1b"
    ) {
        id,
        name,
        description,
        quantity,
        status
    }
}
```

### Update Product

```graphql
mutation {
    updateProduct(
      	id: "5f2411dfd5b8b1928e356a1b",
        name: "My Product",
        description: "This is My Product",
        quantity: 0,
        status: true
    ) {
        id,
        name,
        description,
        quantity,
        status
    }
}
```

## Build

```bash
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .build/app
```

Build Docker image

```bash
docker build -t bramalho/go-graphql -f .docker/prod/Dockerfile .
```

Run Docker image

```bash
docker run -p 80:80 --env APP_PORT=80 bramalho/go-graphql
```

Publish Docker image

```basg
docker login
docker push bramalho/go-graphql
```

The image will be published in your [docker hub](https://hub.docker.com/r/bramalho/go-graphql).

## Infrastructure
