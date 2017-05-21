package models

type Article struct {
	Id         int32
	Cid        int32  `gorm:"type:smallint(5);index"`
	Title      string `gorm:"type:varchar(250);unique"`
	Cover      string `gorm:"type:varchar(250)"`
	Author     string `gorm:"type:varchar(50)"`
	Comeurl    string `gorm:"type:varchar(50)"`
	Remark     string `gorm:"type:text;not null"`
	ShortTitle string `gorm:"type:text"`
	Keywords   string `gorm:"type:text"`
	Content    string
	Hits       int32 `gorm:"type:mediumint(8);index"`
	Star       int8  `gorm:"type:tinyint(1);index"`
	Status     int8  `gorm:"type:tinyint(1)"`
	Up         int32 `gorm:"type:mediumint(8);index"`
	Down       int32 `gorm:"type:mediumint(8);index"`
	Addtime    int64 `gorm:"index"`
}

func (*Article) TableName() string {
	return kDbPrefix + "article"
}
