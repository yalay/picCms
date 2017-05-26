package models

type Ad struct {
	Id      int32  `gorm:"primary_key"`
	Title   string `gorm:"type:varchar(50);not null"`
	Des     string `gorm:"type:varchar(50);not null"`
	Content string `gorm:"type:text;not null"`
	Status  int8   `gorm:"type:tinyint(1)"`
}

func (*Ad) TableName() string {
	return kDbPrefix + "adsense"
}
