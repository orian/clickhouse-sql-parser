clickhouse-cloud :) INSERT INTO query_cache_test SELECT * FROM generateRandom('name String, age UInt8',1,1) LIMIT 100000;

INSERT INTO query_cache_test SELECT *
FROM generateRandom('name String, age UInt8', 1, 1)
LIMIT 100000

Query id: 90369105-bd67-494c-bdaf-d90dbfb6def9

Ok.

0 rows in set. Elapsed: 0.173 sec. Processed 327.05 thousand rows, 3.43 MB (1.89 million rows/s., 19.86 MB/s.)
