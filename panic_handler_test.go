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

/*
- Panic Handler as a giving additional information when it happens an error
*/

func TestPanicHandler(t *testing.T) {
	// (1) Create Router
	router := httprouter.New()

	// (2) Create router panic handler
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		// (3) Print information when panic
		fmt.Fprint(writer, "Panic : ", error)
	}

	// (4) Create endpoint Root
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// (5) Making panic response with argument "Ups"
		panic("Ups")
	})

	// (6) Create test request
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	// (7) Create new recorder for response writer
	recorder := httptest.NewRecorder()

	// (8) Send argument to router
	router.ServeHTTP(recorder, request)

	// (9) Create var response for get response from recorder
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	// (10) Check if error
	if err != nil {
		panic(err)
	}

	// (11) Testing assert with response body "Panic : Ups"
	assert.Equal(t, "Panic : Ups", string(body))
}
