package docker

// Network 网络配置
type Network struct {
	Driver     string         `yaml:"driver,omitempty"`
	DriverOpts map[string]any `yaml:"driver_opts,omitempty"`
	Attachable bool           `yaml:"attachable,omitempty"`
	EnableIpv6 bool           `yaml:"enable_ipv6,omitempty"`
	Ipam       NetworkIpam    `yaml:"ipam,omitempty"`
	// TODO internal
	Labels   map[string]string `yaml:"labels,omitempty"`
	External bool              `yaml:"external,omitempty"` // 是否是外部创建
	Name     string            `yaml:"name,omitempty"`
}

type NetworkIpam struct {
	Driver  string `yaml:"driver,omitempty"`
	Config  []NetworkIpamConfig
	Options map[string]any
}

type NetworkIpamConfig struct {
	Subnet             string `yaml:"subnet,omitempty"`
	Gateway            string `yaml:"gateway,omitempty"`
	IPRange            string `yaml:"ip_range,omitempty"`
	AuxiliaryAddresses map[string]string
}
