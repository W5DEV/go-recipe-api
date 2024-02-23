package models

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/datatypes"
)

type Recipe struct {
	ID            uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Title         string            `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Slug          string            `gorm:"uniqueIndex;not null" json:"slug,omitempty"`
	Description   string            `json:"description,omitempty"`
	Category      string            `json:"category,omitempty"`
	Ingredients   datatypes.JSON    `json:"ingredients,omitempty"`
	Instructions  datatypes.JSON    `json:"instructions,omitempty"`
	Image         string            `json:"image,omitempty"`
	User          uuid.UUID         `json:"user,omitempty"`
	CreatedAt     time.Time         `json:"created_at,omitempty"`
	UpdatedAt     time.Time         `json:"updated_at,omitempty"`
}

type CreateRecipeRequest struct {
	Title         string            `json:"title"  binding:"required"`
	Slug          string            `json:"slug"  binding:"required"`
	Description   string            `json:"description,omitempty"`
	Category      string            `json:"category,omitempty"`
	Ingredients   datatypes.JSON    `json:"ingredients,omitempty"`
	Instructions  datatypes.JSON    `json:"instructions,omitempty"`
	Image         string            `json:"image,omitempty"`
	User          string            `json:"user,omitempty"`
	CreatedAt     time.Time         `json:"created_at,omitempty"`
	UpdatedAt     time.Time         `json:"updated_at,omitempty"`
}

type UpdateRecipe struct {
	Title         string             `json:"title,omitempty"`
	Slug          string             `json:"slug,omitempty"`
	Description   string             `json:"description,omitempty"`
	Category      string             `json:"category,omitempty"`
	Ingredients   datatypes.JSON     `json:"ingredients,omitempty"`
	Instructions  datatypes.JSON     `json:"instructions,omitempty"`
	Image         string             `json:"image,omitempty"`
	User          string             `json:"user,omitempty"`
	CreateAt      time.Time          `json:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty"`
}