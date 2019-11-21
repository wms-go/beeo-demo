package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/wms-go/beeo-demo/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "bigbay_dev:BigbayDev123@tcp(rm-2ze38nkyoj60l0l9zro.mysql.rds.aliyuncs.com:3306)/test")
	beego.Run()
}

