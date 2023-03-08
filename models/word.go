package models

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

var (
	checkWordRegex = regexp.MustCompile("(?m)^[a-zA-Z]+$")
)

type Word struct {
	Model
	IsNull   bool   `gorm:"-" json:"isNull,omitempty"`
	Word     string `gorm:"unique;not null;" json:"word,omitempty"`
	Repeated int    `gorm:"not null;default:1;" json:"repeated,omitempty"`
	Pattern  string `gorm:"-" json:"pattern,omitempty"`
}

func (word *Word) BeforeSave(tx *gorm.DB) error {
	mErr := word.Model.BeforeSave(tx)
	if mErr != nil {
		return mErr
	}
	if !checkWordRegex.MatchString(word.Word) {
		return fmt.Errorf("invalid word provided, expecting pattern ^[a-zA-Z]+$")
	}
	return nil
}
