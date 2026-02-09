package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type EncryptedRecord struct {
	CihpherText string
	Nonce       string
}

func Encrypt(data string, secretKey []byte) (*EncryptedRecord, error) {
	plain_text := []byte(data)
	block, err := aes.NewCipher(secretKey)
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

func Decrypt(data EncryptedRecord, secretKey []byte) (string, error) {
	block, err := aes.NewCipher(secretKey)
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
