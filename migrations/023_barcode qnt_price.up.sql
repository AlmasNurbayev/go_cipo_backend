ALTER TABLE "qnt_price_registry"
  ADD COLUMN IF NOT EXISTS barcode TEXT;

DROP INDEX IF EXISTS qnt_price_idx;

CREATE INDEX IF NOT EXISTS qnt_price_idx
ON qnt_price_registry (registrator_id, product_id, product_create_date, barcode)
INCLUDE (qnt);

