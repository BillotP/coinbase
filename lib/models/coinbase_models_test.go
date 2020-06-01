package models

import "testing"

func TestFilterEmpty(t *testing.T) {
	var testV = Accounts{
		Datas: []Account{
			Account{
				Balance: Balance{Amount: "0.000"},
			},
			Account{
				Balance: Balance{Amount: "0.001"},
			},
		},
	}
	testV.FilterEmpty()
	if len(testV.Datas) != 1 {
		t.Errorf("Error FilterEmpty want 0 got %v", len(testV.Datas))
	}
}
