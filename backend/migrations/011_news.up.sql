CREATE TABLE IF NOT EXISTS news (
  id BIGINT GENERATED BY DEFAULT AS IDENTITY  PRIMARY KEY,

  title TEXT NOT NULL,
  data TEXT NOT NULL,
  image_path TEXT,
  Operation_date TIMESTAMPTZ NOT NULL,

 changed_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, 
 create_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);