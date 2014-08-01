package controllers

import (
    "github.com/revolvingcow/grassfed/app/models"
)

type Application struct {
    DatabaseController
}

func (c Application) Connected() *models.Account {
    if c.RenderArgs["account"] != nil {
        return c.RenderArgs["account"].(*models.Account)
    }

    if id, ok := c.Session["account"]; ok {
        return c.getAccount(id)
    }

    return nil
}

func (c Application) getAccount(id string) *models.Account {
    accounts, err := c.Transaction.Select(models.Account{}, `select * from Account where Profile = ?`, id)

    if err != nil {
        panic(err)
    }

    if len(accounts) == 0 {
        return nil
    }

    return accounts[0].(*models.Account)
}
