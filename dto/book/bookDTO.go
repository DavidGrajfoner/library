package book

type CreateBookRequest struct {
	Title    string `json:"title" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}
