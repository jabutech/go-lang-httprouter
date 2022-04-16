package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

/*
- Serve File a get static file use function ServeGile(Path, FileSystem)
- Catch All Parameter in serve file only use *pathFile
*/

// (1) Load file hello.txt on folder resources with golang embed
//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	// (1) Create Router
	router := httprouter.New()

	// (2) Go to the resources folder with fs.Sub
	directory, _ := fs.Sub(resources, "resources")

	// (3) Create endpoint with Method ServeFiles and name endpoint files with parameter name *filepath as catch all parameter
	router.ServeFiles("/files/*filepath", http.FS(directory))

	// (5) Create test request with send url with params
	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
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

	// (10) Testing assert must be Hello Http Router
	assert.Equal(t, "Hello Http Router", string(body))
}
