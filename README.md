
# Micro Services in Golang.

## Episode 1

- What is *http.HandleFunc*? Handle Func is a convinience method in the http package. It registers a function to a path on a thing called the *default servemux*. What it actually does is 
- Default Serve Mux.  It is an http Handler. Everything related to the server in Go is an HTTP handler. Responsible for redirecting paths.
- http.ListenAndServe
- Handler in http is just a interface. So any
- http request has a field called body. Body is an io.ReadCloser. So it  implements the interface io.ReadCloser. So what that means is that you can use any of the standard Go liraries to reading from that stream.
- ioutil.ReadAll is the simplest way to readAll from anything that implements io.Reader. It will return data, or err.

## Episode 2

- You don't want to directly create concrete objects inside a Handler. Because it impacts testability.
- 