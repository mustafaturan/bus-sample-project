package printer

import (
	"bus-sample-project/config"
	"fmt"

	"github.com/mustafaturan/bus"
)

// Start registers the printer handler
func Start() {
	b := config.Bus
	h := bus.Handler{Handle: print, Matcher: ".*"}
	b.RegisterHandler("printer", &h)
	fmt.Printf("Registered printer handler...\n")
}

// Stop deregisters the printer handler
func Stop() {
	defer fmt.Printf("Deregistered printer handler...\n")

	b := config.Bus
	b.DeregisterHandler("printer")
}

func print(e *bus.Event) {
	fmt.Printf("\nEvent for %s: %+v\n\n", e.Topic, e)
}
