package docker

import (
	"errors"
	"fmt"
	"github.com/docker/cli/cli/compose/types"
	"strconv"
	"strings"
)

// Port 端口(Short Syntax)
type Port struct {
	Mode      string `yaml:"mode,omitempty" json:"mode,omitempty"`
	Target    uint32 `yaml:"target,omitempty" json:"target,omitempty"`
	Published uint32 `yaml:"published,omitempty" json:"published,omitempty"`
	Protocol  string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
}

// NewPort 新建一个端口A到端口B的映射
func NewPort(hostPort uint32, containerPort uint32) Port {
	return Port{
		Published: hostPort,
		Target:    containerPort,
	}
}

// NewPortSame 新建一个相同端口映射
func NewPortSame(port uint32) Port {
	return NewPort(port, port)
}

// MarshalYAML 序列化
func (m Port) MarshalYAML() (result interface{}, err error) {
	if m.Published == 0 {
		err = errors.New("docker: simple-port host can not be empty")
		return
	}
	tmp := strconv.FormatUint(uint64(m.Published), 10)
	if m.Target != 0 {
		tmp += fmt.Sprintf(":%d", m.Target)
		if len(m.Protocol) > 0 {
			tmp += fmt.Sprintf("/%s", m.Protocol)
		}
	}
	result = tmp
	return
}

// UnmarshalYAML 反序列化
func (m *Port) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	err = m.UnmarshalYAMLStr(unmarshal)
	if err == nil {
		return
	}
	err = m.UnmarshalYAMLMap(unmarshal)
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
		parseUint, err := strconv.ParseUint(remain, 10, 32)
		if err != nil {
			return err
		}
		m.Published, m.Target = uint32(parseUint), 0
	} else if loc != len(remain)-1 {
		parseUint, err := strconv.ParseUint(remain[0:loc], 10, 32)
		if err != nil {
			return err
		}
		m.Published = uint32(parseUint)
		parseUint, err = strconv.ParseUint(remain[loc+1:], 10, 32)
		if err != nil {
			return err
		}
		m.Target = uint32(parseUint)
	} else {
		parseUint, err := strconv.ParseUint(remain[0:loc], 10, 32)
		if err != nil {
			return err
		}
		m.Published = uint32(parseUint)
	}
	// 校验
	if m.Published == 0 {
		err = errors.New("docker: simple-port format error")
		return
	}
	if m.Target == 0 && len(m.Protocol) > 0 {
		err = errors.New("docker: simple-port format error")
		return
	}
	return
}

func (m *Port) UnmarshalYAMLMap(unmarshal func(interface{}) error) (err error) {
	var origin types.ServicePortConfig
	if err = unmarshal(&origin); err != nil {
		return err
	}
	m.Target = origin.Target
	m.Published = origin.Published
	m.Protocol = origin.Protocol
	m.Mode = origin.Mode
	return
}
