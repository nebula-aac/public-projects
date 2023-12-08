Follow [Getting Started](https://gqlgen.com/getting-started/)

Initialize `gqlgen`
`go run github.com/99designs/gqlgen init`

Run the GraphQL server
`go run ./server.go`

Example queries and mutations

To register a ninja

```graphql
mutation registerNewNinja {
  registerNewNinja(input: {name: "Ninja1", rank: "black", fight: "", sneak: ""}) {
    message
  }
}
```

Output

```graphql
{
  "data": {
    "registerNewNinja": {
      "message": "Ninja registered successfully"
    }
  }
}
```

Query to use `findNinja`

```graphql
{
  findNinja(name: "Ninja") {
    name
    rank
    fight
    sneak
  }
}
```

Output

```graphql
{
  "data": {
    "findNinja": {
      "name": "Placeholder Ninja",
      "rank": "Unknown",
      "fight": "N/A",
      "sneak": "N/A"
    }
  }
}
```

Return all ninjas
```graphql
{
  returnAllNinjas{
    name
    rank
    fight
    sneak
  }
}
```

Output

```graphql
{
  "data": {
    "returnAllNinjas": [
      {
        "name": "Ninja 1",
        "rank": "Unknown",
        "fight": "N/A",
        "sneak": "N/A"
      },
      {
        "name": "Ninja 2",
        "rank": "Unknown",
        "fight": "N/A",
        "sneak": "N/A"
      }
    ]
  }
}
```