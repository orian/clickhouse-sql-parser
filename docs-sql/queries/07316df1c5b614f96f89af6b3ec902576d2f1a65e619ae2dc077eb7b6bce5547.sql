CREATE TABLE mydata
(
    `A` Int64,
    `B` Int8,
    `C` String
)
ENGINE = MergeTree
ORDER BY A AS
SELECT
    number,
    0,
    if(number between 1000 and 2000, 'x', toString(number))
FROM numbers(10000000);

SELECT count()
FROM mydata
WHERE (B = 0) AND (C = 'x');

1 row in set. Elapsed: 0.074 sec. Processed 10.00 million rows, 168.89 MB (134.98 million rows/s., 2.28 GB/s.)

-- Enable tracing to see which predicates are moved to PREWHERE.
set send_logs_level='debug';

MergeTreeWhereOptimizer: condition "B = 0" moved to PREWHERE
-- ClickHouse automatically moves B = 0 to PREWHERE, but this condition does not filter any rows because B is always 0.

-- Move the more selective C = 'x' predicate to PREWHERE.

SELECT count()
FROM mydata
PREWHERE C = 'x'
WHERE B = 0;

1 row in set. Elapsed: 0.069 sec. Processed 10.00 million rows, 158.89 MB (144.90 million rows/s., 2.30 GB/s.)

-- The query with manually specified PREWHERE processes slightly less data: 158.89 MB instead of 168.89 MB.
