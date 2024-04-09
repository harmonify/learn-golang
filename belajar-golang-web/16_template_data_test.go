package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]string{
		"Title": "Template Data Map",
		"Name":  "Eko",
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	bodyHtml := string(body)

	fmt.Println(bodyHtml)

	assert.HTTPBodyContains(t, TemplateDataMap, http.MethodGet, "http://localhost:8080", nil, "<title>Template Data Map</title>")
	assert.HTTPBodyContains(t, TemplateDataMap, http.MethodGet, "http://localhost:8080", nil, "Hello Eko")
}

type Address struct {
	Street string
}

type Page struct {
	Title string
	Name  string
	Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Map",
		Name:  "Eko",
		Address: Address{
			Street: "Jalan Belum Ada",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	bodyString := assert.HTTPBody(TemplateDataStruct, http.MethodGet, "http://localhost:8080", nil)
	fmt.Println(bodyString)
	assert.Contains(t, bodyString, "<title>Template Data Map</title>")
	assert.Contains(t, bodyString, "Hello Eko")
	assert.Contains(t, bodyString, "Jalan Belum Ada")
}
