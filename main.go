package main

import (
	"coinbase/lib/models"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var err error
	var coinbase = CoinbaseClient
	var value *models.Accounts
	var nativeSum float64
	var nativeCurrency string
	if value, err = coinbase.GetAccounts(); err != nil {
		log.Fatal(err)
	}
	for _, account := range value.Datas {
		if !strings.HasPrefix(account.Balance.Amount, "0.0") {
			nativeCurrency = account.NativeBalance.Currency
			fmt.Printf("Coinbase | [%s] balance [%s] (%s %s)\n",
				account.Balance.Currency,
				account.Balance.Amount,
				account.NativeBalance.Amount,
				nativeCurrency,
			)
			val, _ := strconv.ParseFloat(account.NativeBalance.Amount, 64)
			nativeSum += val
		}
	}
	fmt.Printf("\n\tCoinbase Total Balance : %v %s\n\n", nativeSum, nativeCurrency)
}
