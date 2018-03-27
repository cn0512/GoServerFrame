package Mysqlx

/*
	use sqlx lib
*/

import (
	cfg "github.com/cn0512/GoServerFrame/Config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New() *sqlx.DB {

	db, err := sqlx.Open("mysql", cfg.GetMysqlSrc(0))
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(100)
	//defer db.Close()
	return db
}
