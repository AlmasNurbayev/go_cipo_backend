CREATE TABLE IF NOT EXISTS kaspi_organizations (
  id BIGINT GENERATED BY DEFAULT AS IDENTITY  PRIMARY KEY,

  name TEXT NOT NULL,
  kaspi_id TEXT NOT NULL,
  kaspi_api_token_hash TEXT NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT FALSE,

  changed_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, 
  create_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE kaspi_organizations IS 'Организации заведенные в Kaspi-магазин';