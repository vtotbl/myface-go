package mysql_db

import (
	"errors"
	"path/filepath"

	"github.com/Valeriy-Totubalin/myface-go/pkg/config_manager"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	path, err := filepath.Abs("./db/dbconf.yml")
	if nil != err {
		return nil, err
	}
	config, err := config_manager.GetDbConfig(path)
	if nil != err {
		return nil, err
	}
	db, err := gorm.Open(mysql.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Error connecting to database")
	}

	return db, nil
}
