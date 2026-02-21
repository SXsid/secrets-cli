package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	cli "github.com/SXsid/secrets-cli/internal/commandLine"
)

func printASCII() {
	cmd := exec.Command("figlet", "-f", "ANSIShadow", "vault")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Welcome to VAULT")
	} else {
		fmt.Println(string(output))
	}
}

func usage() {
	printASCII()
	fmt.Println("vault init -f <file_path> -p <user_password>")
	fmt.Println("vault set -k <key> -v <value>")
	fmt.Println("vault get -k <key>")
	fmt.Println("vault ls")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Invalid Commands")
		usage()
		return
	}
	cliFlags := cli.CliParser(args)
	Vaultconfig, err := cli.Verify()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cliFlags)
	switch args[1] {
	case "get":
		cli.Get(Vaultconfig, cliFlags)
	case "set":
		cli.Set(Vaultconfig, cliFlags)

	case "ls":
		cli.List(Vaultconfig)
	default:
		usage()

	}
}
