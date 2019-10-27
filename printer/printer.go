package printer

import (
	"fmt"

	"github.com/mustafaturan/bus"
)

// Register registers the printer handler
func Register() {
	h := bus.Handler{Handle: print, Matcher: ".*"}
	bus.RegisterHandler("printer", &h)
	fmt.Printf("Registered printer handler...\n")
}

func print(e *bus.Event) {
	fmt.Printf("\nEvent for %s: %+v\n\n", e.Topic.Name, e)
}
