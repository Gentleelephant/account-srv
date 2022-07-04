package model

type Account struct {
	Id        int64  `gorm:"primarykey"`
	Mobile    string `gorm:"type:varchar(11);unique;not null;index:idx_mobile"`
	Password  string `gorm:"type:varchar(256);not null"`
	Nickname  string `gorm:"type:varchar(64);not null"`
	Gender    string `gorm:"type:varchar(6);default:male"`
	Salt      string `gorm:"type:varchar(256);not null"`
	Role      string `gorm:"type:int;default:1;comment'1-普通用户 2-管理员'"`
	CreatedAt int64  `gorm:"type:bigint;autoCreateTime"`
	UpdatedAt int64  `gorm:"type:bigint;autoUpdateTime"`
	Status    uint   `gorm:"type:tinyint;default:1;comment'1-正常 2-禁用'"`
}
