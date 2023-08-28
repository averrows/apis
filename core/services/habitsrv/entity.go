package habitsrv

import (
	"github.com/averrows/apis/core/models"
	"github.com/averrows/apis/core/ports"
)

type EntityService struct {
	repo ports.HabitRepository
}

func NewEntityService(repo ports.HabitRepository) *EntityService {
	return &EntityService{repo: repo}
}

func (s *EntityService) GetAll() ([]models.Habit, error) {
	return s.repo.GetAllHabits()
}

func (s *EntityService) Get(id uint) (models.Habit, error) {
	return s.repo.GetHabit(id)
}

func (s *EntityService) Save(habit models.Habit) (models.Habit, error) {
	return s.repo.SaveHabit(habit)
}
