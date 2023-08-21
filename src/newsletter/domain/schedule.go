package domain

import (
	"fmt"
	"os"
	"pulzo/src/shared/infraestructure/util"
	"time"
)

type Schedule struct {
	idCampaign string
	date       time.Time
	hour       time.Time
	message    string
	subject    string
	from       string
}

func (s *Schedule) SetCampaignId(idCampaign string) {
	s.idCampaign = idCampaign
}

func (s *Schedule) CampaignId() string {
	return s.idCampaign
}

func (s *Schedule) SetDate(strDate string) {
	date, err := time.Parse(util.YYYY_MM_DD, strDate)
	if err != nil {
		date = time.Now()
	}
	s.date = date
}

func (s *Schedule) Date() string {
	return s.date.Format(util.YYYY_MM_DD)
}

func (s *Schedule) SetHour(strHour string) {
	hour, err := time.Parse(util.HH_MM, strHour)
	if err != nil {
		hour = time.Now()
	}
	s.hour = hour
}

func (s *Schedule) Hour() string {
	hour := s.hour.Format(util.HH)
	minute := s.hour.Format(util.MM)

	return fmt.Sprintf("%s:%s:00", hour, minute)
}

func (s *Schedule) SetMessage(message string) {
	s.message = message
}

func (s *Schedule) Message() string {
	return s.message
}

func (s *Schedule) SetSubject(subject string) {
	s.subject = subject
}

func (s *Schedule) Subject() string {
	return s.subject
}

func (s *Schedule) SetFrom(from string) {
	if len(from) == 0 {
		from = os.Getenv("APP_SEND_FROM_DEFAULT")
	}
	s.from = from
}

func (s *Schedule) From() string {
	return s.from
}

func (s *Schedule) Exists() bool {
	return s.idCampaign != ""
}

func (s *Schedule) WasExecuted() bool {

	loc, _ := time.LoadLocation("America/Bogota")

	today := time.Now().In(loc)
	if today.Before(s.date) {
		return false
	}

	isEqualDate := (today.Year() == s.date.Year() && today.YearDay() == s.date.YearDay())
	if !isEqualDate {
		return true
	}

	currentHour := time.Date(0000, 1, 1, today.Hour(), today.Minute(), 0, 0, time.UTC)
	scheduleHour := time.Date(0000, 1, 1, s.hour.Hour(), s.hour.Minute(), 0, 0, time.UTC)

	return scheduleHour.Before(currentHour)
}

func (s *Schedule) Status() string {
	if !s.Exists() {
		return "Sin programación"
	}

	if s.WasExecuted() {
		return "enviado"
	}

	return "programado"
}
