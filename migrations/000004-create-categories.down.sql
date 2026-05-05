-- +migrate Down
ALTER TABLE products DROP CONSTRAINT IF EXISTS fk_product_category;
DROP TABLE IF EXISTS categories;
