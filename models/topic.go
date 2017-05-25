package models

type Topic struct {
	Tid          int32  `gorm:"primary_key"`
	Oid          int32  `gorm:"not null"`
	Status       int8   `gorm:"type:tinyint(1)"`
	Name         string `gorm:"type:varchar(30);not null"`
	EngName      string `gorm:"type:varchar(30);not null"`
	Stpl         string `gorm:"type:varchar(30);not null"`
	Stitle       string `gorm:"type:varchar(50);not null"`
	Skeywords    string `gorm:"type:varchar(255);not null"`
	Sdescription string `gorm:"type:varchar(255);not null"`
	Cover        string `gorm:"type:varchar(250);not null"`
	Content      string `gorm:"type:longtext;not null"`
}

func (*Topic) TableName() string {
	return kDbPrefix + "topic"
}
