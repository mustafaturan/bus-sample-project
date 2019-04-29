# Bus Sample Project

## Install

```shell
git clone https://github.com/mustafaturan/bus-sample-project
cd bus-sample-project
go get github.com/mustafaturan/monoton
go get github.com/mustafaturan/bus
```

## Usage

Execute the program:

```shell
go run main.go
```

Sample output:

```shell
Registered counter handler...
Registered printer handler...

Event for order.created: &{ID:0RP0M0Ih00010001 TxID:0RP0M0Ih00000001 Topic:0xc00008e060 Data:{Name:Product #0 Amount:51} OccurredAt:1556514925943127000}


Event for order.created: &{ID:0RP0M0Ih00020001 TxID:0RP0M0Ih00000001 Topic:0xc00008e060 Data:{Name:Product #1 Amount:97} OccurredAt:1556514925943265000}


Event for order.created: &{ID:0RP0M0Ih00030001 TxID:0RP0M0Ih00000001 Topic:0xc00008e060 Data:{Name:Product #2 Amount:57} OccurredAt:1556514925943283000}


Event for order.canceled: &{ID:0RP0M0Ih00040001 TxID:0RP0M0Ih00050001 Topic:0xc00008e150 Data:{Name:Product #N Amount:39} OccurredAt:1556514925943299000}

Total evet count for order.created: 3
Total evet count for order.canceled: 1
```

## License

Apache License 2.0

Copyright (c) 2019 Mustafa Turan
