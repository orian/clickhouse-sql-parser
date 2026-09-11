CREATE TABLE test (d Dynamic) ENGINE = Memory;
INSERT INTO test VALUES ('world'), (123), (456);

-- 默认行为（类型不匹配时抛出异常）：length() 不接受整数类型，因此查询将抛出异常。
SELECT length(d) FROM test;  -- 抛出 ILLEGAL_TYPE_OF_ARGUMENT

-- 禁用异常抛出后：不兼容的行返回 NULL。
SET dynamic_throw_on_type_mismatch = false;
SELECT d, length(d) FROM test ORDER BY d::String NULLS LAST;
