//   Product Api:
//    version: 0.1
//    title: Product Api
//   Schemes: http, https
//   Host:
//   BasePath: /api/v1
//      Consumes:
//      - application/json
//   Produces:
//   - application/json
//   swagger:meta
package main

import (
	"net/http"

	"github.com/averrows/apis/core/services/habitsrv"
	habit_http "github.com/averrows/apis/handlers/http"
	mysql_repo "github.com/averrows/apis/repositories/mysql"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	habitRepo, err := mysql_repo.NewHabitMySQLRepository()
	if err != nil {
		panic(err)
	}
	habitEntityService := habitsrv.NewEntityService(habitRepo)
	habitNestedService := habitsrv.NewNestedService(habitRepo)
	habitLogService := habitsrv.NewLogService(habitRepo)

	habitHandler := habit_http.NewHTTPHandler(habitEntityService, habitNestedService, habitLogService)

	router := gin.Default()

	router.GET("/swagger.yaml", gin.WrapH(http.FileServer(http.Dir("./"))))

	opts := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	router.GET("/docs", gin.WrapH(sh))
	router.Use(CORSMiddleware())

	apiV1 := router.Group("/v1")
	// swagger:operation GET /habits Habits GetHabits
	// Get Habits
	//
	// ---
	// parameters:
	// - name: id
	//   in: query
	//   description: id of habit
	//   type: integer
	//   format: int64
	// produces:
	// - application/json
	// responses:
	// 	200:
	//   schema:
	//   	type: array
	//   	items:
	//         "$ref": "#/definitions/Habit"
	apiV1.GET("/habits", habitHandler.Get)

	// swagger:operation POST /habits Habits PostHabit
	// Create Habit
	//
	// ---
	// parameters:
	// - name: habit
	//   in: body
	//   description: habit to create
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Habit"
	// responses:
	// 	200:
	//   schema:
	//   	"$ref": "#/responses/habitResponse"
	apiV1.POST("/habits", habitHandler.Post)
	// swagger:operation POST /habits/log Habits LogHabit
	// Create Habit
	//
	// ---
	// parameters:
	// - name: habit
	//   in: body
	//   description: habit to create
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Habit"
	// responses:
	// 	200:
	//   schema:
	//   	"$ref": "#/responses/habitResponse"
	apiV1.POST("/habits/log", habitHandler.LogHabit)
	apiV1.GET("/habits/log/check-date", habitHandler.CheckIfDateValid)

	router.Run(":8000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
