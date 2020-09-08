package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	// Define a field called l, its of type log.Logger.
	l *log.Logger
}

// following the idiomatic principles of creating Go Code and define a function
// called NewHello which is goin to take a log.Logger and return a Hello Hander
// as a reference
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}

}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//l is log.Logger and l has Println method
	h.l.Println("Hello World!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
