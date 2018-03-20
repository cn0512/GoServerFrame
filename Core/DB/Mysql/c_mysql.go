package Mysql

/*
	use xorm lib
*/

import (
	"MSvrs/Config"
	"MSvrs/Core/Utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func init() {

}

func mysqlEngine() (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", Config.GetMysqlSrc(0))
}

func GetMysql() (engine *xorm.Engine, err error) {
	engine, err = mysqlEngine()
	if err != nil {
		Utils.Logout("GetMysql err:%v\n", err)
		return nil, err
	}
	return engine, nil
}
