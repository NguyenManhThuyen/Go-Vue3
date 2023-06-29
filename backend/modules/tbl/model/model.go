package model

import (
	"time"
)

type Tbl struct {
	ID         uint       `gorm:"<-:false;column:department_id;primaryKey"`
	NameVN     string     `gorm:"column:department_name_vn;size:100;not null"`
	NameEN     string     `gorm:"column:department_name_en;size:100;not null"`
	NameJP     string     `gorm:"column:department_name_jp;size:100;not null"`
	Shortcut   string     `gorm:"column:department_shortcut;size:5"`
	CreatedAt  time.Time  `gorm:"<-:false;autoCreateTime;column:created_at"`
	CreatedBy  string     `gorm:"column:created_by;size:5"`
	UpdatedAt  time.Time  `gorm:"<-:false;autoUpdateTime:milli;column:updated_at"`
	UpdatedBy  string     `gorm:"column:updated_by;size:5"`
	DeletedAt  *time.Time `gorm:"-:all;column:deleted_at"`
	DeletedBy  string     `gorm:"column:deleted_by;size:5"`
	LogVersion int64      `gorm:"column:log_version;default:0"`
}

func (Tbl) TableName() string {
	return "tbl_department"
}
