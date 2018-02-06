package dbtool

import (
	"log"
	"voting-system-api/base/db"
	"voting-system-api/tool/dbtool/dbconfig"
)

// DB 数据库操作对象
var DB mysql.MYDB

// Init 连接准备好的数据库
func Init() {
	err := DB.NewConnect(dbconfig.Host, dbconfig.Username, dbconfig.Password, dbconfig.DBName)

	if err != nil {
		log.Println(err)
	}
}

// Close 关闭数据库链接
func Close() {
	DB.Close()
}
