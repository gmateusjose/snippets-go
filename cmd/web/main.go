package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command line flag with the name 'addr', a default value of
	// an port and some short help text explaining what the flag controls. The value
	// of flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line args
	// this reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this "before" you use the addr variable
	// otherwise it will always contain the defalut value of ":4000". If any error
	// was encountered during parsing the application will be terminated.
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the "./ui/static"
	// directory. Note that the path given to the http.Dir function is relative
	// to the project directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
