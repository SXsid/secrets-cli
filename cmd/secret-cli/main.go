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
	vault, err := valut.NewValut(os.Getenv("secret_key"), os.Getenv("file_path"))
	if err != nil {
		log.Fatal(err)
	}
	defer vault.Close()
	data, err := vault.Get("key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
