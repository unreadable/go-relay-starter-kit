package main

import (
	"flag"
	"io/ioutil"
	"go-relay-starter-kit/data"

	"os"

	"github.com/playlyfe/go-graphql"
)

func main() {
	flag.Parse()
	gql, err := graphql.NewGraphQL(&graphql.GraphQLParams{
		SchemaDefinition: data.Schema,
		QueryRoot:        "Query",
		MutationRoot:     "Mutation",
		Resolvers:        data.Resolvers,
		Scalars:          nil,
		ResolveType:      nil,
	})

	if err != nil {
		panic(err)
	}

	gql.Debug = true

	var dirname, _ = os.Getwd()
	err = ioutil.WriteFile(dirname+"/data/schema.json", []byte(gql.PrintSchema()), 0644)

	if err != nil {
		panic(err)
	}

}
