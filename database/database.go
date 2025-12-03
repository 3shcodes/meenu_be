package database

import (
	sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
)



func CreateInstance(dbConf string) *sql.DB {

	db, err := sql.Open("mysql", dbConf);
	if err != nil {
		panic(err)
	}

	return db;

}
