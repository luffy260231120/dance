package model

type CourseInfo struct {
	ID        int    `gorm:"column:id" json:"id"`
	Type      int    `gorm:"column:type" json:"type" binding:"required"`
	Title     string `gorm:"column:title" json:"title" binding:"required"`
	TeacherId int    `gorm:"column:teacher_id" json:"teacher_id" binding:"required"`
	MaxNumber int    `gorm:"column:max_number" json:"max_number" binding:"required"`
	Bk        string `gorm:"column:bk" json:"bk" binding:"required"`
	OpenTime  string `gorm:"column:open_time" json:"open_time" binding:"required,time"`
	StartTime string `gorm:"column:start_time" json:"start_time" binding:"required,time"`
	EndTime   string `gorm:"column:end_time" json:"end_time" binding:"required,time"`
	Date      string `gorm:"column:date" json:"date"`
	Class     string `gorm:"column:class" json:"class"`

	CreateTime string `gorm:"column:create_time" json:"-"`
	UpdateTime string `gorm:"column:update_time" json:"-"`
}
