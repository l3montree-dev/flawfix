package models

import (
	"database/sql"
	"time"
)

type CWE struct {
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`

	CWE string `json:"cwe" gorm:"primaryKey;not null;"`

	Description string `json:"description" gorm:"type:text;"`
}

func (m CWE) TableName() string {
	return "cwes"
}
