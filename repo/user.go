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
	Update(user models.User) (*models.User, error)
	GetByID(id int) (*models.User, error)
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

func (r *userRepo) Update(user models.User) (*models.User, error) {
	query := `
		UPDATE users
		SET
			first_name = $1,
			last_name = $2,
			is_shop_owner = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
		RETURNING id, first_name, last_name, email, password, is_shop_owner
	`

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.IsShopOwner, user.ID)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsShopOwner)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetByID(id int) (*models.User, error) {
	var user models.User

	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE id = $1
		LIMIT 1;
	`

	err := r.db.Get(&user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
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
