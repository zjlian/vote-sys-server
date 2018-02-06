package voteQuery

import (
	"voting-system-api/tool/dbtool"
)

// VoteBaseInfo 投票基本信息结构
type VoteBaseInfo struct {
	//创建投票的用户的id
	//UID   uint64 `json:"uid"`
	Vcode string `json:"vcode"`
	Title string `json:"title"`
	//投票参与人数
	Voters uint64 `json:"voters"`
}

func allVcodeByUID(uid uint64) ([]string, error) {
	var (
		sqlstr string
		err    error
		tmp    string
		result []string
	)
	sqlstr = "select vcode from voteSys.Vote where uid=?"
	rows, err := dbtool.DB.Query(sqlstr, uid)
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&tmp)
		if err != nil {
			break
		}
		result = append(result, tmp)
	}
	if err != nil {
		return nil, err
	}

	return result, err
}

// baseInfo 根据投票的 vcode 获取投票的基本信息
func baseInfo(vcode string) VoteBaseInfo {
	var (
		sqlstr string
		err    error
		vbi    VoteBaseInfo
	)
	sqlstr = "select vcode,title from voteSys.Vote where vcode=?"
	row := dbtool.DB.QueryRow(sqlstr, vcode)
	row.Scan(&vbi.Vcode, &vbi.Title)

	sqlstr = "select count(id) from voteSys.votingLog group by uid where vcode=?"
	row = dbtool.DB.QueryRow(sqlstr, vcode)
	err = row.Scan(&vbi.Voters)
	if err != nil {
		vbi.Voters = 0
	}

	return vbi
}

// AllBaseInfoByUID 根据 uid 获取该用户的所有投票基本信息
func AllBaseInfoByUID(uid uint64) ([]VoteBaseInfo, error) {
	var (
		vcodes []string
		err    error
		result []VoteBaseInfo
	)

	dbtool.Init()
	defer dbtool.Close()

	vcodes, err = allVcodeByUID(uid)
	if len(vcodes) == 0 || err != nil {
		return result, err
	}

	for _, vcode := range vcodes {
		result = append(result, baseInfo(vcode))
	}

	return result, err
}
