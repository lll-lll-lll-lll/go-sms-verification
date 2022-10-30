package sms_verification

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envAUTHTOKEN() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv("TWILIO_AUTHTOKEN")
}

func envSERVICESID() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv("TWILIO_SERVICES_ID")
}
