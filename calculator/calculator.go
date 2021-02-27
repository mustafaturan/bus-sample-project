package calculator

import (
	"bus-sample-project/config"
	"bus-sample-project/models"
	"context"
	"fmt"
	"sync"

	"github.com/mustafaturan/bus/v2"
)

var total int64

var c chan *bus.Event

const worker = "calculator"

func init() {
	c = make(chan *bus.Event)
}

// Start registers the calculator handler
func Start(wg *sync.WaitGroup) {
	b := config.Bus
	h := bus.Handler{Handle: sum, Matcher: "^order.(created|canceled)$"}
	b.RegisterHandler(worker, &h)
	fmt.Printf("Registered calculator handler...\n")

	wg.Add(1)
	go calculate(wg)
}

// Stop deregisters the calculator handler
func Stop() {
	defer fmt.Printf("Deregistered calculator handler...\n")

	b := config.Bus
	b.DeregisterHandler(worker)
	c <- nil
}

func sum(_ context.Context, e *bus.Event) {
	c <- e
}

func calculate(wg *sync.WaitGroup) {
	defer wg.Done()
	defer printTotal()
	for {
		// Separating the logic from channels would be better. So, please
		// consider this is an example but do not consider as best practice.
		e := <-c
		if e == nil {
			break
		}

		amount := int64(e.Data.(models.Order).Amount)

		// I personally recommend creating separate consumer for each topic. But
		// in this context, there is an example usage of the same consumer for
		// multiple topics(purposes).
		switch e.Topic {
		case "order.created":
			total += amount
		case "order.canceled":
			total -= amount
		default:
			fmt.Printf("whoops unexpected topic (%s)", e.Topic)
		}
	}
}

func printTotal() {
	fmt.Printf("Order total amount %d\n", total)
}
