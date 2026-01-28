package valut

import (
	"fmt"
	"os"

	"github.com/SXsid/secrets-cli/encrypt"
)

type Vault struct {
	key          string
	file_path    string
	data         map[string]encrypt.EncryptedRecord
	file_pointer *os.File
}

func NewValut(key, file_path string) (*Vault, error) {
	if key == "" || file_path == "" {
		panic("env is not set")
	}
	vault := &Vault{
		key:       key,
		file_path: file_path,
		data:      map[string]encrypt.EncryptedRecord{},
	}
	if err := vault.loadKeyValues(); err != nil {
		return nil, err
	}
	return vault, nil
}

func (v *Vault) Set(key, value string) error {
	encrypted_value, err := encrypt.Encrypt(value, v.key)
	if err != nil {
		return nil
	}
	v.data[key] = *encrypted_value
	if err = v.dumpKeyValues(); err != nil {
		return err
	}
	return nil
}

func (v *Vault) Get(key string) (string, error) {
	store_value, ok := v.data[key]
	if !ok {
		return "", fmt.Errorf("no data for this %s key", key)
	}
	res, err := encrypt.Decrypt(store_value, v.key)
	if err != nil {
		return "", err
	}

	return res, nil
}
