package main

import (
	"bus-sample-project/calculator"
	"bus-sample-project/config"
	"bus-sample-project/counter"
	"bus-sample-project/models"
	"bus-sample-project/printer"
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/mustafaturan/bus"
	"github.com/mustafaturan/monoton"
)

func main() {
	config.Init()

	var wg sync.WaitGroup
	defer wg.Wait()

	// register the event printer handler (synchronous handler)
	printer.Start()
	defer printer.Stop()

	// register the event counter handler (asynchronous handler)
	counter.Start(&wg)
	defer counter.Stop()

	// register the event calculator handler (asynchronous handler)
	calculator.Start(&wg)
	defer calculator.Stop()

	txID := monoton.Next()
	ctx := context.Background()
	context.WithValue(ctx, bus.CtxKeyTxID, txID)

	b := config.Bus

	for i := 0; i < 3; i++ {
		_, err := b.Emit(
			ctx,
			"order.created",
			models.Order{Name: fmt.Sprintf("Product #%d", i), Amount: randomAmount()},
		)
		if err != nil {
			fmt.Println("ERROR >>>>", err)
		}
	}

	// if the txID is not available on the context and bus package sets it
	ctx = context.Background()
	_, err := b.Emit(
		ctx,              // context
		"order.canceled", // topic
		models.Order{Name: "Product #N", Amount: randomAmount()}, // data
	)
	if err != nil {
		fmt.Println("ERROR >>>>", err)
	}
}

func randomAmount() int {
	max := 100
	min := 10
	return rand.Intn(max-min) + min
}
