package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SXsid/secrets-cli/internal/vault"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	vault, err := vault.NewValut(os.Getenv("secret_key"), os.Getenv("file_path"))
	if err != nil {
		log.Fatal(err)
	}
	data, err := vault.Get("sid-aws")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
