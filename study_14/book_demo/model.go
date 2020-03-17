package main

type Book struct {
	ID int64 `db:"id"`
	Title string `db:"title" form:"title"`
	Price string `db:"price" form:"price"`
}
