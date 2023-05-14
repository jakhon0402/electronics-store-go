package controller

import (
	"electronics-store-go/internal/app/models"
	"electronics-store-go/internal/app/payload"
	"electronics-store-go/middleware/handler"
	"electronics-store-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	Db *gorm.DB
}

func (h *Handler) CreateProduct(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		log := logger.NewLogger()
		var body payload.ProductDto

		if err := c.ShouldBind(&body); err != nil {
			log.Error("Failed to bind product data!")
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid body value", vErrs)
			}
		}
		newProduct := &models.Product{
			Name:          body.Name,
			Title:         body.Title,
			Price:         body.Price,
			Specification: body.Specification,
		}
		result := h.Db.Create(newProduct)
		if result.Error != nil {
			return handler.NewErrorResponse(http.StatusConflict, handler.InvalidQueryValue, "invalid create product", result)
		}
		return handler.NewSuccessResponse(http.StatusOK, result)
	})
}

func RouteV1(h *Handler, r *gin.Engine) {
	r.Group("/api/product")
}

func NewProductHandler(db *gorm.DB) *Handler {
	return &Handler{Db: db}
}
