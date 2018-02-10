package voteQuery

import (
	"encoding/json"
	"errors"
)

// QueryPayloadData 查询请求的结构体，
//Action 字段存放示要查询的内容，Data 存放查询所需的数据
type QueryPayloadData struct {
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

// Do 传入一个查询请求附带的json字符串，返回查询结果
func Do(jsonstr string) (string, error) {
	var (
		req    QueryPayloadData
		err    error
		result []byte
	)

	err = json.Unmarshal([]byte(jsonstr), &req)
	if err != nil {
		return "", err
	}

	switch req.Action {
	case "QUERY_USER_ALL_VOTE_BASE_INFO":
		if req.Data["uid"] == nil {
			err = errors.New("请求所需参数不存在")
			break
		}

		uid, ok := req.Data["uid"].(float64)
		if !ok {
			err = errors.New("请求所需参数无效")
			break
		}

		result, err = AllVoteBaseInfoToJSON(uint64(uid))
	default:
		err = errors.New("非法请求")
	}
	return string(result), err
}
