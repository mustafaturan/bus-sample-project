# Bus Sample Project

This is an example project to demonstrate usage of the `bus` package for
internal package communication. Please note the aim of this sample project to
guide on a un-realistic use case.

## Install

```shell
git clone https://github.com/mustafaturan/bus-sample-project
cd bus-sample-project
go get github.com/mustafaturan/monoton
go get github.com/mustafaturan/bus
```

## Scenario

Assume that project has order data model and for each order creation and
cancellation we want to log order with the time of creation, count how many
orders created/cancelled, and also calculate sum of order amounts.

The sample project consist of three seperate consumers which are responsible
for printing events, counting the topic events and calculating the sum of
amounts.

### Configuration

In the example, it is used the same example configuration from the package
readme file.

```go
func init() {
	// configure id generator (it doesn't have to be monoton)
	node := uint(1)
	initialTime := uint(0)
	monoton.Configure(sequencer.NewMillisecond(), node, initialTime)

	// configure bus package
	if err := bus.Configure(bus.Config{Next: monoton.Next}); err != nil {
		panic("Whoops, couldn't configure the bus package!")
	}

	// ...
}
```

### Register topics

Assume that we have two topics which are; `order.created` and `order.canceled`.

```go
func init() {
	// ...
	bus.RegisterTopics("order.created", "order.canceled")
	// ...
}
```

### Registering handlers

For each consumers, handler functions are registered on their `init()` functions
like in `printer/printer.go` consumer:

```go
func init() {
	h := bus.Handler{Handle: print, Matcher: ".*"}
	bus.RegisterHandler("printer", &h)
	fmt.Printf("Registered printer handler...\n")
}
```

### Emitting events

Events can be emitted on any package. As a sample two events created on
`main.go` file like:

```go
// Three order.created events
txID := monoton.Next()
for i := 0; i < 3; i++ {
	name := fmt.Sprintf("Product #%d", i)
	bus.Emit(
		"order.created",
		models.Order{Name: name, Amount: randomAmount()},
		txID,
	)
}

// One order.canceled event
bus.Emit(
	"order.canceled",
	models.Order{Name: "Product #N", Amount: randomAmount()},
	monoton.Next(),
)
```

### Execution

Execute the program:

```shell
go run main.go
```

### Outputs

The execution of the emitting will result similar output:

**On load:**

```shell
Registered calculator handler...
Registered counter handler...
Registered printer handler...
```

**After emitting events:**

```shell
Event for order.created: &{ID:0RPwZrc400010001 TxID:0RPwZrc400000001 Topic:0xc00009a090 Data:{Name:Product #0 Amount:51} OccurredAt:1557375256628182000}


Event for order.created: &{ID:0RPwZrc400020001 TxID:0RPwZrc400000001 Topic:0xc00009a090 Data:{Name:Product #1 Amount:97} OccurredAt:1557375256628257000}


Event for order.created: &{ID:0RPwZrc400030001 TxID:0RPwZrc400000001 Topic:0xc00009a090 Data:{Name:Product #2 Amount:57} OccurredAt:1557375256628283000}


Event for order.canceled: &{ID:0RPwZrc400040001 TxID:0RPwZrc400050001 Topic:0xc00009a210 Data:{Name:Product #N Amount:39} OccurredAt:1557375256628348000}

You should see 4 events printed above!^^^
Total evet count for order.canceled: 1
Total evet count for order.created: 3
Order total amount 166
```

## License

Apache License 2.0

Copyright (c) 2019 Mustafa Turan
