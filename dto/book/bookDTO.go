package book

type CreateBookRequest struct {
	Title    string `json:"title"`
	Quantity int    `json:"quantity"`
}
