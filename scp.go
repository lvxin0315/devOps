package devOps

import (
	"fmt"
	"os/exec"
)

type scpInfo struct {
	Host              string
	User              string
	Password          string
	Port              int
	ReadFilePath      string //可读文件内容
	LocalSaveFilePath string //保存文件内容
}

func NewScpInfo() *scpInfo {
	return new(scpInfo)
}

func (s *scpInfo) Do() {
	//生成结构目录
	err := mkdir(s.LocalSaveFilePath)
	if err != nil {
		errorLog(err.Error())
		panic(err)
	}
	//生成命令
	cmdString := fmt.Sprintf(`set timeout -1
spawn scp -r -P %d %s@%s:%s %s
expect {
 "(yes/no)?"
  {
    send "yes\n"
    expect "*assword:" { send "%s\n"}
  }
 "*assword:"
  {
    send "%s\n"
  }
}
expect "100%%"
expect eof`, s.Port, s.User, s.Host, s.ReadFilePath, s.LocalSaveFilePath, s.Password, s.Password)
	infoLog("命令内容：")
	info2Log(cmdString)
	//写入到shell
	createShellFile, err := createExpectShellFile(cmdString, s.LocalSaveFilePath)
	if err != nil {
		errorLog(err.Error())
		panic(err)
	}
	//执行shell文件
	cmd := exec.Command("expect", createShellFile)
	//开始执行
	infoLog("开始执行")
	info2Log(createShellFile)
	err = cmd.Run()
	if err != nil {
		errorLog(err.Error())
		panic(err)
	}
}
