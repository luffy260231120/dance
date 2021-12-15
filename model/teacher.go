package model

type TeacherInfo struct {
	ID          int    `gorm:"column:id" json:"id"`
	TeacherId   string `gorm:"column:teacher_id" json:"teacher_id"`
	Token       string `gorm:"column:token" json:"token"`
	TeacherName string `gorm:"column:name" json:"name"`
	Sex         int    `gorm:"column:sex" json:"sex"`       // 1-女 2-男
	Type        int    `gorm:"column:type" json:"type"`     // 1 兼职 2 全职
	Status      int    `gorm:"column:status" json:"status"` // 1 在职 2 离职
	Bk          string `gorm:"column:bk" json:"bk"`
	Avatar      string `gorm:"column:avatar" json:"avatar"`
	CreateTime  string `gorm:"column:create_time" json:"create_time"`
	UpdateTime  string `gorm:"column:update_time" json:"update_time"`
}
