package database

import (
	"github.com/guiyones/brimos/internal/entity"
)

func (s *Service) CreateUser(user *entity.User) error {
	stmt, err := s.db.Prepare("INSERT INTO user(id, name, email, password) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FindUserByEmail(email string) (*entity.User, error) {
	stmt, err := s.db.Prepare("SELECT id, name, email FROM user WHERE email = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var u entity.User
	err = stmt.QueryRow(email).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
