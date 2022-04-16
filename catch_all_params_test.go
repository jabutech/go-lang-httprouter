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
- catch all parameter for taking all the parameter
- Catch all parameter must start with * (star symbol), and following with parameter name
- Catch all parameter must be at the end position of the url
*/

func TestCatchAllParameter(t *testing.T) {
	// (1) Create Router
	router := httprouter.New()

	// (2) Create endpoint with parameter name image as catch all parameter
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// (3) Get id from param
		image := params.ByName("image")

		// (4) Print response with id
		fmt.Fprintf(writer, "Images %s", image)
	})

	// (5) Create test request with send url with params
	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
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
	assert.Equal(t, "Images /small/profile.png", string(body))
}
