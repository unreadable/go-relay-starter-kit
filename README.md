# Golang-Relay Starter Kit

This kit includes:
- a **Golang** GraphQL server: to serve the back-end `graphql-go` server that handles GraphQL queries
- a Babel transpiler workflow using `webpack` that you can use to get started building an app with Relay.

For a walkthrough, see the [Relay tutorial](https://facebook.github.io/relay/docs/tutorial.html).

### Notes:
This is based on `playlyfe/graphql` implementation but it can be customized for other implementation as well. 
Be sure to watch the repository for future changes!

## Installation

- Install dependencies for NodeJS app server
```
npm install
```
- Install dependencies for Golang GraphQL server
```
go get -v
```

## Running

Start a local server:

```
npm start
```

The above command will run both the go-graphql server and will display a simple `Hello world!`based on the API.

## Developing

### JavaScript
Any changes you make to files in the `js/` directory will cause the server to
automatically rebuild the app and refresh your browser.

### Golang

#### Schema data
The Schema and Resolvers should be kept in `data/graphql.go`.
Since Golang does not support loading package / module dynamically, after updating the `data/graphql.go` file, make sure to update the schema:
- `graphql.go`
- `utils/updateSchema.go`

For e.g

```go
import (
  ...
  "github.com/sogko/golang-relay-starter-kit/data" // <--- update to package containing schema
)
```

#### Schema updates
If at any time you make changes to `data/graphql.go`, stop the server,
rebuild `data/schema.json` and restart the server:

```
npm run update-schema
npm start
```

`schema.json` is needed by the JS code for `./build/babelRelayPlugin.js`

## Examples
This starter kit provides a simple go-graphql API that displays "Hello world!"

Feel free to submit a PR to add to this list.

## TODOs
- [ ] Make it production-ready

## Credits
This kit is build on top of https://github.com/relayjs/relay-starter-kit
