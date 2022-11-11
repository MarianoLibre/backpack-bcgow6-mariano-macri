package user

import (
	"database/sql"
	"errors"

	"library.com/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.User, error)
	Get(int) (domain.User, error)
	Save(domain.User) (domain.User, error)
	Update(domain.User) error
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

func (r *repository) GetAll() ([]domain.User, error) {
	rows, err := r.database.Query("select * from User")
	if err != nil {
		return []domain.User{}, nil
	}

	var users []domain.User
	for rows.Next() {
		u := domain.User{}
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return []domain.User{}, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *repository) Get(id int) (domain.User, error) {
	row := r.database.QueryRow("select * from User where id = ?", id)
	user := domain.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *repository) Save(user domain.User) (domain.User, error) {
	statement, err := r.database.Prepare("insert into User (Name, Age) values (?, ?)")
	if err != nil {
		return domain.User{}, err
	}

	result, err := statement.Exec(&user.Name, &user.Age)
	if err != nil {
		return domain.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}

	user.Id = int(id)
	return user, nil
}

func (r *repository) Update(user domain.User) error {
	statement, err := r.database.Prepare("update User set Name = ?, Age = ? where id = ?")
	if err != nil {
		return err
	}

	result, err := statement.Exec(&user.Name, &user.Age, &user.Id)
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
	statement, err := r.database.Prepare("delete from User where id = ?")
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
		return errors.New("User not found")
	}

	return nil
}
