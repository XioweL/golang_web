package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type logMiddleware struct {
	Handler http.Handler
}

func (middleware *logMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler") // logic disini
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler") // logic disini

}

// HANDLE ERROR MIDDLEWARE

type ErrorHandler struct {
	HandlerError http.Handler
}

func (handlerErr *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("RECOVER : ", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 Internal Server Error / Error : %s", err)
		}
	}()
	handlerErr.HandlerError.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic Arrgument")
		panic("Crash")
	})

	logMiddleware := &logMiddleware{
		Handler: mux,
	}

	ErrorHandler := &ErrorHandler{
		HandlerError: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: ErrorHandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
