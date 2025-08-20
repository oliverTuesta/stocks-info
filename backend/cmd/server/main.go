package server

import (
	"githug.com/oliverTuesta/stocks-info/backend/internal/db"
	"log"
)

func main() {
	conn, err := db.ConnectGorm()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	log.Println("DB connection established")

	err = db.Migrate(conn)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

}
