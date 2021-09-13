package zhuhe

import "open-git.jonham.cn/Jonham/goeasy/request/raw"

//API 对API的实现
type API struct {
	Token    string
	BasePath string
	Logging  bool //是否logging
}

//NewAPI 获取带有默认值的API信息
func NewAPI(token string) *API {
	return &API{
		BasePath: BaseURL, //末尾不要带/，只填正式API。测试环境通过环境变量处理
		Token:    token,
	}
}

func (api API) Get(url string) ([]byte, error) {
	return raw.Get(api.ParseURL(url), raw.ParseTokenHeader(api.Token))
}

func (api API) Post(url string, data interface{}) ([]byte, error) {
	return raw.Post(api.ParseURL(url), data, raw.ParseTokenHeader(api.Token))
}

func (api API) HandleError() {
}

func (api API) ParseURL(reqPath string) string {
	return api.BasePath + reqPath
}
