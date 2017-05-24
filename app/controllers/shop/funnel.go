package controllers

import (
	"github.com/revel/revel"
	"regexp"
	"github.com/Junbong/go-api-test/app"
	"github.com/Junbong/go-api-test/app/models"
	"errors"
	"strings"
	"strconv"
	"fmt"
)

type ShopFunnel struct {
	*revel.Controller
}

var (
	NUMBERS = regexp.MustCompile(`^[0-9,]+$`)
)


func (c ShopFunnel) Index(shop_query string, from, to, profile string) revel.Result {
	c.Validation.Match(from, regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}"))
	c.Validation.Match(to, regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}"))
	
	if profile == "" {
		profile = "all"
	}
	profiles := strings.Split(profile, ",")
	
	shopIds, err := parseToShopIds(shop_query)
	if err != nil {
		return c.RenderError(err)
	}
	
	var funnels []models.InsightsFunnelDailies
	app.DB.
		Model(&models.InsightsFunnelDailies{}).
		Where("shop_id IN (?)", shopIds).
		Where("date BETWEEN ? AND ?", from, to).
		Where("profile IN (?)", profiles).
		Scan(&funnels)
	
	result := make(map[string]interface{})
	result["shop_id"] = shopIds
	result["from"] = from
	result["to"] = to
	result["funnels"] = funnels
	return c.RenderJSON(result)
}


func (c ShopFunnel) Average(shop_query string, from, to, profile string) revel.Result {
	c.Validation.Match(from, regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}"))
	c.Validation.Match(to, regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}"))
	
	if profile == "" {
		profile = "all"
	}
	profiles := strings.Split(profile, ",")
	
	shopIds, err := parseToShopIds(shop_query)
	if err != nil {
		return c.RenderError(err)
	}
	
	var funnels []models.InsightsBase
	app.DB.
		Table("insights_funnel_dailies").
		Where("shop_id IN (?)", shopIds).
		Where("date BETWEEN ? AND ?", from, to).
		Where("profile IN (?)", profiles).
		Select("shop_id, profile, date, " +
			"AVG(out) AS out, " +
			"AVG(visitors) AS visitors, " +
			"AVG(guests) AS guests, " +
			"AVG(bounce) AS bounce, " +
			"AVG(capture_rate) AS capture_rate, " +
			"AVG(guest_rate) AS guest_rate, " +
			"AVG(bounce_rate) AS bounce_rate, " +
			"AVG(dwell_time_mean) AS dwell_time_mean, " +
			"AVG(dwell_time_median) AS dwell_time_median, " +
			"AVG(revisit_count) AS revisit_count, " +
			"AVG(revisit_period) AS revisit_period").
		Group("shop_id, profile, date").
		Scan(&funnels)
	
	result := make(map[string]interface{})
	result["shop_ids"] = shopIds
	result["from"] = from
	result["to"] = to
	result["funnels"] = funnels
	return c.RenderJSON(result)
}


func parseToShopIds(shopQuery string) (shopIds []uint, err error) {
	switch {
	case len(shopQuery) == 0:
		err = errors.New("Invalid shop query")
		break;
	
	case NUMBERS.MatchString(shopQuery):
		spl := strings.Split(shopQuery, ",")
		var u64 uint64
		for _, tkn := range spl {
			u64, err = strconv.ParseUint(tkn, 10, 32)
			shopIds = append(shopIds, uint(u64))
		}
		break;
		
	default:
		var taggings []models.Taggings
		app.DB.
			Table("taggings").
			Select("taggings.taggable_id").
			Joins("INNER JOIN tags ON taggings.tag_id = tags.id").
			Where("taggings.taggable_type = 'Shop'").
			Where("tags.name = ?", shopQuery).
			Scan(&taggings)
		
		if len(taggings) > 0 {
			for _, t := range taggings {
				shopIds = append(shopIds, t.TaggableId)
			}
		} else {
			err = errors.New(fmt.Sprintf("No shops found with specified tag '%s'", shopQuery))
		}
	}
	return
}
