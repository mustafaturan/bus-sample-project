package counter

import (
	"fmt"
	"sync"

	"github.com/mustafaturan/bus"
)

var topics map[string]uint

var c chan *bus.Event

const worker = "counter"

func init() {
	topics = make(map[string]uint)
	c = make(chan *bus.Event, 5)
}

// Start registers the counter handler
func Start(wg *sync.WaitGroup) {
	h := bus.Handler{Handle: count, Matcher: ".*"}
	bus.RegisterHandler(worker, &h)
	fmt.Printf("Registered counter handler...\n")

	wg.Add(1)
	go increment(wg)
}

// Stop deregister the counter handler
func Stop() {
	bus.DeregisterHandler(worker)
	c <- nil
}

func count(e *bus.Event) {
	c <- e
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	defer printEventCounts()
	for {
		// Separating the logic from channels would be better. So, please
		// consider this is an example but do not consider as best practice.
		e := <-c
		if e == nil {
			break
		}
		topics[e.Topic.Name]++
	}
}

func printEventCounts() {
	// Let's print event counts for each topic
	for _, topic := range bus.ListTopics() {
		fmt.Printf(
			"Total evet count for %s: %d\n",
			topic.Name,
			topics[topic.Name],
		)
	}
}
