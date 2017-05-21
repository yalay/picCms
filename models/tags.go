package models

import ()

type Tags struct {
	Tag       string `gorm:"type:varchar(30);not null;index"`
	Title     string `gorm:"type:varchar(250);not null"`
	ArticleId int32  `gorm:"index"`
}

func (*Tags) TableName() string {
	return kDbPrefix + "tags"
}
