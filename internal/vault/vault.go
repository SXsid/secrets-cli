package vault

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/SXsid/secrets-cli/internal/encrypt"
)

type FileData struct {
	Salt string                             `json:"salt"`
	Data map[string]encrypt.EncryptedRecord `json:"data"`
}
type Vault struct {
	mu          sync.Mutex
	key         []byte
	fileData    map[string]encrypt.EncryptedRecord
	filePointer *os.File
}

func valutUsage() {
	fmt.Printf("vault init -f <full_file_path> -p <your_password>\n")
	fmt.Printf("\t -f : file path where you keys will live\n")
	fmt.Printf("\t -p : you vault password\n")
	fmt.Println("vault help: to get all avaible command ")
}

func NewValut(key []byte) (*Vault, error) {
	if key == nil {
		valutUsage()
		return nil, errors.New("invalid creds")
	}
	Vault := &Vault{
		fileData: make(map[string]encrypt.EncryptedRecord),
		key:      key,
	}
	if err := Vault.loadKeyValues(); err != nil {
		return nil, err
	}
	return Vault, nil
}

func (v *Vault) Set(key, value string) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	encrypted_value, err := encrypt.Encrypt(value, v.key)
	if err != nil {
		return nil
	}
	v.fileData[key] = *encrypted_value
	if err = v.dumpKeyValues(); err != nil {
		return err
	}
	return nil
}

func (v *Vault) Get(key string) (string, error) {
	v.mu.Lock()
	defer v.mu.Unlock()
	store_value, ok := v.fileData[key]
	if !ok {
		return "", fmt.Errorf("no data for this %s key", key)
	}
	res, err := encrypt.Decrypt(store_value, v.key)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (v *Vault) List() []string {
	keys := make([]string, 0, len(v.fileData))
	for k := range v.fileData {
		keys = append(keys, k)
	}
	return keys
}
