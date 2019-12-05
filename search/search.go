package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching.
var matcher = make(map[string]Matcher)

// Run performs the search logic.
func Run(term string) {
	// Retrieve the list of feed to search through.
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create a unbuffere channel to receive match results.
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while -
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// Lauch a goroutine for each feed to find results.
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, term, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()
}
