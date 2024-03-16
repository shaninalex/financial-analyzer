package typedefs

import (
	"time"
)

type Report struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    string    `gorm:"index" json:"user_id"`
	Ticker    string    `json:"ticker"`
	Link      string    `json:"link"` // link for a generated file
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Status    bool      `json:"status"`
	Issues    []Issue   `json:"issues"`
}

// TODO: handle multiple receivers
// EmailSend      string    `gorm:"type:varchar(255)" json:"email_send"`
// EmailRecipient string    `gorm:"type:varchar(255)" json:"email_recipient"`

type Issue struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Closed      bool      `gorm:"default:false" json:"closed"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	ClosedAt    time.Time `json:"closed_at"`
	ReportID    uint
}
