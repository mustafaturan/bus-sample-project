package counter

import (
	"fmt"
	"sync"

	"github.com/mustafaturan/bus"
)

var topics map[string]uint

var c chan *bus.Event

// Load registers the counter handler
func Load(wg *sync.WaitGroup) {
	h := bus.Handler{Handle: count, Matcher: ".*"}
	bus.RegisterHandler("counter", &h)
	fmt.Printf("Registered counter handler...\n")

	topics = make(map[string]uint, 0)
	c = make(chan *bus.Event)

	wg.Add(1)
	go increment(wg)
}

func Close() {
	c <- nil
}

func count(e *bus.Event) {
	c <- e
}

func increment(wg *sync.WaitGroup) {
	defer printEventCounts()
	defer wg.Done()
	for {
		e := <-c
		if e == nil {
			break
		}
		n := e.Topic.Name
		if count, ok := topics[n]; ok {
			topics[n] = count + 1
		} else {
			topics[n] = 1
		}
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
