package mysql_db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	driver   string
	user     string
	password string
	database string
}

var database *sql.DB

func GetDB() DB {
	return newDB("mysql", "root", "", "myface")
}

func newDB(driver, user, password, database string) DB {
	db := DB{driver, user, password, database}
	return db
}

// func (c *DB) Insert(table string, values string) {
// 	database, err := sql.Open(c.driver, c.user+":"+c.password+"@/"+c.database)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Println("insert into " + table + " values (" + values + ")")
// 	database.Query("insert into " + table + " values (" + values + ")")
// 	defer database.Close()
// }

// func (c *DB) Update(table string, values string) {
// 	database, err := sql.Open(c.driver, c.user+":"+c.password+"@/"+c.database)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	database.Query("update " + " " + table + " set " + values)
// 	defer database.Close()
// }

// func (c *DB) Delete(table string, values string) {
// 	database, err := sql.Open(c.driver, c.user+":"+c.password+"@/"+c.database)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	database.Query("delete from " + " " + table + " where " + values)
// 	defer database.Close()
// }

func (c *DB) Query(queryString string) *sql.Rows {
	database, err := sql.Open(c.driver, c.user+":"+c.password+"@/"+c.database)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := database.Query(queryString)
	if err != nil {
		log.Println(err)
	}
	defer database.Close()

	return rows
}
