package controllers

import (
	"got-cep/app/services"

	"github.com/revel/revel"
)

type Zipcode struct {
	App
}

func (c Zipcode) Index(zipcode string) revel.Result {

	zipdata := services.Zipcode(zipcode)
	c.Response.Status = zipdata.Status
	return c.RenderText(zipdata.Json)
}
