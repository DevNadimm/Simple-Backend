-- +migrate Down
ALTER TABLE products 
DROP COLUMN IF EXISTS currency,
DROP COLUMN IF EXISTS stock,
DROP COLUMN IF EXISTS category_id,
DROP COLUMN IF EXISTS image_url;
