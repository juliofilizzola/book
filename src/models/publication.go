package models

type Publication struct {
	ID          uint64 `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	Description string `json:"description"`
	AuthId      string `json:"auth_id,omitempty"`
	Likes       string `json:"likes"`
	CreatedAt   string `json:"created_at,omitempty"`
}
