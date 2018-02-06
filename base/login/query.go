package login

import "voting-system-api/tool/dbtool"

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

// HasUser 检查是否存在该用户（需要先调用依赖的数据库包dbtool.Init()）
func HasUser(uid uint64) bool {
	return hasUserByUID(uid)
}
