package main

import (
	"bus-sample-project/calculator"
	"bus-sample-project/counter"
	"bus-sample-project/models"
	"bus-sample-project/printer"
	"fmt"
	"math/rand"
	"sync"

	"github.com/mustafaturan/bus"
	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"
)

func init() {
	// configure id generator (it doesn't have to be monoton)
	node := uint(1)
	initialTime := uint(0)
	monoton.Configure(sequencer.NewMillisecond(), node, initialTime)

	// configure bus package
	if err := bus.Configure(bus.Config{Next: monoton.Next}); err != nil {
		panic("Whoops, couldn't configure the bus package!")
	}

	// regiter topics
	bus.RegisterTopics("order.created", "order.canceled")

	// load printer package
	printer.Load()

	// no need to load counter and calculator packages since we are running the
	// FetchEventCount, and TotalAmount function from the counter package, it
	// will auto execute the init function on load
}

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	// load printer package
	counter.Start(&wg)
	defer counter.Stop()

	// load printer package
	calculator.Start(&wg)
	defer calculator.Stop()

	txID := monoton.Next()
	for i := 0; i < 3; i++ {
		bus.Emit(
			"order.created",
			models.Order{Name: fmt.Sprintf("Product #%d", i), Amount: randomAmount()},
			txID,
		)
	}

	bus.Emit(
		"order.canceled", // topic
		models.Order{Name: "Product #N", Amount: randomAmount()}, // data
		"", // when blank bus package auto assigns an ID using the provided gen
	)

	// printer consumer processed all events at that moment since it is synchronous
	fmt.Println("You should see 4 events printed above!^^^")
}

func randomAmount() int {
	max := 100
	min := 10
	return rand.Intn(max-min) + min
}
