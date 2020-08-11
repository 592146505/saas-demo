package conf

import (
	"time"
)

// NacosConfig nacos配置
type NacosConfig struct {
	Client NacosClientConfig            `yaml:"client"`
	Server map[string]NacosServerConfig `yaml:"server"`
}

// NacosClientConfig nacos客户端配置
type NacosClientConfig struct {
	TimeoutMs            uint64 `yaml:"timeout_ms"`              // TimeoutMs http请求超时时间，单位毫秒
	ListenInterval       uint64 `yaml:"listen_interval"`         // ListenInterval 监听间隔时间，单位毫秒（仅在ConfigClient中有效）
	BeatInterval         int64  `yaml:"beat_interval"`           // BeatInterval         心跳间隔时间，单位毫秒（仅在ServiceClient中有效）
	NamespaceId          string `yaml:"namespace_id"`            // NamespaceId          nacos命名空间
	Endpoint             string `yaml:"endpoint"`                // Endpoint             获取nacos节点ip的服务地址
	CacheDir             string `yaml:"cacheDir"`                // CacheDir             缓存目录
	LogDir               string `yaml:"logDir"`                  // LogDir               日志目录
	UpdateThreadNum      int    `yaml:"update_thread_num"`       // UpdateThreadNum      更新服务的线程数
	NotLoadCacheAtStart  bool   `yaml:"not_load_cache_at_start"` // NotLoadCacheAtStart  在启动时不读取本地缓存数据，true--不读取，false--读取
	UpdateCacheWhenEmpty bool   `yaml:"update_cache_when_empty"` // UpdateCacheWhenEmpty 当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	Username             string `yaml:"username"`
	Password             string `yaml:"password"`
}

// NacosServerConfig nacos服务器配置
type NacosServerConfig struct {
	IpAddr      string `yaml:"ip_addr"`      // IpAddr      nacos命名空间
	ContextPath string `yaml:"context_path"` // ContextPath 获取nacos节点ip的服务地址
	Port        uint64 `yaml:"port"`         // Port        缓存目录
}

// HTTPServerConfig is http server config.
type HTTPServerConfig struct {
	Port         uint64        `yaml:"port"`
	Network      string        `yaml:"network"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

// DiscoveryConfig is nacos service config.
type DiscoveryConfig struct {
	GroupName   string  `yaml:"group_name"`
	ClusterName string  `yaml:"cluster_name"`
	Weight      float64 `yaml:"weight"`
	Enable      bool    `yaml:"enable"`
	Healthy     bool    `yaml:"healthy"`
	Ephemeral   bool    `yaml:"ephemeral"`
}

// ApplicationConfig is ApplicationConfig service config.
type ApplicationConfig struct {
	ServiceName string `yaml:"name"`
}
