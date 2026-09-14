-- Создать таблицу и вставить данные
CREATE TABLE test_lazy (json JSON) ENGINE = MergeTree ORDER BY tuple();
INSERT INTO test_lazy VALUES ('{"user_id": "123", "score": "95.5"}');

-- Включить экспериментальную настройку
SET allow_experimental_json_lazy_type_hints = 1;

-- Добавить подсказки типов — выполняется мгновенно без мутации
ALTER TABLE test_lazy MODIFY COLUMN json JSON(user_id UInt64, score Float64);

-- Запросить данные — подсказки типов применяются во время чтения
SELECT json.user_id, toTypeName(json.user_id), json.score, toTypeName(json.score) FROM test_lazy;
