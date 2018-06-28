package sqllite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type DbEngine struct {
	Engine *gorm.DB
}

func buildDbEngine() *DbEngine{
	db,err:=initEngine()

	if err!=nil{
		log.Fatal(err.Error())
		return nil
	}

	db.LogMode(true)

	return &DbEngine{Engine:db}

}

var (
	GlobalSqlLiteDbEngine *DbEngine
	DBFileName="./UserRole.db"
)

func GetDbEngine()*DbEngine{
	if GlobalSqlLiteDbEngine==nil {
		GlobalSqlLiteDbEngine=buildDbEngine()
	}

	return GlobalSqlLiteDbEngine
}


func initEngine() (*gorm.DB,error) {
	return gorm.Open("sqlite3",DBFileName)
}
