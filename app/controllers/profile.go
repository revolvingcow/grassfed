package controllers

import (
    "fmt"
    "strings"
    "time"
    "github.com/revel/revel"
    "github.com/revolvingcow/grassfed/app/models"
)

type Profile struct {
    Application
}

func (c Profile) getHistory(account *models.Account) []*models.History {
    if account == nil {
        return nil
    }

    results, err := c.Transaction.Select(
        models.History{},
        `select * from History where AccountId = ? order by Date desc`,
        account.Id)

    if err != nil {
        return nil
    }

    rows := len(results)
    if rows == 0 {
        return nil
    }

    history := make([]*models.History, rows)
    for i := 0; i < rows; i++ {
        history = append(history, results[i].(*models.History))
    }

    return history
}

func (c Profile) getCaloriesForDate(history []*models.History, date time.Time) (current int64) {
    current = 0

    if history != nil {
        for _, moment := range history {
            if moment != nil {
                local := moment.Date.Local()
                if local.Day() == date.Day() && local.Month() == date.Month() && local.Year() == date.Year() {
                    current += moment.Calories
                }
            }
        }
    }

    return current
}

func (c Profile) getStreak(history []*models.History, ceiling int64) (streak int64) {
    now := time.Now()
    streak = 0

    if history != nil && len(history) > 0 {
        interval := 1

        for {
            s := fmt.Sprintf("-%dh", interval * 24)
            duration, _ := time.ParseDuration(s)
            count := c.getCaloriesForDate(history, now.Add(duration))

            if count > 0 && ceiling > count {
                streak += 1
                interval += 1
            } else {
                break
            }
        }
    }

    return streak
}

func (c Profile) getMoment(id int64) *models.History {
    history, err := c.Transaction.Select(models.History{}, `select * from History where Id = ?`, id)
    if err != nil {
        panic(err)
    }

    if len(history) == 0 {
        return nil
    }

    return history[0].(*models.History)
}

func (c Profile) Index() revel.Result {
    return c.Render()
}

func (c Profile) Logon(id string) revel.Result {
    c.Response.ContentType = "application/json"
    c.Validation.Required(id).Message("You must be logged on.")

    if c.Validation.HasErrors() {
        revel.INFO.Println("Validation errors found.")
        c.Validation.Keep()
        c.FlashParams()
        return c.RenderJson(nil)
    }

    revel.INFO.Println("Setting up the variables for storage.")
    now := time.Now()
    account := c.getAccount(id)

    if account == nil {
        revel.INFO.Println("Creating account.")
        account = &models.Account{}
        account.Profile = id
        account.Goal = 2000
        account.Created = now
        account.LastVisit = now
        c.Transaction.Insert(account)
    } else {
        revel.INFO.Println("Updating account.")
        account.LastVisit = now
        c.Transaction.Update(account)
    }

    c.Session["account"] = id
    c.Session.SetDefaultExpiration()

    return c.RenderJson(true)
}

func (c Profile) History() revel.Result {
    account := c.Connected()
    if account == nil {
        return c.RenderJson(nil)
    }

    history := c.getHistory(account)
    return c.RenderJson(history)
}

func (c Profile) Stats() revel.Result {
    account := c.Connected()
    if account == nil {
        return c.RenderJson(nil)
    }

    history := c.getHistory(account)

    response := models.ResponseStatistics {
        Goal: account.Goal,
        Current: c.getCaloriesForDate(history, time.Now()),
        Streak: c.getStreak(history, account.Goal),
    }

    return c.RenderJson(response)
}

func (c Profile) Add(product string, calories int64) revel.Result {
    account := c.Connected()
    if account == nil || strings.TrimSpace(product) == "" {
        return c.RenderJson(nil)
    }

    c.Validation.Required(product).Message("You must include a product.")
    c.Validation.Required(calories).Message("You must provide the amount of calories")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.RenderJson(nil)
    }

    moment := models.History {
        AccountId: account.Id,
        Product: product,
        Calories: calories,
        Date: time.Now(),
    }
    c.Transaction.Insert(&moment)

    return c.RenderJson(moment)
}

func (c Profile) Delete(id int64) revel.Result {
    account := c.Connected()
    if account == nil {
        return c.RenderJson(nil)
    }

    moment := c.getMoment(id)
    if moment == nil {
        return c.RenderJson(nil)
    }
    c.Transaction.Delete(moment)

    return c.RenderJson(true)
}

func (c Profile) Goal(calories int64) revel.Result {
    account := c.Connected()
    if account == nil {
        return c.RenderJson(nil)
    }

    account.Goal = calories
    c.Transaction.Update(account)

    return c.RenderJson(true)
}
