package counter

import (
	"bus-sample-project/config"
	"context"
	"fmt"
	"sync"

	"github.com/mustafaturan/bus/v2"
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
	b := config.Bus
	h := bus.Handler{Handle: count, Matcher: ".*"}
	b.RegisterHandler(worker, &h)
	fmt.Printf("Registered counter handler...\n")

	wg.Add(1)
	go increment(wg)
}

// Stop deregister the counter handler
func Stop() {
	defer fmt.Printf("Deregistered counter handler...\n")

	b := config.Bus
	b.DeregisterHandler(worker)
	c <- nil
}

func count(_ context.Context, e *bus.Event) {
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
		topics[e.Topic]++
	}
}

func printEventCounts() {
	b := config.Bus
	// Print event counts for each topic
	for _, topic := range b.Topics() {
		fmt.Printf(
			"Total evet count for %s: %d\n",
			topic,
			topics[topic],
		)
	}
}
