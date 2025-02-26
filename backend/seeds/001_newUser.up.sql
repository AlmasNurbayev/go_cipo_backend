INSERT INTO "user"
(id, email, name, password)
VALUES 
(1, 'ntldr@mail.ru', 'almas', '1')
ON CONFLICT (id) DO NOTHING;

SELECT setval(pg_get_serial_sequence('user', 'id'), 
              (SELECT MAX(id) FROM "user"));
