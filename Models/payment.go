package Models

type Payment struct {
	ID      string  `json:"id"`
	GymID   string  `json:"gym_id"`
	OwnerID string  `json:"owner_id"`
	Month   string  `json:"month"`
	Status  string  `json:"status"`
	DueDate string  `json:"due_date"`
	Value   float64 `json:"value"`
}
