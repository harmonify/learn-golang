package helper

import (
	"belajar-golang-restful-api/model/web"
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) {
	decoder := json.NewDecoder(request.Body)
	PanicIfError(decoder.Decode(result))
}

func WriteToResponseBody(writer http.ResponseWriter, response web.WebResponse) {
	writer.WriteHeader(response.Code)
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	PanicIfError(encoder.Encode(response))
}
