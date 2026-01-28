package valut

import (
	"encoding/json"
	"io"
	"os"
)

func (v *Vault) loadKeyValues() error {
	file, err := os.OpenFile(v.file_path, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&v.data)
	if err != nil && err != io.EOF {
		return err
	}
	v.file_pointer = file

	return nil
}

func (v *Vault) Close() error {
	if err := v.file_pointer.Close(); err != nil {
		return err
	}
	return nil
}

func (v *Vault) dumpKeyValues() error {
	if err := v.file_pointer.Truncate(0); err != nil {
		return err
	}
	if _, err := v.file_pointer.Seek(0, 0); err != nil {
		return err
	}
	encoder := json.NewEncoder(v.file_pointer)
	if err := encoder.Encode(v.data); err != nil {
		return err
	}
	return nil
}
