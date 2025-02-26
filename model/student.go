package model

import (
	"gorm.io/gorm"
)

// Student โครงสร้างข้อมูลนักเรียน (GORM Model)
type StudentEntity struct {
	Id        uint           `gorm:"primaryKey;autoIncrement"`
	StudentId string         `gorm:"type:varchar(20);unique;not null"`
	Title     string         `gorm:"type:varchar(10);not null"`
	FirstName string         `gorm:"type:varchar(50);not null"`
	LastName  string         `gorm:"type:varchar(50);not null"`
	NickName  string         `gorm:"type:varchar(30)"`
	Gender    string         `gorm:"type:varchar(10);not null"`
	Class     string         `gorm:"type:varchar(10);not null"`
	CreatedAt int64          `gorm:"autoCreateTime:milli"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
