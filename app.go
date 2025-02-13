package main

import (
	"context"
	"fmt"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Prolabore(workDayPerWeek int, hoursPerDay int, pretendedSalary int) string {
	t := time.Date(time.Now().Year(), time.Now().Month(), 32, 0, 0, 0, 0, time.UTC)
	daysInMonth := 32 - t.Day()
	weeksInMonth := float64(daysInMonth / 7)
	salaryPerHour := pretendedSalary / (workDayPerWeek * hoursPerDay * int(weeksInMonth))
	if salaryPerHour <= 0 {
		return "Horas trabalhadas muito baixas para o salário desejado"
	}
	return fmt.Sprintf("Você deveria cobrar %d por hora este mês", salaryPerHour)
}
