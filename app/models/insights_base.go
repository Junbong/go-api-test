package models

import (
	"time"
	"fmt"
	"errors"
	"github.com/Junbong/go-api-test/app/libs"
	"gopkg.in/guregu/null.v3"
)

type InsightsBase struct {
	ShopId          uint            `gorm:"primary_key,column:shop_id"  json:"shop_id"`
	Date            time.Time       `gorm:"primary_key,column:date"     json:"date"`
	Profile         string          `gorm:"primary_key,column:profile"  json:"profile"`
	Out             null.Float      `gorm:"column:out"                  json:"out"`
	Visitors        null.Float      `gorm:"column:visitors"             json:"visitors"`
	Guests          null.Float      `gorm:"column:guests"               json:"guests"`
	Bounce          null.Float      `gorm:"column:bounce"               json:"bounce"`
	CaptureRate     null.Float      `gorm:"column:capture_rate"         json:"capture_rate"`
	GuestRate       null.Float      `gorm:"column:guest_rate"           json:"guest_rate"`
	BounceRate      null.Float      `gorm:"column:bounce_rate"          json:"bounce_rate"`
	DwellTimeMean   null.Float      `gorm:"column:dwell_time_mean"      json:"dwell_time_mean"`
	DwellTimeMedian null.Float      `gorm:"column:dwell_time_median"    json:"dwell_time_median"`
	RevisitCount    null.Float      `gorm:"column:revisit_count"        json:"revisit_count,omitempty"`
	RevisitPeriod   null.Float      `gorm:"column:revisit_period"       json:"revisit_period,omitempty"`
}

func (i *InsightsBase) String() string {
	return fmt.Sprintf(
		"insights_funnel_dailies{shop_id=%s, date=%s, profile=%s, " +
			"out=%s, visitors=%s, guests=%s, bounce=%s, " +
			"capture_rate=%s, guest_rate=%s, bounce_rate=%s, " +
			"dwell_time_mean=%s, dwell_time_median=%s, " +
			"revisit_count=%s, revisit_period=%s}",
		strutil.NullInt64ToS(i.ShopId),
		i.Date.Format("2006-01-02"), i.Profile,
		strutil.NullFloat64ToS(i.Out),
		strutil.NullFloat64ToS(i.Visitors),
		strutil.NullFloat64ToS(i.Guests),
		strutil.NullFloat64ToS(i.Bounce),
		strutil.NullFloat64ToS(i.CaptureRate),
		strutil.NullFloat64ToS(i.GuestRate),
		strutil.NullFloat64ToS(i.BounceRate),
		strutil.NullFloat64ToS(i.DwellTimeMean),
		strutil.NullFloat64ToS(i.DwellTimeMedian),
		strutil.NullFloat64ToS(i.RevisitCount),
		strutil.NullFloat64ToS(i.RevisitPeriod))
}

func (i *InsightsBase) AfterFind() (err error) {
	if i.ShopId == 0 {
		err = errors.New("Not found")
	}
	return
}
