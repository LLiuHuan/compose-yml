package docker

import (
	"errors"
	"fmt"
	"strings"
)

// Port 端口(Short Syntax)
type Port struct {
	Mode      string `yaml:"mode,omitempty" json:"mode,omitempty"`
	Target    string `yaml:"target,omitempty" json:"target,omitempty"`
	Published string `yaml:"published,omitempty" json:"published,omitempty"`
	Protocol  string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
	Desc      string `yaml:"desc,omitempty" json:"desc,omitempty"` // 描述
}

// NewPort 新建一个端口A到端口B的映射
func NewPort(hostPort string, containerPort string, desc string) Port {
	return Port{
		Published: hostPort,
		Target:    containerPort,
		Desc:      desc,
	}
}

// NewPortSame 新建一个相同端口映射
func NewPortSame(port string, desc string) Port {
	return NewPort(port, port, desc)
}

// MarshalYAML 序列化
func (m Port) MarshalYAML() (result interface{}, err error) {
	if m.Published == "" {
		err = errors.New("docker: simple-port host can not be empty")
		return
	}
	tmp := m.Published
	if m.Target != "" {
		tmp += fmt.Sprintf(":%s", m.Target)
		if len(m.Protocol) > 0 {
			tmp += fmt.Sprintf("/%s", m.Protocol)
		}
	}
	result = tmp
	return
}

// UnmarshalYAML 反序列化
func (m *Port) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
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

func (m *Port) UnmarshalYAMLStr(unmarshal func(interface{}) error) (err error) {
	var origin string
	if err = unmarshal(&origin); err != nil {
		return
	}
	// 拆分协议部分
	parts, remain := strings.Split(origin, "/"), ""
	if len(parts) > 2 {
		err = errors.New("docker: simple-port format error")
		return
	}
	if len(parts) > 1 {
		m.Protocol = parts[1]
	}
	remain = parts[0]
	// 拆分主机和容器端口
	loc := strings.LastIndex(remain, ":")
	if loc < 0 {
		m.Published, m.Target = remain, ""
	} else if loc != len(remain)-1 {
		m.Published = remain[0:loc]
		m.Target = remain[loc+1:]
	} else {
		m.Published = remain[0:loc]
	}
	// 校验
	if m.Published == "" {
		err = errors.New("docker: simple-port format error")
		return
	}
	if m.Target == "" && len(m.Protocol) > 0 {
		err = errors.New("docker: simple-port format error")
		return
	}
	return
}

func (m *Port) UnmarshalYAMLMap(unmarshal func(interface{}) error) (err error) {
	type TemporaryPort struct {
		Mode      string `yaml:"mode,omitempty" json:"mode,omitempty"`
		Target    string `yaml:"target,omitempty" json:"target,omitempty"`
		Published string `yaml:"published,omitempty" json:"published,omitempty"`
		Protocol  string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
		Desc      string `yaml:"desc,omitempty" json:"desc,omitempty"` // 描述
	}
	var origin TemporaryPort
	if err = unmarshal(&origin); err != nil {
		return err
	}
	m.Target = origin.Target
	m.Published = origin.Published
	m.Protocol = origin.Protocol
	m.Mode = origin.Mode
	m.Desc = origin.Desc
	return
}
