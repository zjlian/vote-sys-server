package voteQuery

import (
	"encoding/json"
	"errors"
	"voting-system-api/base/createVote"
	"voting-system-api/tool/dbtool"
)

// hasVoteByVcode 检查是否存在对应vcode的投票，调用该函数之前需要创建一个数据库链接
func hasVoteByVcode(vcode string) bool {
	var (
		err error
		id  uint64
	)
	row := dbtool.DB.QueryRow("select id from voteSys.Vote where vcode=?", vcode)
	err = row.Scan(&id)
	if err != nil {
		return false
	}
	return true
}

// HasVote 检查是否存在对应vcode的投票，
// **注意**调用该函数会进行一次数据库连接操作，并在return后关闭数据库
func HasVote(vcode string) bool {
	dbtool.Init()
	defer dbtool.Close()
	return hasVoteByVcode(vcode)
}

// getVoteOptions 通过vode获取投票的所有选项， 返回一个Option结构的切片数组
func getVoteOptions(vcode string) ([]createVote.Option, error) {
	var (
		err    error
		tmp    createVote.Option
		result []createVote.Option
		sqlstr string
	)

	sqlstr = "select id,content from voteSys.Option where vcode=?"
	rows, err := dbtool.DB.Query(sqlstr, vcode)
	if err != nil {
		return nil, err
	}
	//log.Println("queryVote.go: 41", err)

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.Content)
		if err != nil {
			break
		}
		tmp.Vcode = vcode
		result = append(result, tmp)
	}
	defer rows.Close()
	return result, err
}

// GetVoteInfo 同过vcode，从数据库查询该投票的详细信息，并返回Vote结构对象
func GetVoteInfo(vcode string) (createVote.Vote, error) {
	var (
		err    error
		vote   createVote.Vote
		sqlstr string
	)

	dbtool.Init()
	defer dbtool.Close()

	sqlstr = "select uid,vcode,title,`describe`,selectType,createTime,deadline,location from voteSys.Vote where vcode=?"
	row := dbtool.DB.QueryRow(sqlstr, vcode)
	row.Scan(&vote.UID, &vote.Vcode, &vote.Title, &vote.Describe,
		&vote.SelectType, &vote.CreateTime, &vote.Deadline, &vote.Location)

	vote.Options, err = getVoteOptions(vcode)

	return vote, err
}

// VoteInfoToJSON 传入参数为含和vcode属性的map对象，将查询结果放入Vote对象中 并转为json字符串返回
func VoteInfoToJSON(args interface{}) (interface{}, error) {
	var (
		err    error
		vcode  string
		vote   createVote.Vote
		result []byte
	)

	data, ok := args.(map[string]interface{})
	if !ok {
		return nil, errors.New("queryInfo.go VoteInfoToJSON 请求参数的类型不是 map[string]interface{}")
	}

	_, ok = data["vcode"]
	if !ok {
		return nil, errors.New("查询请求缺失参数 vcode")
	}

	vcode, ok = data["vcode"].(string)
	if !ok {
		return nil, errors.New("查询请求参数类型异常 vcode 无法转换字符串类型")
	}

	if !HasVote(vcode) {
		return nil, errors.New("不存在的投票")
	}

	vote, err = GetVoteInfo(vcode)
	if err != nil {
		return nil, err
	}

	result, err = json.Marshal(vote)
	return result, err
}
