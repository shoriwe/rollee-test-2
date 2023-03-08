package controller

import (
	"github.com/shoriwe/rollee-test-2/models"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func (c *Controller) Close() {
	conn, err := c.db.DB()
	if err != nil {
		panic(err)
	}
	conn.Close()
}

func New(db *gorm.DB) *Controller {
	err := db.AutoMigrate(&models.Word{})
	if err != nil {
		panic(err)
	}
	return &Controller{db: db}
}
