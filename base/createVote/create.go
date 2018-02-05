package createVote

import (
	"encoding/json"
	"errors"
	"voting-system-api/tool/dbtool"
	"voting-system-api/tool/rand"
)

// byJSON 通过json字符串创建Vote实例
func jsonToVote(jsonStr string, V *Vote) error {
	err := json.Unmarshal([]byte(jsonStr), V)

	return err
}

func addIntoTabVote(V *Vote) error {
	var (
		sqlstr   string
		randCode string
		err      error
	)
	if !isLegalVote(V) {
		return errors.New("投票数据不合法")
	}

	randCode = rand.GetRS16()
	sqlstr = "voteSys.Vote (vcode, title, `describe`, selectType, createTime, deadline, location) values (?,?,?,?,?,?,?)"
	_, err = dbtool.DB.Insert(sqlstr,
		randCode, V.Title, V.Describe, V.SelectType, V.CreateTime, V.Deadline, V.Location)

	return err
}

func create(V *Vote) {
	dbtool.Init()
}

// Create 创建投票
func Create(data string) error {
	var (
		vote Vote
		err  error
	)
	jsonToVote(data, &vote)
	err = create(&vote)
	return err
}
