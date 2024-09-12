package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello TemplateCaching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))
}

// AUTO ESCAPE

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  "<p>Selamat Belajar Xiowel</p>",
	})
}

func TestAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))

}

func TestServeAutoEscape(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// AUTO DISABLE ESCAPE

func TemplateDisableAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Disable Auto Escape",
		"Body":  template.HTML("<h1>Selamat Belajar Xiowel</h1>"),
		//"Body":  template.JS("<h1>Selamat Belajar Xiowel</h1>"),
		//"Body":  template.CSS("<h1>Selamat Belajar Xiowel</h1>"),
	})
}

func TestDisableAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDisableAutoEscape(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))
}

func TestServeDisableAutoEscape(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateDisableAutoEscape),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
