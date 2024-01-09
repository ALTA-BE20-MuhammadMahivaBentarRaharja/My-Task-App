package main

import (
	"my-task-app/app/configs"
	"my-task-app/app/databases"
	"my-task-app/app/routers"
	_projectData "my-task-app/features/project/data"
	_taskData "my-task-app/features/task/data"
	_userData "my-task-app/features/user/data"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	dbSql := databases.InitDBMysql(cfg)
	dbSql.AutoMigrate(&_userData.User{}, &_projectData.Project{}, &_taskData.Task{})

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	// e.Use(middleware.Logger())

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	routers.InitRouter(dbSql, e)
	//start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
