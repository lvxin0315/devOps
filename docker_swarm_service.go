package devOps

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/swarm"
)

const (
	ReplicasAny = uint64(1)
)

type SwarmService interface {
	Init()
	GetOption() types.ServiceCreateOptions
	GetServiceSpec() swarm.ServiceSpec
	SetNetwork(target string)
	SetID(id string)
}

type GeneralService struct {
	ID                      string
	ServiceSpec             swarm.ServiceSpec
	Option                  types.ServiceCreateOptions
	NetworkAttachmentConfig []swarm.NetworkAttachmentConfig
	PortConfig              []swarm.PortConfig
	MountList               []mount.Mount
	AnnotationsName         string   //service name
	ContainerSpecImageName  string   //使用image
	Replicas                uint64   //重启方式
	EnvList                 []string //环境变量
}

func (service *GeneralService) GetOption() types.ServiceCreateOptions {
	return service.Option
}

func (service *GeneralService) GetServiceSpec() swarm.ServiceSpec {
	return service.ServiceSpec
}

func (service *GeneralService) SetID(id string) {
	service.ID = id
}

func (service *GeneralService) Init() {
	service.ServiceSpec = swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: service.AnnotationsName,
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: &swarm.ContainerSpec{
				Image:  service.Image(),
				Mounts: service.MountList,
				Env:    service.EnvList,
			},
		},
		Mode: swarm.ServiceMode{
			Replicated: &swarm.ReplicatedService{
				Replicas: service.GetReplicas(),
			},
		},
		Networks: service.NetworkAttachmentConfig,
		EndpointSpec: &swarm.EndpointSpec{
			Ports: service.PortConfig,
		},
		UpdateConfig: service.GetUpdateConfig(),
	}
	service.Option = types.ServiceCreateOptions{}
}

/**
设置网络
*/
func (service *GeneralService) SetNetwork(target string) {
	service.NetworkAttachmentConfig = append(service.NetworkAttachmentConfig, swarm.NetworkAttachmentConfig{
		Target: target,
	})
}

/**
设置端口配置
*/
func (service *GeneralService) SetPortConfig(protocol swarm.PortConfigProtocol, targetPort, publishedPort uint32) {
	service.PortConfig = append(service.PortConfig, swarm.PortConfig{
		Protocol:      protocol,
		TargetPort:    targetPort,
		PublishedPort: publishedPort,
	})
}

/**
设置服务名称
*/
func (service *GeneralService) SetAnnotationsName(annotationsName string) {
	service.AnnotationsName = annotationsName
}

/**
设置镜像名称
*/
func (service *GeneralService) SetContainerSpecImage(containerSpecImage string) {
	service.ContainerSpecImageName = containerSpecImage
}

func (service *GeneralService) SetReplicas(replicas uint64) {
	service.Replicas = replicas
}

func (service *GeneralService) GetReplicas() *uint64 {
	if service.Replicas == 0 {
		service.Replicas = ReplicasAny
	}
	return &service.Replicas
}

/**
设置环境变量
*/
func (service *GeneralService) SetEnv(key, value string) {
	service.EnvList = append(service.EnvList, fmt.Sprintf("%s=%s", key, value))
}

func (service *GeneralService) GetEnv() []string {
	if len(service.EnvList) == 0 {
		return nil
	}
	return service.EnvList
}

/**
服务更新方式
*/
func (service *GeneralService) GetUpdateConfig() *swarm.UpdateConfig {
	return &swarm.UpdateConfig{
		FailureAction: "pause",
		Order:         "start-first",
	}
}

/**
容器名称
*/
func (service *GeneralService) Image() string {
	return service.ContainerSpecImageName
}

/**
设置mount
*/
func (service *GeneralService) SetMount(mountType mount.Type, target, source string) {
	service.MountList = append(service.MountList, mount.Mount{
		Type:     mountType,
		Source:   source,
		Target:   target,
		ReadOnly: false,
	})
}
