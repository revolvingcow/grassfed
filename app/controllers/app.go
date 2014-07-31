package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) About() revel.Result {
    return c.Render()
}

func (c App) Me() revel.Result {
    return c.Render()
}

func (c App) Stats() revel.Result {
    return c.Render()
}

func (c App) Add(id int, product string, calories int) revel.Result {
    c.Validation.Required(id).Message("You must be logged on to add more calories.")
    c.Validation.Required(product).Message("You must include a product.")
    c.Validation.Required(calories).Message("You must provide the amount of calories")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Index)
    }

    return c.Render()
}

func (c App) Goal() revel.Result {
    return c.Render()
}

func (c App) SetGoal() revel.Result {
    return c.Render()
}

func (c App) Streak() revel.Result {
    return c.Render()
}
