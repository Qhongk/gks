package controllers

import "github.com/astaxie/beego"
import (
	"encoding/xml"
)

type ApiController struct {
	beego.Controller
}

type mystruct struct {
	FieldOne string `xml:"field_one"`
	FieldTwo string `xml:"field_two"`
}

func (c *ApiController) Get() {
	mystruct := mystruct{ "123", "333" }
	c.Data["xml"] = &mystruct
	c.ServeXML()
}

// StringMap is a map[string]string.
type StringMap map[string]string

// StringMap marshals into XML.
func (s StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tokens := []xml.Token{start}

	for key, value := range s {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}