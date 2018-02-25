package voteQuery

import (
	"encoding/json"
	"errors"
	"event"
)

// QueryPayloadData 查询请求的结构体，
//Action 字段存放示要查询的内容，Data 存放查询所需的数据
type QueryPayloadData struct {
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

// Init 初始化voteQuery，在事件系统上绑定查询命令和处理函数
func Init() {
	event.Subscribe(
		event.CreateEvent(
			"QUERY_USER_ALL_VOTE_BASE_INFO", AllVoteBaseInfoToJSON))

	event.Subscribe(
		event.CreateEvent(
			"QUERY_VOTE_INFO", VoteInfoToJSON))
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

	results, perfect := event.Publish(req.Action, req.Data)
	if !perfect {
		return "", errors.New("handler.go Do\n" + event.Error())
	}
	for _, v := range results {
		var ok bool
		if result, ok = v.([]byte); ok {
			break
		}
	}
	return string(result), err
}
