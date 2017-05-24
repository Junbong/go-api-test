package models

import "fmt"

type Tags struct {
	Id              uint    `gorm:"primary_key, AUTO_INCREMENT"         json:"id"`
	Name            string  `gorm:"column:name"                         json:"name"`
	TaggingsCount   uint    `gorm:"default:0, column:taggings_count"    json:"taggings_count"`
}

func (t *Tags) String() string {
	return fmt.Sprintf(
		"tags{id=%d, name=%s, taggings_count=%d}",
		t.Id, t.Name, t.TaggingsCount)
}
