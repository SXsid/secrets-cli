package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/SXsid/secrets-cli/internal/commandLine"
)

func usage() {
	fmt.Println("vault init -f<file_path> -p<user_password> ")
	fmt.Println("vault set -k<key> -v <value>")
	fmt.Println("vault get -k")
	fmt.Println("vault ls")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Invalid Commands")
		usage()
		return
	}

	switch args[1] {
	case "init":
		cli.Init()
	case "get":
		Vaultconfig, err := cli.Verify()
		if err != nil {
			log.Fatal(err)
		}
		cli.Get(Vaultconfig)
	case "set":
		Vaultconfig, err := cli.Verify()
		if err != nil {
			log.Fatal(err)
		}
		cli.Set(Vaultconfig)

	case "ls":
		Vaultconfig, err := cli.Verify()
		if err != nil {
			log.Fatal(err)
		}
		cli.List(Vaultconfig)
	default:
		usage()

	}
}
