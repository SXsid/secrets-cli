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
	filePath    string
	fileData    FileData
	filePointer *os.File
}

func valutUsage() {
	fmt.Printf("vault init -f <full_file_path> -p <your_password>\n")
	fmt.Printf("\t -f : file path where you keys will live\n")
	fmt.Printf("\t -p : you vault password\n")
	fmt.Println("vault help: to get all avaible command ")
}

func NewValut(password, filePath string) (*Vault, error) {
	if password == "" || filePath == "" {
		valutUsage()
		return nil, errors.New("password and file path are required")
	}
	Vault := &Vault{
		filePath: filePath,
		fileData: FileData{
			Data: make(map[string]encrypt.EncryptedRecord),
		},
	}
	if err := Vault.loadKeyValues(); err != nil {
		return nil, err
	}
	if salt := Vault.fileData.Salt; salt == "" {
		salt, err := encrypt.Genrate_slat()
		if err != nil {
			return nil, err
		}
		Vault.fileData.Salt = salt

	}
	secretKey, err := encrypt.Derive_key(password, Vault.fileData.Salt)
	if err != nil {
		return nil, err
	}
	Vault.key = secretKey
	return Vault, nil
}

func (v *Vault) Set(key, value string) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	encrypted_value, err := encrypt.Encrypt(value, v.key)
	if err != nil {
		return nil
	}
	v.fileData.Data[key] = *encrypted_value
	if err = v.dumpKeyValues(); err != nil {
		return err
	}
	return nil
}

func (v *Vault) Get(key string) (string, error) {
	v.mu.Lock()
	defer v.mu.Unlock()
	store_value, ok := v.fileData.Data[key]
	if !ok {
		return "", fmt.Errorf("no data for this %s key", key)
	}
	res, err := encrypt.Decrypt(store_value, v.key)
	if err != nil {
		return "", err
	}

	return res, nil
}
