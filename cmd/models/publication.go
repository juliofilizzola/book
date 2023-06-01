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
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description"`
	Content     string `json:"content,omitempty"`
	Likes       string `json:"likes"`
	CreatedAt   string `json:"created_at,omitempty"`
	AuthNick    string `json:"auth_nick,omitempty"`
	AuthEmail   string `json:"auth_email,omitempty"`
	AuthId      string `json:"auth_id,omitempty"`
}

type PublicationReturn2 struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description"`
	Content     string `json:"content,omitempty"`
	Likes       int    `json:"likes"`
	CreatedAt   string `json:"created_at,omitempty"`
	AuthNick    string `json:"auth_nick,omitempty"`
	AuthEmail   string `json:"auth_email,omitempty"`
	AuthId      string `json:"auth_id,omitempty"`
}

type Like struct {
	Like int `json:"like"`
}
