CREATE OR REPLACE TABLE test
(
    id UInt64,
    updated_at DateTime DEFAULT now(),
    updated_at_date Date DEFAULT toDate(updated_at)
)
ENGINE = MergeTree
ORDER BY id;

INSERT INTO test (id) VALUES (1);

SELECT * FROM test;
┌─id─┬──────────updated_at─┬─updated_at_date─┐
│  1 │ 2023-02-24 17:06:46 │      2023-02-24 │
└────┴─────────────────────┴─────────────────┘
