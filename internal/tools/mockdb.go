package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"bob": {
		AuthToken: "456DEF",
		Username:  "bob",
	},
	"charlie": {
		AuthToken: "789GHI",
		Username:  "charlie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"bob": {
		Coins:    200,
		Username: "bob",
	},
	"charlie": {
		Coins:    300,
		Username: "charlie",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate database latency
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate database latency
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}
