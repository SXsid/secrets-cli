package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
)

type EncryptedRecord struct {
	CihpherText string
	Nonce       string
}

func Encrypt(data string) (*EncryptedRecord, error) {
	secret_key := derive_key(os.Getenv("secret_key"))
	if len(secret_key) == 0 {
		return nil, errors.New("failed to load secrete key")
	}
	plain_text := []byte(data)
	block, err := aes.NewCipher(secret_key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	chiperText := aesgcm.Seal(nil, nonce, plain_text, nil)

	return &EncryptedRecord{
		CihpherText: encode_into_base64(chiperText),
		Nonce:       encode_into_base64(nonce),
	}, nil
}

func Decrypt(data EncryptedRecord) (string, error) {
	secret_key := derive_key(os.Getenv("secret_key"))
	if len(secret_key) == 0 {
		return "", errors.New("failed to load secrete key")
	}
	block, err := aes.NewCipher(secret_key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	res, err := aesgcm.Open(nil, decode_into_base64(data.Nonce), decode_into_base64(data.CihpherText), nil)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
