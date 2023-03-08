package controller

import (
	"errors"

	"github.com/shoriwe/rollee-test-2/models"
	"gorm.io/gorm"
)

func (c *Controller) AddWord(word *models.Word) error {
	var dbWord models.Word
	fErr := c.db.First(&dbWord, "word = ?", word.Word).Error
	switch {
	case fErr == nil:
		return c.db.Exec("UPDATE words SET repeated = repeated + 1 WHERE uuid = ?", dbWord.UUID).Error
	case fErr != nil && errors.Is(fErr, gorm.ErrRecordNotFound):
		return c.db.Create(word).Error
	default:
		return fErr
	}
}

func (c *Controller) QueryWord(pattern string) (*models.Word, error) {
	result := &models.Word{}
	fErr := c.db.
		Order("repeated DESC").
		First(result, "word LIKE ?", pattern+"%").
		Error
	switch {
	case fErr == nil:
		return result, nil
	case fErr != nil && errors.Is(fErr, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, fErr
	}
}
