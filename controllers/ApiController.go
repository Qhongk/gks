package controllers

import "github.com/astaxie/beego"

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	c.Data["xml"] = &countryCapitalMap
	c.ServeXML()
}