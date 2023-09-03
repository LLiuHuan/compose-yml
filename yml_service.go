package docker

// Service 服务配置
type Service struct {
	Build        Build    `yaml:"build,omitempty"`         // 构建时的配置项
	CapAdd       []string `yaml:"cap_add,omitempty"`       // 添加的容器功能
	CapDrop      []string `yaml:"cap_drop,omitempty"`      // 添加的容器功能
	CgroupParent string   `yaml:"cgroup_parent,omitempty"` // 可选的父控制组
	Command      []string `yaml:"command,omitempty"`       // 默认命令
	// TODO configs
	ContainerName string `yaml:"container_name,omitempty"` // 容器名称
	// TODO credential_spec
	DependsOn []string `yaml:"depends_on,omitempty"` // 服务之间的依赖关系
	// TODO deploy
	Deploy Deploy `yaml:"deploy,omitempty"` // 部署配置
	// TODO devices
	// TODO dns
	// TODO dns_search
	// TODO entrypoint
	EnvFile     []string      `yaml:"env_file,omitempty"`    // 环境变量文件
	Environment []Environment `yaml:"environment,omitempty"` // 环境变量
	// TODO expose
	// TODO external_links
	// TODO extra_hosts
	// TODO healthcheck
	Image Image `yaml:"image"` // 容器启动的镜像
	// TODO init
	// TODO isolation
	// TODO labels
	// TODO links
	// TODO logging
	NetworkMode string                `yaml:"network_mode,omitempty"` // 网络模式，可参见常量
	Networks    map[string]NetworkMap `yaml:"networks,omitempty"`     // 加入的网络
	// TODO pid
	Ports    []Port   `yaml:"ports,omitempty"`    // 暴露的端口号
	Proflies []string `yaml:"profiles,omitempty"` // 服务的配置文件
	Restart  string   `yaml:"restart,omitempty"`  // 重启策略
	Secrets  []Secret `yaml:"secrets,omitempty"`  // 密钥
	// TODO security_opt
	// TODO stop_grace_period
	// TODO stop_signal
	// TODO sysctls
	// TODO tmpfs
	// TODO ulimits
	// TODO userns_mode
	Volumes []VolumeMap `yaml:"volumes,omitempty"` // 挂载卷
	// Each of these is a single value, analogous to its docker run counterpart. Note that mac_address is a legacy option.
	User       string `yaml:"user,omitempty"`        // 用户名
	WorkingDir string `yaml:"working_dir,omitempty"` // 工作目录
	DomainName string `yaml:"domainname,omitempty"`  // 域名
	HostName   string `yaml:"hostname,omitempty"`    // 主机名
	Ipc        string `yaml:"ipc,omitempty"`         // ipc
	MacAddress string `yaml:"mac_address,omitempty"` // mac地址
	Privileged bool   `yaml:"privileged,omitempty"`  // 特权
	ReadOnly   bool   `yaml:"read_only,omitempty"`   // 只读
	ShmSize    string `yaml:"shm_size,omitempty"`    // 共享内存大小
	StdinOpen  bool   `yaml:"stdin_open,omitempty"`  // 标准输入
	Tty        bool   `yaml:"tty,omitempty"`         // 是否开启tty

	// 下面是适配自定义的字段，正常使用本库时可以忽略
	Host []string `yaml:"host,omitempty"` // 主机
	Exec []string `yaml:"exec,omitempty"` // 执行命令
}
