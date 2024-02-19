package main

import (
	"log"
	"os"
	"time"

	"fmt"
	"html/template"
	"net/http"
	_ "newsApp/matchers"
	"newsApp/search"

	"github.com/patrickmn/go-cache"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	intialHandler := func(w http.ResponseWriter, r *http.Request) {

		results := search.Run("");
 
		tmpl, err := template.ParseFiles("template.html")
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}

		fmt.Println(results)
		err = tmpl.Execute(w, struct{ Results []*search.Result }{results})
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		}

}

c := cache.New(5*time.Minute, 10*time.Minute)

searchHandler := func(w http.ResponseWriter, r *http.Request) {
 
	query := r.URL.Query()
 
	queryKey := query.Get("searchQuery")
	
	 cachedData, found := c.Get(queryKey)
	 var results []*search.Result
   results = make([]*search.Result, 0)

	 if found {
		   results = cachedData.([]*search.Result);
			 fmt.Println("Data retrieved from cache:", cachedData)
			 tmpl, err := template.ParseFiles("template.html")
	     if err != nil {
			    http.Error(w, err.Error(), http.StatusInternalServerError)
			    return
	     }


	fmt.Println(results)
	err = tmpl.Execute(w, struct{ Results []*search.Result }{results})
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	 } else {
		   fmt.Println("not ok")
			 results := search.Run(queryKey)
			 fmt.Println("Data fetched:", results)
			 c.Set(queryKey, results, cache.DefaultExpiration)

			 tmpl, err := template.ParseFiles("template.html")
    	if err != nil {
		   	http.Error(w, err.Error(), http.StatusInternalServerError)
		    	return
     	}


	    fmt.Println(results)
	     err = tmpl.Execute(w, struct{ Results []*search.Result }{results})
     	if err != nil {
		   	http.Error(w, err.Error(), http.StatusInternalServerError)
     	}
	 }

	
}


http.HandleFunc("/", intialHandler)
http.HandleFunc("/search", searchHandler)


fmt.Println("Server started on port 8080")
http.ListenAndServe(":8080", nil)
}
