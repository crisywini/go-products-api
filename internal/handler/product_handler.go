package handler

import (
	"log"
	"net/http"

	"github.com/crisywini/go-products-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

type createProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int32   `json:"stock" binding:"required,gte=0"`
}

type updateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int32   `json:"stock" binding:"required,gte=0"`
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req createProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[CreateProduct] bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		log.Printf("[CreateProduct] userID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		log.Printf("[CreateProduct] userID in context has unexpected type: %T", userIDValue)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID in token"})
		return
	}

	log.Printf("[CreateProduct] user=%s name=%q price=%.2f stock=%d", userID, req.Name, req.Price, req.Stock)

	product, err := h.productService.CreateProduct(
		c.Request.Context(),
		req.Name, req.Description, req.Price, req.Stock, userID,
	)
	if err != nil {
		log.Printf("[CreateProduct] service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product, err := h.productService.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.productService.ListProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var req updateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.productService.UpdateProduct(
		c.Request.Context(),
		id, req.Name, req.Description, req.Price, req.Stock,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	if err := h.productService.DeleteProduct(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
