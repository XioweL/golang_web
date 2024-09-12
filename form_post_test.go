package golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	//err := request.ParseForm()
	//if err != nil {
	//	panic(err)
	//}
	//firstname := request.PostFormValue("first_name")
	//lastname := request.PostFormValue("last_name")

	firstName := request.PostFormValue("first_name")
	lastName := request.PostFormValue("last_name")

	fmt.Fprintf(writer, "%s %s", firstName, lastName)

}
func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Ferdi&last_name=Alvan")
	request := httptest.NewRequest("POST", "/localhost:8080", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}
