package main

import (
	"bus-sample-project/counter"
	"bus-sample-project/printer"
	"fmt"

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
	bus.RegisterTopics("user.created", "user.canceled")

	// load printer package
	printer.Load()

	// no need to load counter package since we are running the
	// FetchEventCount function from the counter package, it will auto execute
	// the init function on load
}

// User struct for sample event
type User struct {
	Name  string
	Email string
}

func main() {
	txID := monoton.Next()
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("MyName#%d", i)
		email := fmt.Sprintf("sample%d@example.com", i)
		bus.Emit("user.created", User{Name: name, Email: email}, txID)
	}

	bus.Emit("user.canceled", User{Name: "Another", Email: "sample@example.com"}, "")

	// Let's print event counts for each topic
	for _, topic := range bus.ListTopics() {
		fmt.Printf(
			"Total evet count for %s: %d\n",
			topic.Name,
			counter.FetchEventCount(topic.Name),
		)
	}
}