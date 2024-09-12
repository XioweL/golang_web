package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}

}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hi?name=Ferdi", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstname := request.URL.Query().Get("first_name")
	lastname := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "%s %s", firstname, lastname)

}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hi?first_name=Ferdi&last_name=Alvan", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))

}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hi?name=Ferdi&name=Alvan&name=Xiowel", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}
