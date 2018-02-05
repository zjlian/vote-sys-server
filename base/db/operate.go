package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// MYDB 进行 MySql 相关的操作
type MYDB struct {
	Base     *sql.DB
	User     string
	Password string
	Host     string
	DBName   string
}

// NewConnect 新建一个 MySql 数据库连接
func (db *MYDB) NewConnect(host, username, password, dbname string) error {
	db.User, db.Password = username, password
	db.Host, db.DBName = host, dbname

	return db.Connect()
}

// Connect 启动数据库连接， 前提树已经调用过NewConnect方法的结构体才能使用
func (db *MYDB) Connect() error {
	var (
		e        error
		conninfo string
	)
	conninfo = db.User + ":" + db.Password + "@/" + db.DBName + "?charset=utf8"
	db.Base, e = sql.Open("mysql", conninfo)
	return e
}

// Close 关闭数据库链接
func (db *MYDB) Close() {
	db.Base.Close()
}

func (db *MYDB) execSQL(ss string, args ...interface{}) (sql.Result, error) {
	var (
		res sql.Result
		err error
	)

	stmt, err := db.Base.Prepare(ss)
	checkErr(err)

	res, err = stmt.Exec(args...)
	checkErr(err)

	return res, err
}

// Insert 执行一个数据库简单插入操作，传入的sql语句中不用写开头的 INSERT 命令
func (db *MYDB) Insert(sql string, args ...interface{}) (sql.Result, error) {
	return db.execSQL("insert into "+sql, args...)
}

// Update 执行一个简单的数据库更新操作
func (db *MYDB) Update(sql string, args ...interface{}) (sql.Result, error) {
	return db.execSQL("update "+sql, args...)
}

// Delete 执行一个简单的数据库删除操作
func (db *MYDB) Delete(sql string, args ...interface{}) (sql.Result, error) {
	return db.execSQL("delete from "+sql, args...)
}

// Query 执行一个查询操作
func (db *MYDB) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return db.Base.Query(sql, args...)
}

// QueryRow 执行一个查询操作
func (db *MYDB) QueryRow(sql string, args ...interface{}) *sql.Row {
	return db.Base.QueryRow(sql, args...)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
