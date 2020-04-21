package devOpsTest

import (
	"github.com/lvxin0315/devOps"
	"log"
	"os"
	"testing"
)

func Test_DockerSwarmClientCreateSwarmNetwork(t *testing.T) {
	os.Setenv("DOCKER_HOST", "http://172.16.0.228:2375")
	os.Setenv("DOCKER_API_VERSION", "1.39")
	swarmCli, err := devOps.NewSwarmCli()
	if err != nil {
		t.Error(err)
	}
	//创建一个network
	res, err := swarmCli.CreateSwarmNetwork("test123123")
	if err != nil {
		t.Error(err)
	}
	log.Println(res.ID)
}
