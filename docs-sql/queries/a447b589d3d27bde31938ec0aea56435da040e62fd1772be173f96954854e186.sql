CREATE OR REPLACE TABLE base.t1
(
    n UInt64,
    s Nullable(String)
)
ENGINE = MergeTree
ORDER BY n;

INSERT INTO base.t1 VALUES (2, null);

SELECT * FROM base.t1;

2
