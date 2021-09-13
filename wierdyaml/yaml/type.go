package yaml

type YAMLValue interface{}

type YAMLBook map[string]YAMLValue

type ParsedYAML struct {
	Indent int
	Raw    string
	Name   string //变量名称
	Type   string
	Value  string
}
