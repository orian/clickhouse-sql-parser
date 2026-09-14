CREATE TABLE test (d Dynamic) ENGINE = Memory;
INSERT INTO test VALUES ('world'), (123), (456);

-- 기본값(불일치 시 예외 발생): length()는 정수를 허용하지 않으므로 쿼리에서 예외가 발생합니다.
SELECT length(d) FROM test;  -- ILLEGAL_TYPE_OF_ARGUMENT 예외 발생

-- 예외 발생 비활성화 시: 호환되지 않는 행은 NULL을 반환합니다.
SET dynamic_throw_on_type_mismatch = false;
SELECT d, length(d) FROM test ORDER BY d::String NULLS LAST;
