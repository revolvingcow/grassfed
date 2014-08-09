package controllers

import (
	"github.com/revel/revel"
	"github.com/revolvingcow/grassfed/app/models"
)

type Home struct {
	Application
}

func (c Home) getNumberOfAccounts() (count int64) {
	count, err := c.Transaction.SelectInt(`select count(*) from Account`)
	if err != nil {
		return 0
	}

	return count
}

func (c Home) getNumberOfCalories() (calories int64) {
	calories, err := c.Transaction.SelectInt(`select sum(Calories) from History`)
	if err != nil {
		return 0
	}

	return calories
}

func (c Home) Index() revel.Result {
	return c.Render()
}

func (c Home) About() revel.Result {
	return c.Render()
}

func (c Home) Overview() revel.Result {
	model := models.Overview{
		Accounts: c.getNumberOfAccounts(),
		Calories: c.getNumberOfCalories(),
	}

	return c.RenderJson(model)
}
