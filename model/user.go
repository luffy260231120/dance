package model

type UserInfo struct {
	ID         int    `gorm:"column:id" json:"id"`
	UserId     string `gorm:"column:user_id" json:"user_id"`
	Token      string `gorm:"column:token" json:"token"`
	Name       string `gorm:"column:name" json:"name"`
	Sex        int    `gorm:"column:sex" json:"sex"` // 1-女 2-男
	Bk         string `gorm:"column:bk" json:"bk"`
	CreateTime string `gorm:"column:create_time" json:"create_time"`
	UpdateTime string `gorm:"column:update_time" json:"update_time"`
}
