package models

type Page struct {
	CurrentPage int
	PerPage     int
	TotalPages  int
	PageMap     map[string]int
}

func NewPage(currentPage, perPage, totalPages int, pageMap map[string]int) *Page {
	return &Page{
		CurrentPage: currentPage,
		PerPage:     perPage,
		TotalPages:  totalPages,
		PageMap:     pageMap,
	}
}
