package devOps

import (
	"fmt"
	"os/exec"
	"strings"
)

type dumper struct {
	Host          string
	User          string
	Password      string
	Database      string
	MysqlDumpPath string
	TableNameList []string
	SavePath      string
}

func NewDumper() *dumper {
	return new(dumper)
}

func (d *dumper) Do() {
	//生成结构目录
	err := mkdir(d.SavePath)
	if err != nil {
		errorLog(err.Error())
		panic(err)
	}
	//生成命令
	cmd := exec.Command(d.MysqlDumpPath,
		d.Database,
		"-u", d.User,
		fmt.Sprintf("-p%s", d.Password),
		"-h", d.Host,
		"--tables", strings.Join(d.TableNameList, " "),
		fmt.Sprintf(">%s/mysqldump.sql", d.SavePath))
	infoLog("命令内容：")
	info2Log(cmd.String())
	//写入到shell
	createShellFile, err := createShellFile(cmd.String(), d.SavePath)
	if err != nil {
		errorLog(err.Error())
		panic(err)
	}
	//执行shell文件
	cmd = exec.Command("bash", createShellFile)
	//开始执行
	infoLog("开始执行")
	info2Log(createShellFile)
	err = cmd.Run()
	if err != nil {
		errorLog(err.Error())
		panic(err)
	}
}
