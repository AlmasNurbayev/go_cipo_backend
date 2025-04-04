ALTER TABLE kaspi_categories ADD COLUMN IF NOT EXISTS material_kaspi TEXT[];
ALTER TABLE kaspi_categories ADD COLUMN IF NOT EXISTS season_kaspi TEXT[];
ALTER TABLE kaspi_categories ADD COLUMN IF NOT EXISTS colour_kaspi TEXT[];
ALTER TABLE kaspi_categories ADD COLUMN IF NOT EXISTS attributes_list JSONB;