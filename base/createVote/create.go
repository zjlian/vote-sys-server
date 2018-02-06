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

func addOptionToTab(O *Option, vcode string) error {
	var (
		sqlstr string
		err    error
	)
	if !isLegalOption(O) {
		return errors.New("投票选项为空")
	}

	sqlstr = "voteSys.Option (vcode, content) values (?,?)"
	_, err = dbtool.DB.Insert(sqlstr, vcode, O.Content)
	return err
}

func addVoteToTab(V *Vote) error {
	var (
		sqlstr   string
		randCode string
		err      error
	)
	if yes, err := isLegalVote(V); !yes {
		return err
	}

	randCode = rand.GetRS16()
	sqlstr = "voteSys.Vote (uid, vcode, title, `describe`, selectType, createTime, deadline, location) values (?,?,?,?,?,?,?,?)"
	_, err = dbtool.DB.Insert(sqlstr,
		V.UID, randCode, V.Title, V.Describe, V.SelectType, V.CreateTime, V.Deadline, V.Location)

	if err != nil {
		return err
	}

	for _, v := range V.Options {
		addOptionToTab(&v, randCode)
	}
	return err
}

func create(V *Vote) error {
	dbtool.Init()
	err := addVoteToTab(V)
	return err
}

// Create 创建投票
func Create(data string) error {
	var vote Vote
	jsonToVote(data, &vote)
	return create(&vote)
}
