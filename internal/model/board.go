package model

import "time"

type Board struct {
	ID        int       `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Author    string    `json:"author,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Code      string    `json:"code,omitempty" grom:"-"`
}

func (Board) TableName() string {
	return "board"
}
