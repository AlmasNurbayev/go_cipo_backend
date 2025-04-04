ALTER TABLE product 
  ADD COLUMN IF NOT EXISTS kaspi_in_sale BOOLEAN NOT NULL DEFAULT FALSE,
  ADD COLUMN IF NOT EXISTS kaspi_category TEXT;

COMMENT ON COLUMN product.kaspi_in_sale IS 'Признак допустимости выгрузки товара в Kaspi';
COMMENT ON COLUMN product.kaspi_category IS 'Категория товара для Kaspi указанная в 1С';  

DROP INDEX IF EXISTS product_idx;

CREATE INDEX IF NOT EXISTS product_idx 
ON product (id, registrator_id, id_1c, kaspi_in_sale);