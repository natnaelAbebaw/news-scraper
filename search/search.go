package search

import (
	"log"
	"sync"
)
var matchers = make(map[string]Matcher)


func Run(searchTerm string) ( [] *Result){

	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	
	results := make(chan *Result)


	var waitGroup sync.WaitGroup


	waitGroup.Add(len(feeds))


	for _, feed := range feeds {
	
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}


		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {

		waitGroup.Wait()

		
		close(results)
	}()


	resultList := Display(results)

	return resultList
}


func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
