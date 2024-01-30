// Package docker
// @program: compose-yml
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-10-22 19:13
package docker

type Deploy struct {
	EndpointMode       string              `yaml:"endpoint_mode,omitempty"`         // endpoint_mode 端点模式 用于指定服务的网络模式，可选值为 vip 或 dnsrr
	Labels             []string            `yaml:"labels,omitempty"`                // labels 服务标签
	Mode               string              `yaml:"mode,omitempty"`                  // mode 服务模式 用于指定服务的模式，可选值为 replicated 或 global
	Placement          []map[string]string `yaml:"placement,omitempty"`             // placement 服务部署策略 用于指定服务的部署策略，可选值为 max_replicas_per_node 或 max_replicas_per_node
	MaxReplicasPerNode int                 `yaml:"max_replicas_per_node,omitempty"` // max_replicas_per_node 服务部署策略 用于指定服务的部署策略，可选值为 max_replicas_per_node 或 max_replicas_per_node
	Replicas           int                 `yaml:"replicas,omitempty"`              // replicas 服务副本数 用于指定服务的副本数，可选值为 0 到 1 000 000 000
	Resources          Resources           `yaml:"resources,omitempty"`             // resources 服务资源配置 用于指定服务的资源配置
	RestartPolicy      RestartPolicy       `yaml:"restart_policy,omitempty"`        // restart_policy 服务重启策略 用于指定服务的重启策略
	RollbackConfig     RollbackConfig      `yaml:"rollback_config,omitempty"`       // rollback_config 服务回滚配置 用于指定服务的回滚配置
	UpdateConfig       UpdateConfig        `yaml:"update_config,omitempty"`         // update_config 服务更新配置 用于指定服务的更新配置
}

type Resources struct {
	//Cpus   string `yaml:"cpus,omitempty"`   // cpus 服务 CPU 配置 用于指定服务的 CPU 配置，可选值为 0.000 到 1 000.000
	//Memory string `yaml:"memory,omitempty"` // memory 服务内存配置 用于指定服务的内存配置，可选值为 4M 到 1 000G
	Limits       ResourcesItem `yaml:"limits"`       // limits 服务资源限制 用于指定服务的资源限制，可选值为 cpus、memory
	Reservations ResourcesItem `yaml:"reservations"` // reservations 服务资源预留 用于指定服务的资源预留，可选值为 cpus、memory
}

type ResourcesItem struct {
	Cpus    string             `yaml:"cpus"`
	Memory  string             `yaml:"memory"`
	Pids    int                `yaml:"pids"`
	Devices []ResourcesDevices `yaml:"devices"`
}

type ResourcesDevices struct {
	Capabilities []string       `yaml:"capabilities"`
	Driver       string         `yaml:"driver"`
	Count        int            `yaml:"count"`
	DeviceIds    []string       `yaml:"device_ids"`
	Options      map[string]any `yaml:"options"`
}

type RestartPolicy struct {
	Condition string `yaml:"condition,omitempty"` // condition 服务重启条件 用于指定服务的重启条件，可选值为 none、on-failure、any
}

type RollbackConfig struct {
	Parallelism     int     `yaml:"parallelism,omitempty"`       // parallelism 服务回滚并行数 用于指定服务的回滚并行数，可选值为 0 到 1 000 000 000
	Delay           int     `yaml:"delay,omitempty"`             // delay 服务回滚延迟 用于指定服务的回滚延迟，可选值为 0 到 1 000 000 000
	FailureAction   string  `yaml:"failure_action,omitempty"`    // failure_action 服务回滚失败行为 用于指定服务的回滚失败行为，可选值为 pause、continue
	MaxFailureRatio float64 `yaml:"max_failure_ratio,omitempty"` // max_failure_ratio 服务回滚失败比例 用于指定服务的回滚失败比例，可选值为 0.000 到 1.000
	Monitor         int     `yaml:"monitor,omitempty"`           // monitor 服务回滚监控 用于指定服务的回滚监控，可选值为 0 到 1 000 000 000
	Order           string  `yaml:"order,omitempty"`             // order 服务回滚顺序 用于指定服务的回滚顺序，可选值为 start-first、stop-first
}

type UpdateConfig struct {
	Parallelism     int     `yaml:"parallelism,omitempty"`       // parallelism 服务更新并行度 用于指定服务的更新并行度，可选值为 0 到 1 000
	Delay           int     `yaml:"delay,omitempty"`             // delay 服务更新延迟 用于指定服务的更新延迟，可选值为 0 到 1 000 000 000
	FailureAction   string  `yaml:"failure_action,omitempty"`    // failure_action 服务更新失败行为 用于指定服务的更新失败行为，可选值为 continue、pause、rollback
	MaxFailureRatio float64 `yaml:"max_failure_ratio,omitempty"` // max_failure_ratio 服务更新失败比例 用于指定服务的更新失败比例，可选值为 0.000 到 1.000
	Monitor         int     `yaml:"monitor,omitempty"`           // monitor 服务更新监控 用于指定服务的更新监控，可选值为 0 到 1 000 000 000
	Order           string  `yaml:"order,omitempty"`             // order 服务更新顺序 用于指定服务的更新顺序，可选值为 start-first、stop-first
}
