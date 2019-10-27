package calculator

import (
	"bus-sample-project/models"
	"fmt"
	"sync"

	"github.com/mustafaturan/bus"
)

var total int64

var c chan *bus.Event

const worker = "calculator"

// Start registers the calculator handler
func Start(wg *sync.WaitGroup) {
	h := bus.Handler{Handle: sum, Matcher: "^order.(created|canceled)$"}
	bus.RegisterHandler(worker, &h)
	fmt.Printf("Registered calculator handler...\n")

	total = 0
	c = make(chan *bus.Event)
	wg.Add(1)
	go calculate(wg)
}

// Stop deregisters the calculator handler
func Stop() {
	bus.DeregisterHandler(worker)
	c <- nil
}

func sum(e *bus.Event) {
	c <- e
}

func calculate(wg *sync.WaitGroup) {
	defer wg.Done()
	defer printTotal()
	for {
		e := <-c
		if e == nil {
			break
		}
		amount := int64(e.Data.(models.Order).Amount)
		switch e.Topic.Name {
		case "order.created":
			total += amount
		case "order.canceled":
			total -= amount
		default:
			fmt.Printf("whoops unexpected topic (%s)", e.Topic.Name)
		}
	}
}

func printTotal() {
	fmt.Printf("Order total amount %d\n", total)
}
