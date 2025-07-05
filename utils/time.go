package utils

import (
	"net/http"
	"time"
)

func StringToLocalTime(s string) (*time.Time, error) {
	tz, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, NewCustomError(http.StatusInternalServerError, err.Error())
	}
	parseTime, err := time.Parse(LayoutDateString, s)
	if err != nil {
		return nil, NewCustomError(http.StatusInternalServerError, err.Error())
	}
	timeInTZ := parseTime.In(tz)
	return &timeInTZ, nil
}

func StringToLocalDateOnly(s string) (*time.Time, error) {
	tz, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, NewCustomError(http.StatusInternalServerError, err.Error())
	}
	parseTime, err := time.Parse(LayoutDateOnlyString, s)
	if err != nil {
		return nil, NewCustomError(http.StatusInternalServerError, err.Error())
	}
	timeInTZ := parseTime.In(tz)
	return &timeInTZ, nil
}

func ToLocalTime(t time.Time) time.Time {
	tz, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return t
	}
	timeInTZ := t.In(tz)
	return timeInTZ
}
