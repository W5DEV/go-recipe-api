package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Recipe struct {
	ID            uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Title         string            `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Description   sql.NullString    `json:"description,omitempty"`
	Ingredients   sql.NullString    `gorm:"type:text[]" json:"ingredients,omitempty"`
	Instructions  sql.NullString    `gorm:"type:text[]" json:"instructions,omitempty"`
	Image         sql.NullString    `json:"image,omitempty"`
	User          uuid.UUID         `json:"user,omitempty"`
	CreatedAt     time.Time         `json:"created_at,omitempty"`
	UpdatedAt     time.Time         `json:"updated_at,omitempty"`
}

type CreateRecipeRequest struct {
	Title         string            `json:"title"  binding:"required"`
	Description   sql.NullString    `json:"description,omitempty"`
	Ingredients   sql.NullString    `gorm:"type:text[]" json:"ingredients,omitempty"`
	Instructions  sql.NullString    `gorm:"type:text[]" json:"instructions,omitempty"`
	Image         sql.NullString    `json:"image,omitempty"`
	User          sql.NullString    `json:"user,omitempty"`
	CreatedAt     time.Time         `json:"created_at,omitempty"`
	UpdatedAt     time.Time         `json:"updated_at,omitempty"`
}

type UpdateRecipe struct {
	Title         string             `json:"title,omitempty"`
	Description   sql.NullString     `json:"description,omitempty"`
	Ingredients   sql.NullString     `gorm:"type:text[]" json:"ingredients,omitempty"`
	Instructions  sql.NullString     `gorm:"type:text[]" json:"instructions,omitempty"`
	Image         sql.NullString     `json:"image,omitempty"`
	User          sql.NullString     `json:"user,omitempty"`
	CreateAt      time.Time          `json:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty"`
}