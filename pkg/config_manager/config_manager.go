package config_manager

import (
	"encoding/json"
	"os"
)

type JWTConfig struct {
	SecretKey string
}

type DBConfig struct {
	Driver   string
	User     string
	Password string
	Database string
}

func GetDbConfig(path string) (*DBConfig, error) {
	// наверно логику определения пути лучше сюда положить. Написать типа доки. Что конфиги
	// должны лежать по определенному пути с определенным названием.
	// Или сделать структуру в свойствах которой будет путь к папке
	file, err := os.Open(path)
	if nil != err {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

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
