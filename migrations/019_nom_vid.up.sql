ALTER TABLE product ADD COLUMN IF NOT EXISTS nom_vid TEXT;

ALTER TABLE qnt_price_registry ADD COLUMN IF NOT EXISTS nom_vid TEXT;

CREATE INDEX IF NOT EXISTS qnt_price_registry_nom_vid_idx2 
ON qnt_price_registry (id, registrator_id, product_name, nom_vid);
