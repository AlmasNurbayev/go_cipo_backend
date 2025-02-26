INSERT INTO registrator
	(id, operation_date, name_folder, name_file, user_id, date_schema,
		id_catalog, id_class, name_catalog, name_class, ver_schema, only_change) 
		VALUES 
		(1, '2024-02-21 15:30:00+05', 'seed_name_folder', 'seed_name_file', 1, '2024-02-21 15:30:00+05', 'id_catalog', 
    'id_class', 'name_catalog', 'name_class', 'ver_schema', false) 
		ON CONFLICT (id) DO NOTHING;

		SELECT setval(pg_get_serial_sequence('registrator', 'id'), 
              (SELECT MAX(id) FROM registrator));