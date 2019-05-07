package calculator

import (
	"bus-sample-project/models"
	"fmt"

	"github.com/mustafaturan/bus"
)

var total int64

var c chan *bus.Event

func init() {
	h := bus.Handler{Handle: sum, Matcher: "^order.(created|canceled)$"}
	bus.RegisterHandler("calculator", &h)
	fmt.Printf("Registered calculator handler...\n")

	total = 0
	c = make(chan *bus.Event)

	go calculate()
}

func sum(e *bus.Event) {
	c <- e
}

func calculate() {
	for {
		e := <-c
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

// TotalAmount returns total amount
func TotalAmount() int64 {
	return total
}
