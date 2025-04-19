package storage

import (
	"encoding/json"
	"errors"
	"os"
)

func ReadDatabase(fileName string) (map[string]string, error) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, errors.New("database cannot be read")
	}

	if len(file) == 0 {
		return make(map[string]string), nil
	}

	database := make(map[string]string)
	err = json.Unmarshal(file, &database)

	if err != nil {
		return nil, errors.New("failed to deserialize")
	}

	return database, nil
}

func UpdateDatabase(filename string, database map[string]string) error {
	bytes, err := json.Marshal(database)
	if err != nil {
		return errors.New("failed to deserialize")
	}

	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		return errors.New("failed to write to file")
	}

	return nil
}
