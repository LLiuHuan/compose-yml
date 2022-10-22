package docker

import (
	"errors"
	"fmt"
	"strings"
)

// Image 镜像
type Image struct {
	Name string `yaml:"name,omitempty"` // 镜像名称
	Tag  string `yaml:"tag,omitempty"`  // 镜像标签
}

// NewImage 新建一个镜像
func NewImage(name string, tag string) Image {
	return Image{
		Name: name,
		Tag:  tag,
	}
}

// NewImageFromText 从文本创建镜像
func NewImageFromText(str string) (m Image) {
	parts := strings.Split(str, ":")
	m.Name = parts[0]
	if len(parts) > 1 {
		m.Tag = parts[1]
	} else {
		m.Tag = "latest"
	}
	return
}

// MarshalYAML 序列化
func (m Image) MarshalYAML() (result interface{}, err error) {
	if len(m.Name) == 0 {
		err = errors.New("docker: image name can not be empty")
		return
	}
	if len(m.Tag) > 0 {
		result = fmt.Sprintf("%s:%s", m.Name, m.Tag)
	} else {
		result = m.Name
	}
	return
}

// UnmarshalYAML 反序列化
func (m *Image) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var origin string
	if err = unmarshal(&origin); err != nil {
		return
	}
	parts := strings.Split(origin, ":")
	if len(parts) > 2 {
		err = errors.New("docker: image format error")
		return
	}
	m.Name = parts[0]
	if len(parts) > 1 {
		m.Tag = parts[1]
	}
	return
}
