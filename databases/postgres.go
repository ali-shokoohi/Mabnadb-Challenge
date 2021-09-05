package databases

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetPostgres - Return a valid database client
func GetPostgres() *gorm.DB {
	var (
		host     = os.Getenv("postgres_host")
		port     = os.Getenv("postgres_port")
		user     = os.Getenv("postgres_user")
		password = os.Getenv("postgres_password")
		dbname   = os.Getenv("postgres_dbname")
		sslmode  = os.Getenv("postgres_sslmode")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		host, port, user, password, dbname, sslmode)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	// Check if it failed
	if err != nil {
		log.Fatalln("Can't connect to database:", err.Error())
	}
	// Every things os good!
	log.Println("Successfully connected to database!")

	// Return database client
	return db
}
