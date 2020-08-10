package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/pkg/errors"
	"saas-demo/pkg/conf"
)

type NacosClient struct {
	Conf         *conf.NacosConfig
	namingClient naming_client.INamingClient
	configClient config_client.IConfigClient
}

var (
	emptyClientConf = conf.NacosClientConfig{}
	defaultClient   = constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 30 * 1000,
		BeatInterval:   5 * 1000,
	}
)

func NewNacosClient(conf *conf.NacosConfig) (*NacosClient, error) {
	if len(conf.Server) == 0 {
		return nil, errors.Errorf("nacos server config cannot empty\n")
	}
	serverConfig := newServerConfigs(conf.Server)
	clientConfig := newClientConfig(conf.Client)

	naming, err := newNamingClient(serverConfig, clientConfig)
	if err != nil {
		return nil, err
	}
	config, err := newConfigClient(serverConfig, clientConfig)
	if err != nil {
		return nil, err
	}
	return &NacosClient{
		namingClient: naming,
		configClient: config,
	}, nil
}

func newClientConfig(conf conf.NacosClientConfig) constant.ClientConfig {
	if conf == emptyClientConf {
		return defaultClient
	}
	return constant.ClientConfig{
		NamespaceId:    conf.NamespaceId,
		TimeoutMs:      conf.TimeoutMs,
		ListenInterval: conf.ListenInterval,
		BeatInterval:   conf.BeatInterval,
		LogDir:         conf.LogDir,
		CacheDir:       conf.CacheDir,
		Username:       conf.Username,
		Password:       conf.Password,
	}
}
func newServerConfigs(confs map[string]conf.NacosServerConfig) []constant.ServerConfig {
	var ss []constant.ServerConfig
	for _, conf := range confs {
		ss = append(ss, constant.ServerConfig{
			IpAddr:      conf.IpAddr,
			Port:        conf.Port,
			ContextPath: conf.ContextPath,
		})
	}
	return ss
}

func newNamingClient(ss []constant.ServerConfig, c constant.ClientConfig) (naming_client.INamingClient, error) {
	return clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": ss,
		"clientConfig":  c,
	})
}

func newConfigClient(ss []constant.ServerConfig, c constant.ClientConfig) (config_client.IConfigClient, error) {
	return clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": ss,
		"clientConfig":  c,
	})
}

func (n *NacosClient) RegisterInstance(param vo.RegisterInstanceParam) (bool, error) {
	return n.namingClient.RegisterInstance(param)
}

func (n *NacosClient) DeregisterInstance(param vo.DeregisterInstanceParam) (bool, error) {
	return n.namingClient.DeregisterInstance(param)
}

func (n *NacosClient) GetService(param vo.GetServiceParam) (model.Service, error) {
	return n.namingClient.GetService(param)
}

func (n *NacosClient) GetAllServicesInfo(param vo.GetAllServiceInfoParam) (model.ServiceList, error) {
	return n.namingClient.GetAllServicesInfo(param)
}

func (n *NacosClient) SelectAllInstances(param vo.SelectAllInstancesParam) ([]model.Instance, error) {
	return n.namingClient.SelectAllInstances(param)
}

func (n *NacosClient) SelectInstances(param vo.SelectInstancesParam) ([]model.Instance, error) {
	return n.namingClient.SelectInstances(param)
}

func (n *NacosClient) SelectOneHealthyInstance(param vo.SelectOneHealthInstanceParam) (*model.Instance, error) {
	return n.namingClient.SelectOneHealthyInstance(param)
}

func (n *NacosClient) Subscribe(param *vo.SubscribeParam) error {
	return n.namingClient.Subscribe(param)
}

func (n *NacosClient) Unsubscribe(param *vo.SubscribeParam) error {
	return n.namingClient.Unsubscribe(param)
}
