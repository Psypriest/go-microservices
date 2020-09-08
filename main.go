package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"os/signal"


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
	
	// It is a really good practice to set up timeouts.
	// Resources are finite and there are only so many connections
	// your server can handle. So if there are too many blocking connections
	// Then it is basically like a DoS attack.



	// An http server is just a type in Go
	s := &http.Server{
		Addr: ":9090",
		//Handler is going to be the ServeMux
		Handler: sm,
		// Server time out set to 120 Seconds
		IdleTimeout: 120*time.Second,
		// Set ReadTimeout, since a tiny info, 1 second
		ReadTimeout: 1 *time.Second,
		// Set to 1 second as well.
		WriteTimeout: 1 * time.Second,

	}
	
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful Shutdown",sig)

	tc, _ := context.WithTimeout(context.Background(),30*time.Second)
	// 
	s.Shutdown(tc)


	
	//http.ListenAndServe(":9090", sm)



}
