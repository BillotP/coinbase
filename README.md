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
	"os"

	"github.com/BillotP/coinbase"
	"github.com/BillotP/coinbase/lib/models"
)

func main() {
	var (
		err        error
		pub        = os.Getenv("COINBASE_APIKEY")
		priv       = os.Getenv("COINBASE_APISECRET")
		myaccounts *models.Accounts
		// Load client
		client = coinbase.New(&pub, &priv)
	)
	// Get all your accounts
	if myaccounts, err = client.GetAccounts(); err != nil {
		log.Fatal(err)
	}
	// List them
	for _, account := range myaccounts.Datas {
		fmt.Printf("%s: %s (%s %s)\n",
			account.Balance.Currency,
			account.Balance.Amount,
			account.NativeBalance.Amount,
			account.NativeBalance.Currency,
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