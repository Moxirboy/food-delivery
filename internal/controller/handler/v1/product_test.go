package v1

import (
	"food-delivery/internal/configs"
	"food-delivery/internal/service/usecase"
	"food-delivery/pkg/logger"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNewProductHandler(t *testing.T) {
	type args struct {
		r   *gin.RouterGroup
		log logger.Logger
		uc  usecase.ProductUsecase
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewProductHandler(tt.args.r, tt.args.log, tt.args.uc)
		})
	}
}

func TestProductHandler_Delete(t *testing.T) {
	type fields struct {
		log    logger.Logger
		uc     usecase.ProductUsecase
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
			h := &ProductHandler{
				log:    tt.fields.log,
				uc:     tt.fields.uc,
				config: tt.fields.config,
			}
			h.Delete(tt.args.c)
		})
	}
}

func TestProductHandler_GetByID(t *testing.T) {
	type fields struct {
		log    logger.Logger
		uc     usecase.ProductUsecase
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
			h := &ProductHandler{
				log:    tt.fields.log,
				uc:     tt.fields.uc,
				config: tt.fields.config,
			}
			h.GetByID(tt.args.c)
		})
	}
}

func TestProductHandler_GetList(t *testing.T) {
	type fields struct {
		log    logger.Logger
		uc     usecase.ProductUsecase
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
			h := &ProductHandler{
				log:    tt.fields.log,
				uc:     tt.fields.uc,
				config: tt.fields.config,
			}
			h.GetList(tt.args.c)
		})
	}
}

func TestProductHandler_create(t *testing.T) {
	type fields struct {
		log    logger.Logger
		uc     usecase.ProductUsecase
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
			h := &ProductHandler{
				log:    tt.fields.log,
				uc:     tt.fields.uc,
				config: tt.fields.config,
			}
			h.create(tt.args.c)
		})
	}
}
