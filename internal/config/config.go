/*
	My solutions are commented out.
	They work, but I wanted to use the instructor solutions for reference.
*/

package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(*c)
}

// func Read() (Config, error) {
// 	filePath, err := getConfigFilePath()
// 	if err != nil {
// 		return Config{}, err
// 	}

// 	file, err := os.ReadFile(filePath)

// 	if err != nil {
// 		return Config{}, err
// 	}

// 	config := Config{}
// 	err = json.Unmarshal(file, &config)
// 	if err != nil {
// 		return Config{}, err
// 	}

// 	return config, nil
// }

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}

	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFileName)
	return fullPath, nil
}

// func write(cfg Config) error {
// 	data, err := json.Marshal(cfg)
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}

// 	filePath, err := getConfigFilePath()

// 	if err != nil {
// 		return err
// 	}

// 	err = os.WriteFile(filePath, data, 0666)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)

	if err != nil {
		return err
	}
	return nil
}
