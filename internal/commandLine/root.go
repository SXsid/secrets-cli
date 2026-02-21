package cli

import (
	"fmt"

	"github.com/SXsid/secrets-cli/internal/vault"
)

func Get(cfg *vault.Vault, flags *CliFlags) {
	if flags.Key == "" {
		fmt.Println("Error: -k is missing \n\t value get -k <key>")
		return
	}
	data, err := cfg.Get(flags.Key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func Set(cfg *vault.Vault, flags *CliFlags) {
	if flags.Key == "" || flags.Value == "" {
		fmt.Println("Error: -k  or -v is missing \n\t value set -k <key> -v <value>")
		return
	}
	err := cfg.Set(flags.Key, flags.Value)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func List(cfg *vault.Vault) {
	keys := cfg.List()
	fmt.Println("All Avai")
	for _, k := range keys {
		fmt.Printf("\t%s\n", k)
	}
}
