package docker

// Yml 完整的配置文件
type Yml struct {
	Version  string                `yaml:"version"`            // 版本号
	Services map[string]*Service   `yaml:"services"`           // 服务配置
	Volumes  map[string]*Volume    `yaml:"volumes,omitempty"`  // 挂载卷配置
	Networks map[string]*Network   `yaml:"networks,omitempty"` // 网络配置
	Secrets  map[string]*YmlSecret `yaml:"secrets,omitempty"`  // 密钥
	// TODO configs
	// TODO Variable substitution
	// TODO Extension fields
}

// NewYml 从文本生成Yml结构体
func NewYml(text string) (obj Yml, err error) {
	err = UnmarshalYaml(text, &obj)
	return
}
