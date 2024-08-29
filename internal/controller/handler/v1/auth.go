package v1

import (
	"context"
	"food-delivery/internal/dto"
	"food-delivery/internal/models"
	"food-delivery/internal/service/usecase"
	"food-delivery/pkg/logger"
	"food-delivery/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	log       logger.Logger
	AccountUC usecase.IAccountUseCase
	AuthUC    usecase.IAuthUseCase
}

func NewAuthHandler(
	r *gin.RouterGroup,
	l logger.Logger,
	AuthUC usecase.IAuthUseCase,
	AccountUC usecase.IAccountUseCase,
) {
	handler := &AuthHandler{
		log:       l,
		AuthUC:    AuthUC,
		AccountUC: AccountUC,
	}
	auths := r.Group("/auth")

	auths.POST("/login", handler.login)
	auths.POST("/refresh", handler.renew)
	auths.POST("/sign", handler.signUp)
}

// @Router /v1/auth/refresh [post]
func (a *AuthHandler) renew(c *gin.Context) {
	var body dto.RenewRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	res, err := a.AuthUC.ReNew(
		context.Background(),
		body.RefreshToken,
	)

	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	utils.SendResponse(
		c,
		toRenewResponse(c.Request.Context(), res),
		nil,
	)
}

// @Router /v1/auth/login [post]
func (h *AuthHandler) login(c *gin.Context) {
	var body dto.LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error parsing body: ", err)
		utils.SendResponse(c, nil, err)
		return
	}

	invalidParams := utils.Validate(body)
	if invalidParams != nil {
		utils.SendResponse(c, invalidParams, nil)
		return
	}

	user, err := h.AccountUC.LoginUser(
		c.Request.Context(),
		body.Login,
		body.Password,
	)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	tokens, err := h.AuthUC.New(c.Request.Context(),
		user.ID,
		user.Position,
	)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	utils.SendResponse(c,
		toLoginResponse(c.Request.Context(), user, tokens),
		nil,
	)
}

// @Router /v1/auth/sign [post]
func (h *AuthHandler) signUp(c *gin.Context) {
	var body dto.SignUpRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error parsing body: ", err)
		utils.SendResponse(c, nil, err)
		return
	}

	invalidParams := utils.Validate(body)
	if invalidParams != nil {
		utils.SendResponse(c, invalidParams, nil)
		return
	}
	user := models.NewUser(body)
	err := h.AccountUC.CreateUser(
		c.Request.Context(),
		user,
	)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	tokens, err := h.AuthUC.New(c.Request.Context(),
		user.ID,
		user.Position,
	)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	utils.SendResponse(c,
		toLoginResponse(c.Request.Context(), user, tokens),
		nil,
	)
}
