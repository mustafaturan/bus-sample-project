package config

import (
	"github.com/mustafaturan/bus/v3"
	"github.com/mustafaturan/monoton/v3"
	"github.com/mustafaturan/monoton/v3/sequencer"
)

// Bus is a ref to bus.Bus
var Bus *bus.Bus

// Monoton is an instance of monoton.Monoton
var Monoton monoton.Monoton

// Init inits the app config
func Init() {
	// configure id generator (it doesn't have to be monoton)
	node := uint64(1)
	initialTime := uint64(1577865600000)
	m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil {
		panic(err)
	}

	// init an id generator
	var idGenerator bus.Next = m.Next

	// create a new bus instance
	b, err := bus.NewBus(idGenerator)
	if err != nil {
		panic(err)
	}

	// maybe register topics in here
	b.RegisterTopics("order.created", "order.canceled")

	Bus = b
	Monoton = m
}
