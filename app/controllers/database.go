package controllers

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
	"github.com/revolvingcow/grassfed/app/models"
)

// DbMap is the database mapping used throughout the database controller.
var (
	DbMap *gorp.DbMap
)

// Initialize the database and create a database mapping.
func Initialize() {
	db.Init()
	DbMap = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	DbMap.AddTable(models.Account{}).SetKeys(true, "Id")
	DbMap.AddTable(models.History{}).SetKeys(true, "Id")
	DbMap.AddTable(models.Goal{}).SetKeys(true, "Id")
	DbMap.AddTable(models.Weight{}).SetKeys(true, "Id")

	DbMap.TraceOn("[db]", revel.INFO)
	DbMap.CreateTablesIfNotExists()
}

// DatabaseController wraps the Revel controller together with the database Transaction.
type DatabaseController struct {
	*revel.Controller
	Transaction *gorp.Transaction
}

// Begin creates a new transaction to be used.
func (c *DatabaseController) Begin() revel.Result {
	transaction, err := DbMap.Begin()

	if err != nil {
		panic(err)
	}

	c.Transaction = transaction
	return nil
}

// Commit attempts to finalize all database commands.
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

// Rollback will revert all recent changes returning the database to its original state. 
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
