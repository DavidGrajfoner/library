package borrow

type BorrowReturnRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	BookID uint `json:"book_id" binding:"required"`
}

type BorrowResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	BookID    uint   `json:"book_id"`
	BookTitle string `json:"book_title"`
}