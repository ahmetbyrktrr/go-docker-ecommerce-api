package handler

import (
	"ecommerce/internal/model"
	"ecommerce/internal/response"
	"ecommerce/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	Service *service.CategoryService
}

func (h *CategoryHandler) Register(r *gin.Engine) {
	r.GET("/categories/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"message": "Category service is healthy",
		})
	})
	r.POST("/categories", h.Create)
	r.GET("/categories", h.GetAll)
	r.PUT("/categories/:id", h.Update)
	r.DELETE("/categories/:id", h.Delete)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var cat model.Category
	c.BindJSON(&cat)
	h.Service.Create(cat)
	response.JSON(c, http.StatusCreated, cat)
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	data, _ := h.Service.GetAll()
	response.JSON(c, http.StatusOK, data)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var cat model.Category
	c.BindJSON(&cat)
	err := h.Service.Update(id, cat.Name, cat.Products)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.JSON(c, http.StatusOK, "updated")
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	h.Service.Delete(id)
	response.JSON(c, http.StatusOK, "deleted")
}
