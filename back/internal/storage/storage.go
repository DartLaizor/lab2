package storage

import (
	"errors"
)

type Review struct {
	Name         string   `json:"name"`
	Date         string   `json:"date"`
	Phone        string   `json:"phone"`
	Email        string   `json:"email"`
	Technologies []string `json:"technologies"`
	Rating       int      `json:"rating"`
	Comment      string   `json:"comment"`
}

var (
	ErrURLExists      = errors.New("url exists")
	ErrURLNotFound    = errors.New("url not found")
	ErrAlisNotDeleted = errors.New("url not deleted")
)
