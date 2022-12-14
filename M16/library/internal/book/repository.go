package book

import (
	"database/sql"
	"errors"

	"library.com/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Book, error)
	Get(int) (domain.Book, error)
	Save(domain.Book) (int, error)
	Update(domain.Book) error
	Delete(int) error
}

type repository struct {
	database *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &repository{
		database: database,
	}
}

const (
	selectAllBooks = "select Title, Quantity from Book"
	selectBookById = "select Title, Quantity from Book where id = ?"
	insertBook = "insert into Book (Title, Quantity) values (?, ?)"
	updateBook = "update Book set Title = ?, Quantity = ? where id = ?"
	deleteBook = "delete from Book where id = ?"
)

func (r *repository) GetAll() ([]domain.Book, error) {
	rows, err := r.database.Query(selectAllBooks)
	if err != nil {
		return []domain.Book{}, nil
	}

	var books []domain.Book
	for rows.Next() {
		b := domain.Book{}
		if err := rows.Scan(&b.Id, &b.Title, &b.Quantity); err != nil {
			return []domain.Book{}, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (r *repository) Get(id int) (domain.Book, error) {
	row := r.database.QueryRow(selectBookById, id)
	book := domain.Book{}
	err := row.Scan(&book.Id, &book.Title, &book.Quantity)
	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (r *repository) Save(book domain.Book) (int, error) {
	statement, err := r.database.Prepare(insertBook)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(&book.Title, &book.Quantity)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(book domain.Book) error {
	statement, err := r.database.Prepare(updateBook)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(&book.Title, &book.Quantity, &book.Id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(id int) error {
	statement, err := r.database.Prepare(deleteBook)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("Book not found")
	}

	return nil
}
