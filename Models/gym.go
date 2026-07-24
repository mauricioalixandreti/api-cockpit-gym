package Models

import "time"

type Gym struct {
	ID        string    `json:"id"`
	Nome      string    `json:"nome"`
	Endereco  string    `json:"endereco"`
	Telefone  string    `json:"telefone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
