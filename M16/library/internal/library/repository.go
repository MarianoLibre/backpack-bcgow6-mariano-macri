package library

import (
	"database/sql"
	"errors"

	"library.com/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Library, error)
	Get(int) (domain.Library, error)
	Save(domain.Library) (int, error)
	Update(domain.Library) error
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

func (r *repository) GetAll() ([]domain.Library, error) {
	rows, err := r.database.Query("select * from Library")
	if err != nil {
		return []domain.Library{}, nil
	}

	var libraries []domain.Library
	for rows.Next() {
		l := domain.Library{}
		if err := rows.Scan(&l.Id, &l.Name, &l.Address, &l.PhoneNumber); err != nil {
			return []domain.Library{}, err
		}
		libraries = append(libraries, l)
	}

	return libraries, nil
}

func (r *repository) Get(id int) (domain.Library, error) {
	row := r.database.QueryRow("select * from Library where id = ?", id)
	library := domain.Library{}
	err := row.Scan(&library.Id, &library.Name, &library.Address, &library.PhoneNumber)
	if err != nil {
		return domain.Library{}, err
	}

	return library, nil
}

func (r *repository) Save(library domain.Library) (int, error) {
	statement, err := r.database.Prepare("insert into Library (Name, Address, PhoneNumber) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(&library.Name, &library.Address, &library.PhoneNumber)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(library domain.Library) error {
	statement, err := r.database.Prepare("update Library set Name = ?, Address = ?, PhoneNumber = ? where id = ?")
	if err != nil {
		return err
	}

	result, err := statement.Exec(&library.Name, &library.Address, &library.PhoneNumber, &library.Id)
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
	statement, err := r.database.Prepare("delete from Library where id = ?")
	if err != nil {
		return err
	}

	result, err := statement.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("Library not found")
	}

	return nil
}
