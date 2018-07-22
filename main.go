package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PrinceNorin/monga/controllers"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      controllers.Router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
