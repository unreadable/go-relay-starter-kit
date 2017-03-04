package data

import "github.com/playlyfe/go-graphql"

var Schema = `
        type Query {
		greetings: Greetings
	}

	type Greetings {
		hello: String!
	}
    `

var Resolvers = map[string]interface{}{
	"Query/greetings": func(params *graphql.ResolveParams) (interface{}, error) {
		return map[string]interface{}{
			"hello": "Hello world!",
		}, nil
	},
}
