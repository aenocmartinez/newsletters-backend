package domain_test

import (
	"pulzo/src/newsletter/domain"
	"pulzo/src/shared/infraestructure/util"
	"testing"
	"time"
)

func TestWasExecuted(t *testing.T) {
	currentSchedule := domain.Schedule{}

	currentSchedule.SetDate("2023-07-16")
	if !currentSchedule.WasExecuted() {
		t.Error("se espera que haya sido ejecutado")
	}

	currentSchedule.SetDate("2023-07-18")
	if currentSchedule.WasExecuted() {
		t.Error("se espera que no haya sido ejecutado")
	}

	today := time.Now()
	currentSchedule.SetDate(today.Format(util.YYYY_MM_DD))

	dentroDeUnaHora := today.Add(time.Hour * -1)
	currentSchedule.SetHour(dentroDeUnaHora.Format(util.HH_MM_SS))
	if !currentSchedule.WasExecuted() {
		t.Error("se espera que haya sido ejecutado")
	}
}
