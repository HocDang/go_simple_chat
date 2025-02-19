package postgres

import (
	"chat-server/internal/domain/entities"
	"chat-server/internal/domain/repositories"

	"github.com/go-pg/pg/v10"
)

type userPgRepository struct {
	db *pg.DB
}

func NewUserPgRepository(db *pg.DB) repositories.UserRepository {
	return &userPgRepository{db: db}
}

func (r *userPgRepository) Create(user *entities.User) error {
	_, err := r.db.Model(user).Insert()
	return err
}

func (r *userPgRepository) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userPgRepository) GetAll() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}
