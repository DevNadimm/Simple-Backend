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

	return repo.Get(product.ID)
}

// -------------------- GET --------------------
func (repo *productRepo) Get(productId int) (*models.Product, error) {
	var product models.Product
	var category models.Category
	var categoryID sql.NullInt64
	var categoryName sql.NullString
	var categoryParentID sql.NullInt64

	query := `
		SELECT 
			p.id, p.title, p.description, p.price, p.currency, p.stock, p.category_id, p.image_url,
			c.id as c_id, c.name as c_name, c.parent_id as c_parent_id
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1
		LIMIT 1;
	`

	row := repo.db.QueryRow(query, productId)
	err := row.Scan(
		&product.ID, &product.Title, &product.Description, &product.Price, &product.Currency, &product.Stock, &product.CategoryID, &product.ImageURL,
		&categoryID, &categoryName, &categoryParentID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	if categoryID.Valid {
		category.ID = categoryID.Int64
		category.Name = categoryName.String
		if categoryParentID.Valid {
			pid := categoryParentID.Int64
			category.ParentID = &pid
		}
		product.Category = &category
	}

	return &product, nil
}

// -------------------- LIST --------------------
func (repo *productRepo) List() ([]*models.Product, error) {
	query := `
		SELECT 
			p.id, p.title, p.description, p.price, p.currency, p.stock, p.category_id, p.image_url,
			c.id as c_id, c.name as c_name, c.parent_id as c_parent_id
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		ORDER BY p.id ASC;
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var product models.Product
		var category models.Category
		var categoryID sql.NullInt64
		var categoryName sql.NullString
		var categoryParentID sql.NullInt64

		err := rows.Scan(
			&product.ID, &product.Title, &product.Description, &product.Price, &product.Currency, &product.Stock, &product.CategoryID, &product.ImageURL,
			&categoryID, &categoryName, &categoryParentID,
		)
		if err != nil {
			return nil, err
		}

		if categoryID.Valid {
			category.ID = categoryID.Int64
			category.Name = categoryName.String
			if categoryParentID.Valid {
				pid := categoryParentID.Int64
				category.ParentID = &pid
			}
			product.Category = &category
		}
		products = append(products, &product)
	}

	return products, nil
}

// -------------------- UPDATE --------------------
func (repo *productRepo) Update(product models.Product) (*models.Product, error) {
	query := `
		UPDATE products
		SET title=$1, description=$2, price=$3, currency=$4, stock=$5, category_id=$6, image_url=$7, updated_at=CURRENT_TIMESTAMP
		WHERE id=$8
		RETURNING id;
	`

	var id int
	err := repo.db.QueryRow(query, product.Title, product.Description, product.Price, product.Currency, product.Stock, product.CategoryID, product.ImageURL, product.ID).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	return repo.Get(id)
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
