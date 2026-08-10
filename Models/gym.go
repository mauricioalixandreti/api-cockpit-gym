package Models

import "time"

type Gym struct {
	ID          string    `json:"id"`
	Nome        string    `json:"nome"`
	Owner_Email string    `json:"owner_email"`
	Endereço    string    `json:"endereço"`
	Telefone    string    `json:"telefone"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
