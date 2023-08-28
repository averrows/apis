package handlers

// swagger:response Response
type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type HabitLogRequest struct {
	HabitID     uint   `json:"habit_id"`
	Description string `json:"description"`
}
