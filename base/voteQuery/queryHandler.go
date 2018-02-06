package voteQuery

// QueryPayloadData 查询请求的结构体，
//Action 字段存放示要查询的内容，Data 存放查询所需的数据
type QueryPayloadData struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

// Do 传入一个查询请求附带的json字符串，返回查询结果
func Do(jsonstr string) string {
	return "TODO"
}
