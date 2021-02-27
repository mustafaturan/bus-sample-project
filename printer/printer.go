package printer

import (
	"bus-sample-project/config"
	"context"
	"fmt"

	"github.com/mustafaturan/bus/v2"
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

func print(ctx context.Context, e *bus.Event) {
	fmt.Printf("\nEvent for %s: %+v\n\n", e.Topic, e)
}
