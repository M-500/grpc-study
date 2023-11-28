package go_xorm_study

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/27 17:08
//
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

var dbEngine *xorm.Engine

func InitDb() {
	if dbEngine != nil {
		return
	}
	sourceName := "admin:123456@tcp(192.168.1.52:3306)/test_xorm?charset=utf8mb4"
	if engine, err := xorm.NewEngine("mysql", sourceName); err != nil {
		log.Fatalf("db_helper.InitDb(%s),%v\n", sourceName, err)
	} else {
		dbEngine = engine
	}
	dbEngine.SetMaxIdleConns(2)
	dbEngine.SetMaxOpenConns(10)
	dbEngine.SetConnMaxLifetime(time.Minute * time.Duration(15))

}

func GetDb() *xorm.Engine {
	return dbEngine
}
