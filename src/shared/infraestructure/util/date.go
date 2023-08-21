package util

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
)

var HH string = "15"
var MM string = "04"
var SS string = "05"
var HH_MM string = "15:04"
var HH_MM_SS string = "15:04:05"
var YYYY_MM_DD string = "2006-01-02"
var YYYY_MM_DD_HH_MM_SS string = "2006-01-02 15:04:05"

func ItIsValidHour(hour time.Time) bool {
	isZero := hour.Hour() == 0 && hour.Minute() == 0 && hour.Second() == 0
	return !isZero
}

func ItIsAnHourBeforeTheCurrent(hour string) bool {
	loc, _ := time.LoadLocation("America/Bogota")
	currentHour := time.Now().In(loc)

	arrayHour := strings.Split(hour, ":")
	if len(arrayHour) < 1 {
		return true
	}

	intHour, _ := strconv.Atoi(arrayHour[0])
	intMinute, _ := strconv.Atoi(arrayHour[1])

	if intHour < currentHour.Hour() {
		return true
	}

	if intHour == currentHour.Hour() {
		return intMinute < currentHour.Minute()
	}

	return false
}

func DateFormatYYYYMMDD(date string) bool {
	return regexp.MustCompile(`\d{4}-\d{2}-\d{2}`).MatchString(date)
}

func HourFormatHHMMSS(hour string) bool {
	return regexp.MustCompile(`^([0-1]?[0-9]|[2][0-3]):([0-5][0-9])(:[0-5][0-9])?$`).MatchString(hour)
}

func HourOfPeriodOf15Minutes(hour time.Time) bool {
	minute := hour.Minute()
	if minute == 0 {
		return true
	}

	if minute == 15 {
		return true
	}

	if minute == 30 {
		return true
	}

	if minute == 45 {
		return true
	}

	return false
}

func DateFormatForNewsletter() string {
	fechaActual := time.Now()
	idioma := "es"
	traductor := time.FixedZone(idioma, 0)
	localizador := time.Date(fechaActual.Year(), fechaActual.Month(), fechaActual.Day(), fechaActual.Hour(), fechaActual.Minute(), fechaActual.Second(), fechaActual.Nanosecond(), traductor)
	formato := "Monday, 2 de January de 2006"
	fechaFormateada := localizador.Format(formato)

	fechaFormateada = strings.ReplaceAll(fechaFormateada, "Monday", "Lunes")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "Tuesday", "Martes")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "Wednesday", "Miércoles")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "Thursday", "Jueves")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "Friday", "Viernes")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "Saturday", "Sábado")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "Sunday", "Domingo")

	fechaFormateada = strings.ReplaceAll(fechaFormateada, "January", "Enero")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "February", "Febrero")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "March", "Marzo")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "April", "Abril")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "May", "Mayo")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "June", "Junio")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "July", "Julio")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "August", "Agosto")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "September", "Septiembre")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "Octuber", "Octubre")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "November", "Noviembre")
	fechaFormateada = strings.ReplaceAll(fechaFormateada, "December", "Diciembre")

	return fechaFormateada
}

func ConvertIntToTime(val int) time.Time {
	strValue := strconv.Itoa(val)
	i, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		sentry.CaptureException(err)
		return time.Now()
	}

	return time.Unix(i, 0)
}
