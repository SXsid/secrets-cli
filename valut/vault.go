package valut

import (
	"fmt"

	"github.com/SXsid/secrets-cli/encrypt"
)

type Vault struct {
	key  string
	data map[string]encrypt.EncryptedRecord
}

func NewValut(key string) *Vault {
	if key == "" {
		panic("env is not set")
	}
	return &Vault{
		key:  key,
		data: map[string]encrypt.EncryptedRecord{},
	}
}

func (v *Vault) Set(key, value string) error {
	encrypted_value, err := encrypt.Encrypt(value, v.key)
	if err != nil {
		return nil
	}
	v.data[key] = *encrypted_value
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
