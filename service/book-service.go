package service

import (
    "fmt"
    "log"

    "github.com/Drack112/Crud-Golang-API/dto"
    "github.com/Drack112/Crud-Golang-API/entity"
    "github.com/Drack112/Crud-Golang-API/repository"
    "github.com/mashingan/smapping"
)

type BookService interface {
    Insert(b dto.BookCreateDTO) entity.Book
    Update(b dto.BookUpdateDTO) entity.Book
    Delete(b entity.Book)
    All() []entity.Book
    FindByID(bookID uint64) entity.Book
    IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
    bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
    return &bookService{
        bookRepository: bookRepo,
    }
}

func (service *bookService) Insert(b dto.BookCreateDTO) entity.Book {
    book := entity.Book{}
    err := smapping.FillStruct(&book, smapping.MapFields(&b))
    if err != nil {
        log.Fatalf("Failed map %v", err)
    }

    res := service.bookRepository.InsertBook(book)
    return res

}

func (service *bookService) Update(b dto.BookUpdateDTO) entity.Book {
    book := entity.Book{}
    err := smapping.FillStruct(&book, smapping.MapFields(&b))
    if err != nil {
        log.Fatalf("Failed map %v: ", err)
    }

    res := service.bookRepository.UpdateBook(book)
    return res
}

func (service *bookService) Delete(b entity.Book) {
    service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() []entity.Book {
    return service.bookRepository.AllBook()
}

func (service *bookService) FindByID(bookID uint64) entity.Book {
    return service.bookRepository.FindBookById(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
    b := service.bookRepository.FindBookById(bookID)
    id := fmt.Sprintf("%v", b.UserID)
    return userID == id
}
