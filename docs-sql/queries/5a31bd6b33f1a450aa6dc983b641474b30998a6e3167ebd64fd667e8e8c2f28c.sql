CREATE TABLE test (d Dynamic) ENGINE = Memory;
INSERT INTO test VALUES ('world'), (123), (456);

-- По умолчанию (исключение при несоответствии типов): length() не принимает целые числа, поэтому запрос завершается ошибкой.
SELECT length(d) FROM test;  -- выбрасывает ILLEGAL_TYPE_OF_ARGUMENT

-- При отключённом выбрасывании исключений: несовместимые строки возвращают NULL.
SET dynamic_throw_on_type_mismatch = false;
SELECT d, length(d) FROM test ORDER BY d::String NULLS LAST;
