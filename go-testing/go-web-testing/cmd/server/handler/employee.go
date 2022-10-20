package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/pkg/web"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	employeeService employee.Service
}

type employeeDTO struct {
	CardNumberID string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WarehouseID  int    `json:"warehouse_id"`
}

type employeeToUpdateDTO struct {
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}

var employeeList []employeeDTO

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		employeeService: e,
	}
}

// Get employee by ID
// @Summary Show the required employee
// @Tags    Employees
// @Accept  json
// @Produce json
// @Param   id    path     int          true "Employee ID"
// @Param   token header   string       true "Token"
// @Success 200   {object} web.response "Employee"
// @Failure 400   {object} web.response "Malformed ID"
// @Failure 401   {object} web.response "Unauthorized + invalid token"
// @Failure 404   {object} web.response "Employee not found"
// @Router  /employees/{id} [GET]
func (e *Employee) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidID, id := parseID(c)
		if !isValidID {
			return
		}

		employee, err := e.employeeService.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "%s", err)
			return
		}

		web.Success(c, http.StatusOK, convertModelToDTO(employee))
	}
}

// Get all employees
// @Summary Show the list of employees in db
// @Tags    Employees
// @Produce json
// @Param   token header   string       true "Token"
// @Success 200   {object} web.response "List of employees"
// @Failure 401   {object} web.response "Unauthorized + invalid token"
// @Failure 404   {object} web.response "No employees found"
// @Failure 500   {object} web.response "Failure executing listing service"
// @Router  /employees [GET]
func (e *Employee) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		employees, err := e.employeeService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "%s", err)
			return
		}

		if len(employees) <= 0 {
			web.Error(c, http.StatusNotFound, "no employees were found in the database")
			return
		}

		employeeList = nil
		for _, employee := range employees {
			employeeList = append(employeeList, convertModelToDTO(employee))
		}

		web.Success(c, http.StatusOK, employeeList)
	}
}

// Create a new employee
// @Summary Save a new employee in db
// @Tags    Employees
// @Accept  json
// @Produce json
// @Param   token   header   string       true "Token"
// @Param   product body     employeeDTO  true "Employee to save"
// @Success 201     {object} web.response "Employee created"
// @Failure 401     {object} web.response "Unauthorized + invalid token"
// @Failure 404     {object} web.response "The employee was saved, but it couldn't be retrieved by ID"
// @Failure 409     {object} web.response "The card number ID provided is already registered"
// @Failure 422     {object} web.response "Gin validator error"
// @Failure 500     {object} web.response "Failure executing creation service"
// @Router  /employees [POST]
func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var employeeReceived employeeDTO
		if err := c.ShouldBindJSON(&employeeReceived); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", err)
			return
		}

		employeeToCreate := domain.Employee{
			CardNumberID: employeeReceived.CardNumberID,
			FirstName:    employeeReceived.FirstName,
			LastName:     employeeReceived.LastName,
			WarehouseID:  employeeReceived.WarehouseID,
		}
		savedEmployee, err := e.employeeService.Save(c, employeeToCreate)
		if err != nil {
			switch savedEmployee.ID {
			case 0:
				web.Error(c, http.StatusNotFound, "%s", err)
				return
			case -1:
				web.Error(c, http.StatusConflict, "%s", err)
				return
			case -2:
				web.Error(c, http.StatusInternalServerError, "%s", err)
				return
			}
		}

		web.Success(c, http.StatusCreated, savedEmployee)
	}
}

// Update an employee
// @Summary Update the employee with the required fields
// @Tags    Employees
// @Accept  json
// @Produce json
// @Param   id      path     int                 true  "Employee ID"
// @Param   token   header   string              false "Token"
// @Param   product body     employeeToUpdateDTO true  "Employee fields to update"
// @Success 200     {object} web.response        "Employee"
// @Failure 400     {object} web.response        "Malformed ID"
// @Failure 401     {object} web.response        "Unauthorized: empty or invalid token"
// @Failure 404     {object} web.response        "Employee not found"
// @Failure 409     {object} web.response        "The card number ID provided is already registered"
// @Failure 422     {object} web.response        "Gin validator error"
// @Failure 500     {object} web.response        "Failure executing update service"
// @Router  /employees/{id} [PATCH]
func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidID, id := parseID(c)
		if !isValidID {
			return
		}

		var employeeReceived employeeToUpdateDTO
		if err := c.ShouldBindJSON(&employeeReceived); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", err)
			return
		}

		employeeToUpdate := domain.Employee{
			CardNumberID: employeeReceived.CardNumberID,
			FirstName:    employeeReceived.FirstName,
			LastName:     employeeReceived.LastName,
			WarehouseID:  employeeReceived.WarehouseID,
		}

		updatedEmployee, err := e.employeeService.Update(c, id, employeeToUpdate)
		if err != nil {
			switch updatedEmployee.ID {
			case 0:
				web.Error(c, http.StatusNotFound, "%s", err)
				return
			case -1:
				web.Error(c, http.StatusInternalServerError, "%s", err)
				return
			case -2:
				web.Error(c, http.StatusConflict, "%s", err)
				return
			}
		}

		web.Success(c, http.StatusOK, updatedEmployee)
	}
}

// Delete employee
// @Summary Delete a required employee
// @Tags    Employees
// @Param   id    path   int    true "Employee ID"
// @Param   token header string true "Token"
// @Success 204
// @Failure 400 {object} web.response "Malformed ID"
// @Failure 401 {object} web.response "Unauthorized: empty or invalid token"
// @Failure 404 {object} web.response "Employee not found"
// @Failure 500 {object} web.response "Failure executing delete service"
// @Router  /employees/{id} [DELETE]
func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidID, id := parseID(c)
		if !isValidID {
			return
		}

		statusString, err := e.employeeService.Delete(c, id)
		if err != nil {
			switch statusString {
			case "ID not found":
				web.Error(c, http.StatusNotFound, "%s", err)
				return
			case "delete failed":
				web.Error(c, http.StatusInternalServerError, "%s", err)
				return
			}

		}

		web.Success(c, http.StatusNoContent, statusString)
	}
}

func convertModelToDTO(employee domain.Employee) employeeDTO {
	return employeeDTO{
		CardNumberID: employee.CardNumberID,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		WarehouseID:  employee.WarehouseID,
	}
}

func parseID(ctx *gin.Context) (bool, int) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		web.Error(ctx, http.StatusBadRequest, "invalid ID - %s", err)
		return false, 0
	} else if id <= 0 {
		web.Error(ctx, http.StatusBadRequest, "invalid ID: must be greater than 0")
		return false, 0
	}
	return true, id
}
