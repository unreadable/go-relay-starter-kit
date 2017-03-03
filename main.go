package main

import (
	"learn/data"
	"net/http"

	"github.com/krypton97/GraphiQL"

	"github.com/krypton97/HandleGraphQL"
	"github.com/playlyfe/go-graphql"
)

func main() {
	executor, err := graphql.NewExecutor(data.Schema, "Query", "", data.Resolvers)

	if err != nil {
		panic(err)
	}

	api := handler.New(&handler.Config{
		Executor: executor,
		Context:  "",
		Pretty:   true,
	})

	graphiql := graphiql.New("/graphql")
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", fs)
	http.Handle("/graphql", api)
	http.Handle("/graphiql", graphiql)

	http.ListenAndServe(":3000", nil)
}
