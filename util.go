package docker

import (
	yaml "gopkg.in/yaml.v2"
)

// MarshalYaml 编码为Yaml格式
func MarshalYaml(obj interface{}) (string, error) {
	v, e := yaml.Marshal(obj)
	return string(v), e
}

// UnmarshalYaml 解码Yaml文本
func UnmarshalYaml(content string, obj interface{}) error {
	return yaml.Unmarshal([]byte(content), obj)
}
