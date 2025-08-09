package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// Google Search 3.0: Reduce tail latency using replicated search servers
type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q", kind, query))
	}
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	for i := range replicas {
		go func(i int) { c <- replicas[i](query) }(i)
	}
	return <-c
}

// Create replicas dynamically
func createReplicas(serviceType string, numReplicas int) []Search {
	replicas := make([]Search, numReplicas)
	for i := 0; i < numReplicas; i++ {
		replicas[i] = fakeSearch(fmt.Sprintf("%s%d", serviceType, i+1))
	}
	return replicas
}

func Google(query string) (results []Result) {
	numReplicas := 2 // configurable per service

	webReplicas := createReplicas("web", numReplicas)
	imageReplicas := createReplicas("image", numReplicas)
	videoReplicas := createReplicas("video", numReplicas)

	c := make(chan Result)
	go func() { c <- First(query, webReplicas...) }()
	go func() { c <- First(query, imageReplicas...) }()
	go func() { c <- First(query, videoReplicas...) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		default:
			// If no result is ready, we can continue to wait for the next one
			time.Sleep(10 * time.Millisecond) // Avoid busy waiting
		}
	}
	return
}
