package routers

import (
	"./../controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Get("/aaa", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world..."))
	})

}
