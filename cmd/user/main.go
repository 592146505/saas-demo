package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"saas-demo/common/nacos"
	conf2 "saas-demo/internal/user/conf"
	"saas-demo/internal/user/http"
)

func main() {
	http.New(conf2.HTTPServer).Run()
	nacosClient, err := nacos.NewNacosClient(conf2.Nacos)
	if err != nil {
		panic(err)
	}

	ok, err := nacosClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          conf2.IP,
		Port:        conf2.HTTPServer.Port,
		ServiceName: conf2.Application.ServiceName,
		GroupName:   conf2.Discovery.GroupName,
		ClusterName: conf2.Discovery.ClusterName,
		Weight:      conf2.Discovery.Weight,
		Enable:      conf2.Discovery.Enable,
		Healthy:     conf2.Discovery.Healthy,
		Ephemeral:   conf2.Discovery.Ephemeral,
	})
	if !ok {
		fmt.Printf("RegisterInstance  server[%s] fail\n", "saas-demo.go")
	}
	service, err := nacosClient.GetService(vo.GetServiceParam{
		ServiceName: conf2.Application.ServiceName,
		GroupName:   conf2.Discovery.GroupName,
	})
	if err == nil {
		js, _ := json.MarshalIndent(service, "", "	")
		fmt.Printf("GetService server[%v]\n", string(js))
	}
	services, err := nacosClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		GroupName: conf2.Discovery.GroupName,
		PageNo:    1,
		PageSize:  20,
	})
	if err == nil {
		js, _ := json.MarshalIndent(services, "", "	")
		fmt.Printf("GetAllServicesInfo servers[%v]\n", string(js))
	}
	instances, err := nacosClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: conf2.Application.ServiceName,
		GroupName:   conf2.Discovery.GroupName,
	})
	if err == nil {
		js, _ := json.MarshalIndent(instances, "", "	")
		fmt.Printf("SelectAllInstances instances[%v] \n", string(js))
	}
	instances, err = nacosClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: conf2.Application.ServiceName,
		GroupName:   conf2.Discovery.GroupName,
		HealthyOnly: true,
	})
	if err == nil {
		js, _ := json.MarshalIndent(instances, "", "	")
		fmt.Printf("SelectInstances instances[%v] \n", string(js))
	}

	instance, err := nacosClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: conf2.Application.ServiceName,
		GroupName:   conf2.Discovery.GroupName,
	})
	if err == nil {
		js, _ := json.MarshalIndent(instance, "", "	")
		fmt.Printf("SelectOneHealthyInstance instance[%v] \n", string(js))
	}

	err = nacosClient.Subscribe(&vo.SubscribeParam{
		ServiceName: conf2.Application.ServiceName,
		GroupName:   conf2.Discovery.GroupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			js, _ := json.MarshalIndent(services, "", "	")
			fmt.Printf("Subscribe services[%v] \n", string(js))
		},
	})
	if err != nil {
		fmt.Printf("Subscribe fail[%v] \n", err)
	}
	for {
		fmt.Print()
	}
}
