package search

import (
	"log"
)

// Result contains the result of a search.
type Result struct {
	Field   string
	Content string
	Image  string
	Description string
	PubDate string
}

// Matcher defines the behavior required by types that want
// to implement a new search type.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}


func Display(results chan *Result) ([]*Result) {

	var resultList []*Result

		

	for result := range results {
		resultList = append(resultList, result)
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}

	return resultList
}
