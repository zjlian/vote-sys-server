package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"voting-system-api/base/createVote"
	"voting-system-api/base/login"
	"voting-system-api/base/voteQuery"
	"voting-system-api/tool/reqtool"
)

// LoginCode 第三方登录用
type LoginCode struct {
	Code string
}

var portCode = flag.String("p", "80", "web服务监听的端口")

func defHandler(res http.ResponseWriter, req *http.Request) {
}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		return
	}
	reqtool.SetACAO(&res)

	var (
		lc      LoginCode
		jsonStr string
		err     error
	)
	jsonStr = reqtool.GetDataString(req)
	if strings.Index(jsonStr, "code") == -1 {
		return
	}
	//log.Println("收到 code: " + jsonStr)
	err = json.Unmarshal([]byte(jsonStr), &lc)

	if err != nil && lc.Code == "" {
		log.Println("数据错误")
		return
	}
	//log.Println("调用登录处理函数")
	jsonStr, err = login.Handler(lc.Code)
	if err != nil {
		fmt.Fprintf(res, "{\"status\":%t,\"info\":\"%s\"}", false, err.Error())
	}

	fmt.Fprintf(res, "%s", jsonStr)
}

func loginByGitHub(res http.ResponseWriter, req *http.Request) {
	//log.Println("请求通过github登录")
	fmt.Fprintf(res, "%s", `
		<script>
			window.location.href='https://github.com/login/oauth/authorize?client_id=1727a687d3c6f886b356'
		</script>
	`)
}

func createVoteHandler(res http.ResponseWriter, req *http.Request) {
	var (
		jsonStr string
		err     error
	)
	reqtool.SetACAO(&res)
	jsonStr = reqtool.GetDataString(req)

	err = createVote.Create(jsonStr)
	if err != nil {
		fmt.Fprintf(res, "%s", err.Error())
	} else {
		fmt.Fprintf(res, "%d", 1)
	}
}

func queryHandler(res http.ResponseWriter, req *http.Request) {
	var (
		jsonStr string
		err     error
		result  string
	)
	reqtool.SetACAO(&res)
	jsonStr = reqtool.GetDataString(req)
	result, err = voteQuery.Do(jsonStr)

	if err != nil {
		log.Println(err)
		fmt.Fprintf(res, "%s", err.Error())
	} else {
		fmt.Fprintf(res, "%s", result)
	}
}

func main() {
	flag.Parse()
	var portStr = ":" + (*portCode)

	voteQuery.Init()

	http.HandleFunc("/", defHandler)
	http.HandleFunc("/loginbygithub", loginByGitHub)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/create", createVoteHandler)
	http.HandleFunc("/query", queryHandler)

	http.ListenAndServe(portStr, nil)

}
