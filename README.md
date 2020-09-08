
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
- Graceful Shutdown. If we are doing something like a large file upload. Like a DB transaction. Then we decide that we want to shut the server down because I want to upgrade the version of it or something like that. If I don't do it gracefully, then I am risking disconnecting my client. They are not gonna finish what they are doing and it is going to cause an error on their side.
- With the Go server we can use the ShutDown() function which will wait until the requests that are currently being handled are completed before it shuts down.