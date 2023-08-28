package habit_http

import (
	"strconv"
	"time"

	"github.com/averrows/apis/core/models"
	"github.com/averrows/apis/core/ports"
	"github.com/averrows/apis/handlers"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	habitEntityService ports.HabitService
	habitNestedService ports.HabitService
	habitLogService    ports.HabitLogService
}

func NewHTTPHandler(
	habitEntityService ports.HabitService,
	habitNestedService ports.HabitService,
	habitLogService ports.HabitLogService) *HTTPHandler {
	return &HTTPHandler{
		habitEntityService: habitEntityService,
		habitNestedService: habitNestedService,
		habitLogService:    habitLogService,
	}
}

func (h *HTTPHandler) Get(c *gin.Context) {
	var id = c.Request.URL.Query().Get("id")
	if id != "" {
		parsedID, err := strconv.Atoi(id)
		if err != nil {
			c.IndentedJSON(400, handlers.Response{Message: "Invalid ID"})
			return
		}
		habit, err := h.habitNestedService.Get(uint(parsedID))
		if err != nil {
			c.IndentedJSON(500, handlers.Response{Message: err.Error()})
			return
		}
		c.IndentedJSON(200, handlers.Response{Data: habit})
		return
	}

	habits, err := h.habitEntityService.GetAll()
	if err != nil {
		c.IndentedJSON(500, handlers.Response{Message: err.Error()})
		return
	}
	c.IndentedJSON(200, habits)
}

func (h *HTTPHandler) Post(c *gin.Context) {
	var habit models.Habit
	if err := c.BindJSON(&habit); err != nil {
		c.IndentedJSON(400, handlers.Response{Message: "Invalid JSON"})
		return
	}
	habit, err := h.habitEntityService.Save(habit)
	if err != nil {
		c.IndentedJSON(500, handlers.Response{Message: err.Error()})
		return
	}
	c.IndentedJSON(200, handlers.Response{Data: habit})
}

func (h *HTTPHandler) LogHabit(c *gin.Context) {
	var req handlers.HabitLogRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(400, handlers.Response{Message: "Invalid JSON"})
		return
	}
	if req.HabitID == 0 {
		c.IndentedJSON(400, handlers.Response{Message: "Invalid Habit ID"})
		return
	}
	habitLog, err := h.habitLogService.Log(models.HabitLog{
		HabitID:     req.HabitID,
		Description: req.Description,
		IsDone:      true,
	})
	if err != nil {
		c.IndentedJSON(500, handlers.Response{Message: err.Error()})
		return
	}
	c.IndentedJSON(200, handlers.Response{Data: habitLog})
}

func (h *HTTPHandler) CheckIfDateValid(c *gin.Context) {
	id := c.Request.URL.Query().Get("habitId")
	if id == "" {
		c.IndentedJSON(400, handlers.Response{Message: "Invalid Habit ID"})
		return
	}
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(400, handlers.Response{Message: "Invalid Habit ID"})
		return
	}
	now := time.Now()
	isValid, err := h.habitLogService.CheckDateValidity(uint(parsedId), now.Format("2006-01-02"))
	if err != nil {
		c.IndentedJSON(500, handlers.Response{Message: err.Error()})
		return
	}
	c.IndentedJSON(200, handlers.Response{Data: isValid})
}
