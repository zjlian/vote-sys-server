package createVote

import (
	"errors"
	"strings"
	"voting-system-api/base/login"
)

func isLegalOption(O *Option) bool {
	return strings.TrimSpace(O.Content) != ""
}

func isLegalVote(V *Vote) (bool, error) {
	opCount := 0

	if !login.HasUser(V.UID) {
		return false, errors.New("创建投票的用户不存在系统中")
	}
	if V.Deadline < V.CreateTime {
		return false, errors.New("截至时间小于投票创建时间")
	}
	if strings.TrimSpace(V.Title) == "" {
		return false, errors.New("投票标题为空")
	}
	if V.SelectType != 0 && V.SelectType != 1 {
		return false, errors.New("投票类型异常")
	}

	for _, v := range V.Options {
		if isLegalOption(&v) {
			opCount++
		}
	}
	if opCount < 2 && opCount > 64 {
		return false, errors.New("投票选项数量异常")
	}

	return true, nil
}
