CREATE TABLE test
(
    key Int64,
    some String
)
ENGINE = ReplacingMergeTree
ORDER BY key;

INSERT INTO test FORMAT Values (1, 'first');
INSERT INTO test FORMAT Values (1, 'second');

SELECT * FROM test;
┌─key─┬─some───┐
│   1 │ second │
└─────┴────────┘
┌─key─┬─some──┐
│   1 │ first │
└─────┴───────┘

SELECT * FROM test SETTINGS final = 1;
┌─key─┬─some───┐
│   1 │ second │
└─────┴────────┘

SET final = 1;
SELECT * FROM test;
┌─key─┬─some───┐
│   1 │ second │
└─────┴────────┘
