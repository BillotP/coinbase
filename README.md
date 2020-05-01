# Go-Coinbase - Go Coinbase API unofficial Client


## Getting started

Grab this dep 

```bash
go get "github.com/BillotP/go-coinbase"
```

Write your api key and secrets in a `.env` file for example
```env
COINBASE_APIKEY=myawesomeapikey
COINBASE_APISECRET=mysuperapisecret
```

And then in your code 

```go
package main

import (
    "fmt"
    "log"
    "github.com/BillotP/go-coinbase"
)

func main() {
    // Load client
    var coinbaseClient = coinbase.CoinbaseClient
    // Get all your accounts
    if myaccounts, err := coinbaseClient.Accounts(); err != nil {
        log.Fatal(err)
    }
    // List them
    for _, account := range myaccounts {
        fmt.Printf("%s: %s (%s %s)\n", 
            myaccounts.Balance.Currency, 
            myaccounts.Balance.Amount, 
            myaccounts.NativeBalance.Amount, 
            myaccounts.NativeBalance.Currency,
        )
    }
}
```

And finally run the whole thing 

```bash
source .env
COINBASE_APIKEY=$COINBASE_APIKEY COINBASE_APISECRET=$COINBASEAPISECRET go run .
```

Et voila !