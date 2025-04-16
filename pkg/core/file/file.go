package file

import (
	"encoding/json"
	"os"
)

func ReadModelFromFile[T any](path string, model *T) (*T, error) {
	bytesArray, err := os.ReadFile(path)
	if err != nil {
		return model, err
	}
	unmarshalErr := json.Unmarshal(bytesArray, model)
	return model, unmarshalErr
}

func WriteModelIntoFile[T any](path string, model *T) error {
	byteArray, err := json.Marshal(&model)
	if err != nil {
		return err
	}
	writeErr := os.WriteFile(path, byteArray, 0644)
	return writeErr
}

func ReadFile(path string) (string, error) {
	bytesArray, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytesArray), err
}

func WriteFile(path string, str string) error {
	return os.WriteFile(path, []byte(str), 0644)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}
