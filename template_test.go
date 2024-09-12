package golang_web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TEMPLATE HTML
func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<!DOCTYPE html><body>{{.}}</body></html>`
	//t, err := template.New("SIMPLE").Parse(templateText)
	//if err != nil {
	//	panic(err)
	//}
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

func TestTemplateHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))

}

// TEMPLATE FILE
func TemplateHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template File")

}

func TestTemplateHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateHTMLFile(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))
}

// TEMPLATE DIRECTORY
func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template Directory")

}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))
}

// TEMPLATE GOLANG EMBED

////go:embed templates/*.gohtml
//var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template Embed")

}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))

}
