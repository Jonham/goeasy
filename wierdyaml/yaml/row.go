package yaml

import (
	"log"
	"regexp"
	"strings"
)

//IsComment 是否是注释
func IsComment(c string) bool {
	if strings.HasPrefix(c, "#") {
		return true
	}
	return false
}

//SplitRow 拆解行，返回具体信息
func SplitRow(r string) (parsed *ParsedYAML) {
	leftSpace := countLeftSpace(r)
	rg := regexp.MustCompile("^( +)?(.+?):(.+)")
	strList := rg.FindAllStringSubmatch(r, -1)
	if len(strList) == 1 {
		l := strList[0]
		return &ParsedYAML{
			Indent: leftSpace,
			Raw:    l[0],
			Name:   l[2],
			Value:  strings.Trim(l[3], " "),
			Type:   "string",
		}
	}

	return nil
}

//ParseToBook 注册到书
func ParseToBook(raw string) YAMLBook {
	result := YAMLBook{}
	//parentMap := map[int]string{}
	//currentParentLevel := 0
	currentParent := ""
	currentLevel := 0
	preNode := ""
	joinParent := func(parsed *ParsedYAML) string {
		trim := strings.Trim(parsed.Name, " ")
		if currentParent == "" {
			return trim
		} else {
			return currentParent + "." + trim
		}
	}
	list := strings.Split(raw, "\n")
	for _, str := range list {
		item := SplitRow(str)
		if item != nil {
			log.Println(item)
			if item.Indent > currentLevel {
				currentParent = preNode
				//currentParentLevel = currentLevel
				preNode = ""
				currentLevel = item.Indent
			} else {
			}
			if item.Value != "" {
				result[joinParent(item)] = item.Value
			}
		} else {
			//空行
			//currentParentLevel = 0
			currentParent = ""
		}
	}

	return result
}

func countLeftSpace(raw string) int {
	if len(raw) == 0 {
		return 0
	}
	for i := 0; i < len(raw); i++ {
		if raw[i] != ' ' {
			return i
		}
	}
	return len(raw)
}
