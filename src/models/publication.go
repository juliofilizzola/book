package models

import "time"

type Publication struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	Description string    `json:"description"`
	AuthId      string    `json:"auth_id,omitempty"`
	Likes       int       `json:"likes"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type PublicationReturn struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	Description string    `json:"description"`
	AuthNick    string    `json:"auth_nick,omitempty"`
	AuthId      string    `json:"auth_id,omitempty"`
	AuthEmail   string    `json:"auth_email,omitempty"`
	Likes       string    `json:"likes"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type Like struct {
	Like int64 `json:"like"`
}
