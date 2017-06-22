package models

type Download struct {
	Id        int32
	ArticleId int32  `gorm:"index"`
	Url       string `gorm:"type:varchar(250)"`
	Type      int8   `gorm:"type:tinyint"`
	Status    int8   `gorm:"type:tinyint(1)"`
}

const (
	KdlTypeDirect        = 0
	KdlTypeShortenPrefix = 1
)

func (*Download) TableName() string {
	return kDbPrefix + "download"
}
