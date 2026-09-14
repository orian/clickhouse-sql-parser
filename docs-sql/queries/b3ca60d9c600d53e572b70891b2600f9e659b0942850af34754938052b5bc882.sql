CREATE TABLE test
(
    x Int32,
    e Int32 EPHEMERAL 0,
    m Int32 MATERIALIZED x + e
)
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO test (x, e) VALUES (1, 7);

ALTER TABLE test UPDATE x = 2 WHERE 1;

SELECT x, m FROM test;
