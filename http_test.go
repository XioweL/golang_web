package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hi", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, r)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

}
