clickhouse-cloud :) INSERT INTO raw_data SELECT * FROM generateRandom('`id` UInt32,
`data` String',1,1) LIMIT 1000000;

INSERT INTO raw_data SELECT *
FROM generateRandom('`id` UInt32,
`data` String', 1, 1)
LIMIT 1000000

Query id: c552a34a-b72f-45e1-bed0-778923e1b5c9

Ok.

0 rows in set. Elapsed: 0.438 sec. Processed 1.05 million rows, 10.99 MB (2.39 million rows/s., 25.11 MB/s.)
