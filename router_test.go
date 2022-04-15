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

func TestRouter(t *testing.T) {
	// (1) Create Router
	router := httprouter.New()

	// (2) Create endpoint Root
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// (3) Print response
		fmt.Fprint(writer, "Hello World")
	})

	// (4) Create test request
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	// (5) Create new recorder for response writer
	recorder := httptest.NewRecorder()

	// (6) Send argument to router
	router.ServeHTTP(recorder, request)

	// (7) Create var response for get response from recorder
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	// (8) Check if error
	if err != nil {
		panic(err)
	}

	// (9) Testing assert
	assert.Equal(t, "Hello World", string(body))
}
