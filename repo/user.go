package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"test/models"

	"github.com/jmoiron/sqlx"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepo interface {
	Create(user models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user models.User) (*models.User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
		RETURNING id;
	`

	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.IsShopOwner).Scan(&user.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User

	query := `
		SELECT 
			id,
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		FROM users
		WHERE email = $1
		LIMIT 1;
	`

	err := r.db.Get(&user, query, email)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}
