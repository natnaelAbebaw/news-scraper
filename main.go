package main

import (
	"log"
	"os"

	_ "newsApp/matchers"
	"newsApp/search"
	"net/http"
	"fmt"
	"html/template"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	intialHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

}

searchHandler := func(w http.ResponseWriter, r *http.Request) {
 
	query := r.URL.Query()
 
	name := query.Get("searchQuery")
			
	results := search.Run(name)
 
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	// Execute the template with the data

	fmt.Println(results)
	err = tmpl.Execute(w, struct{ Results []*search.Result }{results})
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}


// Register the handler function with the default ServeMux (multiplexer)
http.HandleFunc("/", intialHandler)
http.HandleFunc("/search", searchHandler)

// Start the server and specify the port to listen on
fmt.Println("Server started on port 8080")
http.ListenAndServe(":8080", nil)
}
