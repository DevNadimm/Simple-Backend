-- +migrate Up
CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    parent_id BIGINT NULL REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add foreign key to products if it doesn't exist
ALTER TABLE products 
ADD CONSTRAINT fk_product_category 
FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL;
