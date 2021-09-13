package zhuhe

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func (api API) logIfPrintf(format string, v ...interface{}) {
	if api.Logging {
		log.Printf(format, v...)
	}
}

//outputTempJSON 缓存JSON到本地
func outputTempJSON(res []byte, debugFilename ...string) {
	if len(debugFilename) > 0 {
		filename := fmt.Sprintf("%s.json", debugFilename)
		os.WriteFile(filename, res, 0777)
	}
}

//GetAPI 发起请求
func (api API) GetAPI(url string, debugFilename ...string) (*ResponseCommonDto, error) {
	api.logIfPrintf("GetAPI(): %s\n", url)

	res, err := api.Get(url)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	outputTempJSON(res, debugFilename...)

	data := &ResponseCommonDto{}
	err = json.Unmarshal(res, data)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return data, err
}

func (api API) PostAPI(url string, body interface{}, debugFilename ...string) (*ResponseCommonDto, error) {
	api.logIfPrintf("PostAPI(): %s\n", url)

	res, err := api.Post(url, body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	outputTempJSON(res, debugFilename...)

	data := &ResponseCommonDto{}
	err = json.Unmarshal(res, data)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return data, err
}

type ResponseCommonDto struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
