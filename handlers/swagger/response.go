package swagger

import "github.com/averrows/apis/core/models"

type SwaggerHabitResponse struct {
	Data    []models.Habit `json:"data"`
	Message string         `json:"message"`
}

// swagger:response habitResponse
type SwaggerHabitResponseWrapper struct {
	Body SwaggerHabitResponse
}
