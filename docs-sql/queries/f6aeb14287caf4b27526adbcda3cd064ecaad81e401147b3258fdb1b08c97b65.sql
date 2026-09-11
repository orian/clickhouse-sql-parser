CREATE TABLE test (d Dynamic) ENGINE = Memory;
INSERT INTO test VALUES ('world'), (123), (456);

-- デフォルト（不一致時に例外をスロー）: length() は整数を受け付けないため、クエリは例外をスローします。
SELECT length(d) FROM test;  -- ILLEGAL_TYPE_OF_ARGUMENT をスロー

-- スローを無効にした場合: 互換性のない行は NULL を返します。
SET dynamic_throw_on_type_mismatch = false;
SELECT d, length(d) FROM test ORDER BY d::String NULLS LAST;
