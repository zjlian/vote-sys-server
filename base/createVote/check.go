package createVote

import (
	"strings"
)

func isLegalOption(O *Option) bool {
	return strings.TrimSpace(O.Content) != ""
}

func isLegalVote(V *Vote) bool {
	opCount := 0

	if V.Deadline < V.CreateTime {
		return false
	}
	if strings.TrimSpace(V.Title) == "" {
		return false
	}
	if V.SelectType != 0 && V.SelectType != 1 {
		return false
	}

	for _, v := range V.Options {
		if isLegalOption(&v) {
			opCount++
		}
	}
	if opCount <= 0 {
		return false
	}

	return true
}
