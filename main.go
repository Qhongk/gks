package main

//import (
//"github.com/astaxie/beego"
//)
import (
	_ "./routers"
	"github.com/astaxie/beego"
)

//
//type MainController struct {
//	beego.Controller
//}

//func (this *MainController) Get() {
//	//this.Ctx.WriteString("hello world")
//	this.Ctx.Redirect(200, "index")
//}

func main() {
	//beego.Router("/", &MainController{})
	beego.Run()
}