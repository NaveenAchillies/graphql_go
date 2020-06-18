// Command graphql-go-example starts an HTTP GraphQL API server which is backed by data

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"./resolver"
	graphqlSchema "./schema"
	"github.com/graph-gophers/graphql-go"
)

type JSON = map[string]interface{}

type ClientQuery struct {
	OpName    string
	Query     string
	Variables JSON
}

var (
	root   *resolver.RootResolver
	schema *graphql.Schema
)

func init() {
	root = &resolver.RootResolver{}
	root.AddClients()
	schema = graphql.MustParseSchema(graphqlSchema.String(), root)
}

func main() {
	// c := swapi.NewClient(http.DefaultClient) // TODO: don't use the default client.
	// root, err := resolver.NewRoot(c)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Create the request handler; inject dependencies.
	// schema := graphql.MustParseSchema(schema.String(), root)

	ctx := context.Background()
	q2 := ClientQuery{
		OpName: "CustomerOrders",
		Query: `query CustomerOrders($number: String!) {
				customerOrder(number: $number) {
					number
					VendorOrders{
						VoonikOrderNumber
					}
				}
			}`,
		Variables: JSON{
			"number": "R6334676811-3",
		},
	}
	resp2 := schema.Exec(ctx, q2.Query, q2.OpName, q2.Variables)
	json2, err := json.MarshalIndent(resp2, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json2))

	q3 := ClientQuery{
		OpName: "GetProductDataOms",
		Query: `query GetProductDataOms($id: Int!) {
			getProductDataOms(id: $id) {
				name
				slug
					Variants{
						sku
					}
				}
			}`,
		Variables: JSON{
			"id": "48896",
		},
	}
	resp3 := schema.Exec(ctx, q3.Query, q3.OpName, q3.Variables)
	json3, err := json.MarshalIndent(resp3, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json3))

	q5 := ClientQuery{
		OpName: "SetDueTime",
		Query: `mutation SetDueTime($voonikOrderNumber: String!, $dueTime: Int!) {
			setDueTime(voonikOrderNumber: $voonikOrderNumber, dueTime: $dueTime) {
				VendorOrders{
					VoonikOrderNumber
				}
			}
		}`,
		Variables: JSON{
			"voonikOrderNumber": "R6334676811-3",
			"dueTime":           1,
		},
	}
	resp5 := schema.Exec(ctx, q5.Query, q5.OpName, q5.Variables)
	json5, err := json.MarshalIndent(resp5, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json5))

	// Tweak configuration values here.
	// var (
	// 	addr              = ":8000"
	// 	readHeaderTimeout = 1 * time.Second
	// 	writeTimeout      = 10 * time.Second
	// 	idleTimeout       = 90 * time.Second
	// 	maxHeaderBytes    = http.DefaultMaxHeaderBytes
	// )

	// log.SetFlags(log.Lshortfile | log.LstdFlags)

	// c := swapi.NewClient(http.DefaultClient) // TODO: don't use the default client.

	// root, err := resolver.NewRoot(c)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Create the request handler; inject dependencies.
	// h := handler.GraphQL{
	// 	// Parse and validate schema. Panic if unable to do so.
	// 	Schema:  graphql.MustParseSchema(schema.String(), root),
	// 	Loaders: loader.Initialize(c),
	// }

	// // Register handlers to routes.
	// mux := http.NewServeMux()
	// mux.Handle("/", handler.GraphiQL{})
	// mux.Handle("/graphql/", h)
	// mux.Handle("/graphql", h) // Register without a trailing slash to avoid redirect.

	// // Configure the HTTP server.
	// s := &http.Server{
	// 	Addr:              addr,
	// 	Handler:           mux,
	// 	ReadHeaderTimeout: readHeaderTimeout,
	// 	WriteTimeout:      writeTimeout,
	// 	IdleTimeout:       idleTimeout,
	// 	MaxHeaderBytes:    maxHeaderBytes,
	// }

	// // Begin listeing for requests.
	// log.Printf("Listening for requests on %s", s.Addr)

	// if err = s.ListenAndServe(); err != nil {
	// 	log.Println("server.ListenAndServe:", err)
	// }

	// // TODO: intercept shutdown signals for cleanup of connections.
	// log.Println("Shut down.")
}
