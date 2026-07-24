package Models

import "time"

type User struct {
	ID        string    `json:"id"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Plano     string    `json:"plano"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
