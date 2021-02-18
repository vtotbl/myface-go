package config_manager

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

type JWTConfig struct {
	SecretKey string
}

type DBConfig struct {
	Development struct {
		Open string `yaml:"open,omitempty"`
	}
}

func (db *DBConfig) GetDSN() string {
	return db.Development.Open
}

func GetDbConfig(path string) (*DBConfig, error) {
	file, err := os.Open(path)
	if nil != err {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)

	configuration := DBConfig{}
	err = decoder.Decode(&configuration)
	if nil != err {
		return nil, err
	}
	return &configuration, nil
}

func GetJWTConfig(path string) (*JWTConfig, error) {
	file, err := os.Open(path)
	if nil != err {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	configuration := JWTConfig{}
	err = decoder.Decode(&configuration)
	if nil != err {
		return nil, err
	}

	return &configuration, nil
}
