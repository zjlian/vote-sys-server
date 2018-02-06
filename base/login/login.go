package login

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"voting-system-api/tool/dbtool"
	"voting-system-api/tool/reqtool"
)

// User 用户基本信息
type User struct {
	ID    uint64 `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

// getAccessToken 根据get请求第三方登录返回的code验证码，请求获取用户信息的验证码
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

	if strings.Index(result, "error=bad_verification_code") != -1 {
		return "", errors.New("登录信息已过期")
	}

	if err != nil {
		// log.Println("GitHub 第三登录信息获取失败")
		return "", err
	}

	return result, err
}

// getUserInfo 同过 getAccessToken() 函数获取到的验证吗，获取用户信息
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

// signUp 注册一个不存在的用户（需要先调用依赖的数据库包dbtool.Init()）
func signUp(u User) {
	result, err := dbtool.DB.Insert(
		"voteSys.Ass (uid,username,login) values (?,?,?)",
		u.ID, u.Name, u.Login,
	)

	if err != nil {
		log.Println(result, err)
	}
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
		return "", err
	}
	//log.Println("获取到 access_token: " + accessToken)
	jsonString, err = getUserInfo(accessToken)
	if err != nil {
		// log.Println("用户信息获取失败")
		return jsonString, err
	}

	json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		return "false", err
	}

	outputJSON, err = json.Marshal(user)
	//log.Println("获取到用户信息: " + string(outputJSON))
	dbtool.Init()
	defer dbtool.Close()
	//log.Println("进行数据库查询，是否已有该用户信息")
	// 查询数据库中是否存在该用户的信息，没有就添加一条

	if !HasUser(user.ID) {
		//log.Println("不存在的用户，登记注册")
		signUp(user)
	}
	//log.Println("完成登录")
	return string(outputJSON), err
}
