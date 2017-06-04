package models

type Lang struct {
	Zh  string `gorm:"type:varchar(250);not null;index"`
	Eng string `gorm:"type:varchar(250);not null;index"`
}

func (*Lang) TableName() string {
	return kDbPrefix + "lang"
}
