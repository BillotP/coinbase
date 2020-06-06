package coinbase

import (
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
	os.Exit(m.Run())
}

func TestGetAccounts(t *testing.T) {
	var pub = creds.Account[0].Pub
	var priv = creds.Account[0].Priv
	var foo = New(&pub, &priv)
	res, fii := foo.GetAccounts()
	if fii != nil {
		t.Errorf("Error : %s\n", fii.Error())
	}
	if res == nil {
		t.Errorf("Error want models.Accounts got %v", res)
	}
	t.Log("Got accounts ", res)
}

func TestGetTransactionsByAccountID(t *testing.T) {
	var pub = creds.Account[0].Pub
	var priv = creds.Account[0].Priv
	var foo = New(&pub, &priv)
	res, fii := foo.GetAccounts()
	if fii != nil {
		log.Fatal(fii)
	}
	fon, fuu := foo.GetTransactionsByAccountID(res.Datas[0].ID)
	if fuu != nil {
		t.Errorf("Error : %s\n", fuu.Error())
	}
	t.Log("Got account transaction ", fon)
}
