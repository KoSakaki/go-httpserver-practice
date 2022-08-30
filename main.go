package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	_ "reflect"
	_ "runtime"
)

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello!")
// }

// func log(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
// 		fmt.Println("Handler function called - " + name)
// 		h(w, r)
// 	}
// }

// func protected(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// 省略
// 		h.ServeHTTP(w ,r)
// 	})
// }

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello!, %s\n", p.ByName("name"))
}

func main() {

	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}

	// http.Handle("/hello/", protected(log(hello)))

	server.ListenAndServe()
}