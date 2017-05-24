package models

import (
	"time"
	"fmt"
	"github.com/Junbong/go-api-test/app/libs"
)

type InsightsFunnelDailies struct {
	InsightsBase
	CreatedAt   time.Time   `gorm:"column:created_at"   json:"created_at"`
	UpdatedAt   time.Time   `gorm:"column:updated_at"   json:"updated_at"`
}

func (i *InsightsFunnelDailies) String() string {
	return fmt.Sprintf(
		"insights_funnel_dailies{shop_id=%s, date=%s, profile=%s, " +
			"out=%s, visitors=%s, guests=%s, bounce=%s, " +
			"capture_rate=%s, guest_rate=%s, bounce_rate=%s, " +
			"dwell_time_mean=%s, dwell_time_median=%s, " +
			"revisit_count=%s, revisit_period=%s, " +
			"created_at=%s, updated_at=%s}",
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
		strutil.NullFloat64ToS(i.RevisitPeriod),
		i.CreatedAt.Format(time.RFC3339Nano),
		i.UpdatedAt.Format(time.RFC3339Nano))
}
