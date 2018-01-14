package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"

	"voting-system-api/base/login"
	"voting-system-api/tool/reqtool"
)

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

	err = json.Unmarshal([]byte(jsonStr), &lc)

	if err != nil && lc.Code == "" {
		fmt.Println("数据错误")
		return
	}
	jsonStr, err = login.Handler(lc.Code)
	if err != nil {
		fmt.Fprintf(res, "%s", "false")
	}
	fmt.Fprintf(res, "%s", jsonStr)
}

func loginByGitHub(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%s", `
		<script>
			window.location.href='https://github.com/login/oauth/authorize?client_id=1727a687d3c6f886b356'
		</script>	
	`)
}

func main() {
	flag.Parse()
	var portStr = ":" + (*portCode)

	http.HandleFunc("/", defHandler)
	http.HandleFunc("/loginbygithub", loginByGitHub)
	http.HandleFunc("/login", loginHandler)

	http.ListenAndServe(portStr, nil)
}
