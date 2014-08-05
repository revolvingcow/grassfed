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

func (c Profile) getGoal(account *models.Account) (goal int64) {
    goal = 2000

    if account == nil {
        return goal
    }

    results := models.Goal{}
    err := c.Transaction.SelectOne(
        &results,
        `select * from Goal where AccountId = ? order by Date desc limit 1`,
        account.Id)

    if err != nil {
        return goal
    }

    goal = results.Calories
    return goal
}

func (c Profile) setGoal(account *models.Account, calories int64) {
    goals, err := c.Transaction.Select(
        models.Goal{},
        `select * from Goal where AccountId = ? order by Date desc limit 1`,
        account.Id)

    if err != nil {
        revel.INFO.Println(err)
        return
    }

    now := time.Now().Local()

    if len(goals) > 0 {
        goal := goals[0].(*models.Goal)
        local := goal.Date.Local()

        if now.Day() == local.Day() && now.Month() == local.Month() && now.Year() == local.Year() {
            goal.Calories = calories
            c.Transaction.Update(goal)
        } else {
            newGoal := models.Goal{ AccountId: account.Id, Calories: calories, Date: now }
            c.Transaction.Insert(&newGoal)
        }
    } else {
        newGoal := models.Goal{ AccountId: account.Id, Calories: calories, Date: now }
        c.Transaction.Insert(&newGoal)
    }
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
    account := c.Connected()
    return c.Render(account)
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

    goal := c.getGoal(account)
    history := c.getHistory(account)

    response := models.ResponseStatistics {
        Goal: goal,
        Current: c.getCaloriesForDate(history, time.Now()),
        Streak: c.getStreak(history, goal),
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

    c.setGoal(account, calories)
    return c.RenderJson(true)
}
