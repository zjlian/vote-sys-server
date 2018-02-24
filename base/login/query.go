package login

import "voting-system-api/tool/dbtool"

// HasUser 检查是否存在该用户，调用该函数前需要创建一个数据库连接
func hasUserByUID(uid uint64) bool {
	var loginName string

	row := dbtool.DB.QueryRow(
		"select login from voteSys.Ass where uid=?",
		uid,
	)
	err := row.Scan(&loginName)
	//log.Println(loginName)
	if err != nil {
		return false
	}

	// if loginName != u.Login {
	// 	return false
	// }
	return true
}

// HasUser 检查是否存在该用户
// **注意**调用该函数会进行一次数据库连接操作，并在return后关闭数据库
func HasUser(uid uint64) bool {
	dbtool.Init()
	defer dbtool.Close()
	return hasUserByUID(uid)
}
