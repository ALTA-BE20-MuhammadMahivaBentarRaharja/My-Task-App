package handler

import (
	"my-task-app/app/middlewares"
	"my-task-app/features/project"
	"my-task-app/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService project.ProjectServiceInterface
}

func New(service project.ProjectServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
	}
}

func (handler *ProjectHandler) GetAllProjects(c echo.Context) error {
	//mengambil informasi id user yang dikirim pada token payload
	userIdLogin := middlewares.ExtractTokenUserId(c)

	// panggil func di service layer
	results, errSelect := handler.projectService.GetAll(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}
	// proses mapping dari core ke response
	projectsResult := CoreToResponseList(results)

	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", projectsResult))
}

func (handler *ProjectHandler) GetProjectsById(c echo.Context) error {
	//mengambil informasi id user yang dikirim pada token payload
	userIdLogin := middlewares.ExtractTokenUserId(c)

	id, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}

	result, errSelect := handler.projectService.GetById(id, userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	// proses mapping dari core ke response
	var projectResult = CoreToResponse(result)

	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", projectResult))
}

func (handler *ProjectHandler) CreateProject(c echo.Context) error {
	//mengambil informasi id user yang dikirim pada token payload
	userIdLogin := middlewares.ExtractTokenUserId(c)

	newProject := ProjectRequest{}
	errBind := c.Bind(&newProject) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	//mapping dari request ke core
	projectCore := RequestToCore(uint(userIdLogin), newProject)
	errInsert := handler.projectService.Create(projectCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *ProjectHandler) UpdateProject(c echo.Context) error {
	//mengambil informasi id user yang dikirim pada token payload
	userIdLogin := middlewares.ExtractTokenUserId(c)

	id := c.Param("project_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	var projectData = ProjectRequestUpdate{}
	errBind := c.Bind(&projectData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	projectCore := RequestToCoreUpdate(projectData)
	errUpdate := handler.projectService.Update(userIdLogin, idParam, projectCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *ProjectHandler) DeleteProject(c echo.Context) error {
	//mengambil informasi id user yang dikirim pada token payload
	userIdLogin := middlewares.ExtractTokenUserId(c)

	id, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}

	errDelete := handler.projectService.Delete(id, userIdLogin)
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data "+errDelete.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
