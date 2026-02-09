package vault

import (
	"encoding/json"
	"io"
	"os"
)

func (v *Vault) loadKeyValues() error {
	file, err := os.OpenFile(v.filePath, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&v.fileData)
	if err != nil && err != io.EOF {
		return err
	}
	v.filePointer = file

	return nil
}

func (v *Vault) Close() error {
	if err := v.filePointer.Close(); err != nil {
		return err
	}
	return nil
}

func (v *Vault) dumpKeyValues() error {
	if err := v.filePointer.Truncate(0); err != nil {
		return err
	}
	if _, err := v.filePointer.Seek(0, 0); err != nil {
		return err
	}
	encoder := json.NewEncoder(v.filePointer)
	if err := encoder.Encode(v.fileData); err != nil {
		return err
	}
	return nil
}
