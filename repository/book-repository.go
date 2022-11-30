package repository

import (
    "github.com/Drack112/Crud-Golang-API/entity"
    "gorm.io/gorm"
)

type BookRepository interface {
    InsertBook(b entity.Book) entity.Book
    UpdateBook(b entity.Book) entity.Book
    DeleteBook(b entity.Book)
    AllBook() []entity.Book
    FindBookById(book_id uint64) entity.Book
}

type bookConnection struct {
    bookCon *gorm.DB
}

func NewBookRepository(connection_book *gorm.DB) BookRepository {
    return &bookConnection{
        bookCon: connection_book,
    }
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
    db.bookCon.Save(&b)
    db.bookCon.Preload("User").Find(&b)
    return b

}

func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
    db.bookCon.Save(&b)
    db.bookCon.Preload("User").Find(&b)
    return b
}

func (db *bookConnection) DeleteBook(b entity.Book) {
    db.bookCon.Delete(&b)
}

func (db *bookConnection) AllBook() []entity.Book {
    var books []entity.Book
    db.bookCon.Preload("User").Find(&books)
    return books
}

func (db *bookConnection) FindBookById(book_id uint64) entity.Book {
    var book entity.Book
    db.bookCon.Preload("User").Find(&book, book_id)
    return book
}
