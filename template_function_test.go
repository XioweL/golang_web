package golang_web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "XioweL"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Ferdi",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))

}

// GLOBAL FUNCTION
func TemplateFuncGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len "XioweL"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Ferdi",
	})
}

func TestTemplateFuncGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFuncGlobal(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))

}

// GLOBAL FUNCTION LAGI

func TemplateFuncCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "mochamad ferdian alvanza",
	})
}

func TestTemplateFuncCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFuncCreateGlobal(recorder, request)

	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))

}

// FUNCTION PIPELINES

func TemplateFuncPipelines(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "mochamad ferdian alvanza",
	})
}

func TestFuncPipeLines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFuncPipelines(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Body)
	fmt.Println(string(body))

}
