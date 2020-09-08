package main

import (
	"log"
	"net/http"
	"os"

	"github.com/psypriest/go-microservices/pkgs/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	// ServeMux implements the handler interface so we want to use our
	// own ServeMux
	http.ListenAndServe(":9090", sm)

}
