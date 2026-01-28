package encrypt

import (
	"crypto/sha256"
	"encoding/base64"
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

func derive_key(secret_key string) []byte {
	hash := sha256.Sum256([]byte(secret_key))
	if len(hash) == 0 {
		panic("Invlaid secret_key")
	}
	return hash[:]
}
