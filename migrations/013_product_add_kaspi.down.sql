DROP INDEX IF EXISTS product_idx;

CREATE INDEX IF NOT EXISTS product_idx 
ON product (id, registrator_id, id_1c);

ALTER TABLE product 
  DROP COLUMN kaspi_in_sale,
  DROP COLUMN kaspi_category;