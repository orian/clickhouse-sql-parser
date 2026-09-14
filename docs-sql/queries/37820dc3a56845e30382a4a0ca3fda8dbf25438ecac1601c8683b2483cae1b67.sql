CREATE OR REPLACE TABLE foo
(
    n Nested(a String, b Int32)
) ENGINE = MergeTree ORDER BY ();
-- flatten_nested=1 é o padrão
INSERT INTO foo VALUES (['foo', 'bar'], [42, 144]);
