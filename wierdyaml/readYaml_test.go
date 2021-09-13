package wierdyaml

import "testing"

func TestReadYAML(t *testing.T) {
	ReadYAML("../temp/yamlFile/.env.yaml")
}
