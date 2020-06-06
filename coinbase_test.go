package coinbase

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

type account struct {
	Name string
	Pub  string
	Priv string
}

type accounts struct {
	Account []account
}

var creds accounts
var defcli *Client

func TestMain(m *testing.M) {
	var (
		err error
		dat []byte
	)
	if dat, err = ioutil.ReadFile("config.toml"); err != nil {
		log.Fatal(err)
	}
	if _, err := toml.Decode(string(dat), &creds); err != nil {
		log.Fatal(err)
	}
	var pub = creds.Account[0].Pub
	var priv = creds.Account[0].Priv
	defcli = New(&pub, &priv)
	os.Exit(m.Run())
}

func TestGetSpotPrice(t *testing.T) {
	res, err := defcli.GetSpotPrice("BTC", "EUR")
	if err != nil {
		t.Errorf("Error : %s\n", err.Error())
	}
	t.Log("Got SpotPrice ", res)
	res, err = defcli.GetSpotPrice("YLO", "KLKLL")
	if err == nil {
		t.Errorf("Error want error got nil")
	}
	t.Log("Got err ", err)

}

func TestGetAccounts(t *testing.T) {
	res, err := defcli.GetAccounts()
	if err != nil {
		t.Errorf("Error : %s\n", err.Error())
	}
	if res == nil {
		t.Errorf("Error want models.Accounts got %v", res)
	}
	t.Log("Got accounts ", res)
}

func TestGetTransactionsByAccountID(t *testing.T) {
	res, fii := defcli.GetAccounts()
	if fii != nil {
		log.Fatal(fii)
	}
	fon, fuu := defcli.GetTransactionsByAccountID(res.Datas[0].ID)
	if fuu != nil {
		t.Errorf("Error : %s\n", fuu.Error())
	}
	t.Log("Got account transaction ", fon)
}

func TestGetNewAccountAddress(t *testing.T) {
	r, err := defcli.GetAccounts()
	if err != nil {
		log.Fatal(err)
	}
	ltcAcc := r.Get("LTC")
	if ltcAcc == nil {
		log.Fatal(fmt.Errorf("Failed to get LTC Account ID"))
	}
	res, err := defcli.GetNewAccountAddress(ltcAcc.ID)
	if err != nil {
		t.Errorf("Error : %s\n", err.Error())
	}
	t.Log("Got account address ", res)
	res, err = defcli.GetNewAccountAddress("invalid")
	if err == nil {
		t.Errorf("Error want error got nil")
	}
	t.Log("Got error ", err)
}
