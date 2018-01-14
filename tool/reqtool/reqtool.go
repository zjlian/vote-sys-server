package reqtool

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func enCodingRequestArgs(form interface{}) string {
	if form == nil {
		return ""
	}
	switch form.(type) {
	case string:
		return form.(string)

	case map[string]string:
		var req http.Request
		req.ParseForm()
		for k, v := range form.(map[string]string) {
			req.Form.Add(k, v)
		}
		return strings.TrimSpace(req.Form.Encode())
	}
	return ""
}

func sendRequest(req *http.Request) (string, error) {
	var (
		resp   *http.Response
		result []byte
		err    error
	)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("on package sendRequest: 请求失败")
		return "", err
	}

	result, err = ioutil.ReadAll(resp.Body)

	return string(result), err
}

// Post 发起一个post请求 第一个参数是一个 url 地址，第二个参数则是一个要发送数据的map
// 返回值为 string, error
func Post(url string, form interface{}) (string, error) {
	var (
		data string
		err  error
	)

	data = enCodingRequestArgs(form)

	request, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		fmt.Println("on package req: 创建post请求失败")
		return "", err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return sendRequest(request)
}

// Get 发起一个get请求 第一个参数是一个 url 地址，第二个参数则是一个要发送数据的map
// 返回值为 string, error
func Get(url string, form interface{}) (string, error) {
	var (
		data string
		err  error
	)

	data = enCodingRequestArgs(form)

	request, err := http.NewRequest("GET", url+"?"+data, nil)
	if err != nil {
		fmt.Println("on package req: 创建get请求失败")
		return "", err
	}

	return sendRequest(request)
}

// GetDataString 获取http请求中的 <entity-body> 部分
func GetDataString(req *http.Request) string {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(result).String()
}

// SetACAO 给response设置跨域属性
func SetACAO(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	(*res).Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	(*res).Header().Set("content-type", "application/json")             //返回数据格式是json
}
