package docker

import (
	"errors"
	"fmt"
	"strings"
)

//// VolumeMap 路径(Short Syntax)
//type VolumeMap struct {
//	Host      string // 外部主机的路径
//	Container string // 内部容器的路径
//	Mode      string // 权限
//}

// VolumeMap 路径(Short Syntax)
type VolumeMap struct {
	Type     string `yaml:"type,omitempty"`      // 外部主机的路径
	Source   string `yaml:"source,omitempty"`    // 内部容器的路径
	Target   string `yaml:"target,omitempty"`    // 权限
	ReadOnly bool   `yaml:"read_only,omitempty"` // 是否只读
	Desc     string `yaml:"desc,omitempty"`      // 描述
}

// NewVolumeMap 新建一个路径A到路径B的映射
func NewVolumeMap(hostVolumeMap string, containerVolumeMap string, desc string) VolumeMap {
	return VolumeMap{
		Source: hostVolumeMap,
		Target: containerVolumeMap,
	}
}

// NewVolumeMapSame 新建一个相同的路径映射
func NewVolumeMapSame(volumeMap string, desc string) VolumeMap {
	vm := NewVolumeMap(volumeMap, volumeMap, desc)

	//vm.Mode = VolumeReadOnly
	return vm
}

// MarshalYAML 序列化
func (m VolumeMap) MarshalYAML() (result interface{}, err error) {
	if len(m.Target) == 0 {
		err = errors.New("docker: simple-volume-map host can not be empty")
		return
	}
	tmp := m.Target
	if len(m.Source) > 0 {
		tmp += fmt.Sprintf(":%s", m.Source)
		if m.ReadOnly {
			tmp += fmt.Sprintf(":%s", VolumeReadOnly)
		}
	}
	result = tmp
	return
}

// UnmarshalYAML 反序列化
func (m *VolumeMap) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
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

func (m *VolumeMap) UnmarshalYAMLStr(unmarshal func(interface{}) error) (err error) {
	var origin string
	if err = unmarshal(&origin); err != nil {
		return err
	}
	// 拆分
	parts := strings.Split(origin, ":")
	if len(parts) > 3 {
		err = errors.New("docker: simple-volume-map format error")
		return
	}
	m.Source = parts[0]
	if len(parts) > 1 {
		m.Target = parts[1]
	}
	if len(parts) > 2 {
		if parts[2] == VolumeReadOnly {
			m.ReadOnly = true
		}
	}
	// 校验
	if len(m.Source) == 0 {
		err = errors.New("docker: simple-volume-map format error")
		return
	}
	if len(m.Target) == 0 {
		err = errors.New("docker: simple-volume-map format error")
		return
	}
	return
}

func (m *VolumeMap) UnmarshalYAMLMap(unmarshal func(interface{}) error) (err error) {
	type TemporaryVolumeMap struct {
		Type     string `yaml:"type,omitempty"`      // 外部主机的路径
		Source   string `yaml:"source,omitempty"`    // 内部容器的路径
		Target   string `yaml:"target,omitempty"`    // 权限
		ReadOnly bool   `yaml:"read_only,omitempty"` // 是否只读
		Desc     string `yaml:"desc,omitempty"`      // 描述
	}

	var origin TemporaryVolumeMap
	if err = unmarshal(&origin); err != nil {
		return err
	}
	m.Type = origin.Type
	m.Source = origin.Source
	m.Target = origin.Target
	m.ReadOnly = origin.ReadOnly
	m.Desc = origin.Desc
	return
}
