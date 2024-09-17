package types

import "github.com/google/uuid"

/*
Define a User struct with the following fields:
- ID: A UUID that represents the user
- Username: A string that represents the username of the user
- Email: A string that represents the email of the user
- Password: A string that represents the password of the user
- isSuperDuper: A boolean that represents if the user is a super user
*/
type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	// Super Dummy User is The Admin Privileged User
	// Can Access All The Resources
	// Can Perform All The Operations
	IsSuperDuper bool `json:"isSuperDuper"`
}

/*
Define a Book struct with the following fields:
- ISBN: A string that represents the ISBN of the book
- ID: A UUID that represents the book
- Title: A string that represents the title of the book
- Author: A string that represents the author of the book
- isAvailable: A boolean that represents if the book is available
- Status: A string that represents the status of the book
*/
type Book struct {
	ISBN        string    `json:"isbn"`
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	IsAvailable bool      `json:"isAvailable"`
	Status      string    `json:"status"`
}

/*
Define a BorrowRequest struct with the following fields:
- UserID: A UUID that represents the user who is borrowing the book
- BookID: A UUID that represents the book that is being borrowed
- Status: A string that represents the status of the borrow request
*/
type BorrowRequest struct {
	UserID uuid.UUID `json:"userId"`
	BookID uuid.UUID `json:"bookId"`
	Status string    `json:"status"`
}
