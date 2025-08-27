package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(200)" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	Category    string    `gorm:"type:varchar(100)" json:"category"`
	CreatedDate time.Time `gorm:"column:created_date" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100)" json:"status"`
}

// BeforeCreate sets timestamps before inserting a new record
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	p.CreatedDate = now
	p.UpdatedDate = now
	return
}

// BeforeUpdate sets UpdatedDate before updating an existing record
func (p *Post) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedDate = time.Now()
	return
}
