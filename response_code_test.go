package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		//writer.WriteHeader(http.StatusBadRequest)
		writer.WriteHeader(400) //Bad Request
		fmt.Fprintf(writer, "error, name is empty")
	} else {
		writer.WriteHeader(200) // OK Sukses
		fmt.Fprintf(writer, "success, name is %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest("GET", "/localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", "/localhost:8080/?name=Ferdi", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}
