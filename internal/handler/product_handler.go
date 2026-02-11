package handler

import (
	"ecommerce/internal/model"
	"ecommerce/internal/response"
	"ecommerce/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Service *service.ProductService
}

func (h *ProductHandler) Register(r *gin.Engine) {
	r.POST("/products", h.Create)
	r.GET("/products", h.GetAll)
	r.PUT("/products/:id", h.Update)
	r.DELETE("/products/:id", h.Delete)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var p model.Product
	c.BindJSON(&p)
	h.Service.Create(p)
	response.JSON(c, http.StatusCreated, p)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	data, _ := h.Service.GetAll()
	response.JSON(c, http.StatusOK, data)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var p model.Product
	c.BindJSON(&p)
	h.Service.Update(id, p)
	response.JSON(c, http.StatusOK, "updated")
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	h.Service.Delete(id)
	response.JSON(c, http.StatusOK, "deleted")
}
