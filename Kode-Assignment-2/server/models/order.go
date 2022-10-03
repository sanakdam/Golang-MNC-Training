package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Order struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"order_id"`
	CustomerName string    `gorm:"size:255;not null" json:"customer_name"`
	Items        []Item    `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"ordered_at"`
}

func (o *Order) Prepare() {
	o.ID = 0
	o.CustomerName = html.EscapeString(strings.TrimSpace(o.CustomerName))
	o.OrderedAt = time.Now()
}

func (o *Order) Validate() error {
	if o.CustomerName == "" {
		return errors.New("Required CustomerName")
	}
	if len(o.Items) == 0 {
		return errors.New("Required Items")
	}
	return nil
}
