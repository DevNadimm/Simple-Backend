package repo

import (
	"database/sql"
	"errors"
	"test/models"

	"github.com/jmoiron/sqlx"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryRepo interface {
	Create(category models.Category) (*models.Category, error)
	List() ([]*models.Category, error)
	GetByID(id int) (*models.Category, error)
	Update(category models.Category) (*models.Category, error)
	Delete(id int64) error
}

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{db: db}
}

func (repo *categoryRepo) Create(category models.Category) (*models.Category, error) {
	query := `
		INSERT INTO categories (name, parent_id)
		VALUES ($1, $2)
		RETURNING id;
	`

	err := repo.db.QueryRow(query, category.Name, category.ParentID).Scan(&category.ID)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (repo *categoryRepo) List() ([]*models.Category, error) {
	var categories []*models.Category
	query := `SELECT id, name, parent_id FROM categories ORDER BY id ASC;`

	err := repo.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (repo *categoryRepo) GetByID(id int) (*models.Category, error) {
	var category models.Category

	query := `
		SELECT id, name, parent_id
		FROM categories
		WHERE id = $1
		LIMIT 1;
	`

	err := repo.db.Get(&category, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	return &category, nil
}

func (repo *categoryRepo) Update(category models.Category) (*models.Category, error) {
	query := `
		UPDATE categories
		SET name=$1, parent_id=$2, updated_at=CURRENT_TIMESTAMP
		WHERE id=$3
		RETURNING id, name, parent_id;
	`

	err := repo.db.QueryRow(query, category.Name, category.ParentID, category.ID).Scan(&category.ID, &category.Name, &category.ParentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	return &category, nil
}

func (repo *categoryRepo) Delete(id int64) error {
	query := `DELETE FROM categories WHERE id=$1;`
	
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrCategoryNotFound
	}

	return nil
}
