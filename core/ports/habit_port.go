package ports

import "github.com/averrows/apis/core/models"

type HabitRepository interface {
	GetAllHabits() ([]models.Habit, error)
	GetHabit(id uint) (models.Habit, error)
	SaveHabit(habit models.Habit) (models.Habit, error)
	GetHabitLogs(habitID uint) ([]models.HabitLog, error)
	SaveHabitLog(habitLog models.HabitLog) (models.HabitLog, error)
}

type HabitService interface {
	GetAll() ([]models.Habit, error)
	Get(id uint) (models.Habit, error)
	Save(habit models.Habit) (models.Habit, error)
}

type HabitLogService interface {
	Log(habitLog models.HabitLog) (models.HabitLog, error)
	CheckDateValidity(habitId uint, date string) (bool, error)
}
