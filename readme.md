# Go GraphQL

## Get Products

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

## Create Product

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

## Get Product

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

## Update Product

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
