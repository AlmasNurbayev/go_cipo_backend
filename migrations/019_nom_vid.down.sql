ALTER TABLE product DROP COLUMN IF EXISTS nom_vid;
ALTER TABLE qnt_price_registry DROP COLUMN IF EXISTS nom_vid;

DROP INDEX IF EXISTS CREATE INDEX IF NOT EXISTS qnt_price_registry_nom_vid_idx2;
