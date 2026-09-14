REPLACE TABLE base.t1 (n UInt64)
ENGINE = MergeTree
ORDER BY n;

INSERT INTO base.t1 VALUES (3);

SELECT * FROM base.t1;

┌─n─┐
│ 3 │
└───┘
