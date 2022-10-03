package models

import (
	"errors"
	"html"
	"strings"
)

type Item struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"item_id"`
	ItemCode    string `gorm:"size:255;not null;unique" json:"item_code"`
	Description string `gorm:"size:255;not null;" json:"description"`
	Quantity    uint64 `gorm:"not null" json:"quantity"`
	OrderID     uint64 `gorm:"not null" json:"order_id"`
}

func (p *Item) Prepare() {
	p.ID = 0
	p.ItemCode = html.EscapeString(strings.TrimSpace(p.ItemCode))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
}

func (p *Item) Validate() error {
	if p.ItemCode == "" {
		return errors.New("Required ItemCode")
	}
	if p.Description == "" {
		return errors.New("Required Description")
	}
	if p.Quantity < 1 {
		return errors.New("Required Quantity")
	}
	if p.OrderID < 1 {
		return errors.New("Required Order")
	}
	return nil
}
