package models

type Cate struct {
	Cid          int32 `gorm:"primary_key"`
	Pid          int32
	Oid          int32  `gorm:"not null"`
	ViewType     int8   `gorm:"type:tinyint(1)"`
	Status       int8   `gorm:"type:tinyint(1)"`
	Index        int8   `gorm:"type:tinyint(1)"`
	Name         string `gorm:"type:varchar(30);not null"`
	EngName      string `gorm:"type:varchar(30);not null"`
	Ctpl         string `gorm:"type:varchar(30);not null"`
	Ctitle       string `gorm:"type:varchar(50);not null"`
	Ckeywords    string `gorm:"type:varchar(255);not null"`
	Cdescription string `gorm:"type:varchar(255);not null"`
}

func (*Cate) TableName() string {
	return kDbPrefix + "cate"
}
