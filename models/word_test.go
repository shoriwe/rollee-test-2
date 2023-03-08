package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/shoriwe/rollee-test-2/common/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestValidWord(t *testing.T) {
	db := sqlite.NewTest()
	assert.Nil(t, db.AutoMigrate(&Word{}))
	assert.Nil(t, db.Create(&Word{Word: "Rollee"}).Error)
	var word Word
	assert.Nil(t, db.First(&word, "word = ?", "Rollee").Error)
	assert.Equal(t, "Rollee", word.Word)
	assert.NotEqual(t, uuid.Nil, word.UUID)
}

func TestInvalidWord(t *testing.T) {
	db := sqlite.NewTest()
	assert.Nil(t, db.AutoMigrate(&Word{}))
	assert.NotNil(t, db.Create(&Word{Word: "234897234"}).Error)
	var word Word
	assert.NotNil(t, db.First(&word, "word = ?", "234897234").Error)
	assert.NotEqual(t, "234897234", word.Word)
}
