package raw

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const HTTPContentTypeJSON = "application/json"
const HTTPHeaderContentType = "Content-Type"

type HeaderInfo map[string]string

//SetHeader 设置请求头
func SetHeader(resp *http.Request, headerMap HeaderInfo) {
	for K, v := range headerMap {
		resp.Header[K] = []string{v}
	}
}

//SetHeaderContentType 设置JSON类型请求头
func SetHeaderContentType(resp *http.Request) {
	resp.Header.Set(HTTPHeaderContentType, HTTPContentTypeJSON)
}

func prepareJSON(data interface{}) *bytes.Buffer {
	jsonStr, _ := json.Marshal(data)
	body := bytes.NewBuffer(jsonStr)
	return body
}

//ParseTokenHeader Token填在header
func ParseTokenHeader(token string) HeaderInfo {
	return HeaderInfo{"token": token}
}

////Get
//// 发送GET请求
//// url：         请求地址
//// headerMap：	 header请求头
//// response：    请求返回的内容
//func Get(url string, headerMap HeaderInfo) string {
//	// 超时时间：5秒
//	client := &http.Client{Timeout: 5 * time.Second}
//	request, _ := http.NewRequest("GET", url, nil)
//	SetHeader(request, headerMap)
//	SetHeaderContentType(request)
//	//发送请求
//	Response, err := client.Do(request)
//	if err != nil {
//		panic(err)
//	}
//	defer Response.Body.Close()
//	respByte, _ := ioutil.ReadAll(Response.Body)
//	return string(respByte)
//}

//Get
// 发送GET请求
// url：         请求地址
// headerMap：	 header请求头
// response：    请求返回的内容
func Get(url string, headerMap HeaderInfo) ([]byte, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	request, _ := http.NewRequest("GET", url, nil)
	SetHeader(request, headerMap)
	SetHeaderContentType(request)
	//发送请求
	Response, err := client.Do(request)
	if err != nil {
		//panic(err)
		return nil, err
	}
	defer Response.Body.Close()
	respByte, _ := ioutil.ReadAll(Response.Body)
	return respByte, nil
}

//Post
// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// headerMap：	 header请求头
// content：     请求放回的内容
func Post(url string, data interface{}, headerMap HeaderInfo) ([]byte, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	body := prepareJSON(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	SetHeader(req, headerMap)
	SetHeaderContentType(req)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	result, _ := ioutil.ReadAll(resp.Body)
	return result, nil
}
