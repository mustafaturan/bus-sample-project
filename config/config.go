package config

import (
	"github.com/mustafaturan/bus"
	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"
)

// Bus is a ref to bus.Bus
var Bus *bus.Bus

// Init inits the app config
func Init() {
	// configure id generator (it doesn't have to be monoton)
	node := uint64(1)
	initialTime := uint64(0)
	monoton.Configure(sequencer.NewMillisecond(), node, initialTime)

	// init an id generator
	var idGenerator bus.Next = monoton.Next

	// create a new bus instance
	b, err := bus.NewBus(idGenerator)
	if err != nil {
		panic(err)
	}

	// maybe register topics in here
	b.RegisterTopics("order.created", "order.canceled")

	Bus = b
}
