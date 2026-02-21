package vault

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type ConfigData struct {
	Salt         string `json:"salt"`
	HashPassword string `json:"password"`
}

func ensureValueDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dirPath := filepath.Join(dir, ".myVault")
	if err := os.MkdirAll(dirPath, 0o755); err != nil {
		return "", err
	}

	return dirPath, nil
}

func LoadConfig() (*ConfigData, *os.File, error) {
	dirPath, err := ensureValueDir()
	if err != nil {
		return nil, nil, err
	}
	ConfigFilePath := filepath.Join(dirPath, "config.json")
	file, err := os.OpenFile(ConfigFilePath, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		file.Close()
		return nil, nil, err
	}
	decoder := json.NewDecoder(file)
	var configData ConfigData
	err = decoder.Decode(&configData)
	if err != nil && err != io.EOF {
		file.Close()
		return nil, nil, err
	}
	return &configData, file, nil
}

func DumpConfig(cnfg ConfigData, filePointer *os.File) error {
	defer filePointer.Close()
	if err := filePointer.Truncate(0); err != nil {
		return err
	}
	if _, err := filePointer.Seek(0, 0); err != nil {
		return err
	}
	encoder := json.NewEncoder(filePointer)
	if err := encoder.Encode(cnfg); err != nil {
		return err
	}
	return nil
}

func (v *Vault) loadKeyValues() error {
	dirPath, err := ensureValueDir()
	if err != nil {
		return err
	}
	DataFilePath := filepath.Join(dirPath, "data.json")
	file, err := os.OpenFile(DataFilePath, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&v.fileData)
	if err != nil && err != io.EOF {
		file.Close()
		return err
	}
	v.filePointer = file

	return nil
}

func (v *Vault) dumpKeyValues() error {
	defer v.filePointer.Close()
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
