package devOps

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"

	"github.com/sirupsen/logrus"
)

type swarmCli struct {
	cli             *client.Client
	networkNameList []string
	serviceList     []SwarmService
	volumeIDList    []string
}

func NewSwarmCli() (*swarmCli, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	swarmCli := new(swarmCli)
	swarmCli.cli = cli
	return swarmCli, nil
}

/**
创建swarm 内部网络
@param string networkName 名称
*/
func (c *swarmCli) CreateSwarmNetwork(networkName string) (types.NetworkCreateResponse, error) {
	networkCreateResponse, err := c.cli.NetworkCreate(context.Background(), networkName, types.NetworkCreate{
		Driver: "overlay",
	})
	if err != nil {
		logrus.Error("CreateSwarmNetWork.err:", err)
		return networkCreateResponse, err
	}
	c.networkNameList = append(c.networkNameList, networkName)
	return networkCreateResponse, nil
}

/**
创建Service
@param Service service
*/
func (c *swarmCli) CreateSwarmService(service SwarmService) (types.ServiceCreateResponse, error) {
	if len(c.networkNameList) > 0 {
		for _, networkName := range c.networkNameList {
			service.SetNetwork(networkName)
		}
	}
	service.Init()
	serviceCreateResponse, err := c.cli.ServiceCreate(context.Background(), service.GetServiceSpec(), service.GetOption())
	if err != nil {
		logrus.Error("CreateService.err:", err)
		return serviceCreateResponse, err
	}
	//返回id
	service.SetID(serviceCreateResponse.ID)
	c.serviceList = append(c.serviceList, service)
	return serviceCreateResponse, nil
}

/**
创建swarm Volume
@param string networkName 名称
*/
func (c *swarmCli) CreateVolume(volumeName string) (types.Volume, error) {
	volumeTypes, err := c.cli.VolumeCreate(context.Background(), volume.VolumeCreateBody{
		Name: volumeName,
	})
	if err != nil {
		logrus.Error("CreateVolume.err:", err)
		return volumeTypes, err
	}
	//id保存
	c.volumeIDList = append(c.volumeIDList, volumeName)
	return volumeTypes, nil
}
