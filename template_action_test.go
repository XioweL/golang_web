package golang_web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "",
	})

}

func TestTemplateIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateIf(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// TEMPLATE OPERATOR

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Data Struct",
		"FinalValue": 70,
	})

}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// FUNCTION RANGE

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Data Struct",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

//TEMPLATE WITH

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title": "Template Data Struct",
		"Name":  "Ferdi",
		"Address": map[string]interface{}{
			"Street": "Jalan Negla No 5A",
			"City":   "Bandung",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
