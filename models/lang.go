package models

type Lang struct {
	Id  int32
	Zh  string `gorm:"type:varchar(250);not null;index"`
	Zht string `gorm:"type:varchar(250);not null;index"`
	Eng string `gorm:"type:varchar(250);not null;index"`
}

func (*Lang) TableName() string {
	return kDbPrefix + "lang"
}
