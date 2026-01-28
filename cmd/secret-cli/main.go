package main

import (
	"fmt"
	"log"

	"github.com/SXsid/secrets-cli/encrypt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	res, err := encrypt.Encrypt("this is the data")
	if err != nil {
		panic(err)
	}
	fmt.Printf("text=>%s\n nonce=>%s\n", res.CihpherText, res.Nonce)
	orignalString, err := encrypt.Decrypt(*res)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", orignalString)
}
