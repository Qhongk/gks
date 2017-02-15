package routers

import (
    "github.com/astaxie/beego"
    "gks/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/tpl", &controllers.MainController{}, "get:HelloTpl")
    beego.Router("/article", &controllers.ArticleController{})
    beego.Router("/api", &controllers.ApiController{})
    beego.Router("/json", &controllers.JsonController{})

    //beego.Get("/aaa", func(ctx *context.Context) {
    //	ctx.Output.Body([]byte("hello world..."))
    //})

    //beego.SetViewsPath("views")
    beego.DelStaticPath("/static")
    beego.SetStaticPath("/img", "img")
    beego.SetStaticPath("/css", "css")
    beego.SetStaticPath("/js", "js")

    beego.AppConfig.Set("ext", "html")
}
