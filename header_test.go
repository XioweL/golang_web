package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contenType := request.Header.Get("content-type")
	fmt.Fprint(writer, contenType)

}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content", "XioweL Golang")
	fmt.Fprint(writer, "OK")

}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/localhost:8080", nil)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("content"))
}
