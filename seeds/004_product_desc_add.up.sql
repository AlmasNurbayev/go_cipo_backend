SELECT setval(pg_get_serial_sequence('product_desc_mapping', 'id'), 
              (SELECT MAX(id) FROM product_desc_mapping));

INSERT INTO product_desc_mapping
(id, id_1c, name_1c, field)
VALUES 
(9, '3f3f6419-0c75-11f0-923a-50ebf624c538', 'КатегорияКаспийМагазин', 'kaspi_category'),
(10, '3f3f641a-0c75-11f0-923a-50ebf624c538', 'ВыгружатьВКаспий', 'kaspi_in_sale')
ON CONFLICT (id) DO NOTHING;

SELECT setval(pg_get_serial_sequence('product_desc_mapping', 'id'), 
              (SELECT MAX(id) FROM product_desc_mapping));