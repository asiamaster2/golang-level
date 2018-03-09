package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Create(username, password string) revel.Result {
    if username == "hylee" && password == "aaron11!" {
        var resultmsg string = "1.1.1.1"
        return c.Render(resultmsg)
    }
    var resultmsg string = "Please check your credential."
    return c.Render(resultmsg)
}

func (c App) Healthcheck() revel.Result {
    return c.Render()
}

