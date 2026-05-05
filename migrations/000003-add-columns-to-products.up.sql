-- +migrate Up
ALTER TABLE products 
ADD COLUMN currency VARCHAR(10) DEFAULT 'BDT',
ADD COLUMN stock INT DEFAULT 0,
ADD COLUMN category_id BIGINT,
ADD COLUMN image_url TEXT;
