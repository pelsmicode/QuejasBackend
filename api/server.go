package api

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func Start() {
	router := Controllers()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8484"
	}

	handlerServer := cors.AllowAll().Handler(router)
	print("Servidor levantado en " + PORT)
	log.Fatalln(http.ListenAndServe(":"+PORT, handlerServer))
}
