package models

import (
	"gopkg.in/guregu/null.v3"
	"time"
	"fmt"
)

type Taggings struct {
	Id              uint        `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	TagId           uint        `gorm:"column:tag_id"               json:"tag_id"`
	TaggableId      uint        `gorm:"column:taggable_id"          json:"taggable_id"`
	TaggableType    string      `gorm:"column:taggable_type"        json:"taggable_type"`
	TaggerId        uint        `gorm:"column:tagger_id"            json:"tagger_id"`
	TaggerType      string      `gorm:"column:tagger_type"          json:"tagger_type"`
	Context         null.String `gorm:"size:255, column:context"    json:"context"`
	CreatedAt       time.Time   `gorm:"column:created_at"           json:"created_at"`
}

func (i *Taggings) String() string {
	return fmt.Sprintf(
		"taggings{id=%d, tag_id=%d, taggable_id=%d, taggable_type=%s, " +
			"tagger_id=%d, tagger_type=%s, context=%s, created_at=%s}",
		i.Id, i.TagId, i.TaggableId, i.TaggableType,
		i.TaggerId, i.TaggerType, i.Context, i.CreatedAt.Format(time.RFC3339Nano))
}
