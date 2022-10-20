package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type SectionRequest struct {
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required"`
	WarehouseID        int `json:"warehouse_id" binding:"required"`
	ProductTypeID      int `json:"product_type_id" binding:"required"`
}

type UpdateSectionRequest struct {
	SectionNumber      int `json:"section_number,omitempty"`
	CurrentTemperature int `json:"current_temperature,omitempty"`
	MinimumTemperature int `json:"minimum_temperature,omitempty"`
	CurrentCapacity    int `json:"current_capacity,omitempty"`
	MinimumCapacity    int `json:"minimum_capacity,omitempty"`
	MaximumCapacity    int `json:"maximum_capacity,omitempty"`
	WarehouseID        int `json:"warehouse_id,omitempty"`
	ProductTypeID      int `json:"product_type_id,omitempty"`
}

type Section struct {
	sectionService section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
		sectionService: s,
	}
}

// Get all sections
// @Summary  Show the list of sections in db
// @Tags     Sections
// @Produce  json
// @Param    token  header    string        true  "Token"
// @Success  200    {object}  web.response  "List of sections"
// @Failure  401    {object}  web.response  "Unauthorized + invalid token"
// @Failure  500    {object}  web.response  "Internal server error "
// @Router   /sections [GET]
func (s *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sections, err := s.sectionService.GetAll(c)
		if err != nil {
			web.Error(c, 500, err.Error())
			log.Fatal(err)
			return
		}
		web.Success(c, 200, sections)
	}
}

// Get section by ID
// @Summary  Show the required section
// @Tags     Sections
// @Accept json
// @Produce  json
// @Param	id	path	int	true "Section ID"
// @Param    token  header    string        true  "Token"
// @Success  200    {object}  web.response  "Section"
// @Failure  401    {object}  web.response  "Unauthorized + invalid token"
// @Failure  404    {object}  web.response  "Section not found"
// @Router   /sections/{id} [GET]
func (s *Section) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, err.Error())
			log.Fatal(err)
			return
		}

		section, err := s.sectionService.GetByID(c, id)
		if err != nil {
			web.Error(c, 404, err.Error())
			return
		}
		web.Success(c, 200, section)

	}
}

// Create a new section
// @Summary  Save a new section in db
// @Tags     Sections
// @Accept   json
// @Produce  json
// @Param    token    header    string          true  "Token"
// @Param    product  body      SectionRequest  true  "Section to save"
// @Success  201      {object}  web.response "Created section"
// @Failure  422      {object}  web.response "Gin validator error"
// @Failure  500      {object}  web.response "Internal server error"
// @Failure  409      {object}  web.response "Conflict"
// @Failure  401    {object}  web.response  "Unauthorized + invalid token"
// @Router   /sections [POST]
func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req SectionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "Missing or malformed fields")
			return
		}

		if s.sectionService.SectionNumberAlreadyExists(c, req.SectionNumber) {
			web.Error(c, 409, fmt.Sprintf("Section number %d is already in use", req.SectionNumber))
			return
		}

		section, err := s.sectionService.Save(c, req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity, req.WarehouseID, req.ProductTypeID)
		if err != nil {
			web.Error(c, 500, err.Error())
		}
		web.Success(c, 201, section)

	}
}

// Update a section
// @Summary  Update the section with the required fields
// @Tags     Sections
// @Accept   json
// @Produce  json
// @Param    id       path      int             true   "Section ID"
// @Param    token    header    string          false  "Token"
// @Param    product  body      UpdateSectionRequest  true   "Section fields to update"
// @Success  200      {object}  web.response "Section"
// @Failure  409      {object}  web.response "Section number is already in use"
// @Failure  400      {object}  web.response "Bad request"
// @Failure  404      {object}  web.response "Section no found"
// @Failure  401    {object}  web.response  "Unauthorized: empty or invalid token"
// @Router   /sections/{id} [PUT]
func (s *Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, err.Error())
			return
		}

		section, err := s.sectionService.GetByID(c, id)
		if err != nil {
			web.Error(c, 404, err.Error())
			return
		}

		var req UpdateSectionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 400, err.Error())
			return
		}

		if req.SectionNumber != 0 {
			if s.sectionService.SectionNumberAlreadyExists(c, req.SectionNumber) {
				web.Error(c, 409, fmt.Sprintf("Section number %d is already in use", req.SectionNumber))
				return
			}
			section.SectionNumber = req.SectionNumber

		}

		if req.CurrentCapacity != 0 {
			section.CurrentCapacity = req.CurrentCapacity
		}

		if req.CurrentTemperature != 0 {
			section.CurrentTemperature = req.CurrentTemperature
		}

		if req.MaximumCapacity != 0 {
			section.CurrentCapacity = req.CurrentCapacity
		}

		if req.MinimumCapacity != 0 {
			section.MinimumCapacity = req.MaximumCapacity
		}

		if req.MinimumTemperature != 0 {
			section.MinimumTemperature = req.MinimumTemperature
		}

		if req.ProductTypeID != 0 {
			section.ProductTypeID = req.ProductTypeID
		}

		if req.WarehouseID != 0 {
			section.WarehouseID = req.WarehouseID
		}

		if err := s.sectionService.Update(c, section); err != nil {
			web.Error(c, 500, err.Error())
		}

		web.Success(c, 200, section)

	}
}

// Delete section
// @Summary  Delete a required section
// @Tags     Sections
// @Param    id     path      int     true  "Section ID"
// @Param    token  header    string  true  "Token"
// @Success  204
// @Failure  401    {object}  web.response "Unauthorized: empty or invalid token"
// @Failure  400    {object}  web.response "Bad request"
// @Failure  404    {object}  web.response "Section not found"
// @Router   /sections/{id} [DELETE]
func (s *Section) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, 400, err.Error())
			return
		}
		if err := s.sectionService.Delete(c, id); err != nil {
			web.Error(c, 404, err.Error())
			return
		}
		web.Empty(c, 204)
	}
}
