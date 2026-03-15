package main

import (
	"fmt"
	"net/http"

	"github.com/Searching96/my-go-api/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handlers(r)

	fmt.Println("Starting GO API service...")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
