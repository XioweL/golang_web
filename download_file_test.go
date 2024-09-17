package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Query().Get("fileName")
	if fileName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(writer, request, "./resources/"+fileName)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
