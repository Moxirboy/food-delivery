package v1

import (
	"food-delivery/internal/configs"
	"food-delivery/internal/dto"
	"food-delivery/internal/models"
	"food-delivery/internal/service/usecase"
	"food-delivery/pkg/logger"
	"food-delivery/pkg/utils"
	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	log    logger.Logger
	uc     usecase.ICartUseCase
	config configs.Config
}

// NewCartHandler initializes the cart routes.
//
//	@BasePath	/api/v1
func NewCartHandler(
	r *gin.RouterGroup,
	log logger.Logger,
	uc usecase.ICartUseCase,
) {
	handler := &cartHandler{
		log: log,
		uc:  uc,
	}

	carts := r.Group("/cart")
	carts.POST("/create", handler.create)
	carts.GET("/get", handler.get)
	carts.DELETE("/remove", handler.removeProduct)
	carts.POST("/add", handler.addProduct)
}

// create creates a new cart for the user.
//
//	@Summary		Create a new cart
//	@Description	Creates a new cart for the authenticated user.
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//
// @Security Bearer
//
//	@Success		200	{object}	dto.BaseResponse	"Cart created successfully"
//	@Failure		401	{object}	dto.BaseResponse	"User not authenticated"
//	@Failure		500	{object}	dto.BaseResponse	"Internal server error"
//	@Router			/cart/create [post]
func (h *cartHandler) create(c *gin.Context) {
	ID, exist := c.Get("ID")
	if !exist {
		utils.SendResponse(c, nil, utils.ErrNotAuthenticated)
		return
	}
	cart := models.NewCart(ID.(string))
	err := h.uc.Create(c.Request.Context(), cart)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, ToCartResponse(*cart), nil)
}

// get retrieves the cart for the user.
//
//	@Summary		Get the cart
//	@Description	Retrieves the cart for the authenticated user.
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//
// @Security Bearer
//
//	@Success		200	{object}	dto.BaseResponse	"Cart retrieved successfully"
//	@Failure		401	{object}	dto.BaseResponse	"User not authenticated"
//	@Failure		500	{object}	dto.BaseResponse	"Internal server error"
//	@Router			/cart/get [get]
func (h *cartHandler) get(c *gin.Context) {
	ID, exist := c.Get("ID")
	if !exist {
		utils.SendResponse(c, nil, utils.ErrNotAuthenticated)
		return
	}
	cart, err := h.uc.GetByID(c.Request.Context(), ID.(string))
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, ToCartResponse(*cart), nil)
}

// addProduct adds a product to the cart.
//
//	@Summary		Add a product to the cart
//	@Description	Adds a specified product to the user's cart.
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//
// @Security Bearer
//
//	@Param			body	body		dto.CartRequest		true	"Product details"
//	@Success		200		{object}	dto.BaseResponse	"Product added successfully"
//	@Failure		401		{object}	dto.BaseResponse	"User not authenticated"
//	@Failure		400		{object}	dto.BaseResponse	"Invalid request"
//	@Failure		500		{object}	dto.BaseResponse	"Internal server error"
//	@Router			/cart/add [post]
func (h *cartHandler) addProduct(c *gin.Context) {
	ID, exist := c.Get("ID")
	if !exist {
		utils.SendResponse(c, nil, utils.ErrNotAuthenticated)
		return
	}
	body := dto.CartRequest{}
	if err := c.BindJSON(&body); err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	cart, err := h.uc.GetByID(c.Request.Context(), ID.(string))
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	productCart := models.NewCartProduct(cart.ID, body.ProductID, body.Quantity)

	err = h.uc.AddProduct(c.Request.Context(), ID.(string), productCart)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, nil, nil)
}

// removeProduct removes or updates the quantity of a product in the cart.
//
//	@Summary		Remove a product from the cart
//	@Description	Removes or updates the quantity of a specified product in the user's cart.
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//
// @Security Bearer
//
//	@Param			body	body		dto.CartRequest		true	"Product details"
//	@Success		200		{object}	dto.BaseResponse	"Product removed or quantity updated successfully"
//	@Failure		401		{object}	dto.BaseResponse	"User not authenticated"
//	@Failure		400		{object}	dto.BaseResponse	"Invalid request"
//	@Failure		500		{object}	dto.BaseResponse	"Internal server error"
//	@Router			/cart/remove [delete]
func (h *cartHandler) removeProduct(c *gin.Context) {
	ID, exist := c.Get("ID")
	if !exist {
		utils.SendResponse(c, nil, utils.ErrNotAuthenticated)
		return
	}
	body := dto.CartRequest{}
	if err := c.BindJSON(&body); err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	cart, err := h.uc.GetByID(c.Request.Context(), ID.(string))
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	productCart := models.NewCartProduct(cart.ID, body.ProductID, body.Quantity)

	err = h.uc.UpdateQuantity(c.Request.Context(), productCart)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, nil, nil)
}
