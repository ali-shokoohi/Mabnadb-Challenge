package cli

import (
	"log"
	"math/rand"
	"time"

	"gitlab.com/shokoohi/mabnadp-challenge/databases"
	"gitlab.com/shokoohi/mabnadp-challenge/models"
)

// AutoCreate - Create multi objects automatically with limit
func AutoCreate(max uint) {
	// Get database client
	db := databases.GetPostgres()
	// First Create a instrument object for parent of trades
	instrument := models.Instrument{Name: "AutoCreated"}
	db.Create(&instrument)
	var trades []models.Trade
	now := time.Now()
	rand.Seed(now.UnixNano())
	count := 0
	for count < int(max) {
		trade := models.Trade{InstrumentID: instrument.ID, Open: rand.Float32(),
			Close: rand.Float32(), High: rand.Float32(), Low: rand.Float32(), DateEn: now}
		trades = append(trades, trade)
		count++
	}
	db.CreateInBatches(&trades, int(max))
	log.Printf("Created %d trades!", max)
}
