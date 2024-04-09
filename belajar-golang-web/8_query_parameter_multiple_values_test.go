package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	// Instead of using the .Get() method on the url.Values above,
	// we could directly access the map that contains the query parameter values
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Eko&name=Kurniawan&name=Khannedy", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	assert.Equal(t, "Eko Kurniawan Khannedy", string(body))
}
