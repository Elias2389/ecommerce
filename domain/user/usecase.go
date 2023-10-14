package user

import (
	"ecomerce-go/model"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	repository Repository
}

func New(r Repository) User {
	return User{repository: r}
}

func (u User) Create(m * model.User) error  {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	m.ID = ID
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword", err)
	}

	m.Password = string(password)

	if m.Details == nil {
		m.Details = []byte("{}")
	}
	m.CreatedAt = time.Now().Unix()

	err = u.repository.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "Repository.Create()", err)
	}
	m.Password = ""
	return nil
}

func (u User) GetByEmail(email string) (model.User, error)  {
	user, err := u.repository.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "repository.GetByEmail()", err)
	}

	return user, nil
}

func (u User) GetAll() (model.Users, error)  {
	users, err := u.repository.GetAll()
	if err != nil {
		return model.Users{}, fmt.Errorf("%s %w", "repository.GetByAll()", err)
	}

	return users, nil
}
