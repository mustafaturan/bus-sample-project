package printer

import (
	"fmt"

	"github.com/mustafaturan/bus"
)

// Load registers the printer handler
func Load() {
	h := bus.Handler{Handle: print, Matcher: ".*"}
	bus.RegisterHandler("printer", &h)
	fmt.Printf("Registered printer handler...\n")
}

func print(e *bus.Event) {
	fmt.Printf("\nEvent for %s: %+v\n\n", e.Topic.Name, e)
}
