package controllers

import "github.com/astaxie/beego"

type JsonController struct {
    beego.Controller
}

type mystructJson struct {
    FieldOne string `json:"field_one"`
    FieldTwo map[string]string `json:"field_two"`
}

func (c *JsonController) Get() {
    countryCapitalMap := make(map[string]string)

    /* map 插入 key-value 对，各个国家对应的首都 */
    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Rome"
    countryCapitalMap["Japan"] = "Tokyo"
    countryCapitalMap["India"] = "New Delhi"

    mystruct := mystructJson{"123", countryCapitalMap }

    c.Data["json"] = &mystruct
    c.ServeJSON(true)
}
