package main

import (
	"net/http"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Greeting server.
// @termsOfService http://swagger.io/terms/

// @host greeting.swagger.io
// @BasePath /v2
func main() {
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"), //The url pointing to API definition"
	))
	r.Get("/greeting", greeting)
	http.ListenAndServe(":1323", r)
}

// Greeting example
// @Summary get greeting
// @Description get "hello world"
// @ID greeting
// @Produce json
// @Success 200 {string} string "ok"
// @Router /greeting [get]
func greeting(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}
