package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) PrepareData(edit bool) error {
	if err := u.validation(edit); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u *User) validation(edit bool) error {
	if u.Name == "" {
		return errors.New("name has required")
	}

	if u.Nick == "" {
		return errors.New("nick has required")
	}

	if u.Email == "" {
		return errors.New("email has required")
	}

	if u.Password == "" && !edit {
		return errors.New("pasaword has required")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)
}
