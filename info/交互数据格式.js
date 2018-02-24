/**
// Option 投票选项结构
type Option struct {
	ID      int    `json:"id"`
	Vcode   string `json:"vcode"`
	Content string `json:"content"`
}

// Vote 投票结构
type Vote struct {
	//ID         int      `json:"id"`
	UID        uint64   `json:"uid"`
	Vcode      string   `json:"vcode"`
	Title      string   `json:"title"`
	Describe   string   `json:"descrobe"`
	SelectType uint8    `json:"selectType"`
	CreateTime uint64   `json:"createTime"`
	Deadline   uint64   `json:"deadline"`
	Location   string   `json:"location"`
	Options    []Option `json:"options"`
}
*/

// 投票创建
// GET/POST 请求url localhost/create
const createVoteData = {
    //用户登录时获取到的 id
    uid:        number,
    title:      string,
    describe:   string,
    //投票选择类型，只有0和1，0是单选，1是多选
    selectType: number,
    createTime: Date,
    deadline:   Date,
    location:   string,
    //这个只需要数组里的对象各拥有一个 content 属性即可
    options:    array
}



// 查询单个用户创建的所有投票的基本信息，用于展示
const queryUserAllVoteBaseInfo = {
    action: 'QUERY_USER_ALL_VOTE_BASE_INFO',
    data: {
        //用户登录时获取到的 id
        uid: number
    }
}

// 查询某个投票的详细信息
const queryVoteInfo = {
    action: 'QUERY_VOTE_INFO',
    data: {
        vcode: string
    }
}

// 查询某个用户在一个投票中所投的选项
const queryUserVotingRecord = {
    action: 'QUERY_USER_VOTING_RECORD',
    data: {
        uid: number,
        vcode: string
    }
}