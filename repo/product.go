package repo

import (
	"database/sql"
	"errors"
	"test/models"

	"github.com/jmoiron/sqlx"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepo interface {
	Create(product models.Product) (*models.Product, error)
	Get(productId int) (*models.Product, error)
	List() ([]*models.Product, error)
	Update(product models.Product) (*models.Product, error)
	Delete(productId int) error
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}
}

// -------------------- CREATE --------------------
func (repo *productRepo) Create(product models.Product) (*models.Product, error) {
	query := `
		INSERT INTO products (title, description, price, currency, stock, category_id, image_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	err := repo.db.QueryRow(query, product.Title, product.Description, product.Price, product.Currency, product.Stock, product.CategoryID, product.ImageURL).Scan(&product.ID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// -------------------- GET --------------------
func (repo *productRepo) Get(productId int) (*models.Product, error) {
	var product models.Product
	query := `SELECT id, title, description, price, currency, stock, category_id, image_url FROM products WHERE id=$1 LIMIT 1;`

	err := repo.db.Get(&product, query, productId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

// -------------------- LIST --------------------
func (repo *productRepo) List() ([]*models.Product, error) {
	var products []*models.Product
	query := `SELECT id, title, description, price, currency, stock, category_id, image_url FROM products ORDER BY id ASC;`

	err := repo.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// -------------------- UPDATE --------------------
func (repo *productRepo) Update(product models.Product) (*models.Product, error) {
	query := `
		UPDATE products
		SET title=$1, description=$2, price=$3, currency=$4, stock=$5, category_id=$6, image_url=$7, updated_at=CURRENT_TIMESTAMP
		WHERE id=$8
		RETURNING id, title, description, price, currency, stock, category_id, image_url;
	`

	row := repo.db.QueryRow(query, product.Title, product.Description, product.Price, product.Currency, product.Stock, product.CategoryID, product.ImageURL, product.ID)
	err := row.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Currency, &product.Stock, &product.CategoryID, &product.ImageURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

// -------------------- DELETE --------------------
func (repo *productRepo) Delete(productId int) error {
	query := `DELETE FROM products WHERE id=$1;`

	result, err := repo.db.Exec(query, productId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}
