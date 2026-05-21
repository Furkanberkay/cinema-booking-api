package main

import (
	"log"

	"github.com/furkanberkay/cinema-api/internal/booking"
)

func main() {

	db, err := booking.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
