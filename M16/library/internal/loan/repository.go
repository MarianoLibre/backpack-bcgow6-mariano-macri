package loan

import (
	"database/sql"
	"errors"

	"library.com/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Loan, error)
	Get(int) (domain.Loan, error)
	Save(domain.Loan) (int, error)
	Update(domain.Loan) error
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
	selectAllLoans = "select * from Loan"
	selectLoanById = "select * from Loan where id = ?"
	insertLoan = "insert into Loan (BookId, UserId) values (?, ?)"
	updateLoan = "update Loan set BookId = ?, UserId = ? where id = ?"
	deleteLoan = "delete from Loan where id = ?"
)

func (r *repository) GetAll() ([]domain.Loan, error) {
	rows, err := r.database.Query(selectAllLoans)
	if err != nil {
		return []domain.Loan{}, nil
	}

	var loans []domain.Loan
	for rows.Next() {
		l := domain.Loan{}
		if err := rows.Scan(&l.Id, &l.BookId, &l.UserId); err != nil {
			return []domain.Loan{}, err
		}
		loans = append(loans, l)
	}

	return loans, nil
}

func (r *repository) Get(id int) (domain.Loan, error) {
	row := r.database.QueryRow(selectLoanById, id)
	loan := domain.Loan{}
	err := row.Scan(&loan.Id, &loan.BookId, &loan.UserId)
	if err != nil {
		return domain.Loan{}, err
	}

	return loan, nil
}

func (r *repository) Save(loan domain.Loan) (int, error) {
	statement, err := r.database.Prepare(insertLoan)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(&loan.BookId, &loan.BookId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(loan domain.Loan) error {
	statement, err := r.database.Prepare(updateLoan)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(&loan.BookId, &loan.UserId)
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
	statement, err := r.database.Prepare(deleteLoan)
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
		return errors.New("Loan not found")
	}

	return nil
}
