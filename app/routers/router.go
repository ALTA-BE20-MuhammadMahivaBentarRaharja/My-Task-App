package routers

import (
	"my-task-app/app/middlewares"
	_projectData "my-task-app/features/project/data"
	_projectHandler "my-task-app/features/project/handler"
	_projectService "my-task-app/features/project/service"
	_taskData "my-task-app/features/task/data"
	_taskHandler "my-task-app/features/task/handler"
	_taskService "my-task-app/features/task/service"
	_userData "my-task-app/features/user/data"
	_userHandler "my-task-app/features/user/handler"
	_userService "my-task-app/features/user/service"
	"my-task-app/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {

	hashService := encrypts.NewHashService()
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService)
	userHandlerAPI := _userHandler.New(userService)

	projectData := _projectData.New(db)
	projectService := _projectService.New(projectData, taskService)
	projectHandlerAPI := _projectHandler.New(projectService)

	taskData := _taskData.New(db)
	taskService := _taskService.New(taskData, projectService)
	taskHandlerAPI := _taskHandler.New(taskService)
	// define routes/ endpoint
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUsersById, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
	e.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
	e.GET("/projects", projectHandlerAPI.GetAllProjects, middlewares.JWTMiddleware())
	e.GET("/projects/:project_id", projectHandlerAPI.GetProjectsById, middlewares.JWTMiddleware())
	e.PUT("/projects/:project_id", projectHandlerAPI.UpdateProject, middlewares.JWTMiddleware())
	e.DELETE("/projects/:project_id", projectHandlerAPI.DeleteProject, middlewares.JWTMiddleware())
	e.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
	e.PUT("/tasks/:task_id", taskHandlerAPI.UpdateTask, middlewares.JWTMiddleware())
	e.DELETE("/tasks/:task_id", taskHandlerAPI.DeleteTask, middlewares.JWTMiddleware())

}
