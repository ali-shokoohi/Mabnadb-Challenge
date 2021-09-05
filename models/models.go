package models

import (
	"gitlab.com/shokoohi/mabnadp-challenge/databases"
	"gitlab.com/shokoohi/mabnadp-challenge/utils"
)

// LoadModels - Doing some staff such as: fill default values, auto migrations and etc...
func LoadModels() {
	// Get database client from our databases package
	db := databases.GetPostgres()
	// Auto migrations
	db.AutoMigrate(&Instrument{}, &Trade{})
	// Make some Instrument objects
	instruments := []Instrument{{Name: "AAPL"}, {Name: "GOOGLE"}}
	// Submit instruments to the database
	db.Create(&instruments)
	// Make some Trade objects
	trades := []Trade{
		{Open: 1001, High: 2001, Low: 301, Close: 4001, DateEn: utils.ParseTime("2006-01-02", "2020-01-01"), InstrumentID: instruments[0].ID},
		{Open: 1002, High: 2002, Low: 302, Close: 4002, DateEn: utils.ParseTime("2006-01-02", "2020-01-02"), InstrumentID: instruments[0].ID},
		{Open: 1003, High: 2003, Low: 303, Close: 4003, DateEn: utils.ParseTime("2006-01-02", "2020-01-03"), InstrumentID: instruments[0].ID},
		{Open: 1004, High: 2004, Low: 304, Close: 4004, DateEn: utils.ParseTime("2006-01-02", "2020-01-01"), InstrumentID: instruments[1].ID},
		{Open: 1005, High: 2005, Low: 305, Close: 4005, DateEn: utils.ParseTime("2006-01-02", "2020-01-03"), InstrumentID: instruments[1].ID},
		{Open: 1006, High: 2006, Low: 306, Close: 4006, DateEn: utils.ParseTime("2006-01-02", "2020-01-01"), InstrumentID: instruments[1].ID},
		{Open: 1007, High: 2007, Low: 307, Close: 4007, DateEn: utils.ParseTime("2006-01-02", "2020-01-01"), InstrumentID: instruments[0].ID},
	}
	// Submit trades to the database
	db.Create(&trades)
}
