package handler

import (
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Address            *string `json:"address,omitempty"`
	Telephone          *string `json:"telephone,omitempty"`
	WarehouseCode      *string `json:"warehouse_code,omitempty"`
	MinimumCapacity    *int    `json:"minimum_capacity,omitempty"`
	MinimumTemperature *int    `json:"minimum_temperature,omitempty"`
}

type Warehouse struct {
	warehouseService warehouse.Service
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: w,
	}
}

// ListWarehouses godoc
// @Summary List warehouses
// @Tags Warehouses
// @Description get warehouses
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Success 200 object web.response
// @Router /api/v1/warehouses [get]
func (w *Warehouse) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, "Invalid Id")
			return
		}
		wh, err := w.warehouseService.Get(c, int(id))
		if err != nil {
			web.Error(c, 404, "Not found")
			return
		}
		web.Success(c, 200, wh)
	}
}

// FindWarehouse godoc
// @Summary Get warehouse by id
// @Tags Warehouses
// @Description get warehouses
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Success 200 object web.response
// @Router /api/v1/warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := w.warehouseService.GetAll(c)
		if err != nil {
			web.Error(c, 500, err.Error())
			return
		}
		web.Success(c, 200, data)
	}
}

// CreateWarehouse godoc
// @Summary Create warehouse
// @Tags Warehouses
// @Description create warehouses
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Success 200 object web.response
// @Router /api/v1/warehouses [post]
func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Warehouse
		if err := c.Bind(&req); err != nil {
			web.Error(c, 400, err.Error())
			return
		}
		switch {
		case req.Address == "":
			web.Error(c, 422, "Error: '%s' is required.", "Address")
			return
		case req.Telephone == "":
			web.Error(c, 422, "Error: '%s' is required.", "Telephone")
			return
		case req.WarehouseCode == "":
			web.Error(c, 422, "Error: '%s' is required.", "WarehouseCode")
			return
		case req.MinimumCapacity == 0:
			web.Error(c, 422, "Error: '%s' is required.", "MinimumCapacity")
			return
		case req.MinimumTemperature == 0:
			web.Error(c, 422, "Error: '%s' is required.", "MinimumTemperature")
			return
		}

		if w.warehouseService.Exists(c, req.WarehouseCode) {
			web.Error(c, 409, "Error: warehouse code '%s' is already in use.", req.WarehouseCode)
			return
		}

		id, err := w.warehouseService.Save(c, req)
		if err != nil {
			web.Error(c, 500, err.Error())
			return
		}
		req.ID = id
		web.Success(c, 201, req)
	}
}

// UpdateWarehouse godoc
// @Summary Update warehouse
// @Tags Warehouses
// @Description update warehouses
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Success 200 object web.response
// @Router /api/v1/warehouses [patch]
func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, err.Error())
			return
		}
		wh, err := w.warehouseService.Get(c, int(id))
		if err != nil {
			web.Error(c, 400, err.Error())
			return
		}

		var req request
		if err := c.Bind(&req); err != nil {
			web.Error(c, 400, "Error: %s", err.Error())
			return
		}

		{
			if req.Address != nil {
				wh.Address = *req.Address
			}

			if req.Telephone != nil {
				wh.Telephone = *req.Telephone
			}

			if req.WarehouseCode != nil {
				wh.WarehouseCode = *req.WarehouseCode
			}

			if req.MinimumCapacity != nil {
				wh.MinimumCapacity = *req.MinimumCapacity
			}

			if req.MinimumTemperature != nil {
				wh.MinimumTemperature = *req.MinimumTemperature
			}
		}

		err = w.warehouseService.Update(c, wh)
		if err != nil {
			web.Error(c, 404, "Error: %s", err.Error())
			return
		}
		web.Success(c, 200, wh)
	}
}

// DeleteWarehouse godoc
// @Summary Delete warehouse
// @Tags Warehouses
// @Description delete warehouses
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Success 200 object web.response
// @Router /api/v1/warehouses [delete]
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, "Error: %s", err.Error())
			return
		}

		err = w.warehouseService.Delete(c, int(id))
		if err != nil {
			web.Error(c, 404, "Error: %s", err.Error())
			return
		}

		web.Success(c, 204, nil)
	}
}
