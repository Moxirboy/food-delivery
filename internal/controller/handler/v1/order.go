package v1

import (
	"fmt"
	"food-delivery/internal/configs"
	"food-delivery/internal/dto"
	"food-delivery/internal/models"
	"food-delivery/internal/service/usecase"
	"food-delivery/pkg/logger"
	"food-delivery/pkg/utils"
	"github.com/gin-gonic/gin"
)

// orderHandler handles order-related operations.
type orderHandler struct {
	log    logger.Logger
	uco    usecase.IOrderUseCase
	ucc    usecase.ICartUseCase
	config configs.Config
}

// NeworderHandler initializes the order routes.
//
//	@BasePath	/api/v1
func NeworderHandler(
	r *gin.RouterGroup,
	log logger.Logger,
	uco usecase.IOrderUseCase,
	ucc usecase.ICartUseCase,
) {
	handler := &orderHandler{
		log: log,
		uco: uco,
		ucc: ucc,
	}

	orders := r.Group("/order")
	orders.POST("/create", handler.create)
	orders.GET("/get", handler.getAll)
	orders.PUT("/update", handler.updateStatus)
}

// create creates a new order.
//
//	@Summary		Create an order
//	@Description	Creates a new order based on the user's cart.
//	@Tags			order
//	@Accept			json
//	@Produce		json
//
// @Security Bearer
//
//	@Success		200	{object}	dto.BaseResponse	"Order created successfully"
//	@Failure		400	{object}	dto.BaseResponse	"Invalid request"
//	@Failure		401	{object}	dto.BaseResponse	"User not authenticated"
//	@Failure		500	{object}	dto.BaseResponse	"Internal server error"
//	@Router			/order/create [post]
func (h *orderHandler) create(c *gin.Context) {
	ID, exist := c.Get("ID")
	fmt.Println(ID)
	if !exist {
		utils.SendResponse(c, nil, utils.ErrNotAuthenticated)
		return
	}
	cart, err := h.ucc.GetByID(c.Request.Context(), ID.(string))
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	order := models.NewOrder(ID.(string), cart.ID, models.OrderStatusPending)
	err = h.uco.CreateOrder(c.Request.Context(), *order)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, nil, nil)
}

// getAll retrieves all orders.
//
//	@Summary		Get all orders
//	@Description	Retrieves a list of all orders.
//	@Tags			order
//	@Accept			json
//	@Produce		json
//
// @Security Bearer
//
//	@Success		200	{array}		dto.OrderResponse	"List of orders"
//	@Failure		500	{object}	dto.BaseResponse	"Internal server error"
//	@Router			/order/get [get]
func (h *orderHandler) getAll(c *gin.Context) {
	orders, err := h.uco.GetOrders(c.Request.Context())
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, ToOrderResponseArray(orders), nil)
}

// updateStatus updates the status of an order.
//
//	@Summary		Update order status
//	@Description	Updates the status of an order based on the provided status.
//	@Tags			order
//	@Accept			json
//	@Produce		json
//
// @Security Bearer
//
//	@Param			ID		header		string				true	"User ID"
//	@Param			body	body		dto.OrderStatus		true	"Order Status"
//	@Success		200		{object}	dto.BaseResponse	"Order status updated successfully"
//	@Failure		400		{object}	dto.BaseResponse	"Invalid request"
//	@Failure		401		{object}	dto.BaseResponse	"User not authenticated"
//	@Failure		500		{object}	dto.BaseResponse	"Internal server error"
//	@Router			/order/update [put]
func (h *orderHandler) updateStatus(c *gin.Context) {
	ID, exist := c.Get("ID")
	body := dto.OrderStatus{}
	if !exist {
		utils.SendResponse(c, nil, utils.ErrNotAuthenticated)
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	invalidParams := utils.Validate(body)
	if invalidParams != nil {
		utils.SendResponse(c, invalidParams, nil)
		return
	}
	status, err := models.StatusMaker(body.Status)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	cart, err := h.ucc.GetByID(c.Request.Context(), ID.(string))
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	order := models.NewOrder(ID.(string), cart.ID, status)
	err = h.uco.UpdateOrder(c.Request.Context(), *order)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, nil, nil)
}
