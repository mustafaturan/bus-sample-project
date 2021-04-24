package counter

import (
	"bus-sample-project/config"
	"context"
	"fmt"
	"sync"

	"github.com/mustafaturan/bus/v3"
)

var topics map[string]uint

var c chan bus.Event
var ctx context.Context
var cancel context.CancelFunc

const worker = "counter"

func init() {
	topics = make(map[string]uint)
	c = make(chan bus.Event, 5)
	ctx, cancel = context.WithCancel(context.Background())
}

// Start registers the counter handler
func Start(wg *sync.WaitGroup) {
	b := config.Bus
	h := bus.Handler{Handle: count, Matcher: ".*"}
	b.RegisterHandler(worker, h)
	fmt.Printf("Registered counter handler...\n")

	wg.Add(1)
	go increment(ctx, wg)
}

// Stop deregister the counter handler
func Stop() {
	defer fmt.Printf("Deregistered counter handler...\n")

	b := config.Bus
	b.DeregisterHandler(worker)
	cancel()
}

func count(_ context.Context, e bus.Event) {
	c <- e
}

func increment(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer printEventCounts()
	for {
		// Separating the logic from channels would be better. So, please
		// consider this is an example but do not consider as best practice.
		select {
		case <-ctx.Done():
			return
		case e := <-c:
			topics[e.Topic]++
		}
	}
}

func printEventCounts() {
	// Print event counts for each topic
	for topic, count := range topics {
		fmt.Printf(
			"Total event count for %s: %d\n",
			topic,
			count,
		)
	}
}
