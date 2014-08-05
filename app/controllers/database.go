package controllers

import (
    "database/sql"
    "github.com/coopernurse/gorp"
    _ "github.com/mattn/go-sqlite3"
    "github.com/revel/revel"
    "github.com/revel/revel/modules/db/app"
    "github.com/revolvingcow/grassfed/app/models"
)

var (
    DbMap *gorp.DbMap
)

func Initialize() {
    db.Init()
    DbMap = &gorp.DbMap{ Db: db.Db, Dialect: gorp.SqliteDialect{} }

    DbMap.AddTable(models.Account{}).SetKeys(true, "Id")
    DbMap.AddTable(models.History{}).SetKeys(true, "Id")
    DbMap.AddTable(models.Goal{}).SetKeys(true, "Id")
    DbMap.AddTable(models.Weight{}).SetKeys(true, "Id")

    DbMap.TraceOn("[db]", revel.INFO)
    DbMap.CreateTablesIfNotExists()
}

type DatabaseController struct {
    *revel.Controller
    Transaction *gorp.Transaction
}

func (c *DatabaseController) Begin() revel.Result {
    transaction, err := DbMap.Begin()

    if err != nil {
        panic(err)
    }

    c.Transaction = transaction
    return nil
}

func (c *DatabaseController) Commit() revel.Result {
    if c.Transaction == nil {
        return nil
    }

    if err := c.Transaction.Commit(); err != nil && err != sql.ErrTxDone {
        panic(err)
    }

    c.Transaction = nil
    return nil
}

func (c *DatabaseController) Rollback() revel.Result {
    if c.Transaction == nil {
        return nil
    }

    if err := c.Transaction.Rollback(); err != nil && err != sql.ErrTxDone {
        panic(err)
    }

    c.Transaction = nil
    return nil
}
