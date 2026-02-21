package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/SXsid/secrets-cli/internal/encrypt"
	"github.com/SXsid/secrets-cli/internal/vault"
	"golang.org/x/term"
)

func Verify() (*vault.Vault, error) {
	fmt.Print("Enter password: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	password := string(bytePassword)
	cnfg, filePointer, err := vault.LoadConfig()
	if err != nil {
		return nil, err
	}
	if cnfg.HashPassword != "" {
		if ok := encrypt.CheckPasswordHash(password, cnfg.HashPassword); !ok {
			return nil, errors.New("invalid creds")
		}
	} else {
		hashPassword, err := encrypt.HashPassword(password)
		if err != nil {
			return nil, err
		}
		cnfg.HashPassword = hashPassword
	}
	if cnfg.Salt == "" {
		salt, err := encrypt.Genrate_slat()
		if err != nil {
			return nil, err
		}
		cnfg.Salt = salt

	}
	if err := vault.DumpConfig(*cnfg, filePointer); err != nil {
		return nil, err
	}
	secretKey, err := encrypt.Derive_key(cnfg.HashPassword, cnfg.Salt)
	if err != nil {
		return nil, err
	}

	vaultCfg, err := vault.NewValut(secretKey)
	if err != nil {
		return nil, err
	}
	return vaultCfg, nil
}
