package routers

import (
    _ "github.com/go-sql-driver/mysql" // import your used driver
    _ "gks/models"
    _ "gks/utils/cache"
    "github.com/astaxie/beego"
    "gks/controllers"
    "github.com/astaxie/beego/orm"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/tpl", &controllers.MainController{}, "get:HelloTpl")
    beego.Router("/article", &controllers.ArticleController{})
    beego.Router("/api", &controllers.ApiController{})
    beego.Router("/json", &controllers.JsonController{})

    beego.DelStaticPath("/static")
    beego.SetStaticPath("/img", "img")
    beego.SetStaticPath("/css", "css")
    beego.SetStaticPath("/js", "js")

    beego.AppConfig.Set("ext", "html")

    beego.ErrorController(&controllers.ErrorController{})

    username := beego.AppConfig.String("db.user")
    pwd := beego.AppConfig.String("db.pwd")
    host := beego.AppConfig.String("db.host")
    port := beego.AppConfig.String("db.port")
    name := beego.AppConfig.String("db.name")

    // set mysql driver
    orm.RegisterDriver("mysql", orm.DRMySQL)

    // max idle conn
    maxIdle := 30
    // max conn
    maxConn := 30
    // set default database
    orm.RegisterDataBase("default", "mysql", username+":"+pwd+"@tcp("+host+":"+port+")/"+name+"?charset=utf8", maxIdle, maxConn)

    // create table
    orm.RunSyncdb("default", false, false)
}
