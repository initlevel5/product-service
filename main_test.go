package main_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	sgraphql "github.com/initlevel5/product-service/service/graphql"
	"github.com/initlevel5/product-service/service/mock"
)

var (
	addr        = ":8080"
	url         = "http://localhost:8080/query"
	contentType = "application/x-www-form-urlencoded"
)

func TestMicroservices(t *testing.T) {
	var tests = []struct {
		input    []byte
		expected []byte
	}{
		{
			[]byte(`{"query":"{product(id:\"1001\"){id title created price}}"}`),
			[]byte(`{"data":{"product":{"id":"1001","title":"Socks","created":"0001-01-01 00:00:00.000000 +0000 UTC","price":2.95}}}`),
		},
		{
			[]byte(`{"query":"{product(id:\"1001\"){id}}"}`),
			[]byte(`{"data":{"product":{"id":"1001"}}}`),
		},
	}

	go func() {
		l := log.New(os.Stdout, "test: product-service: ", log.LstdFlags)

		svc := mock.NewProductService(l)

		schema := graphql.MustParseSchema(sgraphql.Schema, sgraphql.NewResolver(svc, l))

		http.Handle("/query", &relay.Handler{Schema: schema})
		t.Fatal(http.ListenAndServe(addr, nil))
	}()

	for _, test := range tests {
		req := bytes.NewReader(test.input)
		resp, err := http.Post(url, contentType, req)
		if err != nil {
			t.Fatal(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()

		if len(body) != len(test.expected) {
			t.Errorf("body: %s\nexpected: %s\n", body, test.expected)
		}
	}
}
