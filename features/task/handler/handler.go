package handler

import (
	"my-task-app/features/task"
	"my-task-app/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.TaskServiceInterface
}

func New(service task.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: service,
	}
}

func (handler *TaskHandler) CreateTask(c echo.Context) error {
	newTask := TaskRequest{}
	errBind := c.Bind(&newTask) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	//mapping dari request ke core
	taskCore := RequestToCore(newTask)
	errInsert := handler.taskService.Create(taskCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("task_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	var taskData = TaskRequest{}
	errBind := c.Bind(&taskData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	taskCore := RequestToCore(taskData)
	errUpdate := handler.taskService.Update(idParam, taskCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data"+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *TaskHandler) DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}

	errDelete := handler.taskService.Delete(id)
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data"+errDelete.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
