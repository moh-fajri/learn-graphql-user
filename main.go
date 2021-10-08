package main

import (
	"github.com/moh-fajri/learn-garphql-user/repo/userdb"
	"github.com/moh-fajri/learn-garphql-user/transport"
	"log"
	"github.com/joho/godotenv"
)

func main()  {
	// check env exists
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userdb.InitDbCon()

	transport.NewService().RunGRCP()
}
