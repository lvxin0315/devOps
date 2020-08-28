package item

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gormDB *gorm.DB

//初始化数据库
func InitMysqlDB(conn string) {
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(5)
	//plugin
	//db.Callback().Create().After("gorm:create").Register("plugin:gg_after_create", ggAfterCreate)
	//db.Callback().Query().Before("gorm:query_destination").Register("plugin:gg_before_query_destination", ggBeforeQueryDestination)
	gormDB = db
}

func NewDB() *gorm.DB {
	return gormDB.New()
}
