package zhuhe

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"log"
	"open-git.jonham.cn/Jonham/goeasy/checkerror"
	"open-git.jonham.cn/Jonham/goeasy/request/datagetter"
	"testing"
)

type param struct {
	Name string
	ID   int
}

func TestAPI_FillURL(t *testing.T) {
	p := param{"Jonham", 12}
	temp := "https://demo.jonham.cn/${Name}/${ID}/${Age}"
	expectResult := "https://demo.jonham.cn/Jonham/12/${Age}"
	result := API{}.FillURLStruct(temp, p)
	assert.Equal(t, expectResult, result)
}

func TestAPI_FillURLFreeJSON(t *testing.T) {
	val := `{"id":12,"name":"Jonham","ageFormat":121}`

	var v interface{}
	err := jsoniter.UnmarshalFromString(val, &v)
	checkerror.CheckLog(err)

	temp := "https://demo.jonham.cn/${name}/${id}/${ageFormat}"
	expectResult := "https://demo.jonham.cn/Jonham/12/121"
	result := API{}.FillURL(temp, v)
	assert.Equal(t, expectResult, result)
}

func TestAPI_GetAPI(t *testing.T) {
	api := API{
		Token:    "eyJhbGciOiJIUzI1NiJ9.eyJuYmYiOjE2Mjk0NTQ3MTIsInRhZ19uYW1lIjoiYWRtaW4iLCJhZG1pbklkIjo0LCJleHAiOjE2MzIwNDY3MTIsIndoZXRoZXJfanVkZ2Vfd2hpdGVfbGlzdCI6ZmFsc2UsImlhdCI6MTYyOTQ1NDcxMn0.kctBorIqxpV02eB9dN5fkr8WPAXe1KGV8KHUFB2EVI4",
		BasePath: "",
		Logging:  true,
	}
	getter, err := api.GetAPI("https://temping.dev.shxg.tech/api/workbenchSpecialty/project/717389537034108935")
	checkerror.CheckLog(err)

	log.Println(getter.Data)
	g := datagetter.InitProxyGetter(getter.Data)
	log.Println(g.String())
	log.Println(g.GetChild("city"))
}
