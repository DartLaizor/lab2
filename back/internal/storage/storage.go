package storage

import (
	"errors"
)

type Review struct {
	Name         string   `json:"name" validate:"required"`
	Date         string   `json:"date" validate:"required"`
	Phone        string   `json:"phone" validate:"required"`
	Email        string   `json:"email" validate:"required,email"`
	Technologies []string `json:"technologies" validate:"required,min=1"`
	Rating       int      `json:"rating" validate:"required,min=0,max=9"`
	Comment      string   `json:"comment" validate:"required"`
}

var (
	ErrURLExists      = errors.New("url exists")
	ErrURLNotFound    = errors.New("url not found")
	ErrAlisNotDeleted = errors.New("url not deleted")
)
