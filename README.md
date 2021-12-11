# Bus Sample Project

This is an example project to demonstrate usage of the `bus` package for
internal package communication. Please note the aim of this sample project to
guide on a un-realistic use case.

## Install

```shell
git clone https://github.com/mustafaturan/bus-sample-project
cd bus-sample-project
go get github.com/mustafaturan/monoton/v3
go get github.com/mustafaturan/bus/v3
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

**File:** [config/config.go](config/config.go)
```go
var Bus *bus.Bus

func Init() {
	// configure id generator (it doesn't have to be monoton)
	node := uint64(1)
	initialTime := uint64(0)
	m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil {
		panic(err)
	}

	// init an id generator
	var idGenerator bus.Next = (*m).Next

	// create a new bus instance
	b, err := bus.NewBus(idGenerator)
	if err != nil {
		panic(err)
	}

	// maybe register topics in here
	b.RegisterTopics("order.created", "order.canceled")

	Bus = b

	// ...
}
```

### Register topics

Assume that we have two topics which are; `order.created` and `order.canceled`.

```go
config.Bus.RegisterTopics("order.created", "order.canceled")
```

### Registering handlers

For each consumers, handler functions are registered on their `init()` functions
like in `printer/printer.go` consumer:

```go
b := config.Bus
h := bus.Handler{Handle: print, Matcher: ".*"}
b.RegisterHandler("printer", h)
fmt.Printf("Registered printer handler...\n")
```

### Emitting events

Events can be emitted on any package. As a sample, four events (two topics)
created on [main.go](main.go) file like:

```go
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
ctx = context.WithValue(ctx, bus.CtxKeyTxID, txID)

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
```

### Execution

Execute the program with race condition checks:

```shell
go run -race main.go
```

### Outputs

The execution of the emitting will result similar output:

**On load:**

```shell
Registered printer handler...
Registered counter handler...
Registered calculator handler...
```

**After emitting events:**

```shell
Event for order.created: {ID:0SVU68UR00010001 TxID:0SVU68UR00000001 Topic:order.created Source: OccurredAt:2021-04-24 01:04:26.831182 -0700 PDT m=+0.001772439 Data:{Name:Product #0 Amount:51}}


Event for order.created: {ID:0SVU68UR00020001 TxID:0SVU68UR00000001 Topic:order.created Source: OccurredAt:2021-04-24 01:04:26.831743 -0700 PDT m=+0.002333823 Data:{Name:Product #1 Amount:97}}


Event for order.created: {ID:0SVU68UR00030001 TxID:0SVU68UR00000001 Topic:order.created Source: OccurredAt:2021-04-24 01:04:26.831813 -0700 PDT m=+0.002404064 Data:{Name:Product #2 Amount:57}}


Event for order.canceled: {ID:0SVU68UR00050001 TxID:0SVU68UR00040001 Topic:order.canceled Source: OccurredAt:2021-04-24 01:04:26.831871 -0700 PDT m=+0.002462153 Data:{Name:Product #N Amount:39}}

Deregistered calculator handler...
Deregistered counter handler...
Deregistered printer handler...
Order total amount 166
Total event count for order.created: 3
Total event count for order.canceled: 1
```

## License

Apache License 2.0

Copyright (c) 2020 Mustafa Turan
