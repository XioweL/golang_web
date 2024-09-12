package golang_web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MENGGUNAKAN MAP
func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Ferdi",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Ada Lagi",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

//MENGGUNAKAN STRUCT

type Address struct {
	Street string
}
type Page struct {
	Title   string
	Name    string
	Address Address
	Kel     string
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Ferdi",
		Address: Address{
			Street: "Jalan Negla No 5A",
		},
		Kel: "Isola",
	})

}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
