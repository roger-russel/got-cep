package controllers

import (
	"got-cep/app/services"

	"github.com/revel/revel"
)

type Zipcode struct {
	App
}

type sZipCode struct {
	State          string
	City           string
	Neighborhood   string
	Street         string
	IBGE           string
	additionalInfo string
	Sector         string
}

func (c Zipcode) Index(zipcode int) revel.Result {

	body := services.Zipcode(zipcode)
	data := sZipCode{Street: "ss asd s", City: body}
	return c.RenderJSON(data)
}
