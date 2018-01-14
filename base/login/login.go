package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"voting-system-api/tool/reqtool"
)

// UserInfo 用户基本信息
type User struct {
	ID    uint64 `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

func getAccessToken(code string) (string, error) {
	var (
		url    string
		form   map[string]string
		result string
		err    error
	)

	url = "https://github.com/login/oauth/access_token"
	form = map[string]string{
		"client_id":     "1727a687d3c6f886b356",
		"client_secret": "95f8cacd6587c6c136cb6f2627b4242988e69122",
		"code":          code,
	}

	result, err = reqtool.Post(url, form)

	if err != nil {
		fmt.Println("GitHub 第三登录信息获取失败")
		return "", err
	}

	return result, err
}

func getUserInfo(accessToken string) (string, error) {
	var (
		url    string
		result []byte
		err    error
	)
	url = "https://api.github.com/user"
	res, err := http.Get(url + "?" + accessToken)
	if err != nil {
		return "", err
	}
	result, err = ioutil.ReadAll(res.Body)
	return string(result), err
}

// Handler 处理OAuth2.0登录，参数为登录 code
func Handler(code string) (string, error) {
	var (
		accessToken string
		jsonString  string
		outputJSON  []byte
		err         error
		user        User
	)
	accessToken, err = getAccessToken(code)
	if err != nil {
		fmt.Println("access_token 获取失败")
		return "", err
	}
	jsonString, err = getUserInfo(accessToken)
	if err != nil {
		fmt.Println("用户信息获取失败")
		return jsonString, err
	}
	// fmt.Println(jsonString)
	json.Unmarshal([]byte(jsonString), &user)
	outputJSON, err = json.Marshal(user)
	// fmt.Println(jsonString)
	// fmt.Println(user, string(outputJSON))
	return string(outputJSON), err
}
