package borrow

type BorrowReturnRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	BookID uint `json:"book_id" binding:"required"`
}