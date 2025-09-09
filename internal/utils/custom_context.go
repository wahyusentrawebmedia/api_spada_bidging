package utils

import (
	"api/spada/internal/database"
	"api/spada/internal/model"
	"api/spada/internal/response"

	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CustomContext struct {
	*fiber.Ctx
	// Add custom fields here if needed, e.g. UserID, RequestID, etc.
}

func NewCustomContext(c *fiber.Ctx) *CustomContext {
	return &CustomContext{Ctx: c}
}

func (cc *CustomContext) SuccessResponse(data interface{}, message string) error {
	return cc.Status(fiber.StatusOK).JSON(response.DefaultResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func (cc *CustomContext) ErrorResponse(message string) error {
	return cc.Status(fiber.StatusBadRequest).JSON(response.DefaultResponse{
		Status:  false,
		Message: message,
	})
}

func (cc *CustomContext) ErrorResponseUnauthorized(message string) error {
	return cc.Status(fiber.StatusUnauthorized).JSON(response.DefaultResponse{
		Status:  false,
		Message: message,
	})
}

func (cc *CustomContext) GetEndpoint() string {
	return cc.Locals("endpoint").(string)
}

func (cc *CustomContext) GetUsername() string {
	return cc.Locals("username").(string)
}

func (cc *CustomContext) SetLocalsParameter() error {
	dataPerguruanTinggi, err := cc.GetDataPerguruanTinggi()

	if err != nil {
		return err
	}

	cc.Locals("id_perguruan_tinggi", dataPerguruanTinggi.IDPerguruanTinggi)
	cc.Locals("endpoint", dataPerguruanTinggi.Endpoint)
	return nil
}

func (cc *CustomContext) GetDataPerguruanTinggi() (*model.PostgresConfig, error) {
	idStr := cc.Locals("id_perguruan_tinggi")
	if idStr == "" {
		return nil, errors.New("id_perguruan_tinggi is required")
	}

	var config model.PostgresConfig
	if err := database.DB.Where("id_perguruan_tinggi = ?", idStr).First(&config).Error; err != nil {
		return nil, err
	}

	return &config, nil
}

func (cc *CustomContext) GetPerguruanTinggi() string {
	return cc.Locals("id_perguruan_tinggi").(string)
}

var perguruanTinggiConnections = make(map[string]*gorm.DB)

func (cc *CustomContext) GetGormConnectionForPerguruanTinggi() (*gorm.DB, error) {
	config, err := cc.GetDataPerguruanTinggi()
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s:%s", config.Type, config.IDPerguruanTinggi)
	if db, ok := perguruanTinggiConnections[key]; ok {
		return db, nil
	}

	var db *gorm.DB
	if config.Type == "mysql" {
		db, err = ConnectionMySQL(config.User, config.Password, config.Host, fmt.Sprintf("%d", config.Port), config.DBName)
	} else {
		db, err = ConnectionDB(config.User, config.Password, config.Host, fmt.Sprintf("%d", config.Port), config.DBName)
	}
	if err != nil {
		return nil, err
	}

	perguruanTinggiConnections[key] = db
	return db, nil
}
