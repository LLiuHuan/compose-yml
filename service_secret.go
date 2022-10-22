package docker

// Secret 密钥(Long Syntax)
type Secret struct {
	Source string `yaml:"source,omitempty"` // 名称
	Target string `yaml:"target,omitempty"` // 文件名
	Uid    string `yaml:"uid,omitempty"`    // 文件UID
	Gid    string `yaml:"gid,omitempty"`    // 文件GID
	Mode   string `yaml:"mode,omitempty"`   // 文件权限
}
