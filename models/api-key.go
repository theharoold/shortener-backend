package models

import "time"

type ApiKey struct {
	Id       int64  `json:"id,omitempty" db:"id"`
	ApiKey   string `json:"api_key" db:"api_key" binding:"required"`
	IsActive bool   `json:"is_active,omitempty" db:"is_active"`
	// TODO add user id eventually
	CreateTime time.Time `json:"create_time,omitempty" db:"create_time"`
	// TODO add expiration date (?)
}
