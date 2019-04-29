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

Event for user.created: &{ID:0ROyrnX900010001 TxID:0ROyrnX900000001 Topic:0xc00006a1b0 Data:{Name:MyName#0 Email:sample0@example.com} OccurredAt:1556492950691702000}


Event for user.created: &{ID:0ROyrnX900020001 TxID:0ROyrnX900000001 Topic:0xc00006a1b0 Data:{Name:MyName#1 Email:sample1@example.com} OccurredAt:1556492950691827000}


Event for user.created: &{ID:0ROyrnX900030001 TxID:0ROyrnX900000001 Topic:0xc00006a1b0 Data:{Name:MyName#2 Email:sample2@example.com} OccurredAt:1556492950691845000}


Event for user.canceled: &{ID:0ROyrnX900040001 TxID:0ROyrnX900050001 Topic:0xc00006a2a0 Data:{Name:Another Email:sample@example.com} OccurredAt:1556492950691902000}

Total evet count for user.created: 3
Total evet count for user.canceled: 0
```

## License

Apache License 2.0

Copyright (c) 2019 Mustafa Turan
