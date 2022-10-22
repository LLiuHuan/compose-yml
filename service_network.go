package docker

// NetworkMap 网络
type NetworkMap struct {
	Aliases     []string `yaml:"aliases,omitempty"`      // 别名列表
	Ipv4Address string   `yaml:"ipv4_address,omitempty"` // ipv4地址
	Ipv6Address string   `yaml:"ipv6_address,omitempty"` // ipv6地址
	// TODO ipam
}
