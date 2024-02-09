package settings

import (
	"encoding/json"
	"os"
	"sync"
)

const (
	filename = "settings.json"
)

var singleton = sync.OnceValue(func() *Settings {
	var value Settings

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return &value
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(&value); err != nil {
		return &value
	}

	return &value
})

type Settings struct {
	APIKey string `json:"apiKey"`
}

func Get() *Settings {
	return singleton()
}

func Save() error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	if err = encoder.Encode(Get()); err != nil {
		return err
	}

	return file.Sync()
}
