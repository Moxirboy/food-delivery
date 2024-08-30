package v1

import (
	"food-delivery/internal/configs"
	"food-delivery/internal/service/usecase"
	"food-delivery/pkg/logger"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNewCartHandler(t *testing.T) {
	type args struct {
		r   *gin.RouterGroup
		log logger.Logger
		uc  usecase.ICartUseCase
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewCartHandler(tt.args.r, tt.args.log, tt.args.uc)
		})
	}
}

func Test_cartHandler_addProduct(t *testing.T) {
	type fields struct {
		log    logger.Logger
		uc     usecase.ICartUseCase
		config configs.Config
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cartHandler{
				log:    tt.fields.log,
				uc:     tt.fields.uc,
				config: tt.fields.config,
			}
			h.addProduct(tt.args.c)
		})
	}
}

func Test_cartHandler_create(t *testing.T) {
	type fields struct {
		log    logger.Logger
		uc     usecase.ICartUseCase
		config configs.Config
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cartHandler{
				log:    tt.fields.log,
				uc:     tt.fields.uc,
				config: tt.fields.config,
			}
			h.create(tt.args.c)
		})
	}
}

func Test_cartHandler_get(t *testing.T) {
	type fields struct {
		log    logger.Logger
		uc     usecase.ICartUseCase
		config configs.Config
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cartHandler{
				log:    tt.fields.log,
				uc:     tt.fields.uc,
				config: tt.fields.config,
			}
			h.get(tt.args.c)
		})
	}
}

func Test_cartHandler_removeProduct(t *testing.T) {
	type fields struct {
		log    logger.Logger
		uc     usecase.ICartUseCase
		config configs.Config
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cartHandler{
				log:    tt.fields.log,
				uc:     tt.fields.uc,
				config: tt.fields.config,
			}
			h.removeProduct(tt.args.c)
		})
	}
}
