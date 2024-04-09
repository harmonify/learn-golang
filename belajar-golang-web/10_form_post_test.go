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

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	// firstName := request.PostFormValue("first_name") // can be executed without calling `request.ParseForm` method
	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Eko&last_name=Khannedy")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	assert.Equal(t, "Hello Eko Khannedy", string(body))
}
