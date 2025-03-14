// Package docker
// @program: compose-yml
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-10-22 19:27
package docker

import (
	"errors"
	"fmt"
	"strings"
)

type Environment struct {
	Key  string      `yaml:"key,omitempty" json:"key,omitempty"`
	Val  interface{} `yaml:"val,omitempty" json:"val,omitempty"`
	Desc string      `yaml:"desc,omitempty" json:"desc,omitempty"`
}

// NewEnvironmentMap 新建一个环境变量映射
func NewEnvironmentMap(key string, val interface{}, desc string) Environment {
	return Environment{
		Key:  key,
		Val:  val,
		Desc: desc,
	}
}

// NewEnvironmentMapSame 新建一个相同的环境变量映射
func NewEnvironmentMapSame(volumeMap string, desc string) Environment {
	return NewEnvironmentMap(volumeMap, volumeMap, desc)
}

// MarshalYAML 序列化
func (m Environment) MarshalYAML() (result interface{}, err error) {
	if m.Key == "" {
		err = errors.New("docker: environment key can not be empty")
		return
	}
	tmp := m.Key
	if m.Val != "" {
		tmp += fmt.Sprintf("=%s", m.Val)
	} else {
		tmp += "="
	}
	result = tmp
	return
}

// UnmarshalYAML 反序列化
func (m *Environment) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	err = m.UnmarshalYAMLStruct(unmarshal)
	if err == nil {
		return
	}
	err = m.UnmarshalYAMLMap(unmarshal)
	if err == nil {
		return
	}
	err = m.UnmarshalYAMLStr(unmarshal)
	if err != nil {
		return
	}

	return nil
}

func (m *Environment) UnmarshalYAMLStr(unmarshal func(interface{}) error) (err error) {
	var origin string
	if err = unmarshal(&origin); err != nil {
		return
	}
	// 拆分协议部分
	index := strings.Index(origin, "=")

	if index != -1 { // 如果找到了子串
		m.Key = origin[:index]
		m.Val = origin[index+len("="):]
	}

	// 校验
	if m.Key == "" {
		err = errors.New("docker: simple-environment format error")
		return
	}
	if m.Val == "" {
		err = errors.New("docker: simple-environment format error")
		return
	}
	return
}

func (m *Environment) UnmarshalYAMLMap(unmarshal func(interface{}) error) (err error) {
	var origin map[string]interface{}
	if err = unmarshal(&origin); err != nil {
		return err
	}

	for key, i := range origin {
		switch key {
		case "key":
			m.Key = i.(string)
		case "val":
			m.Val = i
		case "desc":
			m.Desc = i.(string)
		}
	}

	return
}

func (m *Environment) UnmarshalYAMLStruct(unmarshal func(interface{}) error) (err error) {
	type TemporaryEnvironment struct {
		Key  string      `yaml:"key,omitempty"`
		Val  interface{} `yaml:"val,omitempty"`
		Desc string      `yaml:"desc,omitempty"`
	}
	var origin TemporaryEnvironment
	if err = unmarshal(&origin); err != nil {
		return err
	}
	m.Key = origin.Key
	m.Val = origin.Val
	m.Desc = origin.Desc
	return
}
