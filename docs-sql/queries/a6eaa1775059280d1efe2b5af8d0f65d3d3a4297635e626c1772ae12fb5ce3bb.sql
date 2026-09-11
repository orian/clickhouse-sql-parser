-- Создадим таблицу, где user_id — уникальный идентификатор пользователя, encrypted — зашифрованное строковое поле, iv — начальный вектор для расшифровки/шифрования.
-- Предполагается, что пользователи знают свой id и ключ для расшифровки поля encrypted:
CREATE TABLE decrypt_null
(
    dt DateTime,
    user_id UInt32,
    encrypted String,
    iv String
)
ENGINE = MergeTree;

-- Вставим некоторые данные:
INSERT INTO decrypt_null VALUES
('2022-08-02 00:00:00', 1, encrypt('aes-256-gcm', 'value1', 'keykeykeykeykeykeykeykeykeykey01', 'iv1'), 'iv1'),
('2022-09-02 00:00:00', 2, encrypt('aes-256-gcm', 'value2', 'keykeykeykeykeykeykeykeykeykey02', 'iv2'), 'iv2'),
('2022-09-02 00:00:01', 3, encrypt('aes-256-gcm', 'value3', 'keykeykeykeykeykeykeykeykeykey03', 'iv3'), 'iv3');

-- Попробуем расшифровать с помощью одного ключа
SELECT
    dt,
    user_id,
    tryDecrypt('aes-256-gcm', encrypted, 'keykeykeykeykeykeykeykeykeykey02', iv) AS value
FROM decrypt_null
ORDER BY user_id ASC
