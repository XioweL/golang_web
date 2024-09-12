package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "XioweL"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Succsess create cookie")

}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("XioweL")
	if err != nil {
		fmt.Fprint(writer, "fail, no cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}

}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/localhost:8080/?name=Ferdi", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookie := recorder.Result().Cookies()

	for _, cookie := range cookie {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}

}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "XioweL"
	cookie.Value = "Ferdi"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
