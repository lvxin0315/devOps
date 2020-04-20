package devOps

import (
	"fmt"
	"io/ioutil"
	"os"
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
