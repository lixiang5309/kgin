package db

import (
	"log"
	"time"

	"kgin/utils"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Db   *xorm.Engine
	Mgdb *mgo.Collection
)

type ApiLog struct {
	Id      bson.ObjectId `bson:"_id"`
	Url     string        `bson:"url"`
	Project string        `bson:"project"`
	Method  string        `bson:"method"`
	Date    string        `bson:"date"`
	Header  interface{}   `bson:"header"`
	Param   interface{}   `bson:"param"`
	Data    interface{}   `bson:"data"`
}

func Init() {
	dbconn, err := utils.LoadDbConfig("test")
	if err == nil {
		Db, err = xorm.NewEngine("mysql", dbconn)
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
		Db.SetMaxOpenConns(1000)
		Db.SetMaxIdleConns(300)
		Db.SetConnMaxLifetime(time.Hour * 1)
	}

	mgdb, err := utils.LoadConfig("mongo")
	if err == nil {
		// 建立mongodb连接
		mgconn, err := mgo.Dial(mgdb["hostname"] + ":" + mgdb["port"])
		if err == nil {
			mgconn.SetPoolLimit(1000)
			//切换到数据库,没有会创建
			mgdb := mgconn.DB("blog")
			err = mgdb.Login("test", "123456")
			if err != nil {
				log.Printf("this:error: %v", err.Error())
			}
			//切换到collection,没有会创建
			Mgdb = mgdb.C("log")
		} else {
			log.Printf("thismongo:error: %v", err.Error())
		}
	}
}

func Close() {
	Db.Close()
	Mgdb.DropCollection()
}
