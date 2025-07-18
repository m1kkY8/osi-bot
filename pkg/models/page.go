package models

import "github.com/m1kkY8/osi-bot/pkg/types"

// Re-export types for backward compatibility
type Page = types.Page

// NewPage creates a new pagination object
func NewPage(currentPage, perPage, totalPages int, pageMap map[string]int) *Page {
	return types.NewPage(currentPage, perPage, totalPages, pageMap)
}
