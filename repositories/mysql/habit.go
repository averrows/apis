package mysql_repo

import (
	"os"

	"github.com/averrows/apis/core/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type HabitMySQLRepository struct {
	db *gorm.DB
}

func NewHabitMySQLRepository() (*HabitMySQLRepository, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})

	db.AutoMigrate(&models.Habit{})
	db.AutoMigrate(&models.HabitLog{})

	if err != nil {
		return nil, err
	}
	return &HabitMySQLRepository{db: db}, nil
}

func (r *HabitMySQLRepository) GetAllHabits() ([]models.Habit, error) {
	var habits []models.Habit
	r.db.Find(&habits)
	return habits, nil
}

func (r *HabitMySQLRepository) GetHabit(id uint) (models.Habit, error) {
	var habit models.Habit
	r.db.First(&habit, id)
	return habit, nil
}

func (r *HabitMySQLRepository) SaveHabit(habit models.Habit) (models.Habit, error) {
	r.db.Save(&habit)
	return habit, nil
}

func (r *HabitMySQLRepository) GetHabitLogs(habitID uint) ([]models.HabitLog, error) {
	var habitLogs []models.HabitLog
	r.db.Where("habit_id = ?", habitID).Find(&habitLogs)
	return habitLogs, nil
}

func (r *HabitMySQLRepository) SaveHabitLog(habitLog models.HabitLog) (models.HabitLog, error) {
	r.db.Save(&habitLog)
	return habitLog, nil
}
