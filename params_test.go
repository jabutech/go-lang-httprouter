package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	// (1) Create Router
	router := httprouter.New()

	// (2) Create endpoint with parameter name id
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// (3) Get id from param
		id := params.ByName("id")

		// (4) Print response with id
		fmt.Fprintf(writer, "Product %s", id)
	})

	// (5) Create test request with send url with params
	request := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	// (6) Create new recorder for response writer
	recorder := httptest.NewRecorder()

	// (7) Send argument to router
	router.ServeHTTP(recorder, request)

	// (8) Create var response for get response from recorder
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	// (9) Check if error
	if err != nil {
		panic(err)
	}

	// (10) Testing assert
	assert.Equal(t, "Product 1", string(body))
}
