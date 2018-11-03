package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	sgraphql "github.com/initlevel5/product-service/service/graphql"
	"github.com/initlevel5/product-service/service/mock"
	_ "github.com/initlevel5/product-service/service/postgres"
)

var (
	addr = flag.String("http_addr", ":8080", "HTTP tcp address to bind on")
)

var page = []byte(`
<!doctype html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title>product-service</title>
    </head>
    <body>
        <h1>product-service</h1>
    </body>
</html>
`)

func main() {
	flag.Parse()

	logger := log.New(os.Stdout, "product-service: ", log.LstdFlags)

	svc := mock.NewProductService(logger)

	schema := graphql.MustParseSchema(sgraphql.Schema, sgraphql.NewResolver(svc, logger))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))
	http.Handle("/query", &relay.Handler{Schema: schema})
	logger.Fatal(http.ListenAndServe(*addr, nil))
}
