package controllers

import (
    "github.com/revel/revel"
)

type Home struct {
    Application
}

func (c Home) Index() revel.Result {
	return c.Render()
}

func (c Home) About() revel.Result {
    return c.Render()
}
