CREATE OR REPLACE TABLE foo
(
    n Nested(a String, b Int32)
) ENGINE = MergeTree ORDER BY ();
-- flatten_nested=1이 기본값입니다
INSERT INTO foo VALUES (['foo', 'bar'], [42, 144]);
