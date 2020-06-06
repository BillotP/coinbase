package models

import "testing"

func TestFilterEmpty(t *testing.T) {
	var testV = Accounts{
		Datas: []Account{
			{
				Balance: Balance{Amount: "0.000"},
			},
			{
				Balance: Balance{Amount: "0.001"},
			},
		},
	}
	testV.FilterEmpty()
	if len(testV.Datas) != 1 {
		t.Errorf("Error FilterEmpty want 0 got %v", len(testV.Datas))
	}
}
