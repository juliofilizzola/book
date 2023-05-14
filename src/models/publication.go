package models

type Publication struct {
	ID        uint64 `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,,omitempty"`
	AuthId    string `json:"auth_id,omitempty"`
	AuthNick  string `json:"auth_nick,omitempty"`
	Likes     string `json:"likes"`
	CreatedAt string `json:"created_at,omitempty"`
}
