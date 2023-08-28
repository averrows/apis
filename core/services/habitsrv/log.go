package habitsrv

import (
	"fmt"
	"strings"
	"time"

	"github.com/averrows/apis/core/models"
	"github.com/averrows/apis/core/ports"
)

type LogService struct {
	repo ports.HabitRepository
}

func NewLogService(repo ports.HabitRepository) *LogService {
	return &LogService{repo: repo}
}

func (s *LogService) Log(habitLog models.HabitLog) (models.HabitLog, error) {
	_, err := s.repo.GetHabit(habitLog.HabitID)
	if err != nil {
		return models.HabitLog{}, err
	}
	return s.repo.SaveHabitLog(habitLog)
}

func (S *LogService) CheckDateValidity(habitId uint, date string) (bool, error) {
	habit, err := S.repo.GetHabit(habitId)
	if err != nil {
		return false, err
	}
	return isDateValid(habit.TimeType, habit.TimeTable, date), nil
}

func isDateValid(timeType string, timeTable string, date string) bool {
	// parse the time table
	mask := strings.Split(timeTable, ",")
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Here")
		return false
	}
	if timeType == "day" {
		for i := len(mask); i < 7; i++ {
			mask = append(mask, "0")
		}
		day := parsedDate.Weekday()
		return mask[int(day)] == "1"
	} else if timeType == "month" {
		for i := len(mask); i < 12; i++ {
			mask = append(mask, "0")
		}
		month := parsedDate.Month()
		return mask[int(month)] == "1"
	} else {
		return false
	}
}

// if type is daily
// 	retrieve logs for the week
// 	build the dummy logs from the start
// 		if the log is on the same date in the logs
// 			append the log
// 		else
// 			append the dummy uninserted
// 	return the logs

// it will ignore the entry outside the day
// create the checker

// if type is monthly
// 	retrieve logs for the month
// 	build the dummy logs from the start
