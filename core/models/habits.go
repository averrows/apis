package models

import "gorm.io/gorm"

// swagger:model Habit
type Habit struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	TimeType    string     `json:"time_type"`
	TimeTable   string     `json:"time_table"`
	HabitLogs   []HabitLog `json:"habit_logs"`
}

// swagger:model HabitLog
type HabitLog struct {
	gorm.Model
	HabitID     uint
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}
