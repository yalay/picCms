package models

type Config struct {
	Name  string `gorm:"type:varchar(40);primary_key"`
	Value string `gorm:"type:text;not null"`
	Des   string `gorm:"type:varchar(50);not null"`
}

func (*Config) TableName() string {
	return kDbPrefix + "config"
}
