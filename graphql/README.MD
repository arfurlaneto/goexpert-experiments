https://gqlgen.com/


Update code when there is a schema change:

```bash
go run github.com/99designs/gqlgen generate
```

Sample queries:

```graphql
query listTodos {
  todos {
    id
    text
    user {
      id
      name
    }
	}
}

mutation createTodo {
  createTodo(input: { text: "Fazer algo", userId: "1"}) {
    id
  }
}
```
