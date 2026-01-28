package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SXsid/secrets-cli/valut"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	vault := valut.NewValut(os.Getenv("secret_key"))
	if err := vault.Set("key1", "value1"); err != nil {
		log.Fatal(err)
	}

	if err := vault.Set("key1", "value2"); err != nil {
		log.Fatal(err)
	}
	res, err := vault.Get("key1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", res)
}
