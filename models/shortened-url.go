package models

import "time"

type ShortenedURL struct {
	Id            int64  `json:"id,omitempty" db:"id"`
	PlaintextURL  string `json:"plaintext_url" db:"plaintext_url" binding:"required"`
	ShortenedPath string `json:"shortened_path,omitempty" db:"shortened_path"`
	// TODO add auto expiration after some time
	// For now just omit isactive
	IsActive bool `json:"is_valid,omitempty" db:"is_valid"`
	// TODO add user id and api key used to create the shortened url
	CreateTime time.Time `json:"create_time,omitempty" db:"create_time"`
}

