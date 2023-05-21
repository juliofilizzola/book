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

type PublicationReturn struct {
	ID          uint64 `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	Description string `json:"description"`
	AuthId      string `json:"auth_id,omitempty"`
	AuthName    string `json:"auth_name,omitempty"`
	AuthNick    string `json:"auth_nick,omitempty"`
	AuthEmail   string `json:"auth_email,omitempty"`
	Likes       string `json:"likes"`
	CreatedAt   string `json:"created_at,omitempty"`
}
