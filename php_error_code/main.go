package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

var allFileList []string
var allPHPFileList []string
var errorMsgList []string

func main() {
	err := tree("/Users/lvxin/PhpstormProjects/new_agent/application", 0)
	if err != nil {
		panic(err)
	}

	err = getPHPFile()
	if err != nil {
		panic(err)
	}

	for _, fp := range allPHPFileList {
		readLineWithCode(fp)
	}

	errorMsgList = removeRepeatedElement(errorMsgList)

	for _, msg := range errorMsgList {
		fmt.Println(msg)
	}
}

func tree(dstPath string, level int) error {
	dstF, err := os.Open(dstPath)
	if err != nil {
		return err
	}
	defer dstF.Close()
	fileInfo, err := dstF.Stat()
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() { //如果dstF是文件
		//for i:=0;i < level;i++ {
		//	fmt.Print("--")
		//}
		//fmt.Println(dstPath)
		allFileList = append(allFileList, dstPath)
		return nil
	} else { //如果dstF是文件夹
		//for i:=0;i < level;i++ {
		//	fmt.Print("--")
		//}
		//fmt.Println(dstF.Name())
		dir, err := dstF.Readdir(0) //获取文件夹下各个文件或文件夹的fileInfo
		if err != nil {
			return err
		}
		for _, fileInfo = range dir {
			err = tree(dstPath+"/"+fileInfo.Name(), level+1)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func getPHPFile() error {
	for _, fp := range allFileList {
		fileExt := path.Ext(fp)
		if fileExt == ".php" {
			allPHPFileList = append(allPHPFileList, fp)
		}
	}
	return nil
}

//按行读文件，找到包含 「 $this->error( 」
func readLineWithCode(path string) {
	code := "$this->error("
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		if strings.Index(line, code) >= 0 {
			lineStr := strings.ReplaceAll(line, " ", "")
			lineStr = strings.ReplaceAll(lineStr, code, "")
			lineStr = strings.ReplaceAll(lineStr, ");", "")
			lineStr = strings.ReplaceAll(lineStr, "//", "")
			lineStr = strings.ReplaceAll(lineStr, string('\n'), "")
			errorMsgList = append(errorMsgList, lineStr)
		}
	}
}

func removeRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
