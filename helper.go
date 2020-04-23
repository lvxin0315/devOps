package devOps

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

var shellHead = "#!/bin/bash"
var expectShellHead = "#!/usr/bin/expect"

//创建目录
func mkdir(path string) error {
	//目录判断
	ok, err := pathExists(path)
	if !ok {
		if err != nil {
			return err
		}
		//新建目录
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

//判断目录是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//生成expectShell文件
func createExpectShellFile(shellCmd string, path string) (shellFileName string, err error) {
	shellFileName = fmt.Sprintf("%s/%d.sh", path, time.Now().UnixNano())
	shellFileContent := fmt.Sprintf(`%s
%s`, expectShellHead, shellCmd)
	err = ioutil.WriteFile(shellFileName, []byte(shellFileContent), os.ModePerm)
	if err != nil {
		return "", err
	}

	return shellFileName, nil
}

func createShellFile(shellCmd string, path string) (shellFileName string, err error) {
	shellFileName = fmt.Sprintf("%s/%d.sh", path, time.Now().UnixNano())
	shellFileContent := fmt.Sprintf(`%s
%s`, shellHead, shellCmd)
	err = ioutil.WriteFile(shellFileName, []byte(shellFileContent), os.ModePerm)
	if err != nil {
		return "", err
	}

	return shellFileName, nil
}

func cmd(cmdTpl string, params map[string]interface{}) error {
	tmpl, err := template.New("cmd_template").Parse(cmdTpl)
	if err != nil {
		logrus.Error("template：", err)
		return err
	}
	hw := new(bytes.Buffer)
	err = tmpl.Execute(hw, params)
	cmd := exec.Command("/bin/bash", "-c", hw.String())
	cmd.Stdout = logrus.New().Out
	if err := cmd.Run(); err != nil {
		logrus.Info("cmd.Run: ", hw.String())
		logrus.Error("cmd.Run: ", err)
		return err
	}
	return nil
}

//根据文件地址截取文件名称
func cutFileName(filePath string) (fileDir, fileName, onlyName, ext string) {
	fileDir, fileName = filepath.Split(filePath)
	ext = filepath.Ext(filePath)
	onlyName = strings.ReplaceAll(fileName, ext, "")
	return
}

//处理输出目录
func dirStat(outputDirPath string) error {
	fileInfo, err := os.Stat(outputDirPath)
	if os.IsNotExist(err) {
		//创建目录
		err := os.MkdirAll(outputDirPath, os.ModePerm)
		if err != nil {
			logrus.Error("MkdirAll：", err)
			return err
		}
	} else if !fileInfo.IsDir() {
		logrus.Error("MkdirAll：", err)
		return errors.New(outputDirPath + " is not dir")
	}
	return nil
}
