package encrypt

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func encode_into_base64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func decode_into_base64(data string) []byte {
	res, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return res
}

func Genrate_slat() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return encode_into_base64(salt), nil
}

func Derive_key(userPassword, salt string) ([]byte, error) {
	return scrypt.Key([]byte(userPassword), decode_into_base64(salt), 32768, 8, 1, 32)
}

func vaultHelp() {
	fmt.Printf("vault set -k <key> -v <value>\n")
	fmt.Printf("vault get -k\n")
}
