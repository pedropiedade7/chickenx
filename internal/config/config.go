package config

import (
	"log"

	"github.com/pedropiedade7/chickenx/internal/database"
)

func Run() {
	dsn := "postgres://admin:admin@localhost:5433/restaurante?sslmode=disable"

	db, err := database.NewPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}

}
