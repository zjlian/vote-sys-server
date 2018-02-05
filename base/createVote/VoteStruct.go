package createVote

// Option 投票选项结构
type Option struct {
	ID      int    `json:"id"`
	Vcode   string `json:"vcode"`
	Content string `json:"content"`
}

// Vote 投票结构
type Vote struct {
	ID         int      `json:"id"`
	Vcode      string   `json:"vcode"`
	Title      string   `json:"title"`
	Describe   string   `json:"descrobe"`
	SelectType uint8    `json:"selectType"`
	CreateTime uint64   `json:"createTime"`
	Deadline   uint64   `json:"deadline"`
	Location   string   `json:"location"`
	Options    []Option `json:"options"`
}
