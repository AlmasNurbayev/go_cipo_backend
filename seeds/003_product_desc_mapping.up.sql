SELECT setval(pg_get_serial_sequence('product_desc_mapping', 'id'), 
              (SELECT MAX(id) FROM product_desc_mapping));

INSERT INTO product_desc_mapping
(id, id_1c, name_1c, field)
VALUES 
(1, '7dac490d-b0d2-11ed-b0f1-50ebf624c538', 'Материал подошвы', 'material_podoshva'),
(2, '7dac490f-b0d2-11ed-b0f1-50ebf624c538', 'Материал внутри', 'material_inside'),
(3, '7dac490e-b0d2-11ed-b0f1-50ebf624c538', 'Материал вверх', 'material_up'),
(4, '28a101cd-b439-11ed-b0f5-50ebf624c538', 'Мальчик/девочка', 'sex'),
(5, '28a101ce-b439-11ed-b0f5-50ebf624c538', 'Основной цвет', 'main_color'),
(6, 'a001d8e0-a3b3-11ed-b0d2-50ebf624c538', 'ВыгружатьВеб', 'public_web'),
(7, '6c2db24d-a792-11ed-b0d5-50ebf624c538', 'ТоварнаяГруппа', 'product_group'),
(8, '7dac4910-b0d2-11ed-b0f1-50ebf624c538', 'ВидМодели', 'vidModeli')
ON CONFLICT (id) DO NOTHING;

SELECT setval(pg_get_serial_sequence('product_desc_mapping', 'id'), 
              (SELECT MAX(id) FROM product_desc_mapping));