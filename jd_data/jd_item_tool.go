package main

import (
	"fmt"
	"github.com/lvxin0315/devOps/jd_data/item"
)

func main() {
	//建表
	item.InitMysqlDB(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"root",
		"127.0.0.1",
		3306,
		"jd_data",
	))
	tableList := []interface{}{
		item.JdItemModel{},
	}
	for i, t := range tableList {
		err := item.NewDB().CreateTable(t).Error
		if err != nil {
			fmt.Println(i)
			panic(err)
		}
	}
}
