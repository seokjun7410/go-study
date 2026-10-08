package model

import "time"

type Board struct {
	ID        int       `json:"id,omitempty"`
	Title     string    `json:"id,omitempty"`
	Content   string    `json:"id,omitempty"`
	Author    string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"id,omitempty"`
	UpdatedAt time.Time `json:"id,omitempty"`
	Code      string    `json:"code,omitempty"`
}
