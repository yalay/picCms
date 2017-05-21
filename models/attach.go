package models

import ()

type Attach struct {
	Id         int32
	ArticleId  int32  `gorm:"index"`
	Uid        int32  `gorm:"not null"`
	Name       string `gorm:"type:varchar(100);not null"`
	Remark     string `gorm:"type:text;not null"`
	Size       int32
	File       string `gorm:"type:varchar(250);not null"`
	Ext        string `gorm:"type:varchar(10);not null"`
	Status     int8   `gorm:"type:tinyint(1)"`
	Type       int8   `gorm:"type:tinyint(1)"`
	TrlCount   int8   `gorm:"type:tinyint(2)"`
	UploadTime int32
}

func (*Attach) TableName() string {
	return kDbPrefix + "attach"
}
