package devOpsTest

import (
	"fmt"
	"github.com/lvxin0315/devOps"
	"testing"
)

func Test_m3u8CMD(t *testing.T) {
	f := new(devOps.DockerFFmpegMp4ToM3u8)
	err := f.DockerMp4ToM3u8(`/Users/lvxin/Downloads/极客时间-Service_Mesh实战/01_课程介绍.mp4`, "123456789")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}

func Test_jpegCMD(t *testing.T) {
	f := new(devOps.DockerFFmpegMp4ToM3u8)
	err := f.DockerMp4CutJpeg("/Users/lvxin/Downloads/极客时间-Service_Mesh实战/01_课程介绍.mp4", "123456789", "00:01:00")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
