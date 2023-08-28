package habitsrv

import (
	"time"

	"github.com/averrows/apis/core/models"
	"github.com/averrows/apis/core/ports"
)

type NestedService struct {
	repo ports.HabitRepository
}

func NewNestedService(repo ports.HabitRepository) *NestedService {
	return &NestedService{repo: repo}
}

func (s *NestedService) GetAll() ([]models.Habit, error) {
	habits, err := s.repo.GetAllHabits()
	if err != nil {
		return nil, err
	}

	for i := range habits {
		habits[i].HabitLogs, err = s.repo.GetHabitLogs(habits[i].ID)
		if err != nil {
			return nil, err
		}
	}

	return habits, nil
}

func (s *NestedService) Get(id uint) (models.Habit, error) {
	habit, err := s.repo.GetHabit(id)
	if err != nil {
		return models.Habit{}, err
	}
	habit.HabitLogs, err = s.repo.GetHabitLogs(habit.ID)
	if err != nil {
		return models.Habit{}, err
	}
	return habit, nil
}

func (s *NestedService) Save(habit models.Habit) (models.Habit, error) {
	habit, err := s.repo.SaveHabit(habit)
	if err != nil {
		return models.Habit{}, err
	}
	habitLogs, err := s.repo.GetHabitLogs(habit.ID)
	firstDayInTheWeek := time.Now().AddDate(0, 0, -int(time.Now().Weekday()))
	var habitLogsToSent []models.HabitLog
	for i := 0; i < 7; i++ {
		date := firstDayInTheWeek.AddDate(0, 0, i)
		if isDateValid(habit.TimeType, habit.TimeTable, date.Format("2006-01-02")) {
			for _, habitLog := range habitLogs {
				if habitLog.CreatedAt.Format("2006-01-02") == date.Format("2006-01-02") {
					habitLogsToSent = append(habitLogsToSent, habitLog)
				}
			}
			habitLogs = append(habitLogs, models.HabitLog{
				HabitID: habit.ID,
				IsDone:  false,
			})
		}
	}
	if err != nil {
		return models.Habit{}, err
	}
	habit.HabitLogs = habitLogsToSent
	return habit, nil
}
