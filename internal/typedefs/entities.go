package typedefs

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	UserId    string `gorm:"index" json:"user_id"`
	RequestId string `json:"request_id"`
	Ticker    string `json:"ticker"`
	Link      string `json:"link"` // link for a generated file
	Status    bool   `json:"status"`
	Issues    []Issue
	Data      []ReportData
}

type ReportData struct {
	gorm.Model
	Type     string `json:"type"`
	Data     string `json:"data"`
	ReportID uint
}

// Closing issue mean that we "delete" Issue in db (add daleted_at in row)
// Docs: https://gorm.io/docs/models.html#gorm-Model
type Issue struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	ReportID    uint
}

// old, comments... Do not understand it completely
// TODO: handle multiple receivers
// EmailSend      string    `gorm:"type:varchar(255)" json:"email_send"`
// EmailRecipient string    `gorm:"type:varchar(255)" json:"email_recipient"`
