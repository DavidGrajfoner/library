package book

type CreateBookRequest struct {
	Title    string `json:"title" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type BookResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Quantity int    `json:"quantity"`
}
