package position

type Position struct {
	ID       string `json:"id"`
	Code     string `json:"code" validate:"required"`
	Name     string `json:"name" validate:"required"`
	IsDelete string `json:"is_delete"`
}
