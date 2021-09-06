package models

import (
	"log"

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

// LoadQueries - Execute queries directly
func LoadQueries() {
	// Get database client from our databases package
	db := databases.GetPostgres()
	// Send queries
	// Create tables
	tx := db.Exec("CREATE TABLE IF NOT EXISTS Instrument‬‬(ID int, Name varchar(255))")
	if tx.Error != nil {
		log.Fatalln("Can't send query: ", tx.Error)
	}
	tx = db.Exec("CREATE TABLE IF NOT Exists Trade(ID int, InstrumentID int, DateEn timestamptz, Open decimal, High decimal, Low decimal, Close decimal)")
	if tx.Error != nil {
		log.Fatalln("Can't send query: ", tx.Error)
	}
	// Insert values
	insertStmt := `insert into Instrument‬‬ values(1, 'AAPL'), (2, 'GOOGLE')`
	tx = db.Exec(insertStmt)
	if tx.Error != nil {
		log.Fatalln("Can't send query: ", tx.Error)
	}
	insertStmt = `insert into "Trade" values
	(1, 1, '2020-01-01', 1001, 2001, 301, 4001)
	(1, 1, '2020-01-01', 1002, 2002, 302, 4002)
	(1, 1, '2020-01-01', 1003, 2003, 303, 4003)
	(1, 2, '2020-01-01', 1004, 2004, 304, 4004)
	(1, 2, '2020-01-01', 1005, 2005, 305, 4005)
	(1, 2, '2020-01-01', 1006, 2006, 306, 4006)
	(1, 1, '2020-01-01', 1007, 2007, 307, 4007)`
}

func GetLasts() (Instrument, Trade) {
	var instrument Instrument
	var trade Trade
	// Get database client from our databases package
	db := databases.GetPostgres()
	// Get last objects from database
	db.Preload("Trades").Last(&instrument)
	db.Last(&trade)
	// Return objects
	return instrument, trade

}
